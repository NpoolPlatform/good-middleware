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

//nolint:gocyclo
func (h *Handler) UpdateLocation(ctx context.Context) (*npool.Location, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetLocation(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid location")
	}

	if h.Country == nil {
		h.Country = &info.Country
	}
	if h.Province == nil {
		h.Province = &info.Province
	}
	if h.City == nil {
		h.City = &info.City
	}
	if h.Address == nil {
		h.Address = &info.Address
	}
	if h.BrandID == nil {
		id, err := uuid.Parse(info.BrandID)
		if err != nil {
			return nil, err
		}
		h.BrandID = &id
	}

	key := h.lockKey()
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &locationcrud.Conds{
		ID:       &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
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
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := locationcrud.UpdateSet(
			cli.VendorLocation.UpdateOneID(*h.ID),
			&locationcrud.Req{
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
