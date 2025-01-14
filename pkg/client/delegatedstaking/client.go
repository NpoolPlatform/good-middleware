package delegatedstaking

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/delegatedstaking"
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

func CreateDelegatedStaking(ctx context.Context, req *npool.DelegatedStakingReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateDelegatedStaking(_ctx, &npool.CreateDelegatedStakingRequest{
			Info: req,
		})
	})
	return err
}

func GetDelegatedStaking(ctx context.Context, goodID string) (*npool.DelegatedStaking, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDelegatedStaking(ctx, &npool.GetDelegatedStakingRequest{
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
	return info.(*npool.DelegatedStaking), nil
}

func GetDelegatedStakings(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.DelegatedStaking, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDelegatedStakings(ctx, &npool.GetDelegatedStakingsRequest{
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
	return _infos.([]*npool.DelegatedStaking), total, nil
}

func ExistDelegatedStakingConds(ctx context.Context, conds *npool.Conds) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistDelegatedStakingConds(ctx, &npool.ExistDelegatedStakingCondsRequest{
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

func GetDelegatedStakingOnly(ctx context.Context, conds *npool.Conds) (*npool.DelegatedStaking, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDelegatedStakings(ctx, &npool.GetDelegatedStakingsRequest{
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
	if len(infos.([]*npool.DelegatedStaking)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.DelegatedStaking)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.DelegatedStaking)[0], nil
}

func DeleteDelegatedStaking(ctx context.Context, id *uint32, entID, goodID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteDelegatedStaking(_ctx, &npool.DeleteDelegatedStakingRequest{
			Info: &npool.DelegatedStakingReq{
				ID:     id,
				EntID:  entID,
				GoodID: goodID,
			},
		})
	})
	return err
}

func UpdateDelegatedStaking(ctx context.Context, req *npool.DelegatedStakingReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateDelegatedStaking(_ctx, &npool.UpdateDelegatedStakingRequest{
			Info: req,
		})
	})
	return err
}
