// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplaycolor"
	"github.com/google/uuid"
)

// AppGoodDisplayColorCreate is the builder for creating a AppGoodDisplayColor entity.
type AppGoodDisplayColorCreate struct {
	config
	mutation *AppGoodDisplayColorMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (agdcc *AppGoodDisplayColorCreate) SetCreatedAt(u uint32) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetCreatedAt(u)
	return agdcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (agdcc *AppGoodDisplayColorCreate) SetNillableCreatedAt(u *uint32) *AppGoodDisplayColorCreate {
	if u != nil {
		agdcc.SetCreatedAt(*u)
	}
	return agdcc
}

// SetUpdatedAt sets the "updated_at" field.
func (agdcc *AppGoodDisplayColorCreate) SetUpdatedAt(u uint32) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetUpdatedAt(u)
	return agdcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (agdcc *AppGoodDisplayColorCreate) SetNillableUpdatedAt(u *uint32) *AppGoodDisplayColorCreate {
	if u != nil {
		agdcc.SetUpdatedAt(*u)
	}
	return agdcc
}

// SetDeletedAt sets the "deleted_at" field.
func (agdcc *AppGoodDisplayColorCreate) SetDeletedAt(u uint32) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetDeletedAt(u)
	return agdcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (agdcc *AppGoodDisplayColorCreate) SetNillableDeletedAt(u *uint32) *AppGoodDisplayColorCreate {
	if u != nil {
		agdcc.SetDeletedAt(*u)
	}
	return agdcc
}

// SetEntID sets the "ent_id" field.
func (agdcc *AppGoodDisplayColorCreate) SetEntID(u uuid.UUID) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetEntID(u)
	return agdcc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (agdcc *AppGoodDisplayColorCreate) SetNillableEntID(u *uuid.UUID) *AppGoodDisplayColorCreate {
	if u != nil {
		agdcc.SetEntID(*u)
	}
	return agdcc
}

// SetAppGoodID sets the "app_good_id" field.
func (agdcc *AppGoodDisplayColorCreate) SetAppGoodID(u uuid.UUID) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetAppGoodID(u)
	return agdcc
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (agdcc *AppGoodDisplayColorCreate) SetNillableAppGoodID(u *uuid.UUID) *AppGoodDisplayColorCreate {
	if u != nil {
		agdcc.SetAppGoodID(*u)
	}
	return agdcc
}

// SetColor sets the "color" field.
func (agdcc *AppGoodDisplayColorCreate) SetColor(s string) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetColor(s)
	return agdcc
}

// SetNillableColor sets the "color" field if the given value is not nil.
func (agdcc *AppGoodDisplayColorCreate) SetNillableColor(s *string) *AppGoodDisplayColorCreate {
	if s != nil {
		agdcc.SetColor(*s)
	}
	return agdcc
}

// SetIndex sets the "index" field.
func (agdcc *AppGoodDisplayColorCreate) SetIndex(u uint8) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetIndex(u)
	return agdcc
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (agdcc *AppGoodDisplayColorCreate) SetNillableIndex(u *uint8) *AppGoodDisplayColorCreate {
	if u != nil {
		agdcc.SetIndex(*u)
	}
	return agdcc
}

// SetID sets the "id" field.
func (agdcc *AppGoodDisplayColorCreate) SetID(u uint32) *AppGoodDisplayColorCreate {
	agdcc.mutation.SetID(u)
	return agdcc
}

// Mutation returns the AppGoodDisplayColorMutation object of the builder.
func (agdcc *AppGoodDisplayColorCreate) Mutation() *AppGoodDisplayColorMutation {
	return agdcc.mutation
}

// Save creates the AppGoodDisplayColor in the database.
func (agdcc *AppGoodDisplayColorCreate) Save(ctx context.Context) (*AppGoodDisplayColor, error) {
	var (
		err  error
		node *AppGoodDisplayColor
	)
	if err := agdcc.defaults(); err != nil {
		return nil, err
	}
	if len(agdcc.hooks) == 0 {
		if err = agdcc.check(); err != nil {
			return nil, err
		}
		node, err = agdcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppGoodDisplayColorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = agdcc.check(); err != nil {
				return nil, err
			}
			agdcc.mutation = mutation
			if node, err = agdcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(agdcc.hooks) - 1; i >= 0; i-- {
			if agdcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = agdcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, agdcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppGoodDisplayColor)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppGoodDisplayColorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (agdcc *AppGoodDisplayColorCreate) SaveX(ctx context.Context) *AppGoodDisplayColor {
	v, err := agdcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agdcc *AppGoodDisplayColorCreate) Exec(ctx context.Context) error {
	_, err := agdcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agdcc *AppGoodDisplayColorCreate) ExecX(ctx context.Context) {
	if err := agdcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (agdcc *AppGoodDisplayColorCreate) defaults() error {
	if _, ok := agdcc.mutation.CreatedAt(); !ok {
		if appgooddisplaycolor.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appgooddisplaycolor.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appgooddisplaycolor.DefaultCreatedAt()
		agdcc.mutation.SetCreatedAt(v)
	}
	if _, ok := agdcc.mutation.UpdatedAt(); !ok {
		if appgooddisplaycolor.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appgooddisplaycolor.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appgooddisplaycolor.DefaultUpdatedAt()
		agdcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := agdcc.mutation.DeletedAt(); !ok {
		if appgooddisplaycolor.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appgooddisplaycolor.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appgooddisplaycolor.DefaultDeletedAt()
		agdcc.mutation.SetDeletedAt(v)
	}
	if _, ok := agdcc.mutation.EntID(); !ok {
		if appgooddisplaycolor.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized appgooddisplaycolor.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := appgooddisplaycolor.DefaultEntID()
		agdcc.mutation.SetEntID(v)
	}
	if _, ok := agdcc.mutation.AppGoodID(); !ok {
		if appgooddisplaycolor.DefaultAppGoodID == nil {
			return fmt.Errorf("ent: uninitialized appgooddisplaycolor.DefaultAppGoodID (forgotten import ent/runtime?)")
		}
		v := appgooddisplaycolor.DefaultAppGoodID()
		agdcc.mutation.SetAppGoodID(v)
	}
	if _, ok := agdcc.mutation.Color(); !ok {
		v := appgooddisplaycolor.DefaultColor
		agdcc.mutation.SetColor(v)
	}
	if _, ok := agdcc.mutation.Index(); !ok {
		v := appgooddisplaycolor.DefaultIndex
		agdcc.mutation.SetIndex(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (agdcc *AppGoodDisplayColorCreate) check() error {
	if _, ok := agdcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppGoodDisplayColor.created_at"`)}
	}
	if _, ok := agdcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppGoodDisplayColor.updated_at"`)}
	}
	if _, ok := agdcc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppGoodDisplayColor.deleted_at"`)}
	}
	if _, ok := agdcc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "AppGoodDisplayColor.ent_id"`)}
	}
	return nil
}

func (agdcc *AppGoodDisplayColorCreate) sqlSave(ctx context.Context) (*AppGoodDisplayColor, error) {
	_node, _spec := agdcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, agdcc.driver, _spec); err != nil {
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

func (agdcc *AppGoodDisplayColorCreate) createSpec() (*AppGoodDisplayColor, *sqlgraph.CreateSpec) {
	var (
		_node = &AppGoodDisplayColor{config: agdcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appgooddisplaycolor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgooddisplaycolor.FieldID,
			},
		}
	)
	_spec.OnConflict = agdcc.conflict
	if id, ok := agdcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := agdcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddisplaycolor.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := agdcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddisplaycolor.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := agdcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appgooddisplaycolor.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := agdcc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgooddisplaycolor.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := agdcc.mutation.AppGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appgooddisplaycolor.FieldAppGoodID,
		})
		_node.AppGoodID = value
	}
	if value, ok := agdcc.mutation.Color(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appgooddisplaycolor.FieldColor,
		})
		_node.Color = value
	}
	if value, ok := agdcc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: appgooddisplaycolor.FieldIndex,
		})
		_node.Index = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppGoodDisplayColor.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppGoodDisplayColorUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (agdcc *AppGoodDisplayColorCreate) OnConflict(opts ...sql.ConflictOption) *AppGoodDisplayColorUpsertOne {
	agdcc.conflict = opts
	return &AppGoodDisplayColorUpsertOne{
		create: agdcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppGoodDisplayColor.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (agdcc *AppGoodDisplayColorCreate) OnConflictColumns(columns ...string) *AppGoodDisplayColorUpsertOne {
	agdcc.conflict = append(agdcc.conflict, sql.ConflictColumns(columns...))
	return &AppGoodDisplayColorUpsertOne{
		create: agdcc,
	}
}

type (
	// AppGoodDisplayColorUpsertOne is the builder for "upsert"-ing
	//  one AppGoodDisplayColor node.
	AppGoodDisplayColorUpsertOne struct {
		create *AppGoodDisplayColorCreate
	}

	// AppGoodDisplayColorUpsert is the "OnConflict" setter.
	AppGoodDisplayColorUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppGoodDisplayColorUpsert) SetCreatedAt(v uint32) *AppGoodDisplayColorUpsert {
	u.Set(appgooddisplaycolor.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsert) UpdateCreatedAt() *AppGoodDisplayColorUpsert {
	u.SetExcluded(appgooddisplaycolor.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppGoodDisplayColorUpsert) AddCreatedAt(v uint32) *AppGoodDisplayColorUpsert {
	u.Add(appgooddisplaycolor.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppGoodDisplayColorUpsert) SetUpdatedAt(v uint32) *AppGoodDisplayColorUpsert {
	u.Set(appgooddisplaycolor.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsert) UpdateUpdatedAt() *AppGoodDisplayColorUpsert {
	u.SetExcluded(appgooddisplaycolor.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppGoodDisplayColorUpsert) AddUpdatedAt(v uint32) *AppGoodDisplayColorUpsert {
	u.Add(appgooddisplaycolor.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppGoodDisplayColorUpsert) SetDeletedAt(v uint32) *AppGoodDisplayColorUpsert {
	u.Set(appgooddisplaycolor.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsert) UpdateDeletedAt() *AppGoodDisplayColorUpsert {
	u.SetExcluded(appgooddisplaycolor.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppGoodDisplayColorUpsert) AddDeletedAt(v uint32) *AppGoodDisplayColorUpsert {
	u.Add(appgooddisplaycolor.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppGoodDisplayColorUpsert) SetEntID(v uuid.UUID) *AppGoodDisplayColorUpsert {
	u.Set(appgooddisplaycolor.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsert) UpdateEntID() *AppGoodDisplayColorUpsert {
	u.SetExcluded(appgooddisplaycolor.FieldEntID)
	return u
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppGoodDisplayColorUpsert) SetAppGoodID(v uuid.UUID) *AppGoodDisplayColorUpsert {
	u.Set(appgooddisplaycolor.FieldAppGoodID, v)
	return u
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsert) UpdateAppGoodID() *AppGoodDisplayColorUpsert {
	u.SetExcluded(appgooddisplaycolor.FieldAppGoodID)
	return u
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppGoodDisplayColorUpsert) ClearAppGoodID() *AppGoodDisplayColorUpsert {
	u.SetNull(appgooddisplaycolor.FieldAppGoodID)
	return u
}

// SetColor sets the "color" field.
func (u *AppGoodDisplayColorUpsert) SetColor(v string) *AppGoodDisplayColorUpsert {
	u.Set(appgooddisplaycolor.FieldColor, v)
	return u
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsert) UpdateColor() *AppGoodDisplayColorUpsert {
	u.SetExcluded(appgooddisplaycolor.FieldColor)
	return u
}

// ClearColor clears the value of the "color" field.
func (u *AppGoodDisplayColorUpsert) ClearColor() *AppGoodDisplayColorUpsert {
	u.SetNull(appgooddisplaycolor.FieldColor)
	return u
}

// SetIndex sets the "index" field.
func (u *AppGoodDisplayColorUpsert) SetIndex(v uint8) *AppGoodDisplayColorUpsert {
	u.Set(appgooddisplaycolor.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsert) UpdateIndex() *AppGoodDisplayColorUpsert {
	u.SetExcluded(appgooddisplaycolor.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *AppGoodDisplayColorUpsert) AddIndex(v uint8) *AppGoodDisplayColorUpsert {
	u.Add(appgooddisplaycolor.FieldIndex, v)
	return u
}

// ClearIndex clears the value of the "index" field.
func (u *AppGoodDisplayColorUpsert) ClearIndex() *AppGoodDisplayColorUpsert {
	u.SetNull(appgooddisplaycolor.FieldIndex)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppGoodDisplayColor.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appgooddisplaycolor.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppGoodDisplayColorUpsertOne) UpdateNewValues() *AppGoodDisplayColorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appgooddisplaycolor.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppGoodDisplayColor.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppGoodDisplayColorUpsertOne) Ignore() *AppGoodDisplayColorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppGoodDisplayColorUpsertOne) DoNothing() *AppGoodDisplayColorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppGoodDisplayColorCreate.OnConflict
// documentation for more info.
func (u *AppGoodDisplayColorUpsertOne) Update(set func(*AppGoodDisplayColorUpsert)) *AppGoodDisplayColorUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppGoodDisplayColorUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppGoodDisplayColorUpsertOne) SetCreatedAt(v uint32) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppGoodDisplayColorUpsertOne) AddCreatedAt(v uint32) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertOne) UpdateCreatedAt() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppGoodDisplayColorUpsertOne) SetUpdatedAt(v uint32) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppGoodDisplayColorUpsertOne) AddUpdatedAt(v uint32) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertOne) UpdateUpdatedAt() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppGoodDisplayColorUpsertOne) SetDeletedAt(v uint32) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppGoodDisplayColorUpsertOne) AddDeletedAt(v uint32) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertOne) UpdateDeletedAt() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppGoodDisplayColorUpsertOne) SetEntID(v uuid.UUID) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertOne) UpdateEntID() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppGoodDisplayColorUpsertOne) SetAppGoodID(v uuid.UUID) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertOne) UpdateAppGoodID() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppGoodDisplayColorUpsertOne) ClearAppGoodID() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.ClearAppGoodID()
	})
}

// SetColor sets the "color" field.
func (u *AppGoodDisplayColorUpsertOne) SetColor(v string) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertOne) UpdateColor() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateColor()
	})
}

// ClearColor clears the value of the "color" field.
func (u *AppGoodDisplayColorUpsertOne) ClearColor() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.ClearColor()
	})
}

// SetIndex sets the "index" field.
func (u *AppGoodDisplayColorUpsertOne) SetIndex(v uint8) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *AppGoodDisplayColorUpsertOne) AddIndex(v uint8) *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertOne) UpdateIndex() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *AppGoodDisplayColorUpsertOne) ClearIndex() *AppGoodDisplayColorUpsertOne {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *AppGoodDisplayColorUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppGoodDisplayColorCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppGoodDisplayColorUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppGoodDisplayColorUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppGoodDisplayColorUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppGoodDisplayColorCreateBulk is the builder for creating many AppGoodDisplayColor entities in bulk.
type AppGoodDisplayColorCreateBulk struct {
	config
	builders []*AppGoodDisplayColorCreate
	conflict []sql.ConflictOption
}

// Save creates the AppGoodDisplayColor entities in the database.
func (agdccb *AppGoodDisplayColorCreateBulk) Save(ctx context.Context) ([]*AppGoodDisplayColor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(agdccb.builders))
	nodes := make([]*AppGoodDisplayColor, len(agdccb.builders))
	mutators := make([]Mutator, len(agdccb.builders))
	for i := range agdccb.builders {
		func(i int, root context.Context) {
			builder := agdccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppGoodDisplayColorMutation)
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
					_, err = mutators[i+1].Mutate(root, agdccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = agdccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, agdccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, agdccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (agdccb *AppGoodDisplayColorCreateBulk) SaveX(ctx context.Context) []*AppGoodDisplayColor {
	v, err := agdccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (agdccb *AppGoodDisplayColorCreateBulk) Exec(ctx context.Context) error {
	_, err := agdccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (agdccb *AppGoodDisplayColorCreateBulk) ExecX(ctx context.Context) {
	if err := agdccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppGoodDisplayColor.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppGoodDisplayColorUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (agdccb *AppGoodDisplayColorCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppGoodDisplayColorUpsertBulk {
	agdccb.conflict = opts
	return &AppGoodDisplayColorUpsertBulk{
		create: agdccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppGoodDisplayColor.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (agdccb *AppGoodDisplayColorCreateBulk) OnConflictColumns(columns ...string) *AppGoodDisplayColorUpsertBulk {
	agdccb.conflict = append(agdccb.conflict, sql.ConflictColumns(columns...))
	return &AppGoodDisplayColorUpsertBulk{
		create: agdccb,
	}
}

// AppGoodDisplayColorUpsertBulk is the builder for "upsert"-ing
// a bulk of AppGoodDisplayColor nodes.
type AppGoodDisplayColorUpsertBulk struct {
	create *AppGoodDisplayColorCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppGoodDisplayColor.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appgooddisplaycolor.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppGoodDisplayColorUpsertBulk) UpdateNewValues() *AppGoodDisplayColorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appgooddisplaycolor.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppGoodDisplayColor.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppGoodDisplayColorUpsertBulk) Ignore() *AppGoodDisplayColorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppGoodDisplayColorUpsertBulk) DoNothing() *AppGoodDisplayColorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppGoodDisplayColorCreateBulk.OnConflict
// documentation for more info.
func (u *AppGoodDisplayColorUpsertBulk) Update(set func(*AppGoodDisplayColorUpsert)) *AppGoodDisplayColorUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppGoodDisplayColorUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppGoodDisplayColorUpsertBulk) SetCreatedAt(v uint32) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppGoodDisplayColorUpsertBulk) AddCreatedAt(v uint32) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertBulk) UpdateCreatedAt() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppGoodDisplayColorUpsertBulk) SetUpdatedAt(v uint32) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppGoodDisplayColorUpsertBulk) AddUpdatedAt(v uint32) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertBulk) UpdateUpdatedAt() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppGoodDisplayColorUpsertBulk) SetDeletedAt(v uint32) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppGoodDisplayColorUpsertBulk) AddDeletedAt(v uint32) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertBulk) UpdateDeletedAt() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppGoodDisplayColorUpsertBulk) SetEntID(v uuid.UUID) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertBulk) UpdateEntID() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppGoodDisplayColorUpsertBulk) SetAppGoodID(v uuid.UUID) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertBulk) UpdateAppGoodID() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppGoodDisplayColorUpsertBulk) ClearAppGoodID() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.ClearAppGoodID()
	})
}

// SetColor sets the "color" field.
func (u *AppGoodDisplayColorUpsertBulk) SetColor(v string) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetColor(v)
	})
}

// UpdateColor sets the "color" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertBulk) UpdateColor() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateColor()
	})
}

// ClearColor clears the value of the "color" field.
func (u *AppGoodDisplayColorUpsertBulk) ClearColor() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.ClearColor()
	})
}

// SetIndex sets the "index" field.
func (u *AppGoodDisplayColorUpsertBulk) SetIndex(v uint8) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *AppGoodDisplayColorUpsertBulk) AddIndex(v uint8) *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *AppGoodDisplayColorUpsertBulk) UpdateIndex() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *AppGoodDisplayColorUpsertBulk) ClearIndex() *AppGoodDisplayColorUpsertBulk {
	return u.Update(func(s *AppGoodDisplayColorUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *AppGoodDisplayColorUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppGoodDisplayColorCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppGoodDisplayColorCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppGoodDisplayColorUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}