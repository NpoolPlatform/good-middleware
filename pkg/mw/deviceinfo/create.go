package deviceinfo

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	deviceinfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/deviceinfo"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"

	"github.com/google/uuid"
)

func (h *Handler) CreateDeviceInfo(ctx context.Context) (*npool.DeviceInfo, error) {
	if h.Type == nil {
		return nil, fmt.Errorf("invalid type")
	}

	key := h.lockKey()
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &deviceinfocrud.Conds{
		Type: &cruder.Cond{Op: cruder.EQ, Val: *h.Type},
	}
	exist, err := h.ExistDeviceInfoConds(ctx)
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
		if _, err := deviceinfocrud.CreateSet(
			cli.DeviceInfo.Create(),
			&deviceinfocrud.Req{
				ID:              h.ID,
				Type:            h.Type,
				Manufacturer:    h.Manufacturer,
				PowerComsuption: h.PowerComsuption,
				ShipmentAt:      h.ShipmentAt,
				Posters:         h.Posters,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetDeviceInfo(ctx)
}
