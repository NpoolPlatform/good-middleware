package label

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	label.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	label.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
