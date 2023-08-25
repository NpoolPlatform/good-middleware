package topmost

import (
	"context"
	"fmt"

	topmostcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

func (h *Handler) checkPromotion(ctx context.Context) error {
	if *h.TopMostType != types.GoodTopMostType_TopMostPromotion {
		return nil
	}

	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := topmostcrud.SetQueryConds(cli.TopMost.Query(), &topmostcrud.Conds{
			AppID:       &cruder.Cond{Op: cruder.EQ, Val: *h.AppID},
			TopMostType: &cruder.Cond{Op: cruder.EQ, Val: *h.TopMostType},
			StartEnd:    &cruder.Cond{Op: cruder.OVERLAP, Val: []uint32{*h.StartAt, *h.EndAt}},
		})
		if err != nil {
			return err
		}
		exist, err := stm.Exist(_ctx)
		if err != nil {
			return err
		}
		if exist {
			return fmt.Errorf("time overlap")
		}
		return nil
	})
}
