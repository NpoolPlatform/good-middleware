package goodbase

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
	sql string
}

func (h *createHandler) constructSql() {
	h.sql = h.ConstructCreateSql()
}

func (h *createHandler) createGoodBase(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("fail create goodbase")
	}
	return nil
}

func (h *Handler) CreateGoodBase(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createGoodBase(_ctx, tx)
	})
}
