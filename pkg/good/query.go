package good

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/good"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good"
)

func GetGood(ctx context.Context, id string) (*npool.Good, error) {
	return nil, nil
}

func GetGoods(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*npool.Good, uint32, error) {
	return nil, 0, nil
}
