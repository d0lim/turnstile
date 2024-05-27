package usecase

import (
	"github.com/d0lim/turnstile/internal/user/domain"
	"github.com/d0lim/turnstile/internal/user/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(account *domain.User) error {
	return u.repo.CreateUser(account)
}

func (u *UserUsecase) GetUserByID(id int) (*domain.User, error) {
	return u.repo.GetUserByID(id)
}

func (u *UserUsecase) GetUserByEmail(email string) (*domain.User, error) {
	return u.repo.GetUserByEmail(email)
}

func (u *UserUsecase) DeleteUser(id int) error {
	return u.repo.DeleteUser(id)
}
