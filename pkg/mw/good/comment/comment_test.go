package comment

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	vendorbrand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	goodmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"

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

var good = goodmwpb.Good{
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

var ret = npool.Comment{
	ID:        uuid.NewString(),
	AppID:     uuid.NewString(),
	UserID:    uuid.NewString(),
	GoodID:    good.ID,
	GoodName:  good.Title,
	OrderID:   uuid.NewString(),
	Content:   uuid.NewString(),
	ReplyToID: uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	h1, err := vendorbrand1.NewHandler(
		context.Background(),
		vendorbrand1.WithName(&good.VendorBrandName),
		vendorbrand1.WithLogo(&good.VendorBrandLogo),
	)
	assert.Nil(t, err)

	info1, err := h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	h2, err := vendorlocation1.NewHandler(
		context.Background(),
		vendorlocation1.WithID(&good.VendorLocationID),
		vendorlocation1.WithCountry(&good.VendorLocationCountry),
		vendorlocation1.WithProvince(&good.VendorLocationProvince),
		vendorlocation1.WithCity(&good.VendorLocationCity),
		vendorlocation1.WithAddress(&good.VendorLocationAddress),
		vendorlocation1.WithBrandID(&info1.ID),
	)
	assert.Nil(t, err)

	_, err = h2.CreateLocation(context.Background())
	assert.Nil(t, err)

	h3, err := deviceinfo1.NewHandler(
		context.Background(),
		deviceinfo1.WithID(&good.DeviceInfoID),
		deviceinfo1.WithType(&good.DeviceType),
		deviceinfo1.WithManufacturer(&good.DeviceManufacturer),
		deviceinfo1.WithPowerComsuption(&good.DevicePowerComsuption),
		deviceinfo1.WithShipmentAt(&good.DeviceShipmentAt),
		deviceinfo1.WithPosters(good.DevicePosters),
	)
	assert.Nil(t, err)

	_, err = h3.CreateDeviceInfo(context.Background())
	assert.Nil(t, err)

	h4, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
		good1.WithDeviceInfoID(&good.DeviceInfoID, true),
		good1.WithDurationDays(&good.DurationDays, true),
		good1.WithCoinTypeID(&good.CoinTypeID, true),
		good1.WithVendorLocationID(&good.VendorLocationID, true),
		good1.WithPrice(&good.Price, true),
		good1.WithBenefitType(&good.BenefitType, true),
		good1.WithGoodType(&good.GoodType, true),
		good1.WithTitle(&good.Title, true),
		good1.WithUnit(&good.Unit, true),
		good1.WithUnitAmount(&good.UnitAmount, true),
		good1.WithSupportCoinTypeIDs(good.SupportCoinTypeIDs, false),
		good1.WithDeliveryAt(&good.DeliveryAt, true),
		good1.WithStartAt(&good.StartAt, true),
		good1.WithTestOnly(&good.TestOnly, false),
		good1.WithBenefitIntervalHours(&good.BenefitIntervalHours, true),
		good1.WithUnitLockDeposit(&good.UnitLockDeposit, false),
		good1.WithTotal(&good.GoodTotal, true),
		good1.WithPosters(good.Posters, false),
		good1.WithLabels(good.Labels, false),
	)
	assert.Nil(t, err)

	_, err = h4.CreateGood(context.Background())
	assert.Nil(t, err)

	h5, err := NewHandler(
		context.Background(),
		WithID(&ret.ReplyToID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithOrderID(&ret.OrderID, true),
		WithContent(&ret.Content, true),
	)
	assert.Nil(t, err)

	_, err = h5.CreateComment(context.Background())

	return func(*testing.T) {
		_, _ = h5.DeleteComment(context.Background())
		_, _ = h4.DeleteGood(context.Background())
		_, _ = h3.DeleteDeviceInfo(context.Background())
		_, _ = h2.DeleteLocation(context.Background())
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func createComment(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithAppID(&ret.AppID, true),
		WithUserID(&ret.UserID, true),
		WithGoodID(&ret.GoodID, true),
		WithOrderID(&ret.OrderID, true),
		WithContent(&ret.Content, true),
		WithReplyToID(&ret.ReplyToID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.CreateComment(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}

	h1, err := good1.NewHandler(
		context.Background(),
		good1.WithID(&good.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := h1.GetGood(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, uint32(2), info.CommentCount)
		}
	}
}

func updateComment(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithContent(&ret.Content, true),
		WithReplyToID(&ret.ReplyToID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.UpdateComment(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, &ret, info)
		}
	}
}

func getComment(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.GetComment(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}
	}
}

func getComments(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithConds(&npool.Conds{
			ID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
			AppID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.AppID},
			UserID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.UserID},
			GoodID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.GoodID},
			GoodIDs:  &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.GoodID}},
			OrderID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.OrderID},
			OrderIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.OrderID}},
		}),
		WithOffset(0),
		WithLimit(0),
	)
	if assert.Nil(t, err) {
		infos, total, err := handler.GetComments(context.Background())
		if assert.Nil(t, err) {
			if assert.Equal(t, uint32(1), total) {
				assert.Equal(t, &ret, infos[0])
			}
		}
	}
}

func deleteComment(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	if assert.Nil(t, err) {
		info, err := handler.DeleteComment(context.Background())
		if assert.Nil(t, err) {
			assert.Equal(t, &ret, info)
		}

		info, err = handler.GetComment(context.Background())
		assert.Nil(t, err)
		assert.Nil(t, info)
	}
}

func TestComment(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createComment", createComment)
	t.Run("updateComment", updateComment)
	t.Run("getComment", getComment)
	t.Run("getComments", getComments)
	t.Run("deleteComment", deleteComment)
}
