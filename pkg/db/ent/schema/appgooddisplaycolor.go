package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// AppGoodDisplayColor holds the schema definition for the AppGoodDisplayColor entity.
type AppGoodDisplayColor struct {
	ent.Schema
}

func (AppGoodDisplayColor) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

//nolint:funlen
func (AppGoodDisplayColor) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("color").
			Optional().
			Default(""),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the AppGoodDisplayColor.
func (AppGoodDisplayColor) Edges() []ent.Edge {
	return nil
}

func (AppGoodDisplayColor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
