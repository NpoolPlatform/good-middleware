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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstocklock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppStockLockQuery is the builder for querying AppStockLock entities.
type AppStockLockQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppStockLock
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppStockLockQuery builder.
func (aslq *AppStockLockQuery) Where(ps ...predicate.AppStockLock) *AppStockLockQuery {
	aslq.predicates = append(aslq.predicates, ps...)
	return aslq
}

// Limit adds a limit step to the query.
func (aslq *AppStockLockQuery) Limit(limit int) *AppStockLockQuery {
	aslq.limit = &limit
	return aslq
}

// Offset adds an offset step to the query.
func (aslq *AppStockLockQuery) Offset(offset int) *AppStockLockQuery {
	aslq.offset = &offset
	return aslq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aslq *AppStockLockQuery) Unique(unique bool) *AppStockLockQuery {
	aslq.unique = &unique
	return aslq
}

// Order adds an order step to the query.
func (aslq *AppStockLockQuery) Order(o ...OrderFunc) *AppStockLockQuery {
	aslq.order = append(aslq.order, o...)
	return aslq
}

// First returns the first AppStockLock entity from the query.
// Returns a *NotFoundError when no AppStockLock was found.
func (aslq *AppStockLockQuery) First(ctx context.Context) (*AppStockLock, error) {
	nodes, err := aslq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appstocklock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aslq *AppStockLockQuery) FirstX(ctx context.Context) *AppStockLock {
	node, err := aslq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppStockLock ID from the query.
// Returns a *NotFoundError when no AppStockLock ID was found.
func (aslq *AppStockLockQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = aslq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appstocklock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aslq *AppStockLockQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := aslq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppStockLock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppStockLock entity is found.
// Returns a *NotFoundError when no AppStockLock entities are found.
func (aslq *AppStockLockQuery) Only(ctx context.Context) (*AppStockLock, error) {
	nodes, err := aslq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appstocklock.Label}
	default:
		return nil, &NotSingularError{appstocklock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aslq *AppStockLockQuery) OnlyX(ctx context.Context) *AppStockLock {
	node, err := aslq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppStockLock ID in the query.
// Returns a *NotSingularError when more than one AppStockLock ID is found.
// Returns a *NotFoundError when no entities are found.
func (aslq *AppStockLockQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = aslq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appstocklock.Label}
	default:
		err = &NotSingularError{appstocklock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aslq *AppStockLockQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := aslq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppStockLocks.
func (aslq *AppStockLockQuery) All(ctx context.Context) ([]*AppStockLock, error) {
	if err := aslq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aslq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aslq *AppStockLockQuery) AllX(ctx context.Context) []*AppStockLock {
	nodes, err := aslq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppStockLock IDs.
func (aslq *AppStockLockQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := aslq.Select(appstocklock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aslq *AppStockLockQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := aslq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aslq *AppStockLockQuery) Count(ctx context.Context) (int, error) {
	if err := aslq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aslq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aslq *AppStockLockQuery) CountX(ctx context.Context) int {
	count, err := aslq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aslq *AppStockLockQuery) Exist(ctx context.Context) (bool, error) {
	if err := aslq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aslq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aslq *AppStockLockQuery) ExistX(ctx context.Context) bool {
	exist, err := aslq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppStockLockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aslq *AppStockLockQuery) Clone() *AppStockLockQuery {
	if aslq == nil {
		return nil
	}
	return &AppStockLockQuery{
		config:     aslq.config,
		limit:      aslq.limit,
		offset:     aslq.offset,
		order:      append([]OrderFunc{}, aslq.order...),
		predicates: append([]predicate.AppStockLock{}, aslq.predicates...),
		// clone intermediate query.
		sql:    aslq.sql.Clone(),
		path:   aslq.path,
		unique: aslq.unique,
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
//	client.AppStockLock.Query().
//		GroupBy(appstocklock.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (aslq *AppStockLockQuery) GroupBy(field string, fields ...string) *AppStockLockGroupBy {
	grbuild := &AppStockLockGroupBy{config: aslq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aslq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aslq.sqlQuery(ctx), nil
	}
	grbuild.label = appstocklock.Label
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
//	client.AppStockLock.Query().
//		Select(appstocklock.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (aslq *AppStockLockQuery) Select(fields ...string) *AppStockLockSelect {
	aslq.fields = append(aslq.fields, fields...)
	selbuild := &AppStockLockSelect{AppStockLockQuery: aslq}
	selbuild.label = appstocklock.Label
	selbuild.flds, selbuild.scan = &aslq.fields, selbuild.Scan
	return selbuild
}

func (aslq *AppStockLockQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aslq.fields {
		if !appstocklock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aslq.path != nil {
		prev, err := aslq.path(ctx)
		if err != nil {
			return err
		}
		aslq.sql = prev
	}
	if appstocklock.Policy == nil {
		return errors.New("ent: uninitialized appstocklock.Policy (forgotten import ent/runtime?)")
	}
	if err := appstocklock.Policy.EvalQuery(ctx, aslq); err != nil {
		return err
	}
	return nil
}

func (aslq *AppStockLockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppStockLock, error) {
	var (
		nodes = []*AppStockLock{}
		_spec = aslq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppStockLock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppStockLock{config: aslq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(aslq.modifiers) > 0 {
		_spec.Modifiers = aslq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aslq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aslq *AppStockLockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aslq.querySpec()
	if len(aslq.modifiers) > 0 {
		_spec.Modifiers = aslq.modifiers
	}
	_spec.Node.Columns = aslq.fields
	if len(aslq.fields) > 0 {
		_spec.Unique = aslq.unique != nil && *aslq.unique
	}
	return sqlgraph.CountNodes(ctx, aslq.driver, _spec)
}

func (aslq *AppStockLockQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aslq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aslq *AppStockLockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appstocklock.Table,
			Columns: appstocklock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appstocklock.FieldID,
			},
		},
		From:   aslq.sql,
		Unique: true,
	}
	if unique := aslq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aslq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appstocklock.FieldID)
		for i := range fields {
			if fields[i] != appstocklock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aslq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aslq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aslq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aslq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aslq *AppStockLockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aslq.driver.Dialect())
	t1 := builder.Table(appstocklock.Table)
	columns := aslq.fields
	if len(columns) == 0 {
		columns = appstocklock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aslq.sql != nil {
		selector = aslq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aslq.unique != nil && *aslq.unique {
		selector.Distinct()
	}
	for _, m := range aslq.modifiers {
		m(selector)
	}
	for _, p := range aslq.predicates {
		p(selector)
	}
	for _, p := range aslq.order {
		p(selector)
	}
	if offset := aslq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aslq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (aslq *AppStockLockQuery) ForUpdate(opts ...sql.LockOption) *AppStockLockQuery {
	if aslq.driver.Dialect() == dialect.Postgres {
		aslq.Unique(false)
	}
	aslq.modifiers = append(aslq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return aslq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (aslq *AppStockLockQuery) ForShare(opts ...sql.LockOption) *AppStockLockQuery {
	if aslq.driver.Dialect() == dialect.Postgres {
		aslq.Unique(false)
	}
	aslq.modifiers = append(aslq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return aslq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aslq *AppStockLockQuery) Modify(modifiers ...func(s *sql.Selector)) *AppStockLockSelect {
	aslq.modifiers = append(aslq.modifiers, modifiers...)
	return aslq.Select()
}

// AppStockLockGroupBy is the group-by builder for AppStockLock entities.
type AppStockLockGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aslgb *AppStockLockGroupBy) Aggregate(fns ...AggregateFunc) *AppStockLockGroupBy {
	aslgb.fns = append(aslgb.fns, fns...)
	return aslgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aslgb *AppStockLockGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aslgb.path(ctx)
	if err != nil {
		return err
	}
	aslgb.sql = query
	return aslgb.sqlScan(ctx, v)
}

func (aslgb *AppStockLockGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aslgb.fields {
		if !appstocklock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aslgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aslgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aslgb *AppStockLockGroupBy) sqlQuery() *sql.Selector {
	selector := aslgb.sql.Select()
	aggregation := make([]string, 0, len(aslgb.fns))
	for _, fn := range aslgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aslgb.fields)+len(aslgb.fns))
		for _, f := range aslgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aslgb.fields...)...)
}

// AppStockLockSelect is the builder for selecting fields of AppStockLock entities.
type AppStockLockSelect struct {
	*AppStockLockQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (asls *AppStockLockSelect) Scan(ctx context.Context, v interface{}) error {
	if err := asls.prepareQuery(ctx); err != nil {
		return err
	}
	asls.sql = asls.AppStockLockQuery.sqlQuery(ctx)
	return asls.sqlScan(ctx, v)
}

func (asls *AppStockLockSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := asls.sql.Query()
	if err := asls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (asls *AppStockLockSelect) Modify(modifiers ...func(s *sql.Selector)) *AppStockLockSelect {
	asls.modifiers = append(asls.modifiers, modifiers...)
	return asls
}
