
package appstock

import (
	"context"
	"fmt"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type reserveHandler struct {
	*stockAppGoodQuery
}

func (h *reserveHandler) constructGoodStockSQL(table string, id uint32) (string, error) {
	sql := fmt.Sprintf(
		`update %v
		set
		  spot_quantity = spot_quantity - %v,
		  app_reserved = app_reserved + %v`,
		table,
		*h.Reserved,
		*h.Reserved,
	)
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  spot_quantity >= %v`,
		id,
		*h.Reserved,
	)
	sql += ` and
		in_service + wait_start + locked + app_reserved + spot_quantity = total`
	return sql, nil
}

func (h *reserveHandler) constructAppGoodStockSQL(table string, id uint32) (string, error) {
	sql := fmt.Sprintf(
		`update %v
		set
		  spot_quantity = spot_quantity + %v,
		  reserved = reserved + %v`,
		table,
		*h.Reserved,
		*h.Reserved,
	)
	sql += fmt.Sprintf(
		` where
		  id = %v
		and
		  deleted_at = 0
		and
		  spot_quantity <= reserved`,
		id,
	)
	return sql, nil
}

func (h *reserveHandler) reserveStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.stock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructGoodStockSQL("stocks_v1", stock.stock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *reserveHandler) reserveMiningGoodStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.miningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructGoodStockSQL("mining_good_stocks", stock.miningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *reserveHandler) reserveAppStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.appGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructAppGoodStockSQL("app_stocks", stock.appGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *reserveHandler) reserveAppMiningGoodStock(ctx context.Context, tx *ent.Tx) error {
	stock, ok := h.stocks[*h.AppGoodID]
	if !ok || stock.appMiningGoodStock == nil {
		return wlog.Errorf("invalid stock")
	}
	sql, err := h.constructAppGoodStockSQL("app_mining_good_stocks", stock.appMiningGoodStock.ID)
	if err != nil {
		return wlog.WrapError(err)
	}
	return h.execSQL(ctx, tx, sql)
}

func (h *Handler) ReserveStock(ctx context.Context) error {
	handler := &reserveHandler{
		stockAppGoodQuery: &stockAppGoodQuery{
			Handler: h,
		},
	}

	if err := handler.getStockAppGoods(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.reserveAppStock(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.reserveAppMiningGoodStock(ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		if err := handler.reserveStock(ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if handler.stockByMiningPool(*h.AppGoodID) {
			if err := handler.reserveMiningGoodStock(ctx, tx); err != nil {
				return wlog.WrapError(err)
			}
		}
		return nil
	})
}
