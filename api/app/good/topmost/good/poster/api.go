package poster

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/poster"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	poster.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	poster.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
