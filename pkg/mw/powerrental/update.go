package powerrental

import (
	"context"
	"fmt"
	"time"

	rewardcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward"
	rewardhistorycrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward/history"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	mininggoodstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock/mining"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	goodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/goodbase"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

type updateHandler struct {
	*powerRentalGoodQueryHandler
	stockValidator      *validateStockHandler
	sqlPowerRental      string
	sqlGoodBase         string
	sqlMiningGoodStocks []string
}

func (h *updateHandler) constructGoodBaseSql(ctx context.Context) error {
	handler, err := goodbase1.NewHandler(
		ctx,
		goodbase1.WithEntID(func() *string { s := h.GoodBaseReq.EntID.String(); return &s }(), false),
		goodbase1.WithGoodType(h.GoodBaseReq.GoodType, false),
		goodbase1.WithBenefitType(h.GoodBaseReq.BenefitType, false),
		goodbase1.WithName(h.GoodBaseReq.Name, false),
		goodbase1.WithServiceStartAt(h.GoodBaseReq.ServiceStartAt, false),
		goodbase1.WithStartMode(h.GoodBaseReq.StartMode, false),
		goodbase1.WithTestOnly(h.GoodBaseReq.TestOnly, false),
		goodbase1.WithBenefitIntervalHours(h.GoodBaseReq.BenefitIntervalHours, false),
		goodbase1.WithPurchasable(h.GoodBaseReq.Purchasable, false),
		goodbase1.WithOnline(h.GoodBaseReq.Online, false),
	)
	if err != nil {
		return err
	}
	h.sqlGoodBase, err = handler.ConstructUpdateSql()
	if err == cruder.ErrUpdateNothing {
		return nil
	}
	return err
}

func (h *updateHandler) constructPowerRentalSql() error {
	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update power_rentals "
	if h.DeviceTypeID != nil {
		_sql += fmt.Sprintf("%vdevice_type_id = '%v', ", set, *h.DeviceTypeID)
		set = ""
	}
	if h.VendorLocationID != nil {
		_sql += fmt.Sprintf("%vvendor_location_id = '%v', ", set, *h.VendorLocationID)
		set = ""
	}
	if h.UnitPrice != nil {
		_sql += fmt.Sprintf("%vunit_price = '%v', ", set, *h.UnitPrice)
		set = ""
	}
	if h.QuantityUnit != nil {
		_sql += fmt.Sprintf("%vquantity_unit = '%v', ", set, *h.QuantityUnit)
		set = ""
	}
	if h.QuantityUnitAmount != nil {
		_sql += fmt.Sprintf("%vquantity_unit_amount = '%v', ", set, *h.QuantityUnitAmount)
		set = ""
	}
	if h.DeliveryAt != nil {
		_sql += fmt.Sprintf("%vdelivery_at = %v, ", set, *h.DeliveryAt)
		set = ""
	}
	if h.DurationType != nil {
		_sql += fmt.Sprintf("%vduration_type = '%v', ", set, h.DurationType.String())
		set = ""
	}
	if h.StockMode != nil {
		_sql += fmt.Sprintf("%vstock_mode = '%v', ", set, h.StockMode.String())
		set = ""
	}
	if set != "" {
		return cruder.ErrUpdateNothing
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += fmt.Sprintf("where id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from ("
	_sql += "select * from power_rentals as pr "
	_sql += fmt.Sprintf("where pr.good_id = '%v'", *h.GoodID)
	_sql += " limit 1) as tmp)"
	if h.DeviceTypeID != nil {
		_sql += " and exists ("
		_sql += "select 1 from device_infos "
		_sql += fmt.Sprintf("where ent_id = '%v'", *h.DeviceTypeID)
		_sql += " limit 1) "
	}
	if h.VendorLocationID != nil {
		_sql += " and exists ("
		_sql += "select 1 from vendor_locations "
		_sql += fmt.Sprintf("where ent_id = '%v'", *h.VendorLocationID)
		_sql += " limit 1) "
	}
	h.sqlPowerRental = _sql
	return nil
}

func (h *updateHandler) execSql(ctx context.Context, tx *ent.Tx, sql string) error {
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updatePowerRental(ctx context.Context, tx *ent.Tx) error {
	return h.execSql(ctx, tx, h.sqlPowerRental)
}

func (h *updateHandler) updateGoodBase(ctx context.Context, tx *ent.Tx) error {
	if h.sqlGoodBase == "" {
		return nil
	}
	return h.execSql(ctx, tx, h.sqlGoodBase)
}

func (h *updateHandler) updateStock(ctx context.Context, tx *ent.Tx) error {
	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(h.stock.ID),
		h.StockReq,
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *updateHandler) updateMiningGoodStocks(ctx context.Context, tx *ent.Tx) error {
	for _, poolStock := range h.MiningGoodStockReqs {
		if _, err := mininggoodstockcrud.UpdateSet(
			tx.MiningGoodStock.UpdateOneID(*poolStock.ID),
			poolStock,
		).Save(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (h *updateHandler) _validateStock() error {
	if h.StockReq.Total == nil {
		h.StockReq.Total = &h.stock.Total
	}
	if h.StockMode == nil {
		h.StockMode = func() *types.GoodStockMode {
			e := types.GoodStockMode(types.GoodStockMode_value[h.powerRental.StockMode])
			return &e
		}()
	}

	miningGoodStockReqs := h.MiningGoodStockReqs
	defer func() { h.MiningGoodStockReqs = miningGoodStockReqs }()

	for _, poolStock := range h.miningGoodStocks {
		if func() bool {
			for _, _poolStock := range h.MiningGoodStockReqs {
				if *_poolStock.EntID == poolStock.EntID {
					_poolStock.ID = &poolStock.ID
					return true
				}
			}
			return false
		}() {
			continue
		}
		h.MiningGoodStockReqs = append(h.MiningGoodStockReqs, &mininggoodstockcrud.Req{
			EntID: &poolStock.EntID,
			Total: &poolStock.Total,
		})
	}

	return h.stockValidator.validateStock()
}

func (h *updateHandler) validateRewardState() error {
	if h.RewardReq.RewardState == nil {
		return nil
	}

	switch h.goodReward.RewardState {
	case types.BenefitState_BenefitWait.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitFail:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitTransferring.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitBookKeeping:
		case types.BenefitState_BenefitFail:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitUserBookKeeping:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitUserBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitSimulateBookKeeping:
		case types.BenefitState_BenefitDone:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitSimulateBookKeeping.String():
		switch *h.RewardReq.RewardState {
		case types.BenefitState_BenefitDone:
		default:
			return fmt.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	case types.BenefitState_BenefitDone.String():
		fallthrough //nolint
	case types.BenefitState_BenefitFail.String():
		if *h.RewardReq.RewardState != types.BenefitState_BenefitWait {
			return fmt.Errorf("broken rewardstate %v -> %v", h.goodReward.RewardState, *h.RewardReq.RewardState)
		}
	default:
		return fmt.Errorf("invalid rewardstate")
	}
	return nil
}

func (h *updateHandler) updateReward(ctx context.Context, tx *ent.Tx) error {
	if h.RewardReq.RewardState == nil {
		return nil
	}

	totalReward := h.goodReward.TotalRewardAmount
	if *h.RewardReq.RewardState == types.BenefitState_BenefitSimulateBookKeeping {
		if h.RewardReq.LastRewardAmount != nil {
			totalReward = h.RewardReq.LastRewardAmount.Add(totalReward)
		}
	}
	h.RewardReq.TotalRewardAmount = &totalReward

	if _, err := rewardcrud.UpdateSet(h.goodReward.Update(), h.RewardReq).Save(ctx); err != nil {
		return err
	}

	if *h.RewardReq.RewardState != types.BenefitState_BenefitSimulateBookKeeping {
		return nil
	}
	if h.RewardReq.LastRewardAt == nil {
		return fmt.Errorf("invalid rewardat")
	}

	stm, err := rewardhistorycrud.SetQueryConds(tx.GoodRewardHistory.Query(), &rewardhistorycrud.Conds{
		GoodID:     &cruder.Cond{Op: cruder.EQ, Val: *h.EntID},
		RewardDate: &cruder.Cond{Op: cruder.EQ, Val: *h.RewardReq.LastRewardAt},
	})
	if err != nil {
		return err
	}
	exist, err := stm.Exist(ctx)
	if err != nil {
		return err
	}
	if exist {
		return fmt.Errorf("already exists")
	}

	if _, err := rewardhistorycrud.CreateSet(
		tx.GoodRewardHistory.Create(),
		&rewardhistorycrud.Req{
			GoodID:     h.EntID,
			RewardDate: h.RewardReq.LastRewardAt,
			TID:        h.RewardReq.RewardTID,
			Amount:     h.RewardReq.LastRewardAmount,
			UnitAmount: h.RewardReq.LastUnitRewardAmount,
		},
	).Save(ctx); err != nil {
		return err
	}

	return nil
}

func (h *Handler) UpdatePowerRental(ctx context.Context) error {
	handler := &updateHandler{
		powerRentalGoodQueryHandler: &powerRentalGoodQueryHandler{
			Handler: h,
		},
		stockValidator: &validateStockHandler{
			Handler: h,
		},
	}

	if err := handler.requirePowerRentalGood(ctx); err != nil {
		return err
	}
	if err := handler._validateStock(); err != nil {
		return err
	}
	if err := handler.validateRewardState(); err != nil {
		return err
	}

	handler.constructPowerRentalSql()
	if err := handler.constructGoodBaseSql(ctx); err != nil {
		return err
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateGoodBase(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateStock(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateMiningGoodStocks(_ctx, tx); err != nil {
			return err
		}
		if err := handler.updateReward(_ctx, tx); err != nil {
			return err
		}
		return handler.updatePowerRental(_ctx, tx)
	})
}
