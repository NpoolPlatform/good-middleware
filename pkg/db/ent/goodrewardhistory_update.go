// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodrewardhistory"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// GoodRewardHistoryUpdate is the builder for updating GoodRewardHistory entities.
type GoodRewardHistoryUpdate struct {
	config
	hooks     []Hook
	mutation  *GoodRewardHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoodRewardHistoryUpdate builder.
func (grhu *GoodRewardHistoryUpdate) Where(ps ...predicate.GoodRewardHistory) *GoodRewardHistoryUpdate {
	grhu.mutation.Where(ps...)
	return grhu
}

// SetCreatedAt sets the "created_at" field.
func (grhu *GoodRewardHistoryUpdate) SetCreatedAt(u uint32) *GoodRewardHistoryUpdate {
	grhu.mutation.ResetCreatedAt()
	grhu.mutation.SetCreatedAt(u)
	return grhu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableCreatedAt(u *uint32) *GoodRewardHistoryUpdate {
	if u != nil {
		grhu.SetCreatedAt(*u)
	}
	return grhu
}

// AddCreatedAt adds u to the "created_at" field.
func (grhu *GoodRewardHistoryUpdate) AddCreatedAt(u int32) *GoodRewardHistoryUpdate {
	grhu.mutation.AddCreatedAt(u)
	return grhu
}

// SetUpdatedAt sets the "updated_at" field.
func (grhu *GoodRewardHistoryUpdate) SetUpdatedAt(u uint32) *GoodRewardHistoryUpdate {
	grhu.mutation.ResetUpdatedAt()
	grhu.mutation.SetUpdatedAt(u)
	return grhu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (grhu *GoodRewardHistoryUpdate) AddUpdatedAt(u int32) *GoodRewardHistoryUpdate {
	grhu.mutation.AddUpdatedAt(u)
	return grhu
}

// SetDeletedAt sets the "deleted_at" field.
func (grhu *GoodRewardHistoryUpdate) SetDeletedAt(u uint32) *GoodRewardHistoryUpdate {
	grhu.mutation.ResetDeletedAt()
	grhu.mutation.SetDeletedAt(u)
	return grhu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableDeletedAt(u *uint32) *GoodRewardHistoryUpdate {
	if u != nil {
		grhu.SetDeletedAt(*u)
	}
	return grhu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (grhu *GoodRewardHistoryUpdate) AddDeletedAt(u int32) *GoodRewardHistoryUpdate {
	grhu.mutation.AddDeletedAt(u)
	return grhu
}

// SetAppID sets the "app_id" field.
func (grhu *GoodRewardHistoryUpdate) SetAppID(u uuid.UUID) *GoodRewardHistoryUpdate {
	grhu.mutation.SetAppID(u)
	return grhu
}

// SetGoodID sets the "good_id" field.
func (grhu *GoodRewardHistoryUpdate) SetGoodID(u uuid.UUID) *GoodRewardHistoryUpdate {
	grhu.mutation.SetGoodID(u)
	return grhu
}

// SetRewardDate sets the "reward_date" field.
func (grhu *GoodRewardHistoryUpdate) SetRewardDate(u uint32) *GoodRewardHistoryUpdate {
	grhu.mutation.ResetRewardDate()
	grhu.mutation.SetRewardDate(u)
	return grhu
}

// SetNillableRewardDate sets the "reward_date" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableRewardDate(u *uint32) *GoodRewardHistoryUpdate {
	if u != nil {
		grhu.SetRewardDate(*u)
	}
	return grhu
}

// AddRewardDate adds u to the "reward_date" field.
func (grhu *GoodRewardHistoryUpdate) AddRewardDate(u int32) *GoodRewardHistoryUpdate {
	grhu.mutation.AddRewardDate(u)
	return grhu
}

// ClearRewardDate clears the value of the "reward_date" field.
func (grhu *GoodRewardHistoryUpdate) ClearRewardDate() *GoodRewardHistoryUpdate {
	grhu.mutation.ClearRewardDate()
	return grhu
}

// SetTid sets the "tid" field.
func (grhu *GoodRewardHistoryUpdate) SetTid(u uuid.UUID) *GoodRewardHistoryUpdate {
	grhu.mutation.SetTid(u)
	return grhu
}

// SetNillableTid sets the "tid" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableTid(u *uuid.UUID) *GoodRewardHistoryUpdate {
	if u != nil {
		grhu.SetTid(*u)
	}
	return grhu
}

// ClearTid clears the value of the "tid" field.
func (grhu *GoodRewardHistoryUpdate) ClearTid() *GoodRewardHistoryUpdate {
	grhu.mutation.ClearTid()
	return grhu
}

// SetAmount sets the "amount" field.
func (grhu *GoodRewardHistoryUpdate) SetAmount(d decimal.Decimal) *GoodRewardHistoryUpdate {
	grhu.mutation.SetAmount(d)
	return grhu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableAmount(d *decimal.Decimal) *GoodRewardHistoryUpdate {
	if d != nil {
		grhu.SetAmount(*d)
	}
	return grhu
}

// ClearAmount clears the value of the "amount" field.
func (grhu *GoodRewardHistoryUpdate) ClearAmount() *GoodRewardHistoryUpdate {
	grhu.mutation.ClearAmount()
	return grhu
}

// SetUnitAmount sets the "unit_amount" field.
func (grhu *GoodRewardHistoryUpdate) SetUnitAmount(d decimal.Decimal) *GoodRewardHistoryUpdate {
	grhu.mutation.SetUnitAmount(d)
	return grhu
}

// SetNillableUnitAmount sets the "unit_amount" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableUnitAmount(d *decimal.Decimal) *GoodRewardHistoryUpdate {
	if d != nil {
		grhu.SetUnitAmount(*d)
	}
	return grhu
}

// ClearUnitAmount clears the value of the "unit_amount" field.
func (grhu *GoodRewardHistoryUpdate) ClearUnitAmount() *GoodRewardHistoryUpdate {
	grhu.mutation.ClearUnitAmount()
	return grhu
}

// SetUnitNetAmount sets the "unit_net_amount" field.
func (grhu *GoodRewardHistoryUpdate) SetUnitNetAmount(d decimal.Decimal) *GoodRewardHistoryUpdate {
	grhu.mutation.SetUnitNetAmount(d)
	return grhu
}

// SetNillableUnitNetAmount sets the "unit_net_amount" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableUnitNetAmount(d *decimal.Decimal) *GoodRewardHistoryUpdate {
	if d != nil {
		grhu.SetUnitNetAmount(*d)
	}
	return grhu
}

// ClearUnitNetAmount clears the value of the "unit_net_amount" field.
func (grhu *GoodRewardHistoryUpdate) ClearUnitNetAmount() *GoodRewardHistoryUpdate {
	grhu.mutation.ClearUnitNetAmount()
	return grhu
}

// SetResult sets the "result" field.
func (grhu *GoodRewardHistoryUpdate) SetResult(s string) *GoodRewardHistoryUpdate {
	grhu.mutation.SetResult(s)
	return grhu
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (grhu *GoodRewardHistoryUpdate) SetNillableResult(s *string) *GoodRewardHistoryUpdate {
	if s != nil {
		grhu.SetResult(*s)
	}
	return grhu
}

// ClearResult clears the value of the "result" field.
func (grhu *GoodRewardHistoryUpdate) ClearResult() *GoodRewardHistoryUpdate {
	grhu.mutation.ClearResult()
	return grhu
}

// Mutation returns the GoodRewardHistoryMutation object of the builder.
func (grhu *GoodRewardHistoryUpdate) Mutation() *GoodRewardHistoryMutation {
	return grhu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (grhu *GoodRewardHistoryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := grhu.defaults(); err != nil {
		return 0, err
	}
	if len(grhu.hooks) == 0 {
		affected, err = grhu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodRewardHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			grhu.mutation = mutation
			affected, err = grhu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(grhu.hooks) - 1; i >= 0; i-- {
			if grhu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = grhu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, grhu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (grhu *GoodRewardHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := grhu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (grhu *GoodRewardHistoryUpdate) Exec(ctx context.Context) error {
	_, err := grhu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grhu *GoodRewardHistoryUpdate) ExecX(ctx context.Context) {
	if err := grhu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (grhu *GoodRewardHistoryUpdate) defaults() error {
	if _, ok := grhu.mutation.UpdatedAt(); !ok {
		if goodrewardhistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodrewardhistory.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodrewardhistory.UpdateDefaultUpdatedAt()
		grhu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (grhu *GoodRewardHistoryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodRewardHistoryUpdate {
	grhu.modifiers = append(grhu.modifiers, modifiers...)
	return grhu
}

func (grhu *GoodRewardHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodrewardhistory.Table,
			Columns: goodrewardhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodrewardhistory.FieldID,
			},
		},
	}
	if ps := grhu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := grhu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldCreatedAt,
		})
	}
	if value, ok := grhu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldCreatedAt,
		})
	}
	if value, ok := grhu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldUpdatedAt,
		})
	}
	if value, ok := grhu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldUpdatedAt,
		})
	}
	if value, ok := grhu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldDeletedAt,
		})
	}
	if value, ok := grhu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldDeletedAt,
		})
	}
	if value, ok := grhu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrewardhistory.FieldAppID,
		})
	}
	if value, ok := grhu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrewardhistory.FieldGoodID,
		})
	}
	if value, ok := grhu.mutation.RewardDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldRewardDate,
		})
	}
	if value, ok := grhu.mutation.AddedRewardDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldRewardDate,
		})
	}
	if grhu.mutation.RewardDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: goodrewardhistory.FieldRewardDate,
		})
	}
	if value, ok := grhu.mutation.Tid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrewardhistory.FieldTid,
		})
	}
	if grhu.mutation.TidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodrewardhistory.FieldTid,
		})
	}
	if value, ok := grhu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodrewardhistory.FieldAmount,
		})
	}
	if grhu.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodrewardhistory.FieldAmount,
		})
	}
	if value, ok := grhu.mutation.UnitAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodrewardhistory.FieldUnitAmount,
		})
	}
	if grhu.mutation.UnitAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodrewardhistory.FieldUnitAmount,
		})
	}
	if value, ok := grhu.mutation.UnitNetAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodrewardhistory.FieldUnitNetAmount,
		})
	}
	if grhu.mutation.UnitNetAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodrewardhistory.FieldUnitNetAmount,
		})
	}
	if value, ok := grhu.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodrewardhistory.FieldResult,
		})
	}
	if grhu.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodrewardhistory.FieldResult,
		})
	}
	_spec.Modifiers = grhu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, grhu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodrewardhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GoodRewardHistoryUpdateOne is the builder for updating a single GoodRewardHistory entity.
type GoodRewardHistoryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoodRewardHistoryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetCreatedAt(u uint32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ResetCreatedAt()
	grhuo.mutation.SetCreatedAt(u)
	return grhuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableCreatedAt(u *uint32) *GoodRewardHistoryUpdateOne {
	if u != nil {
		grhuo.SetCreatedAt(*u)
	}
	return grhuo
}

// AddCreatedAt adds u to the "created_at" field.
func (grhuo *GoodRewardHistoryUpdateOne) AddCreatedAt(u int32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.AddCreatedAt(u)
	return grhuo
}

// SetUpdatedAt sets the "updated_at" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetUpdatedAt(u uint32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ResetUpdatedAt()
	grhuo.mutation.SetUpdatedAt(u)
	return grhuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (grhuo *GoodRewardHistoryUpdateOne) AddUpdatedAt(u int32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.AddUpdatedAt(u)
	return grhuo
}

// SetDeletedAt sets the "deleted_at" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetDeletedAt(u uint32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ResetDeletedAt()
	grhuo.mutation.SetDeletedAt(u)
	return grhuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableDeletedAt(u *uint32) *GoodRewardHistoryUpdateOne {
	if u != nil {
		grhuo.SetDeletedAt(*u)
	}
	return grhuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (grhuo *GoodRewardHistoryUpdateOne) AddDeletedAt(u int32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.AddDeletedAt(u)
	return grhuo
}

// SetAppID sets the "app_id" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetAppID(u uuid.UUID) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.SetAppID(u)
	return grhuo
}

// SetGoodID sets the "good_id" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetGoodID(u uuid.UUID) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.SetGoodID(u)
	return grhuo
}

// SetRewardDate sets the "reward_date" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetRewardDate(u uint32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ResetRewardDate()
	grhuo.mutation.SetRewardDate(u)
	return grhuo
}

// SetNillableRewardDate sets the "reward_date" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableRewardDate(u *uint32) *GoodRewardHistoryUpdateOne {
	if u != nil {
		grhuo.SetRewardDate(*u)
	}
	return grhuo
}

// AddRewardDate adds u to the "reward_date" field.
func (grhuo *GoodRewardHistoryUpdateOne) AddRewardDate(u int32) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.AddRewardDate(u)
	return grhuo
}

// ClearRewardDate clears the value of the "reward_date" field.
func (grhuo *GoodRewardHistoryUpdateOne) ClearRewardDate() *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ClearRewardDate()
	return grhuo
}

// SetTid sets the "tid" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetTid(u uuid.UUID) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.SetTid(u)
	return grhuo
}

// SetNillableTid sets the "tid" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableTid(u *uuid.UUID) *GoodRewardHistoryUpdateOne {
	if u != nil {
		grhuo.SetTid(*u)
	}
	return grhuo
}

// ClearTid clears the value of the "tid" field.
func (grhuo *GoodRewardHistoryUpdateOne) ClearTid() *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ClearTid()
	return grhuo
}

// SetAmount sets the "amount" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetAmount(d decimal.Decimal) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.SetAmount(d)
	return grhuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableAmount(d *decimal.Decimal) *GoodRewardHistoryUpdateOne {
	if d != nil {
		grhuo.SetAmount(*d)
	}
	return grhuo
}

// ClearAmount clears the value of the "amount" field.
func (grhuo *GoodRewardHistoryUpdateOne) ClearAmount() *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ClearAmount()
	return grhuo
}

// SetUnitAmount sets the "unit_amount" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetUnitAmount(d decimal.Decimal) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.SetUnitAmount(d)
	return grhuo
}

// SetNillableUnitAmount sets the "unit_amount" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableUnitAmount(d *decimal.Decimal) *GoodRewardHistoryUpdateOne {
	if d != nil {
		grhuo.SetUnitAmount(*d)
	}
	return grhuo
}

// ClearUnitAmount clears the value of the "unit_amount" field.
func (grhuo *GoodRewardHistoryUpdateOne) ClearUnitAmount() *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ClearUnitAmount()
	return grhuo
}

// SetUnitNetAmount sets the "unit_net_amount" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetUnitNetAmount(d decimal.Decimal) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.SetUnitNetAmount(d)
	return grhuo
}

// SetNillableUnitNetAmount sets the "unit_net_amount" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableUnitNetAmount(d *decimal.Decimal) *GoodRewardHistoryUpdateOne {
	if d != nil {
		grhuo.SetUnitNetAmount(*d)
	}
	return grhuo
}

// ClearUnitNetAmount clears the value of the "unit_net_amount" field.
func (grhuo *GoodRewardHistoryUpdateOne) ClearUnitNetAmount() *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ClearUnitNetAmount()
	return grhuo
}

// SetResult sets the "result" field.
func (grhuo *GoodRewardHistoryUpdateOne) SetResult(s string) *GoodRewardHistoryUpdateOne {
	grhuo.mutation.SetResult(s)
	return grhuo
}

// SetNillableResult sets the "result" field if the given value is not nil.
func (grhuo *GoodRewardHistoryUpdateOne) SetNillableResult(s *string) *GoodRewardHistoryUpdateOne {
	if s != nil {
		grhuo.SetResult(*s)
	}
	return grhuo
}

// ClearResult clears the value of the "result" field.
func (grhuo *GoodRewardHistoryUpdateOne) ClearResult() *GoodRewardHistoryUpdateOne {
	grhuo.mutation.ClearResult()
	return grhuo
}

// Mutation returns the GoodRewardHistoryMutation object of the builder.
func (grhuo *GoodRewardHistoryUpdateOne) Mutation() *GoodRewardHistoryMutation {
	return grhuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (grhuo *GoodRewardHistoryUpdateOne) Select(field string, fields ...string) *GoodRewardHistoryUpdateOne {
	grhuo.fields = append([]string{field}, fields...)
	return grhuo
}

// Save executes the query and returns the updated GoodRewardHistory entity.
func (grhuo *GoodRewardHistoryUpdateOne) Save(ctx context.Context) (*GoodRewardHistory, error) {
	var (
		err  error
		node *GoodRewardHistory
	)
	if err := grhuo.defaults(); err != nil {
		return nil, err
	}
	if len(grhuo.hooks) == 0 {
		node, err = grhuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodRewardHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			grhuo.mutation = mutation
			node, err = grhuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(grhuo.hooks) - 1; i >= 0; i-- {
			if grhuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = grhuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, grhuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoodRewardHistory)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoodRewardHistoryMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (grhuo *GoodRewardHistoryUpdateOne) SaveX(ctx context.Context) *GoodRewardHistory {
	node, err := grhuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (grhuo *GoodRewardHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := grhuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (grhuo *GoodRewardHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := grhuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (grhuo *GoodRewardHistoryUpdateOne) defaults() error {
	if _, ok := grhuo.mutation.UpdatedAt(); !ok {
		if goodrewardhistory.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodrewardhistory.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodrewardhistory.UpdateDefaultUpdatedAt()
		grhuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (grhuo *GoodRewardHistoryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodRewardHistoryUpdateOne {
	grhuo.modifiers = append(grhuo.modifiers, modifiers...)
	return grhuo
}

func (grhuo *GoodRewardHistoryUpdateOne) sqlSave(ctx context.Context) (_node *GoodRewardHistory, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodrewardhistory.Table,
			Columns: goodrewardhistory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: goodrewardhistory.FieldID,
			},
		},
	}
	id, ok := grhuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoodRewardHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := grhuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodrewardhistory.FieldID)
		for _, f := range fields {
			if !goodrewardhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodrewardhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := grhuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := grhuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldCreatedAt,
		})
	}
	if value, ok := grhuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldCreatedAt,
		})
	}
	if value, ok := grhuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldUpdatedAt,
		})
	}
	if value, ok := grhuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldUpdatedAt,
		})
	}
	if value, ok := grhuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldDeletedAt,
		})
	}
	if value, ok := grhuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldDeletedAt,
		})
	}
	if value, ok := grhuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrewardhistory.FieldAppID,
		})
	}
	if value, ok := grhuo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrewardhistory.FieldGoodID,
		})
	}
	if value, ok := grhuo.mutation.RewardDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldRewardDate,
		})
	}
	if value, ok := grhuo.mutation.AddedRewardDate(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodrewardhistory.FieldRewardDate,
		})
	}
	if grhuo.mutation.RewardDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: goodrewardhistory.FieldRewardDate,
		})
	}
	if value, ok := grhuo.mutation.Tid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodrewardhistory.FieldTid,
		})
	}
	if grhuo.mutation.TidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodrewardhistory.FieldTid,
		})
	}
	if value, ok := grhuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodrewardhistory.FieldAmount,
		})
	}
	if grhuo.mutation.AmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodrewardhistory.FieldAmount,
		})
	}
	if value, ok := grhuo.mutation.UnitAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodrewardhistory.FieldUnitAmount,
		})
	}
	if grhuo.mutation.UnitAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodrewardhistory.FieldUnitAmount,
		})
	}
	if value, ok := grhuo.mutation.UnitNetAmount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: goodrewardhistory.FieldUnitNetAmount,
		})
	}
	if grhuo.mutation.UnitNetAmountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: goodrewardhistory.FieldUnitNetAmount,
		})
	}
	if value, ok := grhuo.mutation.Result(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: goodrewardhistory.FieldResult,
		})
	}
	if grhuo.mutation.ResultCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: goodrewardhistory.FieldResult,
		})
	}
	_spec.Modifiers = grhuo.modifiers
	_node = &GoodRewardHistory{config: grhuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, grhuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodrewardhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
