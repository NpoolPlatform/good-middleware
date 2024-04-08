package constraint

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"
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

func CreateTopMostGoodConstraint(ctx context.Context, req *npool.TopMostGoodConstraintReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateTopMostGoodConstraint(_ctx, &npool.CreateTopMostGoodConstraintRequest{
			Info: req,
		})
	})
	return err
}

func GetTopMostGoodConstraints(ctx context.Context, conds *npool.Conds, offset, limit int32) (intos []*npool.TopMostGoodConstraint, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostGoodConstraints(ctx, &npool.GetTopMostGoodConstraintsRequest{
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
	return _infos.([]*npool.TopMostGoodConstraint), total, nil
}

func GetTopMostGoodConstraintOnly(ctx context.Context, conds *npool.Conds) (*npool.TopMostGoodConstraint, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostGoodConstraints(ctx, &npool.GetTopMostGoodConstraintsRequest{
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
	if len(infos.([]*npool.TopMostGoodConstraint)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.TopMostGoodConstraint)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.TopMostGoodConstraint)[0], nil
}

func UpdateTopMostGoodConstraint(ctx context.Context, req *npool.TopMostGoodConstraintReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateTopMostGoodConstraint(_ctx, &npool.UpdateTopMostGoodConstraintRequest{
			Info: req,
		})
	})
	return err
}

func DeleteTopMostGoodConstraint(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteTopMostGoodConstraint(_ctx, &npool.DeleteTopMostGoodConstraintRequest{
			Info: &npool.TopMostGoodConstraintReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func GetTopMostGoodConstraint(ctx context.Context, id string) (*npool.TopMostGoodConstraint, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostGoodConstraint(ctx, &npool.GetTopMostGoodConstraintRequest{
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
	return info.(*npool.TopMostGoodConstraint), nil
}
