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

// Pledge holds the schema definition for the Pledge entity.
type Pledge struct {
	ent.Schema
}

func (Pledge) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Pledge.
func (Pledge) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("contract_code_url").
			Optional().
			Default(""),
		field.
			String("contract_code_branch").
			Optional().
			Default(""),
		field.
			String("contract_state").
			Optional().
			Default(types.ContractState_ContractWaitDeployment.String()),
	}
}

// Edges of the Pledge.
func (Pledge) Edges() []ent.Edge {
	return nil
}

func (Pledge) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("good_id"),
	}
}
