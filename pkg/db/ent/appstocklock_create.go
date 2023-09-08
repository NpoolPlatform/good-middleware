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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstocklock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppStockLockCreate is the builder for creating a AppStockLock entity.
type AppStockLockCreate struct {
	config
	mutation *AppStockLockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (aslc *AppStockLockCreate) SetCreatedAt(u uint32) *AppStockLockCreate {
	aslc.mutation.SetCreatedAt(u)
	return aslc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (aslc *AppStockLockCreate) SetNillableCreatedAt(u *uint32) *AppStockLockCreate {
	if u != nil {
		aslc.SetCreatedAt(*u)
	}
	return aslc
}

// SetUpdatedAt sets the "updated_at" field.
func (aslc *AppStockLockCreate) SetUpdatedAt(u uint32) *AppStockLockCreate {
	aslc.mutation.SetUpdatedAt(u)
	return aslc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (aslc *AppStockLockCreate) SetNillableUpdatedAt(u *uint32) *AppStockLockCreate {
	if u != nil {
		aslc.SetUpdatedAt(*u)
	}
	return aslc
}

// SetDeletedAt sets the "deleted_at" field.
func (aslc *AppStockLockCreate) SetDeletedAt(u uint32) *AppStockLockCreate {
	aslc.mutation.SetDeletedAt(u)
	return aslc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (aslc *AppStockLockCreate) SetNillableDeletedAt(u *uint32) *AppStockLockCreate {
	if u != nil {
		aslc.SetDeletedAt(*u)
	}
	return aslc
}

// SetUnits sets the "units" field.
func (aslc *AppStockLockCreate) SetUnits(d decimal.Decimal) *AppStockLockCreate {
	aslc.mutation.SetUnits(d)
	return aslc
}

// SetNillableUnits sets the "units" field if the given value is not nil.
func (aslc *AppStockLockCreate) SetNillableUnits(d *decimal.Decimal) *AppStockLockCreate {
	if d != nil {
		aslc.SetUnits(*d)
	}
	return aslc
}

// SetID sets the "id" field.
func (aslc *AppStockLockCreate) SetID(u uuid.UUID) *AppStockLockCreate {
	aslc.mutation.SetID(u)
	return aslc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (aslc *AppStockLockCreate) SetNillableID(u *uuid.UUID) *AppStockLockCreate {
	if u != nil {
		aslc.SetID(*u)
	}
	return aslc
}

// Mutation returns the AppStockLockMutation object of the builder.
func (aslc *AppStockLockCreate) Mutation() *AppStockLockMutation {
	return aslc.mutation
}

// Save creates the AppStockLock in the database.
func (aslc *AppStockLockCreate) Save(ctx context.Context) (*AppStockLock, error) {
	var (
		err  error
		node *AppStockLock
	)
	if err := aslc.defaults(); err != nil {
		return nil, err
	}
	if len(aslc.hooks) == 0 {
		if err = aslc.check(); err != nil {
			return nil, err
		}
		node, err = aslc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppStockLockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = aslc.check(); err != nil {
				return nil, err
			}
			aslc.mutation = mutation
			if node, err = aslc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(aslc.hooks) - 1; i >= 0; i-- {
			if aslc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = aslc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, aslc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppStockLock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppStockLockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (aslc *AppStockLockCreate) SaveX(ctx context.Context) *AppStockLock {
	v, err := aslc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aslc *AppStockLockCreate) Exec(ctx context.Context) error {
	_, err := aslc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aslc *AppStockLockCreate) ExecX(ctx context.Context) {
	if err := aslc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (aslc *AppStockLockCreate) defaults() error {
	if _, ok := aslc.mutation.CreatedAt(); !ok {
		if appstocklock.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appstocklock.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appstocklock.DefaultCreatedAt()
		aslc.mutation.SetCreatedAt(v)
	}
	if _, ok := aslc.mutation.UpdatedAt(); !ok {
		if appstocklock.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appstocklock.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appstocklock.DefaultUpdatedAt()
		aslc.mutation.SetUpdatedAt(v)
	}
	if _, ok := aslc.mutation.DeletedAt(); !ok {
		if appstocklock.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appstocklock.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appstocklock.DefaultDeletedAt()
		aslc.mutation.SetDeletedAt(v)
	}
	if _, ok := aslc.mutation.Units(); !ok {
		v := appstocklock.DefaultUnits
		aslc.mutation.SetUnits(v)
	}
	if _, ok := aslc.mutation.ID(); !ok {
		if appstocklock.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appstocklock.DefaultID (forgotten import ent/runtime?)")
		}
		v := appstocklock.DefaultID()
		aslc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (aslc *AppStockLockCreate) check() error {
	if _, ok := aslc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppStockLock.created_at"`)}
	}
	if _, ok := aslc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppStockLock.updated_at"`)}
	}
	if _, ok := aslc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppStockLock.deleted_at"`)}
	}
	return nil
}

func (aslc *AppStockLockCreate) sqlSave(ctx context.Context) (*AppStockLock, error) {
	_node, _spec := aslc.createSpec()
	if err := sqlgraph.CreateNode(ctx, aslc.driver, _spec); err != nil {
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

func (aslc *AppStockLockCreate) createSpec() (*AppStockLock, *sqlgraph.CreateSpec) {
	var (
		_node = &AppStockLock{config: aslc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appstocklock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appstocklock.FieldID,
			},
		}
	)
	_spec.OnConflict = aslc.conflict
	if id, ok := aslc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := aslc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := aslc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := aslc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appstocklock.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := aslc.mutation.Units(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appstocklock.FieldUnits,
		})
		_node.Units = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppStockLock.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppStockLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (aslc *AppStockLockCreate) OnConflict(opts ...sql.ConflictOption) *AppStockLockUpsertOne {
	aslc.conflict = opts
	return &AppStockLockUpsertOne{
		create: aslc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppStockLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (aslc *AppStockLockCreate) OnConflictColumns(columns ...string) *AppStockLockUpsertOne {
	aslc.conflict = append(aslc.conflict, sql.ConflictColumns(columns...))
	return &AppStockLockUpsertOne{
		create: aslc,
	}
}

type (
	// AppStockLockUpsertOne is the builder for "upsert"-ing
	//  one AppStockLock node.
	AppStockLockUpsertOne struct {
		create *AppStockLockCreate
	}

	// AppStockLockUpsert is the "OnConflict" setter.
	AppStockLockUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppStockLockUpsert) SetCreatedAt(v uint32) *AppStockLockUpsert {
	u.Set(appstocklock.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppStockLockUpsert) UpdateCreatedAt() *AppStockLockUpsert {
	u.SetExcluded(appstocklock.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppStockLockUpsert) AddCreatedAt(v uint32) *AppStockLockUpsert {
	u.Add(appstocklock.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppStockLockUpsert) SetUpdatedAt(v uint32) *AppStockLockUpsert {
	u.Set(appstocklock.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppStockLockUpsert) UpdateUpdatedAt() *AppStockLockUpsert {
	u.SetExcluded(appstocklock.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppStockLockUpsert) AddUpdatedAt(v uint32) *AppStockLockUpsert {
	u.Add(appstocklock.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppStockLockUpsert) SetDeletedAt(v uint32) *AppStockLockUpsert {
	u.Set(appstocklock.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppStockLockUpsert) UpdateDeletedAt() *AppStockLockUpsert {
	u.SetExcluded(appstocklock.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppStockLockUpsert) AddDeletedAt(v uint32) *AppStockLockUpsert {
	u.Add(appstocklock.FieldDeletedAt, v)
	return u
}

// SetUnits sets the "units" field.
func (u *AppStockLockUpsert) SetUnits(v decimal.Decimal) *AppStockLockUpsert {
	u.Set(appstocklock.FieldUnits, v)
	return u
}

// UpdateUnits sets the "units" field to the value that was provided on create.
func (u *AppStockLockUpsert) UpdateUnits() *AppStockLockUpsert {
	u.SetExcluded(appstocklock.FieldUnits)
	return u
}

// ClearUnits clears the value of the "units" field.
func (u *AppStockLockUpsert) ClearUnits() *AppStockLockUpsert {
	u.SetNull(appstocklock.FieldUnits)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppStockLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appstocklock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppStockLockUpsertOne) UpdateNewValues() *AppStockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appstocklock.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppStockLock.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppStockLockUpsertOne) Ignore() *AppStockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppStockLockUpsertOne) DoNothing() *AppStockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppStockLockCreate.OnConflict
// documentation for more info.
func (u *AppStockLockUpsertOne) Update(set func(*AppStockLockUpsert)) *AppStockLockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppStockLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppStockLockUpsertOne) SetCreatedAt(v uint32) *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppStockLockUpsertOne) AddCreatedAt(v uint32) *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppStockLockUpsertOne) UpdateCreatedAt() *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppStockLockUpsertOne) SetUpdatedAt(v uint32) *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppStockLockUpsertOne) AddUpdatedAt(v uint32) *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppStockLockUpsertOne) UpdateUpdatedAt() *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppStockLockUpsertOne) SetDeletedAt(v uint32) *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppStockLockUpsertOne) AddDeletedAt(v uint32) *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppStockLockUpsertOne) UpdateDeletedAt() *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUnits sets the "units" field.
func (u *AppStockLockUpsertOne) SetUnits(v decimal.Decimal) *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetUnits(v)
	})
}

// UpdateUnits sets the "units" field to the value that was provided on create.
func (u *AppStockLockUpsertOne) UpdateUnits() *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateUnits()
	})
}

// ClearUnits clears the value of the "units" field.
func (u *AppStockLockUpsertOne) ClearUnits() *AppStockLockUpsertOne {
	return u.Update(func(s *AppStockLockUpsert) {
		s.ClearUnits()
	})
}

// Exec executes the query.
func (u *AppStockLockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppStockLockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppStockLockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppStockLockUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppStockLockUpsertOne.ID is not supported by MySQL driver. Use AppStockLockUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppStockLockUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppStockLockCreateBulk is the builder for creating many AppStockLock entities in bulk.
type AppStockLockCreateBulk struct {
	config
	builders []*AppStockLockCreate
	conflict []sql.ConflictOption
}

// Save creates the AppStockLock entities in the database.
func (aslcb *AppStockLockCreateBulk) Save(ctx context.Context) ([]*AppStockLock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(aslcb.builders))
	nodes := make([]*AppStockLock, len(aslcb.builders))
	mutators := make([]Mutator, len(aslcb.builders))
	for i := range aslcb.builders {
		func(i int, root context.Context) {
			builder := aslcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppStockLockMutation)
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
					_, err = mutators[i+1].Mutate(root, aslcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = aslcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, aslcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, aslcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (aslcb *AppStockLockCreateBulk) SaveX(ctx context.Context) []*AppStockLock {
	v, err := aslcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (aslcb *AppStockLockCreateBulk) Exec(ctx context.Context) error {
	_, err := aslcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (aslcb *AppStockLockCreateBulk) ExecX(ctx context.Context) {
	if err := aslcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppStockLock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppStockLockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (aslcb *AppStockLockCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppStockLockUpsertBulk {
	aslcb.conflict = opts
	return &AppStockLockUpsertBulk{
		create: aslcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppStockLock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (aslcb *AppStockLockCreateBulk) OnConflictColumns(columns ...string) *AppStockLockUpsertBulk {
	aslcb.conflict = append(aslcb.conflict, sql.ConflictColumns(columns...))
	return &AppStockLockUpsertBulk{
		create: aslcb,
	}
}

// AppStockLockUpsertBulk is the builder for "upsert"-ing
// a bulk of AppStockLock nodes.
type AppStockLockUpsertBulk struct {
	create *AppStockLockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppStockLock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appstocklock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppStockLockUpsertBulk) UpdateNewValues() *AppStockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appstocklock.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppStockLock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppStockLockUpsertBulk) Ignore() *AppStockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppStockLockUpsertBulk) DoNothing() *AppStockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppStockLockCreateBulk.OnConflict
// documentation for more info.
func (u *AppStockLockUpsertBulk) Update(set func(*AppStockLockUpsert)) *AppStockLockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppStockLockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppStockLockUpsertBulk) SetCreatedAt(v uint32) *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppStockLockUpsertBulk) AddCreatedAt(v uint32) *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppStockLockUpsertBulk) UpdateCreatedAt() *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppStockLockUpsertBulk) SetUpdatedAt(v uint32) *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppStockLockUpsertBulk) AddUpdatedAt(v uint32) *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppStockLockUpsertBulk) UpdateUpdatedAt() *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppStockLockUpsertBulk) SetDeletedAt(v uint32) *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppStockLockUpsertBulk) AddDeletedAt(v uint32) *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppStockLockUpsertBulk) UpdateDeletedAt() *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetUnits sets the "units" field.
func (u *AppStockLockUpsertBulk) SetUnits(v decimal.Decimal) *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.SetUnits(v)
	})
}

// UpdateUnits sets the "units" field to the value that was provided on create.
func (u *AppStockLockUpsertBulk) UpdateUnits() *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.UpdateUnits()
	})
}

// ClearUnits clears the value of the "units" field.
func (u *AppStockLockUpsertBulk) ClearUnits() *AppStockLockUpsertBulk {
	return u.Update(func(s *AppStockLockUpsert) {
		s.ClearUnits()
	})
}

// Exec executes the query.
func (u *AppStockLockUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppStockLockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppStockLockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppStockLockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
