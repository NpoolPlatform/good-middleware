package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

func (Comment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("order_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("content").
			Optional().
			Default(""),
		field.
			UUID("reply_to_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("anonymous").
			Optional().
			Default(false),
		field.
			Bool("trial_user").
			Optional().
			Default(false),
		field.
			Bool("purchased_user").
			Optional().
			Default(false),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}

func (Comment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "app_good_id"),
	}
}
