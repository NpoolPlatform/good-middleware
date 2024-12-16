// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	stdsql "database/sql"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
)

// Option function to configure the client.
type Option func(*config)

// Config is the configuration for the client and its builder.
type config struct {
	// driver used for executing database requests.
	driver dialect.Driver
	// debug enable a debug logging.
	debug bool
	// log used for logging on debug mode.
	log func(...interface{})
	// hooks to execute on mutations.
	hooks *hooks
}

// hooks per client, for fast access.
type hooks struct {
	AppDefaultGood         []ent.Hook
	AppFee                 []ent.Hook
	AppGood                []ent.Hook
	AppGoodBase            []ent.Hook
	AppGoodDescription     []ent.Hook
	AppGoodDisplayColor    []ent.Hook
	AppGoodDisplayName     []ent.Hook
	AppGoodLabel           []ent.Hook
	AppGoodPoster          []ent.Hook
	AppLegacyPowerRental   []ent.Hook
	AppMiningGoodStock     []ent.Hook
	AppPledge              []ent.Hook
	AppPowerRental         []ent.Hook
	AppSimulatePowerRental []ent.Hook
	AppStock               []ent.Hook
	AppStockLock           []ent.Hook
	Comment                []ent.Hook
	DelegatedStaking       []ent.Hook
	DeviceInfo             []ent.Hook
	DeviceManufacturer     []ent.Hook
	DevicePoster           []ent.Hook
	ExtraInfo              []ent.Hook
	FbmCrowdFunding        []ent.Hook
	Fee                    []ent.Hook
	Good                   []ent.Hook
	GoodBase               []ent.Hook
	GoodCoin               []ent.Hook
	GoodCoinReward         []ent.Hook
	GoodMalfunction        []ent.Hook
	GoodReward             []ent.Hook
	GoodRewardHistory      []ent.Hook
	Like                   []ent.Hook
	MiningGoodStock        []ent.Hook
	Pledge                 []ent.Hook
	PowerRental            []ent.Hook
	Recommend              []ent.Hook
	RequiredAppGood        []ent.Hook
	RequiredGood           []ent.Hook
	Score                  []ent.Hook
	Stock                  []ent.Hook
	TopMost                []ent.Hook
	TopMostConstraint      []ent.Hook
	TopMostGood            []ent.Hook
	TopMostGoodConstraint  []ent.Hook
	TopMostGoodPoster      []ent.Hook
	TopMostPoster          []ent.Hook
	VendorBrand            []ent.Hook
	VendorLocation         []ent.Hook
}

// Options applies the options on the config object.
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
func Log(fn func(...interface{})) Option {
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

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...interface{}) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...interface{}) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...interface{}) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...interface{}) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
