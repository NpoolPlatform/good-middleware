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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodlabel"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// AppGoodLabelQuery is the builder for querying AppGoodLabel entities.
type AppGoodLabelQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AppGoodLabel
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AppGoodLabelQuery builder.
func (aglq *AppGoodLabelQuery) Where(ps ...predicate.AppGoodLabel) *AppGoodLabelQuery {
	aglq.predicates = append(aglq.predicates, ps...)
	return aglq
}

// Limit adds a limit step to the query.
func (aglq *AppGoodLabelQuery) Limit(limit int) *AppGoodLabelQuery {
	aglq.limit = &limit
	return aglq
}

// Offset adds an offset step to the query.
func (aglq *AppGoodLabelQuery) Offset(offset int) *AppGoodLabelQuery {
	aglq.offset = &offset
	return aglq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aglq *AppGoodLabelQuery) Unique(unique bool) *AppGoodLabelQuery {
	aglq.unique = &unique
	return aglq
}

// Order adds an order step to the query.
func (aglq *AppGoodLabelQuery) Order(o ...OrderFunc) *AppGoodLabelQuery {
	aglq.order = append(aglq.order, o...)
	return aglq
}

// First returns the first AppGoodLabel entity from the query.
// Returns a *NotFoundError when no AppGoodLabel was found.
func (aglq *AppGoodLabelQuery) First(ctx context.Context) (*AppGoodLabel, error) {
	nodes, err := aglq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{appgoodlabel.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) FirstX(ctx context.Context) *AppGoodLabel {
	node, err := aglq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AppGoodLabel ID from the query.
// Returns a *NotFoundError when no AppGoodLabel ID was found.
func (aglq *AppGoodLabelQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = aglq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{appgoodlabel.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := aglq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AppGoodLabel entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AppGoodLabel entity is found.
// Returns a *NotFoundError when no AppGoodLabel entities are found.
func (aglq *AppGoodLabelQuery) Only(ctx context.Context) (*AppGoodLabel, error) {
	nodes, err := aglq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{appgoodlabel.Label}
	default:
		return nil, &NotSingularError{appgoodlabel.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) OnlyX(ctx context.Context) *AppGoodLabel {
	node, err := aglq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AppGoodLabel ID in the query.
// Returns a *NotSingularError when more than one AppGoodLabel ID is found.
// Returns a *NotFoundError when no entities are found.
func (aglq *AppGoodLabelQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = aglq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{appgoodlabel.Label}
	default:
		err = &NotSingularError{appgoodlabel.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := aglq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AppGoodLabels.
func (aglq *AppGoodLabelQuery) All(ctx context.Context) ([]*AppGoodLabel, error) {
	if err := aglq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return aglq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) AllX(ctx context.Context) []*AppGoodLabel {
	nodes, err := aglq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AppGoodLabel IDs.
func (aglq *AppGoodLabelQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := aglq.Select(appgoodlabel.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := aglq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aglq *AppGoodLabelQuery) Count(ctx context.Context) (int, error) {
	if err := aglq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return aglq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) CountX(ctx context.Context) int {
	count, err := aglq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aglq *AppGoodLabelQuery) Exist(ctx context.Context) (bool, error) {
	if err := aglq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return aglq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (aglq *AppGoodLabelQuery) ExistX(ctx context.Context) bool {
	exist, err := aglq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AppGoodLabelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aglq *AppGoodLabelQuery) Clone() *AppGoodLabelQuery {
	if aglq == nil {
		return nil
	}
	return &AppGoodLabelQuery{
		config:     aglq.config,
		limit:      aglq.limit,
		offset:     aglq.offset,
		order:      append([]OrderFunc{}, aglq.order...),
		predicates: append([]predicate.AppGoodLabel{}, aglq.predicates...),
		// clone intermediate query.
		sql:    aglq.sql.Clone(),
		path:   aglq.path,
		unique: aglq.unique,
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
//	client.AppGoodLabel.Query().
//		GroupBy(appgoodlabel.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (aglq *AppGoodLabelQuery) GroupBy(field string, fields ...string) *AppGoodLabelGroupBy {
	grbuild := &AppGoodLabelGroupBy{config: aglq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := aglq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return aglq.sqlQuery(ctx), nil
	}
	grbuild.label = appgoodlabel.Label
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
//	client.AppGoodLabel.Query().
//		Select(appgoodlabel.FieldCreatedAt).
//		Scan(ctx, &v)
func (aglq *AppGoodLabelQuery) Select(fields ...string) *AppGoodLabelSelect {
	aglq.fields = append(aglq.fields, fields...)
	selbuild := &AppGoodLabelSelect{AppGoodLabelQuery: aglq}
	selbuild.label = appgoodlabel.Label
	selbuild.flds, selbuild.scan = &aglq.fields, selbuild.Scan
	return selbuild
}

func (aglq *AppGoodLabelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range aglq.fields {
		if !appgoodlabel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if aglq.path != nil {
		prev, err := aglq.path(ctx)
		if err != nil {
			return err
		}
		aglq.sql = prev
	}
	if appgoodlabel.Policy == nil {
		return errors.New("ent: uninitialized appgoodlabel.Policy (forgotten import ent/runtime?)")
	}
	if err := appgoodlabel.Policy.EvalQuery(ctx, aglq); err != nil {
		return err
	}
	return nil
}

func (aglq *AppGoodLabelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AppGoodLabel, error) {
	var (
		nodes = []*AppGoodLabel{}
		_spec = aglq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*AppGoodLabel).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &AppGoodLabel{config: aglq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(aglq.modifiers) > 0 {
		_spec.Modifiers = aglq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, aglq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (aglq *AppGoodLabelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := aglq.querySpec()
	if len(aglq.modifiers) > 0 {
		_spec.Modifiers = aglq.modifiers
	}
	_spec.Node.Columns = aglq.fields
	if len(aglq.fields) > 0 {
		_spec.Unique = aglq.unique != nil && *aglq.unique
	}
	return sqlgraph.CountNodes(ctx, aglq.driver, _spec)
}

func (aglq *AppGoodLabelQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := aglq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (aglq *AppGoodLabelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   appgoodlabel.Table,
			Columns: appgoodlabel.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: appgoodlabel.FieldID,
			},
		},
		From:   aglq.sql,
		Unique: true,
	}
	if unique := aglq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := aglq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, appgoodlabel.FieldID)
		for i := range fields {
			if fields[i] != appgoodlabel.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := aglq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := aglq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := aglq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := aglq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (aglq *AppGoodLabelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(aglq.driver.Dialect())
	t1 := builder.Table(appgoodlabel.Table)
	columns := aglq.fields
	if len(columns) == 0 {
		columns = appgoodlabel.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if aglq.sql != nil {
		selector = aglq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if aglq.unique != nil && *aglq.unique {
		selector.Distinct()
	}
	for _, m := range aglq.modifiers {
		m(selector)
	}
	for _, p := range aglq.predicates {
		p(selector)
	}
	for _, p := range aglq.order {
		p(selector)
	}
	if offset := aglq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := aglq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (aglq *AppGoodLabelQuery) ForUpdate(opts ...sql.LockOption) *AppGoodLabelQuery {
	if aglq.driver.Dialect() == dialect.Postgres {
		aglq.Unique(false)
	}
	aglq.modifiers = append(aglq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return aglq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (aglq *AppGoodLabelQuery) ForShare(opts ...sql.LockOption) *AppGoodLabelQuery {
	if aglq.driver.Dialect() == dialect.Postgres {
		aglq.Unique(false)
	}
	aglq.modifiers = append(aglq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return aglq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (aglq *AppGoodLabelQuery) Modify(modifiers ...func(s *sql.Selector)) *AppGoodLabelSelect {
	aglq.modifiers = append(aglq.modifiers, modifiers...)
	return aglq.Select()
}

// AppGoodLabelGroupBy is the group-by builder for AppGoodLabel entities.
type AppGoodLabelGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (aglgb *AppGoodLabelGroupBy) Aggregate(fns ...AggregateFunc) *AppGoodLabelGroupBy {
	aglgb.fns = append(aglgb.fns, fns...)
	return aglgb
}

// Scan applies the group-by query and scans the result into the given value.
func (aglgb *AppGoodLabelGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := aglgb.path(ctx)
	if err != nil {
		return err
	}
	aglgb.sql = query
	return aglgb.sqlScan(ctx, v)
}

func (aglgb *AppGoodLabelGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range aglgb.fields {
		if !appgoodlabel.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := aglgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aglgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (aglgb *AppGoodLabelGroupBy) sqlQuery() *sql.Selector {
	selector := aglgb.sql.Select()
	aggregation := make([]string, 0, len(aglgb.fns))
	for _, fn := range aglgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(aglgb.fields)+len(aglgb.fns))
		for _, f := range aglgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(aglgb.fields...)...)
}

// AppGoodLabelSelect is the builder for selecting fields of AppGoodLabel entities.
type AppGoodLabelSelect struct {
	*AppGoodLabelQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (agls *AppGoodLabelSelect) Scan(ctx context.Context, v interface{}) error {
	if err := agls.prepareQuery(ctx); err != nil {
		return err
	}
	agls.sql = agls.AppGoodLabelQuery.sqlQuery(ctx)
	return agls.sqlScan(ctx, v)
}

func (agls *AppGoodLabelSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := agls.sql.Query()
	if err := agls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (agls *AppGoodLabelSelect) Modify(modifiers ...func(s *sql.Selector)) *AppGoodLabelSelect {
	agls.modifiers = append(agls.modifiers, modifiers...)
	return agls
}
