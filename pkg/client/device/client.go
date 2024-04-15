package devicetype

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/device"
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

func CreateDeviceType(ctx context.Context, req *npool.DeviceTypeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateDeviceType(_ctx, &npool.CreateDeviceTypeRequest{
			Info: req,
		})
	})
	return err
}

func GetDeviceType(ctx context.Context, id string) (*npool.DeviceType, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDeviceType(ctx, &npool.GetDeviceTypeRequest{
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
	return info.(*npool.DeviceType), nil
}

func GetDeviceTypes(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.DeviceType, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDeviceTypes(ctx, &npool.GetDeviceTypesRequest{
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
	return _infos.([]*npool.DeviceType), total, nil
}

func ExistDeviceTypeConds(ctx context.Context, conds *npool.Conds, offset, limit int32) (exist bool, err error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.ExistDeviceTypeConds(ctx, &npool.ExistDeviceTypeCondsRequest{
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

func GetDeviceTypeOnly(ctx context.Context, conds *npool.Conds) (*npool.DeviceType, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDeviceTypes(ctx, &npool.GetDeviceTypesRequest{
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
	if len(infos.([]*npool.DeviceType)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.DeviceType)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.DeviceType)[0], nil
}

func DeleteDeviceType(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteDeviceType(_ctx, &npool.DeleteDeviceTypeRequest{
			Info: &npool.DeviceTypeReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateDeviceType(ctx context.Context, req *npool.DeviceTypeReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateDeviceType(_ctx, &npool.UpdateDeviceTypeRequest{
			Info: req,
		})
	})
	return err
}
