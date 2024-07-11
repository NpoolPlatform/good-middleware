// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddescription"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppGoodDescriptionUpdate is the builder for updating AppGoodDescription entities.
type AppGoodDescriptionUpdate struct {
	config
	hooks     []Hook
	mutation  *AppGoodDescriptionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppGoodDescriptionUpdate builder.
func (agdu *AppGoodDescriptionUpdate) Where(ps ...predicate.AppGoodDescription) *AppGoodDescriptionUpdate {
	agdu.mutation.Where(ps...)
	return agdu
}

// SetCreatedAt sets the "created_at" field.
func (agdu *AppGoodDescriptionUpdate) SetCreatedAt(u uint32) *AppGoodDescriptionUpdate {
	agdu.mutation.ResetCreatedAt()
	agdu.mutation.SetCreatedAt(u)
	return agdu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (agdu *AppGoodDescriptionUpdate) SetNillableCreatedAt(u *uint32) *AppGoodDescriptionUpdate {
	if u != nil {
		agdu.SetCreatedAt(*u)
	}
	return agdu
}

// AddCreatedAt adds u to the "created_at" field.
func (agdu *AppGoodDescriptionUpdate) AddCreatedAt(u int32) *AppGoodDescriptionUpdate {
	agdu.mutation.AddCreatedAt(u)
	return agdu
}

// SetUpdatedAt sets the "updated_at" field.
func (agdu *AppGoodDescriptionUpdate) SetUpdatedAt(u uint32) *AppGoodDescriptionUpdate {
	agdu.mutation.ResetUpdatedAt()
	agdu.mutation.SetUpdatedAt(u)
	return agdu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (agdu *AppGoodDescriptionUpdate) AddUpdatedAt(u int32) *AppGoodDescriptionUpdate {
	agdu.mutation.AddUpdatedAt(u)
	return agdu
}

// SetDeletedAt sets the "deleted_at" field.
func (agdu *AppGoodDescriptionUpdate) SetDeletedAt(u uint32) *AppGoodDescriptionUpdate {
	agdu.mutation.ResetDeletedAt()
	agdu.mutation.SetDeletedAt(u)
	return agdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (agdu *AppGoodDescriptionUpdate) SetNillableDeletedAt(u *uint32) *AppGoodDescriptionUpdate {
	if u != nil {
		agdu.SetDeletedAt(*u)
	}
	return agdu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (agdu *AppGoodDescriptionUpdate) AddDeletedAt(u int32) *AppGoodDescriptionUpdate {
	agdu.mutation.AddDeletedAt(u)
	return agdu
}

// SetEntID sets the "ent_id" field.
func (agdu *AppGoodDescriptionUpdate) SetEntID(u uuid.UUID) *AppGoodDescriptionUpdate {
	agdu.mutation.SetEntID(u)
	return agdu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (agdu *AppGoodDescriptionUpdate) SetNillableEntID(u *uuid.UUID) *AppGoodDescriptionUpdate {
	if u != nil {
		agdu.SetEntID(*u)
	}
	return agdu
}

// SetAppGoodID sets the "app_good_id" field.
func (agdu *AppGoodDescriptionUpdate) SetAppGoodID(u uuid.UUID) *AppGoodDescriptionUpdate {
	agdu.mutation.SetAppGoodID(u)
	return agdu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (agdu *AppGoodDescriptionUpdate) SetNillableAppGoodID(u *uuid.UUID) *AppGoodDescriptionUpdate {
	if u != nil {
		agdu.SetAppGoodID(*u)
	}
	return agdu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (agdu *AppGoodDescriptionUpdate) ClearAppGoodID() *AppGoodDescriptionUpdate {
	agdu.mutation.ClearAppGoodID()
	return agdu
}

// SetDescription sets the "description" field.
func (agdu *AppGoodDescriptionUpdate) SetDescription(s string) *AppGoodDescriptionUpdate {
	agdu.mutation.SetDescription(s)
	return agdu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (agdu *AppGoodDescriptionUpdate) SetNillableDescription(s *string) *AppGoodDescriptionUpdate {
	if s != nil {
		agdu.SetDescription(*s)
	}
	return agdu
}

// ClearDescription clears the value of the "description" field.
func (agdu *AppGoodDescriptionUpdate) ClearDescription() *AppGoodDescriptionUpdate {
	agdu.mutation.ClearDescription()
	return agdu
}

// SetIndex sets the "index" field.
func (agdu *AppGoodDescriptionUpdate) SetIndex(u uint8) *AppGoodDescriptionUpdate {
	agdu.mutation.ResetIndex()
	agdu.mutation.SetIndex(u)
	return agdu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (agdu *AppGoodDescriptionUpdate) SetNillableIndex(u *uint8) *AppGoodDescriptionUpdate {
	if u != nil {
		agdu.SetIndex(*u)
	}
	return agdu
}

// AddIndex adds u to the "index" field.
func (agdu *AppGoodDescriptionUpdate) AddIndex(u int8) *AppGoodDescriptionUpdate {
	agdu.mutation.AddIndex(u)
	return agdu
}

// ClearIndex clears the value of the "index" field.
func (agdu *AppGoodDescriptionUpdate) ClearIndex() *AppGoodDescriptionUpdate {
	agdu.mutation.ClearIndex()
	return agdu
}

// Mutation returns the AppGoodDescriptionMutation object of the builder.
func (agdu *AppGoodDescriptionUpdate) Mutation() *AppGoodDescriptionMutation {
	return agdu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (agdu *AppGoodDescriptionUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := agdu.defaults(); err != nil {
		return 0, err
	}
	if len(agdu.hooks) == 0 {
		affected, err = agdu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agdu.mutation = mutation
			affected, err = agdu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(agdu.hooks) - 1; i >= 0; i-- {
			if agdu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agdu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, agdu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (agdu *AppGoodDescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := agdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (agdu *AppGoodDescriptionUpdate) Exec(ctx context.Context) error {
	_, err := agdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agdu *AppGoodDescriptionUpdate) ExecX(ctx context.Context) {
	if err := agdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agdu *AppGoodDescriptionUpdate) defaults() error {
	if _, ok := agdu.mutation.UpdatedAt(); !ok {
		if appgooddescription.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appgooddescription.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appgooddescription.UpdateDefaultUpdatedAt()
		agdu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (agdu *AppGoodDescriptionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppGoodDescriptionUpdate {
	agdu.modifiers = append(agdu.modifiers, modifiers...)
	return agdu
}

func (agdu *AppGoodDescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgooddescription.Table,
			Columns: appgooddescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgooddescription.FieldID,
			},
		},
	}
	if ps := agdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := agdu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldCreatedAt,
		})
	}
	if value, ok := agdu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldCreatedAt,
		})
	}
	if value, ok := agdu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldUpdatedAt,
		})
	}
	if value, ok := agdu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldUpdatedAt,
		})
	}
	if value, ok := agdu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldDeletedAt,
		})
	}
	if value, ok := agdu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldDeletedAt,
		})
	}
	if value, ok := agdu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgooddescription.FieldEntID,
		})
	}
	if value, ok := agdu.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgooddescription.FieldAppGoodID,
		})
	}
	if agdu.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appgooddescription.FieldAppGoodID,
		})
	}
	if value, ok := agdu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appgooddescription.FieldDescription,
		})
	}
	if agdu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appgooddescription.FieldDescription,
		})
	}
	if value, ok := agdu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appgooddescription.FieldIndex,
		})
	}
	if value, ok := agdu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appgooddescription.FieldIndex,
		})
	}
	if agdu.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: appgooddescription.FieldIndex,
		})
	}
	_spec.Modifiers = agdu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, agdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appgooddescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppGoodDescriptionUpdateOne is the builder for updating a single AppGoodDescription entity.
type AppGoodDescriptionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppGoodDescriptionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (agduo *AppGoodDescriptionUpdateOne) SetCreatedAt(u uint32) *AppGoodDescriptionUpdateOne {
	agduo.mutation.ResetCreatedAt()
	agduo.mutation.SetCreatedAt(u)
	return agduo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (agduo *AppGoodDescriptionUpdateOne) SetNillableCreatedAt(u *uint32) *AppGoodDescriptionUpdateOne {
	if u != nil {
		agduo.SetCreatedAt(*u)
	}
	return agduo
}

// AddCreatedAt adds u to the "created_at" field.
func (agduo *AppGoodDescriptionUpdateOne) AddCreatedAt(u int32) *AppGoodDescriptionUpdateOne {
	agduo.mutation.AddCreatedAt(u)
	return agduo
}

// SetUpdatedAt sets the "updated_at" field.
func (agduo *AppGoodDescriptionUpdateOne) SetUpdatedAt(u uint32) *AppGoodDescriptionUpdateOne {
	agduo.mutation.ResetUpdatedAt()
	agduo.mutation.SetUpdatedAt(u)
	return agduo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (agduo *AppGoodDescriptionUpdateOne) AddUpdatedAt(u int32) *AppGoodDescriptionUpdateOne {
	agduo.mutation.AddUpdatedAt(u)
	return agduo
}

// SetDeletedAt sets the "deleted_at" field.
func (agduo *AppGoodDescriptionUpdateOne) SetDeletedAt(u uint32) *AppGoodDescriptionUpdateOne {
	agduo.mutation.ResetDeletedAt()
	agduo.mutation.SetDeletedAt(u)
	return agduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (agduo *AppGoodDescriptionUpdateOne) SetNillableDeletedAt(u *uint32) *AppGoodDescriptionUpdateOne {
	if u != nil {
		agduo.SetDeletedAt(*u)
	}
	return agduo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (agduo *AppGoodDescriptionUpdateOne) AddDeletedAt(u int32) *AppGoodDescriptionUpdateOne {
	agduo.mutation.AddDeletedAt(u)
	return agduo
}

// SetEntID sets the "ent_id" field.
func (agduo *AppGoodDescriptionUpdateOne) SetEntID(u uuid.UUID) *AppGoodDescriptionUpdateOne {
	agduo.mutation.SetEntID(u)
	return agduo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (agduo *AppGoodDescriptionUpdateOne) SetNillableEntID(u *uuid.UUID) *AppGoodDescriptionUpdateOne {
	if u != nil {
		agduo.SetEntID(*u)
	}
	return agduo
}

// SetAppGoodID sets the "app_good_id" field.
func (agduo *AppGoodDescriptionUpdateOne) SetAppGoodID(u uuid.UUID) *AppGoodDescriptionUpdateOne {
	agduo.mutation.SetAppGoodID(u)
	return agduo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (agduo *AppGoodDescriptionUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *AppGoodDescriptionUpdateOne {
	if u != nil {
		agduo.SetAppGoodID(*u)
	}
	return agduo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (agduo *AppGoodDescriptionUpdateOne) ClearAppGoodID() *AppGoodDescriptionUpdateOne {
	agduo.mutation.ClearAppGoodID()
	return agduo
}

// SetDescription sets the "description" field.
func (agduo *AppGoodDescriptionUpdateOne) SetDescription(s string) *AppGoodDescriptionUpdateOne {
	agduo.mutation.SetDescription(s)
	return agduo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (agduo *AppGoodDescriptionUpdateOne) SetNillableDescription(s *string) *AppGoodDescriptionUpdateOne {
	if s != nil {
		agduo.SetDescription(*s)
	}
	return agduo
}

// ClearDescription clears the value of the "description" field.
func (agduo *AppGoodDescriptionUpdateOne) ClearDescription() *AppGoodDescriptionUpdateOne {
	agduo.mutation.ClearDescription()
	return agduo
}

// SetIndex sets the "index" field.
func (agduo *AppGoodDescriptionUpdateOne) SetIndex(u uint8) *AppGoodDescriptionUpdateOne {
	agduo.mutation.ResetIndex()
	agduo.mutation.SetIndex(u)
	return agduo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (agduo *AppGoodDescriptionUpdateOne) SetNillableIndex(u *uint8) *AppGoodDescriptionUpdateOne {
	if u != nil {
		agduo.SetIndex(*u)
	}
	return agduo
}

// AddIndex adds u to the "index" field.
func (agduo *AppGoodDescriptionUpdateOne) AddIndex(u int8) *AppGoodDescriptionUpdateOne {
	agduo.mutation.AddIndex(u)
	return agduo
}

// ClearIndex clears the value of the "index" field.
func (agduo *AppGoodDescriptionUpdateOne) ClearIndex() *AppGoodDescriptionUpdateOne {
	agduo.mutation.ClearIndex()
	return agduo
}

// Mutation returns the AppGoodDescriptionMutation object of the builder.
func (agduo *AppGoodDescriptionUpdateOne) Mutation() *AppGoodDescriptionMutation {
	return agduo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (agduo *AppGoodDescriptionUpdateOne) Select(field string, fields ...string) *AppGoodDescriptionUpdateOne {
	agduo.fields = append([]string{field}, fields...)
	return agduo
}

// Save executes the query and returns the updated AppGoodDescription entity.
func (agduo *AppGoodDescriptionUpdateOne) Save(ctx context.Context) (*AppGoodDescription, error) {
	var (
		err  error
		node *AppGoodDescription
	)
	if err := agduo.defaults(); err != nil {
		return nil, err
	}
	if len(agduo.hooks) == 0 {
		node, err = agduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodDescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			agduo.mutation = mutation
			node, err = agduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(agduo.hooks) - 1; i >= 0; i-- {
			if agduo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agduo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, agduo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppGoodDescription)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppGoodDescriptionMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (agduo *AppGoodDescriptionUpdateOne) SaveX(ctx context.Context) *AppGoodDescription {
	node, err := agduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (agduo *AppGoodDescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := agduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agduo *AppGoodDescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := agduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agduo *AppGoodDescriptionUpdateOne) defaults() error {
	if _, ok := agduo.mutation.UpdatedAt(); !ok {
		if appgooddescription.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appgooddescription.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appgooddescription.UpdateDefaultUpdatedAt()
		agduo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (agduo *AppGoodDescriptionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppGoodDescriptionUpdateOne {
	agduo.modifiers = append(agduo.modifiers, modifiers...)
	return agduo
}

func (agduo *AppGoodDescriptionUpdateOne) sqlSave(ctx context.Context) (_node *AppGoodDescription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgooddescription.Table,
			Columns: appgooddescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgooddescription.FieldID,
			},
		},
	}
	id, ok := agduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppGoodDescription.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := agduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appgooddescription.FieldID)
		for _, f := range fields {
			if !appgooddescription.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appgooddescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := agduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := agduo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldCreatedAt,
		})
	}
	if value, ok := agduo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldCreatedAt,
		})
	}
	if value, ok := agduo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldUpdatedAt,
		})
	}
	if value, ok := agduo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldUpdatedAt,
		})
	}
	if value, ok := agduo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldDeletedAt,
		})
	}
	if value, ok := agduo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddescription.FieldDeletedAt,
		})
	}
	if value, ok := agduo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgooddescription.FieldEntID,
		})
	}
	if value, ok := agduo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgooddescription.FieldAppGoodID,
		})
	}
	if agduo.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appgooddescription.FieldAppGoodID,
		})
	}
	if value, ok := agduo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appgooddescription.FieldDescription,
		})
	}
	if agduo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: appgooddescription.FieldDescription,
		})
	}
	if value, ok := agduo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appgooddescription.FieldIndex,
		})
	}
	if value, ok := agduo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appgooddescription.FieldIndex,
		})
	}
	if agduo.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Column: appgooddescription.FieldIndex,
		})
	}
	_spec.Modifiers = agduo.modifiers
	_node = &AppGoodDescription{config: agduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, agduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appgooddescription.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
