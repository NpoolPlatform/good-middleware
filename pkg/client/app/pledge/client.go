package pledge

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/pledge"
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

func CreatePledge(ctx context.Context, req *npool.PledgeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreatePledge(_ctx, &npool.CreatePledgeRequest{
			Info: req,
		})
	})
	return err
}

func GetPledge(ctx context.Context, appGoodID string) (*npool.Pledge, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPledge(ctx, &npool.GetPledgeRequest{
			AppGoodID: appGoodID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Pledge), nil
}

func GetPledges(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Pledge, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPledges(ctx, &npool.GetPledgesRequest{
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
	return _infos.([]*npool.Pledge), total, nil
}

func ExistPledgeConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistPledgeConds(ctx, &npool.ExistPledgeCondsRequest{
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

func GetPledgeOnly(ctx context.Context, conds *npool.Conds) (*npool.Pledge, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPledges(ctx, &npool.GetPledgesRequest{
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
	if len(infos.([]*npool.Pledge)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Pledge)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Pledge)[0], nil
}

func DeletePledge(ctx context.Context, id *uint32, entID, appGoodID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeletePledge(_ctx, &npool.DeletePledgeRequest{
			Info: &npool.PledgeReq{
				ID:        id,
				EntID:     entID,
				AppGoodID: appGoodID,
			},
		})
	})
	return err
}

func UpdatePledge(ctx context.Context, req *npool.PledgeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdatePledge(_ctx, &npool.UpdatePledgeRequest{
			Info: req,
		})
	})
	return err
}
