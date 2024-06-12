package domain

import (
	"errors"
	"fmt"
)

type ErrorCode int

const (
	Unknown ErrorCode = iota
	NotFound
	Internal
	BadRequest
)

func (e ErrorCode) String() string {
	return [...]string{"Unknown", "NotFound", "Internal"}[e]
}

type DomainError struct {
	Code  ErrorCode
	Msg   string
	Cause error
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("%s (code: %d)", e.Msg, e.Code)
}

func NewDomainError(msg string, code ErrorCode, cause error) *DomainError {
	return &DomainError{
		Code:  code,
		Msg:   msg,
		Cause: cause,
	}
}

func IsDomainError(err error) (*DomainError, bool) {
	var domainError *DomainError
	ok := errors.As(err, &domainError)
	return domainError, ok
}
