package appdefaultgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"
	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

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

var good = goodmwpb.Good{
	EntID:                  uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturer:     uuid.NewString(),
	DevicePowerConsumption: 120,
	DeviceShipmentAt:       uint32(time.Now().Unix() - 1000),
	DevicePosters:          []string{uuid.NewString(), uuid.NewString()},
	CoinTypeID:             uuid.NewString(),
	VendorLocationID:       uuid.NewString(),
	VendorLocationCountry:  uuid.NewString(),
	VendorLocationProvince: uuid.NewString(),
	VendorLocationCity:     uuid.NewString(),
	VendorLocationAddress:  uuid.NewString(),
	VendorBrandName:        uuid.NewString(),
	VendorBrandLogo:        uuid.NewString(),
	GoodType:               types.GoodType_PowerRenting,
	BenefitType:            types.BenefitType_BenefitTypePlatform,
	UnitPrice:              decimal.NewFromInt(123).String(),
	Title:                  uuid.NewString(),
	QuantityUnit:           "TiB",
	QuantityUnitAmount:     "1",
	TestOnly:               true,
	Posters:                []string{uuid.NewString(), uuid.NewString()},
	Labels: []types.GoodLabel{
		types.GoodLabel_GoodLabelInnovationStarter,
		types.GoodLabel_GoodLabelNoviceExclusive,
	},
	GoodTotal:            decimal.NewFromInt(1000).String(),
	GoodLocked:           decimal.NewFromInt(0).String(),
	GoodInService:        decimal.NewFromInt(0).String(),
	GoodWaitStart:        decimal.NewFromInt(0).String(),
	GoodSold:             decimal.NewFromInt(0).String(),
	DeliveryAt:           uint32(time.Now().Unix() + 1000),
	StartAt:              uint32(time.Now().Unix() + 1000),
	BenefitIntervalHours: 24,
	GoodAppReserved:      decimal.NewFromInt(0).String(),
	UnitLockDeposit:      decimal.NewFromInt(1).String(),
}

var appgood = appgoodmwpb.Good{
	EntID:                  uuid.NewString(),
	AppID:                  uuid.NewString(),
	GoodID:                 good.EntID,
	Online:                 false,
	Visible:                false,
	GoodName:               good.Title,
	UnitPrice:              decimal.NewFromInt(125).String(),
	PackagePrice:           decimal.NewFromInt(125).String(),
	DeviceInfoID:           good.DeviceInfoID,
	DeviceType:             good.DeviceType,
	DeviceManufacturer:     good.DeviceManufacturer,
	DevicePowerConsumption: good.DevicePowerConsumption,
	DevicePosters:          good.DevicePosters,
	DeviceShipmentAt:       good.DeviceShipmentAt,
	CoinTypeID:             good.CoinTypeID,
	VendorLocationID:       good.VendorLocationID,
	VendorLocationCountry:  good.VendorLocationCountry,
	VendorBrandName:        good.VendorBrandName,
	VendorBrandLogo:        good.VendorBrandLogo,
	GoodType:               good.GoodType,
	QuantityUnit:           good.QuantityUnit,
	QuantityUnitAmount:     good.QuantityUnitAmount,
	TestOnly:               good.TestOnly,
	Posters:                good.Posters,
	Labels:                 good.Labels,
	BenefitIntervalHours:   good.BenefitIntervalHours,
	GoodTotal:              good.GoodTotal,
	GoodSpotQuantity:       good.GoodTotal,
	CancelModeStr:          types.CancelMode_Uncancellable.String(),
	EnableSetCommission:    true,
	EnablePurchase:         true,
	EnableProductPage:      true,
	Score:                  decimal.NewFromInt(0).String(),
	StartAt:                good.StartAt,
	BenefitType:            good.BenefitType,
	BenefitTypeStr:         good.BenefitType.String(),
	GoodTypeStr:            good.GoodType.String(),
}

var ret = npool.Default{
	EntID:      uuid.NewString(),
	AppID:      appgood.AppID,
	GoodID:     good.EntID,
	AppGoodID:  appgood.EntID,
	CoinTypeID: good.CoinTypeID,
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodName = good.Title
	ret.AppGoodName = appgood.GoodName

	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithName(&good.VendorBrandName, true),
		vendorbrand1.WithLogo(&good.VendorBrandLogo, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateBrand(context.Background())
	assert.Nil(t, err)
	h1.ID = &info1.ID

	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithEntID(&good.VendorLocationID, true),
		vendorlocation1.WithCountry(&good.VendorLocationCountry, true),
		vendorlocation1.WithProvince(&good.VendorLocationProvince, true),
		vendorlocation1.WithCity(&good.VendorLocationCity, true),
		vendorlocation1.WithAddress(&good.VendorLocationAddress, true),
		vendorlocation1.WithBrandID(&info1.EntID, true),
	)
	assert.Nil(t, err)

	info2, err := h2.CreateLocation(context.Background())
	assert.Nil(t, err)
	h2.ID = &info2.ID

	h3, err := deviceinfo1.NewHandler(
		context.Background(),
		deviceinfo1.WithEntID(&good.DeviceInfoID, true),
		deviceinfo1.WithType(&good.DeviceType, true),
		deviceinfo1.WithManufacturer(&good.DeviceManufacturer, true),
		deviceinfo1.WithPowerConsumption(&good.DevicePowerConsumption, true),
		deviceinfo1.WithShipmentAt(&good.DeviceShipmentAt, true),
		deviceinfo1.WithPosters(good.DevicePosters, true),
	)
	assert.Nil(t, err)

	info3, err := h3.CreateDeviceInfo(context.Background())
	assert.Nil(t, err)
	h3.ID = &info3.ID

	h4, err := good1.NewHandler(
		context.Background(),
		good1.WithEntID(&good.EntID, true),
		good1.WithDeviceInfoID(&good.DeviceInfoID, true),
		good1.WithCoinTypeID(&good.CoinTypeID, true),
		good1.WithVendorLocationID(&good.VendorLocationID, true),
		good1.WithUnitPrice(&good.UnitPrice, true),
		good1.WithBenefitType(&good.BenefitType, true),
		good1.WithGoodType(&good.GoodType, true),
		good1.WithTitle(&good.Title, true),
		good1.WithQuantityUnit(&good.QuantityUnit, true),
		good1.WithQuantityUnitAmount(&good.QuantityUnitAmount, true),
		good1.WithDeliveryAt(&good.DeliveryAt, true),
		good1.WithStartAt(&good.StartAt, true),
		good1.WithTestOnly(&good.TestOnly, false),
		good1.WithBenefitIntervalHours(&good.BenefitIntervalHours, true),
		good1.WithUnitLockDeposit(&good.UnitLockDeposit, false),
		good1.WithTotal(&good.GoodTotal, true),
		good1.WithPosters(good.Posters, false),
		good1.WithLabels(good.Labels, false),
	)
	assert.Nil(t, err)

	info4, err := h4.CreateGood(context.Background())
	assert.Nil(t, err)
	h4.ID = &info4.ID

	h5, err := appgood1.NewHandler(
		context.Background(),
		appgood1.WithEntID(&appgood.EntID, true),
		appgood1.WithAppID(&appgood.AppID, true),
		appgood1.WithGoodID(&appgood.GoodID, true),
		appgood1.WithUnitPrice(&appgood.UnitPrice, true),
		appgood1.WithPackagePrice(&appgood.PackagePrice, true),
		appgood1.WithMinOrderDuration(&appgood.MinOrderDuration, true),
		appgood1.WithMaxOrderDuration(&appgood.MaxOrderDuration, true),
		appgood1.WithGoodName(&appgood.GoodName, true),
	)
	assert.Nil(t, err)

	info5, err := h5.CreateGood(context.Background())
	assert.Nil(t, err)
	h5.ID = &info5.ID

	return func(*testing.T) {
		_, _ = h5.DeleteGood(context.Background())
		_, _ = h4.DeleteGood(context.Background())
		_, _ = h3.DeleteDeviceInfo(context.Background())
		_, _ = h2.DeleteLocation(context.Background())
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func createDefault(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateDefault(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updateDefault(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppGoodID(&ret.AppGoodID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateDefault(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getDefault(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetDefault(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getDefaults(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:      &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			GoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			GoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetDefaults(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteDefault(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteDefault(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetDefault(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestDefault(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createDefault", createDefault)
	t.Run("updateDefault", updateDefault)
	t.Run("getDefault", getDefault)
	t.Run("getDefaults", getDefaults)
	t.Run("deleteDefault", deleteDefault)
}
