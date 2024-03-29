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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appsimulategood"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppSimulateGoodQuery is the builder for querying AppSimulateGood entities.
type AppSimulateGoodQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppSimulateGood
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppSimulateGoodQuery builder.
func (asgq *AppSimulateGoodQuery) Where(ps ...predicate.AppSimulateGood) *AppSimulateGoodQuery {
	asgq.predicates = append(asgq.predicates, ps...)
	return asgq
}

// Limit adds a limit step to the query.
func (asgq *AppSimulateGoodQuery) Limit(limit int) *AppSimulateGoodQuery {
	asgq.limit = &limit
	return asgq
}

// Offset adds an offset step to the query.
func (asgq *AppSimulateGoodQuery) Offset(offset int) *AppSimulateGoodQuery {
	asgq.offset = &offset
	return asgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asgq *AppSimulateGoodQuery) Unique(unique bool) *AppSimulateGoodQuery {
	asgq.unique = &unique
	return asgq
}

// Order adds an order step to the query.
func (asgq *AppSimulateGoodQuery) Order(o ...OrderFunc) *AppSimulateGoodQuery {
	asgq.order = append(asgq.order, o...)
	return asgq
}

// First returns the first AppSimulateGood entity from the query.
// Returns a *NotFoundError when no AppSimulateGood was found.
func (asgq *AppSimulateGoodQuery) First(ctx context.Context) (*AppSimulateGood, error) {
	nodes, err := asgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appsimulategood.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) FirstX(ctx context.Context) *AppSimulateGood {
	node, err := asgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppSimulateGood ID from the query.
// Returns a *NotFoundError when no AppSimulateGood ID was found.
func (asgq *AppSimulateGoodQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = asgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appsimulategood.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := asgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppSimulateGood entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppSimulateGood entity is found.
// Returns a *NotFoundError when no AppSimulateGood entities are found.
func (asgq *AppSimulateGoodQuery) Only(ctx context.Context) (*AppSimulateGood, error) {
	nodes, err := asgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appsimulategood.Label}
	default:
		return nil, &NotSingularError{appsimulategood.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) OnlyX(ctx context.Context) *AppSimulateGood {
	node, err := asgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppSimulateGood ID in the query.
// Returns a *NotSingularError when more than one AppSimulateGood ID is found.
// Returns a *NotFoundError when no entities are found.
func (asgq *AppSimulateGoodQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = asgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appsimulategood.Label}
	default:
		err = &NotSingularError{appsimulategood.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := asgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppSimulateGoods.
func (asgq *AppSimulateGoodQuery) All(ctx context.Context) ([]*AppSimulateGood, error) {
	if err := asgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return asgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) AllX(ctx context.Context) []*AppSimulateGood {
	nodes, err := asgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppSimulateGood IDs.
func (asgq *AppSimulateGoodQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := asgq.Select(appsimulategood.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := asgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asgq *AppSimulateGoodQuery) Count(ctx context.Context) (int, error) {
	if err := asgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return asgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) CountX(ctx context.Context) int {
	count, err := asgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asgq *AppSimulateGoodQuery) Exist(ctx context.Context) (bool, error) {
	if err := asgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return asgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (asgq *AppSimulateGoodQuery) ExistX(ctx context.Context) bool {
	exist, err := asgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppSimulateGoodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asgq *AppSimulateGoodQuery) Clone() *AppSimulateGoodQuery {
	if asgq == nil {
		return nil
	}
	return &AppSimulateGoodQuery{
		config:     asgq.config,
		limit:      asgq.limit,
		offset:     asgq.offset,
		order:      append([]OrderFunc{}, asgq.order...),
		predicates: append([]predicate.AppSimulateGood{}, asgq.predicates...),
		// clone intermediate query.
		sql:    asgq.sql.Clone(),
		path:   asgq.path,
		unique: asgq.unique,
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
//	client.AppSimulateGood.Query().
//		GroupBy(appsimulategood.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (asgq *AppSimulateGoodQuery) GroupBy(field string, fields ...string) *AppSimulateGoodGroupBy {
	grbuild := &AppSimulateGoodGroupBy{config: asgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := asgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return asgq.sqlQuery(ctx), nil
	}
	grbuild.label = appsimulategood.Label
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
//	client.AppSimulateGood.Query().
//		Select(appsimulategood.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (asgq *AppSimulateGoodQuery) Select(fields ...string) *AppSimulateGoodSelect {
	asgq.fields = append(asgq.fields, fields...)
	selbuild := &AppSimulateGoodSelect{AppSimulateGoodQuery: asgq}
	selbuild.label = appsimulategood.Label
	selbuild.flds, selbuild.scan = &asgq.fields, selbuild.Scan
	return selbuild
}

func (asgq *AppSimulateGoodQuery) prepareQuery(ctx context.Context) error {
	for _, f := range asgq.fields {
		if !appsimulategood.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if asgq.path != nil {
		prev, err := asgq.path(ctx)
		if err != nil {
			return err
		}
		asgq.sql = prev
	}
	if appsimulategood.Policy == nil {
		return errors.New("ent: uninitialized appsimulategood.Policy (forgotten import ent/runtime?)")
	}
	if err := appsimulategood.Policy.EvalQuery(ctx, asgq); err != nil {
		return err
	}
	return nil
}

func (asgq *AppSimulateGoodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppSimulateGood, error) {
	var (
		nodes = []*AppSimulateGood{}
		_spec = asgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppSimulateGood).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppSimulateGood{config: asgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(asgq.modifiers) > 0 {
		_spec.Modifiers = asgq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, asgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (asgq *AppSimulateGoodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asgq.querySpec()
	if len(asgq.modifiers) > 0 {
		_spec.Modifiers = asgq.modifiers
	}
	_spec.Node.Columns = asgq.fields
	if len(asgq.fields) > 0 {
		_spec.Unique = asgq.unique != nil && *asgq.unique
	}
	return sqlgraph.CountNodes(ctx, asgq.driver, _spec)
}

func (asgq *AppSimulateGoodQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := asgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (asgq *AppSimulateGoodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appsimulategood.Table,
			Columns: appsimulategood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appsimulategood.FieldID,
			},
		},
		From:   asgq.sql,
		Unique: true,
	}
	if unique := asgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := asgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appsimulategood.FieldID)
		for i := range fields {
			if fields[i] != appsimulategood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := asgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asgq *AppSimulateGoodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asgq.driver.Dialect())
	t1 := builder.Table(appsimulategood.Table)
	columns := asgq.fields
	if len(columns) == 0 {
		columns = appsimulategood.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asgq.sql != nil {
		selector = asgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if asgq.unique != nil && *asgq.unique {
		selector.Distinct()
	}
	for _, m := range asgq.modifiers {
		m(selector)
	}
	for _, p := range asgq.predicates {
		p(selector)
	}
	for _, p := range asgq.order {
		p(selector)
	}
	if offset := asgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (asgq *AppSimulateGoodQuery) ForUpdate(opts ...sql.LockOption) *AppSimulateGoodQuery {
	if asgq.driver.Dialect() == dialect.Postgres {
		asgq.Unique(false)
	}
	asgq.modifiers = append(asgq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return asgq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (asgq *AppSimulateGoodQuery) ForShare(opts ...sql.LockOption) *AppSimulateGoodQuery {
	if asgq.driver.Dialect() == dialect.Postgres {
		asgq.Unique(false)
	}
	asgq.modifiers = append(asgq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return asgq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asgq *AppSimulateGoodQuery) Modify(modifiers ...func(s *sql.Selector)) *AppSimulateGoodSelect {
	asgq.modifiers = append(asgq.modifiers, modifiers...)
	return asgq.Select()
}

// AppSimulateGoodGroupBy is the group-by builder for AppSimulateGood entities.
type AppSimulateGoodGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asggb *AppSimulateGoodGroupBy) Aggregate(fns ...AggregateFunc) *AppSimulateGoodGroupBy {
	asggb.fns = append(asggb.fns, fns...)
	return asggb
}

// Scan applies the group-by query and scans the result into the given value.
func (asggb *AppSimulateGoodGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := asggb.path(ctx)
	if err != nil {
		return err
	}
	asggb.sql = query
	return asggb.sqlScan(ctx, v)
}

func (asggb *AppSimulateGoodGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range asggb.fields {
		if !appsimulategood.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := asggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (asggb *AppSimulateGoodGroupBy) sqlQuery() *sql.Selector {
	selector := asggb.sql.Select()
	aggregation := make([]string, 0, len(asggb.fns))
	for _, fn := range asggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(asggb.fields)+len(asggb.fns))
		for _, f := range asggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(asggb.fields...)...)
}

// AppSimulateGoodSelect is the builder for selecting fields of AppSimulateGood entities.
type AppSimulateGoodSelect struct {
	*AppSimulateGoodQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (asgs *AppSimulateGoodSelect) Scan(ctx context.Context, v interface{}) error {
	if err := asgs.prepareQuery(ctx); err != nil {
		return err
	}
	asgs.sql = asgs.AppSimulateGoodQuery.sqlQuery(ctx)
	return asgs.sqlScan(ctx, v)
}

func (asgs *AppSimulateGoodSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := asgs.sql.Query()
	if err := asgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asgs *AppSimulateGoodSelect) Modify(modifiers ...func(s *sql.Selector)) *AppSimulateGoodSelect {
	asgs.modifiers = append(asgs.modifiers, modifiers...)
	return asgs
}
