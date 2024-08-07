package topmost

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"

	"github.com/google/uuid"
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

var ret = &npool.TopMost{
	EntID:          uuid.NewString(),
	AppID:          uuid.NewString(),
	TopMostType:    types.GoodTopMostType_TopMostNoviceExclusive,
	TopMostTypeStr: types.GoodTopMostType_TopMostNoviceExclusive.String(),
	Title:          uuid.NewString(),
	Message:        uuid.NewString(),
	StartAt:        uint32(time.Now().Unix() + 1000),
	EndAt:          uint32(time.Now().Unix() + 6000),
}

func createTopMost(t *testing.T) {
	err := CreateTopMost(context.Background(), &npool.TopMostReq{
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		TopMostType: &ret.TopMostType,
		Title:       &ret.Title,
		Message:     &ret.Message,
		StartAt:     &ret.StartAt,
		EndAt:       &ret.EndAt,
	})
	if assert.Nil(t, err) {
		info, err := GetTopMost(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func updateTopMost(t *testing.T) {
	ret.EndAt += 10
	err := UpdateTopMost(context.Background(), &npool.TopMostReq{
		ID:          &ret.ID,
		EntID:       &ret.EntID,
		AppID:       &ret.AppID,
		TopMostType: &ret.TopMostType,
		Title:       &ret.Title,
		Message:     &ret.Message,
		StartAt:     &ret.StartAt,
		EndAt:       &ret.EndAt,
	})
	if assert.Nil(t, err) {
		info, err := GetTopMost(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getTopMost(t *testing.T) {
	info, err := GetTopMost(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getTopMosts(t *testing.T) {
	infos, total, err := GetTopMosts(context.Background(), &npool.Conds{
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		TopMostType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TopMostType)},
		Title:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.Title},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getTopMostOnly(t *testing.T) {
	info, err := GetTopMostOnly(context.Background(), &npool.Conds{
		ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		TopMostType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TopMostType)},
		Title:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.Title},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteTopMost(t *testing.T) {
	err := DeleteTopMost(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetTopMost(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestTopMost(t *testing.T) {
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

	t.Run("createTopMost", createTopMost)
	t.Run("updateTopMost", updateTopMost)
	t.Run("getTopMost", getTopMost)
	t.Run("getTopMosts", getTopMosts)
	t.Run("getTopMostOnly", getTopMostOnly)
	t.Run("deleteTopMost", deleteTopMost)
}
