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

// StockLock holds the schema definition for the StockLock entity.
type StockLock struct {
	ent.Schema
}

func (StockLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the StockLock.
func (StockLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("stock_id", uuid.UUID{}).
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
			String("lock_state").
			Optional().
			Default(types.StockLockState_StockLocked.String()),
		field.
			String("charge_back_state").
			Optional().
			Default(types.StockLockState_DefaultStockLockState.String()),
		field.
			UUID("ex_lock_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
	}
}

func (StockLock) Indexes() []ent.Index {
	return nil
}
