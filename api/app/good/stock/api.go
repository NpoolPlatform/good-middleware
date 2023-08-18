package appstock

import (
	appstock "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appstock.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appstock.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
