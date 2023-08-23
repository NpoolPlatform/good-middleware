package required

import (
	"context"
	"fmt"

	redis2 "github.com/NpoolPlatform/go-service-framework/pkg/redis"
	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"

	"github.com/google/uuid"
)

func (h *Handler) CreateRequired(ctx context.Context) (*npool.Required, error) {
	key := fmt.Sprintf("%v:%v:%v", basetypes.Prefix_PrefixCreateRequiredGood, *h.MainGoodID, *h.RequiredGoodID)
	if err := redis2.TryLock(key, 0); err != nil {
		return nil, err
	}
	defer func() {
		_ = redis2.Unlock(key)
	}()

	h.Conds = &requiredcrud.Conds{
		MainGoodID:     &cruder.Cond{Op: cruder.EQ, Val: *h.MainGoodID},
		RequiredGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.RequiredGoodID},
	}
	exist, err := h.ExistRequiredConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("already exists")
	}

	h.Conds = &requiredcrud.Conds{
		MainGoodID: &cruder.Cond{Op: cruder.EQ, Val: *h.RequiredGoodID},
	}
	exist, err = h.ExistRequiredConds(ctx)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("not allowed")
	}

	id := uuid.New()
	if h.ID == nil {
		h.ID = &id
	}

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := requiredcrud.CreateSet(
			cli.RequiredGood.Create(),
			&requiredcrud.Req{
				ID:             h.ID,
				MainGoodID:     h.MainGoodID,
				RequiredGoodID: h.RequiredGoodID,
				Must:           h.Must,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return h.GetRequired(ctx)
}
