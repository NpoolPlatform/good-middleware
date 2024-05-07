package displaycolor

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	appgooddisplaycolorcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/display/color"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgooddisplaycolor "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplaycolor"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodDisplayColorSelect
}

func (h *baseQueryHandler) selectDisplayColor(stm *ent.AppGoodDisplayColorQuery) *ent.AppGoodDisplayColorSelect {
	return stm.Select(entappgooddisplaycolor.FieldID)
}

func (h *baseQueryHandler) queryDisplayColor(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.AppGoodDisplayColor.Query().Where(entappgooddisplaycolor.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgooddisplaycolor.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgooddisplaycolor.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDisplayColor(stm)
	return nil
}

func (h *baseQueryHandler) queryDisplayColors(cli *ent.Client) (*ent.AppGoodDisplayColorSelect, error) {
	stm, err := appgooddisplaycolorcrud.SetQueryConds(cli.AppGoodDisplayColor.Query(), h.DisplayColorConds)
	if err != nil {
		return nil, err
	}
	return h.selectDisplayColor(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgooddisplaycolor.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgooddisplaycolor.FieldID),
			t.C(entappgooddisplaycolor.FieldID),
		).
		AppendSelect(
			t.C(entappgooddisplaycolor.FieldEntID),
			t.C(entappgooddisplaycolor.FieldAppGoodID),
			t.C(entappgooddisplaycolor.FieldColor),
			t.C(entappgooddisplaycolor.FieldIndex),
			t.C(entappgooddisplaycolor.FieldCreatedAt),
			t.C(entappgooddisplaycolor.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgooddisplaycolor.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		)
	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldAppID), id),
		)
	}
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldEntID), id),
		)
	}
	if h.AppGoodBaseConds.EntIDs != nil {
		ids, ok := h.AppGoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(t1.C(entappgoodbase.FieldEntID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	if h.AppGoodBaseConds.GoodID != nil {
		id, ok := h.AppGoodBaseConds.GoodID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldGoodID), id),
		)
	}
	if h.AppGoodBaseConds.GoodIDs != nil {
		ids, ok := h.AppGoodBaseConds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(t1.C(entappgoodbase.FieldGoodID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}()...),
		)
	}
	s.AppendSelect(
		t1.C(entappgoodbase.FieldAppID),
		sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
		sql.As(t2.C(entgoodbase.FieldEntID), "good_id"),
		sql.As(t2.C(entgoodbase.FieldName), "good_name"),
		t2.C(entgoodbase.FieldGoodType),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinAppGood", "Error", err)
		}
	})
}
