//nolint:dupl
package label

import (
	"context"

	label1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/good/label"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/good/label"
)

func (s *Server) CreateLabel(ctx context.Context, in *npool.CreateLabelRequest) (*npool.CreateLabelResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"CreateLabel",
			"In", in,
		)
		return &npool.CreateLabelResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := label1.NewHandler(
		ctx,
		label1.WithEntID(req.EntID, false),
		label1.WithAppGoodID(req.AppGoodID, true),
		label1.WithIcon(req.Icon, false),
		label1.WithIconBgColor(req.IconBgColor, false),
		label1.WithLabel(req.Label, true),
		label1.WithLabelBgColor(req.LabelBgColor, false),
		label1.WithIndex(func() *uint8 {
			if req.Index == nil {
				return nil
			}
			u := uint8(*req.Index)
			return &u
		}(), false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"CreateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.CreateLabel(ctx); err != nil {
		logger.Sugar().Errorw(
			"CreateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.CreateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.CreateLabelResponse{}, nil
}
