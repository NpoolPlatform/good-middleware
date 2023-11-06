package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TopMost holds the schema definition for the TopMost entity.
type TopMost struct {
	ent.Schema
}

func (TopMost) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the TopMost.
func (TopMost) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}),
		field.
			String("top_most_type").
			Optional().
			Default(types.GoodTopMostType_DefaultGoodTopMostType.String()),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			String("message").
			Optional().
			Default(""),
		field.
			JSON("posters", []string{}).
			Optional().
			Default([]string{}),
		field.
			Uint32("start_at").
			Optional().
			Default(0),
		field.
			Uint32("end_at").
			Optional().
			Default(0),
		field.
			Other("threshold_credits", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("register_elapsed_seconds").
			Optional().
			Default(0),
		field.
			Uint32("threshold_purchases").
			Optional().
			Default(0),
		field.
			Other("threshold_payment_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("kyc_must").
			Optional().
			Default(true),
	}
}

// Edges of the TopMost.
func (TopMost) Edges() []ent.Edge {
	return nil
}
