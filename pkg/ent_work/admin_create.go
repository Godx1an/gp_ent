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
	"github.com/Godx1an/gp_ent/pkg/ent_work/admin"
)

// AdminCreate is the builder for creating a Admin entity.
type AdminCreate struct {
	config
	mutation *AdminMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedBy sets the "created_by" field.
func (ac *AdminCreate) SetCreatedBy(i int64) *AdminCreate {
	ac.mutation.SetCreatedBy(i)
	return ac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ac *AdminCreate) SetNillableCreatedBy(i *int64) *AdminCreate {
	if i != nil {
		ac.SetCreatedBy(*i)
	}
	return ac
}

// SetUpdatedBy sets the "updated_by" field.
func (ac *AdminCreate) SetUpdatedBy(i int64) *AdminCreate {
	ac.mutation.SetUpdatedBy(i)
	return ac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUpdatedBy(i *int64) *AdminCreate {
	if i != nil {
		ac.SetUpdatedBy(*i)
	}
	return ac
}

// SetCreatedAt sets the "created_at" field.
func (ac *AdminCreate) SetCreatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableCreatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AdminCreate) SetUpdatedAt(t time.Time) *AdminCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableUpdatedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AdminCreate) SetDeletedAt(t time.Time) *AdminCreate {
	ac.mutation.SetDeletedAt(t)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AdminCreate) SetNillableDeletedAt(t *time.Time) *AdminCreate {
	if t != nil {
		ac.SetDeletedAt(*t)
	}
	return ac
}

// SetPhone sets the "phone" field.
func (ac *AdminCreate) SetPhone(s string) *AdminCreate {
	ac.mutation.SetPhone(s)
	return ac
}

// SetNickname sets the "nickname" field.
func (ac *AdminCreate) SetNickname(s string) *AdminCreate {
	ac.mutation.SetNickname(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AdminCreate) SetPassword(s string) *AdminCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetSchool sets the "school" field.
func (ac *AdminCreate) SetSchool(s string) *AdminCreate {
	ac.mutation.SetSchool(s)
	return ac
}

// SetNillableSchool sets the "school" field if the given value is not nil.
func (ac *AdminCreate) SetNillableSchool(s *string) *AdminCreate {
	if s != nil {
		ac.SetSchool(*s)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AdminCreate) SetID(i int64) *AdminCreate {
	ac.mutation.SetID(i)
	return ac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ac *AdminCreate) SetNillableID(i *int64) *AdminCreate {
	if i != nil {
		ac.SetID(*i)
	}
	return ac
}

// Mutation returns the AdminMutation object of the builder.
func (ac *AdminCreate) Mutation() *AdminMutation {
	return ac.mutation
}

// Save creates the Admin in the database.
func (ac *AdminCreate) Save(ctx context.Context) (*Admin, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AdminCreate) SaveX(ctx context.Context) *Admin {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AdminCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AdminCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AdminCreate) defaults() {
	if _, ok := ac.mutation.CreatedBy(); !ok {
		v := admin.DefaultCreatedBy
		ac.mutation.SetCreatedBy(v)
	}
	if _, ok := ac.mutation.UpdatedBy(); !ok {
		v := admin.DefaultUpdatedBy
		ac.mutation.SetUpdatedBy(v)
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := admin.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := admin.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		v := admin.DefaultDeletedAt
		ac.mutation.SetDeletedAt(v)
	}
	if _, ok := ac.mutation.School(); !ok {
		v := admin.DefaultSchool
		ac.mutation.SetSchool(v)
	}
	if _, ok := ac.mutation.ID(); !ok {
		v := admin.DefaultID()
		ac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AdminCreate) check() error {
	if _, ok := ac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent_work: missing required field "Admin.created_by"`)}
	}
	if _, ok := ac.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent_work: missing required field "Admin.updated_by"`)}
	}
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent_work: missing required field "Admin.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent_work: missing required field "Admin.updated_at"`)}
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent_work: missing required field "Admin.deleted_at"`)}
	}
	if _, ok := ac.mutation.Phone(); !ok {
		return &ValidationError{Name: "phone", err: errors.New(`ent_work: missing required field "Admin.phone"`)}
	}
	if v, ok := ac.mutation.Phone(); ok {
		if err := admin.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent_work: validator failed for field "Admin.phone": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Nickname(); !ok {
		return &ValidationError{Name: "nickname", err: errors.New(`ent_work: missing required field "Admin.nickname"`)}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent_work: missing required field "Admin.password"`)}
	}
	if _, ok := ac.mutation.School(); !ok {
		return &ValidationError{Name: "school", err: errors.New(`ent_work: missing required field "Admin.school"`)}
	}
	return nil
}

func (ac *AdminCreate) sqlSave(ctx context.Context) (*Admin, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AdminCreate) createSpec() (*Admin, *sqlgraph.CreateSpec) {
	var (
		_node = &Admin{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(admin.Table, sqlgraph.NewFieldSpec(admin.FieldID, field.TypeInt64))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.SetField(admin.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.UpdatedBy(); ok {
		_spec.SetField(admin.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(admin.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(admin.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(admin.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.Phone(); ok {
		_spec.SetField(admin.FieldPhone, field.TypeString, value)
		_node.Phone = value
	}
	if value, ok := ac.mutation.Nickname(); ok {
		_spec.SetField(admin.FieldNickname, field.TypeString, value)
		_node.Nickname = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(admin.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.School(); ok {
		_spec.SetField(admin.FieldSchool, field.TypeString, value)
		_node.School = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Admin.Create().
//		SetCreatedBy(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdminUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (ac *AdminCreate) OnConflict(opts ...sql.ConflictOption) *AdminUpsertOne {
	ac.conflict = opts
	return &AdminUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AdminCreate) OnConflictColumns(columns ...string) *AdminUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AdminUpsertOne{
		create: ac,
	}
}

type (
	// AdminUpsertOne is the builder for "upsert"-ing
	//  one Admin node.
	AdminUpsertOne struct {
		create *AdminCreate
	}

	// AdminUpsert is the "OnConflict" setter.
	AdminUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedBy sets the "created_by" field.
func (u *AdminUpsert) SetCreatedBy(v int64) *AdminUpsert {
	u.Set(admin.FieldCreatedBy, v)
	return u
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AdminUpsert) UpdateCreatedBy() *AdminUpsert {
	u.SetExcluded(admin.FieldCreatedBy)
	return u
}

// AddCreatedBy adds v to the "created_by" field.
func (u *AdminUpsert) AddCreatedBy(v int64) *AdminUpsert {
	u.Add(admin.FieldCreatedBy, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AdminUpsert) SetUpdatedBy(v int64) *AdminUpsert {
	u.Set(admin.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AdminUpsert) UpdateUpdatedBy() *AdminUpsert {
	u.SetExcluded(admin.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AdminUpsert) AddUpdatedBy(v int64) *AdminUpsert {
	u.Add(admin.FieldUpdatedBy, v)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsert) SetUpdatedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateUpdatedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldUpdatedAt)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AdminUpsert) SetDeletedAt(v time.Time) *AdminUpsert {
	u.Set(admin.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AdminUpsert) UpdateDeletedAt() *AdminUpsert {
	u.SetExcluded(admin.FieldDeletedAt)
	return u
}

// SetPhone sets the "phone" field.
func (u *AdminUpsert) SetPhone(v string) *AdminUpsert {
	u.Set(admin.FieldPhone, v)
	return u
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *AdminUpsert) UpdatePhone() *AdminUpsert {
	u.SetExcluded(admin.FieldPhone)
	return u
}

// SetNickname sets the "nickname" field.
func (u *AdminUpsert) SetNickname(v string) *AdminUpsert {
	u.Set(admin.FieldNickname, v)
	return u
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *AdminUpsert) UpdateNickname() *AdminUpsert {
	u.SetExcluded(admin.FieldNickname)
	return u
}

// SetPassword sets the "password" field.
func (u *AdminUpsert) SetPassword(v string) *AdminUpsert {
	u.Set(admin.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AdminUpsert) UpdatePassword() *AdminUpsert {
	u.SetExcluded(admin.FieldPassword)
	return u
}

// SetSchool sets the "school" field.
func (u *AdminUpsert) SetSchool(v string) *AdminUpsert {
	u.Set(admin.FieldSchool, v)
	return u
}

// UpdateSchool sets the "school" field to the value that was provided on create.
func (u *AdminUpsert) UpdateSchool() *AdminUpsert {
	u.SetExcluded(admin.FieldSchool)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(admin.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AdminUpsertOne) UpdateNewValues() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(admin.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(admin.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AdminUpsertOne) Ignore() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdminUpsertOne) DoNothing() *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdminCreate.OnConflict
// documentation for more info.
func (u *AdminUpsertOne) Update(set func(*AdminUpsert)) *AdminUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdminUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *AdminUpsertOne) SetCreatedBy(v int64) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *AdminUpsertOne) AddCreatedBy(v int64) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateCreatedBy() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AdminUpsertOne) SetUpdatedBy(v int64) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AdminUpsertOne) AddUpdatedBy(v int64) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateUpdatedBy() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsertOne) SetUpdatedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateUpdatedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AdminUpsertOne) SetDeletedAt(v time.Time) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateDeletedAt() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetPhone sets the "phone" field.
func (u *AdminUpsertOne) SetPhone(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdatePhone() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePhone()
	})
}

// SetNickname sets the "nickname" field.
func (u *AdminUpsertOne) SetNickname(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateNickname() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateNickname()
	})
}

// SetPassword sets the "password" field.
func (u *AdminUpsertOne) SetPassword(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdatePassword() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePassword()
	})
}

// SetSchool sets the "school" field.
func (u *AdminUpsertOne) SetSchool(v string) *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.SetSchool(v)
	})
}

// UpdateSchool sets the "school" field to the value that was provided on create.
func (u *AdminUpsertOne) UpdateSchool() *AdminUpsertOne {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateSchool()
	})
}

// Exec executes the query.
func (u *AdminUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent_work: missing options for AdminCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdminUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AdminUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AdminUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AdminCreateBulk is the builder for creating many Admin entities in bulk.
type AdminCreateBulk struct {
	config
	err      error
	builders []*AdminCreate
	conflict []sql.ConflictOption
}

// Save creates the Admin entities in the database.
func (acb *AdminCreateBulk) Save(ctx context.Context) ([]*Admin, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Admin, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AdminMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AdminCreateBulk) SaveX(ctx context.Context) []*Admin {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AdminCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AdminCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Admin.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AdminUpsert) {
//			SetCreatedBy(v+v).
//		}).
//		Exec(ctx)
func (acb *AdminCreateBulk) OnConflict(opts ...sql.ConflictOption) *AdminUpsertBulk {
	acb.conflict = opts
	return &AdminUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AdminCreateBulk) OnConflictColumns(columns ...string) *AdminUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AdminUpsertBulk{
		create: acb,
	}
}

// AdminUpsertBulk is the builder for "upsert"-ing
// a bulk of Admin nodes.
type AdminUpsertBulk struct {
	create *AdminCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(admin.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AdminUpsertBulk) UpdateNewValues() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(admin.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(admin.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Admin.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AdminUpsertBulk) Ignore() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AdminUpsertBulk) DoNothing() *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AdminCreateBulk.OnConflict
// documentation for more info.
func (u *AdminUpsertBulk) Update(set func(*AdminUpsert)) *AdminUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AdminUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedBy sets the "created_by" field.
func (u *AdminUpsertBulk) SetCreatedBy(v int64) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetCreatedBy(v)
	})
}

// AddCreatedBy adds v to the "created_by" field.
func (u *AdminUpsertBulk) AddCreatedBy(v int64) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.AddCreatedBy(v)
	})
}

// UpdateCreatedBy sets the "created_by" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateCreatedBy() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateCreatedBy()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AdminUpsertBulk) SetUpdatedBy(v int64) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *AdminUpsertBulk) AddUpdatedBy(v int64) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateUpdatedBy() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AdminUpsertBulk) SetUpdatedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateUpdatedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AdminUpsertBulk) SetDeletedAt(v time.Time) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateDeletedAt() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetPhone sets the "phone" field.
func (u *AdminUpsertBulk) SetPhone(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetPhone(v)
	})
}

// UpdatePhone sets the "phone" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdatePhone() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePhone()
	})
}

// SetNickname sets the "nickname" field.
func (u *AdminUpsertBulk) SetNickname(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetNickname(v)
	})
}

// UpdateNickname sets the "nickname" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateNickname() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateNickname()
	})
}

// SetPassword sets the "password" field.
func (u *AdminUpsertBulk) SetPassword(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdatePassword() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdatePassword()
	})
}

// SetSchool sets the "school" field.
func (u *AdminUpsertBulk) SetSchool(v string) *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.SetSchool(v)
	})
}

// UpdateSchool sets the "school" field to the value that was provided on create.
func (u *AdminUpsertBulk) UpdateSchool() *AdminUpsertBulk {
	return u.Update(func(s *AdminUpsert) {
		s.UpdateSchool()
	})
}

// Exec executes the query.
func (u *AdminUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent_work: OnConflict was set for builder %d. Set it on the AdminCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent_work: missing options for AdminCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AdminUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
