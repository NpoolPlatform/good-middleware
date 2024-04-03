package poster

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	topmostpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/topmost/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	enttopmost "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmost"
	enttopmostposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostposter"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/poster"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.TopMostPosterSelect
	stmCount  *ent.TopMostPosterSelect
	infos     []*npool.Poster
	total     uint32
}

func (h *queryHandler) selectPoster(stm *ent.TopMostPosterQuery) *ent.TopMostPosterSelect {
	return stm.Select(enttopmostposter.FieldID)
}

func (h *queryHandler) queryPoster(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.TopMostPoster.Query().Where(enttopmostposter.DeletedAt(0))
	if h.ID != nil {
		stm.Where(enttopmostposter.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(enttopmostposter.EntID(*h.EntID))
	}
	h.stmSelect = h.selectPoster(stm)
	return nil
}

func (h *queryHandler) queryPosters(cli *ent.Client) (*ent.TopMostPosterSelect, error) {
	stm, err := topmostpostercrud.SetQueryConds(cli.TopMostPoster.Query(), h.PosterConds)
	if err != nil {
		return nil, err
	}
	return h.selectPoster(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(enttopmostposter.Table)
	s.LeftJoin(t).
		On(
			s.C(enttopmostposter.FieldID),
			t.C(enttopmostposter.FieldID),
		).
		AppendSelect(
			t.C(enttopmostposter.FieldEntID),
			t.C(enttopmostposter.FieldTopMostID),
			t.C(enttopmostposter.FieldPoster),
			t.C(enttopmostposter.FieldIndex),
			t.C(enttopmostposter.FieldCreatedAt),
			t.C(enttopmostposter.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinTopMost(s *sql.Selector) {
	t1 := sql.Table(enttopmost.Table)
	s.LeftJoin(t1).
		On(
			s.C(enttopmostposter.FieldTopMostID),
			t1.C(enttopmost.FieldEntID),
		).
		AppendSelect(
			t1.C(enttopmost.FieldAppID),
			sql.As(t1.C(enttopmost.FieldTopMostType), "top_most_type"),
			sql.As(t1.C(enttopmost.FieldTitle), "top_most_title"),
			sql.As(t1.C(enttopmost.FieldMessage), "top_most_message"),
			sql.As(t1.C(enttopmost.FieldTargetURL), "top_most_target_url"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinTopMost(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinTopMost(s)
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
		return nil, fmt.Errorf("invalid topmostposter")
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
