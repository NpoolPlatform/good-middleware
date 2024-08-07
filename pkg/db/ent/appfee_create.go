// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appfee"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppFeeCreate is the builder for creating a AppFee entity.
type AppFeeCreate struct {
	config
	mutation *AppFeeMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (afc *AppFeeCreate) SetCreatedAt(u uint32) *AppFeeCreate {
	afc.mutation.SetCreatedAt(u)
	return afc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableCreatedAt(u *uint32) *AppFeeCreate {
	if u != nil {
		afc.SetCreatedAt(*u)
	}
	return afc
}

// SetUpdatedAt sets the "updated_at" field.
func (afc *AppFeeCreate) SetUpdatedAt(u uint32) *AppFeeCreate {
	afc.mutation.SetUpdatedAt(u)
	return afc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableUpdatedAt(u *uint32) *AppFeeCreate {
	if u != nil {
		afc.SetUpdatedAt(*u)
	}
	return afc
}

// SetDeletedAt sets the "deleted_at" field.
func (afc *AppFeeCreate) SetDeletedAt(u uint32) *AppFeeCreate {
	afc.mutation.SetDeletedAt(u)
	return afc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableDeletedAt(u *uint32) *AppFeeCreate {
	if u != nil {
		afc.SetDeletedAt(*u)
	}
	return afc
}

// SetEntID sets the "ent_id" field.
func (afc *AppFeeCreate) SetEntID(u uuid.UUID) *AppFeeCreate {
	afc.mutation.SetEntID(u)
	return afc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableEntID(u *uuid.UUID) *AppFeeCreate {
	if u != nil {
		afc.SetEntID(*u)
	}
	return afc
}

// SetAppGoodID sets the "app_good_id" field.
func (afc *AppFeeCreate) SetAppGoodID(u uuid.UUID) *AppFeeCreate {
	afc.mutation.SetAppGoodID(u)
	return afc
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableAppGoodID(u *uuid.UUID) *AppFeeCreate {
	if u != nil {
		afc.SetAppGoodID(*u)
	}
	return afc
}

// SetUnitValue sets the "unit_value" field.
func (afc *AppFeeCreate) SetUnitValue(d decimal.Decimal) *AppFeeCreate {
	afc.mutation.SetUnitValue(d)
	return afc
}

// SetNillableUnitValue sets the "unit_value" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableUnitValue(d *decimal.Decimal) *AppFeeCreate {
	if d != nil {
		afc.SetUnitValue(*d)
	}
	return afc
}

// SetCancelMode sets the "cancel_mode" field.
func (afc *AppFeeCreate) SetCancelMode(s string) *AppFeeCreate {
	afc.mutation.SetCancelMode(s)
	return afc
}

// SetNillableCancelMode sets the "cancel_mode" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableCancelMode(s *string) *AppFeeCreate {
	if s != nil {
		afc.SetCancelMode(*s)
	}
	return afc
}

// SetMinOrderDurationSeconds sets the "min_order_duration_seconds" field.
func (afc *AppFeeCreate) SetMinOrderDurationSeconds(u uint32) *AppFeeCreate {
	afc.mutation.SetMinOrderDurationSeconds(u)
	return afc
}

// SetNillableMinOrderDurationSeconds sets the "min_order_duration_seconds" field if the given value is not nil.
func (afc *AppFeeCreate) SetNillableMinOrderDurationSeconds(u *uint32) *AppFeeCreate {
	if u != nil {
		afc.SetMinOrderDurationSeconds(*u)
	}
	return afc
}

// SetID sets the "id" field.
func (afc *AppFeeCreate) SetID(u uint32) *AppFeeCreate {
	afc.mutation.SetID(u)
	return afc
}

// Mutation returns the AppFeeMutation object of the builder.
func (afc *AppFeeCreate) Mutation() *AppFeeMutation {
	return afc.mutation
}

// Save creates the AppFee in the database.
func (afc *AppFeeCreate) Save(ctx context.Context) (*AppFee, error) {
	var (
		err  error
		node *AppFee
	)
	if err := afc.defaults(); err != nil {
		return nil, err
	}
	if len(afc.hooks) == 0 {
		if err = afc.check(); err != nil {
			return nil, err
		}
		node, err = afc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppFeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = afc.check(); err != nil {
				return nil, err
			}
			afc.mutation = mutation
			if node, err = afc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(afc.hooks) - 1; i >= 0; i-- {
			if afc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = afc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, afc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppFee)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppFeeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (afc *AppFeeCreate) SaveX(ctx context.Context) *AppFee {
	v, err := afc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afc *AppFeeCreate) Exec(ctx context.Context) error {
	_, err := afc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afc *AppFeeCreate) ExecX(ctx context.Context) {
	if err := afc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (afc *AppFeeCreate) defaults() error {
	if _, ok := afc.mutation.CreatedAt(); !ok {
		if appfee.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appfee.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appfee.DefaultCreatedAt()
		afc.mutation.SetCreatedAt(v)
	}
	if _, ok := afc.mutation.UpdatedAt(); !ok {
		if appfee.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appfee.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appfee.DefaultUpdatedAt()
		afc.mutation.SetUpdatedAt(v)
	}
	if _, ok := afc.mutation.DeletedAt(); !ok {
		if appfee.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appfee.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appfee.DefaultDeletedAt()
		afc.mutation.SetDeletedAt(v)
	}
	if _, ok := afc.mutation.EntID(); !ok {
		if appfee.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized appfee.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := appfee.DefaultEntID()
		afc.mutation.SetEntID(v)
	}
	if _, ok := afc.mutation.AppGoodID(); !ok {
		if appfee.DefaultAppGoodID == nil {
			return fmt.Errorf("ent: uninitialized appfee.DefaultAppGoodID (forgotten import ent/runtime?)")
		}
		v := appfee.DefaultAppGoodID()
		afc.mutation.SetAppGoodID(v)
	}
	if _, ok := afc.mutation.UnitValue(); !ok {
		v := appfee.DefaultUnitValue
		afc.mutation.SetUnitValue(v)
	}
	if _, ok := afc.mutation.CancelMode(); !ok {
		v := appfee.DefaultCancelMode
		afc.mutation.SetCancelMode(v)
	}
	if _, ok := afc.mutation.MinOrderDurationSeconds(); !ok {
		v := appfee.DefaultMinOrderDurationSeconds
		afc.mutation.SetMinOrderDurationSeconds(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (afc *AppFeeCreate) check() error {
	if _, ok := afc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppFee.created_at"`)}
	}
	if _, ok := afc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppFee.updated_at"`)}
	}
	if _, ok := afc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppFee.deleted_at"`)}
	}
	if _, ok := afc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "AppFee.ent_id"`)}
	}
	return nil
}

func (afc *AppFeeCreate) sqlSave(ctx context.Context) (*AppFee, error) {
	_node, _spec := afc.createSpec()
	if err := sqlgraph.CreateNode(ctx, afc.driver, _spec); err != nil {
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

func (afc *AppFeeCreate) createSpec() (*AppFee, *sqlgraph.CreateSpec) {
	var (
		_node = &AppFee{config: afc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appfee.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appfee.FieldID,
			},
		}
	)
	_spec.OnConflict = afc.conflict
	if id, ok := afc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := afc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := afc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := afc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := afc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appfee.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := afc.mutation.AppGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appfee.FieldAppGoodID,
		})
		_node.AppGoodID = value
	}
	if value, ok := afc.mutation.UnitValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appfee.FieldUnitValue,
		})
		_node.UnitValue = value
	}
	if value, ok := afc.mutation.CancelMode(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appfee.FieldCancelMode,
		})
		_node.CancelMode = value
	}
	if value, ok := afc.mutation.MinOrderDurationSeconds(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldMinOrderDurationSeconds,
		})
		_node.MinOrderDurationSeconds = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppFee.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppFeeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (afc *AppFeeCreate) OnConflict(opts ...sql.ConflictOption) *AppFeeUpsertOne {
	afc.conflict = opts
	return &AppFeeUpsertOne{
		create: afc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppFee.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (afc *AppFeeCreate) OnConflictColumns(columns ...string) *AppFeeUpsertOne {
	afc.conflict = append(afc.conflict, sql.ConflictColumns(columns...))
	return &AppFeeUpsertOne{
		create: afc,
	}
}

type (
	// AppFeeUpsertOne is the builder for "upsert"-ing
	//  one AppFee node.
	AppFeeUpsertOne struct {
		create *AppFeeCreate
	}

	// AppFeeUpsert is the "OnConflict" setter.
	AppFeeUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppFeeUpsert) SetCreatedAt(v uint32) *AppFeeUpsert {
	u.Set(appfee.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateCreatedAt() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppFeeUpsert) AddCreatedAt(v uint32) *AppFeeUpsert {
	u.Add(appfee.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppFeeUpsert) SetUpdatedAt(v uint32) *AppFeeUpsert {
	u.Set(appfee.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateUpdatedAt() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppFeeUpsert) AddUpdatedAt(v uint32) *AppFeeUpsert {
	u.Add(appfee.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppFeeUpsert) SetDeletedAt(v uint32) *AppFeeUpsert {
	u.Set(appfee.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateDeletedAt() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppFeeUpsert) AddDeletedAt(v uint32) *AppFeeUpsert {
	u.Add(appfee.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppFeeUpsert) SetEntID(v uuid.UUID) *AppFeeUpsert {
	u.Set(appfee.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateEntID() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldEntID)
	return u
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppFeeUpsert) SetAppGoodID(v uuid.UUID) *AppFeeUpsert {
	u.Set(appfee.FieldAppGoodID, v)
	return u
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateAppGoodID() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldAppGoodID)
	return u
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppFeeUpsert) ClearAppGoodID() *AppFeeUpsert {
	u.SetNull(appfee.FieldAppGoodID)
	return u
}

// SetUnitValue sets the "unit_value" field.
func (u *AppFeeUpsert) SetUnitValue(v decimal.Decimal) *AppFeeUpsert {
	u.Set(appfee.FieldUnitValue, v)
	return u
}

// UpdateUnitValue sets the "unit_value" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateUnitValue() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldUnitValue)
	return u
}

// ClearUnitValue clears the value of the "unit_value" field.
func (u *AppFeeUpsert) ClearUnitValue() *AppFeeUpsert {
	u.SetNull(appfee.FieldUnitValue)
	return u
}

// SetCancelMode sets the "cancel_mode" field.
func (u *AppFeeUpsert) SetCancelMode(v string) *AppFeeUpsert {
	u.Set(appfee.FieldCancelMode, v)
	return u
}

// UpdateCancelMode sets the "cancel_mode" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateCancelMode() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldCancelMode)
	return u
}

// ClearCancelMode clears the value of the "cancel_mode" field.
func (u *AppFeeUpsert) ClearCancelMode() *AppFeeUpsert {
	u.SetNull(appfee.FieldCancelMode)
	return u
}

// SetMinOrderDurationSeconds sets the "min_order_duration_seconds" field.
func (u *AppFeeUpsert) SetMinOrderDurationSeconds(v uint32) *AppFeeUpsert {
	u.Set(appfee.FieldMinOrderDurationSeconds, v)
	return u
}

// UpdateMinOrderDurationSeconds sets the "min_order_duration_seconds" field to the value that was provided on create.
func (u *AppFeeUpsert) UpdateMinOrderDurationSeconds() *AppFeeUpsert {
	u.SetExcluded(appfee.FieldMinOrderDurationSeconds)
	return u
}

// AddMinOrderDurationSeconds adds v to the "min_order_duration_seconds" field.
func (u *AppFeeUpsert) AddMinOrderDurationSeconds(v uint32) *AppFeeUpsert {
	u.Add(appfee.FieldMinOrderDurationSeconds, v)
	return u
}

// ClearMinOrderDurationSeconds clears the value of the "min_order_duration_seconds" field.
func (u *AppFeeUpsert) ClearMinOrderDurationSeconds() *AppFeeUpsert {
	u.SetNull(appfee.FieldMinOrderDurationSeconds)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppFee.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appfee.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppFeeUpsertOne) UpdateNewValues() *AppFeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appfee.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppFee.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppFeeUpsertOne) Ignore() *AppFeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppFeeUpsertOne) DoNothing() *AppFeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppFeeCreate.OnConflict
// documentation for more info.
func (u *AppFeeUpsertOne) Update(set func(*AppFeeUpsert)) *AppFeeUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppFeeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppFeeUpsertOne) SetCreatedAt(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppFeeUpsertOne) AddCreatedAt(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateCreatedAt() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppFeeUpsertOne) SetUpdatedAt(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppFeeUpsertOne) AddUpdatedAt(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateUpdatedAt() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppFeeUpsertOne) SetDeletedAt(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppFeeUpsertOne) AddDeletedAt(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateDeletedAt() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppFeeUpsertOne) SetEntID(v uuid.UUID) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateEntID() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppFeeUpsertOne) SetAppGoodID(v uuid.UUID) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateAppGoodID() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppFeeUpsertOne) ClearAppGoodID() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearAppGoodID()
	})
}

// SetUnitValue sets the "unit_value" field.
func (u *AppFeeUpsertOne) SetUnitValue(v decimal.Decimal) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetUnitValue(v)
	})
}

// UpdateUnitValue sets the "unit_value" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateUnitValue() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateUnitValue()
	})
}

// ClearUnitValue clears the value of the "unit_value" field.
func (u *AppFeeUpsertOne) ClearUnitValue() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearUnitValue()
	})
}

// SetCancelMode sets the "cancel_mode" field.
func (u *AppFeeUpsertOne) SetCancelMode(v string) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetCancelMode(v)
	})
}

// UpdateCancelMode sets the "cancel_mode" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateCancelMode() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateCancelMode()
	})
}

// ClearCancelMode clears the value of the "cancel_mode" field.
func (u *AppFeeUpsertOne) ClearCancelMode() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearCancelMode()
	})
}

// SetMinOrderDurationSeconds sets the "min_order_duration_seconds" field.
func (u *AppFeeUpsertOne) SetMinOrderDurationSeconds(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetMinOrderDurationSeconds(v)
	})
}

// AddMinOrderDurationSeconds adds v to the "min_order_duration_seconds" field.
func (u *AppFeeUpsertOne) AddMinOrderDurationSeconds(v uint32) *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddMinOrderDurationSeconds(v)
	})
}

// UpdateMinOrderDurationSeconds sets the "min_order_duration_seconds" field to the value that was provided on create.
func (u *AppFeeUpsertOne) UpdateMinOrderDurationSeconds() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateMinOrderDurationSeconds()
	})
}

// ClearMinOrderDurationSeconds clears the value of the "min_order_duration_seconds" field.
func (u *AppFeeUpsertOne) ClearMinOrderDurationSeconds() *AppFeeUpsertOne {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearMinOrderDurationSeconds()
	})
}

// Exec executes the query.
func (u *AppFeeUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppFeeCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppFeeUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppFeeUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppFeeUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppFeeCreateBulk is the builder for creating many AppFee entities in bulk.
type AppFeeCreateBulk struct {
	config
	builders []*AppFeeCreate
	conflict []sql.ConflictOption
}

// Save creates the AppFee entities in the database.
func (afcb *AppFeeCreateBulk) Save(ctx context.Context) ([]*AppFee, error) {
	specs := make([]*sqlgraph.CreateSpec, len(afcb.builders))
	nodes := make([]*AppFee, len(afcb.builders))
	mutators := make([]Mutator, len(afcb.builders))
	for i := range afcb.builders {
		func(i int, root context.Context) {
			builder := afcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppFeeMutation)
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
					_, err = mutators[i+1].Mutate(root, afcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = afcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, afcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, afcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (afcb *AppFeeCreateBulk) SaveX(ctx context.Context) []*AppFee {
	v, err := afcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (afcb *AppFeeCreateBulk) Exec(ctx context.Context) error {
	_, err := afcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afcb *AppFeeCreateBulk) ExecX(ctx context.Context) {
	if err := afcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppFee.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppFeeUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (afcb *AppFeeCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppFeeUpsertBulk {
	afcb.conflict = opts
	return &AppFeeUpsertBulk{
		create: afcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppFee.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (afcb *AppFeeCreateBulk) OnConflictColumns(columns ...string) *AppFeeUpsertBulk {
	afcb.conflict = append(afcb.conflict, sql.ConflictColumns(columns...))
	return &AppFeeUpsertBulk{
		create: afcb,
	}
}

// AppFeeUpsertBulk is the builder for "upsert"-ing
// a bulk of AppFee nodes.
type AppFeeUpsertBulk struct {
	create *AppFeeCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppFee.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appfee.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppFeeUpsertBulk) UpdateNewValues() *AppFeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appfee.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppFee.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppFeeUpsertBulk) Ignore() *AppFeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppFeeUpsertBulk) DoNothing() *AppFeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppFeeCreateBulk.OnConflict
// documentation for more info.
func (u *AppFeeUpsertBulk) Update(set func(*AppFeeUpsert)) *AppFeeUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppFeeUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppFeeUpsertBulk) SetCreatedAt(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppFeeUpsertBulk) AddCreatedAt(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateCreatedAt() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppFeeUpsertBulk) SetUpdatedAt(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppFeeUpsertBulk) AddUpdatedAt(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateUpdatedAt() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppFeeUpsertBulk) SetDeletedAt(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppFeeUpsertBulk) AddDeletedAt(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateDeletedAt() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppFeeUpsertBulk) SetEntID(v uuid.UUID) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateEntID() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppFeeUpsertBulk) SetAppGoodID(v uuid.UUID) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateAppGoodID() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppFeeUpsertBulk) ClearAppGoodID() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearAppGoodID()
	})
}

// SetUnitValue sets the "unit_value" field.
func (u *AppFeeUpsertBulk) SetUnitValue(v decimal.Decimal) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetUnitValue(v)
	})
}

// UpdateUnitValue sets the "unit_value" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateUnitValue() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateUnitValue()
	})
}

// ClearUnitValue clears the value of the "unit_value" field.
func (u *AppFeeUpsertBulk) ClearUnitValue() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearUnitValue()
	})
}

// SetCancelMode sets the "cancel_mode" field.
func (u *AppFeeUpsertBulk) SetCancelMode(v string) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetCancelMode(v)
	})
}

// UpdateCancelMode sets the "cancel_mode" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateCancelMode() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateCancelMode()
	})
}

// ClearCancelMode clears the value of the "cancel_mode" field.
func (u *AppFeeUpsertBulk) ClearCancelMode() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearCancelMode()
	})
}

// SetMinOrderDurationSeconds sets the "min_order_duration_seconds" field.
func (u *AppFeeUpsertBulk) SetMinOrderDurationSeconds(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.SetMinOrderDurationSeconds(v)
	})
}

// AddMinOrderDurationSeconds adds v to the "min_order_duration_seconds" field.
func (u *AppFeeUpsertBulk) AddMinOrderDurationSeconds(v uint32) *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.AddMinOrderDurationSeconds(v)
	})
}

// UpdateMinOrderDurationSeconds sets the "min_order_duration_seconds" field to the value that was provided on create.
func (u *AppFeeUpsertBulk) UpdateMinOrderDurationSeconds() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.UpdateMinOrderDurationSeconds()
	})
}

// ClearMinOrderDurationSeconds clears the value of the "min_order_duration_seconds" field.
func (u *AppFeeUpsertBulk) ClearMinOrderDurationSeconds() *AppFeeUpsertBulk {
	return u.Update(func(s *AppFeeUpsert) {
		s.ClearMinOrderDurationSeconds()
	})
}

// Exec executes the query.
func (u *AppFeeUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppFeeCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppFeeCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppFeeUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
