package brand

import (
	"context"
	"fmt"
	"time"

	brandcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/brand"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

func (h *Handler) DeleteBrand(ctx context.Context) (*npool.Brand, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetBrand(ctx)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := brandcrud.UpdateSet(
			cli.VendorBrand.UpdateOneID(*h.ID),
			&brandcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
