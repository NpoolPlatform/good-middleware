package appstock

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappmininggoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appmininggoodstock"
	entappgoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entmininggoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	entpowerrental "github.com/NpoolPlatform/good-middleware/pkg/db/ent/powerrental"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type stockAppGood struct {
	goodBase           *ent.GoodBase
	powerRental        *ent.PowerRental
	appGoodBase        *ent.AppGoodBase
	appGoodStock       *ent.AppStock
	appMiningGoodStock *ent.AppMiningGoodStock
	stock              *ent.Stock
	miningGoodStock    *ent.MiningGoodStock
}

type stockAppGoodQuery struct {
	*Handler
	stocks                   map[uuid.UUID]*stockAppGood
	appGoodIDs               []uuid.UUID
	appGoodStockEntIDs       []uuid.UUID
	appMiningGoodStockEntIDs []uuid.UUID
	miningGoodStockEntIDs    []uuid.UUID
	stockGoodIDs             []uuid.UUID
}

func (h *stockAppGoodQuery) _getAppGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	appGoodStocks, err := cli.
		AppStock.
		Query().
		Where(
			entappgoodstock.EntIDIn(h.appGoodStockEntIDs...),
			entappgoodstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	for _, stock := range appGoodStocks {
		_stock, ok := h.stocks[stock.AppGoodID]
		if !ok {
			return fmt.Errorf("invalid stock")
		}
		_stock.appGoodStock = stock
	}
	return nil
}

func (h *stockAppGoodQuery) getAppMiningGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	appMiningGoodStocks, err := cli.
		AppMiningGoodStock.
		Query().
		Where(
			entappmininggoodstock.EntIDIn(h.appMiningGoodStockEntIDs...),
			entappmininggoodstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	for _, stock := range appMiningGoodStocks {
		_stock, err := func() (*stockAppGood, error) {
			for _, _stock := range h.Stocks {
				if stock.EntID != *_stock.EntID {
					continue
				}
				__stock, ok := h.stocks[*_stock.AppGoodID]
				if !ok {
					return nil, fmt.Errorf("invalid stock")
				}
				if __stock.powerRental == nil || __stock.powerRental.StockMode != types.GoodStockMode_GoodStockByUnique.String() {
					continue
				}
				return __stock, nil
			}
			return nil, fmt.Errorf("invalid stock")
		}()
		if err != nil {
			return err
		}
		_stock.appMiningGoodStock = stock
		h.appGoodStockEntIDs = append(h.appGoodStockEntIDs, stock.AppGoodStockID)
		h.miningGoodStockEntIDs = append(h.miningGoodStockEntIDs, stock.MiningGoodStockID)
	}
	return nil
}

func (h *stockAppGoodQuery) getAppGoodBases(ctx context.Context, cli *ent.Client) (err error) {
	appGoodBases, err := cli.
		AppGoodBase.
		Query().
		Where(
			entappgoodbase.EntIDIn(h.appGoodIDs...),
			entappgoodbase.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	for _, appGoodBase := range appGoodBases {
		h.stocks[appGoodBase.EntID].appGoodBase = appGoodBase
		h.stockGoodIDs = append(h.stockGoodIDs, appGoodBase.GoodID)
	}
	return nil
}

func (h *stockAppGoodQuery) getMiningGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	stocks, err := cli.
		MiningGoodStock.
		Query().
		Where(
			entmininggoodstock.EntIDIn(h.miningGoodStockEntIDs...),
			entmininggoodstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	for _, stock := range h.stocks {
		for _, _stock := range stocks {
			if _stock.EntID == stock.appGoodStock.EntID {
				stock.miningGoodStock = _stock
				break
			}
		}
	}
	return nil
}

func (h *stockAppGoodQuery) getGoodStocks(ctx context.Context, cli *ent.Client) (err error) {
	stocks, err := cli.
		Stock.
		Query().
		Where(
			entstock.GoodIDIn(h.stockGoodIDs...),
			entstock.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	for _, stock := range stocks {
		for _, _stock := range h.stocks {
			if _stock.goodBase != nil && _stock.goodBase.EntID == stock.GoodID {
				_stock.stock = stock
				break
			}
		}
	}
	return nil
}

func (h *stockAppGoodQuery) getGoodBases(ctx context.Context, cli *ent.Client) (err error) {
	goodIDs := func() (uids []uuid.UUID) {
		for _, stock := range h.stocks {
			uids = append(uids, stock.appGoodBase.GoodID)
		}
		return
	}()

	goodBases, err := cli.
		GoodBase.
		Query().
		Where(
			entgoodbase.EntIDIn(goodIDs...),
			entgoodbase.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	for _, stock := range h.stocks {
		for _, goodBase := range goodBases {
			if stock.appGoodBase != nil && stock.appGoodBase.GoodID == goodBase.EntID {
				stock.goodBase = goodBase
				break
			}
		}
	}
	return nil
}

func (h *stockAppGoodQuery) getPowerRentals(ctx context.Context, cli *ent.Client) (err error) {
	powerRentalGoodIDs := func() (uids []uuid.UUID) {
		for _, stock := range h.stocks {
			if stock.goodBase == nil {
				continue
			}
			switch stock.goodBase.GoodType {
			case types.GoodType_PowerRental.String():
			case types.GoodType_LegacyPowerRental.String():
			default:
				continue
			}
			uids = append(uids, stock.goodBase.EntID)
		}
		return
	}()

	powerRentals, err := cli.
		PowerRental.
		Query().
		Where(
			entpowerrental.GoodIDIn(powerRentalGoodIDs...),
			entpowerrental.DeletedAt(0),
		).
		All(ctx)
	if err != nil {
		return err
	}
	for _, stock := range h.stocks {
		for _, powerRental := range powerRentals {
			if stock.goodBase != nil && stock.goodBase.EntID == powerRental.GoodID {
				stock.powerRental = powerRental
				break
			}
		}
	}
	return nil
}

func (h *stockAppGoodQuery) getAppGoods(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := h.getAppGoodBases(_ctx, cli); err != nil {
			return err
		}
		if err := h.getGoodBases(_ctx, cli); err != nil {
			return err
		}
		return h.getPowerRentals(_ctx, cli)
	})
}

func (h *stockAppGoodQuery) formalizeAppGoodIDs() error {
	if h.AppGoodID != nil {
		h.appGoodIDs = append(h.appGoodIDs, *h.AppGoodID)
	}
	for _, stock := range h.Stocks {
		h.appGoodIDs = append(h.appGoodIDs, *stock.AppGoodID)
	}
	if len(h.appGoodIDs) == 0 {
		return fmt.Errorf("invalid appgoodids")
	}
	for _, appGoodID := range h.appGoodIDs {
		h.stocks[appGoodID] = &stockAppGood{}
	}
	return nil
}

func (h *stockAppGoodQuery) formalizeStockEntID(appGoodID uuid.UUID, entID uuid.UUID) error {
	_stock, ok := h.stocks[appGoodID]
	if !ok {
		return fmt.Errorf("invalid stock")
	}
	if _stock.powerRental == nil || _stock.powerRental.StockMode != types.GoodStockMode_GoodStockByMiningPool.String() {
		h.appGoodStockEntIDs = append(h.appGoodStockEntIDs, entID)
		return nil
	}
	h.appMiningGoodStockEntIDs = append(h.appMiningGoodStockEntIDs, entID)
	return nil
}

func (h *stockAppGoodQuery) formalizeStockEntIDs() error {
	if h.AppGoodID != nil && h.EntID != nil {
		if err := h.formalizeStockEntID(*h.AppGoodID, *h.EntID); err != nil {
			return err
		}
	}
	for _, stock := range h.Stocks {
		if err := h.formalizeStockEntID(*stock.AppGoodID, *stock.EntID); err != nil {
			return err
		}
	}
	fmt.Printf("stock ent ids: %v\n", h.appGoodStockEntIDs)
	fmt.Printf("mining stock ent ids: %v\n", h.appMiningGoodStockEntIDs)
	if len(h.appGoodStockEntIDs) == 0 && len(h.appMiningGoodStockEntIDs) == 0 {
		return fmt.Errorf("invalid stock")
	}
	return nil
}

func (h *stockAppGoodQuery) getAppGoodStocks(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if len(h.appMiningGoodStockEntIDs) > 0 {
			if err := h.getAppMiningGoodStocks(_ctx, cli); err != nil {
				return err
			}
		}
		if len(h.appGoodStockEntIDs) > 0 {
			if err := h._getAppGoodStocks(_ctx, cli); err != nil {
				return err
			}
		}
		if err := h.getGoodStocks(_ctx, cli); err != nil {
			return err
		}
		if len(h.miningGoodStockEntIDs) > 0 {
			return h.getMiningGoodStocks(_ctx, cli)
		}
		return nil
	})
}

func (h *stockAppGoodQuery) getStockAppGoods(ctx context.Context) error {
	h.stocks = map[uuid.UUID]*stockAppGood{}
	if err := h.formalizeAppGoodIDs(); err != nil {
		return err
	}
	if err := h.getAppGoods(ctx); err != nil {
		return err
	}
	if err := h.formalizeStockEntIDs(); err != nil {
		return err
	}
	return h.getAppGoodStocks(ctx)
}

func (h *stockAppGoodQuery) stockByMiningPool(appGoodID uuid.UUID) bool {
	stock, ok := h.stocks[appGoodID]
	if !ok {
		return false
	}
	return stock.powerRental != nil && stock.powerRental.StockMode == types.GoodStockMode_GoodStockByMiningPool.String()
}
