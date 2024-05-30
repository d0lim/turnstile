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
	"net/http"
)

type UserHandler struct {
	oAuthConfig *config.OAuthConfig
	session     *config.SessionConfig
	usecase     *usecase.UserUsecase
}

func NewUserHandler(
	oAuthConfig *config.OAuthConfig,
	session *config.SessionConfig,
	usecase *usecase.UserUsecase,
) *UserHandler {
	return &UserHandler{
		oAuthConfig: oAuthConfig,
		session:     session,
		usecase:     usecase,
	}
}

func (h *UserHandler) RedirectLoginGoogle(c *fiber.Ctx) error {
	session, err := h.session.Store.Get(c)
	if err != nil {
		return err
	}
	state := uuid.NewString()
	session.Set("state", state)
	err = session.Save()
	if err != nil {
		return err
	}

	authCodeURL := h.oAuthConfig.Google.AuthCodeURL(state, oauth2.AccessTypeOffline)

	return c.Redirect(authCodeURL, http.StatusTemporaryRedirect)
}

func (h *UserHandler) CallbackGoogle(c *fiber.Ctx) error {
	session, err := h.session.Store.Get(c)
	if err != nil {
		return err
	}
	state := c.Query("state")
	code := c.Query("code")
	savedState := session.Get("state")
	if savedState != state {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid state")
	}

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

	fmt.Println(user)

	return c.JSON(user)
}
