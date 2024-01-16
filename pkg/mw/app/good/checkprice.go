package appgood

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"

	"github.com/shopspring/decimal"
)

func (h *Handler) checkUnitPrice(ctx context.Context) error {
	if h.UnitPrice == nil {
		return nil
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Good.
			Query().
			Where(
				entgood.EntID(*h.GoodID),
				entgood.DeletedAt(0),
			).
			Only(_ctx)
		if err != nil {
			return err
		}
		if h.UnitPrice.Cmp(info.UnitPrice) < 0 {
			return fmt.Errorf("invalid unitprice")
		}
		return nil
	})
}

func (h *Handler) checkPackagePrice(ctx context.Context) error {
	if h.PackagePrice == nil {
		return nil
	}
	if h.MinOrderDuration == nil || h.MaxOrderDuration == nil {
		return fmt.Errorf("invalid order duration")
	}
	if *h.MinOrderDuration != *h.MaxOrderDuration {
		return fmt.Errorf("invalid packageprice duration")
	}
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			Good.
			Query().
			Where(
				entgood.EntID(*h.GoodID),
				entgood.DeletedAt(0),
			).
			Only(_ctx)
		if err != nil {
			return err
		}
		packagePrice := info.UnitPrice.Mul(decimal.NewFromInt(int64(*h.MinOrderDuration)))
		if h.PackagePrice.Cmp(packagePrice) < 0 {
			return fmt.Errorf("invalid packageprice")
		}
		return nil
	})
}
