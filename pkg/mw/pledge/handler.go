package pledge

import (
	"context"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"
	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	goodcoincrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin"
	goodcoinrewardcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/coin/reward"
	goodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/goodbase"
	goodrewardcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/good/reward"
	pledgecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/pledge"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
	goodcoinrewardmwpb "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin/reward"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID *uint32
	pledgecrud.Req
	GoodBaseReq    *goodbasecrud.Req
	Rollback       *bool
	RewardReq      *goodrewardcrud.Req
	PledgeConds    *pledgecrud.Conds
	GoodBaseConds  *goodbasecrud.Conds
	GoodCoinConds  *goodcoincrud.Conds
	RewardConds    *goodrewardcrud.Conds
	CoinRewardReqs []*goodcoinrewardcrud.Req
	Offset         int32
	Limit          int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		Req:           pledgecrud.Req{},
		GoodBaseReq:   &goodbasecrud.Req{},
		RewardReq:     &goodrewardcrud.Req{},
		PledgeConds:   &pledgecrud.Conds{},
		GoodBaseConds: &goodbasecrud.Conds{},
		GoodCoinConds: &goodcoincrud.Conds{},
		RewardConds:   &goodrewardcrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, wlog.WrapError(err)
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return wlog.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid entid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.EntID = &id
		return nil
	}
}

func WithGoodID(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid goodid")
			}
			return nil
		}
		id, err := uuid.Parse(*s)
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodID = &id
		h.GoodBaseReq.EntID = &id
		h.RewardReq.GoodID = &id
		return nil
	}
}

func WithContractState(e *types.ContractState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid contractstate")
			}
			return nil
		}
		switch *e {
		case types.ContractState_ContractWaitDeployment:
		case types.ContractState_ContractInDeployment:
		case types.ContractState_ContractDeploymentSuccess:
		case types.ContractState_ContractDeploymentFail:
		default:
			return wlog.Errorf("invalid contractstate")
		}
		h.ContractState = e
		return nil
	}
}

func WithGoodType(e *types.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case types.GoodType_Pledge:
		default:
			return wlog.Errorf("invalid goodtype")
		}
		h.GoodBaseReq.GoodType = e
		return nil
	}
}

func WithBenefitType(e *types.BenefitType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid benefittype")
			}
			return nil
		}
		switch *e {
		case types.BenefitType_BenefitTypeContract:
		case types.BenefitType_BenefitTypePlatform:
		case types.BenefitType_BenefitTypePool:
		case types.BenefitType_BenefitTypeOffline:
		default:
			return wlog.Errorf("invalid benefittype")
		}
		h.GoodBaseReq.BenefitType = e
		return nil
	}
}

func WithName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid name")
			}
			return nil
		}
		const leastNameLen = 3
		if len(*s) < leastNameLen {
			return wlog.Errorf("invalid name")
		}
		h.GoodBaseReq.Name = s
		return nil
	}
}

func WithContractCodeURL(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid contractcodeurl")
			}
			return nil
		}
		const leastNameLen = 3
		if len(*s) < leastNameLen {
			return wlog.Errorf("invalid contractcodeurl")
		}
		h.ContractCodeURL = s
		return nil
	}
}

func WithContractCodeBranch(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return wlog.Errorf("invalid contractcodebranch")
			}
			return nil
		}
		const leastNameLen = 3
		if len(*s) < leastNameLen {
			return wlog.Errorf("invalid contractcodebranch")
		}
		h.ContractCodeBranch = s
		return nil
	}
}

func WithServiceStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.ServiceStartAt = n
		return nil
	}
}

func WithStartMode(e *types.GoodStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid startmode")
			}
			return nil
		}
		switch *e {
		case types.GoodStartMode_GoodStartModeTBD:
		case types.GoodStartMode_GoodStartModeConfirmed:
		case types.GoodStartMode_GoodStartModeNextDay:
		case types.GoodStartMode_GoodStartModeInstantly:
		case types.GoodStartMode_GoodStartModePreset:
		default:
			return wlog.Errorf("invalid startmode")
		}
		h.GoodBaseReq.StartMode = e
		return nil
	}
}

func WithTestOnly(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.TestOnly = b
		return nil
	}
}

func WithBenefitIntervalHours(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.BenefitIntervalHours = n
		return nil
	}
}

func WithPurchasable(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.Purchasable = b
		return nil
	}
}

func WithOnline(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.GoodBaseReq.Online = b
		return nil
	}
}

func WithRewardState(e *types.BenefitState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid rewardstate")
			}
			return nil
		}
		switch *e {
		case types.BenefitState_BenefitWait:
		case types.BenefitState_BenefitTransferring:
		case types.BenefitState_BenefitBookKeeping:
		case types.BenefitState_BenefitUserBookKeeping:
		case types.BenefitState_BenefitSimulateBookKeeping:
		case types.BenefitState_BenefitDone:
		case types.BenefitState_BenefitFail:
		default:
			return wlog.Errorf("invalid rewardstate")
		}
		h.RewardReq.RewardState = e
		return nil
	}
}

func WithRewardAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return wlog.Errorf("invalid rewardat")
			}
			return nil
		}
		h.RewardReq.LastRewardAt = n
		return nil
	}
}

func WithRewards(rewards []*goodcoinrewardmwpb.RewardReq, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		coinTypeIDs := map[uuid.UUID]struct{}{}
		rewardTIDs := map[uuid.UUID]struct{}{}

		for _, reward := range rewards {
			_reward := &goodcoinrewardcrud.Req{}

			if reward.EntID != nil {
				_entID, err := uuid.Parse(*reward.EntID)
				if err != nil {
					return wlog.WrapError(err)
				}
				_reward.EntID = &_entID
			}

			_coinTypeID, err := uuid.Parse(reward.GetCoinTypeID())
			if err != nil {
				return wlog.WrapError(err)
			}
			if _, ok := coinTypeIDs[_coinTypeID]; ok {
				return wlog.Errorf("invalid cointypeid")
			}
			coinTypeIDs[_coinTypeID] = struct{}{}
			_reward.CoinTypeID = &_coinTypeID

			if reward.RewardTID != nil {
				_tid, err := uuid.Parse(*reward.RewardTID)
				if err != nil {
					return wlog.WrapError(err)
				}
				if _tid != uuid.Nil {
					if _, ok := rewardTIDs[_tid]; ok {
						return wlog.Errorf("invalid rewardid")
					}
					rewardTIDs[_tid] = struct{}{}
				}
				_reward.RewardTID = &_tid
			}

			if reward.RewardAmount != nil {
				_rewardAmount, err := decimal.NewFromString(*reward.RewardAmount)
				if err != nil {
					return wlog.WrapError(err)
				}
				_reward.LastRewardAmount = &_rewardAmount
			}

			if reward.NextRewardStartAmount != nil {
				_startAmount, err := decimal.NewFromString(*reward.NextRewardStartAmount)
				if err != nil {
					return wlog.WrapError(err)
				}
				_reward.NextRewardStartAmount = &_startAmount
			}

			h.CoinRewardReqs = append(h.CoinRewardReqs, _reward)
		}
		return nil
	}
}

func WithState(e *types.GoodState, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid state")
			}
			return nil
		}
		switch *e {
		case types.GoodState_GoodStatePreWait:
		case types.GoodState_GoodStateWait:
		case types.GoodState_GoodStateCreateGoodUser:
		case types.GoodState_GoodStateCheckHashRate:
		case types.GoodState_GoodStateReady:
		case types.GoodState_GoodStateFail:
		default:
			return wlog.Errorf("invalid state")
		}
		h.GoodBaseReq.State = e
		return nil
	}
}

func WithRollback(e *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return wlog.Errorf("invalid rollback")
			}
			return nil
		}

		h.Rollback = e
		return nil
	}
}

func (h *Handler) withPledgeConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.PledgeConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PledgeConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.PledgeConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.PledgeConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.ContractState != nil {
		h.PledgeConds.ContractState = &cruder.Cond{
			Op:  conds.GetContractState().GetOp(),
			Val: types.ContractState(conds.GetContractState().GetValue()),
		}
	}
	return nil
}

func (h *Handler) withGoodBaseConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodBaseConds.EntID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodBaseConds.EntIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.GoodType != nil {
		h.GoodBaseConds.GoodType = &cruder.Cond{
			Op:  conds.GetGoodType().GetOp(),
			Val: types.GoodType(conds.GetGoodType().GetValue()),
		}
	}
	if conds.GoodTypes != nil {
		es := []types.GoodType{}
		for _, e := range conds.GetGoodTypes().GetValue() {
			es = append(es, types.GoodType(e))
		}
		h.GoodBaseConds.GoodTypes = &cruder.Cond{
			Op:  conds.GetGoodTypes().GetOp(),
			Val: es,
		}
	}
	if conds.State != nil {
		h.GoodBaseConds.State = &cruder.Cond{
			Op:  conds.GetState().GetOp(),
			Val: types.GoodState(conds.GetState().GetValue()),
		}
	}
	return nil
}

func (h *Handler) withGoodCoinConds(conds *npool.Conds) error {
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodCoinConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	if conds.GoodIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodCoinConds.GoodIDs = &cruder.Cond{
			Op:  conds.GetGoodIDs().GetOp(),
			Val: ids,
		}
	}
	if conds.CoinTypeID != nil {
		id, err := uuid.Parse(conds.GetCoinTypeID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.GoodCoinConds.CoinTypeID = &cruder.Cond{
			Op:  conds.GetCoinTypeID().GetOp(),
			Val: id,
		}
	}
	if conds.CoinTypeIDs != nil {
		ids := []uuid.UUID{}
		for _, id := range conds.GetCoinTypeIDs().GetValue() {
			_id, err := uuid.Parse(id)
			if err != nil {
				return wlog.WrapError(err)
			}
			ids = append(ids, _id)
		}
		h.GoodCoinConds.CoinTypeIDs = &cruder.Cond{
			Op:  conds.GetCoinTypeIDs().GetOp(),
			Val: ids,
		}
	}
	return nil
}

func (h *Handler) withRewardConds(conds *npool.Conds) error {
	if conds.RewardState != nil {
		h.RewardConds.RewardState = &cruder.Cond{
			Op:  conds.GetRewardState().GetOp(),
			Val: types.BenefitState(conds.GetRewardState().GetValue()),
		}
	}
	if conds.RewardAt != nil {
		h.RewardConds.RewardAt = &cruder.Cond{
			Op:  conds.GetRewardAt().GetOp(),
			Val: conds.GetRewardAt().GetValue(),
		}
	}
	if conds.GoodID != nil {
		id, err := uuid.Parse(conds.GetGoodID().GetValue())
		if err != nil {
			return wlog.WrapError(err)
		}
		h.RewardConds.GoodID = &cruder.Cond{
			Op:  conds.GetGoodID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withPledgeConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withGoodCoinConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		if err := h.withRewardConds(conds); err != nil {
			return wlog.WrapError(err)
		}
		return h.withGoodBaseConds(conds)
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
