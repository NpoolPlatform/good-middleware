package recommend

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
	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	brand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	location1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
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

var ret = npool.Recommend{
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	RecommenderID:  uuid.NewString(),
	AppGoodID:      uuid.NewString(),
	GoodName:       uuid.NewString(),
	RecommendIndex: decimal.NewFromInt(4).String(),
	Message:        uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	ret.HideReasonStr = ret.HideReason.String()

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

	goodID := uuid.NewString()
	err = powerrental1.CreatePowerRental(context.Background(), &powerrentalmwpb.PowerRentalReq{
		GoodID:               &goodID,
		DeviceTypeID:         &deviceTypeID,
		VendorLocationID:     &locationID,
		UnitPrice:            func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		QuantityUnit:         func() *string { s := "TiB"; return &s }(),
		QuantityUnitAmount:   func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		DeliveryAt:           func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		GoodType:             func() *types.GoodType { e := types.GoodType_PowerRental; return &e }(),
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
		GoodID:                  &goodID,
		AppGoodID:               &ret.AppGoodID,
		Name:                    &ret.GoodName,
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

	return func(*testing.T) {
		_ = apppowerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.AppGoodID)
		_ = powerrental1.DeletePowerRental(context.Background(), nil, nil, &goodID)
		_ = location1.DeleteLocation(context.Background(), nil, &locationID)
		_ = brand1.DeleteBrand(context.Background(), nil, &brandID)
		_ = devicetype1.DeleteDeviceType(context.Background(), nil, &deviceTypeID)
		_ = manufacturer1.DeleteManufacturer(context.Background(), nil, &manufacturerID)
	}
}

func createRecommend(t *testing.T) {
	err := CreateRecommend(context.Background(), &npool.RecommendReq{
		EntID:          &ret.EntID,
		RecommenderID:  &ret.RecommenderID,
		AppGoodID:      &ret.AppGoodID,
		Message:        &ret.Message,
		RecommendIndex: &ret.RecommendIndex,
	})
	if assert.Nil(t, err) {
		info, err := GetRecommend(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updateRecommend(t *testing.T) {
	err := UpdateRecommend(context.Background(), &npool.RecommendReq{
		ID:      &ret.ID,
		Message: &ret.Message,
	})
	if assert.Nil(t, err) {
		info, err := GetRecommend(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

//nolint
func getRecommends(t *testing.T) {
	infos, total, err := GetRecommends(context.Background(), &npool.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RecommenderID},
		AppGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		AppGoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

//nolint
func getRecommendOnly(t *testing.T) {
	info, err := GetRecommendOnly(context.Background(), &npool.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RecommenderID},
		AppGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		AppGoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

//nolint
func deleteRecommend(t *testing.T) {
	err := DeleteRecommend(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetRecommendOnly(context.Background(), &npool.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RecommenderID},
		AppGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		AppGoodIDs:    &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
	})
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestRecommend(t *testing.T) {
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

	t.Run("createRecommend", createRecommend)
	t.Run("updateRecommend", updateRecommend)
	t.Run("getRecommends", getRecommends)
	t.Run("getRecommendOnly", getRecommendOnly)
	t.Run("deleteRecommend", deleteRecommend)
}
