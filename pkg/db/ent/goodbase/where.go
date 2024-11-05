// Code generated by ent, DO NOT EDIT.

package goodbase

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// GoodType applies equality check predicate on the "good_type" field. It's identical to GoodTypeEQ.
func GoodType(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodType), v))
	})
}

// BenefitType applies equality check predicate on the "benefit_type" field. It's identical to BenefitTypeEQ.
func BenefitType(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitType), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// ServiceStartAt applies equality check predicate on the "service_start_at" field. It's identical to ServiceStartAtEQ.
func ServiceStartAt(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceStartAt), v))
	})
}

// StartMode applies equality check predicate on the "start_mode" field. It's identical to StartModeEQ.
func StartMode(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartMode), v))
	})
}

// TestOnly applies equality check predicate on the "test_only" field. It's identical to TestOnlyEQ.
func TestOnly(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestOnly), v))
	})
}

// BenefitIntervalHours applies equality check predicate on the "benefit_interval_hours" field. It's identical to BenefitIntervalHoursEQ.
func BenefitIntervalHours(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitIntervalHours), v))
	})
}

// Purchasable applies equality check predicate on the "purchasable" field. It's identical to PurchasableEQ.
func Purchasable(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPurchasable), v))
	})
}

// Online applies equality check predicate on the "online" field. It's identical to OnlineEQ.
func Online(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnline), v))
	})
}

// State applies equality check predicate on the "state" field. It's identical to StateEQ.
func State(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// GoodTypeEQ applies the EQ predicate on the "good_type" field.
func GoodTypeEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodType), v))
	})
}

// GoodTypeNEQ applies the NEQ predicate on the "good_type" field.
func GoodTypeNEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodType), v))
	})
}

// GoodTypeIn applies the In predicate on the "good_type" field.
func GoodTypeIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodType), v...))
	})
}

// GoodTypeNotIn applies the NotIn predicate on the "good_type" field.
func GoodTypeNotIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodType), v...))
	})
}

// GoodTypeGT applies the GT predicate on the "good_type" field.
func GoodTypeGT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodType), v))
	})
}

// GoodTypeGTE applies the GTE predicate on the "good_type" field.
func GoodTypeGTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodType), v))
	})
}

// GoodTypeLT applies the LT predicate on the "good_type" field.
func GoodTypeLT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodType), v))
	})
}

// GoodTypeLTE applies the LTE predicate on the "good_type" field.
func GoodTypeLTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodType), v))
	})
}

// GoodTypeContains applies the Contains predicate on the "good_type" field.
func GoodTypeContains(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGoodType), v))
	})
}

// GoodTypeHasPrefix applies the HasPrefix predicate on the "good_type" field.
func GoodTypeHasPrefix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGoodType), v))
	})
}

// GoodTypeHasSuffix applies the HasSuffix predicate on the "good_type" field.
func GoodTypeHasSuffix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGoodType), v))
	})
}

// GoodTypeIsNil applies the IsNil predicate on the "good_type" field.
func GoodTypeIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodType)))
	})
}

// GoodTypeNotNil applies the NotNil predicate on the "good_type" field.
func GoodTypeNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodType)))
	})
}

// GoodTypeEqualFold applies the EqualFold predicate on the "good_type" field.
func GoodTypeEqualFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGoodType), v))
	})
}

// GoodTypeContainsFold applies the ContainsFold predicate on the "good_type" field.
func GoodTypeContainsFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGoodType), v))
	})
}

// BenefitTypeEQ applies the EQ predicate on the "benefit_type" field.
func BenefitTypeEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeNEQ applies the NEQ predicate on the "benefit_type" field.
func BenefitTypeNEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeIn applies the In predicate on the "benefit_type" field.
func BenefitTypeIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBenefitType), v...))
	})
}

// BenefitTypeNotIn applies the NotIn predicate on the "benefit_type" field.
func BenefitTypeNotIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBenefitType), v...))
	})
}

// BenefitTypeGT applies the GT predicate on the "benefit_type" field.
func BenefitTypeGT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeGTE applies the GTE predicate on the "benefit_type" field.
func BenefitTypeGTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeLT applies the LT predicate on the "benefit_type" field.
func BenefitTypeLT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeLTE applies the LTE predicate on the "benefit_type" field.
func BenefitTypeLTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeContains applies the Contains predicate on the "benefit_type" field.
func BenefitTypeContains(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeHasPrefix applies the HasPrefix predicate on the "benefit_type" field.
func BenefitTypeHasPrefix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeHasSuffix applies the HasSuffix predicate on the "benefit_type" field.
func BenefitTypeHasSuffix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeIsNil applies the IsNil predicate on the "benefit_type" field.
func BenefitTypeIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBenefitType)))
	})
}

// BenefitTypeNotNil applies the NotNil predicate on the "benefit_type" field.
func BenefitTypeNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBenefitType)))
	})
}

// BenefitTypeEqualFold applies the EqualFold predicate on the "benefit_type" field.
func BenefitTypeEqualFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBenefitType), v))
	})
}

// BenefitTypeContainsFold applies the ContainsFold predicate on the "benefit_type" field.
func BenefitTypeContainsFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBenefitType), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// ServiceStartAtEQ applies the EQ predicate on the "service_start_at" field.
func ServiceStartAtEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtNEQ applies the NEQ predicate on the "service_start_at" field.
func ServiceStartAtNEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtIn applies the In predicate on the "service_start_at" field.
func ServiceStartAtIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldServiceStartAt), v...))
	})
}

// ServiceStartAtNotIn applies the NotIn predicate on the "service_start_at" field.
func ServiceStartAtNotIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldServiceStartAt), v...))
	})
}

// ServiceStartAtGT applies the GT predicate on the "service_start_at" field.
func ServiceStartAtGT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtGTE applies the GTE predicate on the "service_start_at" field.
func ServiceStartAtGTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtLT applies the LT predicate on the "service_start_at" field.
func ServiceStartAtLT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtLTE applies the LTE predicate on the "service_start_at" field.
func ServiceStartAtLTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldServiceStartAt), v))
	})
}

// ServiceStartAtIsNil applies the IsNil predicate on the "service_start_at" field.
func ServiceStartAtIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldServiceStartAt)))
	})
}

// ServiceStartAtNotNil applies the NotNil predicate on the "service_start_at" field.
func ServiceStartAtNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldServiceStartAt)))
	})
}

// StartModeEQ applies the EQ predicate on the "start_mode" field.
func StartModeEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartMode), v))
	})
}

// StartModeNEQ applies the NEQ predicate on the "start_mode" field.
func StartModeNEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartMode), v))
	})
}

// StartModeIn applies the In predicate on the "start_mode" field.
func StartModeIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldStartMode), v...))
	})
}

// StartModeNotIn applies the NotIn predicate on the "start_mode" field.
func StartModeNotIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldStartMode), v...))
	})
}

// StartModeGT applies the GT predicate on the "start_mode" field.
func StartModeGT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartMode), v))
	})
}

// StartModeGTE applies the GTE predicate on the "start_mode" field.
func StartModeGTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartMode), v))
	})
}

// StartModeLT applies the LT predicate on the "start_mode" field.
func StartModeLT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartMode), v))
	})
}

// StartModeLTE applies the LTE predicate on the "start_mode" field.
func StartModeLTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartMode), v))
	})
}

// StartModeContains applies the Contains predicate on the "start_mode" field.
func StartModeContains(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStartMode), v))
	})
}

// StartModeHasPrefix applies the HasPrefix predicate on the "start_mode" field.
func StartModeHasPrefix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStartMode), v))
	})
}

// StartModeHasSuffix applies the HasSuffix predicate on the "start_mode" field.
func StartModeHasSuffix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStartMode), v))
	})
}

// StartModeIsNil applies the IsNil predicate on the "start_mode" field.
func StartModeIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldStartMode)))
	})
}

// StartModeNotNil applies the NotNil predicate on the "start_mode" field.
func StartModeNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldStartMode)))
	})
}

// StartModeEqualFold applies the EqualFold predicate on the "start_mode" field.
func StartModeEqualFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStartMode), v))
	})
}

// StartModeContainsFold applies the ContainsFold predicate on the "start_mode" field.
func StartModeContainsFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStartMode), v))
	})
}

// TestOnlyEQ applies the EQ predicate on the "test_only" field.
func TestOnlyEQ(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTestOnly), v))
	})
}

// TestOnlyNEQ applies the NEQ predicate on the "test_only" field.
func TestOnlyNEQ(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTestOnly), v))
	})
}

// TestOnlyIsNil applies the IsNil predicate on the "test_only" field.
func TestOnlyIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldTestOnly)))
	})
}

// TestOnlyNotNil applies the NotNil predicate on the "test_only" field.
func TestOnlyNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldTestOnly)))
	})
}

// BenefitIntervalHoursEQ applies the EQ predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBenefitIntervalHours), v))
	})
}

// BenefitIntervalHoursNEQ applies the NEQ predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursNEQ(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBenefitIntervalHours), v))
	})
}

// BenefitIntervalHoursIn applies the In predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBenefitIntervalHours), v...))
	})
}

// BenefitIntervalHoursNotIn applies the NotIn predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursNotIn(vs ...uint32) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBenefitIntervalHours), v...))
	})
}

// BenefitIntervalHoursGT applies the GT predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursGT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBenefitIntervalHours), v))
	})
}

// BenefitIntervalHoursGTE applies the GTE predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursGTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBenefitIntervalHours), v))
	})
}

// BenefitIntervalHoursLT applies the LT predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursLT(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBenefitIntervalHours), v))
	})
}

// BenefitIntervalHoursLTE applies the LTE predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursLTE(v uint32) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBenefitIntervalHours), v))
	})
}

// BenefitIntervalHoursIsNil applies the IsNil predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldBenefitIntervalHours)))
	})
}

// BenefitIntervalHoursNotNil applies the NotNil predicate on the "benefit_interval_hours" field.
func BenefitIntervalHoursNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldBenefitIntervalHours)))
	})
}

// PurchasableEQ applies the EQ predicate on the "purchasable" field.
func PurchasableEQ(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPurchasable), v))
	})
}

// PurchasableNEQ applies the NEQ predicate on the "purchasable" field.
func PurchasableNEQ(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPurchasable), v))
	})
}

// PurchasableIsNil applies the IsNil predicate on the "purchasable" field.
func PurchasableIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPurchasable)))
	})
}

// PurchasableNotNil applies the NotNil predicate on the "purchasable" field.
func PurchasableNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPurchasable)))
	})
}

// OnlineEQ applies the EQ predicate on the "online" field.
func OnlineEQ(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnline), v))
	})
}

// OnlineNEQ applies the NEQ predicate on the "online" field.
func OnlineNEQ(v bool) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOnline), v))
	})
}

// OnlineIsNil applies the IsNil predicate on the "online" field.
func OnlineIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldOnline)))
	})
}

// OnlineNotNil applies the NotNil predicate on the "online" field.
func OnlineNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldOnline)))
	})
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldState), v))
	})
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldState), v))
	})
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldState), v...))
	})
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...string) predicate.GoodBase {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldState), v...))
	})
}

// StateGT applies the GT predicate on the "state" field.
func StateGT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldState), v))
	})
}

// StateGTE applies the GTE predicate on the "state" field.
func StateGTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldState), v))
	})
}

// StateLT applies the LT predicate on the "state" field.
func StateLT(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldState), v))
	})
}

// StateLTE applies the LTE predicate on the "state" field.
func StateLTE(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldState), v))
	})
}

// StateContains applies the Contains predicate on the "state" field.
func StateContains(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldState), v))
	})
}

// StateHasPrefix applies the HasPrefix predicate on the "state" field.
func StateHasPrefix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldState), v))
	})
}

// StateHasSuffix applies the HasSuffix predicate on the "state" field.
func StateHasSuffix(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldState), v))
	})
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldState)))
	})
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldState)))
	})
}

// StateEqualFold applies the EqualFold predicate on the "state" field.
func StateEqualFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldState), v))
	})
}

// StateContainsFold applies the ContainsFold predicate on the "state" field.
func StateContainsFold(v string) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldState), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GoodBase) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GoodBase) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
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
func Not(p predicate.GoodBase) predicate.GoodBase {
	return predicate.GoodBase(func(s *sql.Selector) {
		p(s.Not())
	})
}
