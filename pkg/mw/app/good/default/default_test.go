package appdefaultgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	goodcoin1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin"
	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"

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

var ret = npool.Default{
	EntID:       uuid.NewString(),
	AppID:       uuid.NewString(),
	GoodID:      uuid.NewString(),
	GoodName:    uuid.NewString(),
	AppGoodID:   uuid.NewString(),
	AppGoodName: uuid.NewString(),
	CoinTypeID:  uuid.NewString(),
}
var appGoodID2 = uuid.NewString()

func setup(t *testing.T) func(*testing.T) {
	goodType := types.GoodType_PowerRental
	goodCoinEntID := uuid.NewString()

	h1, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&ret.GoodID, true),
		goodbase1.WithGoodType(&goodType, true),
		goodbase1.WithName(&ret.GoodName, true),
	)
	assert.Nil(t, err)

	err = h1.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	h2, err := appgoodbase1.NewHandler(
		context.Background(),
		appgoodbase1.WithEntID(&ret.AppGoodID, true),
		appgoodbase1.WithAppID(&ret.AppID, true),
		appgoodbase1.WithGoodID(&ret.GoodID, true),
		appgoodbase1.WithName(&ret.AppGoodName, true),
	)
	assert.Nil(t, err)

	err = h2.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	h3, err := goodcoin1.NewHandler(
		context.Background(),
		goodcoin1.WithEntID(&goodCoinEntID, true),
		goodcoin1.WithGoodID(&ret.GoodID, true),
		goodcoin1.WithCoinTypeID(&ret.CoinTypeID, true),
	)
	assert.Nil(t, err)

	err = h3.CreateGoodCoin(context.Background())
	assert.Nil(t, err)

	h4, err := appgoodbase1.NewHandler(
		context.Background(),
		appgoodbase1.WithEntID(&appGoodID2, true),
		appgoodbase1.WithAppID(&ret.AppID, true),
		appgoodbase1.WithGoodID(&ret.GoodID, true),
		appgoodbase1.WithName(&ret.AppGoodName, true),
	)
	assert.Nil(t, err)

	err = h4.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h4.DeleteGoodBase(context.Background())
		_ = h3.DeleteGoodCoin(context.Background())
		_ = h2.DeleteGoodBase(context.Background())
		_ = h1.DeleteGoodBase(context.Background())
	}
}

func createDefault(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&ret.AppGoodID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
	)
	assert.Nil(t, err)

	err = h1.CreateDefault(context.Background())
	if assert.Nil(t, err) {
		info, err := h1.GetDefault(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}

	h2, err := NewHandler(
		context.Background(),
		WithAppGoodID(&ret.AppGoodID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
	)
	assert.Nil(t, err)

	err = h2.CreateDefault(context.Background())
	assert.NotNil(t, err)
}

func updateDefault(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithAppGoodID(&appGoodID2, true),
	)
	assert.Nil(t, err)

	err = handler.UpdateDefault(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetDefault(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			ret.AppGoodID = appGoodID2
			assert.Equal(t, info, &ret)
		}
	}
}

func TestDefault(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createDefault", createDefault)
	t.Run("updateDefault", updateDefault)
}
