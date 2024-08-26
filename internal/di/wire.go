//go:build wireinject
// +build wireinject

package di

import (
	"github.com/d0lim/turnstile/internal/adapters/out/db/ent"
	"github.com/d0lim/turnstile/internal/controller"
	"github.com/d0lim/turnstile/internal/framework"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/d0lim/turnstile/internal/repository"
	"github.com/d0lim/turnstile/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
)

func InitializeApp() (*fiber.App, error) {
	wire.Build(
		ent.NewClient,
		config.NewOAuthConfig,
		config.NewJwtConfig,
		config.NewRedisConfig,
		repository.NewUserRepository,
		service.NewTokenService,
		service.NewUserService,
		service.NewOauthService,
		controller.NewUserController,
		framework.NewApp,
	)

	return &fiber.App{}, nil
}
