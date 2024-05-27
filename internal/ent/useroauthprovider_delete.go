// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/d0lim/turnstile/internal/ent/predicate"
	"github.com/d0lim/turnstile/internal/ent/useroauthprovider"
)

// UserOauthProviderDelete is the builder for deleting a UserOauthProvider entity.
type UserOauthProviderDelete struct {
	config
	hooks    []Hook
	mutation *UserOauthProviderMutation
}

// Where appends a list predicates to the UserOauthProviderDelete builder.
func (uopd *UserOauthProviderDelete) Where(ps ...predicate.UserOauthProvider) *UserOauthProviderDelete {
	uopd.mutation.Where(ps...)
	return uopd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (uopd *UserOauthProviderDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, uopd.sqlExec, uopd.mutation, uopd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (uopd *UserOauthProviderDelete) ExecX(ctx context.Context) int {
	n, err := uopd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (uopd *UserOauthProviderDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(useroauthprovider.Table, sqlgraph.NewFieldSpec(useroauthprovider.FieldID, field.TypeInt64))
	if ps := uopd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, uopd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	uopd.mutation.done = true
	return affected, err
}

// UserOauthProviderDeleteOne is the builder for deleting a single UserOauthProvider entity.
type UserOauthProviderDeleteOne struct {
	uopd *UserOauthProviderDelete
}

// Where appends a list predicates to the UserOauthProviderDelete builder.
func (uopdo *UserOauthProviderDeleteOne) Where(ps ...predicate.UserOauthProvider) *UserOauthProviderDeleteOne {
	uopdo.uopd.mutation.Where(ps...)
	return uopdo
}

// Exec executes the deletion query.
func (uopdo *UserOauthProviderDeleteOne) Exec(ctx context.Context) error {
	n, err := uopdo.uopd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{useroauthprovider.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (uopdo *UserOauthProviderDeleteOne) ExecX(ctx context.Context) {
	if err := uopdo.Exec(ctx); err != nil {
		panic(err)
	}
}