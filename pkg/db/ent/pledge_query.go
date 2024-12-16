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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/pledge"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// PledgeQuery is the builder for querying Pledge entities.
type PledgeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Pledge
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PledgeQuery builder.
func (pq *PledgeQuery) Where(ps ...predicate.Pledge) *PledgeQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PledgeQuery) Limit(limit int) *PledgeQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PledgeQuery) Offset(offset int) *PledgeQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PledgeQuery) Unique(unique bool) *PledgeQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PledgeQuery) Order(o ...OrderFunc) *PledgeQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// First returns the first Pledge entity from the query.
// Returns a *NotFoundError when no Pledge was found.
func (pq *PledgeQuery) First(ctx context.Context) (*Pledge, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pledge.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PledgeQuery) FirstX(ctx context.Context) *Pledge {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pledge ID from the query.
// Returns a *NotFoundError when no Pledge ID was found.
func (pq *PledgeQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pledge.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PledgeQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pledge entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pledge entity is found.
// Returns a *NotFoundError when no Pledge entities are found.
func (pq *PledgeQuery) Only(ctx context.Context) (*Pledge, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pledge.Label}
	default:
		return nil, &NotSingularError{pledge.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PledgeQuery) OnlyX(ctx context.Context) *Pledge {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pledge ID in the query.
// Returns a *NotSingularError when more than one Pledge ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PledgeQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pledge.Label}
	default:
		err = &NotSingularError{pledge.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PledgeQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pledges.
func (pq *PledgeQuery) All(ctx context.Context) ([]*Pledge, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PledgeQuery) AllX(ctx context.Context) []*Pledge {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pledge IDs.
func (pq *PledgeQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := pq.Select(pledge.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PledgeQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PledgeQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PledgeQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PledgeQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PledgeQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PledgeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PledgeQuery) Clone() *PledgeQuery {
	if pq == nil {
		return nil
	}
	return &PledgeQuery{
		config:     pq.config,
		limit:      pq.limit,
		offset:     pq.offset,
		order:      append([]OrderFunc{}, pq.order...),
		predicates: append([]predicate.Pledge{}, pq.predicates...),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
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
//	client.Pledge.Query().
//		GroupBy(pledge.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (pq *PledgeQuery) GroupBy(field string, fields ...string) *PledgeGroupBy {
	grbuild := &PledgeGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = pledge.Label
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
//	client.Pledge.Query().
//		Select(pledge.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (pq *PledgeQuery) Select(fields ...string) *PledgeSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PledgeSelect{PledgeQuery: pq}
	selbuild.label = pledge.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

func (pq *PledgeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !pledge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	if pledge.Policy == nil {
		return errors.New("ent: uninitialized pledge.Policy (forgotten import ent/runtime?)")
	}
	if err := pledge.Policy.EvalQuery(ctx, pq); err != nil {
		return err
	}
	return nil
}

func (pq *PledgeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pledge, error) {
	var (
		nodes = []*Pledge{}
		_spec = pq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Pledge).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Pledge{config: pq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (pq *PledgeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PledgeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := pq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (pq *PledgeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pledge.Table,
			Columns: pledge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: pledge.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pledge.FieldID)
		for i := range fields {
			if fields[i] != pledge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PledgeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pledge.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = pledge.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, m := range pq.modifiers {
		m(selector)
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (pq *PledgeQuery) ForUpdate(opts ...sql.LockOption) *PledgeQuery {
	if pq.driver.Dialect() == dialect.Postgres {
		pq.Unique(false)
	}
	pq.modifiers = append(pq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return pq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (pq *PledgeQuery) ForShare(opts ...sql.LockOption) *PledgeQuery {
	if pq.driver.Dialect() == dialect.Postgres {
		pq.Unique(false)
	}
	pq.modifiers = append(pq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return pq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (pq *PledgeQuery) Modify(modifiers ...func(s *sql.Selector)) *PledgeSelect {
	pq.modifiers = append(pq.modifiers, modifiers...)
	return pq.Select()
}

// PledgeGroupBy is the group-by builder for Pledge entities.
type PledgeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PledgeGroupBy) Aggregate(fns ...AggregateFunc) *PledgeGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PledgeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *PledgeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pgb.fields {
		if !pledge.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PledgeGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PledgeSelect is the builder for selecting fields of Pledge entities.
type PledgeSelect struct {
	*PledgeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PledgeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PledgeQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *PledgeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ps *PledgeSelect) Modify(modifiers ...func(s *sql.Selector)) *PledgeSelect {
	ps.modifiers = append(ps.modifiers, modifiers...)
	return ps
}
