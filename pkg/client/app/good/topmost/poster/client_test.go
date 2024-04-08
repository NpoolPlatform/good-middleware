package poster

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

	topmost1 "github.com/NpoolPlatform/good-middleware/pkg/client/app/good/topmost"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	topmostmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/poster"

	"github.com/google/uuid"
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

var ret = npool.Poster{
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	TopMostID:        uuid.NewString(),
	TopMostType:      types.GoodTopMostType_TopMostPromotion,
	TopMostTitle:     uuid.NewString(),
	TopMostMessage:   uuid.NewString(),
	TopMostTargetUrl: uuid.NewString(),
	Poster:           uuid.NewString(),
	Index:            5,
}

func setup(t *testing.T) func(*testing.T) {
	ret.TopMostTypeStr = ret.TopMostType.String()

	err := topmost1.CreateTopMost(context.Background(), &topmostmwpb.TopMostReq{
		EntID:       &ret.TopMostID,
		AppID:       &ret.AppID,
		TopMostType: &ret.TopMostType,
		Title:       &ret.TopMostTitle,
		Message:     &ret.TopMostMessage,
		TargetUrl:   &ret.TopMostTargetUrl,
		StartAt:     func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
		EndAt:       func() *uint32 { u := uint32(time.Now().Unix() + 1000); return &u }(),
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = topmost1.DeleteTopMost(context.Background(), nil, &ret.TopMostID)
	}
}

func createPoster(t *testing.T) {
	err := CreatePoster(context.Background(), &npool.PosterReq{
		EntID:     &ret.EntID,
		TopMostID: &ret.TopMostID,
		Poster:    &ret.Poster,
		Index:     &ret.Index,
	})
	if assert.Nil(t, err) {
		info, err := GetPoster(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updatePoster(t *testing.T) {
	err := UpdatePoster(context.Background(), &npool.PosterReq{
		ID:     &ret.ID,
		Poster: &ret.Poster,
	})
	if assert.Nil(t, err) {
		info, err := GetPoster(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

//nolint
func getPosters(t *testing.T) {
	infos, total, err := GetPosters(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		TopMostID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.TopMostID},
		TopMostIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.TopMostID}},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, &ret, infos[0])
		}
	}
}

//nolint
func getPosterOnly(t *testing.T) {
	info, err := GetPosterOnly(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		TopMostID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.TopMostID},
		TopMostIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.TopMostID}},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, &ret, info)
	}
}

//nolint
func deletePoster(t *testing.T) {
	err := DeletePoster(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetPosterOnly(context.Background(), &npool.Conds{
		ID:         &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		AppID:      &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
		TopMostID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.TopMostID},
		TopMostIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.TopMostID}},
	})
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestPoster(t *testing.T) {
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

	t.Run("createPoster", createPoster)
	t.Run("updatePoster", updatePoster)
	t.Run("getPosters", getPosters)
	t.Run("getPosterOnly", getPosterOnly)
	t.Run("deletePoster", deletePoster)
}
