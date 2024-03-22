package comment

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

	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/client/app/good"
	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	appgoodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
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
	EntID:        uuid.NewString(),
	AppID:        uuid.NewString(),
	GoodID:       good.EntID,
	GoodName:     uuid.NewString(),
	UnitPrice:    decimal.NewFromInt(123).String(),
	PackagePrice: decimal.NewFromInt(123).String(),
}

var comment = npool.Comment{
	EntID:     uuid.NewString(),
	AppID:     appgood.AppID,
	UserID:    uuid.NewString(),
	GoodID:    appgood.GoodID,
	AppGoodID: appgood.EntID,
	OrderID:   uuid.NewString(),
	Content:   uuid.NewString(),
}

var ret = npool.Comment{
	EntID:     uuid.NewString(),
	AppID:     comment.AppID,
	UserID:    uuid.NewString(),
	GoodID:    appgood.GoodID,
	AppGoodID: appgood.EntID,
	Content:   uuid.NewString(),
	OrderID:   uuid.Nil.String(),
	ReplyToID: comment.EntID,
	GoodName:  appgood.GoodName,
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

	device, err := deviceinfo1.CreateDeviceInfo(context.Background(), &deviceinfomwpb.DeviceInfoReq{
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
		GoodName:         &appgood.GoodName,
		UnitPrice:        &appgood.UnitPrice,
		PackagePrice:     &appgood.PackagePrice,
		MinOrderDuration: &appgood.MinOrderDuration,
		MaxOrderDuration: &appgood.MaxOrderDuration,
	})
	assert.Nil(t, err)

	_, err = CreateComment(context.Background(), &npool.CommentReq{
		EntID:     &comment.EntID,
		AppID:     &comment.AppID,
		UserID:    &comment.UserID,
		AppGoodID: &comment.AppGoodID,
		OrderID:   &comment.OrderID,
		Content:   &comment.Content,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = DeleteComment(context.Background(), comment.ID)
		_, _ = appgood1.DeleteGood(context.Background(), appgood.ID)
		_, _ = good1.DeleteGood(context.Background(), good.ID)
		_, _ = deviceinfo1.DeleteDeviceInfo(context.Background(), device.ID)
		_, _ = vendorlocation1.DeleteLocation(context.Background(), vendorLocation.ID)
		_, _ = vendorbrand1.DeleteBrand(context.Background(), info1.ID)
	}
}

func createComment(t *testing.T) {
	info, err := CreateComment(context.Background(), &npool.CommentReq{
		EntID:     &ret.EntID,
		AppID:     &ret.AppID,
		UserID:    &ret.UserID,
		AppGoodID: &ret.AppGoodID,
		Content:   &ret.Content,
		ReplyToID: &ret.ReplyToID,
	})
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, &ret, info)
	}
}

func updateComment(t *testing.T) {
	info, err := UpdateComment(context.Background(), &npool.CommentReq{
		ID: &ret.ID,
	})
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, &ret, info)
	}
}

//nolint
func getComments(t *testing.T) {
	infos, total, err := GetComments(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
		OrderID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		OrderIDs:   &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

//nolint
func getCommentOnly(t *testing.T) {
	info, err := GetCommentOnly(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
		OrderID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		OrderIDs:   &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

//nolint
func deleteComment(t *testing.T) {
	info, err := DeleteComment(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetCommentOnly(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:     &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		AppGoodID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppGoodID},
		AppGoodIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.AppGoodID}},
		OrderID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		OrderIDs:   &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
	})
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestComment(t *testing.T) {
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

	t.Run("createComment", createComment)
	t.Run("updateComment", updateComment)
	t.Run("getComments", getComments)
	t.Run("getCommentOnly", getCommentOnly)
	t.Run("deleteComment", deleteComment)
}
