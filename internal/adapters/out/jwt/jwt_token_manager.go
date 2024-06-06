package jwt

import (
	"fmt"
	"github.com/d0lim/turnstile/internal/core/domain"
	"github.com/d0lim/turnstile/internal/core/ports/out/token"
	"github.com/d0lim/turnstile/internal/framework/config"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtTokenManager struct {
	config *config.JwtConfig
}

func NewJwtTokenManager(config *config.JwtConfig) token.TokenManager {
	return &jwtTokenManager{config: config}
}

func (m *jwtTokenManager) IssueAccessToken(sub string) (string, *domain.DomainError) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := t.SignedString(m.config.AccessSecret)
	if err != nil {
		return "", domain.NewDomainError("Error while signing access token", domain.Internal, err)
	}
	return tokenString, nil
}

func (m *jwtTokenManager) IssueRefreshToken(sub string) (string, *domain.DomainError) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 168).Unix(),
	})
	tokenString, err := t.SignedString(m.config.RefreshSecret)
	if err != nil {
		return "", domain.NewDomainError("Error while signing refresh token", domain.Internal, err)
	}
	return tokenString, nil
}

func (m *jwtTokenManager) VerifyAccessToken(tokenString string) (*domain.Token, *domain.DomainError) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return m.config.AccessSecret, nil
	})
	if err != nil {
		return nil, domain.NewDomainError("Error while parsing parsedToken", domain.Internal, err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return m.convertClaims(claims)
	} else {
		return nil, domain.NewDomainError("Token is not valid", domain.Internal, err)
	}
}

func (m *jwtTokenManager) VerifyRefreshToken(tokenString string) (*domain.Token, *domain.DomainError) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return m.config.RefreshSecret, nil
	})
	if err != nil {
		return nil, domain.NewDomainError("Error while parsing parsedToken", domain.Internal, err)
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		return m.convertClaims(claims)
	} else {
		return nil, domain.NewDomainError("Token is not valid", domain.Internal, err)
	}
}

func (m *jwtTokenManager) convertClaims(claims jwt.Claims) (*domain.Token, *domain.DomainError) {
	sub, err := claims.GetSubject()
	if err != nil {
		return nil, domain.NewDomainError("Error while parsing token subject", domain.Internal, err)
	}
	iat, err := claims.GetIssuedAt()
	if err != nil {
		return nil, domain.NewDomainError("Error while parsing token iat", domain.Internal, err)
	}
	exp, err := claims.GetExpirationTime()
	if err != nil {
		return nil, domain.NewDomainError("Error while parsing token exp", domain.Internal, err)
	}
	return &domain.Token{
		Sub: sub,
		Iat: iat.Unix(),
		Exp: exp.Unix(),
	}, nil
}
