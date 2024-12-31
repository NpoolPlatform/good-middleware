package pledge

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	rewardcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward"
	pledgecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/pledge"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodcoin "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
)

type deleteHandler struct {
	*pledgeGoodQueryHandler
	now uint32
}

func (h *deleteHandler) deleteGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.goodBase == nil {
		return wlog.Errorf("invalid goodbase")
	}
	if _, err := goodbasecrud.UpdateSet(
		tx.GoodBase.UpdateOneID(h.goodBase.ID),
		&goodbasecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deletePledge(ctx context.Context, tx *ent.Tx) error {
	if h.pledge == nil {
		return wlog.Errorf("invalid pledge")
	}
	if _, err := pledgecrud.UpdateSet(
		tx.Pledge.UpdateOneID(h.pledge.ID),
		&pledgecrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteReward(ctx context.Context, tx *ent.Tx) error {
	if h.goodReward == nil {
		return nil
	}
	if _, err := rewardcrud.UpdateSet(
		tx.GoodReward.UpdateOneID(h.goodReward.ID),
		&rewardcrud.Req{
			DeletedAt: &h.now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteGoodCoin(ctx context.Context, tx *ent.Tx) error {
	fmt.Println("--h.goodCoins: ", h.goodCoins)
	if h.goodCoins == nil {
		return nil
	}
	ids := []uint32{}
	for _, coin := range h.goodCoins {
		ids = append(ids, coin.ID)
	}
	if len(ids) == 0 {
		return nil
	}
	if _, err := tx.GoodCoin.
		Update().
		Where(
			entgoodcoin.IDIn(ids...),
			entgoodcoin.DeletedAt(0),
		).
		SetDeletedAt(h.now).
		Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeletePledge(ctx context.Context) error {
	handler := &deleteHandler{
		pledgeGoodQueryHandler: &pledgeGoodQueryHandler{
			Handler: h,
		},
		now: uint32(time.Now().Unix()),
	}

	if err := handler.getPledgeGood(ctx); err != nil {
		return err
	}
	if handler.pledge == nil {
		return nil
	}
	h.ID = &handler.pledge.ID

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteGoodBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteReward(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteGoodCoin(_ctx, tx); err != nil {
			return err
		}
		return handler.deletePledge(_ctx, tx)
	})
}
