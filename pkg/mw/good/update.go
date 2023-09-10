package good

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	goodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	rewardcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward"
	rewardhistorycrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward/history"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entextrainfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"

	"github.com/shopspring/decimal"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateExtraInfo(ctx context.Context, tx *ent.Tx) error {
	labels := []string{}
	for _, label := range h.Labels {
		labels = append(labels, label.String())
	}

	info, err := tx.
		ExtraInfo.
		Query().
		Where(
			entextrainfo.GoodID(*h.ID),
			entextrainfo.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			Posters: h.Posters,
			Labels:  labels,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:gocyclo,funlen
func (h *updateHandler) updateReward(ctx context.Context, tx *ent.Tx) error {
	if h.RewardState == nil {
		return nil
	}

	info, err := tx.
		GoodReward.
		Query().
		Where(
			entgoodreward.GoodID(*h.ID),
			entgoodreward.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	switch info.RewardState {
	case types.BenefitState_BenefitWait.String():
		switch *h.RewardState {
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitFail:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", info.RewardState, *h.RewardState)
		}
	case types.BenefitState_BenefitTransferring.String():
		switch *h.RewardState {
		case types.BenefitState_BenefitBookKeeping:
		case types.BenefitState_BenefitFail:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", info.RewardState, *h.RewardState)
		}
	case types.BenefitState_BenefitBookKeeping.String():
		switch *h.RewardState {
		case types.BenefitState_BenefitDone:
		case types.BenefitState_BenefitFail:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", info.RewardState, *h.RewardState)
		}
	case types.BenefitState_BenefitDone.String():
		fallthrough //nolint
	case types.BenefitState_BenefitFail.String():
		if *h.RewardState != types.BenefitState_BenefitWait {
			return fmt.Errorf("broken rewardstate %v -> %v", info.RewardState, *h.RewardState)
		}
	default:
		return fmt.Errorf("invalid rewardstate")
	}

	totalReward := info.TotalRewardAmount
	if *h.RewardState == types.BenefitState_BenefitDone {
		if h.RewardAmount != nil {
			totalReward = h.RewardAmount.Add(totalReward)
		}
	}

	if _, err := rewardcrud.UpdateSet(
		info.Update(),
		&rewardcrud.Req{
			RewardState:           h.RewardState,
			LastRewardAt:          h.RewardAt,
			RewardTID:             h.RewardTID,
			NextRewardStartAmount: h.NextRewardStartAmount,
			LastRewardAmount:      h.RewardAmount,
			TotalRewardAmount:     &totalReward,
			LastUnitRewardAmount:  h.UnitRewardAmount,
		},
	).Save(ctx); err != nil {
		return err
	}

	if *h.RewardState != types.BenefitState_BenefitDone {
		return nil
	}
	if h.RewardAt == nil {
		return fmt.Errorf("invalid rewardat")
	}

	stm, err := rewardhistorycrud.SetQueryConds(tx.GoodRewardHistory.Query(), &rewardhistorycrud.Conds{
		GoodID:     &cruder.Cond{Op: cruder.EQ, Val: *h.ID},
		RewardDate: &cruder.Cond{Op: cruder.EQ, Val: *h.RewardAt},
	})
	if err != nil {
		return err
	}
	exist, err := stm.Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("already exists")
	}

	if _, err := rewardhistorycrud.CreateSet(
		tx.GoodRewardHistory.Create(),
		&rewardhistorycrud.Req{
			GoodID:     h.ID,
			RewardDate: h.RewardAt,
			TID:        h.RewardTID,
			Amount:     h.RewardAmount,
			UnitAmount: h.UnitRewardAmount,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *updateHandler) updateStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Stock.
		Query().
		Where(
			entstock.GoodID(*h.ID),
			entstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	if h.Total == nil {
		return nil
	}

	sold := info.Total.Sub(info.SpotQuantity)
	spotQuantity := h.Total.Sub(sold)
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid total")
	}

	if _, err := stockcrud.UpdateSet(
		info.Update(),
		&stockcrud.Req{
			Total:        h.Total,
			SpotQuantity: &spotQuantity,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateGood(ctx context.Context, tx *ent.Tx) error {
	if _, err := goodcrud.UpdateSet(
		tx.Good.UpdateOneID(*h.ID),
		&goodcrud.Req{
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
			SupportCoinTypeIDs:   h.SupportCoinTypeIDs,
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

func (h *Handler) UpdateGood(ctx context.Context) (*npool.Good, error) {
	if h.Title != nil {
		key := h.lockKey()
		if err := redis2.TryLock(key, 0); err != nil {
			return nil, err
		}
		defer func() {
			_ = redis2.Unlock(key)
		}()

		h.Conds = &goodcrud.Conds{
			ID:    &cruder.Cond{Op: cruder.NEQ, Val: *h.ID},
			Title: &cruder.Cond{Op: cruder.EQ, Val: *h.Title},
		}
		exist, err := h.ExistGoodConds(ctx)
		if err != nil {
			return nil, err
		}
		if exist {
			return nil, fmt.Errorf("already exists")
		}
	}

	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateExtraInfo(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateReward(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateStock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateGood(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetGood(ctx)
}
