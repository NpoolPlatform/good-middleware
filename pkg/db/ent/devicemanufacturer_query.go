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
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/devicemanufacturer"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/predicate"
)

// DeviceManufacturerQuery is the builder for querying DeviceManufacturer entities.
type DeviceManufacturerQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.DeviceManufacturer
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeviceManufacturerQuery builder.
func (dmq *DeviceManufacturerQuery) Where(ps ...predicate.DeviceManufacturer) *DeviceManufacturerQuery {
	dmq.predicates = append(dmq.predicates, ps...)
	return dmq
}

// Limit adds a limit step to the query.
func (dmq *DeviceManufacturerQuery) Limit(limit int) *DeviceManufacturerQuery {
	dmq.limit = &limit
	return dmq
}

// Offset adds an offset step to the query.
func (dmq *DeviceManufacturerQuery) Offset(offset int) *DeviceManufacturerQuery {
	dmq.offset = &offset
	return dmq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dmq *DeviceManufacturerQuery) Unique(unique bool) *DeviceManufacturerQuery {
	dmq.unique = &unique
	return dmq
}

// Order adds an order step to the query.
func (dmq *DeviceManufacturerQuery) Order(o ...OrderFunc) *DeviceManufacturerQuery {
	dmq.order = append(dmq.order, o...)
	return dmq
}

// First returns the first DeviceManufacturer entity from the query.
// Returns a *NotFoundError when no DeviceManufacturer was found.
func (dmq *DeviceManufacturerQuery) First(ctx context.Context) (*DeviceManufacturer, error) {
	nodes, err := dmq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{devicemanufacturer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) FirstX(ctx context.Context) *DeviceManufacturer {
	node, err := dmq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first DeviceManufacturer ID from the query.
// Returns a *NotFoundError when no DeviceManufacturer ID was found.
func (dmq *DeviceManufacturerQuery) FirstID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = dmq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{devicemanufacturer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) FirstIDX(ctx context.Context) uint32 {
	id, err := dmq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single DeviceManufacturer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one DeviceManufacturer entity is found.
// Returns a *NotFoundError when no DeviceManufacturer entities are found.
func (dmq *DeviceManufacturerQuery) Only(ctx context.Context) (*DeviceManufacturer, error) {
	nodes, err := dmq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{devicemanufacturer.Label}
	default:
		return nil, &NotSingularError{devicemanufacturer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) OnlyX(ctx context.Context) *DeviceManufacturer {
	node, err := dmq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only DeviceManufacturer ID in the query.
// Returns a *NotSingularError when more than one DeviceManufacturer ID is found.
// Returns a *NotFoundError when no entities are found.
func (dmq *DeviceManufacturerQuery) OnlyID(ctx context.Context) (id uint32, err error) {
	var ids []uint32
	if ids, err = dmq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{devicemanufacturer.Label}
	default:
		err = &NotSingularError{devicemanufacturer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) OnlyIDX(ctx context.Context) uint32 {
	id, err := dmq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of DeviceManufacturers.
func (dmq *DeviceManufacturerQuery) All(ctx context.Context) ([]*DeviceManufacturer, error) {
	if err := dmq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dmq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) AllX(ctx context.Context) []*DeviceManufacturer {
	nodes, err := dmq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of DeviceManufacturer IDs.
func (dmq *DeviceManufacturerQuery) IDs(ctx context.Context) ([]uint32, error) {
	var ids []uint32
	if err := dmq.Select(devicemanufacturer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) IDsX(ctx context.Context) []uint32 {
	ids, err := dmq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dmq *DeviceManufacturerQuery) Count(ctx context.Context) (int, error) {
	if err := dmq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dmq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) CountX(ctx context.Context) int {
	count, err := dmq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dmq *DeviceManufacturerQuery) Exist(ctx context.Context) (bool, error) {
	if err := dmq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dmq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dmq *DeviceManufacturerQuery) ExistX(ctx context.Context) bool {
	exist, err := dmq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeviceManufacturerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dmq *DeviceManufacturerQuery) Clone() *DeviceManufacturerQuery {
	if dmq == nil {
		return nil
	}
	return &DeviceManufacturerQuery{
		config:     dmq.config,
		limit:      dmq.limit,
		offset:     dmq.offset,
		order:      append([]OrderFunc{}, dmq.order...),
		predicates: append([]predicate.DeviceManufacturer{}, dmq.predicates...),
		// clone intermediate query.
		sql:    dmq.sql.Clone(),
		path:   dmq.path,
		unique: dmq.unique,
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
//	client.DeviceManufacturer.Query().
//		GroupBy(devicemanufacturer.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (dmq *DeviceManufacturerQuery) GroupBy(field string, fields ...string) *DeviceManufacturerGroupBy {
	grbuild := &DeviceManufacturerGroupBy{config: dmq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dmq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dmq.sqlQuery(ctx), nil
	}
	grbuild.label = devicemanufacturer.Label
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
//	client.DeviceManufacturer.Query().
//		Select(devicemanufacturer.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (dmq *DeviceManufacturerQuery) Select(fields ...string) *DeviceManufacturerSelect {
	dmq.fields = append(dmq.fields, fields...)
	selbuild := &DeviceManufacturerSelect{DeviceManufacturerQuery: dmq}
	selbuild.label = devicemanufacturer.Label
	selbuild.flds, selbuild.scan = &dmq.fields, selbuild.Scan
	return selbuild
}

func (dmq *DeviceManufacturerQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dmq.fields {
		if !devicemanufacturer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dmq.path != nil {
		prev, err := dmq.path(ctx)
		if err != nil {
			return err
		}
		dmq.sql = prev
	}
	if devicemanufacturer.Policy == nil {
		return errors.New("ent: uninitialized devicemanufacturer.Policy (forgotten import ent/runtime?)")
	}
	if err := devicemanufacturer.Policy.EvalQuery(ctx, dmq); err != nil {
		return err
	}
	return nil
}

func (dmq *DeviceManufacturerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*DeviceManufacturer, error) {
	var (
		nodes = []*DeviceManufacturer{}
		_spec = dmq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*DeviceManufacturer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &DeviceManufacturer{config: dmq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(dmq.modifiers) > 0 {
		_spec.Modifiers = dmq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dmq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (dmq *DeviceManufacturerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dmq.querySpec()
	if len(dmq.modifiers) > 0 {
		_spec.Modifiers = dmq.modifiers
	}
	_spec.Node.Columns = dmq.fields
	if len(dmq.fields) > 0 {
		_spec.Unique = dmq.unique != nil && *dmq.unique
	}
	return sqlgraph.CountNodes(ctx, dmq.driver, _spec)
}

func (dmq *DeviceManufacturerQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dmq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dmq *DeviceManufacturerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   devicemanufacturer.Table,
			Columns: devicemanufacturer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint32,
				Column: devicemanufacturer.FieldID,
			},
		},
		From:   dmq.sql,
		Unique: true,
	}
	if unique := dmq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dmq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, devicemanufacturer.FieldID)
		for i := range fields {
			if fields[i] != devicemanufacturer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dmq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dmq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dmq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dmq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dmq *DeviceManufacturerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dmq.driver.Dialect())
	t1 := builder.Table(devicemanufacturer.Table)
	columns := dmq.fields
	if len(columns) == 0 {
		columns = devicemanufacturer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dmq.sql != nil {
		selector = dmq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dmq.unique != nil && *dmq.unique {
		selector.Distinct()
	}
	for _, m := range dmq.modifiers {
		m(selector)
	}
	for _, p := range dmq.predicates {
		p(selector)
	}
	for _, p := range dmq.order {
		p(selector)
	}
	if offset := dmq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dmq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (dmq *DeviceManufacturerQuery) ForUpdate(opts ...sql.LockOption) *DeviceManufacturerQuery {
	if dmq.driver.Dialect() == dialect.Postgres {
		dmq.Unique(false)
	}
	dmq.modifiers = append(dmq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return dmq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (dmq *DeviceManufacturerQuery) ForShare(opts ...sql.LockOption) *DeviceManufacturerQuery {
	if dmq.driver.Dialect() == dialect.Postgres {
		dmq.Unique(false)
	}
	dmq.modifiers = append(dmq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return dmq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dmq *DeviceManufacturerQuery) Modify(modifiers ...func(s *sql.Selector)) *DeviceManufacturerSelect {
	dmq.modifiers = append(dmq.modifiers, modifiers...)
	return dmq.Select()
}

// DeviceManufacturerGroupBy is the group-by builder for DeviceManufacturer entities.
type DeviceManufacturerGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dmgb *DeviceManufacturerGroupBy) Aggregate(fns ...AggregateFunc) *DeviceManufacturerGroupBy {
	dmgb.fns = append(dmgb.fns, fns...)
	return dmgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dmgb *DeviceManufacturerGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dmgb.path(ctx)
	if err != nil {
		return err
	}
	dmgb.sql = query
	return dmgb.sqlScan(ctx, v)
}

func (dmgb *DeviceManufacturerGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dmgb.fields {
		if !devicemanufacturer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dmgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dmgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dmgb *DeviceManufacturerGroupBy) sqlQuery() *sql.Selector {
	selector := dmgb.sql.Select()
	aggregation := make([]string, 0, len(dmgb.fns))
	for _, fn := range dmgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dmgb.fields)+len(dmgb.fns))
		for _, f := range dmgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dmgb.fields...)...)
}

// DeviceManufacturerSelect is the builder for selecting fields of DeviceManufacturer entities.
type DeviceManufacturerSelect struct {
	*DeviceManufacturerQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (dms *DeviceManufacturerSelect) Scan(ctx context.Context, v interface{}) error {
	if err := dms.prepareQuery(ctx); err != nil {
		return err
	}
	dms.sql = dms.DeviceManufacturerQuery.sqlQuery(ctx)
	return dms.sqlScan(ctx, v)
}

func (dms *DeviceManufacturerSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := dms.sql.Query()
	if err := dms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (dms *DeviceManufacturerSelect) Modify(modifiers ...func(s *sql.Selector)) *DeviceManufacturerSelect {
	dms.modifiers = append(dms.modifiers, modifiers...)
	return dms
}