// Code generated by ent, DO NOT EDIT.

package score

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ID filters vertices based on their ID field.
func ID(id uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// EntID applies equality check predicate on the "ent_id" field. It's identical to EntIDEQ.
func EntID(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// AppGoodID applies equality check predicate on the "app_good_id" field. It's identical to AppGoodIDEQ.
func AppGoodID(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// Score applies equality check predicate on the "score" field. It's identical to ScoreEQ.
func Score(v decimal.Decimal) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// CommentID applies equality check predicate on the "comment_id" field. It's identical to CommentIDEQ.
func CommentID(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommentID), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// EntIDEQ applies the EQ predicate on the "ent_id" field.
func EntIDEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEntID), v))
	})
}

// EntIDNEQ applies the NEQ predicate on the "ent_id" field.
func EntIDNEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEntID), v))
	})
}

// EntIDIn applies the In predicate on the "ent_id" field.
func EntIDIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEntID), v...))
	})
}

// EntIDNotIn applies the NotIn predicate on the "ent_id" field.
func EntIDNotIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEntID), v...))
	})
}

// EntIDGT applies the GT predicate on the "ent_id" field.
func EntIDGT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEntID), v))
	})
}

// EntIDGTE applies the GTE predicate on the "ent_id" field.
func EntIDGTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEntID), v))
	})
}

// EntIDLT applies the LT predicate on the "ent_id" field.
func EntIDLT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEntID), v))
	})
}

// EntIDLTE applies the LTE predicate on the "ent_id" field.
func EntIDLTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEntID), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// AppGoodIDEQ applies the EQ predicate on the "app_good_id" field.
func AppGoodIDEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDNEQ applies the NEQ predicate on the "app_good_id" field.
func AppGoodIDNEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIn applies the In predicate on the "app_good_id" field.
func AppGoodIDIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDNotIn applies the NotIn predicate on the "app_good_id" field.
func AppGoodIDNotIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldAppGoodID), v...))
	})
}

// AppGoodIDGT applies the GT predicate on the "app_good_id" field.
func AppGoodIDGT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDGTE applies the GTE predicate on the "app_good_id" field.
func AppGoodIDGTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLT applies the LT predicate on the "app_good_id" field.
func AppGoodIDLT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDLTE applies the LTE predicate on the "app_good_id" field.
func AppGoodIDLTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAppGoodID), v))
	})
}

// AppGoodIDIsNil applies the IsNil predicate on the "app_good_id" field.
func AppGoodIDIsNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAppGoodID)))
	})
}

// AppGoodIDNotNil applies the NotNil predicate on the "app_good_id" field.
func AppGoodIDNotNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAppGoodID)))
	})
}

// ScoreEQ applies the EQ predicate on the "score" field.
func ScoreEQ(v decimal.Decimal) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldScore), v))
	})
}

// ScoreNEQ applies the NEQ predicate on the "score" field.
func ScoreNEQ(v decimal.Decimal) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldScore), v))
	})
}

// ScoreIn applies the In predicate on the "score" field.
func ScoreIn(vs ...decimal.Decimal) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldScore), v...))
	})
}

// ScoreNotIn applies the NotIn predicate on the "score" field.
func ScoreNotIn(vs ...decimal.Decimal) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldScore), v...))
	})
}

// ScoreGT applies the GT predicate on the "score" field.
func ScoreGT(v decimal.Decimal) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldScore), v))
	})
}

// ScoreGTE applies the GTE predicate on the "score" field.
func ScoreGTE(v decimal.Decimal) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldScore), v))
	})
}

// ScoreLT applies the LT predicate on the "score" field.
func ScoreLT(v decimal.Decimal) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldScore), v))
	})
}

// ScoreLTE applies the LTE predicate on the "score" field.
func ScoreLTE(v decimal.Decimal) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldScore), v))
	})
}

// ScoreIsNil applies the IsNil predicate on the "score" field.
func ScoreIsNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldScore)))
	})
}

// ScoreNotNil applies the NotNil predicate on the "score" field.
func ScoreNotNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldScore)))
	})
}

// CommentIDEQ applies the EQ predicate on the "comment_id" field.
func CommentIDEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCommentID), v))
	})
}

// CommentIDNEQ applies the NEQ predicate on the "comment_id" field.
func CommentIDNEQ(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCommentID), v))
	})
}

// CommentIDIn applies the In predicate on the "comment_id" field.
func CommentIDIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCommentID), v...))
	})
}

// CommentIDNotIn applies the NotIn predicate on the "comment_id" field.
func CommentIDNotIn(vs ...uuid.UUID) predicate.Score {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCommentID), v...))
	})
}

// CommentIDGT applies the GT predicate on the "comment_id" field.
func CommentIDGT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCommentID), v))
	})
}

// CommentIDGTE applies the GTE predicate on the "comment_id" field.
func CommentIDGTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCommentID), v))
	})
}

// CommentIDLT applies the LT predicate on the "comment_id" field.
func CommentIDLT(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCommentID), v))
	})
}

// CommentIDLTE applies the LTE predicate on the "comment_id" field.
func CommentIDLTE(v uuid.UUID) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCommentID), v))
	})
}

// CommentIDIsNil applies the IsNil predicate on the "comment_id" field.
func CommentIDIsNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldCommentID)))
	})
}

// CommentIDNotNil applies the NotNil predicate on the "comment_id" field.
func CommentIDNotNil() predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldCommentID)))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Score) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Score) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
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
func Not(p predicate.Score) predicate.Score {
	return predicate.Score(func(s *sql.Selector) {
		p(s.Not())
	})
}
