package deviceinfo

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"

	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

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

var ret = &npool.DeviceInfo{
	EntID:            uuid.NewString(),
	Type:             uuid.NewString(),
	Manufacturer:     uuid.NewString(),
	PowerConsumption: 123,
	ShipmentAt:       uint32(time.Now().Unix()),
	Posters:          []string{uuid.NewString(), uuid.NewString()},
}

var req = &npool.DeviceInfoReq{
	EntID:            &ret.EntID,
	Type:             &ret.Type,
	Manufacturer:     &ret.Manufacturer,
	PowerConsumption: &ret.PowerConsumption,
	ShipmentAt:       &ret.ShipmentAt,
	Posters:          ret.Posters,
}

func createDeviceInfo(t *testing.T) {
	info, err := CreateDeviceInfo(context.Background(), req)
	if assert.Nil(t, err) {
		ret.PostersStr = info.PostersStr
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		ret.ID = info.ID
		assert.Equal(t, ret, info)
	}
}

func updateDeviceInfo(t *testing.T) {
	req.ID = &ret.ID
	info, err := UpdateDeviceInfo(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getDeviceInfo(t *testing.T) {
	info, err := GetDeviceInfo(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getDeviceInfos(t *testing.T) {
	infos, total, err := GetDeviceInfos(context.Background(), &npool.Conds{
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Type:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.Type},
		Manufacturer: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Manufacturer},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getDeviceInfoOnly(t *testing.T) {
	info, err := GetDeviceInfoOnly(context.Background(), &npool.Conds{
		ID:           &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:        &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Type:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.Type},
		Manufacturer: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Manufacturer},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteDeviceInfo(t *testing.T) {
	info, err := DeleteDeviceInfo(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}

	info, err = GetDeviceInfo(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestDeviceInfo(t *testing.T) {
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

	t.Run("createDeviceInfo", createDeviceInfo)
	t.Run("updateDeviceInfo", updateDeviceInfo)
	t.Run("getDeviceInfo", getDeviceInfo)
	t.Run("getDeviceInfos", getDeviceInfos)
	t.Run("getDeviceInfoOnly", getDeviceInfoOnly)
	t.Run("deleteDeviceInfo", deleteDeviceInfo)
}
