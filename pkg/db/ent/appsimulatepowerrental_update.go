// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulatepowerrental"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppSimulatePowerRentalUpdate is the builder for updating AppSimulatePowerRental entities.
type AppSimulatePowerRentalUpdate struct {
	config
	hooks     []Hook
	mutation  *AppSimulatePowerRentalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppSimulatePowerRentalUpdate builder.
func (aspru *AppSimulatePowerRentalUpdate) Where(ps ...predicate.AppSimulatePowerRental) *AppSimulatePowerRentalUpdate {
	aspru.mutation.Where(ps...)
	return aspru
}

// SetCreatedAt sets the "created_at" field.
func (aspru *AppSimulatePowerRentalUpdate) SetCreatedAt(u uint32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.ResetCreatedAt()
	aspru.mutation.SetCreatedAt(u)
	return aspru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aspru *AppSimulatePowerRentalUpdate) SetNillableCreatedAt(u *uint32) *AppSimulatePowerRentalUpdate {
	if u != nil {
		aspru.SetCreatedAt(*u)
	}
	return aspru
}

// AddCreatedAt adds u to the "created_at" field.
func (aspru *AppSimulatePowerRentalUpdate) AddCreatedAt(u int32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.AddCreatedAt(u)
	return aspru
}

// SetUpdatedAt sets the "updated_at" field.
func (aspru *AppSimulatePowerRentalUpdate) SetUpdatedAt(u uint32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.ResetUpdatedAt()
	aspru.mutation.SetUpdatedAt(u)
	return aspru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (aspru *AppSimulatePowerRentalUpdate) AddUpdatedAt(u int32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.AddUpdatedAt(u)
	return aspru
}

// SetDeletedAt sets the "deleted_at" field.
func (aspru *AppSimulatePowerRentalUpdate) SetDeletedAt(u uint32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.ResetDeletedAt()
	aspru.mutation.SetDeletedAt(u)
	return aspru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aspru *AppSimulatePowerRentalUpdate) SetNillableDeletedAt(u *uint32) *AppSimulatePowerRentalUpdate {
	if u != nil {
		aspru.SetDeletedAt(*u)
	}
	return aspru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (aspru *AppSimulatePowerRentalUpdate) AddDeletedAt(u int32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.AddDeletedAt(u)
	return aspru
}

// SetEntID sets the "ent_id" field.
func (aspru *AppSimulatePowerRentalUpdate) SetEntID(u uuid.UUID) *AppSimulatePowerRentalUpdate {
	aspru.mutation.SetEntID(u)
	return aspru
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (aspru *AppSimulatePowerRentalUpdate) SetNillableEntID(u *uuid.UUID) *AppSimulatePowerRentalUpdate {
	if u != nil {
		aspru.SetEntID(*u)
	}
	return aspru
}

// SetAppGoodID sets the "app_good_id" field.
func (aspru *AppSimulatePowerRentalUpdate) SetAppGoodID(u uuid.UUID) *AppSimulatePowerRentalUpdate {
	aspru.mutation.SetAppGoodID(u)
	return aspru
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (aspru *AppSimulatePowerRentalUpdate) SetNillableAppGoodID(u *uuid.UUID) *AppSimulatePowerRentalUpdate {
	if u != nil {
		aspru.SetAppGoodID(*u)
	}
	return aspru
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (aspru *AppSimulatePowerRentalUpdate) ClearAppGoodID() *AppSimulatePowerRentalUpdate {
	aspru.mutation.ClearAppGoodID()
	return aspru
}

// SetCoinTypeID sets the "coin_type_id" field.
func (aspru *AppSimulatePowerRentalUpdate) SetCoinTypeID(u uuid.UUID) *AppSimulatePowerRentalUpdate {
	aspru.mutation.SetCoinTypeID(u)
	return aspru
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (aspru *AppSimulatePowerRentalUpdate) SetNillableCoinTypeID(u *uuid.UUID) *AppSimulatePowerRentalUpdate {
	if u != nil {
		aspru.SetCoinTypeID(*u)
	}
	return aspru
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (aspru *AppSimulatePowerRentalUpdate) ClearCoinTypeID() *AppSimulatePowerRentalUpdate {
	aspru.mutation.ClearCoinTypeID()
	return aspru
}

// SetOrderUnits sets the "order_units" field.
func (aspru *AppSimulatePowerRentalUpdate) SetOrderUnits(d decimal.Decimal) *AppSimulatePowerRentalUpdate {
	aspru.mutation.SetOrderUnits(d)
	return aspru
}

// SetNillableOrderUnits sets the "order_units" field if the given value is not nil.
func (aspru *AppSimulatePowerRentalUpdate) SetNillableOrderUnits(d *decimal.Decimal) *AppSimulatePowerRentalUpdate {
	if d != nil {
		aspru.SetOrderUnits(*d)
	}
	return aspru
}

// ClearOrderUnits clears the value of the "order_units" field.
func (aspru *AppSimulatePowerRentalUpdate) ClearOrderUnits() *AppSimulatePowerRentalUpdate {
	aspru.mutation.ClearOrderUnits()
	return aspru
}

// SetOrderDurationSeconds sets the "order_duration_seconds" field.
func (aspru *AppSimulatePowerRentalUpdate) SetOrderDurationSeconds(u uint32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.ResetOrderDurationSeconds()
	aspru.mutation.SetOrderDurationSeconds(u)
	return aspru
}

// SetNillableOrderDurationSeconds sets the "order_duration_seconds" field if the given value is not nil.
func (aspru *AppSimulatePowerRentalUpdate) SetNillableOrderDurationSeconds(u *uint32) *AppSimulatePowerRentalUpdate {
	if u != nil {
		aspru.SetOrderDurationSeconds(*u)
	}
	return aspru
}

// AddOrderDurationSeconds adds u to the "order_duration_seconds" field.
func (aspru *AppSimulatePowerRentalUpdate) AddOrderDurationSeconds(u int32) *AppSimulatePowerRentalUpdate {
	aspru.mutation.AddOrderDurationSeconds(u)
	return aspru
}

// ClearOrderDurationSeconds clears the value of the "order_duration_seconds" field.
func (aspru *AppSimulatePowerRentalUpdate) ClearOrderDurationSeconds() *AppSimulatePowerRentalUpdate {
	aspru.mutation.ClearOrderDurationSeconds()
	return aspru
}

// Mutation returns the AppSimulatePowerRentalMutation object of the builder.
func (aspru *AppSimulatePowerRentalUpdate) Mutation() *AppSimulatePowerRentalMutation {
	return aspru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (aspru *AppSimulatePowerRentalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := aspru.defaults(); err != nil {
		return 0, err
	}
	if len(aspru.hooks) == 0 {
		affected, err = aspru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSimulatePowerRentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aspru.mutation = mutation
			affected, err = aspru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(aspru.hooks) - 1; i >= 0; i-- {
			if aspru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aspru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, aspru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (aspru *AppSimulatePowerRentalUpdate) SaveX(ctx context.Context) int {
	affected, err := aspru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (aspru *AppSimulatePowerRentalUpdate) Exec(ctx context.Context) error {
	_, err := aspru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aspru *AppSimulatePowerRentalUpdate) ExecX(ctx context.Context) {
	if err := aspru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aspru *AppSimulatePowerRentalUpdate) defaults() error {
	if _, ok := aspru.mutation.UpdatedAt(); !ok {
		if appsimulatepowerrental.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsimulatepowerrental.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appsimulatepowerrental.UpdateDefaultUpdatedAt()
		aspru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aspru *AppSimulatePowerRentalUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppSimulatePowerRentalUpdate {
	aspru.modifiers = append(aspru.modifiers, modifiers...)
	return aspru
}

func (aspru *AppSimulatePowerRentalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appsimulatepowerrental.Table,
			Columns: appsimulatepowerrental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appsimulatepowerrental.FieldID,
			},
		},
	}
	if ps := aspru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aspru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldCreatedAt,
		})
	}
	if value, ok := aspru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldCreatedAt,
		})
	}
	if value, ok := aspru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := aspru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := aspru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldDeletedAt,
		})
	}
	if value, ok := aspru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldDeletedAt,
		})
	}
	if value, ok := aspru.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldEntID,
		})
	}
	if value, ok := aspru.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldAppGoodID,
		})
	}
	if aspru.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appsimulatepowerrental.FieldAppGoodID,
		})
	}
	if value, ok := aspru.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldCoinTypeID,
		})
	}
	if aspru.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appsimulatepowerrental.FieldCoinTypeID,
		})
	}
	if value, ok := aspru.mutation.OrderUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderUnits,
		})
	}
	if aspru.mutation.OrderUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appsimulatepowerrental.FieldOrderUnits,
		})
	}
	if value, ok := aspru.mutation.OrderDurationSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderDurationSeconds,
		})
	}
	if value, ok := aspru.mutation.AddedOrderDurationSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderDurationSeconds,
		})
	}
	if aspru.mutation.OrderDurationSecondsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: appsimulatepowerrental.FieldOrderDurationSeconds,
		})
	}
	_spec.Modifiers = aspru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, aspru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appsimulatepowerrental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppSimulatePowerRentalUpdateOne is the builder for updating a single AppSimulatePowerRental entity.
type AppSimulatePowerRentalUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppSimulatePowerRentalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetCreatedAt(u uint32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ResetCreatedAt()
	aspruo.mutation.SetCreatedAt(u)
	return aspruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetNillableCreatedAt(u *uint32) *AppSimulatePowerRentalUpdateOne {
	if u != nil {
		aspruo.SetCreatedAt(*u)
	}
	return aspruo
}

// AddCreatedAt adds u to the "created_at" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) AddCreatedAt(u int32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.AddCreatedAt(u)
	return aspruo
}

// SetUpdatedAt sets the "updated_at" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetUpdatedAt(u uint32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ResetUpdatedAt()
	aspruo.mutation.SetUpdatedAt(u)
	return aspruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) AddUpdatedAt(u int32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.AddUpdatedAt(u)
	return aspruo
}

// SetDeletedAt sets the "deleted_at" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetDeletedAt(u uint32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ResetDeletedAt()
	aspruo.mutation.SetDeletedAt(u)
	return aspruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetNillableDeletedAt(u *uint32) *AppSimulatePowerRentalUpdateOne {
	if u != nil {
		aspruo.SetDeletedAt(*u)
	}
	return aspruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) AddDeletedAt(u int32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.AddDeletedAt(u)
	return aspruo
}

// SetEntID sets the "ent_id" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetEntID(u uuid.UUID) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.SetEntID(u)
	return aspruo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetNillableEntID(u *uuid.UUID) *AppSimulatePowerRentalUpdateOne {
	if u != nil {
		aspruo.SetEntID(*u)
	}
	return aspruo
}

// SetAppGoodID sets the "app_good_id" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetAppGoodID(u uuid.UUID) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.SetAppGoodID(u)
	return aspruo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *AppSimulatePowerRentalUpdateOne {
	if u != nil {
		aspruo.SetAppGoodID(*u)
	}
	return aspruo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) ClearAppGoodID() *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ClearAppGoodID()
	return aspruo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetCoinTypeID(u uuid.UUID) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.SetCoinTypeID(u)
	return aspruo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *AppSimulatePowerRentalUpdateOne {
	if u != nil {
		aspruo.SetCoinTypeID(*u)
	}
	return aspruo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) ClearCoinTypeID() *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ClearCoinTypeID()
	return aspruo
}

// SetOrderUnits sets the "order_units" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetOrderUnits(d decimal.Decimal) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.SetOrderUnits(d)
	return aspruo
}

// SetNillableOrderUnits sets the "order_units" field if the given value is not nil.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetNillableOrderUnits(d *decimal.Decimal) *AppSimulatePowerRentalUpdateOne {
	if d != nil {
		aspruo.SetOrderUnits(*d)
	}
	return aspruo
}

// ClearOrderUnits clears the value of the "order_units" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) ClearOrderUnits() *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ClearOrderUnits()
	return aspruo
}

// SetOrderDurationSeconds sets the "order_duration_seconds" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetOrderDurationSeconds(u uint32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ResetOrderDurationSeconds()
	aspruo.mutation.SetOrderDurationSeconds(u)
	return aspruo
}

// SetNillableOrderDurationSeconds sets the "order_duration_seconds" field if the given value is not nil.
func (aspruo *AppSimulatePowerRentalUpdateOne) SetNillableOrderDurationSeconds(u *uint32) *AppSimulatePowerRentalUpdateOne {
	if u != nil {
		aspruo.SetOrderDurationSeconds(*u)
	}
	return aspruo
}

// AddOrderDurationSeconds adds u to the "order_duration_seconds" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) AddOrderDurationSeconds(u int32) *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.AddOrderDurationSeconds(u)
	return aspruo
}

// ClearOrderDurationSeconds clears the value of the "order_duration_seconds" field.
func (aspruo *AppSimulatePowerRentalUpdateOne) ClearOrderDurationSeconds() *AppSimulatePowerRentalUpdateOne {
	aspruo.mutation.ClearOrderDurationSeconds()
	return aspruo
}

// Mutation returns the AppSimulatePowerRentalMutation object of the builder.
func (aspruo *AppSimulatePowerRentalUpdateOne) Mutation() *AppSimulatePowerRentalMutation {
	return aspruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (aspruo *AppSimulatePowerRentalUpdateOne) Select(field string, fields ...string) *AppSimulatePowerRentalUpdateOne {
	aspruo.fields = append([]string{field}, fields...)
	return aspruo
}

// Save executes the query and returns the updated AppSimulatePowerRental entity.
func (aspruo *AppSimulatePowerRentalUpdateOne) Save(ctx context.Context) (*AppSimulatePowerRental, error) {
	var (
		err  error
		node *AppSimulatePowerRental
	)
	if err := aspruo.defaults(); err != nil {
		return nil, err
	}
	if len(aspruo.hooks) == 0 {
		node, err = aspruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSimulatePowerRentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			aspruo.mutation = mutation
			node, err = aspruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(aspruo.hooks) - 1; i >= 0; i-- {
			if aspruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aspruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aspruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppSimulatePowerRental)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppSimulatePowerRentalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (aspruo *AppSimulatePowerRentalUpdateOne) SaveX(ctx context.Context) *AppSimulatePowerRental {
	node, err := aspruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (aspruo *AppSimulatePowerRentalUpdateOne) Exec(ctx context.Context) error {
	_, err := aspruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aspruo *AppSimulatePowerRentalUpdateOne) ExecX(ctx context.Context) {
	if err := aspruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aspruo *AppSimulatePowerRentalUpdateOne) defaults() error {
	if _, ok := aspruo.mutation.UpdatedAt(); !ok {
		if appsimulatepowerrental.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsimulatepowerrental.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appsimulatepowerrental.UpdateDefaultUpdatedAt()
		aspruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (aspruo *AppSimulatePowerRentalUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppSimulatePowerRentalUpdateOne {
	aspruo.modifiers = append(aspruo.modifiers, modifiers...)
	return aspruo
}

func (aspruo *AppSimulatePowerRentalUpdateOne) sqlSave(ctx context.Context) (_node *AppSimulatePowerRental, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appsimulatepowerrental.Table,
			Columns: appsimulatepowerrental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appsimulatepowerrental.FieldID,
			},
		},
	}
	id, ok := aspruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppSimulatePowerRental.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := aspruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appsimulatepowerrental.FieldID)
		for _, f := range fields {
			if !appsimulatepowerrental.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appsimulatepowerrental.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := aspruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := aspruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldCreatedAt,
		})
	}
	if value, ok := aspruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldCreatedAt,
		})
	}
	if value, ok := aspruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := aspruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := aspruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldDeletedAt,
		})
	}
	if value, ok := aspruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldDeletedAt,
		})
	}
	if value, ok := aspruo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldEntID,
		})
	}
	if value, ok := aspruo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldAppGoodID,
		})
	}
	if aspruo.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appsimulatepowerrental.FieldAppGoodID,
		})
	}
	if value, ok := aspruo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldCoinTypeID,
		})
	}
	if aspruo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appsimulatepowerrental.FieldCoinTypeID,
		})
	}
	if value, ok := aspruo.mutation.OrderUnits(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderUnits,
		})
	}
	if aspruo.mutation.OrderUnitsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appsimulatepowerrental.FieldOrderUnits,
		})
	}
	if value, ok := aspruo.mutation.OrderDurationSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderDurationSeconds,
		})
	}
	if value, ok := aspruo.mutation.AddedOrderDurationSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderDurationSeconds,
		})
	}
	if aspruo.mutation.OrderDurationSecondsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: appsimulatepowerrental.FieldOrderDurationSeconds,
		})
	}
	_spec.Modifiers = aspruo.modifiers
	_node = &AppSimulatePowerRental{config: aspruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, aspruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appsimulatepowerrental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
