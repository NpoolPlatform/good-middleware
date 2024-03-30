package appstock

import (
	"context"
	"fmt"

	appgoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	appmininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock/mining"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	powerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/powerrental"

	"github.com/google/uuid"
)

type createHandler struct {
	*stockAppGoodQuery
}

func (h *createHandler) _createStock(ctx context.Context, tx *ent.Tx) error {
	if _, err := appgoodstockcrud.CreateSet(
		tx.AppStock.Create(),
		&h.Req,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *createHandler) createAppGoodMiningStocks(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}
	handler, err := powerrental1.NewHandler(
		ctx,
		powerrental1.WithGoodID(func() *string { s := stock.powerRental.GoodID.String(); return &s }(), true),
	)
	if err != nil {
		return err
	}
	powerRental, err := handler.GetPowerRental(ctx)
	if err != nil {
		return err
	}
	for _, stock := range powerRental.MiningGoodStocks {
		if _, err := appmininggoodstockcrud.CreateSet(
			tx.AppMiningGoodStock.Create(),
			&appmininggoodstockcrud.Req{
				AppGoodStockID:    h.EntID,
				MiningGoodStockID: func() *uuid.UUID { uid := uuid.MustParse(stock.EntID); return &uid }(),
			},
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Only for test. Stock should always be created with good
func (h *Handler) createStock(ctx context.Context) error {
	handler := &createHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
	}
	if err := handler.getStockAppGoods(ctx); err != nil {
		return err
	}
	return db.WithDebugTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.createAppGoodMiningStocks(_ctx, tx); err != nil {
				return err
			}
		}
		return handler._createStock(ctx, tx)
	})
}
