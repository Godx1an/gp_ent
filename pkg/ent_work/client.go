// Code generated by ent, DO NOT EDIT.

package ent_work

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Godx1an/gp_ent/pkg/ent_work/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Godx1an/gp_ent/pkg/ent_work/admin"
	"github.com/Godx1an/gp_ent/pkg/ent_work/fitnesstestitem"
	"github.com/Godx1an/gp_ent/pkg/ent_work/school"
	"github.com/Godx1an/gp_ent/pkg/ent_work/schoolfitnesstestitem"
	"github.com/Godx1an/gp_ent/pkg/ent_work/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Admin is the client for interacting with the Admin builders.
	Admin *AdminClient
	// FitnessTestItem is the client for interacting with the FitnessTestItem builders.
	FitnessTestItem *FitnessTestItemClient
	// School is the client for interacting with the School builders.
	School *SchoolClient
	// SchoolFitnessTestItem is the client for interacting with the SchoolFitnessTestItem builders.
	SchoolFitnessTestItem *SchoolFitnessTestItemClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Admin = NewAdminClient(c.config)
	c.FitnessTestItem = NewFitnessTestItemClient(c.config)
	c.School = NewSchoolClient(c.config)
	c.SchoolFitnessTestItem = NewSchoolFitnessTestItemClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent_work: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent_work: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                   ctx,
		config:                cfg,
		Admin:                 NewAdminClient(cfg),
		FitnessTestItem:       NewFitnessTestItemClient(cfg),
		School:                NewSchoolClient(cfg),
		SchoolFitnessTestItem: NewSchoolFitnessTestItemClient(cfg),
		User:                  NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:                   ctx,
		config:                cfg,
		Admin:                 NewAdminClient(cfg),
		FitnessTestItem:       NewFitnessTestItemClient(cfg),
		School:                NewSchoolClient(cfg),
		SchoolFitnessTestItem: NewSchoolFitnessTestItemClient(cfg),
		User:                  NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Admin.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Admin.Use(hooks...)
	c.FitnessTestItem.Use(hooks...)
	c.School.Use(hooks...)
	c.SchoolFitnessTestItem.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Admin.Intercept(interceptors...)
	c.FitnessTestItem.Intercept(interceptors...)
	c.School.Intercept(interceptors...)
	c.SchoolFitnessTestItem.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AdminMutation:
		return c.Admin.mutate(ctx, m)
	case *FitnessTestItemMutation:
		return c.FitnessTestItem.mutate(ctx, m)
	case *SchoolMutation:
		return c.School.mutate(ctx, m)
	case *SchoolFitnessTestItemMutation:
		return c.SchoolFitnessTestItem.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent_work: unknown mutation type %T", m)
	}
}

// AdminClient is a client for the Admin schema.
type AdminClient struct {
	config
}

// NewAdminClient returns a client for the Admin from the given config.
func NewAdminClient(c config) *AdminClient {
	return &AdminClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `admin.Hooks(f(g(h())))`.
func (c *AdminClient) Use(hooks ...Hook) {
	c.hooks.Admin = append(c.hooks.Admin, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `admin.Intercept(f(g(h())))`.
func (c *AdminClient) Intercept(interceptors ...Interceptor) {
	c.inters.Admin = append(c.inters.Admin, interceptors...)
}

// Create returns a builder for creating a Admin entity.
func (c *AdminClient) Create() *AdminCreate {
	mutation := newAdminMutation(c.config, OpCreate)
	return &AdminCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Admin entities.
func (c *AdminClient) CreateBulk(builders ...*AdminCreate) *AdminCreateBulk {
	return &AdminCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AdminClient) MapCreateBulk(slice any, setFunc func(*AdminCreate, int)) *AdminCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AdminCreateBulk{err: fmt.Errorf("calling to AdminClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AdminCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AdminCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Admin.
func (c *AdminClient) Update() *AdminUpdate {
	mutation := newAdminMutation(c.config, OpUpdate)
	return &AdminUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AdminClient) UpdateOne(a *Admin) *AdminUpdateOne {
	mutation := newAdminMutation(c.config, OpUpdateOne, withAdmin(a))
	return &AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AdminClient) UpdateOneID(id int64) *AdminUpdateOne {
	mutation := newAdminMutation(c.config, OpUpdateOne, withAdminID(id))
	return &AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Admin.
func (c *AdminClient) Delete() *AdminDelete {
	mutation := newAdminMutation(c.config, OpDelete)
	return &AdminDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AdminClient) DeleteOne(a *Admin) *AdminDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AdminClient) DeleteOneID(id int64) *AdminDeleteOne {
	builder := c.Delete().Where(admin.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AdminDeleteOne{builder}
}

// Query returns a query builder for Admin.
func (c *AdminClient) Query() *AdminQuery {
	return &AdminQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAdmin},
		inters: c.Interceptors(),
	}
}

// Get returns a Admin entity by its id.
func (c *AdminClient) Get(ctx context.Context, id int64) (*Admin, error) {
	return c.Query().Where(admin.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AdminClient) GetX(ctx context.Context, id int64) *Admin {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AdminClient) Hooks() []Hook {
	return c.hooks.Admin
}

// Interceptors returns the client interceptors.
func (c *AdminClient) Interceptors() []Interceptor {
	return c.inters.Admin
}

func (c *AdminClient) mutate(ctx context.Context, m *AdminMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AdminCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AdminUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AdminUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AdminDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent_work: unknown Admin mutation op: %q", m.Op())
	}
}

// FitnessTestItemClient is a client for the FitnessTestItem schema.
type FitnessTestItemClient struct {
	config
}

// NewFitnessTestItemClient returns a client for the FitnessTestItem from the given config.
func NewFitnessTestItemClient(c config) *FitnessTestItemClient {
	return &FitnessTestItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `fitnesstestitem.Hooks(f(g(h())))`.
func (c *FitnessTestItemClient) Use(hooks ...Hook) {
	c.hooks.FitnessTestItem = append(c.hooks.FitnessTestItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `fitnesstestitem.Intercept(f(g(h())))`.
func (c *FitnessTestItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.FitnessTestItem = append(c.inters.FitnessTestItem, interceptors...)
}

// Create returns a builder for creating a FitnessTestItem entity.
func (c *FitnessTestItemClient) Create() *FitnessTestItemCreate {
	mutation := newFitnessTestItemMutation(c.config, OpCreate)
	return &FitnessTestItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of FitnessTestItem entities.
func (c *FitnessTestItemClient) CreateBulk(builders ...*FitnessTestItemCreate) *FitnessTestItemCreateBulk {
	return &FitnessTestItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *FitnessTestItemClient) MapCreateBulk(slice any, setFunc func(*FitnessTestItemCreate, int)) *FitnessTestItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &FitnessTestItemCreateBulk{err: fmt.Errorf("calling to FitnessTestItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*FitnessTestItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &FitnessTestItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for FitnessTestItem.
func (c *FitnessTestItemClient) Update() *FitnessTestItemUpdate {
	mutation := newFitnessTestItemMutation(c.config, OpUpdate)
	return &FitnessTestItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FitnessTestItemClient) UpdateOne(fti *FitnessTestItem) *FitnessTestItemUpdateOne {
	mutation := newFitnessTestItemMutation(c.config, OpUpdateOne, withFitnessTestItem(fti))
	return &FitnessTestItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FitnessTestItemClient) UpdateOneID(id int64) *FitnessTestItemUpdateOne {
	mutation := newFitnessTestItemMutation(c.config, OpUpdateOne, withFitnessTestItemID(id))
	return &FitnessTestItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for FitnessTestItem.
func (c *FitnessTestItemClient) Delete() *FitnessTestItemDelete {
	mutation := newFitnessTestItemMutation(c.config, OpDelete)
	return &FitnessTestItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *FitnessTestItemClient) DeleteOne(fti *FitnessTestItem) *FitnessTestItemDeleteOne {
	return c.DeleteOneID(fti.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *FitnessTestItemClient) DeleteOneID(id int64) *FitnessTestItemDeleteOne {
	builder := c.Delete().Where(fitnesstestitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FitnessTestItemDeleteOne{builder}
}

// Query returns a query builder for FitnessTestItem.
func (c *FitnessTestItemClient) Query() *FitnessTestItemQuery {
	return &FitnessTestItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeFitnessTestItem},
		inters: c.Interceptors(),
	}
}

// Get returns a FitnessTestItem entity by its id.
func (c *FitnessTestItemClient) Get(ctx context.Context, id int64) (*FitnessTestItem, error) {
	return c.Query().Where(fitnesstestitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FitnessTestItemClient) GetX(ctx context.Context, id int64) *FitnessTestItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *FitnessTestItemClient) Hooks() []Hook {
	return c.hooks.FitnessTestItem
}

// Interceptors returns the client interceptors.
func (c *FitnessTestItemClient) Interceptors() []Interceptor {
	return c.inters.FitnessTestItem
}

func (c *FitnessTestItemClient) mutate(ctx context.Context, m *FitnessTestItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&FitnessTestItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&FitnessTestItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&FitnessTestItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&FitnessTestItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent_work: unknown FitnessTestItem mutation op: %q", m.Op())
	}
}

// SchoolClient is a client for the School schema.
type SchoolClient struct {
	config
}

// NewSchoolClient returns a client for the School from the given config.
func NewSchoolClient(c config) *SchoolClient {
	return &SchoolClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `school.Hooks(f(g(h())))`.
func (c *SchoolClient) Use(hooks ...Hook) {
	c.hooks.School = append(c.hooks.School, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `school.Intercept(f(g(h())))`.
func (c *SchoolClient) Intercept(interceptors ...Interceptor) {
	c.inters.School = append(c.inters.School, interceptors...)
}

// Create returns a builder for creating a School entity.
func (c *SchoolClient) Create() *SchoolCreate {
	mutation := newSchoolMutation(c.config, OpCreate)
	return &SchoolCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of School entities.
func (c *SchoolClient) CreateBulk(builders ...*SchoolCreate) *SchoolCreateBulk {
	return &SchoolCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SchoolClient) MapCreateBulk(slice any, setFunc func(*SchoolCreate, int)) *SchoolCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SchoolCreateBulk{err: fmt.Errorf("calling to SchoolClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SchoolCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SchoolCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for School.
func (c *SchoolClient) Update() *SchoolUpdate {
	mutation := newSchoolMutation(c.config, OpUpdate)
	return &SchoolUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SchoolClient) UpdateOne(s *School) *SchoolUpdateOne {
	mutation := newSchoolMutation(c.config, OpUpdateOne, withSchool(s))
	return &SchoolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SchoolClient) UpdateOneID(id int64) *SchoolUpdateOne {
	mutation := newSchoolMutation(c.config, OpUpdateOne, withSchoolID(id))
	return &SchoolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for School.
func (c *SchoolClient) Delete() *SchoolDelete {
	mutation := newSchoolMutation(c.config, OpDelete)
	return &SchoolDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SchoolClient) DeleteOne(s *School) *SchoolDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SchoolClient) DeleteOneID(id int64) *SchoolDeleteOne {
	builder := c.Delete().Where(school.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SchoolDeleteOne{builder}
}

// Query returns a query builder for School.
func (c *SchoolClient) Query() *SchoolQuery {
	return &SchoolQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSchool},
		inters: c.Interceptors(),
	}
}

// Get returns a School entity by its id.
func (c *SchoolClient) Get(ctx context.Context, id int64) (*School, error) {
	return c.Query().Where(school.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SchoolClient) GetX(ctx context.Context, id int64) *School {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SchoolClient) Hooks() []Hook {
	return c.hooks.School
}

// Interceptors returns the client interceptors.
func (c *SchoolClient) Interceptors() []Interceptor {
	return c.inters.School
}

func (c *SchoolClient) mutate(ctx context.Context, m *SchoolMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SchoolCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SchoolUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SchoolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SchoolDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent_work: unknown School mutation op: %q", m.Op())
	}
}

// SchoolFitnessTestItemClient is a client for the SchoolFitnessTestItem schema.
type SchoolFitnessTestItemClient struct {
	config
}

// NewSchoolFitnessTestItemClient returns a client for the SchoolFitnessTestItem from the given config.
func NewSchoolFitnessTestItemClient(c config) *SchoolFitnessTestItemClient {
	return &SchoolFitnessTestItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `schoolfitnesstestitem.Hooks(f(g(h())))`.
func (c *SchoolFitnessTestItemClient) Use(hooks ...Hook) {
	c.hooks.SchoolFitnessTestItem = append(c.hooks.SchoolFitnessTestItem, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `schoolfitnesstestitem.Intercept(f(g(h())))`.
func (c *SchoolFitnessTestItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.SchoolFitnessTestItem = append(c.inters.SchoolFitnessTestItem, interceptors...)
}

// Create returns a builder for creating a SchoolFitnessTestItem entity.
func (c *SchoolFitnessTestItemClient) Create() *SchoolFitnessTestItemCreate {
	mutation := newSchoolFitnessTestItemMutation(c.config, OpCreate)
	return &SchoolFitnessTestItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SchoolFitnessTestItem entities.
func (c *SchoolFitnessTestItemClient) CreateBulk(builders ...*SchoolFitnessTestItemCreate) *SchoolFitnessTestItemCreateBulk {
	return &SchoolFitnessTestItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SchoolFitnessTestItemClient) MapCreateBulk(slice any, setFunc func(*SchoolFitnessTestItemCreate, int)) *SchoolFitnessTestItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SchoolFitnessTestItemCreateBulk{err: fmt.Errorf("calling to SchoolFitnessTestItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SchoolFitnessTestItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SchoolFitnessTestItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SchoolFitnessTestItem.
func (c *SchoolFitnessTestItemClient) Update() *SchoolFitnessTestItemUpdate {
	mutation := newSchoolFitnessTestItemMutation(c.config, OpUpdate)
	return &SchoolFitnessTestItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SchoolFitnessTestItemClient) UpdateOne(sfti *SchoolFitnessTestItem) *SchoolFitnessTestItemUpdateOne {
	mutation := newSchoolFitnessTestItemMutation(c.config, OpUpdateOne, withSchoolFitnessTestItem(sfti))
	return &SchoolFitnessTestItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SchoolFitnessTestItemClient) UpdateOneID(id int64) *SchoolFitnessTestItemUpdateOne {
	mutation := newSchoolFitnessTestItemMutation(c.config, OpUpdateOne, withSchoolFitnessTestItemID(id))
	return &SchoolFitnessTestItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SchoolFitnessTestItem.
func (c *SchoolFitnessTestItemClient) Delete() *SchoolFitnessTestItemDelete {
	mutation := newSchoolFitnessTestItemMutation(c.config, OpDelete)
	return &SchoolFitnessTestItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SchoolFitnessTestItemClient) DeleteOne(sfti *SchoolFitnessTestItem) *SchoolFitnessTestItemDeleteOne {
	return c.DeleteOneID(sfti.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SchoolFitnessTestItemClient) DeleteOneID(id int64) *SchoolFitnessTestItemDeleteOne {
	builder := c.Delete().Where(schoolfitnesstestitem.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SchoolFitnessTestItemDeleteOne{builder}
}

// Query returns a query builder for SchoolFitnessTestItem.
func (c *SchoolFitnessTestItemClient) Query() *SchoolFitnessTestItemQuery {
	return &SchoolFitnessTestItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSchoolFitnessTestItem},
		inters: c.Interceptors(),
	}
}

// Get returns a SchoolFitnessTestItem entity by its id.
func (c *SchoolFitnessTestItemClient) Get(ctx context.Context, id int64) (*SchoolFitnessTestItem, error) {
	return c.Query().Where(schoolfitnesstestitem.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SchoolFitnessTestItemClient) GetX(ctx context.Context, id int64) *SchoolFitnessTestItem {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SchoolFitnessTestItemClient) Hooks() []Hook {
	return c.hooks.SchoolFitnessTestItem
}

// Interceptors returns the client interceptors.
func (c *SchoolFitnessTestItemClient) Interceptors() []Interceptor {
	return c.inters.SchoolFitnessTestItem
}

func (c *SchoolFitnessTestItemClient) mutate(ctx context.Context, m *SchoolFitnessTestItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SchoolFitnessTestItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SchoolFitnessTestItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SchoolFitnessTestItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SchoolFitnessTestItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent_work: unknown SchoolFitnessTestItem mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int64) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int64) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int64) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int64) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent_work: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Admin, FitnessTestItem, School, SchoolFitnessTestItem, User []ent.Hook
	}
	inters struct {
		Admin, FitnessTestItem, School, SchoolFitnessTestItem, User []ent.Interceptor
	}
)
