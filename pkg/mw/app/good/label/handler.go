package label

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	appgoodlabelcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/label"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	appgoodlabelcrud.Req
	LabelConds       *appgoodlabelcrud.Conds
	GoodBaseConds    *goodbasecrud.Conds
	AppGoodBaseConds *appgoodbasecrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		LabelConds:    &appgoodlabelcrud.Conds{},
		GoodBaseConds: &goodbasecrud.Conds{},
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
				return wlog.Errorf("invalid id")
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
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := appgoodbase1.NewHandler(
			ctx,
			appgoodbase1.WithEntID(id, true),
		)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err := handler.ExistGoodBase(ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		if !exist {
			return wlog.Errorf("invalid appgood")
		}
		h.AppGoodID = handler.EntID
		return nil
	}
}

func WithIcon(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Icon = s
		return nil
	}
}

func WithIconBgColor(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.IconBgColor = s
		return nil
	}
}

func WithLabel(e *types.GoodLabel, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid label")
			}
			return nil
		}
		switch *e {
		case types.GoodLabel_GoodLabelPromotion:
		case types.GoodLabel_GoodLabelNoviceExclusive:
		case types.GoodLabel_GoodLabelInnovationStarter:
		case types.GoodLabel_GoodLabelLoyaltyExclusive:
		default:
			return wlog.Errorf("invalid label")
		}
		h.Label = e
		return nil
	}
}

func WithLabelBgColor(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.LabelBgColor = s
		return nil
	}
}

func WithIndex(n *uint8, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = n
		return nil
	}
}

func (h *Handler) withAppGoodLabelConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.LabelConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.LabelConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.LabelConds.AppGoodID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.LabelConds.AppGoodIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withGoodBaseConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
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
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withAppGoodBaseConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.AppID = &cruder.Cond{
			Op: conds.GetAppID().GetOp(), Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.AppGoodBaseConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.AppGoodBaseConds.EntID = &cruder.Cond{
			Op: conds.GetAppGoodID().GetOp(), Val: id,
		}
	}
	if conds.AppGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.AppGoodBaseConds.EntIDs = &cruder.Cond{
			Op: conds.GetAppGoodIDs().GetOp(), Val: ids,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withAppGoodLabelConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withGoodBaseConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withAppGoodBaseConds(conds)
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
