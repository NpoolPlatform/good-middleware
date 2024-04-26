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

// AppPowerRental holds the schema definition for the AppPowerRental entity.
type AppPowerRental struct {
	ent.Schema
}

func (AppPowerRental) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (AppPowerRental) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("service_start_at").
			Optional().
			Default(0),
		field.
			String("cancel_mode").
			Optional().
			Default(types.CancelMode_Uncancellable.String()),
		field.
			Uint32("cancelable_before_start_seconds").
			Optional().
			Default(0),
		field.
			Bool("enable_set_commission").
			Optional().
			Default(false),
		field.
			Other("min_order_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("max_order_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("max_user_amount", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("min_order_duration_seconds").
			Optional().
			Default(0),
		field.
			Uint32("max_order_duration_seconds").
			Optional().
			Default(0),
		field.
			Other("unit_price", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Uint32("sale_start_at").
			Optional().
			Default(0),
		field.
			Uint32("sale_end_at").
			Optional().
			Default(0),
		field.
			String("sale_mode").
			Optional().
			Default(types.GoodSaleMode_GoodSaleModeMainnetSpot.String()),
		field.
			Bool("fixed_duration").
			Optional().
			Default(true),
		field.
			Bool("package_with_requireds").
			Optional().
			Default(true),
	}
}

// Edges of the AppPowerRental.
func (AppPowerRental) Edges() []ent.Edge {
	return nil
}

func (AppPowerRental) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
