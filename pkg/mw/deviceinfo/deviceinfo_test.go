package deviceinfo

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"
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

var (
	ret = npool.DeviceInfo{
		ID:              uuid.NewString(),
		Type:            uuid.NewString(),
		Manufacturer:    uuid.NewString(),
		PowerComsuption: 120,
		ShipmentAt:      uint32(time.Now().Unix()),
		Posters:         []string{uuid.NewString()},
	}
)

func setup() func(*testing.T) {
	b, _ := json.Marshal(ret.Posters)
	ret.PostersStr = string(b)
	return func(*testing.T) {}
}

func createDeviceInfo(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithType(&ret.Type),
		WithManufacturer(&ret.Manufacturer),
		WithPowerComsuption(&ret.PowerComsuption),
		WithShipmentAt(&ret.ShipmentAt),
		WithPosters(ret.Posters),
	)
	assert.Nil(t, err)

	info, err := handler.CreateDeviceInfo(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateDeviceInfo(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
		WithType(&ret.Type),
		WithManufacturer(&ret.Manufacturer),
		WithPowerComsuption(&ret.PowerComsuption),
		WithShipmentAt(&ret.ShipmentAt),
		WithPosters(ret.Posters),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateDeviceInfo(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getDeviceInfo(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.GetDeviceInfo(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getDeviceInfos(t *testing.T) {
	conds := &npool.Conds{
		ID:           &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		Type:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.Type},
		Manufacturer: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Manufacturer},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetDeviceInfos(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteDeviceInfo(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteDeviceInfo(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetDeviceInfo(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDeviceInfo(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup()
	defer teardown(t)

	t.Run("createDeviceInfo", createDeviceInfo)
	t.Run("updateDeviceInfo", updateDeviceInfo)
	t.Run("getDeviceInfo", getDeviceInfo)
	t.Run("getDeviceInfos", getDeviceInfos)
	t.Run("deleteDeviceInfo", deleteDeviceInfo)
}
