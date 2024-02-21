//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// AppSimulateGood holds the schema definition for the AppSimulateGood entity.
type AppSimulateGood struct {
	ent.Schema
}

func (AppSimulateGood) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (AppSimulateGood) Fields() []ent.Field {
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
	}
}

// Edges of the AppSimulateGood.
func (AppSimulateGood) Edges() []ent.Edge {
	return nil
}
