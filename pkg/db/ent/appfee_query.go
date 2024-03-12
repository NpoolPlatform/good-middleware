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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appfee"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppFeeQuery is the builder for querying AppFee entities.
type AppFeeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppFee
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppFeeQuery builder.
func (afq *AppFeeQuery) Where(ps ...predicate.AppFee) *AppFeeQuery {
	afq.predicates = append(afq.predicates, ps...)
	return afq
}

// Limit adds a limit step to the query.
func (afq *AppFeeQuery) Limit(limit int) *AppFeeQuery {
	afq.limit = &limit
	return afq
}

// Offset adds an offset step to the query.
func (afq *AppFeeQuery) Offset(offset int) *AppFeeQuery {
	afq.offset = &offset
	return afq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (afq *AppFeeQuery) Unique(unique bool) *AppFeeQuery {
	afq.unique = &unique
	return afq
}

// Order adds an order step to the query.
func (afq *AppFeeQuery) Order(o ...OrderFunc) *AppFeeQuery {
	afq.order = append(afq.order, o...)
	return afq
}

// First returns the first AppFee entity from the query.
// Returns a *NotFoundError when no AppFee was found.
func (afq *AppFeeQuery) First(ctx context.Context) (*AppFee, error) {
	nodes, err := afq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appfee.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (afq *AppFeeQuery) FirstX(ctx context.Context) *AppFee {
	node, err := afq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppFee ID from the query.
// Returns a *NotFoundError when no AppFee ID was found.
func (afq *AppFeeQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = afq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appfee.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (afq *AppFeeQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := afq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppFee entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppFee entity is found.
// Returns a *NotFoundError when no AppFee entities are found.
func (afq *AppFeeQuery) Only(ctx context.Context) (*AppFee, error) {
	nodes, err := afq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appfee.Label}
	default:
		return nil, &NotSingularError{appfee.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (afq *AppFeeQuery) OnlyX(ctx context.Context) *AppFee {
	node, err := afq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppFee ID in the query.
// Returns a *NotSingularError when more than one AppFee ID is found.
// Returns a *NotFoundError when no entities are found.
func (afq *AppFeeQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = afq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appfee.Label}
	default:
		err = &NotSingularError{appfee.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (afq *AppFeeQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := afq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppFees.
func (afq *AppFeeQuery) All(ctx context.Context) ([]*AppFee, error) {
	if err := afq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return afq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (afq *AppFeeQuery) AllX(ctx context.Context) []*AppFee {
	nodes, err := afq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppFee IDs.
func (afq *AppFeeQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := afq.Select(appfee.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (afq *AppFeeQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := afq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (afq *AppFeeQuery) Count(ctx context.Context) (int, error) {
	if err := afq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return afq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (afq *AppFeeQuery) CountX(ctx context.Context) int {
	count, err := afq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (afq *AppFeeQuery) Exist(ctx context.Context) (bool, error) {
	if err := afq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return afq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (afq *AppFeeQuery) ExistX(ctx context.Context) bool {
	exist, err := afq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppFeeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (afq *AppFeeQuery) Clone() *AppFeeQuery {
	if afq == nil {
		return nil
	}
	return &AppFeeQuery{
		config:     afq.config,
		limit:      afq.limit,
		offset:     afq.offset,
		order:      append([]OrderFunc{}, afq.order...),
		predicates: append([]predicate.AppFee{}, afq.predicates...),
		// clone intermediate query.
		sql:    afq.sql.Clone(),
		path:   afq.path,
		unique: afq.unique,
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
//	client.AppFee.Query().
//		GroupBy(appfee.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (afq *AppFeeQuery) GroupBy(field string, fields ...string) *AppFeeGroupBy {
	grbuild := &AppFeeGroupBy{config: afq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := afq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return afq.sqlQuery(ctx), nil
	}
	grbuild.label = appfee.Label
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
//	client.AppFee.Query().
//		Select(appfee.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (afq *AppFeeQuery) Select(fields ...string) *AppFeeSelect {
	afq.fields = append(afq.fields, fields...)
	selbuild := &AppFeeSelect{AppFeeQuery: afq}
	selbuild.label = appfee.Label
	selbuild.flds, selbuild.scan = &afq.fields, selbuild.Scan
	return selbuild
}

func (afq *AppFeeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range afq.fields {
		if !appfee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if afq.path != nil {
		prev, err := afq.path(ctx)
		if err != nil {
			return err
		}
		afq.sql = prev
	}
	if appfee.Policy == nil {
		return errors.New("ent: uninitialized appfee.Policy (forgotten import ent/runtime?)")
	}
	if err := appfee.Policy.EvalQuery(ctx, afq); err != nil {
		return err
	}
	return nil
}

func (afq *AppFeeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppFee, error) {
	var (
		nodes = []*AppFee{}
		_spec = afq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppFee).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppFee{config: afq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(afq.modifiers) > 0 {
		_spec.Modifiers = afq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, afq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (afq *AppFeeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := afq.querySpec()
	if len(afq.modifiers) > 0 {
		_spec.Modifiers = afq.modifiers
	}
	_spec.Node.Columns = afq.fields
	if len(afq.fields) > 0 {
		_spec.Unique = afq.unique != nil && *afq.unique
	}
	return sqlgraph.CountNodes(ctx, afq.driver, _spec)
}

func (afq *AppFeeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := afq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (afq *AppFeeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appfee.Table,
			Columns: appfee.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appfee.FieldID,
			},
		},
		From:   afq.sql,
		Unique: true,
	}
	if unique := afq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := afq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appfee.FieldID)
		for i := range fields {
			if fields[i] != appfee.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := afq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := afq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := afq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := afq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (afq *AppFeeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(afq.driver.Dialect())
	t1 := builder.Table(appfee.Table)
	columns := afq.fields
	if len(columns) == 0 {
		columns = appfee.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if afq.sql != nil {
		selector = afq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if afq.unique != nil && *afq.unique {
		selector.Distinct()
	}
	for _, m := range afq.modifiers {
		m(selector)
	}
	for _, p := range afq.predicates {
		p(selector)
	}
	for _, p := range afq.order {
		p(selector)
	}
	if offset := afq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := afq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (afq *AppFeeQuery) ForUpdate(opts ...sql.LockOption) *AppFeeQuery {
	if afq.driver.Dialect() == dialect.Postgres {
		afq.Unique(false)
	}
	afq.modifiers = append(afq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return afq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (afq *AppFeeQuery) ForShare(opts ...sql.LockOption) *AppFeeQuery {
	if afq.driver.Dialect() == dialect.Postgres {
		afq.Unique(false)
	}
	afq.modifiers = append(afq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return afq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (afq *AppFeeQuery) Modify(modifiers ...func(s *sql.Selector)) *AppFeeSelect {
	afq.modifiers = append(afq.modifiers, modifiers...)
	return afq.Select()
}

// AppFeeGroupBy is the group-by builder for AppFee entities.
type AppFeeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (afgb *AppFeeGroupBy) Aggregate(fns ...AggregateFunc) *AppFeeGroupBy {
	afgb.fns = append(afgb.fns, fns...)
	return afgb
}

// Scan applies the group-by query and scans the result into the given value.
func (afgb *AppFeeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := afgb.path(ctx)
	if err != nil {
		return err
	}
	afgb.sql = query
	return afgb.sqlScan(ctx, v)
}

func (afgb *AppFeeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range afgb.fields {
		if !appfee.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := afgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := afgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (afgb *AppFeeGroupBy) sqlQuery() *sql.Selector {
	selector := afgb.sql.Select()
	aggregation := make([]string, 0, len(afgb.fns))
	for _, fn := range afgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(afgb.fields)+len(afgb.fns))
		for _, f := range afgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(afgb.fields...)...)
}

// AppFeeSelect is the builder for selecting fields of AppFee entities.
type AppFeeSelect struct {
	*AppFeeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (afs *AppFeeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := afs.prepareQuery(ctx); err != nil {
		return err
	}
	afs.sql = afs.AppFeeQuery.sqlQuery(ctx)
	return afs.sqlScan(ctx, v)
}

func (afs *AppFeeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := afs.sql.Query()
	if err := afs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (afs *AppFeeSelect) Modify(modifiers ...func(s *sql.Selector)) *AppFeeSelect {
	afs.modifiers = append(afs.modifiers, modifiers...)
	return afs
}
