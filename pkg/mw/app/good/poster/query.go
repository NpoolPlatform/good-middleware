package poster

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgoodpostercrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/poster"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgoodposter "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodposter"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/poster"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppGoodPosterSelect
	stmCount  *ent.AppGoodPosterSelect
	infos     []*npool.Poster
	total     uint32
}

func (h *queryHandler) selectPoster(stm *ent.AppGoodPosterQuery) *ent.AppGoodPosterSelect {
	return stm.Select(entappgoodposter.FieldID)
}

func (h *queryHandler) queryPoster(cli *ent.Client) error {
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

func (h *queryHandler) queryPosters(cli *ent.Client) (*ent.AppGoodPosterSelect, error) {
	stm, err := appgoodpostercrud.SetQueryConds(cli.AppGoodPoster.Query(), h.PosterConds)
	if err != nil {
		return nil, err
	}
	return h.selectPoster(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
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

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgoodposter.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		AppendSelect(
			t1.C(entappgoodbase.FieldAppID),
			sql.As(t1.C(entappgoodbase.FieldName), "app_good_name"),
		)

	t2 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t2).
		On(
			t1.C(entappgoodbase.FieldGoodID),
			t2.C(entgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t2.C(entgoodbase.FieldEntID), "good_id"),
			sql.As(t2.C(entgoodbase.FieldName), "good_name"),
			t2.C(entgoodbase.FieldGoodType),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
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
		return nil, fmt.Errorf("invalid appgoodposter")
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
