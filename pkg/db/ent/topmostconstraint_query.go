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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostconstraint"
)

// TopMostConstraintQuery is the builder for querying TopMostConstraint entities.
type TopMostConstraintQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.TopMostConstraint
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TopMostConstraintQuery builder.
func (tmcq *TopMostConstraintQuery) Where(ps ...predicate.TopMostConstraint) *TopMostConstraintQuery {
	tmcq.predicates = append(tmcq.predicates, ps...)
	return tmcq
}

// Limit adds a limit step to the query.
func (tmcq *TopMostConstraintQuery) Limit(limit int) *TopMostConstraintQuery {
	tmcq.limit = &limit
	return tmcq
}

// Offset adds an offset step to the query.
func (tmcq *TopMostConstraintQuery) Offset(offset int) *TopMostConstraintQuery {
	tmcq.offset = &offset
	return tmcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tmcq *TopMostConstraintQuery) Unique(unique bool) *TopMostConstraintQuery {
	tmcq.unique = &unique
	return tmcq
}

// Order adds an order step to the query.
func (tmcq *TopMostConstraintQuery) Order(o ...OrderFunc) *TopMostConstraintQuery {
	tmcq.order = append(tmcq.order, o...)
	return tmcq
}

// First returns the first TopMostConstraint entity from the query.
// Returns a *NotFoundError when no TopMostConstraint was found.
func (tmcq *TopMostConstraintQuery) First(ctx context.Context) (*TopMostConstraint, error) {
	nodes, err := tmcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{topmostconstraint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) FirstX(ctx context.Context) *TopMostConstraint {
	node, err := tmcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TopMostConstraint ID from the query.
// Returns a *NotFoundError when no TopMostConstraint ID was found.
func (tmcq *TopMostConstraintQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{topmostconstraint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := tmcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TopMostConstraint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TopMostConstraint entity is found.
// Returns a *NotFoundError when no TopMostConstraint entities are found.
func (tmcq *TopMostConstraintQuery) Only(ctx context.Context) (*TopMostConstraint, error) {
	nodes, err := tmcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{topmostconstraint.Label}
	default:
		return nil, &NotSingularError{topmostconstraint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) OnlyX(ctx context.Context) *TopMostConstraint {
	node, err := tmcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TopMostConstraint ID in the query.
// Returns a *NotSingularError when more than one TopMostConstraint ID is found.
// Returns a *NotFoundError when no entities are found.
func (tmcq *TopMostConstraintQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = tmcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{topmostconstraint.Label}
	default:
		err = &NotSingularError{topmostconstraint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := tmcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TopMostConstraints.
func (tmcq *TopMostConstraintQuery) All(ctx context.Context) ([]*TopMostConstraint, error) {
	if err := tmcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tmcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) AllX(ctx context.Context) []*TopMostConstraint {
	nodes, err := tmcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TopMostConstraint IDs.
func (tmcq *TopMostConstraintQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := tmcq.Select(topmostconstraint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := tmcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tmcq *TopMostConstraintQuery) Count(ctx context.Context) (int, error) {
	if err := tmcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tmcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) CountX(ctx context.Context) int {
	count, err := tmcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tmcq *TopMostConstraintQuery) Exist(ctx context.Context) (bool, error) {
	if err := tmcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tmcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tmcq *TopMostConstraintQuery) ExistX(ctx context.Context) bool {
	exist, err := tmcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TopMostConstraintQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tmcq *TopMostConstraintQuery) Clone() *TopMostConstraintQuery {
	if tmcq == nil {
		return nil
	}
	return &TopMostConstraintQuery{
		config:     tmcq.config,
		limit:      tmcq.limit,
		offset:     tmcq.offset,
		order:      append([]OrderFunc{}, tmcq.order...),
		predicates: append([]predicate.TopMostConstraint{}, tmcq.predicates...),
		// clone intermediate query.
		sql:    tmcq.sql.Clone(),
		path:   tmcq.path,
		unique: tmcq.unique,
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
//	client.TopMostConstraint.Query().
//		GroupBy(topmostconstraint.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tmcq *TopMostConstraintQuery) GroupBy(field string, fields ...string) *TopMostConstraintGroupBy {
	grbuild := &TopMostConstraintGroupBy{config: tmcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tmcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tmcq.sqlQuery(ctx), nil
	}
	grbuild.label = topmostconstraint.Label
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
//	client.TopMostConstraint.Query().
//		Select(topmostconstraint.FieldCreatedAt).
//		Scan(ctx, &v)
func (tmcq *TopMostConstraintQuery) Select(fields ...string) *TopMostConstraintSelect {
	tmcq.fields = append(tmcq.fields, fields...)
	selbuild := &TopMostConstraintSelect{TopMostConstraintQuery: tmcq}
	selbuild.label = topmostconstraint.Label
	selbuild.flds, selbuild.scan = &tmcq.fields, selbuild.Scan
	return selbuild
}

func (tmcq *TopMostConstraintQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tmcq.fields {
		if !topmostconstraint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tmcq.path != nil {
		prev, err := tmcq.path(ctx)
		if err != nil {
			return err
		}
		tmcq.sql = prev
	}
	if topmostconstraint.Policy == nil {
		return errors.New("ent: uninitialized topmostconstraint.Policy (forgotten import ent/runtime?)")
	}
	if err := topmostconstraint.Policy.EvalQuery(ctx, tmcq); err != nil {
		return err
	}
	return nil
}

func (tmcq *TopMostConstraintQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TopMostConstraint, error) {
	var (
		nodes = []*TopMostConstraint{}
		_spec = tmcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*TopMostConstraint).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &TopMostConstraint{config: tmcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(tmcq.modifiers) > 0 {
		_spec.Modifiers = tmcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tmcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tmcq *TopMostConstraintQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tmcq.querySpec()
	if len(tmcq.modifiers) > 0 {
		_spec.Modifiers = tmcq.modifiers
	}
	_spec.Node.Columns = tmcq.fields
	if len(tmcq.fields) > 0 {
		_spec.Unique = tmcq.unique != nil && *tmcq.unique
	}
	return sqlgraph.CountNodes(ctx, tmcq.driver, _spec)
}

func (tmcq *TopMostConstraintQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tmcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tmcq *TopMostConstraintQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   topmostconstraint.Table,
			Columns: topmostconstraint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: topmostconstraint.FieldID,
			},
		},
		From:   tmcq.sql,
		Unique: true,
	}
	if unique := tmcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tmcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, topmostconstraint.FieldID)
		for i := range fields {
			if fields[i] != topmostconstraint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tmcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tmcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tmcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tmcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tmcq *TopMostConstraintQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tmcq.driver.Dialect())
	t1 := builder.Table(topmostconstraint.Table)
	columns := tmcq.fields
	if len(columns) == 0 {
		columns = topmostconstraint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tmcq.sql != nil {
		selector = tmcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tmcq.unique != nil && *tmcq.unique {
		selector.Distinct()
	}
	for _, m := range tmcq.modifiers {
		m(selector)
	}
	for _, p := range tmcq.predicates {
		p(selector)
	}
	for _, p := range tmcq.order {
		p(selector)
	}
	if offset := tmcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tmcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tmcq *TopMostConstraintQuery) ForUpdate(opts ...sql.LockOption) *TopMostConstraintQuery {
	if tmcq.driver.Dialect() == dialect.Postgres {
		tmcq.Unique(false)
	}
	tmcq.modifiers = append(tmcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tmcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tmcq *TopMostConstraintQuery) ForShare(opts ...sql.LockOption) *TopMostConstraintQuery {
	if tmcq.driver.Dialect() == dialect.Postgres {
		tmcq.Unique(false)
	}
	tmcq.modifiers = append(tmcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tmcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmcq *TopMostConstraintQuery) Modify(modifiers ...func(s *sql.Selector)) *TopMostConstraintSelect {
	tmcq.modifiers = append(tmcq.modifiers, modifiers...)
	return tmcq.Select()
}

// TopMostConstraintGroupBy is the group-by builder for TopMostConstraint entities.
type TopMostConstraintGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tmcgb *TopMostConstraintGroupBy) Aggregate(fns ...AggregateFunc) *TopMostConstraintGroupBy {
	tmcgb.fns = append(tmcgb.fns, fns...)
	return tmcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tmcgb *TopMostConstraintGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tmcgb.path(ctx)
	if err != nil {
		return err
	}
	tmcgb.sql = query
	return tmcgb.sqlScan(ctx, v)
}

func (tmcgb *TopMostConstraintGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tmcgb.fields {
		if !topmostconstraint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tmcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tmcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tmcgb *TopMostConstraintGroupBy) sqlQuery() *sql.Selector {
	selector := tmcgb.sql.Select()
	aggregation := make([]string, 0, len(tmcgb.fns))
	for _, fn := range tmcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tmcgb.fields)+len(tmcgb.fns))
		for _, f := range tmcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tmcgb.fields...)...)
}

// TopMostConstraintSelect is the builder for selecting fields of TopMostConstraint entities.
type TopMostConstraintSelect struct {
	*TopMostConstraintQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (tmcs *TopMostConstraintSelect) Scan(ctx context.Context, v interface{}) error {
	if err := tmcs.prepareQuery(ctx); err != nil {
		return err
	}
	tmcs.sql = tmcs.TopMostConstraintQuery.sqlQuery(ctx)
	return tmcs.sqlScan(ctx, v)
}

func (tmcs *TopMostConstraintSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := tmcs.sql.Query()
	if err := tmcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tmcs *TopMostConstraintSelect) Modify(modifiers ...func(s *sql.Selector)) *TopMostConstraintSelect {
	tmcs.modifiers = append(tmcs.modifiers, modifiers...)
	return tmcs
}
