package displayname

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgooddisplaynamecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/display/name"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgooddisplayname "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplayname"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppGoodDisplayNameSelect
	stmCount  *ent.AppGoodDisplayNameSelect
	infos     []*npool.DisplayName
	total     uint32
}

func (h *queryHandler) selectDisplayName(stm *ent.AppGoodDisplayNameQuery) *ent.AppGoodDisplayNameSelect {
	return stm.Select(entappgooddisplayname.FieldID)
}

func (h *queryHandler) queryDisplayName(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppGoodDisplayName.Query().Where(entappgooddisplayname.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgooddisplayname.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgooddisplayname.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDisplayName(stm)
	return nil
}

func (h *queryHandler) queryDisplayNames(cli *ent.Client) (*ent.AppGoodDisplayNameSelect, error) {
	stm, err := appgooddisplaynamecrud.SetQueryConds(cli.AppGoodDisplayName.Query(), h.DisplayNameConds)
	if err != nil {
		return nil, err
	}
	return h.selectDisplayName(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgooddisplayname.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgooddisplayname.FieldID),
			t.C(entappgooddisplayname.FieldID),
		).
		AppendSelect(
			t.C(entappgooddisplayname.FieldEntID),
			t.C(entappgooddisplayname.FieldAppGoodID),
			t.C(entappgooddisplayname.FieldName),
			t.C(entappgooddisplayname.FieldIndex),
			t.C(entappgooddisplayname.FieldCreatedAt),
			t.C(entappgooddisplayname.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgooddisplayname.FieldAppGoodID),
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

func (h *Handler) GetDisplayName(ctx context.Context) (*npool.DisplayName, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDisplayName(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid appgooddisplayname")
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetDisplayNames(ctx context.Context) ([]*npool.DisplayName, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryDisplayNames(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryDisplayNames(cli)
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
