// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/d0lim/turnstile/internal/adapters/in/api"
	"github.com/d0lim/turnstile/internal/adapters/out/db"
	"github.com/d0lim/turnstile/internal/adapters/out/db/ent"
	"github.com/d0lim/turnstile/internal/adapters/out/jwt"
	"github.com/d0lim/turnstile/internal/core/ports/in/usecase"
	"github.com/d0lim/turnstile/internal/framework"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/gofiber/fiber/v2"
)

// Injectors from wire.go:

func InitializeApp() (*fiber.App, error) {
	oAuthConfig := config.NewOAuthConfig()
	sessionConfig := config.NewSessionConfig()
	client, err := ent.NewClient()
	if err != nil {
		return nil, err
	}
	userRepository := db.NewUserRepository(client)
	jwtConfig := config.NewJwtConfig()
	tokenManager := jwt.NewJwtTokenManager(jwtConfig)
	userUsecase := usecase.NewUserUsecase(userRepository, tokenManager)
	userHandler := api.NewUserHandler(oAuthConfig, sessionConfig, userUsecase)
	app, err := framework.NewApp(userHandler)
	if err != nil {
		return nil, err
	}
	return app, nil
}
