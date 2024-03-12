// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/appgooddescription"
	"github.com/google/uuid"
)

// AppGoodDescription is the model entity for the AppGoodDescription schema.
type AppGoodDescription struct {
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
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Index holds the value of the "index" field.
	Index uint8 `json:"index,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*AppGoodDescription) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case appgooddescription.FieldID, appgooddescription.FieldCreatedAt, appgooddescription.FieldUpdatedAt, appgooddescription.FieldDeletedAt, appgooddescription.FieldIndex:
			values[i] = new(sql.NullInt64)
		case appgooddescription.FieldDescription:
			values[i] = new(sql.NullString)
		case appgooddescription.FieldEntID, appgooddescription.FieldAppGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type AppGoodDescription", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the AppGoodDescription fields.
func (agd *AppGoodDescription) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case appgooddescription.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			agd.ID = uint32(value.Int64)
		case appgooddescription.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				agd.CreatedAt = uint32(value.Int64)
			}
		case appgooddescription.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				agd.UpdatedAt = uint32(value.Int64)
			}
		case appgooddescription.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				agd.DeletedAt = uint32(value.Int64)
			}
		case appgooddescription.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				agd.EntID = *value
			}
		case appgooddescription.FieldAppGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_good_id", values[i])
			} else if value != nil {
				agd.AppGoodID = *value
			}
		case appgooddescription.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				agd.Description = value.String
			}
		case appgooddescription.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				agd.Index = uint8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this AppGoodDescription.
// Note that you need to call AppGoodDescription.Unwrap() before calling this method if this AppGoodDescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (agd *AppGoodDescription) Update() *AppGoodDescriptionUpdateOne {
	return (&AppGoodDescriptionClient{config: agd.config}).UpdateOne(agd)
}

// Unwrap unwraps the AppGoodDescription entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (agd *AppGoodDescription) Unwrap() *AppGoodDescription {
	_tx, ok := agd.config.driver.(*txDriver)
	if !ok {
		panic("ent: AppGoodDescription is not a transactional entity")
	}
	agd.config.driver = _tx.drv
	return agd
}

// String implements the fmt.Stringer.
func (agd *AppGoodDescription) String() string {
	var builder strings.Builder
	builder.WriteString("AppGoodDescription(")
	builder.WriteString(fmt.Sprintf("id=%v, ", agd.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", agd.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", agd.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", agd.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", agd.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_good_id=")
	builder.WriteString(fmt.Sprintf("%v", agd.AppGoodID))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(agd.Description)
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", agd.Index))
	builder.WriteByte(')')
	return builder.String()
}

// AppGoodDescriptions is a parsable slice of AppGoodDescription.
type AppGoodDescriptions []*AppGoodDescription

func (agd AppGoodDescriptions) config(cfg config) {
	for _i := range agd {
		agd[_i].config = cfg
	}
}
