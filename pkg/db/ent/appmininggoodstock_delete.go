// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appmininggoodstock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppMiningGoodStockDelete is the builder for deleting a AppMiningGoodStock entity.
type AppMiningGoodStockDelete struct {
	config
	hooks    []Hook
	mutation *AppMiningGoodStockMutation
}

// Where appends a list predicates to the AppMiningGoodStockDelete builder.
func (amgsd *AppMiningGoodStockDelete) Where(ps ...predicate.AppMiningGoodStock) *AppMiningGoodStockDelete {
	amgsd.mutation.Where(ps...)
	return amgsd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (amgsd *AppMiningGoodStockDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(amgsd.hooks) == 0 {
		affected, err = amgsd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppMiningGoodStockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			amgsd.mutation = mutation
			affected, err = amgsd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(amgsd.hooks) - 1; i >= 0; i-- {
			if amgsd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = amgsd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, amgsd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (amgsd *AppMiningGoodStockDelete) ExecX(ctx context.Context) int {
	n, err := amgsd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (amgsd *AppMiningGoodStockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appmininggoodstock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appmininggoodstock.FieldID,
			},
		},
	}
	if ps := amgsd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, amgsd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppMiningGoodStockDeleteOne is the builder for deleting a single AppMiningGoodStock entity.
type AppMiningGoodStockDeleteOne struct {
	amgsd *AppMiningGoodStockDelete
}

// Exec executes the deletion query.
func (amgsdo *AppMiningGoodStockDeleteOne) Exec(ctx context.Context) error {
	n, err := amgsdo.amgsd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appmininggoodstock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (amgsdo *AppMiningGoodStockDeleteOne) ExecX(ctx context.Context) {
	amgsdo.amgsd.ExecX(ctx)
}