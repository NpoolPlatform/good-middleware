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
	apppowerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	brand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	location1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
	apppowerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
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

var ret = npool.TopMostGood{
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	GoodID:           uuid.NewString(),
	GoodType:         types.GoodType_PowerRental,
	GoodName:         uuid.NewString(),
	TopMostID:        uuid.NewString(),
	TopMostType:      types.GoodTopMostType_TopMostPromotion,
	TopMostTitle:     uuid.NewString(),
	TopMostMessage:   uuid.NewString(),
	TopMostTargetUrl: uuid.NewString(),
	AppGoodID:        uuid.NewString(),
	AppGoodName:      uuid.NewString(),
	DisplayIndex:     0,
	UnitPrice:        decimal.NewFromInt(230).String(),
}

func setup(t *testing.T) func(*testing.T) {
	ret.TopMostTypeStr = ret.TopMostType.String()
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
		Online:               func() *bool { b := true; return &b }(),
	})
	assert.Nil(t, err)

	err = apppowerrental1.CreatePowerRental(context.Background(), &apppowerrentalmwpb.PowerRentalReq{
		AppID:                   &ret.AppID,
		GoodID:                  &ret.GoodID,
		AppGoodID:               &ret.AppGoodID,
		Name:                    &ret.AppGoodName,
		ServiceStartAt:          func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		StartMode:               func() *types.GoodStartMode { e := types.GoodStartMode_GoodStartModeInstantly; return &e }(),
		UnitPrice:               func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		SaleMode:                func() *types.GoodSaleMode { e := types.GoodSaleMode_GoodSaleModeMainnetSpot; return &e }(),
		MinOrderAmount:          func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		MaxOrderAmount:          func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		MinOrderDurationSeconds: func() *uint32 { u := uint32(86400); return &u }(),
		MaxOrderDurationSeconds: func() *uint32 { u := uint32(86400); return &u }(),
	})
	assert.Nil(t, err)

	err = topmost1.CreateTopMost(context.Background(), &topmostmwpb.TopMostReq{
		EntID:       &ret.TopMostID,
		AppID:       &ret.AppID,
		TopMostType: &ret.TopMostType,
		Title:       &ret.TopMostTitle,
		Message:     &ret.TopMostMessage,
		TargetUrl:   &ret.TopMostTargetUrl,
		StartAt:     func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		EndAt:       func() *uint32 { u := uint32(time.Now().Unix() + 1000); return &u }(),
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = topmost1.DeleteTopMost(context.Background(), nil, &ret.TopMostID)
		_ = apppowerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.AppGoodID)
		_ = powerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.GoodID)
		_ = location1.DeleteLocation(context.Background(), nil, &locationID)
		_ = brand1.DeleteBrand(context.Background(), nil, &brandID)
		_ = devicetype1.DeleteDeviceType(context.Background(), nil, &deviceTypeID)
		_ = manufacturer1.DeleteManufacturer(context.Background(), nil, &manufacturerID)
	}
}

func createTopMostGood(t *testing.T) {
	err := CreateTopMostGood(context.Background(), &npool.TopMostGoodReq{
		EntID:     &ret.EntID,
		AppGoodID: &ret.AppGoodID,
		TopMostID: &ret.TopMostID,
		UnitPrice: &ret.UnitPrice,
	})
	if assert.Nil(t, err) {
		info, err := GetTopMostGood(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updateTopMostGood(t *testing.T) {
	err := UpdateTopMostGood(context.Background(), &npool.TopMostGoodReq{
		ID:        &ret.ID,
		AppGoodID: &ret.AppGoodID,
	})
	if assert.Nil(t, err) {
		info, err := GetTopMostGood(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
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
	err := DeleteTopMostGood(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetTopMostGood(context.Background(), ret.EntID)
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
