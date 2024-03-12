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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// PowerRentalQuery is the builder for querying PowerRental entities.
type PowerRentalQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.PowerRental
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PowerRentalQuery builder.
func (prq *PowerRentalQuery) Where(ps ...predicate.PowerRental) *PowerRentalQuery {
	prq.predicates = append(prq.predicates, ps...)
	return prq
}

// Limit adds a limit step to the query.
func (prq *PowerRentalQuery) Limit(limit int) *PowerRentalQuery {
	prq.limit = &limit
	return prq
}

// Offset adds an offset step to the query.
func (prq *PowerRentalQuery) Offset(offset int) *PowerRentalQuery {
	prq.offset = &offset
	return prq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (prq *PowerRentalQuery) Unique(unique bool) *PowerRentalQuery {
	prq.unique = &unique
	return prq
}

// Order adds an order step to the query.
func (prq *PowerRentalQuery) Order(o ...OrderFunc) *PowerRentalQuery {
	prq.order = append(prq.order, o...)
	return prq
}

// First returns the first PowerRental entity from the query.
// Returns a *NotFoundError when no PowerRental was found.
func (prq *PowerRentalQuery) First(ctx context.Context) (*PowerRental, error) {
	nodes, err := prq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{powerrental.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (prq *PowerRentalQuery) FirstX(ctx context.Context) *PowerRental {
	node, err := prq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PowerRental ID from the query.
// Returns a *NotFoundError when no PowerRental ID was found.
func (prq *PowerRentalQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = prq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{powerrental.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (prq *PowerRentalQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := prq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PowerRental entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PowerRental entity is found.
// Returns a *NotFoundError when no PowerRental entities are found.
func (prq *PowerRentalQuery) Only(ctx context.Context) (*PowerRental, error) {
	nodes, err := prq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{powerrental.Label}
	default:
		return nil, &NotSingularError{powerrental.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (prq *PowerRentalQuery) OnlyX(ctx context.Context) *PowerRental {
	node, err := prq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PowerRental ID in the query.
// Returns a *NotSingularError when more than one PowerRental ID is found.
// Returns a *NotFoundError when no entities are found.
func (prq *PowerRentalQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = prq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{powerrental.Label}
	default:
		err = &NotSingularError{powerrental.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (prq *PowerRentalQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := prq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PowerRentals.
func (prq *PowerRentalQuery) All(ctx context.Context) ([]*PowerRental, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return prq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (prq *PowerRentalQuery) AllX(ctx context.Context) []*PowerRental {
	nodes, err := prq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PowerRental IDs.
func (prq *PowerRentalQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := prq.Select(powerrental.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (prq *PowerRentalQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := prq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (prq *PowerRentalQuery) Count(ctx context.Context) (int, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return prq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (prq *PowerRentalQuery) CountX(ctx context.Context) int {
	count, err := prq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (prq *PowerRentalQuery) Exist(ctx context.Context) (bool, error) {
	if err := prq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return prq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (prq *PowerRentalQuery) ExistX(ctx context.Context) bool {
	exist, err := prq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PowerRentalQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (prq *PowerRentalQuery) Clone() *PowerRentalQuery {
	if prq == nil {
		return nil
	}
	return &PowerRentalQuery{
		config:     prq.config,
		limit:      prq.limit,
		offset:     prq.offset,
		order:      append([]OrderFunc{}, prq.order...),
		predicates: append([]predicate.PowerRental{}, prq.predicates...),
		// clone intermediate query.
		sql:    prq.sql.Clone(),
		path:   prq.path,
		unique: prq.unique,
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
//	client.PowerRental.Query().
//		GroupBy(powerrental.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (prq *PowerRentalQuery) GroupBy(field string, fields ...string) *PowerRentalGroupBy {
	grbuild := &PowerRentalGroupBy{config: prq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := prq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return prq.sqlQuery(ctx), nil
	}
	grbuild.label = powerrental.Label
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
//	client.PowerRental.Query().
//		Select(powerrental.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (prq *PowerRentalQuery) Select(fields ...string) *PowerRentalSelect {
	prq.fields = append(prq.fields, fields...)
	selbuild := &PowerRentalSelect{PowerRentalQuery: prq}
	selbuild.label = powerrental.Label
	selbuild.flds, selbuild.scan = &prq.fields, selbuild.Scan
	return selbuild
}

func (prq *PowerRentalQuery) prepareQuery(ctx context.Context) error {
	for _, f := range prq.fields {
		if !powerrental.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if prq.path != nil {
		prev, err := prq.path(ctx)
		if err != nil {
			return err
		}
		prq.sql = prev
	}
	if powerrental.Policy == nil {
		return errors.New("ent: uninitialized powerrental.Policy (forgotten import ent/runtime?)")
	}
	if err := powerrental.Policy.EvalQuery(ctx, prq); err != nil {
		return err
	}
	return nil
}

func (prq *PowerRentalQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PowerRental, error) {
	var (
		nodes = []*PowerRental{}
		_spec = prq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*PowerRental).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &PowerRental{config: prq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(prq.modifiers) > 0 {
		_spec.Modifiers = prq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, prq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (prq *PowerRentalQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := prq.querySpec()
	if len(prq.modifiers) > 0 {
		_spec.Modifiers = prq.modifiers
	}
	_spec.Node.Columns = prq.fields
	if len(prq.fields) > 0 {
		_spec.Unique = prq.unique != nil && *prq.unique
	}
	return sqlgraph.CountNodes(ctx, prq.driver, _spec)
}

func (prq *PowerRentalQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := prq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (prq *PowerRentalQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   powerrental.Table,
			Columns: powerrental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: powerrental.FieldID,
			},
		},
		From:   prq.sql,
		Unique: true,
	}
	if unique := prq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := prq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, powerrental.FieldID)
		for i := range fields {
			if fields[i] != powerrental.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := prq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := prq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := prq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := prq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (prq *PowerRentalQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(prq.driver.Dialect())
	t1 := builder.Table(powerrental.Table)
	columns := prq.fields
	if len(columns) == 0 {
		columns = powerrental.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if prq.sql != nil {
		selector = prq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if prq.unique != nil && *prq.unique {
		selector.Distinct()
	}
	for _, m := range prq.modifiers {
		m(selector)
	}
	for _, p := range prq.predicates {
		p(selector)
	}
	for _, p := range prq.order {
		p(selector)
	}
	if offset := prq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := prq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (prq *PowerRentalQuery) ForUpdate(opts ...sql.LockOption) *PowerRentalQuery {
	if prq.driver.Dialect() == dialect.Postgres {
		prq.Unique(false)
	}
	prq.modifiers = append(prq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return prq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (prq *PowerRentalQuery) ForShare(opts ...sql.LockOption) *PowerRentalQuery {
	if prq.driver.Dialect() == dialect.Postgres {
		prq.Unique(false)
	}
	prq.modifiers = append(prq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return prq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (prq *PowerRentalQuery) Modify(modifiers ...func(s *sql.Selector)) *PowerRentalSelect {
	prq.modifiers = append(prq.modifiers, modifiers...)
	return prq.Select()
}

// PowerRentalGroupBy is the group-by builder for PowerRental entities.
type PowerRentalGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (prgb *PowerRentalGroupBy) Aggregate(fns ...AggregateFunc) *PowerRentalGroupBy {
	prgb.fns = append(prgb.fns, fns...)
	return prgb
}

// Scan applies the group-by query and scans the result into the given value.
func (prgb *PowerRentalGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := prgb.path(ctx)
	if err != nil {
		return err
	}
	prgb.sql = query
	return prgb.sqlScan(ctx, v)
}

func (prgb *PowerRentalGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range prgb.fields {
		if !powerrental.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := prgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := prgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (prgb *PowerRentalGroupBy) sqlQuery() *sql.Selector {
	selector := prgb.sql.Select()
	aggregation := make([]string, 0, len(prgb.fns))
	for _, fn := range prgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(prgb.fields)+len(prgb.fns))
		for _, f := range prgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(prgb.fields...)...)
}

// PowerRentalSelect is the builder for selecting fields of PowerRental entities.
type PowerRentalSelect struct {
	*PowerRentalQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (prs *PowerRentalSelect) Scan(ctx context.Context, v interface{}) error {
	if err := prs.prepareQuery(ctx); err != nil {
		return err
	}
	prs.sql = prs.PowerRentalQuery.sqlQuery(ctx)
	return prs.sqlScan(ctx, v)
}

func (prs *PowerRentalSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := prs.sql.Query()
	if err := prs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (prs *PowerRentalSelect) Modify(modifiers ...func(s *sql.Selector)) *PowerRentalSelect {
	prs.modifiers = append(prs.modifiers, modifiers...)
	return prs
}
