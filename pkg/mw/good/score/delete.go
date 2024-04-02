package score

import (
	"context"
	"fmt"
	"time"

	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/extrainfo"
	scorecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/score"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type deleteHandler struct {
	*Handler
	score   decimal.Decimal
	appgood *appgoodpb.Good
}

func (h *deleteHandler) getAppGood(ctx context.Context) error {
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
	if err := handler.getAppGood(ctx); err != nil {
		return nil, err
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
