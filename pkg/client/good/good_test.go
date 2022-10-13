package good

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"

	testinit "github.com/NpoolPlatform/good-middleware/pkg/testinit"
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

var (
	inheritGoodID = uuid.NewString()

	supportCoinTypeIDs    = []string{uuid.NewString(), uuid.NewString()}
	supportCoinTypeIDsStr = fmt.Sprintf(`["%v", "%v"]`, supportCoinTypeIDs[0], supportCoinTypeIDs[1])

	posters    = []string{uuid.NewString(), uuid.NewString()}
	postersStr = fmt.Sprintf(`["%v", "%v"]`, posters[0], posters[1])

	labels    = []string{uuid.NewString(), uuid.NewString()}
	labelsStr = fmt.Sprintf(`["%v", "%v"]`, labels[0], labels[1])

	goodInfo = npool.Good{
		ID:                    uuid.NewString(),
		DeviceInfoID:          uuid.NewString(),
		DurationDays:          365,
		CoinTypeID:            uuid.NewString(),
		InheritFromGoodID:     &inheritGoodID,
		VendorLocationID:      uuid.NewString(),
		GoodType:              mgrpb.GoodType_GoodTypeClassicMining,
		GoodTypeStr:           mgrpb.GoodType_GoodTypeClassicMining.String(),
		BenefitType:           mgrpb.BenefitType_BenefitTypePlatform,
		BenefitTypeStr:        mgrpb.BenefitType_BenefitTypePlatform.String(),
		Price:                 "123.000000000000000000",
		Title:                 uuid.NewString(),
		Unit:                  "TiB",
		UnitAmount:            1,
		SupportCoinTypeIDs:    supportCoinTypeIDs,
		SupportCoinTypeIDsStr: supportCoinTypeIDsStr,
		TestOnly:              true,
		Posters:               posters,
		PostersStr:            postersStr,
		Labels:                labels,
		LabelsStr:             labelsStr,
		GoodTotal:             1000,
		DeliveryAt:            uint32(time.Now().Unix() + 1000),
		StartAt:               uint32(time.Now().Unix() + 1000),
	}
)

var (
	goodReq = npool.GoodReq{
		ID:                 &goodInfo.ID,
		DeviceInfoID:       &goodInfo.DeviceInfoID,
		DurationDays:       &goodInfo.DurationDays,
		CoinTypeID:         &goodInfo.CoinTypeID,
		InheritFromGoodID:  goodInfo.InheritFromGoodID,
		VendorLocationID:   &goodInfo.VendorLocationID,
		Price:              &goodInfo.Price,
		BenefitType:        &goodInfo.BenefitType,
		GoodType:           &goodInfo.GoodType,
		Title:              &goodInfo.Title,
		Unit:               &goodInfo.Unit,
		UnitAmount:         &goodInfo.UnitAmount,
		SupportCoinTypeIDs: goodInfo.SupportCoinTypeIDs,
		DeliveryAt:         &goodInfo.DeliveryAt,
		StartAt:            &goodInfo.StartAt,
		TestOnly:           &goodInfo.TestOnly,
		Total:              &goodInfo.GoodTotal,
		Posters:            goodInfo.Posters,
		Labels:             goodInfo.Labels,
	}
)

var info *npool.Good

func createGood(t *testing.T) {
	var err error
	info, err = CreateGood(context.Background(), &goodReq)
	if assert.Nil(t, err) {
		goodInfo.CreatedAt = info.CreatedAt
		goodInfo.UpdatedAt = info.UpdatedAt
		goodInfo.GoodStockID = info.GoodStockID
		assert.Equal(t, info, &goodInfo)
	}
}

func updateGood(t *testing.T) {
	var err error
	info, err = UpdateGood(context.Background(), &goodReq)
	if assert.Nil(t, err) {
		goodInfo.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &goodInfo)
	}
}

func getGood(t *testing.T) {
	var err error
	info, err = GetGood(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &goodInfo)
	}
}

func getGoods(t *testing.T) {
	infos, total, err := GetGoods(context.Background(),
		&mgrpb.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &goodInfo)
	}
}

func getGoodOnly(t *testing.T) {
	var err error
	infos, total, err := GetManyGoods(context.Background(), []string{info.ID}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &goodInfo)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createGood", createGood)
	t.Run("getGood", getGood)
	t.Run("getGoods", getGoods)
	t.Run("getGoodOnly", getGoodOnly)
	t.Run("updateGood", updateGood)
}
