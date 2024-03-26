// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// StockUpdate is the builder for updating Stock entities.
type StockUpdate struct {
	config
	hooks     []Hook
	mutation  *StockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the StockUpdate builder.
func (su *StockUpdate) Where(ps ...predicate.Stock) *StockUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedAt sets the "created_at" field.
func (su *StockUpdate) SetCreatedAt(u uint32) *StockUpdate {
	su.mutation.ResetCreatedAt()
	su.mutation.SetCreatedAt(u)
	return su
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (su *StockUpdate) SetNillableCreatedAt(u *uint32) *StockUpdate {
	if u != nil {
		su.SetCreatedAt(*u)
	}
	return su
}

// AddCreatedAt adds u to the "created_at" field.
func (su *StockUpdate) AddCreatedAt(u int32) *StockUpdate {
	su.mutation.AddCreatedAt(u)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *StockUpdate) SetUpdatedAt(u uint32) *StockUpdate {
	su.mutation.ResetUpdatedAt()
	su.mutation.SetUpdatedAt(u)
	return su
}

// AddUpdatedAt adds u to the "updated_at" field.
func (su *StockUpdate) AddUpdatedAt(u int32) *StockUpdate {
	su.mutation.AddUpdatedAt(u)
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *StockUpdate) SetDeletedAt(u uint32) *StockUpdate {
	su.mutation.ResetDeletedAt()
	su.mutation.SetDeletedAt(u)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *StockUpdate) SetNillableDeletedAt(u *uint32) *StockUpdate {
	if u != nil {
		su.SetDeletedAt(*u)
	}
	return su
}

// AddDeletedAt adds u to the "deleted_at" field.
func (su *StockUpdate) AddDeletedAt(u int32) *StockUpdate {
	su.mutation.AddDeletedAt(u)
	return su
}

// SetEntID sets the "ent_id" field.
func (su *StockUpdate) SetEntID(u uuid.UUID) *StockUpdate {
	su.mutation.SetEntID(u)
	return su
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (su *StockUpdate) SetNillableEntID(u *uuid.UUID) *StockUpdate {
	if u != nil {
		su.SetEntID(*u)
	}
	return su
}

// SetGoodID sets the "good_id" field.
func (su *StockUpdate) SetGoodID(u uuid.UUID) *StockUpdate {
	su.mutation.SetGoodID(u)
	return su
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (su *StockUpdate) SetNillableGoodID(u *uuid.UUID) *StockUpdate {
	if u != nil {
		su.SetGoodID(*u)
	}
	return su
}

// ClearGoodID clears the value of the "good_id" field.
func (su *StockUpdate) ClearGoodID() *StockUpdate {
	su.mutation.ClearGoodID()
	return su
}

// SetTotal sets the "total" field.
func (su *StockUpdate) SetTotal(d decimal.Decimal) *StockUpdate {
	su.mutation.SetTotal(d)
	return su
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (su *StockUpdate) SetNillableTotal(d *decimal.Decimal) *StockUpdate {
	if d != nil {
		su.SetTotal(*d)
	}
	return su
}

// ClearTotal clears the value of the "total" field.
func (su *StockUpdate) ClearTotal() *StockUpdate {
	su.mutation.ClearTotal()
	return su
}

// SetSpotQuantity sets the "spot_quantity" field.
func (su *StockUpdate) SetSpotQuantity(d decimal.Decimal) *StockUpdate {
	su.mutation.SetSpotQuantity(d)
	return su
}

// SetNillableSpotQuantity sets the "spot_quantity" field if the given value is not nil.
func (su *StockUpdate) SetNillableSpotQuantity(d *decimal.Decimal) *StockUpdate {
	if d != nil {
		su.SetSpotQuantity(*d)
	}
	return su
}

// ClearSpotQuantity clears the value of the "spot_quantity" field.
func (su *StockUpdate) ClearSpotQuantity() *StockUpdate {
	su.mutation.ClearSpotQuantity()
	return su
}

// SetLocked sets the "locked" field.
func (su *StockUpdate) SetLocked(d decimal.Decimal) *StockUpdate {
	su.mutation.SetLocked(d)
	return su
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (su *StockUpdate) SetNillableLocked(d *decimal.Decimal) *StockUpdate {
	if d != nil {
		su.SetLocked(*d)
	}
	return su
}

// ClearLocked clears the value of the "locked" field.
func (su *StockUpdate) ClearLocked() *StockUpdate {
	su.mutation.ClearLocked()
	return su
}

// SetInService sets the "in_service" field.
func (su *StockUpdate) SetInService(d decimal.Decimal) *StockUpdate {
	su.mutation.SetInService(d)
	return su
}

// SetNillableInService sets the "in_service" field if the given value is not nil.
func (su *StockUpdate) SetNillableInService(d *decimal.Decimal) *StockUpdate {
	if d != nil {
		su.SetInService(*d)
	}
	return su
}

// ClearInService clears the value of the "in_service" field.
func (su *StockUpdate) ClearInService() *StockUpdate {
	su.mutation.ClearInService()
	return su
}

// SetWaitStart sets the "wait_start" field.
func (su *StockUpdate) SetWaitStart(d decimal.Decimal) *StockUpdate {
	su.mutation.SetWaitStart(d)
	return su
}

// SetNillableWaitStart sets the "wait_start" field if the given value is not nil.
func (su *StockUpdate) SetNillableWaitStart(d *decimal.Decimal) *StockUpdate {
	if d != nil {
		su.SetWaitStart(*d)
	}
	return su
}

// ClearWaitStart clears the value of the "wait_start" field.
func (su *StockUpdate) ClearWaitStart() *StockUpdate {
	su.mutation.ClearWaitStart()
	return su
}

// SetSold sets the "sold" field.
func (su *StockUpdate) SetSold(d decimal.Decimal) *StockUpdate {
	su.mutation.SetSold(d)
	return su
}

// SetNillableSold sets the "sold" field if the given value is not nil.
func (su *StockUpdate) SetNillableSold(d *decimal.Decimal) *StockUpdate {
	if d != nil {
		su.SetSold(*d)
	}
	return su
}

// ClearSold clears the value of the "sold" field.
func (su *StockUpdate) ClearSold() *StockUpdate {
	su.mutation.ClearSold()
	return su
}

// SetAppReserved sets the "app_reserved" field.
func (su *StockUpdate) SetAppReserved(d decimal.Decimal) *StockUpdate {
	su.mutation.SetAppReserved(d)
	return su
}

// SetNillableAppReserved sets the "app_reserved" field if the given value is not nil.
func (su *StockUpdate) SetNillableAppReserved(d *decimal.Decimal) *StockUpdate {
	if d != nil {
		su.SetAppReserved(*d)
	}
	return su
}

// ClearAppReserved clears the value of the "app_reserved" field.
func (su *StockUpdate) ClearAppReserved() *StockUpdate {
	su.mutation.ClearAppReserved()
	return su
}

// Mutation returns the StockMutation object of the builder.
func (su *StockUpdate) Mutation() *StockMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *StockUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := su.defaults(); err != nil {
		return 0, err
	}
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *StockUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StockUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StockUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *StockUpdate) defaults() error {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		if stock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized stock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := stock.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *StockUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *StockUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *StockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stock.Table,
			Columns: stock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: stock.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldCreatedAt,
		})
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldUpdatedAt,
		})
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldDeletedAt,
		})
	}
	if value, ok := su.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldDeletedAt,
		})
	}
	if value, ok := su.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stock.FieldEntID,
		})
	}
	if value, ok := su.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stock.FieldGoodID,
		})
	}
	if su.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: stock.FieldGoodID,
		})
	}
	if value, ok := su.mutation.Total(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldTotal,
		})
	}
	if su.mutation.TotalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldTotal,
		})
	}
	if value, ok := su.mutation.SpotQuantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldSpotQuantity,
		})
	}
	if su.mutation.SpotQuantityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldSpotQuantity,
		})
	}
	if value, ok := su.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldLocked,
		})
	}
	if su.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldLocked,
		})
	}
	if value, ok := su.mutation.InService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldInService,
		})
	}
	if su.mutation.InServiceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldInService,
		})
	}
	if value, ok := su.mutation.WaitStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldWaitStart,
		})
	}
	if su.mutation.WaitStartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldWaitStart,
		})
	}
	if value, ok := su.mutation.Sold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldSold,
		})
	}
	if su.mutation.SoldCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldSold,
		})
	}
	if value, ok := su.mutation.AppReserved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldAppReserved,
		})
	}
	if su.mutation.AppReservedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldAppReserved,
		})
	}
	_spec.Modifiers = su.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// StockUpdateOne is the builder for updating a single Stock entity.
type StockUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *StockMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (suo *StockUpdateOne) SetCreatedAt(u uint32) *StockUpdateOne {
	suo.mutation.ResetCreatedAt()
	suo.mutation.SetCreatedAt(u)
	return suo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableCreatedAt(u *uint32) *StockUpdateOne {
	if u != nil {
		suo.SetCreatedAt(*u)
	}
	return suo
}

// AddCreatedAt adds u to the "created_at" field.
func (suo *StockUpdateOne) AddCreatedAt(u int32) *StockUpdateOne {
	suo.mutation.AddCreatedAt(u)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *StockUpdateOne) SetUpdatedAt(u uint32) *StockUpdateOne {
	suo.mutation.ResetUpdatedAt()
	suo.mutation.SetUpdatedAt(u)
	return suo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (suo *StockUpdateOne) AddUpdatedAt(u int32) *StockUpdateOne {
	suo.mutation.AddUpdatedAt(u)
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *StockUpdateOne) SetDeletedAt(u uint32) *StockUpdateOne {
	suo.mutation.ResetDeletedAt()
	suo.mutation.SetDeletedAt(u)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableDeletedAt(u *uint32) *StockUpdateOne {
	if u != nil {
		suo.SetDeletedAt(*u)
	}
	return suo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (suo *StockUpdateOne) AddDeletedAt(u int32) *StockUpdateOne {
	suo.mutation.AddDeletedAt(u)
	return suo
}

// SetEntID sets the "ent_id" field.
func (suo *StockUpdateOne) SetEntID(u uuid.UUID) *StockUpdateOne {
	suo.mutation.SetEntID(u)
	return suo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableEntID(u *uuid.UUID) *StockUpdateOne {
	if u != nil {
		suo.SetEntID(*u)
	}
	return suo
}

// SetGoodID sets the "good_id" field.
func (suo *StockUpdateOne) SetGoodID(u uuid.UUID) *StockUpdateOne {
	suo.mutation.SetGoodID(u)
	return suo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableGoodID(u *uuid.UUID) *StockUpdateOne {
	if u != nil {
		suo.SetGoodID(*u)
	}
	return suo
}

// ClearGoodID clears the value of the "good_id" field.
func (suo *StockUpdateOne) ClearGoodID() *StockUpdateOne {
	suo.mutation.ClearGoodID()
	return suo
}

// SetTotal sets the "total" field.
func (suo *StockUpdateOne) SetTotal(d decimal.Decimal) *StockUpdateOne {
	suo.mutation.SetTotal(d)
	return suo
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableTotal(d *decimal.Decimal) *StockUpdateOne {
	if d != nil {
		suo.SetTotal(*d)
	}
	return suo
}

// ClearTotal clears the value of the "total" field.
func (suo *StockUpdateOne) ClearTotal() *StockUpdateOne {
	suo.mutation.ClearTotal()
	return suo
}

// SetSpotQuantity sets the "spot_quantity" field.
func (suo *StockUpdateOne) SetSpotQuantity(d decimal.Decimal) *StockUpdateOne {
	suo.mutation.SetSpotQuantity(d)
	return suo
}

// SetNillableSpotQuantity sets the "spot_quantity" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableSpotQuantity(d *decimal.Decimal) *StockUpdateOne {
	if d != nil {
		suo.SetSpotQuantity(*d)
	}
	return suo
}

// ClearSpotQuantity clears the value of the "spot_quantity" field.
func (suo *StockUpdateOne) ClearSpotQuantity() *StockUpdateOne {
	suo.mutation.ClearSpotQuantity()
	return suo
}

// SetLocked sets the "locked" field.
func (suo *StockUpdateOne) SetLocked(d decimal.Decimal) *StockUpdateOne {
	suo.mutation.SetLocked(d)
	return suo
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableLocked(d *decimal.Decimal) *StockUpdateOne {
	if d != nil {
		suo.SetLocked(*d)
	}
	return suo
}

// ClearLocked clears the value of the "locked" field.
func (suo *StockUpdateOne) ClearLocked() *StockUpdateOne {
	suo.mutation.ClearLocked()
	return suo
}

// SetInService sets the "in_service" field.
func (suo *StockUpdateOne) SetInService(d decimal.Decimal) *StockUpdateOne {
	suo.mutation.SetInService(d)
	return suo
}

// SetNillableInService sets the "in_service" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableInService(d *decimal.Decimal) *StockUpdateOne {
	if d != nil {
		suo.SetInService(*d)
	}
	return suo
}

// ClearInService clears the value of the "in_service" field.
func (suo *StockUpdateOne) ClearInService() *StockUpdateOne {
	suo.mutation.ClearInService()
	return suo
}

// SetWaitStart sets the "wait_start" field.
func (suo *StockUpdateOne) SetWaitStart(d decimal.Decimal) *StockUpdateOne {
	suo.mutation.SetWaitStart(d)
	return suo
}

// SetNillableWaitStart sets the "wait_start" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableWaitStart(d *decimal.Decimal) *StockUpdateOne {
	if d != nil {
		suo.SetWaitStart(*d)
	}
	return suo
}

// ClearWaitStart clears the value of the "wait_start" field.
func (suo *StockUpdateOne) ClearWaitStart() *StockUpdateOne {
	suo.mutation.ClearWaitStart()
	return suo
}

// SetSold sets the "sold" field.
func (suo *StockUpdateOne) SetSold(d decimal.Decimal) *StockUpdateOne {
	suo.mutation.SetSold(d)
	return suo
}

// SetNillableSold sets the "sold" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableSold(d *decimal.Decimal) *StockUpdateOne {
	if d != nil {
		suo.SetSold(*d)
	}
	return suo
}

// ClearSold clears the value of the "sold" field.
func (suo *StockUpdateOne) ClearSold() *StockUpdateOne {
	suo.mutation.ClearSold()
	return suo
}

// SetAppReserved sets the "app_reserved" field.
func (suo *StockUpdateOne) SetAppReserved(d decimal.Decimal) *StockUpdateOne {
	suo.mutation.SetAppReserved(d)
	return suo
}

// SetNillableAppReserved sets the "app_reserved" field if the given value is not nil.
func (suo *StockUpdateOne) SetNillableAppReserved(d *decimal.Decimal) *StockUpdateOne {
	if d != nil {
		suo.SetAppReserved(*d)
	}
	return suo
}

// ClearAppReserved clears the value of the "app_reserved" field.
func (suo *StockUpdateOne) ClearAppReserved() *StockUpdateOne {
	suo.mutation.ClearAppReserved()
	return suo
}

// Mutation returns the StockMutation object of the builder.
func (suo *StockUpdateOne) Mutation() *StockMutation {
	return suo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *StockUpdateOne) Select(field string, fields ...string) *StockUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Stock entity.
func (suo *StockUpdateOne) Save(ctx context.Context) (*Stock, error) {
	var (
		err  error
		node *Stock
	)
	if err := suo.defaults(); err != nil {
		return nil, err
	}
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Stock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StockUpdateOne) SaveX(ctx context.Context) *Stock {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *StockUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StockUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *StockUpdateOne) defaults() error {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		if stock.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized stock.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := stock.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *StockUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *StockUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *StockUpdateOne) sqlSave(ctx context.Context) (_node *Stock, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   stock.Table,
			Columns: stock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: stock.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Stock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, stock.FieldID)
		for _, f := range fields {
			if !stock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != stock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldCreatedAt,
		})
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldUpdatedAt,
		})
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldDeletedAt,
		})
	}
	if value, ok := suo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldDeletedAt,
		})
	}
	if value, ok := suo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stock.FieldEntID,
		})
	}
	if value, ok := suo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stock.FieldGoodID,
		})
	}
	if suo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: stock.FieldGoodID,
		})
	}
	if value, ok := suo.mutation.Total(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldTotal,
		})
	}
	if suo.mutation.TotalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldTotal,
		})
	}
	if value, ok := suo.mutation.SpotQuantity(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldSpotQuantity,
		})
	}
	if suo.mutation.SpotQuantityCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldSpotQuantity,
		})
	}
	if value, ok := suo.mutation.Locked(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldLocked,
		})
	}
	if suo.mutation.LockedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldLocked,
		})
	}
	if value, ok := suo.mutation.InService(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldInService,
		})
	}
	if suo.mutation.InServiceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldInService,
		})
	}
	if value, ok := suo.mutation.WaitStart(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldWaitStart,
		})
	}
	if suo.mutation.WaitStartCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldWaitStart,
		})
	}
	if value, ok := suo.mutation.Sold(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldSold,
		})
	}
	if suo.mutation.SoldCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldSold,
		})
	}
	if value, ok := suo.mutation.AppReserved(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldAppReserved,
		})
	}
	if suo.mutation.AppReservedCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: stock.FieldAppReserved,
		})
	}
	_spec.Modifiers = suo.modifiers
	_node = &Stock{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{stock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
