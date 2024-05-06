package required

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) UpdateRequired(ctx context.Context) error {
	info, err := h.GetRequired(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid required")
	}
	h.ID = &info.ID
	return db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if _, err := requiredcrud.UpdateSet(
			cli.RequiredAppGood.UpdateOneID(*h.ID),
			&requiredcrud.Req{
				Must: h.Must,
			},
		).Save(_ctx); err != nil {
			return wlog.WrapError(err)
		}
		return nil
	})
}
