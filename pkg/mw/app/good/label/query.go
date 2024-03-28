package label

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgoodlabelcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/label"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgoodlabel "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodlabel"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppGoodLabelSelect
	stmCount  *ent.AppGoodLabelSelect
	infos     []*npool.Label
	total     uint32
}

func (h *queryHandler) selectLabel(stm *ent.AppGoodLabelQuery) *ent.AppGoodLabelSelect {
	return stm.Select(entappgoodlabel.FieldID)
}

func (h *queryHandler) queryLabel(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppGoodLabel.Query().Where(entappgoodlabel.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgoodlabel.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgoodlabel.EntID(*h.EntID))
	}
	h.stmSelect = h.selectLabel(stm)
	return nil
}

func (h *queryHandler) queryLabels(cli *ent.Client) (*ent.AppGoodLabelSelect, error) {
	stm, err := appgoodlabelcrud.SetQueryConds(cli.AppGoodLabel.Query(), h.LabelConds)
	if err != nil {
		return nil, err
	}
	return h.selectLabel(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgoodlabel.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgoodlabel.FieldID),
			t.C(entappgoodlabel.FieldID),
		).
		AppendSelect(
			t.C(entappgoodlabel.FieldEntID),
			t.C(entappgoodlabel.FieldAppGoodID),
			t.C(entappgoodlabel.FieldIcon),
			t.C(entappgoodlabel.FieldIconBgColor),
			t.C(entappgoodlabel.FieldLabel),
			t.C(entappgoodlabel.FieldLabelBgColor),
			t.C(entappgoodlabel.FieldIndex),
			t.C(entappgoodlabel.FieldCreatedAt),
			t.C(entappgoodlabel.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgoodlabel.FieldAppGoodID),
			t1.C(entappgoodbase.FieldEntID),
		).
		AppendSelect(
			sql.As(t1.C(entappgoodbase.FieldAppID), "app_id"),
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

func (h *Handler) GetLabel(ctx context.Context) (*npool.Label, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryLabel(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid appgoodlabel")
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	return handler.infos[0], nil
}

func (h *Handler) GetLabels(ctx context.Context) ([]*npool.Label, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryLabels(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryLabels(cli)
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

	return handler.infos, handler.total, nil
}
