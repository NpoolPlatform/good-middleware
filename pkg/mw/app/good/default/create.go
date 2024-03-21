package appdefaultgood

import (
	"context"
	"fmt"
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
)

type createHandler struct {
	*Handler
	sql string
}

func (h *createHandler) constructSql() {
	now := uint32(time.Now().Unix())
	comma := ""
	_sql := "insert into app_default_goods ("
	if h.EntID != nil {
		_sql += "ent_id"
		comma = ", "
	}
	_sql += comma + "app_good_id"
	comma = ", "
	_sql += comma + "coin_type_id"
	_sql += comma + "created_at"
	_sql += comma + "updated_at"
	_sql += comma + "deleted_at"
	_sql += ")"

	comma = ""
	_sql += " select * from ( select "
	if h.EntID != nil {
		_sql += fmt.Sprintf("'%v' as ent_id", *h.EntID)
		comma = ", "
	}
	_sql += fmt.Sprintf("%v'%v' as app_good_id", comma, *h.AppGoodID)
	comma = ", "
	_sql += fmt.Sprintf("%v'%v' as coin_type_id", comma, *h.CoinTypeID)
	_sql += fmt.Sprintf("%v%v as created_at", comma, now)
	_sql += fmt.Sprintf("%v%v as updated_at", comma, now)
	_sql += fmt.Sprintf("%v0 as deleted_at", comma)
	_sql += ") as tmp "
	_sql += "where not exists ("
	_sql += "select 1 from app_default_goods as adg "
	_sql += fmt.Sprintf("where adg.app_good_id = '%v' and adg.coin_type_id='%v'", *h.AppGoodID, *h.CoinTypeID)
	_sql += ") and exists ("
	_sql += "select 1 from app_good_bases as agb "
	_sql += "left join good_bases as gb on agb.good_id = gb.ent_id "
	_sql += fmt.Sprintf("left join good_coins as gc on agb.good_id = gc.good_id and gc.coin_type_id = '%v'", *h.CoinTypeID)
	_sql += ")"

	h.sql = _sql
}

func (h *createHandler) createDefault(ctx context.Context, tx *ent.Tx) error {
	rc, err := tx.ExecContext(ctx, h.sql)
	if err != nil {
		return err
	}
	n, err := rc.RowsAffected()
	if err != nil || n != 1 {
		return fmt.Errorf("fail create appdefaultgood: %v", err)
	}
	return nil
}

func (h *Handler) CreateDefault(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	handler.constructSql()
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createDefault(_ctx, tx)
	})
}
