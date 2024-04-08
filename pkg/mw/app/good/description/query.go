package description

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appgooddescriptioncrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/description"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgoodbase"
	entappgooddescription "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddescription"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppGoodDescriptionSelect
	stmCount  *ent.AppGoodDescriptionSelect
	infos     []*npool.Description
	total     uint32
}

func (h *queryHandler) selectDescription(stm *ent.AppGoodDescriptionQuery) *ent.AppGoodDescriptionSelect {
	return stm.Select(entappgooddescription.FieldID)
}

func (h *queryHandler) queryDescription(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppGoodDescription.Query().Where(entappgooddescription.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappgooddescription.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappgooddescription.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDescription(stm)
	return nil
}

func (h *queryHandler) queryDescriptions(cli *ent.Client) (*ent.AppGoodDescriptionSelect, error) {
	stm, err := appgooddescriptioncrud.SetQueryConds(cli.AppGoodDescription.Query(), h.DescriptionConds)
	if err != nil {
		return nil, err
	}
	return h.selectDescription(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappgooddescription.Table)
	s.LeftJoin(t).
		On(
			s.C(entappgooddescription.FieldID),
			t.C(entappgooddescription.FieldID),
		).
		AppendSelect(
			t.C(entappgooddescription.FieldEntID),
			t.C(entappgooddescription.FieldAppGoodID),
			t.C(entappgooddescription.FieldDescription),
			t.C(entappgooddescription.FieldIndex),
			t.C(entappgooddescription.FieldCreatedAt),
			t.C(entappgooddescription.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t1 := sql.Table(entappgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entappgooddescription.FieldAppGoodID),
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

func (h *Handler) GetDescription(ctx context.Context) (*npool.Description, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDescription(cli); err != nil {
			return err
		}
		handler.queryJoin()
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid appgooddescription")
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetDescriptions(ctx context.Context) ([]*npool.Description, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryDescriptions(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryDescriptions(cli)
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
