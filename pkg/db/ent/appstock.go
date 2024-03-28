// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appstock"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppStock is the model entity for the AppStock schema.
type AppStock struct {
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
	// AppGoodID holds the value of the "app_good_id" field.
	AppGoodID uuid.UUID `json:"app_good_id,omitempty"`
	// Reserved holds the value of the "reserved" field.
	Reserved decimal.Decimal `json:"reserved,omitempty"`
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
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppStock) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appstock.FieldReserved, appstock.FieldSpotQuantity, appstock.FieldLocked, appstock.FieldInService, appstock.FieldWaitStart, appstock.FieldSold:
			values[i] = new(decimal.Decimal)
		case appstock.FieldID, appstock.FieldCreatedAt, appstock.FieldUpdatedAt, appstock.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case appstock.FieldEntID, appstock.FieldAppGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppStock", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppStock fields.
func (as *AppStock) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appstock.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			as.ID = uint32(value.Int64)
		case appstock.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				as.CreatedAt = uint32(value.Int64)
			}
		case appstock.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				as.UpdatedAt = uint32(value.Int64)
			}
		case appstock.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				as.DeletedAt = uint32(value.Int64)
			}
		case appstock.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				as.EntID = *value
			}
		case appstock.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				as.AppGoodID = *value
			}
		case appstock.FieldReserved:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field reserved", values[i])
			} else if value != nil {
				as.Reserved = *value
			}
		case appstock.FieldSpotQuantity:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field spot_quantity", values[i])
			} else if value != nil {
				as.SpotQuantity = *value
			}
		case appstock.FieldLocked:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value != nil {
				as.Locked = *value
			}
		case appstock.FieldInService:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field in_service", values[i])
			} else if value != nil {
				as.InService = *value
			}
		case appstock.FieldWaitStart:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field wait_start", values[i])
			} else if value != nil {
				as.WaitStart = *value
			}
		case appstock.FieldSold:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field sold", values[i])
			} else if value != nil {
				as.Sold = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppStock.
// Note that you need to call AppStock.Unwrap() before calling this method if this AppStock
// was returned from a transaction, and the transaction was committed or rolled back.
func (as *AppStock) Update() *AppStockUpdateOne {
	return (&AppStockClient{config: as.config}).UpdateOne(as)
}

// Unwrap unwraps the AppStock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (as *AppStock) Unwrap() *AppStock {
	_tx, ok := as.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppStock is not a transactional entity")
	}
	as.config.driver = _tx.drv
	return as
}

// String implements the fmt.Stringer.
func (as *AppStock) String() string {
	var builder strings.Builder
	builder.WriteString("AppStock(")
	builder.WriteString(fmt.Sprintf("id=%v, ", as.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", as.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", as.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", as.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", as.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", as.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("reserved=")
	builder.WriteString(fmt.Sprintf("%v", as.Reserved))
	builder.WriteString(", ")
	builder.WriteString("spot_quantity=")
	builder.WriteString(fmt.Sprintf("%v", as.SpotQuantity))
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", as.Locked))
	builder.WriteString(", ")
	builder.WriteString("in_service=")
	builder.WriteString(fmt.Sprintf("%v", as.InService))
	builder.WriteString(", ")
	builder.WriteString("wait_start=")
	builder.WriteString(fmt.Sprintf("%v", as.WaitStart))
	builder.WriteString(", ")
	builder.WriteString("sold=")
	builder.WriteString(fmt.Sprintf("%v", as.Sold))
	builder.WriteByte(')')
	return builder.String()
}

// AppStocks is a parsable slice of AppStock.
type AppStocks []*AppStock

func (as AppStocks) config(cfg config) {
	for _i := range as {
		as[_i].config = cfg
	}
}
