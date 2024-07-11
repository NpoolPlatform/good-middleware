package manufacturer

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"bou.ke/monkey"
	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"

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

var ret = &npool.Manufacturer{
	EntID: uuid.NewString(),
	Name:  uuid.NewString(),
	Logo:  uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	return func(*testing.T) {}
}

func createManufacturer(t *testing.T) {
	err := CreateManufacturer(context.Background(), &npool.ManufacturerReq{
		EntID: &ret.EntID,
		Name:  &ret.Name,
		Logo:  &ret.Logo,
	})
	if assert.Nil(t, err) {
		info, err := GetManufacturer(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, info, ret)
		}
	}
}

func updateManufacturer(t *testing.T) {
	ret.Name = uuid.NewString()
	err := UpdateManufacturer(context.Background(), &npool.ManufacturerReq{
		ID:    &ret.ID,
		EntID: &ret.EntID,
		Name:  &ret.Name,
		Logo:  &ret.Logo,
	})
	if assert.Nil(t, err) {
		info, err := GetManufacturer(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, info, ret)
		}
	}
}

func getManufacturer(t *testing.T) {
	info, err := GetManufacturer(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, ret)
	}
}

func getManufacturers(t *testing.T) {
	conds := &npool.Conds{
		ID:    &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID: &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
	}
	infos, _, err := GetManufacturers(context.Background(), conds, 0, 2)
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], ret)
	}
}

func deleteManufacturer(t *testing.T) {
	err := DeleteManufacturer(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetManufacturer(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestManufacturer(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})
	monkey.Patch(grpc2.GetGRPCConnV1, func(service string, recvMsgBytes int, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createManufacturer", createManufacturer)
	t.Run("updateManufacturer", updateManufacturer)
	t.Run("getManufacturer", getManufacturer)
	t.Run("getManufacturers", getManufacturers)
	t.Run("deleteManufacturer", deleteManufacturer)
}
