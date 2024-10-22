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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplaycolor"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppGoodDisplayColorQuery is the builder for querying AppGoodDisplayColor entities.
type AppGoodDisplayColorQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppGoodDisplayColor
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppGoodDisplayColorQuery builder.
func (agdcq *AppGoodDisplayColorQuery) Where(ps ...predicate.AppGoodDisplayColor) *AppGoodDisplayColorQuery {
	agdcq.predicates = append(agdcq.predicates, ps...)
	return agdcq
}

// Limit adds a limit step to the query.
func (agdcq *AppGoodDisplayColorQuery) Limit(limit int) *AppGoodDisplayColorQuery {
	agdcq.limit = &limit
	return agdcq
}

// Offset adds an offset step to the query.
func (agdcq *AppGoodDisplayColorQuery) Offset(offset int) *AppGoodDisplayColorQuery {
	agdcq.offset = &offset
	return agdcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (agdcq *AppGoodDisplayColorQuery) Unique(unique bool) *AppGoodDisplayColorQuery {
	agdcq.unique = &unique
	return agdcq
}

// Order adds an order step to the query.
func (agdcq *AppGoodDisplayColorQuery) Order(o ...OrderFunc) *AppGoodDisplayColorQuery {
	agdcq.order = append(agdcq.order, o...)
	return agdcq
}

// First returns the first AppGoodDisplayColor entity from the query.
// Returns a *NotFoundError when no AppGoodDisplayColor was found.
func (agdcq *AppGoodDisplayColorQuery) First(ctx context.Context) (*AppGoodDisplayColor, error) {
	nodes, err := agdcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appgooddisplaycolor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) FirstX(ctx context.Context) *AppGoodDisplayColor {
	node, err := agdcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppGoodDisplayColor ID from the query.
// Returns a *NotFoundError when no AppGoodDisplayColor ID was found.
func (agdcq *AppGoodDisplayColorQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agdcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appgooddisplaycolor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := agdcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppGoodDisplayColor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppGoodDisplayColor entity is found.
// Returns a *NotFoundError when no AppGoodDisplayColor entities are found.
func (agdcq *AppGoodDisplayColorQuery) Only(ctx context.Context) (*AppGoodDisplayColor, error) {
	nodes, err := agdcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appgooddisplaycolor.Label}
	default:
		return nil, &NotSingularError{appgooddisplaycolor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) OnlyX(ctx context.Context) *AppGoodDisplayColor {
	node, err := agdcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppGoodDisplayColor ID in the query.
// Returns a *NotSingularError when more than one AppGoodDisplayColor ID is found.
// Returns a *NotFoundError when no entities are found.
func (agdcq *AppGoodDisplayColorQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agdcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appgooddisplaycolor.Label}
	default:
		err = &NotSingularError{appgooddisplaycolor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := agdcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppGoodDisplayColors.
func (agdcq *AppGoodDisplayColorQuery) All(ctx context.Context) ([]*AppGoodDisplayColor, error) {
	if err := agdcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return agdcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) AllX(ctx context.Context) []*AppGoodDisplayColor {
	nodes, err := agdcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppGoodDisplayColor IDs.
func (agdcq *AppGoodDisplayColorQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := agdcq.Select(appgooddisplaycolor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := agdcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (agdcq *AppGoodDisplayColorQuery) Count(ctx context.Context) (int, error) {
	if err := agdcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return agdcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) CountX(ctx context.Context) int {
	count, err := agdcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (agdcq *AppGoodDisplayColorQuery) Exist(ctx context.Context) (bool, error) {
	if err := agdcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return agdcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (agdcq *AppGoodDisplayColorQuery) ExistX(ctx context.Context) bool {
	exist, err := agdcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppGoodDisplayColorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (agdcq *AppGoodDisplayColorQuery) Clone() *AppGoodDisplayColorQuery {
	if agdcq == nil {
		return nil
	}
	return &AppGoodDisplayColorQuery{
		config:     agdcq.config,
		limit:      agdcq.limit,
		offset:     agdcq.offset,
		order:      append([]OrderFunc{}, agdcq.order...),
		predicates: append([]predicate.AppGoodDisplayColor{}, agdcq.predicates...),
		// clone intermediate query.
		sql:    agdcq.sql.Clone(),
		path:   agdcq.path,
		unique: agdcq.unique,
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
//	client.AppGoodDisplayColor.Query().
//		GroupBy(appgooddisplaycolor.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (agdcq *AppGoodDisplayColorQuery) GroupBy(field string, fields ...string) *AppGoodDisplayColorGroupBy {
	grbuild := &AppGoodDisplayColorGroupBy{config: agdcq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := agdcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return agdcq.sqlQuery(ctx), nil
	}
	grbuild.label = appgooddisplaycolor.Label
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
//	client.AppGoodDisplayColor.Query().
//		Select(appgooddisplaycolor.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (agdcq *AppGoodDisplayColorQuery) Select(fields ...string) *AppGoodDisplayColorSelect {
	agdcq.fields = append(agdcq.fields, fields...)
	selbuild := &AppGoodDisplayColorSelect{AppGoodDisplayColorQuery: agdcq}
	selbuild.label = appgooddisplaycolor.Label
	selbuild.flds, selbuild.scan = &agdcq.fields, selbuild.Scan
	return selbuild
}

func (agdcq *AppGoodDisplayColorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range agdcq.fields {
		if !appgooddisplaycolor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if agdcq.path != nil {
		prev, err := agdcq.path(ctx)
		if err != nil {
			return err
		}
		agdcq.sql = prev
	}
	if appgooddisplaycolor.Policy == nil {
		return errors.New("ent: uninitialized appgooddisplaycolor.Policy (forgotten import ent/runtime?)")
	}
	if err := appgooddisplaycolor.Policy.EvalQuery(ctx, agdcq); err != nil {
		return err
	}
	return nil
}

func (agdcq *AppGoodDisplayColorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppGoodDisplayColor, error) {
	var (
		nodes = []*AppGoodDisplayColor{}
		_spec = agdcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppGoodDisplayColor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppGoodDisplayColor{config: agdcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(agdcq.modifiers) > 0 {
		_spec.Modifiers = agdcq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, agdcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (agdcq *AppGoodDisplayColorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := agdcq.querySpec()
	if len(agdcq.modifiers) > 0 {
		_spec.Modifiers = agdcq.modifiers
	}
	_spec.Node.Columns = agdcq.fields
	if len(agdcq.fields) > 0 {
		_spec.Unique = agdcq.unique != nil && *agdcq.unique
	}
	return sqlgraph.CountNodes(ctx, agdcq.driver, _spec)
}

func (agdcq *AppGoodDisplayColorQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := agdcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (agdcq *AppGoodDisplayColorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgooddisplaycolor.Table,
			Columns: appgooddisplaycolor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgooddisplaycolor.FieldID,
			},
		},
		From:   agdcq.sql,
		Unique: true,
	}
	if unique := agdcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := agdcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appgooddisplaycolor.FieldID)
		for i := range fields {
			if fields[i] != appgooddisplaycolor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := agdcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := agdcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := agdcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := agdcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (agdcq *AppGoodDisplayColorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(agdcq.driver.Dialect())
	t1 := builder.Table(appgooddisplaycolor.Table)
	columns := agdcq.fields
	if len(columns) == 0 {
		columns = appgooddisplaycolor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if agdcq.sql != nil {
		selector = agdcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if agdcq.unique != nil && *agdcq.unique {
		selector.Distinct()
	}
	for _, m := range agdcq.modifiers {
		m(selector)
	}
	for _, p := range agdcq.predicates {
		p(selector)
	}
	for _, p := range agdcq.order {
		p(selector)
	}
	if offset := agdcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := agdcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (agdcq *AppGoodDisplayColorQuery) ForUpdate(opts ...sql.LockOption) *AppGoodDisplayColorQuery {
	if agdcq.driver.Dialect() == dialect.Postgres {
		agdcq.Unique(false)
	}
	agdcq.modifiers = append(agdcq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return agdcq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (agdcq *AppGoodDisplayColorQuery) ForShare(opts ...sql.LockOption) *AppGoodDisplayColorQuery {
	if agdcq.driver.Dialect() == dialect.Postgres {
		agdcq.Unique(false)
	}
	agdcq.modifiers = append(agdcq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return agdcq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agdcq *AppGoodDisplayColorQuery) Modify(modifiers ...func(s *sql.Selector)) *AppGoodDisplayColorSelect {
	agdcq.modifiers = append(agdcq.modifiers, modifiers...)
	return agdcq.Select()
}

// AppGoodDisplayColorGroupBy is the group-by builder for AppGoodDisplayColor entities.
type AppGoodDisplayColorGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agdcgb *AppGoodDisplayColorGroupBy) Aggregate(fns ...AggregateFunc) *AppGoodDisplayColorGroupBy {
	agdcgb.fns = append(agdcgb.fns, fns...)
	return agdcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (agdcgb *AppGoodDisplayColorGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agdcgb.path(ctx)
	if err != nil {
		return err
	}
	agdcgb.sql = query
	return agdcgb.sqlScan(ctx, v)
}

func (agdcgb *AppGoodDisplayColorGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agdcgb.fields {
		if !appgooddisplaycolor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agdcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agdcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agdcgb *AppGoodDisplayColorGroupBy) sqlQuery() *sql.Selector {
	selector := agdcgb.sql.Select()
	aggregation := make([]string, 0, len(agdcgb.fns))
	for _, fn := range agdcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agdcgb.fields)+len(agdcgb.fns))
		for _, f := range agdcgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agdcgb.fields...)...)
}

// AppGoodDisplayColorSelect is the builder for selecting fields of AppGoodDisplayColor entities.
type AppGoodDisplayColorSelect struct {
	*AppGoodDisplayColorQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (agdcs *AppGoodDisplayColorSelect) Scan(ctx context.Context, v interface{}) error {
	if err := agdcs.prepareQuery(ctx); err != nil {
		return err
	}
	agdcs.sql = agdcs.AppGoodDisplayColorQuery.sqlQuery(ctx)
	return agdcs.sqlScan(ctx, v)
}

func (agdcs *AppGoodDisplayColorSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agdcs.sql.Query()
	if err := agdcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agdcs *AppGoodDisplayColorSelect) Modify(modifiers ...func(s *sql.Selector)) *AppGoodDisplayColorSelect {
	agdcs.modifiers = append(agdcs.modifiers, modifiers...)
	return agdcs
}
