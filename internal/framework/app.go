package framework

import (
	"github.com/d0lim/turnstile/internal/adapters/in/api"
	"github.com/gofiber/fiber/v2"
)

func NewApp(
	userHandler *api.UserHandler,
) (*fiber.App, error) {
	app := fiber.New()
	SetupRoutes(app, userHandler)

	return app, nil
}

func SetupRoutes(app *fiber.App, handler *api.UserHandler) {
	app.Post("/api/v1/auth/login/google", handler.RedirectLoginGoogle)
	app.Get("/api/v1/auth/callback/google", handler.CallbackGoogle)
}
