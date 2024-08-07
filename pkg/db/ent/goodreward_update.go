// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodRewardUpdate is the builder for updating GoodReward entities.
type GoodRewardUpdate struct {
	config
	hooks     []Hook
	mutation  *GoodRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoodRewardUpdate builder.
func (gru *GoodRewardUpdate) Where(ps ...predicate.GoodReward) *GoodRewardUpdate {
	gru.mutation.Where(ps...)
	return gru
}

// SetCreatedAt sets the "created_at" field.
func (gru *GoodRewardUpdate) SetCreatedAt(u uint32) *GoodRewardUpdate {
	gru.mutation.ResetCreatedAt()
	gru.mutation.SetCreatedAt(u)
	return gru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gru *GoodRewardUpdate) SetNillableCreatedAt(u *uint32) *GoodRewardUpdate {
	if u != nil {
		gru.SetCreatedAt(*u)
	}
	return gru
}

// AddCreatedAt adds u to the "created_at" field.
func (gru *GoodRewardUpdate) AddCreatedAt(u int32) *GoodRewardUpdate {
	gru.mutation.AddCreatedAt(u)
	return gru
}

// SetUpdatedAt sets the "updated_at" field.
func (gru *GoodRewardUpdate) SetUpdatedAt(u uint32) *GoodRewardUpdate {
	gru.mutation.ResetUpdatedAt()
	gru.mutation.SetUpdatedAt(u)
	return gru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gru *GoodRewardUpdate) AddUpdatedAt(u int32) *GoodRewardUpdate {
	gru.mutation.AddUpdatedAt(u)
	return gru
}

// SetDeletedAt sets the "deleted_at" field.
func (gru *GoodRewardUpdate) SetDeletedAt(u uint32) *GoodRewardUpdate {
	gru.mutation.ResetDeletedAt()
	gru.mutation.SetDeletedAt(u)
	return gru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gru *GoodRewardUpdate) SetNillableDeletedAt(u *uint32) *GoodRewardUpdate {
	if u != nil {
		gru.SetDeletedAt(*u)
	}
	return gru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gru *GoodRewardUpdate) AddDeletedAt(u int32) *GoodRewardUpdate {
	gru.mutation.AddDeletedAt(u)
	return gru
}

// SetEntID sets the "ent_id" field.
func (gru *GoodRewardUpdate) SetEntID(u uuid.UUID) *GoodRewardUpdate {
	gru.mutation.SetEntID(u)
	return gru
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gru *GoodRewardUpdate) SetNillableEntID(u *uuid.UUID) *GoodRewardUpdate {
	if u != nil {
		gru.SetEntID(*u)
	}
	return gru
}

// SetGoodID sets the "good_id" field.
func (gru *GoodRewardUpdate) SetGoodID(u uuid.UUID) *GoodRewardUpdate {
	gru.mutation.SetGoodID(u)
	return gru
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gru *GoodRewardUpdate) SetNillableGoodID(u *uuid.UUID) *GoodRewardUpdate {
	if u != nil {
		gru.SetGoodID(*u)
	}
	return gru
}

// ClearGoodID clears the value of the "good_id" field.
func (gru *GoodRewardUpdate) ClearGoodID() *GoodRewardUpdate {
	gru.mutation.ClearGoodID()
	return gru
}

// SetRewardState sets the "reward_state" field.
func (gru *GoodRewardUpdate) SetRewardState(s string) *GoodRewardUpdate {
	gru.mutation.SetRewardState(s)
	return gru
}

// SetNillableRewardState sets the "reward_state" field if the given value is not nil.
func (gru *GoodRewardUpdate) SetNillableRewardState(s *string) *GoodRewardUpdate {
	if s != nil {
		gru.SetRewardState(*s)
	}
	return gru
}

// ClearRewardState clears the value of the "reward_state" field.
func (gru *GoodRewardUpdate) ClearRewardState() *GoodRewardUpdate {
	gru.mutation.ClearRewardState()
	return gru
}

// SetLastRewardAt sets the "last_reward_at" field.
func (gru *GoodRewardUpdate) SetLastRewardAt(u uint32) *GoodRewardUpdate {
	gru.mutation.ResetLastRewardAt()
	gru.mutation.SetLastRewardAt(u)
	return gru
}

// SetNillableLastRewardAt sets the "last_reward_at" field if the given value is not nil.
func (gru *GoodRewardUpdate) SetNillableLastRewardAt(u *uint32) *GoodRewardUpdate {
	if u != nil {
		gru.SetLastRewardAt(*u)
	}
	return gru
}

// AddLastRewardAt adds u to the "last_reward_at" field.
func (gru *GoodRewardUpdate) AddLastRewardAt(u int32) *GoodRewardUpdate {
	gru.mutation.AddLastRewardAt(u)
	return gru
}

// ClearLastRewardAt clears the value of the "last_reward_at" field.
func (gru *GoodRewardUpdate) ClearLastRewardAt() *GoodRewardUpdate {
	gru.mutation.ClearLastRewardAt()
	return gru
}

// Mutation returns the GoodRewardMutation object of the builder.
func (gru *GoodRewardUpdate) Mutation() *GoodRewardMutation {
	return gru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gru *GoodRewardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := gru.defaults(); err != nil {
		return 0, err
	}
	if len(gru.hooks) == 0 {
		affected, err = gru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gru.mutation = mutation
			affected, err = gru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gru.hooks) - 1; i >= 0; i-- {
			if gru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gru *GoodRewardUpdate) SaveX(ctx context.Context) int {
	affected, err := gru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gru *GoodRewardUpdate) Exec(ctx context.Context) error {
	_, err := gru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gru *GoodRewardUpdate) ExecX(ctx context.Context) {
	if err := gru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gru *GoodRewardUpdate) defaults() error {
	if _, ok := gru.mutation.UpdatedAt(); !ok {
		if goodreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodreward.UpdateDefaultUpdatedAt()
		gru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gru *GoodRewardUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodRewardUpdate {
	gru.modifiers = append(gru.modifiers, modifiers...)
	return gru
}

func (gru *GoodRewardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodreward.Table,
			Columns: goodreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodreward.FieldID,
			},
		},
	}
	if ps := gru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldCreatedAt,
		})
	}
	if value, ok := gru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldCreatedAt,
		})
	}
	if value, ok := gru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldUpdatedAt,
		})
	}
	if value, ok := gru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldUpdatedAt,
		})
	}
	if value, ok := gru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldDeletedAt,
		})
	}
	if value, ok := gru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldDeletedAt,
		})
	}
	if value, ok := gru.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodreward.FieldEntID,
		})
	}
	if value, ok := gru.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodreward.FieldGoodID,
		})
	}
	if gru.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodreward.FieldGoodID,
		})
	}
	if value, ok := gru.mutation.RewardState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodreward.FieldRewardState,
		})
	}
	if gru.mutation.RewardStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodreward.FieldRewardState,
		})
	}
	if value, ok := gru.mutation.LastRewardAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldLastRewardAt,
		})
	}
	if value, ok := gru.mutation.AddedLastRewardAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldLastRewardAt,
		})
	}
	if gru.mutation.LastRewardAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: goodreward.FieldLastRewardAt,
		})
	}
	_spec.Modifiers = gru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, gru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GoodRewardUpdateOne is the builder for updating a single GoodReward entity.
type GoodRewardUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoodRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (gruo *GoodRewardUpdateOne) SetCreatedAt(u uint32) *GoodRewardUpdateOne {
	gruo.mutation.ResetCreatedAt()
	gruo.mutation.SetCreatedAt(u)
	return gruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gruo *GoodRewardUpdateOne) SetNillableCreatedAt(u *uint32) *GoodRewardUpdateOne {
	if u != nil {
		gruo.SetCreatedAt(*u)
	}
	return gruo
}

// AddCreatedAt adds u to the "created_at" field.
func (gruo *GoodRewardUpdateOne) AddCreatedAt(u int32) *GoodRewardUpdateOne {
	gruo.mutation.AddCreatedAt(u)
	return gruo
}

// SetUpdatedAt sets the "updated_at" field.
func (gruo *GoodRewardUpdateOne) SetUpdatedAt(u uint32) *GoodRewardUpdateOne {
	gruo.mutation.ResetUpdatedAt()
	gruo.mutation.SetUpdatedAt(u)
	return gruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gruo *GoodRewardUpdateOne) AddUpdatedAt(u int32) *GoodRewardUpdateOne {
	gruo.mutation.AddUpdatedAt(u)
	return gruo
}

// SetDeletedAt sets the "deleted_at" field.
func (gruo *GoodRewardUpdateOne) SetDeletedAt(u uint32) *GoodRewardUpdateOne {
	gruo.mutation.ResetDeletedAt()
	gruo.mutation.SetDeletedAt(u)
	return gruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gruo *GoodRewardUpdateOne) SetNillableDeletedAt(u *uint32) *GoodRewardUpdateOne {
	if u != nil {
		gruo.SetDeletedAt(*u)
	}
	return gruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gruo *GoodRewardUpdateOne) AddDeletedAt(u int32) *GoodRewardUpdateOne {
	gruo.mutation.AddDeletedAt(u)
	return gruo
}

// SetEntID sets the "ent_id" field.
func (gruo *GoodRewardUpdateOne) SetEntID(u uuid.UUID) *GoodRewardUpdateOne {
	gruo.mutation.SetEntID(u)
	return gruo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gruo *GoodRewardUpdateOne) SetNillableEntID(u *uuid.UUID) *GoodRewardUpdateOne {
	if u != nil {
		gruo.SetEntID(*u)
	}
	return gruo
}

// SetGoodID sets the "good_id" field.
func (gruo *GoodRewardUpdateOne) SetGoodID(u uuid.UUID) *GoodRewardUpdateOne {
	gruo.mutation.SetGoodID(u)
	return gruo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gruo *GoodRewardUpdateOne) SetNillableGoodID(u *uuid.UUID) *GoodRewardUpdateOne {
	if u != nil {
		gruo.SetGoodID(*u)
	}
	return gruo
}

// ClearGoodID clears the value of the "good_id" field.
func (gruo *GoodRewardUpdateOne) ClearGoodID() *GoodRewardUpdateOne {
	gruo.mutation.ClearGoodID()
	return gruo
}

// SetRewardState sets the "reward_state" field.
func (gruo *GoodRewardUpdateOne) SetRewardState(s string) *GoodRewardUpdateOne {
	gruo.mutation.SetRewardState(s)
	return gruo
}

// SetNillableRewardState sets the "reward_state" field if the given value is not nil.
func (gruo *GoodRewardUpdateOne) SetNillableRewardState(s *string) *GoodRewardUpdateOne {
	if s != nil {
		gruo.SetRewardState(*s)
	}
	return gruo
}

// ClearRewardState clears the value of the "reward_state" field.
func (gruo *GoodRewardUpdateOne) ClearRewardState() *GoodRewardUpdateOne {
	gruo.mutation.ClearRewardState()
	return gruo
}

// SetLastRewardAt sets the "last_reward_at" field.
func (gruo *GoodRewardUpdateOne) SetLastRewardAt(u uint32) *GoodRewardUpdateOne {
	gruo.mutation.ResetLastRewardAt()
	gruo.mutation.SetLastRewardAt(u)
	return gruo
}

// SetNillableLastRewardAt sets the "last_reward_at" field if the given value is not nil.
func (gruo *GoodRewardUpdateOne) SetNillableLastRewardAt(u *uint32) *GoodRewardUpdateOne {
	if u != nil {
		gruo.SetLastRewardAt(*u)
	}
	return gruo
}

// AddLastRewardAt adds u to the "last_reward_at" field.
func (gruo *GoodRewardUpdateOne) AddLastRewardAt(u int32) *GoodRewardUpdateOne {
	gruo.mutation.AddLastRewardAt(u)
	return gruo
}

// ClearLastRewardAt clears the value of the "last_reward_at" field.
func (gruo *GoodRewardUpdateOne) ClearLastRewardAt() *GoodRewardUpdateOne {
	gruo.mutation.ClearLastRewardAt()
	return gruo
}

// Mutation returns the GoodRewardMutation object of the builder.
func (gruo *GoodRewardUpdateOne) Mutation() *GoodRewardMutation {
	return gruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gruo *GoodRewardUpdateOne) Select(field string, fields ...string) *GoodRewardUpdateOne {
	gruo.fields = append([]string{field}, fields...)
	return gruo
}

// Save executes the query and returns the updated GoodReward entity.
func (gruo *GoodRewardUpdateOne) Save(ctx context.Context) (*GoodReward, error) {
	var (
		err  error
		node *GoodReward
	)
	if err := gruo.defaults(); err != nil {
		return nil, err
	}
	if len(gruo.hooks) == 0 {
		node, err = gruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gruo.mutation = mutation
			node, err = gruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gruo.hooks) - 1; i >= 0; i-- {
			if gruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoodReward)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoodRewardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gruo *GoodRewardUpdateOne) SaveX(ctx context.Context) *GoodReward {
	node, err := gruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gruo *GoodRewardUpdateOne) Exec(ctx context.Context) error {
	_, err := gruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gruo *GoodRewardUpdateOne) ExecX(ctx context.Context) {
	if err := gruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gruo *GoodRewardUpdateOne) defaults() error {
	if _, ok := gruo.mutation.UpdatedAt(); !ok {
		if goodreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodreward.UpdateDefaultUpdatedAt()
		gruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gruo *GoodRewardUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodRewardUpdateOne {
	gruo.modifiers = append(gruo.modifiers, modifiers...)
	return gruo
}

func (gruo *GoodRewardUpdateOne) sqlSave(ctx context.Context) (_node *GoodReward, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodreward.Table,
			Columns: goodreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodreward.FieldID,
			},
		},
	}
	id, ok := gruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoodReward.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodreward.FieldID)
		for _, f := range fields {
			if !goodreward.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodreward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldCreatedAt,
		})
	}
	if value, ok := gruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldCreatedAt,
		})
	}
	if value, ok := gruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldUpdatedAt,
		})
	}
	if value, ok := gruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldUpdatedAt,
		})
	}
	if value, ok := gruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldDeletedAt,
		})
	}
	if value, ok := gruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldDeletedAt,
		})
	}
	if value, ok := gruo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodreward.FieldEntID,
		})
	}
	if value, ok := gruo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodreward.FieldGoodID,
		})
	}
	if gruo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodreward.FieldGoodID,
		})
	}
	if value, ok := gruo.mutation.RewardState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodreward.FieldRewardState,
		})
	}
	if gruo.mutation.RewardStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodreward.FieldRewardState,
		})
	}
	if value, ok := gruo.mutation.LastRewardAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldLastRewardAt,
		})
	}
	if value, ok := gruo.mutation.AddedLastRewardAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodreward.FieldLastRewardAt,
		})
	}
	if gruo.mutation.LastRewardAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: goodreward.FieldLastRewardAt,
		})
	}
	_spec.Modifiers = gruo.modifiers
	_node = &GoodReward{config: gruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
