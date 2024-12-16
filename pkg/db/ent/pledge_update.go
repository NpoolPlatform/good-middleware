// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/pledge"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// PledgeUpdate is the builder for updating Pledge entities.
type PledgeUpdate struct {
	config
	hooks     []Hook
	mutation  *PledgeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PledgeUpdate builder.
func (pu *PledgeUpdate) Where(ps ...predicate.Pledge) *PledgeUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedAt sets the "created_at" field.
func (pu *PledgeUpdate) SetCreatedAt(u uint32) *PledgeUpdate {
	pu.mutation.ResetCreatedAt()
	pu.mutation.SetCreatedAt(u)
	return pu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pu *PledgeUpdate) SetNillableCreatedAt(u *uint32) *PledgeUpdate {
	if u != nil {
		pu.SetCreatedAt(*u)
	}
	return pu
}

// AddCreatedAt adds u to the "created_at" field.
func (pu *PledgeUpdate) AddCreatedAt(u int32) *PledgeUpdate {
	pu.mutation.AddCreatedAt(u)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PledgeUpdate) SetUpdatedAt(u uint32) *PledgeUpdate {
	pu.mutation.ResetUpdatedAt()
	pu.mutation.SetUpdatedAt(u)
	return pu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (pu *PledgeUpdate) AddUpdatedAt(u int32) *PledgeUpdate {
	pu.mutation.AddUpdatedAt(u)
	return pu
}

// SetDeletedAt sets the "deleted_at" field.
func (pu *PledgeUpdate) SetDeletedAt(u uint32) *PledgeUpdate {
	pu.mutation.ResetDeletedAt()
	pu.mutation.SetDeletedAt(u)
	return pu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (pu *PledgeUpdate) SetNillableDeletedAt(u *uint32) *PledgeUpdate {
	if u != nil {
		pu.SetDeletedAt(*u)
	}
	return pu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (pu *PledgeUpdate) AddDeletedAt(u int32) *PledgeUpdate {
	pu.mutation.AddDeletedAt(u)
	return pu
}

// SetEntID sets the "ent_id" field.
func (pu *PledgeUpdate) SetEntID(u uuid.UUID) *PledgeUpdate {
	pu.mutation.SetEntID(u)
	return pu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (pu *PledgeUpdate) SetNillableEntID(u *uuid.UUID) *PledgeUpdate {
	if u != nil {
		pu.SetEntID(*u)
	}
	return pu
}

// SetGoodID sets the "good_id" field.
func (pu *PledgeUpdate) SetGoodID(u uuid.UUID) *PledgeUpdate {
	pu.mutation.SetGoodID(u)
	return pu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (pu *PledgeUpdate) SetNillableGoodID(u *uuid.UUID) *PledgeUpdate {
	if u != nil {
		pu.SetGoodID(*u)
	}
	return pu
}

// ClearGoodID clears the value of the "good_id" field.
func (pu *PledgeUpdate) ClearGoodID() *PledgeUpdate {
	pu.mutation.ClearGoodID()
	return pu
}

// SetContractCodeURL sets the "contract_code_url" field.
func (pu *PledgeUpdate) SetContractCodeURL(s string) *PledgeUpdate {
	pu.mutation.SetContractCodeURL(s)
	return pu
}

// SetNillableContractCodeURL sets the "contract_code_url" field if the given value is not nil.
func (pu *PledgeUpdate) SetNillableContractCodeURL(s *string) *PledgeUpdate {
	if s != nil {
		pu.SetContractCodeURL(*s)
	}
	return pu
}

// ClearContractCodeURL clears the value of the "contract_code_url" field.
func (pu *PledgeUpdate) ClearContractCodeURL() *PledgeUpdate {
	pu.mutation.ClearContractCodeURL()
	return pu
}

// SetContractCodeBranch sets the "contract_code_branch" field.
func (pu *PledgeUpdate) SetContractCodeBranch(s string) *PledgeUpdate {
	pu.mutation.SetContractCodeBranch(s)
	return pu
}

// SetNillableContractCodeBranch sets the "contract_code_branch" field if the given value is not nil.
func (pu *PledgeUpdate) SetNillableContractCodeBranch(s *string) *PledgeUpdate {
	if s != nil {
		pu.SetContractCodeBranch(*s)
	}
	return pu
}

// ClearContractCodeBranch clears the value of the "contract_code_branch" field.
func (pu *PledgeUpdate) ClearContractCodeBranch() *PledgeUpdate {
	pu.mutation.ClearContractCodeBranch()
	return pu
}

// SetContractState sets the "contract_state" field.
func (pu *PledgeUpdate) SetContractState(s string) *PledgeUpdate {
	pu.mutation.SetContractState(s)
	return pu
}

// SetNillableContractState sets the "contract_state" field if the given value is not nil.
func (pu *PledgeUpdate) SetNillableContractState(s *string) *PledgeUpdate {
	if s != nil {
		pu.SetContractState(*s)
	}
	return pu
}

// ClearContractState clears the value of the "contract_state" field.
func (pu *PledgeUpdate) ClearContractState() *PledgeUpdate {
	pu.mutation.ClearContractState()
	return pu
}

// Mutation returns the PledgeMutation object of the builder.
func (pu *PledgeUpdate) Mutation() *PledgeMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PledgeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := pu.defaults(); err != nil {
		return 0, err
	}
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PledgeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PledgeUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PledgeUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PledgeUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PledgeUpdate) defaults() error {
	if _, ok := pu.mutation.UpdatedAt(); !ok {
		if pledge.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pledge.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pledge.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PledgeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PledgeUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PledgeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pledge.Table,
			Columns: pledge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pledge.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldCreatedAt,
		})
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldUpdatedAt,
		})
	}
	if value, ok := pu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldDeletedAt,
		})
	}
	if value, ok := pu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldDeletedAt,
		})
	}
	if value, ok := pu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pledge.FieldEntID,
		})
	}
	if value, ok := pu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pledge.FieldGoodID,
		})
	}
	if pu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: pledge.FieldGoodID,
		})
	}
	if value, ok := pu.mutation.ContractCodeURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pledge.FieldContractCodeURL,
		})
	}
	if pu.mutation.ContractCodeURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pledge.FieldContractCodeURL,
		})
	}
	if value, ok := pu.mutation.ContractCodeBranch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pledge.FieldContractCodeBranch,
		})
	}
	if pu.mutation.ContractCodeBranchCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pledge.FieldContractCodeBranch,
		})
	}
	if value, ok := pu.mutation.ContractState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pledge.FieldContractState,
		})
	}
	if pu.mutation.ContractStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pledge.FieldContractState,
		})
	}
	_spec.Modifiers = pu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pledge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PledgeUpdateOne is the builder for updating a single Pledge entity.
type PledgeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PledgeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (puo *PledgeUpdateOne) SetCreatedAt(u uint32) *PledgeUpdateOne {
	puo.mutation.ResetCreatedAt()
	puo.mutation.SetCreatedAt(u)
	return puo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (puo *PledgeUpdateOne) SetNillableCreatedAt(u *uint32) *PledgeUpdateOne {
	if u != nil {
		puo.SetCreatedAt(*u)
	}
	return puo
}

// AddCreatedAt adds u to the "created_at" field.
func (puo *PledgeUpdateOne) AddCreatedAt(u int32) *PledgeUpdateOne {
	puo.mutation.AddCreatedAt(u)
	return puo
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PledgeUpdateOne) SetUpdatedAt(u uint32) *PledgeUpdateOne {
	puo.mutation.ResetUpdatedAt()
	puo.mutation.SetUpdatedAt(u)
	return puo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (puo *PledgeUpdateOne) AddUpdatedAt(u int32) *PledgeUpdateOne {
	puo.mutation.AddUpdatedAt(u)
	return puo
}

// SetDeletedAt sets the "deleted_at" field.
func (puo *PledgeUpdateOne) SetDeletedAt(u uint32) *PledgeUpdateOne {
	puo.mutation.ResetDeletedAt()
	puo.mutation.SetDeletedAt(u)
	return puo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (puo *PledgeUpdateOne) SetNillableDeletedAt(u *uint32) *PledgeUpdateOne {
	if u != nil {
		puo.SetDeletedAt(*u)
	}
	return puo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (puo *PledgeUpdateOne) AddDeletedAt(u int32) *PledgeUpdateOne {
	puo.mutation.AddDeletedAt(u)
	return puo
}

// SetEntID sets the "ent_id" field.
func (puo *PledgeUpdateOne) SetEntID(u uuid.UUID) *PledgeUpdateOne {
	puo.mutation.SetEntID(u)
	return puo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (puo *PledgeUpdateOne) SetNillableEntID(u *uuid.UUID) *PledgeUpdateOne {
	if u != nil {
		puo.SetEntID(*u)
	}
	return puo
}

// SetGoodID sets the "good_id" field.
func (puo *PledgeUpdateOne) SetGoodID(u uuid.UUID) *PledgeUpdateOne {
	puo.mutation.SetGoodID(u)
	return puo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (puo *PledgeUpdateOne) SetNillableGoodID(u *uuid.UUID) *PledgeUpdateOne {
	if u != nil {
		puo.SetGoodID(*u)
	}
	return puo
}

// ClearGoodID clears the value of the "good_id" field.
func (puo *PledgeUpdateOne) ClearGoodID() *PledgeUpdateOne {
	puo.mutation.ClearGoodID()
	return puo
}

// SetContractCodeURL sets the "contract_code_url" field.
func (puo *PledgeUpdateOne) SetContractCodeURL(s string) *PledgeUpdateOne {
	puo.mutation.SetContractCodeURL(s)
	return puo
}

// SetNillableContractCodeURL sets the "contract_code_url" field if the given value is not nil.
func (puo *PledgeUpdateOne) SetNillableContractCodeURL(s *string) *PledgeUpdateOne {
	if s != nil {
		puo.SetContractCodeURL(*s)
	}
	return puo
}

// ClearContractCodeURL clears the value of the "contract_code_url" field.
func (puo *PledgeUpdateOne) ClearContractCodeURL() *PledgeUpdateOne {
	puo.mutation.ClearContractCodeURL()
	return puo
}

// SetContractCodeBranch sets the "contract_code_branch" field.
func (puo *PledgeUpdateOne) SetContractCodeBranch(s string) *PledgeUpdateOne {
	puo.mutation.SetContractCodeBranch(s)
	return puo
}

// SetNillableContractCodeBranch sets the "contract_code_branch" field if the given value is not nil.
func (puo *PledgeUpdateOne) SetNillableContractCodeBranch(s *string) *PledgeUpdateOne {
	if s != nil {
		puo.SetContractCodeBranch(*s)
	}
	return puo
}

// ClearContractCodeBranch clears the value of the "contract_code_branch" field.
func (puo *PledgeUpdateOne) ClearContractCodeBranch() *PledgeUpdateOne {
	puo.mutation.ClearContractCodeBranch()
	return puo
}

// SetContractState sets the "contract_state" field.
func (puo *PledgeUpdateOne) SetContractState(s string) *PledgeUpdateOne {
	puo.mutation.SetContractState(s)
	return puo
}

// SetNillableContractState sets the "contract_state" field if the given value is not nil.
func (puo *PledgeUpdateOne) SetNillableContractState(s *string) *PledgeUpdateOne {
	if s != nil {
		puo.SetContractState(*s)
	}
	return puo
}

// ClearContractState clears the value of the "contract_state" field.
func (puo *PledgeUpdateOne) ClearContractState() *PledgeUpdateOne {
	puo.mutation.ClearContractState()
	return puo
}

// Mutation returns the PledgeMutation object of the builder.
func (puo *PledgeUpdateOne) Mutation() *PledgeMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PledgeUpdateOne) Select(field string, fields ...string) *PledgeUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pledge entity.
func (puo *PledgeUpdateOne) Save(ctx context.Context) (*Pledge, error) {
	var (
		err  error
		node *Pledge
	)
	if err := puo.defaults(); err != nil {
		return nil, err
	}
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PledgeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Pledge)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PledgeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PledgeUpdateOne) SaveX(ctx context.Context) *Pledge {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PledgeUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PledgeUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PledgeUpdateOne) defaults() error {
	if _, ok := puo.mutation.UpdatedAt(); !ok {
		if pledge.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized pledge.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := pledge.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PledgeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PledgeUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PledgeUpdateOne) sqlSave(ctx context.Context) (_node *Pledge, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pledge.Table,
			Columns: pledge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pledge.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pledge.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pledge.FieldID)
		for _, f := range fields {
			if !pledge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pledge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldCreatedAt,
		})
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldUpdatedAt,
		})
	}
	if value, ok := puo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldDeletedAt,
		})
	}
	if value, ok := puo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: pledge.FieldDeletedAt,
		})
	}
	if value, ok := puo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pledge.FieldEntID,
		})
	}
	if value, ok := puo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: pledge.FieldGoodID,
		})
	}
	if puo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: pledge.FieldGoodID,
		})
	}
	if value, ok := puo.mutation.ContractCodeURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pledge.FieldContractCodeURL,
		})
	}
	if puo.mutation.ContractCodeURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pledge.FieldContractCodeURL,
		})
	}
	if value, ok := puo.mutation.ContractCodeBranch(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pledge.FieldContractCodeBranch,
		})
	}
	if puo.mutation.ContractCodeBranchCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pledge.FieldContractCodeBranch,
		})
	}
	if value, ok := puo.mutation.ContractState(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pledge.FieldContractState,
		})
	}
	if puo.mutation.ContractStateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: pledge.FieldContractState,
		})
	}
	_spec.Modifiers = puo.modifiers
	_node = &Pledge{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pledge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
