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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppGoodQuery is the builder for querying AppGood entities.
type AppGoodQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppGood
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppGoodQuery builder.
func (agq *AppGoodQuery) Where(ps ...predicate.AppGood) *AppGoodQuery {
	agq.predicates = append(agq.predicates, ps...)
	return agq
}

// Limit adds a limit step to the query.
func (agq *AppGoodQuery) Limit(limit int) *AppGoodQuery {
	agq.limit = &limit
	return agq
}

// Offset adds an offset step to the query.
func (agq *AppGoodQuery) Offset(offset int) *AppGoodQuery {
	agq.offset = &offset
	return agq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (agq *AppGoodQuery) Unique(unique bool) *AppGoodQuery {
	agq.unique = &unique
	return agq
}

// Order adds an order step to the query.
func (agq *AppGoodQuery) Order(o ...OrderFunc) *AppGoodQuery {
	agq.order = append(agq.order, o...)
	return agq
}

// First returns the first AppGood entity from the query.
// Returns a *NotFoundError when no AppGood was found.
func (agq *AppGoodQuery) First(ctx context.Context) (*AppGood, error) {
	nodes, err := agq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appgood.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (agq *AppGoodQuery) FirstX(ctx context.Context) *AppGood {
	node, err := agq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppGood ID from the query.
// Returns a *NotFoundError when no AppGood ID was found.
func (agq *AppGoodQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appgood.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (agq *AppGoodQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := agq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppGood entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppGood entity is found.
// Returns a *NotFoundError when no AppGood entities are found.
func (agq *AppGoodQuery) Only(ctx context.Context) (*AppGood, error) {
	nodes, err := agq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appgood.Label}
	default:
		return nil, &NotSingularError{appgood.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (agq *AppGoodQuery) OnlyX(ctx context.Context) *AppGood {
	node, err := agq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppGood ID in the query.
// Returns a *NotSingularError when more than one AppGood ID is found.
// Returns a *NotFoundError when no entities are found.
func (agq *AppGoodQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appgood.Label}
	default:
		err = &NotSingularError{appgood.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (agq *AppGoodQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := agq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppGoods.
func (agq *AppGoodQuery) All(ctx context.Context) ([]*AppGood, error) {
	if err := agq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return agq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (agq *AppGoodQuery) AllX(ctx context.Context) []*AppGood {
	nodes, err := agq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppGood IDs.
func (agq *AppGoodQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := agq.Select(appgood.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (agq *AppGoodQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := agq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (agq *AppGoodQuery) Count(ctx context.Context) (int, error) {
	if err := agq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return agq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (agq *AppGoodQuery) CountX(ctx context.Context) int {
	count, err := agq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (agq *AppGoodQuery) Exist(ctx context.Context) (bool, error) {
	if err := agq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return agq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (agq *AppGoodQuery) ExistX(ctx context.Context) bool {
	exist, err := agq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppGoodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (agq *AppGoodQuery) Clone() *AppGoodQuery {
	if agq == nil {
		return nil
	}
	return &AppGoodQuery{
		config:     agq.config,
		limit:      agq.limit,
		offset:     agq.offset,
		order:      append([]OrderFunc{}, agq.order...),
		predicates: append([]predicate.AppGood{}, agq.predicates...),
		// clone intermediate query.
		sql:    agq.sql.Clone(),
		path:   agq.path,
		unique: agq.unique,
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
//	client.AppGood.Query().
//		GroupBy(appgood.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (agq *AppGoodQuery) GroupBy(field string, fields ...string) *AppGoodGroupBy {
	grbuild := &AppGoodGroupBy{config: agq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := agq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return agq.sqlQuery(ctx), nil
	}
	grbuild.label = appgood.Label
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
//	client.AppGood.Query().
//		Select(appgood.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (agq *AppGoodQuery) Select(fields ...string) *AppGoodSelect {
	agq.fields = append(agq.fields, fields...)
	selbuild := &AppGoodSelect{AppGoodQuery: agq}
	selbuild.label = appgood.Label
	selbuild.flds, selbuild.scan = &agq.fields, selbuild.Scan
	return selbuild
}

func (agq *AppGoodQuery) prepareQuery(ctx context.Context) error {
	for _, f := range agq.fields {
		if !appgood.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if agq.path != nil {
		prev, err := agq.path(ctx)
		if err != nil {
			return err
		}
		agq.sql = prev
	}
	if appgood.Policy == nil {
		return errors.New("ent: uninitialized appgood.Policy (forgotten import ent/runtime?)")
	}
	if err := appgood.Policy.EvalQuery(ctx, agq); err != nil {
		return err
	}
	return nil
}

func (agq *AppGoodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppGood, error) {
	var (
		nodes = []*AppGood{}
		_spec = agq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppGood).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppGood{config: agq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(agq.modifiers) > 0 {
		_spec.Modifiers = agq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, agq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (agq *AppGoodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := agq.querySpec()
	if len(agq.modifiers) > 0 {
		_spec.Modifiers = agq.modifiers
	}
	_spec.Node.Columns = agq.fields
	if len(agq.fields) > 0 {
		_spec.Unique = agq.unique != nil && *agq.unique
	}
	return sqlgraph.CountNodes(ctx, agq.driver, _spec)
}

func (agq *AppGoodQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := agq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (agq *AppGoodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgood.Table,
			Columns: appgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgood.FieldID,
			},
		},
		From:   agq.sql,
		Unique: true,
	}
	if unique := agq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := agq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appgood.FieldID)
		for i := range fields {
			if fields[i] != appgood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := agq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := agq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := agq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := agq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (agq *AppGoodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(agq.driver.Dialect())
	t1 := builder.Table(appgood.Table)
	columns := agq.fields
	if len(columns) == 0 {
		columns = appgood.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if agq.sql != nil {
		selector = agq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if agq.unique != nil && *agq.unique {
		selector.Distinct()
	}
	for _, m := range agq.modifiers {
		m(selector)
	}
	for _, p := range agq.predicates {
		p(selector)
	}
	for _, p := range agq.order {
		p(selector)
	}
	if offset := agq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := agq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (agq *AppGoodQuery) ForUpdate(opts ...sql.LockOption) *AppGoodQuery {
	if agq.driver.Dialect() == dialect.Postgres {
		agq.Unique(false)
	}
	agq.modifiers = append(agq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return agq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (agq *AppGoodQuery) ForShare(opts ...sql.LockOption) *AppGoodQuery {
	if agq.driver.Dialect() == dialect.Postgres {
		agq.Unique(false)
	}
	agq.modifiers = append(agq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return agq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agq *AppGoodQuery) Modify(modifiers ...func(s *sql.Selector)) *AppGoodSelect {
	agq.modifiers = append(agq.modifiers, modifiers...)
	return agq.Select()
}

// AppGoodGroupBy is the group-by builder for AppGood entities.
type AppGoodGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aggb *AppGoodGroupBy) Aggregate(fns ...AggregateFunc) *AppGoodGroupBy {
	aggb.fns = append(aggb.fns, fns...)
	return aggb
}

// Scan applies the group-by query and scans the result into the given value.
func (aggb *AppGoodGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aggb.path(ctx)
	if err != nil {
		return err
	}
	aggb.sql = query
	return aggb.sqlScan(ctx, v)
}

func (aggb *AppGoodGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aggb.fields {
		if !appgood.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aggb *AppGoodGroupBy) sqlQuery() *sql.Selector {
	selector := aggb.sql.Select()
	aggregation := make([]string, 0, len(aggb.fns))
	for _, fn := range aggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aggb.fields)+len(aggb.fns))
		for _, f := range aggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aggb.fields...)...)
}

// AppGoodSelect is the builder for selecting fields of AppGood entities.
type AppGoodSelect struct {
	*AppGoodQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ags *AppGoodSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ags.prepareQuery(ctx); err != nil {
		return err
	}
	ags.sql = ags.AppGoodQuery.sqlQuery(ctx)
	return ags.sqlScan(ctx, v)
}

func (ags *AppGoodSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ags.sql.Query()
	if err := ags.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ags *AppGoodSelect) Modify(modifiers ...func(s *sql.Selector)) *AppGoodSelect {
	ags.modifiers = append(ags.modifiers, modifiers...)
	return ags
}
