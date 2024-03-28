// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulatepowerrental"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppSimulatePowerRentalCreate is the builder for creating a AppSimulatePowerRental entity.
type AppSimulatePowerRentalCreate struct {
	config
	mutation *AppSimulatePowerRentalMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (asprc *AppSimulatePowerRentalCreate) SetCreatedAt(u uint32) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetCreatedAt(u)
	return asprc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableCreatedAt(u *uint32) *AppSimulatePowerRentalCreate {
	if u != nil {
		asprc.SetCreatedAt(*u)
	}
	return asprc
}

// SetUpdatedAt sets the "updated_at" field.
func (asprc *AppSimulatePowerRentalCreate) SetUpdatedAt(u uint32) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetUpdatedAt(u)
	return asprc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableUpdatedAt(u *uint32) *AppSimulatePowerRentalCreate {
	if u != nil {
		asprc.SetUpdatedAt(*u)
	}
	return asprc
}

// SetDeletedAt sets the "deleted_at" field.
func (asprc *AppSimulatePowerRentalCreate) SetDeletedAt(u uint32) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetDeletedAt(u)
	return asprc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableDeletedAt(u *uint32) *AppSimulatePowerRentalCreate {
	if u != nil {
		asprc.SetDeletedAt(*u)
	}
	return asprc
}

// SetEntID sets the "ent_id" field.
func (asprc *AppSimulatePowerRentalCreate) SetEntID(u uuid.UUID) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetEntID(u)
	return asprc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableEntID(u *uuid.UUID) *AppSimulatePowerRentalCreate {
	if u != nil {
		asprc.SetEntID(*u)
	}
	return asprc
}

// SetAppGoodID sets the "app_good_id" field.
func (asprc *AppSimulatePowerRentalCreate) SetAppGoodID(u uuid.UUID) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetAppGoodID(u)
	return asprc
}

// SetNillableAppGoodID sets the "app_good_id" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableAppGoodID(u *uuid.UUID) *AppSimulatePowerRentalCreate {
	if u != nil {
		asprc.SetAppGoodID(*u)
	}
	return asprc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (asprc *AppSimulatePowerRentalCreate) SetCoinTypeID(u uuid.UUID) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetCoinTypeID(u)
	return asprc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableCoinTypeID(u *uuid.UUID) *AppSimulatePowerRentalCreate {
	if u != nil {
		asprc.SetCoinTypeID(*u)
	}
	return asprc
}

// SetOrderUnits sets the "order_units" field.
func (asprc *AppSimulatePowerRentalCreate) SetOrderUnits(d decimal.Decimal) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetOrderUnits(d)
	return asprc
}

// SetNillableOrderUnits sets the "order_units" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableOrderUnits(d *decimal.Decimal) *AppSimulatePowerRentalCreate {
	if d != nil {
		asprc.SetOrderUnits(*d)
	}
	return asprc
}

// SetOrderDuration sets the "order_duration" field.
func (asprc *AppSimulatePowerRentalCreate) SetOrderDuration(u uint32) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetOrderDuration(u)
	return asprc
}

// SetNillableOrderDuration sets the "order_duration" field if the given value is not nil.
func (asprc *AppSimulatePowerRentalCreate) SetNillableOrderDuration(u *uint32) *AppSimulatePowerRentalCreate {
	if u != nil {
		asprc.SetOrderDuration(*u)
	}
	return asprc
}

// SetID sets the "id" field.
func (asprc *AppSimulatePowerRentalCreate) SetID(u uint32) *AppSimulatePowerRentalCreate {
	asprc.mutation.SetID(u)
	return asprc
}

// Mutation returns the AppSimulatePowerRentalMutation object of the builder.
func (asprc *AppSimulatePowerRentalCreate) Mutation() *AppSimulatePowerRentalMutation {
	return asprc.mutation
}

// Save creates the AppSimulatePowerRental in the database.
func (asprc *AppSimulatePowerRentalCreate) Save(ctx context.Context) (*AppSimulatePowerRental, error) {
	var (
		err  error
		node *AppSimulatePowerRental
	)
	if err := asprc.defaults(); err != nil {
		return nil, err
	}
	if len(asprc.hooks) == 0 {
		if err = asprc.check(); err != nil {
			return nil, err
		}
		node, err = asprc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppSimulatePowerRentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = asprc.check(); err != nil {
				return nil, err
			}
			asprc.mutation = mutation
			if node, err = asprc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(asprc.hooks) - 1; i >= 0; i-- {
			if asprc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = asprc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, asprc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*AppSimulatePowerRental)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from AppSimulatePowerRentalMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (asprc *AppSimulatePowerRentalCreate) SaveX(ctx context.Context) *AppSimulatePowerRental {
	v, err := asprc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asprc *AppSimulatePowerRentalCreate) Exec(ctx context.Context) error {
	_, err := asprc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asprc *AppSimulatePowerRentalCreate) ExecX(ctx context.Context) {
	if err := asprc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (asprc *AppSimulatePowerRentalCreate) defaults() error {
	if _, ok := asprc.mutation.CreatedAt(); !ok {
		if appsimulatepowerrental.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsimulatepowerrental.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := appsimulatepowerrental.DefaultCreatedAt()
		asprc.mutation.SetCreatedAt(v)
	}
	if _, ok := asprc.mutation.UpdatedAt(); !ok {
		if appsimulatepowerrental.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized appsimulatepowerrental.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := appsimulatepowerrental.DefaultUpdatedAt()
		asprc.mutation.SetUpdatedAt(v)
	}
	if _, ok := asprc.mutation.DeletedAt(); !ok {
		if appsimulatepowerrental.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized appsimulatepowerrental.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := appsimulatepowerrental.DefaultDeletedAt()
		asprc.mutation.SetDeletedAt(v)
	}
	if _, ok := asprc.mutation.EntID(); !ok {
		if appsimulatepowerrental.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized appsimulatepowerrental.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := appsimulatepowerrental.DefaultEntID()
		asprc.mutation.SetEntID(v)
	}
	if _, ok := asprc.mutation.OrderUnits(); !ok {
		v := appsimulatepowerrental.DefaultOrderUnits
		asprc.mutation.SetOrderUnits(v)
	}
	if _, ok := asprc.mutation.OrderDuration(); !ok {
		v := appsimulatepowerrental.DefaultOrderDuration
		asprc.mutation.SetOrderDuration(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (asprc *AppSimulatePowerRentalCreate) check() error {
	if _, ok := asprc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "AppSimulatePowerRental.created_at"`)}
	}
	if _, ok := asprc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "AppSimulatePowerRental.updated_at"`)}
	}
	if _, ok := asprc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "AppSimulatePowerRental.deleted_at"`)}
	}
	if _, ok := asprc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "AppSimulatePowerRental.ent_id"`)}
	}
	return nil
}

func (asprc *AppSimulatePowerRentalCreate) sqlSave(ctx context.Context) (*AppSimulatePowerRental, error) {
	_node, _spec := asprc.createSpec()
	if err := sqlgraph.CreateNode(ctx, asprc.driver, _spec); err != nil {
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

func (asprc *AppSimulatePowerRentalCreate) createSpec() (*AppSimulatePowerRental, *sqlgraph.CreateSpec) {
	var (
		_node = &AppSimulatePowerRental{config: asprc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appsimulatepowerrental.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appsimulatepowerrental.FieldID,
			},
		}
	)
	_spec.OnConflict = asprc.conflict
	if id, ok := asprc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := asprc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := asprc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := asprc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := asprc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := asprc.mutation.AppGoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldAppGoodID,
		})
		_node.AppGoodID = value
	}
	if value, ok := asprc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appsimulatepowerrental.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := asprc.mutation.OrderUnits(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderUnits,
		})
		_node.OrderUnits = value
	}
	if value, ok := asprc.mutation.OrderDuration(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appsimulatepowerrental.FieldOrderDuration,
		})
		_node.OrderDuration = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppSimulatePowerRental.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppSimulatePowerRentalUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (asprc *AppSimulatePowerRentalCreate) OnConflict(opts ...sql.ConflictOption) *AppSimulatePowerRentalUpsertOne {
	asprc.conflict = opts
	return &AppSimulatePowerRentalUpsertOne{
		create: asprc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppSimulatePowerRental.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (asprc *AppSimulatePowerRentalCreate) OnConflictColumns(columns ...string) *AppSimulatePowerRentalUpsertOne {
	asprc.conflict = append(asprc.conflict, sql.ConflictColumns(columns...))
	return &AppSimulatePowerRentalUpsertOne{
		create: asprc,
	}
}

type (
	// AppSimulatePowerRentalUpsertOne is the builder for "upsert"-ing
	//  one AppSimulatePowerRental node.
	AppSimulatePowerRentalUpsertOne struct {
		create *AppSimulatePowerRentalCreate
	}

	// AppSimulatePowerRentalUpsert is the "OnConflict" setter.
	AppSimulatePowerRentalUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *AppSimulatePowerRentalUpsert) SetCreatedAt(v uint32) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateCreatedAt() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSimulatePowerRentalUpsert) AddCreatedAt(v uint32) *AppSimulatePowerRentalUpsert {
	u.Add(appsimulatepowerrental.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSimulatePowerRentalUpsert) SetUpdatedAt(v uint32) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateUpdatedAt() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSimulatePowerRentalUpsert) AddUpdatedAt(v uint32) *AppSimulatePowerRentalUpsert {
	u.Add(appsimulatepowerrental.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSimulatePowerRentalUpsert) SetDeletedAt(v uint32) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateDeletedAt() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSimulatePowerRentalUpsert) AddDeletedAt(v uint32) *AppSimulatePowerRentalUpsert {
	u.Add(appsimulatepowerrental.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *AppSimulatePowerRentalUpsert) SetEntID(v uuid.UUID) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateEntID() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldEntID)
	return u
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppSimulatePowerRentalUpsert) SetAppGoodID(v uuid.UUID) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldAppGoodID, v)
	return u
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateAppGoodID() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldAppGoodID)
	return u
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppSimulatePowerRentalUpsert) ClearAppGoodID() *AppSimulatePowerRentalUpsert {
	u.SetNull(appsimulatepowerrental.FieldAppGoodID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *AppSimulatePowerRentalUpsert) SetCoinTypeID(v uuid.UUID) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateCoinTypeID() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *AppSimulatePowerRentalUpsert) ClearCoinTypeID() *AppSimulatePowerRentalUpsert {
	u.SetNull(appsimulatepowerrental.FieldCoinTypeID)
	return u
}

// SetOrderUnits sets the "order_units" field.
func (u *AppSimulatePowerRentalUpsert) SetOrderUnits(v decimal.Decimal) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldOrderUnits, v)
	return u
}

// UpdateOrderUnits sets the "order_units" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateOrderUnits() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldOrderUnits)
	return u
}

// ClearOrderUnits clears the value of the "order_units" field.
func (u *AppSimulatePowerRentalUpsert) ClearOrderUnits() *AppSimulatePowerRentalUpsert {
	u.SetNull(appsimulatepowerrental.FieldOrderUnits)
	return u
}

// SetOrderDuration sets the "order_duration" field.
func (u *AppSimulatePowerRentalUpsert) SetOrderDuration(v uint32) *AppSimulatePowerRentalUpsert {
	u.Set(appsimulatepowerrental.FieldOrderDuration, v)
	return u
}

// UpdateOrderDuration sets the "order_duration" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsert) UpdateOrderDuration() *AppSimulatePowerRentalUpsert {
	u.SetExcluded(appsimulatepowerrental.FieldOrderDuration)
	return u
}

// AddOrderDuration adds v to the "order_duration" field.
func (u *AppSimulatePowerRentalUpsert) AddOrderDuration(v uint32) *AppSimulatePowerRentalUpsert {
	u.Add(appsimulatepowerrental.FieldOrderDuration, v)
	return u
}

// ClearOrderDuration clears the value of the "order_duration" field.
func (u *AppSimulatePowerRentalUpsert) ClearOrderDuration() *AppSimulatePowerRentalUpsert {
	u.SetNull(appsimulatepowerrental.FieldOrderDuration)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppSimulatePowerRental.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appsimulatepowerrental.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppSimulatePowerRentalUpsertOne) UpdateNewValues() *AppSimulatePowerRentalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appsimulatepowerrental.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppSimulatePowerRental.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppSimulatePowerRentalUpsertOne) Ignore() *AppSimulatePowerRentalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppSimulatePowerRentalUpsertOne) DoNothing() *AppSimulatePowerRentalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppSimulatePowerRentalCreate.OnConflict
// documentation for more info.
func (u *AppSimulatePowerRentalUpsertOne) Update(set func(*AppSimulatePowerRentalUpsert)) *AppSimulatePowerRentalUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppSimulatePowerRentalUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppSimulatePowerRentalUpsertOne) SetCreatedAt(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSimulatePowerRentalUpsertOne) AddCreatedAt(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateCreatedAt() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSimulatePowerRentalUpsertOne) SetUpdatedAt(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSimulatePowerRentalUpsertOne) AddUpdatedAt(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateUpdatedAt() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSimulatePowerRentalUpsertOne) SetDeletedAt(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSimulatePowerRentalUpsertOne) AddDeletedAt(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateDeletedAt() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppSimulatePowerRentalUpsertOne) SetEntID(v uuid.UUID) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateEntID() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppSimulatePowerRentalUpsertOne) SetAppGoodID(v uuid.UUID) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateAppGoodID() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppSimulatePowerRentalUpsertOne) ClearAppGoodID() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearAppGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *AppSimulatePowerRentalUpsertOne) SetCoinTypeID(v uuid.UUID) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateCoinTypeID() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *AppSimulatePowerRentalUpsertOne) ClearCoinTypeID() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetOrderUnits sets the "order_units" field.
func (u *AppSimulatePowerRentalUpsertOne) SetOrderUnits(v decimal.Decimal) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetOrderUnits(v)
	})
}

// UpdateOrderUnits sets the "order_units" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateOrderUnits() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateOrderUnits()
	})
}

// ClearOrderUnits clears the value of the "order_units" field.
func (u *AppSimulatePowerRentalUpsertOne) ClearOrderUnits() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearOrderUnits()
	})
}

// SetOrderDuration sets the "order_duration" field.
func (u *AppSimulatePowerRentalUpsertOne) SetOrderDuration(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetOrderDuration(v)
	})
}

// AddOrderDuration adds v to the "order_duration" field.
func (u *AppSimulatePowerRentalUpsertOne) AddOrderDuration(v uint32) *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddOrderDuration(v)
	})
}

// UpdateOrderDuration sets the "order_duration" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertOne) UpdateOrderDuration() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateOrderDuration()
	})
}

// ClearOrderDuration clears the value of the "order_duration" field.
func (u *AppSimulatePowerRentalUpsertOne) ClearOrderDuration() *AppSimulatePowerRentalUpsertOne {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearOrderDuration()
	})
}

// Exec executes the query.
func (u *AppSimulatePowerRentalUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppSimulatePowerRentalCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppSimulatePowerRentalUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppSimulatePowerRentalUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppSimulatePowerRentalUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppSimulatePowerRentalCreateBulk is the builder for creating many AppSimulatePowerRental entities in bulk.
type AppSimulatePowerRentalCreateBulk struct {
	config
	builders []*AppSimulatePowerRentalCreate
	conflict []sql.ConflictOption
}

// Save creates the AppSimulatePowerRental entities in the database.
func (asprcb *AppSimulatePowerRentalCreateBulk) Save(ctx context.Context) ([]*AppSimulatePowerRental, error) {
	specs := make([]*sqlgraph.CreateSpec, len(asprcb.builders))
	nodes := make([]*AppSimulatePowerRental, len(asprcb.builders))
	mutators := make([]Mutator, len(asprcb.builders))
	for i := range asprcb.builders {
		func(i int, root context.Context) {
			builder := asprcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppSimulatePowerRentalMutation)
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
					_, err = mutators[i+1].Mutate(root, asprcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = asprcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, asprcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, asprcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (asprcb *AppSimulatePowerRentalCreateBulk) SaveX(ctx context.Context) []*AppSimulatePowerRental {
	v, err := asprcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (asprcb *AppSimulatePowerRentalCreateBulk) Exec(ctx context.Context) error {
	_, err := asprcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (asprcb *AppSimulatePowerRentalCreateBulk) ExecX(ctx context.Context) {
	if err := asprcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppSimulatePowerRental.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppSimulatePowerRentalUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (asprcb *AppSimulatePowerRentalCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppSimulatePowerRentalUpsertBulk {
	asprcb.conflict = opts
	return &AppSimulatePowerRentalUpsertBulk{
		create: asprcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppSimulatePowerRental.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (asprcb *AppSimulatePowerRentalCreateBulk) OnConflictColumns(columns ...string) *AppSimulatePowerRentalUpsertBulk {
	asprcb.conflict = append(asprcb.conflict, sql.ConflictColumns(columns...))
	return &AppSimulatePowerRentalUpsertBulk{
		create: asprcb,
	}
}

// AppSimulatePowerRentalUpsertBulk is the builder for "upsert"-ing
// a bulk of AppSimulatePowerRental nodes.
type AppSimulatePowerRentalUpsertBulk struct {
	create *AppSimulatePowerRentalCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppSimulatePowerRental.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appsimulatepowerrental.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppSimulatePowerRentalUpsertBulk) UpdateNewValues() *AppSimulatePowerRentalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appsimulatepowerrental.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppSimulatePowerRental.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppSimulatePowerRentalUpsertBulk) Ignore() *AppSimulatePowerRentalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppSimulatePowerRentalUpsertBulk) DoNothing() *AppSimulatePowerRentalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppSimulatePowerRentalCreateBulk.OnConflict
// documentation for more info.
func (u *AppSimulatePowerRentalUpsertBulk) Update(set func(*AppSimulatePowerRentalUpsert)) *AppSimulatePowerRentalUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppSimulatePowerRentalUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetCreatedAt(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *AppSimulatePowerRentalUpsertBulk) AddCreatedAt(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateCreatedAt() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetUpdatedAt(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AppSimulatePowerRentalUpsertBulk) AddUpdatedAt(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateUpdatedAt() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetDeletedAt(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AppSimulatePowerRentalUpsertBulk) AddDeletedAt(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateDeletedAt() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetEntID(v uuid.UUID) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateEntID() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateEntID()
	})
}

// SetAppGoodID sets the "app_good_id" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetAppGoodID(v uuid.UUID) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetAppGoodID(v)
	})
}

// UpdateAppGoodID sets the "app_good_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateAppGoodID() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateAppGoodID()
	})
}

// ClearAppGoodID clears the value of the "app_good_id" field.
func (u *AppSimulatePowerRentalUpsertBulk) ClearAppGoodID() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearAppGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetCoinTypeID(v uuid.UUID) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateCoinTypeID() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *AppSimulatePowerRentalUpsertBulk) ClearCoinTypeID() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetOrderUnits sets the "order_units" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetOrderUnits(v decimal.Decimal) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetOrderUnits(v)
	})
}

// UpdateOrderUnits sets the "order_units" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateOrderUnits() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateOrderUnits()
	})
}

// ClearOrderUnits clears the value of the "order_units" field.
func (u *AppSimulatePowerRentalUpsertBulk) ClearOrderUnits() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearOrderUnits()
	})
}

// SetOrderDuration sets the "order_duration" field.
func (u *AppSimulatePowerRentalUpsertBulk) SetOrderDuration(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.SetOrderDuration(v)
	})
}

// AddOrderDuration adds v to the "order_duration" field.
func (u *AppSimulatePowerRentalUpsertBulk) AddOrderDuration(v uint32) *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.AddOrderDuration(v)
	})
}

// UpdateOrderDuration sets the "order_duration" field to the value that was provided on create.
func (u *AppSimulatePowerRentalUpsertBulk) UpdateOrderDuration() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.UpdateOrderDuration()
	})
}

// ClearOrderDuration clears the value of the "order_duration" field.
func (u *AppSimulatePowerRentalUpsertBulk) ClearOrderDuration() *AppSimulatePowerRentalUpsertBulk {
	return u.Update(func(s *AppSimulatePowerRentalUpsert) {
		s.ClearOrderDuration()
	})
}

// Exec executes the query.
func (u *AppSimulatePowerRentalUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppSimulatePowerRentalCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppSimulatePowerRentalCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppSimulatePowerRentalUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}