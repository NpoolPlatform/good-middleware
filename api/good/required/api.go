package required

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/good/required"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	required.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	required.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
