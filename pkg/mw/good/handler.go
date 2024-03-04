//nolint:dupl
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
	Posters               []string
	Labels                []types.GoodLabel
	AppReserved           *decimal.Decimal
	RewardState           *types.BenefitState
	RewardAt              *uint32
	RewardTID             *uuid.UUID
	NextRewardStartAmount *decimal.Decimal
	RewardAmount          *decimal.Decimal
	UnitRewardAmount      *decimal.Decimal
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
			deviceinfo1.WithEntID(id, true),
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
		h.DeviceInfoID = handler.EntID
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
			vendorlocation1.WithEntID(id, true),
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
		h.VendorLocationID = handler.EntID
		return nil
	}
}

func WithUnitPrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid unitprice")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.UnitPrice = &amount
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
		case types.BenefitType_BenefitTypeOffline:
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
		case types.GoodType_TechniqueServiceFee:
		case types.GoodType_ElectricityFee:
		case types.GoodType_MachineRenting:
			fallthrough //nolint
		case types.GoodType_MachineHosting:
			fallthrough //nolint
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

func WithQuantityUnit(s *string, must bool) func(context.Context, *Handler) error {
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
		h.QuantityUnit = s
		return nil
	}
}

func WithQuantityUnitAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid unitamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.QuantityUnitAmount = &amount
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

func WithStartMode(e *types.GoodStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid goodstartmode")
			}
			return nil
		}
		switch *e {
		case types.GoodStartMode_GoodStartModeTBD:
		case types.GoodStartMode_GoodStartModeConfirmed:
		case types.GoodStartMode_GoodStartModeNextDay:
		case types.GoodStartMode_GoodStartModeInstantly:
		case types.GoodStartMode_GoodStartModePreset:
		default:
			return fmt.Errorf("invalid goodstartmode")
		}
		h.StartMode = e
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

func WithLabels(es []types.GoodLabel, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, e := range es {
			switch e {
			case types.GoodLabel_GoodLabelPromotion:
			case types.GoodLabel_GoodLabelNoviceExclusive:
			case types.GoodLabel_GoodLabelInnovationStarter:
			case types.GoodLabel_GoodLabelLoyaltyExclusive:
			default:
				return fmt.Errorf("invalid label")
			}
		}
		h.Labels = es
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

func WithAppReserved(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid appreserved")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.AppReserved = &amount
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
		case types.BenefitState_BenefitUserBookKeeping:
		case types.BenefitState_BenefitSimulateBookKeeping:
		case types.BenefitState_BenefitDone:
		case types.BenefitState_BenefitFail:
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

func WithRewardAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid rewardat")
			}
			return nil
		}
		h.RewardAt = n
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

func WithRewardAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid rewardamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.RewardAmount = &amount
		return nil
	}
}

func WithUnitRewardAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid unitrewardamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.UnitRewardAmount = &amount
		return nil
	}
}

func WithUnitType(e *types.GoodUnitType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid unittype")
			}
			return nil
		}
		switch *e {
		case types.GoodUnitType_GoodUnitByDuration:
		case types.GoodUnitType_GoodUnitByQuantity:
		case types.GoodUnitType_GoodUnitByDurationAndQuantity:
		default:
			return fmt.Errorf("invalid unittype")
		}
		h.UnitType = e
		return nil
	}
}

func WithQuantityCalculateType(e *types.GoodUnitCalculateType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid quantitycalculatetype")
			}
			return nil
		}
		switch *e {
		case types.GoodUnitCalculateType_GoodUnitCalculateBySelf:
		case types.GoodUnitCalculateType_GoodUnitCalculateByParent:
		default:
			return fmt.Errorf("invalid quantitycalculatetype")
		}
		h.QuantityCalculateType = e
		return nil
	}
}

func WithDurationType(e *types.GoodDurationType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid durationtype")
			}
			return nil
		}
		switch *e {
		case types.GoodDurationType_GoodDurationByHour:
		case types.GoodDurationType_GoodDurationByDay:
		case types.GoodDurationType_GoodDurationByMonth:
		case types.GoodDurationType_GoodDurationByYear:
		default:
			return fmt.Errorf("invalid durationtype")
		}
		h.DurationType = e
		return nil
	}
}

func WithSettlementType(e *types.GoodSettlementType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid settlementtype")
			}
			return nil
		}
		switch *e {
		case types.GoodSettlementType_GoodSettledByCash:
		case types.GoodSettlementType_GoodSettledByProfit:
		default:
			return fmt.Errorf("invalid settlementtype")
		}
		h.SettlementType = e
		return nil
	}
}

func WithDurationCalculateType(e *types.GoodUnitCalculateType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid durationcalculatetype")
			}
			return nil
		}
		switch *e {
		case types.GoodUnitCalculateType_GoodUnitCalculateBySelf:
		case types.GoodUnitCalculateType_GoodUnitCalculateByParent:
		default:
			return fmt.Errorf("invalid durationcalculatetype")
		}
		h.DurationCalculateType = e
		return nil
	}
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Conds = &goodcrud.Conds{}
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
		if conds.RewardState != nil {
			h.Conds.RewardState = &cruder.Cond{
				Op:  conds.GetRewardState().GetOp(),
				Val: types.BenefitState(conds.GetRewardState().GetValue()),
			}
		}
		if conds.RewardAt != nil {
			h.Conds.RewardAt = &cruder.Cond{
				Op:  conds.GetRewardAt().GetOp(),
				Val: conds.GetRewardAt().GetValue(),
			}
		}
		if conds.EntIDs != nil {
			ids := []uuid.UUID{}
			for _, id := range conds.GetEntIDs().GetValue() {
				_id, err := uuid.Parse(id)
				if err != nil {
					return err
				}
				ids = append(ids, _id)
			}
			h.Conds.EntIDs = &cruder.Cond{Op: conds.GetEntIDs().GetOp(), Val: ids}
		}
		if conds.IDs != nil {
			h.Conds.IDs = &cruder.Cond{Op: conds.GetIDs().GetOp(), Val: conds.GetIDs().GetValue()}
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
