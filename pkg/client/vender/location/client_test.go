package location

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	brand1 "github.com/NpoolPlatform/good-middleware/pkg/client/vender/brand"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

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

var ret = &npool.Location{
	EntID:     uuid.NewString(),
	Country:   uuid.NewString(),
	Province:  uuid.NewString(),
	City:      uuid.NewString(),
	Address:   uuid.NewString(),
	BrandID:   uuid.NewString(),
	BrandName: uuid.NewString(),
	BrandLogo: uuid.NewString(),
}

func setup(t *testing.T) func(*testing.T) {
	err := brand1.CreateBrand(context.Background(), &brandmwpb.BrandReq{
		EntID: &ret.BrandID,
		Name:  &ret.BrandName,
		Logo:  &ret.BrandLogo,
	})
	assert.Nil(t, err)

	return func(*testing.T) {
		_ = brand1.DeleteBrand(context.Background(), nil, &ret.BrandID)
	}
}

func createLocation(t *testing.T) {
	err := CreateLocation(context.Background(), &npool.LocationReq{
		EntID:    &ret.EntID,
		Country:  &ret.Country,
		Province: &ret.Province,
		City:     &ret.City,
		Address:  &ret.Address,
		BrandID:  &ret.BrandID,
	})
	if assert.Nil(t, err) {
		info, err := GetLocation(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.CreatedAt = info.CreatedAt
			ret.UpdatedAt = info.UpdatedAt
			ret.ID = info.ID
			assert.Equal(t, ret, info)
		}
	}
}

func updateLocation(t *testing.T) {
	err := UpdateLocation(context.Background(), &npool.LocationReq{
		ID:       &ret.ID,
		Country:  &ret.Country,
		Province: &ret.Province,
		City:     &ret.City,
		Address:  &ret.Address,
		BrandID:  &ret.BrandID,
	})
	if assert.Nil(t, err) {
		info, err := GetLocation(context.Background(), ret.EntID)
		if assert.Nil(t, err) {
			ret.UpdatedAt = info.UpdatedAt
			assert.Equal(t, ret, info)
		}
	}
}

func getLocation(t *testing.T) {
	info, err := GetLocation(context.Background(), ret.EntID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getLocations(t *testing.T) {
	infos, total, err := GetLocations(context.Background(), &npool.Conds{
		ID:       &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Country:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.Country},
		Province: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Province},
		BrandID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.BrandID},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getLocationOnly(t *testing.T) {
	info, err := GetLocationOnly(context.Background(), &npool.Conds{
		ID:       &basetypes.Uint32Val{Op: cruder.EQ, Value: ret.ID},
		EntID:    &basetypes.StringVal{Op: cruder.EQ, Value: ret.EntID},
		Country:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.Country},
		Province: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Province},
		BrandID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.BrandID},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteLocation(t *testing.T) {
	err := DeleteLocation(context.Background(), &ret.ID, &ret.EntID)
	assert.Nil(t, err)

	info, err := GetLocation(context.Background(), ret.EntID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestLocation(t *testing.T) {
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

	t.Run("createLocation", createLocation)
	t.Run("updateLocation", updateLocation)
	t.Run("getLocation", getLocation)
	t.Run("getLocations", getLocations)
	t.Run("getLocationOnly", getLocationOnly)
	t.Run("deleteLocation", deleteLocation)
}
