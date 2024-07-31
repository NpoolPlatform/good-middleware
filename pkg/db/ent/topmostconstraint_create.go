// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostconstraint"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// TopMostConstraintCreate is the builder for creating a TopMostConstraint entity.
type TopMostConstraintCreate struct {
	config
	mutation *TopMostConstraintMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (tmcc *TopMostConstraintCreate) SetCreatedAt(u uint32) *TopMostConstraintCreate {
	tmcc.mutation.SetCreatedAt(u)
	return tmcc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableCreatedAt(u *uint32) *TopMostConstraintCreate {
	if u != nil {
		tmcc.SetCreatedAt(*u)
	}
	return tmcc
}

// SetUpdatedAt sets the "updated_at" field.
func (tmcc *TopMostConstraintCreate) SetUpdatedAt(u uint32) *TopMostConstraintCreate {
	tmcc.mutation.SetUpdatedAt(u)
	return tmcc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableUpdatedAt(u *uint32) *TopMostConstraintCreate {
	if u != nil {
		tmcc.SetUpdatedAt(*u)
	}
	return tmcc
}

// SetDeletedAt sets the "deleted_at" field.
func (tmcc *TopMostConstraintCreate) SetDeletedAt(u uint32) *TopMostConstraintCreate {
	tmcc.mutation.SetDeletedAt(u)
	return tmcc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableDeletedAt(u *uint32) *TopMostConstraintCreate {
	if u != nil {
		tmcc.SetDeletedAt(*u)
	}
	return tmcc
}

// SetEntID sets the "ent_id" field.
func (tmcc *TopMostConstraintCreate) SetEntID(u uuid.UUID) *TopMostConstraintCreate {
	tmcc.mutation.SetEntID(u)
	return tmcc
}

// SetNillableEntID sets the "ent_id" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableEntID(u *uuid.UUID) *TopMostConstraintCreate {
	if u != nil {
		tmcc.SetEntID(*u)
	}
	return tmcc
}

// SetTopMostID sets the "top_most_id" field.
func (tmcc *TopMostConstraintCreate) SetTopMostID(u uuid.UUID) *TopMostConstraintCreate {
	tmcc.mutation.SetTopMostID(u)
	return tmcc
}

// SetNillableTopMostID sets the "top_most_id" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableTopMostID(u *uuid.UUID) *TopMostConstraintCreate {
	if u != nil {
		tmcc.SetTopMostID(*u)
	}
	return tmcc
}

// SetConstraint sets the "constraint" field.
func (tmcc *TopMostConstraintCreate) SetConstraint(s string) *TopMostConstraintCreate {
	tmcc.mutation.SetConstraint(s)
	return tmcc
}

// SetNillableConstraint sets the "constraint" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableConstraint(s *string) *TopMostConstraintCreate {
	if s != nil {
		tmcc.SetConstraint(*s)
	}
	return tmcc
}

// SetTargetValue sets the "target_value" field.
func (tmcc *TopMostConstraintCreate) SetTargetValue(d decimal.Decimal) *TopMostConstraintCreate {
	tmcc.mutation.SetTargetValue(d)
	return tmcc
}

// SetNillableTargetValue sets the "target_value" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableTargetValue(d *decimal.Decimal) *TopMostConstraintCreate {
	if d != nil {
		tmcc.SetTargetValue(*d)
	}
	return tmcc
}

// SetIndex sets the "index" field.
func (tmcc *TopMostConstraintCreate) SetIndex(u uint8) *TopMostConstraintCreate {
	tmcc.mutation.SetIndex(u)
	return tmcc
}

// SetNillableIndex sets the "index" field if the given value is not nil.
func (tmcc *TopMostConstraintCreate) SetNillableIndex(u *uint8) *TopMostConstraintCreate {
	if u != nil {
		tmcc.SetIndex(*u)
	}
	return tmcc
}

// SetID sets the "id" field.
func (tmcc *TopMostConstraintCreate) SetID(u uint32) *TopMostConstraintCreate {
	tmcc.mutation.SetID(u)
	return tmcc
}

// Mutation returns the TopMostConstraintMutation object of the builder.
func (tmcc *TopMostConstraintCreate) Mutation() *TopMostConstraintMutation {
	return tmcc.mutation
}

// Save creates the TopMostConstraint in the database.
func (tmcc *TopMostConstraintCreate) Save(ctx context.Context) (*TopMostConstraint, error) {
	var (
		err  error
		node *TopMostConstraint
	)
	if err := tmcc.defaults(); err != nil {
		return nil, err
	}
	if len(tmcc.hooks) == 0 {
		if err = tmcc.check(); err != nil {
			return nil, err
		}
		node, err = tmcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TopMostConstraintMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tmcc.check(); err != nil {
				return nil, err
			}
			tmcc.mutation = mutation
			if node, err = tmcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tmcc.hooks) - 1; i >= 0; i-- {
			if tmcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tmcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tmcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*TopMostConstraint)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TopMostConstraintMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tmcc *TopMostConstraintCreate) SaveX(ctx context.Context) *TopMostConstraint {
	v, err := tmcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmcc *TopMostConstraintCreate) Exec(ctx context.Context) error {
	_, err := tmcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmcc *TopMostConstraintCreate) ExecX(ctx context.Context) {
	if err := tmcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tmcc *TopMostConstraintCreate) defaults() error {
	if _, ok := tmcc.mutation.CreatedAt(); !ok {
		if topmostconstraint.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostconstraint.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := topmostconstraint.DefaultCreatedAt()
		tmcc.mutation.SetCreatedAt(v)
	}
	if _, ok := tmcc.mutation.UpdatedAt(); !ok {
		if topmostconstraint.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostconstraint.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := topmostconstraint.DefaultUpdatedAt()
		tmcc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tmcc.mutation.DeletedAt(); !ok {
		if topmostconstraint.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized topmostconstraint.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := topmostconstraint.DefaultDeletedAt()
		tmcc.mutation.SetDeletedAt(v)
	}
	if _, ok := tmcc.mutation.EntID(); !ok {
		if topmostconstraint.DefaultEntID == nil {
			return fmt.Errorf("ent: uninitialized topmostconstraint.DefaultEntID (forgotten import ent/runtime?)")
		}
		v := topmostconstraint.DefaultEntID()
		tmcc.mutation.SetEntID(v)
	}
	if _, ok := tmcc.mutation.TopMostID(); !ok {
		if topmostconstraint.DefaultTopMostID == nil {
			return fmt.Errorf("ent: uninitialized topmostconstraint.DefaultTopMostID (forgotten import ent/runtime?)")
		}
		v := topmostconstraint.DefaultTopMostID()
		tmcc.mutation.SetTopMostID(v)
	}
	if _, ok := tmcc.mutation.Constraint(); !ok {
		v := topmostconstraint.DefaultConstraint
		tmcc.mutation.SetConstraint(v)
	}
	if _, ok := tmcc.mutation.TargetValue(); !ok {
		v := topmostconstraint.DefaultTargetValue
		tmcc.mutation.SetTargetValue(v)
	}
	if _, ok := tmcc.mutation.Index(); !ok {
		v := topmostconstraint.DefaultIndex
		tmcc.mutation.SetIndex(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tmcc *TopMostConstraintCreate) check() error {
	if _, ok := tmcc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "TopMostConstraint.created_at"`)}
	}
	if _, ok := tmcc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "TopMostConstraint.updated_at"`)}
	}
	if _, ok := tmcc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "TopMostConstraint.deleted_at"`)}
	}
	if _, ok := tmcc.mutation.EntID(); !ok {
		return &ValidationError{Name: "ent_id", err: errors.New(`ent: missing required field "TopMostConstraint.ent_id"`)}
	}
	return nil
}

func (tmcc *TopMostConstraintCreate) sqlSave(ctx context.Context) (*TopMostConstraint, error) {
	_node, _spec := tmcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tmcc.driver, _spec); err != nil {
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

func (tmcc *TopMostConstraintCreate) createSpec() (*TopMostConstraint, *sqlgraph.CreateSpec) {
	var (
		_node = &TopMostConstraint{config: tmcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: topmostconstraint.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostconstraint.FieldID,
			},
		}
	)
	_spec.OnConflict = tmcc.conflict
	if id, ok := tmcc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tmcc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := tmcc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := tmcc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: topmostconstraint.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := tmcc.mutation.EntID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostconstraint.FieldEntID,
		})
		_node.EntID = value
	}
	if value, ok := tmcc.mutation.TopMostID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: topmostconstraint.FieldTopMostID,
		})
		_node.TopMostID = value
	}
	if value, ok := tmcc.mutation.Constraint(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: topmostconstraint.FieldConstraint,
		})
		_node.Constraint = value
	}
	if value, ok := tmcc.mutation.TargetValue(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: topmostconstraint.FieldTargetValue,
		})
		_node.TargetValue = value
	}
	if value, ok := tmcc.mutation.Index(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint8,
			Value:  value,
			Column: topmostconstraint.FieldIndex,
		})
		_node.Index = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TopMostConstraint.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TopMostConstraintUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tmcc *TopMostConstraintCreate) OnConflict(opts ...sql.ConflictOption) *TopMostConstraintUpsertOne {
	tmcc.conflict = opts
	return &TopMostConstraintUpsertOne{
		create: tmcc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TopMostConstraint.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tmcc *TopMostConstraintCreate) OnConflictColumns(columns ...string) *TopMostConstraintUpsertOne {
	tmcc.conflict = append(tmcc.conflict, sql.ConflictColumns(columns...))
	return &TopMostConstraintUpsertOne{
		create: tmcc,
	}
}

type (
	// TopMostConstraintUpsertOne is the builder for "upsert"-ing
	//  one TopMostConstraint node.
	TopMostConstraintUpsertOne struct {
		create *TopMostConstraintCreate
	}

	// TopMostConstraintUpsert is the "OnConflict" setter.
	TopMostConstraintUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *TopMostConstraintUpsert) SetCreatedAt(v uint32) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateCreatedAt() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TopMostConstraintUpsert) AddCreatedAt(v uint32) *TopMostConstraintUpsert {
	u.Add(topmostconstraint.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TopMostConstraintUpsert) SetUpdatedAt(v uint32) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateUpdatedAt() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TopMostConstraintUpsert) AddUpdatedAt(v uint32) *TopMostConstraintUpsert {
	u.Add(topmostconstraint.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TopMostConstraintUpsert) SetDeletedAt(v uint32) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateDeletedAt() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TopMostConstraintUpsert) AddDeletedAt(v uint32) *TopMostConstraintUpsert {
	u.Add(topmostconstraint.FieldDeletedAt, v)
	return u
}

// SetEntID sets the "ent_id" field.
func (u *TopMostConstraintUpsert) SetEntID(v uuid.UUID) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldEntID, v)
	return u
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateEntID() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldEntID)
	return u
}

// SetTopMostID sets the "top_most_id" field.
func (u *TopMostConstraintUpsert) SetTopMostID(v uuid.UUID) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldTopMostID, v)
	return u
}

// UpdateTopMostID sets the "top_most_id" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateTopMostID() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldTopMostID)
	return u
}

// ClearTopMostID clears the value of the "top_most_id" field.
func (u *TopMostConstraintUpsert) ClearTopMostID() *TopMostConstraintUpsert {
	u.SetNull(topmostconstraint.FieldTopMostID)
	return u
}

// SetConstraint sets the "constraint" field.
func (u *TopMostConstraintUpsert) SetConstraint(v string) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldConstraint, v)
	return u
}

// UpdateConstraint sets the "constraint" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateConstraint() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldConstraint)
	return u
}

// ClearConstraint clears the value of the "constraint" field.
func (u *TopMostConstraintUpsert) ClearConstraint() *TopMostConstraintUpsert {
	u.SetNull(topmostconstraint.FieldConstraint)
	return u
}

// SetTargetValue sets the "target_value" field.
func (u *TopMostConstraintUpsert) SetTargetValue(v decimal.Decimal) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldTargetValue, v)
	return u
}

// UpdateTargetValue sets the "target_value" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateTargetValue() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldTargetValue)
	return u
}

// ClearTargetValue clears the value of the "target_value" field.
func (u *TopMostConstraintUpsert) ClearTargetValue() *TopMostConstraintUpsert {
	u.SetNull(topmostconstraint.FieldTargetValue)
	return u
}

// SetIndex sets the "index" field.
func (u *TopMostConstraintUpsert) SetIndex(v uint8) *TopMostConstraintUpsert {
	u.Set(topmostconstraint.FieldIndex, v)
	return u
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *TopMostConstraintUpsert) UpdateIndex() *TopMostConstraintUpsert {
	u.SetExcluded(topmostconstraint.FieldIndex)
	return u
}

// AddIndex adds v to the "index" field.
func (u *TopMostConstraintUpsert) AddIndex(v uint8) *TopMostConstraintUpsert {
	u.Add(topmostconstraint.FieldIndex, v)
	return u
}

// ClearIndex clears the value of the "index" field.
func (u *TopMostConstraintUpsert) ClearIndex() *TopMostConstraintUpsert {
	u.SetNull(topmostconstraint.FieldIndex)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.TopMostConstraint.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(topmostconstraint.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TopMostConstraintUpsertOne) UpdateNewValues() *TopMostConstraintUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(topmostconstraint.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TopMostConstraint.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TopMostConstraintUpsertOne) Ignore() *TopMostConstraintUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TopMostConstraintUpsertOne) DoNothing() *TopMostConstraintUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TopMostConstraintCreate.OnConflict
// documentation for more info.
func (u *TopMostConstraintUpsertOne) Update(set func(*TopMostConstraintUpsert)) *TopMostConstraintUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TopMostConstraintUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TopMostConstraintUpsertOne) SetCreatedAt(v uint32) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TopMostConstraintUpsertOne) AddCreatedAt(v uint32) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateCreatedAt() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TopMostConstraintUpsertOne) SetUpdatedAt(v uint32) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TopMostConstraintUpsertOne) AddUpdatedAt(v uint32) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateUpdatedAt() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TopMostConstraintUpsertOne) SetDeletedAt(v uint32) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TopMostConstraintUpsertOne) AddDeletedAt(v uint32) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateDeletedAt() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *TopMostConstraintUpsertOne) SetEntID(v uuid.UUID) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateEntID() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateEntID()
	})
}

// SetTopMostID sets the "top_most_id" field.
func (u *TopMostConstraintUpsertOne) SetTopMostID(v uuid.UUID) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetTopMostID(v)
	})
}

// UpdateTopMostID sets the "top_most_id" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateTopMostID() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateTopMostID()
	})
}

// ClearTopMostID clears the value of the "top_most_id" field.
func (u *TopMostConstraintUpsertOne) ClearTopMostID() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearTopMostID()
	})
}

// SetConstraint sets the "constraint" field.
func (u *TopMostConstraintUpsertOne) SetConstraint(v string) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetConstraint(v)
	})
}

// UpdateConstraint sets the "constraint" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateConstraint() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateConstraint()
	})
}

// ClearConstraint clears the value of the "constraint" field.
func (u *TopMostConstraintUpsertOne) ClearConstraint() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearConstraint()
	})
}

// SetTargetValue sets the "target_value" field.
func (u *TopMostConstraintUpsertOne) SetTargetValue(v decimal.Decimal) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetTargetValue(v)
	})
}

// UpdateTargetValue sets the "target_value" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateTargetValue() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateTargetValue()
	})
}

// ClearTargetValue clears the value of the "target_value" field.
func (u *TopMostConstraintUpsertOne) ClearTargetValue() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearTargetValue()
	})
}

// SetIndex sets the "index" field.
func (u *TopMostConstraintUpsertOne) SetIndex(v uint8) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *TopMostConstraintUpsertOne) AddIndex(v uint8) *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *TopMostConstraintUpsertOne) UpdateIndex() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *TopMostConstraintUpsertOne) ClearIndex() *TopMostConstraintUpsertOne {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *TopMostConstraintUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TopMostConstraintCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TopMostConstraintUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TopMostConstraintUpsertOne) ID(ctx context.Context) (id uint32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TopMostConstraintUpsertOne) IDX(ctx context.Context) uint32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TopMostConstraintCreateBulk is the builder for creating many TopMostConstraint entities in bulk.
type TopMostConstraintCreateBulk struct {
	config
	builders []*TopMostConstraintCreate
	conflict []sql.ConflictOption
}

// Save creates the TopMostConstraint entities in the database.
func (tmccb *TopMostConstraintCreateBulk) Save(ctx context.Context) ([]*TopMostConstraint, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tmccb.builders))
	nodes := make([]*TopMostConstraint, len(tmccb.builders))
	mutators := make([]Mutator, len(tmccb.builders))
	for i := range tmccb.builders {
		func(i int, root context.Context) {
			builder := tmccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TopMostConstraintMutation)
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
					_, err = mutators[i+1].Mutate(root, tmccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tmccb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tmccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tmccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tmccb *TopMostConstraintCreateBulk) SaveX(ctx context.Context) []*TopMostConstraint {
	v, err := tmccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tmccb *TopMostConstraintCreateBulk) Exec(ctx context.Context) error {
	_, err := tmccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tmccb *TopMostConstraintCreateBulk) ExecX(ctx context.Context) {
	if err := tmccb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.TopMostConstraint.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TopMostConstraintUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (tmccb *TopMostConstraintCreateBulk) OnConflict(opts ...sql.ConflictOption) *TopMostConstraintUpsertBulk {
	tmccb.conflict = opts
	return &TopMostConstraintUpsertBulk{
		create: tmccb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.TopMostConstraint.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tmccb *TopMostConstraintCreateBulk) OnConflictColumns(columns ...string) *TopMostConstraintUpsertBulk {
	tmccb.conflict = append(tmccb.conflict, sql.ConflictColumns(columns...))
	return &TopMostConstraintUpsertBulk{
		create: tmccb,
	}
}

// TopMostConstraintUpsertBulk is the builder for "upsert"-ing
// a bulk of TopMostConstraint nodes.
type TopMostConstraintUpsertBulk struct {
	create *TopMostConstraintCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.TopMostConstraint.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(topmostconstraint.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TopMostConstraintUpsertBulk) UpdateNewValues() *TopMostConstraintUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(topmostconstraint.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.TopMostConstraint.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TopMostConstraintUpsertBulk) Ignore() *TopMostConstraintUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TopMostConstraintUpsertBulk) DoNothing() *TopMostConstraintUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TopMostConstraintCreateBulk.OnConflict
// documentation for more info.
func (u *TopMostConstraintUpsertBulk) Update(set func(*TopMostConstraintUpsert)) *TopMostConstraintUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TopMostConstraintUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TopMostConstraintUpsertBulk) SetCreatedAt(v uint32) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *TopMostConstraintUpsertBulk) AddCreatedAt(v uint32) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateCreatedAt() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TopMostConstraintUpsertBulk) SetUpdatedAt(v uint32) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *TopMostConstraintUpsertBulk) AddUpdatedAt(v uint32) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateUpdatedAt() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *TopMostConstraintUpsertBulk) SetDeletedAt(v uint32) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *TopMostConstraintUpsertBulk) AddDeletedAt(v uint32) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateDeletedAt() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetEntID sets the "ent_id" field.
func (u *TopMostConstraintUpsertBulk) SetEntID(v uuid.UUID) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetEntID(v)
	})
}

// UpdateEntID sets the "ent_id" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateEntID() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateEntID()
	})
}

// SetTopMostID sets the "top_most_id" field.
func (u *TopMostConstraintUpsertBulk) SetTopMostID(v uuid.UUID) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetTopMostID(v)
	})
}

// UpdateTopMostID sets the "top_most_id" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateTopMostID() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateTopMostID()
	})
}

// ClearTopMostID clears the value of the "top_most_id" field.
func (u *TopMostConstraintUpsertBulk) ClearTopMostID() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearTopMostID()
	})
}

// SetConstraint sets the "constraint" field.
func (u *TopMostConstraintUpsertBulk) SetConstraint(v string) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetConstraint(v)
	})
}

// UpdateConstraint sets the "constraint" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateConstraint() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateConstraint()
	})
}

// ClearConstraint clears the value of the "constraint" field.
func (u *TopMostConstraintUpsertBulk) ClearConstraint() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearConstraint()
	})
}

// SetTargetValue sets the "target_value" field.
func (u *TopMostConstraintUpsertBulk) SetTargetValue(v decimal.Decimal) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetTargetValue(v)
	})
}

// UpdateTargetValue sets the "target_value" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateTargetValue() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateTargetValue()
	})
}

// ClearTargetValue clears the value of the "target_value" field.
func (u *TopMostConstraintUpsertBulk) ClearTargetValue() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearTargetValue()
	})
}

// SetIndex sets the "index" field.
func (u *TopMostConstraintUpsertBulk) SetIndex(v uint8) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.SetIndex(v)
	})
}

// AddIndex adds v to the "index" field.
func (u *TopMostConstraintUpsertBulk) AddIndex(v uint8) *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.AddIndex(v)
	})
}

// UpdateIndex sets the "index" field to the value that was provided on create.
func (u *TopMostConstraintUpsertBulk) UpdateIndex() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.UpdateIndex()
	})
}

// ClearIndex clears the value of the "index" field.
func (u *TopMostConstraintUpsertBulk) ClearIndex() *TopMostConstraintUpsertBulk {
	return u.Update(func(s *TopMostConstraintUpsert) {
		s.ClearIndex()
	})
}

// Exec executes the query.
func (u *TopMostConstraintUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TopMostConstraintCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TopMostConstraintCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TopMostConstraintUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
