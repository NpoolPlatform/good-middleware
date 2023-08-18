package topmostgood

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	topmostgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createTopMostGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := topmostgoodcrud.CreateSet(
		tx.TopMostGood.Create(),
		&topmostgoodcrud.Req{
			ID:           h.ID,
			AppID:        h.AppID,
			GoodID:       h.GoodID,
			TopMostID:    h.TopMostID,
			AppGoodID:    h.AppGoodID,
			DisplayIndex: h.DisplayIndex,
			Posters:      h.Posters,
			Price:        h.Price,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateTopMostGood(ctx context.Context) (*npool.TopMostGood, error) {
	key := fmt.Sprintf("%v:%v:%v:%v:%v", basetypes.Prefix_PrefixCreateTopMostGood, *h.AppID, *h.TopMostID, *h.GoodID, *h.AppGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createTopMostGood(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetTopMostGood(ctx)
}
