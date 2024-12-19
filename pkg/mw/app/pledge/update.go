//nolint:dupl
package pledge

import (
	"context"
	"fmt"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	appgoodbase1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/goodbase"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
)

type updateHandler struct {
	*pledgeAppGoodQueryHandler
	sqlAppPledge   string
	sqlAppGoodBase string
	updated        bool
}

func (h *updateHandler) constructAppGoodBaseSQL(ctx context.Context) error {
	handler, err := appgoodbase1.NewHandler(
		ctx,
		appgoodbase1.WithEntID(func() *string { s := h.AppGoodBaseReq.EntID.String(); return &s }(), true),
		appgoodbase1.WithAppID(func() *string { s := h.AppGoodBaseReq.AppID.String(); return &s }(), false),
		appgoodbase1.WithGoodID(func() *string { s := h.AppGoodBaseReq.GoodID.String(); return &s }(), false),
		appgoodbase1.WithName(h.AppGoodBaseReq.Name, false),
		appgoodbase1.WithPurchasable(h.AppGoodBaseReq.Purchasable, false),
		appgoodbase1.WithEnableProductPage(h.AppGoodBaseReq.EnableProductPage, false),
		appgoodbase1.WithProductPage(h.AppGoodBaseReq.ProductPage, false),
		appgoodbase1.WithOnline(h.AppGoodBaseReq.Online, false),
		appgoodbase1.WithVisible(h.AppGoodBaseReq.Visible, false),
		appgoodbase1.WithDisplayIndex(h.AppGoodBaseReq.DisplayIndex, false),
		appgoodbase1.WithBanner(h.AppGoodBaseReq.Banner, false),
	)
	if err != nil {
		return wlog.WrapError(err)
	}
	h.sqlAppGoodBase, err = handler.ConstructUpdateSQL()
	if err != nil && !wlog.Equal(err, cruder.ErrUpdateNothing) {
		return wlog.WrapError(err)
	}
	return nil
}

func (h *updateHandler) constructAppPledgeSQL() {
	set := "set "
	now := uint32(time.Now().Unix())
	_sql := "update app_pledges "
	if h.ServiceStartAt != nil {
		_sql += fmt.Sprintf("%vservice_start_at = %v,", set, *h.ServiceStartAt)
		set = ""
	}
	if h.StartMode != nil {
		_sql += fmt.Sprintf("%vstart_mode = '%v',", set, h.StartMode.String())
		set = ""
	}
	if h.EnableSetCommission != nil {
		_sql += fmt.Sprintf("%venable_set_commission = %v,", set, *h.EnableSetCommission)
		set = ""
	}
	if set != "" {
		return
	}
	_sql += fmt.Sprintf("updated_at = %v", now)
	_sql += fmt.Sprintf(" where id = %v ", *h.ID)
	_sql += fmt.Sprintf(" and ent_id = '%v' ", *h.EntID)
	_sql += fmt.Sprintf(" and app_good_id = '%v'", *h.AppGoodID)

	h.sqlAppPledge = _sql
}

func (h *updateHandler) execSQL(ctx context.Context, tx *ent.Tx, sql string) error {
	if sql == "" {
		return nil
	}
	rc, err := tx.ExecContext(ctx, sql)
	if err != nil {
		return wlog.WrapError(err)
	}
	n, err := rc.RowsAffected()
	if err != nil {
		return wlog.Errorf("fail update apppledge: %v", err)
	}
	h.updated = n == 1
	return nil
}

func (h *updateHandler) updateAppPledge(ctx context.Context, tx *ent.Tx) error {
	if h.sqlAppPledge == "" {
		return nil
	}
	return h.execSQL(ctx, tx, h.sqlAppPledge)
}

func (h *updateHandler) updateAppGoodBase(ctx context.Context, tx *ent.Tx) error {
	return h.execSQL(ctx, tx, h.sqlAppGoodBase)
}

//nolint:gocyclo
func (h *Handler) UpdatePledge(ctx context.Context) error {
	handler := &updateHandler{
		pledgeAppGoodQueryHandler: &pledgeAppGoodQueryHandler{
			Handler: h,
		},
	}

	if err := handler.requireAppPledgeAppGood(ctx); err != nil {
		return wlog.WrapError(err)
	}

	if h.ID == nil {
		h.ID = &handler._ent.appPledge.ID
	}
	if h.EntID == nil {
		h.EntID = &handler._ent.appPledge.EntID
	}
	if h.AppGoodBaseReq.GoodID == nil {
		h.AppGoodBaseReq.GoodID = &handler._ent.appGoodBase.GoodID
	}
	if h.AppGoodID == nil {
		h.AppGoodID = &handler._ent.appPledge.AppGoodID
		h.AppGoodBaseReq.EntID = h.AppGoodID
	}
	if h.AppGoodBaseReq.AppID == nil {
		h.AppGoodBaseReq.AppID = &handler._ent.appGoodBase.AppID
	}
	if h.AppGoodBaseReq.Online != nil && (!handler._ent.goodBase.Online && *h.AppGoodBaseReq.Online) {
		return wlog.Errorf("invalid online")
	}
	if h.AppGoodBaseReq.Purchasable != nil && (!handler._ent.goodBase.Purchasable && *h.AppGoodBaseReq.Purchasable) {
		return wlog.Errorf("invalid purchasable")
	}

	handler.constructAppPledgeSQL()
	if err := handler.constructAppGoodBaseSQL(ctx); err != nil {
		return wlog.WrapError(err)
	}

	return db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		if err := handler.updateAppGoodBase(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if err := handler.updateAppPledge(_ctx, tx); err != nil {
			return wlog.WrapError(err)
		}
		if !handler.updated {
			return wlog.WrapError(cruder.ErrUpdateNothing)
		}
		return nil
	})
}
