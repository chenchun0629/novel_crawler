// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/novel_crawler/internal/data/ent/novel"
	"github.com/novel_crawler/internal/data/ent/novelclue"
	"github.com/novel_crawler/internal/data/ent/predicate"
)

// NovelClueQuery is the builder for querying NovelClue entities.
type NovelClueQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.NovelClue
	// eager-loading edges.
	withNovel *NovelQuery
	withFKs   bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NovelClueQuery builder.
func (ncq *NovelClueQuery) Where(ps ...predicate.NovelClue) *NovelClueQuery {
	ncq.predicates = append(ncq.predicates, ps...)
	return ncq
}

// Limit adds a limit step to the query.
func (ncq *NovelClueQuery) Limit(limit int) *NovelClueQuery {
	ncq.limit = &limit
	return ncq
}

// Offset adds an offset step to the query.
func (ncq *NovelClueQuery) Offset(offset int) *NovelClueQuery {
	ncq.offset = &offset
	return ncq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ncq *NovelClueQuery) Unique(unique bool) *NovelClueQuery {
	ncq.unique = &unique
	return ncq
}

// Order adds an order step to the query.
func (ncq *NovelClueQuery) Order(o ...OrderFunc) *NovelClueQuery {
	ncq.order = append(ncq.order, o...)
	return ncq
}

// QueryNovel chains the current query on the "novel" edge.
func (ncq *NovelClueQuery) QueryNovel() *NovelQuery {
	query := &NovelQuery{config: ncq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ncq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(novelclue.Table, novelclue.FieldID, selector),
			sqlgraph.To(novel.Table, novel.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, novelclue.NovelTable, novelclue.NovelColumn),
		)
		fromU = sqlgraph.SetNeighbors(ncq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first NovelClue entity from the query.
// Returns a *NotFoundError when no NovelClue was found.
func (ncq *NovelClueQuery) First(ctx context.Context) (*NovelClue, error) {
	nodes, err := ncq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{novelclue.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ncq *NovelClueQuery) FirstX(ctx context.Context) *NovelClue {
	node, err := ncq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NovelClue ID from the query.
// Returns a *NotFoundError when no NovelClue ID was found.
func (ncq *NovelClueQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ncq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{novelclue.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ncq *NovelClueQuery) FirstIDX(ctx context.Context) int {
	id, err := ncq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NovelClue entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one NovelClue entity is not found.
// Returns a *NotFoundError when no NovelClue entities are found.
func (ncq *NovelClueQuery) Only(ctx context.Context) (*NovelClue, error) {
	nodes, err := ncq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{novelclue.Label}
	default:
		return nil, &NotSingularError{novelclue.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ncq *NovelClueQuery) OnlyX(ctx context.Context) *NovelClue {
	node, err := ncq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NovelClue ID in the query.
// Returns a *NotSingularError when exactly one NovelClue ID is not found.
// Returns a *NotFoundError when no entities are found.
func (ncq *NovelClueQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ncq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = &NotSingularError{novelclue.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ncq *NovelClueQuery) OnlyIDX(ctx context.Context) int {
	id, err := ncq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NovelClues.
func (ncq *NovelClueQuery) All(ctx context.Context) ([]*NovelClue, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ncq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ncq *NovelClueQuery) AllX(ctx context.Context) []*NovelClue {
	nodes, err := ncq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NovelClue IDs.
func (ncq *NovelClueQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ncq.Select(novelclue.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ncq *NovelClueQuery) IDsX(ctx context.Context) []int {
	ids, err := ncq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ncq *NovelClueQuery) Count(ctx context.Context) (int, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ncq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ncq *NovelClueQuery) CountX(ctx context.Context) int {
	count, err := ncq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ncq *NovelClueQuery) Exist(ctx context.Context) (bool, error) {
	if err := ncq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ncq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ncq *NovelClueQuery) ExistX(ctx context.Context) bool {
	exist, err := ncq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NovelClueQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ncq *NovelClueQuery) Clone() *NovelClueQuery {
	if ncq == nil {
		return nil
	}
	return &NovelClueQuery{
		config:     ncq.config,
		limit:      ncq.limit,
		offset:     ncq.offset,
		order:      append([]OrderFunc{}, ncq.order...),
		predicates: append([]predicate.NovelClue{}, ncq.predicates...),
		withNovel:  ncq.withNovel.Clone(),
		// clone intermediate query.
		sql:  ncq.sql.Clone(),
		path: ncq.path,
	}
}

// WithNovel tells the query-builder to eager-load the nodes that are connected to
// the "novel" edge. The optional arguments are used to configure the query builder of the edge.
func (ncq *NovelClueQuery) WithNovel(opts ...func(*NovelQuery)) *NovelClueQuery {
	query := &NovelQuery{config: ncq.config}
	for _, opt := range opts {
		opt(query)
	}
	ncq.withNovel = query
	return ncq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NovelClue.Query().
//		GroupBy(novelclue.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ncq *NovelClueQuery) GroupBy(field string, fields ...string) *NovelClueGroupBy {
	group := &NovelClueGroupBy{config: ncq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ncq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ncq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.NovelClue.Query().
//		Select(novelclue.FieldCreateTime).
//		Scan(ctx, &v)
//
func (ncq *NovelClueQuery) Select(fields ...string) *NovelClueSelect {
	ncq.fields = append(ncq.fields, fields...)
	return &NovelClueSelect{NovelClueQuery: ncq}
}

func (ncq *NovelClueQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ncq.fields {
		if !novelclue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ncq.path != nil {
		prev, err := ncq.path(ctx)
		if err != nil {
			return err
		}
		ncq.sql = prev
	}
	return nil
}

func (ncq *NovelClueQuery) sqlAll(ctx context.Context) ([]*NovelClue, error) {
	var (
		nodes       = []*NovelClue{}
		withFKs     = ncq.withFKs
		_spec       = ncq.querySpec()
		loadedTypes = [1]bool{
			ncq.withNovel != nil,
		}
	)
	if ncq.withNovel != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, novelclue.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &NovelClue{config: ncq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, ncq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ncq.withNovel; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*NovelClue)
		for i := range nodes {
			if nodes[i].novel_clue == nil {
				continue
			}
			fk := *nodes[i].novel_clue
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(novel.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "novel_clue" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Novel = n
			}
		}
	}

	return nodes, nil
}

func (ncq *NovelClueQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ncq.querySpec()
	return sqlgraph.CountNodes(ctx, ncq.driver, _spec)
}

func (ncq *NovelClueQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ncq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ncq *NovelClueQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   novelclue.Table,
			Columns: novelclue.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: novelclue.FieldID,
			},
		},
		From:   ncq.sql,
		Unique: true,
	}
	if unique := ncq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ncq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, novelclue.FieldID)
		for i := range fields {
			if fields[i] != novelclue.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ncq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ncq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ncq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ncq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ncq *NovelClueQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ncq.driver.Dialect())
	t1 := builder.Table(novelclue.Table)
	columns := ncq.fields
	if len(columns) == 0 {
		columns = novelclue.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ncq.sql != nil {
		selector = ncq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range ncq.predicates {
		p(selector)
	}
	for _, p := range ncq.order {
		p(selector)
	}
	if offset := ncq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ncq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NovelClueGroupBy is the group-by builder for NovelClue entities.
type NovelClueGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ncgb *NovelClueGroupBy) Aggregate(fns ...AggregateFunc) *NovelClueGroupBy {
	ncgb.fns = append(ncgb.fns, fns...)
	return ncgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ncgb *NovelClueGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ncgb.path(ctx)
	if err != nil {
		return err
	}
	ncgb.sql = query
	return ncgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ncgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelClueGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) StringsX(ctx context.Context) []string {
	v, err := ncgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ncgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) StringX(ctx context.Context) string {
	v, err := ncgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelClueGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) IntsX(ctx context.Context) []int {
	v, err := ncgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ncgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) IntX(ctx context.Context) int {
	v, err := ncgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelClueGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ncgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ncgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ncgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ncgb.fields) > 1 {
		return nil, errors.New("ent: NovelClueGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ncgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ncgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ncgb *NovelClueGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ncgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ncgb *NovelClueGroupBy) BoolX(ctx context.Context) bool {
	v, err := ncgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ncgb *NovelClueGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ncgb.fields {
		if !novelclue.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ncgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ncgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ncgb *NovelClueGroupBy) sqlQuery() *sql.Selector {
	selector := ncgb.sql.Select()
	aggregation := make([]string, 0, len(ncgb.fns))
	for _, fn := range ncgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ncgb.fields)+len(ncgb.fns))
		for _, f := range ncgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ncgb.fields...)...)
}

// NovelClueSelect is the builder for selecting fields of NovelClue entities.
type NovelClueSelect struct {
	*NovelClueQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ncs *NovelClueSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ncs.prepareQuery(ctx); err != nil {
		return err
	}
	ncs.sql = ncs.NovelClueQuery.sqlQuery(ctx)
	return ncs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ncs *NovelClueSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ncs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelClueSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ncs *NovelClueSelect) StringsX(ctx context.Context) []string {
	v, err := ncs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ncs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ncs *NovelClueSelect) StringX(ctx context.Context) string {
	v, err := ncs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelClueSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ncs *NovelClueSelect) IntsX(ctx context.Context) []int {
	v, err := ncs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ncs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ncs *NovelClueSelect) IntX(ctx context.Context) int {
	v, err := ncs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelClueSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ncs *NovelClueSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ncs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ncs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ncs *NovelClueSelect) Float64X(ctx context.Context) float64 {
	v, err := ncs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ncs.fields) > 1 {
		return nil, errors.New("ent: NovelClueSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ncs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ncs *NovelClueSelect) BoolsX(ctx context.Context) []bool {
	v, err := ncs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ncs *NovelClueSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ncs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{novelclue.Label}
	default:
		err = fmt.Errorf("ent: NovelClueSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ncs *NovelClueSelect) BoolX(ctx context.Context) bool {
	v, err := ncs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ncs *NovelClueSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ncs.sql.Query()
	if err := ncs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
