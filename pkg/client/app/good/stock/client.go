package appstock

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

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

func LockStock(ctx context.Context, req *npool.LockRequest) (*npool.Stock, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.Lock(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Stock), nil
}

func UnlockStock(ctx context.Context, req *npool.UnlockRequest) (*npool.Stock, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.Unlock(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Stock), nil
}

func WaitStartStock(ctx context.Context, req *npool.WaitStartRequest) (*npool.Stock, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.WaitStart(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Stock), nil
}

func InServiceStock(ctx context.Context, req *npool.InServiceRequest) (*npool.Stock, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.InService(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Stock), nil
}

func ChargeBackStock(ctx context.Context, req *npool.ChargeBackRequest) (*npool.Stock, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.ChargeBack(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Stock), nil
}

func ExpireStock(ctx context.Context, req *npool.ExpireRequest) (*npool.Stock, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (cruder.Any, error) {
		resp, err := cli.Expire(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.Stock), nil
}
