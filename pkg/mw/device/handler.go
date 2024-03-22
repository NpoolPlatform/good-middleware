package device

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	devicecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/device"
	manufacturer1 "github.com/NpoolPlatform/good-middleware/pkg/mw/device/manufacturer"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"

	"github.com/google/uuid"
)

type Handler struct {
	devicecrud.Req
	DeviceConds *devicecrud.Conds
	Offset      int32
	Limit       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		DeviceConds: &devicecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithType(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid type")
			}
			return nil
		}
		const leastTypeLen = 3
		if len(*s) <= leastTypeLen {
			return fmt.Errorf("invalid type")
		}
		h.Type = s
		return nil
	}
}

func WithManufacturerID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid manufacturer")
			}
			return nil
		}
		handler, err := manufacturer1.NewHandler(
			ctx,
			manufacturer1.WithEntID(s, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistManufacturer(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid manufacturer")
		}
		h.ManufacturerID = handler.EntID
		return nil
	}
}

func WithPowerConsumption(s *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.PowerConsumption = s
		return nil
	}
}

func WithShipmentAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ShipmentAt = n
		return nil
	}
}

func (h *Handler) withDeviceConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.DeviceConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.DeviceConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.Type != nil {
		h.DeviceConds.Type = &cruder.Cond{
			Op:  conds.GetType().GetOp(),
			Val: conds.GetType().GetValue(),
		}
	}
	if conds.ManufacturerID != nil {
		id, err := uuid.Parse(conds.GetManufacturerID().GetValue())
		if err != nil {
			return err
		}
		h.DeviceConds.ManufacturerID = &cruder.Cond{
			Op:  conds.GetManufacturerID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		return h.withDeviceConds(conds)
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
