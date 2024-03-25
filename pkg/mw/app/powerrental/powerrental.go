package powerrental

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"
)

type PowerRental interface {
	MinOrderAmount() decimal.Decimal
	MaxOrderAmount() decimal.Decimal
	MinOrderDuration() uint32
	MaxOrderDuration() uint32
}

type powerRental struct {
	powerRental    *ent.PowerRental
	goodBase       *ent.GoodBase
	appGoodBase    *ent.AppGoodBase
	appPowerRental *ent.AppPowerRental
}

func (pr *powerRental) MinOrderAmount() decimal.Decimal {
	return pr.appPowerRental.MinOrderAmount
}

func (pr *powerRental) MaxOrderAmount() decimal.Decimal {
	return pr.appPowerRental.MaxOrderAmount
}

func (pr *powerRental) MinOrderDuration() uint32 {
	return pr.appPowerRental.MinOrderDuration
}

func (pr *powerRental) MaxOrderDuration() uint32 {
	return pr.appPowerRental.MaxOrderDuration
}
