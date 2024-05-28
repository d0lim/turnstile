//go:build wireinject
// +build wireinject

package framework

import (
	"github.com/d0lim/turnstile/internal/adapters/in/api"
	"github.com/d0lim/turnstile/internal/adapters/out/db"
	"github.com/d0lim/turnstile/internal/adapters/out/db/ent"
	"github.com/d0lim/turnstile/internal/core/ports/in/usecase"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializeApp() (*fiber.App, error) {
	wire.Build(
		ent.NewClient,
		config.NewSessionStore,
		db.NewUserRepository,
		usecase.NewUserUsecase,
		api.NewUserHandler,
		NewApp,
	)

	return &fiber.App{}, nil
}
