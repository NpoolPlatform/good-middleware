// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/fee"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Fee is the model entity for the Fee schema.
type Fee struct {
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
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// SettlementType holds the value of the "settlement_type" field.
	SettlementType string `json:"settlement_type,omitempty"`
	// UnitValue holds the value of the "unit_value" field.
	UnitValue decimal.Decimal `json:"unit_value,omitempty"`
	// DurationType holds the value of the "duration_type" field.
	DurationType string `json:"duration_type,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Fee) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case fee.FieldUnitValue:
			values[i] = new(decimal.Decimal)
		case fee.FieldID, fee.FieldCreatedAt, fee.FieldUpdatedAt, fee.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case fee.FieldSettlementType, fee.FieldDurationType:
			values[i] = new(sql.NullString)
		case fee.FieldEntID, fee.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Fee", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Fee fields.
func (f *Fee) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fee.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = uint32(value.Int64)
		case fee.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = uint32(value.Int64)
			}
		case fee.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = uint32(value.Int64)
			}
		case fee.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				f.DeletedAt = uint32(value.Int64)
			}
		case fee.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				f.EntID = *value
			}
		case fee.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				f.GoodID = *value
			}
		case fee.FieldSettlementType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field settlement_type", values[i])
			} else if value.Valid {
				f.SettlementType = value.String
			}
		case fee.FieldUnitValue:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field unit_value", values[i])
			} else if value != nil {
				f.UnitValue = *value
			}
		case fee.FieldDurationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field duration_type", values[i])
			} else if value.Valid {
				f.DurationType = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Fee.
// Note that you need to call Fee.Unwrap() before calling this method if this Fee
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Fee) Update() *FeeUpdateOne {
	return (&FeeClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the Fee entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Fee) Unwrap() *Fee {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Fee is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Fee) String() string {
	var builder strings.Builder
	builder.WriteString("Fee(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", f.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", f.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", f.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", f.EntID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", f.GoodID))
	builder.WriteString(", ")
	builder.WriteString("settlement_type=")
	builder.WriteString(f.SettlementType)
	builder.WriteString(", ")
	builder.WriteString("unit_value=")
	builder.WriteString(fmt.Sprintf("%v", f.UnitValue))
	builder.WriteString(", ")
	builder.WriteString("duration_type=")
	builder.WriteString(f.DurationType)
	builder.WriteByte(')')
	return builder.String()
}

// Fees is a parsable slice of Fee.
type Fees []*Fee

func (f Fees) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}