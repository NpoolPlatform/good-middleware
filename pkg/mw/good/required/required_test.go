//nolint:dupl
package required

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
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"

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

var mainGood = goodmwpb.Good{
	ID:                     uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturer:     uuid.NewString(),
	DevicePowerConsumption: 120,
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

var requiredGood = goodmwpb.Good{
	ID:                     uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturer:     uuid.NewString(),
	DevicePowerConsumption: 120,
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

var ret = npool.Required{
	ID:               uuid.NewString(),
	MainGoodID:       mainGood.ID,
	MainGoodName:     mainGood.Title,
	RequiredGoodID:   requiredGood.ID,
	RequiredGoodName: requiredGood.Title,
	Must:             true,
}

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithName(&mainGood.VendorBrandName, true),
		vendorbrand1.WithLogo(&mainGood.VendorBrandLogo, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithID(&mainGood.VendorLocationID, true),
		vendorlocation1.WithCountry(&mainGood.VendorLocationCountry, true),
		vendorlocation1.WithProvince(&mainGood.VendorLocationProvince, true),
		vendorlocation1.WithCity(&mainGood.VendorLocationCity, true),
		vendorlocation1.WithAddress(&mainGood.VendorLocationAddress, true),
		vendorlocation1.WithBrandID(&info1.ID, true),
	)
	assert.Nil(t, err)

	_, err = h2.CreateLocation(context.Background())
	assert.Nil(t, err)

	h3, err := deviceinfo1.NewHandler(
		context.Background(),
		deviceinfo1.WithID(&mainGood.DeviceInfoID, true),
		deviceinfo1.WithType(&mainGood.DeviceType, true),
		deviceinfo1.WithManufacturer(&mainGood.DeviceManufacturer, true),
		deviceinfo1.WithPowerConsumption(&mainGood.DevicePowerConsumption, true),
		deviceinfo1.WithShipmentAt(&mainGood.DeviceShipmentAt, true),
		deviceinfo1.WithPosters(mainGood.DevicePosters, true),
	)
	assert.Nil(t, err)

	_, err = h3.CreateDeviceInfo(context.Background())
	assert.Nil(t, err)

	h4, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&mainGood.ID, true),
		good1.WithDeviceInfoID(&mainGood.DeviceInfoID, true),
		good1.WithDurationDays(&mainGood.DurationDays, true),
		good1.WithCoinTypeID(&mainGood.CoinTypeID, true),
		good1.WithVendorLocationID(&mainGood.VendorLocationID, true),
		good1.WithPrice(&mainGood.Price, true),
		good1.WithBenefitType(&mainGood.BenefitType, true),
		good1.WithGoodType(&mainGood.GoodType, true),
		good1.WithTitle(&mainGood.Title, true),
		good1.WithUnit(&mainGood.Unit, true),
		good1.WithUnitAmount(&mainGood.UnitAmount, true),
		good1.WithSupportCoinTypeIDs(mainGood.SupportCoinTypeIDs, false),
		good1.WithDeliveryAt(&mainGood.DeliveryAt, true),
		good1.WithStartAt(&mainGood.StartAt, true),
		good1.WithTestOnly(&mainGood.TestOnly, false),
		good1.WithBenefitIntervalHours(&mainGood.BenefitIntervalHours, true),
		good1.WithUnitLockDeposit(&mainGood.UnitLockDeposit, false),
		good1.WithTotal(&mainGood.GoodTotal, true),
		good1.WithPosters(mainGood.Posters, false),
		good1.WithLabels(mainGood.Labels, false),
	)
	assert.Nil(t, err)

	_, err = h4.CreateGood(context.Background())
	assert.Nil(t, err)

	h5, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithName(&requiredGood.VendorBrandName, true),
		vendorbrand1.WithLogo(&requiredGood.VendorBrandLogo, true),
	)
	assert.Nil(t, err)

	info2, err := h5.CreateBrand(context.Background())
	assert.Nil(t, err)

	h6, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithID(&requiredGood.VendorLocationID, true),
		vendorlocation1.WithCountry(&requiredGood.VendorLocationCountry, true),
		vendorlocation1.WithProvince(&requiredGood.VendorLocationProvince, true),
		vendorlocation1.WithCity(&requiredGood.VendorLocationCity, true),
		vendorlocation1.WithAddress(&requiredGood.VendorLocationAddress, true),
		vendorlocation1.WithBrandID(&info2.ID, true),
	)
	assert.Nil(t, err)

	_, err = h6.CreateLocation(context.Background())
	assert.Nil(t, err)

	h7, err := deviceinfo1.NewHandler(
		context.Background(),
		deviceinfo1.WithID(&requiredGood.DeviceInfoID, true),
		deviceinfo1.WithType(&requiredGood.DeviceType, true),
		deviceinfo1.WithManufacturer(&requiredGood.DeviceManufacturer, true),
		deviceinfo1.WithPowerConsumption(&requiredGood.DevicePowerConsumption, true),
		deviceinfo1.WithShipmentAt(&requiredGood.DeviceShipmentAt, true),
		deviceinfo1.WithPosters(requiredGood.DevicePosters, true),
	)
	assert.Nil(t, err)

	_, err = h7.CreateDeviceInfo(context.Background())
	assert.Nil(t, err)

	h8, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&requiredGood.ID, true),
		good1.WithDeviceInfoID(&requiredGood.DeviceInfoID, true),
		good1.WithDurationDays(&requiredGood.DurationDays, true),
		good1.WithCoinTypeID(&requiredGood.CoinTypeID, true),
		good1.WithVendorLocationID(&requiredGood.VendorLocationID, true),
		good1.WithPrice(&requiredGood.Price, true),
		good1.WithBenefitType(&requiredGood.BenefitType, true),
		good1.WithGoodType(&requiredGood.GoodType, true),
		good1.WithTitle(&requiredGood.Title, true),
		good1.WithUnit(&requiredGood.Unit, true),
		good1.WithUnitAmount(&requiredGood.UnitAmount, true),
		good1.WithSupportCoinTypeIDs(requiredGood.SupportCoinTypeIDs, false),
		good1.WithDeliveryAt(&requiredGood.DeliveryAt, true),
		good1.WithStartAt(&requiredGood.StartAt, true),
		good1.WithTestOnly(&requiredGood.TestOnly, false),
		good1.WithBenefitIntervalHours(&requiredGood.BenefitIntervalHours, true),
		good1.WithUnitLockDeposit(&requiredGood.UnitLockDeposit, false),
		good1.WithTotal(&requiredGood.GoodTotal, true),
		good1.WithPosters(requiredGood.Posters, false),
		good1.WithLabels(requiredGood.Labels, false),
	)
	assert.Nil(t, err)

	_, err = h8.CreateGood(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h8.DeleteGood(context.Background())
		_, _ = h7.DeleteDeviceInfo(context.Background())
		_, _ = h6.DeleteLocation(context.Background())
		_, _ = h5.DeleteBrand(context.Background())
		_, _ = h4.DeleteGood(context.Background())
		_, _ = h3.DeleteDeviceInfo(context.Background())
		_, _ = h2.DeleteLocation(context.Background())
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func createRequired(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMainGoodID(&ret.MainGoodID, true),
		WithRequiredGoodID(&ret.RequiredGoodID, true),
		WithMust(&ret.Must, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateRequired(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updateRequired(t *testing.T) {
	ret.Must = false

	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithMust(&ret.Must, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateRequired(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getRequired(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetRequired(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getRequireds(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:             &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			MainGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
			RequiredGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RequiredGoodID},
			GoodID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
			GoodIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.MainGoodID, ret.RequiredGoodID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetRequireds(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteRequired(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteRequired(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetRequired(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestRequired(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createRequired", createRequired)
	t.Run("updateRequired", updateRequired)
	t.Run("getRequired", getRequired)
	t.Run("getRequireds", getRequireds)
	t.Run("deleteRequired", deleteRequired)
}
