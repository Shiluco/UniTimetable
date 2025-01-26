// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/Shiluco/UniTimetable/backend/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Shiluco/UniTimetable/backend/ent/department"
	"github.com/Shiluco/UniTimetable/backend/ent/major"
	"github.com/Shiluco/UniTimetable/backend/ent/post"
	"github.com/Shiluco/UniTimetable/backend/ent/schedule"
	"github.com/Shiluco/UniTimetable/backend/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Department is the client for interacting with the Department builders.
	Department *DepartmentClient
	// Major is the client for interacting with the Major builders.
	Major *MajorClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
	// Schedule is the client for interacting with the Schedule builders.
	Schedule *ScheduleClient
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
	c.Department = NewDepartmentClient(c.config)
	c.Major = NewMajorClient(c.config)
	c.Post = NewPostClient(c.config)
	c.Schedule = NewScheduleClient(c.config)
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
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Major:      NewMajorClient(cfg),
		Post:       NewPostClient(cfg),
		Schedule:   NewScheduleClient(cfg),
		User:       NewUserClient(cfg),
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
		ctx:        ctx,
		config:     cfg,
		Department: NewDepartmentClient(cfg),
		Major:      NewMajorClient(cfg),
		Post:       NewPostClient(cfg),
		Schedule:   NewScheduleClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Department.
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
	c.Department.Use(hooks...)
	c.Major.Use(hooks...)
	c.Post.Use(hooks...)
	c.Schedule.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Department.Intercept(interceptors...)
	c.Major.Intercept(interceptors...)
	c.Post.Intercept(interceptors...)
	c.Schedule.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *DepartmentMutation:
		return c.Department.mutate(ctx, m)
	case *MajorMutation:
		return c.Major.mutate(ctx, m)
	case *PostMutation:
		return c.Post.mutate(ctx, m)
	case *ScheduleMutation:
		return c.Schedule.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// DepartmentClient is a client for the Department schema.
type DepartmentClient struct {
	config
}

// NewDepartmentClient returns a client for the Department from the given config.
func NewDepartmentClient(c config) *DepartmentClient {
	return &DepartmentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `department.Hooks(f(g(h())))`.
func (c *DepartmentClient) Use(hooks ...Hook) {
	c.hooks.Department = append(c.hooks.Department, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `department.Intercept(f(g(h())))`.
func (c *DepartmentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Department = append(c.inters.Department, interceptors...)
}

// Create returns a builder for creating a Department entity.
func (c *DepartmentClient) Create() *DepartmentCreate {
	mutation := newDepartmentMutation(c.config, OpCreate)
	return &DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Department entities.
func (c *DepartmentClient) CreateBulk(builders ...*DepartmentCreate) *DepartmentCreateBulk {
	return &DepartmentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *DepartmentClient) MapCreateBulk(slice any, setFunc func(*DepartmentCreate, int)) *DepartmentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &DepartmentCreateBulk{err: fmt.Errorf("calling to DepartmentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*DepartmentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &DepartmentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Department.
func (c *DepartmentClient) Update() *DepartmentUpdate {
	mutation := newDepartmentMutation(c.config, OpUpdate)
	return &DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DepartmentClient) UpdateOne(d *Department) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartment(d))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DepartmentClient) UpdateOneID(id int) *DepartmentUpdateOne {
	mutation := newDepartmentMutation(c.config, OpUpdateOne, withDepartmentID(id))
	return &DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Department.
func (c *DepartmentClient) Delete() *DepartmentDelete {
	mutation := newDepartmentMutation(c.config, OpDelete)
	return &DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *DepartmentClient) DeleteOne(d *Department) *DepartmentDeleteOne {
	return c.DeleteOneID(d.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *DepartmentClient) DeleteOneID(id int) *DepartmentDeleteOne {
	builder := c.Delete().Where(department.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DepartmentDeleteOne{builder}
}

// Query returns a query builder for Department.
func (c *DepartmentClient) Query() *DepartmentQuery {
	return &DepartmentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeDepartment},
		inters: c.Interceptors(),
	}
}

// Get returns a Department entity by its id.
func (c *DepartmentClient) Get(ctx context.Context, id int) (*Department, error) {
	return c.Query().Where(department.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DepartmentClient) GetX(ctx context.Context, id int) *Department {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsers queries the users edge of a Department.
func (c *DepartmentClient) QueryUsers(d *Department) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.UsersTable, department.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMajors queries the majors edge of a Department.
func (c *DepartmentClient) QueryMajors(d *Department) *MajorQuery {
	query := (&MajorClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := d.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(department.Table, department.FieldID, id),
			sqlgraph.To(major.Table, major.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, department.MajorsTable, department.MajorsColumn),
		)
		fromV = sqlgraph.Neighbors(d.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DepartmentClient) Hooks() []Hook {
	return c.hooks.Department
}

// Interceptors returns the client interceptors.
func (c *DepartmentClient) Interceptors() []Interceptor {
	return c.inters.Department
}

func (c *DepartmentClient) mutate(ctx context.Context, m *DepartmentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&DepartmentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&DepartmentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&DepartmentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&DepartmentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Department mutation op: %q", m.Op())
	}
}

// MajorClient is a client for the Major schema.
type MajorClient struct {
	config
}

// NewMajorClient returns a client for the Major from the given config.
func NewMajorClient(c config) *MajorClient {
	return &MajorClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `major.Hooks(f(g(h())))`.
func (c *MajorClient) Use(hooks ...Hook) {
	c.hooks.Major = append(c.hooks.Major, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `major.Intercept(f(g(h())))`.
func (c *MajorClient) Intercept(interceptors ...Interceptor) {
	c.inters.Major = append(c.inters.Major, interceptors...)
}

// Create returns a builder for creating a Major entity.
func (c *MajorClient) Create() *MajorCreate {
	mutation := newMajorMutation(c.config, OpCreate)
	return &MajorCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Major entities.
func (c *MajorClient) CreateBulk(builders ...*MajorCreate) *MajorCreateBulk {
	return &MajorCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *MajorClient) MapCreateBulk(slice any, setFunc func(*MajorCreate, int)) *MajorCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &MajorCreateBulk{err: fmt.Errorf("calling to MajorClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*MajorCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &MajorCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Major.
func (c *MajorClient) Update() *MajorUpdate {
	mutation := newMajorMutation(c.config, OpUpdate)
	return &MajorUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *MajorClient) UpdateOne(m *Major) *MajorUpdateOne {
	mutation := newMajorMutation(c.config, OpUpdateOne, withMajor(m))
	return &MajorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *MajorClient) UpdateOneID(id int) *MajorUpdateOne {
	mutation := newMajorMutation(c.config, OpUpdateOne, withMajorID(id))
	return &MajorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Major.
func (c *MajorClient) Delete() *MajorDelete {
	mutation := newMajorMutation(c.config, OpDelete)
	return &MajorDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *MajorClient) DeleteOne(m *Major) *MajorDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *MajorClient) DeleteOneID(id int) *MajorDeleteOne {
	builder := c.Delete().Where(major.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &MajorDeleteOne{builder}
}

// Query returns a query builder for Major.
func (c *MajorClient) Query() *MajorQuery {
	return &MajorQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeMajor},
		inters: c.Interceptors(),
	}
}

// Get returns a Major entity by its id.
func (c *MajorClient) Get(ctx context.Context, id int) (*Major, error) {
	return c.Query().Where(major.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *MajorClient) GetX(ctx context.Context, id int) *Major {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDepartment queries the department edge of a Major.
func (c *MajorClient) QueryDepartment(m *Major) *DepartmentQuery {
	query := (&DepartmentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, major.DepartmentTable, major.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryUsers queries the users edge of a Major.
func (c *MajorClient) QueryUsers(m *Major) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(major.Table, major.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, major.UsersTable, major.UsersColumn),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *MajorClient) Hooks() []Hook {
	return c.hooks.Major
}

// Interceptors returns the client interceptors.
func (c *MajorClient) Interceptors() []Interceptor {
	return c.inters.Major
}

func (c *MajorClient) mutate(ctx context.Context, m *MajorMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&MajorCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&MajorUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&MajorUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&MajorDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Major mutation op: %q", m.Op())
	}
}

// PostClient is a client for the Post schema.
type PostClient struct {
	config
}

// NewPostClient returns a client for the Post from the given config.
func NewPostClient(c config) *PostClient {
	return &PostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `post.Hooks(f(g(h())))`.
func (c *PostClient) Use(hooks ...Hook) {
	c.hooks.Post = append(c.hooks.Post, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `post.Intercept(f(g(h())))`.
func (c *PostClient) Intercept(interceptors ...Interceptor) {
	c.inters.Post = append(c.inters.Post, interceptors...)
}

// Create returns a builder for creating a Post entity.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
	return &PostCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PostClient) MapCreateBulk(slice any, setFunc func(*PostCreate, int)) *PostCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PostCreateBulk{err: fmt.Errorf("calling to PostClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PostCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Post.
func (c *PostClient) Update() *PostUpdate {
	mutation := newPostMutation(c.config, OpUpdate)
	return &PostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PostClient) UpdateOne(po *Post) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPost(po))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PostClient) UpdateOneID(id int) *PostUpdateOne {
	mutation := newPostMutation(c.config, OpUpdateOne, withPostID(id))
	return &PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Post.
func (c *PostClient) Delete() *PostDelete {
	mutation := newPostMutation(c.config, OpDelete)
	return &PostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PostClient) DeleteOne(po *Post) *PostDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PostClient) DeleteOneID(id int) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePost},
		inters: c.Interceptors(),
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id int) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id int) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Post.
func (c *PostClient) QueryUser(po *Post) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.UserTable, post.UserColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySchedule queries the schedule edge of a Post.
func (c *PostClient) QuerySchedule(po *Post) *ScheduleQuery {
	query := (&ScheduleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(schedule.Table, schedule.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.ScheduleTable, post.ScheduleColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Post.
func (c *PostClient) QueryParent(po *Post) *PostQuery {
	query := (&PostClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, post.ParentTable, post.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReplies queries the replies edge of a Post.
func (c *PostClient) QueryReplies(po *Post) *PostQuery {
	query := (&PostClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, post.RepliesTable, post.RepliesColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PostClient) Hooks() []Hook {
	return c.hooks.Post
}

// Interceptors returns the client interceptors.
func (c *PostClient) Interceptors() []Interceptor {
	return c.inters.Post
}

func (c *PostClient) mutate(ctx context.Context, m *PostMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PostCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PostUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PostDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Post mutation op: %q", m.Op())
	}
}

// ScheduleClient is a client for the Schedule schema.
type ScheduleClient struct {
	config
}

// NewScheduleClient returns a client for the Schedule from the given config.
func NewScheduleClient(c config) *ScheduleClient {
	return &ScheduleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `schedule.Hooks(f(g(h())))`.
func (c *ScheduleClient) Use(hooks ...Hook) {
	c.hooks.Schedule = append(c.hooks.Schedule, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `schedule.Intercept(f(g(h())))`.
func (c *ScheduleClient) Intercept(interceptors ...Interceptor) {
	c.inters.Schedule = append(c.inters.Schedule, interceptors...)
}

// Create returns a builder for creating a Schedule entity.
func (c *ScheduleClient) Create() *ScheduleCreate {
	mutation := newScheduleMutation(c.config, OpCreate)
	return &ScheduleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Schedule entities.
func (c *ScheduleClient) CreateBulk(builders ...*ScheduleCreate) *ScheduleCreateBulk {
	return &ScheduleCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ScheduleClient) MapCreateBulk(slice any, setFunc func(*ScheduleCreate, int)) *ScheduleCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ScheduleCreateBulk{err: fmt.Errorf("calling to ScheduleClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ScheduleCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ScheduleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Schedule.
func (c *ScheduleClient) Update() *ScheduleUpdate {
	mutation := newScheduleMutation(c.config, OpUpdate)
	return &ScheduleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ScheduleClient) UpdateOne(s *Schedule) *ScheduleUpdateOne {
	mutation := newScheduleMutation(c.config, OpUpdateOne, withSchedule(s))
	return &ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ScheduleClient) UpdateOneID(id int) *ScheduleUpdateOne {
	mutation := newScheduleMutation(c.config, OpUpdateOne, withScheduleID(id))
	return &ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Schedule.
func (c *ScheduleClient) Delete() *ScheduleDelete {
	mutation := newScheduleMutation(c.config, OpDelete)
	return &ScheduleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ScheduleClient) DeleteOne(s *Schedule) *ScheduleDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ScheduleClient) DeleteOneID(id int) *ScheduleDeleteOne {
	builder := c.Delete().Where(schedule.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ScheduleDeleteOne{builder}
}

// Query returns a query builder for Schedule.
func (c *ScheduleClient) Query() *ScheduleQuery {
	return &ScheduleQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSchedule},
		inters: c.Interceptors(),
	}
}

// Get returns a Schedule entity by its id.
func (c *ScheduleClient) Get(ctx context.Context, id int) (*Schedule, error) {
	return c.Query().Where(schedule.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ScheduleClient) GetX(ctx context.Context, id int) *Schedule {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Schedule.
func (c *ScheduleClient) QueryUser(s *Schedule) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schedule.Table, schedule.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, schedule.UserTable, schedule.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPosts queries the posts edge of a Schedule.
func (c *ScheduleClient) QueryPosts(s *Schedule) *PostQuery {
	query := (&PostClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(schedule.Table, schedule.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, schedule.PostsTable, schedule.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ScheduleClient) Hooks() []Hook {
	return c.hooks.Schedule
}

// Interceptors returns the client interceptors.
func (c *ScheduleClient) Interceptors() []Interceptor {
	return c.inters.Schedule
}

func (c *ScheduleClient) mutate(ctx context.Context, m *ScheduleMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ScheduleCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ScheduleUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ScheduleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ScheduleDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Schedule mutation op: %q", m.Op())
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
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
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
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
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
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDepartment queries the department edge of a User.
func (c *UserClient) QueryDepartment(u *User) *DepartmentQuery {
	query := (&DepartmentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(department.Table, department.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.DepartmentTable, user.DepartmentColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryMajor queries the major edge of a User.
func (c *UserClient) QueryMajor(u *User) *MajorQuery {
	query := (&MajorClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(major.Table, major.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.MajorTable, user.MajorColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPosts queries the posts edge of a User.
func (c *UserClient) QueryPosts(u *User) *PostQuery {
	query := (&PostClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.PostsTable, user.PostsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySchedules queries the schedules edge of a User.
func (c *UserClient) QuerySchedules(u *User) *ScheduleQuery {
	query := (&ScheduleClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(schedule.Table, schedule.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.SchedulesTable, user.SchedulesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
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
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Department, Major, Post, Schedule, User []ent.Hook
	}
	inters struct {
		Department, Major, Post, Schedule, User []ent.Interceptor
	}
)
