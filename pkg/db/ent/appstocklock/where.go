// Code generated by ent, DO NOT EDIT.

package appstocklock

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// AppStockID applies equality check predicate on the "app_stock_id" field. It's identical to AppStockIDEQ.
func AppStockID(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppStockID), v))
	})
}

// Units applies equality check predicate on the "units" field. It's identical to UnitsEQ.
func Units(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnits), v))
	})
}

// AppSpotUnits applies equality check predicate on the "app_spot_units" field. It's identical to AppSpotUnitsEQ.
func AppSpotUnits(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppSpotUnits), v))
	})
}

// LockState applies equality check predicate on the "lock_state" field. It's identical to LockStateEQ.
func LockState(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockState), v))
	})
}

// ChargeBackState applies equality check predicate on the "charge_back_state" field. It's identical to ChargeBackStateEQ.
func ChargeBackState(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChargeBackState), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// AppStockIDEQ applies the EQ predicate on the "app_stock_id" field.
func AppStockIDEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppStockID), v))
	})
}

// AppStockIDNEQ applies the NEQ predicate on the "app_stock_id" field.
func AppStockIDNEQ(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppStockID), v))
	})
}

// AppStockIDIn applies the In predicate on the "app_stock_id" field.
func AppStockIDIn(vs ...uuid.UUID) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppStockID), v...))
	})
}

// AppStockIDNotIn applies the NotIn predicate on the "app_stock_id" field.
func AppStockIDNotIn(vs ...uuid.UUID) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppStockID), v...))
	})
}

// AppStockIDGT applies the GT predicate on the "app_stock_id" field.
func AppStockIDGT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppStockID), v))
	})
}

// AppStockIDGTE applies the GTE predicate on the "app_stock_id" field.
func AppStockIDGTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppStockID), v))
	})
}

// AppStockIDLT applies the LT predicate on the "app_stock_id" field.
func AppStockIDLT(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppStockID), v))
	})
}

// AppStockIDLTE applies the LTE predicate on the "app_stock_id" field.
func AppStockIDLTE(v uuid.UUID) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppStockID), v))
	})
}

// AppStockIDIsNil applies the IsNil predicate on the "app_stock_id" field.
func AppStockIDIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppStockID)))
	})
}

// AppStockIDNotNil applies the NotNil predicate on the "app_stock_id" field.
func AppStockIDNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppStockID)))
	})
}

// UnitsEQ applies the EQ predicate on the "units" field.
func UnitsEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUnits), v))
	})
}

// UnitsNEQ applies the NEQ predicate on the "units" field.
func UnitsNEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUnits), v))
	})
}

// UnitsIn applies the In predicate on the "units" field.
func UnitsIn(vs ...decimal.Decimal) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUnits), v...))
	})
}

// UnitsNotIn applies the NotIn predicate on the "units" field.
func UnitsNotIn(vs ...decimal.Decimal) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUnits), v...))
	})
}

// UnitsGT applies the GT predicate on the "units" field.
func UnitsGT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUnits), v))
	})
}

// UnitsGTE applies the GTE predicate on the "units" field.
func UnitsGTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUnits), v))
	})
}

// UnitsLT applies the LT predicate on the "units" field.
func UnitsLT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUnits), v))
	})
}

// UnitsLTE applies the LTE predicate on the "units" field.
func UnitsLTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUnits), v))
	})
}

// UnitsIsNil applies the IsNil predicate on the "units" field.
func UnitsIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUnits)))
	})
}

// UnitsNotNil applies the NotNil predicate on the "units" field.
func UnitsNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUnits)))
	})
}

// AppSpotUnitsEQ applies the EQ predicate on the "app_spot_units" field.
func AppSpotUnitsEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppSpotUnits), v))
	})
}

// AppSpotUnitsNEQ applies the NEQ predicate on the "app_spot_units" field.
func AppSpotUnitsNEQ(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppSpotUnits), v))
	})
}

// AppSpotUnitsIn applies the In predicate on the "app_spot_units" field.
func AppSpotUnitsIn(vs ...decimal.Decimal) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppSpotUnits), v...))
	})
}

// AppSpotUnitsNotIn applies the NotIn predicate on the "app_spot_units" field.
func AppSpotUnitsNotIn(vs ...decimal.Decimal) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppSpotUnits), v...))
	})
}

// AppSpotUnitsGT applies the GT predicate on the "app_spot_units" field.
func AppSpotUnitsGT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppSpotUnits), v))
	})
}

// AppSpotUnitsGTE applies the GTE predicate on the "app_spot_units" field.
func AppSpotUnitsGTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppSpotUnits), v))
	})
}

// AppSpotUnitsLT applies the LT predicate on the "app_spot_units" field.
func AppSpotUnitsLT(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppSpotUnits), v))
	})
}

// AppSpotUnitsLTE applies the LTE predicate on the "app_spot_units" field.
func AppSpotUnitsLTE(v decimal.Decimal) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppSpotUnits), v))
	})
}

// AppSpotUnitsIsNil applies the IsNil predicate on the "app_spot_units" field.
func AppSpotUnitsIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppSpotUnits)))
	})
}

// AppSpotUnitsNotNil applies the NotNil predicate on the "app_spot_units" field.
func AppSpotUnitsNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppSpotUnits)))
	})
}

// LockStateEQ applies the EQ predicate on the "lock_state" field.
func LockStateEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLockState), v))
	})
}

// LockStateNEQ applies the NEQ predicate on the "lock_state" field.
func LockStateNEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLockState), v))
	})
}

// LockStateIn applies the In predicate on the "lock_state" field.
func LockStateIn(vs ...string) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLockState), v...))
	})
}

// LockStateNotIn applies the NotIn predicate on the "lock_state" field.
func LockStateNotIn(vs ...string) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLockState), v...))
	})
}

// LockStateGT applies the GT predicate on the "lock_state" field.
func LockStateGT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLockState), v))
	})
}

// LockStateGTE applies the GTE predicate on the "lock_state" field.
func LockStateGTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLockState), v))
	})
}

// LockStateLT applies the LT predicate on the "lock_state" field.
func LockStateLT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLockState), v))
	})
}

// LockStateLTE applies the LTE predicate on the "lock_state" field.
func LockStateLTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLockState), v))
	})
}

// LockStateContains applies the Contains predicate on the "lock_state" field.
func LockStateContains(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldLockState), v))
	})
}

// LockStateHasPrefix applies the HasPrefix predicate on the "lock_state" field.
func LockStateHasPrefix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldLockState), v))
	})
}

// LockStateHasSuffix applies the HasSuffix predicate on the "lock_state" field.
func LockStateHasSuffix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldLockState), v))
	})
}

// LockStateIsNil applies the IsNil predicate on the "lock_state" field.
func LockStateIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLockState)))
	})
}

// LockStateNotNil applies the NotNil predicate on the "lock_state" field.
func LockStateNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLockState)))
	})
}

// LockStateEqualFold applies the EqualFold predicate on the "lock_state" field.
func LockStateEqualFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldLockState), v))
	})
}

// LockStateContainsFold applies the ContainsFold predicate on the "lock_state" field.
func LockStateContainsFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldLockState), v))
	})
}

// ChargeBackStateEQ applies the EQ predicate on the "charge_back_state" field.
func ChargeBackStateEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateNEQ applies the NEQ predicate on the "charge_back_state" field.
func ChargeBackStateNEQ(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateIn applies the In predicate on the "charge_back_state" field.
func ChargeBackStateIn(vs ...string) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldChargeBackState), v...))
	})
}

// ChargeBackStateNotIn applies the NotIn predicate on the "charge_back_state" field.
func ChargeBackStateNotIn(vs ...string) predicate.AppStockLock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldChargeBackState), v...))
	})
}

// ChargeBackStateGT applies the GT predicate on the "charge_back_state" field.
func ChargeBackStateGT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateGTE applies the GTE predicate on the "charge_back_state" field.
func ChargeBackStateGTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateLT applies the LT predicate on the "charge_back_state" field.
func ChargeBackStateLT(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateLTE applies the LTE predicate on the "charge_back_state" field.
func ChargeBackStateLTE(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateContains applies the Contains predicate on the "charge_back_state" field.
func ChargeBackStateContains(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateHasPrefix applies the HasPrefix predicate on the "charge_back_state" field.
func ChargeBackStateHasPrefix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateHasSuffix applies the HasSuffix predicate on the "charge_back_state" field.
func ChargeBackStateHasSuffix(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateIsNil applies the IsNil predicate on the "charge_back_state" field.
func ChargeBackStateIsNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldChargeBackState)))
	})
}

// ChargeBackStateNotNil applies the NotNil predicate on the "charge_back_state" field.
func ChargeBackStateNotNil() predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldChargeBackState)))
	})
}

// ChargeBackStateEqualFold applies the EqualFold predicate on the "charge_back_state" field.
func ChargeBackStateEqualFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldChargeBackState), v))
	})
}

// ChargeBackStateContainsFold applies the ContainsFold predicate on the "charge_back_state" field.
func ChargeBackStateContainsFold(v string) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldChargeBackState), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppStockLock) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppStockLock) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
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
func Not(p predicate.AppStockLock) predicate.AppStockLock {
	return predicate.AppStockLock(func(s *sql.Selector) {
		p(s.Not())
	})
}
