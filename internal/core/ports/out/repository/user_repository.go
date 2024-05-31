package repository

import (
	"context"
	"github.com/d0lim/turnstile/internal/core/domain"
)

type UserRepository interface {
	CreateUser(account *domain.User, ctx context.Context) (*domain.User, *domain.DomainError)
	GetUserByID(id int64, ctx context.Context) (*domain.User, *domain.DomainError)
	GetUserByOAuthProviderAndEmail(oAuthProvider string, email string, ctx context.Context) (*domain.User, *domain.DomainError)
	DeleteUser(id int64, ctx context.Context) *domain.DomainError
}
