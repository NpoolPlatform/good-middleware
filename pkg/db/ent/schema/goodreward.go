package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodReward holds the schema definition for the GoodReward entity.
type GoodReward struct {
	ent.Schema
}

func (GoodReward) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the GoodReward.
func (GoodReward) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("good_id", uuid.UUID{}),
		field.
			String("benefit_state").
			Optional().
			Default(types.BenefitState_BenefitWait.String()),
		field.
			Uint32("last_benefit_at").
			Optional().
			Default(0),
		field.
			JSON("benefit_tids", []uuid.UUID{}).
			Optional().
			Default([]uuid.UUID{}),
		field.
			Other("next_benefit_start_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("last_benefit_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the GoodReward.
func (GoodReward) Edges() []ent.Edge {
	return nil
}

func (GoodReward) Indexes() []ent.Index {
	return nil
}
