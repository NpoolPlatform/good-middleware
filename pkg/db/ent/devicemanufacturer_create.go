// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	"github.com/google/uuid"
)

// DeviceManufacturerCreate is the builder for creating a DeviceManufacturer entity.
type DeviceManufacturerCreate struct {
	config
	mutation *DeviceManufacturerMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (dmc *DeviceManufacturerCreate) SetCreatedAt(u uint32) *DeviceManufacturerCreate {
	dmc.mutation.SetCreatedAt(u)
	return dmc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dmc *DeviceManufacturerCreate) SetNillableCreatedAt(u *uint32) *DeviceManufacturerCreate {
	if u != nil {
		dmc.SetCreatedAt(*u)
	}
	return dmc
}

// SetUpdatedAt sets the "updated_at" field.
func (dmc *DeviceManufacturerCreate) SetUpdatedAt(u uint32) *DeviceManufacturerCreate {
	dmc.mutation.SetUpdatedAt(u)
	return dmc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (dmc *DeviceManufacturerCreate) SetNillableUpdatedAt(u *uint32) *DeviceManufacturerCreate {
	if u != nil {
		dmc.SetUpdatedAt(*u)
	}
	return dmc
}

// SetDeletedAt sets the "deleted_at" field.
func (dmc *DeviceManufacturerCreate) SetDeletedAt(u uint32) *DeviceManufacturerCreate {
	dmc.mutation.SetDeletedAt(u)
	return dmc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (dmc *DeviceManufacturerCreate) SetNillableDeletedAt(u *uint32) *DeviceManufacturerCreate {
	if u != nil {
		dmc.SetDeletedAt(*u)
	}
	return dmc
}

// SetEntID sets the "ent_id" field.
func (dmc *DeviceManufacturerCreate) SetEntID(u uuid.UUID) *DeviceManufacturerCreate {
	dmc.mutation.SetEntID(u)
	return dmc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (dmc *DeviceManufacturerCreate) SetNillableEntID(u *uuid.UUID) *DeviceManufacturerCreate {
	if u != nil {
		dmc.SetEntID(*u)
	}
	return dmc
}

// SetName sets the "name" field.
func (dmc *DeviceManufacturerCreate) SetName(s string) *DeviceManufacturerCreate {
	dmc.mutation.SetName(s)
	return dmc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (dmc *DeviceManufacturerCreate) SetNillableName(s *string) *DeviceManufacturerCreate {
	if s != nil {
		dmc.SetName(*s)
	}
	return dmc
}

// SetLogo sets the "logo" field.
func (dmc *DeviceManufacturerCreate) SetLogo(s string) *DeviceManufacturerCreate {
	dmc.mutation.SetLogo(s)
	return dmc
}

// SetNillableLogo sets the "logo" field if the given value is not nil.
func (dmc *DeviceManufacturerCreate) SetNillableLogo(s *string) *DeviceManufacturerCreate {
	if s != nil {
		dmc.SetLogo(*s)
	}
	return dmc
}

// SetID sets the "id" field.
func (dmc *DeviceManufacturerCreate) SetID(u uint32) *DeviceManufacturerCreate {
	dmc.mutation.SetID(u)
	return dmc
}

// Mutation returns the DeviceManufacturerMutation object of the builder.
func (dmc *DeviceManufacturerCreate) Mutation() *DeviceManufacturerMutation {
	return dmc.mutation
}

// Save creates the DeviceManufacturer in the database.
func (dmc *DeviceManufacturerCreate) Save(ctx context.Context) (*DeviceManufacturer, error) {
	var (
		err  error
		node *DeviceManufacturer
	)
	if err := dmc.defaults(); err != nil {
		return nil, err
	}
	if len(dmc.hooks) == 0 {
		if err = dmc.check(); err != nil {
			return nil, err
		}
		node, err = dmc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceManufacturerMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = dmc.check(); err != nil {
				return nil, err
			}
			dmc.mutation = mutation
			if node, err = dmc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(dmc.hooks) - 1; i >= 0; i-- {
			if dmc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = dmc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, dmc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*DeviceManufacturer)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeviceManufacturerMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dmc *DeviceManufacturerCreate) SaveX(ctx context.Context) *DeviceManufacturer {
	v, err := dmc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmc *DeviceManufacturerCreate) Exec(ctx context.Context) error {
	_, err := dmc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmc *DeviceManufacturerCreate) ExecX(ctx context.Context) {
	if err := dmc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dmc *DeviceManufacturerCreate) defaults() error {
	if _, ok := dmc.mutation.CreatedAt(); !ok {
		if devicemanufacturer.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized devicemanufacturer.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := devicemanufacturer.DefaultCreatedAt()
		dmc.mutation.SetCreatedAt(v)
	}
	if _, ok := dmc.mutation.UpdatedAt(); !ok {
		if devicemanufacturer.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized devicemanufacturer.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := devicemanufacturer.DefaultUpdatedAt()
		dmc.mutation.SetUpdatedAt(v)
	}
	if _, ok := dmc.mutation.DeletedAt(); !ok {
		if devicemanufacturer.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized devicemanufacturer.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := devicemanufacturer.DefaultDeletedAt()
		dmc.mutation.SetDeletedAt(v)
	}
	if _, ok := dmc.mutation.EntID(); !ok {
		if devicemanufacturer.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized devicemanufacturer.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := devicemanufacturer.DefaultEntID()
		dmc.mutation.SetEntID(v)
	}
	if _, ok := dmc.mutation.Name(); !ok {
		v := devicemanufacturer.DefaultName
		dmc.mutation.SetName(v)
	}
	if _, ok := dmc.mutation.Logo(); !ok {
		v := devicemanufacturer.DefaultLogo
		dmc.mutation.SetLogo(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (dmc *DeviceManufacturerCreate) check() error {
	if _, ok := dmc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DeviceManufacturer.created_at"`)}
	}
	if _, ok := dmc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DeviceManufacturer.updated_at"`)}
	}
	if _, ok := dmc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "DeviceManufacturer.deleted_at"`)}
	}
	if _, ok := dmc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "DeviceManufacturer.ent_id"`)}
	}
	if v, ok := dmc.mutation.Name(); ok {
		if err := devicemanufacturer.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "DeviceManufacturer.name": %w`, err)}
		}
	}
	if v, ok := dmc.mutation.Logo(); ok {
		if err := devicemanufacturer.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`ent: validator failed for field "DeviceManufacturer.logo": %w`, err)}
		}
	}
	return nil
}

func (dmc *DeviceManufacturerCreate) sqlSave(ctx context.Context) (*DeviceManufacturer, error) {
	_node, _spec := dmc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dmc.driver, _spec); err != nil {
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

func (dmc *DeviceManufacturerCreate) createSpec() (*DeviceManufacturer, *sqlgraph.CreateSpec) {
	var (
		_node = &DeviceManufacturer{config: dmc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: devicemanufacturer.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: devicemanufacturer.FieldID,
			},
		}
	)
	_spec.OnConflict = dmc.conflict
	if id, ok := dmc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := dmc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: devicemanufacturer.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := dmc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: devicemanufacturer.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := dmc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: devicemanufacturer.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := dmc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: devicemanufacturer.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := dmc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: devicemanufacturer.FieldName,
		})
		_node.Name = value
	}
	if value, ok := dmc.mutation.Logo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: devicemanufacturer.FieldLogo,
		})
		_node.Logo = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceManufacturer.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceManufacturerUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dmc *DeviceManufacturerCreate) OnConflict(opts ...sql.ConflictOption) *DeviceManufacturerUpsertOne {
	dmc.conflict = opts
	return &DeviceManufacturerUpsertOne{
		create: dmc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceManufacturer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmc *DeviceManufacturerCreate) OnConflictColumns(columns ...string) *DeviceManufacturerUpsertOne {
	dmc.conflict = append(dmc.conflict, sql.ConflictColumns(columns...))
	return &DeviceManufacturerUpsertOne{
		create: dmc,
	}
}

type (
	// DeviceManufacturerUpsertOne is the builder for "upsert"-ing
	//  one DeviceManufacturer node.
	DeviceManufacturerUpsertOne struct {
		create *DeviceManufacturerCreate
	}

	// DeviceManufacturerUpsert is the "OnConflict" setter.
	DeviceManufacturerUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *DeviceManufacturerUpsert) SetCreatedAt(v uint32) *DeviceManufacturerUpsert {
	u.Set(devicemanufacturer.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsert) UpdateCreatedAt() *DeviceManufacturerUpsert {
	u.SetExcluded(devicemanufacturer.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceManufacturerUpsert) AddCreatedAt(v uint32) *DeviceManufacturerUpsert {
	u.Add(devicemanufacturer.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceManufacturerUpsert) SetUpdatedAt(v uint32) *DeviceManufacturerUpsert {
	u.Set(devicemanufacturer.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsert) UpdateUpdatedAt() *DeviceManufacturerUpsert {
	u.SetExcluded(devicemanufacturer.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceManufacturerUpsert) AddUpdatedAt(v uint32) *DeviceManufacturerUpsert {
	u.Add(devicemanufacturer.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceManufacturerUpsert) SetDeletedAt(v uint32) *DeviceManufacturerUpsert {
	u.Set(devicemanufacturer.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsert) UpdateDeletedAt() *DeviceManufacturerUpsert {
	u.SetExcluded(devicemanufacturer.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceManufacturerUpsert) AddDeletedAt(v uint32) *DeviceManufacturerUpsert {
	u.Add(devicemanufacturer.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *DeviceManufacturerUpsert) SetEntID(v uuid.UUID) *DeviceManufacturerUpsert {
	u.Set(devicemanufacturer.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DeviceManufacturerUpsert) UpdateEntID() *DeviceManufacturerUpsert {
	u.SetExcluded(devicemanufacturer.FieldEntID)
	return u
}

// SetName sets the "name" field.
func (u *DeviceManufacturerUpsert) SetName(v string) *DeviceManufacturerUpsert {
	u.Set(devicemanufacturer.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeviceManufacturerUpsert) UpdateName() *DeviceManufacturerUpsert {
	u.SetExcluded(devicemanufacturer.FieldName)
	return u
}

// ClearName clears the value of the "name" field.
func (u *DeviceManufacturerUpsert) ClearName() *DeviceManufacturerUpsert {
	u.SetNull(devicemanufacturer.FieldName)
	return u
}

// SetLogo sets the "logo" field.
func (u *DeviceManufacturerUpsert) SetLogo(v string) *DeviceManufacturerUpsert {
	u.Set(devicemanufacturer.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *DeviceManufacturerUpsert) UpdateLogo() *DeviceManufacturerUpsert {
	u.SetExcluded(devicemanufacturer.FieldLogo)
	return u
}

// ClearLogo clears the value of the "logo" field.
func (u *DeviceManufacturerUpsert) ClearLogo() *DeviceManufacturerUpsert {
	u.SetNull(devicemanufacturer.FieldLogo)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.DeviceManufacturer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicemanufacturer.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceManufacturerUpsertOne) UpdateNewValues() *DeviceManufacturerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(devicemanufacturer.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceManufacturer.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DeviceManufacturerUpsertOne) Ignore() *DeviceManufacturerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceManufacturerUpsertOne) DoNothing() *DeviceManufacturerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceManufacturerCreate.OnConflict
// documentation for more info.
func (u *DeviceManufacturerUpsertOne) Update(set func(*DeviceManufacturerUpsert)) *DeviceManufacturerUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceManufacturerUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeviceManufacturerUpsertOne) SetCreatedAt(v uint32) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceManufacturerUpsertOne) AddCreatedAt(v uint32) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertOne) UpdateCreatedAt() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceManufacturerUpsertOne) SetUpdatedAt(v uint32) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceManufacturerUpsertOne) AddUpdatedAt(v uint32) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertOne) UpdateUpdatedAt() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceManufacturerUpsertOne) SetDeletedAt(v uint32) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceManufacturerUpsertOne) AddDeletedAt(v uint32) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertOne) UpdateDeletedAt() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *DeviceManufacturerUpsertOne) SetEntID(v uuid.UUID) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertOne) UpdateEntID() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateEntID()
	})
}

// SetName sets the "name" field.
func (u *DeviceManufacturerUpsertOne) SetName(v string) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertOne) UpdateName() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DeviceManufacturerUpsertOne) ClearName() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.ClearName()
	})
}

// SetLogo sets the "logo" field.
func (u *DeviceManufacturerUpsertOne) SetLogo(v string) *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertOne) UpdateLogo() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *DeviceManufacturerUpsertOne) ClearLogo() *DeviceManufacturerUpsertOne {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.ClearLogo()
	})
}

// Exec executes the query.
func (u *DeviceManufacturerUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceManufacturerCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceManufacturerUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DeviceManufacturerUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DeviceManufacturerUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DeviceManufacturerCreateBulk is the builder for creating many DeviceManufacturer entities in bulk.
type DeviceManufacturerCreateBulk struct {
	config
	builders []*DeviceManufacturerCreate
	conflict []sql.ConflictOption
}

// Save creates the DeviceManufacturer entities in the database.
func (dmcb *DeviceManufacturerCreateBulk) Save(ctx context.Context) ([]*DeviceManufacturer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(dmcb.builders))
	nodes := make([]*DeviceManufacturer, len(dmcb.builders))
	mutators := make([]Mutator, len(dmcb.builders))
	for i := range dmcb.builders {
		func(i int, root context.Context) {
			builder := dmcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeviceManufacturerMutation)
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
					_, err = mutators[i+1].Mutate(root, dmcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dmcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dmcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, dmcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dmcb *DeviceManufacturerCreateBulk) SaveX(ctx context.Context) []*DeviceManufacturer {
	v, err := dmcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dmcb *DeviceManufacturerCreateBulk) Exec(ctx context.Context) error {
	_, err := dmcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmcb *DeviceManufacturerCreateBulk) ExecX(ctx context.Context) {
	if err := dmcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.DeviceManufacturer.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DeviceManufacturerUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (dmcb *DeviceManufacturerCreateBulk) OnConflict(opts ...sql.ConflictOption) *DeviceManufacturerUpsertBulk {
	dmcb.conflict = opts
	return &DeviceManufacturerUpsertBulk{
		create: dmcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.DeviceManufacturer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dmcb *DeviceManufacturerCreateBulk) OnConflictColumns(columns ...string) *DeviceManufacturerUpsertBulk {
	dmcb.conflict = append(dmcb.conflict, sql.ConflictColumns(columns...))
	return &DeviceManufacturerUpsertBulk{
		create: dmcb,
	}
}

// DeviceManufacturerUpsertBulk is the builder for "upsert"-ing
// a bulk of DeviceManufacturer nodes.
type DeviceManufacturerUpsertBulk struct {
	create *DeviceManufacturerCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.DeviceManufacturer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(devicemanufacturer.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *DeviceManufacturerUpsertBulk) UpdateNewValues() *DeviceManufacturerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(devicemanufacturer.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.DeviceManufacturer.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DeviceManufacturerUpsertBulk) Ignore() *DeviceManufacturerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DeviceManufacturerUpsertBulk) DoNothing() *DeviceManufacturerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DeviceManufacturerCreateBulk.OnConflict
// documentation for more info.
func (u *DeviceManufacturerUpsertBulk) Update(set func(*DeviceManufacturerUpsert)) *DeviceManufacturerUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DeviceManufacturerUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *DeviceManufacturerUpsertBulk) SetCreatedAt(v uint32) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *DeviceManufacturerUpsertBulk) AddCreatedAt(v uint32) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertBulk) UpdateCreatedAt() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *DeviceManufacturerUpsertBulk) SetUpdatedAt(v uint32) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *DeviceManufacturerUpsertBulk) AddUpdatedAt(v uint32) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertBulk) UpdateUpdatedAt() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *DeviceManufacturerUpsertBulk) SetDeletedAt(v uint32) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *DeviceManufacturerUpsertBulk) AddDeletedAt(v uint32) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertBulk) UpdateDeletedAt() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *DeviceManufacturerUpsertBulk) SetEntID(v uuid.UUID) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertBulk) UpdateEntID() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateEntID()
	})
}

// SetName sets the "name" field.
func (u *DeviceManufacturerUpsertBulk) SetName(v string) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertBulk) UpdateName() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateName()
	})
}

// ClearName clears the value of the "name" field.
func (u *DeviceManufacturerUpsertBulk) ClearName() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.ClearName()
	})
}

// SetLogo sets the "logo" field.
func (u *DeviceManufacturerUpsertBulk) SetLogo(v string) *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *DeviceManufacturerUpsertBulk) UpdateLogo() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.UpdateLogo()
	})
}

// ClearLogo clears the value of the "logo" field.
func (u *DeviceManufacturerUpsertBulk) ClearLogo() *DeviceManufacturerUpsertBulk {
	return u.Update(func(s *DeviceManufacturerUpsert) {
		s.ClearLogo()
	})
}

// Exec executes the query.
func (u *DeviceManufacturerUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DeviceManufacturerCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DeviceManufacturerCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DeviceManufacturerUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
