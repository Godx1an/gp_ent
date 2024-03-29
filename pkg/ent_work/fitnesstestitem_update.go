// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Godx1an/gp_ent/pkg/ent_work/fitnesstestitem"
	"github.com/Godx1an/gp_ent/pkg/ent_work/predicate"
)

// FitnessTestItemUpdate is the builder for updating FitnessTestItem entities.
type FitnessTestItemUpdate struct {
	config
	hooks     []Hook
	mutation  *FitnessTestItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FitnessTestItemUpdate builder.
func (ftiu *FitnessTestItemUpdate) Where(ps ...predicate.FitnessTestItem) *FitnessTestItemUpdate {
	ftiu.mutation.Where(ps...)
	return ftiu
}

// SetCreatedBy sets the "created_by" field.
func (ftiu *FitnessTestItemUpdate) SetCreatedBy(i int64) *FitnessTestItemUpdate {
	ftiu.mutation.ResetCreatedBy()
	ftiu.mutation.SetCreatedBy(i)
	return ftiu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ftiu *FitnessTestItemUpdate) SetNillableCreatedBy(i *int64) *FitnessTestItemUpdate {
	if i != nil {
		ftiu.SetCreatedBy(*i)
	}
	return ftiu
}

// AddCreatedBy adds i to the "created_by" field.
func (ftiu *FitnessTestItemUpdate) AddCreatedBy(i int64) *FitnessTestItemUpdate {
	ftiu.mutation.AddCreatedBy(i)
	return ftiu
}

// SetUpdatedBy sets the "updated_by" field.
func (ftiu *FitnessTestItemUpdate) SetUpdatedBy(i int64) *FitnessTestItemUpdate {
	ftiu.mutation.ResetUpdatedBy()
	ftiu.mutation.SetUpdatedBy(i)
	return ftiu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ftiu *FitnessTestItemUpdate) SetNillableUpdatedBy(i *int64) *FitnessTestItemUpdate {
	if i != nil {
		ftiu.SetUpdatedBy(*i)
	}
	return ftiu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ftiu *FitnessTestItemUpdate) AddUpdatedBy(i int64) *FitnessTestItemUpdate {
	ftiu.mutation.AddUpdatedBy(i)
	return ftiu
}

// SetUpdatedAt sets the "updated_at" field.
func (ftiu *FitnessTestItemUpdate) SetUpdatedAt(t time.Time) *FitnessTestItemUpdate {
	ftiu.mutation.SetUpdatedAt(t)
	return ftiu
}

// SetDeletedAt sets the "deleted_at" field.
func (ftiu *FitnessTestItemUpdate) SetDeletedAt(t time.Time) *FitnessTestItemUpdate {
	ftiu.mutation.SetDeletedAt(t)
	return ftiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ftiu *FitnessTestItemUpdate) SetNillableDeletedAt(t *time.Time) *FitnessTestItemUpdate {
	if t != nil {
		ftiu.SetDeletedAt(*t)
	}
	return ftiu
}

// SetItem sets the "item" field.
func (ftiu *FitnessTestItemUpdate) SetItem(s string) *FitnessTestItemUpdate {
	ftiu.mutation.SetItem(s)
	return ftiu
}

// SetNillableItem sets the "item" field if the given value is not nil.
func (ftiu *FitnessTestItemUpdate) SetNillableItem(s *string) *FitnessTestItemUpdate {
	if s != nil {
		ftiu.SetItem(*s)
	}
	return ftiu
}

// Mutation returns the FitnessTestItemMutation object of the builder.
func (ftiu *FitnessTestItemUpdate) Mutation() *FitnessTestItemMutation {
	return ftiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftiu *FitnessTestItemUpdate) Save(ctx context.Context) (int, error) {
	ftiu.defaults()
	return withHooks(ctx, ftiu.sqlSave, ftiu.mutation, ftiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftiu *FitnessTestItemUpdate) SaveX(ctx context.Context) int {
	affected, err := ftiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftiu *FitnessTestItemUpdate) Exec(ctx context.Context) error {
	_, err := ftiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftiu *FitnessTestItemUpdate) ExecX(ctx context.Context) {
	if err := ftiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftiu *FitnessTestItemUpdate) defaults() {
	if _, ok := ftiu.mutation.UpdatedAt(); !ok {
		v := fitnesstestitem.UpdateDefaultUpdatedAt()
		ftiu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ftiu *FitnessTestItemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FitnessTestItemUpdate {
	ftiu.modifiers = append(ftiu.modifiers, modifiers...)
	return ftiu
}

func (ftiu *FitnessTestItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(fitnesstestitem.Table, fitnesstestitem.Columns, sqlgraph.NewFieldSpec(fitnesstestitem.FieldID, field.TypeInt64))
	if ps := ftiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftiu.mutation.CreatedBy(); ok {
		_spec.SetField(fitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(fitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiu.mutation.UpdatedBy(); ok {
		_spec.SetField(fitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(fitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiu.mutation.UpdatedAt(); ok {
		_spec.SetField(fitnesstestitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftiu.mutation.DeletedAt(); ok {
		_spec.SetField(fitnesstestitem.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ftiu.mutation.Item(); ok {
		_spec.SetField(fitnesstestitem.FieldItem, field.TypeString, value)
	}
	_spec.AddModifiers(ftiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, ftiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fitnesstestitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ftiu.mutation.done = true
	return n, nil
}

// FitnessTestItemUpdateOne is the builder for updating a single FitnessTestItem entity.
type FitnessTestItemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FitnessTestItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (ftiuo *FitnessTestItemUpdateOne) SetCreatedBy(i int64) *FitnessTestItemUpdateOne {
	ftiuo.mutation.ResetCreatedBy()
	ftiuo.mutation.SetCreatedBy(i)
	return ftiuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ftiuo *FitnessTestItemUpdateOne) SetNillableCreatedBy(i *int64) *FitnessTestItemUpdateOne {
	if i != nil {
		ftiuo.SetCreatedBy(*i)
	}
	return ftiuo
}

// AddCreatedBy adds i to the "created_by" field.
func (ftiuo *FitnessTestItemUpdateOne) AddCreatedBy(i int64) *FitnessTestItemUpdateOne {
	ftiuo.mutation.AddCreatedBy(i)
	return ftiuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ftiuo *FitnessTestItemUpdateOne) SetUpdatedBy(i int64) *FitnessTestItemUpdateOne {
	ftiuo.mutation.ResetUpdatedBy()
	ftiuo.mutation.SetUpdatedBy(i)
	return ftiuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ftiuo *FitnessTestItemUpdateOne) SetNillableUpdatedBy(i *int64) *FitnessTestItemUpdateOne {
	if i != nil {
		ftiuo.SetUpdatedBy(*i)
	}
	return ftiuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ftiuo *FitnessTestItemUpdateOne) AddUpdatedBy(i int64) *FitnessTestItemUpdateOne {
	ftiuo.mutation.AddUpdatedBy(i)
	return ftiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ftiuo *FitnessTestItemUpdateOne) SetUpdatedAt(t time.Time) *FitnessTestItemUpdateOne {
	ftiuo.mutation.SetUpdatedAt(t)
	return ftiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ftiuo *FitnessTestItemUpdateOne) SetDeletedAt(t time.Time) *FitnessTestItemUpdateOne {
	ftiuo.mutation.SetDeletedAt(t)
	return ftiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ftiuo *FitnessTestItemUpdateOne) SetNillableDeletedAt(t *time.Time) *FitnessTestItemUpdateOne {
	if t != nil {
		ftiuo.SetDeletedAt(*t)
	}
	return ftiuo
}

// SetItem sets the "item" field.
func (ftiuo *FitnessTestItemUpdateOne) SetItem(s string) *FitnessTestItemUpdateOne {
	ftiuo.mutation.SetItem(s)
	return ftiuo
}

// SetNillableItem sets the "item" field if the given value is not nil.
func (ftiuo *FitnessTestItemUpdateOne) SetNillableItem(s *string) *FitnessTestItemUpdateOne {
	if s != nil {
		ftiuo.SetItem(*s)
	}
	return ftiuo
}

// Mutation returns the FitnessTestItemMutation object of the builder.
func (ftiuo *FitnessTestItemUpdateOne) Mutation() *FitnessTestItemMutation {
	return ftiuo.mutation
}

// Where appends a list predicates to the FitnessTestItemUpdate builder.
func (ftiuo *FitnessTestItemUpdateOne) Where(ps ...predicate.FitnessTestItem) *FitnessTestItemUpdateOne {
	ftiuo.mutation.Where(ps...)
	return ftiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftiuo *FitnessTestItemUpdateOne) Select(field string, fields ...string) *FitnessTestItemUpdateOne {
	ftiuo.fields = append([]string{field}, fields...)
	return ftiuo
}

// Save executes the query and returns the updated FitnessTestItem entity.
func (ftiuo *FitnessTestItemUpdateOne) Save(ctx context.Context) (*FitnessTestItem, error) {
	ftiuo.defaults()
	return withHooks(ctx, ftiuo.sqlSave, ftiuo.mutation, ftiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ftiuo *FitnessTestItemUpdateOne) SaveX(ctx context.Context) *FitnessTestItem {
	node, err := ftiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftiuo *FitnessTestItemUpdateOne) Exec(ctx context.Context) error {
	_, err := ftiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftiuo *FitnessTestItemUpdateOne) ExecX(ctx context.Context) {
	if err := ftiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ftiuo *FitnessTestItemUpdateOne) defaults() {
	if _, ok := ftiuo.mutation.UpdatedAt(); !ok {
		v := fitnesstestitem.UpdateDefaultUpdatedAt()
		ftiuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (ftiuo *FitnessTestItemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FitnessTestItemUpdateOne {
	ftiuo.modifiers = append(ftiuo.modifiers, modifiers...)
	return ftiuo
}

func (ftiuo *FitnessTestItemUpdateOne) sqlSave(ctx context.Context) (_node *FitnessTestItem, err error) {
	_spec := sqlgraph.NewUpdateSpec(fitnesstestitem.Table, fitnesstestitem.Columns, sqlgraph.NewFieldSpec(fitnesstestitem.FieldID, field.TypeInt64))
	id, ok := ftiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent_work: missing "FitnessTestItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ftiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fitnesstestitem.FieldID)
		for _, f := range fields {
			if !fitnesstestitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent_work: invalid field %q for query", f)}
			}
			if f != fitnesstestitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ftiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ftiuo.mutation.CreatedBy(); ok {
		_spec.SetField(fitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(fitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiuo.mutation.UpdatedBy(); ok {
		_spec.SetField(fitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(fitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ftiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(fitnesstestitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := ftiuo.mutation.DeletedAt(); ok {
		_spec.SetField(fitnesstestitem.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := ftiuo.mutation.Item(); ok {
		_spec.SetField(fitnesstestitem.FieldItem, field.TypeString, value)
	}
	_spec.AddModifiers(ftiuo.modifiers...)
	_node = &FitnessTestItem{config: ftiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ftiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fitnesstestitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ftiuo.mutation.done = true
	return _node, nil
}
