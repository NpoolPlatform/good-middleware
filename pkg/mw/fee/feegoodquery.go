package fee

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entfee "github.com/NpoolPlatform/good-middleware/pkg/db/ent/fee"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
)

type feeGoodQueryHandler struct {
	*Handler
	fee      *ent.Fee
	goodBase *ent.GoodBase
}

func (h *feeGoodQueryHandler) _getFeeGood(ctx context.Context, must bool) (err error) {
	if h.ID == nil && h.EntID == nil && h.GoodID == nil {
		return fmt.Errorf("invalid id")
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm := cli.Fee.Query()
		if h.ID != nil {
			stm.Where(entfee.ID(*h.ID))
		}
		if h.EntID != nil {
			stm.Where(entfee.EntID(*h.EntID))
		}
		if h.GoodID != nil {
			stm.Where(entfee.GoodID(*h.GoodID))
		}
		if h.fee, err = stm.Only(_ctx); err != nil {
			if ent.IsNotFound(err) && !must {
				return nil
			}
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
