package powerrental

import (
	"fmt"

	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/shopspring/decimal"
)

type validateStockHandler struct {
	*Handler
}

func (h *validateStockHandler) validateStock() error {
	if *h.StockMode != types.GoodStockMode_GoodStockByMiningPool {
		return nil
	}
	poolTotal := decimal.NewFromInt(0)
	for _, poolStock := range h.MiningGoodStockReqs {
		poolTotal = poolStock.Total.Add(poolTotal)
	}
	if h.StockReq.Total.Cmp(poolTotal) != 0 {
		return fmt.Errorf("invalid stock total")
	}
	return nil
}
