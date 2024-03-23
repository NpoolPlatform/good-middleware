package appdefaultgood

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type updateHandler struct {
	*Handler
	sql        string
	coinTypeID string
}

func (h *updateHandler) constructSql() {
	now := uint32(time.Now().Unix())

	_sql := "update app_default_goods "
	_sql += fmt.Sprintf("set app_good_id = '%v', ", *h.AppGoodID)
	_sql += fmt.Sprintf("updated_at = %v ", now)
	_sql += "where "
	_sql += fmt.Sprintf("id = %v ", *h.ID)
	_sql += "and not exists ("
	_sql += "select 1 from (select * from app_default_goods) as adg "
	_sql += fmt.Sprintf("where adg.app_good_id = '%v' and adg.coin_type_id = '%v' and adg.id != %v", *h.AppGoodID, h.coinTypeID, *h.ID)
	_sql += " limit 1) and exists ("
	_sql += "select 1 from app_good_bases as agb "
	_sql += "left join good_bases as gb on agb.good_id = gb.ent_id "
	_sql += fmt.Sprintf("left join good_coins as gc on agb.good_id = gc.good_id and gc.coin_type_id = '%v'", h.coinTypeID)
	_sql += " limit 1)"

	h.sql = _sql
}

func (h *updateHandler) updateDefault(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	if _, err := rc.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func (h *Handler) UpdateDefault(ctx context.Context) error {
	handler := &updateHandler{
		Handler: h,
	}

	info, err := h.GetDefault(ctx)
	if err != nil {
		return err
	}
	if info == nil {
		return fmt.Errorf("invalid appdefaultgood")
	}

	handler.coinTypeID = info.CoinTypeID
	handler.constructSql()

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.updateDefault(_ctx, tx)
	})
}
