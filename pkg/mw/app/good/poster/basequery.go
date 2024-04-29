package poster

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	appgoodpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgoodposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodposter"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.AppGoodPosterSelect
}

func (h *baseQueryHandler) selectPoster(stm *ent.AppGoodPosterQuery) *ent.AppGoodPosterSelect {
	return stm.Select(entappgoodposter.FieldID)
}

func (h *baseQueryHandler) queryPoster(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppGoodPoster.Query().Where(entappgoodposter.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgoodposter.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgoodposter.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPoster(stm)
	return nil
}

func (h *baseQueryHandler) queryPosters(cli *ent.Client) (*ent.AppGoodPosterSelect, error) {
	stm, err := appgoodpostercrud.SetQueryConds(cli.AppGoodPoster.Query(), h.PosterConds)
	if err != nil {
		return nil, err
	}
	return h.selectPoster(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgoodposter.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodposter.FieldID),
			t.C(entappgoodposter.FieldID),
		).
		AppendSelect(
			t.C(entappgoodposter.FieldEntID),
			t.C(entappgoodposter.FieldAppGoodID),
			t.C(entappgoodposter.FieldPoster),
			t.C(entappgoodposter.FieldIndex),
			t.C(entappgoodposter.FieldCreatedAt),
			t.C(entappgoodposter.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinAppGood(s *sql.Selector) error {
	t1 := sql.Table(entappgoodbase.Table)
	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgoodposter.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
		).
		LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		)
	if h.GoodBaseConds.EntID != nil {
		id, ok := h.GoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodid")
		}
		s.OnP(
			sql.EQ(t2.C(entgoodbase.FieldEntID), id),
		)
	}
	if h.GoodBaseConds.EntIDs != nil {
		ids, ok := h.GoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodids")
		}
		s.OnP(
			sql.In(t1.C(entgoodbase.FieldEntID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}),
		)
	}
	if h.AppGoodBaseConds.EntID != nil {
		id, ok := h.AppGoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid appgoodid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldEntID), id),
		)
	}
	if h.AppGoodBaseConds.EntIDs != nil {
		ids, ok := h.AppGoodBaseConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid appgoodids")
		}
		s.OnP(
			sql.In(t1.C(entappgoodbase.FieldEntID), func() (_ids []interface{}) {
				for _, id := range ids {
					_ids = append(_ids, interface{}(id))
				}
				return
			}),
		)
	}
	if h.AppGoodBaseConds.AppID != nil {
		id, ok := h.AppGoodBaseConds.AppID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t1.C(entappgoodbase.FieldAppID), id),
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
