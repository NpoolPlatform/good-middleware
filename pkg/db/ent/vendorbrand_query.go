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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
)

// VendorBrandQuery is the builder for querying VendorBrand entities.
type VendorBrandQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.VendorBrand
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VendorBrandQuery builder.
func (vbq *VendorBrandQuery) Where(ps ...predicate.VendorBrand) *VendorBrandQuery {
	vbq.predicates = append(vbq.predicates, ps...)
	return vbq
}

// Limit adds a limit step to the query.
func (vbq *VendorBrandQuery) Limit(limit int) *VendorBrandQuery {
	vbq.limit = &limit
	return vbq
}

// Offset adds an offset step to the query.
func (vbq *VendorBrandQuery) Offset(offset int) *VendorBrandQuery {
	vbq.offset = &offset
	return vbq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vbq *VendorBrandQuery) Unique(unique bool) *VendorBrandQuery {
	vbq.unique = &unique
	return vbq
}

// Order adds an order step to the query.
func (vbq *VendorBrandQuery) Order(o ...OrderFunc) *VendorBrandQuery {
	vbq.order = append(vbq.order, o...)
	return vbq
}

// First returns the first VendorBrand entity from the query.
// Returns a *NotFoundError when no VendorBrand was found.
func (vbq *VendorBrandQuery) First(ctx context.Context) (*VendorBrand, error) {
	nodes, err := vbq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{vendorbrand.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vbq *VendorBrandQuery) FirstX(ctx context.Context) *VendorBrand {
	node, err := vbq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first VendorBrand ID from the query.
// Returns a *NotFoundError when no VendorBrand ID was found.
func (vbq *VendorBrandQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = vbq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{vendorbrand.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vbq *VendorBrandQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := vbq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single VendorBrand entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one VendorBrand entity is found.
// Returns a *NotFoundError when no VendorBrand entities are found.
func (vbq *VendorBrandQuery) Only(ctx context.Context) (*VendorBrand, error) {
	nodes, err := vbq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{vendorbrand.Label}
	default:
		return nil, &NotSingularError{vendorbrand.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vbq *VendorBrandQuery) OnlyX(ctx context.Context) *VendorBrand {
	node, err := vbq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only VendorBrand ID in the query.
// Returns a *NotSingularError when more than one VendorBrand ID is found.
// Returns a *NotFoundError when no entities are found.
func (vbq *VendorBrandQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = vbq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{vendorbrand.Label}
	default:
		err = &NotSingularError{vendorbrand.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vbq *VendorBrandQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := vbq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of VendorBrands.
func (vbq *VendorBrandQuery) All(ctx context.Context) ([]*VendorBrand, error) {
	if err := vbq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return vbq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (vbq *VendorBrandQuery) AllX(ctx context.Context) []*VendorBrand {
	nodes, err := vbq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of VendorBrand IDs.
func (vbq *VendorBrandQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := vbq.Select(vendorbrand.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vbq *VendorBrandQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := vbq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vbq *VendorBrandQuery) Count(ctx context.Context) (int, error) {
	if err := vbq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return vbq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (vbq *VendorBrandQuery) CountX(ctx context.Context) int {
	count, err := vbq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vbq *VendorBrandQuery) Exist(ctx context.Context) (bool, error) {
	if err := vbq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return vbq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (vbq *VendorBrandQuery) ExistX(ctx context.Context) bool {
	exist, err := vbq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VendorBrandQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vbq *VendorBrandQuery) Clone() *VendorBrandQuery {
	if vbq == nil {
		return nil
	}
	return &VendorBrandQuery{
		config:     vbq.config,
		limit:      vbq.limit,
		offset:     vbq.offset,
		order:      append([]OrderFunc{}, vbq.order...),
		predicates: append([]predicate.VendorBrand{}, vbq.predicates...),
		// clone intermediate query.
		sql:    vbq.sql.Clone(),
		path:   vbq.path,
		unique: vbq.unique,
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
//	client.VendorBrand.Query().
//		GroupBy(vendorbrand.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (vbq *VendorBrandQuery) GroupBy(field string, fields ...string) *VendorBrandGroupBy {
	grbuild := &VendorBrandGroupBy{config: vbq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := vbq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return vbq.sqlQuery(ctx), nil
	}
	grbuild.label = vendorbrand.Label
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
//	client.VendorBrand.Query().
//		Select(vendorbrand.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (vbq *VendorBrandQuery) Select(fields ...string) *VendorBrandSelect {
	vbq.fields = append(vbq.fields, fields...)
	selbuild := &VendorBrandSelect{VendorBrandQuery: vbq}
	selbuild.label = vendorbrand.Label
	selbuild.flds, selbuild.scan = &vbq.fields, selbuild.Scan
	return selbuild
}

func (vbq *VendorBrandQuery) prepareQuery(ctx context.Context) error {
	for _, f := range vbq.fields {
		if !vendorbrand.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if vbq.path != nil {
		prev, err := vbq.path(ctx)
		if err != nil {
			return err
		}
		vbq.sql = prev
	}
	if vendorbrand.Policy == nil {
		return errors.New("ent: uninitialized vendorbrand.Policy (forgotten import ent/runtime?)")
	}
	if err := vendorbrand.Policy.EvalQuery(ctx, vbq); err != nil {
		return err
	}
	return nil
}

func (vbq *VendorBrandQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*VendorBrand, error) {
	var (
		nodes = []*VendorBrand{}
		_spec = vbq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*VendorBrand).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &VendorBrand{config: vbq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(vbq.modifiers) > 0 {
		_spec.Modifiers = vbq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vbq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (vbq *VendorBrandQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vbq.querySpec()
	if len(vbq.modifiers) > 0 {
		_spec.Modifiers = vbq.modifiers
	}
	_spec.Node.Columns = vbq.fields
	if len(vbq.fields) > 0 {
		_spec.Unique = vbq.unique != nil && *vbq.unique
	}
	return sqlgraph.CountNodes(ctx, vbq.driver, _spec)
}

func (vbq *VendorBrandQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := vbq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (vbq *VendorBrandQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vendorbrand.Table,
			Columns: vendorbrand.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: vendorbrand.FieldID,
			},
		},
		From:   vbq.sql,
		Unique: true,
	}
	if unique := vbq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := vbq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vendorbrand.FieldID)
		for i := range fields {
			if fields[i] != vendorbrand.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := vbq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vbq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vbq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vbq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vbq *VendorBrandQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vbq.driver.Dialect())
	t1 := builder.Table(vendorbrand.Table)
	columns := vbq.fields
	if len(columns) == 0 {
		columns = vendorbrand.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vbq.sql != nil {
		selector = vbq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vbq.unique != nil && *vbq.unique {
		selector.Distinct()
	}
	for _, m := range vbq.modifiers {
		m(selector)
	}
	for _, p := range vbq.predicates {
		p(selector)
	}
	for _, p := range vbq.order {
		p(selector)
	}
	if offset := vbq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vbq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (vbq *VendorBrandQuery) ForUpdate(opts ...sql.LockOption) *VendorBrandQuery {
	if vbq.driver.Dialect() == dialect.Postgres {
		vbq.Unique(false)
	}
	vbq.modifiers = append(vbq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return vbq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (vbq *VendorBrandQuery) ForShare(opts ...sql.LockOption) *VendorBrandQuery {
	if vbq.driver.Dialect() == dialect.Postgres {
		vbq.Unique(false)
	}
	vbq.modifiers = append(vbq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return vbq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vbq *VendorBrandQuery) Modify(modifiers ...func(s *sql.Selector)) *VendorBrandSelect {
	vbq.modifiers = append(vbq.modifiers, modifiers...)
	return vbq.Select()
}

// VendorBrandGroupBy is the group-by builder for VendorBrand entities.
type VendorBrandGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vbgb *VendorBrandGroupBy) Aggregate(fns ...AggregateFunc) *VendorBrandGroupBy {
	vbgb.fns = append(vbgb.fns, fns...)
	return vbgb
}

// Scan applies the group-by query and scans the result into the given value.
func (vbgb *VendorBrandGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := vbgb.path(ctx)
	if err != nil {
		return err
	}
	vbgb.sql = query
	return vbgb.sqlScan(ctx, v)
}

func (vbgb *VendorBrandGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range vbgb.fields {
		if !vendorbrand.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := vbgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vbgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (vbgb *VendorBrandGroupBy) sqlQuery() *sql.Selector {
	selector := vbgb.sql.Select()
	aggregation := make([]string, 0, len(vbgb.fns))
	for _, fn := range vbgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(vbgb.fields)+len(vbgb.fns))
		for _, f := range vbgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(vbgb.fields...)...)
}

// VendorBrandSelect is the builder for selecting fields of VendorBrand entities.
type VendorBrandSelect struct {
	*VendorBrandQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (vbs *VendorBrandSelect) Scan(ctx context.Context, v interface{}) error {
	if err := vbs.prepareQuery(ctx); err != nil {
		return err
	}
	vbs.sql = vbs.VendorBrandQuery.sqlQuery(ctx)
	return vbs.sqlScan(ctx, v)
}

func (vbs *VendorBrandSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := vbs.sql.Query()
	if err := vbs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vbs *VendorBrandSelect) Modify(modifiers ...func(s *sql.Selector)) *VendorBrandSelect {
	vbs.modifiers = append(vbs.modifiers, modifiers...)
	return vbs
}
