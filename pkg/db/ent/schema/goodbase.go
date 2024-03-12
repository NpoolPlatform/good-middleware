package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

// Good holds the schema definition for the Good entity.
type GoodBase struct {
	ent.Schema
}

func (GoodBase) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the GoodBase.
func (GoodBase) Fields() []ent.Field {
	const benefitHours = 24
	return []ent.Field{
		field.
			String("good_type").
			Optional().
			Default(types.GoodType_DefaultGoodType.String()),
		field.
			String("benefit_type").
			Optional().
			Default(types.BenefitType_DefaultBenefitType.String()),
		field.
			String("name").
			Optional().
			Default(""),
		field.
			Uint32("service_start_at").
			Optional().
			Default(0),
		field.
			String("start_mode").
			Optional().
			Default(types.GoodStartMode_GoodStartModeNextDay.String()),
		field.
			Bool("test_only").
			Optional().
			Default(false),
		field.
			Uint32("benefit_interval_hours").
			Optional().
			Default(benefitHours),
		field.
			Bool("purchasable").
			Optional().
			Default(false),
		field.
			Bool("online").
			Optional().
			Default(false),
	}
}

// Edges of the GoodBase.
func (GoodBase) Edges() []ent.Edge {
	return nil
}

func (GoodBase) Indexes() []ent.Index {
	return nil
}
