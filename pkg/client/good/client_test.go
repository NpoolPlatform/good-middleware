package good

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
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
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

var ret = npool.Good{
	ID:                     uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturer:     uuid.NewString(),
	DevicePowerConsumption: 120,
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
	GoodTotal:             decimal.NewFromInt(1000).String(),
	GoodLocked:            decimal.NewFromInt(0).String(),
	GoodInService:         decimal.NewFromInt(0).String(),
	GoodWaitStart:         decimal.NewFromInt(0).String(),
	GoodSold:              decimal.NewFromInt(0).String(),
	DeliveryAt:            uint32(time.Now().Unix() + 1000),
	StartAt:               uint32(time.Now().Unix() + 1000),
	BenefitIntervalHours:  24,
	GoodAppReserved:       decimal.NewFromInt(0).String(),
	UnitLockDeposit:       decimal.NewFromInt(1).String(),
	Score:                 decimal.NewFromInt(0).String(),
	RewardStateStr:        types.BenefitState_BenefitWait.String(),
	RewardTID:             uuid.Nil.String(),
	NextRewardStartAmount: decimal.NewFromInt(0).String(),
	LastRewardAmount:      decimal.NewFromInt(0).String(),
	LastUnitRewardAmount:  decimal.NewFromInt(0).String(),
	TotalRewardAmount:     decimal.NewFromInt(0).String(),
	GoodTypeStr:           types.GoodType_PowerRenting.String(),
	BenefitTypeStr:        types.BenefitType_BenefitTypePlatform.String(),
}

func setup(t *testing.T) func(*testing.T) {
	info1, err := vendorbrand1.CreateBrand(context.Background(), &vendorbrandmwpb.BrandReq{
		Name: &ret.VendorBrandName,
		Logo: &ret.VendorBrandLogo,
	})
	assert.Nil(t, err)
	assert.NotNil(t, info1)

	_, err = vendorlocation1.CreateLocation(context.Background(), &vendorlocationmwpb.LocationReq{
		ID:       &ret.VendorLocationID,
		Country:  &ret.VendorLocationCountry,
		Province: &ret.VendorLocationProvince,
		City:     &ret.VendorLocationCity,
		Address:  &ret.VendorLocationAddress,
		BrandID:  &info1.ID,
	})
	assert.Nil(t, err)

	_, err = deviceinfo1.CreateDeviceInfo(context.Background(), &deviceinfomwpb.DeviceInfoReq{
		ID:               &ret.DeviceInfoID,
		Type:             &ret.DeviceType,
		Manufacturer:     &ret.DeviceManufacturer,
		PowerConsumption: &ret.DevicePowerConsumption,
		ShipmentAt:       &ret.DeviceShipmentAt,
		Posters:          ret.DevicePosters,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = deviceinfo1.DeleteDeviceInfo(context.Background(), ret.DeviceInfoID)
		_, _ = vendorlocation1.DeleteLocation(context.Background(), ret.VendorLocationID)
		_, _ = vendorbrand1.DeleteBrand(context.Background(), info1.ID)
	}
}

func createGood(t *testing.T) {
	info, err := CreateGood(context.Background(), &npool.GoodReq{
		ID:                   &ret.ID,
		DeviceInfoID:         &ret.DeviceInfoID,
		DurationDays:         &ret.DurationDays,
		CoinTypeID:           &ret.CoinTypeID,
		VendorLocationID:     &ret.VendorLocationID,
		Price:                &ret.Price,
		BenefitType:          &ret.BenefitType,
		GoodType:             &ret.GoodType,
		Title:                &ret.Title,
		Unit:                 &ret.Unit,
		UnitAmount:           &ret.UnitAmount,
		SupportCoinTypeIDs:   ret.SupportCoinTypeIDs,
		DeliveryAt:           &ret.DeliveryAt,
		StartAt:              &ret.StartAt,
		TestOnly:             &ret.TestOnly,
		BenefitIntervalHours: &ret.BenefitIntervalHours,
		UnitLockDeposit:      &ret.UnitLockDeposit,
		Total:                &ret.GoodTotal,
		Posters:              ret.Posters,
		Labels:               ret.Labels,
	})
	if assert.Nil(t, err) {
		ret.GoodStockID = info.GoodStockID
		ret.DevicePostersStr = info.DevicePostersStr
		ret.LabelsStr = info.LabelsStr
		ret.PostersStr = info.PostersStr
		ret.SupportCoinTypeIDsStr = info.SupportCoinTypeIDsStr
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateGood(t *testing.T) {
	info, err := UpdateGood(context.Background(), &npool.GoodReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getGood(t *testing.T) {
	info, err := GetGood(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func getGoods(t *testing.T) {
	infos, total, err := GetGoods(context.Background(), &npool.Conds{
		ID:               &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		DeviceInfoID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.DeviceInfoID},
		CoinTypeID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		VendorLocationID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.VendorLocationID},
		IDs:              &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.ID}},
		BenefitType:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.BenefitType)},
		GoodType:         &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.GoodType)},
		RewardState:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.RewardState)},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func getGoodOnly(t *testing.T) {
	info, err := GetGoodOnly(context.Background(), &npool.Conds{
		ID:               &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		DeviceInfoID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.DeviceInfoID},
		CoinTypeID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.CoinTypeID},
		VendorLocationID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.VendorLocationID},
		IDs:              &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.ID}},
		BenefitType:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.BenefitType)},
		GoodType:         &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.GoodType)},
		RewardState:      &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.RewardState)},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func deleteGood(t *testing.T) {
	info, err := DeleteGood(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetGood(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

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

	t.Run("createGood", createGood)
	t.Run("updateGood", updateGood)
	t.Run("getGood", getGood)
	t.Run("getGoods", getGoods)
	t.Run("getGoodOnly", getGoodOnly)
	t.Run("deleteGood", deleteGood)
}
