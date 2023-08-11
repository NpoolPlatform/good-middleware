package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// RequiredGood holds the schema definition for the RequiredGood entity.
type RequiredGood struct {
	ent.Schema
}

func (RequiredGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the RequiredGood.
func (RequiredGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("main_good_id", uuid.UUID{}),
		field.
			UUID("required_good_id", uuid.UUID{}),
		field.
			Bool("must").
			Optional().
			Default(false),
		field.
			Bool("commission").
			Optional().
			Default(false),
	}
}

// Edges of the RequiredGood.
func (RequiredGood) Edges() []ent.Edge {
	return nil
}

func (RequiredGood) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "main_good_id", "required_good_id"),
	}
}
