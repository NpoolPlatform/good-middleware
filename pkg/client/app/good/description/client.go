package description

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"
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

func CreateDescription(ctx context.Context, req *npool.DescriptionReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateDescription(_ctx, &npool.CreateDescriptionRequest{
			Info: req,
		})
	})
	return err
}

func GetDescriptions(ctx context.Context, conds *npool.Conds, offset, limit int32) (intos []*npool.Description, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDescriptions(ctx, &npool.GetDescriptionsRequest{
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
	return _infos.([]*npool.Description), total, nil
}

func GetDescriptionOnly(ctx context.Context, conds *npool.Conds) (*npool.Description, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDescriptions(ctx, &npool.GetDescriptionsRequest{
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
	if len(infos.([]*npool.Description)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Description)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Description)[0], nil
}

func ExistDescriptionConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistDescriptionConds(ctx, &npool.ExistDescriptionCondsRequest{
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

func UpdateDescription(ctx context.Context, req *npool.DescriptionReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateDescription(_ctx, &npool.UpdateDescriptionRequest{
			Info: req,
		})
	})
	return err
}

func DeleteDescription(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteDescription(_ctx, &npool.DeleteDescriptionRequest{
			Info: &npool.DescriptionReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func GetDescription(ctx context.Context, id string) (*npool.Description, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDescription(ctx, &npool.GetDescriptionRequest{
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
	return info.(*npool.Description), nil
}
