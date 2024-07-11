package like

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/extrainfo"
	likecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
	now uint32
}

func (h *deleteHandler) deleteLike(ctx context.Context, tx *ent.Tx) error {
	if _, err := likecrud.UpdateSet(
		tx.Like.UpdateOneID(*h.ID),
		&likecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) subGoodLike(ctx context.Context, tx *ent.Tx, like bool) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.ForUpdate().Only(ctx)
	if err != nil {
		return err
	}
	if like && info.Likes >= 1 {
		info.Likes -= 1
	} else if info.Dislikes >= 1 {
		info.Dislikes -= 1
	} else {
		return wlog.Errorf("not allowed")
	}
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Likes:    &info.Likes,
			Dislikes: &info.Dislikes,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteLike(ctx context.Context) error {
	info, err := h.GetLike(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	h.AppGoodID = func() *uuid.UUID { uid, _ := uuid.Parse(info.AppGoodID); return &uid }()
	handler := &deleteHandler{
		Handler: h,
		now:     uint32(time.Now().Unix()),
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteLike(ctx, tx); err != nil {
			return err
		}
		return handler.subGoodLike(ctx, tx, info.Like)
	})
}
