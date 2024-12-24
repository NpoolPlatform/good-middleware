package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	"github.com/google/uuid"
)

// AppPledge holds the schema definition for the AppPledge entity.
type AppPledge struct {
	ent.Schema
}

func (AppPledge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (AppPledge) Fields() []ent.Field {
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
			String("start_mode").
			Optional().
			Default(types.GoodStartMode_GoodStartModeNextDay.String()),
		field.
			Bool("enable_set_commission").
			Optional().
			Default(false),
	}
}

// Edges of the AppPledge.
func (AppPledge) Edges() []ent.Edge {
	return nil
}

func (AppPledge) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}