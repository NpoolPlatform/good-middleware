//nolint:dupl
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
)

// VendorBrand holds the schema definition for the VendorBrand entity.
type VendorBrand struct {
	ent.Schema
}

func (VendorBrand) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the VendorBrand.
func (VendorBrand) Fields() []ent.Field {
	const maxLen = 128
	return []ent.Field{
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
