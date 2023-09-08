package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/mixin"
	"github.com/google/uuid"
)

// AppStockLock holds the schema definition for the AppStockLock entity.
type AppStockLock struct {
	ent.Schema
}

func (AppStockLock) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the AppStockLock.
func (AppStockLock) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
	}
}

func (AppStockLock) Indexes() []ent.Index {
	return nil
}
