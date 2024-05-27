// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/d0lim/turnstile/internal/ent/useroauthprovider"
)

// UserOauthProviderCreate is the builder for creating a UserOauthProvider entity.
type UserOauthProviderCreate struct {
	config
	mutation *UserOauthProviderMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (uopc *UserOauthProviderCreate) SetUserID(i int64) *UserOauthProviderCreate {
	uopc.mutation.SetUserID(i)
	return uopc
}

// SetOauthProvider sets the "oauth_provider" field.
func (uopc *UserOauthProviderCreate) SetOauthProvider(s string) *UserOauthProviderCreate {
	uopc.mutation.SetOauthProvider(s)
	return uopc
}

// SetOauthUserID sets the "oauth_user_id" field.
func (uopc *UserOauthProviderCreate) SetOauthUserID(s string) *UserOauthProviderCreate {
	uopc.mutation.SetOauthUserID(s)
	return uopc
}

// SetID sets the "id" field.
func (uopc *UserOauthProviderCreate) SetID(i int64) *UserOauthProviderCreate {
	uopc.mutation.SetID(i)
	return uopc
}

// Mutation returns the UserOauthProviderMutation object of the builder.
func (uopc *UserOauthProviderCreate) Mutation() *UserOauthProviderMutation {
	return uopc.mutation
}

// Save creates the UserOauthProvider in the database.
func (uopc *UserOauthProviderCreate) Save(ctx context.Context) (*UserOauthProvider, error) {
	return withHooks(ctx, uopc.sqlSave, uopc.mutation, uopc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uopc *UserOauthProviderCreate) SaveX(ctx context.Context) *UserOauthProvider {
	v, err := uopc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uopc *UserOauthProviderCreate) Exec(ctx context.Context) error {
	_, err := uopc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uopc *UserOauthProviderCreate) ExecX(ctx context.Context) {
	if err := uopc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uopc *UserOauthProviderCreate) check() error {
	if _, ok := uopc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserOauthProvider.user_id"`)}
	}
	if _, ok := uopc.mutation.OauthProvider(); !ok {
		return &ValidationError{Name: "oauth_provider", err: errors.New(`ent: missing required field "UserOauthProvider.oauth_provider"`)}
	}
	if v, ok := uopc.mutation.OauthProvider(); ok {
		if err := useroauthprovider.OauthProviderValidator(v); err != nil {
			return &ValidationError{Name: "oauth_provider", err: fmt.Errorf(`ent: validator failed for field "UserOauthProvider.oauth_provider": %w`, err)}
		}
	}
	if _, ok := uopc.mutation.OauthUserID(); !ok {
		return &ValidationError{Name: "oauth_user_id", err: errors.New(`ent: missing required field "UserOauthProvider.oauth_user_id"`)}
	}
	if v, ok := uopc.mutation.OauthUserID(); ok {
		if err := useroauthprovider.OauthUserIDValidator(v); err != nil {
			return &ValidationError{Name: "oauth_user_id", err: fmt.Errorf(`ent: validator failed for field "UserOauthProvider.oauth_user_id": %w`, err)}
		}
	}
	return nil
}

func (uopc *UserOauthProviderCreate) sqlSave(ctx context.Context) (*UserOauthProvider, error) {
	if err := uopc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uopc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uopc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	uopc.mutation.id = &_node.ID
	uopc.mutation.done = true
	return _node, nil
}

func (uopc *UserOauthProviderCreate) createSpec() (*UserOauthProvider, *sqlgraph.CreateSpec) {
	var (
		_node = &UserOauthProvider{config: uopc.config}
		_spec = sqlgraph.NewCreateSpec(useroauthprovider.Table, sqlgraph.NewFieldSpec(useroauthprovider.FieldID, field.TypeInt64))
	)
	if id, ok := uopc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uopc.mutation.UserID(); ok {
		_spec.SetField(useroauthprovider.FieldUserID, field.TypeInt64, value)
		_node.UserID = value
	}
	if value, ok := uopc.mutation.OauthProvider(); ok {
		_spec.SetField(useroauthprovider.FieldOauthProvider, field.TypeString, value)
		_node.OauthProvider = value
	}
	if value, ok := uopc.mutation.OauthUserID(); ok {
		_spec.SetField(useroauthprovider.FieldOauthUserID, field.TypeString, value)
		_node.OauthUserID = value
	}
	return _node, _spec
}

// UserOauthProviderCreateBulk is the builder for creating many UserOauthProvider entities in bulk.
type UserOauthProviderCreateBulk struct {
	config
	err      error
	builders []*UserOauthProviderCreate
}

// Save creates the UserOauthProvider entities in the database.
func (uopcb *UserOauthProviderCreateBulk) Save(ctx context.Context) ([]*UserOauthProvider, error) {
	if uopcb.err != nil {
		return nil, uopcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(uopcb.builders))
	nodes := make([]*UserOauthProvider, len(uopcb.builders))
	mutators := make([]Mutator, len(uopcb.builders))
	for i := range uopcb.builders {
		func(i int, root context.Context) {
			builder := uopcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserOauthProviderMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, uopcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uopcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, uopcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uopcb *UserOauthProviderCreateBulk) SaveX(ctx context.Context) []*UserOauthProvider {
	v, err := uopcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uopcb *UserOauthProviderCreateBulk) Exec(ctx context.Context) error {
	_, err := uopcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uopcb *UserOauthProviderCreateBulk) ExecX(ctx context.Context) {
	if err := uopcb.Exec(ctx); err != nil {
		panic(err)
	}
}
