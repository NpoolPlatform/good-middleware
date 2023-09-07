// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/good"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Good is the model entity for the Good schema.
type Good struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// DeviceInfoID holds the value of the "device_info_id" field.
	DeviceInfoID uuid.UUID `json:"device_info_id,omitempty"`
	// DurationDays holds the value of the "duration_days" field.
	DurationDays int32 `json:"duration_days,omitempty"`
	// CoinTypeID holds the value of the "coin_type_id" field.
	CoinTypeID uuid.UUID `json:"coin_type_id,omitempty"`
	// InheritFromGoodID holds the value of the "inherit_from_good_id" field.
	InheritFromGoodID uuid.UUID `json:"inherit_from_good_id,omitempty"`
	// VendorLocationID holds the value of the "vendor_location_id" field.
	VendorLocationID uuid.UUID `json:"vendor_location_id,omitempty"`
	// Price holds the value of the "price" field.
	Price decimal.Decimal `json:"price,omitempty"`
	// BenefitType holds the value of the "benefit_type" field.
	BenefitType string `json:"benefit_type,omitempty"`
	// GoodType holds the value of the "good_type" field.
	GoodType string `json:"good_type,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Unit holds the value of the "unit" field.
	Unit string `json:"unit,omitempty"`
	// UnitAmount holds the value of the "unit_amount" field.
	UnitAmount int32 `json:"unit_amount,omitempty"`
	// SupportCoinTypeIds holds the value of the "support_coin_type_ids" field.
	SupportCoinTypeIds []uuid.UUID `json:"support_coin_type_ids,omitempty"`
	// DeliveryAt holds the value of the "delivery_at" field.
	DeliveryAt uint32 `json:"delivery_at,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// StartMode holds the value of the "start_mode" field.
	StartMode string `json:"start_mode,omitempty"`
	// TestOnly holds the value of the "test_only" field.
	TestOnly bool `json:"test_only,omitempty"`
	// BenefitIntervalHours holds the value of the "benefit_interval_hours" field.
	BenefitIntervalHours uint32 `json:"benefit_interval_hours,omitempty"`
	// UnitLockDeposit holds the value of the "unit_lock_deposit" field.
	UnitLockDeposit decimal.Decimal `json:"unit_lock_deposit,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Good) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case good.FieldSupportCoinTypeIds:
			values[i] = new([]byte)
		case good.FieldPrice, good.FieldUnitLockDeposit:
			values[i] = new(decimal.Decimal)
		case good.FieldTestOnly:
			values[i] = new(sql.NullBool)
		case good.FieldCreatedAt, good.FieldUpdatedAt, good.FieldDeletedAt, good.FieldDurationDays, good.FieldUnitAmount, good.FieldDeliveryAt, good.FieldStartAt, good.FieldBenefitIntervalHours:
			values[i] = new(sql.NullInt64)
		case good.FieldBenefitType, good.FieldGoodType, good.FieldTitle, good.FieldUnit, good.FieldStartMode:
			values[i] = new(sql.NullString)
		case good.FieldID, good.FieldDeviceInfoID, good.FieldCoinTypeID, good.FieldInheritFromGoodID, good.FieldVendorLocationID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Good", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Good fields.
func (_go *Good) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case good.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				_go.ID = *value
			}
		case good.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_go.CreatedAt = uint32(value.Int64)
			}
		case good.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				_go.UpdatedAt = uint32(value.Int64)
			}
		case good.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				_go.DeletedAt = uint32(value.Int64)
			}
		case good.FieldDeviceInfoID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field device_info_id", values[i])
			} else if value != nil {
				_go.DeviceInfoID = *value
			}
		case good.FieldDurationDays:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_days", values[i])
			} else if value.Valid {
				_go.DurationDays = int32(value.Int64)
			}
		case good.FieldCoinTypeID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field coin_type_id", values[i])
			} else if value != nil {
				_go.CoinTypeID = *value
			}
		case good.FieldInheritFromGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field inherit_from_good_id", values[i])
			} else if value != nil {
				_go.InheritFromGoodID = *value
			}
		case good.FieldVendorLocationID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field vendor_location_id", values[i])
			} else if value != nil {
				_go.VendorLocationID = *value
			}
		case good.FieldPrice:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value != nil {
				_go.Price = *value
			}
		case good.FieldBenefitType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field benefit_type", values[i])
			} else if value.Valid {
				_go.BenefitType = value.String
			}
		case good.FieldGoodType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field good_type", values[i])
			} else if value.Valid {
				_go.GoodType = value.String
			}
		case good.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				_go.Title = value.String
			}
		case good.FieldUnit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field unit", values[i])
			} else if value.Valid {
				_go.Unit = value.String
			}
		case good.FieldUnitAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field unit_amount", values[i])
			} else if value.Valid {
				_go.UnitAmount = int32(value.Int64)
			}
		case good.FieldSupportCoinTypeIds:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field support_coin_type_ids", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &_go.SupportCoinTypeIds); err != nil {
					return fmt.Errorf("unmarshal field support_coin_type_ids: %w", err)
				}
			}
		case good.FieldDeliveryAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_at", values[i])
			} else if value.Valid {
				_go.DeliveryAt = uint32(value.Int64)
			}
		case good.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				_go.StartAt = uint32(value.Int64)
			}
		case good.FieldStartMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_mode", values[i])
			} else if value.Valid {
				_go.StartMode = value.String
			}
		case good.FieldTestOnly:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field test_only", values[i])
			} else if value.Valid {
				_go.TestOnly = value.Bool
			}
		case good.FieldBenefitIntervalHours:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field benefit_interval_hours", values[i])
			} else if value.Valid {
				_go.BenefitIntervalHours = uint32(value.Int64)
			}
		case good.FieldUnitLockDeposit:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field unit_lock_deposit", values[i])
			} else if value != nil {
				_go.UnitLockDeposit = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Good.
// Note that you need to call Good.Unwrap() before calling this method if this Good
// was returned from a transaction, and the transaction was committed or rolled back.
func (_go *Good) Update() *GoodUpdateOne {
	return (&GoodClient{config: _go.config}).UpdateOne(_go)
}

// Unwrap unwraps the Good entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_go *Good) Unwrap() *Good {
	_tx, ok := _go.config.driver.(*txDriver)
	if !ok {
		panic("ent: Good is not a transactional entity")
	}
	_go.config.driver = _tx.drv
	return _go
}

// String implements the fmt.Stringer.
func (_go *Good) String() string {
	var builder strings.Builder
	builder.WriteString("Good(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _go.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", _go.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", _go.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", _go.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("device_info_id=")
	builder.WriteString(fmt.Sprintf("%v", _go.DeviceInfoID))
	builder.WriteString(", ")
	builder.WriteString("duration_days=")
	builder.WriteString(fmt.Sprintf("%v", _go.DurationDays))
	builder.WriteString(", ")
	builder.WriteString("coin_type_id=")
	builder.WriteString(fmt.Sprintf("%v", _go.CoinTypeID))
	builder.WriteString(", ")
	builder.WriteString("inherit_from_good_id=")
	builder.WriteString(fmt.Sprintf("%v", _go.InheritFromGoodID))
	builder.WriteString(", ")
	builder.WriteString("vendor_location_id=")
	builder.WriteString(fmt.Sprintf("%v", _go.VendorLocationID))
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", _go.Price))
	builder.WriteString(", ")
	builder.WriteString("benefit_type=")
	builder.WriteString(_go.BenefitType)
	builder.WriteString(", ")
	builder.WriteString("good_type=")
	builder.WriteString(_go.GoodType)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(_go.Title)
	builder.WriteString(", ")
	builder.WriteString("unit=")
	builder.WriteString(_go.Unit)
	builder.WriteString(", ")
	builder.WriteString("unit_amount=")
	builder.WriteString(fmt.Sprintf("%v", _go.UnitAmount))
	builder.WriteString(", ")
	builder.WriteString("support_coin_type_ids=")
	builder.WriteString(fmt.Sprintf("%v", _go.SupportCoinTypeIds))
	builder.WriteString(", ")
	builder.WriteString("delivery_at=")
	builder.WriteString(fmt.Sprintf("%v", _go.DeliveryAt))
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", _go.StartAt))
	builder.WriteString(", ")
	builder.WriteString("start_mode=")
	builder.WriteString(_go.StartMode)
	builder.WriteString(", ")
	builder.WriteString("test_only=")
	builder.WriteString(fmt.Sprintf("%v", _go.TestOnly))
	builder.WriteString(", ")
	builder.WriteString("benefit_interval_hours=")
	builder.WriteString(fmt.Sprintf("%v", _go.BenefitIntervalHours))
	builder.WriteString(", ")
	builder.WriteString("unit_lock_deposit=")
	builder.WriteString(fmt.Sprintf("%v", _go.UnitLockDeposit))
	builder.WriteByte(')')
	return builder.String()
}

// Goods is a parsable slice of Good.
type Goods []*Good

func (_go Goods) config(cfg config) {
	for _i := range _go {
		_go[_i].config = cfg
	}
}
