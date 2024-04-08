package topmostgood

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second, //nolint
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreateTopMostGood(ctx context.Context, req *npool.TopMostGoodReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateTopMostGood(_ctx, &npool.CreateTopMostGoodRequest{
			Info: req,
		})
	})
	return err
}

func GetTopMostGood(ctx context.Context, id string) (*npool.TopMostGood, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostGood(ctx, &npool.GetTopMostGoodRequest{
			EntID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.TopMostGood), nil
}

func GetTopMostGoods(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.TopMostGood, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostGoods(ctx, &npool.GetTopMostGoodsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.Total
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return _infos.([]*npool.TopMostGood), total, nil
}

func GetTopMostGoodOnly(ctx context.Context, conds *npool.Conds) (*npool.TopMostGood, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostGoods(ctx, &npool.GetTopMostGoodsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.TopMostGood)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.TopMostGood)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.TopMostGood)[0], nil
}

func DeleteTopMostGood(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteTopMostGood(_ctx, &npool.DeleteTopMostGoodRequest{
			Info: &npool.TopMostGoodReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateTopMostGood(ctx context.Context, req *npool.TopMostGoodReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateTopMostGood(_ctx, &npool.UpdateTopMostGoodRequest{
			Info: req,
		})
	})
	return err
}
