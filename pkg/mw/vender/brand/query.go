package brand

import (
	"context"
	"fmt"

	brandcrud "github.com/NpoolPlatform/good-middleware/pkg/crud/vender/brand"
	"github.com/NpoolPlatform/good-middleware/pkg/db"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent"
	entbrand "github.com/NpoolPlatform/good-middleware/pkg/db/ent/vendorbrand"
	npool "github.com/NpoolPlatform/message/npool/good/mw/v1/vender/brand"
)

type queryHandler struct {
	*Handler
	stmSelect *ent.VendorBrandSelect
	infos     []*npool.Brand
	total     uint32
}

func (h *queryHandler) selectVendorBrand(stm *ent.VendorBrandQuery) {
	h.stmSelect = stm.Select(
		entbrand.FieldID,
		entbrand.FieldEntID,
		entbrand.FieldName,
		entbrand.FieldLogo,
		entbrand.FieldCreatedAt,
		entbrand.FieldUpdatedAt,
	)
}

func (h *queryHandler) queryVendorBrand(cli *ent.Client) error {
	if h.ID == nil && h.EntID == nil {
		return fmt.Errorf("invalid id")
	}
	stm := cli.VendorBrand.Query().Where(entbrand.DeletedAt(0))
	if h.ID != nil {
		stm.Where(entbrand.ID(*h.ID))
	}
	if h.EntID != nil {
		stm.Where(entbrand.EntID(*h.EntID))
	}
	h.selectVendorBrand(stm)
	return nil
}

func (h *queryHandler) queryVendorBrands(ctx context.Context, cli *ent.Client) error {
	stm, err := brandcrud.SetQueryConds(cli.VendorBrand.Query(), h.Conds)
	if err != nil {
		return err
	}
	total, err := stm.Count(ctx)
	if err != nil {
		return err
	}
	h.total = uint32(total)
	h.selectVendorBrand(stm)
	return nil
}

func (h *queryHandler) scan(ctx context.Context) error {
	return h.stmSelect.Scan(ctx, &h.infos)
}

func (h *Handler) GetBrand(ctx context.Context) (*npool.Brand, error) {
	handler := &queryHandler{
		Handler: h,
	}

	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryVendorBrand(cli); err != nil {
			return err
		}
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, err
	}
	if len(handler.infos) == 0 {
		return nil, nil
	}
	if len(handler.infos) > 1 {
		return nil, fmt.Errorf("too many records")
	}
	return handler.infos[0], nil
}

func (h *Handler) GetBrands(ctx context.Context) ([]*npool.Brand, uint32, error) {
	handler := &queryHandler{
		Handler: h,
	}
	err := db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		if err := handler.queryVendorBrands(_ctx, cli); err != nil {
			return err
		}
		handler.stmSelect.
			Offset(int(h.Offset)).
			Limit(int(h.Limit))
		return handler.scan(_ctx)
	})
	if err != nil {
		return nil, 0, err
	}
	return handler.infos, handler.total, nil
}
