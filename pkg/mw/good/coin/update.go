package coin

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*Handler
	sql    string
	goodID string
}

func (h *updateHandler) constructSQL() error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}

	set := "set "
	now := uint32(time.Now().Unix())

	_sql := "update good_coins "
	if h.Main != nil {
		_sql += fmt.Sprintf("%vmain = %v, ", set, *h.Main)
		set = ""
	}
	if h.Index != nil {
		_sql += fmt.Sprintf("%v`index` = %v, ", set, *h.Index)
		set = ""
	}
	if set != "" {
		return wlog.WrapError(cruder.ErrUpdateNothing)
	}
	_sql += fmt.Sprintf("updated_at = %v ", now)
	whereAnd := "where "
	if h.ID != nil {
		_sql += fmt.Sprintf("where id = %v ", *h.ID)
		whereAnd = "and"
	}
	if h.EntID != nil {
		_sql += fmt.Sprintf("%v ent_id = '%v'", whereAnd, *h.EntID)
	}
	if h.Main != nil && *h.Main {
		_sql += fmt.Sprintf(
			" and not exist (good_id = '%v' and deleted_at = 0 and main = 1)",
			h.goodID,
		)
	}
	h.sql = _sql
	return nil
}

func (h *updateHandler) updateGoodCoin(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return wlog.Errorf("fail update goodcoin: %v", err)
	}
	return nil
}

func (h *Handler) UpdateGoodCoin(ctx context.Context) error {
	info, err := h.GetGoodCoin(ctx)
	if err != nil {
		return wlog.WrapError(err)
	}
	if info == nil {
		return wlog.Errorf("invalid goodcoin")
	}

	h.ID = &info.ID
	handler := &updateHandler{
		Handler: h,
		goodID:  info.GoodID,
	}
	if err := handler.constructSQL(); err != nil {
		return wlog.WrapError(err)
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateGoodCoin(_ctx, tx)
	})
}
