// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/stocklock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// StockLockCreate is the builder for creating a StockLock entity.
type StockLockCreate struct {
	config
	mutation *StockLockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (slc *StockLockCreate) SetCreatedAt(u uint32) *StockLockCreate {
	slc.mutation.SetCreatedAt(u)
	return slc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableCreatedAt(u *uint32) *StockLockCreate {
	if u != nil {
		slc.SetCreatedAt(*u)
	}
	return slc
}

// SetUpdatedAt sets the "updated_at" field.
func (slc *StockLockCreate) SetUpdatedAt(u uint32) *StockLockCreate {
	slc.mutation.SetUpdatedAt(u)
	return slc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableUpdatedAt(u *uint32) *StockLockCreate {
	if u != nil {
		slc.SetUpdatedAt(*u)
	}
	return slc
}

// SetDeletedAt sets the "deleted_at" field.
func (slc *StockLockCreate) SetDeletedAt(u uint32) *StockLockCreate {
	slc.mutation.SetDeletedAt(u)
	return slc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableDeletedAt(u *uint32) *StockLockCreate {
	if u != nil {
		slc.SetDeletedAt(*u)
	}
	return slc
}

// SetEntID sets the "ent_id" field.
func (slc *StockLockCreate) SetEntID(u uuid.UUID) *StockLockCreate {
	slc.mutation.SetEntID(u)
	return slc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableEntID(u *uuid.UUID) *StockLockCreate {
	if u != nil {
		slc.SetEntID(*u)
	}
	return slc
}

// SetStockID sets the "stock_id" field.
func (slc *StockLockCreate) SetStockID(u uuid.UUID) *StockLockCreate {
	slc.mutation.SetStockID(u)
	return slc
}

// SetNillableStockID sets the "stock_id" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableStockID(u *uuid.UUID) *StockLockCreate {
	if u != nil {
		slc.SetStockID(*u)
	}
	return slc
}

// SetUnits sets the "units" field.
func (slc *StockLockCreate) SetUnits(d decimal.Decimal) *StockLockCreate {
	slc.mutation.SetUnits(d)
	return slc
}

// SetNillableUnits sets the "units" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableUnits(d *decimal.Decimal) *StockLockCreate {
	if d != nil {
		slc.SetUnits(*d)
	}
	return slc
}

// SetLockState sets the "lock_state" field.
func (slc *StockLockCreate) SetLockState(s string) *StockLockCreate {
	slc.mutation.SetLockState(s)
	return slc
}

// SetNillableLockState sets the "lock_state" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableLockState(s *string) *StockLockCreate {
	if s != nil {
		slc.SetLockState(*s)
	}
	return slc
}

// SetChargeBackState sets the "charge_back_state" field.
func (slc *StockLockCreate) SetChargeBackState(s string) *StockLockCreate {
	slc.mutation.SetChargeBackState(s)
	return slc
}

// SetNillableChargeBackState sets the "charge_back_state" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableChargeBackState(s *string) *StockLockCreate {
	if s != nil {
		slc.SetChargeBackState(*s)
	}
	return slc
}

// SetExLockID sets the "ex_lock_id" field.
func (slc *StockLockCreate) SetExLockID(u uuid.UUID) *StockLockCreate {
	slc.mutation.SetExLockID(u)
	return slc
}

// SetNillableExLockID sets the "ex_lock_id" field if the given value is not nil.
func (slc *StockLockCreate) SetNillableExLockID(u *uuid.UUID) *StockLockCreate {
	if u != nil {
		slc.SetExLockID(*u)
	}
	return slc
}

// SetID sets the "id" field.
func (slc *StockLockCreate) SetID(u uint32) *StockLockCreate {
	slc.mutation.SetID(u)
	return slc
}

// Mutation returns the StockLockMutation object of the builder.
func (slc *StockLockCreate) Mutation() *StockLockMutation {
	return slc.mutation
}

// Save creates the StockLock in the database.
func (slc *StockLockCreate) Save(ctx context.Context) (*StockLock, error) {
	var (
		err  error
		node *StockLock
	)
	if err := slc.defaults(); err != nil {
		return nil, err
	}
	if len(slc.hooks) == 0 {
		if err = slc.check(); err != nil {
			return nil, err
		}
		node, err = slc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = slc.check(); err != nil {
				return nil, err
			}
			slc.mutation = mutation
			if node, err = slc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(slc.hooks) - 1; i >= 0; i-- {
			if slc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = slc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, slc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*StockLock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StockLockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (slc *StockLockCreate) SaveX(ctx context.Context) *StockLock {
	v, err := slc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slc *StockLockCreate) Exec(ctx context.Context) error {
	_, err := slc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slc *StockLockCreate) ExecX(ctx context.Context) {
	if err := slc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (slc *StockLockCreate) defaults() error {
	if _, ok := slc.mutation.CreatedAt(); !ok {
		if stocklock.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized stocklock.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := stocklock.DefaultCreatedAt()
		slc.mutation.SetCreatedAt(v)
	}
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		if stocklock.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized stocklock.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := stocklock.DefaultUpdatedAt()
		slc.mutation.SetUpdatedAt(v)
	}
	if _, ok := slc.mutation.DeletedAt(); !ok {
		if stocklock.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized stocklock.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := stocklock.DefaultDeletedAt()
		slc.mutation.SetDeletedAt(v)
	}
	if _, ok := slc.mutation.EntID(); !ok {
		if stocklock.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized stocklock.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := stocklock.DefaultEntID()
		slc.mutation.SetEntID(v)
	}
	if _, ok := slc.mutation.StockID(); !ok {
		if stocklock.DefaultStockID == nil {
			return fmt.Errorf("ent: uninitialized stocklock.DefaultStockID (forgotten import ent/runtime?)")
		}
		v := stocklock.DefaultStockID()
		slc.mutation.SetStockID(v)
	}
	if _, ok := slc.mutation.Units(); !ok {
		v := stocklock.DefaultUnits
		slc.mutation.SetUnits(v)
	}
	if _, ok := slc.mutation.LockState(); !ok {
		v := stocklock.DefaultLockState
		slc.mutation.SetLockState(v)
	}
	if _, ok := slc.mutation.ChargeBackState(); !ok {
		v := stocklock.DefaultChargeBackState
		slc.mutation.SetChargeBackState(v)
	}
	if _, ok := slc.mutation.ExLockID(); !ok {
		if stocklock.DefaultExLockID == nil {
			return fmt.Errorf("ent: uninitialized stocklock.DefaultExLockID (forgotten import ent/runtime?)")
		}
		v := stocklock.DefaultExLockID()
		slc.mutation.SetExLockID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (slc *StockLockCreate) check() error {
	if _, ok := slc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "StockLock.created_at"`)}
	}
	if _, ok := slc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "StockLock.updated_at"`)}
	}
	if _, ok := slc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "StockLock.deleted_at"`)}
	}
	if _, ok := slc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "StockLock.ent_id"`)}
	}
	return nil
}

func (slc *StockLockCreate) sqlSave(ctx context.Context) (*StockLock, error) {
	_node, _spec := slc.createSpec()
	if err := sqlgraph.CreateNode(ctx, slc.driver, _spec); err != nil {
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

func (slc *StockLockCreate) createSpec() (*StockLock, *sqlgraph.CreateSpec) {
	var (
		_node = &StockLock{config: slc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: stocklock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: stocklock.FieldID,
			},
		}
	)
	_spec.OnConflict = slc.conflict
	if id, ok := slc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := slc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stocklock.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := slc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stocklock.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := slc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stocklock.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := slc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stocklock.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := slc.mutation.StockID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stocklock.FieldStockID,
		})
		_node.StockID = value
	}
	if value, ok := slc.mutation.Units(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stocklock.FieldUnits,
		})
		_node.Units = value
	}
	if value, ok := slc.mutation.LockState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stocklock.FieldLockState,
		})
		_node.LockState = value
	}
	if value, ok := slc.mutation.ChargeBackState(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: stocklock.FieldChargeBackState,
		})
		_node.ChargeBackState = value
	}
	if value, ok := slc.mutation.ExLockID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stocklock.FieldExLockID,
		})
		_node.ExLockID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StockLock.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StockLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (slc *StockLockCreate) OnConflict(opts ...sql.ConflictOption) *StockLockUpsertOne {
	slc.conflict = opts
	return &StockLockUpsertOne{
		create: slc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StockLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (slc *StockLockCreate) OnConflictColumns(columns ...string) *StockLockUpsertOne {
	slc.conflict = append(slc.conflict, sql.ConflictColumns(columns...))
	return &StockLockUpsertOne{
		create: slc,
	}
}

type (
	// StockLockUpsertOne is the builder for "upsert"-ing
	//  one StockLock node.
	StockLockUpsertOne struct {
		create *StockLockCreate
	}

	// StockLockUpsert is the "OnConflict" setter.
	StockLockUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *StockLockUpsert) SetCreatedAt(v uint32) *StockLockUpsert {
	u.Set(stocklock.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateCreatedAt() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *StockLockUpsert) AddCreatedAt(v uint32) *StockLockUpsert {
	u.Add(stocklock.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StockLockUpsert) SetUpdatedAt(v uint32) *StockLockUpsert {
	u.Set(stocklock.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateUpdatedAt() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *StockLockUpsert) AddUpdatedAt(v uint32) *StockLockUpsert {
	u.Add(stocklock.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *StockLockUpsert) SetDeletedAt(v uint32) *StockLockUpsert {
	u.Set(stocklock.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateDeletedAt() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *StockLockUpsert) AddDeletedAt(v uint32) *StockLockUpsert {
	u.Add(stocklock.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *StockLockUpsert) SetEntID(v uuid.UUID) *StockLockUpsert {
	u.Set(stocklock.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateEntID() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldEntID)
	return u
}

// SetStockID sets the "stock_id" field.
func (u *StockLockUpsert) SetStockID(v uuid.UUID) *StockLockUpsert {
	u.Set(stocklock.FieldStockID, v)
	return u
}

// UpdateStockID sets the "stock_id" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateStockID() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldStockID)
	return u
}

// ClearStockID clears the value of the "stock_id" field.
func (u *StockLockUpsert) ClearStockID() *StockLockUpsert {
	u.SetNull(stocklock.FieldStockID)
	return u
}

// SetUnits sets the "units" field.
func (u *StockLockUpsert) SetUnits(v decimal.Decimal) *StockLockUpsert {
	u.Set(stocklock.FieldUnits, v)
	return u
}

// UpdateUnits sets the "units" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateUnits() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldUnits)
	return u
}

// ClearUnits clears the value of the "units" field.
func (u *StockLockUpsert) ClearUnits() *StockLockUpsert {
	u.SetNull(stocklock.FieldUnits)
	return u
}

// SetLockState sets the "lock_state" field.
func (u *StockLockUpsert) SetLockState(v string) *StockLockUpsert {
	u.Set(stocklock.FieldLockState, v)
	return u
}

// UpdateLockState sets the "lock_state" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateLockState() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldLockState)
	return u
}

// ClearLockState clears the value of the "lock_state" field.
func (u *StockLockUpsert) ClearLockState() *StockLockUpsert {
	u.SetNull(stocklock.FieldLockState)
	return u
}

// SetChargeBackState sets the "charge_back_state" field.
func (u *StockLockUpsert) SetChargeBackState(v string) *StockLockUpsert {
	u.Set(stocklock.FieldChargeBackState, v)
	return u
}

// UpdateChargeBackState sets the "charge_back_state" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateChargeBackState() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldChargeBackState)
	return u
}

// ClearChargeBackState clears the value of the "charge_back_state" field.
func (u *StockLockUpsert) ClearChargeBackState() *StockLockUpsert {
	u.SetNull(stocklock.FieldChargeBackState)
	return u
}

// SetExLockID sets the "ex_lock_id" field.
func (u *StockLockUpsert) SetExLockID(v uuid.UUID) *StockLockUpsert {
	u.Set(stocklock.FieldExLockID, v)
	return u
}

// UpdateExLockID sets the "ex_lock_id" field to the value that was provided on create.
func (u *StockLockUpsert) UpdateExLockID() *StockLockUpsert {
	u.SetExcluded(stocklock.FieldExLockID)
	return u
}

// ClearExLockID clears the value of the "ex_lock_id" field.
func (u *StockLockUpsert) ClearExLockID() *StockLockUpsert {
	u.SetNull(stocklock.FieldExLockID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.StockLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(stocklock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *StockLockUpsertOne) UpdateNewValues() *StockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(stocklock.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.StockLock.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *StockLockUpsertOne) Ignore() *StockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StockLockUpsertOne) DoNothing() *StockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StockLockCreate.OnConflict
// documentation for more info.
func (u *StockLockUpsertOne) Update(set func(*StockLockUpsert)) *StockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StockLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StockLockUpsertOne) SetCreatedAt(v uint32) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *StockLockUpsertOne) AddCreatedAt(v uint32) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateCreatedAt() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StockLockUpsertOne) SetUpdatedAt(v uint32) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *StockLockUpsertOne) AddUpdatedAt(v uint32) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateUpdatedAt() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *StockLockUpsertOne) SetDeletedAt(v uint32) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *StockLockUpsertOne) AddDeletedAt(v uint32) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateDeletedAt() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *StockLockUpsertOne) SetEntID(v uuid.UUID) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateEntID() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateEntID()
	})
}

// SetStockID sets the "stock_id" field.
func (u *StockLockUpsertOne) SetStockID(v uuid.UUID) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetStockID(v)
	})
}

// UpdateStockID sets the "stock_id" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateStockID() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateStockID()
	})
}

// ClearStockID clears the value of the "stock_id" field.
func (u *StockLockUpsertOne) ClearStockID() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearStockID()
	})
}

// SetUnits sets the "units" field.
func (u *StockLockUpsertOne) SetUnits(v decimal.Decimal) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetUnits(v)
	})
}

// UpdateUnits sets the "units" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateUnits() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateUnits()
	})
}

// ClearUnits clears the value of the "units" field.
func (u *StockLockUpsertOne) ClearUnits() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearUnits()
	})
}

// SetLockState sets the "lock_state" field.
func (u *StockLockUpsertOne) SetLockState(v string) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetLockState(v)
	})
}

// UpdateLockState sets the "lock_state" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateLockState() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateLockState()
	})
}

// ClearLockState clears the value of the "lock_state" field.
func (u *StockLockUpsertOne) ClearLockState() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearLockState()
	})
}

// SetChargeBackState sets the "charge_back_state" field.
func (u *StockLockUpsertOne) SetChargeBackState(v string) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetChargeBackState(v)
	})
}

// UpdateChargeBackState sets the "charge_back_state" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateChargeBackState() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateChargeBackState()
	})
}

// ClearChargeBackState clears the value of the "charge_back_state" field.
func (u *StockLockUpsertOne) ClearChargeBackState() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearChargeBackState()
	})
}

// SetExLockID sets the "ex_lock_id" field.
func (u *StockLockUpsertOne) SetExLockID(v uuid.UUID) *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.SetExLockID(v)
	})
}

// UpdateExLockID sets the "ex_lock_id" field to the value that was provided on create.
func (u *StockLockUpsertOne) UpdateExLockID() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateExLockID()
	})
}

// ClearExLockID clears the value of the "ex_lock_id" field.
func (u *StockLockUpsertOne) ClearExLockID() *StockLockUpsertOne {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearExLockID()
	})
}

// Exec executes the query.
func (u *StockLockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StockLockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StockLockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StockLockUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StockLockUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StockLockCreateBulk is the builder for creating many StockLock entities in bulk.
type StockLockCreateBulk struct {
	config
	builders []*StockLockCreate
	conflict []sql.ConflictOption
}

// Save creates the StockLock entities in the database.
func (slcb *StockLockCreateBulk) Save(ctx context.Context) ([]*StockLock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(slcb.builders))
	nodes := make([]*StockLock, len(slcb.builders))
	mutators := make([]Mutator, len(slcb.builders))
	for i := range slcb.builders {
		func(i int, root context.Context) {
			builder := slcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockLockMutation)
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
					_, err = mutators[i+1].Mutate(root, slcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = slcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, slcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, slcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (slcb *StockLockCreateBulk) SaveX(ctx context.Context) []*StockLock {
	v, err := slcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (slcb *StockLockCreateBulk) Exec(ctx context.Context) error {
	_, err := slcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (slcb *StockLockCreateBulk) ExecX(ctx context.Context) {
	if err := slcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.StockLock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StockLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (slcb *StockLockCreateBulk) OnConflict(opts ...sql.ConflictOption) *StockLockUpsertBulk {
	slcb.conflict = opts
	return &StockLockUpsertBulk{
		create: slcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.StockLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (slcb *StockLockCreateBulk) OnConflictColumns(columns ...string) *StockLockUpsertBulk {
	slcb.conflict = append(slcb.conflict, sql.ConflictColumns(columns...))
	return &StockLockUpsertBulk{
		create: slcb,
	}
}

// StockLockUpsertBulk is the builder for "upsert"-ing
// a bulk of StockLock nodes.
type StockLockUpsertBulk struct {
	create *StockLockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.StockLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(stocklock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *StockLockUpsertBulk) UpdateNewValues() *StockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(stocklock.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.StockLock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *StockLockUpsertBulk) Ignore() *StockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StockLockUpsertBulk) DoNothing() *StockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StockLockCreateBulk.OnConflict
// documentation for more info.
func (u *StockLockUpsertBulk) Update(set func(*StockLockUpsert)) *StockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StockLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StockLockUpsertBulk) SetCreatedAt(v uint32) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *StockLockUpsertBulk) AddCreatedAt(v uint32) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateCreatedAt() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StockLockUpsertBulk) SetUpdatedAt(v uint32) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *StockLockUpsertBulk) AddUpdatedAt(v uint32) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateUpdatedAt() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *StockLockUpsertBulk) SetDeletedAt(v uint32) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *StockLockUpsertBulk) AddDeletedAt(v uint32) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateDeletedAt() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *StockLockUpsertBulk) SetEntID(v uuid.UUID) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateEntID() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateEntID()
	})
}

// SetStockID sets the "stock_id" field.
func (u *StockLockUpsertBulk) SetStockID(v uuid.UUID) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetStockID(v)
	})
}

// UpdateStockID sets the "stock_id" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateStockID() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateStockID()
	})
}

// ClearStockID clears the value of the "stock_id" field.
func (u *StockLockUpsertBulk) ClearStockID() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearStockID()
	})
}

// SetUnits sets the "units" field.
func (u *StockLockUpsertBulk) SetUnits(v decimal.Decimal) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetUnits(v)
	})
}

// UpdateUnits sets the "units" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateUnits() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateUnits()
	})
}

// ClearUnits clears the value of the "units" field.
func (u *StockLockUpsertBulk) ClearUnits() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearUnits()
	})
}

// SetLockState sets the "lock_state" field.
func (u *StockLockUpsertBulk) SetLockState(v string) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetLockState(v)
	})
}

// UpdateLockState sets the "lock_state" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateLockState() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateLockState()
	})
}

// ClearLockState clears the value of the "lock_state" field.
func (u *StockLockUpsertBulk) ClearLockState() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearLockState()
	})
}

// SetChargeBackState sets the "charge_back_state" field.
func (u *StockLockUpsertBulk) SetChargeBackState(v string) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetChargeBackState(v)
	})
}

// UpdateChargeBackState sets the "charge_back_state" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateChargeBackState() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateChargeBackState()
	})
}

// ClearChargeBackState clears the value of the "charge_back_state" field.
func (u *StockLockUpsertBulk) ClearChargeBackState() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearChargeBackState()
	})
}

// SetExLockID sets the "ex_lock_id" field.
func (u *StockLockUpsertBulk) SetExLockID(v uuid.UUID) *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.SetExLockID(v)
	})
}

// UpdateExLockID sets the "ex_lock_id" field to the value that was provided on create.
func (u *StockLockUpsertBulk) UpdateExLockID() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.UpdateExLockID()
	})
}

// ClearExLockID clears the value of the "ex_lock_id" field.
func (u *StockLockUpsertBulk) ClearExLockID() *StockLockUpsertBulk {
	return u.Update(func(s *StockLockUpsert) {
		s.ClearExLockID()
	})
}

// Exec executes the query.
func (u *StockLockUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StockLockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StockLockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StockLockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}