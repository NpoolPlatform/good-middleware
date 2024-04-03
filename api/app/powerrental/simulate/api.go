package appsimulategood

import (
	appsimulategood "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	appsimulategood.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	appsimulategood.RegisterMiddlewareServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
