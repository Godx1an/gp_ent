// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Godx1an/gp_ent/pkg/ent_work/fitnesstestitem"
	"github.com/Godx1an/gp_ent/pkg/ent_work/predicate"
)

// FitnessTestItemDelete is the builder for deleting a FitnessTestItem entity.
type FitnessTestItemDelete struct {
	config
	hooks    []Hook
	mutation *FitnessTestItemMutation
}

// Where appends a list predicates to the FitnessTestItemDelete builder.
func (ftid *FitnessTestItemDelete) Where(ps ...predicate.FitnessTestItem) *FitnessTestItemDelete {
	ftid.mutation.Where(ps...)
	return ftid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftid *FitnessTestItemDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ftid.sqlExec, ftid.mutation, ftid.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftid *FitnessTestItemDelete) ExecX(ctx context.Context) int {
	n, err := ftid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftid *FitnessTestItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(fitnesstestitem.Table, sqlgraph.NewFieldSpec(fitnesstestitem.FieldID, field.TypeInt64))
	if ps := ftid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ftid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ftid.mutation.done = true
	return affected, err
}

// FitnessTestItemDeleteOne is the builder for deleting a single FitnessTestItem entity.
type FitnessTestItemDeleteOne struct {
	ftid *FitnessTestItemDelete
}

// Where appends a list predicates to the FitnessTestItemDelete builder.
func (ftido *FitnessTestItemDeleteOne) Where(ps ...predicate.FitnessTestItem) *FitnessTestItemDeleteOne {
	ftido.ftid.mutation.Where(ps...)
	return ftido
}

// Exec executes the deletion query.
func (ftido *FitnessTestItemDeleteOne) Exec(ctx context.Context) error {
	n, err := ftido.ftid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fitnesstestitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftido *FitnessTestItemDeleteOne) ExecX(ctx context.Context) {
	if err := ftido.Exec(ctx); err != nil {
		panic(err)
	}
}
