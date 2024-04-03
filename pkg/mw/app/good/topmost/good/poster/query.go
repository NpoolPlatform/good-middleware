package poster

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	topmostgoodpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/good/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgood"
	enttopmostgoodposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodposter"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TopMostGoodPosterSelect
	stmCount  *ent.TopMostGoodPosterSelect
	infos     []*npool.Poster
	total     uint32
}

func (h *queryHandler) selectPoster(stm *ent.TopMostGoodPosterQuery) *ent.TopMostGoodPosterSelect {
	return stm.Select(enttopmostgoodposter.FieldID)
}

func (h *queryHandler) queryPoster(cli *ent.Client) error {
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

func (h *queryHandler) queryPosters(cli *ent.Client) (*ent.TopMostGoodPosterSelect, error) {
	stm, err := topmostgoodpostercrud.SetQueryConds(cli.TopMostGoodPoster.Query(), h.PosterConds)
	if err != nil {
		return nil, err
	}
	return h.selectPoster(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
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

func (h *queryHandler) queryJoinTopMostGood(s *sql.Selector) {
	t1 := sql.Table(enttopmostgood.Table)
	t2 := sql.Table(enttopmost.Table)
	t3 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(enttopmostgoodposter.FieldTopMostGoodID),
			t1.C(enttopmostgood.FieldEntID),
		).
		LeftJoin(t2).
		On(
			t1.C(enttopmostgood.FieldTopMostID),
			t2.C(enttopmost.FieldEntID),
		).
		LeftJoin(t3).
		On(
			t1.C(enttopmostgood.FieldAppGoodID),
			t3.C(entappgoodbase.FieldEntID),
		).
		AppendSelect(
			t2.C(enttopmost.FieldAppID),
			t1.C(enttopmostgood.FieldTopMostID),
			sql.As(t2.C(enttopmost.FieldTopMostType), "top_most_type"),
			sql.As(t2.C(enttopmost.FieldTitle), "top_most_title"),
			sql.As(t2.C(enttopmost.FieldMessage), "top_most_message"),
			sql.As(t2.C(enttopmost.FieldTargetURL), "top_most_target_url"),
			t1.C(enttopmostgood.FieldAppGoodID),
			sql.As(t3.C(entappgoodbase.FieldName), "app_good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinTopMostGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinTopMostGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.TopMostType = types.GoodTopMostType(types.GoodTopMostType_value[info.TopMostTypeStr])
	}
}

func (h *Handler) GetPoster(ctx context.Context) (*npool.Poster, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryPoster(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid topmostgoodposter")
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetPosters(ctx context.Context) ([]*npool.Poster, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryPosters(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryPosters(cli)
		if err != nil {
			return err
		}

		handler.queryJoin()

		total, err := handler.stmCount.Count(_ctx)
		if err != nil {
			return err
		}
		handler.total = uint32(total)

		handler.stmSelect.
			Offset(int(handler.Offset)).
			Limit(int(handler.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	handler.formalize()

	return handler.infos, handler.total, nil
}
