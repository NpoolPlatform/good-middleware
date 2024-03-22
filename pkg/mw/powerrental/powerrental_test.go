package powerrental

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
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
		EntID:              uuid.NewString(),
		GoodID:             uuid.NewString(),
		DeviceTypeID:       uuid.NewString(),
		VendorLocationID:   uuid.NewString(),
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
	}
)

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.DurationTypeStr = ret.DurationType.String()
	ret.StartModeStr = ret.StartMode.String()

	h1, err := devicetype1.NewHandler(
		context.Background(),
		devicetype1.WithEntID(&ret.DeviceTypeID, true),
		devicetype1.WithType(&ret.)
	)

	return func(*testing.T) {}
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
	assert.Nil(t, err)

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
	assert.Nil(t, err)
}

func TestPowerRental(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createPowerRental", createPowerRental)
}
