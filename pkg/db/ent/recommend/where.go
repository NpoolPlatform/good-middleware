// Code generated by ent, DO NOT EDIT.

package recommend

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppID applies equality check predicate on the "app_id" field. It's identical to AppIDEQ.
func AppID(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// RecommenderID applies equality check predicate on the "recommender_id" field. It's identical to RecommenderIDEQ.
func RecommenderID(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommenderID), v))
	})
}

// Message applies equality check predicate on the "message" field. It's identical to MessageEQ.
func Message(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// RecommendIndex applies equality check predicate on the "recommend_index" field. It's identical to RecommendIndexEQ.
func RecommendIndex(v decimal.Decimal) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommendIndex), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppIDEQ applies the EQ predicate on the "app_id" field.
func AppIDEQ(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppID), v))
	})
}

// AppIDNEQ applies the NEQ predicate on the "app_id" field.
func AppIDNEQ(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppID), v))
	})
}

// AppIDIn applies the In predicate on the "app_id" field.
func AppIDIn(vs ...uuid.UUID) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppID), v...))
	})
}

// AppIDNotIn applies the NotIn predicate on the "app_id" field.
func AppIDNotIn(vs ...uuid.UUID) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppID), v...))
	})
}

// AppIDGT applies the GT predicate on the "app_id" field.
func AppIDGT(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppID), v))
	})
}

// AppIDGTE applies the GTE predicate on the "app_id" field.
func AppIDGTE(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppID), v))
	})
}

// AppIDLT applies the LT predicate on the "app_id" field.
func AppIDLT(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppID), v))
	})
}

// AppIDLTE applies the LTE predicate on the "app_id" field.
func AppIDLTE(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppID), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// RecommenderIDEQ applies the EQ predicate on the "recommender_id" field.
func RecommenderIDEQ(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommenderID), v))
	})
}

// RecommenderIDNEQ applies the NEQ predicate on the "recommender_id" field.
func RecommenderIDNEQ(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecommenderID), v))
	})
}

// RecommenderIDIn applies the In predicate on the "recommender_id" field.
func RecommenderIDIn(vs ...uuid.UUID) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRecommenderID), v...))
	})
}

// RecommenderIDNotIn applies the NotIn predicate on the "recommender_id" field.
func RecommenderIDNotIn(vs ...uuid.UUID) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRecommenderID), v...))
	})
}

// RecommenderIDGT applies the GT predicate on the "recommender_id" field.
func RecommenderIDGT(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecommenderID), v))
	})
}

// RecommenderIDGTE applies the GTE predicate on the "recommender_id" field.
func RecommenderIDGTE(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecommenderID), v))
	})
}

// RecommenderIDLT applies the LT predicate on the "recommender_id" field.
func RecommenderIDLT(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecommenderID), v))
	})
}

// RecommenderIDLTE applies the LTE predicate on the "recommender_id" field.
func RecommenderIDLTE(v uuid.UUID) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecommenderID), v))
	})
}

// RecommenderIDIsNil applies the IsNil predicate on the "recommender_id" field.
func RecommenderIDIsNil() predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecommenderID)))
	})
}

// RecommenderIDNotNil applies the NotNil predicate on the "recommender_id" field.
func RecommenderIDNotNil() predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecommenderID)))
	})
}

// MessageEQ applies the EQ predicate on the "message" field.
func MessageEQ(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMessage), v))
	})
}

// MessageNEQ applies the NEQ predicate on the "message" field.
func MessageNEQ(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMessage), v))
	})
}

// MessageIn applies the In predicate on the "message" field.
func MessageIn(vs ...string) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMessage), v...))
	})
}

// MessageNotIn applies the NotIn predicate on the "message" field.
func MessageNotIn(vs ...string) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMessage), v...))
	})
}

// MessageGT applies the GT predicate on the "message" field.
func MessageGT(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMessage), v))
	})
}

// MessageGTE applies the GTE predicate on the "message" field.
func MessageGTE(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMessage), v))
	})
}

// MessageLT applies the LT predicate on the "message" field.
func MessageLT(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMessage), v))
	})
}

// MessageLTE applies the LTE predicate on the "message" field.
func MessageLTE(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMessage), v))
	})
}

// MessageContains applies the Contains predicate on the "message" field.
func MessageContains(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMessage), v))
	})
}

// MessageHasPrefix applies the HasPrefix predicate on the "message" field.
func MessageHasPrefix(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMessage), v))
	})
}

// MessageHasSuffix applies the HasSuffix predicate on the "message" field.
func MessageHasSuffix(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMessage), v))
	})
}

// MessageIsNil applies the IsNil predicate on the "message" field.
func MessageIsNil() predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMessage)))
	})
}

// MessageNotNil applies the NotNil predicate on the "message" field.
func MessageNotNil() predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMessage)))
	})
}

// MessageEqualFold applies the EqualFold predicate on the "message" field.
func MessageEqualFold(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMessage), v))
	})
}

// MessageContainsFold applies the ContainsFold predicate on the "message" field.
func MessageContainsFold(v string) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMessage), v))
	})
}

// RecommendIndexEQ applies the EQ predicate on the "recommend_index" field.
func RecommendIndexEQ(v decimal.Decimal) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRecommendIndex), v))
	})
}

// RecommendIndexNEQ applies the NEQ predicate on the "recommend_index" field.
func RecommendIndexNEQ(v decimal.Decimal) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRecommendIndex), v))
	})
}

// RecommendIndexIn applies the In predicate on the "recommend_index" field.
func RecommendIndexIn(vs ...decimal.Decimal) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRecommendIndex), v...))
	})
}

// RecommendIndexNotIn applies the NotIn predicate on the "recommend_index" field.
func RecommendIndexNotIn(vs ...decimal.Decimal) predicate.Recommend {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRecommendIndex), v...))
	})
}

// RecommendIndexGT applies the GT predicate on the "recommend_index" field.
func RecommendIndexGT(v decimal.Decimal) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRecommendIndex), v))
	})
}

// RecommendIndexGTE applies the GTE predicate on the "recommend_index" field.
func RecommendIndexGTE(v decimal.Decimal) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRecommendIndex), v))
	})
}

// RecommendIndexLT applies the LT predicate on the "recommend_index" field.
func RecommendIndexLT(v decimal.Decimal) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRecommendIndex), v))
	})
}

// RecommendIndexLTE applies the LTE predicate on the "recommend_index" field.
func RecommendIndexLTE(v decimal.Decimal) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRecommendIndex), v))
	})
}

// RecommendIndexIsNil applies the IsNil predicate on the "recommend_index" field.
func RecommendIndexIsNil() predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRecommendIndex)))
	})
}

// RecommendIndexNotNil applies the NotNil predicate on the "recommend_index" field.
func RecommendIndexNotNil() predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRecommendIndex)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Recommend) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Recommend) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Recommend) predicate.Recommend {
	return predicate.Recommend(func(s *sql.Selector) {
		p(s.Not())
	})
}