package appstock

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"
	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
	goodstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/stock"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Stock{
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	GoodID:      uuid.NewString(),
	GoodName:    uuid.NewString(),
	AppGoodID:   uuid.NewString(),
	AppGoodName: uuid.NewString(),
	Reserved:    "100",
	Locked:      decimal.NewFromInt(0).String(),
	InService:   decimal.NewFromInt(0).String(),
	WaitStart:   decimal.NewFromInt(0).String(),
	Sold:        decimal.NewFromInt(0).String(),
}

var lockID = uuid.NewString()

func setup(t *testing.T) func(*testing.T) {
	brandID := uuid.NewString()
	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithEntID(&brandID, true),
		vendorbrand1.WithName(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorbrand1.WithLogo(func() *string { s := uuid.NewString(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	locationID := uuid.NewString()
	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithEntID(&locationID, true),
		vendorlocation1.WithCountry(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithProvince(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithCity(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithAddress(func() *string { s := uuid.NewString(); return &s }(), true),
		vendorlocation1.WithBrandID(&brandID, true),
	)
	assert.Nil(t, err)

	err = h2.CreateLocation(context.Background())
	assert.Nil(t, err)

	manufacturerID := uuid.NewString()
	h31, err := manufacturer1.NewHandler(
		context.Background(),
		manufacturer1.WithEntID(&manufacturerID, true),
		manufacturer1.WithName(func() *string { s := uuid.NewString(); return &s }(), true),
		manufacturer1.WithLogo(func() *string { s := uuid.NewString(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h31.CreateManufacturer(context.Background())
	assert.Nil(t, err)

	deviceTypeID := uuid.NewString()
	h3, err := devicetype1.NewHandler(
		context.Background(),
		devicetype1.WithEntID(&deviceTypeID, true),
		devicetype1.WithType(func() *string { s := uuid.NewString(); return &s }(), true),
		devicetype1.WithManufacturerID(&manufacturerID, true),
		devicetype1.WithPowerConsumption(func() *uint32 { u := uint32(100); return &u }(), true),
		devicetype1.WithShipmentAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
	)
	assert.Nil(t, err)

	err = h3.CreateDeviceType(context.Background())
	assert.Nil(t, err)

	h4, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&ret.GoodID, true),
		goodbase1.WithBenefitType(func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(), true),
		goodbase1.WithGoodType(func() *types.GoodType { e := types.GoodType_PowerRental; return &e }(), true),
		goodbase1.WithTestOnly(func() *bool { b := true; return &b }(), false),
		goodbase1.WithBenefitIntervalHours(func() *uint32 { u := uint32(24); return &u }(), true),
		goodbase1.WithName(&ret.GoodName, true),
		goodbase1.WithServiceStartAt(func() *uint32 { u := uint32(time.Now().Unix()); return &u }(), true),
		goodbase1.WithStartMode(func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(), true),
	)
	assert.Nil(t, err)

	err = h4.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	goodStockID := uuid.NewString()
	h41, err := goodstock1.NewHandler(
		context.Background(),
		goodstock1.WithEntID(&goodStockID, true),
		goodstock1.WithGoodID(&ret.GoodID, true),
		goodstock1.WithTotal(func() *string { s := decimal.NewFromInt(120).String(); return &s }(), true),
	)
	assert.Nil(t, err)

	err = h41.CreateStock(context.Background())
	assert.Nil(t, err)

	h5, err := appgoodbase1.NewHandler(
		context.Background(),
		appgoodbase1.WithEntID(&ret.AppGoodID, true),
		appgoodbase1.WithAppID(&ret.AppID, true),
		appgoodbase1.WithGoodID(&ret.GoodID, true),
		appgoodbase1.WithName(&ret.AppGoodName, true),
	)
	assert.Nil(t, err)

	err = h5.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	h6, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	assert.Nil(t, err)

	err = h6.createStock(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h6.deleteStock(context.Background())
		_ = h5.DeleteGoodBase(context.Background())
		_ = h41.DeleteStock(context.Background())
		_ = h4.DeleteGoodBase(context.Background())
		_ = h3.DeleteDeviceType(context.Background())
		_ = h31.DeleteManufacturer(context.Background())
		_ = h2.DeleteLocation(context.Background())
		_ = h1.DeleteBrand(context.Background())
	}
}

func reserveStock(t *testing.T) {
	ret.SpotQuantity = ret.Reserved

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithReserved(&ret.Reserved, true),
	)
	if assert.Nil(t, err) {
		err = handler.ReserveStock(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetStock(context.Background())
			if assert.Nil(t, err) {
				ret.CreatedAt = info.CreatedAt
				ret.UpdatedAt = info.UpdatedAt
				ret.ID = info.ID
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func lockStock(t *testing.T) {
	ret.Locked = decimal.NewFromInt(10).String()
	ret.SpotQuantity = decimal.NewFromInt(90).String()

	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithLocked(&ret.Locked, true),
		WithLockID(&lockID, true),
		WithAppSpotLocked(&ret.Locked, true),
	)
	if assert.Nil(t, err) {
		err = handler.LockStock(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetStock(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func waitStartStock(t *testing.T) {
	ret.WaitStart = ret.Locked
	ret.Sold = ret.Locked
	ret.Locked = decimal.NewFromInt(0).String()

	handler, err := NewHandler(
		context.Background(),
		WithLockID(&lockID, true),
	)
	if assert.Nil(t, err) {
		err = handler.WaitStartStock(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetStock(context.Background())
			if assert.Nil(t, err) {
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func lockFailStock(t *testing.T) {
	locked := decimal.NewFromInt(1999).String()
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithLocked(&locked, true),
		WithLockID(&lockID, true),
	)
	if assert.Nil(t, err) {
		err = handler.LockStock(context.Background())
		assert.NotNil(t, err)
	}
}

func chargeBackStock(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithLockID(&lockID, true),
	)
	if assert.Nil(t, err) {
		err = handler.ChargeBackStock(context.Background())
		if assert.Nil(t, err) {
			info, err := handler.GetStock(context.Background())
			if assert.Nil(t, err) {
				ret.WaitStart = decimal.NewFromInt(0).String()
				ret.Sold = decimal.NewFromInt(0).String()
				ret.UpdatedAt = info.UpdatedAt
				assert.Equal(t, &ret, info)
			}
		}
	}
}

func TestStock(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("reserveStock", reserveStock)
	// t.Run("lockStock", lockStock)
	// t.Run("waitStartStock", waitStartStock)
	// t.Run("lockFailStock", lockFailStock)
	// t.Run("chargeBackStock", chargeBackStock)
}
