// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulategood"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppSimulateGoodDelete is the builder for deleting a AppSimulateGood entity.
type AppSimulateGoodDelete struct {
	config
	hooks    []Hook
	mutation *AppSimulateGoodMutation
}

// Where appends a list predicates to the AppSimulateGoodDelete builder.
func (asgd *AppSimulateGoodDelete) Where(ps ...predicate.AppSimulateGood) *AppSimulateGoodDelete {
	asgd.mutation.Where(ps...)
	return asgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asgd *AppSimulateGoodDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(asgd.hooks) == 0 {
		affected, err = asgd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSimulateGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asgd.mutation = mutation
			affected, err = asgd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(asgd.hooks) - 1; i >= 0; i-- {
			if asgd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asgd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asgd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (asgd *AppSimulateGoodDelete) ExecX(ctx context.Context) int {
	n, err := asgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asgd *AppSimulateGoodDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appsimulategood.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appsimulategood.FieldID,
			},
		},
	}
	if ps := asgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppSimulateGoodDeleteOne is the builder for deleting a single AppSimulateGood entity.
type AppSimulateGoodDeleteOne struct {
	asgd *AppSimulateGoodDelete
}

// Exec executes the deletion query.
func (asgdo *AppSimulateGoodDeleteOne) Exec(ctx context.Context) error {
	n, err := asgdo.asgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appsimulategood.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asgdo *AppSimulateGoodDeleteOne) ExecX(ctx context.Context) {
	asgdo.asgd.ExecX(ctx)
}
