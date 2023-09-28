package score

import (
	"context"
	"fmt"
	"time"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	scorecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/score"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
	score decimal.Decimal
}

func (h *deleteHandler) deleteScore(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := scorecrud.UpdateSet(
		tx.Score.UpdateOneID(*h.ID),
		&scorecrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) updateGoodScore(ctx context.Context, tx *ent.Tx) error {
	appGood, err := tx.AppGood.Get(ctx, *h.AppGoodID)
	if err != nil {
		return err
	}
	if appGood == nil {
		return fmt.Errorf("app good not found %v", *h.AppGoodID)
	}

	stm, err := extrainfocrud.SetQueryConds(tx.ExtraInfo.Query(), &extrainfocrud.Conds{
		GoodID: &cruder.Cond{Op: cruder.EQ, Val: appGood.GoodID},
	})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}

	if info.ScoreCount == 0 {
		return fmt.Errorf("invalid scorecount")
	}

	if info.ScoreCount == 1 {
		info.Score = decimal.NewFromInt(0)
	} else {
		info.Score = info.Score.
			Mul(decimal.NewFromInt(int64(info.ScoreCount))).
			Sub(h.score).
			Div(decimal.NewFromInt(int64(info.ScoreCount - 1)))
	}
	info.ScoreCount--

	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Score:      &info.Score,
			ScoreCount: &info.ScoreCount,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteScore(ctx context.Context) (*npool.Score, error) {
	info, err := h.GetScore(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}

	appGoodID, err := uuid.Parse(info.AppGoodID)
	if err != nil {
		return nil, err
	}
	h.AppGoodID = &appGoodID
	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteScore(ctx, tx); err != nil {
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
