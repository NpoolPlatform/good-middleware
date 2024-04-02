package recommend

import (
	"context"
	"fmt"
	"time"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/extrainfo"
	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/recommend"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"

	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteRecommend(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := recommendcrud.UpdateSet(
		tx.Recommend.UpdateOneID(*h.ID),
		&recommendcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) updateGoodScore(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.GoodID},
		},
	)
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	if info.RecommendCount == 0 {
		return fmt.Errorf("invalid recommendcount")
	}
	info.RecommendCount--

	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			RecommendCount: &info.RecommendCount,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteRecommend(ctx context.Context) (*npool.Recommend, error) {
	info, err := h.GetRecommend(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	goodID, err := uuid.Parse(info.GoodID)
	if err != nil {
		return nil, err
	}
	h.GoodID = &goodID
	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteRecommend(ctx, tx); err != nil {
			return err
		}
		if err := handler.updateGoodScore(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
