package like

import (
	"context"
	"fmt"
	"time"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	likecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteLike(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := likecrud.UpdateSet(
		tx.Like.UpdateOneID(*h.ID),
		&likecrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) subGoodLike(ctx context.Context, tx *ent.Tx, like bool) error {
	appGood, err := tx.AppGood.Get(ctx, *h.AppGoodID)
	if err != nil {
		return err
	}
	if appGood == nil {
		return fmt.Errorf("app good not found %v", *h.AppGoodID)
	}

	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: appGood.GoodID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	if like && info.Likes >= 1 {
		info.Likes -= 1
	} else if info.Dislikes >= 1 {
		info.Dislikes -= 1
	} else {
		return fmt.Errorf("not allowed")
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

func (h *Handler) DeleteLike(ctx context.Context) (*npool.Like, error) {
	info, err := h.GetLike(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	appGoodID := uuid.MustParse(info.AppGoodID)
	h.AppGoodID = &appGoodID
	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteLike(ctx, tx); err != nil {
			return err
		}
		if err := handler.subGoodLike(ctx, tx, info.Like); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
