package appgood

import (
	"context"
	"fmt"
	"time"

	"github.com/shopspring/decimal"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	"github.com/NpoolPlatform/good-manager/pkg/db/ent/appgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID                     *uuid.UUID
	AppID                  *uuid.UUID
	GoodID                 *uuid.UUID
	Online                 *bool
	Visible                *bool
	GoodName               *string
	Price                  *decimal.Decimal
	DisplayIndex           *uint32
	PurchaseLimit          uint32
	CommissionPercent      string
	SaleStartAt            *uint32
	SaleEndAt              *uint32
	ServiceStartAt         *uint32
	Descriptions           *string
	GoodBanner             *string
	DisplayName            []string
	EnablePurchse          *bool
	EnableProductPage      *bool
	CancelMode             *types.CancelMode
	UserPurchaseLimit      *string // Single order purchase limit
	DisplayColors          []string
	CancellableBeforeStart *uint32 // Only could be canceled x seconds before start
	ProductPage            *string
	EnableSetCommission    *bool
}

//nolint:funlen,gocyclo
func CreateSet(c *ent.AppGoodCreate, req *Req) *ent.AppGoodCreate {
	if req.ID != nil {
		c.SetID(uuid.MustParse(req.GetID()))
	}
	if req.AppID != nil {
		c.SetAppID(uuid.MustParse(req.GetAppID()))
	}
	if req.GoodID != nil {
		c.SetGoodID(uuid.MustParse(req.GetGoodID()))
	}
	if req.Online != nil {
		c.SetOnline(req.GetOnline())
	}
	if req.Visible != nil {
		c.SetVisible(req.GetVisible())
	}
	if req.GoodName != nil {
		c.SetGoodName(req.GetGoodName())
	}
	if req.Price != nil {
		price, err := decimal.NewFromString(req.GetPrice())
		if err != nil {
			return nil, err
		}
		c.SetPrice(price)
	}
	if req.DisplayIndex != nil {
		c.SetDisplayIndex(req.GetDisplayIndex())
	}
	if req.PurchaseLimit != nil {
		c.SetPurchaseLimit(req.GetPurchaseLimit())
	}
	if req.CommissionPercent != nil {
		c.SetCommissionPercent(req.GetCommissionPercent())
	}
	if req.SaleStartAt != nil {
		c.SetSaleStartAt(req.GetSaleStartAt())
	}
	if req.SaleEndAt != nil {
		c.SetSaleEndAt(req.GetSaleEndAt())
	}
	if req.ServiceStartAt != nil {
		c.SetServiceStartAt(req.GetServiceStartAt())
	}
	if req.TechnicalFeeRatio != nil {
		c.SetTechnicalFeeRatio(req.GetTechnicalFeeRatio())
	}
	if req.ElectricityFeeRatio != nil {
		c.SetElectricityFeeRatio(req.GetElectricityFeeRatio())
	}
	if req.DailyRewardAmount != nil {
		amount, err := decimal.NewFromString(req.GetDailyRewardAmount())
		if err != nil {
			return nil, err
		}
		c.SetDailyRewardAmount(amount)
	}
	if req.CommissionSettleType != nil {
		c.SetCommissionSettleType(req.GetCommissionSettleType().String())
	}
	if req.Descriptions != nil {
		c.SetDescriptions(req.GetDescriptions())
	}
	if req.GoodBanner != nil {
		c.SetGoodBanner(req.GetGoodBanner())
	}
	displayNames := []string{}
	if req.DisplayNames != nil {
		displayNames = req.GetDisplayNames()
	}
	c.SetDisplayNames(displayNames)
	if req.EnablePurchase != nil {
		c.SetEnablePurchase(req.GetEnablePurchase())
	}
	if req.EnableProductPage != nil {
		c.SetEnableProductPage(req.GetEnableProductPage())
	}
	if req.CancelMode != nil {
		c.SetCancelMode(req.GetCancelMode().String())
	}
	userPurchaseLimit := decimal.NewFromInt(0)
	if req.UserPurchaseLimit != nil {
		userPurchaseLimit, err = decimal.NewFromString(req.GetUserPurchaseLimit())
		if err != nil {
			return nil, err
		}
	}
	c.SetUserPurchaseLimit(userPurchaseLimit)
	displayColors := []string{}
	if req.DisplayColors != nil {
		displayColors = req.GetDisplayColors()
	}
	c.SetDisplayColors(displayColors)
	if req.DisplayColors != nil {
		c.SetDisplayColors(req.GetDisplayColors())
	}
	if req.CancellableBeforeStart != nil {
		c.SetCancellableBeforeStart(req.GetCancellableBeforeStart())
	}
	if req.ProductPage != nil {
		c.SetProductPage(req.GetProductPage())
	}
	if req.EnableSetCommission != nil {
		c.SetEnableSetCommission(req.GetEnableSetCommission())
	}
	return c, nil
}

func Create(ctx context.Context, in *npool.AppGoodReq) (*ent.AppGood, error) {
	var info *ent.AppGood
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := cli.AppGood.Create()
		stm, err := CreateSet(c, in)
		if err != nil {
			return err
		}
		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.AppGoodReq) ([]*ent.AppGood, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.AppGood{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.AppGoodCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.AppGood.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.AppGood.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(u *ent.AppGoodUpdateOne, in *npool.AppGoodReq) (*ent.AppGoodUpdateOne, error) { //nolint
	if in.Online != nil {
		u.SetOnline(in.GetOnline())
	}
	if in.Visible != nil {
		u.SetVisible(in.GetVisible())
	}
	if in.GoodName != nil {
		u.SetGoodName(in.GetGoodName())
	}
	if in.Price != nil {
		price, err := decimal.NewFromString(in.GetPrice())
		if err != nil {
			return nil, err
		}
		u.SetPrice(price)
	}
	if in.DisplayIndex != nil {
		u.SetDisplayIndex(in.GetDisplayIndex())
	}
	if in.PurchaseLimit != nil {
		u.SetPurchaseLimit(in.GetPurchaseLimit())
	}
	if in.CommissionPercent != nil {
		u.SetCommissionPercent(in.GetCommissionPercent())
	}
	if in.SaleStartAt != nil {
		u.SetSaleStartAt(in.GetSaleStartAt())
	}
	if in.SaleEndAt != nil {
		u.SetSaleEndAt(in.GetSaleEndAt())
	}
	if in.ServiceStartAt != nil {
		u.SetServiceStartAt(in.GetServiceStartAt())
	}
	if in.TechnicalFeeRatio != nil {
		u.SetTechnicalFeeRatio(in.GetTechnicalFeeRatio())
	}
	if in.ElectricityFeeRatio != nil {
		u.SetElectricityFeeRatio(in.GetElectricityFeeRatio())
	}
	if in.DailyRewardAmount != nil {
		amount, err := decimal.NewFromString(in.GetDailyRewardAmount())
		if err != nil {
			return nil, err
		}
		u.SetDailyRewardAmount(amount)
	}
	if in.CommissionSettleType != nil {
		u.SetCommissionSettleType(in.GetCommissionSettleType().String())
	}
	if len(in.Descriptions) > 0 {
		u.SetDescriptions(in.GetDescriptions())
	}
	if in.GoodBanner != nil {
		u.SetGoodBanner(in.GetGoodBanner())
	}
	if len(in.DisplayNames) > 0 {
		u.SetDisplayNames(in.GetDisplayNames())
	}
	if in.EnablePurchase != nil {
		u.SetEnablePurchase(in.GetEnablePurchase())
	}
	if in.EnableProductPage != nil {
		u.SetEnableProductPage(in.GetEnableProductPage())
	}
	if in.CancelMode != nil {
		u.SetCancelMode(in.GetCancelMode().String())
	}
	if in.UserPurchaseLimit != nil {
		userPurchaseLimit, err := decimal.NewFromString(in.GetUserPurchaseLimit())
		if err != nil {
			return nil, err
		}
		u.SetUserPurchaseLimit(userPurchaseLimit)
	}
	if in.UserPurchaseLimit != nil {
		userPurchaseLimit, err := decimal.NewFromString(in.GetUserPurchaseLimit())
		if err != nil {
			return nil, err
		}
		u.SetUserPurchaseLimit(userPurchaseLimit)
	}
	if in.DisplayColors != nil {
		u.SetDisplayColors(in.GetDisplayColors())
	}
	if in.CancellableBeforeStart != nil {
		u.SetCancellableBeforeStart(in.GetCancellableBeforeStart())
	}
	if in.ProductPage != nil {
		u.SetProductPage(in.GetProductPage())
	}
	if in.EnableSetCommission != nil {
		u.SetEnableSetCommission(in.GetEnableSetCommission())
	}
	return u, nil
}

func Update(ctx context.Context, in *npool.AppGoodReq) (*ent.AppGood, error) {
	var info *ent.AppGood
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.AppGood.Query().Where(appgood.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := UpdateSet(info.Update(), in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.AppGood, error) {
	var info *ent.AppGood
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppGood.Query().Where(appgood.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.AppGoodQuery, error) {
	stm := cli.AppGood.Query()
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(appgood.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(appgood.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.GoodID != nil {
		switch conds.GetGoodID().GetOp() {
		case cruder.EQ:
			stm.Where(appgood.GoodID(uuid.MustParse(conds.GetGoodID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if len(conds.GetGoodIDs().GetValue()) > 0 {
		ids := []uuid.UUID{}
		for _, id := range conds.GetGoodIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
		switch conds.GetGoodIDs().GetOp() {
		case cruder.IN:
			stm.Where(appgood.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if len(conds.GetAppIDs().GetValue()) > 0 {
		ids := []uuid.UUID{}
		for _, id := range conds.GetAppIDs().GetValue() {
			ids = append(ids, uuid.MustParse(id))
		}
		switch conds.GetAppIDs().GetOp() {
		case cruder.IN:
			stm.Where(appgood.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.AppGood, int, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	rows := []*ent.AppGood{}
	var total int
	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(appgood.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.AppGood, error) {
	var info *ent.AppGood
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.AppGood.Query().Where(appgood.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id string) (*ent.AppGood, error) {
	var info *ent.AppGood
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.AppGood.UpdateOneID(uuid.MustParse(id)).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
