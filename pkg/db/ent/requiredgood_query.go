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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredgood"
	"github.com/google/uuid"
)

// RequiredGoodQuery is the builder for querying RequiredGood entities.
type RequiredGoodQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.RequiredGood
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RequiredGoodQuery builder.
func (rgq *RequiredGoodQuery) Where(ps ...predicate.RequiredGood) *RequiredGoodQuery {
	rgq.predicates = append(rgq.predicates, ps...)
	return rgq
}

// Limit adds a limit step to the query.
func (rgq *RequiredGoodQuery) Limit(limit int) *RequiredGoodQuery {
	rgq.limit = &limit
	return rgq
}

// Offset adds an offset step to the query.
func (rgq *RequiredGoodQuery) Offset(offset int) *RequiredGoodQuery {
	rgq.offset = &offset
	return rgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rgq *RequiredGoodQuery) Unique(unique bool) *RequiredGoodQuery {
	rgq.unique = &unique
	return rgq
}

// Order adds an order step to the query.
func (rgq *RequiredGoodQuery) Order(o ...OrderFunc) *RequiredGoodQuery {
	rgq.order = append(rgq.order, o...)
	return rgq
}

// First returns the first RequiredGood entity from the query.
// Returns a *NotFoundError when no RequiredGood was found.
func (rgq *RequiredGoodQuery) First(ctx context.Context) (*RequiredGood, error) {
	nodes, err := rgq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{requiredgood.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rgq *RequiredGoodQuery) FirstX(ctx context.Context) *RequiredGood {
	node, err := rgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RequiredGood ID from the query.
// Returns a *NotFoundError when no RequiredGood ID was found.
func (rgq *RequiredGoodQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rgq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{requiredgood.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rgq *RequiredGoodQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := rgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RequiredGood entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RequiredGood entity is found.
// Returns a *NotFoundError when no RequiredGood entities are found.
func (rgq *RequiredGoodQuery) Only(ctx context.Context) (*RequiredGood, error) {
	nodes, err := rgq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{requiredgood.Label}
	default:
		return nil, &NotSingularError{requiredgood.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rgq *RequiredGoodQuery) OnlyX(ctx context.Context) *RequiredGood {
	node, err := rgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RequiredGood ID in the query.
// Returns a *NotSingularError when more than one RequiredGood ID is found.
// Returns a *NotFoundError when no entities are found.
func (rgq *RequiredGoodQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = rgq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{requiredgood.Label}
	default:
		err = &NotSingularError{requiredgood.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rgq *RequiredGoodQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := rgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RequiredGoods.
func (rgq *RequiredGoodQuery) All(ctx context.Context) ([]*RequiredGood, error) {
	if err := rgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return rgq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (rgq *RequiredGoodQuery) AllX(ctx context.Context) []*RequiredGood {
	nodes, err := rgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RequiredGood IDs.
func (rgq *RequiredGoodQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := rgq.Select(requiredgood.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rgq *RequiredGoodQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := rgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rgq *RequiredGoodQuery) Count(ctx context.Context) (int, error) {
	if err := rgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return rgq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (rgq *RequiredGoodQuery) CountX(ctx context.Context) int {
	count, err := rgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rgq *RequiredGoodQuery) Exist(ctx context.Context) (bool, error) {
	if err := rgq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return rgq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (rgq *RequiredGoodQuery) ExistX(ctx context.Context) bool {
	exist, err := rgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RequiredGoodQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rgq *RequiredGoodQuery) Clone() *RequiredGoodQuery {
	if rgq == nil {
		return nil
	}
	return &RequiredGoodQuery{
		config:     rgq.config,
		limit:      rgq.limit,
		offset:     rgq.offset,
		order:      append([]OrderFunc{}, rgq.order...),
		predicates: append([]predicate.RequiredGood{}, rgq.predicates...),
		// clone intermediate query.
		sql:    rgq.sql.Clone(),
		path:   rgq.path,
		unique: rgq.unique,
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
//	client.RequiredGood.Query().
//		GroupBy(requiredgood.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (rgq *RequiredGoodQuery) GroupBy(field string, fields ...string) *RequiredGoodGroupBy {
	grbuild := &RequiredGoodGroupBy{config: rgq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := rgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return rgq.sqlQuery(ctx), nil
	}
	grbuild.label = requiredgood.Label
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
//	client.RequiredGood.Query().
//		Select(requiredgood.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (rgq *RequiredGoodQuery) Select(fields ...string) *RequiredGoodSelect {
	rgq.fields = append(rgq.fields, fields...)
	selbuild := &RequiredGoodSelect{RequiredGoodQuery: rgq}
	selbuild.label = requiredgood.Label
	selbuild.flds, selbuild.scan = &rgq.fields, selbuild.Scan
	return selbuild
}

func (rgq *RequiredGoodQuery) prepareQuery(ctx context.Context) error {
	for _, f := range rgq.fields {
		if !requiredgood.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rgq.path != nil {
		prev, err := rgq.path(ctx)
		if err != nil {
			return err
		}
		rgq.sql = prev
	}
	if requiredgood.Policy == nil {
		return errors.New("ent: uninitialized requiredgood.Policy (forgotten import ent/runtime?)")
	}
	if err := requiredgood.Policy.EvalQuery(ctx, rgq); err != nil {
		return err
	}
	return nil
}

func (rgq *RequiredGoodQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RequiredGood, error) {
	var (
		nodes = []*RequiredGood{}
		_spec = rgq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*RequiredGood).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &RequiredGood{config: rgq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(rgq.modifiers) > 0 {
		_spec.Modifiers = rgq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (rgq *RequiredGoodQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rgq.querySpec()
	if len(rgq.modifiers) > 0 {
		_spec.Modifiers = rgq.modifiers
	}
	_spec.Node.Columns = rgq.fields
	if len(rgq.fields) > 0 {
		_spec.Unique = rgq.unique != nil && *rgq.unique
	}
	return sqlgraph.CountNodes(ctx, rgq.driver, _spec)
}

func (rgq *RequiredGoodQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := rgq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (rgq *RequiredGoodQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   requiredgood.Table,
			Columns: requiredgood.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: requiredgood.FieldID,
			},
		},
		From:   rgq.sql,
		Unique: true,
	}
	if unique := rgq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := rgq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, requiredgood.FieldID)
		for i := range fields {
			if fields[i] != requiredgood.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rgq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rgq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rgq *RequiredGoodQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rgq.driver.Dialect())
	t1 := builder.Table(requiredgood.Table)
	columns := rgq.fields
	if len(columns) == 0 {
		columns = requiredgood.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rgq.sql != nil {
		selector = rgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rgq.unique != nil && *rgq.unique {
		selector.Distinct()
	}
	for _, m := range rgq.modifiers {
		m(selector)
	}
	for _, p := range rgq.predicates {
		p(selector)
	}
	for _, p := range rgq.order {
		p(selector)
	}
	if offset := rgq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rgq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (rgq *RequiredGoodQuery) ForUpdate(opts ...sql.LockOption) *RequiredGoodQuery {
	if rgq.driver.Dialect() == dialect.Postgres {
		rgq.Unique(false)
	}
	rgq.modifiers = append(rgq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return rgq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (rgq *RequiredGoodQuery) ForShare(opts ...sql.LockOption) *RequiredGoodQuery {
	if rgq.driver.Dialect() == dialect.Postgres {
		rgq.Unique(false)
	}
	rgq.modifiers = append(rgq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return rgq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rgq *RequiredGoodQuery) Modify(modifiers ...func(s *sql.Selector)) *RequiredGoodSelect {
	rgq.modifiers = append(rgq.modifiers, modifiers...)
	return rgq.Select()
}

// RequiredGoodGroupBy is the group-by builder for RequiredGood entities.
type RequiredGoodGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rggb *RequiredGoodGroupBy) Aggregate(fns ...AggregateFunc) *RequiredGoodGroupBy {
	rggb.fns = append(rggb.fns, fns...)
	return rggb
}

// Scan applies the group-by query and scans the result into the given value.
func (rggb *RequiredGoodGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := rggb.path(ctx)
	if err != nil {
		return err
	}
	rggb.sql = query
	return rggb.sqlScan(ctx, v)
}

func (rggb *RequiredGoodGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range rggb.fields {
		if !requiredgood.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := rggb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rggb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (rggb *RequiredGoodGroupBy) sqlQuery() *sql.Selector {
	selector := rggb.sql.Select()
	aggregation := make([]string, 0, len(rggb.fns))
	for _, fn := range rggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(rggb.fields)+len(rggb.fns))
		for _, f := range rggb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(rggb.fields...)...)
}

// RequiredGoodSelect is the builder for selecting fields of RequiredGood entities.
type RequiredGoodSelect struct {
	*RequiredGoodQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (rgs *RequiredGoodSelect) Scan(ctx context.Context, v interface{}) error {
	if err := rgs.prepareQuery(ctx); err != nil {
		return err
	}
	rgs.sql = rgs.RequiredGoodQuery.sqlQuery(ctx)
	return rgs.sqlScan(ctx, v)
}

func (rgs *RequiredGoodSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := rgs.sql.Query()
	if err := rgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (rgs *RequiredGoodSelect) Modify(modifiers ...func(s *sql.Selector)) *RequiredGoodSelect {
	rgs.modifiers = append(rgs.modifiers, modifiers...)
	return rgs
}
