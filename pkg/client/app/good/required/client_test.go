//nolint:dupl
package required

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

	appfee1 "github.com/NpoolPlatform/good-middleware/pkg/client/app/fee"
	apppowerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/app/powerrental"
	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	fee1 "github.com/NpoolPlatform/good-middleware/pkg/client/fee"
	goodrequired1 "github.com/NpoolPlatform/good-middleware/pkg/client/good/required"
	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	brand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	location1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appfeemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/required"
	apppowerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
	feemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"
	goodrequiredmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
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

var ret = npool.Required{
	EntID:               uuid.NewString(),
	AppID:               uuid.NewString(),
	MainGoodID:          uuid.NewString(),
	MainGoodName:        uuid.NewString(),
	MainAppGoodID:       uuid.NewString(),
	MainAppGoodName:     uuid.NewString(),
	RequiredGoodID:      uuid.NewString(),
	RequiredGoodName:    uuid.NewString(),
	RequiredAppGoodID:   uuid.NewString(),
	RequiredAppGoodName: uuid.NewString(),
	Must:                false,
}

//nolint:funlen
func setup(t *testing.T) func(*testing.T) {
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
		GoodID:               &ret.MainGoodID,
		DeviceTypeID:         &deviceTypeID,
		VendorLocationID:     &locationID,
		UnitPrice:            func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		QuantityUnit:         func() *string { s := "TiB"; return &s }(),
		QuantityUnitAmount:   func() *string { s := decimal.NewFromInt(120).String(); return &s }(),
		DeliveryAt:           func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		GoodType:             func() *types.GoodType { e := types.GoodType_PowerRental; return &e }(),
		BenefitType:          func() *types.BenefitType { e := types.BenefitType_BenefitTypePlatform; return &e }(),
		Name:                 &ret.MainGoodName,
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
		GoodID:                  &ret.MainGoodID,
		AppGoodID:               &ret.MainAppGoodID,
		Name:                    &ret.MainAppGoodName,
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

	err = fee1.CreateFee(context.Background(), &feemwpb.FeeReq{
		GoodID:              &ret.RequiredGoodID,
		GoodType:            func() *types.GoodType { e := types.GoodType_TechniqueServiceFee; return &e }(),
		Name:                &ret.RequiredGoodName,
		SettlementType:      func() *types.GoodSettlementType { e := types.GoodSettlementType_GoodSettledByProfitPercent; return &e }(),
		UnitValue:           func() *string { s := decimal.NewFromInt(20).String(); return &s }(),
		DurationDisplayType: func() *types.GoodDurationType { e := types.GoodDurationType_GoodDurationByDay; return &e }(),
	})
	assert.Nil(t, err)

	err = appfee1.CreateFee(context.Background(), &appfeemwpb.FeeReq{
		AppID:                   &ret.AppID,
		GoodID:                  &ret.RequiredGoodID,
		AppGoodID:               &ret.RequiredAppGoodID,
		ProductPage:             func() *string { s := uuid.NewString(); return &s }(),
		Name:                    &ret.RequiredAppGoodName,
		Banner:                  func() *string { s := uuid.NewString(); return &s }(),
		UnitValue:               func() *string { s := decimal.NewFromInt(20).String(); return &s }(),
		MinOrderDurationSeconds: func() *uint32 { u := uint32(86400); return &u }(),
	})
	assert.Nil(t, err)

	requiredID := uuid.NewString()
	err = goodrequired1.CreateRequired(context.Background(), &goodrequiredmwpb.RequiredReq{
		EntID:          &requiredID,
		MainGoodID:     &ret.MainGoodID,
		RequiredGoodID: &ret.RequiredGoodID,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = goodrequired1.DeleteRequired(context.Background(), nil, &requiredID)
		_ = appfee1.DeleteFee(context.Background(), nil, nil, &ret.RequiredAppGoodID)
		_ = fee1.DeleteFee(context.Background(), nil, nil, &ret.RequiredGoodID)
		_ = apppowerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.MainAppGoodID)
		_ = powerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.MainGoodID)
		_ = location1.DeleteLocation(context.Background(), nil, &locationID)
		_ = brand1.DeleteBrand(context.Background(), nil, &brandID)
		_ = devicetype1.DeleteDeviceType(context.Background(), nil, &deviceTypeID)
		_ = manufacturer1.DeleteManufacturer(context.Background(), nil, &manufacturerID)
	}
}

func createRequired(t *testing.T) {
	err := CreateRequired(context.Background(), &npool.RequiredReq{
		EntID:             &ret.EntID,
		MainAppGoodID:     &ret.MainAppGoodID,
		RequiredAppGoodID: &ret.RequiredAppGoodID,
		Must:              &ret.Must,
	})
	if assert.Nil(t, err) {
		info, err := GetRequired(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updateRequired(t *testing.T) {
	ret.Must = true
	err := UpdateRequired(context.Background(), &npool.RequiredReq{
		ID:   &ret.ID,
		Must: &ret.Must,
	})
	if assert.Nil(t, err) {
		info, err := GetRequired(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getRequired(t *testing.T) {
	info, err := GetRequired(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getRequireds(t *testing.T) {
	infos, total, err := GetRequireds(context.Background(), &npool.Conds{
		ID:                &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:             &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppGoodID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainAppGoodID},
		AppGoodIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.MainAppGoodID, ret.RequiredAppGoodID}},
		MainAppGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainAppGoodID},
		RequiredAppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RequiredAppGoodID},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func deleteRequired(t *testing.T) {
	err := DeleteRequired(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetRequiredOnly(context.Background(), &npool.Conds{
		ID:                &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:             &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppGoodID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainAppGoodID},
		AppGoodIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.MainAppGoodID, ret.RequiredAppGoodID}},
		MainAppGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainAppGoodID},
		RequiredAppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RequiredAppGoodID},
	})
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestRequired(t *testing.T) {
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

	t.Run("createRequired", createRequired)
	t.Run("updateRequired", updateRequired)
	t.Run("getRequired", getRequired)
	t.Run("getRequireds", getRequireds)
	t.Run("deleteRequired", deleteRequired)
}
