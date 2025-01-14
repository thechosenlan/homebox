// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/group"
	"github.com/hay-kot/homebox/backend/internal/data/ent/notifier"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
	"github.com/hay-kot/homebox/backend/internal/data/ent/user"
)

// NotifierQuery is the builder for querying Notifier entities.
type NotifierQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Notifier
	withGroup  *GroupQuery
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotifierQuery builder.
func (nq *NotifierQuery) Where(ps ...predicate.Notifier) *NotifierQuery {
	nq.predicates = append(nq.predicates, ps...)
	return nq
}

// Limit the number of records to be returned by this query.
func (nq *NotifierQuery) Limit(limit int) *NotifierQuery {
	nq.ctx.Limit = &limit
	return nq
}

// Offset to start from.
func (nq *NotifierQuery) Offset(offset int) *NotifierQuery {
	nq.ctx.Offset = &offset
	return nq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (nq *NotifierQuery) Unique(unique bool) *NotifierQuery {
	nq.ctx.Unique = &unique
	return nq
}

// Order specifies how the records should be ordered.
func (nq *NotifierQuery) Order(o ...OrderFunc) *NotifierQuery {
	nq.order = append(nq.order, o...)
	return nq
}

// QueryGroup chains the current query on the "group" edge.
func (nq *NotifierQuery) QueryGroup() *GroupQuery {
	query := (&GroupClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notifier.Table, notifier.FieldID, selector),
			sqlgraph.To(group.Table, group.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notifier.GroupTable, notifier.GroupColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (nq *NotifierQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: nq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := nq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := nq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(notifier.Table, notifier.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, notifier.UserTable, notifier.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(nq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Notifier entity from the query.
// Returns a *NotFoundError when no Notifier was found.
func (nq *NotifierQuery) First(ctx context.Context) (*Notifier, error) {
	nodes, err := nq.Limit(1).All(setContextOp(ctx, nq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notifier.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (nq *NotifierQuery) FirstX(ctx context.Context) *Notifier {
	node, err := nq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Notifier ID from the query.
// Returns a *NotFoundError when no Notifier ID was found.
func (nq *NotifierQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(1).IDs(setContextOp(ctx, nq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notifier.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (nq *NotifierQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := nq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Notifier entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Notifier entity is found.
// Returns a *NotFoundError when no Notifier entities are found.
func (nq *NotifierQuery) Only(ctx context.Context) (*Notifier, error) {
	nodes, err := nq.Limit(2).All(setContextOp(ctx, nq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notifier.Label}
	default:
		return nil, &NotSingularError{notifier.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (nq *NotifierQuery) OnlyX(ctx context.Context) *Notifier {
	node, err := nq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Notifier ID in the query.
// Returns a *NotSingularError when more than one Notifier ID is found.
// Returns a *NotFoundError when no entities are found.
func (nq *NotifierQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = nq.Limit(2).IDs(setContextOp(ctx, nq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notifier.Label}
	default:
		err = &NotSingularError{notifier.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (nq *NotifierQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := nq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Notifiers.
func (nq *NotifierQuery) All(ctx context.Context) ([]*Notifier, error) {
	ctx = setContextOp(ctx, nq.ctx, "All")
	if err := nq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Notifier, *NotifierQuery]()
	return withInterceptors[[]*Notifier](ctx, nq, qr, nq.inters)
}

// AllX is like All, but panics if an error occurs.
func (nq *NotifierQuery) AllX(ctx context.Context) []*Notifier {
	nodes, err := nq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Notifier IDs.
func (nq *NotifierQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if nq.ctx.Unique == nil && nq.path != nil {
		nq.Unique(true)
	}
	ctx = setContextOp(ctx, nq.ctx, "IDs")
	if err = nq.Select(notifier.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (nq *NotifierQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := nq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (nq *NotifierQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, nq.ctx, "Count")
	if err := nq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, nq, querierCount[*NotifierQuery](), nq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (nq *NotifierQuery) CountX(ctx context.Context) int {
	count, err := nq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (nq *NotifierQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, nq.ctx, "Exist")
	switch _, err := nq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (nq *NotifierQuery) ExistX(ctx context.Context) bool {
	exist, err := nq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotifierQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (nq *NotifierQuery) Clone() *NotifierQuery {
	if nq == nil {
		return nil
	}
	return &NotifierQuery{
		config:     nq.config,
		ctx:        nq.ctx.Clone(),
		order:      append([]OrderFunc{}, nq.order...),
		inters:     append([]Interceptor{}, nq.inters...),
		predicates: append([]predicate.Notifier{}, nq.predicates...),
		withGroup:  nq.withGroup.Clone(),
		withUser:   nq.withUser.Clone(),
		// clone intermediate query.
		sql:  nq.sql.Clone(),
		path: nq.path,
	}
}

// WithGroup tells the query-builder to eager-load the nodes that are connected to
// the "group" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotifierQuery) WithGroup(opts ...func(*GroupQuery)) *NotifierQuery {
	query := (&GroupClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withGroup = query
	return nq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (nq *NotifierQuery) WithUser(opts ...func(*UserQuery)) *NotifierQuery {
	query := (&UserClient{config: nq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	nq.withUser = query
	return nq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Notifier.Query().
//		GroupBy(notifier.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (nq *NotifierQuery) GroupBy(field string, fields ...string) *NotifierGroupBy {
	nq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotifierGroupBy{build: nq}
	grbuild.flds = &nq.ctx.Fields
	grbuild.label = notifier.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.Notifier.Query().
//		Select(notifier.FieldCreatedAt).
//		Scan(ctx, &v)
func (nq *NotifierQuery) Select(fields ...string) *NotifierSelect {
	nq.ctx.Fields = append(nq.ctx.Fields, fields...)
	sbuild := &NotifierSelect{NotifierQuery: nq}
	sbuild.label = notifier.Label
	sbuild.flds, sbuild.scan = &nq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotifierSelect configured with the given aggregations.
func (nq *NotifierQuery) Aggregate(fns ...AggregateFunc) *NotifierSelect {
	return nq.Select().Aggregate(fns...)
}

func (nq *NotifierQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range nq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, nq); err != nil {
				return err
			}
		}
	}
	for _, f := range nq.ctx.Fields {
		if !notifier.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if nq.path != nil {
		prev, err := nq.path(ctx)
		if err != nil {
			return err
		}
		nq.sql = prev
	}
	return nil
}

func (nq *NotifierQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Notifier, error) {
	var (
		nodes       = []*Notifier{}
		_spec       = nq.querySpec()
		loadedTypes = [2]bool{
			nq.withGroup != nil,
			nq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Notifier).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Notifier{config: nq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, nq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := nq.withGroup; query != nil {
		if err := nq.loadGroup(ctx, query, nodes, nil,
			func(n *Notifier, e *Group) { n.Edges.Group = e }); err != nil {
			return nil, err
		}
	}
	if query := nq.withUser; query != nil {
		if err := nq.loadUser(ctx, query, nodes, nil,
			func(n *Notifier, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (nq *NotifierQuery) loadGroup(ctx context.Context, query *GroupQuery, nodes []*Notifier, init func(*Notifier), assign func(*Notifier, *Group)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Notifier)
	for i := range nodes {
		fk := nodes[i].GroupID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(group.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "group_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (nq *NotifierQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*Notifier, init func(*Notifier), assign func(*Notifier, *User)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Notifier)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (nq *NotifierQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := nq.querySpec()
	_spec.Node.Columns = nq.ctx.Fields
	if len(nq.ctx.Fields) > 0 {
		_spec.Unique = nq.ctx.Unique != nil && *nq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, nq.driver, _spec)
}

func (nq *NotifierQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notifier.Table, notifier.Columns, sqlgraph.NewFieldSpec(notifier.FieldID, field.TypeUUID))
	_spec.From = nq.sql
	if unique := nq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if nq.path != nil {
		_spec.Unique = true
	}
	if fields := nq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifier.FieldID)
		for i := range fields {
			if fields[i] != notifier.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := nq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := nq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := nq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := nq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (nq *NotifierQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(nq.driver.Dialect())
	t1 := builder.Table(notifier.Table)
	columns := nq.ctx.Fields
	if len(columns) == 0 {
		columns = notifier.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if nq.sql != nil {
		selector = nq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if nq.ctx.Unique != nil && *nq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range nq.predicates {
		p(selector)
	}
	for _, p := range nq.order {
		p(selector)
	}
	if offset := nq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := nq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NotifierGroupBy is the group-by builder for Notifier entities.
type NotifierGroupBy struct {
	selector
	build *NotifierQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ngb *NotifierGroupBy) Aggregate(fns ...AggregateFunc) *NotifierGroupBy {
	ngb.fns = append(ngb.fns, fns...)
	return ngb
}

// Scan applies the selector query and scans the result into the given value.
func (ngb *NotifierGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ngb.build.ctx, "GroupBy")
	if err := ngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotifierQuery, *NotifierGroupBy](ctx, ngb.build, ngb, ngb.build.inters, v)
}

func (ngb *NotifierGroupBy) sqlScan(ctx context.Context, root *NotifierQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ngb.fns))
	for _, fn := range ngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ngb.flds)+len(ngb.fns))
		for _, f := range *ngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotifierSelect is the builder for selecting fields of Notifier entities.
type NotifierSelect struct {
	*NotifierQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ns *NotifierSelect) Aggregate(fns ...AggregateFunc) *NotifierSelect {
	ns.fns = append(ns.fns, fns...)
	return ns
}

// Scan applies the selector query and scans the result into the given value.
func (ns *NotifierSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ns.ctx, "Select")
	if err := ns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotifierQuery, *NotifierSelect](ctx, ns.NotifierQuery, ns, ns.inters, v)
}

func (ns *NotifierSelect) sqlScan(ctx context.Context, root *NotifierQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ns.fns))
	for _, fn := range ns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
