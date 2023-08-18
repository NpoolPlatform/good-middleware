package stock

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID          *uuid.UUID
	GoodID      *uuid.UUID
	Locked      *decimal.Decimal
	WaitStart   *decimal.Decimal
	InService   *decimal.Decimal
	AppReserved *decimal.Decimal
	Conds       *stockcrud.Conds
	Offset      int32
	Limit       int32
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

func WithLocked(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid locked")
			}
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.Locked = &amount
		return nil
	}
}

func WithWaitStart(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid waitstart")
			}
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.WaitStart = &amount
		return nil
	}
}

func WithInService(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid inservice")
			}
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.InService = &amount
		return nil
	}
}

func WithAppReserved(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid appreserved")
			}
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.AppReserved = &amount
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &stockcrud.Conds{}
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
