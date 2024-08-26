package service

import (
	"context"
	"fmt"
	"github.com/d0lim/turnstile/ent"
	"github.com/d0lim/turnstile/internal/domain"
	"github.com/d0lim/turnstile/internal/repository"
	"strconv"
)

type UserService interface {
	Login(oAuthId string, oAuthProvider string, name string, email string, profileImageUrl *string, ctx context.Context) (*domain.TokenPair, error)
	Authenticate(tokenString string, ctx context.Context) (*domain.User, error)
	Refresh(refreshTokenString string) (*domain.TokenPair, error)
}

type userService struct {
	repository   repository.UserRepository
	tokenService TokenService
}

func NewUserService(repository repository.UserRepository, tokenService TokenService) UserService {
	return &userService{
		repository:   repository,
		tokenService: tokenService,
	}
}

func (s *userService) Login(oAuthId string, oAuthProvider string, name string, email string, profileImageUrl *string, ctx context.Context) (*domain.TokenPair, error) {
	user, err := s.createUserIfAbsent(
		oAuthId, oAuthProvider, name, email, profileImageUrl, ctx)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenService.IssueAccessToken(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.tokenService.IssueRefreshToken(strconv.FormatInt(user.ID, 10))
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) Authenticate(tokenString string, ctx context.Context) (*domain.User, error) {
	verifiedToken, err := s.tokenService.VerifyAccessToken(tokenString)
	if err != nil {
		return nil, err
	}
	userId, err := strconv.ParseInt(verifiedToken.Sub, 10, 64)
	if err != nil {
		return nil, err
	}
	foundUser, err := s.repository.GetUserByID(userId, ctx)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (s *userService) Refresh(tokenString string) (*domain.TokenPair, error) {
	reason, err := s.tokenService.GetBlacklistReason(tokenString)
	if err != nil {
		return nil, err
	}
	if reason != "" {
		return nil, fmt.Errorf("already used refresh token, reason: %s", reason)
	}

	verifiedRefreshToken, err := s.tokenService.VerifyRefreshToken(tokenString)
	if err != nil {
		return nil, err
	}

	userId := verifiedRefreshToken.Sub
	accessToken, err := s.tokenService.IssueAccessToken(userId)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.tokenService.IssueRefreshToken(userId)
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *userService) createUserIfAbsent(oAuthId string, oAuthProvider string, name string, email string, profileImageUrl *string, ctx context.Context) (*domain.User, error) {
	foundUser, err := s.repository.GetUserByOAuthProviderAndEmail(oAuthProvider, email, ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			newUser, err := s.repository.CreateUser(&domain.User{
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
			return newUser, nil
		}
		return nil, err
	}
	return foundUser, nil
}
