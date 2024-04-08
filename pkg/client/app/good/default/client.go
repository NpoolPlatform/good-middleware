package appdefaultgood

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"
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

func CreateDefault(ctx context.Context, req *npool.DefaultReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreateDefault(_ctx, &npool.CreateDefaultRequest{
			Info: req,
		})
	})
	return err
}

func GetDefault(ctx context.Context, id string) (*npool.Default, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDefault(ctx, &npool.GetDefaultRequest{
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
	return info.(*npool.Default), nil
}

func GetDefaults(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.Default, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDefaults(ctx, &npool.GetDefaultsRequest{
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
	return _infos.([]*npool.Default), total, nil
}

func GetDefaultOnly(ctx context.Context, conds *npool.Conds) (*npool.Default, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetDefaults(ctx, &npool.GetDefaultsRequest{
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
	if len(infos.([]*npool.Default)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Default)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Default)[0], nil
}

func DeleteDefault(ctx context.Context, id *uint32, entID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeleteDefault(_ctx, &npool.DeleteDefaultRequest{
			Info: &npool.DefaultReq{
				ID:    id,
				EntID: entID,
			},
		})
	})
	return err
}

func UpdateDefault(ctx context.Context, req *npool.DefaultReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdateDefault(_ctx, &npool.UpdateDefaultRequest{
			Info: req,
		})
	})
	return err
}
