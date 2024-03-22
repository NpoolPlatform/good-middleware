package device

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/device"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	device.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	device.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
