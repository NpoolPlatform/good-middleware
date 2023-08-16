// Code generated by ent, DO NOT EDIT.

package goodrewardhistory

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// RewardDate applies equality check predicate on the "reward_date" field. It's identical to RewardDateEQ.
func RewardDate(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRewardDate), v))
	})
}

// Tid applies equality check predicate on the "tid" field. It's identical to TidEQ.
func Tid(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTid), v))
	})
}

// Amount applies equality check predicate on the "amount" field. It's identical to AmountEQ.
func Amount(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// UnitAmount applies equality check predicate on the "unit_amount" field. It's identical to UnitAmountEQ.
func UnitAmount(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitAmount), v))
	})
}

// UnitNetAmount applies equality check predicate on the "unit_net_amount" field. It's identical to UnitNetAmountEQ.
func UnitNetAmount(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitNetAmount), v))
	})
}

// Result applies equality check predicate on the "result" field. It's identical to ResultEQ.
func Result(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodID)))
	})
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodID)))
	})
}

// RewardDateEQ applies the EQ predicate on the "reward_date" field.
func RewardDateEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRewardDate), v))
	})
}

// RewardDateNEQ applies the NEQ predicate on the "reward_date" field.
func RewardDateNEQ(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRewardDate), v))
	})
}

// RewardDateIn applies the In predicate on the "reward_date" field.
func RewardDateIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRewardDate), v...))
	})
}

// RewardDateNotIn applies the NotIn predicate on the "reward_date" field.
func RewardDateNotIn(vs ...uint32) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRewardDate), v...))
	})
}

// RewardDateGT applies the GT predicate on the "reward_date" field.
func RewardDateGT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldRewardDate), v))
	})
}

// RewardDateGTE applies the GTE predicate on the "reward_date" field.
func RewardDateGTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldRewardDate), v))
	})
}

// RewardDateLT applies the LT predicate on the "reward_date" field.
func RewardDateLT(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldRewardDate), v))
	})
}

// RewardDateLTE applies the LTE predicate on the "reward_date" field.
func RewardDateLTE(v uint32) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldRewardDate), v))
	})
}

// RewardDateIsNil applies the IsNil predicate on the "reward_date" field.
func RewardDateIsNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldRewardDate)))
	})
}

// RewardDateNotNil applies the NotNil predicate on the "reward_date" field.
func RewardDateNotNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldRewardDate)))
	})
}

// TidEQ applies the EQ predicate on the "tid" field.
func TidEQ(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTid), v))
	})
}

// TidNEQ applies the NEQ predicate on the "tid" field.
func TidNEQ(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTid), v))
	})
}

// TidIn applies the In predicate on the "tid" field.
func TidIn(vs ...uuid.UUID) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTid), v...))
	})
}

// TidNotIn applies the NotIn predicate on the "tid" field.
func TidNotIn(vs ...uuid.UUID) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTid), v...))
	})
}

// TidGT applies the GT predicate on the "tid" field.
func TidGT(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTid), v))
	})
}

// TidGTE applies the GTE predicate on the "tid" field.
func TidGTE(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTid), v))
	})
}

// TidLT applies the LT predicate on the "tid" field.
func TidLT(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTid), v))
	})
}

// TidLTE applies the LTE predicate on the "tid" field.
func TidLTE(v uuid.UUID) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTid), v))
	})
}

// TidIsNil applies the IsNil predicate on the "tid" field.
func TidIsNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTid)))
	})
}

// TidNotNil applies the NotNil predicate on the "tid" field.
func TidNotNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTid)))
	})
}

// AmountEQ applies the EQ predicate on the "amount" field.
func AmountEQ(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmount), v))
	})
}

// AmountNEQ applies the NEQ predicate on the "amount" field.
func AmountNEQ(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmount), v))
	})
}

// AmountIn applies the In predicate on the "amount" field.
func AmountIn(vs ...decimal.Decimal) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAmount), v...))
	})
}

// AmountNotIn applies the NotIn predicate on the "amount" field.
func AmountNotIn(vs ...decimal.Decimal) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAmount), v...))
	})
}

// AmountGT applies the GT predicate on the "amount" field.
func AmountGT(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmount), v))
	})
}

// AmountGTE applies the GTE predicate on the "amount" field.
func AmountGTE(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmount), v))
	})
}

// AmountLT applies the LT predicate on the "amount" field.
func AmountLT(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmount), v))
	})
}

// AmountLTE applies the LTE predicate on the "amount" field.
func AmountLTE(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmount), v))
	})
}

// AmountIsNil applies the IsNil predicate on the "amount" field.
func AmountIsNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAmount)))
	})
}

// AmountNotNil applies the NotNil predicate on the "amount" field.
func AmountNotNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAmount)))
	})
}

// UnitAmountEQ applies the EQ predicate on the "unit_amount" field.
func UnitAmountEQ(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitAmount), v))
	})
}

// UnitAmountNEQ applies the NEQ predicate on the "unit_amount" field.
func UnitAmountNEQ(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnitAmount), v))
	})
}

// UnitAmountIn applies the In predicate on the "unit_amount" field.
func UnitAmountIn(vs ...decimal.Decimal) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnitAmount), v...))
	})
}

// UnitAmountNotIn applies the NotIn predicate on the "unit_amount" field.
func UnitAmountNotIn(vs ...decimal.Decimal) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnitAmount), v...))
	})
}

// UnitAmountGT applies the GT predicate on the "unit_amount" field.
func UnitAmountGT(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnitAmount), v))
	})
}

// UnitAmountGTE applies the GTE predicate on the "unit_amount" field.
func UnitAmountGTE(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnitAmount), v))
	})
}

// UnitAmountLT applies the LT predicate on the "unit_amount" field.
func UnitAmountLT(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnitAmount), v))
	})
}

// UnitAmountLTE applies the LTE predicate on the "unit_amount" field.
func UnitAmountLTE(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnitAmount), v))
	})
}

// UnitAmountIsNil applies the IsNil predicate on the "unit_amount" field.
func UnitAmountIsNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnitAmount)))
	})
}

// UnitAmountNotNil applies the NotNil predicate on the "unit_amount" field.
func UnitAmountNotNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnitAmount)))
	})
}

// UnitNetAmountEQ applies the EQ predicate on the "unit_net_amount" field.
func UnitNetAmountEQ(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitNetAmount), v))
	})
}

// UnitNetAmountNEQ applies the NEQ predicate on the "unit_net_amount" field.
func UnitNetAmountNEQ(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnitNetAmount), v))
	})
}

// UnitNetAmountIn applies the In predicate on the "unit_net_amount" field.
func UnitNetAmountIn(vs ...decimal.Decimal) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnitNetAmount), v...))
	})
}

// UnitNetAmountNotIn applies the NotIn predicate on the "unit_net_amount" field.
func UnitNetAmountNotIn(vs ...decimal.Decimal) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnitNetAmount), v...))
	})
}

// UnitNetAmountGT applies the GT predicate on the "unit_net_amount" field.
func UnitNetAmountGT(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnitNetAmount), v))
	})
}

// UnitNetAmountGTE applies the GTE predicate on the "unit_net_amount" field.
func UnitNetAmountGTE(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnitNetAmount), v))
	})
}

// UnitNetAmountLT applies the LT predicate on the "unit_net_amount" field.
func UnitNetAmountLT(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnitNetAmount), v))
	})
}

// UnitNetAmountLTE applies the LTE predicate on the "unit_net_amount" field.
func UnitNetAmountLTE(v decimal.Decimal) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnitNetAmount), v))
	})
}

// UnitNetAmountIsNil applies the IsNil predicate on the "unit_net_amount" field.
func UnitNetAmountIsNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnitNetAmount)))
	})
}

// UnitNetAmountNotNil applies the NotNil predicate on the "unit_net_amount" field.
func UnitNetAmountNotNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnitNetAmount)))
	})
}

// ResultEQ applies the EQ predicate on the "result" field.
func ResultEQ(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldResult), v))
	})
}

// ResultNEQ applies the NEQ predicate on the "result" field.
func ResultNEQ(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldResult), v))
	})
}

// ResultIn applies the In predicate on the "result" field.
func ResultIn(vs ...string) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldResult), v...))
	})
}

// ResultNotIn applies the NotIn predicate on the "result" field.
func ResultNotIn(vs ...string) predicate.GoodRewardHistory {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldResult), v...))
	})
}

// ResultGT applies the GT predicate on the "result" field.
func ResultGT(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldResult), v))
	})
}

// ResultGTE applies the GTE predicate on the "result" field.
func ResultGTE(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldResult), v))
	})
}

// ResultLT applies the LT predicate on the "result" field.
func ResultLT(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldResult), v))
	})
}

// ResultLTE applies the LTE predicate on the "result" field.
func ResultLTE(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldResult), v))
	})
}

// ResultContains applies the Contains predicate on the "result" field.
func ResultContains(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldResult), v))
	})
}

// ResultHasPrefix applies the HasPrefix predicate on the "result" field.
func ResultHasPrefix(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldResult), v))
	})
}

// ResultHasSuffix applies the HasSuffix predicate on the "result" field.
func ResultHasSuffix(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldResult), v))
	})
}

// ResultIsNil applies the IsNil predicate on the "result" field.
func ResultIsNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldResult)))
	})
}

// ResultNotNil applies the NotNil predicate on the "result" field.
func ResultNotNil() predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldResult)))
	})
}

// ResultEqualFold applies the EqualFold predicate on the "result" field.
func ResultEqualFold(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldResult), v))
	})
}

// ResultContainsFold applies the ContainsFold predicate on the "result" field.
func ResultContainsFold(v string) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldResult), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodRewardHistory) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodRewardHistory) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
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
func Not(p predicate.GoodRewardHistory) predicate.GoodRewardHistory {
	return predicate.GoodRewardHistory(func(s *sql.Selector) {
		p(s.Not())
	})
}
