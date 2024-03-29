// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	"github.com/google/uuid"
)

// VendorBrandUpdate is the builder for updating VendorBrand entities.
type VendorBrandUpdate struct {
	config
	hooks     []Hook
	mutation  *VendorBrandMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the VendorBrandUpdate builder.
func (vbu *VendorBrandUpdate) Where(ps ...predicate.VendorBrand) *VendorBrandUpdate {
	vbu.mutation.Where(ps...)
	return vbu
}

// SetCreatedAt sets the "created_at" field.
func (vbu *VendorBrandUpdate) SetCreatedAt(u uint32) *VendorBrandUpdate {
	vbu.mutation.ResetCreatedAt()
	vbu.mutation.SetCreatedAt(u)
	return vbu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vbu *VendorBrandUpdate) SetNillableCreatedAt(u *uint32) *VendorBrandUpdate {
	if u != nil {
		vbu.SetCreatedAt(*u)
	}
	return vbu
}

// AddCreatedAt adds u to the "created_at" field.
func (vbu *VendorBrandUpdate) AddCreatedAt(u int32) *VendorBrandUpdate {
	vbu.mutation.AddCreatedAt(u)
	return vbu
}

// SetUpdatedAt sets the "updated_at" field.
func (vbu *VendorBrandUpdate) SetUpdatedAt(u uint32) *VendorBrandUpdate {
	vbu.mutation.ResetUpdatedAt()
	vbu.mutation.SetUpdatedAt(u)
	return vbu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (vbu *VendorBrandUpdate) AddUpdatedAt(u int32) *VendorBrandUpdate {
	vbu.mutation.AddUpdatedAt(u)
	return vbu
}

// SetDeletedAt sets the "deleted_at" field.
func (vbu *VendorBrandUpdate) SetDeletedAt(u uint32) *VendorBrandUpdate {
	vbu.mutation.ResetDeletedAt()
	vbu.mutation.SetDeletedAt(u)
	return vbu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vbu *VendorBrandUpdate) SetNillableDeletedAt(u *uint32) *VendorBrandUpdate {
	if u != nil {
		vbu.SetDeletedAt(*u)
	}
	return vbu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (vbu *VendorBrandUpdate) AddDeletedAt(u int32) *VendorBrandUpdate {
	vbu.mutation.AddDeletedAt(u)
	return vbu
}

// SetEntID sets the "ent_id" field.
func (vbu *VendorBrandUpdate) SetEntID(u uuid.UUID) *VendorBrandUpdate {
	vbu.mutation.SetEntID(u)
	return vbu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (vbu *VendorBrandUpdate) SetNillableEntID(u *uuid.UUID) *VendorBrandUpdate {
	if u != nil {
		vbu.SetEntID(*u)
	}
	return vbu
}

// SetName sets the "name" field.
func (vbu *VendorBrandUpdate) SetName(s string) *VendorBrandUpdate {
	vbu.mutation.SetName(s)
	return vbu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vbu *VendorBrandUpdate) SetNillableName(s *string) *VendorBrandUpdate {
	if s != nil {
		vbu.SetName(*s)
	}
	return vbu
}

// ClearName clears the value of the "name" field.
func (vbu *VendorBrandUpdate) ClearName() *VendorBrandUpdate {
	vbu.mutation.ClearName()
	return vbu
}

// SetLogo sets the "logo" field.
func (vbu *VendorBrandUpdate) SetLogo(s string) *VendorBrandUpdate {
	vbu.mutation.SetLogo(s)
	return vbu
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (vbu *VendorBrandUpdate) SetNillableLogo(s *string) *VendorBrandUpdate {
	if s != nil {
		vbu.SetLogo(*s)
	}
	return vbu
}

// ClearLogo clears the value of the "logo" field.
func (vbu *VendorBrandUpdate) ClearLogo() *VendorBrandUpdate {
	vbu.mutation.ClearLogo()
	return vbu
}

// Mutation returns the VendorBrandMutation object of the builder.
func (vbu *VendorBrandUpdate) Mutation() *VendorBrandMutation {
	return vbu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vbu *VendorBrandUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := vbu.defaults(); err != nil {
		return 0, err
	}
	if len(vbu.hooks) == 0 {
		if err = vbu.check(); err != nil {
			return 0, err
		}
		affected, err = vbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VendorBrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vbu.check(); err != nil {
				return 0, err
			}
			vbu.mutation = mutation
			affected, err = vbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vbu.hooks) - 1; i >= 0; i-- {
			if vbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vbu *VendorBrandUpdate) SaveX(ctx context.Context) int {
	affected, err := vbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vbu *VendorBrandUpdate) Exec(ctx context.Context) error {
	_, err := vbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vbu *VendorBrandUpdate) ExecX(ctx context.Context) {
	if err := vbu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vbu *VendorBrandUpdate) defaults() error {
	if _, ok := vbu.mutation.UpdatedAt(); !ok {
		if vendorbrand.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized vendorbrand.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := vendorbrand.UpdateDefaultUpdatedAt()
		vbu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vbu *VendorBrandUpdate) check() error {
	if v, ok := vbu.mutation.Name(); ok {
		if err := vendorbrand.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "VendorBrand.name": %w`, err)}
		}
	}
	if v, ok := vbu.mutation.Logo(); ok {
		if err := vendorbrand.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "VendorBrand.logo": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vbu *VendorBrandUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VendorBrandUpdate {
	vbu.modifiers = append(vbu.modifiers, modifiers...)
	return vbu
}

func (vbu *VendorBrandUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vendorbrand.Table,
			Columns: vendorbrand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: vendorbrand.FieldID,
			},
		},
	}
	if ps := vbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vbu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldCreatedAt,
		})
	}
	if value, ok := vbu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldCreatedAt,
		})
	}
	if value, ok := vbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldUpdatedAt,
		})
	}
	if value, ok := vbu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldUpdatedAt,
		})
	}
	if value, ok := vbu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldDeletedAt,
		})
	}
	if value, ok := vbu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldDeletedAt,
		})
	}
	if value, ok := vbu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: vendorbrand.FieldEntID,
		})
	}
	if value, ok := vbu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorbrand.FieldName,
		})
	}
	if vbu.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vendorbrand.FieldName,
		})
	}
	if value, ok := vbu.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorbrand.FieldLogo,
		})
	}
	if vbu.mutation.LogoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vendorbrand.FieldLogo,
		})
	}
	_spec.Modifiers = vbu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, vbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vendorbrand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// VendorBrandUpdateOne is the builder for updating a single VendorBrand entity.
type VendorBrandUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *VendorBrandMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (vbuo *VendorBrandUpdateOne) SetCreatedAt(u uint32) *VendorBrandUpdateOne {
	vbuo.mutation.ResetCreatedAt()
	vbuo.mutation.SetCreatedAt(u)
	return vbuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (vbuo *VendorBrandUpdateOne) SetNillableCreatedAt(u *uint32) *VendorBrandUpdateOne {
	if u != nil {
		vbuo.SetCreatedAt(*u)
	}
	return vbuo
}

// AddCreatedAt adds u to the "created_at" field.
func (vbuo *VendorBrandUpdateOne) AddCreatedAt(u int32) *VendorBrandUpdateOne {
	vbuo.mutation.AddCreatedAt(u)
	return vbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (vbuo *VendorBrandUpdateOne) SetUpdatedAt(u uint32) *VendorBrandUpdateOne {
	vbuo.mutation.ResetUpdatedAt()
	vbuo.mutation.SetUpdatedAt(u)
	return vbuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (vbuo *VendorBrandUpdateOne) AddUpdatedAt(u int32) *VendorBrandUpdateOne {
	vbuo.mutation.AddUpdatedAt(u)
	return vbuo
}

// SetDeletedAt sets the "deleted_at" field.
func (vbuo *VendorBrandUpdateOne) SetDeletedAt(u uint32) *VendorBrandUpdateOne {
	vbuo.mutation.ResetDeletedAt()
	vbuo.mutation.SetDeletedAt(u)
	return vbuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (vbuo *VendorBrandUpdateOne) SetNillableDeletedAt(u *uint32) *VendorBrandUpdateOne {
	if u != nil {
		vbuo.SetDeletedAt(*u)
	}
	return vbuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (vbuo *VendorBrandUpdateOne) AddDeletedAt(u int32) *VendorBrandUpdateOne {
	vbuo.mutation.AddDeletedAt(u)
	return vbuo
}

// SetEntID sets the "ent_id" field.
func (vbuo *VendorBrandUpdateOne) SetEntID(u uuid.UUID) *VendorBrandUpdateOne {
	vbuo.mutation.SetEntID(u)
	return vbuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (vbuo *VendorBrandUpdateOne) SetNillableEntID(u *uuid.UUID) *VendorBrandUpdateOne {
	if u != nil {
		vbuo.SetEntID(*u)
	}
	return vbuo
}

// SetName sets the "name" field.
func (vbuo *VendorBrandUpdateOne) SetName(s string) *VendorBrandUpdateOne {
	vbuo.mutation.SetName(s)
	return vbuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (vbuo *VendorBrandUpdateOne) SetNillableName(s *string) *VendorBrandUpdateOne {
	if s != nil {
		vbuo.SetName(*s)
	}
	return vbuo
}

// ClearName clears the value of the "name" field.
func (vbuo *VendorBrandUpdateOne) ClearName() *VendorBrandUpdateOne {
	vbuo.mutation.ClearName()
	return vbuo
}

// SetLogo sets the "logo" field.
func (vbuo *VendorBrandUpdateOne) SetLogo(s string) *VendorBrandUpdateOne {
	vbuo.mutation.SetLogo(s)
	return vbuo
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (vbuo *VendorBrandUpdateOne) SetNillableLogo(s *string) *VendorBrandUpdateOne {
	if s != nil {
		vbuo.SetLogo(*s)
	}
	return vbuo
}

// ClearLogo clears the value of the "logo" field.
func (vbuo *VendorBrandUpdateOne) ClearLogo() *VendorBrandUpdateOne {
	vbuo.mutation.ClearLogo()
	return vbuo
}

// Mutation returns the VendorBrandMutation object of the builder.
func (vbuo *VendorBrandUpdateOne) Mutation() *VendorBrandMutation {
	return vbuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vbuo *VendorBrandUpdateOne) Select(field string, fields ...string) *VendorBrandUpdateOne {
	vbuo.fields = append([]string{field}, fields...)
	return vbuo
}

// Save executes the query and returns the updated VendorBrand entity.
func (vbuo *VendorBrandUpdateOne) Save(ctx context.Context) (*VendorBrand, error) {
	var (
		err  error
		node *VendorBrand
	)
	if err := vbuo.defaults(); err != nil {
		return nil, err
	}
	if len(vbuo.hooks) == 0 {
		if err = vbuo.check(); err != nil {
			return nil, err
		}
		node, err = vbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VendorBrandMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vbuo.check(); err != nil {
				return nil, err
			}
			vbuo.mutation = mutation
			node, err = vbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vbuo.hooks) - 1; i >= 0; i-- {
			if vbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vbuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, vbuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*VendorBrand)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from VendorBrandMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vbuo *VendorBrandUpdateOne) SaveX(ctx context.Context) *VendorBrand {
	node, err := vbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vbuo *VendorBrandUpdateOne) Exec(ctx context.Context) error {
	_, err := vbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vbuo *VendorBrandUpdateOne) ExecX(ctx context.Context) {
	if err := vbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (vbuo *VendorBrandUpdateOne) defaults() error {
	if _, ok := vbuo.mutation.UpdatedAt(); !ok {
		if vendorbrand.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized vendorbrand.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := vendorbrand.UpdateDefaultUpdatedAt()
		vbuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (vbuo *VendorBrandUpdateOne) check() error {
	if v, ok := vbuo.mutation.Name(); ok {
		if err := vendorbrand.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "VendorBrand.name": %w`, err)}
		}
	}
	if v, ok := vbuo.mutation.Logo(); ok {
		if err := vendorbrand.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "VendorBrand.logo": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (vbuo *VendorBrandUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *VendorBrandUpdateOne {
	vbuo.modifiers = append(vbuo.modifiers, modifiers...)
	return vbuo
}

func (vbuo *VendorBrandUpdateOne) sqlSave(ctx context.Context) (_node *VendorBrand, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vendorbrand.Table,
			Columns: vendorbrand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: vendorbrand.FieldID,
			},
		},
	}
	id, ok := vbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VendorBrand.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vendorbrand.FieldID)
		for _, f := range fields {
			if !vendorbrand.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vendorbrand.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vbuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldCreatedAt,
		})
	}
	if value, ok := vbuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldCreatedAt,
		})
	}
	if value, ok := vbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldUpdatedAt,
		})
	}
	if value, ok := vbuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldUpdatedAt,
		})
	}
	if value, ok := vbuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldDeletedAt,
		})
	}
	if value, ok := vbuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: vendorbrand.FieldDeletedAt,
		})
	}
	if value, ok := vbuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: vendorbrand.FieldEntID,
		})
	}
	if value, ok := vbuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorbrand.FieldName,
		})
	}
	if vbuo.mutation.NameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vendorbrand.FieldName,
		})
	}
	if value, ok := vbuo.mutation.Logo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: vendorbrand.FieldLogo,
		})
	}
	if vbuo.mutation.LogoCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: vendorbrand.FieldLogo,
		})
	}
	_spec.Modifiers = vbuo.modifiers
	_node = &VendorBrand{config: vbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vendorbrand.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
