package repository

import "github.com/d0lim/turnstile/internal/core/domain"

type UserRepository interface {
	CreateUser(account *domain.User) error
	GetUserByID(id int64) (*domain.User, error)
	GetUserByOAuthProviderAndEmail(oAuthProvider string, email string) (*domain.User, error)
	DeleteUser(id int64) error
}
