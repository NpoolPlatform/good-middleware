// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// GoodCoinUpdate is the builder for updating GoodCoin entities.
type GoodCoinUpdate struct {
	config
	hooks     []Hook
	mutation  *GoodCoinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GoodCoinUpdate builder.
func (gcu *GoodCoinUpdate) Where(ps ...predicate.GoodCoin) *GoodCoinUpdate {
	gcu.mutation.Where(ps...)
	return gcu
}

// SetCreatedAt sets the "created_at" field.
func (gcu *GoodCoinUpdate) SetCreatedAt(u uint32) *GoodCoinUpdate {
	gcu.mutation.ResetCreatedAt()
	gcu.mutation.SetCreatedAt(u)
	return gcu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gcu *GoodCoinUpdate) SetNillableCreatedAt(u *uint32) *GoodCoinUpdate {
	if u != nil {
		gcu.SetCreatedAt(*u)
	}
	return gcu
}

// AddCreatedAt adds u to the "created_at" field.
func (gcu *GoodCoinUpdate) AddCreatedAt(u int32) *GoodCoinUpdate {
	gcu.mutation.AddCreatedAt(u)
	return gcu
}

// SetUpdatedAt sets the "updated_at" field.
func (gcu *GoodCoinUpdate) SetUpdatedAt(u uint32) *GoodCoinUpdate {
	gcu.mutation.ResetUpdatedAt()
	gcu.mutation.SetUpdatedAt(u)
	return gcu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gcu *GoodCoinUpdate) AddUpdatedAt(u int32) *GoodCoinUpdate {
	gcu.mutation.AddUpdatedAt(u)
	return gcu
}

// SetDeletedAt sets the "deleted_at" field.
func (gcu *GoodCoinUpdate) SetDeletedAt(u uint32) *GoodCoinUpdate {
	gcu.mutation.ResetDeletedAt()
	gcu.mutation.SetDeletedAt(u)
	return gcu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcu *GoodCoinUpdate) SetNillableDeletedAt(u *uint32) *GoodCoinUpdate {
	if u != nil {
		gcu.SetDeletedAt(*u)
	}
	return gcu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gcu *GoodCoinUpdate) AddDeletedAt(u int32) *GoodCoinUpdate {
	gcu.mutation.AddDeletedAt(u)
	return gcu
}

// SetEntID sets the "ent_id" field.
func (gcu *GoodCoinUpdate) SetEntID(u uuid.UUID) *GoodCoinUpdate {
	gcu.mutation.SetEntID(u)
	return gcu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gcu *GoodCoinUpdate) SetNillableEntID(u *uuid.UUID) *GoodCoinUpdate {
	if u != nil {
		gcu.SetEntID(*u)
	}
	return gcu
}

// SetGoodID sets the "good_id" field.
func (gcu *GoodCoinUpdate) SetGoodID(u uuid.UUID) *GoodCoinUpdate {
	gcu.mutation.SetGoodID(u)
	return gcu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gcu *GoodCoinUpdate) SetNillableGoodID(u *uuid.UUID) *GoodCoinUpdate {
	if u != nil {
		gcu.SetGoodID(*u)
	}
	return gcu
}

// ClearGoodID clears the value of the "good_id" field.
func (gcu *GoodCoinUpdate) ClearGoodID() *GoodCoinUpdate {
	gcu.mutation.ClearGoodID()
	return gcu
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gcu *GoodCoinUpdate) SetCoinTypeID(u uuid.UUID) *GoodCoinUpdate {
	gcu.mutation.SetCoinTypeID(u)
	return gcu
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gcu *GoodCoinUpdate) SetNillableCoinTypeID(u *uuid.UUID) *GoodCoinUpdate {
	if u != nil {
		gcu.SetCoinTypeID(*u)
	}
	return gcu
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gcu *GoodCoinUpdate) ClearCoinTypeID() *GoodCoinUpdate {
	gcu.mutation.ClearCoinTypeID()
	return gcu
}

// SetMain sets the "main" field.
func (gcu *GoodCoinUpdate) SetMain(b bool) *GoodCoinUpdate {
	gcu.mutation.SetMain(b)
	return gcu
}

// SetNillableMain sets the "main" field if the given value is not nil.
func (gcu *GoodCoinUpdate) SetNillableMain(b *bool) *GoodCoinUpdate {
	if b != nil {
		gcu.SetMain(*b)
	}
	return gcu
}

// ClearMain clears the value of the "main" field.
func (gcu *GoodCoinUpdate) ClearMain() *GoodCoinUpdate {
	gcu.mutation.ClearMain()
	return gcu
}

// SetIndex sets the "index" field.
func (gcu *GoodCoinUpdate) SetIndex(i int32) *GoodCoinUpdate {
	gcu.mutation.ResetIndex()
	gcu.mutation.SetIndex(i)
	return gcu
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (gcu *GoodCoinUpdate) SetNillableIndex(i *int32) *GoodCoinUpdate {
	if i != nil {
		gcu.SetIndex(*i)
	}
	return gcu
}

// AddIndex adds i to the "index" field.
func (gcu *GoodCoinUpdate) AddIndex(i int32) *GoodCoinUpdate {
	gcu.mutation.AddIndex(i)
	return gcu
}

// ClearIndex clears the value of the "index" field.
func (gcu *GoodCoinUpdate) ClearIndex() *GoodCoinUpdate {
	gcu.mutation.ClearIndex()
	return gcu
}

// Mutation returns the GoodCoinMutation object of the builder.
func (gcu *GoodCoinUpdate) Mutation() *GoodCoinMutation {
	return gcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gcu *GoodCoinUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := gcu.defaults(); err != nil {
		return 0, err
	}
	if len(gcu.hooks) == 0 {
		affected, err = gcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodCoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcu.mutation = mutation
			affected, err = gcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gcu.hooks) - 1; i >= 0; i-- {
			if gcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (gcu *GoodCoinUpdate) SaveX(ctx context.Context) int {
	affected, err := gcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gcu *GoodCoinUpdate) Exec(ctx context.Context) error {
	_, err := gcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcu *GoodCoinUpdate) ExecX(ctx context.Context) {
	if err := gcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcu *GoodCoinUpdate) defaults() error {
	if _, ok := gcu.mutation.UpdatedAt(); !ok {
		if goodcoin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodcoin.UpdateDefaultUpdatedAt()
		gcu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gcu *GoodCoinUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodCoinUpdate {
	gcu.modifiers = append(gcu.modifiers, modifiers...)
	return gcu
}

func (gcu *GoodCoinUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodcoin.Table,
			Columns: goodcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodcoin.FieldID,
			},
		},
	}
	if ps := gcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldCreatedAt,
		})
	}
	if value, ok := gcu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldCreatedAt,
		})
	}
	if value, ok := gcu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldUpdatedAt,
		})
	}
	if value, ok := gcu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldUpdatedAt,
		})
	}
	if value, ok := gcu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldDeletedAt,
		})
	}
	if value, ok := gcu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldDeletedAt,
		})
	}
	if value, ok := gcu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldEntID,
		})
	}
	if value, ok := gcu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldGoodID,
		})
	}
	if gcu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoin.FieldGoodID,
		})
	}
	if value, ok := gcu.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldCoinTypeID,
		})
	}
	if gcu.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoin.FieldCoinTypeID,
		})
	}
	if value, ok := gcu.mutation.Main(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodcoin.FieldMain,
		})
	}
	if gcu.mutation.MainCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: goodcoin.FieldMain,
		})
	}
	if value, ok := gcu.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: goodcoin.FieldIndex,
		})
	}
	if value, ok := gcu.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: goodcoin.FieldIndex,
		})
	}
	if gcu.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: goodcoin.FieldIndex,
		})
	}
	_spec.Modifiers = gcu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, gcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodcoin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// GoodCoinUpdateOne is the builder for updating a single GoodCoin entity.
type GoodCoinUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GoodCoinMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (gcuo *GoodCoinUpdateOne) SetCreatedAt(u uint32) *GoodCoinUpdateOne {
	gcuo.mutation.ResetCreatedAt()
	gcuo.mutation.SetCreatedAt(u)
	return gcuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gcuo *GoodCoinUpdateOne) SetNillableCreatedAt(u *uint32) *GoodCoinUpdateOne {
	if u != nil {
		gcuo.SetCreatedAt(*u)
	}
	return gcuo
}

// AddCreatedAt adds u to the "created_at" field.
func (gcuo *GoodCoinUpdateOne) AddCreatedAt(u int32) *GoodCoinUpdateOne {
	gcuo.mutation.AddCreatedAt(u)
	return gcuo
}

// SetUpdatedAt sets the "updated_at" field.
func (gcuo *GoodCoinUpdateOne) SetUpdatedAt(u uint32) *GoodCoinUpdateOne {
	gcuo.mutation.ResetUpdatedAt()
	gcuo.mutation.SetUpdatedAt(u)
	return gcuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (gcuo *GoodCoinUpdateOne) AddUpdatedAt(u int32) *GoodCoinUpdateOne {
	gcuo.mutation.AddUpdatedAt(u)
	return gcuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gcuo *GoodCoinUpdateOne) SetDeletedAt(u uint32) *GoodCoinUpdateOne {
	gcuo.mutation.ResetDeletedAt()
	gcuo.mutation.SetDeletedAt(u)
	return gcuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcuo *GoodCoinUpdateOne) SetNillableDeletedAt(u *uint32) *GoodCoinUpdateOne {
	if u != nil {
		gcuo.SetDeletedAt(*u)
	}
	return gcuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (gcuo *GoodCoinUpdateOne) AddDeletedAt(u int32) *GoodCoinUpdateOne {
	gcuo.mutation.AddDeletedAt(u)
	return gcuo
}

// SetEntID sets the "ent_id" field.
func (gcuo *GoodCoinUpdateOne) SetEntID(u uuid.UUID) *GoodCoinUpdateOne {
	gcuo.mutation.SetEntID(u)
	return gcuo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gcuo *GoodCoinUpdateOne) SetNillableEntID(u *uuid.UUID) *GoodCoinUpdateOne {
	if u != nil {
		gcuo.SetEntID(*u)
	}
	return gcuo
}

// SetGoodID sets the "good_id" field.
func (gcuo *GoodCoinUpdateOne) SetGoodID(u uuid.UUID) *GoodCoinUpdateOne {
	gcuo.mutation.SetGoodID(u)
	return gcuo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gcuo *GoodCoinUpdateOne) SetNillableGoodID(u *uuid.UUID) *GoodCoinUpdateOne {
	if u != nil {
		gcuo.SetGoodID(*u)
	}
	return gcuo
}

// ClearGoodID clears the value of the "good_id" field.
func (gcuo *GoodCoinUpdateOne) ClearGoodID() *GoodCoinUpdateOne {
	gcuo.mutation.ClearGoodID()
	return gcuo
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gcuo *GoodCoinUpdateOne) SetCoinTypeID(u uuid.UUID) *GoodCoinUpdateOne {
	gcuo.mutation.SetCoinTypeID(u)
	return gcuo
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gcuo *GoodCoinUpdateOne) SetNillableCoinTypeID(u *uuid.UUID) *GoodCoinUpdateOne {
	if u != nil {
		gcuo.SetCoinTypeID(*u)
	}
	return gcuo
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (gcuo *GoodCoinUpdateOne) ClearCoinTypeID() *GoodCoinUpdateOne {
	gcuo.mutation.ClearCoinTypeID()
	return gcuo
}

// SetMain sets the "main" field.
func (gcuo *GoodCoinUpdateOne) SetMain(b bool) *GoodCoinUpdateOne {
	gcuo.mutation.SetMain(b)
	return gcuo
}

// SetNillableMain sets the "main" field if the given value is not nil.
func (gcuo *GoodCoinUpdateOne) SetNillableMain(b *bool) *GoodCoinUpdateOne {
	if b != nil {
		gcuo.SetMain(*b)
	}
	return gcuo
}

// ClearMain clears the value of the "main" field.
func (gcuo *GoodCoinUpdateOne) ClearMain() *GoodCoinUpdateOne {
	gcuo.mutation.ClearMain()
	return gcuo
}

// SetIndex sets the "index" field.
func (gcuo *GoodCoinUpdateOne) SetIndex(i int32) *GoodCoinUpdateOne {
	gcuo.mutation.ResetIndex()
	gcuo.mutation.SetIndex(i)
	return gcuo
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (gcuo *GoodCoinUpdateOne) SetNillableIndex(i *int32) *GoodCoinUpdateOne {
	if i != nil {
		gcuo.SetIndex(*i)
	}
	return gcuo
}

// AddIndex adds i to the "index" field.
func (gcuo *GoodCoinUpdateOne) AddIndex(i int32) *GoodCoinUpdateOne {
	gcuo.mutation.AddIndex(i)
	return gcuo
}

// ClearIndex clears the value of the "index" field.
func (gcuo *GoodCoinUpdateOne) ClearIndex() *GoodCoinUpdateOne {
	gcuo.mutation.ClearIndex()
	return gcuo
}

// Mutation returns the GoodCoinMutation object of the builder.
func (gcuo *GoodCoinUpdateOne) Mutation() *GoodCoinMutation {
	return gcuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gcuo *GoodCoinUpdateOne) Select(field string, fields ...string) *GoodCoinUpdateOne {
	gcuo.fields = append([]string{field}, fields...)
	return gcuo
}

// Save executes the query and returns the updated GoodCoin entity.
func (gcuo *GoodCoinUpdateOne) Save(ctx context.Context) (*GoodCoin, error) {
	var (
		err  error
		node *GoodCoin
	)
	if err := gcuo.defaults(); err != nil {
		return nil, err
	}
	if len(gcuo.hooks) == 0 {
		node, err = gcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodCoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gcuo.mutation = mutation
			node, err = gcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(gcuo.hooks) - 1; i >= 0; i-- {
			if gcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gcuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoodCoin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoodCoinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (gcuo *GoodCoinUpdateOne) SaveX(ctx context.Context) *GoodCoin {
	node, err := gcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gcuo *GoodCoinUpdateOne) Exec(ctx context.Context) error {
	_, err := gcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcuo *GoodCoinUpdateOne) ExecX(ctx context.Context) {
	if err := gcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcuo *GoodCoinUpdateOne) defaults() error {
	if _, ok := gcuo.mutation.UpdatedAt(); !ok {
		if goodcoin.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodcoin.UpdateDefaultUpdatedAt()
		gcuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (gcuo *GoodCoinUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GoodCoinUpdateOne {
	gcuo.modifiers = append(gcuo.modifiers, modifiers...)
	return gcuo
}

func (gcuo *GoodCoinUpdateOne) sqlSave(ctx context.Context) (_node *GoodCoin, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goodcoin.Table,
			Columns: goodcoin.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodcoin.FieldID,
			},
		},
	}
	id, ok := gcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GoodCoin.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goodcoin.FieldID)
		for _, f := range fields {
			if !goodcoin.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != goodcoin.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gcuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldCreatedAt,
		})
	}
	if value, ok := gcuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldCreatedAt,
		})
	}
	if value, ok := gcuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldUpdatedAt,
		})
	}
	if value, ok := gcuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldUpdatedAt,
		})
	}
	if value, ok := gcuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldDeletedAt,
		})
	}
	if value, ok := gcuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldDeletedAt,
		})
	}
	if value, ok := gcuo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldEntID,
		})
	}
	if value, ok := gcuo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldGoodID,
		})
	}
	if gcuo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoin.FieldGoodID,
		})
	}
	if value, ok := gcuo.mutation.CoinTypeID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldCoinTypeID,
		})
	}
	if gcuo.mutation.CoinTypeIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: goodcoin.FieldCoinTypeID,
		})
	}
	if value, ok := gcuo.mutation.Main(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodcoin.FieldMain,
		})
	}
	if gcuo.mutation.MainCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: goodcoin.FieldMain,
		})
	}
	if value, ok := gcuo.mutation.Index(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: goodcoin.FieldIndex,
		})
	}
	if value, ok := gcuo.mutation.AddedIndex(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: goodcoin.FieldIndex,
		})
	}
	if gcuo.mutation.IndexCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Column: goodcoin.FieldIndex,
		})
	}
	_spec.Modifiers = gcuo.modifiers
	_node = &GoodCoin{config: gcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{goodcoin.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
