package goodstm

import (
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

var StateMap = map[types.GoodState]types.MiningGoodStockState{
	types.GoodState_DefaultGoodState:        types.MiningGoodStockState_DefaultMiningGoodStockState,
	types.GoodState_GoodStateWait:           types.MiningGoodStockState_MiningGoodStockStateWait,
	types.GoodState_GoodStateFail:           types.MiningGoodStockState_MiningGoodStockStateFail,
	types.GoodState_GoodStateReady:          types.MiningGoodStockState_MiningGoodStockStateReady,
	types.GoodState_GoodStateCheckHashRate:  types.MiningGoodStockState_MiningGoodStockStateCheckHashRate,
	types.GoodState_GoodStateCreateGoodUser: types.MiningGoodStockState_MiningGoodStockStateCreateGoodUser,
}
