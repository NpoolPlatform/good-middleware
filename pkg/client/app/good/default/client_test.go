package appdefaultgood

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

	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
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

var ret = npool.Default{
	EntID:      uuid.NewString(),
	AppID:      appgood.AppID,
	GoodID:     appgood.GoodID,
	AppGoodID:  appgood.EntID,
	CoinTypeID: good.CoinTypeID,
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

	ret.GoodName = good.Title
	ret.AppGoodName = appgood.GoodName

	return func(*testing.T) {
		_, _ = appgood1.DeleteGood(context.Background(), appgood.ID)
		_, _ = good1.DeleteGood(context.Background(), good.ID)
		_, _ = devicetype1.DeleteDeviceInfo(context.Background(), device.ID)
		_, _ = vendorlocation1.DeleteLocation(context.Background(), vendorLocation.ID)
		_, _ = vendorbrand1.DeleteBrand(context.Background(), info1.ID)
	}
}

func createDefault(t *testing.T) {
	info, err := CreateDefault(context.Background(), &npool.DefaultReq{
		EntID:     &ret.EntID,
		AppGoodID: &ret.AppGoodID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, &ret, info)
	}
}

func updateDefault(t *testing.T) {
	info, err := UpdateDefault(context.Background(), &npool.DefaultReq{
		ID:        &ret.ID,
		AppGoodID: &ret.AppGoodID,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getDefault(t *testing.T) {
	info, err := GetDefault(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getDefaults(t *testing.T) {
	infos, total, err := GetDefaults(context.Background(), &npool.Conds{
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		GoodIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.CoinTypeID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func getDefaultOnly(t *testing.T) {
	info, err := GetDefaultOnly(context.Background(), &npool.Conds{
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		GoodID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		CoinTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		GoodIDs:     &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
		CoinTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.CoinTypeID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func deleteDefault(t *testing.T) {
	info, err := DeleteDefault(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetDefault(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDefault(t *testing.T) {
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

	t.Run("createDefault", createDefault)
	t.Run("updateDefault", updateDefault)
	t.Run("getDefault", getDefault)
	t.Run("getDefaults", getDefaults)
	t.Run("getDefaultOnly", getDefaultOnly)
	t.Run("deleteDefault", deleteDefault)
}
