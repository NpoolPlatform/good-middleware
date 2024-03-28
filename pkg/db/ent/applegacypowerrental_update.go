// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/applegacypowerrental"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppLegacyPowerRentalUpdate is the builder for updating AppLegacyPowerRental entities.
type AppLegacyPowerRentalUpdate struct {
	config
	hooks     []Hook
	mutation  *AppLegacyPowerRentalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppLegacyPowerRentalUpdate builder.
func (alpru *AppLegacyPowerRentalUpdate) Where(ps ...predicate.AppLegacyPowerRental) *AppLegacyPowerRentalUpdate {
	alpru.mutation.Where(ps...)
	return alpru
}

// SetCreatedAt sets the "created_at" field.
func (alpru *AppLegacyPowerRentalUpdate) SetCreatedAt(u uint32) *AppLegacyPowerRentalUpdate {
	alpru.mutation.ResetCreatedAt()
	alpru.mutation.SetCreatedAt(u)
	return alpru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alpru *AppLegacyPowerRentalUpdate) SetNillableCreatedAt(u *uint32) *AppLegacyPowerRentalUpdate {
	if u != nil {
		alpru.SetCreatedAt(*u)
	}
	return alpru
}

// AddCreatedAt adds u to the "created_at" field.
func (alpru *AppLegacyPowerRentalUpdate) AddCreatedAt(u int32) *AppLegacyPowerRentalUpdate {
	alpru.mutation.AddCreatedAt(u)
	return alpru
}

// SetUpdatedAt sets the "updated_at" field.
func (alpru *AppLegacyPowerRentalUpdate) SetUpdatedAt(u uint32) *AppLegacyPowerRentalUpdate {
	alpru.mutation.ResetUpdatedAt()
	alpru.mutation.SetUpdatedAt(u)
	return alpru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (alpru *AppLegacyPowerRentalUpdate) AddUpdatedAt(u int32) *AppLegacyPowerRentalUpdate {
	alpru.mutation.AddUpdatedAt(u)
	return alpru
}

// SetDeletedAt sets the "deleted_at" field.
func (alpru *AppLegacyPowerRentalUpdate) SetDeletedAt(u uint32) *AppLegacyPowerRentalUpdate {
	alpru.mutation.ResetDeletedAt()
	alpru.mutation.SetDeletedAt(u)
	return alpru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (alpru *AppLegacyPowerRentalUpdate) SetNillableDeletedAt(u *uint32) *AppLegacyPowerRentalUpdate {
	if u != nil {
		alpru.SetDeletedAt(*u)
	}
	return alpru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (alpru *AppLegacyPowerRentalUpdate) AddDeletedAt(u int32) *AppLegacyPowerRentalUpdate {
	alpru.mutation.AddDeletedAt(u)
	return alpru
}

// SetEntID sets the "ent_id" field.
func (alpru *AppLegacyPowerRentalUpdate) SetEntID(u uuid.UUID) *AppLegacyPowerRentalUpdate {
	alpru.mutation.SetEntID(u)
	return alpru
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (alpru *AppLegacyPowerRentalUpdate) SetNillableEntID(u *uuid.UUID) *AppLegacyPowerRentalUpdate {
	if u != nil {
		alpru.SetEntID(*u)
	}
	return alpru
}

// SetAppGoodID sets the "app_good_id" field.
func (alpru *AppLegacyPowerRentalUpdate) SetAppGoodID(u uuid.UUID) *AppLegacyPowerRentalUpdate {
	alpru.mutation.SetAppGoodID(u)
	return alpru
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (alpru *AppLegacyPowerRentalUpdate) SetNillableAppGoodID(u *uuid.UUID) *AppLegacyPowerRentalUpdate {
	if u != nil {
		alpru.SetAppGoodID(*u)
	}
	return alpru
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (alpru *AppLegacyPowerRentalUpdate) ClearAppGoodID() *AppLegacyPowerRentalUpdate {
	alpru.mutation.ClearAppGoodID()
	return alpru
}

// SetTechniqueFeeRatio sets the "technique_fee_ratio" field.
func (alpru *AppLegacyPowerRentalUpdate) SetTechniqueFeeRatio(d decimal.Decimal) *AppLegacyPowerRentalUpdate {
	alpru.mutation.SetTechniqueFeeRatio(d)
	return alpru
}

// SetNillableTechniqueFeeRatio sets the "technique_fee_ratio" field if the given value is not nil.
func (alpru *AppLegacyPowerRentalUpdate) SetNillableTechniqueFeeRatio(d *decimal.Decimal) *AppLegacyPowerRentalUpdate {
	if d != nil {
		alpru.SetTechniqueFeeRatio(*d)
	}
	return alpru
}

// ClearTechniqueFeeRatio clears the value of the "technique_fee_ratio" field.
func (alpru *AppLegacyPowerRentalUpdate) ClearTechniqueFeeRatio() *AppLegacyPowerRentalUpdate {
	alpru.mutation.ClearTechniqueFeeRatio()
	return alpru
}

// Mutation returns the AppLegacyPowerRentalMutation object of the builder.
func (alpru *AppLegacyPowerRentalUpdate) Mutation() *AppLegacyPowerRentalMutation {
	return alpru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (alpru *AppLegacyPowerRentalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := alpru.defaults(); err != nil {
		return 0, err
	}
	if len(alpru.hooks) == 0 {
		affected, err = alpru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppLegacyPowerRentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			alpru.mutation = mutation
			affected, err = alpru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(alpru.hooks) - 1; i >= 0; i-- {
			if alpru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = alpru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, alpru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (alpru *AppLegacyPowerRentalUpdate) SaveX(ctx context.Context) int {
	affected, err := alpru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (alpru *AppLegacyPowerRentalUpdate) Exec(ctx context.Context) error {
	_, err := alpru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alpru *AppLegacyPowerRentalUpdate) ExecX(ctx context.Context) {
	if err := alpru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alpru *AppLegacyPowerRentalUpdate) defaults() error {
	if _, ok := alpru.mutation.UpdatedAt(); !ok {
		if applegacypowerrental.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized applegacypowerrental.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := applegacypowerrental.UpdateDefaultUpdatedAt()
		alpru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (alpru *AppLegacyPowerRentalUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppLegacyPowerRentalUpdate {
	alpru.modifiers = append(alpru.modifiers, modifiers...)
	return alpru
}

func (alpru *AppLegacyPowerRentalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applegacypowerrental.Table,
			Columns: applegacypowerrental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: applegacypowerrental.FieldID,
			},
		},
	}
	if ps := alpru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := alpru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldCreatedAt,
		})
	}
	if value, ok := alpru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldCreatedAt,
		})
	}
	if value, ok := alpru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := alpru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := alpru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldDeletedAt,
		})
	}
	if value, ok := alpru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldDeletedAt,
		})
	}
	if value, ok := alpru.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applegacypowerrental.FieldEntID,
		})
	}
	if value, ok := alpru.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applegacypowerrental.FieldAppGoodID,
		})
	}
	if alpru.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: applegacypowerrental.FieldAppGoodID,
		})
	}
	if value, ok := alpru.mutation.TechniqueFeeRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: applegacypowerrental.FieldTechniqueFeeRatio,
		})
	}
	if alpru.mutation.TechniqueFeeRatioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: applegacypowerrental.FieldTechniqueFeeRatio,
		})
	}
	_spec.Modifiers = alpru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, alpru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applegacypowerrental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppLegacyPowerRentalUpdateOne is the builder for updating a single AppLegacyPowerRental entity.
type AppLegacyPowerRentalUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppLegacyPowerRentalMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetCreatedAt(u uint32) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.ResetCreatedAt()
	alpruo.mutation.SetCreatedAt(u)
	return alpruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetNillableCreatedAt(u *uint32) *AppLegacyPowerRentalUpdateOne {
	if u != nil {
		alpruo.SetCreatedAt(*u)
	}
	return alpruo
}

// AddCreatedAt adds u to the "created_at" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) AddCreatedAt(u int32) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.AddCreatedAt(u)
	return alpruo
}

// SetUpdatedAt sets the "updated_at" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetUpdatedAt(u uint32) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.ResetUpdatedAt()
	alpruo.mutation.SetUpdatedAt(u)
	return alpruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) AddUpdatedAt(u int32) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.AddUpdatedAt(u)
	return alpruo
}

// SetDeletedAt sets the "deleted_at" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetDeletedAt(u uint32) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.ResetDeletedAt()
	alpruo.mutation.SetDeletedAt(u)
	return alpruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetNillableDeletedAt(u *uint32) *AppLegacyPowerRentalUpdateOne {
	if u != nil {
		alpruo.SetDeletedAt(*u)
	}
	return alpruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) AddDeletedAt(u int32) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.AddDeletedAt(u)
	return alpruo
}

// SetEntID sets the "ent_id" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetEntID(u uuid.UUID) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.SetEntID(u)
	return alpruo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetNillableEntID(u *uuid.UUID) *AppLegacyPowerRentalUpdateOne {
	if u != nil {
		alpruo.SetEntID(*u)
	}
	return alpruo
}

// SetAppGoodID sets the "app_good_id" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetAppGoodID(u uuid.UUID) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.SetAppGoodID(u)
	return alpruo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *AppLegacyPowerRentalUpdateOne {
	if u != nil {
		alpruo.SetAppGoodID(*u)
	}
	return alpruo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) ClearAppGoodID() *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.ClearAppGoodID()
	return alpruo
}

// SetTechniqueFeeRatio sets the "technique_fee_ratio" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetTechniqueFeeRatio(d decimal.Decimal) *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.SetTechniqueFeeRatio(d)
	return alpruo
}

// SetNillableTechniqueFeeRatio sets the "technique_fee_ratio" field if the given value is not nil.
func (alpruo *AppLegacyPowerRentalUpdateOne) SetNillableTechniqueFeeRatio(d *decimal.Decimal) *AppLegacyPowerRentalUpdateOne {
	if d != nil {
		alpruo.SetTechniqueFeeRatio(*d)
	}
	return alpruo
}

// ClearTechniqueFeeRatio clears the value of the "technique_fee_ratio" field.
func (alpruo *AppLegacyPowerRentalUpdateOne) ClearTechniqueFeeRatio() *AppLegacyPowerRentalUpdateOne {
	alpruo.mutation.ClearTechniqueFeeRatio()
	return alpruo
}

// Mutation returns the AppLegacyPowerRentalMutation object of the builder.
func (alpruo *AppLegacyPowerRentalUpdateOne) Mutation() *AppLegacyPowerRentalMutation {
	return alpruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (alpruo *AppLegacyPowerRentalUpdateOne) Select(field string, fields ...string) *AppLegacyPowerRentalUpdateOne {
	alpruo.fields = append([]string{field}, fields...)
	return alpruo
}

// Save executes the query and returns the updated AppLegacyPowerRental entity.
func (alpruo *AppLegacyPowerRentalUpdateOne) Save(ctx context.Context) (*AppLegacyPowerRental, error) {
	var (
		err  error
		node *AppLegacyPowerRental
	)
	if err := alpruo.defaults(); err != nil {
		return nil, err
	}
	if len(alpruo.hooks) == 0 {
		node, err = alpruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppLegacyPowerRentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			alpruo.mutation = mutation
			node, err = alpruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(alpruo.hooks) - 1; i >= 0; i-- {
			if alpruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = alpruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, alpruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppLegacyPowerRental)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppLegacyPowerRentalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (alpruo *AppLegacyPowerRentalUpdateOne) SaveX(ctx context.Context) *AppLegacyPowerRental {
	node, err := alpruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (alpruo *AppLegacyPowerRentalUpdateOne) Exec(ctx context.Context) error {
	_, err := alpruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (alpruo *AppLegacyPowerRentalUpdateOne) ExecX(ctx context.Context) {
	if err := alpruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (alpruo *AppLegacyPowerRentalUpdateOne) defaults() error {
	if _, ok := alpruo.mutation.UpdatedAt(); !ok {
		if applegacypowerrental.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized applegacypowerrental.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := applegacypowerrental.UpdateDefaultUpdatedAt()
		alpruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (alpruo *AppLegacyPowerRentalUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppLegacyPowerRentalUpdateOne {
	alpruo.modifiers = append(alpruo.modifiers, modifiers...)
	return alpruo
}

func (alpruo *AppLegacyPowerRentalUpdateOne) sqlSave(ctx context.Context) (_node *AppLegacyPowerRental, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   applegacypowerrental.Table,
			Columns: applegacypowerrental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: applegacypowerrental.FieldID,
			},
		},
	}
	id, ok := alpruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppLegacyPowerRental.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := alpruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, applegacypowerrental.FieldID)
		for _, f := range fields {
			if !applegacypowerrental.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != applegacypowerrental.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := alpruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := alpruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldCreatedAt,
		})
	}
	if value, ok := alpruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldCreatedAt,
		})
	}
	if value, ok := alpruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := alpruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldUpdatedAt,
		})
	}
	if value, ok := alpruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldDeletedAt,
		})
	}
	if value, ok := alpruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: applegacypowerrental.FieldDeletedAt,
		})
	}
	if value, ok := alpruo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applegacypowerrental.FieldEntID,
		})
	}
	if value, ok := alpruo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: applegacypowerrental.FieldAppGoodID,
		})
	}
	if alpruo.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: applegacypowerrental.FieldAppGoodID,
		})
	}
	if value, ok := alpruo.mutation.TechniqueFeeRatio(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: applegacypowerrental.FieldTechniqueFeeRatio,
		})
	}
	if alpruo.mutation.TechniqueFeeRatioCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: applegacypowerrental.FieldTechniqueFeeRatio,
		})
	}
	_spec.Modifiers = alpruo.modifiers
	_node = &AppLegacyPowerRental{config: alpruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, alpruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{applegacypowerrental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}