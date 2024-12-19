package pledge

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	pledge.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	pledge.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
