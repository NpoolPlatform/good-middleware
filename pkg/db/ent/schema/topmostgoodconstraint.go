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

// TopMostGoodConstraint holds the schema definition for the TopMostGoodConstraint entity.
type TopMostGoodConstraint struct {
	ent.Schema
}

func (TopMostGoodConstraint) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

//nolint:funlen
func (TopMostGoodConstraint) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("top_most_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("constraint").
			Optional().
			Default(types.GoodTopMostConstraint_TopMostKycMust.String()),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the TopMostGoodConstraint.
func (TopMostGoodConstraint) Edges() []ent.Edge {
	return nil
}

func (TopMostGoodConstraint) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("top_most_good_id"),
	}
}
