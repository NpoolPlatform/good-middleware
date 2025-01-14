package delegatedstaking

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
	goodcoinrewardmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin/reward"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
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

var ret = &npool.DelegatedStaking{
	EntID:                uuid.NewString(),
	GoodID:               uuid.NewString(),
	GoodType:             types.GoodType_DelegatedStaking,
	BenefitType:          types.BenefitType_BenefitTypeContract,
	Name:                 uuid.NewString(),
	ServiceStartAt:       uint32(time.Now().Unix()),
	StartMode:            types.GoodStartMode_GoodStartModeInstantly,
	TestOnly:             true,
	BenefitIntervalHours: 24,
	Purchasable:          true,
	Online:               true,
	RewardState:          types.BenefitState_BenefitWait,
	ContractCodeURL:      uuid.NewString(),
	ContractCodeBranch:   uuid.NewString(),
	ContractState:        types.ContractState_ContractWaitDeployment,
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.BenefitTypeStr = ret.BenefitType.String()
	ret.RewardStateStr = ret.RewardState.String()
	ret.StartModeStr = ret.StartMode.String()
	ret.ContractStateStr = ret.ContractState.String()
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

func createDelegatedStaking(t *testing.T) {
	err := CreateDelegatedStaking(context.Background(), &npool.DelegatedStakingReq{
		EntID:                &ret.EntID,
		GoodID:               &ret.GoodID,
		GoodType:             &ret.GoodType,
		Name:                 &ret.Name,
		ServiceStartAt:       &ret.ServiceStartAt,
		StartMode:            &ret.StartMode,
		TestOnly:             &ret.TestOnly,
		BenefitIntervalHours: &ret.BenefitIntervalHours,
		Purchasable:          &ret.Purchasable,
		Online:               &ret.Online,
		ContractCodeURL:      &ret.ContractCodeURL,
		ContractCodeBranch:   &ret.ContractCodeBranch,
		ContractState:        &ret.ContractState,
	})
	if assert.Nil(t, err) {
		info, err := GetDelegatedStaking(context.Background(), ret.GoodID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			ret.State = info.State
			ret.StateStr = info.StateStr
			assert.Equal(t, ret, info)
		}
	}
}

func updateDelegatedStaking(t *testing.T) {
	ret.Name = uuid.NewString()
	err := UpdateDelegatedStaking(context.Background(), &npool.DelegatedStakingReq{
		ID:                   &ret.ID,
		EntID:                &ret.EntID,
		GoodType:             &ret.GoodType,
		Name:                 &ret.Name,
		ServiceStartAt:       &ret.ServiceStartAt,
		StartMode:            &ret.StartMode,
		TestOnly:             &ret.TestOnly,
		BenefitIntervalHours: &ret.BenefitIntervalHours,
		Purchasable:          &ret.Purchasable,
		Online:               &ret.Online,
	})
	if assert.Nil(t, err) {
		info, err := GetDelegatedStaking(context.Background(), ret.GoodID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getDelegatedStaking(t *testing.T) {
	info, err := GetDelegatedStaking(context.Background(), ret.GoodID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getDelegatedStakings(t *testing.T) {
	infos, total, err := GetDelegatedStakings(context.Background(), &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getDelegatedStakingOnly(t *testing.T) {
	info, err := GetDelegatedStakingOnly(context.Background(), &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteDelegatedStaking(t *testing.T) {
	err := DeleteDelegatedStaking(context.Background(), &ret.ID, &ret.EntID, &ret.GoodID)
	assert.Nil(t, err)

	info, err := GetDelegatedStaking(context.Background(), ret.GoodID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDelegatedStaking(t *testing.T) {
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

	t.Run("createDelegatedStaking", createDelegatedStaking)
	t.Run("updateDelegatedStaking", updateDelegatedStaking)
	t.Run("getDelegatedStaking", getDelegatedStaking)
	t.Run("getDelegatedStakings", getDelegatedStakings)
	t.Run("getDelegatedStakingOnly", getDelegatedStakingOnly)
	t.Run("deleteDelegatedStaking", deleteDelegatedStaking)
}
