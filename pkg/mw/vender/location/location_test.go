package location

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	brand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	brandmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
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
	brand = brandmwpb.Brand{
		ID:   uuid.NewString(),
		Name: uuid.NewString(),
		Logo: uuid.NewString(),
	}

	ret = npool.Location{
		ID:       uuid.NewString(),
		Country:  uuid.NewString(),
		Province: uuid.NewString(),
		City:     uuid.NewString(),
		Address:  uuid.NewString(),
		BrandID:  brand.ID,
	}
)

func setup(t *testing.T) func(*testing.T) {
	h1, err := brand1.NewHandler(
		context.Background(),
		brand1.WithID(&brand.ID, true),
		brand1.WithName(&brand.Name, true),
		brand1.WithLogo(&brand.Logo, true),
	)
	assert.Nil(t, err)

	_, err = h1.CreateBrand(context.Background())
	assert.Nil(t, err)

	return func(*testing.T) {
		_, _ = h1.DeleteBrand(context.Background())
	}
}

func createLocation(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCountry(&ret.Country, true),
		WithProvince(&ret.Province, true),
		WithCity(&ret.City, true),
		WithAddress(&ret.Address, true),
		WithBrandID(&ret.BrandID, true),
	)
	assert.Nil(t, err)

	info, err := handler.CreateLocation(context.Background())
	if assert.Nil(t, err) {
		ret.CreatedAt = info.CreatedAt
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func updateLocation(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
		WithCountry(&ret.Country, true),
		WithProvince(&ret.Province, true),
		WithCity(&ret.City, true),
		WithAddress(&ret.Address, true),
		WithBrandID(&ret.BrandID, true),
	)
	assert.Nil(t, err)

	info, err := handler.UpdateLocation(context.Background())
	if assert.Nil(t, err) {
		ret.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info, &ret)
	}
}

func getLocation(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.GetLocation(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}
}

func getLocations(t *testing.T) {
	conds := &npool.Conds{
		ID:       &basetypes.StringVal{Op: cruder.EQ, Value: ret.ID},
		Country:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.Country},
		Province: &basetypes.StringVal{Op: cruder.EQ, Value: ret.Province},
		BrandID:  &basetypes.StringVal{Op: cruder.EQ, Value: ret.BrandID},
	}

	handler, err := NewHandler(
		context.Background(),
		WithConds(conds),
		WithOffset(0),
		WithLimit(0),
	)
	assert.Nil(t, err)

	infos, _, err := handler.GetLocations(context.Background())
	if !assert.Nil(t, err) {
		assert.Equal(t, len(infos), 1)
		assert.Equal(t, infos[0], &ret)
	}
}

func deleteLocation(t *testing.T) {
	handler, err := NewHandler(
		context.Background(),
		WithID(&ret.ID, true),
	)
	assert.Nil(t, err)

	info, err := handler.DeleteLocation(context.Background())
	if assert.Nil(t, err) {
		assert.Equal(t, info, &ret)
	}

	info, err = handler.GetLocation(context.Background())
	assert.Nil(t, err)
	assert.Nil(t, info)
}

func TestLocation(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	teardown := setup(t)
	defer teardown(t)

	t.Run("createLocation", createLocation)
	t.Run("updateLocation", updateLocation)
	t.Run("getLocation", getLocation)
	t.Run("getLocations", getLocations)
	t.Run("deleteLocation", deleteLocation)
}
