package fee

import (
	"context"

	feecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/fee"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	"github.com/google/uuid"
)

type createHandler struct {
	*Handler
}

func (h *createHandler) createGoodBase(ctx context.Context, tx *ent.Tx) error {
	valueTrue := true
	if _, err := goodbasecrud.CreateSet(
		tx.GoodBase.Create(),
		&goodbasecrud.Req{
			EntID:       h.GoodID,
			GoodType:    h.GoodBaseReq.GoodType,
			Name:        h.GoodBaseReq.Name,
			Purchasable: &valueTrue,
			Online:      &valueTrue,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createFee(ctx context.Context, tx *ent.Tx) error {
	if _, err := feecrud.CreateSet(
		tx.Fee.Create(),
		&feecrud.Req{
			EntID:          h.EntID,
			GoodID:         h.GoodID,
			SettlementType: h.SettlementType,
			UnitValue:      h.UnitValue,
			DurationType:   h.DurationType,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) CreateFee(ctx context.Context) error {
	entID := uuid.New()
	if h.EntID == nil {
		h.EntID = &entID
	}
	goodID := uuid.New()
	if h.GoodID == nil {
		h.GoodID = &goodID
	}

	handler := &createHandler{
		Handler: h,
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.createGoodBase(ctx, tx); err != nil {
			return err
		}
		return handler.createFee(ctx, tx)
	})
}
