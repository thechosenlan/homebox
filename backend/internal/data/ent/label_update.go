



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
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/label"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/group"
		 "github.com/thechosenlan/homebox/backend/internal/data/ent/item"
)






// LabelUpdate is the builder for updating Label entities.
type LabelUpdate struct {
	config
	hooks []Hook
	mutation *LabelMutation
}

// Where appends a list predicates to the LabelUpdate builder.
func (lu *LabelUpdate) Where(ps ...predicate.Label) *LabelUpdate {
	lu.mutation.Where(ps...)
	return lu
}


	




	
	


	
	
	// SetUpdatedAt sets the "updated_at" field.
	func (lu *LabelUpdate) SetUpdatedAt(t time.Time) *LabelUpdate {
		lu.mutation.SetUpdatedAt(t)
		return lu
	}

	
	

	

	

	

	
	
	// SetName sets the "name" field.
	func (lu *LabelUpdate) SetName(s string) *LabelUpdate {
		lu.mutation.SetName(s)
		return lu
	}

	
	

	

	

	

	
	
	// SetDescription sets the "description" field.
	func (lu *LabelUpdate) SetDescription(s string) *LabelUpdate {
		lu.mutation.SetDescription(s)
		return lu
	}

	
	
		
		// SetNillableDescription sets the "description" field if the given value is not nil.
		func (lu *LabelUpdate) SetNillableDescription(s *string) *LabelUpdate {
			if s != nil {
				lu.SetDescription(*s)
			}
			return lu
		}
	

	

	

	
		
		// ClearDescription clears the value of the "description" field.
		func (lu *LabelUpdate) ClearDescription() *LabelUpdate {
			lu.mutation.ClearDescription()
			return lu
		}
	

	
	
	// SetColor sets the "color" field.
	func (lu *LabelUpdate) SetColor(s string) *LabelUpdate {
		lu.mutation.SetColor(s)
		return lu
	}

	
	
		
		// SetNillableColor sets the "color" field if the given value is not nil.
		func (lu *LabelUpdate) SetNillableColor(s *string) *LabelUpdate {
			if s != nil {
				lu.SetColor(*s)
			}
			return lu
		}
	

	

	

	
		
		// ClearColor clears the value of the "color" field.
		func (lu *LabelUpdate) ClearColor() *LabelUpdate {
			lu.mutation.ClearColor()
			return lu
		}
	



	
	
	
	
	
		// SetGroupID sets the "group" edge to the Group entity by ID.
		func (lu *LabelUpdate) SetGroupID(id uuid.UUID) *LabelUpdate {
			lu.mutation.SetGroupID(id)
			return lu
		}
	
	
	
	
	
	// SetGroup sets the "group" edge to the Group entity.
	func (lu *LabelUpdate) SetGroup(g *Group) *LabelUpdate {
		return lu.SetGroupID(g.ID)
	}

	
	
	
	
	
		// AddItemIDs adds the "items" edge to the Item entity by IDs.
		func (lu *LabelUpdate) AddItemIDs(ids ... uuid.UUID) *LabelUpdate {
			lu.mutation.AddItemIDs(ids ...)
			return lu
		}
	
	
	
	
	
	// AddItems adds the "items" edges to the Item entity.
	func (lu *LabelUpdate) AddItems(i ...*Item) *LabelUpdate {
		ids := make([]uuid.UUID, len(i))
			for j := range i {
				ids[j] = i[j].ID
			}
			return lu.AddItemIDs(ids...)
	}


// Mutation returns the LabelMutation object of the builder.
func (lu *LabelUpdate) Mutation() *LabelMutation {
	return lu.mutation
}





	





	
	
	// ClearGroup clears the "group" edge to the Group entity.
	func (lu *LabelUpdate) ClearGroup() *LabelUpdate {
		lu.mutation.ClearGroup()
		return lu
	}
	

	
	
	// ClearItems clears all "items" edges to the Item entity.
	func (lu *LabelUpdate) ClearItems() *LabelUpdate {
		lu.mutation.ClearItems()
		return lu
	}
	
		
		
		
		
		// RemoveItemIDs removes the "items" edge to Item entities by IDs.
		func (lu *LabelUpdate) RemoveItemIDs(ids ...uuid.UUID) *LabelUpdate {
			lu.mutation.RemoveItemIDs(ids...)
			return lu
		}
		
		// RemoveItems removes "items" edges to Item entities.
		func (lu *LabelUpdate) RemoveItems(i ...*Item) *LabelUpdate {
			ids := make([]uuid.UUID, len(i))
			for j := range i {
				ids[j] = i[j].ID
			}
			return lu.RemoveItemIDs(ids...)
		}
	




// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LabelUpdate) Save(ctx context.Context) (int, error) {
			lu.defaults()
	return withHooks[int, LabelMutation](ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LabelUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LabelUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LabelUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}


	







	// defaults sets the default values of the builder before save.
	func (lu *LabelUpdate) defaults() {
				if _, ok := lu.mutation.UpdatedAt(); !ok  {
					v := label.UpdateDefaultUpdatedAt()
					lu.mutation.SetUpdatedAt(v)
				}
	}



	// check runs all checks and user-defined validators on the builder.
	func (lu *LabelUpdate) check() error {
				if v, ok := lu.mutation.Name(); ok {
					if err := label.NameValidator(v); err != nil {
						return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Label.name": %w`, err)}
					}
				}
				if v, ok := lu.mutation.Description(); ok {
					if err := label.DescriptionValidator(v); err != nil {
						return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Label.description": %w`, err)}
					}
				}
				if v, ok := lu.mutation.Color(); ok {
					if err := label.ColorValidator(v); err != nil {
						return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "Label.color": %w`, err)}
					}
				}
				if _, ok := lu.mutation.GroupID(); lu.mutation.GroupCleared() && !ok {
					return errors.New(`ent: clearing a required unique edge "Label.group"`)
				}
		return nil
	}






	
	





    


func (lu *LabelUpdate) sqlSave(ctx context.Context) (n int, err error) {
		if err := lu.check(); err != nil {
			return n, err
		}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table: label.Table,
			Columns: label.Columns,
				ID: &sqlgraph.FieldSpec{
					Type: field.TypeUUID,
					Column: label.FieldID,
				},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := lu.mutation.UpdatedAt(); ok {
					_spec.SetField(label.FieldUpdatedAt, field.TypeTime, value)
				}
				if value, ok := lu.mutation.Name(); ok {
					_spec.SetField(label.FieldName, field.TypeString, value)
				}
				if value, ok := lu.mutation.Description(); ok {
					_spec.SetField(label.FieldDescription, field.TypeString, value)
				}
				if lu.mutation.DescriptionCleared() {
					_spec.ClearField(label.FieldDescription, field.TypeString)
				}
				if value, ok := lu.mutation.Color(); ok {
					_spec.SetField(label.FieldColor, field.TypeString, value)
				}
				if lu.mutation.ColorCleared() {
					_spec.ClearField(label.FieldColor, field.TypeString)
				}
		if lu.mutation.GroupCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: label.GroupTable,
		Columns: []string{ label.GroupColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: group.FieldID,
			},
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := lu.mutation.GroupIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: label.GroupTable,
		Columns: []string{ label.GroupColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: group.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if lu.mutation.ItemsCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: false,
		Table: label.ItemsTable,
		Columns: label.ItemsPrimaryKey,
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
			if nodes := lu.mutation.RemovedItemsIDs(); len(nodes) > 0 && !lu.mutation.ItemsCleared() {
					edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: false,
		Table: label.ItemsTable,
		Columns: label.ItemsPrimaryKey,
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
				_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
			}
		if nodes := lu.mutation.ItemsIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: false,
		Table: label.ItemsTable,
		Columns: label.ItemsPrimaryKey,
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
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}







// LabelUpdateOne is the builder for updating a single Label entity.
type LabelUpdateOne struct {
	config
	fields []string
	hooks []Hook
	mutation *LabelMutation

}


	




	
	


	
	
	// SetUpdatedAt sets the "updated_at" field.
	func (luo *LabelUpdateOne) SetUpdatedAt(t time.Time) *LabelUpdateOne {
		luo.mutation.SetUpdatedAt(t)
		return luo
	}

	
	

	

	

	

	
	
	// SetName sets the "name" field.
	func (luo *LabelUpdateOne) SetName(s string) *LabelUpdateOne {
		luo.mutation.SetName(s)
		return luo
	}

	
	

	

	

	

	
	
	// SetDescription sets the "description" field.
	func (luo *LabelUpdateOne) SetDescription(s string) *LabelUpdateOne {
		luo.mutation.SetDescription(s)
		return luo
	}

	
	
		
		// SetNillableDescription sets the "description" field if the given value is not nil.
		func (luo *LabelUpdateOne) SetNillableDescription(s *string) *LabelUpdateOne {
			if s != nil {
				luo.SetDescription(*s)
			}
			return luo
		}
	

	

	

	
		
		// ClearDescription clears the value of the "description" field.
		func (luo *LabelUpdateOne) ClearDescription() *LabelUpdateOne {
			luo.mutation.ClearDescription()
			return luo
		}
	

	
	
	// SetColor sets the "color" field.
	func (luo *LabelUpdateOne) SetColor(s string) *LabelUpdateOne {
		luo.mutation.SetColor(s)
		return luo
	}

	
	
		
		// SetNillableColor sets the "color" field if the given value is not nil.
		func (luo *LabelUpdateOne) SetNillableColor(s *string) *LabelUpdateOne {
			if s != nil {
				luo.SetColor(*s)
			}
			return luo
		}
	

	

	

	
		
		// ClearColor clears the value of the "color" field.
		func (luo *LabelUpdateOne) ClearColor() *LabelUpdateOne {
			luo.mutation.ClearColor()
			return luo
		}
	



	
	
	
	
	
		// SetGroupID sets the "group" edge to the Group entity by ID.
		func (luo *LabelUpdateOne) SetGroupID(id uuid.UUID) *LabelUpdateOne {
			luo.mutation.SetGroupID(id)
			return luo
		}
	
	
	
	
	
	// SetGroup sets the "group" edge to the Group entity.
	func (luo *LabelUpdateOne) SetGroup(g *Group) *LabelUpdateOne {
		return luo.SetGroupID(g.ID)
	}

	
	
	
	
	
		// AddItemIDs adds the "items" edge to the Item entity by IDs.
		func (luo *LabelUpdateOne) AddItemIDs(ids ... uuid.UUID) *LabelUpdateOne {
			luo.mutation.AddItemIDs(ids ...)
			return luo
		}
	
	
	
	
	
	// AddItems adds the "items" edges to the Item entity.
	func (luo *LabelUpdateOne) AddItems(i ...*Item) *LabelUpdateOne {
		ids := make([]uuid.UUID, len(i))
			for j := range i {
				ids[j] = i[j].ID
			}
			return luo.AddItemIDs(ids...)
	}


// Mutation returns the LabelMutation object of the builder.
func (luo *LabelUpdateOne) Mutation() *LabelMutation {
	return luo.mutation
}






	





	
	
	// ClearGroup clears the "group" edge to the Group entity.
	func (luo *LabelUpdateOne) ClearGroup() *LabelUpdateOne {
		luo.mutation.ClearGroup()
		return luo
	}
	

	
	
	// ClearItems clears all "items" edges to the Item entity.
	func (luo *LabelUpdateOne) ClearItems() *LabelUpdateOne {
		luo.mutation.ClearItems()
		return luo
	}
	
		
		
		
		
		// RemoveItemIDs removes the "items" edge to Item entities by IDs.
		func (luo *LabelUpdateOne) RemoveItemIDs(ids ...uuid.UUID) *LabelUpdateOne {
			luo.mutation.RemoveItemIDs(ids...)
			return luo
		}
		
		// RemoveItems removes "items" edges to Item entities.
		func (luo *LabelUpdateOne) RemoveItems(i ...*Item) *LabelUpdateOne {
			ids := make([]uuid.UUID, len(i))
			for j := range i {
				ids[j] = i[j].ID
			}
			return luo.RemoveItemIDs(ids...)
		}
	




// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LabelUpdateOne) Select(field string, fields ...string) *LabelUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Label entity.
func (luo *LabelUpdateOne ) Save(ctx context.Context) (*Label, error) {
			luo.defaults()
	return withHooks[*Label, LabelMutation](ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LabelUpdateOne) SaveX(ctx context.Context) *Label {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LabelUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LabelUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}


	







	// defaults sets the default values of the builder before save.
	func (luo *LabelUpdateOne) defaults() {
				if _, ok := luo.mutation.UpdatedAt(); !ok  {
					v := label.UpdateDefaultUpdatedAt()
					luo.mutation.SetUpdatedAt(v)
				}
	}



	// check runs all checks and user-defined validators on the builder.
	func (luo *LabelUpdateOne) check() error {
				if v, ok := luo.mutation.Name(); ok {
					if err := label.NameValidator(v); err != nil {
						return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Label.name": %w`, err)}
					}
				}
				if v, ok := luo.mutation.Description(); ok {
					if err := label.DescriptionValidator(v); err != nil {
						return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Label.description": %w`, err)}
					}
				}
				if v, ok := luo.mutation.Color(); ok {
					if err := label.ColorValidator(v); err != nil {
						return &ValidationError{Name: "color", err: fmt.Errorf(`ent: validator failed for field "Label.color": %w`, err)}
					}
				}
				if _, ok := luo.mutation.GroupID(); luo.mutation.GroupCleared() && !ok {
					return errors.New(`ent: clearing a required unique edge "Label.group"`)
				}
		return nil
	}






	
	





    


func (luo *LabelUpdateOne) sqlSave(ctx context.Context) (_node *Label, err error) {
		if err := luo.check(); err != nil {
			return _node, err
		}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table: label.Table,
			Columns: label.Columns,
				ID: &sqlgraph.FieldSpec{
					Type: field.TypeUUID,
					Column: label.FieldID,
				},
		},
	}
			id, ok := luo.mutation.ID()
			if !ok {
				return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Label.id" for update`)}
			}
			_spec.Node.ID.Value = id
			if fields := luo.fields; len(fields) > 0 {
				_spec.Node.Columns = make([]string, 0, len(fields))
				_spec.Node.Columns = append(_spec.Node.Columns, label.FieldID)
				for _, f := range fields {
					if !label.ValidColumn(f) {
						return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
					}
					if f != label.FieldID {
						_spec.Node.Columns = append(_spec.Node.Columns, f)
					}
				}
			}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
				if value, ok := luo.mutation.UpdatedAt(); ok {
					_spec.SetField(label.FieldUpdatedAt, field.TypeTime, value)
				}
				if value, ok := luo.mutation.Name(); ok {
					_spec.SetField(label.FieldName, field.TypeString, value)
				}
				if value, ok := luo.mutation.Description(); ok {
					_spec.SetField(label.FieldDescription, field.TypeString, value)
				}
				if luo.mutation.DescriptionCleared() {
					_spec.ClearField(label.FieldDescription, field.TypeString)
				}
				if value, ok := luo.mutation.Color(); ok {
					_spec.SetField(label.FieldColor, field.TypeString, value)
				}
				if luo.mutation.ColorCleared() {
					_spec.ClearField(label.FieldColor, field.TypeString)
				}
		if luo.mutation.GroupCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: label.GroupTable,
		Columns: []string{ label.GroupColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: group.FieldID,
			},
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
		if nodes := luo.mutation.GroupIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2O,
		Inverse: true,
		Table: label.GroupTable,
		Columns: []string{ label.GroupColumn },
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: group.FieldID,
			},
		},
	}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		if luo.mutation.ItemsCleared() {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: false,
		Table: label.ItemsTable,
		Columns: label.ItemsPrimaryKey,
		Bidi: false,
		Target: &sqlgraph.EdgeTarget{
			IDSpec: &sqlgraph.FieldSpec{
				Type: field.TypeUUID,
				Column: item.FieldID,
			},
		},
	}
			_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
		}
			if nodes := luo.mutation.RemovedItemsIDs(); len(nodes) > 0 && !luo.mutation.ItemsCleared() {
					edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: false,
		Table: label.ItemsTable,
		Columns: label.ItemsPrimaryKey,
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
				_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
			}
		if nodes := luo.mutation.ItemsIDs(); len(nodes) > 0 {
				edge := &sqlgraph.EdgeSpec{
		Rel: sqlgraph.M2M,
		Inverse: false,
		Table: label.ItemsTable,
		Columns: label.ItemsPrimaryKey,
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
			_spec.Edges.Add = append(_spec.Edges.Add, edge)
		}
		_node = &Label{config: luo.config}
		_spec.Assign = _node.assignValues
		_spec.ScanValues = _node.scanValues
		if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ label.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}



