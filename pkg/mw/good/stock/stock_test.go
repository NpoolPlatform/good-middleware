package stock

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"

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
	ID:                     uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturer:     uuid.NewString(),
	DevicePowerComsuption:  120,
	DeviceShipmentAt:       uint32(time.Now().Unix() - 1000),
	DevicePosters:          []string{uuid.NewString(), uuid.NewString()},
	DurationDays:           14,
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
	Price:                  decimal.NewFromInt(123).String(),
	Title:                  uuid.NewString(),
	Unit:                   "TiB",
	UnitAmount:             1,
	SupportCoinTypeIDs:     []string{uuid.NewString(), uuid.NewString()},
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

var ret = npool.Stock{
	GoodID:      good.ID,
	GoodName:    good.Title,
	Total:       good.GoodTotal,
	Locked:      decimal.NewFromInt(10).String(),
	InService:   decimal.NewFromInt(0).String(),
	WaitStart:   decimal.NewFromInt(10).String(),
	Sold:        decimal.NewFromInt(10).String(),
	AppReserved: decimal.NewFromInt(10).String(),
}

func setup(t *testing.T) func(*testing.T) {
	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithName(&good.VendorBrandName),
		vendorbrand1.WithLogo(&good.VendorBrandLogo),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithID(&good.VendorLocationID),
		vendorlocation1.WithCountry(&good.VendorLocationCountry),
		vendorlocation1.WithProvince(&good.VendorLocationProvince),
		vendorlocation1.WithCity(&good.VendorLocationCity),
		vendorlocation1.WithAddress(&good.VendorLocationAddress),
		vendorlocation1.WithBrandID(&info1.ID),
	)
	assert.Nil(t, err)

	_, err = h2.CreateLocation(context.Background())
	assert.Nil(t, err)

	h3, err := deviceinfo1.NewHandler(
		context.Background(),
		deviceinfo1.WithID(&good.DeviceInfoID),
		deviceinfo1.WithType(&good.DeviceType),
		deviceinfo1.WithManufacturer(&good.DeviceManufacturer),
		deviceinfo1.WithPowerComsuption(&good.DevicePowerComsuption),
		deviceinfo1.WithShipmentAt(&good.DeviceShipmentAt),
		deviceinfo1.WithPosters(good.DevicePosters),
	)
	assert.Nil(t, err)

	_, err = h3.CreateDeviceInfo(context.Background())
	assert.Nil(t, err)

	h4, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithDeviceInfoID(&good.DeviceInfoID, true),
		good1.WithDurationDays(&good.DurationDays, true),
		good1.WithCoinTypeID(&good.CoinTypeID, true),
		good1.WithVendorLocationID(&good.VendorLocationID, true),
		good1.WithPrice(&good.Price, true),
		good1.WithBenefitType(&good.BenefitType, true),
		good1.WithGoodType(&good.GoodType, true),
		good1.WithTitle(&good.Title, true),
		good1.WithUnit(&good.Unit, true),
		good1.WithUnitAmount(&good.UnitAmount, true),
		good1.WithSupportCoinTypeIDs(good.SupportCoinTypeIDs, false),
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

	good, err := h4.CreateGood(context.Background())
	assert.Nil(t, err)

	ret.ID = good.GoodStockID

	return func(*testing.T) {
		_, _ = h4.DeleteGood(context.Background())
		_, _ = h3.DeleteDeviceInfo(context.Background())
		_, _ = h2.DeleteLocation(context.Background())
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func addStock(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithLocked(&ret.Locked, true),
		WithInService(&ret.InService, true),
		WithWaitStart(&ret.WaitStart, true),
		WithAppReserved(&ret.AppReserved, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.AddStock(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}

	ret.InService = ret.WaitStart
	ret.WaitStart = decimal.NewFromInt(-10).String()

	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithInService(&ret.InService, true),
		WithWaitStart(&ret.WaitStart, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.AddStock(context.Background())
		if assert.Nil(t, err) {
			ret.WaitStart = decimal.NewFromInt(0).String()
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}

	h1, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := h1.GetGood(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, ret.Locked, info.GoodLocked)
			assert.Equal(t, ret.InService, info.GoodInService)
			assert.Equal(t, ret.WaitStart, info.GoodWaitStart)
			assert.Equal(t, ret.Sold, info.GoodSold)
			assert.Equal(t, ret.AppReserved, info.GoodAppReserved)
		}
	}

	ret.Locked = decimal.NewFromInt(1999).String()
	handler, err = NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithLocked(&ret.Locked, true),
	)
	if assert.Nil(t, err) {
		_, err := handler.AddStock(context.Background())
		assert.NotNil(t, err)
	}

	ret.Locked = decimal.NewFromInt(10).String()
}

func subStock(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithLocked(&ret.Locked, true),
		WithInService(&ret.InService, true),
		WithWaitStart(&ret.WaitStart, true),
		WithAppReserved(&ret.AppReserved, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.SubStock(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			ret.Locked = decimal.NewFromInt(0).String()
			ret.InService = decimal.NewFromInt(0).String()
			ret.WaitStart = decimal.NewFromInt(0).String()
			ret.Sold = decimal.NewFromInt(0).String()
			ret.AppReserved = decimal.NewFromInt(0).String()
			assert.Equal(t, &ret, info)
		}
	}

	h1, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := h1.GetGood(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, decimal.NewFromInt(0).String(), info.GoodLocked)
		}
	}
}

func TestStock(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("addStock", addStock)
	t.Run("subStock", subStock)
}
