package brand

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	brandcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/brand"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (h *Handler) UpdateBrand(ctx context.Context) (*npool.Brand, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetBrand(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid brand")
	}

	if h.Name == nil {
		h.Name = &info.Name
	}

	key := h.lockKey()
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &brandcrud.Conds{
		ID:   &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
		Name: &cruder.Cond{Op: cruder.EQ, Val: *h.Name},
	}
	exist, err := h.ExistBrandConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("arleady exists")
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := brandcrud.UpdateSet(
			cli.VendorBrand.UpdateOneID(*h.ID),
			&brandcrud.Req{
				Name: h.Name,
				Logo: h.Logo,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetBrand(ctx)
}
