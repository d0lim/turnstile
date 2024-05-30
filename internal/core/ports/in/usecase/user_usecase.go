package usecase

import (
	"context"
	"github.com/d0lim/turnstile/internal/core/domain"
	"github.com/d0lim/turnstile/internal/core/ports/out/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) CreateUser(account *domain.User, ctx context.Context) error {
	return u.repo.CreateUser(account, ctx)
}

func (u *UserUsecase) GetUserByID(id int64, ctx context.Context) (*domain.User, error) {
	return u.repo.GetUserByID(id, ctx)
}

func (u *UserUsecase) GetUserByOAuthProviderAndEmail(oAuthProvider string, email string, ctx context.Context) (*domain.User, error) {
	return u.repo.GetUserByOAuthProviderAndEmail(oAuthProvider, email, ctx)
}

func (u *UserUsecase) DeleteUser(id int64, ctx context.Context) error {
	return u.repo.DeleteUser(id, ctx)
}
