package score

import (
	"context"
	"fmt"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	scorecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/score"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entscore "github.com/NpoolPlatform/good-middleware/pkg/db/ent/score"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
	score decimal.Decimal
}

func (h *updateHandler) updateScore(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Score.
		Query().
		Where(
			entscore.ID(*h.ID),
			entscore.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	h.score = info.Score

	if _, err := scorecrud.UpdateSet(
		info.Update(),
		&scorecrud.Req{
			Score: h.Score,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateGoodScore(ctx context.Context, tx *ent.Tx) error {
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

	info.Score = info.Score.
		Mul(decimal.NewFromInt(int64(info.ScoreCount))).
		Sub(h.score).
		Add(*h.Score).
		Div(decimal.NewFromInt(int64(info.ScoreCount)))

	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Score: &info.Score,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateScore(ctx context.Context) (*npool.Score, error) {
	info, err := h.GetScore(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, fmt.Errorf("invalid score")
	}

	if h.Score == nil {
		return info, nil
	}

	goodID := uuid.MustParse(info.GoodID)
	h.GoodID = &goodID
	handler := &updateHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateScore(ctx, tx); err != nil {
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

	return h.GetScore(ctx)
}
