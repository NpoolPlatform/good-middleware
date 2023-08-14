package good

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

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

var ret = npool.Good{
	ID:                   uuid.NewString(),
	DeviceInfoID:         uuid.NewString(),
	CoinTypeID:           uuid.NewString(),
	VendorLocationID:     uuid.NewString(),
	GoodType:             types.GoodType_PowerRenting,
	BenefitType:          types.BenefitType_BenefitTypePlatform,
	Price:                decimal.NewFromInt(123).String(),
	Title:                uuid.NewString(),
	Unit:                 "TiB",
	UnitAmount:           1,
	SupportCoinTypeIDs:   []string{uuid.NewString(), uuid.NewString()},
	TestOnly:             true,
	Posters:              []string{uuid.NewString(), uuid.NewString()},
	Labels:               []string{uuid.NewString(), uuid.NewString()},
	GoodTotal:            decimal.NewFromInt(1000).String(),
	DeliveryAt:           uint32(time.Now().Unix() + 1000),
	StartAt:              uint32(time.Now().Unix() + 1000),
	BenefitIntervalHours: 24,
	GoodAppLocked:        decimal.NewFromInt(0).String(),
}

func setup(t *testing.T) func(*testing.T) {

}
