package coin

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"

	goodcoincrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entgoodcoin "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"

	"github.com/google/uuid"
)

type queryHandler struct {
	*Handler
	stmCount  *ent.GoodCoinSelect
	stmSelect *ent.GoodCoinSelect
	infos     []*npool.GoodCoin
	total     uint32
}

func (h *queryHandler) selectGoodCoin(stm *ent.GoodCoinQuery) *ent.GoodCoinSelect {
	return stm.Select(entgoodcoin.FieldID)
}

func (h *queryHandler) queryGoodCoin(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid goodcoinid")
	}
	stm := cli.GoodCoin.
		Query().
		Where(
			entgoodcoin.DeletedAt(0),
		)
	if h.ID != nil {
		stm.Where(entgoodcoin.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entgoodcoin.EntID(*h.EntID))
	}
	h.stmSelect = h.selectGoodCoin(stm)
	return nil
}

func (h *queryHandler) queryGoodCoins(cli *ent.Client) (*ent.GoodCoinSelect, error) {
	stm, err := goodcoincrud.SetQueryConds(cli.GoodCoin.Query(), h.GoodCoinConds)
	if err != nil {
		return nil, err
	}
	return h.selectGoodCoin(stm), nil
}

func (h *queryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entgoodcoin.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodcoin.FieldID),
			t1.C(entgoodcoin.FieldID),
		).
		AppendSelect(
			t1.C(entgoodcoin.FieldEntID),
			t1.C(entgoodcoin.FieldGoodID),
			t1.C(entgoodcoin.FieldCoinTypeID),
			t1.C(entgoodcoin.FieldMain),
			t1.C(entgoodcoin.FieldIndex),
			t1.C(entgoodcoin.FieldCreatedAt),
			t1.C(entgoodcoin.FieldUpdatedAt),
		)
}

func (h *queryHandler) queryJoinGoodBase(s *sql.Selector) error {
	t1 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodcoin.FieldGoodID),
			t1.C(entgoodbase.FieldEntID),
		)
	if h.GoodBaseConds != nil && h.GoodBaseConds.EntID != nil {
		uid, ok := h.GoodBaseConds.EntID.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodid")
		}
		s.OnP(sql.EQ(t1.C(entgoodbase.FieldEntID), uid))
	}
	if h.GoodBaseConds != nil && h.GoodBaseConds.EntIDs != nil {
		uid, ok := h.GoodBaseConds.EntIDs.Val.(uuid.UUID)
		if !ok {
			return fmt.Errorf("invalid goodids")
		}
		s.OnP(sql.In(t1.C(entgoodbase.FieldEntID), uid))
	}
	s.AppendSelect(
		sql.As(t1.C(entgoodbase.FieldName), "good_name"),
		t1.C(entgoodbase.FieldGoodType),
	)
	return nil
}

func (h *queryHandler) queryJoin() (err error) {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		err = h.queryJoinGoodBase(s)
	})
	if err != nil {
		return err
	}
	if h.stmCount == nil {
		return nil
	}
	h.stmCount.Modify(func(s *sql.Selector) {
		err = h.queryJoinGoodBase(s)
	})
	return err
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *queryHandler) formalize() {
	for _, info := range h.infos {
		info.GoodType = types.GoodType(types.GoodType_value[info.GoodTypeStr])
	}
}

func (h *Handler) GetGoodCoin(ctx context.Context) (*npool.GoodCoin, error) {
	handler := &queryHandler{
		Handler: h,
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryGoodCoin(cli); err != nil {
			return nil
		}
		if err := handler.queryJoin(); err != nil {
			return nil
		}
		handler.stmSelect.
			Offset(0).
			Limit(2)
		return handler.scan(ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	if len(handler.infos) == 0 {
		return nil, fmt.Errorf("invalid goodcoin")
	}

	handler.formalize()

	return handler.infos[0], nil
}

func (h *Handler) GetGoodCoins(ctx context.Context) ([]*npool.GoodCoin, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	var err error
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		handler.stmSelect, err = handler.queryGoodCoins(cli)
		if err != nil {
			return nil
		}
		handler.stmCount, err = handler.queryGoodCoins(cli)
		if err != nil {
			return nil
		}
		if err := handler.queryJoin(); err != nil {
			return nil
		}

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
