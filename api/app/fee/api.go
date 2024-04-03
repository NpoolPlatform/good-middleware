package fee

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	fee.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	fee.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
