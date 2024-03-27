package powerrental

import (
	"fmt"

	"github.com/shopspring/decimal"
)

type validateStockHandler struct {
	*Handler
}

func (h *validateStockHandler) validateStock() error {
	poolTotal := decimal.NewFromInt(0)
	for _, poolStock := range h.MiningGoodStockReqs {
		poolTotal = poolStock.Total.Add(poolTotal)
	}
	if h.StockReq.Total.Cmp(poolTotal) < 0 {
		return fmt.Errorf("invalid stock total")
	}
	return nil
}
