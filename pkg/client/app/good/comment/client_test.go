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

	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
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

var comment = npool.Comment{
	EntID:     uuid.NewString(),
	AppID:     uuid.NewString(),
	UserID:    uuid.NewString(),
	GoodID:    uuid.NewString(),
	AppGoodID: uuid.NewString(),
	OrderID:   uuid.NewString(),
	Content:   uuid.NewString(),
}

var ret = npool.Comment{
	EntID:             uuid.NewString(),
	AppID:             comment.AppID,
	UserID:            uuid.NewString(),
	GoodID:            comment.GoodID,
	AppGoodID:         comment.AppGoodID,
	GoodName:          uuid.NewString(),
	Content:           uuid.NewString(),
	OrderID:           uuid.Nil.String(),
	ReplyToID:         comment.EntID,
	OrderFirstComment: false,
}

func setup(t *testing.T) func(*testing.T) {

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
		_, _ = devicetype1.DeleteDeviceInfo(context.Background(), device.ID)
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
