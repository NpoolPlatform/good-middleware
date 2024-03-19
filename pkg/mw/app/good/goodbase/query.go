package goodbase

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
)

func (h *Handler) GetGoodBase(ctx context.Context) (AppGoodBase, error) {
	var _goodBase *ent.AppGoodBase
	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppGoodBase.Query()
		if h.ID != nil {
			stm.Where(entappgoodbase.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entappgoodbase.EntID(*h.EntID))
		}
		_goodBase, err = stm.Only(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &goodBase{
		_ent: _goodBase,
	}, nil
}
