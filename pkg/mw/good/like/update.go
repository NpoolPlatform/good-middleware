package like

import (
	"context"
	"fmt"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	likecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entlike "github.com/NpoolPlatform/good-middleware/pkg/db/ent/like"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"

	"github.com/google/uuid"
)

type updateHandler struct {
	*Handler
	like int
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
