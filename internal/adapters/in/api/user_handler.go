package api

import (
	"github.com/d0lim/turnstile/internal/core/ports/in/usecase"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	session *config.SessionConfig
	usecase *usecase.UserUsecase
}

func NewUserHandler(
	session *config.SessionConfig,
	usecase *usecase.UserUsecase,
) *UserHandler {
	return &UserHandler{session: session, usecase: usecase}
}

func (h *UserHandler) RedirectLoginGoogle(c *fiber.Ctx) error {
	return nil
}

func (h *UserHandler) CallbackGoogle(c *fiber.Ctx) error {
	return nil
}

func SetupRoutes(app *fiber.App, handler *UserHandler) {
	app.Post("/api/v1/auth/login/google", handler.RedirectLoginGoogle)
	app.Get("/api/v1/auth/callback/google", handler.CallbackGoogle)
}
