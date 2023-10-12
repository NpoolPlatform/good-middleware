package score

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	scorecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/score"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/score"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type createHandler struct {
	*Handler
	appgood *ent.AppGood
}

func (h *createHandler) getAppGood(ctx context.Context, tx *ent.Tx) error {
	appGood, err := tx.AppGood.Get(ctx, *h.AppGoodID)
	if err != nil {
		return err
	}
	if appGood == nil {
		return fmt.Errorf("app good not found %v", *h.AppGoodID)
	}
	h.appgood = appGood
	return nil
}

func (h *createHandler) createScore(ctx context.Context, tx *ent.Tx) error {
	if _, err := scorecrud.CreateSet(
		tx.Score.Create(),
		&scorecrud.Req{
			ID:        h.ID,
			AppID:     h.AppID,
			UserID:    h.UserID,
			GoodID:    &h.appgood.GoodID,
			AppGoodID: h.AppGoodID,
			Score:     h.Score,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) updateGoodScore(ctx context.Context, tx *ent.Tx) error {
	stm, err := extrainfocrud.SetQueryConds(
		tx.ExtraInfo.Query(),
		&extrainfocrud.Conds{
			GoodID: &cruder.Cond{Op: cruder.EQ, Val: h.appgood.GoodID},
		})
	if err != nil {
		return err
	}
	info, err := stm.Only(ctx)
	if err != nil {
		return err
	}
	scoreCount := info.ScoreCount + 1
	score := info.Score.
		Mul(decimal.NewFromInt(int64(info.ScoreCount))).
		Add(*h.Score).
		Div(decimal.NewFromInt(int64(scoreCount)))
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			ScoreCount: &scoreCount,
			Score:      &score,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateScore(ctx context.Context) (*npool.Score, error) {
	key := fmt.Sprintf("%v:%v:%v:%v", basetypes.Prefix_PrefixScoreGood, *h.AppID, *h.UserID, *h.AppGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &scorecrud.Conds{
		AppID:     &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
		UserID:    &cruder.Cond{Op: cruder.EQ, Val: *h.UserID},
		AppGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.AppGoodID},
	}
	exist, err := h.ExistScoreConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.getAppGood(ctx, tx); err != nil {
			return err
		}
		if err := handler.createScore(ctx, tx); err != nil {
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
