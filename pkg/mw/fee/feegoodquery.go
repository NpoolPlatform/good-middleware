package fee

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
)

type feeGoodQueryHandler struct {
	*Handler
	fee      *ent.Fee
	goodBase *ent.GoodBase
}

func (h *feeGoodQueryHandler) _getFeeGood(ctx context.Context, must bool) (err error) {
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if h.fee, err = cli.Fee.Get(_ctx, *h.ID); err != nil {
			return err
		}
		if h.goodBase, err = cli.
			GoodBase.
			Query().
			Where(
				entgoodbase.EntID(h.fee.GoodID),
				entgoodbase.DeletedAt(0),
			).Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
			return err
		}
		return nil
	})
}

func (h *feeGoodQueryHandler) getFeeGood(ctx context.Context) error {
	return h._getFeeGood(ctx, false)
}

func (h *feeGoodQueryHandler) requireFeeGood(ctx context.Context) error {
	return h._getFeeGood(ctx, true)
}
