package good

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	goodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	rewardcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createExtraInfo(ctx context.Context, tx *ent.Tx) error {
	labels := []string{}
	for _, label := range h.Labels {
		labels = append(labels, label.String())
	}

	if _, err := extrainfocrud.CreateSet(
		tx.ExtraInfo.Create(),
		&extrainfocrud.Req{
			GoodID:  h.EntID,
			Posters: h.Posters,
			Labels:  labels,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createReward(ctx context.Context, tx *ent.Tx) error {
	if _, err := rewardcrud.CreateSet(
		tx.GoodReward.Create(),
		&rewardcrud.Req{
			GoodID: h.EntID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createStock(ctx context.Context, tx *ent.Tx) error {
	if _, err := stockcrud.CreateSet(
		tx.Stock.Create(),
		&stockcrud.Req{
			GoodID:       h.EntID,
			Total:        h.Total,
			SpotQuantity: h.Total,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := goodcrud.CreateSet(
		tx.Good.Create(),
		&goodcrud.Req{
			EntID:                h.EntID,
			DeviceInfoID:         h.DeviceInfoID,
			DurationDays:         h.DurationDays,
			CoinTypeID:           h.CoinTypeID,
			VendorLocationID:     h.VendorLocationID,
			Price:                h.Price,
			BenefitType:          h.BenefitType,
			GoodType:             h.GoodType,
			Title:                h.Title,
			Unit:                 h.Unit,
			UnitAmount:           h.UnitAmount,
			DeliveryAt:           h.DeliveryAt,
			StartAt:              h.StartAt,
			StartMode:            h.StartMode,
			TestOnly:             h.TestOnly,
			BenefitIntervalHours: h.BenefitIntervalHours,
			UnitLockDeposit:      h.UnitLockDeposit,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateGood(ctx context.Context) (*npool.Good, error) {
	key := h.lockKey()
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &goodcrud.Conds{
		Title: &cruder.Cond{Op: cruder.EQ, Val: *h.Title},
	}
	exist, err := h.ExistGoodConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	id := uuid.New()
	if h.EntID == nil {
		h.EntID = &id
	}

	handler := &createHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createExtraInfo(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createReward(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createStock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.createGood(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
