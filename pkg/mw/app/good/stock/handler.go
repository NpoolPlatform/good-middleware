//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	appmininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/mining"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
	appmininggoodstockpwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock/mining"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	appstockcrud.Req
	AppMiningGoodStockReqs []*appmininggoodstockcrud.Req
	AppSpotLocked          *decimal.Decimal
	AppStockLockState      *types.AppStockLockState
	LockID                 *uuid.UUID
	Rollback               *bool
	Stocks                 []*LockStock
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
		if amount.Cmp(decimal.NewFromInt(0)) < 0 {
			return fmt.Errorf("invalid reserved")
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
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid locked")
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
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid locked")
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
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid locked")
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
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid locked")
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

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appgoodid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.AppGoodID = &_id
		return nil
	}
}

func WithStocks(stocks []*npool.LocksRequest_XStock, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, stock := range stocks {
			_stock := &LockStock{}

			entID, err := uuid.Parse(stock.EntID)
			if err != nil {
				return err
			}
			_stock.EntID = &entID

			appGoodID, err := uuid.Parse(stock.AppGoodID)
			if err != nil {
				return err
			}
			_stock.AppGoodID = &appGoodID

			units, err := decimal.NewFromString(stock.Units)
			if err != nil {
				return err
			}
			_stock.Locked = &units

			appSpotUnits, err := decimal.NewFromString(stock.AppSpotUnits)
			if err != nil {
				return err
			}
			_stock.AppSpotLocked = &appSpotUnits

			h.Stocks = append(h.Stocks, _stock)
		}
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
		case types.AppStockLockState_AppStockCanceled:
		default:
			return fmt.Errorf("invalid appstocklockstate")
		}
		h.AppStockLockState = e
		return nil
	}
}

func WithAppMiningGoodStocks(stocks []*appmininggoodstockpwpb.StockReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, _stock := range stocks {
			entID := func() *uuid.UUID {
				uid, err := uuid.Parse(_stock.GetEntID())
				if err != nil {
					return nil
				}
				return &uid
			}()
			miningGoodStockID, err := uuid.Parse(_stock.GetMiningGoodStockID())
			if err != nil {
				return err
			}
			amount, err := decimal.NewFromString(_stock.GetReserved())
			if err != nil {
				return err
			}
			h.AppMiningGoodStockReqs = append(h.AppMiningGoodStockReqs, &appmininggoodstockcrud.Req{
				EntID:             entID,
				MiningGoodStockID: &miningGoodStockID,
				Reserved:          &amount,
			})
		}
		return nil
	}
}
