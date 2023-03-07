package api

import (
	"context"

	"github.com/NpoolPlatform/good-middleware/api/appdefaultgood"

	"github.com/NpoolPlatform/good-middleware/api/recommend"

	goodmw "github.com/NpoolPlatform/message/npool/good/mw/v1"

	"github.com/NpoolPlatform/good-middleware/api/appgood"
	"github.com/NpoolPlatform/good-middleware/api/good"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodmw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodmw.RegisterMiddlewareServer(server, &Server{})
	good.Register(server)
	appgood.Register(server)
	recommend.Register(server)
	appdefaultgood.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := goodmw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	if err := good.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	if err := appgood.RegisterGateway(mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
