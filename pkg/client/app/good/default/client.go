package appdefaultgood

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"

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

func CreateDefault(ctx context.Context, in *npool.DefaultReq) (*npool.Default, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.CreateDefault(ctx, &npool.CreateDefaultRequest{
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
	return info.(*npool.Default), nil
}

func GetDefault(ctx context.Context, id string) (*npool.Default, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDefault(ctx, &npool.GetDefaultRequest{
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
	return info.(*npool.Default), nil
}

func GetDefaults(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Default, uint32, error) {
	total := uint32(0)

	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
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
	return infos.([]*npool.Default), total, nil
}

func GetDefaultOnly(ctx context.Context, conds *npool.Conds) (*npool.Default, error) {
	const limit = 2
	infos, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.GetDefaults(ctx, &npool.GetDefaultsRequest{
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
	if len(infos.([]*npool.Default)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.Default)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.Default)[0], nil
}

func DeleteDefault(ctx context.Context, in *npool.DefaultReq) (*npool.Default, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.DeleteDefault(ctx, &npool.DeleteDefaultRequest{
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
	return info.(*npool.Default), nil
}
