package api

import (
	"context"

	appfee "github.com/NpoolPlatform/good-middleware/api/app/fee"
	"github.com/NpoolPlatform/good-middleware/api/app/good/comment"
	appdefaultgood "github.com/NpoolPlatform/good-middleware/api/app/good/default"
	"github.com/NpoolPlatform/good-middleware/api/app/good/description"
	displaycolor "github.com/NpoolPlatform/good-middleware/api/app/good/display/color"
	displayname "github.com/NpoolPlatform/good-middleware/api/app/good/display/name"
	"github.com/NpoolPlatform/good-middleware/api/app/good/label"
	"github.com/NpoolPlatform/good-middleware/api/app/good/like"
	appgoodposter "github.com/NpoolPlatform/good-middleware/api/app/good/poster"
	"github.com/NpoolPlatform/good-middleware/api/app/good/recommend"
	appgoodrequired "github.com/NpoolPlatform/good-middleware/api/app/good/required"
	"github.com/NpoolPlatform/good-middleware/api/app/good/score"
	appstock "github.com/NpoolPlatform/good-middleware/api/app/good/stock"
	topmost "github.com/NpoolPlatform/good-middleware/api/app/good/topmost"
	topmostconstraint "github.com/NpoolPlatform/good-middleware/api/app/good/topmost/constraint"
	topmostgood "github.com/NpoolPlatform/good-middleware/api/app/good/topmost/good"
	topmostgoodconstraint "github.com/NpoolPlatform/good-middleware/api/app/good/topmost/good/constraint"
	topmostgoodposter "github.com/NpoolPlatform/good-middleware/api/app/good/topmost/good/poster"
	topmostposter "github.com/NpoolPlatform/good-middleware/api/app/good/topmost/poster"
	apppowerrental "github.com/NpoolPlatform/good-middleware/api/app/powerrental"
	appsimulatepowerrental "github.com/NpoolPlatform/good-middleware/api/app/powerrental/simulate"
	"github.com/NpoolPlatform/good-middleware/api/device"
	manufacturer "github.com/NpoolPlatform/good-middleware/api/device/manufacturer"
	deviceposter "github.com/NpoolPlatform/good-middleware/api/device/poster"
	fee "github.com/NpoolPlatform/good-middleware/api/fee"
	good "github.com/NpoolPlatform/good-middleware/api/good"
	goodcoin "github.com/NpoolPlatform/good-middleware/api/good/coin"
	goodrequired "github.com/NpoolPlatform/good-middleware/api/good/required"
	"github.com/NpoolPlatform/good-middleware/api/good/reward/history"
	powerrental "github.com/NpoolPlatform/good-middleware/api/powerrental"
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
	appsimulatepowerrental.Register(server)
	apppowerrental.Register(server)
	appstock.Register(server)
	topmost.Register(server)
	topmostconstraint.Register(server)
	topmostgoodconstraint.Register(server)
	topmostgoodposter.Register(server)
	topmostposter.Register(server)
	topmostgood.Register(server)
	comment.Register(server)
	appgoodposter.Register(server)
	label.Register(server)
	displayname.Register(server)
	displaycolor.Register(server)
	description.Register(server)
	like.Register(server)
	recommend.Register(server)
	goodrequired.Register(server)
	goodcoin.Register(server)
	appgoodrequired.Register(server)
	history.Register(server)
	score.Register(server)
	device.Register(server)
	manufacturer.Register(server)
	deviceposter.Register(server)
	brand.Register(server)
	location.Register(server)
	appfee.Register(server)
	fee.Register(server)
	powerrental.Register(server)
	good.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := goodmw.RegisterMiddlewareHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
