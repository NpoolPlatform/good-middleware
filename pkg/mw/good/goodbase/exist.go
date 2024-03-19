package goodbase

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
)

func (h *Handler) ExistGoodBase(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.GoodBase.Query()
		if h.ID != nil {
			stm.Where(entgoodbase.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entgoodbase.EntID(*h.EntID))
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
