package appstock

import (
	"context"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second,
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func LockStock(ctx context.Context, req *npool.LockRequest) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.Lock(_ctx, req)
	})
	return err
}

func UnlockStock(ctx context.Context, req *npool.UnlockRequest) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.Unlock(_ctx, req)
	})
	return err
}

func WaitStartStock(ctx context.Context, req *npool.WaitStartRequest) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.WaitStart(_ctx, req)
	})
	return err
}

func InServiceStock(ctx context.Context, req *npool.InServiceRequest) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.InService(_ctx, req)
	})
	return err
}

func ChargeBackStock(ctx context.Context, req *npool.ChargeBackRequest) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.ChargeBack(_ctx, req)
	})
	return err
}

func ExpireStock(ctx context.Context, req *npool.ExpireRequest) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.Expire(_ctx, req)
	})
	return err
}
