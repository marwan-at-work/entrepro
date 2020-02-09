// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"entrepro/ent/migrate"

	"entrepro/ent/spec"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Spec is the client for interacting with the Spec builders.
	Spec *SpecClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		Spec:   NewSpecClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
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

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
		config: cfg,
		Spec:   NewSpecClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Spec.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config: cfg,
		Schema: migrate.NewSchema(cfg.driver),
		Spec:   NewSpecClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// SpecClient is a client for the Spec schema.
type SpecClient struct {
	config
}

// NewSpecClient returns a client for the Spec from the given config.
func NewSpecClient(c config) *SpecClient {
	return &SpecClient{config: c}
}

// Create returns a create builder for Spec.
func (c *SpecClient) Create() *SpecCreate {
	return &SpecCreate{config: c.config}
}

// Update returns an update builder for Spec.
func (c *SpecClient) Update() *SpecUpdate {
	return &SpecUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *SpecClient) UpdateOne(s *Spec) *SpecUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *SpecClient) UpdateOneID(id string) *SpecUpdateOne {
	return &SpecUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Spec.
func (c *SpecClient) Delete() *SpecDelete {
	return &SpecDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SpecClient) DeleteOne(s *Spec) *SpecDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SpecClient) DeleteOneID(id string) *SpecDeleteOne {
	return &SpecDeleteOne{c.Delete().Where(spec.ID(id))}
}

// Create returns a query builder for Spec.
func (c *SpecClient) Query() *SpecQuery {
	return &SpecQuery{config: c.config}
}

// Get returns a Spec entity by its id.
func (c *SpecClient) Get(ctx context.Context, id string) (*Spec, error) {
	return c.Query().Where(spec.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SpecClient) GetX(ctx context.Context, id string) *Spec {
	s, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return s
}