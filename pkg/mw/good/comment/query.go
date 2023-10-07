package comment

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	commentcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/comment"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entcomment "github.com/NpoolPlatform/good-middleware/pkg/db/ent/comment"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/comment"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.CommentSelect
	stmCount  *ent.CommentSelect
	infos     []*npool.Comment
	total     uint32
}

func (h *queryHandler) selectComment(stm *ent.CommentQuery) *ent.CommentSelect {
	return stm.Select(entcomment.FieldID)
}

func (h *queryHandler) queryComment(cli *ent.Client) {
	h.stmSelect = h.selectComment(
		cli.Comment.
			Query().
			Where(
				entcomment.ID(*h.ID),
				entcomment.DeletedAt(0),
			),
	)
}

func (h *queryHandler) queryComments(cli *ent.Client) (*ent.CommentSelect, error) {
	stm, err := commentcrud.SetQueryConds(cli.Comment.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectComment(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entcomment.Table)
	s.LeftJoin(t).
		On(
			s.C(entcomment.FieldID),
			t.C(entcomment.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entcomment.FieldAppID), "app_id"),
			sql.As(t.C(entcomment.FieldUserID), "user_id"),
			sql.As(t.C(entcomment.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entcomment.FieldOrderID), "order_id"),
			sql.As(t.C(entcomment.FieldContent), "content"),
			sql.As(t.C(entcomment.FieldReplyToID), "reply_to_id"),
			sql.As(t.C(entcomment.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcomment.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t := sql.Table(entappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entcomment.FieldAppGoodID),
			t.C(entappgood.FieldID),
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

func (h *Handler) GetComment(ctx context.Context) (*npool.Comment, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.queryComment(cli)
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

func (h *Handler) GetComments(ctx context.Context) ([]*npool.Comment, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryComments(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryComments(cli)
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
