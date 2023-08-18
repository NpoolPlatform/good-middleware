package deviceinfo

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/deviceinfo"

	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.MiddlewareClient) (cruder.Any, error)

func do(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(servicename.ServiceDomain, grpc2.GRPCTAG)
	if err != nil {
		return nil, err
	}

	defer conn.Close()

	cli := npool.NewMiddlewareClient(conn)

	return handler(_ctx, cli)
}

func CreateDeviceInfo(ctx context.Context, in *npool.DeviceInfoReq) (*npool.DeviceInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateDeviceInfo(ctx, &npool.CreateDeviceInfoRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}

func GetDeviceInfo(ctx context.Context, id string) (*npool.DeviceInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDeviceInfo(ctx, &npool.GetDeviceInfoRequest{
			ID: id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}

func GetDeviceInfos(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.DeviceInfo, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDeviceInfos(ctx, &npool.GetDeviceInfosRequest{
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
	return infos.([]*npool.DeviceInfo), total, nil
}

func GetDeviceInfoOnly(ctx context.Context, conds *npool.Conds) (*npool.DeviceInfo, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDeviceInfos(ctx, &npool.GetDeviceInfosRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.DeviceInfo)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.DeviceInfo)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.DeviceInfo)[0], nil
}

func DeleteDeviceInfo(ctx context.Context, in *npool.DeviceInfoReq) (*npool.DeviceInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteDeviceInfo(ctx, &npool.DeleteDeviceInfoRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}

func UpdateDeviceInfo(ctx context.Context, in *npool.DeviceInfoReq) (*npool.DeviceInfo, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.UpdateDeviceInfo(ctx, &npool.UpdateDeviceInfoRequest{
			Info: in,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.DeviceInfo), nil
}
