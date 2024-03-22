package device

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
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
	ret = npool.DeviceType{
		EntID:            uuid.NewString(),
		Type:             uuid.NewString(),
		Manufacturer:     uuid.NewString(),
		PowerConsumption: 120,
		ShipmentAt:       uint32(time.Now().Unix()),
	}
)

func setup() func(*testing.T) {
	return func(*testing.T) {}
}

func createDeviceType(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
		WithType(&ret.Type, true),
		WithManufacturer(&ret.Manufacturer, true),
		WithPowerConsumption(&ret.PowerConsumption, true),
		WithShipmentAt(&ret.ShipmentAt, true),
	)
	assert.Nil(t, err)

	err = handler.CreateDeviceType(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetDeviceType(context.Background())
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, &ret)
		}
	}
}

func updateDeviceType(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithType(&ret.Type, false),
		WithManufacturer(&ret.Manufacturer, false),
		WithPowerConsumption(&ret.PowerConsumption, false),
		WithShipmentAt(&ret.ShipmentAt, false),
	)
	assert.Nil(t, err)

	err = handler.UpdateDeviceType(context.Background())
	if assert.Nil(t, err) {
		info, err := handler.GetDeviceType(context.Background())
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, &ret)
		}
	}
}

func getDeviceType(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithEntID(&ret.EntID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetDeviceType(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getDeviceTypes(t *testing.T) {
	conds := &npool.Conds{
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
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

	infos, _, err := handler.GetDeviceTypes(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteDeviceType(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	err = handler.DeleteDeviceType(context.Background())
	assert.Nil(t, err)

	info, err := handler.GetDeviceType(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDeviceType(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup()
	defer teardown(t)

	t.Run("createDeviceType", createDeviceType)
	t.Run("updateDeviceType", updateDeviceType)
	t.Run("getDeviceType", getDeviceType)
	t.Run("getDeviceTypes", getDeviceTypes)
	t.Run("deleteDeviceType", deleteDeviceType)
}
