package pledge

import (
	"entgo.io/ent/dialect/sql"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entgoodbase "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodbase"
	entgoodcoin "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodcoin"
	entgoodreward "github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodreward"
	entpledge "github.com/NpoolPlatform/good-middleware/pkg/db/ent/pledge"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type baseQueryHandler struct {
	*Handler
	stmSelect *ent.GoodBaseSelect
}

func (h *baseQueryHandler) selectGoodBase(stm *ent.GoodBaseQuery) *ent.GoodBaseSelect {
	return stm.Select(entgoodbase.FieldID)
}

func (h *baseQueryHandler) queryGoodBase(cli *ent.Client) error {
	if h.GoodID == nil {
		return wlog.Errorf("invalid goodid")
	}
	h.stmSelect = h.selectGoodBase(
		cli.GoodBase.
			Query().
			Where(
				entgoodbase.DeletedAt(0),
				entgoodbase.EntID(*h.GoodID),
				entgoodbase.Or(
					entgoodbase.GoodType(types.GoodType_Pledge.String()),
				),
			),
	)
	return nil
}

func (h *baseQueryHandler) queryGoodBases(cli *ent.Client) (*ent.GoodBaseSelect, error) {
	stm, err := goodbasecrud.SetQueryConds(cli.GoodBase.Query(), h.GoodBaseConds)
	if err != nil {
		return nil, wlog.WrapError(err)
	}
	stm.Where(
		entgoodbase.Or(
			entgoodbase.GoodType(types.GoodType_Pledge.String()),
		),
	)
	return h.selectGoodBase(stm), nil
}

func (h *baseQueryHandler) queryJoinMyself(s *sql.Selector) {
	t1 := sql.Table(entgoodbase.Table)
	s.LeftJoin(t1).
		On(
			s.C(entgoodbase.FieldID),
			t1.C(entgoodbase.FieldID),
		).
		AppendSelect(
			t1.C(entgoodbase.FieldGoodType),
			t1.C(entgoodbase.FieldBenefitType),
			t1.C(entgoodbase.FieldName),
			t1.C(entgoodbase.FieldServiceStartAt),
			t1.C(entgoodbase.FieldStartMode),
			t1.C(entgoodbase.FieldTestOnly),
			t1.C(entgoodbase.FieldBenefitIntervalHours),
			t1.C(entgoodbase.FieldPurchasable),
			t1.C(entgoodbase.FieldOnline),
			t1.C(entgoodbase.FieldState),
			t1.C(entgoodbase.FieldCreatedAt),
			t1.C(entgoodbase.FieldUpdatedAt),
		)
}

func (h *baseQueryHandler) queryJoinPledge(s *sql.Selector) error {
	t1 := sql.Table(entpledge.Table)

	s.Join(t1).
		On(
			s.C(entgoodbase.FieldEntID),
			t1.C(entpledge.FieldGoodID),
		).
		OnP(
			sql.EQ(t1.C(entpledge.FieldDeletedAt), 0),
		)
	if h.PledgeConds.ID != nil {
		u, ok := h.PledgeConds.ID.Val.(uint32)
		if !ok {
			return wlog.Errorf("invalid id")
		}
		s.OnP(sql.EQ(t1.C(entpledge.FieldID), u))
	}
	if h.PledgeConds.IDs != nil {
		ids, ok := h.PledgeConds.IDs.Val.([]uint32)
		if !ok {
			return wlog.Errorf("invalid ids")
		}
		s.OnP(sql.In(t1.C(entpledge.FieldID), func() (_ids []interface{}) {
			for _, id := range ids {
				_ids = append(_ids, interface{}(id))
			}
			return
		}()...))
	}
	if h.PledgeConds.EntID != nil {
		uid, ok := h.PledgeConds.EntID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entid")
		}
		s.OnP(sql.EQ(t1.C(entpledge.FieldEntID), uid))
	}
	if h.PledgeConds.EntIDs != nil {
		uids, ok := h.PledgeConds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid entids")
		}
		s.OnP(sql.In(t1.C(entpledge.FieldEntID), func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()...))
	}
	if h.PledgeConds.ContractState != nil {
		state, ok := h.PledgeConds.ContractState.Val.(types.ContractState)
		if !ok {
			return wlog.Errorf("invalid contractstate")
		}
		s.OnP(
			sql.EQ(t1.C(entpledge.FieldContractState), state.String()),
		)
	}
	if h.PledgeConds.ContractStates != nil {
		states, ok := h.PledgeConds.ContractStates.Val.([]types.ContractState)
		if !ok {
			return wlog.Errorf("invalid contractstates")
		}
		s.OnP(
			sql.In(t1.C(entpledge.FieldContractState), func() (_states []interface{}) {
				for _, state := range states {
					_states = append(_states, interface{}(state.String()))
				}
				return
			}()...),
		)
	}

	s.AppendSelect(
		t1.C(entpledge.FieldID),
		t1.C(entpledge.FieldEntID),
		t1.C(entpledge.FieldGoodID),
		t1.C(entpledge.FieldContractCodeURL),
		t1.C(entpledge.FieldContractCodeBranch),
		t1.C(entpledge.FieldContractState),
	)
	return nil
}

func (h *baseQueryHandler) queryJoinReward(s *sql.Selector) {
	t := sql.Table(entgoodreward.Table)
	s.Join(t).
		On(
			s.C(entgoodbase.FieldEntID),
			t.C(entgoodreward.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entgoodreward.FieldDeletedAt), 0),
		)
	if h.RewardConds.RewardState != nil {
		s.OnP(
			sql.EQ(t.C(entgoodreward.FieldRewardState), h.RewardConds.RewardState.Val.(types.BenefitState).String()),
		)
	}
	if h.RewardConds.RewardAt != nil {
		switch h.RewardConds.RewardAt.Op {
		case cruder.EQ:
			s.OnP(sql.EQ(t.C(entgoodreward.FieldLastRewardAt), h.RewardConds.RewardAt.Val))
		case cruder.NEQ:
			s.OnP(sql.NEQ(t.C(entgoodreward.FieldLastRewardAt), h.RewardConds.RewardAt.Val))
		}
	}
	s.AppendSelect(
		t.C(entgoodreward.FieldRewardState),
		t.C(entgoodreward.FieldLastRewardAt),
	)
}

func (h *baseQueryHandler) queryJoinGoodCoin(s *sql.Selector) error {
	t := sql.Table(entgoodcoin.Table)
	s.LeftJoin(t).
		On(
			s.C(entgoodbase.FieldEntID),
			t.C(entgoodcoin.FieldGoodID),
		).
		OnP(
			sql.EQ(t.C(entgoodcoin.FieldDeletedAt), 0),
		).
		Distinct()
	if h.GoodCoinConds.CoinTypeID != nil {
		id, ok := h.GoodCoinConds.CoinTypeID.Val.(uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeid")
		}
		s.OnP(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
		s.Where(
			sql.EQ(t.C(entgoodcoin.FieldCoinTypeID), id),
		)
	}
	if h.GoodCoinConds.CoinTypeIDs != nil {
		uids, ok := h.GoodCoinConds.CoinTypeIDs.Val.([]uuid.UUID)
		if !ok {
			return wlog.Errorf("invalid cointypeids")
		}
		_uids := func() (_uids []interface{}) {
			for _, uid := range uids {
				_uids = append(_uids, interface{}(uid))
			}
			return
		}()
		s.OnP(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
		s.Where(
			sql.In(t.C(entgoodcoin.FieldCoinTypeID), _uids...),
		)
	}
	return nil
}

func (h *baseQueryHandler) queryJoin() {
	h.stmSelect.Modify(func(s *sql.Selector) {
		h.queryJoinMyself(s)
		h.queryJoinReward(s)
		if err := h.queryJoinGoodCoin(s); err != nil {
			logger.Sugar().Errorw("queryJoinGoodCoin", "Error", err)
		}
		if err := h.queryJoinPledge(s); err != nil {
			logger.Sugar().Errorw("queryJoinPledge", "Error", err)
		}
	})
}
