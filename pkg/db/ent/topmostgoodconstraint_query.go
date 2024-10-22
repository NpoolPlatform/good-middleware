// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodconstraint"
)

// TopMostGoodConstraintQuery is the builder for querying TopMostGoodConstraint entities.
type TopMostGoodConstraintQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TopMostGoodConstraint
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TopMostGoodConstraintQuery builder.
func (tmgcq *TopMostGoodConstraintQuery) Where(ps ...predicate.TopMostGoodConstraint) *TopMostGoodConstraintQuery {
	tmgcq.predicates = append(tmgcq.predicates, ps...)
	return tmgcq
}

// Limit adds a limit step to the query.
func (tmgcq *TopMostGoodConstraintQuery) Limit(limit int) *TopMostGoodConstraintQuery {
	tmgcq.limit = &limit
	return tmgcq
}

// Offset adds an offset step to the query.
func (tmgcq *TopMostGoodConstraintQuery) Offset(offset int) *TopMostGoodConstraintQuery {
	tmgcq.offset = &offset
	return tmgcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tmgcq *TopMostGoodConstraintQuery) Unique(unique bool) *TopMostGoodConstraintQuery {
	tmgcq.unique = &unique
	return tmgcq
}

// Order adds an order step to the query.
func (tmgcq *TopMostGoodConstraintQuery) Order(o ...OrderFunc) *TopMostGoodConstraintQuery {
	tmgcq.order = append(tmgcq.order, o...)
	return tmgcq
}

// First returns the first TopMostGoodConstraint entity from the query.
// Returns a *NotFoundError when no TopMostGoodConstraint was found.
func (tmgcq *TopMostGoodConstraintQuery) First(ctx context.Context) (*TopMostGoodConstraint, error) {
	nodes, err := tmgcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{topmostgoodconstraint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) FirstX(ctx context.Context) *TopMostGoodConstraint {
	node, err := tmgcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TopMostGoodConstraint ID from the query.
// Returns a *NotFoundError when no TopMostGoodConstraint ID was found.
func (tmgcq *TopMostGoodConstraintQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmgcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{topmostgoodconstraint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := tmgcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TopMostGoodConstraint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TopMostGoodConstraint entity is found.
// Returns a *NotFoundError when no TopMostGoodConstraint entities are found.
func (tmgcq *TopMostGoodConstraintQuery) Only(ctx context.Context) (*TopMostGoodConstraint, error) {
	nodes, err := tmgcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{topmostgoodconstraint.Label}
	default:
		return nil, &NotSingularError{topmostgoodconstraint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) OnlyX(ctx context.Context) *TopMostGoodConstraint {
	node, err := tmgcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TopMostGoodConstraint ID in the query.
// Returns a *NotSingularError when more than one TopMostGoodConstraint ID is found.
// Returns a *NotFoundError when no entities are found.
func (tmgcq *TopMostGoodConstraintQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmgcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{topmostgoodconstraint.Label}
	default:
		err = &NotSingularError{topmostgoodconstraint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := tmgcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TopMostGoodConstraints.
func (tmgcq *TopMostGoodConstraintQuery) All(ctx context.Context) ([]*TopMostGoodConstraint, error) {
	if err := tmgcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tmgcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) AllX(ctx context.Context) []*TopMostGoodConstraint {
	nodes, err := tmgcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TopMostGoodConstraint IDs.
func (tmgcq *TopMostGoodConstraintQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := tmgcq.Select(topmostgoodconstraint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := tmgcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tmgcq *TopMostGoodConstraintQuery) Count(ctx context.Context) (int, error) {
	if err := tmgcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tmgcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) CountX(ctx context.Context) int {
	count, err := tmgcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tmgcq *TopMostGoodConstraintQuery) Exist(ctx context.Context) (bool, error) {
	if err := tmgcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tmgcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tmgcq *TopMostGoodConstraintQuery) ExistX(ctx context.Context) bool {
	exist, err := tmgcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TopMostGoodConstraintQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tmgcq *TopMostGoodConstraintQuery) Clone() *TopMostGoodConstraintQuery {
	if tmgcq == nil {
		return nil
	}
	return &TopMostGoodConstraintQuery{
		config:     tmgcq.config,
		limit:      tmgcq.limit,
		offset:     tmgcq.offset,
		order:      append([]OrderFunc{}, tmgcq.order...),
		predicates: append([]predicate.TopMostGoodConstraint{}, tmgcq.predicates...),
		// clone intermediate query.
		sql:    tmgcq.sql.Clone(),
		path:   tmgcq.path,
		unique: tmgcq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TopMostGoodConstraint.Query().
//		GroupBy(topmostgoodconstraint.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tmgcq *TopMostGoodConstraintQuery) GroupBy(field string, fields ...string) *TopMostGoodConstraintGroupBy {
	grbuild := &TopMostGoodConstraintGroupBy{config: tmgcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tmgcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tmgcq.sqlQuery(ctx), nil
	}
	grbuild.label = topmostgoodconstraint.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt uint32 `json:"created_at,omitempty"`
//	}
//
//	client.TopMostGoodConstraint.Query().
//		Select(topmostgoodconstraint.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (tmgcq *TopMostGoodConstraintQuery) Select(fields ...string) *TopMostGoodConstraintSelect {
	tmgcq.fields = append(tmgcq.fields, fields...)
	selbuild := &TopMostGoodConstraintSelect{TopMostGoodConstraintQuery: tmgcq}
	selbuild.label = topmostgoodconstraint.Label
	selbuild.flds, selbuild.scan = &tmgcq.fields, selbuild.Scan
	return selbuild
}

func (tmgcq *TopMostGoodConstraintQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tmgcq.fields {
		if !topmostgoodconstraint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tmgcq.path != nil {
		prev, err := tmgcq.path(ctx)
		if err != nil {
			return err
		}
		tmgcq.sql = prev
	}
	if topmostgoodconstraint.Policy == nil {
		return errors.New("ent: uninitialized topmostgoodconstraint.Policy (forgotten import ent/runtime?)")
	}
	if err := topmostgoodconstraint.Policy.EvalQuery(ctx, tmgcq); err != nil {
		return err
	}
	return nil
}

func (tmgcq *TopMostGoodConstraintQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TopMostGoodConstraint, error) {
	var (
		nodes = []*TopMostGoodConstraint{}
		_spec = tmgcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TopMostGoodConstraint).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TopMostGoodConstraint{config: tmgcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tmgcq.modifiers) > 0 {
		_spec.Modifiers = tmgcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tmgcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tmgcq *TopMostGoodConstraintQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tmgcq.querySpec()
	if len(tmgcq.modifiers) > 0 {
		_spec.Modifiers = tmgcq.modifiers
	}
	_spec.Node.Columns = tmgcq.fields
	if len(tmgcq.fields) > 0 {
		_spec.Unique = tmgcq.unique != nil && *tmgcq.unique
	}
	return sqlgraph.CountNodes(ctx, tmgcq.driver, _spec)
}

func (tmgcq *TopMostGoodConstraintQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tmgcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tmgcq *TopMostGoodConstraintQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topmostgoodconstraint.Table,
			Columns: topmostgoodconstraint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostgoodconstraint.FieldID,
			},
		},
		From:   tmgcq.sql,
		Unique: true,
	}
	if unique := tmgcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tmgcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topmostgoodconstraint.FieldID)
		for i := range fields {
			if fields[i] != topmostgoodconstraint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tmgcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tmgcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tmgcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tmgcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tmgcq *TopMostGoodConstraintQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tmgcq.driver.Dialect())
	t1 := builder.Table(topmostgoodconstraint.Table)
	columns := tmgcq.fields
	if len(columns) == 0 {
		columns = topmostgoodconstraint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tmgcq.sql != nil {
		selector = tmgcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tmgcq.unique != nil && *tmgcq.unique {
		selector.Distinct()
	}
	for _, m := range tmgcq.modifiers {
		m(selector)
	}
	for _, p := range tmgcq.predicates {
		p(selector)
	}
	for _, p := range tmgcq.order {
		p(selector)
	}
	if offset := tmgcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tmgcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tmgcq *TopMostGoodConstraintQuery) ForUpdate(opts ...sql.LockOption) *TopMostGoodConstraintQuery {
	if tmgcq.driver.Dialect() == dialect.Postgres {
		tmgcq.Unique(false)
	}
	tmgcq.modifiers = append(tmgcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tmgcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tmgcq *TopMostGoodConstraintQuery) ForShare(opts ...sql.LockOption) *TopMostGoodConstraintQuery {
	if tmgcq.driver.Dialect() == dialect.Postgres {
		tmgcq.Unique(false)
	}
	tmgcq.modifiers = append(tmgcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tmgcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmgcq *TopMostGoodConstraintQuery) Modify(modifiers ...func(s *sql.Selector)) *TopMostGoodConstraintSelect {
	tmgcq.modifiers = append(tmgcq.modifiers, modifiers...)
	return tmgcq.Select()
}

// TopMostGoodConstraintGroupBy is the group-by builder for TopMostGoodConstraint entities.
type TopMostGoodConstraintGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tmgcgb *TopMostGoodConstraintGroupBy) Aggregate(fns ...AggregateFunc) *TopMostGoodConstraintGroupBy {
	tmgcgb.fns = append(tmgcgb.fns, fns...)
	return tmgcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tmgcgb *TopMostGoodConstraintGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tmgcgb.path(ctx)
	if err != nil {
		return err
	}
	tmgcgb.sql = query
	return tmgcgb.sqlScan(ctx, v)
}

func (tmgcgb *TopMostGoodConstraintGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tmgcgb.fields {
		if !topmostgoodconstraint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tmgcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmgcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tmgcgb *TopMostGoodConstraintGroupBy) sqlQuery() *sql.Selector {
	selector := tmgcgb.sql.Select()
	aggregation := make([]string, 0, len(tmgcgb.fns))
	for _, fn := range tmgcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tmgcgb.fields)+len(tmgcgb.fns))
		for _, f := range tmgcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tmgcgb.fields...)...)
}

// TopMostGoodConstraintSelect is the builder for selecting fields of TopMostGoodConstraint entities.
type TopMostGoodConstraintSelect struct {
	*TopMostGoodConstraintQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tmgcs *TopMostGoodConstraintSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tmgcs.prepareQuery(ctx); err != nil {
		return err
	}
	tmgcs.sql = tmgcs.TopMostGoodConstraintQuery.sqlQuery(ctx)
	return tmgcs.sqlScan(ctx, v)
}

func (tmgcs *TopMostGoodConstraintSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tmgcs.sql.Query()
	if err := tmgcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmgcs *TopMostGoodConstraintSelect) Modify(modifiers ...func(s *sql.Selector)) *TopMostGoodConstraintSelect {
	tmgcs.modifiers = append(tmgcs.modifiers, modifiers...)
	return tmgcs
}
