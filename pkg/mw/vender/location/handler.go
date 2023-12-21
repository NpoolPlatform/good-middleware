package location

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	locationcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/location"
	brand1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/brand"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

	"github.com/google/uuid"
)

type Handler struct {
	locationcrud.Req
	Conds  *locationcrud.Conds
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

func WithCountry(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid country")
			}
			return nil
		}
		const leastCountryLen = 3
		if len(*s) <= leastCountryLen {
			return fmt.Errorf("invalid country")
		}
		h.Country = s
		return nil
	}
}

func WithProvince(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid province")
			}
			return nil
		}
		const leastProvinceLen = 3
		if len(*s) <= leastProvinceLen {
			return fmt.Errorf("invalid province")
		}
		h.Province = s
		return nil
	}
}

func WithCity(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid city")
			}
			return nil
		}
		const leastCityLen = 3
		if len(*s) <= leastCityLen {
			return fmt.Errorf("invalid city")
		}
		h.City = s
		return nil
	}
}

func WithAddress(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid address")
			}
			return nil
		}
		const leastAddressLen = 3
		if len(*s) <= leastAddressLen {
			return fmt.Errorf("invalid address")
		}
		h.Address = s
		return nil
	}
}

func WithBrandID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid brandid")
			}
			return nil
		}
		handler, err := brand1.NewHandler(
			ctx,
			brand1.WithEntID(id, true),
		)
		if err != nil {
			return fmt.Errorf("invalid brandid")
		}
		exist, err := handler.ExistBrand(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid brand")
		}
		h.BrandID = handler.EntID
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &locationcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.EntID != nil {
			id, err := uuid.Parse(conds.GetEntID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
		}
		if conds.ID != nil {
			h.Conds.ID = &cruder.Cond{Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue()}
		}
		if conds.Country != nil {
			h.Conds.Country = &cruder.Cond{
				Op:  conds.GetCountry().GetOp(),
				Val: conds.GetCountry().GetValue(),
			}
		}
		if conds.Province != nil {
			h.Conds.Province = &cruder.Cond{
				Op:  conds.GetProvince().GetOp(),
				Val: conds.GetProvince().GetValue(),
			}
		}
		if conds.BrandID != nil {
			id, err := uuid.Parse(conds.GetBrandID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.BrandID = &cruder.Cond{
				Op:  conds.GetBrandID().GetOp(),
				Val: id,
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
