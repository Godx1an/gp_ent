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
	"github.com/Godx1an/gp_ent/pkg/ent_work/schoolfitnesstestitem"
)

// SchoolFitnessTestItemUpdate is the builder for updating SchoolFitnessTestItem entities.
type SchoolFitnessTestItemUpdate struct {
	config
	hooks     []Hook
	mutation  *SchoolFitnessTestItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SchoolFitnessTestItemUpdate builder.
func (sftiu *SchoolFitnessTestItemUpdate) Where(ps ...predicate.SchoolFitnessTestItem) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.Where(ps...)
	return sftiu
}

// SetCreatedBy sets the "created_by" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetCreatedBy(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.ResetCreatedBy()
	sftiu.mutation.SetCreatedBy(i)
	return sftiu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sftiu *SchoolFitnessTestItemUpdate) SetNillableCreatedBy(i *int64) *SchoolFitnessTestItemUpdate {
	if i != nil {
		sftiu.SetCreatedBy(*i)
	}
	return sftiu
}

// AddCreatedBy adds i to the "created_by" field.
func (sftiu *SchoolFitnessTestItemUpdate) AddCreatedBy(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.AddCreatedBy(i)
	return sftiu
}

// SetUpdatedBy sets the "updated_by" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetUpdatedBy(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.ResetUpdatedBy()
	sftiu.mutation.SetUpdatedBy(i)
	return sftiu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sftiu *SchoolFitnessTestItemUpdate) SetNillableUpdatedBy(i *int64) *SchoolFitnessTestItemUpdate {
	if i != nil {
		sftiu.SetUpdatedBy(*i)
	}
	return sftiu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (sftiu *SchoolFitnessTestItemUpdate) AddUpdatedBy(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.AddUpdatedBy(i)
	return sftiu
}

// SetUpdatedAt sets the "updated_at" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetUpdatedAt(t time.Time) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.SetUpdatedAt(t)
	return sftiu
}

// SetDeletedAt sets the "deleted_at" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetDeletedAt(t time.Time) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.SetDeletedAt(t)
	return sftiu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sftiu *SchoolFitnessTestItemUpdate) SetNillableDeletedAt(t *time.Time) *SchoolFitnessTestItemUpdate {
	if t != nil {
		sftiu.SetDeletedAt(*t)
	}
	return sftiu
}

// SetMaxParticipants sets the "max_participants" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetMaxParticipants(i int) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.ResetMaxParticipants()
	sftiu.mutation.SetMaxParticipants(i)
	return sftiu
}

// SetNillableMaxParticipants sets the "max_participants" field if the given value is not nil.
func (sftiu *SchoolFitnessTestItemUpdate) SetNillableMaxParticipants(i *int) *SchoolFitnessTestItemUpdate {
	if i != nil {
		sftiu.SetMaxParticipants(*i)
	}
	return sftiu
}

// AddMaxParticipants adds i to the "max_participants" field.
func (sftiu *SchoolFitnessTestItemUpdate) AddMaxParticipants(i int) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.AddMaxParticipants(i)
	return sftiu
}

// SetAvgTimePerPerson sets the "avg_time_per_person" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetAvgTimePerPerson(i int) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.ResetAvgTimePerPerson()
	sftiu.mutation.SetAvgTimePerPerson(i)
	return sftiu
}

// SetNillableAvgTimePerPerson sets the "avg_time_per_person" field if the given value is not nil.
func (sftiu *SchoolFitnessTestItemUpdate) SetNillableAvgTimePerPerson(i *int) *SchoolFitnessTestItemUpdate {
	if i != nil {
		sftiu.SetAvgTimePerPerson(*i)
	}
	return sftiu
}

// AddAvgTimePerPerson adds i to the "avg_time_per_person" field.
func (sftiu *SchoolFitnessTestItemUpdate) AddAvgTimePerPerson(i int) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.AddAvgTimePerPerson(i)
	return sftiu
}

// SetSchoolID sets the "school_id" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetSchoolID(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.ResetSchoolID()
	sftiu.mutation.SetSchoolID(i)
	return sftiu
}

// SetNillableSchoolID sets the "school_id" field if the given value is not nil.
func (sftiu *SchoolFitnessTestItemUpdate) SetNillableSchoolID(i *int64) *SchoolFitnessTestItemUpdate {
	if i != nil {
		sftiu.SetSchoolID(*i)
	}
	return sftiu
}

// AddSchoolID adds i to the "school_id" field.
func (sftiu *SchoolFitnessTestItemUpdate) AddSchoolID(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.AddSchoolID(i)
	return sftiu
}

// SetItemID sets the "item_id" field.
func (sftiu *SchoolFitnessTestItemUpdate) SetItemID(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.ResetItemID()
	sftiu.mutation.SetItemID(i)
	return sftiu
}

// SetNillableItemID sets the "item_id" field if the given value is not nil.
func (sftiu *SchoolFitnessTestItemUpdate) SetNillableItemID(i *int64) *SchoolFitnessTestItemUpdate {
	if i != nil {
		sftiu.SetItemID(*i)
	}
	return sftiu
}

// AddItemID adds i to the "item_id" field.
func (sftiu *SchoolFitnessTestItemUpdate) AddItemID(i int64) *SchoolFitnessTestItemUpdate {
	sftiu.mutation.AddItemID(i)
	return sftiu
}

// Mutation returns the SchoolFitnessTestItemMutation object of the builder.
func (sftiu *SchoolFitnessTestItemUpdate) Mutation() *SchoolFitnessTestItemMutation {
	return sftiu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sftiu *SchoolFitnessTestItemUpdate) Save(ctx context.Context) (int, error) {
	sftiu.defaults()
	return withHooks(ctx, sftiu.sqlSave, sftiu.mutation, sftiu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sftiu *SchoolFitnessTestItemUpdate) SaveX(ctx context.Context) int {
	affected, err := sftiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sftiu *SchoolFitnessTestItemUpdate) Exec(ctx context.Context) error {
	_, err := sftiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sftiu *SchoolFitnessTestItemUpdate) ExecX(ctx context.Context) {
	if err := sftiu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sftiu *SchoolFitnessTestItemUpdate) defaults() {
	if _, ok := sftiu.mutation.UpdatedAt(); !ok {
		v := schoolfitnesstestitem.UpdateDefaultUpdatedAt()
		sftiu.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sftiu *SchoolFitnessTestItemUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SchoolFitnessTestItemUpdate {
	sftiu.modifiers = append(sftiu.modifiers, modifiers...)
	return sftiu
}

func (sftiu *SchoolFitnessTestItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(schoolfitnesstestitem.Table, schoolfitnesstestitem.Columns, sqlgraph.NewFieldSpec(schoolfitnesstestitem.FieldID, field.TypeInt64))
	if ps := sftiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sftiu.mutation.CreatedBy(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiu.mutation.UpdatedBy(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiu.mutation.UpdatedAt(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sftiu.mutation.DeletedAt(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := sftiu.mutation.MaxParticipants(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldMaxParticipants, field.TypeInt, value)
	}
	if value, ok := sftiu.mutation.AddedMaxParticipants(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldMaxParticipants, field.TypeInt, value)
	}
	if value, ok := sftiu.mutation.AvgTimePerPerson(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldAvgTimePerPerson, field.TypeInt, value)
	}
	if value, ok := sftiu.mutation.AddedAvgTimePerPerson(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldAvgTimePerPerson, field.TypeInt, value)
	}
	if value, ok := sftiu.mutation.SchoolID(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldSchoolID, field.TypeInt64, value)
	}
	if value, ok := sftiu.mutation.AddedSchoolID(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldSchoolID, field.TypeInt64, value)
	}
	if value, ok := sftiu.mutation.ItemID(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldItemID, field.TypeInt64, value)
	}
	if value, ok := sftiu.mutation.AddedItemID(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldItemID, field.TypeInt64, value)
	}
	_spec.AddModifiers(sftiu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, sftiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schoolfitnesstestitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	sftiu.mutation.done = true
	return n, nil
}

// SchoolFitnessTestItemUpdateOne is the builder for updating a single SchoolFitnessTestItem entity.
type SchoolFitnessTestItemUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SchoolFitnessTestItemMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedBy sets the "created_by" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetCreatedBy(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.ResetCreatedBy()
	sftiuo.mutation.SetCreatedBy(i)
	return sftiuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetNillableCreatedBy(i *int64) *SchoolFitnessTestItemUpdateOne {
	if i != nil {
		sftiuo.SetCreatedBy(*i)
	}
	return sftiuo
}

// AddCreatedBy adds i to the "created_by" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) AddCreatedBy(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.AddCreatedBy(i)
	return sftiuo
}

// SetUpdatedBy sets the "updated_by" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetUpdatedBy(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.ResetUpdatedBy()
	sftiuo.mutation.SetUpdatedBy(i)
	return sftiuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetNillableUpdatedBy(i *int64) *SchoolFitnessTestItemUpdateOne {
	if i != nil {
		sftiuo.SetUpdatedBy(*i)
	}
	return sftiuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) AddUpdatedBy(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.AddUpdatedBy(i)
	return sftiuo
}

// SetUpdatedAt sets the "updated_at" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetUpdatedAt(t time.Time) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.SetUpdatedAt(t)
	return sftiuo
}

// SetDeletedAt sets the "deleted_at" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetDeletedAt(t time.Time) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.SetDeletedAt(t)
	return sftiuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetNillableDeletedAt(t *time.Time) *SchoolFitnessTestItemUpdateOne {
	if t != nil {
		sftiuo.SetDeletedAt(*t)
	}
	return sftiuo
}

// SetMaxParticipants sets the "max_participants" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetMaxParticipants(i int) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.ResetMaxParticipants()
	sftiuo.mutation.SetMaxParticipants(i)
	return sftiuo
}

// SetNillableMaxParticipants sets the "max_participants" field if the given value is not nil.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetNillableMaxParticipants(i *int) *SchoolFitnessTestItemUpdateOne {
	if i != nil {
		sftiuo.SetMaxParticipants(*i)
	}
	return sftiuo
}

// AddMaxParticipants adds i to the "max_participants" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) AddMaxParticipants(i int) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.AddMaxParticipants(i)
	return sftiuo
}

// SetAvgTimePerPerson sets the "avg_time_per_person" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetAvgTimePerPerson(i int) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.ResetAvgTimePerPerson()
	sftiuo.mutation.SetAvgTimePerPerson(i)
	return sftiuo
}

// SetNillableAvgTimePerPerson sets the "avg_time_per_person" field if the given value is not nil.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetNillableAvgTimePerPerson(i *int) *SchoolFitnessTestItemUpdateOne {
	if i != nil {
		sftiuo.SetAvgTimePerPerson(*i)
	}
	return sftiuo
}

// AddAvgTimePerPerson adds i to the "avg_time_per_person" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) AddAvgTimePerPerson(i int) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.AddAvgTimePerPerson(i)
	return sftiuo
}

// SetSchoolID sets the "school_id" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetSchoolID(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.ResetSchoolID()
	sftiuo.mutation.SetSchoolID(i)
	return sftiuo
}

// SetNillableSchoolID sets the "school_id" field if the given value is not nil.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetNillableSchoolID(i *int64) *SchoolFitnessTestItemUpdateOne {
	if i != nil {
		sftiuo.SetSchoolID(*i)
	}
	return sftiuo
}

// AddSchoolID adds i to the "school_id" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) AddSchoolID(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.AddSchoolID(i)
	return sftiuo
}

// SetItemID sets the "item_id" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetItemID(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.ResetItemID()
	sftiuo.mutation.SetItemID(i)
	return sftiuo
}

// SetNillableItemID sets the "item_id" field if the given value is not nil.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SetNillableItemID(i *int64) *SchoolFitnessTestItemUpdateOne {
	if i != nil {
		sftiuo.SetItemID(*i)
	}
	return sftiuo
}

// AddItemID adds i to the "item_id" field.
func (sftiuo *SchoolFitnessTestItemUpdateOne) AddItemID(i int64) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.AddItemID(i)
	return sftiuo
}

// Mutation returns the SchoolFitnessTestItemMutation object of the builder.
func (sftiuo *SchoolFitnessTestItemUpdateOne) Mutation() *SchoolFitnessTestItemMutation {
	return sftiuo.mutation
}

// Where appends a list predicates to the SchoolFitnessTestItemUpdate builder.
func (sftiuo *SchoolFitnessTestItemUpdateOne) Where(ps ...predicate.SchoolFitnessTestItem) *SchoolFitnessTestItemUpdateOne {
	sftiuo.mutation.Where(ps...)
	return sftiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sftiuo *SchoolFitnessTestItemUpdateOne) Select(field string, fields ...string) *SchoolFitnessTestItemUpdateOne {
	sftiuo.fields = append([]string{field}, fields...)
	return sftiuo
}

// Save executes the query and returns the updated SchoolFitnessTestItem entity.
func (sftiuo *SchoolFitnessTestItemUpdateOne) Save(ctx context.Context) (*SchoolFitnessTestItem, error) {
	sftiuo.defaults()
	return withHooks(ctx, sftiuo.sqlSave, sftiuo.mutation, sftiuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (sftiuo *SchoolFitnessTestItemUpdateOne) SaveX(ctx context.Context) *SchoolFitnessTestItem {
	node, err := sftiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sftiuo *SchoolFitnessTestItemUpdateOne) Exec(ctx context.Context) error {
	_, err := sftiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sftiuo *SchoolFitnessTestItemUpdateOne) ExecX(ctx context.Context) {
	if err := sftiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sftiuo *SchoolFitnessTestItemUpdateOne) defaults() {
	if _, ok := sftiuo.mutation.UpdatedAt(); !ok {
		v := schoolfitnesstestitem.UpdateDefaultUpdatedAt()
		sftiuo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (sftiuo *SchoolFitnessTestItemUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SchoolFitnessTestItemUpdateOne {
	sftiuo.modifiers = append(sftiuo.modifiers, modifiers...)
	return sftiuo
}

func (sftiuo *SchoolFitnessTestItemUpdateOne) sqlSave(ctx context.Context) (_node *SchoolFitnessTestItem, err error) {
	_spec := sqlgraph.NewUpdateSpec(schoolfitnesstestitem.Table, schoolfitnesstestitem.Columns, sqlgraph.NewFieldSpec(schoolfitnesstestitem.FieldID, field.TypeInt64))
	id, ok := sftiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent_work: missing "SchoolFitnessTestItem.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sftiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, schoolfitnesstestitem.FieldID)
		for _, f := range fields {
			if !schoolfitnesstestitem.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent_work: invalid field %q for query", f)}
			}
			if f != schoolfitnesstestitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sftiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sftiuo.mutation.CreatedBy(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldCreatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiuo.mutation.UpdatedBy(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := sftiuo.mutation.UpdatedAt(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := sftiuo.mutation.DeletedAt(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldDeletedAt, field.TypeTime, value)
	}
	if value, ok := sftiuo.mutation.MaxParticipants(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldMaxParticipants, field.TypeInt, value)
	}
	if value, ok := sftiuo.mutation.AddedMaxParticipants(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldMaxParticipants, field.TypeInt, value)
	}
	if value, ok := sftiuo.mutation.AvgTimePerPerson(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldAvgTimePerPerson, field.TypeInt, value)
	}
	if value, ok := sftiuo.mutation.AddedAvgTimePerPerson(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldAvgTimePerPerson, field.TypeInt, value)
	}
	if value, ok := sftiuo.mutation.SchoolID(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldSchoolID, field.TypeInt64, value)
	}
	if value, ok := sftiuo.mutation.AddedSchoolID(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldSchoolID, field.TypeInt64, value)
	}
	if value, ok := sftiuo.mutation.ItemID(); ok {
		_spec.SetField(schoolfitnesstestitem.FieldItemID, field.TypeInt64, value)
	}
	if value, ok := sftiuo.mutation.AddedItemID(); ok {
		_spec.AddField(schoolfitnesstestitem.FieldItemID, field.TypeInt64, value)
	}
	_spec.AddModifiers(sftiuo.modifiers...)
	_node = &SchoolFitnessTestItem{config: sftiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sftiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{schoolfitnesstestitem.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	sftiuo.mutation.done = true
	return _node, nil
}