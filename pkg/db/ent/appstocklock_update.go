// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstocklock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppStockLockUpdate is the builder for updating AppStockLock entities.
type AppStockLockUpdate struct {
	config
	hooks     []Hook
	mutation  *AppStockLockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppStockLockUpdate builder.
func (aslu *AppStockLockUpdate) Where(ps ...predicate.AppStockLock) *AppStockLockUpdate {
	aslu.mutation.Where(ps...)
	return aslu
}

// SetCreatedAt sets the "created_at" field.
func (aslu *AppStockLockUpdate) SetCreatedAt(u uint32) *AppStockLockUpdate {
	aslu.mutation.ResetCreatedAt()
	aslu.mutation.SetCreatedAt(u)
	return aslu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aslu *AppStockLockUpdate) SetNillableCreatedAt(u *uint32) *AppStockLockUpdate {
	if u != nil {
		aslu.SetCreatedAt(*u)
	}
	return aslu
}

// AddCreatedAt adds u to the "created_at" field.
func (aslu *AppStockLockUpdate) AddCreatedAt(u int32) *AppStockLockUpdate {
	aslu.mutation.AddCreatedAt(u)
	return aslu
}

// SetUpdatedAt sets the "updated_at" field.
func (aslu *AppStockLockUpdate) SetUpdatedAt(u uint32) *AppStockLockUpdate {
	aslu.mutation.ResetUpdatedAt()
	aslu.mutation.SetUpdatedAt(u)
	return aslu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (aslu *AppStockLockUpdate) AddUpdatedAt(u int32) *AppStockLockUpdate {
	aslu.mutation.AddUpdatedAt(u)
	return aslu
}

// SetDeletedAt sets the "deleted_at" field.
func (aslu *AppStockLockUpdate) SetDeletedAt(u uint32) *AppStockLockUpdate {
	aslu.mutation.ResetDeletedAt()
	aslu.mutation.SetDeletedAt(u)
	return aslu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aslu *AppStockLockUpdate) SetNillableDeletedAt(u *uint32) *AppStockLockUpdate {
	if u != nil {
		aslu.SetDeletedAt(*u)
	}
	return aslu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (aslu *AppStockLockUpdate) AddDeletedAt(u int32) *AppStockLockUpdate {
	aslu.mutation.AddDeletedAt(u)
	return aslu
}

// SetAppStockID sets the "app_stock_id" field.
func (aslu *AppStockLockUpdate) SetAppStockID(u uuid.UUID) *AppStockLockUpdate {
	aslu.mutation.SetAppStockID(u)
	return aslu
}

// SetNillableAppStockID sets the "app_stock_id" field if the given value is not nil.
func (aslu *AppStockLockUpdate) SetNillableAppStockID(u *uuid.UUID) *AppStockLockUpdate {
	if u != nil {
		aslu.SetAppStockID(*u)
	}
	return aslu
}

// ClearAppStockID clears the value of the "app_stock_id" field.
func (aslu *AppStockLockUpdate) ClearAppStockID() *AppStockLockUpdate {
	aslu.mutation.ClearAppStockID()
	return aslu
}

// SetUnits sets the "units" field.
func (aslu *AppStockLockUpdate) SetUnits(d decimal.Decimal) *AppStockLockUpdate {
	aslu.mutation.SetUnits(d)
	return aslu
}

// SetNillableUnits sets the "units" field if the given value is not nil.
func (aslu *AppStockLockUpdate) SetNillableUnits(d *decimal.Decimal) *AppStockLockUpdate {
	if d != nil {
		aslu.SetUnits(*d)
	}
	return aslu
}

// ClearUnits clears the value of the "units" field.
func (aslu *AppStockLockUpdate) ClearUnits() *AppStockLockUpdate {
	aslu.mutation.ClearUnits()
	return aslu
}

// SetAppSpotUnits sets the "app_spot_units" field.
func (aslu *AppStockLockUpdate) SetAppSpotUnits(d decimal.Decimal) *AppStockLockUpdate {
	aslu.mutation.SetAppSpotUnits(d)
	return aslu
}

// SetNillableAppSpotUnits sets the "app_spot_units" field if the given value is not nil.
func (aslu *AppStockLockUpdate) SetNillableAppSpotUnits(d *decimal.Decimal) *AppStockLockUpdate {
	if d != nil {
		aslu.SetAppSpotUnits(*d)
	}
	return aslu
}

// ClearAppSpotUnits clears the value of the "app_spot_units" field.
func (aslu *AppStockLockUpdate) ClearAppSpotUnits() *AppStockLockUpdate {
	aslu.mutation.ClearAppSpotUnits()
	return aslu
}

// SetLockState sets the "lock_state" field.
func (aslu *AppStockLockUpdate) SetLockState(s string) *AppStockLockUpdate {
	aslu.mutation.SetLockState(s)
	return aslu
}

// SetNillableLockState sets the "lock_state" field if the given value is not nil.
func (aslu *AppStockLockUpdate) SetNillableLockState(s *string) *AppStockLockUpdate {
	if s != nil {
		aslu.SetLockState(*s)
	}
	return aslu
}

// ClearLockState clears the value of the "lock_state" field.
func (aslu *AppStockLockUpdate) ClearLockState() *AppStockLockUpdate {
	aslu.mutation.ClearLockState()
	return aslu
}

// SetChargeBackState sets the "charge_back_state" field.
func (aslu *AppStockLockUpdate) SetChargeBackState(s string) *AppStockLockUpdate {
	aslu.mutation.SetChargeBackState(s)
	return aslu
}

// SetNillableChargeBackState sets the "charge_back_state" field if the given value is not nil.
func (aslu *AppStockLockUpdate) SetNillableChargeBackState(s *string) *AppStockLockUpdate {
	if s != nil {
		aslu.SetChargeBackState(*s)
	}
	return aslu
}

// ClearChargeBackState clears the value of the "charge_back_state" field.
func (aslu *AppStockLockUpdate) ClearChargeBackState() *AppStockLockUpdate {
	aslu.mutation.ClearChargeBackState()
	return aslu
}

// Mutation returns the AppStockLockMutation object of the builder.
func (aslu *AppStockLockUpdate) Mutation() *AppStockLockMutation {
	return aslu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aslu *AppStockLockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := aslu.defaults(); err != nil {
		return 0, err
	}
	if len(aslu.hooks) == 0 {
		affected, err = aslu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppStockLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aslu.mutation = mutation
			affected, err = aslu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aslu.hooks) - 1; i >= 0; i-- {
			if aslu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aslu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aslu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aslu *AppStockLockUpdate) SaveX(ctx context.Context) int {
	affected, err := aslu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aslu *AppStockLockUpdate) Exec(ctx context.Context) error {
	_, err := aslu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aslu *AppStockLockUpdate) ExecX(ctx context.Context) {
	if err := aslu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aslu *AppStockLockUpdate) defaults() error {
	if _, ok := aslu.mutation.UpdatedAt(); !ok {
		if appstocklock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appstocklock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appstocklock.UpdateDefaultUpdatedAt()
		aslu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aslu *AppStockLockUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppStockLockUpdate {
	aslu.modifiers = append(aslu.modifiers, modifiers...)
	return aslu
}

func (aslu *AppStockLockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appstocklock.Table,
			Columns: appstocklock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appstocklock.FieldID,
			},
		},
	}
	if ps := aslu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aslu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldCreatedAt,
		})
	}
	if value, ok := aslu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldCreatedAt,
		})
	}
	if value, ok := aslu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldUpdatedAt,
		})
	}
	if value, ok := aslu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldUpdatedAt,
		})
	}
	if value, ok := aslu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldDeletedAt,
		})
	}
	if value, ok := aslu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldDeletedAt,
		})
	}
	if value, ok := aslu.mutation.AppStockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appstocklock.FieldAppStockID,
		})
	}
	if aslu.mutation.AppStockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appstocklock.FieldAppStockID,
		})
	}
	if value, ok := aslu.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appstocklock.FieldUnits,
		})
	}
	if aslu.mutation.UnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appstocklock.FieldUnits,
		})
	}
	if value, ok := aslu.mutation.AppSpotUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appstocklock.FieldAppSpotUnits,
		})
	}
	if aslu.mutation.AppSpotUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appstocklock.FieldAppSpotUnits,
		})
	}
	if value, ok := aslu.mutation.LockState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appstocklock.FieldLockState,
		})
	}
	if aslu.mutation.LockStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appstocklock.FieldLockState,
		})
	}
	if value, ok := aslu.mutation.ChargeBackState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appstocklock.FieldChargeBackState,
		})
	}
	if aslu.mutation.ChargeBackStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appstocklock.FieldChargeBackState,
		})
	}
	_spec.Modifiers = aslu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, aslu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appstocklock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppStockLockUpdateOne is the builder for updating a single AppStockLock entity.
type AppStockLockUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppStockLockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (asluo *AppStockLockUpdateOne) SetCreatedAt(u uint32) *AppStockLockUpdateOne {
	asluo.mutation.ResetCreatedAt()
	asluo.mutation.SetCreatedAt(u)
	return asluo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asluo *AppStockLockUpdateOne) SetNillableCreatedAt(u *uint32) *AppStockLockUpdateOne {
	if u != nil {
		asluo.SetCreatedAt(*u)
	}
	return asluo
}

// AddCreatedAt adds u to the "created_at" field.
func (asluo *AppStockLockUpdateOne) AddCreatedAt(u int32) *AppStockLockUpdateOne {
	asluo.mutation.AddCreatedAt(u)
	return asluo
}

// SetUpdatedAt sets the "updated_at" field.
func (asluo *AppStockLockUpdateOne) SetUpdatedAt(u uint32) *AppStockLockUpdateOne {
	asluo.mutation.ResetUpdatedAt()
	asluo.mutation.SetUpdatedAt(u)
	return asluo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (asluo *AppStockLockUpdateOne) AddUpdatedAt(u int32) *AppStockLockUpdateOne {
	asluo.mutation.AddUpdatedAt(u)
	return asluo
}

// SetDeletedAt sets the "deleted_at" field.
func (asluo *AppStockLockUpdateOne) SetDeletedAt(u uint32) *AppStockLockUpdateOne {
	asluo.mutation.ResetDeletedAt()
	asluo.mutation.SetDeletedAt(u)
	return asluo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (asluo *AppStockLockUpdateOne) SetNillableDeletedAt(u *uint32) *AppStockLockUpdateOne {
	if u != nil {
		asluo.SetDeletedAt(*u)
	}
	return asluo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (asluo *AppStockLockUpdateOne) AddDeletedAt(u int32) *AppStockLockUpdateOne {
	asluo.mutation.AddDeletedAt(u)
	return asluo
}

// SetAppStockID sets the "app_stock_id" field.
func (asluo *AppStockLockUpdateOne) SetAppStockID(u uuid.UUID) *AppStockLockUpdateOne {
	asluo.mutation.SetAppStockID(u)
	return asluo
}

// SetNillableAppStockID sets the "app_stock_id" field if the given value is not nil.
func (asluo *AppStockLockUpdateOne) SetNillableAppStockID(u *uuid.UUID) *AppStockLockUpdateOne {
	if u != nil {
		asluo.SetAppStockID(*u)
	}
	return asluo
}

// ClearAppStockID clears the value of the "app_stock_id" field.
func (asluo *AppStockLockUpdateOne) ClearAppStockID() *AppStockLockUpdateOne {
	asluo.mutation.ClearAppStockID()
	return asluo
}

// SetUnits sets the "units" field.
func (asluo *AppStockLockUpdateOne) SetUnits(d decimal.Decimal) *AppStockLockUpdateOne {
	asluo.mutation.SetUnits(d)
	return asluo
}

// SetNillableUnits sets the "units" field if the given value is not nil.
func (asluo *AppStockLockUpdateOne) SetNillableUnits(d *decimal.Decimal) *AppStockLockUpdateOne {
	if d != nil {
		asluo.SetUnits(*d)
	}
	return asluo
}

// ClearUnits clears the value of the "units" field.
func (asluo *AppStockLockUpdateOne) ClearUnits() *AppStockLockUpdateOne {
	asluo.mutation.ClearUnits()
	return asluo
}

// SetAppSpotUnits sets the "app_spot_units" field.
func (asluo *AppStockLockUpdateOne) SetAppSpotUnits(d decimal.Decimal) *AppStockLockUpdateOne {
	asluo.mutation.SetAppSpotUnits(d)
	return asluo
}

// SetNillableAppSpotUnits sets the "app_spot_units" field if the given value is not nil.
func (asluo *AppStockLockUpdateOne) SetNillableAppSpotUnits(d *decimal.Decimal) *AppStockLockUpdateOne {
	if d != nil {
		asluo.SetAppSpotUnits(*d)
	}
	return asluo
}

// ClearAppSpotUnits clears the value of the "app_spot_units" field.
func (asluo *AppStockLockUpdateOne) ClearAppSpotUnits() *AppStockLockUpdateOne {
	asluo.mutation.ClearAppSpotUnits()
	return asluo
}

// SetLockState sets the "lock_state" field.
func (asluo *AppStockLockUpdateOne) SetLockState(s string) *AppStockLockUpdateOne {
	asluo.mutation.SetLockState(s)
	return asluo
}

// SetNillableLockState sets the "lock_state" field if the given value is not nil.
func (asluo *AppStockLockUpdateOne) SetNillableLockState(s *string) *AppStockLockUpdateOne {
	if s != nil {
		asluo.SetLockState(*s)
	}
	return asluo
}

// ClearLockState clears the value of the "lock_state" field.
func (asluo *AppStockLockUpdateOne) ClearLockState() *AppStockLockUpdateOne {
	asluo.mutation.ClearLockState()
	return asluo
}

// SetChargeBackState sets the "charge_back_state" field.
func (asluo *AppStockLockUpdateOne) SetChargeBackState(s string) *AppStockLockUpdateOne {
	asluo.mutation.SetChargeBackState(s)
	return asluo
}

// SetNillableChargeBackState sets the "charge_back_state" field if the given value is not nil.
func (asluo *AppStockLockUpdateOne) SetNillableChargeBackState(s *string) *AppStockLockUpdateOne {
	if s != nil {
		asluo.SetChargeBackState(*s)
	}
	return asluo
}

// ClearChargeBackState clears the value of the "charge_back_state" field.
func (asluo *AppStockLockUpdateOne) ClearChargeBackState() *AppStockLockUpdateOne {
	asluo.mutation.ClearChargeBackState()
	return asluo
}

// Mutation returns the AppStockLockMutation object of the builder.
func (asluo *AppStockLockUpdateOne) Mutation() *AppStockLockMutation {
	return asluo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (asluo *AppStockLockUpdateOne) Select(field string, fields ...string) *AppStockLockUpdateOne {
	asluo.fields = append([]string{field}, fields...)
	return asluo
}

// Save executes the query and returns the updated AppStockLock entity.
func (asluo *AppStockLockUpdateOne) Save(ctx context.Context) (*AppStockLock, error) {
	var (
		err  error
		node *AppStockLock
	)
	if err := asluo.defaults(); err != nil {
		return nil, err
	}
	if len(asluo.hooks) == 0 {
		node, err = asluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppStockLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asluo.mutation = mutation
			node, err = asluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(asluo.hooks) - 1; i >= 0; i-- {
			if asluo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asluo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, asluo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppStockLock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppStockLockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (asluo *AppStockLockUpdateOne) SaveX(ctx context.Context) *AppStockLock {
	node, err := asluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (asluo *AppStockLockUpdateOne) Exec(ctx context.Context) error {
	_, err := asluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asluo *AppStockLockUpdateOne) ExecX(ctx context.Context) {
	if err := asluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asluo *AppStockLockUpdateOne) defaults() error {
	if _, ok := asluo.mutation.UpdatedAt(); !ok {
		if appstocklock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appstocklock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appstocklock.UpdateDefaultUpdatedAt()
		asluo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (asluo *AppStockLockUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppStockLockUpdateOne {
	asluo.modifiers = append(asluo.modifiers, modifiers...)
	return asluo
}

func (asluo *AppStockLockUpdateOne) sqlSave(ctx context.Context) (_node *AppStockLock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appstocklock.Table,
			Columns: appstocklock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appstocklock.FieldID,
			},
		},
	}
	id, ok := asluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppStockLock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := asluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appstocklock.FieldID)
		for _, f := range fields {
			if !appstocklock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appstocklock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := asluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := asluo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldCreatedAt,
		})
	}
	if value, ok := asluo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldCreatedAt,
		})
	}
	if value, ok := asluo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldUpdatedAt,
		})
	}
	if value, ok := asluo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldUpdatedAt,
		})
	}
	if value, ok := asluo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldDeletedAt,
		})
	}
	if value, ok := asluo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldDeletedAt,
		})
	}
	if value, ok := asluo.mutation.AppStockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appstocklock.FieldAppStockID,
		})
	}
	if asluo.mutation.AppStockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appstocklock.FieldAppStockID,
		})
	}
	if value, ok := asluo.mutation.Units(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appstocklock.FieldUnits,
		})
	}
	if asluo.mutation.UnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appstocklock.FieldUnits,
		})
	}
	if value, ok := asluo.mutation.AppSpotUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appstocklock.FieldAppSpotUnits,
		})
	}
	if asluo.mutation.AppSpotUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appstocklock.FieldAppSpotUnits,
		})
	}
	if value, ok := asluo.mutation.LockState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appstocklock.FieldLockState,
		})
	}
	if asluo.mutation.LockStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appstocklock.FieldLockState,
		})
	}
	if value, ok := asluo.mutation.ChargeBackState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appstocklock.FieldChargeBackState,
		})
	}
	if asluo.mutation.ChargeBackStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appstocklock.FieldChargeBackState,
		})
	}
	_spec.Modifiers = asluo.modifiers
	_node = &AppStockLock{config: asluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, asluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appstocklock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
