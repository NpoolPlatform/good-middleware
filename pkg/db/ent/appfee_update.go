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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppFeeUpdate is the builder for updating AppFee entities.
type AppFeeUpdate struct {
	config
	hooks     []Hook
	mutation  *AppFeeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppFeeUpdate builder.
func (afu *AppFeeUpdate) Where(ps ...predicate.AppFee) *AppFeeUpdate {
	afu.mutation.Where(ps...)
	return afu
}

// SetCreatedAt sets the "created_at" field.
func (afu *AppFeeUpdate) SetCreatedAt(u uint32) *AppFeeUpdate {
	afu.mutation.ResetCreatedAt()
	afu.mutation.SetCreatedAt(u)
	return afu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (afu *AppFeeUpdate) SetNillableCreatedAt(u *uint32) *AppFeeUpdate {
	if u != nil {
		afu.SetCreatedAt(*u)
	}
	return afu
}

// AddCreatedAt adds u to the "created_at" field.
func (afu *AppFeeUpdate) AddCreatedAt(u int32) *AppFeeUpdate {
	afu.mutation.AddCreatedAt(u)
	return afu
}

// SetUpdatedAt sets the "updated_at" field.
func (afu *AppFeeUpdate) SetUpdatedAt(u uint32) *AppFeeUpdate {
	afu.mutation.ResetUpdatedAt()
	afu.mutation.SetUpdatedAt(u)
	return afu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (afu *AppFeeUpdate) AddUpdatedAt(u int32) *AppFeeUpdate {
	afu.mutation.AddUpdatedAt(u)
	return afu
}

// SetDeletedAt sets the "deleted_at" field.
func (afu *AppFeeUpdate) SetDeletedAt(u uint32) *AppFeeUpdate {
	afu.mutation.ResetDeletedAt()
	afu.mutation.SetDeletedAt(u)
	return afu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (afu *AppFeeUpdate) SetNillableDeletedAt(u *uint32) *AppFeeUpdate {
	if u != nil {
		afu.SetDeletedAt(*u)
	}
	return afu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (afu *AppFeeUpdate) AddDeletedAt(u int32) *AppFeeUpdate {
	afu.mutation.AddDeletedAt(u)
	return afu
}

// SetEntID sets the "ent_id" field.
func (afu *AppFeeUpdate) SetEntID(u uuid.UUID) *AppFeeUpdate {
	afu.mutation.SetEntID(u)
	return afu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (afu *AppFeeUpdate) SetNillableEntID(u *uuid.UUID) *AppFeeUpdate {
	if u != nil {
		afu.SetEntID(*u)
	}
	return afu
}

// SetAppGoodID sets the "app_good_id" field.
func (afu *AppFeeUpdate) SetAppGoodID(u uuid.UUID) *AppFeeUpdate {
	afu.mutation.SetAppGoodID(u)
	return afu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (afu *AppFeeUpdate) SetNillableAppGoodID(u *uuid.UUID) *AppFeeUpdate {
	if u != nil {
		afu.SetAppGoodID(*u)
	}
	return afu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (afu *AppFeeUpdate) ClearAppGoodID() *AppFeeUpdate {
	afu.mutation.ClearAppGoodID()
	return afu
}

// SetUnitValue sets the "unit_value" field.
func (afu *AppFeeUpdate) SetUnitValue(d decimal.Decimal) *AppFeeUpdate {
	afu.mutation.SetUnitValue(d)
	return afu
}

// SetNillableUnitValue sets the "unit_value" field if the given value is not nil.
func (afu *AppFeeUpdate) SetNillableUnitValue(d *decimal.Decimal) *AppFeeUpdate {
	if d != nil {
		afu.SetUnitValue(*d)
	}
	return afu
}

// ClearUnitValue clears the value of the "unit_value" field.
func (afu *AppFeeUpdate) ClearUnitValue() *AppFeeUpdate {
	afu.mutation.ClearUnitValue()
	return afu
}

// SetCancelMode sets the "cancel_mode" field.
func (afu *AppFeeUpdate) SetCancelMode(s string) *AppFeeUpdate {
	afu.mutation.SetCancelMode(s)
	return afu
}

// SetNillableCancelMode sets the "cancel_mode" field if the given value is not nil.
func (afu *AppFeeUpdate) SetNillableCancelMode(s *string) *AppFeeUpdate {
	if s != nil {
		afu.SetCancelMode(*s)
	}
	return afu
}

// ClearCancelMode clears the value of the "cancel_mode" field.
func (afu *AppFeeUpdate) ClearCancelMode() *AppFeeUpdate {
	afu.mutation.ClearCancelMode()
	return afu
}

// SetMinOrderDurationSeconds sets the "min_order_duration_seconds" field.
func (afu *AppFeeUpdate) SetMinOrderDurationSeconds(u uint32) *AppFeeUpdate {
	afu.mutation.ResetMinOrderDurationSeconds()
	afu.mutation.SetMinOrderDurationSeconds(u)
	return afu
}

// SetNillableMinOrderDurationSeconds sets the "min_order_duration_seconds" field if the given value is not nil.
func (afu *AppFeeUpdate) SetNillableMinOrderDurationSeconds(u *uint32) *AppFeeUpdate {
	if u != nil {
		afu.SetMinOrderDurationSeconds(*u)
	}
	return afu
}

// AddMinOrderDurationSeconds adds u to the "min_order_duration_seconds" field.
func (afu *AppFeeUpdate) AddMinOrderDurationSeconds(u int32) *AppFeeUpdate {
	afu.mutation.AddMinOrderDurationSeconds(u)
	return afu
}

// ClearMinOrderDurationSeconds clears the value of the "min_order_duration_seconds" field.
func (afu *AppFeeUpdate) ClearMinOrderDurationSeconds() *AppFeeUpdate {
	afu.mutation.ClearMinOrderDurationSeconds()
	return afu
}

// Mutation returns the AppFeeMutation object of the builder.
func (afu *AppFeeUpdate) Mutation() *AppFeeMutation {
	return afu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (afu *AppFeeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := afu.defaults(); err != nil {
		return 0, err
	}
	if len(afu.hooks) == 0 {
		affected, err = afu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppFeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			afu.mutation = mutation
			affected, err = afu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(afu.hooks) - 1; i >= 0; i-- {
			if afu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = afu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, afu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (afu *AppFeeUpdate) SaveX(ctx context.Context) int {
	affected, err := afu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (afu *AppFeeUpdate) Exec(ctx context.Context) error {
	_, err := afu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afu *AppFeeUpdate) ExecX(ctx context.Context) {
	if err := afu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (afu *AppFeeUpdate) defaults() error {
	if _, ok := afu.mutation.UpdatedAt(); !ok {
		if appfee.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appfee.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appfee.UpdateDefaultUpdatedAt()
		afu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (afu *AppFeeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppFeeUpdate {
	afu.modifiers = append(afu.modifiers, modifiers...)
	return afu
}

func (afu *AppFeeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appfee.Table,
			Columns: appfee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appfee.FieldID,
			},
		},
	}
	if ps := afu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldCreatedAt,
		})
	}
	if value, ok := afu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldCreatedAt,
		})
	}
	if value, ok := afu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldUpdatedAt,
		})
	}
	if value, ok := afu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldUpdatedAt,
		})
	}
	if value, ok := afu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldDeletedAt,
		})
	}
	if value, ok := afu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldDeletedAt,
		})
	}
	if value, ok := afu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appfee.FieldEntID,
		})
	}
	if value, ok := afu.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appfee.FieldAppGoodID,
		})
	}
	if afu.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appfee.FieldAppGoodID,
		})
	}
	if value, ok := afu.mutation.UnitValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appfee.FieldUnitValue,
		})
	}
	if afu.mutation.UnitValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appfee.FieldUnitValue,
		})
	}
	if value, ok := afu.mutation.CancelMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appfee.FieldCancelMode,
		})
	}
	if afu.mutation.CancelModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appfee.FieldCancelMode,
		})
	}
	if value, ok := afu.mutation.MinOrderDurationSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldMinOrderDurationSeconds,
		})
	}
	if value, ok := afu.mutation.AddedMinOrderDurationSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldMinOrderDurationSeconds,
		})
	}
	if afu.mutation.MinOrderDurationSecondsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: appfee.FieldMinOrderDurationSeconds,
		})
	}
	_spec.Modifiers = afu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, afu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appfee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppFeeUpdateOne is the builder for updating a single AppFee entity.
type AppFeeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppFeeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (afuo *AppFeeUpdateOne) SetCreatedAt(u uint32) *AppFeeUpdateOne {
	afuo.mutation.ResetCreatedAt()
	afuo.mutation.SetCreatedAt(u)
	return afuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (afuo *AppFeeUpdateOne) SetNillableCreatedAt(u *uint32) *AppFeeUpdateOne {
	if u != nil {
		afuo.SetCreatedAt(*u)
	}
	return afuo
}

// AddCreatedAt adds u to the "created_at" field.
func (afuo *AppFeeUpdateOne) AddCreatedAt(u int32) *AppFeeUpdateOne {
	afuo.mutation.AddCreatedAt(u)
	return afuo
}

// SetUpdatedAt sets the "updated_at" field.
func (afuo *AppFeeUpdateOne) SetUpdatedAt(u uint32) *AppFeeUpdateOne {
	afuo.mutation.ResetUpdatedAt()
	afuo.mutation.SetUpdatedAt(u)
	return afuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (afuo *AppFeeUpdateOne) AddUpdatedAt(u int32) *AppFeeUpdateOne {
	afuo.mutation.AddUpdatedAt(u)
	return afuo
}

// SetDeletedAt sets the "deleted_at" field.
func (afuo *AppFeeUpdateOne) SetDeletedAt(u uint32) *AppFeeUpdateOne {
	afuo.mutation.ResetDeletedAt()
	afuo.mutation.SetDeletedAt(u)
	return afuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (afuo *AppFeeUpdateOne) SetNillableDeletedAt(u *uint32) *AppFeeUpdateOne {
	if u != nil {
		afuo.SetDeletedAt(*u)
	}
	return afuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (afuo *AppFeeUpdateOne) AddDeletedAt(u int32) *AppFeeUpdateOne {
	afuo.mutation.AddDeletedAt(u)
	return afuo
}

// SetEntID sets the "ent_id" field.
func (afuo *AppFeeUpdateOne) SetEntID(u uuid.UUID) *AppFeeUpdateOne {
	afuo.mutation.SetEntID(u)
	return afuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (afuo *AppFeeUpdateOne) SetNillableEntID(u *uuid.UUID) *AppFeeUpdateOne {
	if u != nil {
		afuo.SetEntID(*u)
	}
	return afuo
}

// SetAppGoodID sets the "app_good_id" field.
func (afuo *AppFeeUpdateOne) SetAppGoodID(u uuid.UUID) *AppFeeUpdateOne {
	afuo.mutation.SetAppGoodID(u)
	return afuo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (afuo *AppFeeUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *AppFeeUpdateOne {
	if u != nil {
		afuo.SetAppGoodID(*u)
	}
	return afuo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (afuo *AppFeeUpdateOne) ClearAppGoodID() *AppFeeUpdateOne {
	afuo.mutation.ClearAppGoodID()
	return afuo
}

// SetUnitValue sets the "unit_value" field.
func (afuo *AppFeeUpdateOne) SetUnitValue(d decimal.Decimal) *AppFeeUpdateOne {
	afuo.mutation.SetUnitValue(d)
	return afuo
}

// SetNillableUnitValue sets the "unit_value" field if the given value is not nil.
func (afuo *AppFeeUpdateOne) SetNillableUnitValue(d *decimal.Decimal) *AppFeeUpdateOne {
	if d != nil {
		afuo.SetUnitValue(*d)
	}
	return afuo
}

// ClearUnitValue clears the value of the "unit_value" field.
func (afuo *AppFeeUpdateOne) ClearUnitValue() *AppFeeUpdateOne {
	afuo.mutation.ClearUnitValue()
	return afuo
}

// SetCancelMode sets the "cancel_mode" field.
func (afuo *AppFeeUpdateOne) SetCancelMode(s string) *AppFeeUpdateOne {
	afuo.mutation.SetCancelMode(s)
	return afuo
}

// SetNillableCancelMode sets the "cancel_mode" field if the given value is not nil.
func (afuo *AppFeeUpdateOne) SetNillableCancelMode(s *string) *AppFeeUpdateOne {
	if s != nil {
		afuo.SetCancelMode(*s)
	}
	return afuo
}

// ClearCancelMode clears the value of the "cancel_mode" field.
func (afuo *AppFeeUpdateOne) ClearCancelMode() *AppFeeUpdateOne {
	afuo.mutation.ClearCancelMode()
	return afuo
}

// SetMinOrderDurationSeconds sets the "min_order_duration_seconds" field.
func (afuo *AppFeeUpdateOne) SetMinOrderDurationSeconds(u uint32) *AppFeeUpdateOne {
	afuo.mutation.ResetMinOrderDurationSeconds()
	afuo.mutation.SetMinOrderDurationSeconds(u)
	return afuo
}

// SetNillableMinOrderDurationSeconds sets the "min_order_duration_seconds" field if the given value is not nil.
func (afuo *AppFeeUpdateOne) SetNillableMinOrderDurationSeconds(u *uint32) *AppFeeUpdateOne {
	if u != nil {
		afuo.SetMinOrderDurationSeconds(*u)
	}
	return afuo
}

// AddMinOrderDurationSeconds adds u to the "min_order_duration_seconds" field.
func (afuo *AppFeeUpdateOne) AddMinOrderDurationSeconds(u int32) *AppFeeUpdateOne {
	afuo.mutation.AddMinOrderDurationSeconds(u)
	return afuo
}

// ClearMinOrderDurationSeconds clears the value of the "min_order_duration_seconds" field.
func (afuo *AppFeeUpdateOne) ClearMinOrderDurationSeconds() *AppFeeUpdateOne {
	afuo.mutation.ClearMinOrderDurationSeconds()
	return afuo
}

// Mutation returns the AppFeeMutation object of the builder.
func (afuo *AppFeeUpdateOne) Mutation() *AppFeeMutation {
	return afuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (afuo *AppFeeUpdateOne) Select(field string, fields ...string) *AppFeeUpdateOne {
	afuo.fields = append([]string{field}, fields...)
	return afuo
}

// Save executes the query and returns the updated AppFee entity.
func (afuo *AppFeeUpdateOne) Save(ctx context.Context) (*AppFee, error) {
	var (
		err  error
		node *AppFee
	)
	if err := afuo.defaults(); err != nil {
		return nil, err
	}
	if len(afuo.hooks) == 0 {
		node, err = afuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppFeeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			afuo.mutation = mutation
			node, err = afuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(afuo.hooks) - 1; i >= 0; i-- {
			if afuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = afuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, afuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (afuo *AppFeeUpdateOne) SaveX(ctx context.Context) *AppFee {
	node, err := afuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (afuo *AppFeeUpdateOne) Exec(ctx context.Context) error {
	_, err := afuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (afuo *AppFeeUpdateOne) ExecX(ctx context.Context) {
	if err := afuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (afuo *AppFeeUpdateOne) defaults() error {
	if _, ok := afuo.mutation.UpdatedAt(); !ok {
		if appfee.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appfee.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appfee.UpdateDefaultUpdatedAt()
		afuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (afuo *AppFeeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppFeeUpdateOne {
	afuo.modifiers = append(afuo.modifiers, modifiers...)
	return afuo
}

func (afuo *AppFeeUpdateOne) sqlSave(ctx context.Context) (_node *AppFee, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appfee.Table,
			Columns: appfee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appfee.FieldID,
			},
		},
	}
	id, ok := afuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppFee.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := afuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appfee.FieldID)
		for _, f := range fields {
			if !appfee.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appfee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := afuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := afuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldCreatedAt,
		})
	}
	if value, ok := afuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldCreatedAt,
		})
	}
	if value, ok := afuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldUpdatedAt,
		})
	}
	if value, ok := afuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldUpdatedAt,
		})
	}
	if value, ok := afuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldDeletedAt,
		})
	}
	if value, ok := afuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldDeletedAt,
		})
	}
	if value, ok := afuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appfee.FieldEntID,
		})
	}
	if value, ok := afuo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appfee.FieldAppGoodID,
		})
	}
	if afuo.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appfee.FieldAppGoodID,
		})
	}
	if value, ok := afuo.mutation.UnitValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appfee.FieldUnitValue,
		})
	}
	if afuo.mutation.UnitValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: appfee.FieldUnitValue,
		})
	}
	if value, ok := afuo.mutation.CancelMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appfee.FieldCancelMode,
		})
	}
	if afuo.mutation.CancelModeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appfee.FieldCancelMode,
		})
	}
	if value, ok := afuo.mutation.MinOrderDurationSeconds(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldMinOrderDurationSeconds,
		})
	}
	if value, ok := afuo.mutation.AddedMinOrderDurationSeconds(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appfee.FieldMinOrderDurationSeconds,
		})
	}
	if afuo.mutation.MinOrderDurationSecondsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: appfee.FieldMinOrderDurationSeconds,
		})
	}
	_spec.Modifiers = afuo.modifiers
	_node = &AppFee{config: afuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, afuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appfee.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
