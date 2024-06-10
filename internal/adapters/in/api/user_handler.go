package api

import (
	"encoding/json"
	"fmt"
	"github.com/d0lim/turnstile/internal/adapters/in/api/dto"
	"github.com/d0lim/turnstile/internal/core/ports/in/usecase"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"strings"
	"time"
)

type UserHandler struct {
	oAuthConfig *config.OAuthConfig
	usecase     *usecase.UserUsecase
}

func NewUserHandler(
	oAuthConfig *config.OAuthConfig,
	usecase *usecase.UserUsecase,
) *UserHandler {
	return &UserHandler{
		oAuthConfig: oAuthConfig,
		usecase:     usecase,
	}
}

func (h *UserHandler) GetRedirectLoginGoogle(c *fiber.Ctx) error {
	state := uuid.NewString()

	authCodeURL := h.oAuthConfig.Google.AuthCodeURL(state, oauth2.AccessTypeOffline)

	return c.JSON(&dto.RedirectUriResponse{RedirectUri: authCodeURL})
}

func (h *UserHandler) CallbackGoogle(c *fiber.Ctx) error {
	code := c.Query("code")

	token, err := h.oAuthConfig.Google.Exchange(c.Context(), code)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	client := h.oAuthConfig.Google.Client(c.Context(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer userInfo.Body.Close()

	var user dto.GoogleUserResponse
	if err := json.NewDecoder(userInfo.Body).Decode(&user); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	tokenPair, derr := h.usecase.Login(
		user.ID,
		"GOOGLE",
		user.Name,
		user.Email,
		&user.Picture,
		c.Context(),
	)

	if derr != nil {
		fmt.Println(derr.Cause.Error())
		return fiber.NewError(fiber.StatusInternalServerError, derr.Error())
	}

	cookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Expires:  time.Now().Add(168 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	}

	c.Cookie(cookie)

	return c.JSON(&dto.LoginResponse{AccessToken: tokenPair.AccessToken})
}

func (h *UserHandler) Authenticate(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or malformed JWT",
		})
	}

	// The token normally comes as "Bearer <token>"
	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Malformed JWT",
		})
	}

	user, dErr := h.usecase.Authenticate(tokenString, c.Context())
	if dErr != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": dErr.Error(),
		})
	}

	internalUserResponse := &dto.InternalUserResponse{
		ID:              user.ID,
		OAuthProvider:   user.OAuthProvider,
		Email:           user.Email,
		Name:            user.Name,
		ProfileImageUrl: user.ProfileImageUrl,
	}

	marshaled, err := json.Marshal(internalUserResponse)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{})
	}

	c.Set("X-Auth-User", string(marshaled))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"authenticated": true})
}
