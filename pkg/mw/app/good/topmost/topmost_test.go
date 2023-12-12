package topmost

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"

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

var ret = npool.TopMost{
	EntID:                  uuid.NewString(),
	AppID:                  uuid.NewString(),
	TopMostType:            types.GoodTopMostType_TopMostInnovationStarter,
	TopMostTypeStr:         types.GoodTopMostType_TopMostInnovationStarter.String(),
	Title:                  uuid.NewString(),
	Message:                uuid.NewString(),
	Posters:                []string{uuid.NewString(), uuid.NewString()},
	StartAt:                uint32(time.Now().Unix() + 1000),
	EndAt:                  uint32(time.Now().Unix() + 6000),
	ThresholdCredits:       decimal.NewFromInt(0).String(),
	RegisterElapsedSeconds: 3000,
	ThresholdPurchases:     3000,
	ThresholdPaymentAmount: "3000",
	KycMust:                true,
}

func createTopMost(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppID(&ret.AppID, true),
		WithTopMostType(&ret.TopMostType, true),
		WithTitle(&ret.Title, true),
		WithMessage(&ret.Message, true),
		WithPosters(ret.Posters, true),
		WithStartAt(&ret.StartAt, true),
		WithEndAt(&ret.EndAt, true),
		WithThresholdCredits(&ret.ThresholdCredits, true),
		WithRegisterElapsedSeconds(&ret.RegisterElapsedSeconds, true),
		WithThresholdPurchases(&ret.ThresholdPurchases, true),
		WithThresholdPaymentAmount(&ret.ThresholdPaymentAmount, true),
		WithKycMust(&ret.KycMust, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateTopMost(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, &ret, info)
		}
	}
}

func updateTopMost(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateTopMost(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getTopMost(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetTopMost(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			assert.Equal(t, &ret, info)
		}
	}
}

func getTopMosts(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:          &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
			EntID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
			AppID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			TopMostType: &basetypes.Uint32Val{Op: cruder.EQ, Value: uint32(ret.TopMostType)},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetTopMosts(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				infos[0].PostersStr = ret.PostersStr
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteTopMost(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteTopMost(context.Background())
		if assert.Nil(t, err) {
			info.PostersStr = ret.PostersStr
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetTopMost(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestTopMost(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	t.Run("createTopMost", createTopMost)
	t.Run("updateTopMost", updateTopMost)
	t.Run("getTopMost", getTopMost)
	t.Run("getTopMosts", getTopMosts)
	t.Run("deleteTopMost", deleteTopMost)
}
