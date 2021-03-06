// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"enttest/ent/migrate"

	"enttest/ent/hashtag"
	"enttest/ent/hashtagtopost"
	"enttest/ent/post"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Hashtag is the client for interacting with the Hashtag builders.
	Hashtag *HashtagClient
	// HashtagToPost is the client for interacting with the HashtagToPost builders.
	HashtagToPost *HashtagToPostClient
	// Post is the client for interacting with the Post builders.
	Post *PostClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Hashtag = NewHashtagClient(c.config)
	c.HashtagToPost = NewHashtagToPostClient(c.config)
	c.Post = NewPostClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		Hashtag:       NewHashtagClient(cfg),
		HashtagToPost: NewHashtagToPostClient(cfg),
		Post:          NewPostClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:           ctx,
		config:        cfg,
		Hashtag:       NewHashtagClient(cfg),
		HashtagToPost: NewHashtagToPostClient(cfg),
		Post:          NewPostClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Hashtag.
//		Query().
//		Count(ctx)
//
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
	c.Hashtag.Use(hooks...)
	c.HashtagToPost.Use(hooks...)
	c.Post.Use(hooks...)
}

// HashtagClient is a client for the Hashtag schema.
type HashtagClient struct {
	config
}

// NewHashtagClient returns a client for the Hashtag from the given config.
func NewHashtagClient(c config) *HashtagClient {
	return &HashtagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `hashtag.Hooks(f(g(h())))`.
func (c *HashtagClient) Use(hooks ...Hook) {
	c.hooks.Hashtag = append(c.hooks.Hashtag, hooks...)
}

// Create returns a builder for creating a Hashtag entity.
func (c *HashtagClient) Create() *HashtagCreate {
	mutation := newHashtagMutation(c.config, OpCreate)
	return &HashtagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Hashtag entities.
func (c *HashtagClient) CreateBulk(builders ...*HashtagCreate) *HashtagCreateBulk {
	return &HashtagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Hashtag.
func (c *HashtagClient) Update() *HashtagUpdate {
	mutation := newHashtagMutation(c.config, OpUpdate)
	return &HashtagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HashtagClient) UpdateOne(h *Hashtag) *HashtagUpdateOne {
	mutation := newHashtagMutation(c.config, OpUpdateOne, withHashtag(h))
	return &HashtagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HashtagClient) UpdateOneID(id string) *HashtagUpdateOne {
	mutation := newHashtagMutation(c.config, OpUpdateOne, withHashtagID(id))
	return &HashtagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Hashtag.
func (c *HashtagClient) Delete() *HashtagDelete {
	mutation := newHashtagMutation(c.config, OpDelete)
	return &HashtagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HashtagClient) DeleteOne(h *Hashtag) *HashtagDeleteOne {
	return c.DeleteOneID(h.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *HashtagClient) DeleteOneID(id string) *HashtagDeleteOne {
	builder := c.Delete().Where(hashtag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HashtagDeleteOne{builder}
}

// Query returns a query builder for Hashtag.
func (c *HashtagClient) Query() *HashtagQuery {
	return &HashtagQuery{
		config: c.config,
	}
}

// Get returns a Hashtag entity by its id.
func (c *HashtagClient) Get(ctx context.Context, id string) (*Hashtag, error) {
	return c.Query().Where(hashtag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HashtagClient) GetX(ctx context.Context, id string) *Hashtag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPost queries the post edge of a Hashtag.
func (c *HashtagClient) QueryPost(h *Hashtag) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hashtag.Table, hashtag.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, hashtag.PostTable, hashtag.PostPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHashtagToPost queries the hashtag_to_post edge of a Hashtag.
func (c *HashtagClient) QueryHashtagToPost(h *Hashtag) *HashtagToPostQuery {
	query := &HashtagToPostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := h.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hashtag.Table, hashtag.FieldID, id),
			sqlgraph.To(hashtagtopost.Table, hashtagtopost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, hashtag.HashtagToPostTable, hashtag.HashtagToPostColumn),
		)
		fromV = sqlgraph.Neighbors(h.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HashtagClient) Hooks() []Hook {
	return c.hooks.Hashtag
}

// HashtagToPostClient is a client for the HashtagToPost schema.
type HashtagToPostClient struct {
	config
}

// NewHashtagToPostClient returns a client for the HashtagToPost from the given config.
func NewHashtagToPostClient(c config) *HashtagToPostClient {
	return &HashtagToPostClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `hashtagtopost.Hooks(f(g(h())))`.
func (c *HashtagToPostClient) Use(hooks ...Hook) {
	c.hooks.HashtagToPost = append(c.hooks.HashtagToPost, hooks...)
}

// Create returns a builder for creating a HashtagToPost entity.
func (c *HashtagToPostClient) Create() *HashtagToPostCreate {
	mutation := newHashtagToPostMutation(c.config, OpCreate)
	return &HashtagToPostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of HashtagToPost entities.
func (c *HashtagToPostClient) CreateBulk(builders ...*HashtagToPostCreate) *HashtagToPostCreateBulk {
	return &HashtagToPostCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for HashtagToPost.
func (c *HashtagToPostClient) Update() *HashtagToPostUpdate {
	mutation := newHashtagToPostMutation(c.config, OpUpdate)
	return &HashtagToPostUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *HashtagToPostClient) UpdateOne(htp *HashtagToPost) *HashtagToPostUpdateOne {
	mutation := newHashtagToPostMutation(c.config, OpUpdateOne, withHashtagToPost(htp))
	return &HashtagToPostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *HashtagToPostClient) UpdateOneID(id string) *HashtagToPostUpdateOne {
	mutation := newHashtagToPostMutation(c.config, OpUpdateOne, withHashtagToPostID(id))
	return &HashtagToPostUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for HashtagToPost.
func (c *HashtagToPostClient) Delete() *HashtagToPostDelete {
	mutation := newHashtagToPostMutation(c.config, OpDelete)
	return &HashtagToPostDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *HashtagToPostClient) DeleteOne(htp *HashtagToPost) *HashtagToPostDeleteOne {
	return c.DeleteOneID(htp.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *HashtagToPostClient) DeleteOneID(id string) *HashtagToPostDeleteOne {
	builder := c.Delete().Where(hashtagtopost.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &HashtagToPostDeleteOne{builder}
}

// Query returns a query builder for HashtagToPost.
func (c *HashtagToPostClient) Query() *HashtagToPostQuery {
	return &HashtagToPostQuery{
		config: c.config,
	}
}

// Get returns a HashtagToPost entity by its id.
func (c *HashtagToPostClient) Get(ctx context.Context, id string) (*HashtagToPost, error) {
	return c.Query().Where(hashtagtopost.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *HashtagToPostClient) GetX(ctx context.Context, id string) *HashtagToPost {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPost queries the post edge of a HashtagToPost.
func (c *HashtagToPostClient) QueryPost(htp *HashtagToPost) *PostQuery {
	query := &PostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := htp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hashtagtopost.Table, hashtagtopost.FieldID, id),
			sqlgraph.To(post.Table, post.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, hashtagtopost.PostTable, hashtagtopost.PostColumn),
		)
		fromV = sqlgraph.Neighbors(htp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHashtag queries the hashtag edge of a HashtagToPost.
func (c *HashtagToPostClient) QueryHashtag(htp *HashtagToPost) *HashtagQuery {
	query := &HashtagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := htp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(hashtagtopost.Table, hashtagtopost.FieldID, id),
			sqlgraph.To(hashtag.Table, hashtag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, hashtagtopost.HashtagTable, hashtagtopost.HashtagColumn),
		)
		fromV = sqlgraph.Neighbors(htp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *HashtagToPostClient) Hooks() []Hook {
	return c.hooks.HashtagToPost
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

// Create returns a builder for creating a Post entity.
func (c *PostClient) Create() *PostCreate {
	mutation := newPostMutation(c.config, OpCreate)
	return &PostCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Post entities.
func (c *PostClient) CreateBulk(builders ...*PostCreate) *PostCreateBulk {
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
func (c *PostClient) UpdateOneID(id string) *PostUpdateOne {
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

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *PostClient) DeleteOneID(id string) *PostDeleteOne {
	builder := c.Delete().Where(post.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PostDeleteOne{builder}
}

// Query returns a query builder for Post.
func (c *PostClient) Query() *PostQuery {
	return &PostQuery{
		config: c.config,
	}
}

// Get returns a Post entity by its id.
func (c *PostClient) Get(ctx context.Context, id string) (*Post, error) {
	return c.Query().Where(post.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PostClient) GetX(ctx context.Context, id string) *Post {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryHashtags queries the hashtags edge of a Post.
func (c *PostClient) QueryHashtags(po *Post) *HashtagQuery {
	query := &HashtagQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(hashtag.Table, hashtag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, post.HashtagsTable, post.HashtagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryHashtagToPost queries the hashtag_to_post edge of a Post.
func (c *PostClient) QueryHashtagToPost(po *Post) *HashtagToPostQuery {
	query := &HashtagToPostQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(post.Table, post.FieldID, id),
			sqlgraph.To(hashtagtopost.Table, hashtagtopost.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, post.HashtagToPostTable, post.HashtagToPostColumn),
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
