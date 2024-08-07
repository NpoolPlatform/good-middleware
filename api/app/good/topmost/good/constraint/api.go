package constraint

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/topmost/good/constraint"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	constraint.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	constraint.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
