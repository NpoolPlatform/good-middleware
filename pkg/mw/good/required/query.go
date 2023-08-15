package required

import (
	"context"
	"fmt"

	// requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entrequiredgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredgood"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.RequiredGoodSelect
	infos     []*npool.Required
	total     uint32
}

func (h *queryHandler) selectRequired(stm *ent.RequiredGoodQuery) {
	h.stmSelect = stm.Select(
		entrequiredgood.FieldID,
		entrequiredgood.FieldMainGoodID,
		entrequiredgood.FieldRequiredGoodID,
		entrequiredgood.FieldMust,
		entrequiredgood.FieldCommission,
		entrequiredgood.FieldCreatedAt,
		entrequiredgood.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryRequired(cli *ent.Client) {
	h.selectRequired(
		cli.RequiredGood.
			Query().
			Where(
				entrequiredgood.ID(*h.ID),
				entrequiredgood.DeletedAt(0),
			),
	)
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetRequired(ctx context.Context) (*npool.Required, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryRequired(cli)
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}
