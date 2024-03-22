package device

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	devicecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"

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

	h.Conds = &devicecrud.Conds{
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
	if h.EntID == nil {
		h.EntID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := devicecrud.CreateSet(
			cli.DeviceInfo.Create(),
			&devicecrud.Req{
				EntID:            h.EntID,
				Type:             h.Type,
				Manufacturer:     h.Manufacturer,
				PowerConsumption: h.PowerConsumption,
				ShipmentAt:       h.ShipmentAt,
				Posters:          h.Posters,
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
