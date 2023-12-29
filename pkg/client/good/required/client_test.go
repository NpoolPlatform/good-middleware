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

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
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

var _good1 = goodmwpb.Good{
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
	GoodType:               types.GoodType_PowerRenting,
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

var _good2 = goodmwpb.Good{
	EntID:                  uuid.NewString(),
	DeviceInfoID:           _good1.DeviceInfoID,
	DeviceType:             _good1.DeviceType,
	DeviceManufacturer:     _good1.DeviceManufacturer,
	DevicePowerConsumption: _good1.DevicePowerConsumption,
	DeviceShipmentAt:       _good1.DeviceShipmentAt,
	DevicePosters:          _good1.DevicePosters,
	CoinTypeID:             uuid.NewString(),
	VendorLocationID:       _good1.VendorLocationID,
	VendorLocationCountry:  _good1.VendorLocationCountry,
	VendorLocationProvince: _good1.VendorLocationProvince,
	VendorLocationCity:     _good1.VendorLocationCity,
	VendorLocationAddress:  _good1.VendorLocationAddress,
	VendorBrandName:        _good1.VendorBrandName,
	VendorBrandLogo:        _good1.VendorBrandLogo,
	GoodType:               types.GoodType_PowerRenting,
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

var ret = npool.Required{
	EntID:            uuid.NewString(),
	MainGoodID:       _good1.EntID,
	RequiredGoodID:   _good2.EntID,
	Must:             true,
	MainGoodName:     _good1.Title,
	RequiredGoodName: _good2.Title,
}

func setup(t *testing.T) func(*testing.T) {
	info1, err := vendorbrand1.CreateBrand(context.Background(), &vendorbrandmwpb.BrandReq{
		Name: &_good1.VendorBrandName,
		Logo: &_good1.VendorBrandLogo,
	})
	assert.Nil(t, err)
	assert.NotNil(t, info1)

	vendorLocation, err := vendorlocation1.CreateLocation(context.Background(), &vendorlocationmwpb.LocationReq{
		EntID:    &_good1.VendorLocationID,
		Country:  &_good1.VendorLocationCountry,
		Province: &_good1.VendorLocationProvince,
		City:     &_good1.VendorLocationCity,
		Address:  &_good1.VendorLocationAddress,
		BrandID:  &info1.EntID,
	})
	assert.Nil(t, err)

	device, err := deviceinfo1.CreateDeviceInfo(context.Background(), &deviceinfomwpb.DeviceInfoReq{
		EntID:            &_good1.DeviceInfoID,
		Type:             &_good1.DeviceType,
		Manufacturer:     &_good1.DeviceManufacturer,
		PowerConsumption: &_good1.DevicePowerConsumption,
		ShipmentAt:       &_good1.DeviceShipmentAt,
		Posters:          _good1.DevicePosters,
	})
	assert.Nil(t, err)

	_, err = good1.CreateGood(context.Background(), &goodmwpb.GoodReq{
		EntID:                &_good1.EntID,
		DeviceInfoID:         &_good1.DeviceInfoID,
		CoinTypeID:           &_good1.CoinTypeID,
		VendorLocationID:     &_good1.VendorLocationID,
		UnitPrice:            &_good1.UnitPrice,
		BenefitType:          &_good1.BenefitType,
		GoodType:             &_good1.GoodType,
		Title:                &_good1.Title,
		QuantityUnit:         &_good1.QuantityUnit,
		QuantityUnitAmount:   &_good1.QuantityUnitAmount,
		DeliveryAt:           &_good1.DeliveryAt,
		StartAt:              &_good1.StartAt,
		TestOnly:             &_good1.TestOnly,
		BenefitIntervalHours: &_good1.BenefitIntervalHours,
		UnitLockDeposit:      &_good1.UnitLockDeposit,
		Total:                &_good1.GoodTotal,
		Posters:              _good1.Posters,
		Labels:               _good1.Labels,
	})
	assert.Nil(t, err)

	_, err = good1.CreateGood(context.Background(), &goodmwpb.GoodReq{
		EntID:                &_good2.EntID,
		DeviceInfoID:         &_good2.DeviceInfoID,
		CoinTypeID:           &_good2.CoinTypeID,
		VendorLocationID:     &_good2.VendorLocationID,
		UnitPrice:            &_good2.UnitPrice,
		BenefitType:          &_good2.BenefitType,
		GoodType:             &_good2.GoodType,
		Title:                &_good2.Title,
		QuantityUnit:         &_good2.QuantityUnit,
		QuantityUnitAmount:   &_good2.QuantityUnitAmount,
		DeliveryAt:           &_good2.DeliveryAt,
		StartAt:              &_good2.StartAt,
		TestOnly:             &_good2.TestOnly,
		BenefitIntervalHours: &_good2.BenefitIntervalHours,
		UnitLockDeposit:      &_good2.UnitLockDeposit,
		Total:                &_good2.GoodTotal,
		Posters:              _good2.Posters,
		Labels:               _good2.Labels,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = good1.DeleteGood(context.Background(), _good2.ID)
		_, _ = good1.DeleteGood(context.Background(), _good1.ID)
		_, _ = deviceinfo1.DeleteDeviceInfo(context.Background(), device.ID)
		_, _ = vendorlocation1.DeleteLocation(context.Background(), vendorLocation.ID)
		_, _ = vendorbrand1.DeleteBrand(context.Background(), info1.ID)
	}
}

func createRequired(t *testing.T) {
	info, err := CreateRequired(context.Background(), &npool.RequiredReq{
		EntID:          &ret.EntID,
		MainGoodID:     &ret.MainGoodID,
		RequiredGoodID: &ret.RequiredGoodID,
		Must:           &ret.Must,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, &ret, info)
	}
}

func updateRequired(t *testing.T) {
	ret.Must = false
	info, err := UpdateRequired(context.Background(), &npool.RequiredReq{
		ID:   &ret.ID,
		Must: &ret.Must,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
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
		ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
		GoodIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.MainGoodID, ret.RequiredGoodID}},
		MainGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
		RequiredGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RequiredGoodID},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func deleteRequired(t *testing.T) {
	info, err := DeleteRequired(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetRequiredOnly(context.Background(), &npool.Conds{
		ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
		GoodIDs:        &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.MainGoodID, ret.RequiredGoodID}},
		MainGoodID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.MainGoodID},
		RequiredGoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RequiredGoodID},
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
