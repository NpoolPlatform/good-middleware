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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appdefaultgood"
	"github.com/google/uuid"
)

// AppDefaultGoodCreate is the builder for creating a AppDefaultGood entity.
type AppDefaultGoodCreate struct {
	config
	mutation *AppDefaultGoodMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (adgc *AppDefaultGoodCreate) SetCreatedAt(u uint32) *AppDefaultGoodCreate {
	adgc.mutation.SetCreatedAt(u)
	return adgc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableCreatedAt(u *uint32) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetCreatedAt(*u)
	}
	return adgc
}

// SetUpdatedAt sets the "updated_at" field.
func (adgc *AppDefaultGoodCreate) SetUpdatedAt(u uint32) *AppDefaultGoodCreate {
	adgc.mutation.SetUpdatedAt(u)
	return adgc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableUpdatedAt(u *uint32) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetUpdatedAt(*u)
	}
	return adgc
}

// SetDeletedAt sets the "deleted_at" field.
func (adgc *AppDefaultGoodCreate) SetDeletedAt(u uint32) *AppDefaultGoodCreate {
	adgc.mutation.SetDeletedAt(u)
	return adgc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableDeletedAt(u *uint32) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetDeletedAt(*u)
	}
	return adgc
}

// SetAppID sets the "app_id" field.
func (adgc *AppDefaultGoodCreate) SetAppID(u uuid.UUID) *AppDefaultGoodCreate {
	adgc.mutation.SetAppID(u)
	return adgc
}

// SetNillableAppID sets the "app_id" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableAppID(u *uuid.UUID) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetAppID(*u)
	}
	return adgc
}

// SetGoodID sets the "good_id" field.
func (adgc *AppDefaultGoodCreate) SetGoodID(u uuid.UUID) *AppDefaultGoodCreate {
	adgc.mutation.SetGoodID(u)
	return adgc
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableGoodID(u *uuid.UUID) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetGoodID(*u)
	}
	return adgc
}

// SetAppGoodID sets the "app_good_id" field.
func (adgc *AppDefaultGoodCreate) SetAppGoodID(u uuid.UUID) *AppDefaultGoodCreate {
	adgc.mutation.SetAppGoodID(u)
	return adgc
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableAppGoodID(u *uuid.UUID) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetAppGoodID(*u)
	}
	return adgc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (adgc *AppDefaultGoodCreate) SetCoinTypeID(u uuid.UUID) *AppDefaultGoodCreate {
	adgc.mutation.SetCoinTypeID(u)
	return adgc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableCoinTypeID(u *uuid.UUID) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetCoinTypeID(*u)
	}
	return adgc
}

// SetID sets the "id" field.
func (adgc *AppDefaultGoodCreate) SetID(u uuid.UUID) *AppDefaultGoodCreate {
	adgc.mutation.SetID(u)
	return adgc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (adgc *AppDefaultGoodCreate) SetNillableID(u *uuid.UUID) *AppDefaultGoodCreate {
	if u != nil {
		adgc.SetID(*u)
	}
	return adgc
}

// Mutation returns the AppDefaultGoodMutation object of the builder.
func (adgc *AppDefaultGoodCreate) Mutation() *AppDefaultGoodMutation {
	return adgc.mutation
}

// Save creates the AppDefaultGood in the database.
func (adgc *AppDefaultGoodCreate) Save(ctx context.Context) (*AppDefaultGood, error) {
	var (
		err  error
		node *AppDefaultGood
	)
	if err := adgc.defaults(); err != nil {
		return nil, err
	}
	if len(adgc.hooks) == 0 {
		if err = adgc.check(); err != nil {
			return nil, err
		}
		node, err = adgc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppDefaultGoodMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = adgc.check(); err != nil {
				return nil, err
			}
			adgc.mutation = mutation
			if node, err = adgc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(adgc.hooks) - 1; i >= 0; i-- {
			if adgc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = adgc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, adgc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppDefaultGood)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppDefaultGoodMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (adgc *AppDefaultGoodCreate) SaveX(ctx context.Context) *AppDefaultGood {
	v, err := adgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adgc *AppDefaultGoodCreate) Exec(ctx context.Context) error {
	_, err := adgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adgc *AppDefaultGoodCreate) ExecX(ctx context.Context) {
	if err := adgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (adgc *AppDefaultGoodCreate) defaults() error {
	if _, ok := adgc.mutation.CreatedAt(); !ok {
		if appdefaultgood.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appdefaultgood.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appdefaultgood.DefaultCreatedAt()
		adgc.mutation.SetCreatedAt(v)
	}
	if _, ok := adgc.mutation.UpdatedAt(); !ok {
		if appdefaultgood.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appdefaultgood.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appdefaultgood.DefaultUpdatedAt()
		adgc.mutation.SetUpdatedAt(v)
	}
	if _, ok := adgc.mutation.DeletedAt(); !ok {
		if appdefaultgood.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appdefaultgood.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appdefaultgood.DefaultDeletedAt()
		adgc.mutation.SetDeletedAt(v)
	}
	if _, ok := adgc.mutation.ID(); !ok {
		if appdefaultgood.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized appdefaultgood.DefaultID (forgotten import ent/runtime?)")
		}
		v := appdefaultgood.DefaultID()
		adgc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (adgc *AppDefaultGoodCreate) check() error {
	if _, ok := adgc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppDefaultGood.created_at"`)}
	}
	if _, ok := adgc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppDefaultGood.updated_at"`)}
	}
	if _, ok := adgc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppDefaultGood.deleted_at"`)}
	}
	return nil
}

func (adgc *AppDefaultGoodCreate) sqlSave(ctx context.Context) (*AppDefaultGood, error) {
	_node, _spec := adgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, adgc.driver, _spec); err != nil {
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

func (adgc *AppDefaultGoodCreate) createSpec() (*AppDefaultGood, *sqlgraph.CreateSpec) {
	var (
		_node = &AppDefaultGood{config: adgc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appdefaultgood.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appdefaultgood.FieldID,
			},
		}
	)
	_spec.OnConflict = adgc.conflict
	if id, ok := adgc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := adgc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := adgc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := adgc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appdefaultgood.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := adgc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := adgc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := adgc.mutation.AppGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldAppGoodID,
		})
		_node.AppGoodID = value
	}
	if value, ok := adgc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appdefaultgood.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppDefaultGood.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppDefaultGoodUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (adgc *AppDefaultGoodCreate) OnConflict(opts ...sql.ConflictOption) *AppDefaultGoodUpsertOne {
	adgc.conflict = opts
	return &AppDefaultGoodUpsertOne{
		create: adgc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppDefaultGood.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (adgc *AppDefaultGoodCreate) OnConflictColumns(columns ...string) *AppDefaultGoodUpsertOne {
	adgc.conflict = append(adgc.conflict, sql.ConflictColumns(columns...))
	return &AppDefaultGoodUpsertOne{
		create: adgc,
	}
}

type (
	// AppDefaultGoodUpsertOne is the builder for "upsert"-ing
	//  one AppDefaultGood node.
	AppDefaultGoodUpsertOne struct {
		create *AppDefaultGoodCreate
	}

	// AppDefaultGoodUpsert is the "OnConflict" setter.
	AppDefaultGoodUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppDefaultGoodUpsert) SetCreatedAt(v uint32) *AppDefaultGoodUpsert {
	u.Set(appdefaultgood.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsert) UpdateCreatedAt() *AppDefaultGoodUpsert {
	u.SetExcluded(appdefaultgood.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppDefaultGoodUpsert) AddCreatedAt(v uint32) *AppDefaultGoodUpsert {
	u.Add(appdefaultgood.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppDefaultGoodUpsert) SetUpdatedAt(v uint32) *AppDefaultGoodUpsert {
	u.Set(appdefaultgood.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsert) UpdateUpdatedAt() *AppDefaultGoodUpsert {
	u.SetExcluded(appdefaultgood.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppDefaultGoodUpsert) AddUpdatedAt(v uint32) *AppDefaultGoodUpsert {
	u.Add(appdefaultgood.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppDefaultGoodUpsert) SetDeletedAt(v uint32) *AppDefaultGoodUpsert {
	u.Set(appdefaultgood.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsert) UpdateDeletedAt() *AppDefaultGoodUpsert {
	u.SetExcluded(appdefaultgood.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppDefaultGoodUpsert) AddDeletedAt(v uint32) *AppDefaultGoodUpsert {
	u.Add(appdefaultgood.FieldDeletedAt, v)
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppDefaultGoodUpsert) SetAppID(v uuid.UUID) *AppDefaultGoodUpsert {
	u.Set(appdefaultgood.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsert) UpdateAppID() *AppDefaultGoodUpsert {
	u.SetExcluded(appdefaultgood.FieldAppID)
	return u
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppDefaultGoodUpsert) ClearAppID() *AppDefaultGoodUpsert {
	u.SetNull(appdefaultgood.FieldAppID)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *AppDefaultGoodUpsert) SetGoodID(v uuid.UUID) *AppDefaultGoodUpsert {
	u.Set(appdefaultgood.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsert) UpdateGoodID() *AppDefaultGoodUpsert {
	u.SetExcluded(appdefaultgood.FieldGoodID)
	return u
}

// ClearGoodID clears the value of the "good_id" field.
func (u *AppDefaultGoodUpsert) ClearGoodID() *AppDefaultGoodUpsert {
	u.SetNull(appdefaultgood.FieldGoodID)
	return u
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppDefaultGoodUpsert) SetAppGoodID(v uuid.UUID) *AppDefaultGoodUpsert {
	u.Set(appdefaultgood.FieldAppGoodID, v)
	return u
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsert) UpdateAppGoodID() *AppDefaultGoodUpsert {
	u.SetExcluded(appdefaultgood.FieldAppGoodID)
	return u
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppDefaultGoodUpsert) ClearAppGoodID() *AppDefaultGoodUpsert {
	u.SetNull(appdefaultgood.FieldAppGoodID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *AppDefaultGoodUpsert) SetCoinTypeID(v uuid.UUID) *AppDefaultGoodUpsert {
	u.Set(appdefaultgood.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsert) UpdateCoinTypeID() *AppDefaultGoodUpsert {
	u.SetExcluded(appdefaultgood.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *AppDefaultGoodUpsert) ClearCoinTypeID() *AppDefaultGoodUpsert {
	u.SetNull(appdefaultgood.FieldCoinTypeID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppDefaultGood.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appdefaultgood.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppDefaultGoodUpsertOne) UpdateNewValues() *AppDefaultGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appdefaultgood.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppDefaultGood.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppDefaultGoodUpsertOne) Ignore() *AppDefaultGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppDefaultGoodUpsertOne) DoNothing() *AppDefaultGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppDefaultGoodCreate.OnConflict
// documentation for more info.
func (u *AppDefaultGoodUpsertOne) Update(set func(*AppDefaultGoodUpsert)) *AppDefaultGoodUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppDefaultGoodUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppDefaultGoodUpsertOne) SetCreatedAt(v uint32) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppDefaultGoodUpsertOne) AddCreatedAt(v uint32) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertOne) UpdateCreatedAt() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppDefaultGoodUpsertOne) SetUpdatedAt(v uint32) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppDefaultGoodUpsertOne) AddUpdatedAt(v uint32) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertOne) UpdateUpdatedAt() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppDefaultGoodUpsertOne) SetDeletedAt(v uint32) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppDefaultGoodUpsertOne) AddDeletedAt(v uint32) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertOne) UpdateDeletedAt() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppDefaultGoodUpsertOne) SetAppID(v uuid.UUID) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertOne) UpdateAppID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppDefaultGoodUpsertOne) ClearAppID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *AppDefaultGoodUpsertOne) SetGoodID(v uuid.UUID) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertOne) UpdateGoodID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *AppDefaultGoodUpsertOne) ClearGoodID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearGoodID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppDefaultGoodUpsertOne) SetAppGoodID(v uuid.UUID) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertOne) UpdateAppGoodID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppDefaultGoodUpsertOne) ClearAppGoodID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearAppGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *AppDefaultGoodUpsertOne) SetCoinTypeID(v uuid.UUID) *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertOne) UpdateCoinTypeID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *AppDefaultGoodUpsertOne) ClearCoinTypeID() *AppDefaultGoodUpsertOne {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearCoinTypeID()
	})
}

// Exec executes the query.
func (u *AppDefaultGoodUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppDefaultGoodCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppDefaultGoodUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppDefaultGoodUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppDefaultGoodUpsertOne.ID is not supported by MySQL driver. Use AppDefaultGoodUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppDefaultGoodUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppDefaultGoodCreateBulk is the builder for creating many AppDefaultGood entities in bulk.
type AppDefaultGoodCreateBulk struct {
	config
	builders []*AppDefaultGoodCreate
	conflict []sql.ConflictOption
}

// Save creates the AppDefaultGood entities in the database.
func (adgcb *AppDefaultGoodCreateBulk) Save(ctx context.Context) ([]*AppDefaultGood, error) {
	specs := make([]*sqlgraph.CreateSpec, len(adgcb.builders))
	nodes := make([]*AppDefaultGood, len(adgcb.builders))
	mutators := make([]Mutator, len(adgcb.builders))
	for i := range adgcb.builders {
		func(i int, root context.Context) {
			builder := adgcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppDefaultGoodMutation)
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
					_, err = mutators[i+1].Mutate(root, adgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = adgcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, adgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, adgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (adgcb *AppDefaultGoodCreateBulk) SaveX(ctx context.Context) []*AppDefaultGood {
	v, err := adgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (adgcb *AppDefaultGoodCreateBulk) Exec(ctx context.Context) error {
	_, err := adgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (adgcb *AppDefaultGoodCreateBulk) ExecX(ctx context.Context) {
	if err := adgcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppDefaultGood.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppDefaultGoodUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (adgcb *AppDefaultGoodCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppDefaultGoodUpsertBulk {
	adgcb.conflict = opts
	return &AppDefaultGoodUpsertBulk{
		create: adgcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppDefaultGood.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (adgcb *AppDefaultGoodCreateBulk) OnConflictColumns(columns ...string) *AppDefaultGoodUpsertBulk {
	adgcb.conflict = append(adgcb.conflict, sql.ConflictColumns(columns...))
	return &AppDefaultGoodUpsertBulk{
		create: adgcb,
	}
}

// AppDefaultGoodUpsertBulk is the builder for "upsert"-ing
// a bulk of AppDefaultGood nodes.
type AppDefaultGoodUpsertBulk struct {
	create *AppDefaultGoodCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppDefaultGood.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appdefaultgood.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppDefaultGoodUpsertBulk) UpdateNewValues() *AppDefaultGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appdefaultgood.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppDefaultGood.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppDefaultGoodUpsertBulk) Ignore() *AppDefaultGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppDefaultGoodUpsertBulk) DoNothing() *AppDefaultGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppDefaultGoodCreateBulk.OnConflict
// documentation for more info.
func (u *AppDefaultGoodUpsertBulk) Update(set func(*AppDefaultGoodUpsert)) *AppDefaultGoodUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppDefaultGoodUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppDefaultGoodUpsertBulk) SetCreatedAt(v uint32) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppDefaultGoodUpsertBulk) AddCreatedAt(v uint32) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertBulk) UpdateCreatedAt() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppDefaultGoodUpsertBulk) SetUpdatedAt(v uint32) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppDefaultGoodUpsertBulk) AddUpdatedAt(v uint32) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertBulk) UpdateUpdatedAt() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppDefaultGoodUpsertBulk) SetDeletedAt(v uint32) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppDefaultGoodUpsertBulk) AddDeletedAt(v uint32) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertBulk) UpdateDeletedAt() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetAppID sets the "app_id" field.
func (u *AppDefaultGoodUpsertBulk) SetAppID(v uuid.UUID) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertBulk) UpdateAppID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateAppID()
	})
}

// ClearAppID clears the value of the "app_id" field.
func (u *AppDefaultGoodUpsertBulk) ClearAppID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearAppID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *AppDefaultGoodUpsertBulk) SetGoodID(v uuid.UUID) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertBulk) UpdateGoodID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *AppDefaultGoodUpsertBulk) ClearGoodID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearGoodID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppDefaultGoodUpsertBulk) SetAppGoodID(v uuid.UUID) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertBulk) UpdateAppGoodID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppDefaultGoodUpsertBulk) ClearAppGoodID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearAppGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *AppDefaultGoodUpsertBulk) SetCoinTypeID(v uuid.UUID) *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *AppDefaultGoodUpsertBulk) UpdateCoinTypeID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *AppDefaultGoodUpsertBulk) ClearCoinTypeID() *AppDefaultGoodUpsertBulk {
	return u.Update(func(s *AppDefaultGoodUpsert) {
		s.ClearCoinTypeID()
	})
}

// Exec executes the query.
func (u *AppDefaultGoodUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppDefaultGoodCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppDefaultGoodCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppDefaultGoodUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
