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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// MiningGoodStockQuery is the builder for querying MiningGoodStock entities.
type MiningGoodStockQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.MiningGoodStock
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MiningGoodStockQuery builder.
func (mgsq *MiningGoodStockQuery) Where(ps ...predicate.MiningGoodStock) *MiningGoodStockQuery {
	mgsq.predicates = append(mgsq.predicates, ps...)
	return mgsq
}

// Limit adds a limit step to the query.
func (mgsq *MiningGoodStockQuery) Limit(limit int) *MiningGoodStockQuery {
	mgsq.limit = &limit
	return mgsq
}

// Offset adds an offset step to the query.
func (mgsq *MiningGoodStockQuery) Offset(offset int) *MiningGoodStockQuery {
	mgsq.offset = &offset
	return mgsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mgsq *MiningGoodStockQuery) Unique(unique bool) *MiningGoodStockQuery {
	mgsq.unique = &unique
	return mgsq
}

// Order adds an order step to the query.
func (mgsq *MiningGoodStockQuery) Order(o ...OrderFunc) *MiningGoodStockQuery {
	mgsq.order = append(mgsq.order, o...)
	return mgsq
}

// First returns the first MiningGoodStock entity from the query.
// Returns a *NotFoundError when no MiningGoodStock was found.
func (mgsq *MiningGoodStockQuery) First(ctx context.Context) (*MiningGoodStock, error) {
	nodes, err := mgsq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{mininggoodstock.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) FirstX(ctx context.Context) *MiningGoodStock {
	node, err := mgsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MiningGoodStock ID from the query.
// Returns a *NotFoundError when no MiningGoodStock ID was found.
func (mgsq *MiningGoodStockQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = mgsq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{mininggoodstock.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := mgsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MiningGoodStock entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MiningGoodStock entity is found.
// Returns a *NotFoundError when no MiningGoodStock entities are found.
func (mgsq *MiningGoodStockQuery) Only(ctx context.Context) (*MiningGoodStock, error) {
	nodes, err := mgsq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{mininggoodstock.Label}
	default:
		return nil, &NotSingularError{mininggoodstock.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) OnlyX(ctx context.Context) *MiningGoodStock {
	node, err := mgsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MiningGoodStock ID in the query.
// Returns a *NotSingularError when more than one MiningGoodStock ID is found.
// Returns a *NotFoundError when no entities are found.
func (mgsq *MiningGoodStockQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = mgsq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{mininggoodstock.Label}
	default:
		err = &NotSingularError{mininggoodstock.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := mgsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MiningGoodStocks.
func (mgsq *MiningGoodStockQuery) All(ctx context.Context) ([]*MiningGoodStock, error) {
	if err := mgsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mgsq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) AllX(ctx context.Context) []*MiningGoodStock {
	nodes, err := mgsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MiningGoodStock IDs.
func (mgsq *MiningGoodStockQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := mgsq.Select(mininggoodstock.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := mgsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mgsq *MiningGoodStockQuery) Count(ctx context.Context) (int, error) {
	if err := mgsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mgsq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) CountX(ctx context.Context) int {
	count, err := mgsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mgsq *MiningGoodStockQuery) Exist(ctx context.Context) (bool, error) {
	if err := mgsq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mgsq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mgsq *MiningGoodStockQuery) ExistX(ctx context.Context) bool {
	exist, err := mgsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MiningGoodStockQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mgsq *MiningGoodStockQuery) Clone() *MiningGoodStockQuery {
	if mgsq == nil {
		return nil
	}
	return &MiningGoodStockQuery{
		config:     mgsq.config,
		limit:      mgsq.limit,
		offset:     mgsq.offset,
		order:      append([]OrderFunc{}, mgsq.order...),
		predicates: append([]predicate.MiningGoodStock{}, mgsq.predicates...),
		// clone intermediate query.
		sql:    mgsq.sql.Clone(),
		path:   mgsq.path,
		unique: mgsq.unique,
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
//	client.MiningGoodStock.Query().
//		GroupBy(mininggoodstock.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (mgsq *MiningGoodStockQuery) GroupBy(field string, fields ...string) *MiningGoodStockGroupBy {
	grbuild := &MiningGoodStockGroupBy{config: mgsq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mgsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mgsq.sqlQuery(ctx), nil
	}
	grbuild.label = mininggoodstock.Label
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
//	client.MiningGoodStock.Query().
//		Select(mininggoodstock.FieldCreatedAt).
//		Scan(ctx, &v)
func (mgsq *MiningGoodStockQuery) Select(fields ...string) *MiningGoodStockSelect {
	mgsq.fields = append(mgsq.fields, fields...)
	selbuild := &MiningGoodStockSelect{MiningGoodStockQuery: mgsq}
	selbuild.label = mininggoodstock.Label
	selbuild.flds, selbuild.scan = &mgsq.fields, selbuild.Scan
	return selbuild
}

func (mgsq *MiningGoodStockQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mgsq.fields {
		if !mininggoodstock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if mgsq.path != nil {
		prev, err := mgsq.path(ctx)
		if err != nil {
			return err
		}
		mgsq.sql = prev
	}
	if mininggoodstock.Policy == nil {
		return errors.New("ent: uninitialized mininggoodstock.Policy (forgotten import ent/runtime?)")
	}
	if err := mininggoodstock.Policy.EvalQuery(ctx, mgsq); err != nil {
		return err
	}
	return nil
}

func (mgsq *MiningGoodStockQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MiningGoodStock, error) {
	var (
		nodes = []*MiningGoodStock{}
		_spec = mgsq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*MiningGoodStock).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &MiningGoodStock{config: mgsq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(mgsq.modifiers) > 0 {
		_spec.Modifiers = mgsq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mgsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (mgsq *MiningGoodStockQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mgsq.querySpec()
	if len(mgsq.modifiers) > 0 {
		_spec.Modifiers = mgsq.modifiers
	}
	_spec.Node.Columns = mgsq.fields
	if len(mgsq.fields) > 0 {
		_spec.Unique = mgsq.unique != nil && *mgsq.unique
	}
	return sqlgraph.CountNodes(ctx, mgsq.driver, _spec)
}

func (mgsq *MiningGoodStockQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := mgsq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (mgsq *MiningGoodStockQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mininggoodstock.Table,
			Columns: mininggoodstock.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: mininggoodstock.FieldID,
			},
		},
		From:   mgsq.sql,
		Unique: true,
	}
	if unique := mgsq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mgsq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mininggoodstock.FieldID)
		for i := range fields {
			if fields[i] != mininggoodstock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mgsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mgsq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mgsq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mgsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mgsq *MiningGoodStockQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mgsq.driver.Dialect())
	t1 := builder.Table(mininggoodstock.Table)
	columns := mgsq.fields
	if len(columns) == 0 {
		columns = mininggoodstock.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mgsq.sql != nil {
		selector = mgsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mgsq.unique != nil && *mgsq.unique {
		selector.Distinct()
	}
	for _, m := range mgsq.modifiers {
		m(selector)
	}
	for _, p := range mgsq.predicates {
		p(selector)
	}
	for _, p := range mgsq.order {
		p(selector)
	}
	if offset := mgsq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mgsq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (mgsq *MiningGoodStockQuery) ForUpdate(opts ...sql.LockOption) *MiningGoodStockQuery {
	if mgsq.driver.Dialect() == dialect.Postgres {
		mgsq.Unique(false)
	}
	mgsq.modifiers = append(mgsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return mgsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (mgsq *MiningGoodStockQuery) ForShare(opts ...sql.LockOption) *MiningGoodStockQuery {
	if mgsq.driver.Dialect() == dialect.Postgres {
		mgsq.Unique(false)
	}
	mgsq.modifiers = append(mgsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return mgsq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mgsq *MiningGoodStockQuery) Modify(modifiers ...func(s *sql.Selector)) *MiningGoodStockSelect {
	mgsq.modifiers = append(mgsq.modifiers, modifiers...)
	return mgsq.Select()
}

// MiningGoodStockGroupBy is the group-by builder for MiningGoodStock entities.
type MiningGoodStockGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgsgb *MiningGoodStockGroupBy) Aggregate(fns ...AggregateFunc) *MiningGoodStockGroupBy {
	mgsgb.fns = append(mgsgb.fns, fns...)
	return mgsgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgsgb *MiningGoodStockGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := mgsgb.path(ctx)
	if err != nil {
		return err
	}
	mgsgb.sql = query
	return mgsgb.sqlScan(ctx, v)
}

func (mgsgb *MiningGoodStockGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range mgsgb.fields {
		if !mininggoodstock.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgsgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgsgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgsgb *MiningGoodStockGroupBy) sqlQuery() *sql.Selector {
	selector := mgsgb.sql.Select()
	aggregation := make([]string, 0, len(mgsgb.fns))
	for _, fn := range mgsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgsgb.fields)+len(mgsgb.fns))
		for _, f := range mgsgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgsgb.fields...)...)
}

// MiningGoodStockSelect is the builder for selecting fields of MiningGoodStock entities.
type MiningGoodStockSelect struct {
	*MiningGoodStockQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (mgss *MiningGoodStockSelect) Scan(ctx context.Context, v interface{}) error {
	if err := mgss.prepareQuery(ctx); err != nil {
		return err
	}
	mgss.sql = mgss.MiningGoodStockQuery.sqlQuery(ctx)
	return mgss.sqlScan(ctx, v)
}

func (mgss *MiningGoodStockSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := mgss.sql.Query()
	if err := mgss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (mgss *MiningGoodStockSelect) Modify(modifiers ...func(s *sql.Selector)) *MiningGoodStockSelect {
	mgss.modifiers = append(mgss.modifiers, modifiers...)
	return mgss
}
