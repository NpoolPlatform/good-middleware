package appsimulatepowerrental

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-middleware/pkg/const"
	appgoodbasecrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/good/goodbase"
	appsimulatepowerrentalcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/app/powerrental/simulate"
	apppowerrental1 "github.com/NpoolPlatform/good-middleware/pkg/mw/app/powerrental"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/app/powerrental/simulate"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Here we actually only have simulate order for PowerRental good

type Handler struct {
	ID *uint32
	appsimulatepowerrentalcrud.Req
	AppSimulatePowerRentalConds *appsimulatepowerrentalcrud.Conds
	AppGoodBaseConds            *appgoodbasecrud.Conds
	Offset                      int32
	Limit                       int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{
		AppSimulatePowerRentalConds: &appsimulatepowerrentalcrud.Conds{},
		AppGoodBaseConds:            &appgoodbasecrud.Conds{},
	}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid entid")
			}
			return nil
		}
		_id, err := uuid.Parse(*id)
		if err != nil {
			return err
		}
		h.EntID = &_id
		return nil
	}
}

func WithAppGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appgoodid")
			}
			return nil
		}
		handler, err := apppowerrental1.NewHandler(
			ctx,
			apppowerrental1.WithAppGoodID(id, true),
		)
		if err != nil {
			return err
		}
		exist, err := handler.ExistPowerRental(ctx)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid appgoodid")
		}
		h.AppGoodID = handler.AppGoodID
		return nil
	}
}

func WithOrderUnits(value *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == nil {
			if must {
				return fmt.Errorf("invalid units")
			}
			return nil
		}
		amount, err := decimal.NewFromString(*value)
		if err != nil {
			return err
		}
		if amount.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("units is less than or equal to 0")
		}
		h.OrderUnits = &amount
		return nil
	}
}

func WithOrderDurationSeconds(duration *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if duration == nil {
			if must {
				return fmt.Errorf("invalid duration")
			}
			return nil
		}
		if *duration <= uint32(0) {
			return fmt.Errorf("duration is less than or equal to 0")
		}
		h.OrderDurationSeconds = duration
		return nil
	}
}

func (h *Handler) withAppGoodBaseConds(conds *npool.Conds) error {
	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return err
		}
		h.AppGoodBaseConds.AppID = &cruder.Cond{
			Op:  conds.GetAppID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func (h *Handler) withAppSimulatePowerRentalConds(conds *npool.Conds) error {
	if conds.ID != nil {
		h.AppSimulatePowerRentalConds.ID = &cruder.Cond{
			Op:  conds.GetID().GetOp(),
			Val: conds.GetID().GetValue(),
		}
	}
	if conds.EntID != nil {
		id, err := uuid.Parse(conds.GetEntID().GetValue())
		if err != nil {
			return err
		}
		h.AppSimulatePowerRentalConds.EntID = &cruder.Cond{
			Op:  conds.GetEntID().GetOp(),
			Val: id,
		}
	}
	if conds.AppGoodID != nil {
		id, err := uuid.Parse(conds.GetAppGoodID().GetValue())
		if err != nil {
			return err
		}
		h.AppSimulatePowerRentalConds.AppGoodID = &cruder.Cond{
			Op:  conds.GetAppGoodID().GetOp(),
			Val: id,
		}
	}
	return nil
}

func WithConds(conds *npool.Conds) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if conds == nil {
			return nil
		}
		if err := h.withAppGoodBaseConds(conds); err != nil {
			return err
		}
		return h.withAppSimulatePowerRentalConds(conds)
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
