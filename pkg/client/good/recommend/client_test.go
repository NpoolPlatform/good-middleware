//nolint:dupl
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

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/recommend"
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

var ret = npool.Recommend{
	ID:             uuid.NewString(),
	AppID:          uuid.NewString(),
	RecommenderID:  uuid.NewString(),
	GoodID:         good.ID,
	RecommendIndex: decimal.RequireFromString("4.9").String(),
	Message:        uuid.NewString(),
	GoodName:       good.Title,
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
		ID:               &good.DeviceInfoID,
		Type:             &good.DeviceType,
		Manufacturer:     &good.DeviceManufacturer,
		PowerConsumption: &good.DevicePowerConsumption,
		ShipmentAt:       &good.DeviceShipmentAt,
		Posters:          good.DevicePosters,
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

func createRecommend(t *testing.T) {
	info, err := CreateRecommend(context.Background(), &npool.RecommendReq{
		ID:             &ret.ID,
		AppID:          &ret.AppID,
		RecommenderID:  &ret.RecommenderID,
		GoodID:         &ret.GoodID,
		Message:        &ret.Message,
		RecommendIndex: &ret.RecommendIndex,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func updateRecommend(t *testing.T) {
	ret.Message = uuid.NewString()
	info, err := UpdateRecommend(context.Background(), &npool.RecommendReq{
		ID:      &ret.ID,
		Message: &ret.Message,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

func getRecommends(t *testing.T) {
	infos, total, err := GetRecommends(context.Background(), &npool.Conds{
		ID:            &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RecommenderID},
		GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func getRecommendOnly(t *testing.T) {
	info, err := GetRecommendOnly(context.Background(), &npool.Conds{
		ID:            &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RecommenderID},
		GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func deleteRecommend(t *testing.T) {
	info, err := DeleteRecommend(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetRecommendOnly(context.Background(), &npool.Conds{
		ID:            &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		RecommenderID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.RecommenderID},
		GoodID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:       &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
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