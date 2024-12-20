package pledge

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	goodcoin1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/coin"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
	goodcoinrewardmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin/reward"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var ret = npool.Pledge{
	EntID:              uuid.NewString(),
	GoodID:             uuid.NewString(),
	ContractCodeURL:    uuid.NewString(),
	ContractCodeBranch: uuid.NewString(),
	ContractState:      types.ContractState_ContractWaitDeployment,
	ContractStateStr:   types.ContractState_ContractWaitDeployment.String(),

	UnitPrice:          decimal.NewFromInt(0).String(),
	QuantityUnitAmount: decimal.NewFromInt(0).String(),

	GoodType:             types.GoodType_Pledge,
	GoodTypeStr:          types.GoodType_Pledge.String(),
	BenefitType:          types.BenefitType_BenefitTypeContract,
	BenefitTypeStr:       types.BenefitType_BenefitTypeContract.String(),
	Name:                 uuid.NewString(),
	ServiceStartAt:       uint32(time.Now().Unix()),
	StartMode:            types.GoodStartMode_GoodStartModeInstantly,
	StartModeStr:         types.GoodStartMode_GoodStartModeInstantly.String(),
	BenefitIntervalHours: 20,
	UnitLockDeposit:      decimal.NewFromInt(0).String(),

	GoodCoins: []*goodcoinmwpb.GoodCoinInfo{
		{
			CoinTypeID: uuid.NewString(),
			Main:       true,
		},
		{
			CoinTypeID: uuid.NewString(),
			Main:       false,
		},
		{
			CoinTypeID: uuid.NewString(),
			Main:       false,
		},
	},

	RewardState:    types.BenefitState_BenefitWait,
	RewardStateStr: types.BenefitState_BenefitWait.String(),
	State:          types.GoodState_GoodStatePreWait,
	StateStr:       types.GoodState_GoodStatePreWait.String(),
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	for _, goodCoin := range ret.GoodCoins {
		ret.Rewards = append(ret.Rewards, &goodcoinrewardmwpb.RewardInfo{
			GoodID:                ret.GoodID,
			CoinTypeID:            goodCoin.CoinTypeID,
			RewardTID:             uuid.Nil.String(),
			LastRewardAmount:      decimal.NewFromInt(0).String(),
			NextRewardStartAmount: decimal.NewFromInt(0).String(),
			LastUnitRewardAmount:  decimal.NewFromInt(0).String(),
			TotalRewardAmount:     decimal.NewFromInt(0).String(),
			MainCoin:              goodCoin.Main,
		})
	}
	return func(*testing.T) {}
}

func createPledge(t *testing.T) {
	h1, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
		WithState(&ret.State, true),
		WithContractCodeURL(&ret.ContractCodeURL, true),
		WithContractCodeBranch(&ret.ContractCodeBranch, true),
		WithContractState(&ret.ContractState, true),
	)
	assert.Nil(t, err)

	err = h1.CreatePledge(context.Background())
	if assert.Nil(t, err) {
		for _, goodCoin := range ret.GoodCoins {
			goodCoin.GoodID = ret.GoodID
			h5, err := goodcoin1.NewHandler(
				context.Background(),
				goodcoin1.WithGoodID(&ret.GoodID, true),
				goodcoin1.WithCoinTypeID(&goodCoin.CoinTypeID, true),
				goodcoin1.WithMain(&goodCoin.Main, true),
				goodcoin1.WithIndex(&goodCoin.Index, true),
			)
			assert.Nil(t, err)

			err = h5.CreateGoodCoin(context.Background())
			assert.Nil(t, err)
		}
		handler, _ := NewHandler(
			context.Background(),
			WithGoodID(&ret.GoodID, true),
		)
		info, err := handler.GetPledge(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			ret.State = info.State
			ret.StateStr = info.StateStr
			assert.Equal(t, &ret, info)
		}
	}

	h2, err := NewHandler(
		context.Background(),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
		WithState(&ret.State, true),
		WithContractCodeURL(&ret.ContractCodeURL, true),
		WithContractCodeBranch(&ret.ContractCodeBranch, true),
		WithContractState(&ret.ContractState, true),
	)
	assert.Nil(t, err)

	err = h2.CreatePledge(context.Background())
	assert.NotNil(t, err)
}

func updatePledge(t *testing.T) {
	ret.Name = uuid.NewString()

	h1, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)
	_, err = h1.GetPledge(context.Background())
	assert.Nil(t, err)

	h2, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
		WithGoodType(&ret.GoodType, true),
		WithBenefitType(&ret.BenefitType, true),
		WithName(&ret.Name, true),
		WithServiceStartAt(&ret.ServiceStartAt, true),
		WithStartMode(&ret.StartMode, true),
		WithTestOnly(&ret.TestOnly, true),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithPurchasable(&ret.Purchasable, true),
		WithOnline(&ret.Online, true),
	)
	assert.Nil(t, err)

	err = h2.UpdatePledge(context.Background())
	assert.Nil(t, err)

	info, err := h2.GetPledge(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getPledge(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetPledge(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getPledges(t *testing.T) {
	conds := &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetPledges(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deletePledge(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithEntID(&ret.EntID, true),
		WithGoodID(&ret.GoodID, true),
	)
	assert.Nil(t, err)

	err = handler.DeletePledge(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetPledge(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestPledge(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createPledge", createPledge)
	t.Run("updatePledge", updatePledge)
	t.Run("getPledge", getPledge)
	t.Run("getPledges", getPledges)
	t.Run("deletePledge", deletePledge)
}
