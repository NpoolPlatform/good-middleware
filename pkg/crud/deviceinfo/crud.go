package deviceinfo

import (
	"fmt"

	"github.com/NpoolPlatform/good-manager/pkg/db/ent"
	entdeviceinfo "github.com/NpoolPlatform/good-manager/pkg/db/ent/deviceinfo"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	"github.com/google/uuid"
)

type Req struct {
	ID              *uuid.UUID
	Type            *string
	Manufacturer    *string
	PowerComsuption *uint32
	ShipmentAt      *uint32
	Posters         []string
}

func CreateSet(c *ent.DeviceInfoCreate, req *Req) *ent.DeviceInfoCreate {
	if req.ID != nil {
		c.SetID(*req.ID)
	}
	if req.Type != nil {
		c.SetType(*req.Type)
	}
	if req.Manufacturer != nil {
		c.SetManufacturer(*req.Manufacturer)
	}
	if req.PowerComsuption != nil {
		c.SetPowerComsuption(*req.PowerComsuption)
	}
	if req.ShipmentAt != nil {
		c.SetShipmentAt(*req.ShipmentAt)
	}
	if len(req.Posters) > 0 {
		c.SetPosters(req.Posters)
	}
	return c
}

func UpdateSet(u *ent.DeviceInfoUpdateOne, req *Req) *ent.DeviceInfoUpdateOne {
	if len(req.Posters) > 0 {
		u.SetPosters(req.Posters)
	}
	return u
}

type Conds struct {
	ID           *cruder.Cond
	Type         *cruder.Cond
	Manufacturer *cruder.Cond
}

func SetQueryConds(q *ent.DeviceInfoQuery, conds *Conds) (*ent.DeviceInfoQuery, error) {
	q.Where(entdeviceinfo.DeletedAt(0))
	if conds == nil {
		return q, nil
	}
	if conds.ID != nil {
		id, ok := conds.ID.Val.(uuid.UUID)
		if !ok {
			return nil, fmt.Errorf("invalid id")
		}
		switch conds.ID.Op {
		case cruder.EQ:
			q.Where(entdeviceinfo.ID(id))
		default:
			return nil, fmt.Errorf("invalid deviceinfo field")
		}
	}
	if conds.Type != nil {
		_type, ok := conds.Type.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid type")
		}
		switch conds.Type.Op {
		case cruder.EQ:
			q.Where(entdeviceinfo.Type(_type))
		default:
			return nil, fmt.Errorf("invalid deviceinfo field")
		}
	}
	if conds.Manufacturer != nil {
		manufacturer, ok := conds.Manufacturer.Val.(string)
		if !ok {
			return nil, fmt.Errorf("invalid manufacturer")
		}
		switch conds.Manufacturer.Op {
		case cruder.EQ:
			q.Where(entdeviceinfo.Manufacturer(manufacturer))
		default:
			return nil, fmt.Errorf("invalid deviceinfo field")
		}
	}
	return q, nil
}
