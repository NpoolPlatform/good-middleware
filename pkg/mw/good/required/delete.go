package required

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"
)

type deleteHandler struct {
	*Handler
	info *npool.Required
	sql  string
}


func (h *deleteHandler) constructSQL() {
	_sql := "select 1"
	_sql += "from required_goods rg"
	_sql += "join app_good_bases agb on rg.main_good_id = agb.good_id or rg.required_good_id = agb.good_id"
	_sql += "join required_app_goods rag on agb.ent_id = rag.main_app_good_id or agb.ent_id = rag.required_app_good_id"
	_sql += fmt.Sprintf("where rg.main_good_id = '%v' and rg.required_good_id = '%v' ", h.info.MainGoodID, h.info.RequiredGoodID)
	_sql += "and rg.deleted_at=0 and agb.deleted_at=0 and rag.deleted_at=0 limit 1"

	h.sql = _sql
}

func (h *deleteHandler) checkAppGoodRequired(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil || n == 1 {
		return wlog.Errorf("fail delete requiredgood: %v", err)
	}
	return nil
}

func (h *Handler) DeleteRequired(ctx context.Context) error {
	info, err := h.GetRequired(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return nil
	}

	h.ID = &info.ID
	now := uint32(time.Now().Unix())

	handler := &deleteHandler{
		Handler: h,
		info:    info,
	}

	handler.constructSQL()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.checkAppGoodRequired(ctx, tx); err != nil {
			return err
		}
		if _, err := requiredcrud.UpdateSet(
			tx.RequiredGood.UpdateOneID(*h.ID),
			&requiredcrud.Req{
				DeletedAt: &now,
			},
		).Save(_ctx); err != nil {
			return err
		}
		return nil
	})
}
