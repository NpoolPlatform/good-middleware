// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appdefaultgood"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppDefaultGoodUpdate is the builder for updating AppDefaultGood entities.
type AppDefaultGoodUpdate struct {
	config
	hooks     []Hook
	mutation  *AppDefaultGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the AppDefaultGoodUpdate builder.
func (adgu *AppDefaultGoodUpdate) Where(ps ...predicate.AppDefaultGood) *AppDefaultGoodUpdate {
	adgu.mutation.Where(ps...)
	return adgu
}

// SetCreatedAt sets the "created_at" field.
func (adgu *AppDefaultGoodUpdate) SetCreatedAt(u uint32) *AppDefaultGoodUpdate {
	adgu.mutation.ResetCreatedAt()
	adgu.mutation.SetCreatedAt(u)
	return adgu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (adgu *AppDefaultGoodUpdate) SetNillableCreatedAt(u *uint32) *AppDefaultGoodUpdate {
	if u != nil {
		adgu.SetCreatedAt(*u)
	}
	return adgu
}

// AddCreatedAt adds u to the "created_at" field.
func (adgu *AppDefaultGoodUpdate) AddCreatedAt(u int32) *AppDefaultGoodUpdate {
	adgu.mutation.AddCreatedAt(u)
	return adgu
}

// SetUpdatedAt sets the "updated_at" field.
func (adgu *AppDefaultGoodUpdate) SetUpdatedAt(u uint32) *AppDefaultGoodUpdate {
	adgu.mutation.ResetUpdatedAt()
	adgu.mutation.SetUpdatedAt(u)
	return adgu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (adgu *AppDefaultGoodUpdate) AddUpdatedAt(u int32) *AppDefaultGoodUpdate {
	adgu.mutation.AddUpdatedAt(u)
	return adgu
}

// SetDeletedAt sets the "deleted_at" field.
func (adgu *AppDefaultGoodUpdate) SetDeletedAt(u uint32) *AppDefaultGoodUpdate {
	adgu.mutation.ResetDeletedAt()
	adgu.mutation.SetDeletedAt(u)
	return adgu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (adgu *AppDefaultGoodUpdate) SetNillableDeletedAt(u *uint32) *AppDefaultGoodUpdate {
	if u != nil {
		adgu.SetDeletedAt(*u)
	}
	return adgu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (adgu *AppDefaultGoodUpdate) AddDeletedAt(u int32) *AppDefaultGoodUpdate {
	adgu.mutation.AddDeletedAt(u)
	return adgu
}

// SetEntID sets the "ent_id" field.
func (adgu *AppDefaultGoodUpdate) SetEntID(u uuid.UUID) *AppDefaultGoodUpdate {
	adgu.mutation.SetEntID(u)
	return adgu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (adgu *AppDefaultGoodUpdate) SetNillableEntID(u *uuid.UUID) *AppDefaultGoodUpdate {
	if u != nil {
		adgu.SetEntID(*u)
	}
	return adgu
}

// SetAppGoodID sets the "app_good_id" field.
func (adgu *AppDefaultGoodUpdate) SetAppGoodID(u uuid.UUID) *AppDefaultGoodUpdate {
	adgu.mutation.SetAppGoodID(u)
	return adgu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (adgu *AppDefaultGoodUpdate) SetNillableAppGoodID(u *uuid.UUID) *AppDefaultGoodUpdate {
	if u != nil {
		adgu.SetAppGoodID(*u)
	}
	return adgu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (adgu *AppDefaultGoodUpdate) ClearAppGoodID() *AppDefaultGoodUpdate {
	adgu.mutation.ClearAppGoodID()
	return adgu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (adgu *AppDefaultGoodUpdate) SetCoinTypeID(u uuid.UUID) *AppDefaultGoodUpdate {
	adgu.mutation.SetCoinTypeID(u)
	return adgu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (adgu *AppDefaultGoodUpdate) SetNillableCoinTypeID(u *uuid.UUID) *AppDefaultGoodUpdate {
	if u != nil {
		adgu.SetCoinTypeID(*u)
	}
	return adgu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (adgu *AppDefaultGoodUpdate) ClearCoinTypeID() *AppDefaultGoodUpdate {
	adgu.mutation.ClearCoinTypeID()
	return adgu
}

// Mutation returns the AppDefaultGoodMutation object of the builder.
func (adgu *AppDefaultGoodUpdate) Mutation() *AppDefaultGoodMutation {
	return adgu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (adgu *AppDefaultGoodUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := adgu.defaults(); err != nil {
		return 0, err
	}
	if len(adgu.hooks) == 0 {
		affected, err = adgu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppDefaultGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			adgu.mutation = mutation
			affected, err = adgu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(adgu.hooks) - 1; i >= 0; i-- {
			if adgu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = adgu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, adgu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (adgu *AppDefaultGoodUpdate) SaveX(ctx context.Context) int {
	affected, err := adgu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (adgu *AppDefaultGoodUpdate) Exec(ctx context.Context) error {
	_, err := adgu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adgu *AppDefaultGoodUpdate) ExecX(ctx context.Context) {
	if err := adgu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (adgu *AppDefaultGoodUpdate) defaults() error {
	if _, ok := adgu.mutation.UpdatedAt(); !ok {
		if appdefaultgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appdefaultgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appdefaultgood.UpdateDefaultUpdatedAt()
		adgu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (adgu *AppDefaultGoodUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppDefaultGoodUpdate {
	adgu.modifiers = append(adgu.modifiers, modifiers...)
	return adgu
}

func (adgu *AppDefaultGoodUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appdefaultgood.Table,
			Columns: appdefaultgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appdefaultgood.FieldID,
			},
		},
	}
	if ps := adgu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := adgu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldCreatedAt,
		})
	}
	if value, ok := adgu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldCreatedAt,
		})
	}
	if value, ok := adgu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldUpdatedAt,
		})
	}
	if value, ok := adgu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldUpdatedAt,
		})
	}
	if value, ok := adgu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldDeletedAt,
		})
	}
	if value, ok := adgu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldDeletedAt,
		})
	}
	if value, ok := adgu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldEntID,
		})
	}
	if value, ok := adgu.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldAppGoodID,
		})
	}
	if adgu.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appdefaultgood.FieldAppGoodID,
		})
	}
	if value, ok := adgu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldCoinTypeID,
		})
	}
	if adgu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appdefaultgood.FieldCoinTypeID,
		})
	}
	_spec.Modifiers = adgu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, adgu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appdefaultgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// AppDefaultGoodUpdateOne is the builder for updating a single AppDefaultGood entity.
type AppDefaultGoodUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *AppDefaultGoodMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (adguo *AppDefaultGoodUpdateOne) SetCreatedAt(u uint32) *AppDefaultGoodUpdateOne {
	adguo.mutation.ResetCreatedAt()
	adguo.mutation.SetCreatedAt(u)
	return adguo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (adguo *AppDefaultGoodUpdateOne) SetNillableCreatedAt(u *uint32) *AppDefaultGoodUpdateOne {
	if u != nil {
		adguo.SetCreatedAt(*u)
	}
	return adguo
}

// AddCreatedAt adds u to the "created_at" field.
func (adguo *AppDefaultGoodUpdateOne) AddCreatedAt(u int32) *AppDefaultGoodUpdateOne {
	adguo.mutation.AddCreatedAt(u)
	return adguo
}

// SetUpdatedAt sets the "updated_at" field.
func (adguo *AppDefaultGoodUpdateOne) SetUpdatedAt(u uint32) *AppDefaultGoodUpdateOne {
	adguo.mutation.ResetUpdatedAt()
	adguo.mutation.SetUpdatedAt(u)
	return adguo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (adguo *AppDefaultGoodUpdateOne) AddUpdatedAt(u int32) *AppDefaultGoodUpdateOne {
	adguo.mutation.AddUpdatedAt(u)
	return adguo
}

// SetDeletedAt sets the "deleted_at" field.
func (adguo *AppDefaultGoodUpdateOne) SetDeletedAt(u uint32) *AppDefaultGoodUpdateOne {
	adguo.mutation.ResetDeletedAt()
	adguo.mutation.SetDeletedAt(u)
	return adguo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (adguo *AppDefaultGoodUpdateOne) SetNillableDeletedAt(u *uint32) *AppDefaultGoodUpdateOne {
	if u != nil {
		adguo.SetDeletedAt(*u)
	}
	return adguo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (adguo *AppDefaultGoodUpdateOne) AddDeletedAt(u int32) *AppDefaultGoodUpdateOne {
	adguo.mutation.AddDeletedAt(u)
	return adguo
}

// SetEntID sets the "ent_id" field.
func (adguo *AppDefaultGoodUpdateOne) SetEntID(u uuid.UUID) *AppDefaultGoodUpdateOne {
	adguo.mutation.SetEntID(u)
	return adguo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (adguo *AppDefaultGoodUpdateOne) SetNillableEntID(u *uuid.UUID) *AppDefaultGoodUpdateOne {
	if u != nil {
		adguo.SetEntID(*u)
	}
	return adguo
}

// SetAppGoodID sets the "app_good_id" field.
func (adguo *AppDefaultGoodUpdateOne) SetAppGoodID(u uuid.UUID) *AppDefaultGoodUpdateOne {
	adguo.mutation.SetAppGoodID(u)
	return adguo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (adguo *AppDefaultGoodUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *AppDefaultGoodUpdateOne {
	if u != nil {
		adguo.SetAppGoodID(*u)
	}
	return adguo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (adguo *AppDefaultGoodUpdateOne) ClearAppGoodID() *AppDefaultGoodUpdateOne {
	adguo.mutation.ClearAppGoodID()
	return adguo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (adguo *AppDefaultGoodUpdateOne) SetCoinTypeID(u uuid.UUID) *AppDefaultGoodUpdateOne {
	adguo.mutation.SetCoinTypeID(u)
	return adguo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (adguo *AppDefaultGoodUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *AppDefaultGoodUpdateOne {
	if u != nil {
		adguo.SetCoinTypeID(*u)
	}
	return adguo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (adguo *AppDefaultGoodUpdateOne) ClearCoinTypeID() *AppDefaultGoodUpdateOne {
	adguo.mutation.ClearCoinTypeID()
	return adguo
}

// Mutation returns the AppDefaultGoodMutation object of the builder.
func (adguo *AppDefaultGoodUpdateOne) Mutation() *AppDefaultGoodMutation {
	return adguo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (adguo *AppDefaultGoodUpdateOne) Select(field string, fields ...string) *AppDefaultGoodUpdateOne {
	adguo.fields = append([]string{field}, fields...)
	return adguo
}

// Save executes the query and returns the updated AppDefaultGood entity.
func (adguo *AppDefaultGoodUpdateOne) Save(ctx context.Context) (*AppDefaultGood, error) {
	var (
		err  error
		node *AppDefaultGood
	)
	if err := adguo.defaults(); err != nil {
		return nil, err
	}
	if len(adguo.hooks) == 0 {
		node, err = adguo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppDefaultGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			adguo.mutation = mutation
			node, err = adguo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(adguo.hooks) - 1; i >= 0; i-- {
			if adguo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = adguo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, adguo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppDefaultGood)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppDefaultGoodMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (adguo *AppDefaultGoodUpdateOne) SaveX(ctx context.Context) *AppDefaultGood {
	node, err := adguo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (adguo *AppDefaultGoodUpdateOne) Exec(ctx context.Context) error {
	_, err := adguo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adguo *AppDefaultGoodUpdateOne) ExecX(ctx context.Context) {
	if err := adguo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (adguo *AppDefaultGoodUpdateOne) defaults() error {
	if _, ok := adguo.mutation.UpdatedAt(); !ok {
		if appdefaultgood.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appdefaultgood.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appdefaultgood.UpdateDefaultUpdatedAt()
		adguo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (adguo *AppDefaultGoodUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *AppDefaultGoodUpdateOne {
	adguo.modifiers = append(adguo.modifiers, modifiers...)
	return adguo
}

func (adguo *AppDefaultGoodUpdateOne) sqlSave(ctx context.Context) (_node *AppDefaultGood, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appdefaultgood.Table,
			Columns: appdefaultgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appdefaultgood.FieldID,
			},
		},
	}
	id, ok := adguo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "AppDefaultGood.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := adguo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appdefaultgood.FieldID)
		for _, f := range fields {
			if !appdefaultgood.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != appdefaultgood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := adguo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := adguo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldCreatedAt,
		})
	}
	if value, ok := adguo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldCreatedAt,
		})
	}
	if value, ok := adguo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldUpdatedAt,
		})
	}
	if value, ok := adguo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldUpdatedAt,
		})
	}
	if value, ok := adguo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldDeletedAt,
		})
	}
	if value, ok := adguo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldDeletedAt,
		})
	}
	if value, ok := adguo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldEntID,
		})
	}
	if value, ok := adguo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldAppGoodID,
		})
	}
	if adguo.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appdefaultgood.FieldAppGoodID,
		})
	}
	if value, ok := adguo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldCoinTypeID,
		})
	}
	if adguo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: appdefaultgood.FieldCoinTypeID,
		})
	}
	_spec.Modifiers = adguo.modifiers
	_node = &AppDefaultGood{config: adguo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, adguo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{appdefaultgood.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
