package history

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/good/coin/reward/history"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	history.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	history.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
