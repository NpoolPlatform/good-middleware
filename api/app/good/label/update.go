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

func (s *Server) UpdateLabel(ctx context.Context, in *npool.UpdateLabelRequest) (*npool.UpdateLabelResponse, error) {
	req := in.GetInfo()
	if req == nil {
		logger.Sugar().Errorw(
			"UpdateLabel",
			"In", in,
		)
		return &npool.UpdateLabelResponse{}, status.Error(codes.Aborted, "invalid argument")
	}
	handler, err := label1.NewHandler(
		ctx,
		label1.WithID(req.ID, false),
		label1.WithEntID(req.EntID, false),
		label1.WithIcon(req.Icon, false),
		label1.WithIconBgColor(req.IconBgColor, false),
		label1.WithLabel(req.Label, false),
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
			"UpdateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	if err := handler.UpdateLabel(ctx); err != nil {
		logger.Sugar().Errorw(
			"UpdateLabel",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateLabelResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateLabelResponse{}, nil
}
