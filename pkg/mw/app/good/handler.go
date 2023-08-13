package appgood

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	appgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	appgoodcrud.Req
	Conds  *appgoodcrud.Conds
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

func WithAppID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppID = &_id
		return nil
	}
}

func WithGoodID(id *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.GoodID = &_id
		return nil
	}
}

func WithOnline(b *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Online = b
		return nil
	}
}

func WithVisible(b *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Visible = b
		return nil
	}
}

func WithGoodName(s *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			return nil
		}
		const leastNameLen = 3
		if len(*s) < leastNameLen {
			return fmt.Errorf("invalid goodname")
		}
		h.GoodName = s
		return nil
	}
}

func WithPrice(s *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.Price = &amount
		return nil
	}
}

func WithDisplayIndex(n *int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.DisplayIndex = n
		return nil
	}
}

func WithPurchaseLimit(n *int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.PurchaseLimit = n
		return nil
	}
}

func WithSaleStartAt(n *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SaleStartAt = n
		return nil
	}
}

func WithSaleEndAt(n *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SaleEndAt = n
		return nil
	}
}

func WithServiceStartAt(n *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ServiceStartAt = n
		return nil
	}
}

func WithDescriptions(ss []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Descriptions = ss
		return nil
	}
}

func WithGoodBanner(s *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBanner = s
		return nil
	}
}

func WithEnablePurchase(b *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnablePurchase = b
		return nil
	}
}

func WithEnableProductPage(b *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnableProductPage = b
		return nil
	}
}

func WithCancelMode(e *types.CancelMode) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			return nil
		}
		switch *e {
		case types.CancelMode_CancellableBeforeStart:
		case types.CancelMode_CancellableBeforeBenefit:
		case types.CancelMode_Uncancellable:
		default:
			return fmt.Errorf("invalid cancelmode")
		}
		h.CancelMode = e
		return nil
	}
}

func WithUserPurchaseLimit(s *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.UserPurchaseLimit = &amount
		return nil
	}
}

func WithDisplayColors(ss []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.DisplayColors = ss
		return nil
	}
}

func WithCancellableBeforeStart(n *uint32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.CancellableBeforeStart = n
		return nil
	}
}

func WithProductPage(s *string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ProductPage = s
		return nil
	}
}

func WithEnableSetCommission(b *bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnableSetCommission = b
		return nil
	}
}

func WithPosters(ss []string) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Posters = ss
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appgoodcrud.Conds{}
		if conds == nil {
			return nil
		}
		if conds.ID != nil {
			id, err := uuid.Parse(conds.GetID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.ID = &cruder.Cond{
				Op:  conds.GetID().GetOp(),
				Val: id,
			}
		}
		if conds.AppID != nil {
			id, err := uuid.Parse(conds.GetAppID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppID = &cruder.Cond{
				Op:  conds.GetAppID().GetOp(),
				Val: id,
			}
		}
		if conds.GoodID != nil {
			id, err := uuid.Parse(conds.GetGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.GoodID = &cruder.Cond{
				Op:  conds.GetGoodID().GetOp(),
				Val: id,
			}
		}
		if conds.GoodIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetGoodIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.GoodIDs = &cruder.Cond{
				Op:  conds.GetGoodIDs().GetOp(),
				Val: ids,
			}
		}
		if conds.AppIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetAppIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.AppIDs = &cruder.Cond{
				Op:  conds.GetAppIDs().GetOp(),
				Val: ids,
			}
		}
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}