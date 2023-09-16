package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	appgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good"
	good1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	appstockcrud.Req
	AppSpotLocked     *decimal.Decimal
	AppStockLockState *types.AppStockLockState
	LockID            *uuid.UUID
	Rollback          *bool
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
				return fmt.Errorf("invalid reserved")
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
				return fmt.Errorf("invalid spotquantity")
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
				return fmt.Errorf("invalid locked")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid locked")
		}
		h.Locked = &amount
		return nil
	}
}

func WithInService(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid inservice")
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
				return fmt.Errorf("invalid waitstart")
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
				return fmt.Errorf("invalid sold")
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

func WithAppSpotLocked(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid appspotlocked")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid appspotlocked")
		}
		h.AppSpotLocked = &amount
		return nil
	}
}

func WithLockID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid lockid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.LockID = &_id
		return nil
	}
}

func WithRollback(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Rollback = b
		return nil
	}
}

func WithAppStockLockState(e *types.AppStockLockState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid appstocklockstate")
			}
		}
		switch *e {
		case types.AppStockLockState_AppStockLocked:
		case types.AppStockLockState_AppStockWaitStart:
		case types.AppStockLockState_AppStockInService:
		case types.AppStockLockState_AppStockExpired:
		case types.AppStockLockState_AppStockChargeBack:
		case types.AppStockLockState_AppStockRollback:
		default:
			return fmt.Errorf("invalid appstocklockstate")
		}
		h.AppStockLockState = e
		return nil
	}
}
