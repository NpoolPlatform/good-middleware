package topmostgood

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

	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
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
	GoodType:               types.GoodType_PowerRental,
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

var topmost = topmostmwpb.TopMost{
	EntID:                  uuid.NewString(),
	AppID:                  appgood.AppID,
	TopMostType:            types.GoodTopMostType_TopMostNoviceExclusive,
	TopMostTypeStr:         types.GoodTopMostType_TopMostNoviceExclusive.String(),
	Title:                  uuid.NewString(),
	Message:                uuid.NewString(),
	Posters:                []string{uuid.NewString(), uuid.NewString()},
	StartAt:                uint32(time.Now().Unix() + 1000),
	EndAt:                  uint32(time.Now().Unix() + 6000),
	ThresholdCredits:       "0",
	RegisterElapsedSeconds: 3000,
	ThresholdPurchases:     3000,
	ThresholdPaymentAmount: "3000",
	KycMust:                true,
}

var ret = npool.TopMostGood{
	EntID:          uuid.NewString(),
	AppID:          appgood.AppID,
	GoodID:         appgood.GoodID,
	CoinTypeID:     appgood.CoinTypeID,
	AppGoodID:      appgood.EntID,
	TopMostID:      topmost.EntID,
	DisplayIndex:   0,
	Posters:        []string{uuid.NewString(), uuid.NewString()},
	UnitPrice:      decimal.NewFromInt(230).String(),
	PackagePrice:   decimal.NewFromInt(230).String(),
	TopMostTypeStr: topmost.TopMostType.String(),
	TopMostType:    topmost.TopMostType,
	TopMostTitle:   topmost.Title,
	TopMostMessage: topmost.Message,
}

func setup(t *testing.T) func(*testing.T) {
	info1, err := vendorbrand1.CreateBrand(context.Background(), &vendorbrandmwpb.BrandReq{
		Name: &good.VendorBrandName,
		Logo: &good.VendorBrandLogo,
	})
	assert.Nil(t, err)
	assert.NotNil(t, info1)

	vendorLocation, err := vendorlocation1.CreateLocation(context.Background(), &vendorlocationmwpb.LocationReq{
		EntID:    &good.VendorLocationID,
		Country:  &good.VendorLocationCountry,
		Province: &good.VendorLocationProvince,
		City:     &good.VendorLocationCity,
		Address:  &good.VendorLocationAddress,
		BrandID:  &info1.EntID,
	})
	assert.Nil(t, err)

	device, err := devicetype1.CreateDeviceInfo(context.Background(), &devicetypemwpb.DeviceInfoReq{
		EntID:            &good.DeviceInfoID,
		Type:             &good.DeviceType,
		Manufacturer:     &good.DeviceManufacturer,
		PowerConsumption: &good.DevicePowerConsumption,
		ShipmentAt:       &good.DeviceShipmentAt,
		Posters:          good.DevicePosters,
	})
	assert.Nil(t, err)

	_, err = good1.CreateGood(context.Background(), &goodmwpb.GoodReq{
		EntID:                &good.EntID,
		DeviceInfoID:         &good.DeviceInfoID,
		CoinTypeID:           &good.CoinTypeID,
		VendorLocationID:     &good.VendorLocationID,
		UnitPrice:            &good.UnitPrice,
		BenefitType:          &good.BenefitType,
		GoodType:             &good.GoodType,
		Title:                &good.Title,
		QuantityUnit:         &good.QuantityUnit,
		QuantityUnitAmount:   &good.QuantityUnitAmount,
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

	_, err = appgood1.CreateGood(context.Background(), &appgoodmwpb.GoodReq{
		EntID:            &appgood.EntID,
		AppID:            &appgood.AppID,
		GoodID:           &appgood.GoodID,
		UnitPrice:        &appgood.UnitPrice,
		PackagePrice:     &appgood.PackagePrice,
		MinOrderDuration: &appgood.MinOrderDuration,
		MaxOrderDuration: &appgood.MaxOrderDuration,
		GoodName:         &appgood.GoodName,
	})
	assert.Nil(t, err)

	_, err = topmost1.CreateTopMost(context.Background(), &topmostmwpb.TopMostReq{
		EntID:                  &topmost.EntID,
		AppID:                  &topmost.AppID,
		TopMostType:            &topmost.TopMostType,
		Title:                  &topmost.Title,
		Message:                &topmost.Message,
		Posters:                topmost.Posters,
		StartAt:                &topmost.StartAt,
		EndAt:                  &topmost.EndAt,
		ThresholdCredits:       &topmost.ThresholdCredits,
		RegisterElapsedSeconds: &topmost.RegisterElapsedSeconds,
		ThresholdPurchases:     &topmost.ThresholdPurchases,
		ThresholdPaymentAmount: &topmost.ThresholdPaymentAmount,
		KycMust:                &topmost.KycMust,
	})
	assert.Nil(t, err)

	ret.GoodName = good.Title
	ret.AppGoodName = appgood.GoodName

	return func(*testing.T) {
		_, _ = topmost1.DeleteTopMost(context.Background(), topmost.ID)
		_, _ = appgood1.DeleteGood(context.Background(), appgood.ID)
		_, _ = good1.DeleteGood(context.Background(), good.ID)
		_, _ = devicetype1.DeleteDeviceInfo(context.Background(), device.ID)
		_, _ = vendorlocation1.DeleteLocation(context.Background(), vendorLocation.ID)
		_, _ = vendorbrand1.DeleteBrand(context.Background(), info1.ID)
	}
}

func createTopMostGood(t *testing.T) {
	info, err := CreateTopMostGood(context.Background(), &npool.TopMostGoodReq{
		EntID:        &ret.EntID,
		AppGoodID:    &ret.AppGoodID,
		TopMostID:    &ret.TopMostID,
		UnitPrice:    &ret.UnitPrice,
		PackagePrice: &ret.PackagePrice,
		Posters:      ret.Posters,
	})
	if assert.Nil(t, err) {
		ret.PostersStr = info.PostersStr
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, &ret, info)
	}
}

func updateTopMostGood(t *testing.T) {
	info, err := UpdateTopMostGood(context.Background(), &npool.TopMostGoodReq{
		ID:        &ret.ID,
		AppGoodID: &ret.AppGoodID,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getTopMostGood(t *testing.T) {
	info, err := GetTopMostGood(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getTopMostGoods(t *testing.T) {
	infos, total, err := GetTopMostGoods(context.Background(), &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		TopMostID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.TopMostID},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func getTopMostGoodOnly(t *testing.T) {
	info, err := GetTopMostGoodOnly(context.Background(), &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		TopMostID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.TopMostID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func deleteTopMostGood(t *testing.T) {
	info, err := DeleteTopMostGood(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetTopMostGood(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestTopMostGood(t *testing.T) {
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

	t.Run("createTopMostGood", createTopMostGood)
	t.Run("updateTopMostGood", updateTopMostGood)
	t.Run("getTopMostGood", getTopMostGood)
	t.Run("getTopMostGoods", getTopMostGoods)
	t.Run("getTopMostGoodOnly", getTopMostGoodOnly)
	t.Run("deleteTopMostGood", deleteTopMostGood)
}
