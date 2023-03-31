package appdefaultgood

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	testinit "github.com/NpoolPlatform/good-middleware/pkg/testinit"
	npool "github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	mgrconverter "github.com/NpoolPlatform/good-manager/pkg/converter/appdefaultgood"
	mgrcrud "github.com/NpoolPlatform/good-manager/pkg/crud/appdefaultgood"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appDate = npool.AppDefaultGood{
	ID:         uuid.NewString(),
	AppID:      uuid.NewString(),
	GoodID:     uuid.NewString(),
	CoinTypeID: uuid.NewString(),
}

var (
	appInfo = npool.AppDefaultGoodReq{
		ID:         &appDate.ID,
		AppID:      &appDate.AppID,
		GoodID:     &appDate.GoodID,
		CoinTypeID: &appDate.CoinTypeID,
	}
)

var info *npool.AppDefaultGood

func getAppDefaultGood(t *testing.T) {
	entInfo, err := mgrcrud.Create(context.Background(), &appInfo)
	assert.Nil(t, err)
	info = mgrconverter.Ent2Grpc(entInfo)
	info, err = GetAppDefaultGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getAppDefaultGoods(t *testing.T) {
	infos, total, err := GetAppDefaultGoods(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getAppDefaultGoodOnly(t *testing.T) {
	var err error
	info, err = GetAppDefaultGoodOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		appDate.CreatedAt = info.CreatedAt
		appDate.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &appDate)
	}
}

func deleteAppDefaultGood(t *testing.T) {
	_, err := DeleteAppDefaultGood(context.Background(), info.ID)
	assert.Nil(t, err)
	_, err = GetAppDefaultGood(context.Background(), info.ID)
	assert.NotNil(t, err)
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	t.Run("getAppDefaultGood", getAppDefaultGood)
	t.Run("getAppDefaultGoods", getAppDefaultGoods)
	t.Run("getAppDefaultGoodOnly", getAppDefaultGoodOnly)
	t.Run("deleteAppDefaultGood", deleteAppDefaultGood)
}
