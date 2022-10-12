package appgood

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"
)

func GetGoods(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Good, uint32, error) {
	return nil, 0, nil
}
