// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appfee"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// AppFee is the model entity for the AppFee schema.
type AppFee struct {
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
	// UnitValue holds the value of the "unit_value" field.
	UnitValue decimal.Decimal `json:"unit_value,omitempty"`
	// CancelMode holds the value of the "cancel_mode" field.
	CancelMode string `json:"cancel_mode,omitempty"`
	// MinOrderDurationSeconds holds the value of the "min_order_duration_seconds" field.
	MinOrderDurationSeconds uint32 `json:"min_order_duration_seconds,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppFee) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appfee.FieldUnitValue:
			values[i] = new(decimal.Decimal)
		case appfee.FieldID, appfee.FieldCreatedAt, appfee.FieldUpdatedAt, appfee.FieldDeletedAt, appfee.FieldMinOrderDurationSeconds:
			values[i] = new(sql.NullInt64)
		case appfee.FieldCancelMode:
			values[i] = new(sql.NullString)
		case appfee.FieldEntID, appfee.FieldAppGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppFee", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppFee fields.
func (af *AppFee) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appfee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			af.ID = uint32(value.Int64)
		case appfee.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				af.CreatedAt = uint32(value.Int64)
			}
		case appfee.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				af.UpdatedAt = uint32(value.Int64)
			}
		case appfee.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				af.DeletedAt = uint32(value.Int64)
			}
		case appfee.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				af.EntID = *value
			}
		case appfee.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				af.AppGoodID = *value
			}
		case appfee.FieldUnitValue:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field unit_value", values[i])
			} else if value != nil {
				af.UnitValue = *value
			}
		case appfee.FieldCancelMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cancel_mode", values[i])
			} else if value.Valid {
				af.CancelMode = value.String
			}
		case appfee.FieldMinOrderDurationSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field min_order_duration_seconds", values[i])
			} else if value.Valid {
				af.MinOrderDurationSeconds = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppFee.
// Note that you need to call AppFee.Unwrap() before calling this method if this AppFee
// was returned from a transaction, and the transaction was committed or rolled back.
func (af *AppFee) Update() *AppFeeUpdateOne {
	return (&AppFeeClient{config: af.config}).UpdateOne(af)
}

// Unwrap unwraps the AppFee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (af *AppFee) Unwrap() *AppFee {
	_tx, ok := af.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppFee is not a transactional entity")
	}
	af.config.driver = _tx.drv
	return af
}

// String implements the fmt.Stringer.
func (af *AppFee) String() string {
	var builder strings.Builder
	builder.WriteString("AppFee(")
	builder.WriteString(fmt.Sprintf("id=%v, ", af.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", af.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", af.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", af.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", af.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", af.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("unit_value=")
	builder.WriteString(fmt.Sprintf("%v", af.UnitValue))
	builder.WriteString(", ")
	builder.WriteString("cancel_mode=")
	builder.WriteString(af.CancelMode)
	builder.WriteString(", ")
	builder.WriteString("min_order_duration_seconds=")
	builder.WriteString(fmt.Sprintf("%v", af.MinOrderDurationSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// AppFees is a parsable slice of AppFee.
type AppFees []*AppFee

func (af AppFees) config(cfg config) {
	for _i := range af {
		af[_i].config = cfg
	}
}
