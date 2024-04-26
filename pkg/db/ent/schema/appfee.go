package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppFee holds the schema definition for the AppFee entity.
type AppFee struct {
	ent.Schema
}

func (AppFee) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the AppFee.
func (AppFee) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.UUID{}
			}),
		field.
			Other("unit_value", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		// Here we do not need duration calculate type
		// If you need the child order duration same as parent
		// just set min_order_duration_seconds the the same as the parent duration
		field.
			Uint32("min_order_duration_seconds").
			Optional().
			Default(3),
	}
}

// Edges of the AppFee.
func (AppFee) Edges() []ent.Edge {
	return nil
}

func (AppFee) Indexes() []ent.Index {
	return nil
}
