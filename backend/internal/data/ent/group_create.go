




// Code generated by ent, DO NOT EDIT.



package ent



import (
	"context"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
		"github.com/thechosenlan/homebox/backend/internal/data/ent/predicate"
				"github.com/google/uuid"
			"github.com/google/uuid"
	"entgo.io/ent"
	"entgo.io/ent/dialect"
			"database/sql/driver"
			"entgo.io/ent/dialect/sql"
			"entgo.io/ent/dialect/sql/sqlgraph"
			"entgo.io/ent/dialect/sql/sqljson"
			"entgo.io/ent/schema/field"

)


import (
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/group"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/user"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/location"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/item"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/label"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/document"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/groupinvitationtoken"
)





// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks []Hook
}


	




	


	
	
	// SetCreatedAt sets the "created_at" field.
	func (gc *GroupCreate) SetCreatedAt(t time.Time) *GroupCreate {
		gc.mutation.SetCreatedAt(t)
		return gc
	}

	
	
		
		// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
		func (gc *GroupCreate) SetNillableCreatedAt(t *time.Time) *GroupCreate {
			if t != nil {
				gc.SetCreatedAt(*t)
			}
			return gc
		}
	

	

	

	

	
	
	// SetUpdatedAt sets the "updated_at" field.
	func (gc *GroupCreate) SetUpdatedAt(t time.Time) *GroupCreate {
		gc.mutation.SetUpdatedAt(t)
		return gc
	}

	
	
		
		// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
		func (gc *GroupCreate) SetNillableUpdatedAt(t *time.Time) *GroupCreate {
			if t != nil {
				gc.SetUpdatedAt(*t)
			}
			return gc
		}
	

	

	

	

	
	
	// SetName sets the "name" field.
	func (gc *GroupCreate) SetName(s string) *GroupCreate {
		gc.mutation.SetName(s)
		return gc
	}

	
	

	

	

	

	
	
	// SetCurrency sets the "currency" field.
	func (gc *GroupCreate) SetCurrency(gr group.Currency) *GroupCreate {
		gc.mutation.SetCurrency(gr)
		return gc
	}

	
	
		
		// SetNillableCurrency sets the "currency" field if the given value is not nil.
		func (gc *GroupCreate) SetNillableCurrency(gr *group.Currency) *GroupCreate {
			if gr != nil {
				gc.SetCurrency(*gr)
			}
			return gc
		}
	

	

	

	

	
	
	// SetID sets the "id" field.
	func (gc *GroupCreate) SetID(u uuid.UUID) *GroupCreate {
		gc.mutation.SetID(u)
		return gc
	}

	
	
		
		// SetNillableID sets the "id" field if the given value is not nil.
		func (gc *GroupCreate) SetNillableID(u *uuid.UUID) *GroupCreate {
			if u != nil {
				gc.SetID(*u)
			}
			return gc
		}
	

	

	

	



	
	
	
	
	
		// AddUserIDs adds the "users" edge to the User entity by IDs.
		func (gc *GroupCreate) AddUserIDs(ids ... uuid.UUID) *GroupCreate {
			gc.mutation.AddUserIDs(ids ...)
			return gc
		}
	
	
	
	
	
	// AddUsers adds the "users" edges to the User entity.
	func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
		ids := make([]uuid.UUID, len(u))
			for i := range u {
				ids[i] = u[i].ID
			}
			return gc.AddUserIDs(ids...)
	}

	
	
	
	
	
		// AddLocationIDs adds the "locations" edge to the Location entity by IDs.
		func (gc *GroupCreate) AddLocationIDs(ids ... uuid.UUID) *GroupCreate {
			gc.mutation.AddLocationIDs(ids ...)
			return gc
		}
	
	
	
	
	
	// AddLocations adds the "locations" edges to the Location entity.
	func (gc *GroupCreate) AddLocations(l ...*Location) *GroupCreate {
		ids := make([]uuid.UUID, len(l))
			for i := range l {
				ids[i] = l[i].ID
			}
			return gc.AddLocationIDs(ids...)
	}

	
	
	
	
	
		// AddItemIDs adds the "items" edge to the Item entity by IDs.
		func (gc *GroupCreate) AddItemIDs(ids ... uuid.UUID) *GroupCreate {
			gc.mutation.AddItemIDs(ids ...)
			return gc
		}
	
	
	
	
	
	// AddItems adds the "items" edges to the Item entity.
	func (gc *GroupCreate) AddItems(i ...*Item) *GroupCreate {
		ids := make([]uuid.UUID, len(i))
			for j := range i {
				ids[j] = i[j].ID
			}
			return gc.AddItemIDs(ids...)
	}

	
	
	
	
	
		// AddLabelIDs adds the "labels" edge to the Label entity by IDs.
		func (gc *GroupCreate) AddLabelIDs(ids ... uuid.UUID) *GroupCreate {
			gc.mutation.AddLabelIDs(ids ...)
			return gc
		}
	
	
	
	
	
	// AddLabels adds the "labels" edges to the Label entity.
	func (gc *GroupCreate) AddLabels(l ...*Label) *GroupCreate {
		ids := make([]uuid.UUID, len(l))
			for i := range l {
				ids[i] = l[i].ID
			}
			return gc.AddLabelIDs(ids...)
	}

	
	
	
	
	
		// AddDocumentIDs adds the "documents" edge to the Document entity by IDs.
		func (gc *GroupCreate) AddDocumentIDs(ids ... uuid.UUID) *GroupCreate {
			gc.mutation.AddDocumentIDs(ids ...)
			return gc
		}
	
	
	
	
	
	// AddDocuments adds the "documents" edges to the Document entity.
	func (gc *GroupCreate) AddDocuments(d ...*Document) *GroupCreate {
		ids := make([]uuid.UUID, len(d))
			for i := range d {
				ids[i] = d[i].ID
			}
			return gc.AddDocumentIDs(ids...)
	}

	
	
	
	
	
		// AddInvitationTokenIDs adds the "invitation_tokens" edge to the GroupInvitationToken entity by IDs.
		func (gc *GroupCreate) AddInvitationTokenIDs(ids ... uuid.UUID) *GroupCreate {
			gc.mutation.AddInvitationTokenIDs(ids ...)
			return gc
		}
	
	
	
	
	
	// AddInvitationTokens adds the "invitation_tokens" edges to the GroupInvitationToken entity.
	func (gc *GroupCreate) AddInvitationTokens(g ...*GroupInvitationToken) *GroupCreate {
		ids := make([]uuid.UUID, len(g))
			for i := range g {
				ids[i] = g[i].ID
			}
			return gc.AddInvitationTokenIDs(ids...)
	}


// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}




// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
			gc.defaults()
	return withHooks[*Group, GroupMutation](ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

	// defaults sets the default values of the builder before save.
	func (gc *GroupCreate) defaults() {
				if _, ok := gc.mutation.CreatedAt(); !ok {
					v := group.DefaultCreatedAt()
					gc.mutation.SetCreatedAt(v)
				}
				if _, ok := gc.mutation.UpdatedAt(); !ok {
					v := group.DefaultUpdatedAt()
					gc.mutation.SetUpdatedAt(v)
				}
				if _, ok := gc.mutation.Currency(); !ok {
					v := group.DefaultCurrency
					gc.mutation.SetCurrency(v)
				}
				if _, ok := gc.mutation.ID(); !ok {
					v := group.DefaultID()
					gc.mutation.SetID(v)
				}
	}


// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
					if _, ok := gc.mutation.CreatedAt(); !ok {
						return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Group.created_at"`)}
					}
					if _, ok := gc.mutation.UpdatedAt(); !ok {
						return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Group.updated_at"`)}
					}
					if _, ok := gc.mutation.Name(); !ok {
						return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
					}
			if v, ok := gc.mutation.Name(); ok {
				if err := group.NameValidator(v); err != nil {
					return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Group.name": %w`, err)}
				}
			}
					if _, ok := gc.mutation.Currency(); !ok {
						return &ValidationError{Name: "currency", err: errors.New(`ent: missing required field "Group.currency"`)}
					}
			if v, ok := gc.mutation.Currency(); ok {
				if err := group.CurrencyValidator(v); err != nil {
					return &ValidationError{Name: "currency", err: fmt.Errorf(`ent: validator failed for field "Group.currency": %w`, err)}
				}
			}
	return nil
}


	
	




func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
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
		gc.mutation.id = &_node.ID
		gc.mutation.done = true
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: group.Table,
				ID: &sqlgraph.FieldSpec{
					Type: field.TypeUUID,
					Column: group.FieldID,
				},
		}
	)
		if id, ok := gc.mutation.ID(); ok {
			_node.ID = id
			_spec.ID.Value = &id
		}
		if value, ok := gc.mutation.CreatedAt(); ok {
			_spec.SetField(group.FieldCreatedAt, field.TypeTime, value)
			_node.CreatedAt = value
		}
		if value, ok := gc.mutation.UpdatedAt(); ok {
			_spec.SetField(group.FieldUpdatedAt, field.TypeTime, value)
			_node.UpdatedAt = value
		}
		if value, ok := gc.mutation.Name(); ok {
			_spec.SetField(group.FieldName, field.TypeString, value)
			_node.Name = value
		}
		if value, ok := gc.mutation.Currency(); ok {
			_spec.SetField(group.FieldCurrency, field.TypeEnum, value)
			_node.Currency = value
		}
		if nodes := gc.mutation.UsersIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.O2M,
		Inverse: false,
		Table: group.UsersTable,
		Columns: []string{ group.UsersColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: user.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := gc.mutation.LocationsIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.O2M,
		Inverse: false,
		Table: group.LocationsTable,
		Columns: []string{ group.LocationsColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: location.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := gc.mutation.ItemsIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.O2M,
		Inverse: false,
		Table: group.ItemsTable,
		Columns: []string{ group.ItemsColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := gc.mutation.LabelsIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.O2M,
		Inverse: false,
		Table: group.LabelsTable,
		Columns: []string{ group.LabelsColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: label.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := gc.mutation.DocumentsIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.O2M,
		Inverse: false,
		Table: group.DocumentsTable,
		Columns: []string{ group.DocumentsColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: document.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
		if nodes := gc.mutation.InvitationTokensIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.O2M,
		Inverse: false,
		Table: group.InvitationTokensTable,
		Columns: []string{ group.InvitationTokensColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: groupinvitationtoken.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges = append(_spec.Edges, edge)
		}
	return _node, _spec
}
	








// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	builders []*GroupCreate
}




	
		



// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
				builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}
	


	

