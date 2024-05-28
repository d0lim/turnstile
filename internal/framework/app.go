package framework

import (
	"github.com/d0lim/turnstile/internal/adapters/in/api"
	"github.com/gofiber/fiber/v2"
)

func NewApp(
	userHandler *api.UserHandler,
) (*fiber.App, error) {
	app := fiber.New()
	api.SetupRoutes(app, userHandler)

	return app, nil
}
