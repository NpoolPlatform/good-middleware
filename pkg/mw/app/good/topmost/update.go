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
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateTopMost(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostcrud.UpdateSet(
		tx.TopMost.UpdateOneID(*h.ID),
		&topmostcrud.Req{
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

func (h *Handler) UpdateTopMost(ctx context.Context) (*npool.TopMost, error) {
	if h.TopMostType != nil {
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
				TopMostType: &cruder.Cond{Op: cruder.EQ, Val: uint32(*h.TopMostType)},
			}
			exist, err := h.ExistTopMostConds(ctx)
			if err != nil {
				return nil, err
			}
			if exist {
				return nil, fmt.Errorf("already exists")
			}
		}
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateTopMost(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTopMost(ctx)
}
