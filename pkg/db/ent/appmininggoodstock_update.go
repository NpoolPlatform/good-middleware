// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appmininggoodstock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppMiningGoodStockUpdate is the builder for updating AppMiningGoodStock entities.
type AppMiningGoodStockUpdate struct {
	config
	hooks     []Hook
	mutation  *AppMiningGoodStockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppMiningGoodStockUpdate builder.
func (amgsu *AppMiningGoodStockUpdate) Where(ps ...predicate.AppMiningGoodStock) *AppMiningGoodStockUpdate {
	amgsu.mutation.Where(ps...)
	return amgsu
}

// SetCreatedAt sets the "created_at" field.
func (amgsu *AppMiningGoodStockUpdate) SetCreatedAt(u uint32) *AppMiningGoodStockUpdate {
	amgsu.mutation.ResetCreatedAt()
	amgsu.mutation.SetCreatedAt(u)
	return amgsu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableCreatedAt(u *uint32) *AppMiningGoodStockUpdate {
	if u != nil {
		amgsu.SetCreatedAt(*u)
	}
	return amgsu
}

// AddCreatedAt adds u to the "created_at" field.
func (amgsu *AppMiningGoodStockUpdate) AddCreatedAt(u int32) *AppMiningGoodStockUpdate {
	amgsu.mutation.AddCreatedAt(u)
	return amgsu
}

// SetUpdatedAt sets the "updated_at" field.
func (amgsu *AppMiningGoodStockUpdate) SetUpdatedAt(u uint32) *AppMiningGoodStockUpdate {
	amgsu.mutation.ResetUpdatedAt()
	amgsu.mutation.SetUpdatedAt(u)
	return amgsu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (amgsu *AppMiningGoodStockUpdate) AddUpdatedAt(u int32) *AppMiningGoodStockUpdate {
	amgsu.mutation.AddUpdatedAt(u)
	return amgsu
}

// SetDeletedAt sets the "deleted_at" field.
func (amgsu *AppMiningGoodStockUpdate) SetDeletedAt(u uint32) *AppMiningGoodStockUpdate {
	amgsu.mutation.ResetDeletedAt()
	amgsu.mutation.SetDeletedAt(u)
	return amgsu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableDeletedAt(u *uint32) *AppMiningGoodStockUpdate {
	if u != nil {
		amgsu.SetDeletedAt(*u)
	}
	return amgsu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (amgsu *AppMiningGoodStockUpdate) AddDeletedAt(u int32) *AppMiningGoodStockUpdate {
	amgsu.mutation.AddDeletedAt(u)
	return amgsu
}

// SetEntID sets the "ent_id" field.
func (amgsu *AppMiningGoodStockUpdate) SetEntID(u uuid.UUID) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetEntID(u)
	return amgsu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableEntID(u *uuid.UUID) *AppMiningGoodStockUpdate {
	if u != nil {
		amgsu.SetEntID(*u)
	}
	return amgsu
}

// SetAppGoodStockID sets the "app_good_stock_id" field.
func (amgsu *AppMiningGoodStockUpdate) SetAppGoodStockID(u uuid.UUID) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetAppGoodStockID(u)
	return amgsu
}

// SetNillableAppGoodStockID sets the "app_good_stock_id" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableAppGoodStockID(u *uuid.UUID) *AppMiningGoodStockUpdate {
	if u != nil {
		amgsu.SetAppGoodStockID(*u)
	}
	return amgsu
}

// ClearAppGoodStockID clears the value of the "app_good_stock_id" field.
func (amgsu *AppMiningGoodStockUpdate) ClearAppGoodStockID() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearAppGoodStockID()
	return amgsu
}

// SetMiningGoodStockID sets the "mining_good_stock_id" field.
func (amgsu *AppMiningGoodStockUpdate) SetMiningGoodStockID(u uuid.UUID) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetMiningGoodStockID(u)
	return amgsu
}

// SetNillableMiningGoodStockID sets the "mining_good_stock_id" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableMiningGoodStockID(u *uuid.UUID) *AppMiningGoodStockUpdate {
	if u != nil {
		amgsu.SetMiningGoodStockID(*u)
	}
	return amgsu
}

// ClearMiningGoodStockID clears the value of the "mining_good_stock_id" field.
func (amgsu *AppMiningGoodStockUpdate) ClearMiningGoodStockID() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearMiningGoodStockID()
	return amgsu
}

// SetReserved sets the "reserved" field.
func (amgsu *AppMiningGoodStockUpdate) SetReserved(d decimal.Decimal) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetReserved(d)
	return amgsu
}

// SetNillableReserved sets the "reserved" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableReserved(d *decimal.Decimal) *AppMiningGoodStockUpdate {
	if d != nil {
		amgsu.SetReserved(*d)
	}
	return amgsu
}

// ClearReserved clears the value of the "reserved" field.
func (amgsu *AppMiningGoodStockUpdate) ClearReserved() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearReserved()
	return amgsu
}

// SetSpotQuantity sets the "spot_quantity" field.
func (amgsu *AppMiningGoodStockUpdate) SetSpotQuantity(d decimal.Decimal) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetSpotQuantity(d)
	return amgsu
}

// SetNillableSpotQuantity sets the "spot_quantity" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableSpotQuantity(d *decimal.Decimal) *AppMiningGoodStockUpdate {
	if d != nil {
		amgsu.SetSpotQuantity(*d)
	}
	return amgsu
}

// ClearSpotQuantity clears the value of the "spot_quantity" field.
func (amgsu *AppMiningGoodStockUpdate) ClearSpotQuantity() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearSpotQuantity()
	return amgsu
}

// SetLocked sets the "locked" field.
func (amgsu *AppMiningGoodStockUpdate) SetLocked(d decimal.Decimal) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetLocked(d)
	return amgsu
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableLocked(d *decimal.Decimal) *AppMiningGoodStockUpdate {
	if d != nil {
		amgsu.SetLocked(*d)
	}
	return amgsu
}

// ClearLocked clears the value of the "locked" field.
func (amgsu *AppMiningGoodStockUpdate) ClearLocked() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearLocked()
	return amgsu
}

// SetInService sets the "in_service" field.
func (amgsu *AppMiningGoodStockUpdate) SetInService(d decimal.Decimal) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetInService(d)
	return amgsu
}

// SetNillableInService sets the "in_service" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableInService(d *decimal.Decimal) *AppMiningGoodStockUpdate {
	if d != nil {
		amgsu.SetInService(*d)
	}
	return amgsu
}

// ClearInService clears the value of the "in_service" field.
func (amgsu *AppMiningGoodStockUpdate) ClearInService() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearInService()
	return amgsu
}

// SetWaitStart sets the "wait_start" field.
func (amgsu *AppMiningGoodStockUpdate) SetWaitStart(d decimal.Decimal) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetWaitStart(d)
	return amgsu
}

// SetNillableWaitStart sets the "wait_start" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableWaitStart(d *decimal.Decimal) *AppMiningGoodStockUpdate {
	if d != nil {
		amgsu.SetWaitStart(*d)
	}
	return amgsu
}

// ClearWaitStart clears the value of the "wait_start" field.
func (amgsu *AppMiningGoodStockUpdate) ClearWaitStart() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearWaitStart()
	return amgsu
}

// SetSold sets the "sold" field.
func (amgsu *AppMiningGoodStockUpdate) SetSold(d decimal.Decimal) *AppMiningGoodStockUpdate {
	amgsu.mutation.SetSold(d)
	return amgsu
}

// SetNillableSold sets the "sold" field if the given value is not nil.
func (amgsu *AppMiningGoodStockUpdate) SetNillableSold(d *decimal.Decimal) *AppMiningGoodStockUpdate {
	if d != nil {
		amgsu.SetSold(*d)
	}
	return amgsu
}

// ClearSold clears the value of the "sold" field.
func (amgsu *AppMiningGoodStockUpdate) ClearSold() *AppMiningGoodStockUpdate {
	amgsu.mutation.ClearSold()
	return amgsu
}

// Mutation returns the AppMiningGoodStockMutation object of the builder.
func (amgsu *AppMiningGoodStockUpdate) Mutation() *AppMiningGoodStockMutation {
	return amgsu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (amgsu *AppMiningGoodStockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := amgsu.defaults(); err != nil {
		return 0, err
	}
	if len(amgsu.hooks) == 0 {
		affected, err = amgsu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMiningGoodStockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amgsu.mutation = mutation
			affected, err = amgsu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(amgsu.hooks) - 1; i >= 0; i-- {
			if amgsu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = amgsu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amgsu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (amgsu *AppMiningGoodStockUpdate) SaveX(ctx context.Context) int {
	affected, err := amgsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (amgsu *AppMiningGoodStockUpdate) Exec(ctx context.Context) error {
	_, err := amgsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amgsu *AppMiningGoodStockUpdate) ExecX(ctx context.Context) {
	if err := amgsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amgsu *AppMiningGoodStockUpdate) defaults() error {
	if _, ok := amgsu.mutation.UpdatedAt(); !ok {
		if appmininggoodstock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appmininggoodstock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appmininggoodstock.UpdateDefaultUpdatedAt()
		amgsu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (amgsu *AppMiningGoodStockUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppMiningGoodStockUpdate {
	amgsu.modifiers = append(amgsu.modifiers, modifiers...)
	return amgsu
}

func (amgsu *AppMiningGoodStockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appmininggoodstock.Table,
			Columns: appmininggoodstock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appmininggoodstock.FieldID,
			},
		},
	}
	if ps := amgsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amgsu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldCreatedAt,
		})
	}
	if value, ok := amgsu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldCreatedAt,
		})
	}
	if value, ok := amgsu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldUpdatedAt,
		})
	}
	if value, ok := amgsu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldUpdatedAt,
		})
	}
	if value, ok := amgsu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldDeletedAt,
		})
	}
	if value, ok := amgsu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldDeletedAt,
		})
	}
	if value, ok := amgsu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appmininggoodstock.FieldEntID,
		})
	}
	if value, ok := amgsu.mutation.AppGoodStockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appmininggoodstock.FieldAppGoodStockID,
		})
	}
	if amgsu.mutation.AppGoodStockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appmininggoodstock.FieldAppGoodStockID,
		})
	}
	if value, ok := amgsu.mutation.MiningGoodStockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appmininggoodstock.FieldMiningGoodStockID,
		})
	}
	if amgsu.mutation.MiningGoodStockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appmininggoodstock.FieldMiningGoodStockID,
		})
	}
	if value, ok := amgsu.mutation.Reserved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldReserved,
		})
	}
	if amgsu.mutation.ReservedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldReserved,
		})
	}
	if value, ok := amgsu.mutation.SpotQuantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldSpotQuantity,
		})
	}
	if amgsu.mutation.SpotQuantityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldSpotQuantity,
		})
	}
	if value, ok := amgsu.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldLocked,
		})
	}
	if amgsu.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldLocked,
		})
	}
	if value, ok := amgsu.mutation.InService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldInService,
		})
	}
	if amgsu.mutation.InServiceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldInService,
		})
	}
	if value, ok := amgsu.mutation.WaitStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldWaitStart,
		})
	}
	if amgsu.mutation.WaitStartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldWaitStart,
		})
	}
	if value, ok := amgsu.mutation.Sold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldSold,
		})
	}
	if amgsu.mutation.SoldCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldSold,
		})
	}
	_spec.Modifiers = amgsu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, amgsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appmininggoodstock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppMiningGoodStockUpdateOne is the builder for updating a single AppMiningGoodStock entity.
type AppMiningGoodStockUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppMiningGoodStockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetCreatedAt(u uint32) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ResetCreatedAt()
	amgsuo.mutation.SetCreatedAt(u)
	return amgsuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableCreatedAt(u *uint32) *AppMiningGoodStockUpdateOne {
	if u != nil {
		amgsuo.SetCreatedAt(*u)
	}
	return amgsuo
}

// AddCreatedAt adds u to the "created_at" field.
func (amgsuo *AppMiningGoodStockUpdateOne) AddCreatedAt(u int32) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.AddCreatedAt(u)
	return amgsuo
}

// SetUpdatedAt sets the "updated_at" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetUpdatedAt(u uint32) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ResetUpdatedAt()
	amgsuo.mutation.SetUpdatedAt(u)
	return amgsuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (amgsuo *AppMiningGoodStockUpdateOne) AddUpdatedAt(u int32) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.AddUpdatedAt(u)
	return amgsuo
}

// SetDeletedAt sets the "deleted_at" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetDeletedAt(u uint32) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ResetDeletedAt()
	amgsuo.mutation.SetDeletedAt(u)
	return amgsuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableDeletedAt(u *uint32) *AppMiningGoodStockUpdateOne {
	if u != nil {
		amgsuo.SetDeletedAt(*u)
	}
	return amgsuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (amgsuo *AppMiningGoodStockUpdateOne) AddDeletedAt(u int32) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.AddDeletedAt(u)
	return amgsuo
}

// SetEntID sets the "ent_id" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetEntID(u uuid.UUID) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetEntID(u)
	return amgsuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableEntID(u *uuid.UUID) *AppMiningGoodStockUpdateOne {
	if u != nil {
		amgsuo.SetEntID(*u)
	}
	return amgsuo
}

// SetAppGoodStockID sets the "app_good_stock_id" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetAppGoodStockID(u uuid.UUID) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetAppGoodStockID(u)
	return amgsuo
}

// SetNillableAppGoodStockID sets the "app_good_stock_id" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableAppGoodStockID(u *uuid.UUID) *AppMiningGoodStockUpdateOne {
	if u != nil {
		amgsuo.SetAppGoodStockID(*u)
	}
	return amgsuo
}

// ClearAppGoodStockID clears the value of the "app_good_stock_id" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearAppGoodStockID() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearAppGoodStockID()
	return amgsuo
}

// SetMiningGoodStockID sets the "mining_good_stock_id" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetMiningGoodStockID(u uuid.UUID) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetMiningGoodStockID(u)
	return amgsuo
}

// SetNillableMiningGoodStockID sets the "mining_good_stock_id" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableMiningGoodStockID(u *uuid.UUID) *AppMiningGoodStockUpdateOne {
	if u != nil {
		amgsuo.SetMiningGoodStockID(*u)
	}
	return amgsuo
}

// ClearMiningGoodStockID clears the value of the "mining_good_stock_id" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearMiningGoodStockID() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearMiningGoodStockID()
	return amgsuo
}

// SetReserved sets the "reserved" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetReserved(d decimal.Decimal) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetReserved(d)
	return amgsuo
}

// SetNillableReserved sets the "reserved" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableReserved(d *decimal.Decimal) *AppMiningGoodStockUpdateOne {
	if d != nil {
		amgsuo.SetReserved(*d)
	}
	return amgsuo
}

// ClearReserved clears the value of the "reserved" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearReserved() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearReserved()
	return amgsuo
}

// SetSpotQuantity sets the "spot_quantity" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetSpotQuantity(d decimal.Decimal) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetSpotQuantity(d)
	return amgsuo
}

// SetNillableSpotQuantity sets the "spot_quantity" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableSpotQuantity(d *decimal.Decimal) *AppMiningGoodStockUpdateOne {
	if d != nil {
		amgsuo.SetSpotQuantity(*d)
	}
	return amgsuo
}

// ClearSpotQuantity clears the value of the "spot_quantity" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearSpotQuantity() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearSpotQuantity()
	return amgsuo
}

// SetLocked sets the "locked" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetLocked(d decimal.Decimal) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetLocked(d)
	return amgsuo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableLocked(d *decimal.Decimal) *AppMiningGoodStockUpdateOne {
	if d != nil {
		amgsuo.SetLocked(*d)
	}
	return amgsuo
}

// ClearLocked clears the value of the "locked" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearLocked() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearLocked()
	return amgsuo
}

// SetInService sets the "in_service" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetInService(d decimal.Decimal) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetInService(d)
	return amgsuo
}

// SetNillableInService sets the "in_service" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableInService(d *decimal.Decimal) *AppMiningGoodStockUpdateOne {
	if d != nil {
		amgsuo.SetInService(*d)
	}
	return amgsuo
}

// ClearInService clears the value of the "in_service" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearInService() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearInService()
	return amgsuo
}

// SetWaitStart sets the "wait_start" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetWaitStart(d decimal.Decimal) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetWaitStart(d)
	return amgsuo
}

// SetNillableWaitStart sets the "wait_start" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableWaitStart(d *decimal.Decimal) *AppMiningGoodStockUpdateOne {
	if d != nil {
		amgsuo.SetWaitStart(*d)
	}
	return amgsuo
}

// ClearWaitStart clears the value of the "wait_start" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearWaitStart() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearWaitStart()
	return amgsuo
}

// SetSold sets the "sold" field.
func (amgsuo *AppMiningGoodStockUpdateOne) SetSold(d decimal.Decimal) *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.SetSold(d)
	return amgsuo
}

// SetNillableSold sets the "sold" field if the given value is not nil.
func (amgsuo *AppMiningGoodStockUpdateOne) SetNillableSold(d *decimal.Decimal) *AppMiningGoodStockUpdateOne {
	if d != nil {
		amgsuo.SetSold(*d)
	}
	return amgsuo
}

// ClearSold clears the value of the "sold" field.
func (amgsuo *AppMiningGoodStockUpdateOne) ClearSold() *AppMiningGoodStockUpdateOne {
	amgsuo.mutation.ClearSold()
	return amgsuo
}

// Mutation returns the AppMiningGoodStockMutation object of the builder.
func (amgsuo *AppMiningGoodStockUpdateOne) Mutation() *AppMiningGoodStockMutation {
	return amgsuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (amgsuo *AppMiningGoodStockUpdateOne) Select(field string, fields ...string) *AppMiningGoodStockUpdateOne {
	amgsuo.fields = append([]string{field}, fields...)
	return amgsuo
}

// Save executes the query and returns the updated AppMiningGoodStock entity.
func (amgsuo *AppMiningGoodStockUpdateOne) Save(ctx context.Context) (*AppMiningGoodStock, error) {
	var (
		err  error
		node *AppMiningGoodStock
	)
	if err := amgsuo.defaults(); err != nil {
		return nil, err
	}
	if len(amgsuo.hooks) == 0 {
		node, err = amgsuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMiningGoodStockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amgsuo.mutation = mutation
			node, err = amgsuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(amgsuo.hooks) - 1; i >= 0; i-- {
			if amgsuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = amgsuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, amgsuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppMiningGoodStock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppMiningGoodStockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (amgsuo *AppMiningGoodStockUpdateOne) SaveX(ctx context.Context) *AppMiningGoodStock {
	node, err := amgsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (amgsuo *AppMiningGoodStockUpdateOne) Exec(ctx context.Context) error {
	_, err := amgsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (amgsuo *AppMiningGoodStockUpdateOne) ExecX(ctx context.Context) {
	if err := amgsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (amgsuo *AppMiningGoodStockUpdateOne) defaults() error {
	if _, ok := amgsuo.mutation.UpdatedAt(); !ok {
		if appmininggoodstock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appmininggoodstock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appmininggoodstock.UpdateDefaultUpdatedAt()
		amgsuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (amgsuo *AppMiningGoodStockUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppMiningGoodStockUpdateOne {
	amgsuo.modifiers = append(amgsuo.modifiers, modifiers...)
	return amgsuo
}

func (amgsuo *AppMiningGoodStockUpdateOne) sqlSave(ctx context.Context) (_node *AppMiningGoodStock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appmininggoodstock.Table,
			Columns: appmininggoodstock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appmininggoodstock.FieldID,
			},
		},
	}
	id, ok := amgsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppMiningGoodStock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := amgsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appmininggoodstock.FieldID)
		for _, f := range fields {
			if !appmininggoodstock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appmininggoodstock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := amgsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := amgsuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldCreatedAt,
		})
	}
	if value, ok := amgsuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldCreatedAt,
		})
	}
	if value, ok := amgsuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldUpdatedAt,
		})
	}
	if value, ok := amgsuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldUpdatedAt,
		})
	}
	if value, ok := amgsuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldDeletedAt,
		})
	}
	if value, ok := amgsuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appmininggoodstock.FieldDeletedAt,
		})
	}
	if value, ok := amgsuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appmininggoodstock.FieldEntID,
		})
	}
	if value, ok := amgsuo.mutation.AppGoodStockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appmininggoodstock.FieldAppGoodStockID,
		})
	}
	if amgsuo.mutation.AppGoodStockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appmininggoodstock.FieldAppGoodStockID,
		})
	}
	if value, ok := amgsuo.mutation.MiningGoodStockID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appmininggoodstock.FieldMiningGoodStockID,
		})
	}
	if amgsuo.mutation.MiningGoodStockIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appmininggoodstock.FieldMiningGoodStockID,
		})
	}
	if value, ok := amgsuo.mutation.Reserved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldReserved,
		})
	}
	if amgsuo.mutation.ReservedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldReserved,
		})
	}
	if value, ok := amgsuo.mutation.SpotQuantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldSpotQuantity,
		})
	}
	if amgsuo.mutation.SpotQuantityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldSpotQuantity,
		})
	}
	if value, ok := amgsuo.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldLocked,
		})
	}
	if amgsuo.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldLocked,
		})
	}
	if value, ok := amgsuo.mutation.InService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldInService,
		})
	}
	if amgsuo.mutation.InServiceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldInService,
		})
	}
	if value, ok := amgsuo.mutation.WaitStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldWaitStart,
		})
	}
	if amgsuo.mutation.WaitStartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldWaitStart,
		})
	}
	if value, ok := amgsuo.mutation.Sold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appmininggoodstock.FieldSold,
		})
	}
	if amgsuo.mutation.SoldCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appmininggoodstock.FieldSold,
		})
	}
	_spec.Modifiers = amgsuo.modifiers
	_node = &AppMiningGoodStock{config: amgsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, amgsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appmininggoodstock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}