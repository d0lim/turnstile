package db

import (
	"context"
	"github.com/d0lim/turnstile/ent"
	"github.com/d0lim/turnstile/ent/user"
	"github.com/d0lim/turnstile/internal/core/domain"
	"github.com/d0lim/turnstile/internal/core/ports/out/repository"
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
		SetOAuthID(user.OAuthId).
		SetOAuthProvider(user.OAuthProvider).
		SetEmail(user.Email).
		SetName(user.Name).
		SetNillableProfileImageURL(user.ProfileImageUrl).
		Save(context.Background())
	return err
}

func (r *userRepository) GetUserByID(id int64) (*domain.User, error) {
	a, err := r.client.User.Get(context.Background(), id)
	if err != nil {
		return nil, errors.Wrap(err, "getting core by ID")
	}

	return &domain.User{
		ID:              a.ID,
		OAuthId:         a.OAuthID,
		OAuthProvider:   a.OAuthProvider,
		Email:           a.Email,
		Name:            a.Name,
		ProfileImageUrl: a.ProfileImageURL,
	}, nil
}

func (r *userRepository) GetUserByOAuthProviderAndEmail(oAuthProvider string, email string) (*domain.User, error) {
	u, err := r.client.User.Query().Where(user.OAuthProvider(oAuthProvider), user.Email(email)).Only(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "getting core by Email")
	}

	return &domain.User{
		ID:              u.ID,
		OAuthId:         u.OAuthID,
		OAuthProvider:   u.OAuthProvider,
		Email:           u.Email,
		Name:            u.Name,
		ProfileImageUrl: u.ProfileImageURL,
	}, nil
}

func (r *userRepository) DeleteUser(id int64) error {
	return r.client.User.DeleteOneID(id).Exec(context.Background())
}
