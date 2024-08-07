// Code generated by ent, DO NOT EDIT.

package appmininggoodstock

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppGoodStockID applies equality check predicate on the "app_good_stock_id" field. It's identical to AppGoodStockIDEQ.
func AppGoodStockID(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodStockID), v))
	})
}

// MiningGoodStockID applies equality check predicate on the "mining_good_stock_id" field. It's identical to MiningGoodStockIDEQ.
func MiningGoodStockID(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiningGoodStockID), v))
	})
}

// Reserved applies equality check predicate on the "reserved" field. It's identical to ReservedEQ.
func Reserved(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReserved), v))
	})
}

// SpotQuantity applies equality check predicate on the "spot_quantity" field. It's identical to SpotQuantityEQ.
func SpotQuantity(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpotQuantity), v))
	})
}

// Locked applies equality check predicate on the "locked" field. It's identical to LockedEQ.
func Locked(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// InService applies equality check predicate on the "in_service" field. It's identical to InServiceEQ.
func InService(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInService), v))
	})
}

// WaitStart applies equality check predicate on the "wait_start" field. It's identical to WaitStartEQ.
func WaitStart(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWaitStart), v))
	})
}

// Sold applies equality check predicate on the "sold" field. It's identical to SoldEQ.
func Sold(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSold), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppGoodStockIDEQ applies the EQ predicate on the "app_good_stock_id" field.
func AppGoodStockIDEQ(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodStockID), v))
	})
}

// AppGoodStockIDNEQ applies the NEQ predicate on the "app_good_stock_id" field.
func AppGoodStockIDNEQ(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppGoodStockID), v))
	})
}

// AppGoodStockIDIn applies the In predicate on the "app_good_stock_id" field.
func AppGoodStockIDIn(vs ...uuid.UUID) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppGoodStockID), v...))
	})
}

// AppGoodStockIDNotIn applies the NotIn predicate on the "app_good_stock_id" field.
func AppGoodStockIDNotIn(vs ...uuid.UUID) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppGoodStockID), v...))
	})
}

// AppGoodStockIDGT applies the GT predicate on the "app_good_stock_id" field.
func AppGoodStockIDGT(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppGoodStockID), v))
	})
}

// AppGoodStockIDGTE applies the GTE predicate on the "app_good_stock_id" field.
func AppGoodStockIDGTE(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppGoodStockID), v))
	})
}

// AppGoodStockIDLT applies the LT predicate on the "app_good_stock_id" field.
func AppGoodStockIDLT(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppGoodStockID), v))
	})
}

// AppGoodStockIDLTE applies the LTE predicate on the "app_good_stock_id" field.
func AppGoodStockIDLTE(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppGoodStockID), v))
	})
}

// AppGoodStockIDIsNil applies the IsNil predicate on the "app_good_stock_id" field.
func AppGoodStockIDIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppGoodStockID)))
	})
}

// AppGoodStockIDNotNil applies the NotNil predicate on the "app_good_stock_id" field.
func AppGoodStockIDNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppGoodStockID)))
	})
}

// MiningGoodStockIDEQ applies the EQ predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDEQ(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMiningGoodStockID), v))
	})
}

// MiningGoodStockIDNEQ applies the NEQ predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDNEQ(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMiningGoodStockID), v))
	})
}

// MiningGoodStockIDIn applies the In predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDIn(vs ...uuid.UUID) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMiningGoodStockID), v...))
	})
}

// MiningGoodStockIDNotIn applies the NotIn predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDNotIn(vs ...uuid.UUID) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMiningGoodStockID), v...))
	})
}

// MiningGoodStockIDGT applies the GT predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDGT(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMiningGoodStockID), v))
	})
}

// MiningGoodStockIDGTE applies the GTE predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDGTE(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMiningGoodStockID), v))
	})
}

// MiningGoodStockIDLT applies the LT predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDLT(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMiningGoodStockID), v))
	})
}

// MiningGoodStockIDLTE applies the LTE predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDLTE(v uuid.UUID) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMiningGoodStockID), v))
	})
}

// MiningGoodStockIDIsNil applies the IsNil predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMiningGoodStockID)))
	})
}

// MiningGoodStockIDNotNil applies the NotNil predicate on the "mining_good_stock_id" field.
func MiningGoodStockIDNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMiningGoodStockID)))
	})
}

// ReservedEQ applies the EQ predicate on the "reserved" field.
func ReservedEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReserved), v))
	})
}

// ReservedNEQ applies the NEQ predicate on the "reserved" field.
func ReservedNEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReserved), v))
	})
}

// ReservedIn applies the In predicate on the "reserved" field.
func ReservedIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldReserved), v...))
	})
}

// ReservedNotIn applies the NotIn predicate on the "reserved" field.
func ReservedNotIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldReserved), v...))
	})
}

// ReservedGT applies the GT predicate on the "reserved" field.
func ReservedGT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReserved), v))
	})
}

// ReservedGTE applies the GTE predicate on the "reserved" field.
func ReservedGTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReserved), v))
	})
}

// ReservedLT applies the LT predicate on the "reserved" field.
func ReservedLT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReserved), v))
	})
}

// ReservedLTE applies the LTE predicate on the "reserved" field.
func ReservedLTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReserved), v))
	})
}

// ReservedIsNil applies the IsNil predicate on the "reserved" field.
func ReservedIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldReserved)))
	})
}

// ReservedNotNil applies the NotNil predicate on the "reserved" field.
func ReservedNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldReserved)))
	})
}

// SpotQuantityEQ applies the EQ predicate on the "spot_quantity" field.
func SpotQuantityEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSpotQuantity), v))
	})
}

// SpotQuantityNEQ applies the NEQ predicate on the "spot_quantity" field.
func SpotQuantityNEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSpotQuantity), v))
	})
}

// SpotQuantityIn applies the In predicate on the "spot_quantity" field.
func SpotQuantityIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSpotQuantity), v...))
	})
}

// SpotQuantityNotIn applies the NotIn predicate on the "spot_quantity" field.
func SpotQuantityNotIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSpotQuantity), v...))
	})
}

// SpotQuantityGT applies the GT predicate on the "spot_quantity" field.
func SpotQuantityGT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSpotQuantity), v))
	})
}

// SpotQuantityGTE applies the GTE predicate on the "spot_quantity" field.
func SpotQuantityGTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSpotQuantity), v))
	})
}

// SpotQuantityLT applies the LT predicate on the "spot_quantity" field.
func SpotQuantityLT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSpotQuantity), v))
	})
}

// SpotQuantityLTE applies the LTE predicate on the "spot_quantity" field.
func SpotQuantityLTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSpotQuantity), v))
	})
}

// SpotQuantityIsNil applies the IsNil predicate on the "spot_quantity" field.
func SpotQuantityIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSpotQuantity)))
	})
}

// SpotQuantityNotNil applies the NotNil predicate on the "spot_quantity" field.
func SpotQuantityNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSpotQuantity)))
	})
}

// LockedEQ applies the EQ predicate on the "locked" field.
func LockedEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldLocked), v))
	})
}

// LockedNEQ applies the NEQ predicate on the "locked" field.
func LockedNEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldLocked), v))
	})
}

// LockedIn applies the In predicate on the "locked" field.
func LockedIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldLocked), v...))
	})
}

// LockedNotIn applies the NotIn predicate on the "locked" field.
func LockedNotIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldLocked), v...))
	})
}

// LockedGT applies the GT predicate on the "locked" field.
func LockedGT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldLocked), v))
	})
}

// LockedGTE applies the GTE predicate on the "locked" field.
func LockedGTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldLocked), v))
	})
}

// LockedLT applies the LT predicate on the "locked" field.
func LockedLT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldLocked), v))
	})
}

// LockedLTE applies the LTE predicate on the "locked" field.
func LockedLTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldLocked), v))
	})
}

// LockedIsNil applies the IsNil predicate on the "locked" field.
func LockedIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldLocked)))
	})
}

// LockedNotNil applies the NotNil predicate on the "locked" field.
func LockedNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldLocked)))
	})
}

// InServiceEQ applies the EQ predicate on the "in_service" field.
func InServiceEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInService), v))
	})
}

// InServiceNEQ applies the NEQ predicate on the "in_service" field.
func InServiceNEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInService), v))
	})
}

// InServiceIn applies the In predicate on the "in_service" field.
func InServiceIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInService), v...))
	})
}

// InServiceNotIn applies the NotIn predicate on the "in_service" field.
func InServiceNotIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInService), v...))
	})
}

// InServiceGT applies the GT predicate on the "in_service" field.
func InServiceGT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInService), v))
	})
}

// InServiceGTE applies the GTE predicate on the "in_service" field.
func InServiceGTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInService), v))
	})
}

// InServiceLT applies the LT predicate on the "in_service" field.
func InServiceLT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInService), v))
	})
}

// InServiceLTE applies the LTE predicate on the "in_service" field.
func InServiceLTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInService), v))
	})
}

// InServiceIsNil applies the IsNil predicate on the "in_service" field.
func InServiceIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInService)))
	})
}

// InServiceNotNil applies the NotNil predicate on the "in_service" field.
func InServiceNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInService)))
	})
}

// WaitStartEQ applies the EQ predicate on the "wait_start" field.
func WaitStartEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWaitStart), v))
	})
}

// WaitStartNEQ applies the NEQ predicate on the "wait_start" field.
func WaitStartNEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWaitStart), v))
	})
}

// WaitStartIn applies the In predicate on the "wait_start" field.
func WaitStartIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWaitStart), v...))
	})
}

// WaitStartNotIn applies the NotIn predicate on the "wait_start" field.
func WaitStartNotIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWaitStart), v...))
	})
}

// WaitStartGT applies the GT predicate on the "wait_start" field.
func WaitStartGT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWaitStart), v))
	})
}

// WaitStartGTE applies the GTE predicate on the "wait_start" field.
func WaitStartGTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWaitStart), v))
	})
}

// WaitStartLT applies the LT predicate on the "wait_start" field.
func WaitStartLT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWaitStart), v))
	})
}

// WaitStartLTE applies the LTE predicate on the "wait_start" field.
func WaitStartLTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWaitStart), v))
	})
}

// WaitStartIsNil applies the IsNil predicate on the "wait_start" field.
func WaitStartIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldWaitStart)))
	})
}

// WaitStartNotNil applies the NotNil predicate on the "wait_start" field.
func WaitStartNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldWaitStart)))
	})
}

// SoldEQ applies the EQ predicate on the "sold" field.
func SoldEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSold), v))
	})
}

// SoldNEQ applies the NEQ predicate on the "sold" field.
func SoldNEQ(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSold), v))
	})
}

// SoldIn applies the In predicate on the "sold" field.
func SoldIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldSold), v...))
	})
}

// SoldNotIn applies the NotIn predicate on the "sold" field.
func SoldNotIn(vs ...decimal.Decimal) predicate.AppMiningGoodStock {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldSold), v...))
	})
}

// SoldGT applies the GT predicate on the "sold" field.
func SoldGT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSold), v))
	})
}

// SoldGTE applies the GTE predicate on the "sold" field.
func SoldGTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSold), v))
	})
}

// SoldLT applies the LT predicate on the "sold" field.
func SoldLT(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSold), v))
	})
}

// SoldLTE applies the LTE predicate on the "sold" field.
func SoldLTE(v decimal.Decimal) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSold), v))
	})
}

// SoldIsNil applies the IsNil predicate on the "sold" field.
func SoldIsNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldSold)))
	})
}

// SoldNotNil applies the NotNil predicate on the "sold" field.
func SoldNotNil() predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldSold)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppMiningGoodStock) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppMiningGoodStock) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
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
func Not(p predicate.AppMiningGoodStock) predicate.AppMiningGoodStock {
	return predicate.AppMiningGoodStock(func(s *sql.Selector) {
		p(s.Not())
	})
}
