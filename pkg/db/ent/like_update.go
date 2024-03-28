// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// LikeUpdate is the builder for updating Like entities.
type LikeUpdate struct {
	config
	hooks     []Hook
	mutation  *LikeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the LikeUpdate builder.
func (lu *LikeUpdate) Where(ps ...predicate.Like) *LikeUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetCreatedAt sets the "created_at" field.
func (lu *LikeUpdate) SetCreatedAt(u uint32) *LikeUpdate {
	lu.mutation.ResetCreatedAt()
	lu.mutation.SetCreatedAt(u)
	return lu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableCreatedAt(u *uint32) *LikeUpdate {
	if u != nil {
		lu.SetCreatedAt(*u)
	}
	return lu
}

// AddCreatedAt adds u to the "created_at" field.
func (lu *LikeUpdate) AddCreatedAt(u int32) *LikeUpdate {
	lu.mutation.AddCreatedAt(u)
	return lu
}

// SetUpdatedAt sets the "updated_at" field.
func (lu *LikeUpdate) SetUpdatedAt(u uint32) *LikeUpdate {
	lu.mutation.ResetUpdatedAt()
	lu.mutation.SetUpdatedAt(u)
	return lu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (lu *LikeUpdate) AddUpdatedAt(u int32) *LikeUpdate {
	lu.mutation.AddUpdatedAt(u)
	return lu
}

// SetDeletedAt sets the "deleted_at" field.
func (lu *LikeUpdate) SetDeletedAt(u uint32) *LikeUpdate {
	lu.mutation.ResetDeletedAt()
	lu.mutation.SetDeletedAt(u)
	return lu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableDeletedAt(u *uint32) *LikeUpdate {
	if u != nil {
		lu.SetDeletedAt(*u)
	}
	return lu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (lu *LikeUpdate) AddDeletedAt(u int32) *LikeUpdate {
	lu.mutation.AddDeletedAt(u)
	return lu
}

// SetEntID sets the "ent_id" field.
func (lu *LikeUpdate) SetEntID(u uuid.UUID) *LikeUpdate {
	lu.mutation.SetEntID(u)
	return lu
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableEntID(u *uuid.UUID) *LikeUpdate {
	if u != nil {
		lu.SetEntID(*u)
	}
	return lu
}

// SetAppID sets the "app_id" field.
func (lu *LikeUpdate) SetAppID(u uuid.UUID) *LikeUpdate {
	lu.mutation.SetAppID(u)
	return lu
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableAppID(u *uuid.UUID) *LikeUpdate {
	if u != nil {
		lu.SetAppID(*u)
	}
	return lu
}

// ClearAppID clears the value of the "app_id" field.
func (lu *LikeUpdate) ClearAppID() *LikeUpdate {
	lu.mutation.ClearAppID()
	return lu
}

// SetUserID sets the "user_id" field.
func (lu *LikeUpdate) SetUserID(u uuid.UUID) *LikeUpdate {
	lu.mutation.SetUserID(u)
	return lu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableUserID(u *uuid.UUID) *LikeUpdate {
	if u != nil {
		lu.SetUserID(*u)
	}
	return lu
}

// ClearUserID clears the value of the "user_id" field.
func (lu *LikeUpdate) ClearUserID() *LikeUpdate {
	lu.mutation.ClearUserID()
	return lu
}

// SetAppGoodID sets the "app_good_id" field.
func (lu *LikeUpdate) SetAppGoodID(u uuid.UUID) *LikeUpdate {
	lu.mutation.SetAppGoodID(u)
	return lu
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (lu *LikeUpdate) SetNillableAppGoodID(u *uuid.UUID) *LikeUpdate {
	if u != nil {
		lu.SetAppGoodID(*u)
	}
	return lu
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (lu *LikeUpdate) ClearAppGoodID() *LikeUpdate {
	lu.mutation.ClearAppGoodID()
	return lu
}

// SetLike sets the "like" field.
func (lu *LikeUpdate) SetLike(b bool) *LikeUpdate {
	lu.mutation.SetLike(b)
	return lu
}

// Mutation returns the LikeMutation object of the builder.
func (lu *LikeUpdate) Mutation() *LikeMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LikeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := lu.defaults(); err != nil {
		return 0, err
	}
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LikeUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LikeUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LikeUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lu *LikeUpdate) defaults() error {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		if like.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized like.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := like.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lu *LikeUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LikeUpdate {
	lu.modifiers = append(lu.modifiers, modifiers...)
	return lu
}

func (lu *LikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   like.Table,
			Columns: like.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: like.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldCreatedAt,
		})
	}
	if value, ok := lu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldCreatedAt,
		})
	}
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldUpdatedAt,
		})
	}
	if value, ok := lu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldUpdatedAt,
		})
	}
	if value, ok := lu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldDeletedAt,
		})
	}
	if value, ok := lu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldDeletedAt,
		})
	}
	if value, ok := lu.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldEntID,
		})
	}
	if value, ok := lu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldAppID,
		})
	}
	if lu.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: like.FieldAppID,
		})
	}
	if value, ok := lu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldUserID,
		})
	}
	if lu.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: like.FieldUserID,
		})
	}
	if value, ok := lu.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldAppGoodID,
		})
	}
	if lu.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: like.FieldAppGoodID,
		})
	}
	if value, ok := lu.mutation.Like(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: like.FieldLike,
		})
	}
	_spec.Modifiers = lu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// LikeUpdateOne is the builder for updating a single Like entity.
type LikeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *LikeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (luo *LikeUpdateOne) SetCreatedAt(u uint32) *LikeUpdateOne {
	luo.mutation.ResetCreatedAt()
	luo.mutation.SetCreatedAt(u)
	return luo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableCreatedAt(u *uint32) *LikeUpdateOne {
	if u != nil {
		luo.SetCreatedAt(*u)
	}
	return luo
}

// AddCreatedAt adds u to the "created_at" field.
func (luo *LikeUpdateOne) AddCreatedAt(u int32) *LikeUpdateOne {
	luo.mutation.AddCreatedAt(u)
	return luo
}

// SetUpdatedAt sets the "updated_at" field.
func (luo *LikeUpdateOne) SetUpdatedAt(u uint32) *LikeUpdateOne {
	luo.mutation.ResetUpdatedAt()
	luo.mutation.SetUpdatedAt(u)
	return luo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (luo *LikeUpdateOne) AddUpdatedAt(u int32) *LikeUpdateOne {
	luo.mutation.AddUpdatedAt(u)
	return luo
}

// SetDeletedAt sets the "deleted_at" field.
func (luo *LikeUpdateOne) SetDeletedAt(u uint32) *LikeUpdateOne {
	luo.mutation.ResetDeletedAt()
	luo.mutation.SetDeletedAt(u)
	return luo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableDeletedAt(u *uint32) *LikeUpdateOne {
	if u != nil {
		luo.SetDeletedAt(*u)
	}
	return luo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (luo *LikeUpdateOne) AddDeletedAt(u int32) *LikeUpdateOne {
	luo.mutation.AddDeletedAt(u)
	return luo
}

// SetEntID sets the "ent_id" field.
func (luo *LikeUpdateOne) SetEntID(u uuid.UUID) *LikeUpdateOne {
	luo.mutation.SetEntID(u)
	return luo
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableEntID(u *uuid.UUID) *LikeUpdateOne {
	if u != nil {
		luo.SetEntID(*u)
	}
	return luo
}

// SetAppID sets the "app_id" field.
func (luo *LikeUpdateOne) SetAppID(u uuid.UUID) *LikeUpdateOne {
	luo.mutation.SetAppID(u)
	return luo
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableAppID(u *uuid.UUID) *LikeUpdateOne {
	if u != nil {
		luo.SetAppID(*u)
	}
	return luo
}

// ClearAppID clears the value of the "app_id" field.
func (luo *LikeUpdateOne) ClearAppID() *LikeUpdateOne {
	luo.mutation.ClearAppID()
	return luo
}

// SetUserID sets the "user_id" field.
func (luo *LikeUpdateOne) SetUserID(u uuid.UUID) *LikeUpdateOne {
	luo.mutation.SetUserID(u)
	return luo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableUserID(u *uuid.UUID) *LikeUpdateOne {
	if u != nil {
		luo.SetUserID(*u)
	}
	return luo
}

// ClearUserID clears the value of the "user_id" field.
func (luo *LikeUpdateOne) ClearUserID() *LikeUpdateOne {
	luo.mutation.ClearUserID()
	return luo
}

// SetAppGoodID sets the "app_good_id" field.
func (luo *LikeUpdateOne) SetAppGoodID(u uuid.UUID) *LikeUpdateOne {
	luo.mutation.SetAppGoodID(u)
	return luo
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableAppGoodID(u *uuid.UUID) *LikeUpdateOne {
	if u != nil {
		luo.SetAppGoodID(*u)
	}
	return luo
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (luo *LikeUpdateOne) ClearAppGoodID() *LikeUpdateOne {
	luo.mutation.ClearAppGoodID()
	return luo
}

// SetLike sets the "like" field.
func (luo *LikeUpdateOne) SetLike(b bool) *LikeUpdateOne {
	luo.mutation.SetLike(b)
	return luo
}

// Mutation returns the LikeMutation object of the builder.
func (luo *LikeUpdateOne) Mutation() *LikeMutation {
	return luo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LikeUpdateOne) Select(field string, fields ...string) *LikeUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Like entity.
func (luo *LikeUpdateOne) Save(ctx context.Context) (*Like, error) {
	var (
		err  error
		node *Like
	)
	if err := luo.defaults(); err != nil {
		return nil, err
	}
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LikeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Like)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from LikeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LikeUpdateOne) SaveX(ctx context.Context) *Like {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LikeUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LikeUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (luo *LikeUpdateOne) defaults() error {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		if like.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized like.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := like.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (luo *LikeUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *LikeUpdateOne {
	luo.modifiers = append(luo.modifiers, modifiers...)
	return luo
}

func (luo *LikeUpdateOne) sqlSave(ctx context.Context) (_node *Like, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   like.Table,
			Columns: like.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: like.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Like.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, like.FieldID)
		for _, f := range fields {
			if !like.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != like.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldCreatedAt,
		})
	}
	if value, ok := luo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldCreatedAt,
		})
	}
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldUpdatedAt,
		})
	}
	if value, ok := luo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldUpdatedAt,
		})
	}
	if value, ok := luo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldDeletedAt,
		})
	}
	if value, ok := luo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: like.FieldDeletedAt,
		})
	}
	if value, ok := luo.mutation.EntID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldEntID,
		})
	}
	if value, ok := luo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldAppID,
		})
	}
	if luo.mutation.AppIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: like.FieldAppID,
		})
	}
	if value, ok := luo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldUserID,
		})
	}
	if luo.mutation.UserIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: like.FieldUserID,
		})
	}
	if value, ok := luo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: like.FieldAppGoodID,
		})
	}
	if luo.mutation.AppGoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: like.FieldAppGoodID,
		})
	}
	if value, ok := luo.mutation.Like(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: like.FieldLike,
		})
	}
	_spec.Modifiers = luo.modifiers
	_node = &Like{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
