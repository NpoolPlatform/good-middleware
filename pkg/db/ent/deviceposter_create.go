// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/deviceposter"
	"github.com/google/uuid"
)

// DevicePosterCreate is the builder for creating a DevicePoster entity.
type DevicePosterCreate struct {
	config
	mutation *DevicePosterMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dpc *DevicePosterCreate) SetCreatedAt(u uint32) *DevicePosterCreate {
	dpc.mutation.SetCreatedAt(u)
	return dpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dpc *DevicePosterCreate) SetNillableCreatedAt(u *uint32) *DevicePosterCreate {
	if u != nil {
		dpc.SetCreatedAt(*u)
	}
	return dpc
}

// SetUpdatedAt sets the "updated_at" field.
func (dpc *DevicePosterCreate) SetUpdatedAt(u uint32) *DevicePosterCreate {
	dpc.mutation.SetUpdatedAt(u)
	return dpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dpc *DevicePosterCreate) SetNillableUpdatedAt(u *uint32) *DevicePosterCreate {
	if u != nil {
		dpc.SetUpdatedAt(*u)
	}
	return dpc
}

// SetDeletedAt sets the "deleted_at" field.
func (dpc *DevicePosterCreate) SetDeletedAt(u uint32) *DevicePosterCreate {
	dpc.mutation.SetDeletedAt(u)
	return dpc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dpc *DevicePosterCreate) SetNillableDeletedAt(u *uint32) *DevicePosterCreate {
	if u != nil {
		dpc.SetDeletedAt(*u)
	}
	return dpc
}

// SetEntID sets the "ent_id" field.
func (dpc *DevicePosterCreate) SetEntID(u uuid.UUID) *DevicePosterCreate {
	dpc.mutation.SetEntID(u)
	return dpc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (dpc *DevicePosterCreate) SetNillableEntID(u *uuid.UUID) *DevicePosterCreate {
	if u != nil {
		dpc.SetEntID(*u)
	}
	return dpc
}

// SetDeviceTypeID sets the "device_type_id" field.
func (dpc *DevicePosterCreate) SetDeviceTypeID(u uuid.UUID) *DevicePosterCreate {
	dpc.mutation.SetDeviceTypeID(u)
	return dpc
}

// SetNillableDeviceTypeID sets the "device_type_id" field if the given value is not nil.
func (dpc *DevicePosterCreate) SetNillableDeviceTypeID(u *uuid.UUID) *DevicePosterCreate {
	if u != nil {
		dpc.SetDeviceTypeID(*u)
	}
	return dpc
}

// SetPoster sets the "poster" field.
func (dpc *DevicePosterCreate) SetPoster(s string) *DevicePosterCreate {
	dpc.mutation.SetPoster(s)
	return dpc
}

// SetNillablePoster sets the "poster" field if the given value is not nil.
func (dpc *DevicePosterCreate) SetNillablePoster(s *string) *DevicePosterCreate {
	if s != nil {
		dpc.SetPoster(*s)
	}
	return dpc
}

// SetIndex sets the "index" field.
func (dpc *DevicePosterCreate) SetIndex(u uint8) *DevicePosterCreate {
	dpc.mutation.SetIndex(u)
	return dpc
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (dpc *DevicePosterCreate) SetNillableIndex(u *uint8) *DevicePosterCreate {
	if u != nil {
		dpc.SetIndex(*u)
	}
	return dpc
}

// SetID sets the "id" field.
func (dpc *DevicePosterCreate) SetID(u uint32) *DevicePosterCreate {
	dpc.mutation.SetID(u)
	return dpc
}

// Mutation returns the DevicePosterMutation object of the builder.
func (dpc *DevicePosterCreate) Mutation() *DevicePosterMutation {
	return dpc.mutation
}

// Save creates the DevicePoster in the database.
func (dpc *DevicePosterCreate) Save(ctx context.Context) (*DevicePoster, error) {
	var (
		err  error
		node *DevicePoster
	)
	if err := dpc.defaults(); err != nil {
		return nil, err
	}
	if len(dpc.hooks) == 0 {
		if err = dpc.check(); err != nil {
			return nil, err
		}
		node, err = dpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DevicePosterMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dpc.check(); err != nil {
				return nil, err
			}
			dpc.mutation = mutation
			if node, err = dpc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dpc.hooks) - 1; i >= 0; i-- {
			if dpc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dpc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dpc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DevicePoster)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DevicePosterMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dpc *DevicePosterCreate) SaveX(ctx context.Context) *DevicePoster {
	v, err := dpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dpc *DevicePosterCreate) Exec(ctx context.Context) error {
	_, err := dpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpc *DevicePosterCreate) ExecX(ctx context.Context) {
	if err := dpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dpc *DevicePosterCreate) defaults() error {
	if _, ok := dpc.mutation.CreatedAt(); !ok {
		if deviceposter.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceposter.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := deviceposter.DefaultCreatedAt()
		dpc.mutation.SetCreatedAt(v)
	}
	if _, ok := dpc.mutation.UpdatedAt(); !ok {
		if deviceposter.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceposter.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := deviceposter.DefaultUpdatedAt()
		dpc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dpc.mutation.DeletedAt(); !ok {
		if deviceposter.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized deviceposter.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := deviceposter.DefaultDeletedAt()
		dpc.mutation.SetDeletedAt(v)
	}
	if _, ok := dpc.mutation.EntID(); !ok {
		if deviceposter.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized deviceposter.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := deviceposter.DefaultEntID()
		dpc.mutation.SetEntID(v)
	}
	if _, ok := dpc.mutation.DeviceTypeID(); !ok {
		if deviceposter.DefaultDeviceTypeID == nil {
			return fmt.Errorf("ent: uninitialized deviceposter.DefaultDeviceTypeID (forgotten import ent/runtime?)")
		}
		v := deviceposter.DefaultDeviceTypeID()
		dpc.mutation.SetDeviceTypeID(v)
	}
	if _, ok := dpc.mutation.Poster(); !ok {
		v := deviceposter.DefaultPoster
		dpc.mutation.SetPoster(v)
	}
	if _, ok := dpc.mutation.Index(); !ok {
		v := deviceposter.DefaultIndex
		dpc.mutation.SetIndex(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dpc *DevicePosterCreate) check() error {
	if _, ok := dpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DevicePoster.created_at"`)}
	}
	if _, ok := dpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DevicePoster.updated_at"`)}
	}
	if _, ok := dpc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "DevicePoster.deleted_at"`)}
	}
	if _, ok := dpc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "DevicePoster.ent_id"`)}
	}
	return nil
}

func (dpc *DevicePosterCreate) sqlSave(ctx context.Context) (*DevicePoster, error) {
	_node, _spec := dpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dpc.driver, _spec); err != nil {
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

func (dpc *DevicePosterCreate) createSpec() (*DevicePoster, *sqlgraph.CreateSpec) {
	var (
		_node = &DevicePoster{config: dpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: deviceposter.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: deviceposter.FieldID,
			},
		}
	)
	_spec.OnConflict = dpc.conflict
	if id, ok := dpc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dpc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceposter.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dpc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceposter.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dpc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: deviceposter.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := dpc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deviceposter.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := dpc.mutation.DeviceTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: deviceposter.FieldDeviceTypeID,
		})
		_node.DeviceTypeID = value
	}
	if value, ok := dpc.mutation.Poster(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: deviceposter.FieldPoster,
		})
		_node.Poster = value
	}
	if value, ok := dpc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: deviceposter.FieldIndex,
		})
		_node.Index = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DevicePoster.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DevicePosterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dpc *DevicePosterCreate) OnConflict(opts ...sql.ConflictOption) *DevicePosterUpsertOne {
	dpc.conflict = opts
	return &DevicePosterUpsertOne{
		create: dpc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DevicePoster.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dpc *DevicePosterCreate) OnConflictColumns(columns ...string) *DevicePosterUpsertOne {
	dpc.conflict = append(dpc.conflict, sql.ConflictColumns(columns...))
	return &DevicePosterUpsertOne{
		create: dpc,
	}
}

type (
	// DevicePosterUpsertOne is the builder for "upsert"-ing
	//  one DevicePoster node.
	DevicePosterUpsertOne struct {
		create *DevicePosterCreate
	}

	// DevicePosterUpsert is the "OnConflict" setter.
	DevicePosterUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *DevicePosterUpsert) SetCreatedAt(v uint32) *DevicePosterUpsert {
	u.Set(deviceposter.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DevicePosterUpsert) UpdateCreatedAt() *DevicePosterUpsert {
	u.SetExcluded(deviceposter.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DevicePosterUpsert) AddCreatedAt(v uint32) *DevicePosterUpsert {
	u.Add(deviceposter.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DevicePosterUpsert) SetUpdatedAt(v uint32) *DevicePosterUpsert {
	u.Set(deviceposter.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DevicePosterUpsert) UpdateUpdatedAt() *DevicePosterUpsert {
	u.SetExcluded(deviceposter.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DevicePosterUpsert) AddUpdatedAt(v uint32) *DevicePosterUpsert {
	u.Add(deviceposter.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DevicePosterUpsert) SetDeletedAt(v uint32) *DevicePosterUpsert {
	u.Set(deviceposter.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DevicePosterUpsert) UpdateDeletedAt() *DevicePosterUpsert {
	u.SetExcluded(deviceposter.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DevicePosterUpsert) AddDeletedAt(v uint32) *DevicePosterUpsert {
	u.Add(deviceposter.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *DevicePosterUpsert) SetEntID(v uuid.UUID) *DevicePosterUpsert {
	u.Set(deviceposter.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DevicePosterUpsert) UpdateEntID() *DevicePosterUpsert {
	u.SetExcluded(deviceposter.FieldEntID)
	return u
}

// SetDeviceTypeID sets the "device_type_id" field.
func (u *DevicePosterUpsert) SetDeviceTypeID(v uuid.UUID) *DevicePosterUpsert {
	u.Set(deviceposter.FieldDeviceTypeID, v)
	return u
}

// UpdateDeviceTypeID sets the "device_type_id" field to the value that was provided on create.
func (u *DevicePosterUpsert) UpdateDeviceTypeID() *DevicePosterUpsert {
	u.SetExcluded(deviceposter.FieldDeviceTypeID)
	return u
}

// ClearDeviceTypeID clears the value of the "device_type_id" field.
func (u *DevicePosterUpsert) ClearDeviceTypeID() *DevicePosterUpsert {
	u.SetNull(deviceposter.FieldDeviceTypeID)
	return u
}

// SetPoster sets the "poster" field.
func (u *DevicePosterUpsert) SetPoster(v string) *DevicePosterUpsert {
	u.Set(deviceposter.FieldPoster, v)
	return u
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *DevicePosterUpsert) UpdatePoster() *DevicePosterUpsert {
	u.SetExcluded(deviceposter.FieldPoster)
	return u
}

// ClearPoster clears the value of the "poster" field.
func (u *DevicePosterUpsert) ClearPoster() *DevicePosterUpsert {
	u.SetNull(deviceposter.FieldPoster)
	return u
}

// SetIndex sets the "index" field.
func (u *DevicePosterUpsert) SetIndex(v uint8) *DevicePosterUpsert {
	u.Set(deviceposter.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *DevicePosterUpsert) UpdateIndex() *DevicePosterUpsert {
	u.SetExcluded(deviceposter.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *DevicePosterUpsert) AddIndex(v uint8) *DevicePosterUpsert {
	u.Add(deviceposter.FieldIndex, v)
	return u
}

// ClearIndex clears the value of the "index" field.
func (u *DevicePosterUpsert) ClearIndex() *DevicePosterUpsert {
	u.SetNull(deviceposter.FieldIndex)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DevicePoster.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceposter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DevicePosterUpsertOne) UpdateNewValues() *DevicePosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(deviceposter.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DevicePoster.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DevicePosterUpsertOne) Ignore() *DevicePosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DevicePosterUpsertOne) DoNothing() *DevicePosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DevicePosterCreate.OnConflict
// documentation for more info.
func (u *DevicePosterUpsertOne) Update(set func(*DevicePosterUpsert)) *DevicePosterUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DevicePosterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DevicePosterUpsertOne) SetCreatedAt(v uint32) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DevicePosterUpsertOne) AddCreatedAt(v uint32) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DevicePosterUpsertOne) UpdateCreatedAt() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DevicePosterUpsertOne) SetUpdatedAt(v uint32) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DevicePosterUpsertOne) AddUpdatedAt(v uint32) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DevicePosterUpsertOne) UpdateUpdatedAt() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DevicePosterUpsertOne) SetDeletedAt(v uint32) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DevicePosterUpsertOne) AddDeletedAt(v uint32) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DevicePosterUpsertOne) UpdateDeletedAt() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *DevicePosterUpsertOne) SetEntID(v uuid.UUID) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DevicePosterUpsertOne) UpdateEntID() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateEntID()
	})
}

// SetDeviceTypeID sets the "device_type_id" field.
func (u *DevicePosterUpsertOne) SetDeviceTypeID(v uuid.UUID) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetDeviceTypeID(v)
	})
}

// UpdateDeviceTypeID sets the "device_type_id" field to the value that was provided on create.
func (u *DevicePosterUpsertOne) UpdateDeviceTypeID() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateDeviceTypeID()
	})
}

// ClearDeviceTypeID clears the value of the "device_type_id" field.
func (u *DevicePosterUpsertOne) ClearDeviceTypeID() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.ClearDeviceTypeID()
	})
}

// SetPoster sets the "poster" field.
func (u *DevicePosterUpsertOne) SetPoster(v string) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetPoster(v)
	})
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *DevicePosterUpsertOne) UpdatePoster() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdatePoster()
	})
}

// ClearPoster clears the value of the "poster" field.
func (u *DevicePosterUpsertOne) ClearPoster() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.ClearPoster()
	})
}

// SetIndex sets the "index" field.
func (u *DevicePosterUpsertOne) SetIndex(v uint8) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *DevicePosterUpsertOne) AddIndex(v uint8) *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *DevicePosterUpsertOne) UpdateIndex() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *DevicePosterUpsertOne) ClearIndex() *DevicePosterUpsertOne {
	return u.Update(func(s *DevicePosterUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *DevicePosterUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DevicePosterCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DevicePosterUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DevicePosterUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DevicePosterUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DevicePosterCreateBulk is the builder for creating many DevicePoster entities in bulk.
type DevicePosterCreateBulk struct {
	config
	builders []*DevicePosterCreate
	conflict []sql.ConflictOption
}

// Save creates the DevicePoster entities in the database.
func (dpcb *DevicePosterCreateBulk) Save(ctx context.Context) ([]*DevicePoster, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dpcb.builders))
	nodes := make([]*DevicePoster, len(dpcb.builders))
	mutators := make([]Mutator, len(dpcb.builders))
	for i := range dpcb.builders {
		func(i int, root context.Context) {
			builder := dpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DevicePosterMutation)
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
					_, err = mutators[i+1].Mutate(root, dpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dpcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dpcb *DevicePosterCreateBulk) SaveX(ctx context.Context) []*DevicePoster {
	v, err := dpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dpcb *DevicePosterCreateBulk) Exec(ctx context.Context) error {
	_, err := dpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dpcb *DevicePosterCreateBulk) ExecX(ctx context.Context) {
	if err := dpcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DevicePoster.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DevicePosterUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dpcb *DevicePosterCreateBulk) OnConflict(opts ...sql.ConflictOption) *DevicePosterUpsertBulk {
	dpcb.conflict = opts
	return &DevicePosterUpsertBulk{
		create: dpcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DevicePoster.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dpcb *DevicePosterCreateBulk) OnConflictColumns(columns ...string) *DevicePosterUpsertBulk {
	dpcb.conflict = append(dpcb.conflict, sql.ConflictColumns(columns...))
	return &DevicePosterUpsertBulk{
		create: dpcb,
	}
}

// DevicePosterUpsertBulk is the builder for "upsert"-ing
// a bulk of DevicePoster nodes.
type DevicePosterUpsertBulk struct {
	create *DevicePosterCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DevicePoster.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(deviceposter.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DevicePosterUpsertBulk) UpdateNewValues() *DevicePosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(deviceposter.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DevicePoster.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DevicePosterUpsertBulk) Ignore() *DevicePosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DevicePosterUpsertBulk) DoNothing() *DevicePosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DevicePosterCreateBulk.OnConflict
// documentation for more info.
func (u *DevicePosterUpsertBulk) Update(set func(*DevicePosterUpsert)) *DevicePosterUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DevicePosterUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DevicePosterUpsertBulk) SetCreatedAt(v uint32) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DevicePosterUpsertBulk) AddCreatedAt(v uint32) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DevicePosterUpsertBulk) UpdateCreatedAt() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DevicePosterUpsertBulk) SetUpdatedAt(v uint32) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DevicePosterUpsertBulk) AddUpdatedAt(v uint32) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DevicePosterUpsertBulk) UpdateUpdatedAt() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DevicePosterUpsertBulk) SetDeletedAt(v uint32) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DevicePosterUpsertBulk) AddDeletedAt(v uint32) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DevicePosterUpsertBulk) UpdateDeletedAt() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *DevicePosterUpsertBulk) SetEntID(v uuid.UUID) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DevicePosterUpsertBulk) UpdateEntID() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateEntID()
	})
}

// SetDeviceTypeID sets the "device_type_id" field.
func (u *DevicePosterUpsertBulk) SetDeviceTypeID(v uuid.UUID) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetDeviceTypeID(v)
	})
}

// UpdateDeviceTypeID sets the "device_type_id" field to the value that was provided on create.
func (u *DevicePosterUpsertBulk) UpdateDeviceTypeID() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateDeviceTypeID()
	})
}

// ClearDeviceTypeID clears the value of the "device_type_id" field.
func (u *DevicePosterUpsertBulk) ClearDeviceTypeID() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.ClearDeviceTypeID()
	})
}

// SetPoster sets the "poster" field.
func (u *DevicePosterUpsertBulk) SetPoster(v string) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetPoster(v)
	})
}

// UpdatePoster sets the "poster" field to the value that was provided on create.
func (u *DevicePosterUpsertBulk) UpdatePoster() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdatePoster()
	})
}

// ClearPoster clears the value of the "poster" field.
func (u *DevicePosterUpsertBulk) ClearPoster() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.ClearPoster()
	})
}

// SetIndex sets the "index" field.
func (u *DevicePosterUpsertBulk) SetIndex(v uint8) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *DevicePosterUpsertBulk) AddIndex(v uint8) *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *DevicePosterUpsertBulk) UpdateIndex() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *DevicePosterUpsertBulk) ClearIndex() *DevicePosterUpsertBulk {
	return u.Update(func(s *DevicePosterUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *DevicePosterUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DevicePosterCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DevicePosterCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DevicePosterUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
