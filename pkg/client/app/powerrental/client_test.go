package powerrental

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/client/powerrental"
	brand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	location1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
	goodcoinrewardmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin/reward"
	powerrentalmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
	locationmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = &npool.PowerRental{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	GoodID:    uuid.NewString(),
	AppGoodID: uuid.NewString(),

	DeviceTypeID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturerName: uuid.NewString(),
	DeviceManufacturerLogo: uuid.NewString(),
	DevicePowerConsumption: 123,
	DeviceShipmentAt:       uint32(time.Now().Unix()),

	VendorLocationID: uuid.NewString(),
	VendorBrand:      uuid.NewString(),
	VendorLogo:       uuid.NewString(),
	VendorCountry:    uuid.NewString(),
	VendorProvince:   uuid.NewString(),

	UnitPrice:           decimal.NewFromInt(120).String(),
	QuantityUnit:        "TiB",
	QuantityUnitAmount:  decimal.NewFromInt(10).String(),
	DeliveryAt:          uint32(time.Now().Unix()),
	UnitLockDeposit:     decimal.NewFromInt(100).String(),
	DurationDisplayType: types.GoodDurationType_GoodDurationByDay,

	GoodType:             types.GoodType_PowerRental,
	BenefitType:          types.BenefitType_BenefitTypePlatform,
	GoodName:             uuid.NewString(),
	GoodServiceStartAt:   uint32(time.Now().Unix()),
	GoodStartMode:        types.GoodStartMode_GoodStartModeInstantly,
	TestOnly:             true,
	BenefitIntervalHours: 24,
	GoodPurchasable:      true,
	GoodOnline:           true,
	StockMode:            types.GoodStockMode_GoodStockByUnique,

	AppGoodPurchasable:           true,
	AppGoodOnline:                true,
	EnableProductPage:            true,
	ProductPage:                  uuid.NewString(),
	Visible:                      true,
	AppGoodName:                  uuid.NewString(),
	DisplayIndex:                 1,
	Banner:                       uuid.NewString(),
	CancelMode:                   types.CancelMode_Uncancellable,
	CancelableBeforeStartSeconds: 100000,
	EnableSetCommission:          true,
	MinOrderAmount:               decimal.NewFromInt(10).String(),
	MaxOrderAmount:               decimal.NewFromInt(100).String(),
	MaxUserAmount:                decimal.NewFromInt(200).String(),
	MinOrderDurationSeconds:      uint32(86400),
	MaxOrderDurationSeconds:      uint32(86400),
	SaleStartAt:                  uint32(time.Now().Unix()),
	SaleEndAt:                    uint32(time.Now().Unix() + 10000),
	SaleMode:                     types.GoodSaleMode_GoodSaleModeMainnetSpot,
	FixedDuration:                true,
	AppGoodServiceStartAt:        uint32(time.Now().Unix()),
	AppGoodStartMode:             types.GoodStartMode_GoodStartModeInstantly,
	PackageWithRequireds:         true,

	TechniqueFeeRatio: decimal.NewFromInt(0).String(),

	GoodStockID:      uuid.NewString(),
	GoodTotal:        decimal.NewFromInt(200).String(),
	GoodSpotQuantity: decimal.NewFromInt(200).String(),

	AppGoodStockID:      uuid.NewString(),
	AppGoodReserved:     decimal.NewFromInt(0).String(),
	AppGoodSpotQuantity: decimal.NewFromInt(0).String(),
	AppGoodLocked:       decimal.NewFromInt(0).String(),
	AppGoodWaitStart:    decimal.NewFromInt(0).String(),
	AppGoodInService:    decimal.NewFromInt(0).String(),
	AppGoodSold:         decimal.NewFromInt(0).String(),

	Score: decimal.NewFromInt(0).String(),
}

func setup(t *testing.T) func(*testing.T) {
	ret.DurationDisplayTypeStr = ret.DurationDisplayType.String()
	ret.GoodTypeStr = ret.GoodType.String()
	ret.BenefitTypeStr = ret.BenefitType.String()
	ret.StockModeStr = ret.StockMode.String()
	ret.GoodStartModeStr = ret.GoodStartMode.String()
	ret.AppGoodStartModeStr = ret.AppGoodStartMode.String()
	ret.CancelModeStr = ret.CancelMode.String()
	ret.SaleModeStr = ret.SaleMode.String()
	for _, goodCoin := range ret.GoodCoins {
		ret.Rewards = append(ret.Rewards, &goodcoinrewardmwpb.RewardInfo{
			GoodID:                ret.GoodID,
			CoinTypeID:            goodCoin.CoinTypeID,
			RewardTID:             uuid.Nil.String(),
			LastRewardAmount:      decimal.NewFromInt(0).String(),
			NextRewardStartAmount: decimal.NewFromInt(0).String(),
			LastUnitRewardAmount:  decimal.NewFromInt(0).String(),
			TotalRewardAmount:     decimal.NewFromInt(0).String(),
			MainCoin:              goodCoin.Main,
		})
	}

	manufacturerID := uuid.NewString()
	err := manufacturer1.CreateManufacturer(context.Background(), &manufacturermwpb.ManufacturerReq{
		EntID: &manufacturerID,
		Name:  &ret.DeviceManufacturerName,
		Logo:  &ret.DeviceManufacturerLogo,
	})
	assert.Nil(t, err)

	err = devicetype1.CreateDeviceType(context.Background(), &devicetypemwpb.DeviceTypeReq{
		EntID:            &ret.DeviceTypeID,
		Type:             &ret.DeviceType,
		ManufacturerID:   &manufacturerID,
		PowerConsumption: &ret.DevicePowerConsumption,
		ShipmentAt:       &ret.DeviceShipmentAt,
	})
	assert.Nil(t, err)

	brandID := uuid.NewString()
	err = brand1.CreateBrand(context.Background(), &brandmwpb.BrandReq{
		EntID: &brandID,
		Name:  &ret.VendorBrand,
		Logo:  &ret.VendorLogo,
	})
	assert.Nil(t, err)

	err = location1.CreateLocation(context.Background(), &locationmwpb.LocationReq{
		EntID:    &ret.VendorLocationID,
		Country:  &ret.VendorCountry,
		Province: &ret.VendorProvince,
		City:     func() *string { s := uuid.NewString(); return &s }(),
		Address:  func() *string { s := uuid.NewString(); return &s }(),
		BrandID:  &brandID,
	})
	assert.Nil(t, err)

	err = powerrental1.CreatePowerRental(context.Background(), &powerrentalmwpb.PowerRentalReq{
		GoodID:               &ret.GoodID,
		DeviceTypeID:         &ret.DeviceTypeID,
		VendorLocationID:     &ret.VendorLocationID,
		UnitPrice:            &ret.UnitPrice,
		QuantityUnit:         &ret.QuantityUnit,
		QuantityUnitAmount:   &ret.QuantityUnitAmount,
		DeliveryAt:           &ret.DeliveryAt,
		GoodType:             &ret.GoodType,
		BenefitType:          &ret.BenefitType,
		Name:                 &ret.GoodName,
		ServiceStartAt:       &ret.GoodServiceStartAt,
		StartMode:            &ret.GoodStartMode,
		TestOnly:             &ret.TestOnly,
		BenefitIntervalHours: &ret.BenefitIntervalHours,
		StockMode:            &ret.StockMode,
		StockID:              &ret.GoodStockID,
		Total:                &ret.GoodTotal,
		UnitLockDeposit:      &ret.UnitLockDeposit,
		Purchasable:          &ret.GoodPurchasable,
		Online:               &ret.GoodOnline,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = powerrental1.DeletePowerRental(context.Background(), nil, nil, &ret.GoodID)
		_ = location1.DeleteLocation(context.Background(), nil, &ret.VendorLocationID)
		_ = brand1.DeleteBrand(context.Background(), nil, &brandID)
		_ = devicetype1.DeleteDeviceType(context.Background(), nil, &ret.DeviceTypeID)
		_ = manufacturer1.DeleteManufacturer(context.Background(), nil, &manufacturerID)
	}
}

func createPowerRental(t *testing.T) {
	err := CreatePowerRental(context.Background(), &npool.PowerRentalReq{
		EntID:                        &ret.EntID,
		AppID:                        &ret.AppID,
		GoodID:                       &ret.GoodID,
		AppGoodID:                    &ret.AppGoodID,
		Name:                         &ret.AppGoodName,
		UnitPrice:                    &ret.UnitPrice,
		SaleMode:                     &ret.SaleMode,
		ServiceStartAt:               &ret.AppGoodServiceStartAt,
		StartMode:                    &ret.AppGoodStartMode,
		AppGoodStockID:               &ret.AppGoodStockID,
		MinOrderAmount:               &ret.MinOrderAmount,
		MaxOrderAmount:               &ret.MaxOrderAmount,
		MaxUserAmount:                &ret.MaxUserAmount,
		MinOrderDurationSeconds:      &ret.MinOrderDurationSeconds,
		MaxOrderDurationSeconds:      &ret.MaxOrderDurationSeconds,
		EnableProductPage:            &ret.EnableProductPage,
		ProductPage:                  &ret.ProductPage,
		Purchasable:                  &ret.AppGoodPurchasable,
		Online:                       &ret.AppGoodOnline,
		Visible:                      &ret.Visible,
		Banner:                       &ret.Banner,
		DisplayIndex:                 &ret.DisplayIndex,
		CancelableBeforeStartSeconds: &ret.CancelableBeforeStartSeconds,
		EnableSetCommission:          &ret.EnableSetCommission,
		SaleStartAt:                  &ret.SaleStartAt,
		SaleEndAt:                    &ret.SaleEndAt,
	})
	if assert.Nil(t, err) {
		info, err := GetPowerRental(context.Background(), ret.AppGoodID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			ret.State = info.State
			ret.StateStr = info.StateStr
			assert.Equal(t, ret, info)
		}
	}
}

func updatePowerRental(t *testing.T) {
	ret.SaleEndAt += 1000
	err := UpdatePowerRental(context.Background(), &npool.PowerRentalReq{
		ID:        &ret.ID,
		EntID:     &ret.EntID,
		AppID:     &ret.AppID,
		GoodID:    &ret.GoodID,
		AppGoodID: &ret.AppGoodID,
		SaleEndAt: &ret.SaleEndAt,
	})
	if assert.Nil(t, err) {
		info, err := GetPowerRental(context.Background(), ret.AppGoodID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getPowerRental(t *testing.T) {
	info, err := GetPowerRental(context.Background(), ret.AppGoodID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getPowerRentals(t *testing.T) {
	infos, total, err := GetPowerRentals(context.Background(), &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getPowerRentalOnly(t *testing.T) {
	info, err := GetPowerRentalOnly(context.Background(), &npool.Conds{
		ID:        &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		AppGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deletePowerRental(t *testing.T) {
	err := DeletePowerRental(context.Background(), &ret.ID, &ret.EntID, &ret.AppGoodID)
	assert.Nil(t, err)

	info, err := GetPowerRental(context.Background(), ret.AppGoodID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestPowerRental(t *testing.T) {
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

	t.Run("createPowerRental", createPowerRental)
	t.Run("updatePowerRental", updatePowerRental)
	t.Run("getPowerRental", getPowerRental)
	t.Run("getPowerRentals", getPowerRentals)
	t.Run("getPowerRentalOnly", getPowerRentalOnly)
	t.Run("deletePowerRental", deletePowerRental)
}
