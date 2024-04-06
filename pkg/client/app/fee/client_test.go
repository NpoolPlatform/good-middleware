package appfee

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

	fee1 "github.com/NpoolPlatform/good-middleware/pkg/client/fee"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"

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
	EntID:            uuid.NewString(),
	AppID:            uuid.NewString(),
	GoodID:           uuid.NewString(),
	AppGoodID:        uuid.NewString(),
	GoodType:         types.GoodType_TechniqueServiceFee,
	Name:             uuid.NewString(),
	SettlementType:   types.GoodSettlementType_GoodSettledByProfitPercent,
	UnitValue:        decimal.NewFromInt(20).String(),
	DurationType:     types.GoodDurationType_GoodDurationByDay,
	MinOrderDuration: 20,
}

func setup(t *testing.T) func(*testing.T) {
	ret.GoodTypeStr = ret.GoodType.String()
	ret.SettlementTypeStr = ret.SettlementType.String()
	ret.DurationTypeStr = ret.DurationType.String()

	feeEntID := uuid.NewString()
	h1, err := fee1.NewHandler(
		context.Background(),
		fee1.WithEntID(&feeEntID, true),
		fee1.WithGoodID(&ret.GoodID, true),
		fee1.WithGoodType(&ret.GoodType, true),
		fee1.WithName(&ret.Name, true),
		fee1.WithSettlementType(&ret.SettlementType, true),
		fee1.WithUnitValue(&ret.UnitValue, true),
		fee1.WithDurationType(&ret.DurationType, true),
	)
	assert.Nil(t, err)

	err = h1.CreateFee(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = h1.DeleteFee(context.Background())
	}
}
