package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TopMostGood holds the schema definition for the TopMostGood entity.
type TopMostGood struct {
	ent.Schema
}

func (TopMostGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the TopMostGood.
func (TopMostGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("top_most_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("display_index").
			Optional().
			Default(0),
		field.
			Other("unit_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the TopMostGood.
func (TopMostGood) Edges() []ent.Edge {
	return nil
}
