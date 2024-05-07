package required

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	requiredcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/required"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entrequiredappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredappgood"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.RequiredAppGoodSelect
}

func (h *baseQueryHandler) selectRequired(stm *ent.RequiredAppGoodQuery) *ent.RequiredAppGoodSelect {
	return stm.Select(entrequiredappgood.FieldID)
}

func (h *baseQueryHandler) queryRequired(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return wlog.Errorf("invalid id")
	}
	stm := cli.RequiredAppGood.Query().Where(entrequiredappgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrequiredappgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrequiredappgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRequired(stm)
	return nil
}

func (h *baseQueryHandler) queryRequireds(cli *ent.Client) (*ent.RequiredAppGoodSelect, error) {
	stm, err := requiredcrud.SetQueryConds(cli.RequiredAppGood.Query(), h.RequiredConds)
	if err != nil {
		return nil, err
	}
	return h.selectRequired(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrequiredappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entrequiredappgood.FieldID),
			t.C(entrequiredappgood.FieldID),
		).
		AppendSelect(
			t.C(entrequiredappgood.FieldEntID),
			t.C(entrequiredappgood.FieldMainAppGoodID),
			t.C(entrequiredappgood.FieldRequiredAppGoodID),
			t.C(entrequiredappgood.FieldMust),
			t.C(entrequiredappgood.FieldCreatedAt),
			t.C(entrequiredappgood.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryWithAppGoodBaseConds(t1 *sql.SelectTable, s *sql.Selector) error {
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
	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldAppID), id),
		)
	}
	return nil
}

func (h *baseQueryHandler) queryJoinMainAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entrequiredappgood.FieldMainAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		)
	if err := h.queryWithAppGoodBaseConds(t1, s); err != nil {
		return wlog.WrapError(err)
	}
	s.AppendSelect(
		t1.C(entappgoodbase.FieldAppID),
		sql.As(t1.C(entappgoodbase.FieldName), "main_app_good_name"),
		sql.As(t1.C(entappgoodbase.FieldGoodID), "main_good_id"),
		sql.As(t2.C(entgoodbase.FieldName), "main_good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinRequiredAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entrequiredappgood.FieldRequiredAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		)
	if err := h.queryWithAppGoodBaseConds(t1, s); err != nil {
		return wlog.WrapError(err)
	}
	s.AppendSelect(
		sql.As(t1.C(entappgoodbase.FieldName), "required_app_good_name"),
		sql.As(t1.C(entappgoodbase.FieldGoodID), "required_good_id"),
		sql.As(t2.C(entgoodbase.FieldName), "required_good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinMainAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinMainAppGood", "Error", err)
		}
		if err := h.queryJoinRequiredAppGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinRequiredAppGood", "Error", err)
		}
	})
}
