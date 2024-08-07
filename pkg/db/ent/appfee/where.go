// Code generated by ent, DO NOT EDIT.

package appfee

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// UnitValue applies equality check predicate on the "unit_value" field. It's identical to UnitValueEQ.
func UnitValue(v decimal.Decimal) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitValue), v))
	})
}

// CancelMode applies equality check predicate on the "cancel_mode" field. It's identical to CancelModeEQ.
func CancelMode(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCancelMode), v))
	})
}

// MinOrderDurationSeconds applies equality check predicate on the "min_order_duration_seconds" field. It's identical to MinOrderDurationSecondsEQ.
func MinOrderDurationSeconds(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinOrderDurationSeconds), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppGoodID)))
	})
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppGoodID)))
	})
}

// UnitValueEQ applies the EQ predicate on the "unit_value" field.
func UnitValueEQ(v decimal.Decimal) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnitValue), v))
	})
}

// UnitValueNEQ applies the NEQ predicate on the "unit_value" field.
func UnitValueNEQ(v decimal.Decimal) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnitValue), v))
	})
}

// UnitValueIn applies the In predicate on the "unit_value" field.
func UnitValueIn(vs ...decimal.Decimal) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnitValue), v...))
	})
}

// UnitValueNotIn applies the NotIn predicate on the "unit_value" field.
func UnitValueNotIn(vs ...decimal.Decimal) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnitValue), v...))
	})
}

// UnitValueGT applies the GT predicate on the "unit_value" field.
func UnitValueGT(v decimal.Decimal) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnitValue), v))
	})
}

// UnitValueGTE applies the GTE predicate on the "unit_value" field.
func UnitValueGTE(v decimal.Decimal) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnitValue), v))
	})
}

// UnitValueLT applies the LT predicate on the "unit_value" field.
func UnitValueLT(v decimal.Decimal) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnitValue), v))
	})
}

// UnitValueLTE applies the LTE predicate on the "unit_value" field.
func UnitValueLTE(v decimal.Decimal) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnitValue), v))
	})
}

// UnitValueIsNil applies the IsNil predicate on the "unit_value" field.
func UnitValueIsNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnitValue)))
	})
}

// UnitValueNotNil applies the NotNil predicate on the "unit_value" field.
func UnitValueNotNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnitValue)))
	})
}

// CancelModeEQ applies the EQ predicate on the "cancel_mode" field.
func CancelModeEQ(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCancelMode), v))
	})
}

// CancelModeNEQ applies the NEQ predicate on the "cancel_mode" field.
func CancelModeNEQ(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCancelMode), v))
	})
}

// CancelModeIn applies the In predicate on the "cancel_mode" field.
func CancelModeIn(vs ...string) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCancelMode), v...))
	})
}

// CancelModeNotIn applies the NotIn predicate on the "cancel_mode" field.
func CancelModeNotIn(vs ...string) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCancelMode), v...))
	})
}

// CancelModeGT applies the GT predicate on the "cancel_mode" field.
func CancelModeGT(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCancelMode), v))
	})
}

// CancelModeGTE applies the GTE predicate on the "cancel_mode" field.
func CancelModeGTE(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCancelMode), v))
	})
}

// CancelModeLT applies the LT predicate on the "cancel_mode" field.
func CancelModeLT(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCancelMode), v))
	})
}

// CancelModeLTE applies the LTE predicate on the "cancel_mode" field.
func CancelModeLTE(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCancelMode), v))
	})
}

// CancelModeContains applies the Contains predicate on the "cancel_mode" field.
func CancelModeContains(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCancelMode), v))
	})
}

// CancelModeHasPrefix applies the HasPrefix predicate on the "cancel_mode" field.
func CancelModeHasPrefix(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCancelMode), v))
	})
}

// CancelModeHasSuffix applies the HasSuffix predicate on the "cancel_mode" field.
func CancelModeHasSuffix(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCancelMode), v))
	})
}

// CancelModeIsNil applies the IsNil predicate on the "cancel_mode" field.
func CancelModeIsNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCancelMode)))
	})
}

// CancelModeNotNil applies the NotNil predicate on the "cancel_mode" field.
func CancelModeNotNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCancelMode)))
	})
}

// CancelModeEqualFold applies the EqualFold predicate on the "cancel_mode" field.
func CancelModeEqualFold(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCancelMode), v))
	})
}

// CancelModeContainsFold applies the ContainsFold predicate on the "cancel_mode" field.
func CancelModeContainsFold(v string) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCancelMode), v))
	})
}

// MinOrderDurationSecondsEQ applies the EQ predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMinOrderDurationSeconds), v))
	})
}

// MinOrderDurationSecondsNEQ applies the NEQ predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsNEQ(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMinOrderDurationSeconds), v))
	})
}

// MinOrderDurationSecondsIn applies the In predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMinOrderDurationSeconds), v...))
	})
}

// MinOrderDurationSecondsNotIn applies the NotIn predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsNotIn(vs ...uint32) predicate.AppFee {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMinOrderDurationSeconds), v...))
	})
}

// MinOrderDurationSecondsGT applies the GT predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsGT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMinOrderDurationSeconds), v))
	})
}

// MinOrderDurationSecondsGTE applies the GTE predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsGTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMinOrderDurationSeconds), v))
	})
}

// MinOrderDurationSecondsLT applies the LT predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsLT(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMinOrderDurationSeconds), v))
	})
}

// MinOrderDurationSecondsLTE applies the LTE predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsLTE(v uint32) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMinOrderDurationSeconds), v))
	})
}

// MinOrderDurationSecondsIsNil applies the IsNil predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsIsNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMinOrderDurationSeconds)))
	})
}

// MinOrderDurationSecondsNotNil applies the NotNil predicate on the "min_order_duration_seconds" field.
func MinOrderDurationSecondsNotNil() predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMinOrderDurationSeconds)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppFee) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppFee) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
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
func Not(p predicate.AppFee) predicate.AppFee {
	return predicate.AppFee(func(s *sql.Selector) {
		p(s.Not())
	})
}
