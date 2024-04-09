package coin

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin"
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

func CreateGoodCoin(ctx context.Context, req *npool.GoodCoinReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateGoodCoin(_ctx, &npool.CreateGoodCoinRequest{
			Info: req,
		})
	})
	return err
}

func UpdateGoodCoin(ctx context.Context, req *npool.GoodCoinReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateGoodCoin(_ctx, &npool.UpdateGoodCoinRequest{
			Info: req,
		})
	})
	return err
}

func GetGoodCoin(ctx context.Context, entID string) (*npool.GoodCoin, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetGoodCoin(_ctx, &npool.GetGoodCoinRequest{
			EntID: entID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.GoodCoin), nil
}

func GetGoodCoins(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.GoodCoin, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetGoodCoins(_ctx, &npool.GetGoodCoinsRequest{
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
	return _infos.([]*npool.GoodCoin), total, nil
}

func GetGoodCoinOnly(ctx context.Context, conds *npool.Conds) (info *npool.GoodCoin, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetGoodCoins(_ctx, &npool.GetGoodCoinsRequest{
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
	if len(_infos.([]*npool.GoodCoin)) == 0 {
		return nil, fmt.Errorf("invalid goodcoin")
	}
	if len(_infos.([]*npool.GoodCoin)) > 1 {
		return nil, fmt.Errorf("too many goodcoins")
	}
	return _infos.([]*npool.GoodCoin)[0], nil
}

func ExistGoodCoinConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistGoodCoinConds(_ctx, &npool.ExistGoodCoinCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}

func DeleteGoodCoin(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteGoodCoin(_ctx, &npool.DeleteGoodCoinRequest{
			Info: &npool.GoodCoinReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
