// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"repro/ent/predicate"
	"repro/ent/role"
	"repro/ent/user"
	"repro/ent/userrole"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserRoleUpdate is the builder for updating UserRole entities.
type UserRoleUpdate struct {
	config
	hooks    []Hook
	mutation *UserRoleMutation
}

// Where adds a new predicate for the UserRoleUpdate builder.
func (uru *UserRoleUpdate) Where(ps ...predicate.UserRole) *UserRoleUpdate {
	uru.mutation.predicates = append(uru.mutation.predicates, ps...)
	return uru
}

// SetTargetID sets the "target_id" field.
func (uru *UserRoleUpdate) SetTargetID(i int) *UserRoleUpdate {
	uru.mutation.ResetTargetID()
	uru.mutation.SetTargetID(i)
	return uru
}

// AddTargetID adds i to the "target_id" field.
func (uru *UserRoleUpdate) AddTargetID(i int) *UserRoleUpdate {
	uru.mutation.AddTargetID(i)
	return uru
}

// SetTargetType sets the "target_type" field.
func (uru *UserRoleUpdate) SetTargetType(s string) *UserRoleUpdate {
	uru.mutation.SetTargetType(s)
	return uru
}

// SetActorType sets the "actor_type" field.
func (uru *UserRoleUpdate) SetActorType(s string) *UserRoleUpdate {
	uru.mutation.SetActorType(s)
	return uru
}

// SetActorID sets the "actor_id" field.
func (uru *UserRoleUpdate) SetActorID(i int) *UserRoleUpdate {
	uru.mutation.ResetActorID()
	uru.mutation.SetActorID(i)
	return uru
}

// AddActorID adds i to the "actor_id" field.
func (uru *UserRoleUpdate) AddActorID(i int) *UserRoleUpdate {
	uru.mutation.AddActorID(i)
	return uru
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (uru *UserRoleUpdate) SetRoleID(id int) *UserRoleUpdate {
	uru.mutation.SetRoleID(id)
	return uru
}

// SetNillableRoleID sets the "role" edge to the Role entity by ID if the given value is not nil.
func (uru *UserRoleUpdate) SetNillableRoleID(id *int) *UserRoleUpdate {
	if id != nil {
		uru = uru.SetRoleID(*id)
	}
	return uru
}

// SetRole sets the "role" edge to the Role entity.
func (uru *UserRoleUpdate) SetRole(r *Role) *UserRoleUpdate {
	return uru.SetRoleID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uru *UserRoleUpdate) SetUserID(id int) *UserRoleUpdate {
	uru.mutation.SetUserID(id)
	return uru
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (uru *UserRoleUpdate) SetNillableUserID(id *int) *UserRoleUpdate {
	if id != nil {
		uru = uru.SetUserID(*id)
	}
	return uru
}

// SetUser sets the "user" edge to the User entity.
func (uru *UserRoleUpdate) SetUser(u *User) *UserRoleUpdate {
	return uru.SetUserID(u.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (uru *UserRoleUpdate) Mutation() *UserRoleMutation {
	return uru.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (uru *UserRoleUpdate) ClearRole() *UserRoleUpdate {
	uru.mutation.ClearRole()
	return uru
}

// ClearUser clears the "user" edge to the User entity.
func (uru *UserRoleUpdate) ClearUser() *UserRoleUpdate {
	uru.mutation.ClearUser()
	return uru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uru *UserRoleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uru.hooks) == 0 {
		affected, err = uru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uru.mutation = mutation
			affected, err = uru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uru.hooks) - 1; i >= 0; i-- {
			mut = uru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uru *UserRoleUpdate) SaveX(ctx context.Context) int {
	affected, err := uru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uru *UserRoleUpdate) Exec(ctx context.Context) error {
	_, err := uru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uru *UserRoleUpdate) ExecX(ctx context.Context) {
	if err := uru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uru *UserRoleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userrole.Table,
			Columns: userrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userrole.FieldID,
			},
		},
	}
	if ps := uru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uru.mutation.TargetID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldTargetID,
		})
	}
	if value, ok := uru.mutation.AddedTargetID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldTargetID,
		})
	}
	if value, ok := uru.mutation.TargetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userrole.FieldTargetType,
		})
	}
	if value, ok := uru.mutation.ActorType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userrole.FieldActorType,
		})
	}
	if value, ok := uru.mutation.ActorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldActorID,
		})
	}
	if value, ok := uru.mutation.AddedActorID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldActorID,
		})
	}
	if uru.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UserRoleUpdateOne is the builder for updating a single UserRole entity.
type UserRoleUpdateOne struct {
	config
	hooks    []Hook
	mutation *UserRoleMutation
}

// SetTargetID sets the "target_id" field.
func (uruo *UserRoleUpdateOne) SetTargetID(i int) *UserRoleUpdateOne {
	uruo.mutation.ResetTargetID()
	uruo.mutation.SetTargetID(i)
	return uruo
}

// AddTargetID adds i to the "target_id" field.
func (uruo *UserRoleUpdateOne) AddTargetID(i int) *UserRoleUpdateOne {
	uruo.mutation.AddTargetID(i)
	return uruo
}

// SetTargetType sets the "target_type" field.
func (uruo *UserRoleUpdateOne) SetTargetType(s string) *UserRoleUpdateOne {
	uruo.mutation.SetTargetType(s)
	return uruo
}

// SetActorType sets the "actor_type" field.
func (uruo *UserRoleUpdateOne) SetActorType(s string) *UserRoleUpdateOne {
	uruo.mutation.SetActorType(s)
	return uruo
}

// SetActorID sets the "actor_id" field.
func (uruo *UserRoleUpdateOne) SetActorID(i int) *UserRoleUpdateOne {
	uruo.mutation.ResetActorID()
	uruo.mutation.SetActorID(i)
	return uruo
}

// AddActorID adds i to the "actor_id" field.
func (uruo *UserRoleUpdateOne) AddActorID(i int) *UserRoleUpdateOne {
	uruo.mutation.AddActorID(i)
	return uruo
}

// SetRoleID sets the "role" edge to the Role entity by ID.
func (uruo *UserRoleUpdateOne) SetRoleID(id int) *UserRoleUpdateOne {
	uruo.mutation.SetRoleID(id)
	return uruo
}

// SetNillableRoleID sets the "role" edge to the Role entity by ID if the given value is not nil.
func (uruo *UserRoleUpdateOne) SetNillableRoleID(id *int) *UserRoleUpdateOne {
	if id != nil {
		uruo = uruo.SetRoleID(*id)
	}
	return uruo
}

// SetRole sets the "role" edge to the Role entity.
func (uruo *UserRoleUpdateOne) SetRole(r *Role) *UserRoleUpdateOne {
	return uruo.SetRoleID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (uruo *UserRoleUpdateOne) SetUserID(id int) *UserRoleUpdateOne {
	uruo.mutation.SetUserID(id)
	return uruo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (uruo *UserRoleUpdateOne) SetNillableUserID(id *int) *UserRoleUpdateOne {
	if id != nil {
		uruo = uruo.SetUserID(*id)
	}
	return uruo
}

// SetUser sets the "user" edge to the User entity.
func (uruo *UserRoleUpdateOne) SetUser(u *User) *UserRoleUpdateOne {
	return uruo.SetUserID(u.ID)
}

// Mutation returns the UserRoleMutation object of the builder.
func (uruo *UserRoleUpdateOne) Mutation() *UserRoleMutation {
	return uruo.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (uruo *UserRoleUpdateOne) ClearRole() *UserRoleUpdateOne {
	uruo.mutation.ClearRole()
	return uruo
}

// ClearUser clears the "user" edge to the User entity.
func (uruo *UserRoleUpdateOne) ClearUser() *UserRoleUpdateOne {
	uruo.mutation.ClearUser()
	return uruo
}

// Save executes the query and returns the updated UserRole entity.
func (uruo *UserRoleUpdateOne) Save(ctx context.Context) (*UserRole, error) {
	var (
		err  error
		node *UserRole
	)
	if len(uruo.hooks) == 0 {
		node, err = uruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserRoleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uruo.mutation = mutation
			node, err = uruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uruo.hooks) - 1; i >= 0; i-- {
			mut = uruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uruo *UserRoleUpdateOne) SaveX(ctx context.Context) *UserRole {
	node, err := uruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uruo *UserRoleUpdateOne) Exec(ctx context.Context) error {
	_, err := uruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uruo *UserRoleUpdateOne) ExecX(ctx context.Context) {
	if err := uruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uruo *UserRoleUpdateOne) sqlSave(ctx context.Context) (_node *UserRole, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userrole.Table,
			Columns: userrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: userrole.FieldID,
			},
		},
	}
	id, ok := uruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserRole.ID for update")}
	}
	_spec.Node.ID.Value = id
	if ps := uruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uruo.mutation.TargetID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldTargetID,
		})
	}
	if value, ok := uruo.mutation.AddedTargetID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldTargetID,
		})
	}
	if value, ok := uruo.mutation.TargetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userrole.FieldTargetType,
		})
	}
	if value, ok := uruo.mutation.ActorType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userrole.FieldActorType,
		})
	}
	if value, ok := uruo.mutation.ActorID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldActorID,
		})
	}
	if value, ok := uruo.mutation.AddedActorID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userrole.FieldActorID,
		})
	}
	if uruo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.RoleTable,
			Columns: []string{userrole.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userrole.UserTable,
			Columns: []string{userrole.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserRole{config: uruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userrole.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
