package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodRewardHistory holds the schema definition for the GoodRewardHistory entity.
type GoodRewardHistory struct {
	ent.Schema
}

func (GoodRewardHistory) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the GoodRewardHistory.
func (GoodRewardHistory) Fields() []ent.Field {
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
			Uint32("reward_date").
			Optional().
			Default(uint32(time.Now().Unix())),
		field.
			UUID("tid", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("unit_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("unit_net_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("result").
			Optional().
			Default(basetypes.Result_DefaultResult.String()),
	}
}

// Edges of the GoodRewardHistory.
func (GoodRewardHistory) Edges() []ent.Edge {
	return nil
}

func (GoodRewardHistory) Indexes() []ent.Index {
	return nil
}