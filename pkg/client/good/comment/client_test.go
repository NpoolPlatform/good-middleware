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

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/client/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/client/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	deviceinfomwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
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
	ID:                     uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturer:     uuid.NewString(),
	DevicePowerComsuption:  120,
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

var comment = npool.Comment{
	ID:      uuid.NewString(),
	AppID:   uuid.NewString(),
	UserID:  uuid.NewString(),
	GoodID:  good.ID,
	OrderID: uuid.NewString(),
	Content: uuid.NewString(),
}

var ret = npool.Comment{
	ID:        uuid.NewString(),
	AppID:     comment.AppID,
	UserID:    uuid.NewString(),
	GoodID:    good.ID,
	Content:   uuid.NewString(),
	ReplyToID: comment.ID,
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
		ID:              &good.DeviceInfoID,
		Type:            &good.DeviceType,
		Manufacturer:    &good.DeviceManufacturer,
		PowerComsuption: &good.DevicePowerComsuption,
		ShipmentAt:      &good.DeviceShipmentAt,
		Posters:         good.DevicePosters,
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

	_, err = CreateComment(context.Background(), &npool.CommentReq{
		ID:      &comment.ID,
		AppID:   &comment.AppID,
		UserID:  &comment.UserID,
		GoodID:  &comment.GoodID,
		OrderID: &comment.OrderID,
		Content: &comment.Content,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = DeleteComment(context.Background(), comment.ID)
		_, _ = good1.DeleteGood(context.Background(), good.ID)
		_, _ = deviceinfo1.DeleteDeviceInfo(context.Background(), good.DeviceInfoID)
		_, _ = vendorlocation1.DeleteLocation(context.Background(), good.VendorLocationID)
		_, _ = vendorbrand1.DeleteBrand(context.Background(), info1.ID)
	}
}

func createComment(t *testing.T) {
	info, err := CreateComment(context.Background(), &npool.CommentReq{
		ID:        &ret.ID,
		AppID:     &ret.AppID,
		UserID:    &ret.UserID,
		GoodID:    &ret.GoodID,
		Content:   &ret.Content,
		ReplyToID: &ret.ReplyToID,
	})
	fmt.Printf("---------------- %v\n", err)
	if assert.Nil(t, err) {
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

func getComments(t *testing.T) {
	infos, total, err := GetComments(context.Background(), &npool.Conds{
		ID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		GoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
		OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

func getCommentOnly(t *testing.T) {
	info, err := GetCommentOnly(context.Background(), &npool.Conds{
		ID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		GoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
		OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

func deleteComment(t *testing.T) {
	info, err := DeleteComment(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}

	info, err = GetCommentOnly(context.Background(), &npool.Conds{
		ID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		UserID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
		GoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
		GoodIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
		OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
		OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
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
	// t.Run("updateComment", updateComment)
	// t.Run("getComments", getComments)
	// t.Run("getCommentOnly", getCommentOnly)
	// t.Run("deleteComment", deleteComment)
}