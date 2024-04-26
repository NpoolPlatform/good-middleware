package fee

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"

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

var ret = &npool.Fee{
	EntID:               uuid.NewString(),
	GoodID:              uuid.NewString(),
	GoodType:            types.GoodType_TechniqueServiceFee,
	Name:                uuid.NewString(),
	SettlementType:      types.GoodSettlementType_GoodSettledByProfitPercent,
	UnitValue:           decimal.NewFromInt(20).String(),
	DurationDisplayType: types.GoodDurationType_GoodDurationByDay,
}

//nolint:unparam
func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.SettlementTypeStr = ret.SettlementType.String()
	ret.DurationDisplayTypeStr = ret.DurationDisplayType.String()

	return func(*testing.T) {}
}

func createFee(t *testing.T) {
	err := CreateFee(context.Background(), &npool.FeeReq{
		EntID:               &ret.EntID,
		GoodID:              &ret.GoodID,
		GoodType:            &ret.GoodType,
		Name:                &ret.Name,
		SettlementType:      &ret.SettlementType,
		UnitValue:           &ret.UnitValue,
		DurationDisplayType: &ret.DurationDisplayType,
	})
	if assert.Nil(t, err) {
		info, err := GetFee(context.Background(), ret.GoodID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, ret)
		}
	}
}

func updateFee(t *testing.T) {
	err := UpdateFee(context.Background(), &npool.FeeReq{
		ID:                  &ret.ID,
		EntID:               &ret.EntID,
		GoodID:              &ret.GoodID,
		GoodType:            &ret.GoodType,
		Name:                &ret.Name,
		SettlementType:      &ret.SettlementType,
		UnitValue:           &ret.UnitValue,
		DurationDisplayType: &ret.DurationDisplayType,
	})
	if assert.Nil(t, err) {
		info, err := GetFee(context.Background(), ret.GoodID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, ret)
		}
	}
}

func getFee(t *testing.T) {
	info, err := GetFee(context.Background(), ret.GoodID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getFees(t *testing.T) {
	conds := &npool.Conds{
		ID:     &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		GoodID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
	}
	infos, _, err := GetFees(context.Background(), conds, 0, 2)
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func deleteFee(t *testing.T) {
	err := DeleteFee(context.Background(), &ret.ID, &ret.EntID, &ret.GoodID)
	assert.Nil(t, err)

	info, err := GetFee(context.Background(), ret.GoodID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestFee(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createFee", createFee)
	t.Run("updateFee", updateFee)
	t.Run("getFee", getFee)
	t.Run("getFees", getFees)
	t.Run("deleteFee", deleteFee)
}
