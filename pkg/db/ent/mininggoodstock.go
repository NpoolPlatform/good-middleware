// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/mininggoodstock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// MiningGoodStock is the model entity for the MiningGoodStock schema.
type MiningGoodStock struct {
	config `json:"-"`
	// ID of the ent.
	ID uint32 `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// EntID holds the value of the "ent_id" field.
	EntID uuid.UUID `json:"ent_id,omitempty"`
	// GoodStockID holds the value of the "good_stock_id" field.
	GoodStockID uuid.UUID `json:"good_stock_id,omitempty"`
	// PoolRootUserID holds the value of the "pool_root_user_id" field.
	PoolRootUserID uuid.UUID `json:"pool_root_user_id,omitempty"`
	// PoolGoodUserID holds the value of the "pool_good_user_id" field.
	PoolGoodUserID uuid.UUID `json:"pool_good_user_id,omitempty"`
	// Total holds the value of the "total" field.
	Total decimal.Decimal `json:"total,omitempty"`
	// SpotQuantity holds the value of the "spot_quantity" field.
	SpotQuantity decimal.Decimal `json:"spot_quantity,omitempty"`
	// Locked holds the value of the "locked" field.
	Locked decimal.Decimal `json:"locked,omitempty"`
	// InService holds the value of the "in_service" field.
	InService decimal.Decimal `json:"in_service,omitempty"`
	// WaitStart holds the value of the "wait_start" field.
	WaitStart decimal.Decimal `json:"wait_start,omitempty"`
	// Sold holds the value of the "sold" field.
	Sold decimal.Decimal `json:"sold,omitempty"`
	// AppReserved holds the value of the "app_reserved" field.
	AppReserved decimal.Decimal `json:"app_reserved,omitempty"`
	// State holds the value of the "state" field.
	State string `json:"state,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MiningGoodStock) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mininggoodstock.FieldTotal, mininggoodstock.FieldSpotQuantity, mininggoodstock.FieldLocked, mininggoodstock.FieldInService, mininggoodstock.FieldWaitStart, mininggoodstock.FieldSold, mininggoodstock.FieldAppReserved:
			values[i] = new(decimal.Decimal)
		case mininggoodstock.FieldID, mininggoodstock.FieldCreatedAt, mininggoodstock.FieldUpdatedAt, mininggoodstock.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case mininggoodstock.FieldState:
			values[i] = new(sql.NullString)
		case mininggoodstock.FieldEntID, mininggoodstock.FieldGoodStockID, mininggoodstock.FieldPoolRootUserID, mininggoodstock.FieldPoolGoodUserID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MiningGoodStock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MiningGoodStock fields.
func (mgs *MiningGoodStock) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mininggoodstock.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			mgs.ID = uint32(value.Int64)
		case mininggoodstock.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				mgs.CreatedAt = uint32(value.Int64)
			}
		case mininggoodstock.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				mgs.UpdatedAt = uint32(value.Int64)
			}
		case mininggoodstock.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				mgs.DeletedAt = uint32(value.Int64)
			}
		case mininggoodstock.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				mgs.EntID = *value
			}
		case mininggoodstock.FieldGoodStockID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_stock_id", values[i])
			} else if value != nil {
				mgs.GoodStockID = *value
			}
		case mininggoodstock.FieldPoolRootUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field pool_root_user_id", values[i])
			} else if value != nil {
				mgs.PoolRootUserID = *value
			}
		case mininggoodstock.FieldPoolGoodUserID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field pool_good_user_id", values[i])
			} else if value != nil {
				mgs.PoolGoodUserID = *value
			}
		case mininggoodstock.FieldTotal:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value != nil {
				mgs.Total = *value
			}
		case mininggoodstock.FieldSpotQuantity:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field spot_quantity", values[i])
			} else if value != nil {
				mgs.SpotQuantity = *value
			}
		case mininggoodstock.FieldLocked:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value != nil {
				mgs.Locked = *value
			}
		case mininggoodstock.FieldInService:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field in_service", values[i])
			} else if value != nil {
				mgs.InService = *value
			}
		case mininggoodstock.FieldWaitStart:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field wait_start", values[i])
			} else if value != nil {
				mgs.WaitStart = *value
			}
		case mininggoodstock.FieldSold:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field sold", values[i])
			} else if value != nil {
				mgs.Sold = *value
			}
		case mininggoodstock.FieldAppReserved:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field app_reserved", values[i])
			} else if value != nil {
				mgs.AppReserved = *value
			}
		case mininggoodstock.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				mgs.State = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MiningGoodStock.
// Note that you need to call MiningGoodStock.Unwrap() before calling this method if this MiningGoodStock
// was returned from a transaction, and the transaction was committed or rolled back.
func (mgs *MiningGoodStock) Update() *MiningGoodStockUpdateOne {
	return (&MiningGoodStockClient{config: mgs.config}).UpdateOne(mgs)
}

// Unwrap unwraps the MiningGoodStock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mgs *MiningGoodStock) Unwrap() *MiningGoodStock {
	_tx, ok := mgs.config.driver.(*txDriver)
	if !ok {
		panic("ent: MiningGoodStock is not a transactional entity")
	}
	mgs.config.driver = _tx.drv
	return mgs
}

// String implements the fmt.Stringer.
func (mgs *MiningGoodStock) String() string {
	var builder strings.Builder
	builder.WriteString("MiningGoodStock(")
	builder.WriteString(fmt.Sprintf("id=%v, ", mgs.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", mgs.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", mgs.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", mgs.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", mgs.EntID))
	builder.WriteString(", ")
	builder.WriteString("good_stock_id=")
	builder.WriteString(fmt.Sprintf("%v", mgs.GoodStockID))
	builder.WriteString(", ")
	builder.WriteString("pool_root_user_id=")
	builder.WriteString(fmt.Sprintf("%v", mgs.PoolRootUserID))
	builder.WriteString(", ")
	builder.WriteString("pool_good_user_id=")
	builder.WriteString(fmt.Sprintf("%v", mgs.PoolGoodUserID))
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", mgs.Total))
	builder.WriteString(", ")
	builder.WriteString("spot_quantity=")
	builder.WriteString(fmt.Sprintf("%v", mgs.SpotQuantity))
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", mgs.Locked))
	builder.WriteString(", ")
	builder.WriteString("in_service=")
	builder.WriteString(fmt.Sprintf("%v", mgs.InService))
	builder.WriteString(", ")
	builder.WriteString("wait_start=")
	builder.WriteString(fmt.Sprintf("%v", mgs.WaitStart))
	builder.WriteString(", ")
	builder.WriteString("sold=")
	builder.WriteString(fmt.Sprintf("%v", mgs.Sold))
	builder.WriteString(", ")
	builder.WriteString("app_reserved=")
	builder.WriteString(fmt.Sprintf("%v", mgs.AppReserved))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(mgs.State)
	builder.WriteByte(')')
	return builder.String()
}

// MiningGoodStocks is a parsable slice of MiningGoodStock.
type MiningGoodStocks []*MiningGoodStock

func (mgs MiningGoodStocks) config(cfg config) {
	for _i := range mgs {
		mgs[_i].config = cfg
	}
}
