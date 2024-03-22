package device

import (
	"context"
	"fmt"
	"time"

	devicecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
)

func (h *Handler) DeleteDeviceInfo(ctx context.Context) (*npool.DeviceInfo, error) {
	if h.ID == nil {
		return nil, fmt.Errorf("invalid id")
	}

	info, err := h.GetDeviceInfo(ctx)
	if err != nil {
		return nil, err
	}

	now := uint32(time.Now().Unix())
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := devicecrud.UpdateSet(
			cli.DeviceInfo.UpdateOneID(*h.ID),
			&devicecrud.Req{
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
