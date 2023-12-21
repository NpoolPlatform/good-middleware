package location

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	locationcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/location"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

	"github.com/google/uuid"
)

func (h *Handler) CreateLocation(ctx context.Context) (*npool.Location, error) {
	if h.Country == nil {
		return nil, fmt.Errorf("invalid country")
	}
	if h.Province == nil {
		return nil, fmt.Errorf("invalid province")
	}
	if h.City == nil {
		return nil, fmt.Errorf("invalid city")
	}
	if h.Address == nil {
		return nil, fmt.Errorf("invalid address")
	}
	if h.BrandID == nil {
		return nil, fmt.Errorf("invalid brandid")
	}

	key := h.lockKey()
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &locationcrud.Conds{
		Country:  &cruder.Cond{Op: cruder.EQ, Val: *h.Country},
		Province: &cruder.Cond{Op: cruder.EQ, Val: *h.Province},
		City:     &cruder.Cond{Op: cruder.EQ, Val: *h.City},
		Address:  &cruder.Cond{Op: cruder.EQ, Val: *h.Address},
		BrandID:  &cruder.Cond{Op: cruder.EQ, Val: *h.BrandID},
	}
	exist, err := h.ExistLocationConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("arleady exists")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := locationcrud.CreateSet(
			cli.VendorLocation.Create(),
			&locationcrud.Req{
				EntID:    h.EntID,
				Country:  h.Country,
				Province: h.Province,
				City:     h.City,
				Address:  h.Address,
				BrandID:  h.BrandID,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLocation(ctx)
}
