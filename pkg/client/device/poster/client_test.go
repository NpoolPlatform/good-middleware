package poster

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	devicetype1 "github.com/NpoolPlatform/good-middleware/pkg/client/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/client/device/manufacturer"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	devicetypemwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
	manufacturermwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/poster"

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

var ret = &npool.Poster{
	EntID:        uuid.NewString(),
	DeviceTypeID: uuid.NewString(),
	DeviceType:   uuid.NewString(),
	Manufacturer: uuid.NewString(),
	Poster:       uuid.NewString(),
	Index:        5,
}

func setup(t *testing.T) func(*testing.T) {
	manufacturerID := uuid.NewString()
	err := manufacturer1.CreateManufacturer(context.Background(), &manufacturermwpb.ManufacturerReq{
		EntID: &manufacturerID,
		Name:  &ret.Manufacturer,
		Logo:  func() *string { s := uuid.NewString(); return &s }(),
	})
	assert.Nil(t, err)

	err = devicetype1.CreateDeviceType(context.Background(), &devicetypemwpb.DeviceTypeReq{
		EntID:            &ret.DeviceTypeID,
		Type:             &ret.DeviceType,
		ManufacturerID:   &manufacturerID,
		PowerConsumption: func() *uint32 { u := uint32(120); return &u }(),
		ShipmentAt:       func() *uint32 { u := uint32(time.Now().Unix()); return &u }(),
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = devicetype1.DeleteDeviceType(context.Background(), nil, &ret.DeviceTypeID)
		_ = manufacturer1.DeleteManufacturer(context.Background(), nil, &manufacturerID)
	}
}

func createPoster(t *testing.T) {
	err := CreatePoster(context.Background(), &npool.PosterReq{
		EntID:        &ret.EntID,
		DeviceTypeID: &ret.DeviceTypeID,
		Poster:       &ret.Poster,
		Index:        &ret.Index,
	})
	if assert.Nil(t, err) {
		info, err := GetPoster(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, ret)
		}
	}
}

func updatePoster(t *testing.T) {
	err := UpdatePoster(context.Background(), &npool.PosterReq{
		ID:           &ret.ID,
		EntID:        &ret.EntID,
		DeviceTypeID: &ret.DeviceTypeID,
		Poster:       &ret.Poster,
		Index:        &ret.Index,
	})
	if assert.Nil(t, err) {
		info, err := GetPoster(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, ret)
		}
	}
}

func getPoster(t *testing.T) {
	info, err := GetPoster(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getPosters(t *testing.T) {
	conds := &npool.Conds{
		ID:            &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:         &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		DeviceTypeID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.DeviceTypeID},
		DeviceTypeIDs: &basetypes.StringSliceVal{Op: cruder.IN, Value: []string{ret.DeviceTypeID}},
	}
	infos, _, err := GetPosters(context.Background(), conds, 0, 2)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func deletePoster(t *testing.T) {
	err := DeletePoster(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetPoster(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestPoster(t *testing.T) {
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

	t.Run("createPoster", createPoster)
	t.Run("updatePoster", updatePoster)
	t.Run("getPoster", getPoster)
	t.Run("getPosters", getPosters)
	t.Run("deletePoster", deletePoster)
}
