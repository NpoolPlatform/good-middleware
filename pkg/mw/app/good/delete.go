package appgood

import (
	"context"
	"time"

	appgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good"
	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good"
	"github.com/google/uuid"
)

type deleteHandler struct {
	*Handler
}

func (h *deleteHandler) deleteStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppStock.
		Query().
		Where(
			entappstock.AppGoodID(*h.EntID),
			entappstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}

	now := uint32(time.Now().Unix())
	if _, err := appstockcrud.UpdateSet(
		info.Update(),
		&appstockcrud.Req{
			DeletedAt: &now,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *deleteHandler) deleteAppGood(ctx context.Context, tx *ent.Tx) error {
	now := uint32(time.Now().Unix())
	if _, err := appgoodcrud.UpdateSet(
		tx.AppGood.UpdateOneID(*h.ID),
		&appgoodcrud.Req{
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
		if err := handler.deleteStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.deleteAppGood(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
