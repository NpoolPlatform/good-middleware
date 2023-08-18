//nolint:dupl
package stock

import (
	"context"

	stock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/good/stock"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/good/stock"
)

func (s *Server) AddStock(ctx context.Context, in *npool.AddStockRequest) (*npool.AddStockResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"AddStock",
			"In", in,
		)
		return &npool.AddStockResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := stock1.NewHandler(
		ctx,
		stock1.WithID(req.ID, true),
		stock1.WithGoodID(req.GoodID, true),
		stock1.WithAppReserved(req.AppReserved, false),
		stock1.WithLocked(req.Locked, false),
		stock1.WithWaitStart(req.WaitStart, false),
		stock1.WithInService(req.InService, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AddStock",
			"In", in,
			"Error", err,
		)
		return &npool.AddStockResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.AddStock(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AddStock",
			"In", in,
			"Error", err,
		)
		return &npool.AddStockResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AddStockResponse{
		Info: info,
	}, nil
}
