package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// GoodCoin holds the schema definition for the GoodCoin entity.
type GoodCoin struct {
	ent.Schema
}

func (GoodCoin) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the GoodCoin.
func (GoodCoin) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Bool("main").
			Optional().
			Default(true),
		field.
			Int32("index").
			Optional().
			Default(0),
	}
}

// Edges of the GoodCoin.
func (GoodCoin) Edges() []ent.Edge {
	return nil
}

func (GoodCoin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id"),
	}
}
