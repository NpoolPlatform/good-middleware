package displaycolor

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/display/color"
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

func CreateDisplayColor(ctx context.Context, req *npool.DisplayColorReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateDisplayColor(_ctx, &npool.CreateDisplayColorRequest{
			Info: req,
		})
	})
	return err
}

func GetDisplayColors(ctx context.Context, conds *npool.Conds, offset, limit int32) (intos []*npool.DisplayColor, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDisplayColors(ctx, &npool.GetDisplayColorsRequest{
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
	return _infos.([]*npool.DisplayColor), total, nil
}

func GetDisplayColorOnly(ctx context.Context, conds *npool.Conds) (*npool.DisplayColor, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDisplayColors(ctx, &npool.GetDisplayColorsRequest{
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
	if len(infos.([]*npool.DisplayColor)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.DisplayColor)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.DisplayColor)[0], nil
}

func UpdateDisplayColor(ctx context.Context, req *npool.DisplayColorReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateDisplayColor(_ctx, &npool.UpdateDisplayColorRequest{
			Info: req,
		})
	})
	return err
}

func DeleteDisplayColor(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteDisplayColor(_ctx, &npool.DeleteDisplayColorRequest{
			Info: &npool.DisplayColorReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func GetDisplayColor(ctx context.Context, id string) (*npool.DisplayColor, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDisplayColor(ctx, &npool.GetDisplayColorRequest{
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
	return info.(*npool.DisplayColor), nil
}
