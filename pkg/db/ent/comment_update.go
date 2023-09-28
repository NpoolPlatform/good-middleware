// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/comment"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks     []Hook
	mutation  *CommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetCreatedAt sets the "created_at" field.
func (cu *CommentUpdate) SetCreatedAt(u uint32) *CommentUpdate {
	cu.mutation.ResetCreatedAt()
	cu.mutation.SetCreatedAt(u)
	return cu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableCreatedAt(u *uint32) *CommentUpdate {
	if u != nil {
		cu.SetCreatedAt(*u)
	}
	return cu
}

// AddCreatedAt adds u to the "created_at" field.
func (cu *CommentUpdate) AddCreatedAt(u int32) *CommentUpdate {
	cu.mutation.AddCreatedAt(u)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(u uint32) *CommentUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(u)
	return cu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cu *CommentUpdate) AddUpdatedAt(u int32) *CommentUpdate {
	cu.mutation.AddUpdatedAt(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CommentUpdate) SetDeletedAt(u uint32) *CommentUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(u)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDeletedAt(u *uint32) *CommentUpdate {
	if u != nil {
		cu.SetDeletedAt(*u)
	}
	return cu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cu *CommentUpdate) AddDeletedAt(u int32) *CommentUpdate {
	cu.mutation.AddDeletedAt(u)
	return cu
}

// SetAppID sets the "app_id" field.
func (cu *CommentUpdate) SetAppID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetAppID(u)
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *CommentUpdate) SetUserID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetUserID(u)
	return cu
}

// SetGoodID sets the "good_id" field.
func (cu *CommentUpdate) SetGoodID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetGoodID(u)
	return cu
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableGoodID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetGoodID(*u)
	}
	return cu
}

// ClearGoodID clears the value of the "good_id" field.
func (cu *CommentUpdate) ClearGoodID() *CommentUpdate {
	cu.mutation.ClearGoodID()
	return cu
}

// SetAppGoodID sets the "app_good_id" field.
func (cu *CommentUpdate) SetAppGoodID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetAppGoodID(u)
	return cu
}

// SetOrderID sets the "order_id" field.
func (cu *CommentUpdate) SetOrderID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetOrderID(u)
	return cu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableOrderID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetOrderID(*u)
	}
	return cu
}

// ClearOrderID clears the value of the "order_id" field.
func (cu *CommentUpdate) ClearOrderID() *CommentUpdate {
	cu.mutation.ClearOrderID()
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableContent(s *string) *CommentUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// ClearContent clears the value of the "content" field.
func (cu *CommentUpdate) ClearContent() *CommentUpdate {
	cu.mutation.ClearContent()
	return cu
}

// SetReplyToID sets the "reply_to_id" field.
func (cu *CommentUpdate) SetReplyToID(u uuid.UUID) *CommentUpdate {
	cu.mutation.SetReplyToID(u)
	return cu
}

// SetNillableReplyToID sets the "reply_to_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableReplyToID(u *uuid.UUID) *CommentUpdate {
	if u != nil {
		cu.SetReplyToID(*u)
	}
	return cu
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (cu *CommentUpdate) ClearReplyToID() *CommentUpdate {
	cu.mutation.ClearReplyToID()
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cu *CommentUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommentUpdate {
	cu.modifiers = append(cu.modifiers, modifiers...)
	return cu
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldDeletedAt,
		})
	}
	if value, ok := cu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldAppID,
		})
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldUserID,
		})
	}
	if value, ok := cu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldGoodID,
		})
	}
	if cu.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: comment.FieldGoodID,
		})
	}
	if value, ok := cu.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldAppGoodID,
		})
	}
	if value, ok := cu.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldOrderID,
		})
	}
	if cu.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: comment.FieldOrderID,
		})
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if cu.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cu.mutation.ReplyToID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldReplyToID,
		})
	}
	if cu.mutation.ReplyToIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: comment.FieldReplyToID,
		})
	}
	_spec.Modifiers = cu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *CommentMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (cuo *CommentUpdateOne) SetCreatedAt(u uint32) *CommentUpdateOne {
	cuo.mutation.ResetCreatedAt()
	cuo.mutation.SetCreatedAt(u)
	return cuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableCreatedAt(u *uint32) *CommentUpdateOne {
	if u != nil {
		cuo.SetCreatedAt(*u)
	}
	return cuo
}

// AddCreatedAt adds u to the "created_at" field.
func (cuo *CommentUpdateOne) AddCreatedAt(u int32) *CommentUpdateOne {
	cuo.mutation.AddCreatedAt(u)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(u uint32) *CommentUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(u)
	return cuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (cuo *CommentUpdateOne) AddUpdatedAt(u int32) *CommentUpdateOne {
	cuo.mutation.AddUpdatedAt(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CommentUpdateOne) SetDeletedAt(u uint32) *CommentUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(u)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDeletedAt(u *uint32) *CommentUpdateOne {
	if u != nil {
		cuo.SetDeletedAt(*u)
	}
	return cuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (cuo *CommentUpdateOne) AddDeletedAt(u int32) *CommentUpdateOne {
	cuo.mutation.AddDeletedAt(u)
	return cuo
}

// SetAppID sets the "app_id" field.
func (cuo *CommentUpdateOne) SetAppID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetAppID(u)
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *CommentUpdateOne) SetUserID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetUserID(u)
	return cuo
}

// SetGoodID sets the "good_id" field.
func (cuo *CommentUpdateOne) SetGoodID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetGoodID(u)
	return cuo
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableGoodID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetGoodID(*u)
	}
	return cuo
}

// ClearGoodID clears the value of the "good_id" field.
func (cuo *CommentUpdateOne) ClearGoodID() *CommentUpdateOne {
	cuo.mutation.ClearGoodID()
	return cuo
}

// SetAppGoodID sets the "app_good_id" field.
func (cuo *CommentUpdateOne) SetAppGoodID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetAppGoodID(u)
	return cuo
}

// SetOrderID sets the "order_id" field.
func (cuo *CommentUpdateOne) SetOrderID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetOrderID(u)
	return cuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableOrderID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetOrderID(*u)
	}
	return cuo
}

// ClearOrderID clears the value of the "order_id" field.
func (cuo *CommentUpdateOne) ClearOrderID() *CommentUpdateOne {
	cuo.mutation.ClearOrderID()
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableContent(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// ClearContent clears the value of the "content" field.
func (cuo *CommentUpdateOne) ClearContent() *CommentUpdateOne {
	cuo.mutation.ClearContent()
	return cuo
}

// SetReplyToID sets the "reply_to_id" field.
func (cuo *CommentUpdateOne) SetReplyToID(u uuid.UUID) *CommentUpdateOne {
	cuo.mutation.SetReplyToID(u)
	return cuo
}

// SetNillableReplyToID sets the "reply_to_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableReplyToID(u *uuid.UUID) *CommentUpdateOne {
	if u != nil {
		cuo.SetReplyToID(*u)
	}
	return cuo
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (cuo *CommentUpdateOne) ClearReplyToID() *CommentUpdateOne {
	cuo.mutation.ClearReplyToID()
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Comment)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CommentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (cuo *CommentUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *CommentUpdateOne {
	cuo.modifiers = append(cuo.modifiers, modifiers...)
	return cuo
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: comment.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldDeletedAt,
		})
	}
	if value, ok := cuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldAppID,
		})
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldUserID,
		})
	}
	if value, ok := cuo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldGoodID,
		})
	}
	if cuo.mutation.GoodIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: comment.FieldGoodID,
		})
	}
	if value, ok := cuo.mutation.AppGoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldAppGoodID,
		})
	}
	if value, ok := cuo.mutation.OrderID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldOrderID,
		})
	}
	if cuo.mutation.OrderIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: comment.FieldOrderID,
		})
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
	}
	if cuo.mutation.ContentCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: comment.FieldContent,
		})
	}
	if value, ok := cuo.mutation.ReplyToID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldReplyToID,
		})
	}
	if cuo.mutation.ReplyToIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Column: comment.FieldReplyToID,
		})
	}
	_spec.Modifiers = cuo.modifiers
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
