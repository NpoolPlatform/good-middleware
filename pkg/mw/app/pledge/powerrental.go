package pledge

import (
	"time"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"
)

type Pledge interface {
	GoodServiceStartAt() uint32
	GoodStartMode() types.GoodStartMode
}

type pledge struct {
	pledge      *ent.Pledge
	goodBase    *ent.GoodBase
	appGoodBase *ent.AppGoodBase
	appPledge   *ent.AppPledge
}

func (pr *pledge) GoodServiceStartAt() uint32 {
	now := uint32(time.Now().Unix())
	if now < pr.goodBase.ServiceStartAt {
		return pr.goodBase.ServiceStartAt
	}
	return now
}

func (pr *pledge) GoodStartMode() types.GoodStartMode {
	return types.GoodStartMode(types.GoodStartMode_value[pr.goodBase.StartMode])
}
