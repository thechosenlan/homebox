// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/group"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/item"
	"github.com/thechosenlan/homebox/backend/internal/data/ent/location"
)

// LocationCreate is the builder for creating a Location entity.
type LocationCreate struct {
	config
	mutation *LocationMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (lc *LocationCreate) SetCreatedAt(t time.Time) *LocationCreate {
	lc.mutation.SetCreatedAt(t)
	return lc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (lc *LocationCreate) SetNillableCreatedAt(t *time.Time) *LocationCreate {
	if t != nil {
		lc.SetCreatedAt(*t)
	}
	return lc
}

// SetUpdatedAt sets the "updated_at" field.
func (lc *LocationCreate) SetUpdatedAt(t time.Time) *LocationCreate {
	lc.mutation.SetUpdatedAt(t)
	return lc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (lc *LocationCreate) SetNillableUpdatedAt(t *time.Time) *LocationCreate {
	if t != nil {
		lc.SetUpdatedAt(*t)
	}
	return lc
}

// SetName sets the "name" field.
func (lc *LocationCreate) SetName(s string) *LocationCreate {
	lc.mutation.SetName(s)
	return lc
}

// SetDescription sets the "description" field.
func (lc *LocationCreate) SetDescription(s string) *LocationCreate {
	lc.mutation.SetDescription(s)
	return lc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (lc *LocationCreate) SetNillableDescription(s *string) *LocationCreate {
	if s != nil {
		lc.SetDescription(*s)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LocationCreate) SetID(u uuid.UUID) *LocationCreate {
	lc.mutation.SetID(u)
	return lc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (lc *LocationCreate) SetNillableID(u *uuid.UUID) *LocationCreate {
	if u != nil {
		lc.SetID(*u)
	}
	return lc
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (lc *LocationCreate) SetGroupID(id uuid.UUID) *LocationCreate {
	lc.mutation.SetGroupID(id)
	return lc
}

// SetGroup sets the "group" edge to the Group entity.
func (lc *LocationCreate) SetGroup(g *Group) *LocationCreate {
	return lc.SetGroupID(g.ID)
}

// SetParentID sets the "parent" edge to the Location entity by ID.
func (lc *LocationCreate) SetParentID(id uuid.UUID) *LocationCreate {
	lc.mutation.SetParentID(id)
	return lc
}

// SetNillableParentID sets the "parent" edge to the Location entity by ID if the given value is not nil.
func (lc *LocationCreate) SetNillableParentID(id *uuid.UUID) *LocationCreate {
	if id != nil {
		lc = lc.SetParentID(*id)
	}
	return lc
}

// SetParent sets the "parent" edge to the Location entity.
func (lc *LocationCreate) SetParent(l *Location) *LocationCreate {
	return lc.SetParentID(l.ID)
}

// AddChildIDs adds the "children" edge to the Location entity by IDs.
func (lc *LocationCreate) AddChildIDs(ids ...uuid.UUID) *LocationCreate {
	lc.mutation.AddChildIDs(ids...)
	return lc
}

// AddChildren adds the "children" edges to the Location entity.
func (lc *LocationCreate) AddChildren(l ...*Location) *LocationCreate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return lc.AddChildIDs(ids...)
}

// AddItemIDs adds the "items" edge to the Item entity by IDs.
func (lc *LocationCreate) AddItemIDs(ids ...uuid.UUID) *LocationCreate {
	lc.mutation.AddItemIDs(ids...)
	return lc
}

// AddItems adds the "items" edges to the Item entity.
func (lc *LocationCreate) AddItems(i ...*Item) *LocationCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return lc.AddItemIDs(ids...)
}

// Mutation returns the LocationMutation object of the builder.
func (lc *LocationCreate) Mutation() *LocationMutation {
	return lc.mutation
}

// Save creates the Location in the database.
func (lc *LocationCreate) Save(ctx context.Context) (*Location, error) {
	lc.defaults()
	return withHooks[*Location, LocationMutation](ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LocationCreate) SaveX(ctx context.Context) *Location {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LocationCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LocationCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LocationCreate) defaults() {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		v := location.DefaultCreatedAt()
		lc.mutation.SetCreatedAt(v)
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		v := location.DefaultUpdatedAt()
		lc.mutation.SetUpdatedAt(v)
	}
	if _, ok := lc.mutation.ID(); !ok {
		v := location.DefaultID()
		lc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LocationCreate) check() error {
	if _, ok := lc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Location.created_at"`)}
	}
	if _, ok := lc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Location.updated_at"`)}
	}
	if _, ok := lc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Location.name"`)}
	}
	if v, ok := lc.mutation.Name(); ok {
		if err := location.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Location.name": %w`, err)}
		}
	}
	if v, ok := lc.mutation.Description(); ok {
		if err := location.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Location.description": %w`, err)}
		}
	}
	if _, ok := lc.mutation.GroupID(); !ok {
		return &ValidationError{Name: "group", err: errors.New(`ent: missing required edge "Location.group"`)}
	}
	return nil
}

func (lc *LocationCreate) sqlSave(ctx context.Context) (*Location, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LocationCreate) createSpec() (*Location, *sqlgraph.CreateSpec) {
	var (
		_node = &Location{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(location.Table, sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID))
	)
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := lc.mutation.CreatedAt(); ok {
		_spec.SetField(location.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := lc.mutation.UpdatedAt(); ok {
		_spec.SetField(location.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := lc.mutation.Name(); ok {
		_spec.SetField(location.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := lc.mutation.Description(); ok {
		_spec.SetField(location.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := lc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.GroupTable,
			Columns: []string{location.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_locations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   location.ParentTable,
			Columns: []string{location.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.location_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ChildrenTable,
			Columns: []string{location.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(location.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lc.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   location.ItemsTable,
			Columns: []string{location.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// LocationCreateBulk is the builder for creating many Location entities in bulk.
type LocationCreateBulk struct {
	config
	builders []*LocationCreate
}

// Save creates the Location entities in the database.
func (lcb *LocationCreateBulk) Save(ctx context.Context) ([]*Location, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*Location, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LocationMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LocationCreateBulk) SaveX(ctx context.Context) []*Location {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LocationCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LocationCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}
