package appdefaultgood

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
)

type queryAppGoodHandler struct {
	*Handler
}

func (h *queryAppGoodHandler) getAppGood(ctx context.Context) error {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err := cli.
			AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(*h.AppGoodID),
				entappgoodbase.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}

		info1, err := cli.
			GoodBase.
			Query().
			Where(
				entgoodbase.EntID(info.GoodID),
				entgoodbase.DeletedAt(0),
			).
			Only(ctx)
		if err != nil {
			return err
		}
		h.AppID = &info.AppID
		h.GoodID = &info.GoodID
		h.CoinTypeID = &info1.CoinTypeID
		return nil
	})
}
