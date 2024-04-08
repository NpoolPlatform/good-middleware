package location

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"
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

func CreateLocation(ctx context.Context, req *npool.LocationReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateLocation(_ctx, &npool.CreateLocationRequest{
			Info: req,
		})
	})
	return err
}

func GetLocation(ctx context.Context, id string) (*npool.Location, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetLocation(ctx, &npool.GetLocationRequest{
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
	return info.(*npool.Location), nil
}

func GetLocations(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Location, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetLocations(ctx, &npool.GetLocationsRequest{
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
	return _infos.([]*npool.Location), total, nil
}

func GetLocationOnly(ctx context.Context, conds *npool.Conds) (*npool.Location, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetLocations(ctx, &npool.GetLocationsRequest{
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
	if len(infos.([]*npool.Location)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Location)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Location)[0], nil
}

func UpdateLocation(ctx context.Context, req *npool.LocationReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateLocation(_ctx, &npool.UpdateLocationRequest{
			Info: req,
		})
	})
	return err
}

func DeleteLocation(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteLocation(_ctx, &npool.DeleteLocationRequest{
			Info: &npool.LocationReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}
