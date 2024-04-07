package powerrental

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/extrainfo"
	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	appgoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	apppowerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental"
	goodcoincrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	powerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/powerrental"
	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	apppowerrentalcrud.Req
	AppGoodBaseReq      *appgoodbasecrud.Req
	AppGoodStockReq     *appgoodstockcrud.Req
	ExtraInfoReq        *extrainfocrud.Req
	AppPowerRentalConds *apppowerrentalcrud.Conds
	PowerRentalConds    *powerrentalcrud.Conds
	AppGoodBaseConds    *appgoodbasecrud.Conds
	GoodBaseConds       *goodbasecrud.Conds
	GoodCoinConds       *goodcoincrud.Conds
	Offset              int32
	Limit               int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		AppGoodBaseReq:      &appgoodbasecrud.Req{},
		AppGoodStockReq:     &appgoodstockcrud.Req{},
		ExtraInfoReq:        &extrainfocrud.Req{},
		AppPowerRentalConds: &apppowerrentalcrud.Conds{},
		PowerRentalConds:    &powerrentalcrud.Conds{},
		AppGoodBaseConds:    &appgoodbasecrud.Conds{},
		GoodBaseConds:       &goodbasecrud.Conds{},
		GoodCoinConds:       &goodcoincrud.Conds{},
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

func WithAppID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return err
		}
		h.AppGoodBaseReq.AppID = &id
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
		handler, err := goodbase1.NewHandler(
			ctx,
			goodbase1.WithEntID(s, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistGoodBase(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid goodid")
		}
		h.AppGoodBaseReq.GoodID = handler.EntID
		return nil
	}
}

func WithAppGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid appgoodid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return err
		}
		h.AppGoodID = &id
		h.AppGoodBaseReq.EntID = &id
		h.AppGoodStockReq.AppGoodID = &id
		h.ExtraInfoReq.AppGoodID = &id
		return nil
	}
}

func WithPurchasable(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.Purchasable = b
		return nil
	}
}

func WithEnableProductPage(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.EnableProductPage = b
		return nil
	}
}

func WithProductPage(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.ProductPage = s
		return nil
	}
}

func WithOnline(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.Online = b
		return nil
	}
}

func WithVisible(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.Visible = b
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
		if len(*s) < 3 {
			return fmt.Errorf("invalid name")
		}
		h.AppGoodBaseReq.Name = s
		return nil
	}
}

func WithDisplayIndex(n *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.DisplayIndex = n
		return nil
	}
}

func WithBanner(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.AppGoodBaseReq.Banner = s
		return nil
	}
}

func WithServiceStartAt(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if u == nil {
			if must {
				return fmt.Errorf("invalid servicestartat")
			}
			return nil
		}
		h.ServiceStartAt = u
		return nil
	}
}

func WithCancelMode(e *types.CancelMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid cancelmode")
			}
			return nil
		}
		switch *e {
		case types.CancelMode_Uncancellable:
		case types.CancelMode_CancellableBeforeStart:
		case types.CancelMode_CancellableBeforeBenefit:
		default:
			return fmt.Errorf("invalid cancelmode")
		}
		h.CancelMode = e
		return nil
	}
}

func WithCancelableBeforeStartSeconds(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.CancelableBeforeStartSeconds = u
		return nil
	}
}

func WithEnableSetCommission(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnableSetCommission = b
		return nil
	}
}

func WithMinOrderAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid minorderamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.MinOrderAmount = &amount
		return nil
	}
}

func WithMaxOrderAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid maxorderamount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid maxorderamount")
		}
		h.MaxOrderAmount = &amount
		return nil
	}
}

func WithMaxUserAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid maxuseramount")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid maxuseramount")
		}
		h.MaxUserAmount = &amount
		return nil
	}
}

func WithMinOrderDuration(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MinOrderDuration = u
		return nil
	}
}

func WithMaxOrderDuration(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxOrderDuration = u
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

func WithSaleStartAt(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SaleStartAt = u
		return nil
	}
}

func WithSaleEndAt(u *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SaleEndAt = u
		return nil
	}
}

func WithSaleMode(e *types.GoodSaleMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid salemode")
			}
			return nil
		}
		switch *e {
		case types.GoodSaleMode_GoodSaleModeMainnetSpot:
		case types.GoodSaleMode_GoodSaleModeMainnetPresaleSpot:
		case types.GoodSaleMode_GoodSaleModeMainnetPresaleScratch:
		case types.GoodSaleMode_GoodSaleModeTestnetPresale:
		default:
			return fmt.Errorf("invalid salemode")
		}
		h.SaleMode = e
		return nil
	}
}

func WithFixedDuration(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.FixedDuration = b
		return nil
	}
}

func WithPackageWithRequireds(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.PackageWithRequireds = b
		return nil
	}
}

func WithAppGoodStockID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid appgoodstockid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return err
		}
		h.AppGoodStockReq.EntID = &id
		return nil
	}
}

func (h *Handler) withAppPowerRentalConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.AppPowerRentalConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.AppPowerRentalConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return err
		}
		h.AppPowerRentalConds.AppGoodID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
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
		h.AppPowerRentalConds.AppGoodIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withPowerRentalConds(conds *npool.Conds) error {
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

func (h *Handler) withAppGoodBaseConds(conds *npool.Conds) error {
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return err
		}
		h.AppGoodBaseConds.EntID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
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
		h.AppGoodBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
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

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withAppPowerRentalConds(conds); err != nil {
			return err
		}
		if err := h.withAppGoodBaseConds(conds); err != nil {
			return err
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
