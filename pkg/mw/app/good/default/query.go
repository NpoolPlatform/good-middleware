package appdefaultgood

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	appdefaultgoodcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/default"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappdefaultgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appdefaultgood"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	entgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.AppDefaultGoodSelect
	stmCount  *ent.AppDefaultGoodSelect
	infos     []*npool.Default
	total     uint32
}

func (h *queryHandler) selectDefault(stm *ent.AppDefaultGoodQuery) *ent.AppDefaultGoodSelect {
	return stm.Select(entappdefaultgood.FieldID)
}

func (h *queryHandler) queryDefault(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.AppDefaultGood.Query().Where(entappdefaultgood.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entappdefaultgood.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entappdefaultgood.EntID(*h.EntID))
	}
	h.stmSelect = h.selectDefault(stm)
	return nil
}

func (h *queryHandler) queryDefaults(cli *ent.Client) (*ent.AppDefaultGoodSelect, error) {
	stm, err := appdefaultgoodcrud.SetQueryConds(cli.AppDefaultGood.Query(), h.Conds)
	if err != nil {
		return nil, err
	}
	return h.selectDefault(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t := sql.Table(entappdefaultgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappdefaultgood.FieldID),
			t.C(entappdefaultgood.FieldID),
		).
		AppendSelect(
			sql.As(t.C(entappdefaultgood.FieldEntID), "ent_id"),
			sql.As(t.C(entappdefaultgood.FieldAppID), "app_id"),
			sql.As(t.C(entappdefaultgood.FieldGoodID), "good_id"),
			sql.As(t.C(entappdefaultgood.FieldAppGoodID), "app_good_id"),
			sql.As(t.C(entappdefaultgood.FieldCoinTypeID), "coin_type_id"),
			sql.As(t.C(entappdefaultgood.FieldCreatedAt), "created_at"),
			sql.As(t.C(entappdefaultgood.FieldUpdatedAt), "updated_at"),
		)
}

func (h *queryHandler) queryJoinGood(s *sql.Selector) {
	t := sql.Table(entgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappdefaultgood.FieldGoodID),
			t.C(entgood.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entgood.FieldTitle), "good_name"),
		)
}

func (h *queryHandler) queryJoinAppGood(s *sql.Selector) {
	t := sql.Table(entappgood.Table)
	s.LeftJoin(t).
		On(
			s.C(entappdefaultgood.FieldAppGoodID),
			t.C(entappgood.FieldEntID),
		).
		AppendSelect(
			sql.As(t.C(entappgood.FieldGoodName), "app_good_name"),
		)
}

func (h *queryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinGood(s)
		h.queryJoinAppGood(s)
	})
	if h.stmCount == nil {
		return
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		h.queryJoinGood(s)
		h.queryJoinAppGood(s)
	})
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetDefault(ctx context.Context) (*npool.Default, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryDefault(cli); err != nil {
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

func (h *Handler) GetDefaults(ctx context.Context) ([]*npool.Default, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}

	var err error

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryDefaults(cli)
		if err != nil {
			return err
		}
		handler.stmCount, err = handler.queryDefaults(cli)
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
