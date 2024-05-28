package schema

import (
	"context"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/d0lim/turnstile/internal/ent/hook"
	"time"
)

// TimestampMixin defines the schema for timestamp fields.
type TimestampMixin struct {
	mixin.Schema
}

// Fields of the TimestampMixin.
func (TimestampMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Hooks of the TimestampMixin.
func (TimestampMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					if m.Op().Is(ent.OpCreate) {
						if _, ok := m.Field("created_at"); !ok {
							m.SetField("created_at", time.Now())
						}
						if _, ok := m.Field("updated_at"); !ok {
							m.SetField("updated_at", time.Now())
						}
					} else if m.Op().Is(ent.OpUpdate) || m.Op().Is(ent.OpUpdateOne) {
						m.SetField("updated_at", time.Now())
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}
