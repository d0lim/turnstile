package service

import (
	"fmt"
	"github.com/d0lim/turnstile/internal/domain"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type TokenService interface {
	IssueAccessToken(sub string) (string, error)
	IssueRefreshToken(sub string) (string, error)
	VerifyAccessToken(tokenString string) (*domain.Token, error)
	VerifyRefreshToken(tokenString string) (*domain.Token, error)
	GetBlacklistReason(tokenString string) (string, error)
	AddToRefreshTokenBlacklist(tokenString string, reason string) error
}

type tokenService struct {
	config *config.JwtConfig
	redis  *config.RedisConfig
}

func NewTokenService(config *config.JwtConfig, redis *config.RedisConfig) TokenService {
	return &tokenService{
		config: config,
		redis:  redis,
	}
}

func (s *tokenService) IssueAccessToken(sub string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString(s.config.AccessSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *tokenService) IssueRefreshToken(sub string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	})
	tokenString, err := token.SignedString(s.config.RefreshSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (s *tokenService) VerifyAccessToken(tokenString string) (*domain.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.config.AccessSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return s.convertClaims(claims)
	} else {
		return nil, err
	}
}

func (s *tokenService) VerifyRefreshToken(tokenString string) (*domain.Token, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return s.config.RefreshSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return s.convertClaims(claims)
	} else {
		return nil, err
	}
}

func (s *tokenService) GetBlacklistReason(tokenString string) (string, error) {
	prefix := "refresh_token:"
	reason, err := s.redis.Store.Get(prefix + tokenString)
	if err != nil {
		return "", err
	}
	reasonStr := string(reason)

	return reasonStr, nil
}

func (s *tokenService) AddToRefreshTokenBlacklist(tokenString string, reason string) error {
	prefix := "refresh_token:"
	err := s.redis.Store.Set(prefix+tokenString, []byte(reason), time.Hour*168)
	if err != nil {
		return err
	}
	return nil
}

func (s *tokenService) convertClaims(claims jwt.Claims) (*domain.Token, error) {
	sub, err := claims.GetSubject()
	if err != nil {
		return nil, err
	}
	iat, err := claims.GetIssuedAt()
	if err != nil {
		return nil, err
	}
	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, err
	}

	return &domain.Token{
		Sub: sub,
		Iat: iat.Unix(),
		Exp: exp.Unix(),
	}, nil
}
