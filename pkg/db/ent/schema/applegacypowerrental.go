package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppLegacyPowerRental holds the schema definition for the AppLegacyPowerRental entity.
type AppLegacyPowerRental struct {
	ent.Schema
}

func (AppLegacyPowerRental) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

//nolint:funlen
func (AppLegacyPowerRental) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("technique_fee_ratio", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
	}
}

// Edges of the AppLegacyPowerRental.
func (AppLegacyPowerRental) Edges() []ent.Edge {
	return nil
}

func (AppLegacyPowerRental) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id", "app_id"),
	}
}
