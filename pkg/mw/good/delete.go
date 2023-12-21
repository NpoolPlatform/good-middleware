package good

import (
	"context"
	"time"

	goodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good"
	extrainfocrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/extrainfo"
	rewardcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entextrainfo "github.com/NpoolPlatform/good-middleware/pkg/db/ent/extrainfo"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

//nolint:dupl
func (h *deleteHandler) deleteExtraInfo(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		ExtraInfo.
		Query().
		Where(
			entextrainfo.GoodID(*h.EntID),
			entextrainfo.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := extrainfocrud.UpdateSet(
		info.Update(),
		&extrainfocrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *deleteHandler) deleteReward(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		GoodReward.
		Query().
		Where(
			entgoodreward.GoodID(*h.EntID),
			entgoodreward.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := rewardcrud.UpdateSet(
		info.Update(),
		&rewardcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:dupl
func (h *deleteHandler) deleteStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Stock.
		Query().
		Where(
			entstock.GoodID(*h.EntID),
			entstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := stockcrud.UpdateSet(
		info.Update(),
		&stockcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteGood(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := goodcrud.UpdateSet(
		tx.Good.UpdateOneID(*h.ID),
		&goodcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) DeleteGood(ctx context.Context) (*npool.Good, error) {
	info, err := h.GetGood(ctx)
	if err != nil {
		return nil, err
	}
	if info == nil {
		return nil, nil
	}
	entID := uuid.MustParse(info.EntID)
	h.EntID = &entID

	handler := &deleteHandler{
		Handler: h,
	}

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.deleteExtraInfo(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteReward(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteStock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteGood(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
