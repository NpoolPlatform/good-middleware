//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// DevicePoster holds the schema definition for the DevicePoster entity.
type DevicePoster struct {
	ent.Schema
}

func (DevicePoster) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (DevicePoster) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("device_type_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("poster").
			Optional().
			Default(""),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the DevicePoster.
func (DevicePoster) Edges() []ent.Edge {
	return nil
}

func (DevicePoster) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("device_type_id"),
	}
}
