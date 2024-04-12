package displayname

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/name"
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

func CreateDisplayName(ctx context.Context, req *npool.DisplayNameReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateDisplayName(_ctx, &npool.CreateDisplayNameRequest{
			Info: req,
		})
	})
	return err
}

func GetDisplayNames(ctx context.Context, conds *npool.Conds, offset, limit int32) (intos []*npool.DisplayName, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDisplayNames(ctx, &npool.GetDisplayNamesRequest{
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
	return _infos.([]*npool.DisplayName), total, nil
}

func GetDisplayNameOnly(ctx context.Context, conds *npool.Conds) (*npool.DisplayName, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDisplayNames(ctx, &npool.GetDisplayNamesRequest{
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
	if len(infos.([]*npool.DisplayName)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.DisplayName)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.DisplayName)[0], nil
}

func ExistDisplayNameConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistDisplayNameConds(ctx, &npool.ExistDisplayNameCondsRequest{
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

func UpdateDisplayName(ctx context.Context, req *npool.DisplayNameReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateDisplayName(_ctx, &npool.UpdateDisplayNameRequest{
			Info: req,
		})
	})
	return err
}

func DeleteDisplayName(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteDisplayName(_ctx, &npool.DeleteDisplayNameRequest{
			Info: &npool.DisplayNameReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func GetDisplayName(ctx context.Context, id string) (*npool.DisplayName, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDisplayName(ctx, &npool.GetDisplayNameRequest{
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
	return info.(*npool.DisplayName), nil
}
