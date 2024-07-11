package recommend

import (
	"context"
	"time"

	wlog "github.com/NpoolPlatform/go-service-framework/pkg/wlog"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"
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

func CreateRecommend(ctx context.Context, req *npool.RecommendReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateRecommend(_ctx, &npool.CreateRecommendRequest{
			Info: req,
		})
	})
	return err
}

func GetRecommends(ctx context.Context, conds *npool.Conds, offset, limit int32) (intos []*npool.Recommend, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRecommends(ctx, &npool.GetRecommendsRequest{
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
	return _infos.([]*npool.Recommend), total, nil
}

func GetRecommendOnly(ctx context.Context, conds *npool.Conds) (*npool.Recommend, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRecommends(ctx, &npool.GetRecommendsRequest{
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
	if len(infos.([]*npool.Recommend)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Recommend)) > 1 {
		return nil, wlog.Errorf("too many records")
	}
	return infos.([]*npool.Recommend)[0], nil
}

func ExistRecommendConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistRecommendConds(ctx, &npool.ExistRecommendCondsRequest{
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

func UpdateRecommend(ctx context.Context, req *npool.RecommendReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateRecommend(_ctx, &npool.UpdateRecommendRequest{
			Info: req,
		})
	})
	return err
}

func DeleteRecommend(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteRecommend(_ctx, &npool.DeleteRecommendRequest{
			Info: &npool.RecommendReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func GetRecommend(ctx context.Context, id string) (*npool.Recommend, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetRecommend(ctx, &npool.GetRecommendRequest{
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
	return info.(*npool.Recommend), nil
}
