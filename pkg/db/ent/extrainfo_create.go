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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	"github.com/google/uuid"
)

// ExtraInfoCreate is the builder for creating a ExtraInfo entity.
type ExtraInfoCreate struct {
	config
	mutation *ExtraInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (eic *ExtraInfoCreate) SetCreatedAt(u uint32) *ExtraInfoCreate {
	eic.mutation.SetCreatedAt(u)
	return eic
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eic *ExtraInfoCreate) SetNillableCreatedAt(u *uint32) *ExtraInfoCreate {
	if u != nil {
		eic.SetCreatedAt(*u)
	}
	return eic
}

// SetUpdatedAt sets the "updated_at" field.
func (eic *ExtraInfoCreate) SetUpdatedAt(u uint32) *ExtraInfoCreate {
	eic.mutation.SetUpdatedAt(u)
	return eic
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (eic *ExtraInfoCreate) SetNillableUpdatedAt(u *uint32) *ExtraInfoCreate {
	if u != nil {
		eic.SetUpdatedAt(*u)
	}
	return eic
}

// SetDeletedAt sets the "deleted_at" field.
func (eic *ExtraInfoCreate) SetDeletedAt(u uint32) *ExtraInfoCreate {
	eic.mutation.SetDeletedAt(u)
	return eic
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eic *ExtraInfoCreate) SetNillableDeletedAt(u *uint32) *ExtraInfoCreate {
	if u != nil {
		eic.SetDeletedAt(*u)
	}
	return eic
}

// SetGoodID sets the "good_id" field.
func (eic *ExtraInfoCreate) SetGoodID(u uuid.UUID) *ExtraInfoCreate {
	eic.mutation.SetGoodID(u)
	return eic
}

// SetPosters sets the "posters" field.
func (eic *ExtraInfoCreate) SetPosters(s []string) *ExtraInfoCreate {
	eic.mutation.SetPosters(s)
	return eic
}

// SetLabels sets the "labels" field.
func (eic *ExtraInfoCreate) SetLabels(s []string) *ExtraInfoCreate {
	eic.mutation.SetLabels(s)
	return eic
}

// SetVoteCount sets the "vote_count" field.
func (eic *ExtraInfoCreate) SetVoteCount(u uint32) *ExtraInfoCreate {
	eic.mutation.SetVoteCount(u)
	return eic
}

// SetNillableVoteCount sets the "vote_count" field if the given value is not nil.
func (eic *ExtraInfoCreate) SetNillableVoteCount(u *uint32) *ExtraInfoCreate {
	if u != nil {
		eic.SetVoteCount(*u)
	}
	return eic
}

// SetRating sets the "rating" field.
func (eic *ExtraInfoCreate) SetRating(f float32) *ExtraInfoCreate {
	eic.mutation.SetRating(f)
	return eic
}

// SetNillableRating sets the "rating" field if the given value is not nil.
func (eic *ExtraInfoCreate) SetNillableRating(f *float32) *ExtraInfoCreate {
	if f != nil {
		eic.SetRating(*f)
	}
	return eic
}

// SetID sets the "id" field.
func (eic *ExtraInfoCreate) SetID(u uuid.UUID) *ExtraInfoCreate {
	eic.mutation.SetID(u)
	return eic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (eic *ExtraInfoCreate) SetNillableID(u *uuid.UUID) *ExtraInfoCreate {
	if u != nil {
		eic.SetID(*u)
	}
	return eic
}

// Mutation returns the ExtraInfoMutation object of the builder.
func (eic *ExtraInfoCreate) Mutation() *ExtraInfoMutation {
	return eic.mutation
}

// Save creates the ExtraInfo in the database.
func (eic *ExtraInfoCreate) Save(ctx context.Context) (*ExtraInfo, error) {
	var (
		err  error
		node *ExtraInfo
	)
	if err := eic.defaults(); err != nil {
		return nil, err
	}
	if len(eic.hooks) == 0 {
		if err = eic.check(); err != nil {
			return nil, err
		}
		node, err = eic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExtraInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eic.check(); err != nil {
				return nil, err
			}
			eic.mutation = mutation
			if node, err = eic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(eic.hooks) - 1; i >= 0; i-- {
			if eic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eic.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eic.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ExtraInfo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ExtraInfoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (eic *ExtraInfoCreate) SaveX(ctx context.Context) *ExtraInfo {
	v, err := eic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eic *ExtraInfoCreate) Exec(ctx context.Context) error {
	_, err := eic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eic *ExtraInfoCreate) ExecX(ctx context.Context) {
	if err := eic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eic *ExtraInfoCreate) defaults() error {
	if _, ok := eic.mutation.CreatedAt(); !ok {
		if extrainfo.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized extrainfo.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := extrainfo.DefaultCreatedAt()
		eic.mutation.SetCreatedAt(v)
	}
	if _, ok := eic.mutation.UpdatedAt(); !ok {
		if extrainfo.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized extrainfo.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := extrainfo.DefaultUpdatedAt()
		eic.mutation.SetUpdatedAt(v)
	}
	if _, ok := eic.mutation.DeletedAt(); !ok {
		if extrainfo.DefaultDeletedAt == nil {
			return fmt.Errorf("ent: uninitialized extrainfo.DefaultDeletedAt (forgotten import ent/runtime?)")
		}
		v := extrainfo.DefaultDeletedAt()
		eic.mutation.SetDeletedAt(v)
	}
	if _, ok := eic.mutation.Posters(); !ok {
		v := extrainfo.DefaultPosters
		eic.mutation.SetPosters(v)
	}
	if _, ok := eic.mutation.Labels(); !ok {
		v := extrainfo.DefaultLabels
		eic.mutation.SetLabels(v)
	}
	if _, ok := eic.mutation.VoteCount(); !ok {
		v := extrainfo.DefaultVoteCount
		eic.mutation.SetVoteCount(v)
	}
	if _, ok := eic.mutation.Rating(); !ok {
		v := extrainfo.DefaultRating
		eic.mutation.SetRating(v)
	}
	if _, ok := eic.mutation.ID(); !ok {
		if extrainfo.DefaultID == nil {
			return fmt.Errorf("ent: uninitialized extrainfo.DefaultID (forgotten import ent/runtime?)")
		}
		v := extrainfo.DefaultID()
		eic.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (eic *ExtraInfoCreate) check() error {
	if _, ok := eic.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ExtraInfo.created_at"`)}
	}
	if _, ok := eic.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ExtraInfo.updated_at"`)}
	}
	if _, ok := eic.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "ExtraInfo.deleted_at"`)}
	}
	if _, ok := eic.mutation.GoodID(); !ok {
		return &ValidationError{Name: "good_id", err: errors.New(`ent: missing required field "ExtraInfo.good_id"`)}
	}
	return nil
}

func (eic *ExtraInfoCreate) sqlSave(ctx context.Context) (*ExtraInfo, error) {
	_node, _spec := eic.createSpec()
	if err := sqlgraph.CreateNode(ctx, eic.driver, _spec); err != nil {
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

func (eic *ExtraInfoCreate) createSpec() (*ExtraInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &ExtraInfo{config: eic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: extrainfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: extrainfo.FieldID,
			},
		}
	)
	_spec.OnConflict = eic.conflict
	if id, ok := eic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := eic.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := eic.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := eic.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := eic.mutation.GoodID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: extrainfo.FieldGoodID,
		})
		_node.GoodID = value
	}
	if value, ok := eic.mutation.Posters(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: extrainfo.FieldPosters,
		})
		_node.Posters = value
	}
	if value, ok := eic.mutation.Labels(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: extrainfo.FieldLabels,
		})
		_node.Labels = value
	}
	if value, ok := eic.mutation.VoteCount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldVoteCount,
		})
		_node.VoteCount = value
	}
	if value, ok := eic.mutation.Rating(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: extrainfo.FieldRating,
		})
		_node.Rating = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ExtraInfo.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExtraInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (eic *ExtraInfoCreate) OnConflict(opts ...sql.ConflictOption) *ExtraInfoUpsertOne {
	eic.conflict = opts
	return &ExtraInfoUpsertOne{
		create: eic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ExtraInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (eic *ExtraInfoCreate) OnConflictColumns(columns ...string) *ExtraInfoUpsertOne {
	eic.conflict = append(eic.conflict, sql.ConflictColumns(columns...))
	return &ExtraInfoUpsertOne{
		create: eic,
	}
}

type (
	// ExtraInfoUpsertOne is the builder for "upsert"-ing
	//  one ExtraInfo node.
	ExtraInfoUpsertOne struct {
		create *ExtraInfoCreate
	}

	// ExtraInfoUpsert is the "OnConflict" setter.
	ExtraInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *ExtraInfoUpsert) SetCreatedAt(v uint32) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdateCreatedAt() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldCreatedAt)
	return u
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ExtraInfoUpsert) AddCreatedAt(v uint32) *ExtraInfoUpsert {
	u.Add(extrainfo.FieldCreatedAt, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ExtraInfoUpsert) SetUpdatedAt(v uint32) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdateUpdatedAt() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ExtraInfoUpsert) AddUpdatedAt(v uint32) *ExtraInfoUpsert {
	u.Add(extrainfo.FieldUpdatedAt, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ExtraInfoUpsert) SetDeletedAt(v uint32) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdateDeletedAt() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ExtraInfoUpsert) AddDeletedAt(v uint32) *ExtraInfoUpsert {
	u.Add(extrainfo.FieldDeletedAt, v)
	return u
}

// SetGoodID sets the "good_id" field.
func (u *ExtraInfoUpsert) SetGoodID(v uuid.UUID) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldGoodID, v)
	return u
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdateGoodID() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldGoodID)
	return u
}

// SetPosters sets the "posters" field.
func (u *ExtraInfoUpsert) SetPosters(v []string) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldPosters, v)
	return u
}

// UpdatePosters sets the "posters" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdatePosters() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldPosters)
	return u
}

// ClearPosters clears the value of the "posters" field.
func (u *ExtraInfoUpsert) ClearPosters() *ExtraInfoUpsert {
	u.SetNull(extrainfo.FieldPosters)
	return u
}

// SetLabels sets the "labels" field.
func (u *ExtraInfoUpsert) SetLabels(v []string) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldLabels, v)
	return u
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdateLabels() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldLabels)
	return u
}

// ClearLabels clears the value of the "labels" field.
func (u *ExtraInfoUpsert) ClearLabels() *ExtraInfoUpsert {
	u.SetNull(extrainfo.FieldLabels)
	return u
}

// SetVoteCount sets the "vote_count" field.
func (u *ExtraInfoUpsert) SetVoteCount(v uint32) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldVoteCount, v)
	return u
}

// UpdateVoteCount sets the "vote_count" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdateVoteCount() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldVoteCount)
	return u
}

// AddVoteCount adds v to the "vote_count" field.
func (u *ExtraInfoUpsert) AddVoteCount(v uint32) *ExtraInfoUpsert {
	u.Add(extrainfo.FieldVoteCount, v)
	return u
}

// ClearVoteCount clears the value of the "vote_count" field.
func (u *ExtraInfoUpsert) ClearVoteCount() *ExtraInfoUpsert {
	u.SetNull(extrainfo.FieldVoteCount)
	return u
}

// SetRating sets the "rating" field.
func (u *ExtraInfoUpsert) SetRating(v float32) *ExtraInfoUpsert {
	u.Set(extrainfo.FieldRating, v)
	return u
}

// UpdateRating sets the "rating" field to the value that was provided on create.
func (u *ExtraInfoUpsert) UpdateRating() *ExtraInfoUpsert {
	u.SetExcluded(extrainfo.FieldRating)
	return u
}

// AddRating adds v to the "rating" field.
func (u *ExtraInfoUpsert) AddRating(v float32) *ExtraInfoUpsert {
	u.Add(extrainfo.FieldRating, v)
	return u
}

// ClearRating clears the value of the "rating" field.
func (u *ExtraInfoUpsert) ClearRating() *ExtraInfoUpsert {
	u.SetNull(extrainfo.FieldRating)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ExtraInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(extrainfo.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ExtraInfoUpsertOne) UpdateNewValues() *ExtraInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(extrainfo.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.ExtraInfo.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ExtraInfoUpsertOne) Ignore() *ExtraInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExtraInfoUpsertOne) DoNothing() *ExtraInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExtraInfoCreate.OnConflict
// documentation for more info.
func (u *ExtraInfoUpsertOne) Update(set func(*ExtraInfoUpsert)) *ExtraInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExtraInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ExtraInfoUpsertOne) SetCreatedAt(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ExtraInfoUpsertOne) AddCreatedAt(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdateCreatedAt() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ExtraInfoUpsertOne) SetUpdatedAt(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ExtraInfoUpsertOne) AddUpdatedAt(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdateUpdatedAt() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ExtraInfoUpsertOne) SetDeletedAt(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ExtraInfoUpsertOne) AddDeletedAt(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdateDeletedAt() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *ExtraInfoUpsertOne) SetGoodID(v uuid.UUID) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdateGoodID() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateGoodID()
	})
}

// SetPosters sets the "posters" field.
func (u *ExtraInfoUpsertOne) SetPosters(v []string) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetPosters(v)
	})
}

// UpdatePosters sets the "posters" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdatePosters() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdatePosters()
	})
}

// ClearPosters clears the value of the "posters" field.
func (u *ExtraInfoUpsertOne) ClearPosters() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearPosters()
	})
}

// SetLabels sets the "labels" field.
func (u *ExtraInfoUpsertOne) SetLabels(v []string) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdateLabels() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *ExtraInfoUpsertOne) ClearLabels() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearLabels()
	})
}

// SetVoteCount sets the "vote_count" field.
func (u *ExtraInfoUpsertOne) SetVoteCount(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetVoteCount(v)
	})
}

// AddVoteCount adds v to the "vote_count" field.
func (u *ExtraInfoUpsertOne) AddVoteCount(v uint32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddVoteCount(v)
	})
}

// UpdateVoteCount sets the "vote_count" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdateVoteCount() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateVoteCount()
	})
}

// ClearVoteCount clears the value of the "vote_count" field.
func (u *ExtraInfoUpsertOne) ClearVoteCount() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearVoteCount()
	})
}

// SetRating sets the "rating" field.
func (u *ExtraInfoUpsertOne) SetRating(v float32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetRating(v)
	})
}

// AddRating adds v to the "rating" field.
func (u *ExtraInfoUpsertOne) AddRating(v float32) *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddRating(v)
	})
}

// UpdateRating sets the "rating" field to the value that was provided on create.
func (u *ExtraInfoUpsertOne) UpdateRating() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateRating()
	})
}

// ClearRating clears the value of the "rating" field.
func (u *ExtraInfoUpsertOne) ClearRating() *ExtraInfoUpsertOne {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearRating()
	})
}

// Exec executes the query.
func (u *ExtraInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExtraInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExtraInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ExtraInfoUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ExtraInfoUpsertOne.ID is not supported by MySQL driver. Use ExtraInfoUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ExtraInfoUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ExtraInfoCreateBulk is the builder for creating many ExtraInfo entities in bulk.
type ExtraInfoCreateBulk struct {
	config
	builders []*ExtraInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the ExtraInfo entities in the database.
func (eicb *ExtraInfoCreateBulk) Save(ctx context.Context) ([]*ExtraInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eicb.builders))
	nodes := make([]*ExtraInfo, len(eicb.builders))
	mutators := make([]Mutator, len(eicb.builders))
	for i := range eicb.builders {
		func(i int, root context.Context) {
			builder := eicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ExtraInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, eicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = eicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, eicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eicb *ExtraInfoCreateBulk) SaveX(ctx context.Context) []*ExtraInfo {
	v, err := eicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eicb *ExtraInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := eicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eicb *ExtraInfoCreateBulk) ExecX(ctx context.Context) {
	if err := eicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ExtraInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ExtraInfoUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
//
func (eicb *ExtraInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *ExtraInfoUpsertBulk {
	eicb.conflict = opts
	return &ExtraInfoUpsertBulk{
		create: eicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ExtraInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (eicb *ExtraInfoCreateBulk) OnConflictColumns(columns ...string) *ExtraInfoUpsertBulk {
	eicb.conflict = append(eicb.conflict, sql.ConflictColumns(columns...))
	return &ExtraInfoUpsertBulk{
		create: eicb,
	}
}

// ExtraInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of ExtraInfo nodes.
type ExtraInfoUpsertBulk struct {
	create *ExtraInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ExtraInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(extrainfo.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ExtraInfoUpsertBulk) UpdateNewValues() *ExtraInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(extrainfo.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ExtraInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ExtraInfoUpsertBulk) Ignore() *ExtraInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ExtraInfoUpsertBulk) DoNothing() *ExtraInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ExtraInfoCreateBulk.OnConflict
// documentation for more info.
func (u *ExtraInfoUpsertBulk) Update(set func(*ExtraInfoUpsert)) *ExtraInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ExtraInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ExtraInfoUpsertBulk) SetCreatedAt(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetCreatedAt(v)
	})
}

// AddCreatedAt adds v to the "created_at" field.
func (u *ExtraInfoUpsertBulk) AddCreatedAt(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdateCreatedAt() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ExtraInfoUpsertBulk) SetUpdatedAt(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *ExtraInfoUpsertBulk) AddUpdatedAt(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdateUpdatedAt() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *ExtraInfoUpsertBulk) SetDeletedAt(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *ExtraInfoUpsertBulk) AddDeletedAt(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdateDeletedAt() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetGoodID sets the "good_id" field.
func (u *ExtraInfoUpsertBulk) SetGoodID(v uuid.UUID) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetGoodID(v)
	})
}

// UpdateGoodID sets the "good_id" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdateGoodID() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateGoodID()
	})
}

// SetPosters sets the "posters" field.
func (u *ExtraInfoUpsertBulk) SetPosters(v []string) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetPosters(v)
	})
}

// UpdatePosters sets the "posters" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdatePosters() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdatePosters()
	})
}

// ClearPosters clears the value of the "posters" field.
func (u *ExtraInfoUpsertBulk) ClearPosters() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearPosters()
	})
}

// SetLabels sets the "labels" field.
func (u *ExtraInfoUpsertBulk) SetLabels(v []string) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdateLabels() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *ExtraInfoUpsertBulk) ClearLabels() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearLabels()
	})
}

// SetVoteCount sets the "vote_count" field.
func (u *ExtraInfoUpsertBulk) SetVoteCount(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetVoteCount(v)
	})
}

// AddVoteCount adds v to the "vote_count" field.
func (u *ExtraInfoUpsertBulk) AddVoteCount(v uint32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddVoteCount(v)
	})
}

// UpdateVoteCount sets the "vote_count" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdateVoteCount() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateVoteCount()
	})
}

// ClearVoteCount clears the value of the "vote_count" field.
func (u *ExtraInfoUpsertBulk) ClearVoteCount() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearVoteCount()
	})
}

// SetRating sets the "rating" field.
func (u *ExtraInfoUpsertBulk) SetRating(v float32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.SetRating(v)
	})
}

// AddRating adds v to the "rating" field.
func (u *ExtraInfoUpsertBulk) AddRating(v float32) *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.AddRating(v)
	})
}

// UpdateRating sets the "rating" field to the value that was provided on create.
func (u *ExtraInfoUpsertBulk) UpdateRating() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.UpdateRating()
	})
}

// ClearRating clears the value of the "rating" field.
func (u *ExtraInfoUpsertBulk) ClearRating() *ExtraInfoUpsertBulk {
	return u.Update(func(s *ExtraInfoUpsert) {
		s.ClearRating()
	})
}

// Exec executes the query.
func (u *ExtraInfoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ExtraInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ExtraInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ExtraInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
