package powerrental

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	goodcoincrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	powerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/powerrental"
	deviceinfo1 "github.com/NpoolPlatform/good-middleware/pkg/mw/deviceinfo"
	vendorlocation1 "github.com/NpoolPlatform/good-middleware/pkg/mw/vender/location"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	powerrentalcrud.Req
	GoodBaseReq      *goodbasecrud.Req
	PowerRentalConds *powerrentalcrud.Conds
	GoodBaseConds    *goodbasecrud.Conds
	GoodCoinConds    *goodcoincrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		GoodBaseReq:      &goodbasecrud.Req{},
		PowerRentalConds: &powerrentalcrud.Conds{},
		GoodBaseConds:    &goodbasecrud.Conds{},
		GoodCoinConds:    &goodcoincrud.Conds{},
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

func WithEntID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return err
		}
		h.EntID = &id
		return nil
	}
}

func WithGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return err
		}
		h.GoodID = &id
		return nil
	}
}

func WithDeviceTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid devicetypeid")
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
			return fmt.Errorf("invalid devicetypeid")
		}
		h.DeviceTypeID = handler.EntID
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
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid unitprice")
		}
		h.UnitPrice = &amount
		return nil
	}
}

func WithQuantityUnit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid quantityunit")
			}
			return nil
		}
		const leastUnitLen = 2
		if len(*s) < leastUnitLen {
			return fmt.Errorf("invalid quantityunit")
		}
		h.QuantityUnit = s
		return nil
	}
}

func WithQuantityUnitAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid quantityunitamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid quantityunitamount")
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
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid unitlockdeposit")
		}
		h.UnitLockDeposit = &amount
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

func WithGoodType(e *types.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case types.GoodType_PowerRental:
		case types.GoodType_LegacyPowerRental:
		default:
			return fmt.Errorf("invalid goodtype")
		}
		h.GoodBaseReq.GoodType = e
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
		h.GoodBaseReq.BenefitType = e
		return nil
	}
}

func WithName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid name")
			}
			return nil
		}
		const leastNameLen = 3
		if len(*s) < leastNameLen {
			return fmt.Errorf("invalid name")
		}
		h.GoodBaseReq.Name = s
		return nil
	}
}

func WithServiceStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.ServiceStartAt = n
		return nil
	}
}

func WithStartMode(e *types.GoodStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid startmode")
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
			return fmt.Errorf("invalid startmode")
		}
		h.GoodBaseReq.StartMode = e
		return nil
	}
}

func WithTestOnly(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.TestOnly = b
		return nil
	}
}

func WithBenefitIntervalHours(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.BenefitIntervalHours = n
		return nil
	}
}

func WithPurchasable(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.Purchasable = b
		return nil
	}
}

func WithOnline(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.Online = b
		return nil
	}
}

func (h *Handler) withPowerRentalConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.PowerRentalConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.PowerRentalConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return err
		}
		h.PowerRentalConds.GoodID = &cruder.Cond{
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
		h.PowerRentalConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withGoodBaseConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return err
		}
		h.GoodBaseConds.EntID = &cruder.Cond{
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
		h.GoodBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.GoodType != nil {
		h.GoodBaseConds.GoodType = &cruder.Cond{
			Op:  conds.GetGoodType().GetOp(),
			Val: types.GoodType(conds.GetGoodType().GetValue()),
		}
	}
	if conds.GoodTypes != nil {
		es := []types.GoodType{}
		for _, e := range conds.GetGoodTypes().GetValue() {
			es = append(es, types.GoodType(e))
		}
		h.GoodBaseConds.GoodTypes = &cruder.Cond{
			Op:  conds.GetGoodTypes().GetOp(),
			Val: es,
		}
	}
	return nil
}

func (h *Handler) withGoodCoinConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return err
		}
		h.GoodCoinConds.GoodID = &cruder.Cond{
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
		h.GoodCoinConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.CoinTypeID != nil {
		id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
		if err != nil {
			return err
		}
		h.GoodCoinConds.CoinTypeID = &cruder.Cond{
			Op:  conds.GetCoinTypeID().GetOp(),
			Val: id,
		}
	}
	if conds.CoinTypeIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetCoinTypeIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			ids = append(ids, _id)
		}
		h.GoodCoinConds.CoinTypeIDs = &cruder.Cond{
			Op:  conds.GetCoinTypeIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withPowerRentalConds(conds); err != nil {
			return err
		}
		if err := h.withGoodCoinConds(conds); err != nil {
			return err
		}
		return h.withGoodBaseConds(conds)
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
