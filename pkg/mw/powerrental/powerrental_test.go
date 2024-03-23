package powerrental

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.PowerRental{
		EntID:  uuid.NewString(),
		GoodID: uuid.NewString(),

		DeviceTypeID:           uuid.NewString(),
		DeviceType:             uuid.NewString(),
		DeviceManufacturerName: uuid.NewString(),
		DeviceManufacturerLogo: uuid.NewString(),
		DevicePowerConsumption: 120,
		DeviceShipmentAt:       uint32(time.Now().Unix()),

		VendorLocationID: uuid.NewString(),
		VendorBrand:      uuid.NewString(),
		VendorLogo:       uuid.NewString(),
		VendorCountry:    uuid.NewString(),
		VendorProvince:   uuid.NewString(),

		UnitPrice:          decimal.NewFromInt(120).String(),
		QuantityUnit:       "TiB",
		QuantityUnitAmount: decimal.NewFromInt(2).String(),
		DeliveryAt:         uint32(time.Now().Unix()),
		DurationType:       types.GoodDurationType_GoodDurationByDay,

		GoodType:             types.GoodType_PowerRental,
		BenefitType:          types.BenefitType_BenefitTypePlatform,
		Name:                 uuid.NewString(),
		ServiceStartAt:       uint32(time.Now().Unix()),
		StartMode:            types.GoodStartMode_GoodStartModeInstantly,
		BenefitIntervalHours: 20,
		UnitLockDeposit:      decimal.NewFromInt(1).String(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.BenefitTypeStr = ret.BenefitType.String()
	ret.StartModeStr = ret.StartMode.String()
	ret.DurationTypeStr = ret.DurationType.String()
	ret.StartModeStr = ret.StartMode.String()

	manufacturerID := uuid.NewString()
	h1, err := manufacturer1.NewHandler(
		context.Background(),
		manufacturer1.WithEntID(&manufacturerID, true),
		manufacturer1.WithName(&ret.DeviceManufacturerName, true),
		manufacturer1.WithLogo(&ret.DeviceManufacturerLogo, true),
	)
	assert.Nil(t, err)

	err = h1.CreateManufacturer(context.Background())
	assert.Nil(t, err)

	h2, err := devicetype1.NewHandler(
		context.Background(),
		devicetype1.WithEntID(&ret.DeviceTypeID, true),
		devicetype1.WithType(&ret.DeviceType, true),
		devicetype1.WithManufacturerID(&manufacturerID, true),
		devicetype1.WithPowerConsumption(&ret.DevicePowerConsumption, true),
		devicetype1.WithShipmentAt(&ret.DeviceShipmentAt, true),
	)
	assert.Nil(t, err)

	err = h2.CreateDeviceType(context.Background())
	assert.Nil(t, err)

	brandID := uuid.NewString()
	h3, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithEntID(&brandID, true),
		vendorbrand1.WithName(&ret.VendorBrand, true),
		vendorbrand1.WithLogo(&ret.VendorLogo, true),
	)
	assert.Nil(t, err)

	err = h3.CreateBrand(context.Background())
	assert.Nil(t, err)

	h4, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithEntID(&ret.VendorLocationID, true),
		vendorlocation1.WithBrandID(&brandID, true),
		vendorlocation1.WithCountry(&ret.VendorCountry, true),
		vendorlocation1.WithProvince(&ret.VendorProvince, true),
		vendorlocation1.WithCity(&ret.VendorProvince, true),
		vendorlocation1.WithAddress(&ret.VendorProvince, true),
	)
	assert.Nil(t, err)

	err = h4.CreateLocation(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h4.DeleteLocation(context.Background())
		_ = h3.DeleteBrand(context.Background())
		_ = h2.DeleteDeviceType(context.Background())
		_ = h1.DeleteManufacturer(context.Background())
	}
}

func createPowerRental(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithVendorLocationID(&ret.VendorLocationID, true),
		WithUnitPrice(&ret.UnitPrice, true),
		WithQuantityUnit(&ret.QuantityUnit, true),
		WithQuantityUnitAmount(&ret.QuantityUnitAmount, true),
		WithDeliveryAt(&ret.DeliveryAt, true),
		WithUnitLockDeposit(&ret.UnitLockDeposit, true),
		WithDurationType(&ret.DurationType, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
	)
	assert.Nil(t, err)

	err = h1.CreatePowerRental(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetPowerRental(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}

	h2, err := NewHandler(
		context.Background(),
		WithGoodID(&ret.GoodID, true),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithVendorLocationID(&ret.VendorLocationID, true),
		WithUnitPrice(&ret.UnitPrice, true),
		WithQuantityUnit(&ret.QuantityUnit, true),
		WithQuantityUnitAmount(&ret.QuantityUnitAmount, true),
		WithDeliveryAt(&ret.DeliveryAt, true),
		WithUnitLockDeposit(&ret.UnitLockDeposit, true),
		WithDurationType(&ret.DurationType, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
	)
	assert.Nil(t, err)

	err = h2.CreatePowerRental(context.Background())
	assert.NotNil(t, err)
}

func updatePowerRental(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithDeviceTypeID(&ret.DeviceTypeID, true),
		WithVendorLocationID(&ret.VendorLocationID, true),
		WithUnitPrice(&ret.UnitPrice, true),
		WithQuantityUnit(&ret.QuantityUnit, true),
		WithQuantityUnitAmount(&ret.QuantityUnitAmount, true),
		WithDeliveryAt(&ret.DeliveryAt, true),
		WithUnitLockDeposit(&ret.UnitLockDeposit, true),
		WithDurationType(&ret.DurationType, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
	)
	assert.Nil(t, err)

	err = h1.UpdatePowerRental(context.Background())
	assert.Nil(t, err)
}

func getPowerRental(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetPowerRental(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getPowerRentals(t *testing.T) {
	conds := &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetPowerRentals(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deletePowerRental(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	err = handler.DeletePowerRental(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetPowerRental(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestPowerRental(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createPowerRental", createPowerRental)
	t.Run("updatePowerRental", updatePowerRental)
	t.Run("getPowerRental", getPowerRental)
	t.Run("getPowerRentals", getPowerRentals)
	t.Run("deletePowerRental", deletePowerRental)
}
