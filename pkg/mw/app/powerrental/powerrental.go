package powerrental

import (
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/shopspring/decimal"
)

type PowerRental interface {
	MinOrderAmount() decimal.Decimal
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
