package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

func (Like) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("good_id", uuid.UUID{}).
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
			Bool("like"),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return nil
}

func (Like) Indexes() []ent.Index {
	return nil
}
