// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	stdsql "database/sql"
	"fmt"
	"sync"

	"entgo.io/ent/dialect"
)

// Tx is a transactional client that is created by calling Client.Tx().
type Tx struct {
	config
	// AppDefaultGood is the client for interacting with the AppDefaultGood builders.
	AppDefaultGood *AppDefaultGoodClient
	// AppFee is the client for interacting with the AppFee builders.
	AppFee *AppFeeClient
	// AppGood is the client for interacting with the AppGood builders.
	AppGood *AppGoodClient
	// AppGoodBase is the client for interacting with the AppGoodBase builders.
	AppGoodBase *AppGoodBaseClient
	// AppGoodDescription is the client for interacting with the AppGoodDescription builders.
	AppGoodDescription *AppGoodDescriptionClient
	// AppGoodDisplayColor is the client for interacting with the AppGoodDisplayColor builders.
	AppGoodDisplayColor *AppGoodDisplayColorClient
	// AppGoodDisplayName is the client for interacting with the AppGoodDisplayName builders.
	AppGoodDisplayName *AppGoodDisplayNameClient
	// AppGoodPoster is the client for interacting with the AppGoodPoster builders.
	AppGoodPoster *AppGoodPosterClient
	// AppLegacyPowerRental is the client for interacting with the AppLegacyPowerRental builders.
	AppLegacyPowerRental *AppLegacyPowerRentalClient
	// AppPowerRental is the client for interacting with the AppPowerRental builders.
	AppPowerRental *AppPowerRentalClient
	// AppSimulatePowerRental is the client for interacting with the AppSimulatePowerRental builders.
	AppSimulatePowerRental *AppSimulatePowerRentalClient
	// AppStock is the client for interacting with the AppStock builders.
	AppStock *AppStockClient
	// AppStockLock is the client for interacting with the AppStockLock builders.
	AppStockLock *AppStockLockClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// DelegatedStaking is the client for interacting with the DelegatedStaking builders.
	DelegatedStaking *DelegatedStakingClient
	// DeviceInfo is the client for interacting with the DeviceInfo builders.
	DeviceInfo *DeviceInfoClient
	// DeviceManufacturer is the client for interacting with the DeviceManufacturer builders.
	DeviceManufacturer *DeviceManufacturerClient
	// DevicePoster is the client for interacting with the DevicePoster builders.
	DevicePoster *DevicePosterClient
	// ExtraInfo is the client for interacting with the ExtraInfo builders.
	ExtraInfo *ExtraInfoClient
	// FbmCrowdFunding is the client for interacting with the FbmCrowdFunding builders.
	FbmCrowdFunding *FbmCrowdFundingClient
	// Fee is the client for interacting with the Fee builders.
	Fee *FeeClient
	// Good is the client for interacting with the Good builders.
	Good *GoodClient
	// GoodBase is the client for interacting with the GoodBase builders.
	GoodBase *GoodBaseClient
	// GoodCoin is the client for interacting with the GoodCoin builders.
	GoodCoin *GoodCoinClient
	// GoodReward is the client for interacting with the GoodReward builders.
	GoodReward *GoodRewardClient
	// GoodRewardHistory is the client for interacting with the GoodRewardHistory builders.
	GoodRewardHistory *GoodRewardHistoryClient
	// Like is the client for interacting with the Like builders.
	Like *LikeClient
	// PowerRental is the client for interacting with the PowerRental builders.
	PowerRental *PowerRentalClient
	// Recommend is the client for interacting with the Recommend builders.
	Recommend *RecommendClient
	// RequiredAppGood is the client for interacting with the RequiredAppGood builders.
	RequiredAppGood *RequiredAppGoodClient
	// RequiredGood is the client for interacting with the RequiredGood builders.
	RequiredGood *RequiredGoodClient
	// Score is the client for interacting with the Score builders.
	Score *ScoreClient
	// Stock is the client for interacting with the Stock builders.
	Stock *StockClient
	// TopMost is the client for interacting with the TopMost builders.
	TopMost *TopMostClient
	// TopMostGood is the client for interacting with the TopMostGood builders.
	TopMostGood *TopMostGoodClient
	// VendorBrand is the client for interacting with the VendorBrand builders.
	VendorBrand *VendorBrandClient
	// VendorLocation is the client for interacting with the VendorLocation builders.
	VendorLocation *VendorLocationClient

	// lazily loaded.
	client     *Client
	clientOnce sync.Once

	// completion callbacks.
	mu         sync.Mutex
	onCommit   []CommitHook
	onRollback []RollbackHook

	// ctx lives for the life of the transaction. It is
	// the same context used by the underlying connection.
	ctx context.Context
}

type (
	// Committer is the interface that wraps the Commit method.
	Committer interface {
		Commit(context.Context, *Tx) error
	}

	// The CommitFunc type is an adapter to allow the use of ordinary
	// function as a Committer. If f is a function with the appropriate
	// signature, CommitFunc(f) is a Committer that calls f.
	CommitFunc func(context.Context, *Tx) error

	// CommitHook defines the "commit middleware". A function that gets a Committer
	// and returns a Committer. For example:
	//
	//	hook := func(next ent.Committer) ent.Committer {
	//		return ent.CommitFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Commit(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	CommitHook func(Committer) Committer
)

// Commit calls f(ctx, m).
func (f CommitFunc) Commit(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Commit commits the transaction.
func (tx *Tx) Commit() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Committer = CommitFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Commit()
	})
	tx.mu.Lock()
	hooks := append([]CommitHook(nil), tx.onCommit...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Commit(tx.ctx, tx)
}

// OnCommit adds a hook to call on commit.
func (tx *Tx) OnCommit(f CommitHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onCommit = append(tx.onCommit, f)
}

type (
	// Rollbacker is the interface that wraps the Rollback method.
	Rollbacker interface {
		Rollback(context.Context, *Tx) error
	}

	// The RollbackFunc type is an adapter to allow the use of ordinary
	// function as a Rollbacker. If f is a function with the appropriate
	// signature, RollbackFunc(f) is a Rollbacker that calls f.
	RollbackFunc func(context.Context, *Tx) error

	// RollbackHook defines the "rollback middleware". A function that gets a Rollbacker
	// and returns a Rollbacker. For example:
	//
	//	hook := func(next ent.Rollbacker) ent.Rollbacker {
	//		return ent.RollbackFunc(func(ctx context.Context, tx *ent.Tx) error {
	//			// Do some stuff before.
	//			if err := next.Rollback(ctx, tx); err != nil {
	//				return err
	//			}
	//			// Do some stuff after.
	//			return nil
	//		})
	//	}
	//
	RollbackHook func(Rollbacker) Rollbacker
)

// Rollback calls f(ctx, m).
func (f RollbackFunc) Rollback(ctx context.Context, tx *Tx) error {
	return f(ctx, tx)
}

// Rollback rollbacks the transaction.
func (tx *Tx) Rollback() error {
	txDriver := tx.config.driver.(*txDriver)
	var fn Rollbacker = RollbackFunc(func(context.Context, *Tx) error {
		return txDriver.tx.Rollback()
	})
	tx.mu.Lock()
	hooks := append([]RollbackHook(nil), tx.onRollback...)
	tx.mu.Unlock()
	for i := len(hooks) - 1; i >= 0; i-- {
		fn = hooks[i](fn)
	}
	return fn.Rollback(tx.ctx, tx)
}

// OnRollback adds a hook to call on rollback.
func (tx *Tx) OnRollback(f RollbackHook) {
	tx.mu.Lock()
	defer tx.mu.Unlock()
	tx.onRollback = append(tx.onRollback, f)
}

// Client returns a Client that binds to current transaction.
func (tx *Tx) Client() *Client {
	tx.clientOnce.Do(func() {
		tx.client = &Client{config: tx.config}
		tx.client.init()
	})
	return tx.client
}

func (tx *Tx) init() {
	tx.AppDefaultGood = NewAppDefaultGoodClient(tx.config)
	tx.AppFee = NewAppFeeClient(tx.config)
	tx.AppGood = NewAppGoodClient(tx.config)
	tx.AppGoodBase = NewAppGoodBaseClient(tx.config)
	tx.AppGoodDescription = NewAppGoodDescriptionClient(tx.config)
	tx.AppGoodDisplayColor = NewAppGoodDisplayColorClient(tx.config)
	tx.AppGoodDisplayName = NewAppGoodDisplayNameClient(tx.config)
	tx.AppGoodPoster = NewAppGoodPosterClient(tx.config)
	tx.AppLegacyPowerRental = NewAppLegacyPowerRentalClient(tx.config)
	tx.AppPowerRental = NewAppPowerRentalClient(tx.config)
	tx.AppSimulatePowerRental = NewAppSimulatePowerRentalClient(tx.config)
	tx.AppStock = NewAppStockClient(tx.config)
	tx.AppStockLock = NewAppStockLockClient(tx.config)
	tx.Comment = NewCommentClient(tx.config)
	tx.DelegatedStaking = NewDelegatedStakingClient(tx.config)
	tx.DeviceInfo = NewDeviceInfoClient(tx.config)
	tx.DeviceManufacturer = NewDeviceManufacturerClient(tx.config)
	tx.DevicePoster = NewDevicePosterClient(tx.config)
	tx.ExtraInfo = NewExtraInfoClient(tx.config)
	tx.FbmCrowdFunding = NewFbmCrowdFundingClient(tx.config)
	tx.Fee = NewFeeClient(tx.config)
	tx.Good = NewGoodClient(tx.config)
	tx.GoodBase = NewGoodBaseClient(tx.config)
	tx.GoodCoin = NewGoodCoinClient(tx.config)
	tx.GoodReward = NewGoodRewardClient(tx.config)
	tx.GoodRewardHistory = NewGoodRewardHistoryClient(tx.config)
	tx.Like = NewLikeClient(tx.config)
	tx.PowerRental = NewPowerRentalClient(tx.config)
	tx.Recommend = NewRecommendClient(tx.config)
	tx.RequiredAppGood = NewRequiredAppGoodClient(tx.config)
	tx.RequiredGood = NewRequiredGoodClient(tx.config)
	tx.Score = NewScoreClient(tx.config)
	tx.Stock = NewStockClient(tx.config)
	tx.TopMost = NewTopMostClient(tx.config)
	tx.TopMostGood = NewTopMostGoodClient(tx.config)
	tx.VendorBrand = NewVendorBrandClient(tx.config)
	tx.VendorLocation = NewVendorLocationClient(tx.config)
}

// txDriver wraps the given dialect.Tx with a nop dialect.Driver implementation.
// The idea is to support transactions without adding any extra code to the builders.
// When a builder calls to driver.Tx(), it gets the same dialect.Tx instance.
// Commit and Rollback are nop for the internal builders and the user must call one
// of them in order to commit or rollback the transaction.
//
// If a closed transaction is embedded in one of the generated entities, and the entity
// applies a query, for example: AppDefaultGood.QueryXXX(), the query will be executed
// through the driver which created this transaction.
//
// Note that txDriver is not goroutine safe.
type txDriver struct {
	// the driver we started the transaction from.
	drv dialect.Driver
	// tx is the underlying transaction.
	tx dialect.Tx
}

// newTx creates a new transactional driver.
func newTx(ctx context.Context, drv dialect.Driver) (*txDriver, error) {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return nil, err
	}
	return &txDriver{tx: tx, drv: drv}, nil
}

// Tx returns the transaction wrapper (txDriver) to avoid Commit or Rollback calls
// from the internal builders. Should be called only by the internal builders.
func (tx *txDriver) Tx(context.Context) (dialect.Tx, error) { return tx, nil }

// Dialect returns the dialect of the driver we started the transaction from.
func (tx *txDriver) Dialect() string { return tx.drv.Dialect() }

// Close is a nop close.
func (*txDriver) Close() error { return nil }

// Commit is a nop commit for the internal builders.
// User must call `Tx.Commit` in order to commit the transaction.
func (*txDriver) Commit() error { return nil }

// Rollback is a nop rollback for the internal builders.
// User must call `Tx.Rollback` in order to rollback the transaction.
func (*txDriver) Rollback() error { return nil }

// Exec calls tx.Exec.
func (tx *txDriver) Exec(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Exec(ctx, query, args, v)
}

// Query calls tx.Query.
func (tx *txDriver) Query(ctx context.Context, query string, args, v interface{}) error {
	return tx.tx.Query(ctx, query, args, v)
}

var _ dialect.Driver = (*txDriver)(nil)

// ExecContext allows calling the underlying ExecContext method of the transaction if it is supported by it.
// See, database/sql#Tx.ExecContext for more information.
func (tx *txDriver) ExecContext(ctx context.Context, query string, args ...interface{}) (stdsql.Result, error) {
	ex, ok := tx.tx.(interface {
		ExecContext(context.Context, string, ...interface{}) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the transaction if it is supported by it.
// See, database/sql#Tx.QueryContext for more information.
func (tx *txDriver) QueryContext(ctx context.Context, query string, args ...interface{}) (*stdsql.Rows, error) {
	q, ok := tx.tx.(interface {
		QueryContext(context.Context, string, ...interface{}) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Tx.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
