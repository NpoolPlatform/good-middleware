package recommend

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/recommend"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	recommendcrud.Req
	RecommendConds   *recommendcrud.Conds
	AppGoodBaseConds *appgoodbasecrud.Conds
	GoodBaseConds    *goodbasecrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		RecommendConds:   &recommendcrud.Conds{},
		AppGoodBaseConds: &appgoodbasecrud.Conds{},
		GoodBaseConds:    &goodbasecrud.Conds{},
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

func WithRecommenderID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid recommenderid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.RecommenderID = &_id
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
		handler, err := appgoodbase1.NewHandler(
			ctx,
			appgoodbase1.WithEntID(id, true),
		)
		if err != nil {
			return fmt.Errorf("invalid appgoodid")
		}
		exist, err := handler.ExistGoodBase(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid appgood")
		}
		h.AppGoodID = handler.EntID
		return nil
	}
}

func WithMessage(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		const leastMessageLen = 10
		if s == nil {
			if must {
				return fmt.Errorf("invalid message")
			}
			return nil
		}
		if len(*s) < leastMessageLen {
			return fmt.Errorf("invalid message")
		}
		h.Message = s
		return nil
	}
}

func WithRecommendIndex(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid recommendindex")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.RecommendIndex = &amount
		return nil
	}
}

func WithHide(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Hide = b
		return nil
	}
}

func WithHideReason(e *types.GoodCommentHideReason, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid hidereason")
			}
			return nil
		}
		switch *e {
		case types.GoodCommentHideReason_GoodCommentHideBySpam:
		case types.GoodCommentHideReason_GoodCommentHideByNotThisGood:
		case types.GoodCommentHideReason_GoodCommentHideByFalseDescription:
		default:
			return fmt.Errorf("invalid hidereason")
		}
		h.HideReason = e
		return nil
	}
}

func (h *Handler) withRecommendConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.RecommendConds.ID = &cruder.Cond{
			Op: conds.GetID().GetOp(), Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.RecommendConds.EntID = &cruder.Cond{Op: conds.GetEntID().GetOp(), Val: id}
	}
	if conds.RecommenderID != nil {
		id, err := uuid.Parse(conds.GetRecommenderID().GetValue())
		if err != nil {
			return err
		}
		h.RecommendConds.RecommenderID = &cruder.Cond{
			Op:  conds.GetRecommenderID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return err
		}
		h.RecommendConds.AppGoodID = &cruder.Cond{
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
		h.RecommendConds.AppGoodIDs = &cruder.Cond{
			Op:  conds.GetAppGoodIDs().GetOp(),
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
			Op: conds.GetAppGoodID().GetOp(), Val: id,
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
		if err := h.withRecommendConds(conds); err != nil {
			return err
		}
		return h.withAppGoodBaseConds(conds)
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
