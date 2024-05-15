package goodbase

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
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
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}

func (h *Handler) ExistGoodBaseConds(ctx context.Context) (exist bool, err error) {
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
		if err != nil {
			return wlog.WrapError(err)
		}
		exist, err = stm.Exist(_ctx)
		return wlog.WrapError(err)
	})
	if err != nil {
		return false, wlog.WrapError(err)
	}
	return exist, nil
}
