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

// AppStockLock holds the schema definition for the AppStockLock entity.
type AppStockLock struct {
	ent.Schema
}

func (AppStockLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppStockLock.
func (AppStockLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_stock_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Other("units", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Other("app_spot_units", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			String("lock_state").
			Optional().
			Default(types.AppStockLockState_AppStockLocked.String()),
		field.
			String("charge_back_state").
			Optional().
			Default(types.AppStockLockState_DefaultAppStockLockState.String()),
		field.
			UUID("ex_lock_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
	}
}

func (AppStockLock) Indexes() []ent.Index {
	return nil
}
