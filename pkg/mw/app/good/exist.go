package good

import (
	"context"

	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) ExistGoodConds(ctx context.Context) (exist bool, err error) {
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := appgoodbasecrud.SetQueryConds(cli.AppGoodBase.Query(), h.AppGoodConds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		return err
	}); err != nil {
		return false, err
	}
	return exist, nil
}
