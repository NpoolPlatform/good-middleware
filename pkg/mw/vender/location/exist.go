package location

import (
	"context"
	"fmt"

	locationcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/location"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entvendorlocation "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorlocation"
)

func (h *Handler) ExistLocation(ctx context.Context) (bool, error) {
	if h.EntID == nil {
		return false, fmt.Errorf("invalid entid")
	}

	exist := false
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.
			VendorLocation.
			Query().
			Where(
				entvendorlocation.EntID(*h.EntID),
				entvendorlocation.DeletedAt(0),
			).
			Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistLocationConds(ctx context.Context) (bool, error) {
	exist := false
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := locationcrud.SetQueryConds(cli.VendorLocation.Query(), h.Conds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
