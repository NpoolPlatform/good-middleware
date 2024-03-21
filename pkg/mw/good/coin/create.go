package coin

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
	_sql := "insert into good_coins "
	_sql += "  (ent_id, good_id, coin_type_id) "
	_sql += "values ("
	if h.EntID != nil {
		_sql += fmt.Sprintf("`%v`, ", *h.EntID)
	}
	_sql += fmt.Sprintf("`%v`, ", *h.GoodID)
	_sql += fmt.Sprintf("`%v`", *h.CoinTypeID)
	_sql += ") "
	_sql += "where not exists ("
	_sql += fmt.Sprintf("  select 1 from good_coins where good_id='%v' and coin_type_id='%v'", *h.GoodID, *h.CoinTypeID)
	_sql += ")"
}

func (h *createHandler) createGoodCoin(ctx context.Context, tx *ent.Tx) error {
	_sql := "insert into good_coins "
	_sql += "  (ent_id, good_id, coin_type_id) "
	_sql += "values ("
	if h.EntID != nil {
		_sql += fmt.Sprintf("`%v`, ", *h.EntID)
	}
	_sql += fmt.Sprintf("`%v`, ", *h.GoodID)
	_sql += fmt.Sprintf("`%v`", *h.CoinTypeID)
	_sql += ") "
	_sql += "where not exists ("
	_sql += fmt.Sprintf("  select 1 from good_coins where good_id='%v' and coin_type_id='%v'", *h.GoodID, *h.CoinTypeID)
	_sql += ")"

	rc, err := tx.ExecContext(ctx, _sql)
	if err != nil {
		return err
	}
	if n, err := rc.RowsAffected(); err != nil || n != 1 {
		return fmt.Errorf("fail create goodcoin")
	}
	return nil
}

func (h *Handler) CreateGoodCoin(ctx context.Context) error {
	handler := &createHandler{
		Handler: h,
	}
	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		return handler.createGoodCoin(_ctx, tx)
	})
}
