package repository

import "github.com/d0lim/turnstile/internal/user/domain"

type UserRepository interface {
	CreateUser(account *domain.User) error
	GetUserByID(id int64) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	DeleteUser(id int64) error
}
