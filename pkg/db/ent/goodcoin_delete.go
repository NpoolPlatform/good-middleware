// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// GoodCoinDelete is the builder for deleting a GoodCoin entity.
type GoodCoinDelete struct {
	config
	hooks    []Hook
	mutation *GoodCoinMutation
}

// Where appends a list predicates to the GoodCoinDelete builder.
func (gcd *GoodCoinDelete) Where(ps ...predicate.GoodCoin) *GoodCoinDelete {
	gcd.mutation.Where(ps...)
	return gcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gcd *GoodCoinDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gcd.hooks) == 0 {
		affected, err = gcd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodCoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcd.mutation = mutation
			affected, err = gcd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcd.hooks) - 1; i >= 0; i-- {
			if gcd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gcd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcd *GoodCoinDelete) ExecX(ctx context.Context) int {
	n, err := gcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gcd *GoodCoinDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: goodcoin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodcoin.FieldID,
			},
		},
	}
	if ps := gcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// GoodCoinDeleteOne is the builder for deleting a single GoodCoin entity.
type GoodCoinDeleteOne struct {
	gcd *GoodCoinDelete
}

// Exec executes the deletion query.
func (gcdo *GoodCoinDeleteOne) Exec(ctx context.Context) error {
	n, err := gcdo.gcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goodcoin.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gcdo *GoodCoinDeleteOne) ExecX(ctx context.Context) {
	gcdo.gcd.ExecX(ctx)
}
