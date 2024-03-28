// Code generated by ent, DO NOT EDIT.

package appsimulatepowerrental

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// OrderUnits applies equality check predicate on the "order_units" field. It's identical to OrderUnitsEQ.
func OrderUnits(v decimal.Decimal) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderUnits), v))
	})
}

// OrderDuration applies equality check predicate on the "order_duration" field. It's identical to OrderDurationEQ.
func OrderDuration(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderDuration), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppGoodID)))
	})
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppGoodID)))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIsNil applies the IsNil predicate on the "coin_type_id" field.
func CoinTypeIDIsNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCoinTypeID)))
	})
}

// CoinTypeIDNotNil applies the NotNil predicate on the "coin_type_id" field.
func CoinTypeIDNotNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCoinTypeID)))
	})
}

// OrderUnitsEQ applies the EQ predicate on the "order_units" field.
func OrderUnitsEQ(v decimal.Decimal) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderUnits), v))
	})
}

// OrderUnitsNEQ applies the NEQ predicate on the "order_units" field.
func OrderUnitsNEQ(v decimal.Decimal) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderUnits), v))
	})
}

// OrderUnitsIn applies the In predicate on the "order_units" field.
func OrderUnitsIn(vs ...decimal.Decimal) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderUnits), v...))
	})
}

// OrderUnitsNotIn applies the NotIn predicate on the "order_units" field.
func OrderUnitsNotIn(vs ...decimal.Decimal) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderUnits), v...))
	})
}

// OrderUnitsGT applies the GT predicate on the "order_units" field.
func OrderUnitsGT(v decimal.Decimal) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderUnits), v))
	})
}

// OrderUnitsGTE applies the GTE predicate on the "order_units" field.
func OrderUnitsGTE(v decimal.Decimal) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderUnits), v))
	})
}

// OrderUnitsLT applies the LT predicate on the "order_units" field.
func OrderUnitsLT(v decimal.Decimal) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderUnits), v))
	})
}

// OrderUnitsLTE applies the LTE predicate on the "order_units" field.
func OrderUnitsLTE(v decimal.Decimal) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderUnits), v))
	})
}

// OrderUnitsIsNil applies the IsNil predicate on the "order_units" field.
func OrderUnitsIsNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderUnits)))
	})
}

// OrderUnitsNotNil applies the NotNil predicate on the "order_units" field.
func OrderUnitsNotNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderUnits)))
	})
}

// OrderDurationEQ applies the EQ predicate on the "order_duration" field.
func OrderDurationEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOrderDuration), v))
	})
}

// OrderDurationNEQ applies the NEQ predicate on the "order_duration" field.
func OrderDurationNEQ(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOrderDuration), v))
	})
}

// OrderDurationIn applies the In predicate on the "order_duration" field.
func OrderDurationIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldOrderDuration), v...))
	})
}

// OrderDurationNotIn applies the NotIn predicate on the "order_duration" field.
func OrderDurationNotIn(vs ...uint32) predicate.AppSimulatePowerRental {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldOrderDuration), v...))
	})
}

// OrderDurationGT applies the GT predicate on the "order_duration" field.
func OrderDurationGT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOrderDuration), v))
	})
}

// OrderDurationGTE applies the GTE predicate on the "order_duration" field.
func OrderDurationGTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOrderDuration), v))
	})
}

// OrderDurationLT applies the LT predicate on the "order_duration" field.
func OrderDurationLT(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOrderDuration), v))
	})
}

// OrderDurationLTE applies the LTE predicate on the "order_duration" field.
func OrderDurationLTE(v uint32) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOrderDuration), v))
	})
}

// OrderDurationIsNil applies the IsNil predicate on the "order_duration" field.
func OrderDurationIsNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOrderDuration)))
	})
}

// OrderDurationNotNil applies the NotNil predicate on the "order_duration" field.
func OrderDurationNotNil() predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOrderDuration)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppSimulatePowerRental) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppSimulatePowerRental) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
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
func Not(p predicate.AppSimulatePowerRental) predicate.AppSimulatePowerRental {
	return predicate.AppSimulatePowerRental(func(s *sql.Selector) {
		p(s.Not())
	})
}