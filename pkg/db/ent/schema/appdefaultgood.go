package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	"github.com/google/uuid"
)

// AppDefaultGood holds the schema definition for the AppDefaultGood entity.
type AppDefaultGood struct {
	ent.Schema
}

func (AppDefaultGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (AppDefaultGood) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_id", uuid.UUID{}).
			Optional(),
		field.
			UUID("good_id", uuid.UUID{}).
			Optional(),
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional(),
		field.
			String("good_type").
			Optional().
			Default(types.GoodType_PowerRental.String()),
	}
}

// Edges of the AppDefaultGood.
func (AppDefaultGood) Edges() []ent.Edge {
	return nil
}
