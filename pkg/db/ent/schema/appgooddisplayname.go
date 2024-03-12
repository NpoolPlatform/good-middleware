package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// AppGoodDisplayName holds the schema definition for the AppGoodDisplayName entity.
type AppGoodDisplayName struct {
	ent.Schema
}

func (AppGoodDisplayName) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

//nolint:funlen
func (AppGoodDisplayName) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the AppGoodDisplayName.
func (AppGoodDisplayName) Edges() []ent.Edge {
	return nil
}

func (AppGoodDisplayName) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
