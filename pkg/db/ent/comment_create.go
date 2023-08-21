// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/comment"
	"github.com/google/uuid"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommentCreate) SetCreatedAt(u uint32) *CommentCreate {
	cc.mutation.SetCreatedAt(u)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableCreatedAt(u *uint32) *CommentCreate {
	if u != nil {
		cc.SetCreatedAt(*u)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CommentCreate) SetUpdatedAt(u uint32) *CommentCreate {
	cc.mutation.SetUpdatedAt(u)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableUpdatedAt(u *uint32) *CommentCreate {
	if u != nil {
		cc.SetUpdatedAt(*u)
	}
	return cc
}

// SetDeletedAt sets the "deleted_at" field.
func (cc *CommentCreate) SetDeletedAt(u uint32) *CommentCreate {
	cc.mutation.SetDeletedAt(u)
	return cc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cc *CommentCreate) SetNillableDeletedAt(u *uint32) *CommentCreate {
	if u != nil {
		cc.SetDeletedAt(*u)
	}
	return cc
}

// SetAppID sets the "app_id" field.
func (cc *CommentCreate) SetAppID(u uuid.UUID) *CommentCreate {
	cc.mutation.SetAppID(u)
	return cc
}

// SetUserID sets the "user_id" field.
func (cc *CommentCreate) SetUserID(u uuid.UUID) *CommentCreate {
	cc.mutation.SetUserID(u)
	return cc
}

// SetGoodID sets the "good_id" field.
func (cc *CommentCreate) SetGoodID(u uuid.UUID) *CommentCreate {
	cc.mutation.SetGoodID(u)
	return cc
}

// SetOrderID sets the "order_id" field.
func (cc *CommentCreate) SetOrderID(u uuid.UUID) *CommentCreate {
	cc.mutation.SetOrderID(u)
	return cc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableOrderID(u *uuid.UUID) *CommentCreate {
	if u != nil {
		cc.SetOrderID(*u)
	}
	return cc
}

// SetContent sets the "content" field.
func (cc *CommentCreate) SetContent(s string) *CommentCreate {
	cc.mutation.SetContent(s)
	return cc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cc *CommentCreate) SetNillableContent(s *string) *CommentCreate {
	if s != nil {
		cc.SetContent(*s)
	}
	return cc
}

// SetReplyToID sets the "reply_to_id" field.
func (cc *CommentCreate) SetReplyToID(u uuid.UUID) *CommentCreate {
	cc.mutation.SetReplyToID(u)
	return cc
}

// SetNillableReplyToID sets the "reply_to_id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableReplyToID(u *uuid.UUID) *CommentCreate {
	if u != nil {
		cc.SetReplyToID(*u)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(u uuid.UUID) *CommentCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *CommentCreate) SetNillableID(u *uuid.UUID) *CommentCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if err := cc.defaults(); err != nil {
		return nil, err
	}
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CommentCreate) defaults() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		if comment.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := comment.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		if comment.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		if comment.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := comment.DefaultDeletedAt()
		cc.mutation.SetDeletedAt(v)
	}
	if _, ok := cc.mutation.OrderID(); !ok {
		if comment.DefaultOrderID == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultOrderID (forgotten import ent/runtime?)")
		}
		v := comment.DefaultOrderID()
		cc.mutation.SetOrderID(v)
	}
	if _, ok := cc.mutation.Content(); !ok {
		v := comment.DefaultContent
		cc.mutation.SetContent(v)
	}
	if _, ok := cc.mutation.ReplyToID(); !ok {
		if comment.DefaultReplyToID == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultReplyToID (forgotten import ent/runtime?)")
		}
		v := comment.DefaultReplyToID()
		cc.mutation.SetReplyToID(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		if comment.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized comment.DefaultID (forgotten import ent/runtime?)")
		}
		v := comment.DefaultID()
		cc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Comment.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Comment.updated_at"`)}
	}
	if _, ok := cc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Comment.deleted_at"`)}
	}
	if _, ok := cc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "Comment.app_id"`)}
	}
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Comment.user_id"`)}
	}
	if _, ok := cc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "Comment.good_id"`)}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: comment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: comment.FieldID,
			},
		}
	)
	_spec.OnConflict = cc.conflict
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: comment.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := cc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := cc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := cc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := cc.mutation.OrderID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldOrderID,
		})
		_node.OrderID = value
	}
	if value, ok := cc.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: comment.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := cc.mutation.ReplyToID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: comment.FieldReplyToID,
		})
		_node.ReplyToID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Comment.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CommentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (cc *CommentCreate) OnConflict(opts ...sql.ConflictOption) *CommentUpsertOne {
	cc.conflict = opts
	return &CommentUpsertOne{
		create: cc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (cc *CommentCreate) OnConflictColumns(columns ...string) *CommentUpsertOne {
	cc.conflict = append(cc.conflict, sql.ConflictColumns(columns...))
	return &CommentUpsertOne{
		create: cc,
	}
}

type (
	// CommentUpsertOne is the builder for "upsert"-ing
	//  one Comment node.
	CommentUpsertOne struct {
		create *CommentCreate
	}

	// CommentUpsert is the "OnConflict" setter.
	CommentUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *CommentUpsert) SetCreatedAt(v uint32) *CommentUpsert {
	u.Set(comment.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CommentUpsert) UpdateCreatedAt() *CommentUpsert {
	u.SetExcluded(comment.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CommentUpsert) AddCreatedAt(v uint32) *CommentUpsert {
	u.Add(comment.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CommentUpsert) SetUpdatedAt(v uint32) *CommentUpsert {
	u.Set(comment.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CommentUpsert) UpdateUpdatedAt() *CommentUpsert {
	u.SetExcluded(comment.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CommentUpsert) AddUpdatedAt(v uint32) *CommentUpsert {
	u.Add(comment.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CommentUpsert) SetDeletedAt(v uint32) *CommentUpsert {
	u.Set(comment.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CommentUpsert) UpdateDeletedAt() *CommentUpsert {
	u.SetExcluded(comment.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CommentUpsert) AddDeletedAt(v uint32) *CommentUpsert {
	u.Add(comment.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *CommentUpsert) SetAppID(v uuid.UUID) *CommentUpsert {
	u.Set(comment.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CommentUpsert) UpdateAppID() *CommentUpsert {
	u.SetExcluded(comment.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *CommentUpsert) SetUserID(v uuid.UUID) *CommentUpsert {
	u.Set(comment.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CommentUpsert) UpdateUserID() *CommentUpsert {
	u.SetExcluded(comment.FieldUserID)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *CommentUpsert) SetGoodID(v uuid.UUID) *CommentUpsert {
	u.Set(comment.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *CommentUpsert) UpdateGoodID() *CommentUpsert {
	u.SetExcluded(comment.FieldGoodID)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *CommentUpsert) SetOrderID(v uuid.UUID) *CommentUpsert {
	u.Set(comment.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CommentUpsert) UpdateOrderID() *CommentUpsert {
	u.SetExcluded(comment.FieldOrderID)
	return u
}

// ClearOrderID clears the value of the "order_id" field.
func (u *CommentUpsert) ClearOrderID() *CommentUpsert {
	u.SetNull(comment.FieldOrderID)
	return u
}

// SetContent sets the "content" field.
func (u *CommentUpsert) SetContent(v string) *CommentUpsert {
	u.Set(comment.FieldContent, v)
	return u
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *CommentUpsert) UpdateContent() *CommentUpsert {
	u.SetExcluded(comment.FieldContent)
	return u
}

// ClearContent clears the value of the "content" field.
func (u *CommentUpsert) ClearContent() *CommentUpsert {
	u.SetNull(comment.FieldContent)
	return u
}

// SetReplyToID sets the "reply_to_id" field.
func (u *CommentUpsert) SetReplyToID(v uuid.UUID) *CommentUpsert {
	u.Set(comment.FieldReplyToID, v)
	return u
}

// UpdateReplyToID sets the "reply_to_id" field to the value that was provided on create.
func (u *CommentUpsert) UpdateReplyToID() *CommentUpsert {
	u.SetExcluded(comment.FieldReplyToID)
	return u
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (u *CommentUpsert) ClearReplyToID() *CommentUpsert {
	u.SetNull(comment.FieldReplyToID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(comment.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CommentUpsertOne) UpdateNewValues() *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(comment.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Comment.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *CommentUpsertOne) Ignore() *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CommentUpsertOne) DoNothing() *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CommentCreate.OnConflict
// documentation for more info.
func (u *CommentUpsertOne) Update(set func(*CommentUpsert)) *CommentUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CommentUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CommentUpsertOne) SetCreatedAt(v uint32) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CommentUpsertOne) AddCreatedAt(v uint32) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateCreatedAt() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CommentUpsertOne) SetUpdatedAt(v uint32) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CommentUpsertOne) AddUpdatedAt(v uint32) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateUpdatedAt() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CommentUpsertOne) SetDeletedAt(v uint32) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CommentUpsertOne) AddDeletedAt(v uint32) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateDeletedAt() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *CommentUpsertOne) SetAppID(v uuid.UUID) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateAppID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *CommentUpsertOne) SetUserID(v uuid.UUID) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateUserID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUserID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *CommentUpsertOne) SetGoodID(v uuid.UUID) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateGoodID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateGoodID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *CommentUpsertOne) SetOrderID(v uuid.UUID) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateOrderID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *CommentUpsertOne) ClearOrderID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.ClearOrderID()
	})
}

// SetContent sets the "content" field.
func (u *CommentUpsertOne) SetContent(v string) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateContent() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateContent()
	})
}

// ClearContent clears the value of the "content" field.
func (u *CommentUpsertOne) ClearContent() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.ClearContent()
	})
}

// SetReplyToID sets the "reply_to_id" field.
func (u *CommentUpsertOne) SetReplyToID(v uuid.UUID) *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.SetReplyToID(v)
	})
}

// UpdateReplyToID sets the "reply_to_id" field to the value that was provided on create.
func (u *CommentUpsertOne) UpdateReplyToID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateReplyToID()
	})
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (u *CommentUpsertOne) ClearReplyToID() *CommentUpsertOne {
	return u.Update(func(s *CommentUpsert) {
		s.ClearReplyToID()
	})
}

// Exec executes the query.
func (u *CommentUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CommentCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommentUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *CommentUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: CommentUpsertOne.ID is not supported by MySQL driver. Use CommentUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *CommentUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	builders []*CommentCreate
	conflict []sql.ConflictOption
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Comment.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.CommentUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (ccb *CommentCreateBulk) OnConflict(opts ...sql.ConflictOption) *CommentUpsertBulk {
	ccb.conflict = opts
	return &CommentUpsertBulk{
		create: ccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ccb *CommentCreateBulk) OnConflictColumns(columns ...string) *CommentUpsertBulk {
	ccb.conflict = append(ccb.conflict, sql.ConflictColumns(columns...))
	return &CommentUpsertBulk{
		create: ccb,
	}
}

// CommentUpsertBulk is the builder for "upsert"-ing
// a bulk of Comment nodes.
type CommentUpsertBulk struct {
	create *CommentCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(comment.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *CommentUpsertBulk) UpdateNewValues() *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(comment.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Comment.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *CommentUpsertBulk) Ignore() *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *CommentUpsertBulk) DoNothing() *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the CommentCreateBulk.OnConflict
// documentation for more info.
func (u *CommentUpsertBulk) Update(set func(*CommentUpsert)) *CommentUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&CommentUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *CommentUpsertBulk) SetCreatedAt(v uint32) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *CommentUpsertBulk) AddCreatedAt(v uint32) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateCreatedAt() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *CommentUpsertBulk) SetUpdatedAt(v uint32) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *CommentUpsertBulk) AddUpdatedAt(v uint32) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateUpdatedAt() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *CommentUpsertBulk) SetDeletedAt(v uint32) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *CommentUpsertBulk) AddDeletedAt(v uint32) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateDeletedAt() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *CommentUpsertBulk) SetAppID(v uuid.UUID) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateAppID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *CommentUpsertBulk) SetUserID(v uuid.UUID) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateUserID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateUserID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *CommentUpsertBulk) SetGoodID(v uuid.UUID) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateGoodID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateGoodID()
	})
}

// SetOrderID sets the "order_id" field.
func (u *CommentUpsertBulk) SetOrderID(v uuid.UUID) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateOrderID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *CommentUpsertBulk) ClearOrderID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.ClearOrderID()
	})
}

// SetContent sets the "content" field.
func (u *CommentUpsertBulk) SetContent(v string) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetContent(v)
	})
}

// UpdateContent sets the "content" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateContent() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateContent()
	})
}

// ClearContent clears the value of the "content" field.
func (u *CommentUpsertBulk) ClearContent() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.ClearContent()
	})
}

// SetReplyToID sets the "reply_to_id" field.
func (u *CommentUpsertBulk) SetReplyToID(v uuid.UUID) *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.SetReplyToID(v)
	})
}

// UpdateReplyToID sets the "reply_to_id" field to the value that was provided on create.
func (u *CommentUpsertBulk) UpdateReplyToID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.UpdateReplyToID()
	})
}

// ClearReplyToID clears the value of the "reply_to_id" field.
func (u *CommentUpsertBulk) ClearReplyToID() *CommentUpsertBulk {
	return u.Update(func(s *CommentUpsert) {
		s.ClearReplyToID()
	})
}

// Exec executes the query.
func (u *CommentUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the CommentCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for CommentCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommentUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
