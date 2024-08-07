// Code generated by ent, DO NOT EDIT.

package delegatedstaking

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// GoodID applies equality check predicate on the "good_id" field. It's identical to GoodIDEQ.
func GoodID(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// NoStakeRedeemDelayHours applies equality check predicate on the "no_stake_redeem_delay_hours" field. It's identical to NoStakeRedeemDelayHoursEQ.
func NoStakeRedeemDelayHours(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNoStakeRedeemDelayHours), v))
	})
}

// MaxRedeemDelayHours applies equality check predicate on the "max_redeem_delay_hours" field. It's identical to MaxRedeemDelayHoursEQ.
func MaxRedeemDelayHours(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxRedeemDelayHours), v))
	})
}

// ContractAddress applies equality check predicate on the "contract_address" field. It's identical to ContractAddressEQ.
func ContractAddress(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractAddress), v))
	})
}

// NoStakeBenefitDelayHours applies equality check predicate on the "no_stake_benefit_delay_hours" field. It's identical to NoStakeBenefitDelayHoursEQ.
func NoStakeBenefitDelayHours(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNoStakeBenefitDelayHours), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// GoodIDEQ applies the EQ predicate on the "good_id" field.
func GoodIDEQ(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGoodID), v))
	})
}

// GoodIDNEQ applies the NEQ predicate on the "good_id" field.
func GoodIDNEQ(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGoodID), v))
	})
}

// GoodIDIn applies the In predicate on the "good_id" field.
func GoodIDIn(vs ...uuid.UUID) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGoodID), v...))
	})
}

// GoodIDNotIn applies the NotIn predicate on the "good_id" field.
func GoodIDNotIn(vs ...uuid.UUID) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGoodID), v...))
	})
}

// GoodIDGT applies the GT predicate on the "good_id" field.
func GoodIDGT(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGoodID), v))
	})
}

// GoodIDGTE applies the GTE predicate on the "good_id" field.
func GoodIDGTE(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGoodID), v))
	})
}

// GoodIDLT applies the LT predicate on the "good_id" field.
func GoodIDLT(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGoodID), v))
	})
}

// GoodIDLTE applies the LTE predicate on the "good_id" field.
func GoodIDLTE(v uuid.UUID) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGoodID), v))
	})
}

// GoodIDIsNil applies the IsNil predicate on the "good_id" field.
func GoodIDIsNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGoodID)))
	})
}

// GoodIDNotNil applies the NotNil predicate on the "good_id" field.
func GoodIDNotNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGoodID)))
	})
}

// NoStakeRedeemDelayHoursEQ applies the EQ predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNoStakeRedeemDelayHours), v))
	})
}

// NoStakeRedeemDelayHoursNEQ applies the NEQ predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursNEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNoStakeRedeemDelayHours), v))
	})
}

// NoStakeRedeemDelayHoursIn applies the In predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNoStakeRedeemDelayHours), v...))
	})
}

// NoStakeRedeemDelayHoursNotIn applies the NotIn predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursNotIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNoStakeRedeemDelayHours), v...))
	})
}

// NoStakeRedeemDelayHoursGT applies the GT predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursGT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNoStakeRedeemDelayHours), v))
	})
}

// NoStakeRedeemDelayHoursGTE applies the GTE predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursGTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNoStakeRedeemDelayHours), v))
	})
}

// NoStakeRedeemDelayHoursLT applies the LT predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursLT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNoStakeRedeemDelayHours), v))
	})
}

// NoStakeRedeemDelayHoursLTE applies the LTE predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursLTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNoStakeRedeemDelayHours), v))
	})
}

// NoStakeRedeemDelayHoursIsNil applies the IsNil predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursIsNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNoStakeRedeemDelayHours)))
	})
}

// NoStakeRedeemDelayHoursNotNil applies the NotNil predicate on the "no_stake_redeem_delay_hours" field.
func NoStakeRedeemDelayHoursNotNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNoStakeRedeemDelayHours)))
	})
}

// MaxRedeemDelayHoursEQ applies the EQ predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMaxRedeemDelayHours), v))
	})
}

// MaxRedeemDelayHoursNEQ applies the NEQ predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursNEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMaxRedeemDelayHours), v))
	})
}

// MaxRedeemDelayHoursIn applies the In predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMaxRedeemDelayHours), v...))
	})
}

// MaxRedeemDelayHoursNotIn applies the NotIn predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursNotIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMaxRedeemDelayHours), v...))
	})
}

// MaxRedeemDelayHoursGT applies the GT predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursGT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMaxRedeemDelayHours), v))
	})
}

// MaxRedeemDelayHoursGTE applies the GTE predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursGTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMaxRedeemDelayHours), v))
	})
}

// MaxRedeemDelayHoursLT applies the LT predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursLT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMaxRedeemDelayHours), v))
	})
}

// MaxRedeemDelayHoursLTE applies the LTE predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursLTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMaxRedeemDelayHours), v))
	})
}

// MaxRedeemDelayHoursIsNil applies the IsNil predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursIsNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldMaxRedeemDelayHours)))
	})
}

// MaxRedeemDelayHoursNotNil applies the NotNil predicate on the "max_redeem_delay_hours" field.
func MaxRedeemDelayHoursNotNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldMaxRedeemDelayHours)))
	})
}

// ContractAddressEQ applies the EQ predicate on the "contract_address" field.
func ContractAddressEQ(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContractAddress), v))
	})
}

// ContractAddressNEQ applies the NEQ predicate on the "contract_address" field.
func ContractAddressNEQ(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContractAddress), v))
	})
}

// ContractAddressIn applies the In predicate on the "contract_address" field.
func ContractAddressIn(vs ...string) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContractAddress), v...))
	})
}

// ContractAddressNotIn applies the NotIn predicate on the "contract_address" field.
func ContractAddressNotIn(vs ...string) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContractAddress), v...))
	})
}

// ContractAddressGT applies the GT predicate on the "contract_address" field.
func ContractAddressGT(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContractAddress), v))
	})
}

// ContractAddressGTE applies the GTE predicate on the "contract_address" field.
func ContractAddressGTE(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContractAddress), v))
	})
}

// ContractAddressLT applies the LT predicate on the "contract_address" field.
func ContractAddressLT(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContractAddress), v))
	})
}

// ContractAddressLTE applies the LTE predicate on the "contract_address" field.
func ContractAddressLTE(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContractAddress), v))
	})
}

// ContractAddressContains applies the Contains predicate on the "contract_address" field.
func ContractAddressContains(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContractAddress), v))
	})
}

// ContractAddressHasPrefix applies the HasPrefix predicate on the "contract_address" field.
func ContractAddressHasPrefix(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContractAddress), v))
	})
}

// ContractAddressHasSuffix applies the HasSuffix predicate on the "contract_address" field.
func ContractAddressHasSuffix(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContractAddress), v))
	})
}

// ContractAddressIsNil applies the IsNil predicate on the "contract_address" field.
func ContractAddressIsNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldContractAddress)))
	})
}

// ContractAddressNotNil applies the NotNil predicate on the "contract_address" field.
func ContractAddressNotNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldContractAddress)))
	})
}

// ContractAddressEqualFold applies the EqualFold predicate on the "contract_address" field.
func ContractAddressEqualFold(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContractAddress), v))
	})
}

// ContractAddressContainsFold applies the ContainsFold predicate on the "contract_address" field.
func ContractAddressContainsFold(v string) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContractAddress), v))
	})
}

// NoStakeBenefitDelayHoursEQ applies the EQ predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNoStakeBenefitDelayHours), v))
	})
}

// NoStakeBenefitDelayHoursNEQ applies the NEQ predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursNEQ(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNoStakeBenefitDelayHours), v))
	})
}

// NoStakeBenefitDelayHoursIn applies the In predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNoStakeBenefitDelayHours), v...))
	})
}

// NoStakeBenefitDelayHoursNotIn applies the NotIn predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursNotIn(vs ...uint32) predicate.DelegatedStaking {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNoStakeBenefitDelayHours), v...))
	})
}

// NoStakeBenefitDelayHoursGT applies the GT predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursGT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNoStakeBenefitDelayHours), v))
	})
}

// NoStakeBenefitDelayHoursGTE applies the GTE predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursGTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNoStakeBenefitDelayHours), v))
	})
}

// NoStakeBenefitDelayHoursLT applies the LT predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursLT(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNoStakeBenefitDelayHours), v))
	})
}

// NoStakeBenefitDelayHoursLTE applies the LTE predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursLTE(v uint32) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNoStakeBenefitDelayHours), v))
	})
}

// NoStakeBenefitDelayHoursIsNil applies the IsNil predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursIsNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldNoStakeBenefitDelayHours)))
	})
}

// NoStakeBenefitDelayHoursNotNil applies the NotNil predicate on the "no_stake_benefit_delay_hours" field.
func NoStakeBenefitDelayHoursNotNil() predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldNoStakeBenefitDelayHours)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DelegatedStaking) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DelegatedStaking) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
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
func Not(p predicate.DelegatedStaking) predicate.DelegatedStaking {
	return predicate.DelegatedStaking(func(s *sql.Selector) {
		p(s.Not())
	})
}
