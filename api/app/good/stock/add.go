//nolint:dupl
package appstock

import (
	"context"

	appstock1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/stock"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/stock"
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
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithID(req.ID, true),
		appstock1.WithAppID(req.AppID, true),
		appstock1.WithGoodID(req.GoodID, true),
		appstock1.WithAppGoodID(req.AppGoodID, true),
		appstock1.WithLocked(req.Locked, false),
		appstock1.WithWaitStart(req.WaitStart, false),
		appstock1.WithInService(req.InService, false),
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

func (s *Server) AddReserved(ctx context.Context, in *npool.AddReservedRequest) (*npool.AddReservedResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"AddReserved",
			"In", in,
		)
		return &npool.AddReservedResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := appstock1.NewHandler(
		ctx,
		appstock1.WithID(req.ID, true),
		appstock1.WithAppID(req.AppID, true),
		appstock1.WithGoodID(req.GoodID, true),
		appstock1.WithAppGoodID(req.AppGoodID, true),
		appstock1.WithReserved(req.Reserved, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"AddReserved",
			"In", in,
			"Error", err,
		)
		return &npool.AddReservedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.AddReserved(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"AddReserved",
			"In", in,
			"Error", err,
		)
		return &npool.AddReservedResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.AddReservedResponse{
		Info: info,
	}, nil
}