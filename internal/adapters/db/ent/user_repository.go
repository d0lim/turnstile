package ent

import (
	"context"
	"github.com/d0lim/turnstile/internal/core/domain"
	"github.com/d0lim/turnstile/internal/core/ports/out/repository"
	"github.com/d0lim/turnstile/internal/ent"
	"github.com/d0lim/turnstile/internal/ent/user"
	"github.com/pkg/errors"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) repository.UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	_, err := r.client.User.Create().
		SetEmail(user.Email).
		SetNickname(user.Nickname).
		SetNillableProfileImageURL(nil).
		Save(context.Background())
	return err
}

func (r *userRepository) GetUserByID(id int64) (*domain.User, error) {
	a, err := r.client.User.Get(context.Background(), id)
	if err != nil {
		return nil, errors.Wrap(err, "getting core by ID")
	}

	return &domain.User{
		ID:       a.ID,
		Nickname: a.Nickname,
		Email:    a.Email,
	}, nil
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	u, err := r.client.User.Query().Where(user.Email(email)).Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "getting core by Email")
	}

	return &domain.User{
		ID:       u.ID,
		Nickname: u.Nickname,
		Email:    u.Email,
	}, nil
}

func (r *userRepository) DeleteUser(id int64) error {
	return r.client.User.DeleteOneID(id).Exec(context.Background())
}
