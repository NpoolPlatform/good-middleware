package good

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	goodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good"
	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	goodcrud.Req
	Total                 *decimal.Decimal
	Locked                *decimal.Decimal
	InService             *decimal.Decimal
	WaitStart             *decimal.Decimal
	Posters               []string
	Labels                []string
	AppLocked             *decimal.Decimal
	Likes                 *uint32
	Dislikes              *uint32
	Rating                *decimal.Decimal
	RewardState           *types.BenefitState
	LastRewardAt          *uint32
	RewardTID             *uuid.UUID
	NextRewardStartAmount *decimal.Decimal
	LastRewardAmount      *decimal.Decimal
	Conds                 *goodcrud.Conds
	Offset                int32
	Limit                 int32
}

const leastStrLen = 3

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

func WithDeviceInfoID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid deviceinfoid")
			}
			return nil
		}
		handler, err := deviceinfo1.NewHandler(
			ctx,
			deviceinfo1.WithID(id),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistDeviceInfo(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid deviceinfo")
		}
		h.DeviceInfoID = handler.ID
		return nil
	}
}

func WithDurationDays(n *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid durationdays")
			}
			return nil
		}
		h.DurationDays = n
		return nil
	}
}

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.CoinTypeID = &_id
		return nil
	}
}

func WithVendorLocationID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid vendorlocationid")
			}
			return nil
		}
		handler, err := vendorlocation1.NewHandler(
			ctx,
			vendorlocation1.WithID(id),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistLocation(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid vendorlocation")
		}
		h.VendorLocationID = handler.ID
		return nil
	}
}

func WithPrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid price")
			}
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

func WithBenefitType(e *types.BenefitType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid benefittype")
			}
			return nil
		}
		switch *e {
		case types.BenefitType_BenefitTypePlatform:
		case types.BenefitType_BenefitTypePool:
		default:
			return fmt.Errorf("invalid benefittype")
		}
		h.BenefitType = e
		return nil
	}
}

func WithGoodType(e *types.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case types.GoodType_PowerRenting:
		case types.GoodType_MachineRenting:
			fallthrough //nolint
		case types.GoodType_MachineHosting:
			fallthrough //nolint
		case types.GoodType_TechniqueServiceFee:
			fallthrough //nolint
		case types.GoodType_ElectricityFee:
			fallthrough //nolint
		case types.GoodType_CrowdFunding:
			fallthrough //nolint
		case types.GoodType_Arbitrage:
			return fmt.Errorf("not implemented")
		default:
			return fmt.Errorf("invalid goodtype")
		}
		h.GoodType = e
		return nil
	}
}

func WithTitle(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid title")
			}
			return nil
		}
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid title")
		}
		h.Title = s
		return nil
	}
}

func WithUnit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid unit")
			}
			return nil
		}
		const leastUnitLen = 2
		if len(*s) < leastUnitLen {
			return fmt.Errorf("invalid unit")
		}
		h.Unit = s
		return nil
	}
}

func WithUnitAmount(n *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid unitamount")
			}
			return nil
		}
		h.UnitAmount = n
		return nil
	}
}

func WithSupportCoinTypeIDs(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, s := range ss {
			id, err := uuid.Parse(s)
			if err != nil {
				return err
			}
			h.SupportCoinTypeIDs = append(h.SupportCoinTypeIDs, id)
		}
		return nil
	}
}

func WithDeliveryAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid deliveryat")
			}
			return nil
		}
		h.DeliveryAt = n
		return nil
	}
}

func WithStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid startat")
			}
			return nil
		}
		h.StartAt = n
		return nil
	}
}

func WithTestOnly(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.TestOnly = b
		return nil
	}
}

func WithTotal(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid total")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.Total = &amount
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

func WithPosters(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, s := range ss {
			if len(s) < leastStrLen {
				return fmt.Errorf("invalid poster")
			}
		}
		h.Posters = ss
		return nil
	}
}

func WithLabels(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, s := range ss {
			if len(s) < leastStrLen {
				return fmt.Errorf("invalid labels")
			}
		}
		h.Labels = ss
		return nil
	}
}

func WithBenefitIntervalHours(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid benefitintervalhours")
			}
			return nil
		}
		h.BenefitIntervalHours = n
		return nil
	}
}

func WithUnitLockDeposit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid unitlockdeposit")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.UnitLockDeposit = &amount
		return nil
	}
}

func WithAppLocked(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid applocked")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.AppLocked = &amount
		return nil
	}
}

func WithLikes(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid likes")
			}
			return nil
		}
		h.Likes = n
		return nil
	}
}

func WithDislikes(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid dislikes")
			}
			return nil
		}
		h.Dislikes = n
		return nil
	}
}

func WithRating(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid rating")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.Rating = &amount
		return nil
	}
}

func WithRewardState(e *types.BenefitState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid rewardstate")
			}
			return nil
		}
		switch *e {
		case types.BenefitState_BenefitWait:
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitBookKeeping:
		default:
			return fmt.Errorf("invalid rewardstate")
		}
		h.RewardState = e
		return nil
	}
}

func WithRewardTID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid rewardtid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.RewardTID = &_id
		return nil
	}
}

func WithLastRewardAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid lastrewardat")
			}
			return nil
		}
		h.LastRewardAt = n
		return nil
	}
}

func WithNextRewardStartAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid nextrewardstartamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.NextRewardStartAmount = &amount
		return nil
	}
}

func WithLastRewardAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid lastrewardamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.LastRewardAmount = &amount
		return nil
	}
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &goodcrud.Conds{}
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
		if conds.DeviceInfoID != nil {
			id, err := uuid.Parse(conds.GetDeviceInfoID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.DeviceInfoID = &cruder.Cond{
				Op:  conds.GetDeviceInfoID().GetOp(),
				Val: id,
			}
		}
		if conds.CoinTypeID != nil {
			id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.CoinTypeID = &cruder.Cond{
				Op:  conds.GetCoinTypeID().GetOp(),
				Val: id,
			}
		}
		if conds.VendorLocationID != nil {
			id, err := uuid.Parse(conds.GetVendorLocationID().GetValue())
			if err != nil {
				return err
			}
			h.Conds.VendorLocationID = &cruder.Cond{
				Op:  conds.GetVendorLocationID().GetOp(),
				Val: id,
			}
		}
		if conds.BenefitType != nil {
			h.Conds.BenefitType = &cruder.Cond{
				Op:  conds.GetBenefitType().GetOp(),
				Val: types.BenefitType(conds.GetBenefitType().GetValue()),
			}
		}
		if conds.GoodType != nil {
			h.Conds.GoodType = &cruder.Cond{
				Op:  conds.GetGoodType().GetOp(),
				Val: types.GoodType(conds.GetGoodType().GetValue()),
			}
		}
		if conds.IDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.IDs = &cruder.Cond{
				Op:  conds.GetIDs().GetOp(),
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
