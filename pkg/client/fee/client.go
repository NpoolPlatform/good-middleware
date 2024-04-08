package fee

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"
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

func CreateFee(ctx context.Context, req *npool.FeeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateFee(_ctx, &npool.CreateFeeRequest{
			Info: req,
		})
	})
	return err
}

func UpdateFee(ctx context.Context, req *npool.FeeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateFee(_ctx, &npool.UpdateFeeRequest{
			Info: req,
		})
	})
	return err
}

func GetFee(ctx context.Context, goodID string) (*npool.Fee, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetFee(_ctx, &npool.GetFeeRequest{
			GoodID: goodID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Fee), nil
}

func GetFees(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Fee, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetFees(_ctx, &npool.GetFeesRequest{
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
	return _infos.([]*npool.Fee), total, nil
}

func DeleteFee(ctx context.Context, id *uint32, entID, goodID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteFee(_ctx, &npool.DeleteFeeRequest{
			Info: &npool.FeeReq{
				ID:     id,
				EntID:  entID,
				GoodID: goodID,
			},
		})
	})
	return err
}
