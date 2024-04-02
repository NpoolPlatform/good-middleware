package comment

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	commentcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/comment"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entcomment "github.com/NpoolPlatform/good-middleware/pkg/db/ent/comment"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/comment"
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

func (h *queryHandler) queryComment(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Comment.Query().Where(entcomment.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entcomment.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entcomment.EntID(*h.EntID))
	}
	h.stmSelect = h.selectComment(stm)
	return nil
}

func (h *queryHandler) queryComments(cli *ent.Client) (*ent.CommentSelect, error) {
	stm, err := commentcrud.SetQueryConds(cli.Comment.Query(), h.CommentConds)
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
			sql.As(t.C(entcomment.FieldEntID), "ent_id"),
			sql.As(t.C(entcomment.FieldUserID), "user_id"),
			sql.As(t.C(entcomment.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entcomment.FieldOrderID), "order_id"),
			sql.As(t.C(entcomment.FieldContent), "content"),
			sql.As(t.C(entcomment.FieldReplyToID), "reply_to_id"),
			sql.As(t.C(entcomment.FieldAnonymous), "anonymous"),
			sql.As(t.C(entcomment.FieldPurchasedUser), "purchased_user"),
			sql.As(t.C(entcomment.FieldTrialUser), "trial_user"),
			sql.As(t.C(entcomment.FieldOrderFirstComment), "order_first_comment"),
			sql.As(t.C(entcomment.FieldCreatedAt), "created_at"),
			sql.As(t.C(entcomment.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinAppGoodBase(s *sql.Selector) {
	t := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entcomment.FieldAppGoodID),
			t.C(entappgoodbase.FieldEntID),
		).
		AppendSelect(
			t.C(entappgoodbase.FieldAppID),
			t.C(entappgoodbase.FieldGoodID),
			sql.As(t.C(entappgoodbase.FieldName), "good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinAppGoodBase(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinAppGoodBase(s)
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
		if err := handler.queryComment(cli); err != nil {
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
