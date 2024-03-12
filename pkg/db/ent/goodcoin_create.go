// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	"github.com/google/uuid"
)

// GoodCoinCreate is the builder for creating a GoodCoin entity.
type GoodCoinCreate struct {
	config
	mutation *GoodCoinMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (gcc *GoodCoinCreate) SetCreatedAt(u uint32) *GoodCoinCreate {
	gcc.mutation.SetCreatedAt(u)
	return gcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (gcc *GoodCoinCreate) SetNillableCreatedAt(u *uint32) *GoodCoinCreate {
	if u != nil {
		gcc.SetCreatedAt(*u)
	}
	return gcc
}

// SetUpdatedAt sets the "updated_at" field.
func (gcc *GoodCoinCreate) SetUpdatedAt(u uint32) *GoodCoinCreate {
	gcc.mutation.SetUpdatedAt(u)
	return gcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (gcc *GoodCoinCreate) SetNillableUpdatedAt(u *uint32) *GoodCoinCreate {
	if u != nil {
		gcc.SetUpdatedAt(*u)
	}
	return gcc
}

// SetDeletedAt sets the "deleted_at" field.
func (gcc *GoodCoinCreate) SetDeletedAt(u uint32) *GoodCoinCreate {
	gcc.mutation.SetDeletedAt(u)
	return gcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gcc *GoodCoinCreate) SetNillableDeletedAt(u *uint32) *GoodCoinCreate {
	if u != nil {
		gcc.SetDeletedAt(*u)
	}
	return gcc
}

// SetEntID sets the "ent_id" field.
func (gcc *GoodCoinCreate) SetEntID(u uuid.UUID) *GoodCoinCreate {
	gcc.mutation.SetEntID(u)
	return gcc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (gcc *GoodCoinCreate) SetNillableEntID(u *uuid.UUID) *GoodCoinCreate {
	if u != nil {
		gcc.SetEntID(*u)
	}
	return gcc
}

// SetGoodID sets the "good_id" field.
func (gcc *GoodCoinCreate) SetGoodID(u uuid.UUID) *GoodCoinCreate {
	gcc.mutation.SetGoodID(u)
	return gcc
}

// SetNillableGoodID sets the "good_id" field if the given value is not nil.
func (gcc *GoodCoinCreate) SetNillableGoodID(u *uuid.UUID) *GoodCoinCreate {
	if u != nil {
		gcc.SetGoodID(*u)
	}
	return gcc
}

// SetCoinTypeID sets the "coin_type_id" field.
func (gcc *GoodCoinCreate) SetCoinTypeID(u uuid.UUID) *GoodCoinCreate {
	gcc.mutation.SetCoinTypeID(u)
	return gcc
}

// SetNillableCoinTypeID sets the "coin_type_id" field if the given value is not nil.
func (gcc *GoodCoinCreate) SetNillableCoinTypeID(u *uuid.UUID) *GoodCoinCreate {
	if u != nil {
		gcc.SetCoinTypeID(*u)
	}
	return gcc
}

// SetMain sets the "main" field.
func (gcc *GoodCoinCreate) SetMain(b bool) *GoodCoinCreate {
	gcc.mutation.SetMain(b)
	return gcc
}

// SetNillableMain sets the "main" field if the given value is not nil.
func (gcc *GoodCoinCreate) SetNillableMain(b *bool) *GoodCoinCreate {
	if b != nil {
		gcc.SetMain(*b)
	}
	return gcc
}

// SetID sets the "id" field.
func (gcc *GoodCoinCreate) SetID(u uint32) *GoodCoinCreate {
	gcc.mutation.SetID(u)
	return gcc
}

// Mutation returns the GoodCoinMutation object of the builder.
func (gcc *GoodCoinCreate) Mutation() *GoodCoinMutation {
	return gcc.mutation
}

// Save creates the GoodCoin in the database.
func (gcc *GoodCoinCreate) Save(ctx context.Context) (*GoodCoin, error) {
	var (
		err  error
		node *GoodCoin
	)
	if err := gcc.defaults(); err != nil {
		return nil, err
	}
	if len(gcc.hooks) == 0 {
		if err = gcc.check(); err != nil {
			return nil, err
		}
		node, err = gcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GoodCoinMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gcc.check(); err != nil {
				return nil, err
			}
			gcc.mutation = mutation
			if node, err = gcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gcc.hooks) - 1; i >= 0; i-- {
			if gcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, gcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*GoodCoin)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from GoodCoinMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gcc *GoodCoinCreate) SaveX(ctx context.Context) *GoodCoin {
	v, err := gcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcc *GoodCoinCreate) Exec(ctx context.Context) error {
	_, err := gcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcc *GoodCoinCreate) ExecX(ctx context.Context) {
	if err := gcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gcc *GoodCoinCreate) defaults() error {
	if _, ok := gcc.mutation.CreatedAt(); !ok {
		if goodcoin.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := goodcoin.DefaultCreatedAt()
		gcc.mutation.SetCreatedAt(v)
	}
	if _, ok := gcc.mutation.UpdatedAt(); !ok {
		if goodcoin.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := goodcoin.DefaultUpdatedAt()
		gcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := gcc.mutation.DeletedAt(); !ok {
		if goodcoin.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := goodcoin.DefaultDeletedAt()
		gcc.mutation.SetDeletedAt(v)
	}
	if _, ok := gcc.mutation.EntID(); !ok {
		if goodcoin.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := goodcoin.DefaultEntID()
		gcc.mutation.SetEntID(v)
	}
	if _, ok := gcc.mutation.GoodID(); !ok {
		if goodcoin.DefaultGoodID == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.DefaultGoodID (forgotten import ent/runtime?)")
		}
		v := goodcoin.DefaultGoodID()
		gcc.mutation.SetGoodID(v)
	}
	if _, ok := gcc.mutation.CoinTypeID(); !ok {
		if goodcoin.DefaultCoinTypeID == nil {
			return fmt.Errorf("ent: uninitialized goodcoin.DefaultCoinTypeID (forgotten import ent/runtime?)")
		}
		v := goodcoin.DefaultCoinTypeID()
		gcc.mutation.SetCoinTypeID(v)
	}
	if _, ok := gcc.mutation.Main(); !ok {
		v := goodcoin.DefaultMain
		gcc.mutation.SetMain(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gcc *GoodCoinCreate) check() error {
	if _, ok := gcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "GoodCoin.created_at"`)}
	}
	if _, ok := gcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "GoodCoin.updated_at"`)}
	}
	if _, ok := gcc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "GoodCoin.deleted_at"`)}
	}
	if _, ok := gcc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "GoodCoin.ent_id"`)}
	}
	return nil
}

func (gcc *GoodCoinCreate) sqlSave(ctx context.Context) (*GoodCoin, error) {
	_node, _spec := gcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gcc.driver, _spec); err != nil {
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

func (gcc *GoodCoinCreate) createSpec() (*GoodCoin, *sqlgraph.CreateSpec) {
	var (
		_node = &GoodCoin{config: gcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: goodcoin.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: goodcoin.FieldID,
			},
		}
	)
	_spec.OnConflict = gcc.conflict
	if id, ok := gcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := gcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := gcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: goodcoin.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := gcc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := gcc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := gcc.mutation.CoinTypeID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: goodcoin.FieldCoinTypeID,
		})
		_node.CoinTypeID = value
	}
	if value, ok := gcc.mutation.Main(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: goodcoin.FieldMain,
		})
		_node.Main = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodCoin.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodCoinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (gcc *GoodCoinCreate) OnConflict(opts ...sql.ConflictOption) *GoodCoinUpsertOne {
	gcc.conflict = opts
	return &GoodCoinUpsertOne{
		create: gcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodCoin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gcc *GoodCoinCreate) OnConflictColumns(columns ...string) *GoodCoinUpsertOne {
	gcc.conflict = append(gcc.conflict, sql.ConflictColumns(columns...))
	return &GoodCoinUpsertOne{
		create: gcc,
	}
}

type (
	// GoodCoinUpsertOne is the builder for "upsert"-ing
	//  one GoodCoin node.
	GoodCoinUpsertOne struct {
		create *GoodCoinCreate
	}

	// GoodCoinUpsert is the "OnConflict" setter.
	GoodCoinUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *GoodCoinUpsert) SetCreatedAt(v uint32) *GoodCoinUpsert {
	u.Set(goodcoin.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GoodCoinUpsert) UpdateCreatedAt() *GoodCoinUpsert {
	u.SetExcluded(goodcoin.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *GoodCoinUpsert) AddCreatedAt(v uint32) *GoodCoinUpsert {
	u.Add(goodcoin.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GoodCoinUpsert) SetUpdatedAt(v uint32) *GoodCoinUpsert {
	u.Set(goodcoin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GoodCoinUpsert) UpdateUpdatedAt() *GoodCoinUpsert {
	u.SetExcluded(goodcoin.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *GoodCoinUpsert) AddUpdatedAt(v uint32) *GoodCoinUpsert {
	u.Add(goodcoin.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GoodCoinUpsert) SetDeletedAt(v uint32) *GoodCoinUpsert {
	u.Set(goodcoin.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GoodCoinUpsert) UpdateDeletedAt() *GoodCoinUpsert {
	u.SetExcluded(goodcoin.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *GoodCoinUpsert) AddDeletedAt(v uint32) *GoodCoinUpsert {
	u.Add(goodcoin.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *GoodCoinUpsert) SetEntID(v uuid.UUID) *GoodCoinUpsert {
	u.Set(goodcoin.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *GoodCoinUpsert) UpdateEntID() *GoodCoinUpsert {
	u.SetExcluded(goodcoin.FieldEntID)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *GoodCoinUpsert) SetGoodID(v uuid.UUID) *GoodCoinUpsert {
	u.Set(goodcoin.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *GoodCoinUpsert) UpdateGoodID() *GoodCoinUpsert {
	u.SetExcluded(goodcoin.FieldGoodID)
	return u
}

// ClearGoodID clears the value of the "good_id" field.
func (u *GoodCoinUpsert) ClearGoodID() *GoodCoinUpsert {
	u.SetNull(goodcoin.FieldGoodID)
	return u
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *GoodCoinUpsert) SetCoinTypeID(v uuid.UUID) *GoodCoinUpsert {
	u.Set(goodcoin.FieldCoinTypeID, v)
	return u
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *GoodCoinUpsert) UpdateCoinTypeID() *GoodCoinUpsert {
	u.SetExcluded(goodcoin.FieldCoinTypeID)
	return u
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *GoodCoinUpsert) ClearCoinTypeID() *GoodCoinUpsert {
	u.SetNull(goodcoin.FieldCoinTypeID)
	return u
}

// SetMain sets the "main" field.
func (u *GoodCoinUpsert) SetMain(v bool) *GoodCoinUpsert {
	u.Set(goodcoin.FieldMain, v)
	return u
}

// UpdateMain sets the "main" field to the value that was provided on create.
func (u *GoodCoinUpsert) UpdateMain() *GoodCoinUpsert {
	u.SetExcluded(goodcoin.FieldMain)
	return u
}

// ClearMain clears the value of the "main" field.
func (u *GoodCoinUpsert) ClearMain() *GoodCoinUpsert {
	u.SetNull(goodcoin.FieldMain)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.GoodCoin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goodcoin.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GoodCoinUpsertOne) UpdateNewValues() *GoodCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(goodcoin.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.GoodCoin.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *GoodCoinUpsertOne) Ignore() *GoodCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodCoinUpsertOne) DoNothing() *GoodCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodCoinCreate.OnConflict
// documentation for more info.
func (u *GoodCoinUpsertOne) Update(set func(*GoodCoinUpsert)) *GoodCoinUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodCoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GoodCoinUpsertOne) SetCreatedAt(v uint32) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *GoodCoinUpsertOne) AddCreatedAt(v uint32) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GoodCoinUpsertOne) UpdateCreatedAt() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GoodCoinUpsertOne) SetUpdatedAt(v uint32) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *GoodCoinUpsertOne) AddUpdatedAt(v uint32) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GoodCoinUpsertOne) UpdateUpdatedAt() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GoodCoinUpsertOne) SetDeletedAt(v uint32) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *GoodCoinUpsertOne) AddDeletedAt(v uint32) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GoodCoinUpsertOne) UpdateDeletedAt() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *GoodCoinUpsertOne) SetEntID(v uuid.UUID) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *GoodCoinUpsertOne) UpdateEntID() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateEntID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *GoodCoinUpsertOne) SetGoodID(v uuid.UUID) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *GoodCoinUpsertOne) UpdateGoodID() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *GoodCoinUpsertOne) ClearGoodID() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *GoodCoinUpsertOne) SetCoinTypeID(v uuid.UUID) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *GoodCoinUpsertOne) UpdateCoinTypeID() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *GoodCoinUpsertOne) ClearCoinTypeID() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetMain sets the "main" field.
func (u *GoodCoinUpsertOne) SetMain(v bool) *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetMain(v)
	})
}

// UpdateMain sets the "main" field to the value that was provided on create.
func (u *GoodCoinUpsertOne) UpdateMain() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateMain()
	})
}

// ClearMain clears the value of the "main" field.
func (u *GoodCoinUpsertOne) ClearMain() *GoodCoinUpsertOne {
	return u.Update(func(s *GoodCoinUpsert) {
		s.ClearMain()
	})
}

// Exec executes the query.
func (u *GoodCoinUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodCoinCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodCoinUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GoodCoinUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GoodCoinUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GoodCoinCreateBulk is the builder for creating many GoodCoin entities in bulk.
type GoodCoinCreateBulk struct {
	config
	builders []*GoodCoinCreate
	conflict []sql.ConflictOption
}

// Save creates the GoodCoin entities in the database.
func (gccb *GoodCoinCreateBulk) Save(ctx context.Context) ([]*GoodCoin, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gccb.builders))
	nodes := make([]*GoodCoin, len(gccb.builders))
	mutators := make([]Mutator, len(gccb.builders))
	for i := range gccb.builders {
		func(i int, root context.Context) {
			builder := gccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GoodCoinMutation)
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
					_, err = mutators[i+1].Mutate(root, gccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gccb *GoodCoinCreateBulk) SaveX(ctx context.Context) []*GoodCoin {
	v, err := gccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gccb *GoodCoinCreateBulk) Exec(ctx context.Context) error {
	_, err := gccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gccb *GoodCoinCreateBulk) ExecX(ctx context.Context) {
	if err := gccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GoodCoin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GoodCoinUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (gccb *GoodCoinCreateBulk) OnConflict(opts ...sql.ConflictOption) *GoodCoinUpsertBulk {
	gccb.conflict = opts
	return &GoodCoinUpsertBulk{
		create: gccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GoodCoin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gccb *GoodCoinCreateBulk) OnConflictColumns(columns ...string) *GoodCoinUpsertBulk {
	gccb.conflict = append(gccb.conflict, sql.ConflictColumns(columns...))
	return &GoodCoinUpsertBulk{
		create: gccb,
	}
}

// GoodCoinUpsertBulk is the builder for "upsert"-ing
// a bulk of GoodCoin nodes.
type GoodCoinUpsertBulk struct {
	create *GoodCoinCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GoodCoin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(goodcoin.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *GoodCoinUpsertBulk) UpdateNewValues() *GoodCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(goodcoin.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GoodCoin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *GoodCoinUpsertBulk) Ignore() *GoodCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GoodCoinUpsertBulk) DoNothing() *GoodCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GoodCoinCreateBulk.OnConflict
// documentation for more info.
func (u *GoodCoinUpsertBulk) Update(set func(*GoodCoinUpsert)) *GoodCoinUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GoodCoinUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *GoodCoinUpsertBulk) SetCreatedAt(v uint32) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *GoodCoinUpsertBulk) AddCreatedAt(v uint32) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *GoodCoinUpsertBulk) UpdateCreatedAt() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *GoodCoinUpsertBulk) SetUpdatedAt(v uint32) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *GoodCoinUpsertBulk) AddUpdatedAt(v uint32) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *GoodCoinUpsertBulk) UpdateUpdatedAt() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *GoodCoinUpsertBulk) SetDeletedAt(v uint32) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *GoodCoinUpsertBulk) AddDeletedAt(v uint32) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *GoodCoinUpsertBulk) UpdateDeletedAt() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *GoodCoinUpsertBulk) SetEntID(v uuid.UUID) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *GoodCoinUpsertBulk) UpdateEntID() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateEntID()
	})
}

// SetGoodID sets the "good_id" field.
func (u *GoodCoinUpsertBulk) SetGoodID(v uuid.UUID) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *GoodCoinUpsertBulk) UpdateGoodID() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateGoodID()
	})
}

// ClearGoodID clears the value of the "good_id" field.
func (u *GoodCoinUpsertBulk) ClearGoodID() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.ClearGoodID()
	})
}

// SetCoinTypeID sets the "coin_type_id" field.
func (u *GoodCoinUpsertBulk) SetCoinTypeID(v uuid.UUID) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetCoinTypeID(v)
	})
}

// UpdateCoinTypeID sets the "coin_type_id" field to the value that was provided on create.
func (u *GoodCoinUpsertBulk) UpdateCoinTypeID() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateCoinTypeID()
	})
}

// ClearCoinTypeID clears the value of the "coin_type_id" field.
func (u *GoodCoinUpsertBulk) ClearCoinTypeID() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.ClearCoinTypeID()
	})
}

// SetMain sets the "main" field.
func (u *GoodCoinUpsertBulk) SetMain(v bool) *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.SetMain(v)
	})
}

// UpdateMain sets the "main" field to the value that was provided on create.
func (u *GoodCoinUpsertBulk) UpdateMain() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.UpdateMain()
	})
}

// ClearMain clears the value of the "main" field.
func (u *GoodCoinUpsertBulk) ClearMain() *GoodCoinUpsertBulk {
	return u.Update(func(s *GoodCoinUpsert) {
		s.ClearMain()
	})
}

// Exec executes the query.
func (u *GoodCoinUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GoodCoinCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GoodCoinCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GoodCoinUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
