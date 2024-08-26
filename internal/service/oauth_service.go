package service

import (
	"context"
	"encoding/json"
	"github.com/d0lim/turnstile/internal/client/dto"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type OauthService interface {
	GoogleLoginRedirectUri(redirectUri string) string
	GoogleHandleCallback(code string, ctx context.Context) (*dto.GoogleUserResponse, error)
}

type oauthService struct {
	oauthConfig *config.OAuthConfig
}

func NewOauthService(oauthConfig *config.OAuthConfig) OauthService {
	return &oauthService{oauthConfig: oauthConfig}
}

func (s *oauthService) GoogleLoginRedirectUri(redirectUri string) string {
	state := uuid.NewString()

	oAuthConfig := s.oauthConfig.Google
	oAuthConfig.RedirectURL = redirectUri

	return oAuthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline)
}

func (s *oauthService) GoogleHandleCallback(code string, ctx context.Context) (*dto.GoogleUserResponse, error) {
	token, err := s.oauthConfig.Google.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	client := s.oauthConfig.Google.Client(ctx, token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer userInfo.Body.Close()

	var user dto.GoogleUserResponse
	if err := json.NewDecoder(userInfo.Body).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}
