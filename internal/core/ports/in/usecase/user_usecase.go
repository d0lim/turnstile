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

func (u *UserUsecase) GetUserByOAuthProviderAndEmailOrCreateIfAbsent(
	oAuthProvider string,
	email string,
	user *domain.User,
	ctx context.Context,
) (*domain.User, error) {
	userFromDb, err := u.GetUserByOAuthProviderAndEmail(oAuthProvider, email, ctx)
	if err != nil {
		if domainErr, ok := domain.IsDomainError(err); ok {
			if domainErr.Code == domain.NotFound {
				createdUser, err := u.CreateUser(user, ctx)
				if err != nil {
					return nil, err
				}
				return createdUser, nil
			}
		}
		return nil, err
	}
	return userFromDb, nil
}

func (u *UserUsecase) CreateUser(user *domain.User, ctx context.Context) (*domain.User, error) {
	return u.repo.CreateUser(user, ctx)
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
