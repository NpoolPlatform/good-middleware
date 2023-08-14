package good

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"

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

var ret = npool.Good{
	ID:                     uuid.NewString(),
	DeviceInfoID:           uuid.NewString(),
	DeviceType:             uuid.NewString(),
	DeviceManufacturer:     uuid.NewString(),
	DevicePowerComsuption:  120,
	DeviceShipmentAt:       uint32(time.Now().Unix() - 1000),
	DevicePosters:          []string{uuid.NewString(), uuid.NewString()},
	DurationDays:           14,
	CoinTypeID:             uuid.NewString(),
	VendorLocationID:       uuid.NewString(),
	VendorLocationCountry:  uuid.NewString(),
	VendorLocationProvince: uuid.NewString(),
	VendorLocationCity:     uuid.NewString(),
	VendorLocationAddress:  uuid.NewString(),
	VendorBrandName:        uuid.NewString(),
	VendorBrandLogo:        uuid.NewString(),
	GoodType:               types.GoodType_PowerRenting,
	BenefitType:            types.BenefitType_BenefitTypePlatform,
	Price:                  decimal.NewFromInt(123).String(),
	Title:                  uuid.NewString(),
	Unit:                   "TiB",
	UnitAmount:             1,
	SupportCoinTypeIDs:     []string{uuid.NewString(), uuid.NewString()},
	TestOnly:               true,
	Posters:                []string{uuid.NewString(), uuid.NewString()},
	Labels:                 []string{uuid.NewString(), uuid.NewString()},
	GoodTotal:              decimal.NewFromInt(1000).String(),
	GoodLocked:             decimal.NewFromInt(0).String(),
	GoodInService:          decimal.NewFromInt(0).String(),
	GoodWaitStart:          decimal.NewFromInt(0).String(),
	GoodSold:               decimal.NewFromInt(0).String(),
	DeliveryAt:             uint32(time.Now().Unix() + 1000),
	StartAt:                uint32(time.Now().Unix() + 1000),
	BenefitIntervalHours:   24,
	GoodAppLocked:          decimal.NewFromInt(0).String(),
	UnitLockDeposit:        decimal.NewFromInt(1).String(),
}

func setup(t *testing.T) func(*testing.T) {
	b, _ := json.Marshal(ret.Labels)
	ret.LabelsStr = strings.ReplaceAll(string(b), " ", "")
	b, _ = json.Marshal(ret.Posters)
	ret.PostersStr = strings.ReplaceAll(string(b), " ", "")
	b, _ = json.Marshal(ret.SupportCoinTypeIDs)
	ret.SupportCoinTypeIDsStr = strings.ReplaceAll(string(b), " ", "")
	b, _ = json.Marshal(ret.DevicePosters)
	ret.DevicePostersStr = strings.ReplaceAll(string(b), " ", "")

	ret.BenefitTypeStr = ret.BenefitType.String()
	ret.GoodTypeStr = ret.GoodType.String()
	ret.RewardStateStr = types.BenefitState_BenefitWait.String()
	ret.RewardTID = uuid.Nil.String()
	ret.NextRewardStartAmount = decimal.NewFromInt(0).String()
	ret.LastRewardAmount = decimal.NewFromInt(0).String()

	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithName(&ret.VendorBrandName),
		vendorbrand1.WithLogo(&ret.VendorBrandLogo),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithID(&ret.VendorLocationID),
		vendorlocation1.WithCountry(&ret.VendorLocationCountry),
		vendorlocation1.WithProvince(&ret.VendorLocationProvince),
		vendorlocation1.WithCity(&ret.VendorLocationCity),
		vendorlocation1.WithAddress(&ret.VendorLocationAddress),
		vendorlocation1.WithBrandID(&info1.ID),
	)
	assert.Nil(t, err)

	_, err = h2.CreateLocation(context.Background())
	assert.Nil(t, err)

	h3, err := deviceinfo1.NewHandler(
		context.Background(),
		deviceinfo1.WithID(&ret.DeviceInfoID),
		deviceinfo1.WithType(&ret.DeviceType),
		deviceinfo1.WithManufacturer(&ret.DeviceManufacturer),
		deviceinfo1.WithPowerComsuption(&ret.DevicePowerComsuption),
		deviceinfo1.WithShipmentAt(&ret.DeviceShipmentAt),
		deviceinfo1.WithPosters(ret.DevicePosters),
	)
	assert.Nil(t, err)

	_, err = h3.CreateDeviceInfo(context.Background())

	return func(*testing.T) {
		_, _ = h3.DeleteDeviceInfo(context.Background())
		_, _ = h2.DeleteLocation(context.Background())
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func createGood(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithDeviceInfoID(&ret.DeviceInfoID, true),
		WithDurationDays(&ret.DurationDays, true),
		WithCoinTypeID(&ret.CoinTypeID, true),
		WithVendorLocationID(&ret.VendorLocationID, true),
		WithPrice(&ret.Price, true),
		WithBenefitType(&ret.BenefitType, true),
		WithGoodType(&ret.GoodType, true),
		WithTitle(&ret.Title, true),
		WithUnit(&ret.Unit, true),
		WithUnitAmount(&ret.UnitAmount, true),
		WithSupportCoinTypeIDs(ret.SupportCoinTypeIDs, false),
		WithDeliveryAt(&ret.DeliveryAt, true),
		WithStartAt(&ret.StartAt, true),
		WithTestOnly(&ret.TestOnly, false),
		WithBenefitIntervalHours(&ret.BenefitIntervalHours, true),
		WithUnitLockDeposit(&ret.UnitLockDeposit, false),
		WithTotal(&ret.GoodTotal, true),
		WithPosters(ret.Posters, false),
		WithLabels(ret.Labels, false),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateGood(context.Background())
		if assert.Nil(t, err) {
			info.LabelsStr = strings.ReplaceAll(info.LabelsStr, " ", "")
			info.PostersStr = strings.ReplaceAll(info.PostersStr, " ", "")
			info.SupportCoinTypeIDsStr = strings.ReplaceAll(info.SupportCoinTypeIDsStr, " ", "")
			info.DevicePostersStr = strings.ReplaceAll(info.DevicePostersStr, " ", "")
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.GoodStockID = info.GoodStockID
			assert.Equal(t, &ret, info)
		}
	}
}

func TestGood(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createGood", createGood)
}
