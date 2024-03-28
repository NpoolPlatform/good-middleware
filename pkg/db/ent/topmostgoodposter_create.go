// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodposter"
	"github.com/google/uuid"
)

// TopMostGoodPosterCreate is the builder for creating a TopMostGoodPoster entity.
type TopMostGoodPosterCreate struct {
	config
	mutation *TopMostGoodPosterMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tmgpc *TopMostGoodPosterCreate) SetCreatedAt(u uint32) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetCreatedAt(u)
	return tmgpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmgpc *TopMostGoodPosterCreate) SetNillableCreatedAt(u *uint32) *TopMostGoodPosterCreate {
	if u != nil {
		tmgpc.SetCreatedAt(*u)
	}
	return tmgpc
}

// SetUpdatedAt sets the "updated_at" field.
func (tmgpc *TopMostGoodPosterCreate) SetUpdatedAt(u uint32) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetUpdatedAt(u)
	return tmgpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tmgpc *TopMostGoodPosterCreate) SetNillableUpdatedAt(u *uint32) *TopMostGoodPosterCreate {
	if u != nil {
		tmgpc.SetUpdatedAt(*u)
	}
	return tmgpc
}

// SetDeletedAt sets the "deleted_at" field.
func (tmgpc *TopMostGoodPosterCreate) SetDeletedAt(u uint32) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetDeletedAt(u)
	return tmgpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmgpc *TopMostGoodPosterCreate) SetNillableDeletedAt(u *uint32) *TopMostGoodPosterCreate {
	if u != nil {
		tmgpc.SetDeletedAt(*u)
	}
	return tmgpc
}

// SetEntID sets the "ent_id" field.
func (tmgpc *TopMostGoodPosterCreate) SetEntID(u uuid.UUID) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetEntID(u)
	return tmgpc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmgpc *TopMostGoodPosterCreate) SetNillableEntID(u *uuid.UUID) *TopMostGoodPosterCreate {
	if u != nil {
		tmgpc.SetEntID(*u)
	}
	return tmgpc
}

// SetTopMostGoodID sets the "top_most_good_id" field.
func (tmgpc *TopMostGoodPosterCreate) SetTopMostGoodID(u uuid.UUID) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetTopMostGoodID(u)
	return tmgpc
}

// SetNillableTopMostGoodID sets the "top_most_good_id" field if the given value is not nil.
func (tmgpc *TopMostGoodPosterCreate) SetNillableTopMostGoodID(u *uuid.UUID) *TopMostGoodPosterCreate {
	if u != nil {
		tmgpc.SetTopMostGoodID(*u)
	}
	return tmgpc
}

// SetPoster sets the "poster" field.
func (tmgpc *TopMostGoodPosterCreate) SetPoster(s string) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetPoster(s)
	return tmgpc
}

// SetNillablePoster sets the "poster" field if the given value is not nil.
func (tmgpc *TopMostGoodPosterCreate) SetNillablePoster(s *string) *TopMostGoodPosterCreate {
	if s != nil {
		tmgpc.SetPoster(*s)
	}
	return tmgpc
}

// SetIndex sets the "index" field.
func (tmgpc *TopMostGoodPosterCreate) SetIndex(u uint8) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetIndex(u)
	return tmgpc
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tmgpc *TopMostGoodPosterCreate) SetNillableIndex(u *uint8) *TopMostGoodPosterCreate {
	if u != nil {
		tmgpc.SetIndex(*u)
	}
	return tmgpc
}

// SetID sets the "id" field.
func (tmgpc *TopMostGoodPosterCreate) SetID(u uint32) *TopMostGoodPosterCreate {
	tmgpc.mutation.SetID(u)
	return tmgpc
}

// Mutation returns the TopMostGoodPosterMutation object of the builder.
func (tmgpc *TopMostGoodPosterCreate) Mutation() *TopMostGoodPosterMutation {
	return tmgpc.mutation
}

// Save creates the TopMostGoodPoster in the database.
func (tmgpc *TopMostGoodPosterCreate) Save(ctx context.Context) (*TopMostGoodPoster, error) {
	var (
		err  error
		node *TopMostGoodPoster
	)
	if err := tmgpc.defaults(); err != nil {
		return nil, err
	}
	if len(tmgpc.hooks) == 0 {
		if err = tmgpc.check(); err != nil {
			return nil, err
		}
		node, err = tmgpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopMostGoodPosterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tmgpc.check(); err != nil {
				return nil, err
			}
			tmgpc.mutation = mutation
			if node, err = tmgpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tmgpc.hooks) - 1; i >= 0; i-- {
			if tmgpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmgpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tmgpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TopMostGoodPoster)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TopMostGoodPosterMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tmgpc *TopMostGoodPosterCreate) SaveX(ctx context.Context) *TopMostGoodPoster {
	v, err := tmgpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmgpc *TopMostGoodPosterCreate) Exec(ctx context.Context) error {
	_, err := tmgpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmgpc *TopMostGoodPosterCreate) ExecX(ctx context.Context) {
	if err := tmgpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmgpc *TopMostGoodPosterCreate) defaults() error {
	if _, ok := tmgpc.mutation.CreatedAt(); !ok {
		if topmostgoodposter.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostgoodposter.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := topmostgoodposter.DefaultCreatedAt()
		tmgpc.mutation.SetCreatedAt(v)
	}
	if _, ok := tmgpc.mutation.UpdatedAt(); !ok {
		if topmostgoodposter.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostgoodposter.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topmostgoodposter.DefaultUpdatedAt()
		tmgpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tmgpc.mutation.DeletedAt(); !ok {
		if topmostgoodposter.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostgoodposter.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := topmostgoodposter.DefaultDeletedAt()
		tmgpc.mutation.SetDeletedAt(v)
	}
	if _, ok := tmgpc.mutation.EntID(); !ok {
		if topmostgoodposter.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized topmostgoodposter.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := topmostgoodposter.DefaultEntID()
		tmgpc.mutation.SetEntID(v)
	}
	if _, ok := tmgpc.mutation.TopMostGoodID(); !ok {
		if topmostgoodposter.DefaultTopMostGoodID == nil {
			return fmt.Errorf("ent: uninitialized topmostgoodposter.DefaultTopMostGoodID (forgotten import ent/runtime?)")
		}
		v := topmostgoodposter.DefaultTopMostGoodID()
		tmgpc.mutation.SetTopMostGoodID(v)
	}
	if _, ok := tmgpc.mutation.Poster(); !ok {
		v := topmostgoodposter.DefaultPoster
		tmgpc.mutation.SetPoster(v)
	}
	if _, ok := tmgpc.mutation.Index(); !ok {
		v := topmostgoodposter.DefaultIndex
		tmgpc.mutation.SetIndex(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tmgpc *TopMostGoodPosterCreate) check() error {
	if _, ok := tmgpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TopMostGoodPoster.created_at"`)}
	}
	if _, ok := tmgpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TopMostGoodPoster.updated_at"`)}
	}
	if _, ok := tmgpc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "TopMostGoodPoster.deleted_at"`)}
	}
	if _, ok := tmgpc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "TopMostGoodPoster.ent_id"`)}
	}
	return nil
}

func (tmgpc *TopMostGoodPosterCreate) sqlSave(ctx context.Context) (*TopMostGoodPoster, error) {
	_node, _spec := tmgpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tmgpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint32(id)
	}
	return _node, nil
}

func (tmgpc *TopMostGoodPosterCreate) createSpec() (*TopMostGoodPoster, *sqlgraph.CreateSpec) {
	var (
		_node = &TopMostGoodPoster{config: tmgpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: topmostgoodposter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostgoodposter.FieldID,
			},
		}
	)
	_spec.OnConflict = tmgpc.conflict
	if id, ok := tmgpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tmgpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgoodposter.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tmgpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgoodposter.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tmgpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostgoodposter.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := tmgpc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgoodposter.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := tmgpc.mutation.TopMostGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostgoodposter.FieldTopMostGoodID,
		})
		_node.TopMostGoodID = value
	}
	if value, ok := tmgpc.mutation.Poster(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topmostgoodposter.FieldPoster,
		})
		_node.Poster = value
	}
	if value, ok := tmgpc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: topmostgoodposter.FieldIndex,
		})
		_node.Index = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TopMostGoodPoster.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TopMostGoodPosterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tmgpc *TopMostGoodPosterCreate) OnConflict(opts ...sql.ConflictOption) *TopMostGoodPosterUpsertOne {
	tmgpc.conflict = opts
	return &TopMostGoodPosterUpsertOne{
		create: tmgpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TopMostGoodPoster.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tmgpc *TopMostGoodPosterCreate) OnConflictColumns(columns ...string) *TopMostGoodPosterUpsertOne {
	tmgpc.conflict = append(tmgpc.conflict, sql.ConflictColumns(columns...))
	return &TopMostGoodPosterUpsertOne{
		create: tmgpc,
	}
}

type (
	// TopMostGoodPosterUpsertOne is the builder for "upsert"-ing
	//  one TopMostGoodPoster node.
	TopMostGoodPosterUpsertOne struct {
		create *TopMostGoodPosterCreate
	}

	// TopMostGoodPosterUpsert is the "OnConflict" setter.
	TopMostGoodPosterUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *TopMostGoodPosterUpsert) SetCreatedAt(v uint32) *TopMostGoodPosterUpsert {
	u.Set(topmostgoodposter.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsert) UpdateCreatedAt() *TopMostGoodPosterUpsert {
	u.SetExcluded(topmostgoodposter.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TopMostGoodPosterUpsert) AddCreatedAt(v uint32) *TopMostGoodPosterUpsert {
	u.Add(topmostgoodposter.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TopMostGoodPosterUpsert) SetUpdatedAt(v uint32) *TopMostGoodPosterUpsert {
	u.Set(topmostgoodposter.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsert) UpdateUpdatedAt() *TopMostGoodPosterUpsert {
	u.SetExcluded(topmostgoodposter.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TopMostGoodPosterUpsert) AddUpdatedAt(v uint32) *TopMostGoodPosterUpsert {
	u.Add(topmostgoodposter.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TopMostGoodPosterUpsert) SetDeletedAt(v uint32) *TopMostGoodPosterUpsert {
	u.Set(topmostgoodposter.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsert) UpdateDeletedAt() *TopMostGoodPosterUpsert {
	u.SetExcluded(topmostgoodposter.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TopMostGoodPosterUpsert) AddDeletedAt(v uint32) *TopMostGoodPosterUpsert {
	u.Add(topmostgoodposter.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *TopMostGoodPosterUpsert) SetEntID(v uuid.UUID) *TopMostGoodPosterUpsert {
	u.Set(topmostgoodposter.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsert) UpdateEntID() *TopMostGoodPosterUpsert {
	u.SetExcluded(topmostgoodposter.FieldEntID)
	return u
}

// SetTopMostGoodID sets the "top_most_good_id" field.
func (u *TopMostGoodPosterUpsert) SetTopMostGoodID(v uuid.UUID) *TopMostGoodPosterUpsert {
	u.Set(topmostgoodposter.FieldTopMostGoodID, v)
	return u
}

// UpdateTopMostGoodID sets the "top_most_good_id" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsert) UpdateTopMostGoodID() *TopMostGoodPosterUpsert {
	u.SetExcluded(topmostgoodposter.FieldTopMostGoodID)
	return u
}

// ClearTopMostGoodID clears the value of the "top_most_good_id" field.
func (u *TopMostGoodPosterUpsert) ClearTopMostGoodID() *TopMostGoodPosterUpsert {
	u.SetNull(topmostgoodposter.FieldTopMostGoodID)
	return u
}

// SetPoster sets the "poster" field.
func (u *TopMostGoodPosterUpsert) SetPoster(v string) *TopMostGoodPosterUpsert {
	u.Set(topmostgoodposter.FieldPoster, v)
	return u
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsert) UpdatePoster() *TopMostGoodPosterUpsert {
	u.SetExcluded(topmostgoodposter.FieldPoster)
	return u
}

// ClearPoster clears the value of the "poster" field.
func (u *TopMostGoodPosterUpsert) ClearPoster() *TopMostGoodPosterUpsert {
	u.SetNull(topmostgoodposter.FieldPoster)
	return u
}

// SetIndex sets the "index" field.
func (u *TopMostGoodPosterUpsert) SetIndex(v uint8) *TopMostGoodPosterUpsert {
	u.Set(topmostgoodposter.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsert) UpdateIndex() *TopMostGoodPosterUpsert {
	u.SetExcluded(topmostgoodposter.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *TopMostGoodPosterUpsert) AddIndex(v uint8) *TopMostGoodPosterUpsert {
	u.Add(topmostgoodposter.FieldIndex, v)
	return u
}

// ClearIndex clears the value of the "index" field.
func (u *TopMostGoodPosterUpsert) ClearIndex() *TopMostGoodPosterUpsert {
	u.SetNull(topmostgoodposter.FieldIndex)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TopMostGoodPoster.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(topmostgoodposter.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TopMostGoodPosterUpsertOne) UpdateNewValues() *TopMostGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(topmostgoodposter.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.TopMostGoodPoster.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TopMostGoodPosterUpsertOne) Ignore() *TopMostGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TopMostGoodPosterUpsertOne) DoNothing() *TopMostGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TopMostGoodPosterCreate.OnConflict
// documentation for more info.
func (u *TopMostGoodPosterUpsertOne) Update(set func(*TopMostGoodPosterUpsert)) *TopMostGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TopMostGoodPosterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TopMostGoodPosterUpsertOne) SetCreatedAt(v uint32) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TopMostGoodPosterUpsertOne) AddCreatedAt(v uint32) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertOne) UpdateCreatedAt() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TopMostGoodPosterUpsertOne) SetUpdatedAt(v uint32) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TopMostGoodPosterUpsertOne) AddUpdatedAt(v uint32) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertOne) UpdateUpdatedAt() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TopMostGoodPosterUpsertOne) SetDeletedAt(v uint32) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TopMostGoodPosterUpsertOne) AddDeletedAt(v uint32) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertOne) UpdateDeletedAt() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *TopMostGoodPosterUpsertOne) SetEntID(v uuid.UUID) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertOne) UpdateEntID() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateEntID()
	})
}

// SetTopMostGoodID sets the "top_most_good_id" field.
func (u *TopMostGoodPosterUpsertOne) SetTopMostGoodID(v uuid.UUID) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetTopMostGoodID(v)
	})
}

// UpdateTopMostGoodID sets the "top_most_good_id" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertOne) UpdateTopMostGoodID() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateTopMostGoodID()
	})
}

// ClearTopMostGoodID clears the value of the "top_most_good_id" field.
func (u *TopMostGoodPosterUpsertOne) ClearTopMostGoodID() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.ClearTopMostGoodID()
	})
}

// SetPoster sets the "poster" field.
func (u *TopMostGoodPosterUpsertOne) SetPoster(v string) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetPoster(v)
	})
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertOne) UpdatePoster() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdatePoster()
	})
}

// ClearPoster clears the value of the "poster" field.
func (u *TopMostGoodPosterUpsertOne) ClearPoster() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.ClearPoster()
	})
}

// SetIndex sets the "index" field.
func (u *TopMostGoodPosterUpsertOne) SetIndex(v uint8) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *TopMostGoodPosterUpsertOne) AddIndex(v uint8) *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertOne) UpdateIndex() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *TopMostGoodPosterUpsertOne) ClearIndex() *TopMostGoodPosterUpsertOne {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *TopMostGoodPosterUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TopMostGoodPosterCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TopMostGoodPosterUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TopMostGoodPosterUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TopMostGoodPosterUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TopMostGoodPosterCreateBulk is the builder for creating many TopMostGoodPoster entities in bulk.
type TopMostGoodPosterCreateBulk struct {
	config
	builders []*TopMostGoodPosterCreate
	conflict []sql.ConflictOption
}

// Save creates the TopMostGoodPoster entities in the database.
func (tmgpcb *TopMostGoodPosterCreateBulk) Save(ctx context.Context) ([]*TopMostGoodPoster, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tmgpcb.builders))
	nodes := make([]*TopMostGoodPoster, len(tmgpcb.builders))
	mutators := make([]Mutator, len(tmgpcb.builders))
	for i := range tmgpcb.builders {
		func(i int, root context.Context) {
			builder := tmgpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopMostGoodPosterMutation)
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
					_, err = mutators[i+1].Mutate(root, tmgpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tmgpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tmgpcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint32(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tmgpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tmgpcb *TopMostGoodPosterCreateBulk) SaveX(ctx context.Context) []*TopMostGoodPoster {
	v, err := tmgpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmgpcb *TopMostGoodPosterCreateBulk) Exec(ctx context.Context) error {
	_, err := tmgpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmgpcb *TopMostGoodPosterCreateBulk) ExecX(ctx context.Context) {
	if err := tmgpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TopMostGoodPoster.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TopMostGoodPosterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (tmgpcb *TopMostGoodPosterCreateBulk) OnConflict(opts ...sql.ConflictOption) *TopMostGoodPosterUpsertBulk {
	tmgpcb.conflict = opts
	return &TopMostGoodPosterUpsertBulk{
		create: tmgpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TopMostGoodPoster.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tmgpcb *TopMostGoodPosterCreateBulk) OnConflictColumns(columns ...string) *TopMostGoodPosterUpsertBulk {
	tmgpcb.conflict = append(tmgpcb.conflict, sql.ConflictColumns(columns...))
	return &TopMostGoodPosterUpsertBulk{
		create: tmgpcb,
	}
}

// TopMostGoodPosterUpsertBulk is the builder for "upsert"-ing
// a bulk of TopMostGoodPoster nodes.
type TopMostGoodPosterUpsertBulk struct {
	create *TopMostGoodPosterCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TopMostGoodPoster.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(topmostgoodposter.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TopMostGoodPosterUpsertBulk) UpdateNewValues() *TopMostGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(topmostgoodposter.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TopMostGoodPoster.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TopMostGoodPosterUpsertBulk) Ignore() *TopMostGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TopMostGoodPosterUpsertBulk) DoNothing() *TopMostGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TopMostGoodPosterCreateBulk.OnConflict
// documentation for more info.
func (u *TopMostGoodPosterUpsertBulk) Update(set func(*TopMostGoodPosterUpsert)) *TopMostGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TopMostGoodPosterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TopMostGoodPosterUpsertBulk) SetCreatedAt(v uint32) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TopMostGoodPosterUpsertBulk) AddCreatedAt(v uint32) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertBulk) UpdateCreatedAt() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TopMostGoodPosterUpsertBulk) SetUpdatedAt(v uint32) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TopMostGoodPosterUpsertBulk) AddUpdatedAt(v uint32) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertBulk) UpdateUpdatedAt() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TopMostGoodPosterUpsertBulk) SetDeletedAt(v uint32) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TopMostGoodPosterUpsertBulk) AddDeletedAt(v uint32) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertBulk) UpdateDeletedAt() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *TopMostGoodPosterUpsertBulk) SetEntID(v uuid.UUID) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertBulk) UpdateEntID() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateEntID()
	})
}

// SetTopMostGoodID sets the "top_most_good_id" field.
func (u *TopMostGoodPosterUpsertBulk) SetTopMostGoodID(v uuid.UUID) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetTopMostGoodID(v)
	})
}

// UpdateTopMostGoodID sets the "top_most_good_id" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertBulk) UpdateTopMostGoodID() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateTopMostGoodID()
	})
}

// ClearTopMostGoodID clears the value of the "top_most_good_id" field.
func (u *TopMostGoodPosterUpsertBulk) ClearTopMostGoodID() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.ClearTopMostGoodID()
	})
}

// SetPoster sets the "poster" field.
func (u *TopMostGoodPosterUpsertBulk) SetPoster(v string) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetPoster(v)
	})
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertBulk) UpdatePoster() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdatePoster()
	})
}

// ClearPoster clears the value of the "poster" field.
func (u *TopMostGoodPosterUpsertBulk) ClearPoster() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.ClearPoster()
	})
}

// SetIndex sets the "index" field.
func (u *TopMostGoodPosterUpsertBulk) SetIndex(v uint8) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *TopMostGoodPosterUpsertBulk) AddIndex(v uint8) *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *TopMostGoodPosterUpsertBulk) UpdateIndex() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *TopMostGoodPosterUpsertBulk) ClearIndex() *TopMostGoodPosterUpsertBulk {
	return u.Update(func(s *TopMostGoodPosterUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *TopMostGoodPosterUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TopMostGoodPosterCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TopMostGoodPosterCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TopMostGoodPosterUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}