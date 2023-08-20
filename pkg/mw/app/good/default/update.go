package appdefaultgood

import (
	"context"

	appdefaultgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/default"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

type updateHandler struct {
	*Handler
}

func (h *updateHandler) updateDefault(ctx context.Context, tx *ent.Tx) error {
	if _, err := appdefaultgoodcrud.UpdateSet(
		tx.AppDefaultGood.UpdateOneID(*h.ID),
		&appdefaultgoodcrud.Req{
			GoodID:    h.GoodID,
			AppGoodID: h.AppGoodID,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDefault(ctx context.Context) (*npool.Default, error) {
	handler := &updateHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateDefault(_ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetDefault(ctx)
}
