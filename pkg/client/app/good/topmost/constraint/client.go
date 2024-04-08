package constraint

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/constraint"
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

func CreateTopMostConstraint(ctx context.Context, req *npool.TopMostConstraintReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateTopMostConstraint(_ctx, &npool.CreateTopMostConstraintRequest{
			Info: req,
		})
	})
	return err
}

func GetTopMostConstraints(ctx context.Context, conds *npool.Conds, offset, limit int32) (intos []*npool.TopMostConstraint, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostConstraints(ctx, &npool.GetTopMostConstraintsRequest{
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
	return _infos.([]*npool.TopMostConstraint), total, nil
}

func GetTopMostConstraintOnly(ctx context.Context, conds *npool.Conds) (*npool.TopMostConstraint, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostConstraints(ctx, &npool.GetTopMostConstraintsRequest{
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
	if len(infos.([]*npool.TopMostConstraint)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.TopMostConstraint)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.TopMostConstraint)[0], nil
}

func UpdateTopMostConstraint(ctx context.Context, req *npool.TopMostConstraintReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateTopMostConstraint(_ctx, &npool.UpdateTopMostConstraintRequest{
			Info: req,
		})
	})
	return err
}

func DeleteTopMostConstraint(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteTopMostConstraint(_ctx, &npool.DeleteTopMostConstraintRequest{
			Info: &npool.TopMostConstraintReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func GetTopMostConstraint(ctx context.Context, id string) (*npool.TopMostConstraint, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetTopMostConstraint(ctx, &npool.GetTopMostConstraintRequest{
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
	return info.(*npool.TopMostConstraint), nil
}
