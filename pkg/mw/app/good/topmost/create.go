package topmost

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	topmostcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createTopMost(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostcrud.CreateSet(
		tx.TopMost.Create(),
		&topmostcrud.Req{
			ID:                     h.ID,
			AppID:                  h.AppID,
			TopMostType:            h.TopMostType,
			Title:                  h.Title,
			Message:                h.Message,
			Posters:                h.Posters,
			StartAt:                h.StartAt,
			EndAt:                  h.EndAt,
			ThresholdCredits:       h.ThresholdCredits,
			RegisterElapsedSeconds: h.RegisterElapsedSeconds,
			ThresholdPurchases:     h.ThresholdPurchases,
			ThresholdPaymentAmount: h.ThresholdPaymentAmount,
			KycMust:                h.KycMust,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateTopMost(ctx context.Context) (*npool.TopMost, error) {
	if err := h.formalizeStartEnd(); err != nil {
		return nil, err
	}
	if err := h.checkPromotion(ctx); err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateTopMost, *h.AppID, h.TopMostType)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	if *h.TopMostType != types.GoodTopMostType_TopMostPromotion {
		h.Conds = &topmostcrud.Conds{
			AppID:       &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			TopMostType: &cruder.Cond{Op: cruder.EQ, Val: *h.TopMostType},
		}
		exist, err := h.ExistTopMostConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("already exists")
		}
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createTopMost(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTopMost(ctx)
}
