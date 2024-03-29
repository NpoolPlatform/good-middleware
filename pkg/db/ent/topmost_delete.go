// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
)

// TopMostDelete is the builder for deleting a TopMost entity.
type TopMostDelete struct {
	config
	hooks    []Hook
	mutation *TopMostMutation
}

// Where appends a list predicates to the TopMostDelete builder.
func (tmd *TopMostDelete) Where(ps ...predicate.TopMost) *TopMostDelete {
	tmd.mutation.Where(ps...)
	return tmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (tmd *TopMostDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tmd.hooks) == 0 {
		affected, err = tmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopMostMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tmd.mutation = mutation
			affected, err = tmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tmd.hooks) - 1; i >= 0; i-- {
			if tmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmd *TopMostDelete) ExecX(ctx context.Context) int {
	n, err := tmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (tmd *TopMostDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: topmost.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmost.FieldID,
			},
		},
	}
	if ps := tmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, tmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// TopMostDeleteOne is the builder for deleting a single TopMost entity.
type TopMostDeleteOne struct {
	tmd *TopMostDelete
}

// Exec executes the deletion query.
func (tmdo *TopMostDeleteOne) Exec(ctx context.Context) error {
	n, err := tmdo.tmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{topmost.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tmdo *TopMostDeleteOne) ExecX(ctx context.Context) {
	tmdo.tmd.ExecX(ctx)
}
