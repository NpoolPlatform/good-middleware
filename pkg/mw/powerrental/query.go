package powerrental

import (
	"context"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodcoincrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin"
	mininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock/mining"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodcoin "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	entmininggoodstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	goodcoinmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
	stockmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*baseQueryHandler
	stmCount         *ent.GoodBaseSelect
	infos            []*npool.PowerRental
	miningGoodStocks []*stockmwpb.MiningGoodStock
	goodCoins        []*goodcoinmwpb.GoodCoinInfo
	total            uint32
}

func (h *queryHandler) queryJoin() {
	h.baseQueryHandler.queryJoin()
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinStock(s)
		h.queryJoinReward(s)
		if err := h.queryJoinPowerRental(s); err != nil {
			logger.Sugar().Errorw("queryJoinPowerRental", "Error", err)
		}
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) getMiningGoodStocks(ctx context.Context, cli *ent.Client) error {
	goodStockIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.GoodStockID))
		}
		return
	}()

	stm, err := mininggoodstockcrud.SetQueryConds(
		cli.MiningGoodStock.Query(),
		&mininggoodstockcrud.Conds{
			GoodStockIDs: &cruder.Cond{Op: cruder.IN, Val: goodStockIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entmininggoodstock.FieldEntID,
		entmininggoodstock.FieldGoodStockID,
		entmininggoodstock.FieldMiningPoolID,
		entmininggoodstock.FieldPoolGoodUserID,
		entmininggoodstock.FieldTotal,
		entmininggoodstock.FieldSpotQuantity,
		entmininggoodstock.FieldLocked,
		entmininggoodstock.FieldWaitStart,
		entmininggoodstock.FieldInService,
		entmininggoodstock.FieldSold,
	).Scan(ctx, &h.miningGoodStocks)
}

func (h *queryHandler) getGoodCoins(ctx context.Context, cli *ent.Client) error {
	goodIDs := func() (uids []uuid.UUID) {
		for _, info := range h.infos {
			uids = append(uids, uuid.MustParse(info.GoodID))
		}
		return
	}()

	stm, err := goodcoincrud.SetQueryConds(
		cli.GoodCoin.Query(),
		&goodcoincrud.Conds{
			GoodIDs: &cruder.Cond{Op: cruder.IN, Val: goodIDs},
		},
	)
	if err != nil {
		return wlog.WrapError(err)
	}

	return stm.Select(
		entgoodcoin.FieldGoodID,
		entgoodcoin.FieldCoinTypeID,
		entgoodcoin.FieldMain,
		entgoodcoin.FieldIndex,
	).Scan(ctx, &h.goodCoins)
}

func (h *queryHandler) formalize() {
	goodMiningStocks := map[string][]*stockmwpb.MiningGoodStock{}
	goodCoins := map[string][]*goodcoinmwpb.GoodCoinInfo{}

	for _, stock := range h.miningGoodStocks {
		stock.Total = func() string { amount, _ := decimal.NewFromString(stock.Total); return amount.String() }()
		stock.SpotQuantity = func() string { amount, _ := decimal.NewFromString(stock.SpotQuantity); return amount.String() }()
		stock.Locked = func() string { amount, _ := decimal.NewFromString(stock.Locked); return amount.String() }()
		stock.WaitStart = func() string { amount, _ := decimal.NewFromString(stock.WaitStart); return amount.String() }()
		stock.InService = func() string { amount, _ := decimal.NewFromString(stock.InService); return amount.String() }()
		stock.Sold = func() string { amount, _ := decimal.NewFromString(stock.Sold); return amount.String() }()
		goodMiningStocks[stock.GoodStockID] = append(goodMiningStocks[stock.GoodStockID], stock)
	}
	for _, goodCoin := range h.goodCoins {
		goodCoins[goodCoin.GoodID] = append(goodCoins[goodCoin.GoodID], goodCoin)
	}
	for _, info := range h.infos {
		info.UnitPrice = func() string { amount, _ := decimal.NewFromString(info.UnitPrice); return amount.String() }()
		info.QuantityUnitAmount = func() string { amount, _ := decimal.NewFromString(info.QuantityUnitAmount); return amount.String() }()
		info.UnitLockDeposit = func() string { amount, _ := decimal.NewFromString(info.UnitLockDeposit); return amount.String() }()
		info.GoodTotal = func() string { amount, _ := decimal.NewFromString(info.GoodTotal); return amount.String() }()
		info.GoodSpotQuantity = func() string { amount, _ := decimal.NewFromString(info.GoodSpotQuantity); return amount.String() }()
		info.GoodLocked = func() string { amount, _ := decimal.NewFromString(info.GoodLocked); return amount.String() }()
		info.GoodInService = func() string { amount, _ := decimal.NewFromString(info.GoodInService); return amount.String() }()
		info.GoodWaitStart = func() string { amount, _ := decimal.NewFromString(info.GoodWaitStart); return amount.String() }()
		info.GoodSold = func() string { amount, _ := decimal.NewFromString(info.GoodSold); return amount.String() }()
		info.GoodAppReserved = func() string { amount, _ := decimal.NewFromString(info.GoodAppReserved); return amount.String() }()
		info.LastRewardAmount = func() string { amount, _ := decimal.NewFromString(info.LastRewardAmount); return amount.String() }()
		info.NextRewardStartAmount = func() string { amount, _ := decimal.NewFromString(info.NextRewardStartAmount); return amount.String() }()
		info.LastUnitRewardAmount = func() string { amount, _ := decimal.NewFromString(info.LastUnitRewardAmount); return amount.String() }()
		info.TotalRewardAmount = func() string { amount, _ := decimal.NewFromString(info.TotalRewardAmount); return amount.String() }()
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
		info.BenefitType = types.BenefitType(types.BenefitType_value[info.BenefitTypeStr])
		info.DurationDisplayType = types.GoodDurationType(types.GoodDurationType_value[info.DurationDisplayTypeStr])
		info.StartMode = types.GoodStartMode(types.GoodStartMode_value[info.StartModeStr])
		info.StockMode = types.GoodStockMode(types.GoodStockMode_value[info.StockModeStr])
		info.RewardState = types.BenefitState(types.BenefitState_value[info.RewardStateStr])
		info.MiningGoodStocks = goodMiningStocks[info.GoodStockID]
		info.GoodCoins = goodCoins[info.GoodID]
	}
}

func (h *Handler) GetPowerRental(ctx context.Context) (*npool.PowerRental, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodBase(cli); err != nil {
			return wlog.WrapError(err)
		}
		handler.queryJoin()
		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getGoodCoins(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.getMiningGoodStocks(_ctx, cli)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, wlog.Errorf("too many records")
	}

	handler.formalize()
	return handler.infos[0], nil
}

func (h *Handler) GetPowerRentals(ctx context.Context) ([]*npool.PowerRental, uint32, error) {
	handler := &queryHandler{
		baseQueryHandler: &baseQueryHandler{
			Handler: h,
		},
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryGoodBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.stmCount, err = handler.queryGoodBases(cli)
		if err != nil {
			return wlog.WrapError(err)
		}

		handler.queryJoin()
		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return wlog.WrapError(err)
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))

		if err := handler.scan(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.getGoodCoins(_ctx, cli); err != nil {
			return wlog.WrapError(err)
		}
		return handler.getMiningGoodStocks(_ctx, cli)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
