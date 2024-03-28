package appstock

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
)

type stockAppGoodQuery struct {
	*Handler
	appGoodBase  *ent.AppGoodBase
	appGoodStock *ent.AppStock
	stock        *ent.Stock
}

func (h *stockAppGoodQuery) getAppGoodStock(ctx context.Context, cli *ent.Client, must bool) (err error) {
	stm := cli.AppStock.Query().Where(entappgoodstock.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgoodstock.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgoodstock.EntID(*h.EntID))
	}
	if h.AppGoodID != nil {
		stm.Where(entappgoodstock.AppGoodID(*h.AppGoodID))
	}
	if h.appGoodStock, err = stm.Only(ctx); err != nil {
		return err
	}
	return nil
}

func (h *stockAppGoodQuery) getAppGoodBase(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.appGoodBase, err = cli.
		AppGoodBase.
		Query().
		Where(
			entappgoodbase.EntID(h.appGoodStock.AppGoodID),
			entappgoodbase.DeletedAt(0),
		).
		Only(ctx); err != nil {
		return err
	}
	return nil
}

func (h *stockAppGoodQuery) getGoodStock(ctx context.Context, cli *ent.Client, must bool) (err error) {
	if h.stock, err = cli.
		Stock.
		Query().
		Where(
			entstock.GoodID(h.appGoodBase.GoodID),
			entstock.DeletedAt(0),
		).
		Only(ctx); err != nil {
		return err
	}
	return nil
}

func (h *stockAppGoodQuery) _getStockAppGood(ctx context.Context, must bool) error {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return fmt.Errorf("invalid stockid")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getAppGoodStock(_ctx, cli, must); err != nil {
			return err
		}
		if h.appGoodStock == nil {
			return nil
		}
		if err := h.getAppGoodBase(_ctx, cli, must); err != nil {
			return err
		}
		if h.appGoodBase == nil {
			return nil
		}
		return h.getGoodStock(_ctx, cli, must)
	})
}

func (h *stockAppGoodQuery) getStockAppGood(ctx context.Context) error {
	return h._getStockAppGood(ctx, false)
}

func (h *stockAppGoodQuery) requireStockAppGood(ctx context.Context) error {
	return h._getStockAppGood(ctx, true)
}
