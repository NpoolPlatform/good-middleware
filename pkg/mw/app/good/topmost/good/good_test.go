package topmostgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"
	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost"
	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
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
	GoodAppLocked:        decimal.NewFromInt(0).String(),
	UnitLockDeposit:      decimal.NewFromInt(1).String(),
}

var appgood = appgoodmwpb.Good{
	ID:                    uuid.NewString(),
	AppID:                 uuid.NewString(),
	GoodID:                good.ID,
	Online:                false,
	Visible:               false,
	GoodName:              good.Title,
	Price:                 decimal.NewFromInt(125).String(),
	DeviceInfoID:          good.DeviceInfoID,
	DeviceType:            good.DeviceType,
	DeviceManufacturer:    good.DeviceManufacturer,
	DevicePowerComsuption: good.DevicePowerComsuption,
	DevicePosters:         good.DevicePosters,
	DeviceShipmentAt:      good.DeviceShipmentAt,
	DurationDays:          good.DurationDays,
	CoinTypeID:            good.CoinTypeID,
	VendorLocationID:      good.VendorLocationID,
	VendorLocationCountry: good.VendorLocationCountry,
	VendorBrandName:       good.VendorBrandName,
	VendorBrandLogo:       good.VendorBrandLogo,
	GoodType:              good.GoodType,
	Unit:                  good.Unit,
	UnitAmount:            good.UnitAmount,
	SupportCoinTypeIDs:    good.SupportCoinTypeIDs,
	TestOnly:              good.TestOnly,
	Posters:               good.Posters,
	Labels:                good.Labels,
	BenefitIntervalHours:  good.BenefitIntervalHours,
	GoodTotal:             good.GoodTotal,
	GoodLocked:            good.GoodLocked,
	GoodInService:         good.GoodInService,
	GoodWaitStart:         good.GoodWaitStart,
	GoodSold:              good.GoodSold,
	CancelModeStr:         types.CancelMode_Uncancellable.String(),
	PurchaseLimit:         3000,
	EnableSetCommission:   true,
	EnablePurchase:        true,
	EnableProductPage:     true,
	Score:                 decimal.NewFromInt(0).String(),
	StartAt:               good.StartAt,
	BenefitType:           good.BenefitType,
	BenefitTypeStr:        good.BenefitType.String(),
	GoodTypeStr:           good.GoodType.String(),
	UserPurchaseLimit:     decimal.NewFromInt(0).String(),
}

var topmost = topmostmwpb.TopMost{
	ID:                     uuid.NewString(),
	AppID:                  appgood.AppID,
	TopMostType:            types.GoodTopMostType_TopMostInnovationStarter,
	TopMostTypeStr:         types.GoodTopMostType_TopMostInnovationStarter.String(),
	Title:                  uuid.NewString(),
	Message:                uuid.NewString(),
	Posters:                []string{uuid.NewString(), uuid.NewString()},
	StartAt:                uint32(time.Now().Unix() + 1000),
	EndAt:                  uint32(time.Now().Unix() + 6000),
	ThresholdCredits:       uuid.NewString(),
	RegisterElapsedSeconds: 3000,
	ThresholdPurchases:     3000,
	ThresholdPaymentAmount: "3000",
	KycMust:                true,
}

var ret = npool.TopMostGood{
	ID:           uuid.NewString(),
	AppID:        topmost.AppID,
	GoodID:       good.ID,
	AppGoodID:    appgood.ID,
	TopMostID:    topmost.ID,
	DisplayIndex: 1,
	Posters:      []string{uuid.NewString(), uuid.NewString()},
	Price:        decimal.NewFromInt(234).String(),
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

	_, err = h4.CreateGood(context.Background())
	assert.Nil(t, err)

	h5, err := appgood1.NewHandler(
		context.Background(),
		appgood1.WithID(&appgood.ID, true),
		appgood1.WithAppID(&appgood.AppID, true),
		appgood1.WithGoodID(&appgood.GoodID, true),
		appgood1.WithPrice(&appgood.Price, true),
		appgood1.WithGoodName(&appgood.GoodName, true),
	)
	assert.Nil(t, err)

	_, err = h5.CreateGood(context.Background())
	assert.Nil(t, err)

	h6, err := topmost1.NewHandler(
		context.Background(),
		topmost1.WithID(&topmost.ID, true),
		topmost1.WithAppID(&topmost.AppID, true),
		topmost1.WithTopMostType(&topmost.TopMostType, true),
		topmost1.WithTitle(&topmost.Title, true),
		topmost1.WithMessage(&topmost.Message, true),
		topmost1.WithPosters(topmost.Posters, true),
		topmost1.WithStartAt(&topmost.StartAt, true),
		topmost1.WithEndAt(&topmost.EndAt, true),
		topmost1.WithThresholdCredits(&topmost.ThresholdCredits, true),
		topmost1.WithRegisterElapsedSeconds(&topmost.RegisterElapsedSeconds, true),
		topmost1.WithThresholdPurchases(&topmost.ThresholdPurchases, true),
		topmost1.WithThresholdPaymentAmount(&topmost.ThresholdPaymentAmount, true),
		topmost1.WithKycMust(&topmost.KycMust, true),
	)
	assert.Nil(t, err)

	_, err = h6.CreateTopMost(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h6.DeleteTopMost(context.Background())
		_, _ = h5.DeleteGood(context.Background())
		_, _ = h4.DeleteGood(context.Background())
		_, _ = h3.DeleteDeviceInfo(context.Background())
		_, _ = h2.DeleteLocation(context.Background())
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func createTopMostGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithTopMostID(&ret.TopMostID, true),
		WithDisplayIndex(&ret.DisplayIndex, true),
		WithPosters(ret.Posters, true),
		WithPrice(&ret.Price, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateTopMostGood(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func updateTopMostGood(t *testing.T) {
	ret.DisplayIndex = 2
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithTopMostID(&ret.TopMostID, true),
		WithDisplayIndex(&ret.DisplayIndex, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateTopMostGood(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getTopMostGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetTopMostGood(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			assert.Equal(t, &ret, info)
		}
	}
}

func getTopMostGoods(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetTopMostGoods(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				infos[0].PostersStr = ret.PostersStr
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteTopMostGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteTopMostGood(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetTopMostGood(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestTopMostGood(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createTopMostGood", createTopMostGood)
	t.Run("updateTopMostGood", updateTopMostGood)
	t.Run("getTopMostGood", getTopMostGood)
	t.Run("getTopMostGoods", getTopMostGoods)
	t.Run("deleteTopMostGood", deleteTopMostGood)
}
