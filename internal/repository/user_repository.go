package repository

import (
	"context"
	"github.com/d0lim/turnstile/ent"
	"github.com/d0lim/turnstile/ent/user"
	"github.com/d0lim/turnstile/internal/domain"
)

type UserRepository interface {
	CreateUser(user *domain.User, ctx context.Context) (*domain.User, error)
	GetUserByID(id int64, ctx context.Context) (*domain.User, error)
	GetUserByOAuthProviderAndEmail(oAuthProvider string, email string, ctx context.Context) (*domain.User, error)
	DeleteUser(id int64, ctx context.Context) error
}

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) UserRepository {
	return &userRepository{client: client}
}

func (r *userRepository) CreateUser(user *domain.User, ctx context.Context) (*domain.User, error) {
	u, err := r.client.User.Create().
		SetOAuthID(user.OAuthId).
		SetOAuthProvider(user.OAuthProvider).
		SetEmail(user.Email).
		SetName(user.Name).
		SetNillableProfileImageURL(user.ProfileImageUrl).
		Save(ctx)
	if err != nil {
		return nil, err
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

func (r *userRepository) GetUserByID(id int64, ctx context.Context) (*domain.User, error) {
	a, err := r.client.User.Get(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}
		return nil, err
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

func (r *userRepository) GetUserByOAuthProviderAndEmail(oAuthProvider string, email string, ctx context.Context) (*domain.User, error) {
	u, err := r.client.User.Query().Where(user.OAuthProvider(oAuthProvider), user.Email(email)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, err
		}
		return nil, err
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

func (r *userRepository) DeleteUser(id int64, ctx context.Context) error {
	err := r.client.User.DeleteOneID(id).Exec(ctx)
	return err
}
