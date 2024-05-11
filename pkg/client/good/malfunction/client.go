package malfunction

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/malfunction"
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

func CreateMalfunction(ctx context.Context, req *npool.MalfunctionReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateMalfunction(_ctx, &npool.CreateMalfunctionRequest{
			Info: req,
		})
	})
	return err
}

func GetMalfunction(ctx context.Context, id string) (*npool.Malfunction, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetMalfunction(ctx, &npool.GetMalfunctionRequest{
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
	return info.(*npool.Malfunction), nil
}

func GetMalfunctions(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Malfunction, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetMalfunctions(ctx, &npool.GetMalfunctionsRequest{
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
	return _infos.([]*npool.Malfunction), total, nil
}

func GetMalfunctionOnly(ctx context.Context, conds *npool.Conds) (*npool.Malfunction, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetMalfunctions(ctx, &npool.GetMalfunctionsRequest{
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
	if len(infos.([]*npool.Malfunction)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Malfunction)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Malfunction)[0], nil
}

func UpdateMalfunction(ctx context.Context, req *npool.MalfunctionReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateMalfunction(_ctx, &npool.UpdateMalfunctionRequest{
			Info: req,
		})
	})
	return err
}

func DeleteMalfunction(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteMalfunction(_ctx, &npool.DeleteMalfunctionRequest{
			Info: &npool.MalfunctionReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func ExistMalfunctionConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistMalfunctionConds(ctx, &npool.ExistMalfunctionCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return false, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, err
	}
	return info.(bool), nil
}
