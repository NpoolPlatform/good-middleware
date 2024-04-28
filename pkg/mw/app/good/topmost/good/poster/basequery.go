package poster

import (
	"fmt"

	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	topmostgoodpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	enttopmostgoodposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodposter"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.TopMostGoodPosterSelect
}

func (h *baseQueryHandler) selectPoster(stm *ent.TopMostGoodPosterQuery) *ent.TopMostGoodPosterSelect {
	return stm.Select(enttopmostgoodposter.FieldID)
}

func (h *baseQueryHandler) queryPoster(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.TopMostGoodPoster.Query().Where(enttopmostgoodposter.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostgoodposter.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostgoodposter.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPoster(stm)
	return nil
}

func (h *baseQueryHandler) queryPosters(cli *ent.Client) (*ent.TopMostGoodPosterSelect, error) {
	stm, err := topmostgoodpostercrud.SetQueryConds(cli.TopMostGoodPoster.Query(), h.PosterConds)
	if err != nil {
		return nil, err
	}
	return h.selectPoster(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostgoodposter.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostgoodposter.FieldID),
			t.C(enttopmostgoodposter.FieldID),
		).
		AppendSelect(
			t.C(enttopmostgoodposter.FieldEntID),
			t.C(enttopmostgoodposter.FieldTopMostGoodID),
			t.C(enttopmostgoodposter.FieldPoster),
			t.C(enttopmostgoodposter.FieldIndex),
			t.C(enttopmostgoodposter.FieldCreatedAt),
			t.C(enttopmostgoodposter.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinTopMostGood(s *sql.Selector) error {
	t1 := sql.Table(enttopmostgood.Table)
	t2 := sql.Table(enttopmost.Table)
	t3 := sql.Table(entappgoodbase.Table)
	s.Join(t1).
		On(
			s.C(enttopmostgoodposter.FieldTopMostGoodID),
			t1.C(enttopmostgood.FieldEntID),
		).
		Join(t2).
		On(
			t1.C(enttopmostgood.FieldTopMostID),
			t2.C(enttopmost.FieldEntID),
		).
		Join(t3).
		On(
			t1.C(enttopmostgood.FieldAppGoodID),
			t3.C(entappgoodbase.FieldEntID),
		)
	if h.TopMostConds.EntID != nil {
		entID, ok := h.TopMostConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid topmostentid")
		}
		s.OnP(
			sql.EQ(t2.C(enttopmost.FieldEntID), entID),
		)
	}
	if h.TopMostConds.AppID != nil {
		appID, ok := h.TopMostConds.AppID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid appid")
		}
		s.OnP(
			sql.EQ(t2.C(enttopmost.FieldAppID), appID),
		)
	}
	s.AppendSelect(
		t2.C(enttopmost.FieldAppID),
		t1.C(enttopmostgood.FieldTopMostID),
		sql.As(t2.C(enttopmost.FieldTopMostType), "top_most_type"),
		sql.As(t2.C(enttopmost.FieldTitle), "top_most_title"),
		sql.As(t2.C(enttopmost.FieldMessage), "top_most_message"),
		sql.As(t2.C(enttopmost.FieldTargetURL), "top_most_target_url"),
		t1.C(enttopmostgood.FieldAppGoodID),
		sql.As(t3.C(entappgoodbase.FieldName), "app_good_name"),
	)
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		if err := h.queryJoinTopMostGood(s); err != nil {
			logger.Sugar().Errorw("queryJoinTopMostGood", "Error", err)
		}
	})
}
