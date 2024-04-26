package powerrental

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"
)

type PowerRental interface {
	MinOrderAmount() decimal.Decimal
	MaxOrderAmount() decimal.Decimal
	MinOrderDurationSeconds() uint32
	MaxOrderDurationSeconds() uint32
}

type powerRental struct {
	powerRental         *ent.PowerRental
	goodBase            *ent.GoodBase
	appGoodBase         *ent.AppGoodBase
	appPowerRental      *ent.AppPowerRental
	stock               *ent.Stock
	miningGoodStocks    []*ent.MiningGoodStock
	appGoodStock        *ent.AppStock
	appMiningGoodStocks []*ent.AppMiningGoodStock
}

func (pr *powerRental) MinOrderAmount() decimal.Decimal {
	return pr.appPowerRental.MinOrderAmount
}

func (pr *powerRental) MaxOrderAmount() decimal.Decimal {
	return pr.appPowerRental.MaxOrderAmount
}

func (pr *powerRental) MinOrderDurationSeconds() uint32 {
	return pr.appPowerRental.MinOrderDurationSeconds
}

func (pr *powerRental) MaxOrderDurationSeconds() uint32 {
	return pr.appPowerRental.MaxOrderDurationSeconds
}
