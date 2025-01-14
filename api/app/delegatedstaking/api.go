package delegatedstaking

import (
	"github.com/NpoolPlatform/message/npool/good/mw/v1/app/delegatedstaking"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	delegatedstaking.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	delegatedstaking.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
