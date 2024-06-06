package usecase

import (
	"context"
	"github.com/d0lim/turnstile/internal/core/domain"
	"github.com/d0lim/turnstile/internal/core/ports/out/repository"
	"github.com/d0lim/turnstile/internal/core/ports/out/token"
	"strconv"
)

type UserUsecase struct {
	repo    repository.UserRepository
	manager token.TokenManager
}

func NewUserUsecase(repo repository.UserRepository, manager token.TokenManager) *UserUsecase {
	return &UserUsecase{repo: repo, manager: manager}
}

func (u *UserUsecase) Login(
	oAuthId string,
	oAuthProvider string,
	name string,
	email string,
	profileImageUrl *string,
	ctx context.Context,
) (*domain.TokenPair, *domain.DomainError) {
	user, err := u.GetUserByOAuthProviderAndEmailOrCreateIfAbsent(
		oAuthId,
		oAuthProvider,
		name,
		email,
		profileImageUrl,
		ctx,
	)
	if err != nil {
		return nil, err
	}

	accessToken, err := u.manager.IssueAccessToken(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, err
	}
	refreshToken, err := u.manager.IssueRefreshToken(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UserUsecase) Authenticate(tokenString string, ctx context.Context) (*domain.User, *domain.DomainError) {
	verifiedToken, err := u.manager.VerifyAccessToken(tokenString)
	if err != nil {
		return nil, err
	}
	userId, pErr := strconv.ParseInt(verifiedToken.Sub, 10, 64)
	if pErr != nil {
		return nil, domain.NewDomainError("ParseInt failed", domain.Internal, pErr)
	}
	user, err := u.GetUserByID(userId, ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserUsecase) GetUserByOAuthProviderAndEmailOrCreateIfAbsent(
	oAuthId string,
	oAuthProvider string,
	name string,
	email string,
	profileImageUrl *string,
	ctx context.Context,
) (*domain.User, *domain.DomainError) {
	userFromDb, err := u.GetUserByOAuthProviderAndEmail(oAuthProvider, email, ctx)
	if err != nil {
		if domainErr, ok := domain.IsDomainError(err); ok {
			if domainErr.Code == domain.NotFound {
				createdUser, err := u.CreateUser(&domain.User{
					ID:              0,
					OAuthId:         oAuthId,
					OAuthProvider:   oAuthProvider,
					Name:            name,
					Email:           email,
					ProfileImageUrl: profileImageUrl,
				}, ctx)
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

func (u *UserUsecase) CreateUser(user *domain.User, ctx context.Context) (*domain.User, *domain.DomainError) {
	return u.repo.CreateUser(user, ctx)
}

func (u *UserUsecase) GetUserByID(id int64, ctx context.Context) (*domain.User, *domain.DomainError) {
	return u.repo.GetUserByID(id, ctx)
}

func (u *UserUsecase) GetUserByOAuthProviderAndEmail(oAuthProvider string, email string, ctx context.Context) (*domain.User, *domain.DomainError) {
	return u.repo.GetUserByOAuthProviderAndEmail(oAuthProvider, email, ctx)
}

func (u *UserUsecase) DeleteUser(id int64, ctx context.Context) *domain.DomainError {
	return u.repo.DeleteUser(id, ctx)
}
