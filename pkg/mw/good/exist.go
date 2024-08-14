package good

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) ExistGoodConds(ctx context.Context) (exist bool, err error) {
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodConds)
		if err != nil {
			return err
		}
		exist, err = stm.Exist(_ctx)
		return err
	}); err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
