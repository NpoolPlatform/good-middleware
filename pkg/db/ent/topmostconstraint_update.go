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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostconstraint"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TopMostConstraintUpdate is the builder for updating TopMostConstraint entities.
type TopMostConstraintUpdate struct {
	config
	hooks     []Hook
	mutation  *TopMostConstraintMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TopMostConstraintUpdate builder.
func (tmcu *TopMostConstraintUpdate) Where(ps ...predicate.TopMostConstraint) *TopMostConstraintUpdate {
	tmcu.mutation.Where(ps...)
	return tmcu
}

// SetCreatedAt sets the "created_at" field.
func (tmcu *TopMostConstraintUpdate) SetCreatedAt(u uint32) *TopMostConstraintUpdate {
	tmcu.mutation.ResetCreatedAt()
	tmcu.mutation.SetCreatedAt(u)
	return tmcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmcu *TopMostConstraintUpdate) SetNillableCreatedAt(u *uint32) *TopMostConstraintUpdate {
	if u != nil {
		tmcu.SetCreatedAt(*u)
	}
	return tmcu
}

// AddCreatedAt adds u to the "created_at" field.
func (tmcu *TopMostConstraintUpdate) AddCreatedAt(u int32) *TopMostConstraintUpdate {
	tmcu.mutation.AddCreatedAt(u)
	return tmcu
}

// SetUpdatedAt sets the "updated_at" field.
func (tmcu *TopMostConstraintUpdate) SetUpdatedAt(u uint32) *TopMostConstraintUpdate {
	tmcu.mutation.ResetUpdatedAt()
	tmcu.mutation.SetUpdatedAt(u)
	return tmcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tmcu *TopMostConstraintUpdate) AddUpdatedAt(u int32) *TopMostConstraintUpdate {
	tmcu.mutation.AddUpdatedAt(u)
	return tmcu
}

// SetDeletedAt sets the "deleted_at" field.
func (tmcu *TopMostConstraintUpdate) SetDeletedAt(u uint32) *TopMostConstraintUpdate {
	tmcu.mutation.ResetDeletedAt()
	tmcu.mutation.SetDeletedAt(u)
	return tmcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmcu *TopMostConstraintUpdate) SetNillableDeletedAt(u *uint32) *TopMostConstraintUpdate {
	if u != nil {
		tmcu.SetDeletedAt(*u)
	}
	return tmcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tmcu *TopMostConstraintUpdate) AddDeletedAt(u int32) *TopMostConstraintUpdate {
	tmcu.mutation.AddDeletedAt(u)
	return tmcu
}

// SetEntID sets the "ent_id" field.
func (tmcu *TopMostConstraintUpdate) SetEntID(u uuid.UUID) *TopMostConstraintUpdate {
	tmcu.mutation.SetEntID(u)
	return tmcu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmcu *TopMostConstraintUpdate) SetNillableEntID(u *uuid.UUID) *TopMostConstraintUpdate {
	if u != nil {
		tmcu.SetEntID(*u)
	}
	return tmcu
}

// SetTopMostID sets the "top_most_id" field.
func (tmcu *TopMostConstraintUpdate) SetTopMostID(u uuid.UUID) *TopMostConstraintUpdate {
	tmcu.mutation.SetTopMostID(u)
	return tmcu
}

// SetNillableTopMostID sets the "top_most_id" field if the given value is not nil.
func (tmcu *TopMostConstraintUpdate) SetNillableTopMostID(u *uuid.UUID) *TopMostConstraintUpdate {
	if u != nil {
		tmcu.SetTopMostID(*u)
	}
	return tmcu
}

// ClearTopMostID clears the value of the "top_most_id" field.
func (tmcu *TopMostConstraintUpdate) ClearTopMostID() *TopMostConstraintUpdate {
	tmcu.mutation.ClearTopMostID()
	return tmcu
}

// SetConstraint sets the "constraint" field.
func (tmcu *TopMostConstraintUpdate) SetConstraint(s string) *TopMostConstraintUpdate {
	tmcu.mutation.SetConstraint(s)
	return tmcu
}

// SetNillableConstraint sets the "constraint" field if the given value is not nil.
func (tmcu *TopMostConstraintUpdate) SetNillableConstraint(s *string) *TopMostConstraintUpdate {
	if s != nil {
		tmcu.SetConstraint(*s)
	}
	return tmcu
}

// ClearConstraint clears the value of the "constraint" field.
func (tmcu *TopMostConstraintUpdate) ClearConstraint() *TopMostConstraintUpdate {
	tmcu.mutation.ClearConstraint()
	return tmcu
}

// SetTargetValue sets the "target_value" field.
func (tmcu *TopMostConstraintUpdate) SetTargetValue(d decimal.Decimal) *TopMostConstraintUpdate {
	tmcu.mutation.SetTargetValue(d)
	return tmcu
}

// SetNillableTargetValue sets the "target_value" field if the given value is not nil.
func (tmcu *TopMostConstraintUpdate) SetNillableTargetValue(d *decimal.Decimal) *TopMostConstraintUpdate {
	if d != nil {
		tmcu.SetTargetValue(*d)
	}
	return tmcu
}

// ClearTargetValue clears the value of the "target_value" field.
func (tmcu *TopMostConstraintUpdate) ClearTargetValue() *TopMostConstraintUpdate {
	tmcu.mutation.ClearTargetValue()
	return tmcu
}

// SetIndex sets the "index" field.
func (tmcu *TopMostConstraintUpdate) SetIndex(u uint8) *TopMostConstraintUpdate {
	tmcu.mutation.ResetIndex()
	tmcu.mutation.SetIndex(u)
	return tmcu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tmcu *TopMostConstraintUpdate) SetNillableIndex(u *uint8) *TopMostConstraintUpdate {
	if u != nil {
		tmcu.SetIndex(*u)
	}
	return tmcu
}

// AddIndex adds u to the "index" field.
func (tmcu *TopMostConstraintUpdate) AddIndex(u int8) *TopMostConstraintUpdate {
	tmcu.mutation.AddIndex(u)
	return tmcu
}

// ClearIndex clears the value of the "index" field.
func (tmcu *TopMostConstraintUpdate) ClearIndex() *TopMostConstraintUpdate {
	tmcu.mutation.ClearIndex()
	return tmcu
}

// Mutation returns the TopMostConstraintMutation object of the builder.
func (tmcu *TopMostConstraintUpdate) Mutation() *TopMostConstraintMutation {
	return tmcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tmcu *TopMostConstraintUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := tmcu.defaults(); err != nil {
		return 0, err
	}
	if len(tmcu.hooks) == 0 {
		affected, err = tmcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopMostConstraintMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmcu.mutation = mutation
			affected, err = tmcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tmcu.hooks) - 1; i >= 0; i-- {
			if tmcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmcu *TopMostConstraintUpdate) SaveX(ctx context.Context) int {
	affected, err := tmcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tmcu *TopMostConstraintUpdate) Exec(ctx context.Context) error {
	_, err := tmcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcu *TopMostConstraintUpdate) ExecX(ctx context.Context) {
	if err := tmcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmcu *TopMostConstraintUpdate) defaults() error {
	if _, ok := tmcu.mutation.UpdatedAt(); !ok {
		if topmostconstraint.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostconstraint.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topmostconstraint.UpdateDefaultUpdatedAt()
		tmcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmcu *TopMostConstraintUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TopMostConstraintUpdate {
	tmcu.modifiers = append(tmcu.modifiers, modifiers...)
	return tmcu
}

func (tmcu *TopMostConstraintUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topmostconstraint.Table,
			Columns: topmostconstraint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostconstraint.FieldID,
			},
		},
	}
	if ps := tmcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldCreatedAt,
		})
	}
	if value, ok := tmcu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldCreatedAt,
		})
	}
	if value, ok := tmcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldUpdatedAt,
		})
	}
	if value, ok := tmcu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldUpdatedAt,
		})
	}
	if value, ok := tmcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldDeletedAt,
		})
	}
	if value, ok := tmcu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldDeletedAt,
		})
	}
	if value, ok := tmcu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostconstraint.FieldEntID,
		})
	}
	if value, ok := tmcu.mutation.TopMostID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostconstraint.FieldTopMostID,
		})
	}
	if tmcu.mutation.TopMostIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: topmostconstraint.FieldTopMostID,
		})
	}
	if value, ok := tmcu.mutation.Constraint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topmostconstraint.FieldConstraint,
		})
	}
	if tmcu.mutation.ConstraintCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: topmostconstraint.FieldConstraint,
		})
	}
	if value, ok := tmcu.mutation.TargetValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: topmostconstraint.FieldTargetValue,
		})
	}
	if tmcu.mutation.TargetValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: topmostconstraint.FieldTargetValue,
		})
	}
	if value, ok := tmcu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: topmostconstraint.FieldIndex,
		})
	}
	if value, ok := tmcu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: topmostconstraint.FieldIndex,
		})
	}
	if tmcu.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: topmostconstraint.FieldIndex,
		})
	}
	_spec.Modifiers = tmcu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, tmcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topmostconstraint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// TopMostConstraintUpdateOne is the builder for updating a single TopMostConstraint entity.
type TopMostConstraintUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TopMostConstraintMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (tmcuo *TopMostConstraintUpdateOne) SetCreatedAt(u uint32) *TopMostConstraintUpdateOne {
	tmcuo.mutation.ResetCreatedAt()
	tmcuo.mutation.SetCreatedAt(u)
	return tmcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmcuo *TopMostConstraintUpdateOne) SetNillableCreatedAt(u *uint32) *TopMostConstraintUpdateOne {
	if u != nil {
		tmcuo.SetCreatedAt(*u)
	}
	return tmcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (tmcuo *TopMostConstraintUpdateOne) AddCreatedAt(u int32) *TopMostConstraintUpdateOne {
	tmcuo.mutation.AddCreatedAt(u)
	return tmcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (tmcuo *TopMostConstraintUpdateOne) SetUpdatedAt(u uint32) *TopMostConstraintUpdateOne {
	tmcuo.mutation.ResetUpdatedAt()
	tmcuo.mutation.SetUpdatedAt(u)
	return tmcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (tmcuo *TopMostConstraintUpdateOne) AddUpdatedAt(u int32) *TopMostConstraintUpdateOne {
	tmcuo.mutation.AddUpdatedAt(u)
	return tmcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tmcuo *TopMostConstraintUpdateOne) SetDeletedAt(u uint32) *TopMostConstraintUpdateOne {
	tmcuo.mutation.ResetDeletedAt()
	tmcuo.mutation.SetDeletedAt(u)
	return tmcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmcuo *TopMostConstraintUpdateOne) SetNillableDeletedAt(u *uint32) *TopMostConstraintUpdateOne {
	if u != nil {
		tmcuo.SetDeletedAt(*u)
	}
	return tmcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (tmcuo *TopMostConstraintUpdateOne) AddDeletedAt(u int32) *TopMostConstraintUpdateOne {
	tmcuo.mutation.AddDeletedAt(u)
	return tmcuo
}

// SetEntID sets the "ent_id" field.
func (tmcuo *TopMostConstraintUpdateOne) SetEntID(u uuid.UUID) *TopMostConstraintUpdateOne {
	tmcuo.mutation.SetEntID(u)
	return tmcuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmcuo *TopMostConstraintUpdateOne) SetNillableEntID(u *uuid.UUID) *TopMostConstraintUpdateOne {
	if u != nil {
		tmcuo.SetEntID(*u)
	}
	return tmcuo
}

// SetTopMostID sets the "top_most_id" field.
func (tmcuo *TopMostConstraintUpdateOne) SetTopMostID(u uuid.UUID) *TopMostConstraintUpdateOne {
	tmcuo.mutation.SetTopMostID(u)
	return tmcuo
}

// SetNillableTopMostID sets the "top_most_id" field if the given value is not nil.
func (tmcuo *TopMostConstraintUpdateOne) SetNillableTopMostID(u *uuid.UUID) *TopMostConstraintUpdateOne {
	if u != nil {
		tmcuo.SetTopMostID(*u)
	}
	return tmcuo
}

// ClearTopMostID clears the value of the "top_most_id" field.
func (tmcuo *TopMostConstraintUpdateOne) ClearTopMostID() *TopMostConstraintUpdateOne {
	tmcuo.mutation.ClearTopMostID()
	return tmcuo
}

// SetConstraint sets the "constraint" field.
func (tmcuo *TopMostConstraintUpdateOne) SetConstraint(s string) *TopMostConstraintUpdateOne {
	tmcuo.mutation.SetConstraint(s)
	return tmcuo
}

// SetNillableConstraint sets the "constraint" field if the given value is not nil.
func (tmcuo *TopMostConstraintUpdateOne) SetNillableConstraint(s *string) *TopMostConstraintUpdateOne {
	if s != nil {
		tmcuo.SetConstraint(*s)
	}
	return tmcuo
}

// ClearConstraint clears the value of the "constraint" field.
func (tmcuo *TopMostConstraintUpdateOne) ClearConstraint() *TopMostConstraintUpdateOne {
	tmcuo.mutation.ClearConstraint()
	return tmcuo
}

// SetTargetValue sets the "target_value" field.
func (tmcuo *TopMostConstraintUpdateOne) SetTargetValue(d decimal.Decimal) *TopMostConstraintUpdateOne {
	tmcuo.mutation.SetTargetValue(d)
	return tmcuo
}

// SetNillableTargetValue sets the "target_value" field if the given value is not nil.
func (tmcuo *TopMostConstraintUpdateOne) SetNillableTargetValue(d *decimal.Decimal) *TopMostConstraintUpdateOne {
	if d != nil {
		tmcuo.SetTargetValue(*d)
	}
	return tmcuo
}

// ClearTargetValue clears the value of the "target_value" field.
func (tmcuo *TopMostConstraintUpdateOne) ClearTargetValue() *TopMostConstraintUpdateOne {
	tmcuo.mutation.ClearTargetValue()
	return tmcuo
}

// SetIndex sets the "index" field.
func (tmcuo *TopMostConstraintUpdateOne) SetIndex(u uint8) *TopMostConstraintUpdateOne {
	tmcuo.mutation.ResetIndex()
	tmcuo.mutation.SetIndex(u)
	return tmcuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tmcuo *TopMostConstraintUpdateOne) SetNillableIndex(u *uint8) *TopMostConstraintUpdateOne {
	if u != nil {
		tmcuo.SetIndex(*u)
	}
	return tmcuo
}

// AddIndex adds u to the "index" field.
func (tmcuo *TopMostConstraintUpdateOne) AddIndex(u int8) *TopMostConstraintUpdateOne {
	tmcuo.mutation.AddIndex(u)
	return tmcuo
}

// ClearIndex clears the value of the "index" field.
func (tmcuo *TopMostConstraintUpdateOne) ClearIndex() *TopMostConstraintUpdateOne {
	tmcuo.mutation.ClearIndex()
	return tmcuo
}

// Mutation returns the TopMostConstraintMutation object of the builder.
func (tmcuo *TopMostConstraintUpdateOne) Mutation() *TopMostConstraintMutation {
	return tmcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tmcuo *TopMostConstraintUpdateOne) Select(field string, fields ...string) *TopMostConstraintUpdateOne {
	tmcuo.fields = append([]string{field}, fields...)
	return tmcuo
}

// Save executes the query and returns the updated TopMostConstraint entity.
func (tmcuo *TopMostConstraintUpdateOne) Save(ctx context.Context) (*TopMostConstraint, error) {
	var (
		err  error
		node *TopMostConstraint
	)
	if err := tmcuo.defaults(); err != nil {
		return nil, err
	}
	if len(tmcuo.hooks) == 0 {
		node, err = tmcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopMostConstraintMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmcuo.mutation = mutation
			node, err = tmcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tmcuo.hooks) - 1; i >= 0; i-- {
			if tmcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tmcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TopMostConstraint)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TopMostConstraintMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tmcuo *TopMostConstraintUpdateOne) SaveX(ctx context.Context) *TopMostConstraint {
	node, err := tmcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tmcuo *TopMostConstraintUpdateOne) Exec(ctx context.Context) error {
	_, err := tmcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcuo *TopMostConstraintUpdateOne) ExecX(ctx context.Context) {
	if err := tmcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmcuo *TopMostConstraintUpdateOne) defaults() error {
	if _, ok := tmcuo.mutation.UpdatedAt(); !ok {
		if topmostconstraint.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostconstraint.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topmostconstraint.UpdateDefaultUpdatedAt()
		tmcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tmcuo *TopMostConstraintUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TopMostConstraintUpdateOne {
	tmcuo.modifiers = append(tmcuo.modifiers, modifiers...)
	return tmcuo
}

func (tmcuo *TopMostConstraintUpdateOne) sqlSave(ctx context.Context) (_node *TopMostConstraint, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topmostconstraint.Table,
			Columns: topmostconstraint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostconstraint.FieldID,
			},
		},
	}
	id, ok := tmcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TopMostConstraint.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tmcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topmostconstraint.FieldID)
		for _, f := range fields {
			if !topmostconstraint.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != topmostconstraint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tmcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tmcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldCreatedAt,
		})
	}
	if value, ok := tmcuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldCreatedAt,
		})
	}
	if value, ok := tmcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldUpdatedAt,
		})
	}
	if value, ok := tmcuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldUpdatedAt,
		})
	}
	if value, ok := tmcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldDeletedAt,
		})
	}
	if value, ok := tmcuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldDeletedAt,
		})
	}
	if value, ok := tmcuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostconstraint.FieldEntID,
		})
	}
	if value, ok := tmcuo.mutation.TopMostID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostconstraint.FieldTopMostID,
		})
	}
	if tmcuo.mutation.TopMostIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: topmostconstraint.FieldTopMostID,
		})
	}
	if value, ok := tmcuo.mutation.Constraint(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topmostconstraint.FieldConstraint,
		})
	}
	if tmcuo.mutation.ConstraintCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: topmostconstraint.FieldConstraint,
		})
	}
	if value, ok := tmcuo.mutation.TargetValue(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: topmostconstraint.FieldTargetValue,
		})
	}
	if tmcuo.mutation.TargetValueCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: topmostconstraint.FieldTargetValue,
		})
	}
	if value, ok := tmcuo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: topmostconstraint.FieldIndex,
		})
	}
	if value, ok := tmcuo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: topmostconstraint.FieldIndex,
		})
	}
	if tmcuo.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: topmostconstraint.FieldIndex,
		})
	}
	_spec.Modifiers = tmcuo.modifiers
	_node = &TopMostConstraint{config: tmcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tmcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{topmostconstraint.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
