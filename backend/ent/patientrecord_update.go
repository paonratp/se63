// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/patientrecord"
	"github.com/team10/app/ent/predicate"
)

// PatientrecordUpdate is the builder for updating Patientrecord entities.
type PatientrecordUpdate struct {
	config
	hooks      []Hook
	mutation   *PatientrecordMutation
	predicates []predicate.Patientrecord
}

// Where adds a new predicate for the builder.
func (pu *PatientrecordUpdate) Where(ps ...predicate.Patientrecord) *PatientrecordUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the Name field.
func (pu *PatientrecordUpdate) SetName(s string) *PatientrecordUpdate {
	pu.mutation.SetName(s)
	return pu
}

// AddHistorytakingIDs adds the historytaking edge to Historytaking by ids.
func (pu *PatientrecordUpdate) AddHistorytakingIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.AddHistorytakingIDs(ids...)
	return pu
}

// AddHistorytaking adds the historytaking edges to Historytaking.
func (pu *PatientrecordUpdate) AddHistorytaking(h ...*Historytaking) *PatientrecordUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pu.AddHistorytakingIDs(ids...)
}

// Mutation returns the PatientrecordMutation object of the builder.
func (pu *PatientrecordUpdate) Mutation() *PatientrecordMutation {
	return pu.mutation
}

// RemoveHistorytakingIDs removes the historytaking edge to Historytaking by ids.
func (pu *PatientrecordUpdate) RemoveHistorytakingIDs(ids ...int) *PatientrecordUpdate {
	pu.mutation.RemoveHistorytakingIDs(ids...)
	return pu
}

// RemoveHistorytaking removes historytaking edges to Historytaking.
func (pu *PatientrecordUpdate) RemoveHistorytaking(h ...*Historytaking) *PatientrecordUpdate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pu.RemoveHistorytakingIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PatientrecordUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PatientrecordUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PatientrecordUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PatientrecordUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PatientrecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientrecord.Table,
			Columns: patientrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrecord.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldName,
		})
	}
	if nodes := pu.mutation.RemovedHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.HistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patientrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PatientrecordUpdateOne is the builder for updating a single Patientrecord entity.
type PatientrecordUpdateOne struct {
	config
	hooks    []Hook
	mutation *PatientrecordMutation
}

// SetName sets the Name field.
func (puo *PatientrecordUpdateOne) SetName(s string) *PatientrecordUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// AddHistorytakingIDs adds the historytaking edge to Historytaking by ids.
func (puo *PatientrecordUpdateOne) AddHistorytakingIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.AddHistorytakingIDs(ids...)
	return puo
}

// AddHistorytaking adds the historytaking edges to Historytaking.
func (puo *PatientrecordUpdateOne) AddHistorytaking(h ...*Historytaking) *PatientrecordUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return puo.AddHistorytakingIDs(ids...)
}

// Mutation returns the PatientrecordMutation object of the builder.
func (puo *PatientrecordUpdateOne) Mutation() *PatientrecordMutation {
	return puo.mutation
}

// RemoveHistorytakingIDs removes the historytaking edge to Historytaking by ids.
func (puo *PatientrecordUpdateOne) RemoveHistorytakingIDs(ids ...int) *PatientrecordUpdateOne {
	puo.mutation.RemoveHistorytakingIDs(ids...)
	return puo
}

// RemoveHistorytaking removes historytaking edges to Historytaking.
func (puo *PatientrecordUpdateOne) RemoveHistorytaking(h ...*Historytaking) *PatientrecordUpdateOne {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return puo.RemoveHistorytakingIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (puo *PatientrecordUpdateOne) Save(ctx context.Context) (*Patientrecord, error) {

	var (
		err  error
		node *Patientrecord
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PatientrecordUpdateOne) SaveX(ctx context.Context) *Patientrecord {
	pa, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pa
}

// Exec executes the query on the entity.
func (puo *PatientrecordUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PatientrecordUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PatientrecordUpdateOne) sqlSave(ctx context.Context) (pa *Patientrecord, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   patientrecord.Table,
			Columns: patientrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrecord.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Patientrecord.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldName,
		})
	}
	if nodes := puo.mutation.RemovedHistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.HistorytakingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   patientrecord.HistorytakingTable,
			Columns: []string{patientrecord.HistorytakingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: historytaking.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pa = &Patientrecord{config: puo.config}
	_spec.Assign = pa.assignValues
	_spec.ScanValues = pa.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patientrecord.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pa, nil
}
