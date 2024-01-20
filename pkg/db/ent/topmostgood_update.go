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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TopMostGoodUpdate is the builder for updating TopMostGood entities.
type TopMostGoodUpdate struct {
	config
	hooks     []Hook
	mutation  *TopMostGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TopMostGoodUpdate builder.
func (tmgu *TopMostGoodUpdate) Where(ps ...predicate.TopMostGood) *TopMostGoodUpdate {
	tmgu.mutation.Where(ps...)
	return tmgu
}

// SetCreatedAt sets the "created_at" field.
func (tmgu *TopMostGoodUpdate) SetCreatedAt(u uint32) *TopMostGoodUpdate {
	tmgu.mutation.ResetCreatedAt()
	tmgu.mutation.SetCreatedAt(u)
	return tmgu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmgu *TopMostGoodUpdate) SetNillableCreatedAt(u *uint32) *TopMostGoodUpdate {
	if u != nil {
		tmgu.SetCreatedAt(*u)
	}
	return tmgu
}

// AddCreatedAt adds u to the "created_at" field.
func (tmgu *TopMostGoodUpdate) AddCreatedAt(u int32) *TopMostGoodUpdate {
	tmgu.mutation.AddCreatedAt(u)
	return tmgu
}

// SetUpdatedAt sets the "updated_at" field.
func (tmgu *TopMostGoodUpdate) SetUpdatedAt(u uint32) *TopMostGoodUpdate {
	tmgu.mutation.ResetUpdatedAt()
	tmgu.mutation.SetUpdatedAt(u)
	return tmgu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tmgu *TopMostGoodUpdate) AddUpdatedAt(u int32) *TopMostGoodUpdate {
	tmgu.mutation.AddUpdatedAt(u)
	return tmgu
}

// SetDeletedAt sets the "deleted_at" field.
func (tmgu *TopMostGoodUpdate) SetDeletedAt(u uint32) *TopMostGoodUpdate {
	tmgu.mutation.ResetDeletedAt()
	tmgu.mutation.SetDeletedAt(u)
	return tmgu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmgu *TopMostGoodUpdate) SetNillableDeletedAt(u *uint32) *TopMostGoodUpdate {
	if u != nil {
		tmgu.SetDeletedAt(*u)
	}
	return tmgu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tmgu *TopMostGoodUpdate) AddDeletedAt(u int32) *TopMostGoodUpdate {
	tmgu.mutation.AddDeletedAt(u)
	return tmgu
}

// SetEntID sets the "ent_id" field.
func (tmgu *TopMostGoodUpdate) SetEntID(u uuid.UUID) *TopMostGoodUpdate {
	tmgu.mutation.SetEntID(u)
	return tmgu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmgu *TopMostGoodUpdate) SetNillableEntID(u *uuid.UUID) *TopMostGoodUpdate {
	if u != nil {
		tmgu.SetEntID(*u)
	}
	return tmgu
}

// SetAppID sets the "app_id" field.
func (tmgu *TopMostGoodUpdate) SetAppID(u uuid.UUID) *TopMostGoodUpdate {
	tmgu.mutation.SetAppID(u)
	return tmgu
}

// SetGoodID sets the "good_id" field.
func (tmgu *TopMostGoodUpdate) SetGoodID(u uuid.UUID) *TopMostGoodUpdate {
	tmgu.mutation.SetGoodID(u)
	return tmgu
}

// SetAppGoodID sets the "app_good_id" field.
func (tmgu *TopMostGoodUpdate) SetAppGoodID(u uuid.UUID) *TopMostGoodUpdate {
	tmgu.mutation.SetAppGoodID(u)
	return tmgu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (tmgu *TopMostGoodUpdate) SetCoinTypeID(u uuid.UUID) *TopMostGoodUpdate {
	tmgu.mutation.SetCoinTypeID(u)
	return tmgu
}

// SetTopMostID sets the "top_most_id" field.
func (tmgu *TopMostGoodUpdate) SetTopMostID(u uuid.UUID) *TopMostGoodUpdate {
	tmgu.mutation.SetTopMostID(u)
	return tmgu
}

// SetDisplayIndex sets the "display_index" field.
func (tmgu *TopMostGoodUpdate) SetDisplayIndex(u uint32) *TopMostGoodUpdate {
	tmgu.mutation.ResetDisplayIndex()
	tmgu.mutation.SetDisplayIndex(u)
	return tmgu
}

// SetNillableDisplayIndex sets the "display_index" field if the given value is not nil.
func (tmgu *TopMostGoodUpdate) SetNillableDisplayIndex(u *uint32) *TopMostGoodUpdate {
	if u != nil {
		tmgu.SetDisplayIndex(*u)
	}
	return tmgu
}

// AddDisplayIndex adds u to the "display_index" field.
func (tmgu *TopMostGoodUpdate) AddDisplayIndex(u int32) *TopMostGoodUpdate {
	tmgu.mutation.AddDisplayIndex(u)
	return tmgu
}

// ClearDisplayIndex clears the value of the "display_index" field.
func (tmgu *TopMostGoodUpdate) ClearDisplayIndex() *TopMostGoodUpdate {
	tmgu.mutation.ClearDisplayIndex()
	return tmgu
}

// SetPosters sets the "posters" field.
func (tmgu *TopMostGoodUpdate) SetPosters(s []string) *TopMostGoodUpdate {
	tmgu.mutation.SetPosters(s)
	return tmgu
}

// ClearPosters clears the value of the "posters" field.
func (tmgu *TopMostGoodUpdate) ClearPosters() *TopMostGoodUpdate {
	tmgu.mutation.ClearPosters()
	return tmgu
}

// SetUnitPrice sets the "unit_price" field.
func (tmgu *TopMostGoodUpdate) SetUnitPrice(d decimal.Decimal) *TopMostGoodUpdate {
	tmgu.mutation.SetUnitPrice(d)
	return tmgu
}

// SetNillableUnitPrice sets the "unit_price" field if the given value is not nil.
func (tmgu *TopMostGoodUpdate) SetNillableUnitPrice(d *decimal.Decimal) *TopMostGoodUpdate {
	if d != nil {
		tmgu.SetUnitPrice(*d)
	}
	return tmgu
}

// ClearUnitPrice clears the value of the "unit_price" field.
func (tmgu *TopMostGoodUpdate) ClearUnitPrice() *TopMostGoodUpdate {
	tmgu.mutation.ClearUnitPrice()
	return tmgu
}

// SetPackagePrice sets the "package_price" field.
func (tmgu *TopMostGoodUpdate) SetPackagePrice(d decimal.Decimal) *TopMostGoodUpdate {
	tmgu.mutation.SetPackagePrice(d)
	return tmgu
}

// SetNillablePackagePrice sets the "package_price" field if the given value is not nil.
func (tmgu *TopMostGoodUpdate) SetNillablePackagePrice(d *decimal.Decimal) *TopMostGoodUpdate {
	if d != nil {
		tmgu.SetPackagePrice(*d)
	}
	return tmgu
}

// ClearPackagePrice clears the value of the "package_price" field.
func (tmgu *TopMostGoodUpdate) ClearPackagePrice() *TopMostGoodUpdate {
	tmgu.mutation.ClearPackagePrice()
	return tmgu
}

// Mutation returns the TopMostGoodMutation object of the builder.
func (tmgu *TopMostGoodUpdate) Mutation() *TopMostGoodMutation {
	return tmgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tmgu *TopMostGoodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := tmgu.defaults(); err != nil {
		return 0, err
	}
	if len(tmgu.hooks) == 0 {
		affected, err = tmgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopMostGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmgu.mutation = mutation
			affected, err = tmgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tmgu.hooks) - 1; i >= 0; i-- {
			if tmgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmgu *TopMostGoodUpdate) SaveX(ctx context.Context) int {
	affected, err := tmgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tmgu *TopMostGoodUpdate) Exec(ctx context.Context) error {
	_, err := tmgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmgu *TopMostGoodUpdate) ExecX(ctx context.Context) {
	if err := tmgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmgu *TopMostGoodUpdate) defaults() error {
	if _, ok := tmgu.mutation.UpdatedAt(); !ok {
		if topmostgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topmostgood.UpdateDefaultUpdatedAt()
		tmgu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmgu *TopMostGoodUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TopMostGoodUpdate {
	tmgu.modifiers = append(tmgu.modifiers, modifiers...)
	return tmgu
}

func (tmgu *TopMostGoodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topmostgood.Table,
			Columns: topmostgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostgood.FieldID,
			},
		},
	}
	if ps := tmgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmgu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldCreatedAt,
		})
	}
	if value, ok := tmgu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldCreatedAt,
		})
	}
	if value, ok := tmgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldUpdatedAt,
		})
	}
	if value, ok := tmgu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldUpdatedAt,
		})
	}
	if value, ok := tmgu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDeletedAt,
		})
	}
	if value, ok := tmgu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDeletedAt,
		})
	}
	if value, ok := tmgu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldEntID,
		})
	}
	if value, ok := tmgu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldAppID,
		})
	}
	if value, ok := tmgu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldGoodID,
		})
	}
	if value, ok := tmgu.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldAppGoodID,
		})
	}
	if value, ok := tmgu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldCoinTypeID,
		})
	}
	if value, ok := tmgu.mutation.TopMostID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldTopMostID,
		})
	}
	if value, ok := tmgu.mutation.DisplayIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDisplayIndex,
		})
	}
	if value, ok := tmgu.mutation.AddedDisplayIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDisplayIndex,
		})
	}
	if tmgu.mutation.DisplayIndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: topmostgood.FieldDisplayIndex,
		})
	}
	if value, ok := tmgu.mutation.Posters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: topmostgood.FieldPosters,
		})
	}
	if tmgu.mutation.PostersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: topmostgood.FieldPosters,
		})
	}
	if value, ok := tmgu.mutation.UnitPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: topmostgood.FieldUnitPrice,
		})
	}
	if tmgu.mutation.UnitPriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: topmostgood.FieldUnitPrice,
		})
	}
	if value, ok := tmgu.mutation.PackagePrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: topmostgood.FieldPackagePrice,
		})
	}
	if tmgu.mutation.PackagePriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: topmostgood.FieldPackagePrice,
		})
	}
	_spec.Modifiers = tmgu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, tmgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topmostgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TopMostGoodUpdateOne is the builder for updating a single TopMostGood entity.
type TopMostGoodUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TopMostGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (tmguo *TopMostGoodUpdateOne) SetCreatedAt(u uint32) *TopMostGoodUpdateOne {
	tmguo.mutation.ResetCreatedAt()
	tmguo.mutation.SetCreatedAt(u)
	return tmguo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmguo *TopMostGoodUpdateOne) SetNillableCreatedAt(u *uint32) *TopMostGoodUpdateOne {
	if u != nil {
		tmguo.SetCreatedAt(*u)
	}
	return tmguo
}

// AddCreatedAt adds u to the "created_at" field.
func (tmguo *TopMostGoodUpdateOne) AddCreatedAt(u int32) *TopMostGoodUpdateOne {
	tmguo.mutation.AddCreatedAt(u)
	return tmguo
}

// SetUpdatedAt sets the "updated_at" field.
func (tmguo *TopMostGoodUpdateOne) SetUpdatedAt(u uint32) *TopMostGoodUpdateOne {
	tmguo.mutation.ResetUpdatedAt()
	tmguo.mutation.SetUpdatedAt(u)
	return tmguo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tmguo *TopMostGoodUpdateOne) AddUpdatedAt(u int32) *TopMostGoodUpdateOne {
	tmguo.mutation.AddUpdatedAt(u)
	return tmguo
}

// SetDeletedAt sets the "deleted_at" field.
func (tmguo *TopMostGoodUpdateOne) SetDeletedAt(u uint32) *TopMostGoodUpdateOne {
	tmguo.mutation.ResetDeletedAt()
	tmguo.mutation.SetDeletedAt(u)
	return tmguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmguo *TopMostGoodUpdateOne) SetNillableDeletedAt(u *uint32) *TopMostGoodUpdateOne {
	if u != nil {
		tmguo.SetDeletedAt(*u)
	}
	return tmguo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tmguo *TopMostGoodUpdateOne) AddDeletedAt(u int32) *TopMostGoodUpdateOne {
	tmguo.mutation.AddDeletedAt(u)
	return tmguo
}

// SetEntID sets the "ent_id" field.
func (tmguo *TopMostGoodUpdateOne) SetEntID(u uuid.UUID) *TopMostGoodUpdateOne {
	tmguo.mutation.SetEntID(u)
	return tmguo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmguo *TopMostGoodUpdateOne) SetNillableEntID(u *uuid.UUID) *TopMostGoodUpdateOne {
	if u != nil {
		tmguo.SetEntID(*u)
	}
	return tmguo
}

// SetAppID sets the "app_id" field.
func (tmguo *TopMostGoodUpdateOne) SetAppID(u uuid.UUID) *TopMostGoodUpdateOne {
	tmguo.mutation.SetAppID(u)
	return tmguo
}

// SetGoodID sets the "good_id" field.
func (tmguo *TopMostGoodUpdateOne) SetGoodID(u uuid.UUID) *TopMostGoodUpdateOne {
	tmguo.mutation.SetGoodID(u)
	return tmguo
}

// SetAppGoodID sets the "app_good_id" field.
func (tmguo *TopMostGoodUpdateOne) SetAppGoodID(u uuid.UUID) *TopMostGoodUpdateOne {
	tmguo.mutation.SetAppGoodID(u)
	return tmguo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (tmguo *TopMostGoodUpdateOne) SetCoinTypeID(u uuid.UUID) *TopMostGoodUpdateOne {
	tmguo.mutation.SetCoinTypeID(u)
	return tmguo
}

// SetTopMostID sets the "top_most_id" field.
func (tmguo *TopMostGoodUpdateOne) SetTopMostID(u uuid.UUID) *TopMostGoodUpdateOne {
	tmguo.mutation.SetTopMostID(u)
	return tmguo
}

// SetDisplayIndex sets the "display_index" field.
func (tmguo *TopMostGoodUpdateOne) SetDisplayIndex(u uint32) *TopMostGoodUpdateOne {
	tmguo.mutation.ResetDisplayIndex()
	tmguo.mutation.SetDisplayIndex(u)
	return tmguo
}

// SetNillableDisplayIndex sets the "display_index" field if the given value is not nil.
func (tmguo *TopMostGoodUpdateOne) SetNillableDisplayIndex(u *uint32) *TopMostGoodUpdateOne {
	if u != nil {
		tmguo.SetDisplayIndex(*u)
	}
	return tmguo
}

// AddDisplayIndex adds u to the "display_index" field.
func (tmguo *TopMostGoodUpdateOne) AddDisplayIndex(u int32) *TopMostGoodUpdateOne {
	tmguo.mutation.AddDisplayIndex(u)
	return tmguo
}

// ClearDisplayIndex clears the value of the "display_index" field.
func (tmguo *TopMostGoodUpdateOne) ClearDisplayIndex() *TopMostGoodUpdateOne {
	tmguo.mutation.ClearDisplayIndex()
	return tmguo
}

// SetPosters sets the "posters" field.
func (tmguo *TopMostGoodUpdateOne) SetPosters(s []string) *TopMostGoodUpdateOne {
	tmguo.mutation.SetPosters(s)
	return tmguo
}

// ClearPosters clears the value of the "posters" field.
func (tmguo *TopMostGoodUpdateOne) ClearPosters() *TopMostGoodUpdateOne {
	tmguo.mutation.ClearPosters()
	return tmguo
}

// SetUnitPrice sets the "unit_price" field.
func (tmguo *TopMostGoodUpdateOne) SetUnitPrice(d decimal.Decimal) *TopMostGoodUpdateOne {
	tmguo.mutation.SetUnitPrice(d)
	return tmguo
}

// SetNillableUnitPrice sets the "unit_price" field if the given value is not nil.
func (tmguo *TopMostGoodUpdateOne) SetNillableUnitPrice(d *decimal.Decimal) *TopMostGoodUpdateOne {
	if d != nil {
		tmguo.SetUnitPrice(*d)
	}
	return tmguo
}

// ClearUnitPrice clears the value of the "unit_price" field.
func (tmguo *TopMostGoodUpdateOne) ClearUnitPrice() *TopMostGoodUpdateOne {
	tmguo.mutation.ClearUnitPrice()
	return tmguo
}

// SetPackagePrice sets the "package_price" field.
func (tmguo *TopMostGoodUpdateOne) SetPackagePrice(d decimal.Decimal) *TopMostGoodUpdateOne {
	tmguo.mutation.SetPackagePrice(d)
	return tmguo
}

// SetNillablePackagePrice sets the "package_price" field if the given value is not nil.
func (tmguo *TopMostGoodUpdateOne) SetNillablePackagePrice(d *decimal.Decimal) *TopMostGoodUpdateOne {
	if d != nil {
		tmguo.SetPackagePrice(*d)
	}
	return tmguo
}

// ClearPackagePrice clears the value of the "package_price" field.
func (tmguo *TopMostGoodUpdateOne) ClearPackagePrice() *TopMostGoodUpdateOne {
	tmguo.mutation.ClearPackagePrice()
	return tmguo
}

// Mutation returns the TopMostGoodMutation object of the builder.
func (tmguo *TopMostGoodUpdateOne) Mutation() *TopMostGoodMutation {
	return tmguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tmguo *TopMostGoodUpdateOne) Select(field string, fields ...string) *TopMostGoodUpdateOne {
	tmguo.fields = append([]string{field}, fields...)
	return tmguo
}

// Save executes the query and returns the updated TopMostGood entity.
func (tmguo *TopMostGoodUpdateOne) Save(ctx context.Context) (*TopMostGood, error) {
	var (
		err  error
		node *TopMostGood
	)
	if err := tmguo.defaults(); err != nil {
		return nil, err
	}
	if len(tmguo.hooks) == 0 {
		node, err = tmguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopMostGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmguo.mutation = mutation
			node, err = tmguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tmguo.hooks) - 1; i >= 0; i-- {
			if tmguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tmguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TopMostGood)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TopMostGoodMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmguo *TopMostGoodUpdateOne) SaveX(ctx context.Context) *TopMostGood {
	node, err := tmguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tmguo *TopMostGoodUpdateOne) Exec(ctx context.Context) error {
	_, err := tmguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmguo *TopMostGoodUpdateOne) ExecX(ctx context.Context) {
	if err := tmguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmguo *TopMostGoodUpdateOne) defaults() error {
	if _, ok := tmguo.mutation.UpdatedAt(); !ok {
		if topmostgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topmostgood.UpdateDefaultUpdatedAt()
		tmguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmguo *TopMostGoodUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TopMostGoodUpdateOne {
	tmguo.modifiers = append(tmguo.modifiers, modifiers...)
	return tmguo
}

func (tmguo *TopMostGoodUpdateOne) sqlSave(ctx context.Context) (_node *TopMostGood, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topmostgood.Table,
			Columns: topmostgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostgood.FieldID,
			},
		},
	}
	id, ok := tmguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TopMostGood.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tmguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topmostgood.FieldID)
		for _, f := range fields {
			if !topmostgood.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topmostgood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tmguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldCreatedAt,
		})
	}
	if value, ok := tmguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldCreatedAt,
		})
	}
	if value, ok := tmguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldUpdatedAt,
		})
	}
	if value, ok := tmguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldUpdatedAt,
		})
	}
	if value, ok := tmguo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDeletedAt,
		})
	}
	if value, ok := tmguo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDeletedAt,
		})
	}
	if value, ok := tmguo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldEntID,
		})
	}
	if value, ok := tmguo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldAppID,
		})
	}
	if value, ok := tmguo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldGoodID,
		})
	}
	if value, ok := tmguo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldAppGoodID,
		})
	}
	if value, ok := tmguo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldCoinTypeID,
		})
	}
	if value, ok := tmguo.mutation.TopMostID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgood.FieldTopMostID,
		})
	}
	if value, ok := tmguo.mutation.DisplayIndex(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDisplayIndex,
		})
	}
	if value, ok := tmguo.mutation.AddedDisplayIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgood.FieldDisplayIndex,
		})
	}
	if tmguo.mutation.DisplayIndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: topmostgood.FieldDisplayIndex,
		})
	}
	if value, ok := tmguo.mutation.Posters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: topmostgood.FieldPosters,
		})
	}
	if tmguo.mutation.PostersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: topmostgood.FieldPosters,
		})
	}
	if value, ok := tmguo.mutation.UnitPrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: topmostgood.FieldUnitPrice,
		})
	}
	if tmguo.mutation.UnitPriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: topmostgood.FieldUnitPrice,
		})
	}
	if value, ok := tmguo.mutation.PackagePrice(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: topmostgood.FieldPackagePrice,
		})
	}
	if tmguo.mutation.PackagePriceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: topmostgood.FieldPackagePrice,
		})
	}
	_spec.Modifiers = tmguo.modifiers
	_node = &TopMostGood{config: tmguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tmguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topmostgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
