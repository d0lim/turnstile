package token

import "github.com/d0lim/turnstile/internal/core/domain"

type TokenManager interface {
	IssueAccessToken(sub string) (string, *domain.DomainError)
	IssueRefreshToken(sub string) (string, *domain.DomainError)
	VerifyAccessToken(tokenString string) (*domain.Token, *domain.DomainError)
	VerifyRefreshToken(tokenString string) (*domain.Token, *domain.DomainError)
}
