package delegatedstaking

import (
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

type DelegatedStaking interface {
	GoodServiceStartAt() uint32
	GoodStartMode() types.GoodStartMode
}

type delegatedstaking struct {
	delegatedstaking    *ent.DelegatedStaking
	goodBase            *ent.GoodBase
	appGoodBase         *ent.AppGoodBase
	appDelegatedStaking *ent.AppDelegatedStaking
}

func (pr *delegatedstaking) GoodServiceStartAt() uint32 {
	now := uint32(time.Now().Unix())
	if now < pr.goodBase.ServiceStartAt {
		return pr.goodBase.ServiceStartAt
	}
	return now
}

func (pr *delegatedstaking) GoodStartMode() types.GoodStartMode {
	return types.GoodStartMode(types.GoodStartMode_value[pr.goodBase.StartMode])
}
