package repository

import "github.com/d0lim/turnstile/internal/user/domain"

type UserRepository interface {
	CreateUser(account *domain.User) error
	GetUserByID(id int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	DeleteUser(id int) error
}
