package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	crudermixin "github.com/NpoolPlatform/libent-cruder/pkg/mixin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Recommend holds the schema definition for the Recommend entity.
type Recommend struct {
	ent.Schema
}

func (Recommend) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

// Fields of the Recommend.
func (Recommend) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			UUID("recommender_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("message").
			Optional().
			Default(""),
		field.
			Other("recommend_index", decimal.Decimal{}).
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(37,18)",
			}).
			Optional().
			Default(decimal.Decimal{}),
		field.
			Bool("hide").
			Optional().
			Default(false),
		field.
			String("hide_reason").
			Optional().
			Default(types.GoodCommentHideReason_DefaultGoodCommentHideReason.String()),
	}
}

// Indexes of the Recommend.
func (Recommend) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}

func (Recommend) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "app_good_recommends"},
	}
}
