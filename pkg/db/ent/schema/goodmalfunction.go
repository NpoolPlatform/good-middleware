package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// GoodMalfunction holds the schema definition for the GoodMalfunction entity.
type GoodMalfunction struct {
	ent.Schema
}

func (GoodMalfunction) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the GoodMalfunction.
func (GoodMalfunction) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			Text("message").
			Optional().
			Default(""),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("duration_seconds").
			Optional().
			Default(0),
		field.
			Uint32("compensate_seconds").
			Optional().
			Default(0),
	}
}

// Edges of the GoodMalfunction.
func (GoodMalfunction) Edges() []ent.Edge {
	return nil
}

func (GoodMalfunction) Indexes() []ent.Index {
	return nil
}
