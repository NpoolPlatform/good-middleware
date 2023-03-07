//nolint:nolintlint,dupl
package appdefaultgood

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	mgrpb "github.com/NpoolPlatform/message/npool/good/mgr/v1/appdefaultgood"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/appdefaultgood"

	constant "github.com/NpoolPlatform/good-middleware/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func GetAppDefaultGood(ctx context.Context, id string) (*mgrpb.AppDefaultGood, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppDefaultGood(ctx, &npool.GetAppDefaultGoodRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*mgrpb.AppDefaultGood), nil
}

func GetAppDefaultGoods(ctx context.Context, conds *mgrpb.Conds, offset, limit int32) ([]*mgrpb.AppDefaultGood, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppDefaultGoods(ctx, &npool.GetAppDefaultGoodsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, err
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return infos.([]*mgrpb.AppDefaultGood), total, nil
}

func GetAppDefaultGoodOnly(ctx context.Context, conds *mgrpb.Conds) (*mgrpb.AppDefaultGood, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetAppDefaultGoodOnly(ctx, &npool.GetAppDefaultGoodOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*mgrpb.AppDefaultGood), nil
}
