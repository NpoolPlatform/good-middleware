// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstocklock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppStockLockDelete is the builder for deleting a AppStockLock entity.
type AppStockLockDelete struct {
	config
	hooks    []Hook
	mutation *AppStockLockMutation
}

// Where appends a list predicates to the AppStockLockDelete builder.
func (asld *AppStockLockDelete) Where(ps ...predicate.AppStockLock) *AppStockLockDelete {
	asld.mutation.Where(ps...)
	return asld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (asld *AppStockLockDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(asld.hooks) == 0 {
		affected, err = asld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppStockLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			asld.mutation = mutation
			affected, err = asld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(asld.hooks) - 1; i >= 0; i-- {
			if asld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, asld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (asld *AppStockLockDelete) ExecX(ctx context.Context) int {
	n, err := asld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (asld *AppStockLockDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: appstocklock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appstocklock.FieldID,
			},
		},
	}
	if ps := asld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, asld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// AppStockLockDeleteOne is the builder for deleting a single AppStockLock entity.
type AppStockLockDeleteOne struct {
	asld *AppStockLockDelete
}

// Exec executes the deletion query.
func (asldo *AppStockLockDeleteOne) Exec(ctx context.Context) error {
	n, err := asldo.asld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{appstocklock.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (asldo *AppStockLockDeleteOne) ExecX(ctx context.Context) {
	asldo.asld.ExecX(ctx)
}