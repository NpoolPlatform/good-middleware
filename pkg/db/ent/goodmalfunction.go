// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/goodmalfunction"
	"github.com/google/uuid"
)

// GoodMalfunction is the model entity for the GoodMalfunction schema.
type GoodMalfunction struct {
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
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// StartAt holds the value of the "start_at" field.
	StartAt uint32 `json:"start_at,omitempty"`
	// DurationSeconds holds the value of the "duration_seconds" field.
	DurationSeconds uint32 `json:"duration_seconds,omitempty"`
	// CompensateSeconds holds the value of the "compensate_seconds" field.
	CompensateSeconds uint32 `json:"compensate_seconds,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GoodMalfunction) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case goodmalfunction.FieldID, goodmalfunction.FieldCreatedAt, goodmalfunction.FieldUpdatedAt, goodmalfunction.FieldDeletedAt, goodmalfunction.FieldStartAt, goodmalfunction.FieldDurationSeconds, goodmalfunction.FieldCompensateSeconds:
			values[i] = new(sql.NullInt64)
		case goodmalfunction.FieldTitle, goodmalfunction.FieldMessage:
			values[i] = new(sql.NullString)
		case goodmalfunction.FieldEntID, goodmalfunction.FieldGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GoodMalfunction", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GoodMalfunction fields.
func (gm *GoodMalfunction) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goodmalfunction.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gm.ID = uint32(value.Int64)
		case goodmalfunction.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				gm.CreatedAt = uint32(value.Int64)
			}
		case goodmalfunction.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				gm.UpdatedAt = uint32(value.Int64)
			}
		case goodmalfunction.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				gm.DeletedAt = uint32(value.Int64)
			}
		case goodmalfunction.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				gm.EntID = *value
			}
		case goodmalfunction.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				gm.GoodID = *value
			}
		case goodmalfunction.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				gm.Title = value.String
			}
		case goodmalfunction.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				gm.Message = value.String
			}
		case goodmalfunction.FieldStartAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field start_at", values[i])
			} else if value.Valid {
				gm.StartAt = uint32(value.Int64)
			}
		case goodmalfunction.FieldDurationSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration_seconds", values[i])
			} else if value.Valid {
				gm.DurationSeconds = uint32(value.Int64)
			}
		case goodmalfunction.FieldCompensateSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field compensate_seconds", values[i])
			} else if value.Valid {
				gm.CompensateSeconds = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this GoodMalfunction.
// Note that you need to call GoodMalfunction.Unwrap() before calling this method if this GoodMalfunction
// was returned from a transaction, and the transaction was committed or rolled back.
func (gm *GoodMalfunction) Update() *GoodMalfunctionUpdateOne {
	return (&GoodMalfunctionClient{config: gm.config}).UpdateOne(gm)
}

// Unwrap unwraps the GoodMalfunction entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gm *GoodMalfunction) Unwrap() *GoodMalfunction {
	_tx, ok := gm.config.driver.(*txDriver)
	if !ok {
		panic("ent: GoodMalfunction is not a transactional entity")
	}
	gm.config.driver = _tx.drv
	return gm
}

// String implements the fmt.Stringer.
func (gm *GoodMalfunction) String() string {
	var builder strings.Builder
	builder.WriteString("GoodMalfunction(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", gm.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", gm.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", gm.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.EntID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", gm.GoodID))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(gm.Title)
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(gm.Message)
	builder.WriteString(", ")
	builder.WriteString("start_at=")
	builder.WriteString(fmt.Sprintf("%v", gm.StartAt))
	builder.WriteString(", ")
	builder.WriteString("duration_seconds=")
	builder.WriteString(fmt.Sprintf("%v", gm.DurationSeconds))
	builder.WriteString(", ")
	builder.WriteString("compensate_seconds=")
	builder.WriteString(fmt.Sprintf("%v", gm.CompensateSeconds))
	builder.WriteByte(')')
	return builder.String()
}

// GoodMalfunctions is a parsable slice of GoodMalfunction.
type GoodMalfunctions []*GoodMalfunction

func (gm GoodMalfunctions) config(cfg config) {
	for _i := range gm {
		gm[_i].config = cfg
	}
}