package framework

import (
	_ "github.com/d0lim/turnstile/ent/runtime"
	"github.com/d0lim/turnstile/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func NewApp(
	userController controller.UserController,
) (*fiber.App, error) {
	app := fiber.New()
	SetupRoutes(app, userController)

	return app, nil
}

func SetupRoutes(app *fiber.App, handler controller.UserController) {
	app.Get("/api/v1/ping", func(ctx *fiber.Ctx) error {
		return ctx.Send([]byte("PONG"))
	})
	app.Get("/api/v1/auth/login", handler.GetGoogleLoginRedirect)
	app.Get("/api/v1/auth/callback", handler.CallbackGoogle)
	app.Get("/api/v1/auth/authenticate", handler.Authenticate)
	app.Get("/api/v1/auth/refresh", handler.Refresh)
}
