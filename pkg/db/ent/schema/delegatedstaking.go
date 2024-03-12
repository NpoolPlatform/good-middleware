package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	"github.com/google/uuid"
)

// DelegatedStaking holds the schema definition for the DelegatedStaking entity.
type DelegatedStaking struct {
	ent.Schema
}

func (DelegatedStaking) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the DelegatedStaking.
func (DelegatedStaking) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			Uint32("no_stake_redeem_delay_hours").
			Optional().
			Default(8),
		field.
			Uint32("max_redeem_delay_hours").
			Optional().
			Default(96),
		field.
			String("contract_address").
			Optional().
			Default(""),
		field.
			Uint32("no_stake_benefit_delay_hours").
			Optional().
			Default(24),
	}
}

// Edges of the DelegatedStaking.
func (DelegatedStaking) Edges() []ent.Edge {
	return nil
}

func (DelegatedStaking) Indexes() []ent.Index {
	return nil
}
