package fee

import (
	"context"

	fee1 "github.com/NpoolPlatform/good-middleware/pkg/mw/fee"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/fee"
)

func (s *Server) UpdateFee(ctx context.Context, in *npool.UpdateFeeRequest) (*npool.UpdateFeeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateFee",
			"In", in,
		)
		return &npool.UpdateFeeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := fee1.NewHandler(
		ctx,
		fee1.WithID(req.ID, false),
		fee1.WithEntID(req.EntID, false),
		fee1.WithGoodID(req.GoodID, false),
		fee1.WithGoodType(req.GoodType, false),
		fee1.WithName(req.Name, false),
		fee1.WithSettlementType(req.SettlementType, false),
		fee1.WithUnitValue(req.UnitValue, false),
		fee1.WithDurationDisplayType(req.DurationDisplayType, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateFee",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateFee(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateFee",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateFeeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateFeeResponse{}, nil
}
