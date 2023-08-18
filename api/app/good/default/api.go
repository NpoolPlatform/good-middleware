package appdefaultgood

import (
	appdefaultgood "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/default"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appdefaultgood.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appdefaultgood.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
