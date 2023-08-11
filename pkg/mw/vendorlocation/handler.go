package vendorlocation

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	vendorlocationcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vendorlocation"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vendorlocation"

	"github.com/google/uuid"
)

type Handler struct {
	vendorlocationcrud.Req
	Conds  *vendorlocationcrud.Conds
	Offset int32
	Limit  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.ID = &_id
		return nil
	}
}

func WithType(s *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
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

func WithManufacturer(s *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			return nil
		}
		const leastManufacturerLen = 3
		if len(*s) <= leastManufacturerLen {
			return fmt.Errorf("invalid manufacturer")
		}
		h.Manufacturer = s
		return nil
	}
}

func WithPowerComsuption(s *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.PowerComsuption = s
		return nil
	}
}

func WithShipmentAt(n *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ShipmentAt = n
		return nil
	}
}

func WithPosters(ss []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, s := range ss {
			if s == "" {
				return fmt.Errorf("invalid poster")
			}
		}
		h.Posters = ss
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &vendorlocationcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: id}
		}
		if conds.Type != nil {
			h.Conds.Type = &cruder.Cond{
				Op:  conds.GetType().GetOp(),
				Val: conds.GetType().GetValue(),
			}
		}
		if conds.Manufacturer != nil {
			h.Conds.Manufacturer = &cruder.Cond{
				Op:  conds.GetManufacturer().GetOp(),
				Val: conds.GetManufacturer().GetValue(),
			}
		}
		return nil
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
