package appgood

import (
	"context"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appgood"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appgood"

	appgoodcli "github.com/NpoolPlatform/good-manager/pkg/client/appgood"
)

func UpdateGood(ctx context.Context, in *mgrpb.AppGoodReq) (*npool.Good, error) {
	info, err := appgoodcli.UpdateAppGood(ctx, in)
	if err != nil {
		return nil, err
	}
	return GetGood(ctx, info.ID)
}
