package brand

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"

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

var ret = &npool.Brand{
	ID:   uuid.NewString(),
	Name: uuid.NewString(),
	Logo: uuid.NewString(),
}

var req = &npool.BrandReq{
	ID:   &ret.ID,
	Name: &ret.Name,
	Logo: &ret.Logo,
}

func createBrand(t *testing.T) {
	info, err := CreateBrand(context.Background(), req)
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func updateBrand(t *testing.T) {
	info, err := UpdateBrand(context.Background(), req)
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, ret, info)
	}
}

func getBrand(t *testing.T) {
	info, err := GetBrand(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func getBrands(t *testing.T) {
	infos, total, err := GetBrands(context.Background(), &npool.Conds{
		ID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		Name: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
	}, int32(0), int32(2))
	if assert.Nil(t, err) {
		if assert.Equal(t, uint32(1), total) {
			assert.Equal(t, ret, infos[0])
		}
	}
}

func getBrandOnly(t *testing.T) {
	info, err := GetBrandOnly(context.Background(), &npool.Conds{
		ID:   &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		Name: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Name},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}
}

func deleteBrand(t *testing.T) {
	info, err := DeleteBrand(context.Background(), ret.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, ret, info)
	}

	info, err = GetBrand(context.Background(), ret.ID)
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestBrand(t *testing.T) {
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

	t.Run("createBrand", createBrand)
	t.Run("updateBrand", updateBrand)
	t.Run("getBrand", getBrand)
	t.Run("getBrands", getBrands)
	t.Run("getBrandOnly", getBrandOnly)
	t.Run("deleteBrand", deleteBrand)
}
