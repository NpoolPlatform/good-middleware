package powerrental

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	powerrental.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	powerrental.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
