package good

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

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

var (
	inheritGoodID = uuid.NewString()

	goodInfo = npool.Good{
		ID:                    uuid.NewString(),
		DeviceInfoID:          uuid.NewString(),
		DeviceType:            uuid.NewString(),
		DeviceManufacturer:    uuid.NewString(),
		DevicePowerComsuption: 123,
		DeviceShipmentAt:      uint32(time.Now().Unix() - 1000),
		DevicePosters:         []string{uuid.NewString(), uuid.NewString()},
		DurationDays:          365,
		CoinTypeID:            uuid.NewString(),
		InheritFromGoodID:     &inheritGoodID,
		VendorLocationID:      uuid.NewString(),
		GoodType:              mgrpb.GoodType_GoodTypeClassicMining,
		BenefitType:           mgrpb.BenefitType_BenefitTypePlatform,
		Price:                 "123.0",
		Title:                 uuid.NewString(),
		Unit:                  "TiB",
		UnitAmount:            1,
		SupportCoinTypeIDs:    []string{uuid.NewString(), uuid.NewString()},
		TestOnly:              true,
		Posters:               []string{uuid.NewString(), uuid.NewString()},
		Labels:                []string{uuid.NewString(), uuid.NewString()},
		GoodTotal:             1000,
		DeliveryAt:            uint32(time.Now().Unix() + 1000),
		StartAt:               uint32(time.Now().Unix() + 1000),
	}
)

func creatGood(t *testing.T) {
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
	info, err := CreateGood(context.Background(), &goodReq)
	if assert.Nil(t, err) {
		goodInfo.CreatedAt = info.CreatedAt
		goodInfo.UpdatedAt = info.UpdatedAt
		goodInfo.GoodStockID = info.GoodStockID
		goodInfo.SupportCoinTypeIDsStr = info.SupportCoinTypeIDsStr
		goodInfo.PostersStr = info.PostersStr
		goodInfo.LabelsStr = info.LabelsStr
		assert.Equal(t, info, &goodInfo)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("creatGood", creatGood)
}
