package appgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	vendorbrandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
	vendorlocationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

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

var ret = npool.Good{
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

func setup(t *testing.T) func(*testing.T) {
	info1, err := vendorbrand1.CreateBrand(context.Background(), &vendorbrandmwpb.BrandReq{
		Name: &good.VendorBrandName,
		Logo: &good.VendorBrandLogo,
	})
	assert.Nil(t, err)
	assert.NotNil(t, info1)

	_, err = vendorlocation1.CreateLocation(context.Background(), &vendorlocationmwpb.LocationReq{
		ID:       &good.VendorLocationID,
		Country:  &good.VendorLocationCountry,
		Province: &good.VendorLocationProvince,
		City:     &good.VendorLocationCity,
		Address:  &good.VendorLocationAddress,
		BrandID:  &info1.ID,
	})
	assert.Nil(t, err)

	_, err = deviceinfo1.CreateDeviceInfo(context.Background(), &deviceinfomwpb.DeviceInfoReq{
		ID:              &good.DeviceInfoID,
		Type:            &good.DeviceType,
		Manufacturer:    &good.DeviceManufacturer,
		PowerComsuption: &good.DevicePowerComsuption,
		ShipmentAt:      &good.DeviceShipmentAt,
		Posters:         good.DevicePosters,
	})
	assert.Nil(t, err)

	_, err = good1.CreateGood(context.Background(), &goodmwpb.GoodReq{
		ID:                   &good.ID,
		DeviceInfoID:         &good.DeviceInfoID,
		DurationDays:         &good.DurationDays,
		CoinTypeID:           &good.CoinTypeID,
		VendorLocationID:     &good.VendorLocationID,
		Price:                &good.Price,
		BenefitType:          &good.BenefitType,
		GoodType:             &good.GoodType,
		Title:                &good.Title,
		Unit:                 &good.Unit,
		UnitAmount:           &good.UnitAmount,
		SupportCoinTypeIDs:   good.SupportCoinTypeIDs,
		DeliveryAt:           &good.DeliveryAt,
		StartAt:              &good.StartAt,
		TestOnly:             &good.TestOnly,
		BenefitIntervalHours: &good.BenefitIntervalHours,
		UnitLockDeposit:      &good.UnitLockDeposit,
		Total:                &good.GoodTotal,
		Posters:              good.Posters,
		Labels:               good.Labels,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = good1.DeleteGood(context.Background(), good.ID)
		_, _ = deviceinfo1.DeleteDeviceInfo(context.Background(), good.DeviceInfoID)
		_, _ = vendorlocation1.DeleteLocation(context.Background(), good.VendorLocationID)
		_, _ = vendorbrand1.DeleteBrand(context.Background(), info1.ID)
	}
}

/*
func createGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithGoodID(&ret.GoodID, true),
		WithPrice(&ret.Price, true),
		WithGoodName(&ret.GoodName, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateGood(context.Background())
		if assert.Nil(t, err) {
			ret.DevicePostersStr = info.DevicePostersStr
			ret.DisplayColorsStr = info.DisplayColorsStr
			ret.DisplayNamesStr = info.DisplayNamesStr
			ret.DescriptionsStr = info.DescriptionsStr
			ret.LabelsStr = info.LabelsStr
			ret.PostersStr = info.PostersStr
			ret.SupportCoinTypeIDsStr = info.SupportCoinTypeIDsStr
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.AppGoodStockID = info.AppGoodStockID
			ret.AppGoodReserved = info.AppGoodReserved
			ret.AppSpotQuantity = info.AppSpotQuantity
			ret.AppGoodLocked = info.AppGoodLocked
			ret.AppGoodWaitStart = info.AppGoodWaitStart
			ret.AppGoodInService = info.AppGoodInService
			ret.AppGoodSold = info.AppGoodSold
			assert.Equal(t, &ret, info)
		}
	}
}

func updateGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateGood(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetGood(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getGoods(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			AppID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			GoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			GoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
			AppIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetGoods(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteGood(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetGood(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}
*/

func TestGood(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setup(t)
	defer teardown(t)

	// t.Run("createGood", createGood)
	// t.Run("updateGood", updateGood)
	// t.Run("getGood", getGood)
	// t.Run("getGoods", getGoods)
	// t.Run("deleteGood", deleteGood)
}
