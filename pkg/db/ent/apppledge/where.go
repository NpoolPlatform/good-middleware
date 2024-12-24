// Code generated by ent, DO NOT EDIT.

package apppledge

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// ServiceStartAt applies equality check predicate on the "service_start_at" field. It's identical to ServiceStartAtEQ.
func ServiceStartAt(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceStartAt), v))
	})
}

// StartMode applies equality check predicate on the "start_mode" field. It's identical to StartModeEQ.
func StartMode(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartMode), v))
	})
}

// EnableSetCommission applies equality check predicate on the "enable_set_commission" field. It's identical to EnableSetCommissionEQ.
func EnableSetCommission(v bool) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableSetCommission), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppGoodID)))
	})
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppGoodID)))
	})
}

// ServiceStartAtEQ applies the EQ predicate on the "service_start_at" field.
func ServiceStartAtEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtNEQ applies the NEQ predicate on the "service_start_at" field.
func ServiceStartAtNEQ(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtIn applies the In predicate on the "service_start_at" field.
func ServiceStartAtIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldServiceStartAt), v...))
	})
}

// ServiceStartAtNotIn applies the NotIn predicate on the "service_start_at" field.
func ServiceStartAtNotIn(vs ...uint32) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldServiceStartAt), v...))
	})
}

// ServiceStartAtGT applies the GT predicate on the "service_start_at" field.
func ServiceStartAtGT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtGTE applies the GTE predicate on the "service_start_at" field.
func ServiceStartAtGTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtLT applies the LT predicate on the "service_start_at" field.
func ServiceStartAtLT(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtLTE applies the LTE predicate on the "service_start_at" field.
func ServiceStartAtLTE(v uint32) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtIsNil applies the IsNil predicate on the "service_start_at" field.
func ServiceStartAtIsNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldServiceStartAt)))
	})
}

// ServiceStartAtNotNil applies the NotNil predicate on the "service_start_at" field.
func ServiceStartAtNotNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldServiceStartAt)))
	})
}

// StartModeEQ applies the EQ predicate on the "start_mode" field.
func StartModeEQ(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartMode), v))
	})
}

// StartModeNEQ applies the NEQ predicate on the "start_mode" field.
func StartModeNEQ(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartMode), v))
	})
}

// StartModeIn applies the In predicate on the "start_mode" field.
func StartModeIn(vs ...string) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartMode), v...))
	})
}

// StartModeNotIn applies the NotIn predicate on the "start_mode" field.
func StartModeNotIn(vs ...string) predicate.AppPledge {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartMode), v...))
	})
}

// StartModeGT applies the GT predicate on the "start_mode" field.
func StartModeGT(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartMode), v))
	})
}

// StartModeGTE applies the GTE predicate on the "start_mode" field.
func StartModeGTE(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartMode), v))
	})
}

// StartModeLT applies the LT predicate on the "start_mode" field.
func StartModeLT(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartMode), v))
	})
}

// StartModeLTE applies the LTE predicate on the "start_mode" field.
func StartModeLTE(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartMode), v))
	})
}

// StartModeContains applies the Contains predicate on the "start_mode" field.
func StartModeContains(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStartMode), v))
	})
}

// StartModeHasPrefix applies the HasPrefix predicate on the "start_mode" field.
func StartModeHasPrefix(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStartMode), v))
	})
}

// StartModeHasSuffix applies the HasSuffix predicate on the "start_mode" field.
func StartModeHasSuffix(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStartMode), v))
	})
}

// StartModeIsNil applies the IsNil predicate on the "start_mode" field.
func StartModeIsNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartMode)))
	})
}

// StartModeNotNil applies the NotNil predicate on the "start_mode" field.
func StartModeNotNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartMode)))
	})
}

// StartModeEqualFold applies the EqualFold predicate on the "start_mode" field.
func StartModeEqualFold(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStartMode), v))
	})
}

// StartModeContainsFold applies the ContainsFold predicate on the "start_mode" field.
func StartModeContainsFold(v string) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStartMode), v))
	})
}

// EnableSetCommissionEQ applies the EQ predicate on the "enable_set_commission" field.
func EnableSetCommissionEQ(v bool) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEnableSetCommission), v))
	})
}

// EnableSetCommissionNEQ applies the NEQ predicate on the "enable_set_commission" field.
func EnableSetCommissionNEQ(v bool) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEnableSetCommission), v))
	})
}

// EnableSetCommissionIsNil applies the IsNil predicate on the "enable_set_commission" field.
func EnableSetCommissionIsNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldEnableSetCommission)))
	})
}

// EnableSetCommissionNotNil applies the NotNil predicate on the "enable_set_commission" field.
func EnableSetCommissionNotNil() predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldEnableSetCommission)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AppPledge) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AppPledge) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
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
func Not(p predicate.AppPledge) predicate.AppPledge {
	return predicate.AppPledge(func(s *sql.Selector) {
		p(s.Not())
	})
}