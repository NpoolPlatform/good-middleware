// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddisplaycolor"
	"github.com/google/uuid"
)

// AppGoodDisplayColor is the model entity for the AppGoodDisplayColor schema.
type AppGoodDisplayColor struct {
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
	// Color holds the value of the "color" field.
	Color string `json:"color,omitempty"`
	// Index holds the value of the "index" field.
	Index uint8 `json:"index,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppGoodDisplayColor) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appgooddisplaycolor.FieldID, appgooddisplaycolor.FieldCreatedAt, appgooddisplaycolor.FieldUpdatedAt, appgooddisplaycolor.FieldDeletedAt, appgooddisplaycolor.FieldIndex:
			values[i] = new(sql.NullInt64)
		case appgooddisplaycolor.FieldColor:
			values[i] = new(sql.NullString)
		case appgooddisplaycolor.FieldEntID, appgooddisplaycolor.FieldAppGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppGoodDisplayColor", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppGoodDisplayColor fields.
func (agdc *AppGoodDisplayColor) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appgooddisplaycolor.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			agdc.ID = uint32(value.Int64)
		case appgooddisplaycolor.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				agdc.CreatedAt = uint32(value.Int64)
			}
		case appgooddisplaycolor.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				agdc.UpdatedAt = uint32(value.Int64)
			}
		case appgooddisplaycolor.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				agdc.DeletedAt = uint32(value.Int64)
			}
		case appgooddisplaycolor.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				agdc.EntID = *value
			}
		case appgooddisplaycolor.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				agdc.AppGoodID = *value
			}
		case appgooddisplaycolor.FieldColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field color", values[i])
			} else if value.Valid {
				agdc.Color = value.String
			}
		case appgooddisplaycolor.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				agdc.Index = uint8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppGoodDisplayColor.
// Note that you need to call AppGoodDisplayColor.Unwrap() before calling this method if this AppGoodDisplayColor
// was returned from a transaction, and the transaction was committed or rolled back.
func (agdc *AppGoodDisplayColor) Update() *AppGoodDisplayColorUpdateOne {
	return (&AppGoodDisplayColorClient{config: agdc.config}).UpdateOne(agdc)
}

// Unwrap unwraps the AppGoodDisplayColor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (agdc *AppGoodDisplayColor) Unwrap() *AppGoodDisplayColor {
	_tx, ok := agdc.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppGoodDisplayColor is not a transactional entity")
	}
	agdc.config.driver = _tx.drv
	return agdc
}

// String implements the fmt.Stringer.
func (agdc *AppGoodDisplayColor) String() string {
	var builder strings.Builder
	builder.WriteString("AppGoodDisplayColor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", agdc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", agdc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", agdc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", agdc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", agdc.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", agdc.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("color=")
	builder.WriteString(agdc.Color)
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", agdc.Index))
	builder.WriteByte(')')
	return builder.String()
}

// AppGoodDisplayColors is a parsable slice of AppGoodDisplayColor.
type AppGoodDisplayColors []*AppGoodDisplayColor

func (agdc AppGoodDisplayColors) config(cfg config) {
	for _i := range agdc {
		agdc[_i].config = cfg
	}
}
