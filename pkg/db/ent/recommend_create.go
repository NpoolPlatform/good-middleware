// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/recommend"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// RecommendCreate is the builder for creating a Recommend entity.
type RecommendCreate struct {
	config
	mutation *RecommendMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (rc *RecommendCreate) SetCreatedAt(u uint32) *RecommendCreate {
	rc.mutation.SetCreatedAt(u)
	return rc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableCreatedAt(u *uint32) *RecommendCreate {
	if u != nil {
		rc.SetCreatedAt(*u)
	}
	return rc
}

// SetUpdatedAt sets the "updated_at" field.
func (rc *RecommendCreate) SetUpdatedAt(u uint32) *RecommendCreate {
	rc.mutation.SetUpdatedAt(u)
	return rc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableUpdatedAt(u *uint32) *RecommendCreate {
	if u != nil {
		rc.SetUpdatedAt(*u)
	}
	return rc
}

// SetDeletedAt sets the "deleted_at" field.
func (rc *RecommendCreate) SetDeletedAt(u uint32) *RecommendCreate {
	rc.mutation.SetDeletedAt(u)
	return rc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableDeletedAt(u *uint32) *RecommendCreate {
	if u != nil {
		rc.SetDeletedAt(*u)
	}
	return rc
}

// SetEntID sets the "ent_id" field.
func (rc *RecommendCreate) SetEntID(u uuid.UUID) *RecommendCreate {
	rc.mutation.SetEntID(u)
	return rc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableEntID(u *uuid.UUID) *RecommendCreate {
	if u != nil {
		rc.SetEntID(*u)
	}
	return rc
}

// SetAppGoodID sets the "app_good_id" field.
func (rc *RecommendCreate) SetAppGoodID(u uuid.UUID) *RecommendCreate {
	rc.mutation.SetAppGoodID(u)
	return rc
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableAppGoodID(u *uuid.UUID) *RecommendCreate {
	if u != nil {
		rc.SetAppGoodID(*u)
	}
	return rc
}

// SetRecommenderID sets the "recommender_id" field.
func (rc *RecommendCreate) SetRecommenderID(u uuid.UUID) *RecommendCreate {
	rc.mutation.SetRecommenderID(u)
	return rc
}

// SetNillableRecommenderID sets the "recommender_id" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableRecommenderID(u *uuid.UUID) *RecommendCreate {
	if u != nil {
		rc.SetRecommenderID(*u)
	}
	return rc
}

// SetMessage sets the "message" field.
func (rc *RecommendCreate) SetMessage(s string) *RecommendCreate {
	rc.mutation.SetMessage(s)
	return rc
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableMessage(s *string) *RecommendCreate {
	if s != nil {
		rc.SetMessage(*s)
	}
	return rc
}

// SetRecommendIndex sets the "recommend_index" field.
func (rc *RecommendCreate) SetRecommendIndex(d decimal.Decimal) *RecommendCreate {
	rc.mutation.SetRecommendIndex(d)
	return rc
}

// SetNillableRecommendIndex sets the "recommend_index" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableRecommendIndex(d *decimal.Decimal) *RecommendCreate {
	if d != nil {
		rc.SetRecommendIndex(*d)
	}
	return rc
}

// SetHide sets the "hide" field.
func (rc *RecommendCreate) SetHide(b bool) *RecommendCreate {
	rc.mutation.SetHide(b)
	return rc
}

// SetNillableHide sets the "hide" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableHide(b *bool) *RecommendCreate {
	if b != nil {
		rc.SetHide(*b)
	}
	return rc
}

// SetHideReason sets the "hide_reason" field.
func (rc *RecommendCreate) SetHideReason(s string) *RecommendCreate {
	rc.mutation.SetHideReason(s)
	return rc
}

// SetNillableHideReason sets the "hide_reason" field if the given value is not nil.
func (rc *RecommendCreate) SetNillableHideReason(s *string) *RecommendCreate {
	if s != nil {
		rc.SetHideReason(*s)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *RecommendCreate) SetID(u uint32) *RecommendCreate {
	rc.mutation.SetID(u)
	return rc
}

// Mutation returns the RecommendMutation object of the builder.
func (rc *RecommendCreate) Mutation() *RecommendMutation {
	return rc.mutation
}

// Save creates the Recommend in the database.
func (rc *RecommendCreate) Save(ctx context.Context) (*Recommend, error) {
	var (
		err  error
		node *Recommend
	)
	if err := rc.defaults(); err != nil {
		return nil, err
	}
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecommendMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Recommend)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RecommendMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecommendCreate) SaveX(ctx context.Context) *Recommend {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecommendCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecommendCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RecommendCreate) defaults() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		if recommend.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized recommend.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := recommend.DefaultCreatedAt()
		rc.mutation.SetCreatedAt(v)
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		if recommend.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized recommend.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := recommend.DefaultUpdatedAt()
		rc.mutation.SetUpdatedAt(v)
	}
	if _, ok := rc.mutation.DeletedAt(); !ok {
		if recommend.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized recommend.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := recommend.DefaultDeletedAt()
		rc.mutation.SetDeletedAt(v)
	}
	if _, ok := rc.mutation.EntID(); !ok {
		if recommend.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized recommend.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := recommend.DefaultEntID()
		rc.mutation.SetEntID(v)
	}
	if _, ok := rc.mutation.AppGoodID(); !ok {
		if recommend.DefaultAppGoodID == nil {
			return fmt.Errorf("ent: uninitialized recommend.DefaultAppGoodID (forgotten import ent/runtime?)")
		}
		v := recommend.DefaultAppGoodID()
		rc.mutation.SetAppGoodID(v)
	}
	if _, ok := rc.mutation.RecommenderID(); !ok {
		if recommend.DefaultRecommenderID == nil {
			return fmt.Errorf("ent: uninitialized recommend.DefaultRecommenderID (forgotten import ent/runtime?)")
		}
		v := recommend.DefaultRecommenderID()
		rc.mutation.SetRecommenderID(v)
	}
	if _, ok := rc.mutation.Message(); !ok {
		v := recommend.DefaultMessage
		rc.mutation.SetMessage(v)
	}
	if _, ok := rc.mutation.RecommendIndex(); !ok {
		v := recommend.DefaultRecommendIndex
		rc.mutation.SetRecommendIndex(v)
	}
	if _, ok := rc.mutation.Hide(); !ok {
		v := recommend.DefaultHide
		rc.mutation.SetHide(v)
	}
	if _, ok := rc.mutation.HideReason(); !ok {
		v := recommend.DefaultHideReason
		rc.mutation.SetHideReason(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecommendCreate) check() error {
	if _, ok := rc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Recommend.created_at"`)}
	}
	if _, ok := rc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Recommend.updated_at"`)}
	}
	if _, ok := rc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Recommend.deleted_at"`)}
	}
	if _, ok := rc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "Recommend.ent_id"`)}
	}
	return nil
}

func (rc *RecommendCreate) sqlSave(ctx context.Context) (*Recommend, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (rc *RecommendCreate) createSpec() (*Recommend, *sqlgraph.CreateSpec) {
	var (
		_node = &Recommend{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recommend.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: recommend.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: recommend.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := rc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: recommend.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := rc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: recommend.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := rc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := rc.mutation.AppGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldAppGoodID,
		})
		_node.AppGoodID = value
	}
	if value, ok := rc.mutation.RecommenderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: recommend.FieldRecommenderID,
		})
		_node.RecommenderID = value
	}
	if value, ok := rc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recommend.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := rc.mutation.RecommendIndex(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: recommend.FieldRecommendIndex,
		})
		_node.RecommendIndex = value
	}
	if value, ok := rc.mutation.Hide(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: recommend.FieldHide,
		})
		_node.Hide = value
	}
	if value, ok := rc.mutation.HideReason(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: recommend.FieldHideReason,
		})
		_node.HideReason = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Recommend.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecommendUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rc *RecommendCreate) OnConflict(opts ...sql.ConflictOption) *RecommendUpsertOne {
	rc.conflict = opts
	return &RecommendUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Recommend.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rc *RecommendCreate) OnConflictColumns(columns ...string) *RecommendUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &RecommendUpsertOne{
		create: rc,
	}
}

type (
	// RecommendUpsertOne is the builder for "upsert"-ing
	//  one Recommend node.
	RecommendUpsertOne struct {
		create *RecommendCreate
	}

	// RecommendUpsert is the "OnConflict" setter.
	RecommendUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *RecommendUpsert) SetCreatedAt(v uint32) *RecommendUpsert {
	u.Set(recommend.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateCreatedAt() *RecommendUpsert {
	u.SetExcluded(recommend.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *RecommendUpsert) AddCreatedAt(v uint32) *RecommendUpsert {
	u.Add(recommend.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RecommendUpsert) SetUpdatedAt(v uint32) *RecommendUpsert {
	u.Set(recommend.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateUpdatedAt() *RecommendUpsert {
	u.SetExcluded(recommend.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *RecommendUpsert) AddUpdatedAt(v uint32) *RecommendUpsert {
	u.Add(recommend.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RecommendUpsert) SetDeletedAt(v uint32) *RecommendUpsert {
	u.Set(recommend.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateDeletedAt() *RecommendUpsert {
	u.SetExcluded(recommend.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RecommendUpsert) AddDeletedAt(v uint32) *RecommendUpsert {
	u.Add(recommend.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *RecommendUpsert) SetEntID(v uuid.UUID) *RecommendUpsert {
	u.Set(recommend.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateEntID() *RecommendUpsert {
	u.SetExcluded(recommend.FieldEntID)
	return u
}

// SetAppGoodID sets the "app_good_id" field.
func (u *RecommendUpsert) SetAppGoodID(v uuid.UUID) *RecommendUpsert {
	u.Set(recommend.FieldAppGoodID, v)
	return u
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateAppGoodID() *RecommendUpsert {
	u.SetExcluded(recommend.FieldAppGoodID)
	return u
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *RecommendUpsert) ClearAppGoodID() *RecommendUpsert {
	u.SetNull(recommend.FieldAppGoodID)
	return u
}

// SetRecommenderID sets the "recommender_id" field.
func (u *RecommendUpsert) SetRecommenderID(v uuid.UUID) *RecommendUpsert {
	u.Set(recommend.FieldRecommenderID, v)
	return u
}

// UpdateRecommenderID sets the "recommender_id" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateRecommenderID() *RecommendUpsert {
	u.SetExcluded(recommend.FieldRecommenderID)
	return u
}

// ClearRecommenderID clears the value of the "recommender_id" field.
func (u *RecommendUpsert) ClearRecommenderID() *RecommendUpsert {
	u.SetNull(recommend.FieldRecommenderID)
	return u
}

// SetMessage sets the "message" field.
func (u *RecommendUpsert) SetMessage(v string) *RecommendUpsert {
	u.Set(recommend.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateMessage() *RecommendUpsert {
	u.SetExcluded(recommend.FieldMessage)
	return u
}

// ClearMessage clears the value of the "message" field.
func (u *RecommendUpsert) ClearMessage() *RecommendUpsert {
	u.SetNull(recommend.FieldMessage)
	return u
}

// SetRecommendIndex sets the "recommend_index" field.
func (u *RecommendUpsert) SetRecommendIndex(v decimal.Decimal) *RecommendUpsert {
	u.Set(recommend.FieldRecommendIndex, v)
	return u
}

// UpdateRecommendIndex sets the "recommend_index" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateRecommendIndex() *RecommendUpsert {
	u.SetExcluded(recommend.FieldRecommendIndex)
	return u
}

// ClearRecommendIndex clears the value of the "recommend_index" field.
func (u *RecommendUpsert) ClearRecommendIndex() *RecommendUpsert {
	u.SetNull(recommend.FieldRecommendIndex)
	return u
}

// SetHide sets the "hide" field.
func (u *RecommendUpsert) SetHide(v bool) *RecommendUpsert {
	u.Set(recommend.FieldHide, v)
	return u
}

// UpdateHide sets the "hide" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateHide() *RecommendUpsert {
	u.SetExcluded(recommend.FieldHide)
	return u
}

// ClearHide clears the value of the "hide" field.
func (u *RecommendUpsert) ClearHide() *RecommendUpsert {
	u.SetNull(recommend.FieldHide)
	return u
}

// SetHideReason sets the "hide_reason" field.
func (u *RecommendUpsert) SetHideReason(v string) *RecommendUpsert {
	u.Set(recommend.FieldHideReason, v)
	return u
}

// UpdateHideReason sets the "hide_reason" field to the value that was provided on create.
func (u *RecommendUpsert) UpdateHideReason() *RecommendUpsert {
	u.SetExcluded(recommend.FieldHideReason)
	return u
}

// ClearHideReason clears the value of the "hide_reason" field.
func (u *RecommendUpsert) ClearHideReason() *RecommendUpsert {
	u.SetNull(recommend.FieldHideReason)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Recommend.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(recommend.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RecommendUpsertOne) UpdateNewValues() *RecommendUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(recommend.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Recommend.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RecommendUpsertOne) Ignore() *RecommendUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecommendUpsertOne) DoNothing() *RecommendUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecommendCreate.OnConflict
// documentation for more info.
func (u *RecommendUpsertOne) Update(set func(*RecommendUpsert)) *RecommendUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecommendUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *RecommendUpsertOne) SetCreatedAt(v uint32) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *RecommendUpsertOne) AddCreatedAt(v uint32) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateCreatedAt() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RecommendUpsertOne) SetUpdatedAt(v uint32) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *RecommendUpsertOne) AddUpdatedAt(v uint32) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateUpdatedAt() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RecommendUpsertOne) SetDeletedAt(v uint32) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RecommendUpsertOne) AddDeletedAt(v uint32) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateDeletedAt() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *RecommendUpsertOne) SetEntID(v uuid.UUID) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateEntID() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *RecommendUpsertOne) SetAppGoodID(v uuid.UUID) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateAppGoodID() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *RecommendUpsertOne) ClearAppGoodID() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearAppGoodID()
	})
}

// SetRecommenderID sets the "recommender_id" field.
func (u *RecommendUpsertOne) SetRecommenderID(v uuid.UUID) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetRecommenderID(v)
	})
}

// UpdateRecommenderID sets the "recommender_id" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateRecommenderID() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateRecommenderID()
	})
}

// ClearRecommenderID clears the value of the "recommender_id" field.
func (u *RecommendUpsertOne) ClearRecommenderID() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearRecommenderID()
	})
}

// SetMessage sets the "message" field.
func (u *RecommendUpsertOne) SetMessage(v string) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateMessage() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *RecommendUpsertOne) ClearMessage() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearMessage()
	})
}

// SetRecommendIndex sets the "recommend_index" field.
func (u *RecommendUpsertOne) SetRecommendIndex(v decimal.Decimal) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetRecommendIndex(v)
	})
}

// UpdateRecommendIndex sets the "recommend_index" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateRecommendIndex() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateRecommendIndex()
	})
}

// ClearRecommendIndex clears the value of the "recommend_index" field.
func (u *RecommendUpsertOne) ClearRecommendIndex() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearRecommendIndex()
	})
}

// SetHide sets the "hide" field.
func (u *RecommendUpsertOne) SetHide(v bool) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetHide(v)
	})
}

// UpdateHide sets the "hide" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateHide() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateHide()
	})
}

// ClearHide clears the value of the "hide" field.
func (u *RecommendUpsertOne) ClearHide() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearHide()
	})
}

// SetHideReason sets the "hide_reason" field.
func (u *RecommendUpsertOne) SetHideReason(v string) *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.SetHideReason(v)
	})
}

// UpdateHideReason sets the "hide_reason" field to the value that was provided on create.
func (u *RecommendUpsertOne) UpdateHideReason() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateHideReason()
	})
}

// ClearHideReason clears the value of the "hide_reason" field.
func (u *RecommendUpsertOne) ClearHideReason() *RecommendUpsertOne {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearHideReason()
	})
}

// Exec executes the query.
func (u *RecommendUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecommendCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecommendUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RecommendUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RecommendUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RecommendCreateBulk is the builder for creating many Recommend entities in bulk.
type RecommendCreateBulk struct {
	config
	builders []*RecommendCreate
	conflict []sql.ConflictOption
}

// Save creates the Recommend entities in the database.
func (rcb *RecommendCreateBulk) Save(ctx context.Context) ([]*Recommend, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Recommend, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecommendMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecommendCreateBulk) SaveX(ctx context.Context) []*Recommend {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecommendCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecommendCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Recommend.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RecommendUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (rcb *RecommendCreateBulk) OnConflict(opts ...sql.ConflictOption) *RecommendUpsertBulk {
	rcb.conflict = opts
	return &RecommendUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Recommend.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rcb *RecommendCreateBulk) OnConflictColumns(columns ...string) *RecommendUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &RecommendUpsertBulk{
		create: rcb,
	}
}

// RecommendUpsertBulk is the builder for "upsert"-ing
// a bulk of Recommend nodes.
type RecommendUpsertBulk struct {
	create *RecommendCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Recommend.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(recommend.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *RecommendUpsertBulk) UpdateNewValues() *RecommendUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(recommend.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Recommend.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RecommendUpsertBulk) Ignore() *RecommendUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RecommendUpsertBulk) DoNothing() *RecommendUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RecommendCreateBulk.OnConflict
// documentation for more info.
func (u *RecommendUpsertBulk) Update(set func(*RecommendUpsert)) *RecommendUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RecommendUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *RecommendUpsertBulk) SetCreatedAt(v uint32) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *RecommendUpsertBulk) AddCreatedAt(v uint32) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateCreatedAt() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *RecommendUpsertBulk) SetUpdatedAt(v uint32) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *RecommendUpsertBulk) AddUpdatedAt(v uint32) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateUpdatedAt() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *RecommendUpsertBulk) SetDeletedAt(v uint32) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *RecommendUpsertBulk) AddDeletedAt(v uint32) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateDeletedAt() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *RecommendUpsertBulk) SetEntID(v uuid.UUID) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateEntID() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *RecommendUpsertBulk) SetAppGoodID(v uuid.UUID) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateAppGoodID() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *RecommendUpsertBulk) ClearAppGoodID() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearAppGoodID()
	})
}

// SetRecommenderID sets the "recommender_id" field.
func (u *RecommendUpsertBulk) SetRecommenderID(v uuid.UUID) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetRecommenderID(v)
	})
}

// UpdateRecommenderID sets the "recommender_id" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateRecommenderID() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateRecommenderID()
	})
}

// ClearRecommenderID clears the value of the "recommender_id" field.
func (u *RecommendUpsertBulk) ClearRecommenderID() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearRecommenderID()
	})
}

// SetMessage sets the "message" field.
func (u *RecommendUpsertBulk) SetMessage(v string) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateMessage() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateMessage()
	})
}

// ClearMessage clears the value of the "message" field.
func (u *RecommendUpsertBulk) ClearMessage() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearMessage()
	})
}

// SetRecommendIndex sets the "recommend_index" field.
func (u *RecommendUpsertBulk) SetRecommendIndex(v decimal.Decimal) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetRecommendIndex(v)
	})
}

// UpdateRecommendIndex sets the "recommend_index" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateRecommendIndex() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateRecommendIndex()
	})
}

// ClearRecommendIndex clears the value of the "recommend_index" field.
func (u *RecommendUpsertBulk) ClearRecommendIndex() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearRecommendIndex()
	})
}

// SetHide sets the "hide" field.
func (u *RecommendUpsertBulk) SetHide(v bool) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetHide(v)
	})
}

// UpdateHide sets the "hide" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateHide() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateHide()
	})
}

// ClearHide clears the value of the "hide" field.
func (u *RecommendUpsertBulk) ClearHide() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearHide()
	})
}

// SetHideReason sets the "hide_reason" field.
func (u *RecommendUpsertBulk) SetHideReason(v string) *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.SetHideReason(v)
	})
}

// UpdateHideReason sets the "hide_reason" field to the value that was provided on create.
func (u *RecommendUpsertBulk) UpdateHideReason() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.UpdateHideReason()
	})
}

// ClearHideReason clears the value of the "hide_reason" field.
func (u *RecommendUpsertBulk) ClearHideReason() *RecommendUpsertBulk {
	return u.Update(func(s *RecommendUpsert) {
		s.ClearHideReason()
	})
}

// Exec executes the query.
func (u *RecommendUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RecommendCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RecommendCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RecommendUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
