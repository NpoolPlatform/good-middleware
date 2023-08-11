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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// StockCreate is the builder for creating a Stock entity.
type StockCreate struct {
	config
	mutation *StockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (sc *StockCreate) SetCreatedAt(u uint32) *StockCreate {
	sc.mutation.SetCreatedAt(u)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *StockCreate) SetNillableCreatedAt(u *uint32) *StockCreate {
	if u != nil {
		sc.SetCreatedAt(*u)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StockCreate) SetUpdatedAt(u uint32) *StockCreate {
	sc.mutation.SetUpdatedAt(u)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StockCreate) SetNillableUpdatedAt(u *uint32) *StockCreate {
	if u != nil {
		sc.SetUpdatedAt(*u)
	}
	return sc
}

// SetDeletedAt sets the "deleted_at" field.
func (sc *StockCreate) SetDeletedAt(u uint32) *StockCreate {
	sc.mutation.SetDeletedAt(u)
	return sc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sc *StockCreate) SetNillableDeletedAt(u *uint32) *StockCreate {
	if u != nil {
		sc.SetDeletedAt(*u)
	}
	return sc
}

// SetGoodID sets the "good_id" field.
func (sc *StockCreate) SetGoodID(u uuid.UUID) *StockCreate {
	sc.mutation.SetGoodID(u)
	return sc
}

// SetTotal sets the "total" field.
func (sc *StockCreate) SetTotal(d decimal.Decimal) *StockCreate {
	sc.mutation.SetTotal(d)
	return sc
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (sc *StockCreate) SetNillableTotal(d *decimal.Decimal) *StockCreate {
	if d != nil {
		sc.SetTotal(*d)
	}
	return sc
}

// SetLocked sets the "locked" field.
func (sc *StockCreate) SetLocked(d decimal.Decimal) *StockCreate {
	sc.mutation.SetLocked(d)
	return sc
}

// SetNillableLocked sets the "locked" field if the given value is not nil.
func (sc *StockCreate) SetNillableLocked(d *decimal.Decimal) *StockCreate {
	if d != nil {
		sc.SetLocked(*d)
	}
	return sc
}

// SetInService sets the "in_service" field.
func (sc *StockCreate) SetInService(d decimal.Decimal) *StockCreate {
	sc.mutation.SetInService(d)
	return sc
}

// SetNillableInService sets the "in_service" field if the given value is not nil.
func (sc *StockCreate) SetNillableInService(d *decimal.Decimal) *StockCreate {
	if d != nil {
		sc.SetInService(*d)
	}
	return sc
}

// SetWaitStart sets the "wait_start" field.
func (sc *StockCreate) SetWaitStart(d decimal.Decimal) *StockCreate {
	sc.mutation.SetWaitStart(d)
	return sc
}

// SetNillableWaitStart sets the "wait_start" field if the given value is not nil.
func (sc *StockCreate) SetNillableWaitStart(d *decimal.Decimal) *StockCreate {
	if d != nil {
		sc.SetWaitStart(*d)
	}
	return sc
}

// SetSold sets the "sold" field.
func (sc *StockCreate) SetSold(d decimal.Decimal) *StockCreate {
	sc.mutation.SetSold(d)
	return sc
}

// SetNillableSold sets the "sold" field if the given value is not nil.
func (sc *StockCreate) SetNillableSold(d *decimal.Decimal) *StockCreate {
	if d != nil {
		sc.SetSold(*d)
	}
	return sc
}

// SetChannelLocked sets the "channel_locked" field.
func (sc *StockCreate) SetChannelLocked(d decimal.Decimal) *StockCreate {
	sc.mutation.SetChannelLocked(d)
	return sc
}

// SetNillableChannelLocked sets the "channel_locked" field if the given value is not nil.
func (sc *StockCreate) SetNillableChannelLocked(d *decimal.Decimal) *StockCreate {
	if d != nil {
		sc.SetChannelLocked(*d)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StockCreate) SetID(u uuid.UUID) *StockCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *StockCreate) SetNillableID(u *uuid.UUID) *StockCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// Mutation returns the StockMutation object of the builder.
func (sc *StockCreate) Mutation() *StockMutation {
	return sc.mutation
}

// Save creates the Stock in the database.
func (sc *StockCreate) Save(ctx context.Context) (*Stock, error) {
	var (
		err  error
		node *Stock
	)
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*StockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Stock)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from StockMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StockCreate) SaveX(ctx context.Context) *Stock {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StockCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StockCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StockCreate) defaults() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		if stock.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized stock.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := stock.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		if stock.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized stock.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := stock.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		if stock.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized stock.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := stock.DefaultDeletedAt()
		sc.mutation.SetDeletedAt(v)
	}
	if _, ok := sc.mutation.Total(); !ok {
		v := stock.DefaultTotal
		sc.mutation.SetTotal(v)
	}
	if _, ok := sc.mutation.Locked(); !ok {
		v := stock.DefaultLocked
		sc.mutation.SetLocked(v)
	}
	if _, ok := sc.mutation.InService(); !ok {
		v := stock.DefaultInService
		sc.mutation.SetInService(v)
	}
	if _, ok := sc.mutation.WaitStart(); !ok {
		v := stock.DefaultWaitStart
		sc.mutation.SetWaitStart(v)
	}
	if _, ok := sc.mutation.Sold(); !ok {
		v := stock.DefaultSold
		sc.mutation.SetSold(v)
	}
	if _, ok := sc.mutation.ChannelLocked(); !ok {
		v := stock.DefaultChannelLocked
		sc.mutation.SetChannelLocked(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		if stock.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized stock.DefaultID (forgotten import ent/runtime?)")
		}
		v := stock.DefaultID()
		sc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *StockCreate) check() error {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Stock.created_at"`)}
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Stock.updated_at"`)}
	}
	if _, ok := sc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Stock.deleted_at"`)}
	}
	if _, ok := sc.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "Stock.good_id"`)}
	}
	return nil
}

func (sc *StockCreate) sqlSave(ctx context.Context) (*Stock, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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

func (sc *StockCreate) createSpec() (*Stock, *sqlgraph.CreateSpec) {
	var (
		_node = &Stock{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: stock.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: stock.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := sc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: stock.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := sc.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: stock.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := sc.mutation.Total(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldTotal,
		})
		_node.Total = value
	}
	if value, ok := sc.mutation.Locked(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldLocked,
		})
		_node.Locked = value
	}
	if value, ok := sc.mutation.InService(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldInService,
		})
		_node.InService = value
	}
	if value, ok := sc.mutation.WaitStart(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldWaitStart,
		})
		_node.WaitStart = value
	}
	if value, ok := sc.mutation.Sold(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldSold,
		})
		_node.Sold = value
	}
	if value, ok := sc.mutation.ChannelLocked(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: stock.FieldChannelLocked,
		})
		_node.ChannelLocked = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Stock.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (sc *StockCreate) OnConflict(opts ...sql.ConflictOption) *StockUpsertOne {
	sc.conflict = opts
	return &StockUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Stock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *StockCreate) OnConflictColumns(columns ...string) *StockUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &StockUpsertOne{
		create: sc,
	}
}

type (
	// StockUpsertOne is the builder for "upsert"-ing
	//  one Stock node.
	StockUpsertOne struct {
		create *StockCreate
	}

	// StockUpsert is the "OnConflict" setter.
	StockUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *StockUpsert) SetCreatedAt(v uint32) *StockUpsert {
	u.Set(stock.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StockUpsert) UpdateCreatedAt() *StockUpsert {
	u.SetExcluded(stock.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *StockUpsert) AddCreatedAt(v uint32) *StockUpsert {
	u.Add(stock.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StockUpsert) SetUpdatedAt(v uint32) *StockUpsert {
	u.Set(stock.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StockUpsert) UpdateUpdatedAt() *StockUpsert {
	u.SetExcluded(stock.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *StockUpsert) AddUpdatedAt(v uint32) *StockUpsert {
	u.Add(stock.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *StockUpsert) SetDeletedAt(v uint32) *StockUpsert {
	u.Set(stock.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *StockUpsert) UpdateDeletedAt() *StockUpsert {
	u.SetExcluded(stock.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *StockUpsert) AddDeletedAt(v uint32) *StockUpsert {
	u.Add(stock.FieldDeletedAt, v)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *StockUpsert) SetGoodID(v uuid.UUID) *StockUpsert {
	u.Set(stock.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *StockUpsert) UpdateGoodID() *StockUpsert {
	u.SetExcluded(stock.FieldGoodID)
	return u
}

// SetTotal sets the "total" field.
func (u *StockUpsert) SetTotal(v decimal.Decimal) *StockUpsert {
	u.Set(stock.FieldTotal, v)
	return u
}

// UpdateTotal sets the "total" field to the value that was provided on create.
func (u *StockUpsert) UpdateTotal() *StockUpsert {
	u.SetExcluded(stock.FieldTotal)
	return u
}

// ClearTotal clears the value of the "total" field.
func (u *StockUpsert) ClearTotal() *StockUpsert {
	u.SetNull(stock.FieldTotal)
	return u
}

// SetLocked sets the "locked" field.
func (u *StockUpsert) SetLocked(v decimal.Decimal) *StockUpsert {
	u.Set(stock.FieldLocked, v)
	return u
}

// UpdateLocked sets the "locked" field to the value that was provided on create.
func (u *StockUpsert) UpdateLocked() *StockUpsert {
	u.SetExcluded(stock.FieldLocked)
	return u
}

// ClearLocked clears the value of the "locked" field.
func (u *StockUpsert) ClearLocked() *StockUpsert {
	u.SetNull(stock.FieldLocked)
	return u
}

// SetInService sets the "in_service" field.
func (u *StockUpsert) SetInService(v decimal.Decimal) *StockUpsert {
	u.Set(stock.FieldInService, v)
	return u
}

// UpdateInService sets the "in_service" field to the value that was provided on create.
func (u *StockUpsert) UpdateInService() *StockUpsert {
	u.SetExcluded(stock.FieldInService)
	return u
}

// ClearInService clears the value of the "in_service" field.
func (u *StockUpsert) ClearInService() *StockUpsert {
	u.SetNull(stock.FieldInService)
	return u
}

// SetWaitStart sets the "wait_start" field.
func (u *StockUpsert) SetWaitStart(v decimal.Decimal) *StockUpsert {
	u.Set(stock.FieldWaitStart, v)
	return u
}

// UpdateWaitStart sets the "wait_start" field to the value that was provided on create.
func (u *StockUpsert) UpdateWaitStart() *StockUpsert {
	u.SetExcluded(stock.FieldWaitStart)
	return u
}

// ClearWaitStart clears the value of the "wait_start" field.
func (u *StockUpsert) ClearWaitStart() *StockUpsert {
	u.SetNull(stock.FieldWaitStart)
	return u
}

// SetSold sets the "sold" field.
func (u *StockUpsert) SetSold(v decimal.Decimal) *StockUpsert {
	u.Set(stock.FieldSold, v)
	return u
}

// UpdateSold sets the "sold" field to the value that was provided on create.
func (u *StockUpsert) UpdateSold() *StockUpsert {
	u.SetExcluded(stock.FieldSold)
	return u
}

// ClearSold clears the value of the "sold" field.
func (u *StockUpsert) ClearSold() *StockUpsert {
	u.SetNull(stock.FieldSold)
	return u
}

// SetChannelLocked sets the "channel_locked" field.
func (u *StockUpsert) SetChannelLocked(v decimal.Decimal) *StockUpsert {
	u.Set(stock.FieldChannelLocked, v)
	return u
}

// UpdateChannelLocked sets the "channel_locked" field to the value that was provided on create.
func (u *StockUpsert) UpdateChannelLocked() *StockUpsert {
	u.SetExcluded(stock.FieldChannelLocked)
	return u
}

// ClearChannelLocked clears the value of the "channel_locked" field.
func (u *StockUpsert) ClearChannelLocked() *StockUpsert {
	u.SetNull(stock.FieldChannelLocked)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Stock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(stock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *StockUpsertOne) UpdateNewValues() *StockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(stock.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Stock.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *StockUpsertOne) Ignore() *StockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StockUpsertOne) DoNothing() *StockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StockCreate.OnConflict
// documentation for more info.
func (u *StockUpsertOne) Update(set func(*StockUpsert)) *StockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StockUpsertOne) SetCreatedAt(v uint32) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *StockUpsertOne) AddCreatedAt(v uint32) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateCreatedAt() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StockUpsertOne) SetUpdatedAt(v uint32) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *StockUpsertOne) AddUpdatedAt(v uint32) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateUpdatedAt() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *StockUpsertOne) SetDeletedAt(v uint32) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *StockUpsertOne) AddDeletedAt(v uint32) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateDeletedAt() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *StockUpsertOne) SetGoodID(v uuid.UUID) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateGoodID() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateGoodID()
	})
}

// SetTotal sets the "total" field.
func (u *StockUpsertOne) SetTotal(v decimal.Decimal) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetTotal(v)
	})
}

// UpdateTotal sets the "total" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateTotal() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateTotal()
	})
}

// ClearTotal clears the value of the "total" field.
func (u *StockUpsertOne) ClearTotal() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.ClearTotal()
	})
}

// SetLocked sets the "locked" field.
func (u *StockUpsertOne) SetLocked(v decimal.Decimal) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetLocked(v)
	})
}

// UpdateLocked sets the "locked" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateLocked() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateLocked()
	})
}

// ClearLocked clears the value of the "locked" field.
func (u *StockUpsertOne) ClearLocked() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.ClearLocked()
	})
}

// SetInService sets the "in_service" field.
func (u *StockUpsertOne) SetInService(v decimal.Decimal) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetInService(v)
	})
}

// UpdateInService sets the "in_service" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateInService() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateInService()
	})
}

// ClearInService clears the value of the "in_service" field.
func (u *StockUpsertOne) ClearInService() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.ClearInService()
	})
}

// SetWaitStart sets the "wait_start" field.
func (u *StockUpsertOne) SetWaitStart(v decimal.Decimal) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetWaitStart(v)
	})
}

// UpdateWaitStart sets the "wait_start" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateWaitStart() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateWaitStart()
	})
}

// ClearWaitStart clears the value of the "wait_start" field.
func (u *StockUpsertOne) ClearWaitStart() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.ClearWaitStart()
	})
}

// SetSold sets the "sold" field.
func (u *StockUpsertOne) SetSold(v decimal.Decimal) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetSold(v)
	})
}

// UpdateSold sets the "sold" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateSold() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateSold()
	})
}

// ClearSold clears the value of the "sold" field.
func (u *StockUpsertOne) ClearSold() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.ClearSold()
	})
}

// SetChannelLocked sets the "channel_locked" field.
func (u *StockUpsertOne) SetChannelLocked(v decimal.Decimal) *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.SetChannelLocked(v)
	})
}

// UpdateChannelLocked sets the "channel_locked" field to the value that was provided on create.
func (u *StockUpsertOne) UpdateChannelLocked() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.UpdateChannelLocked()
	})
}

// ClearChannelLocked clears the value of the "channel_locked" field.
func (u *StockUpsertOne) ClearChannelLocked() *StockUpsertOne {
	return u.Update(func(s *StockUpsert) {
		s.ClearChannelLocked()
	})
}

// Exec executes the query.
func (u *StockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *StockUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: StockUpsertOne.ID is not supported by MySQL driver. Use StockUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *StockUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// StockCreateBulk is the builder for creating many Stock entities in bulk.
type StockCreateBulk struct {
	config
	builders []*StockCreate
	conflict []sql.ConflictOption
}

// Save creates the Stock entities in the database.
func (scb *StockCreateBulk) Save(ctx context.Context) ([]*Stock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stock, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StockMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StockCreateBulk) SaveX(ctx context.Context) []*Stock {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StockCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StockCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Stock.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.StockUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (scb *StockCreateBulk) OnConflict(opts ...sql.ConflictOption) *StockUpsertBulk {
	scb.conflict = opts
	return &StockUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Stock.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *StockCreateBulk) OnConflictColumns(columns ...string) *StockUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &StockUpsertBulk{
		create: scb,
	}
}

// StockUpsertBulk is the builder for "upsert"-ing
// a bulk of Stock nodes.
type StockUpsertBulk struct {
	create *StockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Stock.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(stock.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *StockUpsertBulk) UpdateNewValues() *StockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(stock.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Stock.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *StockUpsertBulk) Ignore() *StockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *StockUpsertBulk) DoNothing() *StockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the StockCreateBulk.OnConflict
// documentation for more info.
func (u *StockUpsertBulk) Update(set func(*StockUpsert)) *StockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&StockUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *StockUpsertBulk) SetCreatedAt(v uint32) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *StockUpsertBulk) AddCreatedAt(v uint32) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateCreatedAt() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *StockUpsertBulk) SetUpdatedAt(v uint32) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *StockUpsertBulk) AddUpdatedAt(v uint32) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateUpdatedAt() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *StockUpsertBulk) SetDeletedAt(v uint32) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *StockUpsertBulk) AddDeletedAt(v uint32) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateDeletedAt() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *StockUpsertBulk) SetGoodID(v uuid.UUID) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateGoodID() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateGoodID()
	})
}

// SetTotal sets the "total" field.
func (u *StockUpsertBulk) SetTotal(v decimal.Decimal) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetTotal(v)
	})
}

// UpdateTotal sets the "total" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateTotal() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateTotal()
	})
}

// ClearTotal clears the value of the "total" field.
func (u *StockUpsertBulk) ClearTotal() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.ClearTotal()
	})
}

// SetLocked sets the "locked" field.
func (u *StockUpsertBulk) SetLocked(v decimal.Decimal) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetLocked(v)
	})
}

// UpdateLocked sets the "locked" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateLocked() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateLocked()
	})
}

// ClearLocked clears the value of the "locked" field.
func (u *StockUpsertBulk) ClearLocked() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.ClearLocked()
	})
}

// SetInService sets the "in_service" field.
func (u *StockUpsertBulk) SetInService(v decimal.Decimal) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetInService(v)
	})
}

// UpdateInService sets the "in_service" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateInService() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateInService()
	})
}

// ClearInService clears the value of the "in_service" field.
func (u *StockUpsertBulk) ClearInService() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.ClearInService()
	})
}

// SetWaitStart sets the "wait_start" field.
func (u *StockUpsertBulk) SetWaitStart(v decimal.Decimal) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetWaitStart(v)
	})
}

// UpdateWaitStart sets the "wait_start" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateWaitStart() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateWaitStart()
	})
}

// ClearWaitStart clears the value of the "wait_start" field.
func (u *StockUpsertBulk) ClearWaitStart() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.ClearWaitStart()
	})
}

// SetSold sets the "sold" field.
func (u *StockUpsertBulk) SetSold(v decimal.Decimal) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetSold(v)
	})
}

// UpdateSold sets the "sold" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateSold() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateSold()
	})
}

// ClearSold clears the value of the "sold" field.
func (u *StockUpsertBulk) ClearSold() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.ClearSold()
	})
}

// SetChannelLocked sets the "channel_locked" field.
func (u *StockUpsertBulk) SetChannelLocked(v decimal.Decimal) *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.SetChannelLocked(v)
	})
}

// UpdateChannelLocked sets the "channel_locked" field to the value that was provided on create.
func (u *StockUpsertBulk) UpdateChannelLocked() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.UpdateChannelLocked()
	})
}

// ClearChannelLocked clears the value of the "channel_locked" field.
func (u *StockUpsertBulk) ClearChannelLocked() *StockUpsertBulk {
	return u.Update(func(s *StockUpsert) {
		s.ClearChannelLocked()
	})
}

// Exec executes the query.
func (u *StockUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the StockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for StockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *StockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
