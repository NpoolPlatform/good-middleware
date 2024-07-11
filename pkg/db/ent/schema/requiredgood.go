//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// RequiredGood holds the schema definition for the RequiredGood entity.
type RequiredGood struct {
	ent.Schema
}

func (RequiredGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the RequiredGood.
func (RequiredGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("main_good_id", uuid.UUID{}),
		field.
			UUID("required_good_id", uuid.UUID{}),
		field.
			Bool("must").
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
		index.Fields("main_good_id", "required_good_id"),
	}
}
