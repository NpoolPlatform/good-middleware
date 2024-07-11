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

// AppGoodLabel holds the schema definition for the AppGoodLabel entity.
type AppGoodLabel struct {
	ent.Schema
}

func (AppGoodLabel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		crudermixin.AutoIDMixin{},
	}
}

func (AppGoodLabel) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("app_good_id", uuid.UUID{}).
			Optional().
			Default(func() uuid.UUID {
				return uuid.Nil
			}),
		field.
			String("icon").
			Optional().
			Default(""),
		field.
			String("icon_bg_color").
			Optional().
			Default(""),
		field.
			String("label").
			Optional().
			Default(types.GoodLabel_DefaultGoodLabel.String()),
		field.
			String("label_bg_color").
			Optional().
			Default(""),
		field.
			Uint8("index").
			Optional().
			Default(0),
	}
}

// Edges of the AppGoodLabel.
func (AppGoodLabel) Edges() []ent.Edge {
	return nil
}

func (AppGoodLabel) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_good_id"),
	}
}
