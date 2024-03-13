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
	"github.com/Godx1an/gp_ent/pkg/ent_work/predicate"
	"github.com/Godx1an/gp_ent/pkg/ent_work/school"
)

// SchoolUpdate is the builder for updating School entities.
type SchoolUpdate struct {
	config
	hooks     []Hook
	mutation  *SchoolMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SchoolUpdate builder.
func (su *SchoolUpdate) Where(ps ...predicate.School) *SchoolUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetCreatedBy sets the "created_by" field.
func (su *SchoolUpdate) SetCreatedBy(i int64) *SchoolUpdate {
	su.mutation.ResetCreatedBy()
	su.mutation.SetCreatedBy(i)
	return su
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableCreatedBy(i *int64) *SchoolUpdate {
	if i != nil {
		su.SetCreatedBy(*i)
	}
	return su
}

// AddCreatedBy adds i to the "created_by" field.
func (su *SchoolUpdate) AddCreatedBy(i int64) *SchoolUpdate {
	su.mutation.AddCreatedBy(i)
	return su
}

// SetUpdatedBy sets the "updated_by" field.
func (su *SchoolUpdate) SetUpdatedBy(i int64) *SchoolUpdate {
	su.mutation.ResetUpdatedBy()
	su.mutation.SetUpdatedBy(i)
	return su
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableUpdatedBy(i *int64) *SchoolUpdate {
	if i != nil {
		su.SetUpdatedBy(*i)
	}
	return su
}

// AddUpdatedBy adds i to the "updated_by" field.
func (su *SchoolUpdate) AddUpdatedBy(i int64) *SchoolUpdate {
	su.mutation.AddUpdatedBy(i)
	return su
}

// SetUpdatedAt sets the "updated_at" field.
func (su *SchoolUpdate) SetUpdatedAt(t time.Time) *SchoolUpdate {
	su.mutation.SetUpdatedAt(t)
	return su
}

// SetDeletedAt sets the "deleted_at" field.
func (su *SchoolUpdate) SetDeletedAt(t time.Time) *SchoolUpdate {
	su.mutation.SetDeletedAt(t)
	return su
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableDeletedAt(t *time.Time) *SchoolUpdate {
	if t != nil {
		su.SetDeletedAt(*t)
	}
	return su
}

// SetName sets the "name" field.
func (su *SchoolUpdate) SetName(s string) *SchoolUpdate {
	su.mutation.SetName(s)
	return su
}

// SetNillableName sets the "name" field if the given value is not nil.
func (su *SchoolUpdate) SetNillableName(s *string) *SchoolUpdate {
	if s != nil {
		su.SetName(*s)
	}
	return su
}

// Mutation returns the SchoolMutation object of the builder.
func (su *SchoolUpdate) Mutation() *SchoolMutation {
	return su.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SchoolUpdate) Save(ctx context.Context) (int, error) {
	su.defaults()
	return withHooks(ctx, su.sqlSave, su.mutation, su.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (su *SchoolUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SchoolUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SchoolUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SchoolUpdate) defaults() {
	if _, ok := su.mutation.UpdatedAt(); !ok {
		v := school.UpdateDefaultUpdatedAt()
		su.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (su *SchoolUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SchoolUpdate {
	su.modifiers = append(su.modifiers, modifiers...)
	return su
}

func (su *SchoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(school.Table, school.Columns, sqlgraph.NewFieldSpec(school.FieldID, field.TypeInt64))
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.CreatedBy(); ok {
		_spec.SetField(school.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedCreatedBy(); ok {
		_spec.AddField(school.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := su.mutation.UpdatedBy(); ok {
		_spec.SetField(school.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := su.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(school.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := su.mutation.UpdatedAt(); ok {
		_spec.SetField(school.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.DeletedAt(); ok {
		_spec.SetField(school.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.SetField(school.FieldName, field.TypeString, value)
	}
	_spec.AddModifiers(su.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{school.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	su.mutation.done = true
	return n, nil
}

// SchoolUpdateOne is the builder for updating a single School entity.
type SchoolUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SchoolMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (suo *SchoolUpdateOne) SetCreatedBy(i int64) *SchoolUpdateOne {
	suo.mutation.ResetCreatedBy()
	suo.mutation.SetCreatedBy(i)
	return suo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableCreatedBy(i *int64) *SchoolUpdateOne {
	if i != nil {
		suo.SetCreatedBy(*i)
	}
	return suo
}

// AddCreatedBy adds i to the "created_by" field.
func (suo *SchoolUpdateOne) AddCreatedBy(i int64) *SchoolUpdateOne {
	suo.mutation.AddCreatedBy(i)
	return suo
}

// SetUpdatedBy sets the "updated_by" field.
func (suo *SchoolUpdateOne) SetUpdatedBy(i int64) *SchoolUpdateOne {
	suo.mutation.ResetUpdatedBy()
	suo.mutation.SetUpdatedBy(i)
	return suo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableUpdatedBy(i *int64) *SchoolUpdateOne {
	if i != nil {
		suo.SetUpdatedBy(*i)
	}
	return suo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (suo *SchoolUpdateOne) AddUpdatedBy(i int64) *SchoolUpdateOne {
	suo.mutation.AddUpdatedBy(i)
	return suo
}

// SetUpdatedAt sets the "updated_at" field.
func (suo *SchoolUpdateOne) SetUpdatedAt(t time.Time) *SchoolUpdateOne {
	suo.mutation.SetUpdatedAt(t)
	return suo
}

// SetDeletedAt sets the "deleted_at" field.
func (suo *SchoolUpdateOne) SetDeletedAt(t time.Time) *SchoolUpdateOne {
	suo.mutation.SetDeletedAt(t)
	return suo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableDeletedAt(t *time.Time) *SchoolUpdateOne {
	if t != nil {
		suo.SetDeletedAt(*t)
	}
	return suo
}

// SetName sets the "name" field.
func (suo *SchoolUpdateOne) SetName(s string) *SchoolUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (suo *SchoolUpdateOne) SetNillableName(s *string) *SchoolUpdateOne {
	if s != nil {
		suo.SetName(*s)
	}
	return suo
}

// Mutation returns the SchoolMutation object of the builder.
func (suo *SchoolUpdateOne) Mutation() *SchoolMutation {
	return suo.mutation
}

// Where appends a list predicates to the SchoolUpdate builder.
func (suo *SchoolUpdateOne) Where(ps ...predicate.School) *SchoolUpdateOne {
	suo.mutation.Where(ps...)
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SchoolUpdateOne) Select(field string, fields ...string) *SchoolUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated School entity.
func (suo *SchoolUpdateOne) Save(ctx context.Context) (*School, error) {
	suo.defaults()
	return withHooks(ctx, suo.sqlSave, suo.mutation, suo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SchoolUpdateOne) SaveX(ctx context.Context) *School {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SchoolUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SchoolUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SchoolUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdatedAt(); !ok {
		v := school.UpdateDefaultUpdatedAt()
		suo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (suo *SchoolUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SchoolUpdateOne {
	suo.modifiers = append(suo.modifiers, modifiers...)
	return suo
}

func (suo *SchoolUpdateOne) sqlSave(ctx context.Context) (_node *School, err error) {
	_spec := sqlgraph.NewUpdateSpec(school.Table, school.Columns, sqlgraph.NewFieldSpec(school.FieldID, field.TypeInt64))
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent_work: missing "School.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, school.FieldID)
		for _, f := range fields {
			if !school.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent_work: invalid field %q for query", f)}
			}
			if f != school.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.CreatedBy(); ok {
		_spec.SetField(school.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(school.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.UpdatedBy(); ok {
		_spec.SetField(school.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(school.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := suo.mutation.UpdatedAt(); ok {
		_spec.SetField(school.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.DeletedAt(); ok {
		_spec.SetField(school.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := suo.mutation.Name(); ok {
		_spec.SetField(school.FieldName, field.TypeString, value)
	}
	_spec.AddModifiers(suo.modifiers...)
	_node = &School{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{school.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	suo.mutation.done = true
	return _node, nil
}
