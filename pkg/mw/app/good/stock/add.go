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

type addHandler struct {
	*Handler
}

func (h *addHandler) addStock(ctx context.Context, tx *ent.Tx) error {
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
		return fmt.Errorf("invalid stock")
	}

	locked := info.Locked
	if h.Locked != nil {
		locked = h.Locked.Add(locked)
	}
	inService := info.InService
	sold := info.Sold
	if h.InService != nil {
		inService = h.InService.Add(inService)
		sold = h.InService.Add(sold)
	}
	waitStart := info.WaitStart
	if h.WaitStart != nil {
		waitStart = h.WaitStart.Add(waitStart)
		sold = h.WaitStart.Add(sold)
	}

	if locked.Add(inService).
		Add(waitStart).
		Add(info.AppLocked).
		Cmp(info.Total) > 0 {
		return fmt.Errorf("stock exhausted")
	}

	if _, err := stockcrud.UpdateSet(
		tx.Stock.UpdateOneID(info.ID),
		&stockcrud.Req{
			Locked:    &locked,
			InService: &inService,
			WaitStart: &waitStart,
			Sold:      &sold,
		},
	).Save(ctx); err != nil {
		return err
	}
	return nil
}

func (h *addHandler) addAppStock(ctx context.Context, tx *ent.Tx) error {
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
		return fmt.Errorf("invalid app stock")
	}
	h.GoodID = &info.GoodID
	spotQuantity := info.SpotQuantity

	locked := info.Locked
	if h.Locked != nil {
		locked = h.Locked.Add(locked)
		spotQuantity = spotQuantity.Sub(*h.Locked)
	}
	inService := info.InService
	sold := info.Sold

	if h.InService != nil {
		inService = h.InService.Add(inService)
		sold = h.InService.Add(sold)
		spotQuantity = spotQuantity.Sub(*h.InService)
	}
	waitStart := info.WaitStart
	if h.WaitStart != nil {
		waitStart = h.WaitStart.Add(waitStart)
		sold = h.WaitStart.Add(sold)
		spotQuantity = spotQuantity.Sub(*h.WaitStart)
	}

	if locked.Add(inService).
		Add(waitStart).
		Cmp(info.SpotQuantity) > 0 {
		spotQuantity = decimal.NewFromInt(0)
	}

	if _, err := appstockcrud.UpdateSet(
		tx.AppStock.UpdateOneID(info.ID),
		&appstockcrud.Req{
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

func (h *Handler) AddStock(ctx context.Context) (*npool.Stock, error) {
	handler := &addHandler{
		Handler: h,
	}
	err := db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.addStock(ctx, tx); err != nil {
			return err
		}
		if err := handler.addAppStock(ctx, tx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetStock(ctx)
}
