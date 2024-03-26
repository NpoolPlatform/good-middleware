// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// MiningGoodStockDelete is the builder for deleting a MiningGoodStock entity.
type MiningGoodStockDelete struct {
	config
	hooks    []Hook
	mutation *MiningGoodStockMutation
}

// Where appends a list predicates to the MiningGoodStockDelete builder.
func (mgsd *MiningGoodStockDelete) Where(ps ...predicate.MiningGoodStock) *MiningGoodStockDelete {
	mgsd.mutation.Where(ps...)
	return mgsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mgsd *MiningGoodStockDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mgsd.hooks) == 0 {
		affected, err = mgsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MiningGoodStockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mgsd.mutation = mutation
			affected, err = mgsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mgsd.hooks) - 1; i >= 0; i-- {
			if mgsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mgsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mgsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mgsd *MiningGoodStockDelete) ExecX(ctx context.Context) int {
	n, err := mgsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mgsd *MiningGoodStockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: mininggoodstock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: mininggoodstock.FieldID,
			},
		},
	}
	if ps := mgsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mgsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// MiningGoodStockDeleteOne is the builder for deleting a single MiningGoodStock entity.
type MiningGoodStockDeleteOne struct {
	mgsd *MiningGoodStockDelete
}

// Exec executes the deletion query.
func (mgsdo *MiningGoodStockDeleteOne) Exec(ctx context.Context) error {
	n, err := mgsdo.mgsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mininggoodstock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mgsdo *MiningGoodStockDeleteOne) ExecX(ctx context.Context) {
	mgsdo.mgsd.ExecX(ctx)
}
