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

	apppowerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	goodcoin1 "github.com/NpoolPlatform/good-middleware/pkg/client/good/coin"
	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	brand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	location1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
	apppowerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
	powerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

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
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	GoodID:      uuid.NewString(),
	GoodName:    uuid.NewString(),
	AppGoodID:   uuid.NewString(),
	AppGoodName: uuid.NewString(),
	CoinTypeID:  uuid.NewString(),
	GoodType:    types.GoodType_PowerRental,
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()

	manufacturerID := uuid.NewString()
	err := manufacturer1.CreateManufacturer(context.Background(), &manufacturermwpb.ManufacturerReq{
		EntID: &manufacturerID,
		Name:  func() *string { s := uuid.NewString(); return &s }(),
		Logo:  func() *string { s := uuid.NewString(); return &s }(),
	})
	assert.Nil(t, err)

	deviceTypeID := uuid.NewString()
	err = devicetype1.CreateDeviceType(context.Background(), &devicetypemwpb.DeviceTypeReq{
		EntID:            &deviceTypeID,
		Type:             func() *string { s := uuid.NewString(); return &s }(),
		ManufacturerID:   &manufacturerID,
		PowerConsumption: func() *uint32 { u := uint32(120); return &u }(),
		ShipmentAt:       func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
	})
	assert.Nil(t, err)

	brandID := uuid.NewString()
	err = brand1.CreateBrand(context.Background(), &brandmwpb.BrandReq{
		EntID: &brandID,
		Name:  func() *string { s := uuid.NewString(); return &s }(),
		Logo:  func() *string { s := uuid.NewString(); return &s }(),
	})
	assert.Nil(t, err)

	locationID := uuid.NewString()
	err = location1.CreateLocation(context.Background(), &locationmwpb.LocationReq{
		EntID:    &locationID,
		Country:  func() *string { s := uuid.NewString(); return &s }(),
		Province: func() *string { s := uuid.NewString(); return &s }(),
		City:     func() *string { s := uuid.NewString(); return &s }(),
		Address:  func() *string { s := uuid.NewString(); return &s }(),
		BrandID:  &brandID,
	})
	assert.Nil(t, err)

	err = powerrental1.CreatePowerRental(context.Background(), &powerrentalmwpb.PowerRentalReq{
		GoodID:               &ret.GoodID,
		DeviceTypeID:         &deviceTypeID,
		VendorLocationID:     &locationID,
		UnitPrice:            func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		QuantityUnit:         func() *string { s := "TiB"; return &s }(),
		QuantityUnitAmount:   func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		DeliveryAt:           func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		GoodType:             &ret.GoodType,
		BenefitType:          func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(),
		Name:                 &ret.GoodName,
		ServiceStartAt:       func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		StartMode:            func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(),
		TestOnly:             func() *bool { b := true; return &b }(),
		BenefitIntervalHours: func() *uint32 { u := uint32(24); return &u }(),
		StockMode:            func() *types.GoodStockMode { e := types.GoodStockMode_GoodStockByUnique; return &e }(),
		Total:                func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
	})
	assert.Nil(t, err)

	goodCoinID := uuid.NewString()
	err = goodcoin1.CreateGoodCoin(context.Background(), &goodcoinmwpb.GoodCoinReq{
		EntID:      &goodCoinID,
		GoodID:     &ret.GoodID,
		CoinTypeID: &ret.CoinTypeID,
	})
	assert.Nil(t, err)

	err = apppowerrental1.CreatePowerRental(context.Background(), &apppowerrentalmwpb.PowerRentalReq{
		AppID:                   &ret.AppID,
		GoodID:                  &ret.GoodID,
		AppGoodID:               &ret.AppGoodID,
		Name:                    &ret.AppGoodName,
		ServiceStartAt:          func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		StartMode:               func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(),
		UnitPrice:               func() *string { s := decimal.NewFromInt(10368000).String(); return &s }(),
		SaleMode:                func() *types.GoodSaleMode { e := types.GoodSaleMode_GoodSaleModeMainnetSpot; return &e }(),
		MinOrderAmount:          func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		MaxOrderAmount:          func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		MinOrderDurationSeconds: func() *uint32 { u := uint32(86400); return &u }(),
		MaxOrderDurationSeconds: func() *uint32 { u := uint32(86400); return &u }(),
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = apppowerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.AppGoodID)
		_ = goodcoin1.DeleteGoodCoin(context.Background(), nil, &goodCoinID)
		_ = powerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.GoodID)
		_ = location1.DeleteLocation(context.Background(), nil, &locationID)
		_ = brand1.DeleteBrand(context.Background(), nil, &brandID)
		_ = devicetype1.DeleteDeviceType(context.Background(), nil, &deviceTypeID)
		_ = manufacturer1.DeleteManufacturer(context.Background(), nil, &manufacturerID)
	}
}

func createDefault(t *testing.T) {
	err := CreateDefault(context.Background(), &npool.DefaultReq{
		EntID:      &ret.EntID,
		AppGoodID:  &ret.AppGoodID,
		CoinTypeID: &ret.CoinTypeID,
	})
	if assert.Nil(t, err) {
		info, err := GetDefault(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updateDefault(t *testing.T) {
	err := UpdateDefault(context.Background(), &npool.DefaultReq{
		ID:        &ret.ID,
		EntID:     &ret.EntID,
		AppGoodID: &ret.AppGoodID,
	})
	if assert.Nil(t, err) {
		info, err := GetDefault(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
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
	err := DeleteDefault(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetDefault(context.Background(), ret.EntID)
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
