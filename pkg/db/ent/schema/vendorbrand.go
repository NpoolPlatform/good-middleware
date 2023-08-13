package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"

	"github.com/google/uuid"
)

// VendorBrand holds the schema definition for the VendorBrand entity.
type VendorBrand struct {
	ent.Schema
}

func (VendorBrand) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the VendorBrand.
func (VendorBrand) Fields() []ent.Field {
	const maxLen = 128
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.String("name").
			Optional().
			Default("").
			MaxLen(maxLen),
		field.String("logo").
			Optional().
			Default("").
			MaxLen(maxLen),
	}
}

// Edges of the VendorBrand.
func (VendorBrand) Edges() []ent.Edge {
	return nil
}