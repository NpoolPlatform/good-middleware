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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppGoodBaseQuery is the builder for querying AppGoodBase entities.
type AppGoodBaseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppGoodBase
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppGoodBaseQuery builder.
func (agbq *AppGoodBaseQuery) Where(ps ...predicate.AppGoodBase) *AppGoodBaseQuery {
	agbq.predicates = append(agbq.predicates, ps...)
	return agbq
}

// Limit adds a limit step to the query.
func (agbq *AppGoodBaseQuery) Limit(limit int) *AppGoodBaseQuery {
	agbq.limit = &limit
	return agbq
}

// Offset adds an offset step to the query.
func (agbq *AppGoodBaseQuery) Offset(offset int) *AppGoodBaseQuery {
	agbq.offset = &offset
	return agbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (agbq *AppGoodBaseQuery) Unique(unique bool) *AppGoodBaseQuery {
	agbq.unique = &unique
	return agbq
}

// Order adds an order step to the query.
func (agbq *AppGoodBaseQuery) Order(o ...OrderFunc) *AppGoodBaseQuery {
	agbq.order = append(agbq.order, o...)
	return agbq
}

// First returns the first AppGoodBase entity from the query.
// Returns a *NotFoundError when no AppGoodBase was found.
func (agbq *AppGoodBaseQuery) First(ctx context.Context) (*AppGoodBase, error) {
	nodes, err := agbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appgoodbase.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) FirstX(ctx context.Context) *AppGoodBase {
	node, err := agbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppGoodBase ID from the query.
// Returns a *NotFoundError when no AppGoodBase ID was found.
func (agbq *AppGoodBaseQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appgoodbase.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := agbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppGoodBase entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppGoodBase entity is found.
// Returns a *NotFoundError when no AppGoodBase entities are found.
func (agbq *AppGoodBaseQuery) Only(ctx context.Context) (*AppGoodBase, error) {
	nodes, err := agbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appgoodbase.Label}
	default:
		return nil, &NotSingularError{appgoodbase.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) OnlyX(ctx context.Context) *AppGoodBase {
	node, err := agbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppGoodBase ID in the query.
// Returns a *NotSingularError when more than one AppGoodBase ID is found.
// Returns a *NotFoundError when no entities are found.
func (agbq *AppGoodBaseQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appgoodbase.Label}
	default:
		err = &NotSingularError{appgoodbase.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := agbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppGoodBases.
func (agbq *AppGoodBaseQuery) All(ctx context.Context) ([]*AppGoodBase, error) {
	if err := agbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return agbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) AllX(ctx context.Context) []*AppGoodBase {
	nodes, err := agbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppGoodBase IDs.
func (agbq *AppGoodBaseQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := agbq.Select(appgoodbase.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := agbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (agbq *AppGoodBaseQuery) Count(ctx context.Context) (int, error) {
	if err := agbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return agbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) CountX(ctx context.Context) int {
	count, err := agbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (agbq *AppGoodBaseQuery) Exist(ctx context.Context) (bool, error) {
	if err := agbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return agbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (agbq *AppGoodBaseQuery) ExistX(ctx context.Context) bool {
	exist, err := agbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppGoodBaseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (agbq *AppGoodBaseQuery) Clone() *AppGoodBaseQuery {
	if agbq == nil {
		return nil
	}
	return &AppGoodBaseQuery{
		config:     agbq.config,
		limit:      agbq.limit,
		offset:     agbq.offset,
		order:      append([]OrderFunc{}, agbq.order...),
		predicates: append([]predicate.AppGoodBase{}, agbq.predicates...),
		// clone intermediate query.
		sql:    agbq.sql.Clone(),
		path:   agbq.path,
		unique: agbq.unique,
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
//	client.AppGoodBase.Query().
//		GroupBy(appgoodbase.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (agbq *AppGoodBaseQuery) GroupBy(field string, fields ...string) *AppGoodBaseGroupBy {
	grbuild := &AppGoodBaseGroupBy{config: agbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := agbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return agbq.sqlQuery(ctx), nil
	}
	grbuild.label = appgoodbase.Label
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
//	client.AppGoodBase.Query().
//		Select(appgoodbase.FieldCreatedAt).
//		Scan(ctx, &v)
func (agbq *AppGoodBaseQuery) Select(fields ...string) *AppGoodBaseSelect {
	agbq.fields = append(agbq.fields, fields...)
	selbuild := &AppGoodBaseSelect{AppGoodBaseQuery: agbq}
	selbuild.label = appgoodbase.Label
	selbuild.flds, selbuild.scan = &agbq.fields, selbuild.Scan
	return selbuild
}

func (agbq *AppGoodBaseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range agbq.fields {
		if !appgoodbase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if agbq.path != nil {
		prev, err := agbq.path(ctx)
		if err != nil {
			return err
		}
		agbq.sql = prev
	}
	if appgoodbase.Policy == nil {
		return errors.New("ent: uninitialized appgoodbase.Policy (forgotten import ent/runtime?)")
	}
	if err := appgoodbase.Policy.EvalQuery(ctx, agbq); err != nil {
		return err
	}
	return nil
}

func (agbq *AppGoodBaseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppGoodBase, error) {
	var (
		nodes = []*AppGoodBase{}
		_spec = agbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppGoodBase).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppGoodBase{config: agbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(agbq.modifiers) > 0 {
		_spec.Modifiers = agbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, agbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (agbq *AppGoodBaseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := agbq.querySpec()
	if len(agbq.modifiers) > 0 {
		_spec.Modifiers = agbq.modifiers
	}
	_spec.Node.Columns = agbq.fields
	if len(agbq.fields) > 0 {
		_spec.Unique = agbq.unique != nil && *agbq.unique
	}
	return sqlgraph.CountNodes(ctx, agbq.driver, _spec)
}

func (agbq *AppGoodBaseQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := agbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (agbq *AppGoodBaseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgoodbase.Table,
			Columns: appgoodbase.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgoodbase.FieldID,
			},
		},
		From:   agbq.sql,
		Unique: true,
	}
	if unique := agbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := agbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appgoodbase.FieldID)
		for i := range fields {
			if fields[i] != appgoodbase.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := agbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := agbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := agbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := agbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (agbq *AppGoodBaseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(agbq.driver.Dialect())
	t1 := builder.Table(appgoodbase.Table)
	columns := agbq.fields
	if len(columns) == 0 {
		columns = appgoodbase.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if agbq.sql != nil {
		selector = agbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if agbq.unique != nil && *agbq.unique {
		selector.Distinct()
	}
	for _, m := range agbq.modifiers {
		m(selector)
	}
	for _, p := range agbq.predicates {
		p(selector)
	}
	for _, p := range agbq.order {
		p(selector)
	}
	if offset := agbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := agbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (agbq *AppGoodBaseQuery) ForUpdate(opts ...sql.LockOption) *AppGoodBaseQuery {
	if agbq.driver.Dialect() == dialect.Postgres {
		agbq.Unique(false)
	}
	agbq.modifiers = append(agbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return agbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (agbq *AppGoodBaseQuery) ForShare(opts ...sql.LockOption) *AppGoodBaseQuery {
	if agbq.driver.Dialect() == dialect.Postgres {
		agbq.Unique(false)
	}
	agbq.modifiers = append(agbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return agbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agbq *AppGoodBaseQuery) Modify(modifiers ...func(s *sql.Selector)) *AppGoodBaseSelect {
	agbq.modifiers = append(agbq.modifiers, modifiers...)
	return agbq.Select()
}

// AppGoodBaseGroupBy is the group-by builder for AppGoodBase entities.
type AppGoodBaseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agbgb *AppGoodBaseGroupBy) Aggregate(fns ...AggregateFunc) *AppGoodBaseGroupBy {
	agbgb.fns = append(agbgb.fns, fns...)
	return agbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (agbgb *AppGoodBaseGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agbgb.path(ctx)
	if err != nil {
		return err
	}
	agbgb.sql = query
	return agbgb.sqlScan(ctx, v)
}

func (agbgb *AppGoodBaseGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agbgb.fields {
		if !appgoodbase.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agbgb *AppGoodBaseGroupBy) sqlQuery() *sql.Selector {
	selector := agbgb.sql.Select()
	aggregation := make([]string, 0, len(agbgb.fns))
	for _, fn := range agbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agbgb.fields)+len(agbgb.fns))
		for _, f := range agbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agbgb.fields...)...)
}

// AppGoodBaseSelect is the builder for selecting fields of AppGoodBase entities.
type AppGoodBaseSelect struct {
	*AppGoodBaseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (agbs *AppGoodBaseSelect) Scan(ctx context.Context, v interface{}) error {
	if err := agbs.prepareQuery(ctx); err != nil {
		return err
	}
	agbs.sql = agbs.AppGoodBaseQuery.sqlQuery(ctx)
	return agbs.sqlScan(ctx, v)
}

func (agbs *AppGoodBaseSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agbs.sql.Query()
	if err := agbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agbs *AppGoodBaseSelect) Modify(modifiers ...func(s *sql.Selector)) *AppGoodBaseSelect {
	agbs.modifiers = append(agbs.modifiers, modifiers...)
	return agbs
}
