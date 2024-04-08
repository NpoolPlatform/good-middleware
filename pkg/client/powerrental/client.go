package powerrental

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	servicename "github.com/NpoolPlatform/good-middleware/pkg/servicename"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/powerrental"
	"google.golang.org/grpc"
)

func withClient(ctx context.Context, handler func(context.Context, npool.MiddlewareClient) (interface{}, error)) (interface{}, error) {
	return grpc2.WithGRPCConn(
		ctx,
		servicename.ServiceDomain,
		10*time.Second, //nolint
		func(_ctx context.Context, conn *grpc.ClientConn) (interface{}, error) {
			return handler(_ctx, npool.NewMiddlewareClient(conn))
		},
		grpc2.GRPCTAG,
	)
}

func CreatePowerRental(ctx context.Context, req *npool.PowerRentalReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.CreatePowerRental(_ctx, &npool.CreatePowerRentalRequest{
			Info: req,
		})
	})
	return err
}

func GetPowerRental(ctx context.Context, goodID string) (*npool.PowerRental, error) {
	info, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPowerRental(ctx, &npool.GetPowerRentalRequest{
			GoodID: goodID,
		})
		if err != nil {
			return nil, err
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, err
	}
	return info.(*npool.PowerRental), nil
}

func GetPowerRentals(ctx context.Context, conds *npool.Conds, offset, limit int32) (infos []*npool.PowerRental, total uint32, err error) {
	_infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPowerRentals(ctx, &npool.GetPowerRentalsRequest{
			Conds:  conds,
			Offset: offset,
			Limit:  limit,
		})
		if err != nil {
			return nil, err
		}
		total = resp.Total
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, err
	}
	return _infos.([]*npool.PowerRental), total, nil
}

func GetPowerRentalOnly(ctx context.Context, conds *npool.Conds) (*npool.PowerRental, error) {
	infos, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		resp, err := cli.GetPowerRentals(ctx, &npool.GetPowerRentalsRequest{
			Conds:  conds,
			Offset: 0,
			Limit:  2,
		})
		if err != nil {
			return nil, err
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, err
	}
	if len(infos.([]*npool.PowerRental)) == 0 {
		return nil, nil
	}
	if len(infos.([]*npool.PowerRental)) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return infos.([]*npool.PowerRental)[0], nil
}

func DeletePowerRental(ctx context.Context, id *uint32, entID, goodID *string) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.DeletePowerRental(_ctx, &npool.DeletePowerRentalRequest{
			Info: &npool.PowerRentalReq{
				ID:     id,
				EntID:  entID,
				GoodID: goodID,
			},
		})
	})
	return err
}

func UpdatePowerRental(ctx context.Context, req *npool.PowerRentalReq) error {
	_, err := withClient(ctx, func(_ctx context.Context, cli npool.MiddlewareClient) (interface{}, error) {
		return cli.UpdatePowerRental(_ctx, &npool.UpdatePowerRentalRequest{
			Info: req,
		})
	})
	return err
}
