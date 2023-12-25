package appgood

import (
	"fmt"

	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entappgood "github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgood"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Req struct {
	ID                     *uint32
	EntID                  *uuid.UUID
	AppID                  *uuid.UUID
	GoodID                 *uuid.UUID
	Online                 *bool
	Visible                *bool
	GoodName               *string
	Price                  *decimal.Decimal
	DisplayIndex           *int32
	PurchaseLimit          *int32
	SaleStartAt            *uint32
	SaleEndAt              *uint32
	ServiceStartAt         *uint32
	Descriptions           []string
	GoodBanner             *string
	DisplayNames           []string
	EnablePurchase         *bool
	EnableProductPage      *bool
	CancelMode             *types.CancelMode
	UserPurchaseLimit      *decimal.Decimal // Single order purchase limit
	DisplayColors          []string
	CancellableBeforeStart *uint32 // Only could be canceled x seconds before start
	ProductPage            *string
	EnableSetCommission    *bool
	Posters                []string
	DeletedAt              *uint32
	TechniqueFeeRatio      *decimal.Decimal
	ElectricityFeeRatio    *decimal.Decimal
	MinOrderAmount         *decimal.Decimal
	MaxOrderAmount         *decimal.Decimal
	MaxUserAmount          *decimal.Decimal
	MinOrderDuration       *uint32
	MaxOrderDuration       *uint32
}

//nolint:gocyclo,funlen
func CreateSet(c *ent.AppGoodCreate, req *Req) *ent.AppGoodCreate {
	if req.EntID != nil {
		c.SetEntID(*req.EntID)
	}
	if req.AppID != nil {
		c.SetAppID(*req.AppID)
	}
	if req.GoodID != nil {
		c.SetGoodID(*req.GoodID)
	}
	if req.Online != nil {
		c.SetOnline(*req.Online)
	}
	if req.Visible != nil {
		c.SetVisible(*req.Visible)
	}
	if req.GoodName != nil {
		c.SetGoodName(*req.GoodName)
	}
	if req.Price != nil {
		c.SetPrice(*req.Price)
	}
	if req.DisplayIndex != nil {
		c.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.PurchaseLimit != nil {
		c.SetPurchaseLimit(*req.PurchaseLimit)
	}
	if req.SaleStartAt != nil {
		c.SetSaleStartAt(*req.SaleStartAt)
	}
	if req.SaleEndAt != nil {
		c.SetSaleEndAt(*req.SaleEndAt)
	}
	if req.ServiceStartAt != nil {
		c.SetServiceStartAt(*req.ServiceStartAt)
	}
	if len(req.Descriptions) > 0 {
		c.SetDescriptions(req.Descriptions)
	}
	if req.GoodBanner != nil {
		c.SetGoodBanner(*req.GoodBanner)
	}
	if len(req.DisplayNames) > 0 {
		c.SetDisplayNames(req.DisplayNames)
	}
	if req.EnablePurchase != nil {
		c.SetEnablePurchase(*req.EnablePurchase)
	}
	if req.EnableProductPage != nil {
		c.SetEnableProductPage(*req.EnableProductPage)
	}
	if req.CancelMode != nil {
		c.SetCancelMode(req.CancelMode.String())
	}
	if req.UserPurchaseLimit != nil {
		c.SetUserPurchaseLimit(*req.UserPurchaseLimit)
	}
	if len(req.DisplayColors) > 0 {
		c.SetDisplayColors(req.DisplayColors)
	}
	if req.CancellableBeforeStart != nil {
		c.SetCancellableBeforeStart(*req.CancellableBeforeStart)
	}
	if req.ProductPage != nil {
		c.SetProductPage(*req.ProductPage)
	}
	if req.EnableSetCommission != nil {
		c.SetEnableSetCommission(*req.EnableSetCommission)
	}
	if len(req.Posters) > 0 {
		c.SetPosters(req.Posters)
	}
	if req.TechniqueFeeRatio != nil {
		c.SetTechnicalFeeRatio(*req.TechniqueFeeRatio)
	}
	if req.ElectricityFeeRatio != nil {
		c.SetElectricityFeeRatio(*req.ElectricityFeeRatio)
	}
	if req.MinOrderAmount != nil {
		c.SetMinOrderAmount(*req.MinOrderAmount)
	}
	if req.MaxOrderAmount != nil {
		c.SetMaxOrderAmount(*req.MaxOrderAmount)
	}
	if req.MaxUserAmount != nil {
		c.SetMaxUserAmount(*req.MaxUserAmount)
	}
	if req.MinOrderDuration != nil {
		c.SetMinOrderDuration(*req.MinOrderDuration)
	}
	if req.MaxOrderDuration != nil {
		c.SetMaxOrderDuration(*req.MaxOrderDuration)
	}
	return c
}

//nolint:gocyclo,funlen
func UpdateSet(u *ent.AppGoodUpdateOne, req *Req) *ent.AppGoodUpdateOne {
	if req.Online != nil {
		u.SetOnline(*req.Online)
	}
	if req.Visible != nil {
		u.SetVisible(*req.Visible)
	}
	if req.GoodName != nil {
		u.SetGoodName(*req.GoodName)
	}
	if req.Price != nil {
		u.SetPrice(*req.Price)
	}
	if req.DisplayIndex != nil {
		u.SetDisplayIndex(*req.DisplayIndex)
	}
	if req.PurchaseLimit != nil {
		u.SetPurchaseLimit(*req.PurchaseLimit)
	}
	if req.SaleStartAt != nil {
		u.SetSaleStartAt(*req.SaleStartAt)
	}
	if req.SaleEndAt != nil {
		u.SetSaleEndAt(*req.SaleEndAt)
	}
	if req.ServiceStartAt != nil {
		u.SetServiceStartAt(*req.ServiceStartAt)
	}
	if len(req.Descriptions) > 0 {
		u.SetDescriptions(req.Descriptions)
	}
	if req.GoodBanner != nil {
		u.SetGoodBanner(*req.GoodBanner)
	}
	if len(req.DisplayNames) > 0 {
		u.SetDisplayNames(req.DisplayNames)
	}
	if req.EnablePurchase != nil {
		u.SetEnablePurchase(*req.EnablePurchase)
	}
	if req.EnableProductPage != nil {
		u.SetEnableProductPage(*req.EnableProductPage)
	}
	if req.CancelMode != nil {
		u.SetCancelMode(req.CancelMode.String())
	}
	if req.UserPurchaseLimit != nil {
		u.SetUserPurchaseLimit(*req.UserPurchaseLimit)
	}
	if req.DisplayColors != nil {
		u.SetDisplayColors(req.DisplayColors)
	}
	if req.CancellableBeforeStart != nil {
		u.SetCancellableBeforeStart(*req.CancellableBeforeStart)
	}
	if req.ProductPage != nil {
		u.SetProductPage(*req.ProductPage)
	}
	if req.EnableSetCommission != nil {
		u.SetEnableSetCommission(*req.EnableSetCommission)
	}
	if len(req.Posters) > 0 {
		u.SetPosters(req.Posters)
	}
	if req.DeletedAt != nil {
		u.SetDeletedAt(*req.DeletedAt)
	}
	if req.TechniqueFeeRatio != nil {
		u.SetTechnicalFeeRatio(*req.TechniqueFeeRatio)
	}
	if req.ElectricityFeeRatio != nil {
		u.SetElectricityFeeRatio(*req.ElectricityFeeRatio)
	}
	if req.MinOrderAmount != nil {
		u.SetMinOrderAmount(*req.MinOrderAmount)
	}
	if req.MaxOrderAmount != nil {
		u.SetMaxOrderAmount(*req.MaxOrderAmount)
	}
	if req.MaxUserAmount != nil {
		u.SetMaxUserAmount(*req.MaxUserAmount)
	}
	if req.MinOrderDuration != nil {
		u.SetMinOrderDuration(*req.MinOrderDuration)
	}
	if req.MaxOrderDuration != nil {
		u.SetMaxOrderDuration(*req.MaxOrderDuration)
	}
	return u
}

type Conds struct {
	ID      *cruder.Cond
	EntID   *cruder.Cond
	AppID   *cruder.Cond
	GoodID  *cruder.Cond
	GoodIDs *cruder.Cond
	AppIDs  *cruder.Cond
	EntIDs  *cruder.Cond
	IDs     *cruder.Cond
}

//nolint:gocyclo,funlen
func SetQueryConds(q *ent.AppGoodQuery, conds *Conds) (*ent.AppGoodQuery, error) {
	q.Where(entappgood.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uint32)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entappgood.ID(id))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.EntID != nil {
		id, ok := conds.EntID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.EntID.Op {
		case cruder.EQ:
			q.Where(entappgood.EntID(id))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.AppID != nil {
		id, ok := conds.AppID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appid")
		}
		switch conds.AppID.Op {
		case cruder.EQ:
			q.Where(entappgood.AppID(id))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.GoodID != nil {
		id, ok := conds.GoodID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodid")
		}
		switch conds.GoodID.Op {
		case cruder.EQ:
			q.Where(entappgood.GoodID(id))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.GoodIDs != nil {
		ids, ok := conds.GoodIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid goodids")
		}
		switch conds.GoodIDs.Op {
		case cruder.IN:
			q.Where(entappgood.GoodIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.AppIDs != nil {
		ids, ok := conds.AppIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid appids")
		}
		switch conds.AppIDs.Op {
		case cruder.IN:
			q.Where(entappgood.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.EntIDs != nil {
		ids, ok := conds.EntIDs.Val.([]uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.EntIDs.Op {
		case cruder.IN:
			q.Where(entappgood.EntIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	if conds.IDs != nil {
		ids, ok := conds.IDs.Val.([]uint32)
		if !ok {
			return nil, fmt.Errorf("invalid ids")
		}
		switch conds.IDs.Op {
		case cruder.IN:
			q.Where(entappgood.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid appgood field")
		}
	}
	return q, nil
}
