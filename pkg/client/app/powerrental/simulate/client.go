package simulate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second,
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreateSimulate(ctx context.Context, req *npool.SimulateReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateSimulate(_ctx, &npool.CreateSimulateRequest{
			Info: req,
		})
	})
	return err
}

func GetSimulate(ctx context.Context, id string) (*npool.Simulate, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSimulate(ctx, &npool.GetSimulateRequest{
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
	return info.(*npool.Simulate), nil
}

func GetSimulates(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Simulate, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSimulates(ctx, &npool.GetSimulatesRequest{
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
	return _infos.([]*npool.Simulate), total, nil
}

func GetSimulateOnly(ctx context.Context, conds *npool.Conds) (*npool.Simulate, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetSimulates(ctx, &npool.GetSimulatesRequest{
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
	if len(infos.([]*npool.Simulate)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Simulate)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Simulate)[0], nil
}

func DeleteSimulate(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteSimulate(_ctx, &npool.DeleteSimulateRequest{
			Info: &npool.SimulateReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateSimulate(ctx context.Context, req *npool.SimulateReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateSimulate(_ctx, &npool.UpdateSimulateRequest{
			Info: req,
		})
	})
	return err
}
