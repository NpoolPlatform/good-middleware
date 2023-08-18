package like

import (
	"context"
	"fmt"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	likecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateLike(ctx context.Context, tx *ent.Tx) error {
	if _, err := likecrud.UpdateSet(
		tx.Like.UpdateOneID(*h.ID),
		&likecrud.Req{
			Like: h.Like,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateGoodLike(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(tx.ExtraInfo.Query(), &extrainfocrud.Conds{
		GoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
	})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	if *h.Like && info.Likes > 1 {
		info.Likes -= 1
	} else if info.Dislikes > 1 {
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

func (h *Handler) UpdateLike(ctx context.Context) (*npool.Like, error) {
	info, err := h.GetLike(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid like")
	}

	goodID := uuid.MustParse(info.GoodID)
	h.GoodID = &goodID
	handler := &updateHandler{
		Handler: h,
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