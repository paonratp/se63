// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/predicate"
	"github.com/team10/app/ent/typetreatment"
)

// TypetreatmentDelete is the builder for deleting a Typetreatment entity.
type TypetreatmentDelete struct {
	config
	hooks      []Hook
	mutation   *TypetreatmentMutation
	predicates []predicate.Typetreatment
}

// Where adds a new predicate to the delete builder.
func (td *TypetreatmentDelete) Where(ps ...predicate.Typetreatment) *TypetreatmentDelete {
	td.predicates = append(td.predicates, ps...)
	return td
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (td *TypetreatmentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(td.hooks) == 0 {
		affected, err = td.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TypetreatmentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			td.mutation = mutation
			affected, err = td.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(td.hooks) - 1; i >= 0; i-- {
			mut = td.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, td.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (td *TypetreatmentDelete) ExecX(ctx context.Context) int {
	n, err := td.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (td *TypetreatmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: typetreatment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: typetreatment.FieldID,
			},
		},
	}
	if ps := td.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, td.driver, _spec)
}

// TypetreatmentDeleteOne is the builder for deleting a single Typetreatment entity.
type TypetreatmentDeleteOne struct {
	td *TypetreatmentDelete
}

// Exec executes the deletion query.
func (tdo *TypetreatmentDeleteOne) Exec(ctx context.Context) error {
	n, err := tdo.td.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{typetreatment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (tdo *TypetreatmentDeleteOne) ExecX(ctx context.Context) {
	tdo.td.ExecX(ctx)
}
