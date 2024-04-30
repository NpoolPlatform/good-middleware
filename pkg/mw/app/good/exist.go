package good

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) ExistGoodConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	if err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryGoods(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(_ctx)
		return err
	}); err != nil {
		return false, err
	}
	return exist, nil
}
