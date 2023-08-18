package appstock

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	appstockcrud.Req
	Conds  *appstockcrud.Conds
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

func WithID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
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

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
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

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := good1.NewHandler(
			ctx,
			good1.WithID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistGood(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid good")
		}
		h.GoodID = handler.ID
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := appgood1.NewHandler(
			ctx,
			appgood1.WithID(id, true),
		)
		if err != nil {
			return err
		}
		info, err := handler.GetGood(ctx)
		if err != nil {
			return err
		}
		if info == nil {
			return fmt.Errorf("invalid appgood")
		}
		h.AppGoodID = handler.ID
		return nil
	}
}

func WithReserved(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid Reserved")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.Reserved = &amount
		return nil
	}
}

func WithSpotQuantity(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid SpotQuantity")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.SpotQuantity = &amount
		return nil
	}
}

func WithLocked(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid Locked")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.Locked = &amount
		return nil
	}
}

func WithInService(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid InService")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.InService = &amount
		return nil
	}
}

func WithWaitStart(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid WaitStart")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.WaitStart = &amount
		return nil
	}
}

func WithSold(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid Sold")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.Sold = &amount
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &appstockcrud.Conds{}
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
		if conds.AppGoodID != nil {
			id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.AppGoodID = &cruder.Cond{
				Op:  conds.GetAppGoodID().GetOp(),
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
		if conds.AppGoodIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetAppGoodIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.AppGoodIDs = &cruder.Cond{
				Op:  conds.GetAppGoodIDs().GetOp(),
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
