// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoinreward"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodCoinRewardUpdate is the builder for updating GoodCoinReward entities.
type GoodCoinRewardUpdate struct {
	config
	hooks     []Hook
	mutation  *GoodCoinRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoodCoinRewardUpdate builder.
func (gcru *GoodCoinRewardUpdate) Where(ps ...predicate.GoodCoinReward) *GoodCoinRewardUpdate {
	gcru.mutation.Where(ps...)
	return gcru
}

// SetCreatedAt sets the "created_at" field.
func (gcru *GoodCoinRewardUpdate) SetCreatedAt(u uint32) *GoodCoinRewardUpdate {
	gcru.mutation.ResetCreatedAt()
	gcru.mutation.SetCreatedAt(u)
	return gcru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableCreatedAt(u *uint32) *GoodCoinRewardUpdate {
	if u != nil {
		gcru.SetCreatedAt(*u)
	}
	return gcru
}

// AddCreatedAt adds u to the "created_at" field.
func (gcru *GoodCoinRewardUpdate) AddCreatedAt(u int32) *GoodCoinRewardUpdate {
	gcru.mutation.AddCreatedAt(u)
	return gcru
}

// SetUpdatedAt sets the "updated_at" field.
func (gcru *GoodCoinRewardUpdate) SetUpdatedAt(u uint32) *GoodCoinRewardUpdate {
	gcru.mutation.ResetUpdatedAt()
	gcru.mutation.SetUpdatedAt(u)
	return gcru
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gcru *GoodCoinRewardUpdate) AddUpdatedAt(u int32) *GoodCoinRewardUpdate {
	gcru.mutation.AddUpdatedAt(u)
	return gcru
}

// SetDeletedAt sets the "deleted_at" field.
func (gcru *GoodCoinRewardUpdate) SetDeletedAt(u uint32) *GoodCoinRewardUpdate {
	gcru.mutation.ResetDeletedAt()
	gcru.mutation.SetDeletedAt(u)
	return gcru
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableDeletedAt(u *uint32) *GoodCoinRewardUpdate {
	if u != nil {
		gcru.SetDeletedAt(*u)
	}
	return gcru
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gcru *GoodCoinRewardUpdate) AddDeletedAt(u int32) *GoodCoinRewardUpdate {
	gcru.mutation.AddDeletedAt(u)
	return gcru
}

// SetEntID sets the "ent_id" field.
func (gcru *GoodCoinRewardUpdate) SetEntID(u uuid.UUID) *GoodCoinRewardUpdate {
	gcru.mutation.SetEntID(u)
	return gcru
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableEntID(u *uuid.UUID) *GoodCoinRewardUpdate {
	if u != nil {
		gcru.SetEntID(*u)
	}
	return gcru
}

// SetGoodID sets the "good_id" field.
func (gcru *GoodCoinRewardUpdate) SetGoodID(u uuid.UUID) *GoodCoinRewardUpdate {
	gcru.mutation.SetGoodID(u)
	return gcru
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableGoodID(u *uuid.UUID) *GoodCoinRewardUpdate {
	if u != nil {
		gcru.SetGoodID(*u)
	}
	return gcru
}

// ClearGoodID clears the value of the "good_id" field.
func (gcru *GoodCoinRewardUpdate) ClearGoodID() *GoodCoinRewardUpdate {
	gcru.mutation.ClearGoodID()
	return gcru
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gcru *GoodCoinRewardUpdate) SetCoinTypeID(u uuid.UUID) *GoodCoinRewardUpdate {
	gcru.mutation.SetCoinTypeID(u)
	return gcru
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableCoinTypeID(u *uuid.UUID) *GoodCoinRewardUpdate {
	if u != nil {
		gcru.SetCoinTypeID(*u)
	}
	return gcru
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gcru *GoodCoinRewardUpdate) ClearCoinTypeID() *GoodCoinRewardUpdate {
	gcru.mutation.ClearCoinTypeID()
	return gcru
}

// SetRewardTid sets the "reward_tid" field.
func (gcru *GoodCoinRewardUpdate) SetRewardTid(u uuid.UUID) *GoodCoinRewardUpdate {
	gcru.mutation.SetRewardTid(u)
	return gcru
}

// SetNillableRewardTid sets the "reward_tid" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableRewardTid(u *uuid.UUID) *GoodCoinRewardUpdate {
	if u != nil {
		gcru.SetRewardTid(*u)
	}
	return gcru
}

// ClearRewardTid clears the value of the "reward_tid" field.
func (gcru *GoodCoinRewardUpdate) ClearRewardTid() *GoodCoinRewardUpdate {
	gcru.mutation.ClearRewardTid()
	return gcru
}

// SetNextRewardStartAmount sets the "next_reward_start_amount" field.
func (gcru *GoodCoinRewardUpdate) SetNextRewardStartAmount(d decimal.Decimal) *GoodCoinRewardUpdate {
	gcru.mutation.SetNextRewardStartAmount(d)
	return gcru
}

// SetNillableNextRewardStartAmount sets the "next_reward_start_amount" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableNextRewardStartAmount(d *decimal.Decimal) *GoodCoinRewardUpdate {
	if d != nil {
		gcru.SetNextRewardStartAmount(*d)
	}
	return gcru
}

// ClearNextRewardStartAmount clears the value of the "next_reward_start_amount" field.
func (gcru *GoodCoinRewardUpdate) ClearNextRewardStartAmount() *GoodCoinRewardUpdate {
	gcru.mutation.ClearNextRewardStartAmount()
	return gcru
}

// SetLastRewardAmount sets the "last_reward_amount" field.
func (gcru *GoodCoinRewardUpdate) SetLastRewardAmount(d decimal.Decimal) *GoodCoinRewardUpdate {
	gcru.mutation.SetLastRewardAmount(d)
	return gcru
}

// SetNillableLastRewardAmount sets the "last_reward_amount" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableLastRewardAmount(d *decimal.Decimal) *GoodCoinRewardUpdate {
	if d != nil {
		gcru.SetLastRewardAmount(*d)
	}
	return gcru
}

// ClearLastRewardAmount clears the value of the "last_reward_amount" field.
func (gcru *GoodCoinRewardUpdate) ClearLastRewardAmount() *GoodCoinRewardUpdate {
	gcru.mutation.ClearLastRewardAmount()
	return gcru
}

// SetLastUnitRewardAmount sets the "last_unit_reward_amount" field.
func (gcru *GoodCoinRewardUpdate) SetLastUnitRewardAmount(d decimal.Decimal) *GoodCoinRewardUpdate {
	gcru.mutation.SetLastUnitRewardAmount(d)
	return gcru
}

// SetNillableLastUnitRewardAmount sets the "last_unit_reward_amount" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableLastUnitRewardAmount(d *decimal.Decimal) *GoodCoinRewardUpdate {
	if d != nil {
		gcru.SetLastUnitRewardAmount(*d)
	}
	return gcru
}

// ClearLastUnitRewardAmount clears the value of the "last_unit_reward_amount" field.
func (gcru *GoodCoinRewardUpdate) ClearLastUnitRewardAmount() *GoodCoinRewardUpdate {
	gcru.mutation.ClearLastUnitRewardAmount()
	return gcru
}

// SetTotalRewardAmount sets the "total_reward_amount" field.
func (gcru *GoodCoinRewardUpdate) SetTotalRewardAmount(d decimal.Decimal) *GoodCoinRewardUpdate {
	gcru.mutation.SetTotalRewardAmount(d)
	return gcru
}

// SetNillableTotalRewardAmount sets the "total_reward_amount" field if the given value is not nil.
func (gcru *GoodCoinRewardUpdate) SetNillableTotalRewardAmount(d *decimal.Decimal) *GoodCoinRewardUpdate {
	if d != nil {
		gcru.SetTotalRewardAmount(*d)
	}
	return gcru
}

// ClearTotalRewardAmount clears the value of the "total_reward_amount" field.
func (gcru *GoodCoinRewardUpdate) ClearTotalRewardAmount() *GoodCoinRewardUpdate {
	gcru.mutation.ClearTotalRewardAmount()
	return gcru
}

// Mutation returns the GoodCoinRewardMutation object of the builder.
func (gcru *GoodCoinRewardUpdate) Mutation() *GoodCoinRewardMutation {
	return gcru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gcru *GoodCoinRewardUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := gcru.defaults(); err != nil {
		return 0, err
	}
	if len(gcru.hooks) == 0 {
		affected, err = gcru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodCoinRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcru.mutation = mutation
			affected, err = gcru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcru.hooks) - 1; i >= 0; i-- {
			if gcru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gcru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gcru *GoodCoinRewardUpdate) SaveX(ctx context.Context) int {
	affected, err := gcru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gcru *GoodCoinRewardUpdate) Exec(ctx context.Context) error {
	_, err := gcru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcru *GoodCoinRewardUpdate) ExecX(ctx context.Context) {
	if err := gcru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcru *GoodCoinRewardUpdate) defaults() error {
	if _, ok := gcru.mutation.UpdatedAt(); !ok {
		if goodcoinreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodcoinreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodcoinreward.UpdateDefaultUpdatedAt()
		gcru.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gcru *GoodCoinRewardUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodCoinRewardUpdate {
	gcru.modifiers = append(gcru.modifiers, modifiers...)
	return gcru
}

func (gcru *GoodCoinRewardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodcoinreward.Table,
			Columns: goodcoinreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodcoinreward.FieldID,
			},
		},
	}
	if ps := gcru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcru.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldCreatedAt,
		})
	}
	if value, ok := gcru.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldCreatedAt,
		})
	}
	if value, ok := gcru.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := gcru.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := gcru.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldDeletedAt,
		})
	}
	if value, ok := gcru.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldDeletedAt,
		})
	}
	if value, ok := gcru.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldEntID,
		})
	}
	if value, ok := gcru.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldGoodID,
		})
	}
	if gcru.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoinreward.FieldGoodID,
		})
	}
	if value, ok := gcru.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldCoinTypeID,
		})
	}
	if gcru.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoinreward.FieldCoinTypeID,
		})
	}
	if value, ok := gcru.mutation.RewardTid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldRewardTid,
		})
	}
	if gcru.mutation.RewardTidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoinreward.FieldRewardTid,
		})
	}
	if value, ok := gcru.mutation.NextRewardStartAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldNextRewardStartAmount,
		})
	}
	if gcru.mutation.NextRewardStartAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldNextRewardStartAmount,
		})
	}
	if value, ok := gcru.mutation.LastRewardAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldLastRewardAmount,
		})
	}
	if gcru.mutation.LastRewardAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldLastRewardAmount,
		})
	}
	if value, ok := gcru.mutation.LastUnitRewardAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldLastUnitRewardAmount,
		})
	}
	if gcru.mutation.LastUnitRewardAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldLastUnitRewardAmount,
		})
	}
	if value, ok := gcru.mutation.TotalRewardAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldTotalRewardAmount,
		})
	}
	if gcru.mutation.TotalRewardAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldTotalRewardAmount,
		})
	}
	_spec.Modifiers = gcru.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, gcru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodcoinreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GoodCoinRewardUpdateOne is the builder for updating a single GoodCoinReward entity.
type GoodCoinRewardUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoodCoinRewardMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (gcruo *GoodCoinRewardUpdateOne) SetCreatedAt(u uint32) *GoodCoinRewardUpdateOne {
	gcruo.mutation.ResetCreatedAt()
	gcruo.mutation.SetCreatedAt(u)
	return gcruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableCreatedAt(u *uint32) *GoodCoinRewardUpdateOne {
	if u != nil {
		gcruo.SetCreatedAt(*u)
	}
	return gcruo
}

// AddCreatedAt adds u to the "created_at" field.
func (gcruo *GoodCoinRewardUpdateOne) AddCreatedAt(u int32) *GoodCoinRewardUpdateOne {
	gcruo.mutation.AddCreatedAt(u)
	return gcruo
}

// SetUpdatedAt sets the "updated_at" field.
func (gcruo *GoodCoinRewardUpdateOne) SetUpdatedAt(u uint32) *GoodCoinRewardUpdateOne {
	gcruo.mutation.ResetUpdatedAt()
	gcruo.mutation.SetUpdatedAt(u)
	return gcruo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gcruo *GoodCoinRewardUpdateOne) AddUpdatedAt(u int32) *GoodCoinRewardUpdateOne {
	gcruo.mutation.AddUpdatedAt(u)
	return gcruo
}

// SetDeletedAt sets the "deleted_at" field.
func (gcruo *GoodCoinRewardUpdateOne) SetDeletedAt(u uint32) *GoodCoinRewardUpdateOne {
	gcruo.mutation.ResetDeletedAt()
	gcruo.mutation.SetDeletedAt(u)
	return gcruo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableDeletedAt(u *uint32) *GoodCoinRewardUpdateOne {
	if u != nil {
		gcruo.SetDeletedAt(*u)
	}
	return gcruo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gcruo *GoodCoinRewardUpdateOne) AddDeletedAt(u int32) *GoodCoinRewardUpdateOne {
	gcruo.mutation.AddDeletedAt(u)
	return gcruo
}

// SetEntID sets the "ent_id" field.
func (gcruo *GoodCoinRewardUpdateOne) SetEntID(u uuid.UUID) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetEntID(u)
	return gcruo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableEntID(u *uuid.UUID) *GoodCoinRewardUpdateOne {
	if u != nil {
		gcruo.SetEntID(*u)
	}
	return gcruo
}

// SetGoodID sets the "good_id" field.
func (gcruo *GoodCoinRewardUpdateOne) SetGoodID(u uuid.UUID) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetGoodID(u)
	return gcruo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableGoodID(u *uuid.UUID) *GoodCoinRewardUpdateOne {
	if u != nil {
		gcruo.SetGoodID(*u)
	}
	return gcruo
}

// ClearGoodID clears the value of the "good_id" field.
func (gcruo *GoodCoinRewardUpdateOne) ClearGoodID() *GoodCoinRewardUpdateOne {
	gcruo.mutation.ClearGoodID()
	return gcruo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gcruo *GoodCoinRewardUpdateOne) SetCoinTypeID(u uuid.UUID) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetCoinTypeID(u)
	return gcruo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *GoodCoinRewardUpdateOne {
	if u != nil {
		gcruo.SetCoinTypeID(*u)
	}
	return gcruo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gcruo *GoodCoinRewardUpdateOne) ClearCoinTypeID() *GoodCoinRewardUpdateOne {
	gcruo.mutation.ClearCoinTypeID()
	return gcruo
}

// SetRewardTid sets the "reward_tid" field.
func (gcruo *GoodCoinRewardUpdateOne) SetRewardTid(u uuid.UUID) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetRewardTid(u)
	return gcruo
}

// SetNillableRewardTid sets the "reward_tid" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableRewardTid(u *uuid.UUID) *GoodCoinRewardUpdateOne {
	if u != nil {
		gcruo.SetRewardTid(*u)
	}
	return gcruo
}

// ClearRewardTid clears the value of the "reward_tid" field.
func (gcruo *GoodCoinRewardUpdateOne) ClearRewardTid() *GoodCoinRewardUpdateOne {
	gcruo.mutation.ClearRewardTid()
	return gcruo
}

// SetNextRewardStartAmount sets the "next_reward_start_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) SetNextRewardStartAmount(d decimal.Decimal) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetNextRewardStartAmount(d)
	return gcruo
}

// SetNillableNextRewardStartAmount sets the "next_reward_start_amount" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableNextRewardStartAmount(d *decimal.Decimal) *GoodCoinRewardUpdateOne {
	if d != nil {
		gcruo.SetNextRewardStartAmount(*d)
	}
	return gcruo
}

// ClearNextRewardStartAmount clears the value of the "next_reward_start_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) ClearNextRewardStartAmount() *GoodCoinRewardUpdateOne {
	gcruo.mutation.ClearNextRewardStartAmount()
	return gcruo
}

// SetLastRewardAmount sets the "last_reward_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) SetLastRewardAmount(d decimal.Decimal) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetLastRewardAmount(d)
	return gcruo
}

// SetNillableLastRewardAmount sets the "last_reward_amount" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableLastRewardAmount(d *decimal.Decimal) *GoodCoinRewardUpdateOne {
	if d != nil {
		gcruo.SetLastRewardAmount(*d)
	}
	return gcruo
}

// ClearLastRewardAmount clears the value of the "last_reward_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) ClearLastRewardAmount() *GoodCoinRewardUpdateOne {
	gcruo.mutation.ClearLastRewardAmount()
	return gcruo
}

// SetLastUnitRewardAmount sets the "last_unit_reward_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) SetLastUnitRewardAmount(d decimal.Decimal) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetLastUnitRewardAmount(d)
	return gcruo
}

// SetNillableLastUnitRewardAmount sets the "last_unit_reward_amount" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableLastUnitRewardAmount(d *decimal.Decimal) *GoodCoinRewardUpdateOne {
	if d != nil {
		gcruo.SetLastUnitRewardAmount(*d)
	}
	return gcruo
}

// ClearLastUnitRewardAmount clears the value of the "last_unit_reward_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) ClearLastUnitRewardAmount() *GoodCoinRewardUpdateOne {
	gcruo.mutation.ClearLastUnitRewardAmount()
	return gcruo
}

// SetTotalRewardAmount sets the "total_reward_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) SetTotalRewardAmount(d decimal.Decimal) *GoodCoinRewardUpdateOne {
	gcruo.mutation.SetTotalRewardAmount(d)
	return gcruo
}

// SetNillableTotalRewardAmount sets the "total_reward_amount" field if the given value is not nil.
func (gcruo *GoodCoinRewardUpdateOne) SetNillableTotalRewardAmount(d *decimal.Decimal) *GoodCoinRewardUpdateOne {
	if d != nil {
		gcruo.SetTotalRewardAmount(*d)
	}
	return gcruo
}

// ClearTotalRewardAmount clears the value of the "total_reward_amount" field.
func (gcruo *GoodCoinRewardUpdateOne) ClearTotalRewardAmount() *GoodCoinRewardUpdateOne {
	gcruo.mutation.ClearTotalRewardAmount()
	return gcruo
}

// Mutation returns the GoodCoinRewardMutation object of the builder.
func (gcruo *GoodCoinRewardUpdateOne) Mutation() *GoodCoinRewardMutation {
	return gcruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gcruo *GoodCoinRewardUpdateOne) Select(field string, fields ...string) *GoodCoinRewardUpdateOne {
	gcruo.fields = append([]string{field}, fields...)
	return gcruo
}

// Save executes the query and returns the updated GoodCoinReward entity.
func (gcruo *GoodCoinRewardUpdateOne) Save(ctx context.Context) (*GoodCoinReward, error) {
	var (
		err  error
		node *GoodCoinReward
	)
	if err := gcruo.defaults(); err != nil {
		return nil, err
	}
	if len(gcruo.hooks) == 0 {
		node, err = gcruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodCoinRewardMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcruo.mutation = mutation
			node, err = gcruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gcruo.hooks) - 1; i >= 0; i-- {
			if gcruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gcruo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gcruo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoodCoinReward)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoodCoinRewardMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gcruo *GoodCoinRewardUpdateOne) SaveX(ctx context.Context) *GoodCoinReward {
	node, err := gcruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gcruo *GoodCoinRewardUpdateOne) Exec(ctx context.Context) error {
	_, err := gcruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcruo *GoodCoinRewardUpdateOne) ExecX(ctx context.Context) {
	if err := gcruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcruo *GoodCoinRewardUpdateOne) defaults() error {
	if _, ok := gcruo.mutation.UpdatedAt(); !ok {
		if goodcoinreward.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodcoinreward.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodcoinreward.UpdateDefaultUpdatedAt()
		gcruo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gcruo *GoodCoinRewardUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodCoinRewardUpdateOne {
	gcruo.modifiers = append(gcruo.modifiers, modifiers...)
	return gcruo
}

func (gcruo *GoodCoinRewardUpdateOne) sqlSave(ctx context.Context) (_node *GoodCoinReward, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodcoinreward.Table,
			Columns: goodcoinreward.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodcoinreward.FieldID,
			},
		},
	}
	id, ok := gcruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoodCoinReward.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gcruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodcoinreward.FieldID)
		for _, f := range fields {
			if !goodcoinreward.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodcoinreward.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gcruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcruo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldCreatedAt,
		})
	}
	if value, ok := gcruo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldCreatedAt,
		})
	}
	if value, ok := gcruo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := gcruo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldUpdatedAt,
		})
	}
	if value, ok := gcruo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldDeletedAt,
		})
	}
	if value, ok := gcruo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoinreward.FieldDeletedAt,
		})
	}
	if value, ok := gcruo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldEntID,
		})
	}
	if value, ok := gcruo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldGoodID,
		})
	}
	if gcruo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoinreward.FieldGoodID,
		})
	}
	if value, ok := gcruo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldCoinTypeID,
		})
	}
	if gcruo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoinreward.FieldCoinTypeID,
		})
	}
	if value, ok := gcruo.mutation.RewardTid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoinreward.FieldRewardTid,
		})
	}
	if gcruo.mutation.RewardTidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoinreward.FieldRewardTid,
		})
	}
	if value, ok := gcruo.mutation.NextRewardStartAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldNextRewardStartAmount,
		})
	}
	if gcruo.mutation.NextRewardStartAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldNextRewardStartAmount,
		})
	}
	if value, ok := gcruo.mutation.LastRewardAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldLastRewardAmount,
		})
	}
	if gcruo.mutation.LastRewardAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldLastRewardAmount,
		})
	}
	if value, ok := gcruo.mutation.LastUnitRewardAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldLastUnitRewardAmount,
		})
	}
	if gcruo.mutation.LastUnitRewardAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldLastUnitRewardAmount,
		})
	}
	if value, ok := gcruo.mutation.TotalRewardAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodcoinreward.FieldTotalRewardAmount,
		})
	}
	if gcruo.mutation.TotalRewardAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodcoinreward.FieldTotalRewardAmount,
		})
	}
	_spec.Modifiers = gcruo.modifiers
	_node = &GoodCoinReward{config: gcruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gcruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodcoinreward.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
