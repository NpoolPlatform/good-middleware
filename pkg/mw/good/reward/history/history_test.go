package history

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
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/reward/history"

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

var ret = npool.History{
	ID:         uuid.NewString(),
	GoodID:     good.ID,
	GoodName:   good.Title,
	RewardDate: uint32(time.Now().Unix() - 1000),
	TID:        uuid.NewString(),
	Amount:     decimal.RequireFromString("25.1").String(),
	UnitAmount: decimal.RequireFromString("2.51").String(),
}

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithName(&good.VendorBrandName, true),
		vendorbrand1.WithLogo(&good.VendorBrandLogo, true),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithID(&good.VendorLocationID, true),
		vendorlocation1.WithCountry(&good.VendorLocationCountry, true),
		vendorlocation1.WithProvince(&good.VendorLocationProvince, true),
		vendorlocation1.WithCity(&good.VendorLocationCity, true),
		vendorlocation1.WithAddress(&good.VendorLocationAddress, true),
		vendorlocation1.WithBrandID(&info1.ID, true),
	)
	assert.Nil(t, err)

	_, err = h2.CreateLocation(context.Background())
	assert.Nil(t, err)

	h3, err := deviceinfo1.NewHandler(
		context.Background(),
		deviceinfo1.WithID(&good.DeviceInfoID, true),
		deviceinfo1.WithType(&good.DeviceType, true),
		deviceinfo1.WithManufacturer(&good.DeviceManufacturer, true),
		deviceinfo1.WithPowerConsumption(&good.DevicePowerConsumption, true),
		deviceinfo1.WithShipmentAt(&good.DeviceShipmentAt, true),
		deviceinfo1.WithPosters(good.DevicePosters, true),
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

	_, err = h4.CreateGood(context.Background())
	assert.Nil(t, err)

	state := types.BenefitState_BenefitCheckWait
	h5, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitTransferring
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitCheckTransferring
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitBookKeeping
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitCheckBookKeeping
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitDone
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
		good1.WithRewardAt(&ret.RewardDate, true),
		good1.WithRewardTID(&ret.TID, true),
		good1.WithRewardAmount(&ret.Amount, true),
		good1.WithUnitRewardAmount(&ret.UnitAmount, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitCheckDone
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitWait
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitCheckWait
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitTransferring
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitCheckTransferring
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitBookKeeping
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitCheckBookKeeping
	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	state = types.BenefitState_BenefitDone
	now := uint32(time.Now().Unix() - 500)
	tid := uuid.NewString()

	h5, err = good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithRewardState(&state, true),
		good1.WithRewardAt(&now, true),
		good1.WithRewardTID(&tid, true),
		good1.WithRewardAmount(&ret.Amount, true),
		good1.WithUnitRewardAmount(&ret.UnitAmount, true),
	)
	assert.Nil(t, err)

	_, err = h5.UpdateGood(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h4.DeleteGood(context.Background())
		_, _ = h3.DeleteDeviceInfo(context.Background())
		_, _ = h2.DeleteLocation(context.Background())
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func getHistories(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			GoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			GoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
			RewardDate: &basetypes.Uint32Val{Op: cruder.LTE, Value: uint32(time.Now().Unix())},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetHistories(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(2), total) {
				found := false
				for _, info := range infos {
					if ret.RewardDate == info.RewardDate && ret.TID == info.TID {
						found = true
					}
				}
				assert.Equal(t, true, found)
			}
		}
	}
}

func TestHistory(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("getHistories", getHistories)
}
