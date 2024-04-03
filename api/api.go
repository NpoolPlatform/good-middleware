package api

import (
	"context"

	appfee "github.com/NpoolPlatform/good-middleware/api/app/fee"
	"github.com/NpoolPlatform/good-middleware/api/app/good/comment"
	appdefaultgood "github.com/NpoolPlatform/good-middleware/api/app/good/default"
	"github.com/NpoolPlatform/good-middleware/api/app/good/like"
	"github.com/NpoolPlatform/good-middleware/api/app/good/recommend"
	"github.com/NpoolPlatform/good-middleware/api/app/good/score"
	appsimulategood "github.com/NpoolPlatform/good-middleware/api/app/good/simulate"
	appstock "github.com/NpoolPlatform/good-middleware/api/app/good/stock"
	topmost "github.com/NpoolPlatform/good-middleware/api/app/good/topmost"
	topmostgood "github.com/NpoolPlatform/good-middleware/api/app/good/topmost/good"
	"github.com/NpoolPlatform/good-middleware/api/device"
	"github.com/NpoolPlatform/good-middleware/api/good/required"
	"github.com/NpoolPlatform/good-middleware/api/good/reward/history"
	"github.com/NpoolPlatform/good-middleware/api/vender/brand"
	"github.com/NpoolPlatform/good-middleware/api/vender/location"
	goodmw "github.com/NpoolPlatform/message/npool/good/mw/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	goodmw.UnimplementedMiddlewareServer
}

func Register(server grpc.ServiceRegistrar) {
	goodmw.RegisterMiddlewareServer(server, &Server{})
	appdefaultgood.Register(server)
	appsimulategood.Register(server)
	appstock.Register(server)
	topmost.Register(server)
	topmostgood.Register(server)
	comment.Register(server)
	like.Register(server)
	recommend.Register(server)
	required.Register(server)
	history.Register(server)
	score.Register(server)
	device.Register(server)
	brand.Register(server)
	location.Register(server)
	appfee.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := goodmw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
