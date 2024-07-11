package description

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/description"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	description.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	description.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
