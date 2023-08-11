package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExtraInfo holds the schema definition for the ExtraInfo entity.
type ExtraInfo struct {
	ent.Schema
}

func (ExtraInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the ExtraInfo.
func (ExtraInfo) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			JSON("posters", []string{}).
			Optional().
			Default([]string{}),
		field.
			JSON("labels", []string{}).
			Optional().
			Default([]string{}),
		field.
			Uint32("vote_count").
			Optional().
			Default(0),
		field.
			Other("rating_v1", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the ExtraInfo.
func (ExtraInfo) Edges() []ent.Edge {
	return nil
}

func (ExtraInfo) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id"),
	}
}
