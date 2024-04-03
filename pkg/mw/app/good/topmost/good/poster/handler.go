package poster

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	topmostcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost"
	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	topmostgoodpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/poster"
	topmostgood1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/topmost/good"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"

	"github.com/google/uuid"
)

type Handler struct {
	ID *uint32
	topmostgoodpostercrud.Req
	PosterConds      *topmostgoodpostercrud.Conds
	TopMostConds     *topmostcrud.Conds
	TopMostGoodConds *topmostgoodcrud.Conds
	Offset           int32
	Limit            int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		PosterConds:      &topmostgoodpostercrud.Conds{},
		TopMostGoodConds: &topmostgoodcrud.Conds{},
		TopMostConds:     &topmostcrud.Conds{},
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

func WithTopMostGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		handler, err := topmostgood1.NewHandler(
			ctx,
			topmostgood1.WithEntID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistTopMostGood(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid topmostgood")
		}
		h.TopMostGoodID = handler.EntID
		return nil
	}
}

func WithPoster(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Poster = s
		return nil
	}
}

func WithIndex(n *uint8, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Index = n
		return nil
	}
}

func (h *Handler) withTopMostGoodPosterConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.PosterConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.PosterConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.TopMostGoodID != nil {
		id, err := uuid.Parse(conds.GetTopMostGoodID().GetValue())
		if err != nil {
			return err
		}
		h.PosterConds.TopMostGoodID = &cruder.Cond{
			Op:  conds.GetTopMostGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.TopMostGoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetTopMostGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return err
			}
			ids = append(ids, _id)
		}
		h.PosterConds.TopMostGoodIDs = &cruder.Cond{
			Op:  conds.GetTopMostGoodIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withTopMostGoodConds(conds *npool.Conds) error {
	if conds.TopMostGoodID != nil {
		id, err := uuid.Parse(conds.GetTopMostGoodID().GetValue())
		if err != nil {
			return err
		}
		h.TopMostGoodConds.EntID = &cruder.Cond{
			Op:  conds.GetTopMostGoodID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func (h *Handler) withTopMostConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return err
		}
		h.TopMostConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	return nil
}

//nolint:gocyclo
func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withTopMostGoodPosterConds(conds); err != nil {
			return err
		}
		if err := h.withTopMostConds(conds); err != nil {
			return err
		}
		return h.withTopMostGoodConds(conds)
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
