package recommend

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/recommend"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	recommend.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	recommend.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
