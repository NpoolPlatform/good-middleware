package manufacturer

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device/manufacturer"
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

func CreateManufacturer(ctx context.Context, req *npool.ManufacturerReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateManufacturer(_ctx, &npool.CreateManufacturerRequest{
			Info: req,
		})
	})
	return err
}

func UpdateManufacturer(ctx context.Context, req *npool.ManufacturerReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateManufacturer(_ctx, &npool.UpdateManufacturerRequest{
			Info: req,
		})
	})
	return err
}

func GetManufacturer(ctx context.Context, entID string) (*npool.Manufacturer, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetManufacturer(_ctx, &npool.GetManufacturerRequest{
			EntID: entID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Manufacturer), nil
}

func GetManufacturers(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Manufacturer, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetManufacturers(_ctx, &npool.GetManufacturersRequest{
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
	return _infos.([]*npool.Manufacturer), total, nil
}

func DeleteManufacturer(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteManufacturer(_ctx, &npool.DeleteManufacturerRequest{
			Info: &npool.ManufacturerReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
