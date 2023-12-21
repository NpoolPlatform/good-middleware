package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// DeviceInfo holds the schema definition for the DeviceInfo entity.
type DeviceInfo struct {
	ent.Schema
}

func (DeviceInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the DeviceInfo.
func (DeviceInfo) Fields() []ent.Field {
	maxLen := 64
	return []ent.Field{
		field.
			String("type").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.
			String("manufacturer").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.
			Uint32("power_consumption").
			Optional().
			Default(0),
		field.
			Uint32("shipment_at").
			Optional().
			Default(0),
		field.
			JSON("posters", []string{}).
			Optional().
			Default([]string{}),
	}
}

// Edges of the DeviceInfo.
func (DeviceInfo) Edges() []ent.Edge {
	return nil
}
