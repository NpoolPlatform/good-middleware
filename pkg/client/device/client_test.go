package devicetype

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/good-middleware/pkg/testinit"

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

var ret = &npool.DeviceType{
	EntID:            uuid.NewString(),
	Type:             uuid.NewString(),
	ManufacturerID:   uuid.NewString(),
	ManufacturerName: uuid.NewString(),
	ManufacturerLogo: uuid.NewString(),
	PowerConsumption: 123,
	ShipmentAt:       uint32(time.Now().Unix()),
}

func setup(t *testing.T) func(*testing.T) {
	err := manufacturer1.CreateManufacturer(context.Background(), &manufacturermwpb.ManufacturerReq{
		EntID: &ret.ManufacturerID,
		Name:  &ret.ManufacturerName,
		Logo:  &ret.ManufacturerLogo,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = manufacturer1.DeleteManufacturer(context.Background(), nil, &ret.ManufacturerID)
	}
}

func createDeviceType(t *testing.T) {
	err := CreateDeviceType(context.Background(), &npool.DeviceTypeReq{
		EntID:            &ret.EntID,
		Type:             &ret.Type,
		ManufacturerID:   &ret.ManufacturerID,
		PowerConsumption: &ret.PowerConsumption,
		ShipmentAt:       &ret.ShipmentAt,
	})
	if assert.Nil(t, err) {
		info, err := GetDeviceType(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func updateDeviceType(t *testing.T) {
	err := UpdateDeviceType(context.Background(), &npool.DeviceTypeReq{
		ID:               &ret.ID,
		EntID:            &ret.EntID,
		Type:             &ret.Type,
		ManufacturerID:   &ret.ManufacturerID,
		PowerConsumption: &ret.PowerConsumption,
		ShipmentAt:       &ret.ShipmentAt,
	})
	if assert.Nil(t, err) {
		info, err := GetDeviceType(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getDeviceType(t *testing.T) {
	info, err := GetDeviceType(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getDeviceTypes(t *testing.T) {
	infos, total, err := GetDeviceTypes(context.Background(), &npool.Conds{
		ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Type:           &basetypes.StringVal{Op: cruder.EQ, Value: ret.Type},
		ManufacturerID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ManufacturerID},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getDeviceTypeOnly(t *testing.T) {
	info, err := GetDeviceTypeOnly(context.Background(), &npool.Conds{
		ID:             &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:          &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Type:           &basetypes.StringVal{Op: cruder.EQ, Value: ret.Type},
		ManufacturerID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.ManufacturerID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteDeviceType(t *testing.T) {
	err := DeleteDeviceType(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetDeviceType(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDeviceType(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	teardown := setup(t)
	defer teardown(t)

	t.Run("createDeviceType", createDeviceType)
	t.Run("updateDeviceType", updateDeviceType)
	t.Run("getDeviceType", getDeviceType)
	t.Run("getDeviceTypes", getDeviceTypes)
	t.Run("getDeviceTypeOnly", getDeviceTypeOnly)
	t.Run("deleteDeviceType", deleteDeviceType)
}
