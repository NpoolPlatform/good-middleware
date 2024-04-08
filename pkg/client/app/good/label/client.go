package label

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"
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

func CreateLabel(ctx context.Context, req *npool.LabelReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateLabel(_ctx, &npool.CreateLabelRequest{
			Info: req,
		})
	})
	return err
}

func GetLabels(ctx context.Context, conds *npool.Conds, offset, limit int32) (intos []*npool.Label, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetLabels(ctx, &npool.GetLabelsRequest{
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
	return _infos.([]*npool.Label), total, nil
}

func GetLabelOnly(ctx context.Context, conds *npool.Conds) (*npool.Label, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetLabels(ctx, &npool.GetLabelsRequest{
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
	if len(infos.([]*npool.Label)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Label)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Label)[0], nil
}

func UpdateLabel(ctx context.Context, req *npool.LabelReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateLabel(_ctx, &npool.UpdateLabelRequest{
			Info: req,
		})
	})
	return err
}

func DeleteLabel(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteLabel(_ctx, &npool.DeleteLabelRequest{
			Info: &npool.LabelReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func GetLabel(ctx context.Context, id string) (*npool.Label, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetLabel(ctx, &npool.GetLabelRequest{
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
	return info.(*npool.Label), nil
}
