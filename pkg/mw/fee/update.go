package fee

import (
	"context"

	feecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/fee"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"
)

type updateHandler struct {
	*Handler
	fee      *ent.Fee
	goodBase *ent.GoodBase
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if _, err := goodbasecrud.UpdateSet(
		tx.GoodBase.UpdateOneID(h.goodBase.ID),
		&goodbasecrud.Req{
			GoodType: h.GoodBaseReq.GoodType,
			Name:     h.GoodBaseReq.Name,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := feecrud.UpdateSet(
		tx.Fee.UpdateOneID(*h.ID),
		&feecrud.Req{
			SettlementType: h.SettlementType,
			UnitValue:      h.UnitValue,
			DurationType:   h.DurationType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateFee(ctx context.Context) (*npool.Fee, error) {
	handler := &updateHandler{
		Handler: h,
	}

	var err error
	if err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.fee, err = cli.Fee.Get(_ctx, *h.ID); err != nil {
			return err
		}
		if handler.goodBase, err = cli.
			GoodBase.
			Query().
			Where(
				entgoodbase.EntID(handler.fee.GoodID),
			).Only(_ctx); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return nil, db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return err
		}
		return handler.updateFee(_ctx, tx)
	})
}
