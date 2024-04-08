package displaycolor

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgooddisplaycolorcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/display/color"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgooddisplaycolor "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplaycolor"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppGoodDisplayColorSelect
	stmCount  *ent.AppGoodDisplayColorSelect
	infos     []*npool.DisplayColor
	total     uint32
}

func (h *queryHandler) selectDisplayColor(stm *ent.AppGoodDisplayColorQuery) *ent.AppGoodDisplayColorSelect {
	return stm.Select(entappgooddisplaycolor.FieldID)
}

func (h *queryHandler) queryDisplayColor(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppGoodDisplayColor.Query().Where(entappgooddisplaycolor.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgooddisplaycolor.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgooddisplaycolor.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDisplayColor(stm)
	return nil
}

func (h *queryHandler) queryDisplayColors(cli *ent.Client) (*ent.AppGoodDisplayColorSelect, error) {
	stm, err := appgooddisplaycolorcrud.SetQueryConds(cli.AppGoodDisplayColor.Query(), h.DisplayColorConds)
	if err != nil {
		return nil, err
	}
	return h.selectDisplayColor(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgooddisplaycolor.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgooddisplaycolor.FieldID),
			t.C(entappgooddisplaycolor.FieldID),
		).
		AppendSelect(
			t.C(entappgooddisplaycolor.FieldEntID),
			t.C(entappgooddisplaycolor.FieldAppGoodID),
			t.C(entappgooddisplaycolor.FieldColor),
			t.C(entappgooddisplaycolor.FieldIndex),
			t.C(entappgooddisplaycolor.FieldCreatedAt),
			t.C(entappgooddisplaycolor.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgooddisplaycolor.FieldAppGoodID),
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

func (h *Handler) GetDisplayColor(ctx context.Context) (*npool.DisplayColor, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDisplayColor(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid appgooddisplaycolor")
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetDisplayColors(ctx context.Context) ([]*npool.DisplayColor, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryDisplayColors(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryDisplayColors(cli)
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
