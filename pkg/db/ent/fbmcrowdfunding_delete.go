// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/fbmcrowdfunding"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// FbmCrowdFundingDelete is the builder for deleting a FbmCrowdFunding entity.
type FbmCrowdFundingDelete struct {
	config
	hooks    []Hook
	mutation *FbmCrowdFundingMutation
}

// Where appends a list predicates to the FbmCrowdFundingDelete builder.
func (fcfd *FbmCrowdFundingDelete) Where(ps ...predicate.FbmCrowdFunding) *FbmCrowdFundingDelete {
	fcfd.mutation.Where(ps...)
	return fcfd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fcfd *FbmCrowdFundingDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fcfd.hooks) == 0 {
		affected, err = fcfd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FbmCrowdFundingMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fcfd.mutation = mutation
			affected, err = fcfd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fcfd.hooks) - 1; i >= 0; i-- {
			if fcfd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fcfd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fcfd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcfd *FbmCrowdFundingDelete) ExecX(ctx context.Context) int {
	n, err := fcfd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fcfd *FbmCrowdFundingDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fbmcrowdfunding.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: fbmcrowdfunding.FieldID,
			},
		},
	}
	if ps := fcfd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fcfd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// FbmCrowdFundingDeleteOne is the builder for deleting a single FbmCrowdFunding entity.
type FbmCrowdFundingDeleteOne struct {
	fcfd *FbmCrowdFundingDelete
}

// Exec executes the deletion query.
func (fcfdo *FbmCrowdFundingDeleteOne) Exec(ctx context.Context) error {
	n, err := fcfdo.fcfd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fbmcrowdfunding.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fcfdo *FbmCrowdFundingDeleteOne) ExecX(ctx context.Context) {
	fcfdo.fcfd.ExecX(ctx)
}
