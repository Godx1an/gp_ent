// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Godx1an/gp_ent/pkg/ent_work/fitnesstestitem"
	"github.com/Godx1an/gp_ent/pkg/ent_work/predicate"
)

// FitnessTestItemQuery is the builder for querying FitnessTestItem entities.
type FitnessTestItemQuery struct {
	config
	ctx        *QueryContext
	order      []fitnesstestitem.OrderOption
	inters     []Interceptor
	predicates []predicate.FitnessTestItem
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FitnessTestItemQuery builder.
func (ftiq *FitnessTestItemQuery) Where(ps ...predicate.FitnessTestItem) *FitnessTestItemQuery {
	ftiq.predicates = append(ftiq.predicates, ps...)
	return ftiq
}

// Limit the number of records to be returned by this query.
func (ftiq *FitnessTestItemQuery) Limit(limit int) *FitnessTestItemQuery {
	ftiq.ctx.Limit = &limit
	return ftiq
}

// Offset to start from.
func (ftiq *FitnessTestItemQuery) Offset(offset int) *FitnessTestItemQuery {
	ftiq.ctx.Offset = &offset
	return ftiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ftiq *FitnessTestItemQuery) Unique(unique bool) *FitnessTestItemQuery {
	ftiq.ctx.Unique = &unique
	return ftiq
}

// Order specifies how the records should be ordered.
func (ftiq *FitnessTestItemQuery) Order(o ...fitnesstestitem.OrderOption) *FitnessTestItemQuery {
	ftiq.order = append(ftiq.order, o...)
	return ftiq
}

// First returns the first FitnessTestItem entity from the query.
// Returns a *NotFoundError when no FitnessTestItem was found.
func (ftiq *FitnessTestItemQuery) First(ctx context.Context) (*FitnessTestItem, error) {
	nodes, err := ftiq.Limit(1).All(setContextOp(ctx, ftiq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fitnesstestitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) FirstX(ctx context.Context) *FitnessTestItem {
	node, err := ftiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FitnessTestItem ID from the query.
// Returns a *NotFoundError when no FitnessTestItem ID was found.
func (ftiq *FitnessTestItemQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ftiq.Limit(1).IDs(setContextOp(ctx, ftiq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fitnesstestitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) FirstIDX(ctx context.Context) int64 {
	id, err := ftiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FitnessTestItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FitnessTestItem entity is found.
// Returns a *NotFoundError when no FitnessTestItem entities are found.
func (ftiq *FitnessTestItemQuery) Only(ctx context.Context) (*FitnessTestItem, error) {
	nodes, err := ftiq.Limit(2).All(setContextOp(ctx, ftiq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fitnesstestitem.Label}
	default:
		return nil, &NotSingularError{fitnesstestitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) OnlyX(ctx context.Context) *FitnessTestItem {
	node, err := ftiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FitnessTestItem ID in the query.
// Returns a *NotSingularError when more than one FitnessTestItem ID is found.
// Returns a *NotFoundError when no entities are found.
func (ftiq *FitnessTestItemQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = ftiq.Limit(2).IDs(setContextOp(ctx, ftiq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fitnesstestitem.Label}
	default:
		err = &NotSingularError{fitnesstestitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := ftiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FitnessTestItems.
func (ftiq *FitnessTestItemQuery) All(ctx context.Context) ([]*FitnessTestItem, error) {
	ctx = setContextOp(ctx, ftiq.ctx, "All")
	if err := ftiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FitnessTestItem, *FitnessTestItemQuery]()
	return withInterceptors[[]*FitnessTestItem](ctx, ftiq, qr, ftiq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) AllX(ctx context.Context) []*FitnessTestItem {
	nodes, err := ftiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FitnessTestItem IDs.
func (ftiq *FitnessTestItemQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if ftiq.ctx.Unique == nil && ftiq.path != nil {
		ftiq.Unique(true)
	}
	ctx = setContextOp(ctx, ftiq.ctx, "IDs")
	if err = ftiq.Select(fitnesstestitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) IDsX(ctx context.Context) []int64 {
	ids, err := ftiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftiq *FitnessTestItemQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ftiq.ctx, "Count")
	if err := ftiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ftiq, querierCount[*FitnessTestItemQuery](), ftiq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) CountX(ctx context.Context) int {
	count, err := ftiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftiq *FitnessTestItemQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ftiq.ctx, "Exist")
	switch _, err := ftiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent_work: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ftiq *FitnessTestItemQuery) ExistX(ctx context.Context) bool {
	exist, err := ftiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FitnessTestItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftiq *FitnessTestItemQuery) Clone() *FitnessTestItemQuery {
	if ftiq == nil {
		return nil
	}
	return &FitnessTestItemQuery{
		config:     ftiq.config,
		ctx:        ftiq.ctx.Clone(),
		order:      append([]fitnesstestitem.OrderOption{}, ftiq.order...),
		inters:     append([]Interceptor{}, ftiq.inters...),
		predicates: append([]predicate.FitnessTestItem{}, ftiq.predicates...),
		// clone intermediate query.
		sql:  ftiq.sql.Clone(),
		path: ftiq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FitnessTestItem.Query().
//		GroupBy(fitnesstestitem.FieldCreatedBy).
//		Aggregate(ent_work.Count()).
//		Scan(ctx, &v)
func (ftiq *FitnessTestItemQuery) GroupBy(field string, fields ...string) *FitnessTestItemGroupBy {
	ftiq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FitnessTestItemGroupBy{build: ftiq}
	grbuild.flds = &ftiq.ctx.Fields
	grbuild.label = fitnesstestitem.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedBy int64 `json:"created_by"`
//	}
//
//	client.FitnessTestItem.Query().
//		Select(fitnesstestitem.FieldCreatedBy).
//		Scan(ctx, &v)
func (ftiq *FitnessTestItemQuery) Select(fields ...string) *FitnessTestItemSelect {
	ftiq.ctx.Fields = append(ftiq.ctx.Fields, fields...)
	sbuild := &FitnessTestItemSelect{FitnessTestItemQuery: ftiq}
	sbuild.label = fitnesstestitem.Label
	sbuild.flds, sbuild.scan = &ftiq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FitnessTestItemSelect configured with the given aggregations.
func (ftiq *FitnessTestItemQuery) Aggregate(fns ...AggregateFunc) *FitnessTestItemSelect {
	return ftiq.Select().Aggregate(fns...)
}

func (ftiq *FitnessTestItemQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ftiq.inters {
		if inter == nil {
			return fmt.Errorf("ent_work: uninitialized interceptor (forgotten import ent_work/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ftiq); err != nil {
				return err
			}
		}
	}
	for _, f := range ftiq.ctx.Fields {
		if !fitnesstestitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent_work: invalid field %q for query", f)}
		}
	}
	if ftiq.path != nil {
		prev, err := ftiq.path(ctx)
		if err != nil {
			return err
		}
		ftiq.sql = prev
	}
	return nil
}

func (ftiq *FitnessTestItemQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FitnessTestItem, error) {
	var (
		nodes = []*FitnessTestItem{}
		_spec = ftiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FitnessTestItem).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FitnessTestItem{config: ftiq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(ftiq.modifiers) > 0 {
		_spec.Modifiers = ftiq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ftiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ftiq *FitnessTestItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ftiq.querySpec()
	if len(ftiq.modifiers) > 0 {
		_spec.Modifiers = ftiq.modifiers
	}
	_spec.Node.Columns = ftiq.ctx.Fields
	if len(ftiq.ctx.Fields) > 0 {
		_spec.Unique = ftiq.ctx.Unique != nil && *ftiq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ftiq.driver, _spec)
}

func (ftiq *FitnessTestItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(fitnesstestitem.Table, fitnesstestitem.Columns, sqlgraph.NewFieldSpec(fitnesstestitem.FieldID, field.TypeInt64))
	_spec.From = ftiq.sql
	if unique := ftiq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ftiq.path != nil {
		_spec.Unique = true
	}
	if fields := ftiq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fitnesstestitem.FieldID)
		for i := range fields {
			if fields[i] != fitnesstestitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ftiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ftiq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ftiq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ftiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ftiq *FitnessTestItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ftiq.driver.Dialect())
	t1 := builder.Table(fitnesstestitem.Table)
	columns := ftiq.ctx.Fields
	if len(columns) == 0 {
		columns = fitnesstestitem.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ftiq.sql != nil {
		selector = ftiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ftiq.ctx.Unique != nil && *ftiq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range ftiq.modifiers {
		m(selector)
	}
	for _, p := range ftiq.predicates {
		p(selector)
	}
	for _, p := range ftiq.order {
		p(selector)
	}
	if offset := ftiq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ftiq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ftiq *FitnessTestItemQuery) Modify(modifiers ...func(s *sql.Selector)) *FitnessTestItemSelect {
	ftiq.modifiers = append(ftiq.modifiers, modifiers...)
	return ftiq.Select()
}

// FitnessTestItemGroupBy is the group-by builder for FitnessTestItem entities.
type FitnessTestItemGroupBy struct {
	selector
	build *FitnessTestItemQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftigb *FitnessTestItemGroupBy) Aggregate(fns ...AggregateFunc) *FitnessTestItemGroupBy {
	ftigb.fns = append(ftigb.fns, fns...)
	return ftigb
}

// Scan applies the selector query and scans the result into the given value.
func (ftigb *FitnessTestItemGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ftigb.build.ctx, "GroupBy")
	if err := ftigb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FitnessTestItemQuery, *FitnessTestItemGroupBy](ctx, ftigb.build, ftigb, ftigb.build.inters, v)
}

func (ftigb *FitnessTestItemGroupBy) sqlScan(ctx context.Context, root *FitnessTestItemQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ftigb.fns))
	for _, fn := range ftigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ftigb.flds)+len(ftigb.fns))
		for _, f := range *ftigb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ftigb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftigb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FitnessTestItemSelect is the builder for selecting fields of FitnessTestItem entities.
type FitnessTestItemSelect struct {
	*FitnessTestItemQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ftis *FitnessTestItemSelect) Aggregate(fns ...AggregateFunc) *FitnessTestItemSelect {
	ftis.fns = append(ftis.fns, fns...)
	return ftis
}

// Scan applies the selector query and scans the result into the given value.
func (ftis *FitnessTestItemSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ftis.ctx, "Select")
	if err := ftis.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FitnessTestItemQuery, *FitnessTestItemSelect](ctx, ftis.FitnessTestItemQuery, ftis, ftis.inters, v)
}

func (ftis *FitnessTestItemSelect) sqlScan(ctx context.Context, root *FitnessTestItemQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ftis.fns))
	for _, fn := range ftis.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ftis.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ftis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ftis *FitnessTestItemSelect) Modify(modifiers ...func(s *sql.Selector)) *FitnessTestItemSelect {
	ftis.modifiers = append(ftis.modifiers, modifiers...)
	return ftis
}