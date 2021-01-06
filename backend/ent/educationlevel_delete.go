// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/team10/app/ent/educationlevel"
	"github.com/team10/app/ent/predicate"
)

// EducationlevelDelete is the builder for deleting a Educationlevel entity.
type EducationlevelDelete struct {
	config
	hooks      []Hook
	mutation   *EducationlevelMutation
	predicates []predicate.Educationlevel
}

// Where adds a new predicate to the delete builder.
func (ed *EducationlevelDelete) Where(ps ...predicate.Educationlevel) *EducationlevelDelete {
	ed.predicates = append(ed.predicates, ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EducationlevelDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ed.hooks) == 0 {
		affected, err = ed.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EducationlevelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ed.mutation = mutation
			affected, err = ed.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ed.hooks) - 1; i >= 0; i-- {
			mut = ed.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ed.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EducationlevelDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EducationlevelDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: educationlevel.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: educationlevel.FieldID,
			},
		},
	}
	if ps := ed.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
}

// EducationlevelDeleteOne is the builder for deleting a single Educationlevel entity.
type EducationlevelDeleteOne struct {
	ed *EducationlevelDelete
}

// Exec executes the deletion query.
func (edo *EducationlevelDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{educationlevel.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EducationlevelDeleteOne) ExecX(ctx context.Context) {
	edo.ed.ExecX(ctx)
}
