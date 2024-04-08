package topmost

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost"
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

func CreateTopMost(ctx context.Context, req *npool.TopMostReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateTopMost(_ctx, &npool.CreateTopMostRequest{
			Info: req,
		})
	})
	return err
}

func GetTopMost(ctx context.Context, id string) (*npool.TopMost, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMost(ctx, &npool.GetTopMostRequest{
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
	return info.(*npool.TopMost), nil
}

func GetTopMosts(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.TopMost, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMosts(ctx, &npool.GetTopMostsRequest{
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
	return _infos.([]*npool.TopMost), total, nil
}

func GetTopMostOnly(ctx context.Context, conds *npool.Conds) (*npool.TopMost, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMosts(ctx, &npool.GetTopMostsRequest{
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
	if len(infos.([]*npool.TopMost)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.TopMost)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.TopMost)[0], nil
}

func DeleteTopMost(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteTopMost(_ctx, &npool.DeleteTopMostRequest{
			Info: &npool.TopMostReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateTopMost(ctx context.Context, req *npool.TopMostReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateTopMost(_ctx, &npool.UpdateTopMostRequest{
			Info: req,
		})
	})
	return err
}
