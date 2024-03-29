package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"

	"github.com/shopspring/decimal"
)

type reserveHandler struct {
	*stockAppGoodQuery
}

func (h *reserveHandler) reserveStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := stock.stock.SpotQuantity
	appReserved := stock.stock.AppReserved

	appReserved = h.Reserved.Add(appReserved)
	spotQuantity = spotQuantity.Sub(*h.Reserved)

	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(stock.stock.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if appReserved.Add(stock.stock.Locked).
		Add(stock.stock.WaitStart).
		Add(stock.stock.InService).
		Add(spotQuantity).
		Cmp(stock.stock.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(stock.stock.ID),
		&stockcrud.Req{
			SpotQuantity: &spotQuantity,
			AppReserved:  &appReserved,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *reserveHandler) reserveAppStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}

	spotQuantity := stock.appGoodStock.SpotQuantity
	reserved := stock.appGoodStock.Reserved
	spotQuantity = h.Reserved.Add(spotQuantity)
	reserved = h.Reserved.Add(reserved)
	if spotQuantity.Cmp(reserved) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(stock.appGoodStock.ID),
		&appstockcrud.Req{
			SpotQuantity: &spotQuantity,
			Reserved:     &reserved,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) ReserveStock(ctx context.Context) error {
	handler := &reserveHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
	}

	if err := handler.getStockAppGoods(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.reserveAppStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.reserveStock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
}
