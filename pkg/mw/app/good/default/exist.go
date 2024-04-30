package appdefaultgood

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

func (h *Handler) ExistDefault(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDefault(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (h *Handler) ExistDefaultConds(ctx context.Context) (exist bool, err error) {
	handler := &baseQueryHandler{
		Handler: h,
	}
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if handler.stmSelect, err = handler.queryDefaults(cli); err != nil {
			return err
		}
		handler.queryJoin()
		exist, err = handler.stmSelect.Exist(ctx)
		return err
	})
	if err != nil {
		return false, err
	}
	return exist, nil
}
