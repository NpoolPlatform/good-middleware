// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodposter"
	"github.com/google/uuid"
)

// AppGoodPosterCreate is the builder for creating a AppGoodPoster entity.
type AppGoodPosterCreate struct {
	config
	mutation *AppGoodPosterMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (agpc *AppGoodPosterCreate) SetCreatedAt(u uint32) *AppGoodPosterCreate {
	agpc.mutation.SetCreatedAt(u)
	return agpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (agpc *AppGoodPosterCreate) SetNillableCreatedAt(u *uint32) *AppGoodPosterCreate {
	if u != nil {
		agpc.SetCreatedAt(*u)
	}
	return agpc
}

// SetUpdatedAt sets the "updated_at" field.
func (agpc *AppGoodPosterCreate) SetUpdatedAt(u uint32) *AppGoodPosterCreate {
	agpc.mutation.SetUpdatedAt(u)
	return agpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (agpc *AppGoodPosterCreate) SetNillableUpdatedAt(u *uint32) *AppGoodPosterCreate {
	if u != nil {
		agpc.SetUpdatedAt(*u)
	}
	return agpc
}

// SetDeletedAt sets the "deleted_at" field.
func (agpc *AppGoodPosterCreate) SetDeletedAt(u uint32) *AppGoodPosterCreate {
	agpc.mutation.SetDeletedAt(u)
	return agpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (agpc *AppGoodPosterCreate) SetNillableDeletedAt(u *uint32) *AppGoodPosterCreate {
	if u != nil {
		agpc.SetDeletedAt(*u)
	}
	return agpc
}

// SetEntID sets the "ent_id" field.
func (agpc *AppGoodPosterCreate) SetEntID(u uuid.UUID) *AppGoodPosterCreate {
	agpc.mutation.SetEntID(u)
	return agpc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (agpc *AppGoodPosterCreate) SetNillableEntID(u *uuid.UUID) *AppGoodPosterCreate {
	if u != nil {
		agpc.SetEntID(*u)
	}
	return agpc
}

// SetAppGoodID sets the "app_good_id" field.
func (agpc *AppGoodPosterCreate) SetAppGoodID(u uuid.UUID) *AppGoodPosterCreate {
	agpc.mutation.SetAppGoodID(u)
	return agpc
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (agpc *AppGoodPosterCreate) SetNillableAppGoodID(u *uuid.UUID) *AppGoodPosterCreate {
	if u != nil {
		agpc.SetAppGoodID(*u)
	}
	return agpc
}

// SetPoster sets the "poster" field.
func (agpc *AppGoodPosterCreate) SetPoster(s string) *AppGoodPosterCreate {
	agpc.mutation.SetPoster(s)
	return agpc
}

// SetNillablePoster sets the "poster" field if the given value is not nil.
func (agpc *AppGoodPosterCreate) SetNillablePoster(s *string) *AppGoodPosterCreate {
	if s != nil {
		agpc.SetPoster(*s)
	}
	return agpc
}

// SetIndex sets the "index" field.
func (agpc *AppGoodPosterCreate) SetIndex(u uint8) *AppGoodPosterCreate {
	agpc.mutation.SetIndex(u)
	return agpc
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (agpc *AppGoodPosterCreate) SetNillableIndex(u *uint8) *AppGoodPosterCreate {
	if u != nil {
		agpc.SetIndex(*u)
	}
	return agpc
}

// SetID sets the "id" field.
func (agpc *AppGoodPosterCreate) SetID(u uint32) *AppGoodPosterCreate {
	agpc.mutation.SetID(u)
	return agpc
}

// Mutation returns the AppGoodPosterMutation object of the builder.
func (agpc *AppGoodPosterCreate) Mutation() *AppGoodPosterMutation {
	return agpc.mutation
}

// Save creates the AppGoodPoster in the database.
func (agpc *AppGoodPosterCreate) Save(ctx context.Context) (*AppGoodPoster, error) {
	var (
		err  error
		node *AppGoodPoster
	)
	if err := agpc.defaults(); err != nil {
		return nil, err
	}
	if len(agpc.hooks) == 0 {
		if err = agpc.check(); err != nil {
			return nil, err
		}
		node, err = agpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodPosterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = agpc.check(); err != nil {
				return nil, err
			}
			agpc.mutation = mutation
			if node, err = agpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(agpc.hooks) - 1; i >= 0; i-- {
			if agpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, agpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppGoodPoster)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppGoodPosterMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (agpc *AppGoodPosterCreate) SaveX(ctx context.Context) *AppGoodPoster {
	v, err := agpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agpc *AppGoodPosterCreate) Exec(ctx context.Context) error {
	_, err := agpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agpc *AppGoodPosterCreate) ExecX(ctx context.Context) {
	if err := agpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agpc *AppGoodPosterCreate) defaults() error {
	if _, ok := agpc.mutation.CreatedAt(); !ok {
		if appgoodposter.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appgoodposter.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appgoodposter.DefaultCreatedAt()
		agpc.mutation.SetCreatedAt(v)
	}
	if _, ok := agpc.mutation.UpdatedAt(); !ok {
		if appgoodposter.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appgoodposter.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appgoodposter.DefaultUpdatedAt()
		agpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := agpc.mutation.DeletedAt(); !ok {
		if appgoodposter.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appgoodposter.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appgoodposter.DefaultDeletedAt()
		agpc.mutation.SetDeletedAt(v)
	}
	if _, ok := agpc.mutation.EntID(); !ok {
		if appgoodposter.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized appgoodposter.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := appgoodposter.DefaultEntID()
		agpc.mutation.SetEntID(v)
	}
	if _, ok := agpc.mutation.AppGoodID(); !ok {
		if appgoodposter.DefaultAppGoodID == nil {
			return fmt.Errorf("ent: uninitialized appgoodposter.DefaultAppGoodID (forgotten import ent/runtime?)")
		}
		v := appgoodposter.DefaultAppGoodID()
		agpc.mutation.SetAppGoodID(v)
	}
	if _, ok := agpc.mutation.Poster(); !ok {
		v := appgoodposter.DefaultPoster
		agpc.mutation.SetPoster(v)
	}
	if _, ok := agpc.mutation.Index(); !ok {
		v := appgoodposter.DefaultIndex
		agpc.mutation.SetIndex(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (agpc *AppGoodPosterCreate) check() error {
	if _, ok := agpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppGoodPoster.created_at"`)}
	}
	if _, ok := agpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppGoodPoster.updated_at"`)}
	}
	if _, ok := agpc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppGoodPoster.deleted_at"`)}
	}
	if _, ok := agpc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "AppGoodPoster.ent_id"`)}
	}
	return nil
}

func (agpc *AppGoodPosterCreate) sqlSave(ctx context.Context) (*AppGoodPoster, error) {
	_node, _spec := agpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, agpc.driver, _spec); err != nil {
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

func (agpc *AppGoodPosterCreate) createSpec() (*AppGoodPoster, *sqlgraph.CreateSpec) {
	var (
		_node = &AppGoodPoster{config: agpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appgoodposter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgoodposter.FieldID,
			},
		}
	)
	_spec.OnConflict = agpc.conflict
	if id, ok := agpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := agpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgoodposter.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := agpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgoodposter.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := agpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgoodposter.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := agpc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodposter.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := agpc.mutation.AppGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgoodposter.FieldAppGoodID,
		})
		_node.AppGoodID = value
	}
	if value, ok := agpc.mutation.Poster(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appgoodposter.FieldPoster,
		})
		_node.Poster = value
	}
	if value, ok := agpc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appgoodposter.FieldIndex,
		})
		_node.Index = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppGoodPoster.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppGoodPosterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (agpc *AppGoodPosterCreate) OnConflict(opts ...sql.ConflictOption) *AppGoodPosterUpsertOne {
	agpc.conflict = opts
	return &AppGoodPosterUpsertOne{
		create: agpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppGoodPoster.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (agpc *AppGoodPosterCreate) OnConflictColumns(columns ...string) *AppGoodPosterUpsertOne {
	agpc.conflict = append(agpc.conflict, sql.ConflictColumns(columns...))
	return &AppGoodPosterUpsertOne{
		create: agpc,
	}
}

type (
	// AppGoodPosterUpsertOne is the builder for "upsert"-ing
	//  one AppGoodPoster node.
	AppGoodPosterUpsertOne struct {
		create *AppGoodPosterCreate
	}

	// AppGoodPosterUpsert is the "OnConflict" setter.
	AppGoodPosterUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppGoodPosterUpsert) SetCreatedAt(v uint32) *AppGoodPosterUpsert {
	u.Set(appgoodposter.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsert) UpdateCreatedAt() *AppGoodPosterUpsert {
	u.SetExcluded(appgoodposter.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppGoodPosterUpsert) AddCreatedAt(v uint32) *AppGoodPosterUpsert {
	u.Add(appgoodposter.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppGoodPosterUpsert) SetUpdatedAt(v uint32) *AppGoodPosterUpsert {
	u.Set(appgoodposter.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsert) UpdateUpdatedAt() *AppGoodPosterUpsert {
	u.SetExcluded(appgoodposter.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppGoodPosterUpsert) AddUpdatedAt(v uint32) *AppGoodPosterUpsert {
	u.Add(appgoodposter.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppGoodPosterUpsert) SetDeletedAt(v uint32) *AppGoodPosterUpsert {
	u.Set(appgoodposter.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsert) UpdateDeletedAt() *AppGoodPosterUpsert {
	u.SetExcluded(appgoodposter.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppGoodPosterUpsert) AddDeletedAt(v uint32) *AppGoodPosterUpsert {
	u.Add(appgoodposter.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppGoodPosterUpsert) SetEntID(v uuid.UUID) *AppGoodPosterUpsert {
	u.Set(appgoodposter.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppGoodPosterUpsert) UpdateEntID() *AppGoodPosterUpsert {
	u.SetExcluded(appgoodposter.FieldEntID)
	return u
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppGoodPosterUpsert) SetAppGoodID(v uuid.UUID) *AppGoodPosterUpsert {
	u.Set(appgoodposter.FieldAppGoodID, v)
	return u
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppGoodPosterUpsert) UpdateAppGoodID() *AppGoodPosterUpsert {
	u.SetExcluded(appgoodposter.FieldAppGoodID)
	return u
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppGoodPosterUpsert) ClearAppGoodID() *AppGoodPosterUpsert {
	u.SetNull(appgoodposter.FieldAppGoodID)
	return u
}

// SetPoster sets the "poster" field.
func (u *AppGoodPosterUpsert) SetPoster(v string) *AppGoodPosterUpsert {
	u.Set(appgoodposter.FieldPoster, v)
	return u
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *AppGoodPosterUpsert) UpdatePoster() *AppGoodPosterUpsert {
	u.SetExcluded(appgoodposter.FieldPoster)
	return u
}

// ClearPoster clears the value of the "poster" field.
func (u *AppGoodPosterUpsert) ClearPoster() *AppGoodPosterUpsert {
	u.SetNull(appgoodposter.FieldPoster)
	return u
}

// SetIndex sets the "index" field.
func (u *AppGoodPosterUpsert) SetIndex(v uint8) *AppGoodPosterUpsert {
	u.Set(appgoodposter.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *AppGoodPosterUpsert) UpdateIndex() *AppGoodPosterUpsert {
	u.SetExcluded(appgoodposter.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *AppGoodPosterUpsert) AddIndex(v uint8) *AppGoodPosterUpsert {
	u.Add(appgoodposter.FieldIndex, v)
	return u
}

// ClearIndex clears the value of the "index" field.
func (u *AppGoodPosterUpsert) ClearIndex() *AppGoodPosterUpsert {
	u.SetNull(appgoodposter.FieldIndex)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppGoodPoster.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appgoodposter.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppGoodPosterUpsertOne) UpdateNewValues() *AppGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appgoodposter.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppGoodPoster.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppGoodPosterUpsertOne) Ignore() *AppGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppGoodPosterUpsertOne) DoNothing() *AppGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppGoodPosterCreate.OnConflict
// documentation for more info.
func (u *AppGoodPosterUpsertOne) Update(set func(*AppGoodPosterUpsert)) *AppGoodPosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppGoodPosterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppGoodPosterUpsertOne) SetCreatedAt(v uint32) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppGoodPosterUpsertOne) AddCreatedAt(v uint32) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsertOne) UpdateCreatedAt() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppGoodPosterUpsertOne) SetUpdatedAt(v uint32) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppGoodPosterUpsertOne) AddUpdatedAt(v uint32) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsertOne) UpdateUpdatedAt() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppGoodPosterUpsertOne) SetDeletedAt(v uint32) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppGoodPosterUpsertOne) AddDeletedAt(v uint32) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsertOne) UpdateDeletedAt() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppGoodPosterUpsertOne) SetEntID(v uuid.UUID) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppGoodPosterUpsertOne) UpdateEntID() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppGoodPosterUpsertOne) SetAppGoodID(v uuid.UUID) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppGoodPosterUpsertOne) UpdateAppGoodID() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppGoodPosterUpsertOne) ClearAppGoodID() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.ClearAppGoodID()
	})
}

// SetPoster sets the "poster" field.
func (u *AppGoodPosterUpsertOne) SetPoster(v string) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetPoster(v)
	})
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *AppGoodPosterUpsertOne) UpdatePoster() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdatePoster()
	})
}

// ClearPoster clears the value of the "poster" field.
func (u *AppGoodPosterUpsertOne) ClearPoster() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.ClearPoster()
	})
}

// SetIndex sets the "index" field.
func (u *AppGoodPosterUpsertOne) SetIndex(v uint8) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *AppGoodPosterUpsertOne) AddIndex(v uint8) *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *AppGoodPosterUpsertOne) UpdateIndex() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *AppGoodPosterUpsertOne) ClearIndex() *AppGoodPosterUpsertOne {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *AppGoodPosterUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppGoodPosterCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppGoodPosterUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppGoodPosterUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppGoodPosterUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppGoodPosterCreateBulk is the builder for creating many AppGoodPoster entities in bulk.
type AppGoodPosterCreateBulk struct {
	config
	builders []*AppGoodPosterCreate
	conflict []sql.ConflictOption
}

// Save creates the AppGoodPoster entities in the database.
func (agpcb *AppGoodPosterCreateBulk) Save(ctx context.Context) ([]*AppGoodPoster, error) {
	specs := make([]*sqlgraph.CreateSpec, len(agpcb.builders))
	nodes := make([]*AppGoodPoster, len(agpcb.builders))
	mutators := make([]Mutator, len(agpcb.builders))
	for i := range agpcb.builders {
		func(i int, root context.Context) {
			builder := agpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppGoodPosterMutation)
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
					_, err = mutators[i+1].Mutate(root, agpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = agpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, agpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, agpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (agpcb *AppGoodPosterCreateBulk) SaveX(ctx context.Context) []*AppGoodPoster {
	v, err := agpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agpcb *AppGoodPosterCreateBulk) Exec(ctx context.Context) error {
	_, err := agpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agpcb *AppGoodPosterCreateBulk) ExecX(ctx context.Context) {
	if err := agpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppGoodPoster.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppGoodPosterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (agpcb *AppGoodPosterCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppGoodPosterUpsertBulk {
	agpcb.conflict = opts
	return &AppGoodPosterUpsertBulk{
		create: agpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppGoodPoster.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (agpcb *AppGoodPosterCreateBulk) OnConflictColumns(columns ...string) *AppGoodPosterUpsertBulk {
	agpcb.conflict = append(agpcb.conflict, sql.ConflictColumns(columns...))
	return &AppGoodPosterUpsertBulk{
		create: agpcb,
	}
}

// AppGoodPosterUpsertBulk is the builder for "upsert"-ing
// a bulk of AppGoodPoster nodes.
type AppGoodPosterUpsertBulk struct {
	create *AppGoodPosterCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppGoodPoster.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appgoodposter.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppGoodPosterUpsertBulk) UpdateNewValues() *AppGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appgoodposter.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppGoodPoster.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppGoodPosterUpsertBulk) Ignore() *AppGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppGoodPosterUpsertBulk) DoNothing() *AppGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppGoodPosterCreateBulk.OnConflict
// documentation for more info.
func (u *AppGoodPosterUpsertBulk) Update(set func(*AppGoodPosterUpsert)) *AppGoodPosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppGoodPosterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppGoodPosterUpsertBulk) SetCreatedAt(v uint32) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppGoodPosterUpsertBulk) AddCreatedAt(v uint32) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsertBulk) UpdateCreatedAt() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppGoodPosterUpsertBulk) SetUpdatedAt(v uint32) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppGoodPosterUpsertBulk) AddUpdatedAt(v uint32) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsertBulk) UpdateUpdatedAt() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppGoodPosterUpsertBulk) SetDeletedAt(v uint32) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppGoodPosterUpsertBulk) AddDeletedAt(v uint32) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppGoodPosterUpsertBulk) UpdateDeletedAt() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppGoodPosterUpsertBulk) SetEntID(v uuid.UUID) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppGoodPosterUpsertBulk) UpdateEntID() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppGoodPosterUpsertBulk) SetAppGoodID(v uuid.UUID) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppGoodPosterUpsertBulk) UpdateAppGoodID() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppGoodPosterUpsertBulk) ClearAppGoodID() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.ClearAppGoodID()
	})
}

// SetPoster sets the "poster" field.
func (u *AppGoodPosterUpsertBulk) SetPoster(v string) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetPoster(v)
	})
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *AppGoodPosterUpsertBulk) UpdatePoster() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdatePoster()
	})
}

// ClearPoster clears the value of the "poster" field.
func (u *AppGoodPosterUpsertBulk) ClearPoster() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.ClearPoster()
	})
}

// SetIndex sets the "index" field.
func (u *AppGoodPosterUpsertBulk) SetIndex(v uint8) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *AppGoodPosterUpsertBulk) AddIndex(v uint8) *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *AppGoodPosterUpsertBulk) UpdateIndex() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *AppGoodPosterUpsertBulk) ClearIndex() *AppGoodPosterUpsertBulk {
	return u.Update(func(s *AppGoodPosterUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *AppGoodPosterUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppGoodPosterCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppGoodPosterCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppGoodPosterUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
