// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// ExtraInfoUpdate is the builder for updating ExtraInfo entities.
type ExtraInfoUpdate struct {
	config
	hooks     []Hook
	mutation  *ExtraInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ExtraInfoUpdate builder.
func (eiu *ExtraInfoUpdate) Where(ps ...predicate.ExtraInfo) *ExtraInfoUpdate {
	eiu.mutation.Where(ps...)
	return eiu
}

// SetCreatedAt sets the "created_at" field.
func (eiu *ExtraInfoUpdate) SetCreatedAt(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetCreatedAt()
	eiu.mutation.SetCreatedAt(u)
	return eiu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableCreatedAt(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetCreatedAt(*u)
	}
	return eiu
}

// AddCreatedAt adds u to the "created_at" field.
func (eiu *ExtraInfoUpdate) AddCreatedAt(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddCreatedAt(u)
	return eiu
}

// SetUpdatedAt sets the "updated_at" field.
func (eiu *ExtraInfoUpdate) SetUpdatedAt(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetUpdatedAt()
	eiu.mutation.SetUpdatedAt(u)
	return eiu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (eiu *ExtraInfoUpdate) AddUpdatedAt(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddUpdatedAt(u)
	return eiu
}

// SetDeletedAt sets the "deleted_at" field.
func (eiu *ExtraInfoUpdate) SetDeletedAt(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetDeletedAt()
	eiu.mutation.SetDeletedAt(u)
	return eiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableDeletedAt(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetDeletedAt(*u)
	}
	return eiu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (eiu *ExtraInfoUpdate) AddDeletedAt(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddDeletedAt(u)
	return eiu
}

// SetGoodID sets the "good_id" field.
func (eiu *ExtraInfoUpdate) SetGoodID(u uuid.UUID) *ExtraInfoUpdate {
	eiu.mutation.SetGoodID(u)
	return eiu
}

// SetPosters sets the "posters" field.
func (eiu *ExtraInfoUpdate) SetPosters(s []string) *ExtraInfoUpdate {
	eiu.mutation.SetPosters(s)
	return eiu
}

// ClearPosters clears the value of the "posters" field.
func (eiu *ExtraInfoUpdate) ClearPosters() *ExtraInfoUpdate {
	eiu.mutation.ClearPosters()
	return eiu
}

// SetLabels sets the "labels" field.
func (eiu *ExtraInfoUpdate) SetLabels(s []string) *ExtraInfoUpdate {
	eiu.mutation.SetLabels(s)
	return eiu
}

// ClearLabels clears the value of the "labels" field.
func (eiu *ExtraInfoUpdate) ClearLabels() *ExtraInfoUpdate {
	eiu.mutation.ClearLabels()
	return eiu
}

// SetLikes sets the "likes" field.
func (eiu *ExtraInfoUpdate) SetLikes(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetLikes()
	eiu.mutation.SetLikes(u)
	return eiu
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableLikes(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetLikes(*u)
	}
	return eiu
}

// AddLikes adds u to the "likes" field.
func (eiu *ExtraInfoUpdate) AddLikes(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddLikes(u)
	return eiu
}

// ClearLikes clears the value of the "likes" field.
func (eiu *ExtraInfoUpdate) ClearLikes() *ExtraInfoUpdate {
	eiu.mutation.ClearLikes()
	return eiu
}

// SetDislikes sets the "dislikes" field.
func (eiu *ExtraInfoUpdate) SetDislikes(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetDislikes()
	eiu.mutation.SetDislikes(u)
	return eiu
}

// SetNillableDislikes sets the "dislikes" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableDislikes(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetDislikes(*u)
	}
	return eiu
}

// AddDislikes adds u to the "dislikes" field.
func (eiu *ExtraInfoUpdate) AddDislikes(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddDislikes(u)
	return eiu
}

// ClearDislikes clears the value of the "dislikes" field.
func (eiu *ExtraInfoUpdate) ClearDislikes() *ExtraInfoUpdate {
	eiu.mutation.ClearDislikes()
	return eiu
}

// SetScoreCount sets the "score_count" field.
func (eiu *ExtraInfoUpdate) SetScoreCount(u uint32) *ExtraInfoUpdate {
	eiu.mutation.ResetScoreCount()
	eiu.mutation.SetScoreCount(u)
	return eiu
}

// SetNillableScoreCount sets the "score_count" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableScoreCount(u *uint32) *ExtraInfoUpdate {
	if u != nil {
		eiu.SetScoreCount(*u)
	}
	return eiu
}

// AddScoreCount adds u to the "score_count" field.
func (eiu *ExtraInfoUpdate) AddScoreCount(u int32) *ExtraInfoUpdate {
	eiu.mutation.AddScoreCount(u)
	return eiu
}

// ClearScoreCount clears the value of the "score_count" field.
func (eiu *ExtraInfoUpdate) ClearScoreCount() *ExtraInfoUpdate {
	eiu.mutation.ClearScoreCount()
	return eiu
}

// SetScore sets the "score" field.
func (eiu *ExtraInfoUpdate) SetScore(d decimal.Decimal) *ExtraInfoUpdate {
	eiu.mutation.SetScore(d)
	return eiu
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (eiu *ExtraInfoUpdate) SetNillableScore(d *decimal.Decimal) *ExtraInfoUpdate {
	if d != nil {
		eiu.SetScore(*d)
	}
	return eiu
}

// ClearScore clears the value of the "score" field.
func (eiu *ExtraInfoUpdate) ClearScore() *ExtraInfoUpdate {
	eiu.mutation.ClearScore()
	return eiu
}

// Mutation returns the ExtraInfoMutation object of the builder.
func (eiu *ExtraInfoUpdate) Mutation() *ExtraInfoMutation {
	return eiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eiu *ExtraInfoUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := eiu.defaults(); err != nil {
		return 0, err
	}
	if len(eiu.hooks) == 0 {
		affected, err = eiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExtraInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eiu.mutation = mutation
			affected, err = eiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eiu.hooks) - 1; i >= 0; i-- {
			if eiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eiu *ExtraInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := eiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eiu *ExtraInfoUpdate) Exec(ctx context.Context) error {
	_, err := eiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiu *ExtraInfoUpdate) ExecX(ctx context.Context) {
	if err := eiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eiu *ExtraInfoUpdate) defaults() error {
	if _, ok := eiu.mutation.UpdatedAt(); !ok {
		if extrainfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized extrainfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := extrainfo.UpdateDefaultUpdatedAt()
		eiu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eiu *ExtraInfoUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraInfoUpdate {
	eiu.modifiers = append(eiu.modifiers, modifiers...)
	return eiu
}

func (eiu *ExtraInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   extrainfo.Table,
			Columns: extrainfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: extrainfo.FieldID,
			},
		},
	}
	if ps := eiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldCreatedAt,
		})
	}
	if value, ok := eiu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldCreatedAt,
		})
	}
	if value, ok := eiu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldUpdatedAt,
		})
	}
	if value, ok := eiu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldUpdatedAt,
		})
	}
	if value, ok := eiu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDeletedAt,
		})
	}
	if value, ok := eiu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDeletedAt,
		})
	}
	if value, ok := eiu.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: extrainfo.FieldGoodID,
		})
	}
	if value, ok := eiu.mutation.Posters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: extrainfo.FieldPosters,
		})
	}
	if eiu.mutation.PostersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: extrainfo.FieldPosters,
		})
	}
	if value, ok := eiu.mutation.Labels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: extrainfo.FieldLabels,
		})
	}
	if eiu.mutation.LabelsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: extrainfo.FieldLabels,
		})
	}
	if value, ok := eiu.mutation.Likes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldLikes,
		})
	}
	if value, ok := eiu.mutation.AddedLikes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldLikes,
		})
	}
	if eiu.mutation.LikesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: extrainfo.FieldLikes,
		})
	}
	if value, ok := eiu.mutation.Dislikes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDislikes,
		})
	}
	if value, ok := eiu.mutation.AddedDislikes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDislikes,
		})
	}
	if eiu.mutation.DislikesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: extrainfo.FieldDislikes,
		})
	}
	if value, ok := eiu.mutation.ScoreCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldScoreCount,
		})
	}
	if value, ok := eiu.mutation.AddedScoreCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldScoreCount,
		})
	}
	if eiu.mutation.ScoreCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: extrainfo.FieldScoreCount,
		})
	}
	if value, ok := eiu.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: extrainfo.FieldScore,
		})
	}
	if eiu.mutation.ScoreCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: extrainfo.FieldScore,
		})
	}
	_spec.Modifiers = eiu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, eiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extrainfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ExtraInfoUpdateOne is the builder for updating a single ExtraInfo entity.
type ExtraInfoUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ExtraInfoMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (eiuo *ExtraInfoUpdateOne) SetCreatedAt(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetCreatedAt()
	eiuo.mutation.SetCreatedAt(u)
	return eiuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableCreatedAt(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetCreatedAt(*u)
	}
	return eiuo
}

// AddCreatedAt adds u to the "created_at" field.
func (eiuo *ExtraInfoUpdateOne) AddCreatedAt(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddCreatedAt(u)
	return eiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (eiuo *ExtraInfoUpdateOne) SetUpdatedAt(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetUpdatedAt()
	eiuo.mutation.SetUpdatedAt(u)
	return eiuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (eiuo *ExtraInfoUpdateOne) AddUpdatedAt(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddUpdatedAt(u)
	return eiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (eiuo *ExtraInfoUpdateOne) SetDeletedAt(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetDeletedAt()
	eiuo.mutation.SetDeletedAt(u)
	return eiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableDeletedAt(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetDeletedAt(*u)
	}
	return eiuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (eiuo *ExtraInfoUpdateOne) AddDeletedAt(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddDeletedAt(u)
	return eiuo
}

// SetGoodID sets the "good_id" field.
func (eiuo *ExtraInfoUpdateOne) SetGoodID(u uuid.UUID) *ExtraInfoUpdateOne {
	eiuo.mutation.SetGoodID(u)
	return eiuo
}

// SetPosters sets the "posters" field.
func (eiuo *ExtraInfoUpdateOne) SetPosters(s []string) *ExtraInfoUpdateOne {
	eiuo.mutation.SetPosters(s)
	return eiuo
}

// ClearPosters clears the value of the "posters" field.
func (eiuo *ExtraInfoUpdateOne) ClearPosters() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearPosters()
	return eiuo
}

// SetLabels sets the "labels" field.
func (eiuo *ExtraInfoUpdateOne) SetLabels(s []string) *ExtraInfoUpdateOne {
	eiuo.mutation.SetLabels(s)
	return eiuo
}

// ClearLabels clears the value of the "labels" field.
func (eiuo *ExtraInfoUpdateOne) ClearLabels() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearLabels()
	return eiuo
}

// SetLikes sets the "likes" field.
func (eiuo *ExtraInfoUpdateOne) SetLikes(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetLikes()
	eiuo.mutation.SetLikes(u)
	return eiuo
}

// SetNillableLikes sets the "likes" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableLikes(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetLikes(*u)
	}
	return eiuo
}

// AddLikes adds u to the "likes" field.
func (eiuo *ExtraInfoUpdateOne) AddLikes(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddLikes(u)
	return eiuo
}

// ClearLikes clears the value of the "likes" field.
func (eiuo *ExtraInfoUpdateOne) ClearLikes() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearLikes()
	return eiuo
}

// SetDislikes sets the "dislikes" field.
func (eiuo *ExtraInfoUpdateOne) SetDislikes(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetDislikes()
	eiuo.mutation.SetDislikes(u)
	return eiuo
}

// SetNillableDislikes sets the "dislikes" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableDislikes(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetDislikes(*u)
	}
	return eiuo
}

// AddDislikes adds u to the "dislikes" field.
func (eiuo *ExtraInfoUpdateOne) AddDislikes(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddDislikes(u)
	return eiuo
}

// ClearDislikes clears the value of the "dislikes" field.
func (eiuo *ExtraInfoUpdateOne) ClearDislikes() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearDislikes()
	return eiuo
}

// SetScoreCount sets the "score_count" field.
func (eiuo *ExtraInfoUpdateOne) SetScoreCount(u uint32) *ExtraInfoUpdateOne {
	eiuo.mutation.ResetScoreCount()
	eiuo.mutation.SetScoreCount(u)
	return eiuo
}

// SetNillableScoreCount sets the "score_count" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableScoreCount(u *uint32) *ExtraInfoUpdateOne {
	if u != nil {
		eiuo.SetScoreCount(*u)
	}
	return eiuo
}

// AddScoreCount adds u to the "score_count" field.
func (eiuo *ExtraInfoUpdateOne) AddScoreCount(u int32) *ExtraInfoUpdateOne {
	eiuo.mutation.AddScoreCount(u)
	return eiuo
}

// ClearScoreCount clears the value of the "score_count" field.
func (eiuo *ExtraInfoUpdateOne) ClearScoreCount() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearScoreCount()
	return eiuo
}

// SetScore sets the "score" field.
func (eiuo *ExtraInfoUpdateOne) SetScore(d decimal.Decimal) *ExtraInfoUpdateOne {
	eiuo.mutation.SetScore(d)
	return eiuo
}

// SetNillableScore sets the "score" field if the given value is not nil.
func (eiuo *ExtraInfoUpdateOne) SetNillableScore(d *decimal.Decimal) *ExtraInfoUpdateOne {
	if d != nil {
		eiuo.SetScore(*d)
	}
	return eiuo
}

// ClearScore clears the value of the "score" field.
func (eiuo *ExtraInfoUpdateOne) ClearScore() *ExtraInfoUpdateOne {
	eiuo.mutation.ClearScore()
	return eiuo
}

// Mutation returns the ExtraInfoMutation object of the builder.
func (eiuo *ExtraInfoUpdateOne) Mutation() *ExtraInfoMutation {
	return eiuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eiuo *ExtraInfoUpdateOne) Select(field string, fields ...string) *ExtraInfoUpdateOne {
	eiuo.fields = append([]string{field}, fields...)
	return eiuo
}

// Save executes the query and returns the updated ExtraInfo entity.
func (eiuo *ExtraInfoUpdateOne) Save(ctx context.Context) (*ExtraInfo, error) {
	var (
		err  error
		node *ExtraInfo
	)
	if err := eiuo.defaults(); err != nil {
		return nil, err
	}
	if len(eiuo.hooks) == 0 {
		node, err = eiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ExtraInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eiuo.mutation = mutation
			node, err = eiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eiuo.hooks) - 1; i >= 0; i-- {
			if eiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eiuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, eiuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (eiuo *ExtraInfoUpdateOne) SaveX(ctx context.Context) *ExtraInfo {
	node, err := eiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eiuo *ExtraInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := eiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiuo *ExtraInfoUpdateOne) ExecX(ctx context.Context) {
	if err := eiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eiuo *ExtraInfoUpdateOne) defaults() error {
	if _, ok := eiuo.mutation.UpdatedAt(); !ok {
		if extrainfo.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized extrainfo.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := extrainfo.UpdateDefaultUpdatedAt()
		eiuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (eiuo *ExtraInfoUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ExtraInfoUpdateOne {
	eiuo.modifiers = append(eiuo.modifiers, modifiers...)
	return eiuo
}

func (eiuo *ExtraInfoUpdateOne) sqlSave(ctx context.Context) (_node *ExtraInfo, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   extrainfo.Table,
			Columns: extrainfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: extrainfo.FieldID,
			},
		},
	}
	id, ok := eiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ExtraInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, extrainfo.FieldID)
		for _, f := range fields {
			if !extrainfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != extrainfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldCreatedAt,
		})
	}
	if value, ok := eiuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldCreatedAt,
		})
	}
	if value, ok := eiuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldUpdatedAt,
		})
	}
	if value, ok := eiuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldUpdatedAt,
		})
	}
	if value, ok := eiuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDeletedAt,
		})
	}
	if value, ok := eiuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDeletedAt,
		})
	}
	if value, ok := eiuo.mutation.GoodID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: extrainfo.FieldGoodID,
		})
	}
	if value, ok := eiuo.mutation.Posters(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: extrainfo.FieldPosters,
		})
	}
	if eiuo.mutation.PostersCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: extrainfo.FieldPosters,
		})
	}
	if value, ok := eiuo.mutation.Labels(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: extrainfo.FieldLabels,
		})
	}
	if eiuo.mutation.LabelsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: extrainfo.FieldLabels,
		})
	}
	if value, ok := eiuo.mutation.Likes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldLikes,
		})
	}
	if value, ok := eiuo.mutation.AddedLikes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldLikes,
		})
	}
	if eiuo.mutation.LikesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: extrainfo.FieldLikes,
		})
	}
	if value, ok := eiuo.mutation.Dislikes(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDislikes,
		})
	}
	if value, ok := eiuo.mutation.AddedDislikes(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldDislikes,
		})
	}
	if eiuo.mutation.DislikesCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: extrainfo.FieldDislikes,
		})
	}
	if value, ok := eiuo.mutation.ScoreCount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldScoreCount,
		})
	}
	if value, ok := eiuo.mutation.AddedScoreCount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: extrainfo.FieldScoreCount,
		})
	}
	if eiuo.mutation.ScoreCountCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Column: extrainfo.FieldScoreCount,
		})
	}
	if value, ok := eiuo.mutation.Score(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Value:  value,
			Column: extrainfo.FieldScore,
		})
	}
	if eiuo.mutation.ScoreCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeOther,
			Column: extrainfo.FieldScore,
		})
	}
	_spec.Modifiers = eiuo.modifiers
	_node = &ExtraInfo{config: eiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{extrainfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
