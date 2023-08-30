//nolint:dupl
package appstock

import (
	"context"
	"fmt"

	appstockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/stock"
	stockcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/stock"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	entstock "github.com/NpoolPlatform/good-middleware/pkg/db/ent/stock"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/shopspring/decimal"
)

type subHandler struct {
	*Handler
	rollbackAppSpotQuantity decimal.Decimal
}

//nolint:gocyclo
func (h *subHandler) subStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		Stock.
		Query().
		Where(
			entstock.GoodID(*h.GoodID),
			entstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("stock not found")
	}

	spotQuantity := info.SpotQuantity
	locked := info.Locked
	inService := info.InService
	sold := info.Sold
	waitStart := info.WaitStart
	appReserved := info.AppReserved

	if h.Locked != nil {
		locked = locked.Sub(*h.Locked)
		if h.Rollback != nil && *h.Rollback {
			spotQuantity = h.Locked.Add(spotQuantity).Sub(h.rollbackAppSpotQuantity)
			appReserved = appReserved.Add(h.rollbackAppSpotQuantity)
		}
	}
	if h.WaitStart != nil {
		waitStart = waitStart.Sub(*h.WaitStart)
		sold = sold.Sub(*h.WaitStart)
		if h.Rollback != nil && *h.Rollback {
			locked = h.WaitStart.Add(locked)
		} else {
			spotQuantity = h.WaitStart.Add(spotQuantity)
		}
	}
	if h.InService != nil {
		inService = inService.Sub(*h.InService)
		if h.ChargeBack != nil && *h.ChargeBack {
			sold = sold.Sub(*h.InService)
		}
		if h.Rollback != nil && *h.Rollback {
			waitStart = h.InService.Add(waitStart)
		} else {
			spotQuantity = h.InService.Add(spotQuantity)
		}
	}
	if h.Reserved != nil {
		appReserved = appReserved.Sub(*h.Reserved)
		spotQuantity = h.Reserved.Add(spotQuantity)
	}

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid waitstart")
	}
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid inservice")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid sold")
	}
	if appReserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid appreserved")
	}

	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid stock")
	}
	if spotQuantity.Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if locked.Add(inService).
		Add(waitStart).
		Add(appReserved).
		Add(spotQuantity).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("invalid stock")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(info.ID),
		&stockcrud.Req{
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			InService:    &inService,
			WaitStart:    &waitStart,
			AppReserved:  &appReserved,
			Sold:         &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

//nolint:gocyclo
func (h *subHandler) subAppStock(ctx context.Context, tx *ent.Tx) error {
	info, err := tx.
		AppStock.
		Query().
		Where(
			entappstock.ID(*h.ID),
			entappstock.DeletedAt(0),
		).
		ForUpdate().
		Only(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("stock not found")
	}

	h.GoodID = &info.GoodID
	locked := info.Locked
	inService := info.InService
	sold := info.Sold
	waitStart := info.WaitStart
	reserved := info.Reserved
	spotQuantity := info.SpotQuantity

	if h.Locked != nil {
		locked = locked.Sub(*h.Locked)
		if h.Rollback != nil && *h.Rollback {
			if spotQuantity.Cmp(decimal.NewFromInt(0)) == 0 {
				if reserved.Cmp(inService.Add(waitStart)) > 0 {
					spotQuantity = reserved.Sub(inService.Add(waitStart))
					if h.Locked.Cmp(spotQuantity) < 0 {
						spotQuantity = *h.Locked
					} else {
						h.rollbackAppSpotQuantity = h.Locked.Sub(spotQuantity)
					}
				}
			} else {
				spotQuantity = h.Locked.Add(spotQuantity)
			}
		}
	}
	if h.WaitStart != nil {
		waitStart = waitStart.Sub(*h.WaitStart)
		sold = sold.Sub(*h.WaitStart)
		if h.Rollback != nil && *h.Rollback {
			locked = h.WaitStart.Add(locked)
		}
	}
	if h.InService != nil {
		inService = inService.Sub(*h.InService)
		if h.ChargeBack != nil && *h.ChargeBack {
			sold = sold.Sub(*h.InService)
		}
		if h.Rollback != nil && *h.Rollback {
			waitStart = h.InService.Add(waitStart)
		}
	}
	if h.Reserved != nil {
		if h.Reserved.Cmp(spotQuantity) > 0 {
			return fmt.Errorf("invalid reserved")
		}
		reserved = reserved.Sub(*h.Reserved)
		spotQuantity = spotQuantity.Sub(*h.Reserved)
	}

	if locked.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid locked")
	}
	if waitStart.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid waitstart")
	}
	if inService.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid inservice")
	}
	if sold.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid sold")
	}
	if spotQuantity.Cmp(info.Reserved) > 0 {
		return fmt.Errorf("invalid spotquantity")
	}
	if spotQuantity.Cmp(decimal.NewFromInt(0)) < 0 {
		spotQuantity = decimal.NewFromInt(0)
	}
	if reserved.Cmp(decimal.NewFromInt(0)) < 0 {
		return fmt.Errorf("invalid reserved")
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(info.ID),
		&appstockcrud.Req{
			Reserved:     &reserved,
			SpotQuantity: &spotQuantity,
			Locked:       &locked,
			InService:    &inService,
			WaitStart:    &waitStart,
			Sold:         &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *Handler) SubStock(ctx context.Context) (*npool.Stock, error) {
	if err := h.validate(); err != nil {
		return nil, err
	}

	handler := &subHandler{
		Handler: h,
	}

	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.subAppStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.subStock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
