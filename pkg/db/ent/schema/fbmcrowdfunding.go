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

// FbmCrowdFunding holds the schema definition for the FbmCrowdFunding entity.
type FbmCrowdFunding struct {
	ent.Schema
}

func (FbmCrowdFunding) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the FbmCrowdFunding.
func (FbmCrowdFunding) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("min_deposit_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("delivery_at").
			Optional().
			Default(0),
		field.
			Other("target_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("deposit_start_at").
			Optional().
			Default(0),
		field.
			Uint32("deposit_end_at").
			Optional().
			Default(0),
		field.
			String("contract_address").
			Optional().
			Default(""),
		field.
			UUID("deposit_coin_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Bool("redeemable").
			Optional().
			Default(true),
		field.
			Uint32("redeem_delay_hours").
			Optional().
			Default(8), //nolint
		field.
			String("duration_display_type").
			Optional().
			Default(types.GoodDurationType_GoodDurationByDay.String()),
		field.
			Uint32("duration_seconds").
			Optional().
			Default(365), //nolint
	}
}

// Edges of the FbmCrowdFunding.
func (FbmCrowdFunding) Edges() []ent.Edge {
	return nil
}

func (FbmCrowdFunding) Indexes() []ent.Index {
	return nil
}
