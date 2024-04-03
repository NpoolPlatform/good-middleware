package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TopMostConstraint holds the schema definition for the TopMostConstraint entity.
type TopMostConstraint struct {
	ent.Schema
}

func (TopMostConstraint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

//nolint:funlen
func (TopMostConstraint) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("top_most_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("constraint").
			Optional().
			Default(types.GoodTopMostConstraint_TopMostKycMust.String()),
		field.
			Other("target_value", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the TopMostConstraint.
func (TopMostConstraint) Edges() []ent.Edge {
	return nil
}

func (TopMostConstraint) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("top_most_id"),
	}
}
