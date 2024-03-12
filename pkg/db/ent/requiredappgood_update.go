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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredappgood"
	"github.com/google/uuid"
)

// RequiredAppGoodUpdate is the builder for updating RequiredAppGood entities.
type RequiredAppGoodUpdate struct {
	config
	hooks     []Hook
	mutation  *RequiredAppGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the RequiredAppGoodUpdate builder.
func (ragu *RequiredAppGoodUpdate) Where(ps ...predicate.RequiredAppGood) *RequiredAppGoodUpdate {
	ragu.mutation.Where(ps...)
	return ragu
}

// SetCreatedAt sets the "created_at" field.
func (ragu *RequiredAppGoodUpdate) SetCreatedAt(u uint32) *RequiredAppGoodUpdate {
	ragu.mutation.ResetCreatedAt()
	ragu.mutation.SetCreatedAt(u)
	return ragu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ragu *RequiredAppGoodUpdate) SetNillableCreatedAt(u *uint32) *RequiredAppGoodUpdate {
	if u != nil {
		ragu.SetCreatedAt(*u)
	}
	return ragu
}

// AddCreatedAt adds u to the "created_at" field.
func (ragu *RequiredAppGoodUpdate) AddCreatedAt(u int32) *RequiredAppGoodUpdate {
	ragu.mutation.AddCreatedAt(u)
	return ragu
}

// SetUpdatedAt sets the "updated_at" field.
func (ragu *RequiredAppGoodUpdate) SetUpdatedAt(u uint32) *RequiredAppGoodUpdate {
	ragu.mutation.ResetUpdatedAt()
	ragu.mutation.SetUpdatedAt(u)
	return ragu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (ragu *RequiredAppGoodUpdate) AddUpdatedAt(u int32) *RequiredAppGoodUpdate {
	ragu.mutation.AddUpdatedAt(u)
	return ragu
}

// SetDeletedAt sets the "deleted_at" field.
func (ragu *RequiredAppGoodUpdate) SetDeletedAt(u uint32) *RequiredAppGoodUpdate {
	ragu.mutation.ResetDeletedAt()
	ragu.mutation.SetDeletedAt(u)
	return ragu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ragu *RequiredAppGoodUpdate) SetNillableDeletedAt(u *uint32) *RequiredAppGoodUpdate {
	if u != nil {
		ragu.SetDeletedAt(*u)
	}
	return ragu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (ragu *RequiredAppGoodUpdate) AddDeletedAt(u int32) *RequiredAppGoodUpdate {
	ragu.mutation.AddDeletedAt(u)
	return ragu
}

// SetEntID sets the "ent_id" field.
func (ragu *RequiredAppGoodUpdate) SetEntID(u uuid.UUID) *RequiredAppGoodUpdate {
	ragu.mutation.SetEntID(u)
	return ragu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (ragu *RequiredAppGoodUpdate) SetNillableEntID(u *uuid.UUID) *RequiredAppGoodUpdate {
	if u != nil {
		ragu.SetEntID(*u)
	}
	return ragu
}

// SetMainAppGoodID sets the "main_app_good_id" field.
func (ragu *RequiredAppGoodUpdate) SetMainAppGoodID(u uuid.UUID) *RequiredAppGoodUpdate {
	ragu.mutation.SetMainAppGoodID(u)
	return ragu
}

// SetRequiredAppGoodID sets the "required_app_good_id" field.
func (ragu *RequiredAppGoodUpdate) SetRequiredAppGoodID(u uuid.UUID) *RequiredAppGoodUpdate {
	ragu.mutation.SetRequiredAppGoodID(u)
	return ragu
}

// SetMust sets the "must" field.
func (ragu *RequiredAppGoodUpdate) SetMust(b bool) *RequiredAppGoodUpdate {
	ragu.mutation.SetMust(b)
	return ragu
}

// SetNillableMust sets the "must" field if the given value is not nil.
func (ragu *RequiredAppGoodUpdate) SetNillableMust(b *bool) *RequiredAppGoodUpdate {
	if b != nil {
		ragu.SetMust(*b)
	}
	return ragu
}

// ClearMust clears the value of the "must" field.
func (ragu *RequiredAppGoodUpdate) ClearMust() *RequiredAppGoodUpdate {
	ragu.mutation.ClearMust()
	return ragu
}

// Mutation returns the RequiredAppGoodMutation object of the builder.
func (ragu *RequiredAppGoodUpdate) Mutation() *RequiredAppGoodMutation {
	return ragu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ragu *RequiredAppGoodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := ragu.defaults(); err != nil {
		return 0, err
	}
	if len(ragu.hooks) == 0 {
		affected, err = ragu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequiredAppGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ragu.mutation = mutation
			affected, err = ragu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ragu.hooks) - 1; i >= 0; i-- {
			if ragu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ragu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ragu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ragu *RequiredAppGoodUpdate) SaveX(ctx context.Context) int {
	affected, err := ragu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ragu *RequiredAppGoodUpdate) Exec(ctx context.Context) error {
	_, err := ragu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ragu *RequiredAppGoodUpdate) ExecX(ctx context.Context) {
	if err := ragu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ragu *RequiredAppGoodUpdate) defaults() error {
	if _, ok := ragu.mutation.UpdatedAt(); !ok {
		if requiredappgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized requiredappgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := requiredappgood.UpdateDefaultUpdatedAt()
		ragu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ragu *RequiredAppGoodUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RequiredAppGoodUpdate {
	ragu.modifiers = append(ragu.modifiers, modifiers...)
	return ragu
}

func (ragu *RequiredAppGoodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requiredappgood.Table,
			Columns: requiredappgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: requiredappgood.FieldID,
			},
		},
	}
	if ps := ragu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ragu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldCreatedAt,
		})
	}
	if value, ok := ragu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldCreatedAt,
		})
	}
	if value, ok := ragu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldUpdatedAt,
		})
	}
	if value, ok := ragu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldUpdatedAt,
		})
	}
	if value, ok := ragu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldDeletedAt,
		})
	}
	if value, ok := ragu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldDeletedAt,
		})
	}
	if value, ok := ragu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: requiredappgood.FieldEntID,
		})
	}
	if value, ok := ragu.mutation.MainAppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: requiredappgood.FieldMainAppGoodID,
		})
	}
	if value, ok := ragu.mutation.RequiredAppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: requiredappgood.FieldRequiredAppGoodID,
		})
	}
	if value, ok := ragu.mutation.Must(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: requiredappgood.FieldMust,
		})
	}
	if ragu.mutation.MustCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: requiredappgood.FieldMust,
		})
	}
	_spec.Modifiers = ragu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, ragu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{requiredappgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RequiredAppGoodUpdateOne is the builder for updating a single RequiredAppGood entity.
type RequiredAppGoodUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *RequiredAppGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (raguo *RequiredAppGoodUpdateOne) SetCreatedAt(u uint32) *RequiredAppGoodUpdateOne {
	raguo.mutation.ResetCreatedAt()
	raguo.mutation.SetCreatedAt(u)
	return raguo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (raguo *RequiredAppGoodUpdateOne) SetNillableCreatedAt(u *uint32) *RequiredAppGoodUpdateOne {
	if u != nil {
		raguo.SetCreatedAt(*u)
	}
	return raguo
}

// AddCreatedAt adds u to the "created_at" field.
func (raguo *RequiredAppGoodUpdateOne) AddCreatedAt(u int32) *RequiredAppGoodUpdateOne {
	raguo.mutation.AddCreatedAt(u)
	return raguo
}

// SetUpdatedAt sets the "updated_at" field.
func (raguo *RequiredAppGoodUpdateOne) SetUpdatedAt(u uint32) *RequiredAppGoodUpdateOne {
	raguo.mutation.ResetUpdatedAt()
	raguo.mutation.SetUpdatedAt(u)
	return raguo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (raguo *RequiredAppGoodUpdateOne) AddUpdatedAt(u int32) *RequiredAppGoodUpdateOne {
	raguo.mutation.AddUpdatedAt(u)
	return raguo
}

// SetDeletedAt sets the "deleted_at" field.
func (raguo *RequiredAppGoodUpdateOne) SetDeletedAt(u uint32) *RequiredAppGoodUpdateOne {
	raguo.mutation.ResetDeletedAt()
	raguo.mutation.SetDeletedAt(u)
	return raguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (raguo *RequiredAppGoodUpdateOne) SetNillableDeletedAt(u *uint32) *RequiredAppGoodUpdateOne {
	if u != nil {
		raguo.SetDeletedAt(*u)
	}
	return raguo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (raguo *RequiredAppGoodUpdateOne) AddDeletedAt(u int32) *RequiredAppGoodUpdateOne {
	raguo.mutation.AddDeletedAt(u)
	return raguo
}

// SetEntID sets the "ent_id" field.
func (raguo *RequiredAppGoodUpdateOne) SetEntID(u uuid.UUID) *RequiredAppGoodUpdateOne {
	raguo.mutation.SetEntID(u)
	return raguo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (raguo *RequiredAppGoodUpdateOne) SetNillableEntID(u *uuid.UUID) *RequiredAppGoodUpdateOne {
	if u != nil {
		raguo.SetEntID(*u)
	}
	return raguo
}

// SetMainAppGoodID sets the "main_app_good_id" field.
func (raguo *RequiredAppGoodUpdateOne) SetMainAppGoodID(u uuid.UUID) *RequiredAppGoodUpdateOne {
	raguo.mutation.SetMainAppGoodID(u)
	return raguo
}

// SetRequiredAppGoodID sets the "required_app_good_id" field.
func (raguo *RequiredAppGoodUpdateOne) SetRequiredAppGoodID(u uuid.UUID) *RequiredAppGoodUpdateOne {
	raguo.mutation.SetRequiredAppGoodID(u)
	return raguo
}

// SetMust sets the "must" field.
func (raguo *RequiredAppGoodUpdateOne) SetMust(b bool) *RequiredAppGoodUpdateOne {
	raguo.mutation.SetMust(b)
	return raguo
}

// SetNillableMust sets the "must" field if the given value is not nil.
func (raguo *RequiredAppGoodUpdateOne) SetNillableMust(b *bool) *RequiredAppGoodUpdateOne {
	if b != nil {
		raguo.SetMust(*b)
	}
	return raguo
}

// ClearMust clears the value of the "must" field.
func (raguo *RequiredAppGoodUpdateOne) ClearMust() *RequiredAppGoodUpdateOne {
	raguo.mutation.ClearMust()
	return raguo
}

// Mutation returns the RequiredAppGoodMutation object of the builder.
func (raguo *RequiredAppGoodUpdateOne) Mutation() *RequiredAppGoodMutation {
	return raguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (raguo *RequiredAppGoodUpdateOne) Select(field string, fields ...string) *RequiredAppGoodUpdateOne {
	raguo.fields = append([]string{field}, fields...)
	return raguo
}

// Save executes the query and returns the updated RequiredAppGood entity.
func (raguo *RequiredAppGoodUpdateOne) Save(ctx context.Context) (*RequiredAppGood, error) {
	var (
		err  error
		node *RequiredAppGood
	)
	if err := raguo.defaults(); err != nil {
		return nil, err
	}
	if len(raguo.hooks) == 0 {
		node, err = raguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RequiredAppGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			raguo.mutation = mutation
			node, err = raguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(raguo.hooks) - 1; i >= 0; i-- {
			if raguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = raguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, raguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RequiredAppGood)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RequiredAppGoodMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (raguo *RequiredAppGoodUpdateOne) SaveX(ctx context.Context) *RequiredAppGood {
	node, err := raguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (raguo *RequiredAppGoodUpdateOne) Exec(ctx context.Context) error {
	_, err := raguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (raguo *RequiredAppGoodUpdateOne) ExecX(ctx context.Context) {
	if err := raguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (raguo *RequiredAppGoodUpdateOne) defaults() error {
	if _, ok := raguo.mutation.UpdatedAt(); !ok {
		if requiredappgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized requiredappgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := requiredappgood.UpdateDefaultUpdatedAt()
		raguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (raguo *RequiredAppGoodUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *RequiredAppGoodUpdateOne {
	raguo.modifiers = append(raguo.modifiers, modifiers...)
	return raguo
}

func (raguo *RequiredAppGoodUpdateOne) sqlSave(ctx context.Context) (_node *RequiredAppGood, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requiredappgood.Table,
			Columns: requiredappgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: requiredappgood.FieldID,
			},
		},
	}
	id, ok := raguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RequiredAppGood.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := raguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, requiredappgood.FieldID)
		for _, f := range fields {
			if !requiredappgood.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != requiredappgood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := raguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := raguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldCreatedAt,
		})
	}
	if value, ok := raguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldCreatedAt,
		})
	}
	if value, ok := raguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldUpdatedAt,
		})
	}
	if value, ok := raguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldUpdatedAt,
		})
	}
	if value, ok := raguo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldDeletedAt,
		})
	}
	if value, ok := raguo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: requiredappgood.FieldDeletedAt,
		})
	}
	if value, ok := raguo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: requiredappgood.FieldEntID,
		})
	}
	if value, ok := raguo.mutation.MainAppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: requiredappgood.FieldMainAppGoodID,
		})
	}
	if value, ok := raguo.mutation.RequiredAppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: requiredappgood.FieldRequiredAppGoodID,
		})
	}
	if value, ok := raguo.mutation.Must(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: requiredappgood.FieldMust,
		})
	}
	if raguo.mutation.MustCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: requiredappgood.FieldMust,
		})
	}
	_spec.Modifiers = raguo.modifiers
	_node = &RequiredAppGood{config: raguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, raguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{requiredappgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
