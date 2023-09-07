package location

import (
	"context"
	"fmt"
	"time"

	locationcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/location"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
)

func (h *Handler) DeleteLocation(ctx context.Context) (*npool.Location, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetLocation(ctx)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := locationcrud.UpdateSet(
			cli.VendorLocation.UpdateOneID(*h.ID),
			&locationcrud.Req{
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
