package coin

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	// "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	// basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var (
	ret = npool.GoodCoin{
		EntID:      uuid.NewString(),
		GoodID:     uuid.NewString(),
		GoodType:   types.GoodType_PowerRental,
		GoodName:   uuid.NewString(),
		CoinTypeID: uuid.NewString(),
	}
)

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()

	h1, err := goodbase1.NewHandler(
		context.Background(),
		goodbase1.WithEntID(&ret.GoodID, true),
		goodbase1.WithGoodType(&ret.GoodType, true),
		goodbase1.WithName(&ret.GoodName, true),
	)
	assert.Nil(t, err)

	err = h1.CreateGoodBase(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteGoodBase(context.Background())
	}
}

func createFee(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
	)
	assert.Nil(t, err)

	err = handler.CreateGoodCoin(context.Background())
	assert.Nil(t, err)
}

func TestFee(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createFee", createFee)
}
