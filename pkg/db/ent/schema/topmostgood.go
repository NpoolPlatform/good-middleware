package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
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
	}
}

// Fields of the TopMostGood.
func (TopMostGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			UUID("app_good_id", uuid.UUID{}),
		field.
			UUID("coin_type_id", uuid.UUID{}),
		field.
			UUID("top_most_id", uuid.UUID{}),
		field.
			Uint32("display_index").
			Optional().
			Default(0),
		field.
			JSON("posters", []string{}).
			Optional().
			Default([]string{}),
		field.
			Other("price", decimal.Decimal{}).
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
