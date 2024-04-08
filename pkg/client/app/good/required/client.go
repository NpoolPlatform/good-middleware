package required

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/required"
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

func CreateRequired(ctx context.Context, req *npool.RequiredReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateRequired(_ctx, &npool.CreateRequiredRequest{
			Info: req,
		})
	})
	return err
}

func GetRequired(ctx context.Context, id string) (*npool.Required, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRequired(ctx, &npool.GetRequiredRequest{
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
	return info.(*npool.Required), nil
}

func GetRequireds(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Required, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRequireds(ctx, &npool.GetRequiredsRequest{
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
	return _infos.([]*npool.Required), total, nil
}

func GetRequiredOnly(ctx context.Context, conds *npool.Conds) (*npool.Required, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRequireds(ctx, &npool.GetRequiredsRequest{
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
	if len(infos.([]*npool.Required)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Required)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Required)[0], nil
}

func UpdateRequired(ctx context.Context, req *npool.RequiredReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateRequired(_ctx, &npool.UpdateRequiredRequest{
			Info: req,
		})
	})
	return err
}

func DeleteRequired(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteRequired(_ctx, &npool.DeleteRequiredRequest{
			Info: &npool.RequiredReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func ExistRequiredConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistRequiredConds(ctx, &npool.ExistRequiredCondsRequest{
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
