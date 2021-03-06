// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/historytaking"
	"github.com/team10/app/ent/patientrecord"
)

// PatientrecordCreate is the builder for creating a Patientrecord entity.
type PatientrecordCreate struct {
	config
	mutation *PatientrecordMutation
	hooks    []Hook
}

// SetName sets the Name field.
func (pc *PatientrecordCreate) SetName(s string) *PatientrecordCreate {
	pc.mutation.SetName(s)
	return pc
}

// AddHistorytakingIDs adds the historytaking edge to Historytaking by ids.
func (pc *PatientrecordCreate) AddHistorytakingIDs(ids ...int) *PatientrecordCreate {
	pc.mutation.AddHistorytakingIDs(ids...)
	return pc
}

// AddHistorytaking adds the historytaking edges to Historytaking.
func (pc *PatientrecordCreate) AddHistorytaking(h ...*Historytaking) *PatientrecordCreate {
	ids := make([]int, len(h))
	for i := range h {
		ids[i] = h[i].ID
	}
	return pc.AddHistorytakingIDs(ids...)
}

// Mutation returns the PatientrecordMutation object of the builder.
func (pc *PatientrecordCreate) Mutation() *PatientrecordMutation {
	return pc.mutation
}

// Save creates the Patientrecord in the database.
func (pc *PatientrecordCreate) Save(ctx context.Context) (*Patientrecord, error) {
	if _, ok := pc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "Name", err: errors.New("ent: missing required field \"Name\"")}
	}
	var (
		err  error
		node *Patientrecord
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PatientrecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PatientrecordCreate) SaveX(ctx context.Context) *Patientrecord {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PatientrecordCreate) sqlSave(ctx context.Context) (*Patientrecord, error) {
	pa, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pa.ID = int(id)
	return pa, nil
}

func (pc *PatientrecordCreate) createSpec() (*Patientrecord, *sqlgraph.CreateSpec) {
	var (
		pa    = &Patientrecord{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: patientrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: patientrecord.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: patientrecord.FieldName,
		})
		pa.Name = value
	}
	if nodes := pc.mutation.HistorytakingIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return pa, _spec
}
