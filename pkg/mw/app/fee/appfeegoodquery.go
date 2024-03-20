package fee

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appfee"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
)

type appFeeGoodQueryHandler struct {
	*Handler
	appFee      *ent.AppFee
	appGoodBase *ent.AppGoodBase
}

func (h *appFeeGoodQueryHandler) _getAppFeeGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.AppGoodID == nil {
		return fmt.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.AppFee.Query()
		if h.ID != nil {
			stm.Where(entappfee.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entappfee.EntID(*h.EntID))
		}
		if h.AppGoodID != nil {
			stm.Where(entappfee.AppGoodID(*h.AppGoodID))
		}
		if h.appFee, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return err
		}
		if h.appGoodBase, err = cli.
			AppGoodBase.
			Query().
			Where(
				entappgoodbase.EntID(h.appFee.AppGoodID),
				entappgoodbase.DeletedAt(0),
			).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return err
		}
		return nil
	})
}

func (h *appFeeGoodQueryHandler) getAppFeeGood(ctx context.Context) error {
	return h._getAppFeeGood(ctx, false)
}

func (h *appFeeGoodQueryHandler) requireAppFeeGood(ctx context.Context) error {
	return h._getAppFeeGood(ctx, true)
}
