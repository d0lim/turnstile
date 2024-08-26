package controller

import (
	"encoding/json"
	"github.com/d0lim/turnstile/internal/controller/dto"
	"github.com/d0lim/turnstile/internal/service"
	"github.com/gofiber/fiber/v2"
	"strings"
	"time"
)

type UserController interface {
	GetGoogleLoginRedirect(c *fiber.Ctx) error
	CallbackGoogle(c *fiber.Ctx) error
	Authenticate(c *fiber.Ctx) error
	Refresh(c *fiber.Ctx) error
}

type userController struct {
	oauthService service.OauthService
	userService  service.UserService
}

func NewUserController(oauthService service.OauthService, userService service.UserService) UserController {
	return &userController{
		oauthService: oauthService,
		userService:  userService,
	}
}

func (u userController) GetGoogleLoginRedirect(c *fiber.Ctx) error {
	callbackRedirectUri := c.Query("redirect_uri")
	loginRedirectUri := u.oauthService.GoogleLoginRedirectUri(callbackRedirectUri)

	return c.JSON(&dto.RedirectUriResponse{RedirectUri: loginRedirectUri})
}

func (u userController) CallbackGoogle(c *fiber.Ctx) error {
	code := c.Query("code")
	userResponse, err := u.oauthService.GoogleHandleCallback(code, c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	tokenPair, err := u.userService.Login(userResponse.ID, "GOOGLE", userResponse.Name, userResponse.Email, &userResponse.Picture, c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	cookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Expires:  time.Now().Add(time.Hour * 168),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	}

	c.Cookie(cookie)

	return c.JSON(&dto.LoginResponse{AccessToken: tokenPair.AccessToken})
}

func (u userController) Authenticate(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Authorization header is empty")
	}

	// Token normally comes as "Bearer <token>"
	tokenString := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer "))
	if tokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Malformed token")
	}

	user, err := u.userService.Authenticate(tokenString, c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
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
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	c.Set("X-Auth-User", string(marshaled))

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"authenticated": true})
}

func (u userController) Refresh(c *fiber.Ctx) error {
	refreshToken := c.Cookies("refresh_token")
	if refreshToken == "" {
		return fiber.NewError(fiber.StatusInternalServerError, "Refresh token is empty")
	}
	tokenPair, err := u.userService.Refresh(refreshToken)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	cookie := &fiber.Cookie{
		Name:     "refresh_token",
		Value:    tokenPair.RefreshToken,
		Expires:  time.Now().Add(time.Hour * 168),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "None",
		Path:     "/",
	}

	c.Cookie(cookie)

	return c.JSON(&dto.LoginResponse{AccessToken: tokenPair.AccessToken})
}
