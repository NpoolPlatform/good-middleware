package pledge

import (
	"context"

	pledge1 "github.com/NpoolPlatform/good-middleware/pkg/mw/pledge"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/pledge"
)

func (s *Server) DeletePledge(ctx context.Context, in *npool.DeletePledgeRequest) (*npool.DeletePledgeResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"DeletePledge",
			"In", in,
		)
		return &npool.DeletePledgeResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := pledge1.NewHandler(
		ctx,
		pledge1.WithID(req.ID, false),
		pledge1.WithEntID(req.EntID, false),
		pledge1.WithGoodID(req.GoodID, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"DeletePledge",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.DeletePledge(ctx); err != nil {
		logger.Sugar().Errorw(
			"DeletePledge",
			"In", in,
			"Error", err,
		)
		return &npool.DeletePledgeResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.DeletePledgeResponse{}, nil
}
