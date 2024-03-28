// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// DeviceManufacturerDelete is the builder for deleting a DeviceManufacturer entity.
type DeviceManufacturerDelete struct {
	config
	hooks    []Hook
	mutation *DeviceManufacturerMutation
}

// Where appends a list predicates to the DeviceManufacturerDelete builder.
func (dmd *DeviceManufacturerDelete) Where(ps ...predicate.DeviceManufacturer) *DeviceManufacturerDelete {
	dmd.mutation.Where(ps...)
	return dmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dmd *DeviceManufacturerDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dmd.hooks) == 0 {
		affected, err = dmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceManufacturerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dmd.mutation = mutation
			affected, err = dmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dmd.hooks) - 1; i >= 0; i-- {
			if dmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmd *DeviceManufacturerDelete) ExecX(ctx context.Context) int {
	n, err := dmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dmd *DeviceManufacturerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: devicemanufacturer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: devicemanufacturer.FieldID,
			},
		},
	}
	if ps := dmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// DeviceManufacturerDeleteOne is the builder for deleting a single DeviceManufacturer entity.
type DeviceManufacturerDeleteOne struct {
	dmd *DeviceManufacturerDelete
}

// Exec executes the deletion query.
func (dmdo *DeviceManufacturerDeleteOne) Exec(ctx context.Context) error {
	n, err := dmdo.dmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{devicemanufacturer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dmdo *DeviceManufacturerDeleteOne) ExecX(ctx context.Context) {
	dmdo.dmd.ExecX(ctx)
}