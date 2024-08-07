package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
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
			UUID("app_good_id", uuid.UUID{}).
			Optional(),
		field.
			UUID("coin_type_id", uuid.UUID{}).
			Optional(),
	}
}

// Edges of the AppDefaultGood.
func (AppDefaultGood) Edges() []ent.Edge {
	return nil
}
