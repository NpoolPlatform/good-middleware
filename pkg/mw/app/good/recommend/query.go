package recommend

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	recommendcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/recommend"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entrecommend "github.com/NpoolPlatform/good-middleware/pkg/db/ent/recommend"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"

	"github.com/shopspring/decimal"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.RecommendSelect
	stmCount  *ent.RecommendSelect
	infos     []*npool.Recommend
	total     uint32
}

func (h *queryHandler) selectRecommend(stm *ent.RecommendQuery) *ent.RecommendSelect {
	return stm.Select(entrecommend.FieldID)
}

func (h *queryHandler) queryRecommend(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.Recommend.Query().Where(entrecommend.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entrecommend.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entrecommend.EntID(*h.EntID))
	}
	h.stmSelect = h.selectRecommend(stm)
	return nil
}

func (h *queryHandler) queryRecommends(cli *ent.Client) (*ent.RecommendSelect, error) {
	stm, err := recommendcrud.SetQueryConds(cli.Recommend.Query(), h.RecommendConds)
	if err != nil {
		return nil, err
	}
	return h.selectRecommend(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entrecommend.Table)
	s.LeftJoin(t).
		On(
			s.C(entrecommend.FieldID),
			t.C(entrecommend.FieldID),
		).
		AppendSelect(
			t.C(entrecommend.FieldEntID),
			t.C(entrecommend.FieldRecommenderID),
			t.C(entrecommend.FieldAppGoodID),
			t.C(entrecommend.FieldMessage),
			t.C(entrecommend.FieldRecommendIndex),
			t.C(entrecommend.FieldHide),
			t.C(entrecommend.FieldHideReason),
			t.C(entrecommend.FieldCreatedAt),
			t.C(entrecommend.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t).
		On(
			s.C(entrecommend.FieldAppGoodID),
			t.C(entappgoodbase.FieldEntID),
		).
		AppendSelect(
			t.C(entappgoodbase.FieldAppID),
			sql.As(t.C(entappgoodbase.FieldName), "good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.RecommendIndex = func() string { amount, _ := decimal.NewFromString(info.RecommendIndex); return amount.String() }()
		info.HideReason = types.GoodCommentHideReason(types.GoodCommentHideReason_value[info.HideReasonStr])
	}
}

func (h *Handler) GetRecommend(ctx context.Context) (*npool.Recommend, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryRecommend(cli); err != nil {
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

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetRecommends(ctx context.Context) ([]*npool.Recommend, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryRecommends(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryRecommends(cli)
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

	handler.formalize()

	return handler.infos, handler.total, nil
}
