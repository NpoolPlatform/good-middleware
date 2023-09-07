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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// AppStockQuery is the builder for querying AppStock entities.
type AppStockQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppStock
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppStockQuery builder.
func (asq *AppStockQuery) Where(ps ...predicate.AppStock) *AppStockQuery {
	asq.predicates = append(asq.predicates, ps...)
	return asq
}

// Limit adds a limit step to the query.
func (asq *AppStockQuery) Limit(limit int) *AppStockQuery {
	asq.limit = &limit
	return asq
}

// Offset adds an offset step to the query.
func (asq *AppStockQuery) Offset(offset int) *AppStockQuery {
	asq.offset = &offset
	return asq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (asq *AppStockQuery) Unique(unique bool) *AppStockQuery {
	asq.unique = &unique
	return asq
}

// Order adds an order step to the query.
func (asq *AppStockQuery) Order(o ...OrderFunc) *AppStockQuery {
	asq.order = append(asq.order, o...)
	return asq
}

// First returns the first AppStock entity from the query.
// Returns a *NotFoundError when no AppStock was found.
func (asq *AppStockQuery) First(ctx context.Context) (*AppStock, error) {
	nodes, err := asq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appstock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (asq *AppStockQuery) FirstX(ctx context.Context) *AppStock {
	node, err := asq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppStock ID from the query.
// Returns a *NotFoundError when no AppStock ID was found.
func (asq *AppStockQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = asq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appstock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (asq *AppStockQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := asq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppStock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppStock entity is found.
// Returns a *NotFoundError when no AppStock entities are found.
func (asq *AppStockQuery) Only(ctx context.Context) (*AppStock, error) {
	nodes, err := asq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appstock.Label}
	default:
		return nil, &NotSingularError{appstock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (asq *AppStockQuery) OnlyX(ctx context.Context) *AppStock {
	node, err := asq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppStock ID in the query.
// Returns a *NotSingularError when more than one AppStock ID is found.
// Returns a *NotFoundError when no entities are found.
func (asq *AppStockQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = asq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appstock.Label}
	default:
		err = &NotSingularError{appstock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (asq *AppStockQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := asq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppStocks.
func (asq *AppStockQuery) All(ctx context.Context) ([]*AppStock, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return asq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (asq *AppStockQuery) AllX(ctx context.Context) []*AppStock {
	nodes, err := asq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppStock IDs.
func (asq *AppStockQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := asq.Select(appstock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (asq *AppStockQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := asq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (asq *AppStockQuery) Count(ctx context.Context) (int, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return asq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (asq *AppStockQuery) CountX(ctx context.Context) int {
	count, err := asq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (asq *AppStockQuery) Exist(ctx context.Context) (bool, error) {
	if err := asq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return asq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (asq *AppStockQuery) ExistX(ctx context.Context) bool {
	exist, err := asq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppStockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (asq *AppStockQuery) Clone() *AppStockQuery {
	if asq == nil {
		return nil
	}
	return &AppStockQuery{
		config:     asq.config,
		limit:      asq.limit,
		offset:     asq.offset,
		order:      append([]OrderFunc{}, asq.order...),
		predicates: append([]predicate.AppStock{}, asq.predicates...),
		// clone intermediate query.
		sql:    asq.sql.Clone(),
		path:   asq.path,
		unique: asq.unique,
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
//	client.AppStock.Query().
//		GroupBy(appstock.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (asq *AppStockQuery) GroupBy(field string, fields ...string) *AppStockGroupBy {
	grbuild := &AppStockGroupBy{config: asq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := asq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return asq.sqlQuery(ctx), nil
	}
	grbuild.label = appstock.Label
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
//	client.AppStock.Query().
//		Select(appstock.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (asq *AppStockQuery) Select(fields ...string) *AppStockSelect {
	asq.fields = append(asq.fields, fields...)
	selbuild := &AppStockSelect{AppStockQuery: asq}
	selbuild.label = appstock.Label
	selbuild.flds, selbuild.scan = &asq.fields, selbuild.Scan
	return selbuild
}

func (asq *AppStockQuery) prepareQuery(ctx context.Context) error {
	for _, f := range asq.fields {
		if !appstock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if asq.path != nil {
		prev, err := asq.path(ctx)
		if err != nil {
			return err
		}
		asq.sql = prev
	}
	if appstock.Policy == nil {
		return errors.New("ent: uninitialized appstock.Policy (forgotten import ent/runtime?)")
	}
	if err := appstock.Policy.EvalQuery(ctx, asq); err != nil {
		return err
	}
	return nil
}

func (asq *AppStockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppStock, error) {
	var (
		nodes = []*AppStock{}
		_spec = asq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppStock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppStock{config: asq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(asq.modifiers) > 0 {
		_spec.Modifiers = asq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, asq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (asq *AppStockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := asq.querySpec()
	if len(asq.modifiers) > 0 {
		_spec.Modifiers = asq.modifiers
	}
	_spec.Node.Columns = asq.fields
	if len(asq.fields) > 0 {
		_spec.Unique = asq.unique != nil && *asq.unique
	}
	return sqlgraph.CountNodes(ctx, asq.driver, _spec)
}

func (asq *AppStockQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := asq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (asq *AppStockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appstock.Table,
			Columns: appstock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appstock.FieldID,
			},
		},
		From:   asq.sql,
		Unique: true,
	}
	if unique := asq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := asq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appstock.FieldID)
		for i := range fields {
			if fields[i] != appstock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := asq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := asq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := asq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := asq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (asq *AppStockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(asq.driver.Dialect())
	t1 := builder.Table(appstock.Table)
	columns := asq.fields
	if len(columns) == 0 {
		columns = appstock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if asq.sql != nil {
		selector = asq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if asq.unique != nil && *asq.unique {
		selector.Distinct()
	}
	for _, m := range asq.modifiers {
		m(selector)
	}
	for _, p := range asq.predicates {
		p(selector)
	}
	for _, p := range asq.order {
		p(selector)
	}
	if offset := asq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := asq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (asq *AppStockQuery) ForUpdate(opts ...sql.LockOption) *AppStockQuery {
	if asq.driver.Dialect() == dialect.Postgres {
		asq.Unique(false)
	}
	asq.modifiers = append(asq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return asq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (asq *AppStockQuery) ForShare(opts ...sql.LockOption) *AppStockQuery {
	if asq.driver.Dialect() == dialect.Postgres {
		asq.Unique(false)
	}
	asq.modifiers = append(asq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return asq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asq *AppStockQuery) Modify(modifiers ...func(s *sql.Selector)) *AppStockSelect {
	asq.modifiers = append(asq.modifiers, modifiers...)
	return asq.Select()
}

// AppStockGroupBy is the group-by builder for AppStock entities.
type AppStockGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (asgb *AppStockGroupBy) Aggregate(fns ...AggregateFunc) *AppStockGroupBy {
	asgb.fns = append(asgb.fns, fns...)
	return asgb
}

// Scan applies the group-by query and scans the result into the given value.
func (asgb *AppStockGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := asgb.path(ctx)
	if err != nil {
		return err
	}
	asgb.sql = query
	return asgb.sqlScan(ctx, v)
}

func (asgb *AppStockGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range asgb.fields {
		if !appstock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := asgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := asgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (asgb *AppStockGroupBy) sqlQuery() *sql.Selector {
	selector := asgb.sql.Select()
	aggregation := make([]string, 0, len(asgb.fns))
	for _, fn := range asgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(asgb.fields)+len(asgb.fns))
		for _, f := range asgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(asgb.fields...)...)
}

// AppStockSelect is the builder for selecting fields of AppStock entities.
type AppStockSelect struct {
	*AppStockQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ass *AppStockSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ass.prepareQuery(ctx); err != nil {
		return err
	}
	ass.sql = ass.AppStockQuery.sqlQuery(ctx)
	return ass.sqlScan(ctx, v)
}

func (ass *AppStockSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ass.sql.Query()
	if err := ass.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ass *AppStockSelect) Modify(modifiers ...func(s *sql.Selector)) *AppStockSelect {
	ass.modifiers = append(ass.modifiers, modifiers...)
	return ass
}
