package like

import (
	"context"
	"fmt"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	likecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entlike "github.com/NpoolPlatform/good-middleware/pkg/db/ent/like"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	like    int
	appgood *appgoodpb.Good
}

func (h *updateHandler) getAppGood(ctx context.Context) error {
	handler, err := appgood1.NewHandler(ctx)
	if err != nil {
		return err
	}
	handler.EntID = h.AppGoodID
	info, err := handler.GetGood(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("app good not found %v", *h.AppGoodID)
	}
	h.appgood = info
	return nil
}

func (h *updateHandler) updateLike(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Like.
		Query().
		Where(
			entlike.ID(*h.ID),
			entlike.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if *h.Like && !info.Like {
		h.like = 1
	} else if !*h.Like && info.Like {
		h.like = -1
	}

	if _, err := likecrud.UpdateSet(
		info.Update(),
		&likecrud.Req{
			Like: h.Like,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateGoodLike(ctx context.Context, tx *ent.Tx) error {
	goodid := uuid.MustParse(h.appgood.GoodID)
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: goodid},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	info.Likes += uint32(h.like)
	info.Dislikes -= uint32(h.like)

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

func (h *Handler) UpdateLike(ctx context.Context) (*npool.Like, error) {
	info, err := h.GetLike(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid like")
	}

	if h.Like == nil {
		return info, nil
	}

	appGoodID := uuid.MustParse(info.AppGoodID)
	h.AppGoodID = &appGoodID
	handler := &updateHandler{
		Handler: h,
	}
	if err := handler.getAppGood(ctx); err != nil {
		return nil, err
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateLike(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateGoodLike(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetLike(ctx)
}
