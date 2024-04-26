package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/fee"
)

func (s *Server) CreateFee(ctx context.Context, in *npool.CreateFeeRequest) (*npool.CreateFeeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateFee",
			"In", in,
		)
		return &npool.CreateFeeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithEntID(req.EntID, false),
		fee1.WithAppID(req.AppID, true),
		fee1.WithGoodID(req.GoodID, true),
		fee1.WithAppGoodID(req.AppGoodID, true),
		fee1.WithProductPage(req.ProductPage, false),
		fee1.WithName(req.Name, true),
		fee1.WithBanner(req.Banner, false),
		fee1.WithUnitValue(req.UnitValue, false),
		fee1.WithMinOrderDurationSeconds(req.MinOrderDurationSeconds, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateFee",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err = handler.CreateFee(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateFee",
			"In", in,
			"Error", err,
		)
		return &npool.CreateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateFeeResponse{}, nil
}
