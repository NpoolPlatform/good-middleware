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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddescription"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppGoodDescriptionQuery is the builder for querying AppGoodDescription entities.
type AppGoodDescriptionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppGoodDescription
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppGoodDescriptionQuery builder.
func (agdq *AppGoodDescriptionQuery) Where(ps ...predicate.AppGoodDescription) *AppGoodDescriptionQuery {
	agdq.predicates = append(agdq.predicates, ps...)
	return agdq
}

// Limit adds a limit step to the query.
func (agdq *AppGoodDescriptionQuery) Limit(limit int) *AppGoodDescriptionQuery {
	agdq.limit = &limit
	return agdq
}

// Offset adds an offset step to the query.
func (agdq *AppGoodDescriptionQuery) Offset(offset int) *AppGoodDescriptionQuery {
	agdq.offset = &offset
	return agdq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (agdq *AppGoodDescriptionQuery) Unique(unique bool) *AppGoodDescriptionQuery {
	agdq.unique = &unique
	return agdq
}

// Order adds an order step to the query.
func (agdq *AppGoodDescriptionQuery) Order(o ...OrderFunc) *AppGoodDescriptionQuery {
	agdq.order = append(agdq.order, o...)
	return agdq
}

// First returns the first AppGoodDescription entity from the query.
// Returns a *NotFoundError when no AppGoodDescription was found.
func (agdq *AppGoodDescriptionQuery) First(ctx context.Context) (*AppGoodDescription, error) {
	nodes, err := agdq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appgooddescription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) FirstX(ctx context.Context) *AppGoodDescription {
	node, err := agdq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppGoodDescription ID from the query.
// Returns a *NotFoundError when no AppGoodDescription ID was found.
func (agdq *AppGoodDescriptionQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agdq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appgooddescription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := agdq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppGoodDescription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppGoodDescription entity is found.
// Returns a *NotFoundError when no AppGoodDescription entities are found.
func (agdq *AppGoodDescriptionQuery) Only(ctx context.Context) (*AppGoodDescription, error) {
	nodes, err := agdq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appgooddescription.Label}
	default:
		return nil, &NotSingularError{appgooddescription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) OnlyX(ctx context.Context) *AppGoodDescription {
	node, err := agdq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppGoodDescription ID in the query.
// Returns a *NotSingularError when more than one AppGoodDescription ID is found.
// Returns a *NotFoundError when no entities are found.
func (agdq *AppGoodDescriptionQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = agdq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appgooddescription.Label}
	default:
		err = &NotSingularError{appgooddescription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := agdq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppGoodDescriptions.
func (agdq *AppGoodDescriptionQuery) All(ctx context.Context) ([]*AppGoodDescription, error) {
	if err := agdq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return agdq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) AllX(ctx context.Context) []*AppGoodDescription {
	nodes, err := agdq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppGoodDescription IDs.
func (agdq *AppGoodDescriptionQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := agdq.Select(appgooddescription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := agdq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (agdq *AppGoodDescriptionQuery) Count(ctx context.Context) (int, error) {
	if err := agdq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return agdq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) CountX(ctx context.Context) int {
	count, err := agdq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (agdq *AppGoodDescriptionQuery) Exist(ctx context.Context) (bool, error) {
	if err := agdq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return agdq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (agdq *AppGoodDescriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := agdq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppGoodDescriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (agdq *AppGoodDescriptionQuery) Clone() *AppGoodDescriptionQuery {
	if agdq == nil {
		return nil
	}
	return &AppGoodDescriptionQuery{
		config:     agdq.config,
		limit:      agdq.limit,
		offset:     agdq.offset,
		order:      append([]OrderFunc{}, agdq.order...),
		predicates: append([]predicate.AppGoodDescription{}, agdq.predicates...),
		// clone intermediate query.
		sql:    agdq.sql.Clone(),
		path:   agdq.path,
		unique: agdq.unique,
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
//	client.AppGoodDescription.Query().
//		GroupBy(appgooddescription.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (agdq *AppGoodDescriptionQuery) GroupBy(field string, fields ...string) *AppGoodDescriptionGroupBy {
	grbuild := &AppGoodDescriptionGroupBy{config: agdq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := agdq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return agdq.sqlQuery(ctx), nil
	}
	grbuild.label = appgooddescription.Label
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
//	client.AppGoodDescription.Query().
//		Select(appgooddescription.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (agdq *AppGoodDescriptionQuery) Select(fields ...string) *AppGoodDescriptionSelect {
	agdq.fields = append(agdq.fields, fields...)
	selbuild := &AppGoodDescriptionSelect{AppGoodDescriptionQuery: agdq}
	selbuild.label = appgooddescription.Label
	selbuild.flds, selbuild.scan = &agdq.fields, selbuild.Scan
	return selbuild
}

func (agdq *AppGoodDescriptionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range agdq.fields {
		if !appgooddescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if agdq.path != nil {
		prev, err := agdq.path(ctx)
		if err != nil {
			return err
		}
		agdq.sql = prev
	}
	if appgooddescription.Policy == nil {
		return errors.New("ent: uninitialized appgooddescription.Policy (forgotten import ent/runtime?)")
	}
	if err := appgooddescription.Policy.EvalQuery(ctx, agdq); err != nil {
		return err
	}
	return nil
}

func (agdq *AppGoodDescriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppGoodDescription, error) {
	var (
		nodes = []*AppGoodDescription{}
		_spec = agdq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppGoodDescription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppGoodDescription{config: agdq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(agdq.modifiers) > 0 {
		_spec.Modifiers = agdq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, agdq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (agdq *AppGoodDescriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := agdq.querySpec()
	if len(agdq.modifiers) > 0 {
		_spec.Modifiers = agdq.modifiers
	}
	_spec.Node.Columns = agdq.fields
	if len(agdq.fields) > 0 {
		_spec.Unique = agdq.unique != nil && *agdq.unique
	}
	return sqlgraph.CountNodes(ctx, agdq.driver, _spec)
}

func (agdq *AppGoodDescriptionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := agdq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (agdq *AppGoodDescriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgooddescription.Table,
			Columns: appgooddescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgooddescription.FieldID,
			},
		},
		From:   agdq.sql,
		Unique: true,
	}
	if unique := agdq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := agdq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appgooddescription.FieldID)
		for i := range fields {
			if fields[i] != appgooddescription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := agdq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := agdq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := agdq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := agdq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (agdq *AppGoodDescriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(agdq.driver.Dialect())
	t1 := builder.Table(appgooddescription.Table)
	columns := agdq.fields
	if len(columns) == 0 {
		columns = appgooddescription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if agdq.sql != nil {
		selector = agdq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if agdq.unique != nil && *agdq.unique {
		selector.Distinct()
	}
	for _, m := range agdq.modifiers {
		m(selector)
	}
	for _, p := range agdq.predicates {
		p(selector)
	}
	for _, p := range agdq.order {
		p(selector)
	}
	if offset := agdq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := agdq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (agdq *AppGoodDescriptionQuery) ForUpdate(opts ...sql.LockOption) *AppGoodDescriptionQuery {
	if agdq.driver.Dialect() == dialect.Postgres {
		agdq.Unique(false)
	}
	agdq.modifiers = append(agdq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return agdq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (agdq *AppGoodDescriptionQuery) ForShare(opts ...sql.LockOption) *AppGoodDescriptionQuery {
	if agdq.driver.Dialect() == dialect.Postgres {
		agdq.Unique(false)
	}
	agdq.modifiers = append(agdq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return agdq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agdq *AppGoodDescriptionQuery) Modify(modifiers ...func(s *sql.Selector)) *AppGoodDescriptionSelect {
	agdq.modifiers = append(agdq.modifiers, modifiers...)
	return agdq.Select()
}

// AppGoodDescriptionGroupBy is the group-by builder for AppGoodDescription entities.
type AppGoodDescriptionGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agdgb *AppGoodDescriptionGroupBy) Aggregate(fns ...AggregateFunc) *AppGoodDescriptionGroupBy {
	agdgb.fns = append(agdgb.fns, fns...)
	return agdgb
}

// Scan applies the group-by query and scans the result into the given value.
func (agdgb *AppGoodDescriptionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := agdgb.path(ctx)
	if err != nil {
		return err
	}
	agdgb.sql = query
	return agdgb.sqlScan(ctx, v)
}

func (agdgb *AppGoodDescriptionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range agdgb.fields {
		if !appgooddescription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := agdgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := agdgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (agdgb *AppGoodDescriptionGroupBy) sqlQuery() *sql.Selector {
	selector := agdgb.sql.Select()
	aggregation := make([]string, 0, len(agdgb.fns))
	for _, fn := range agdgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(agdgb.fields)+len(agdgb.fns))
		for _, f := range agdgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(agdgb.fields...)...)
}

// AppGoodDescriptionSelect is the builder for selecting fields of AppGoodDescription entities.
type AppGoodDescriptionSelect struct {
	*AppGoodDescriptionQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (agds *AppGoodDescriptionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := agds.prepareQuery(ctx); err != nil {
		return err
	}
	agds.sql = agds.AppGoodDescriptionQuery.sqlQuery(ctx)
	return agds.sqlScan(ctx, v)
}

func (agds *AppGoodDescriptionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agds.sql.Query()
	if err := agds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agds *AppGoodDescriptionSelect) Modify(modifiers ...func(s *sql.Selector)) *AppGoodDescriptionSelect {
	agds.modifiers = append(agds.modifiers, modifiers...)
	return agds
}
