package like

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	likecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/like"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entlike "github.com/NpoolPlatform/good-middleware/pkg/db/ent/like"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/like"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.LikeSelect
	stmCount  *ent.LikeSelect
	infos     []*npool.Like
	total     uint32
}

func (h *queryHandler) selectLike(stm *ent.LikeQuery) *ent.LikeSelect {
	return stm.Select(entlike.FieldID)
}

func (h *queryHandler) queryLike(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Like.Query().Where(entlike.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entlike.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entlike.EntID(*h.EntID))
	}
	h.stmSelect = h.selectLike(stm)
	return nil
}

func (h *queryHandler) queryLikes(cli *ent.Client) (*ent.LikeSelect, error) {
	stm, err := likecrud.SetQueryConds(cli.Like.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectLike(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entlike.Table)
	s.LeftJoin(t).
		On(
			s.C(entlike.FieldID),
			t.C(entlike.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entlike.FieldEntID), "ent_id"),
			sql.As(t.C(entlike.FieldAppID), "app_id"),
			sql.As(t.C(entlike.FieldUserID), "user_id"),
			sql.As(t.C(entlike.FieldGoodID), "good_id"),
			sql.As(t.C(entlike.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entlike.FieldLike), "like"),
			sql.As(t.C(entlike.FieldCreatedAt), "created_at"),
			sql.As(t.C(entlike.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t := sql.Table(entappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entlike.FieldAppGoodID),
			t.C(entappgood.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappgood.FieldGoodName), "good_name"),
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

func (h *Handler) GetLike(ctx context.Context) (*npool.Like, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLike(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetLikes(ctx context.Context) ([]*npool.Like, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryLikes(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryLikes(cli)
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

		return handler.scan(ctx)
	})
	if err != nil {
		return nil, 0, err
	}

	return handler.infos, handler.total, nil
}
