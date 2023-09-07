package location

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/vender/location"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	location.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	location.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
