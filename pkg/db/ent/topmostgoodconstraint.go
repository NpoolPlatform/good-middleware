// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostgoodconstraint"
	"github.com/google/uuid"
)

// TopMostGoodConstraint is the model entity for the TopMostGoodConstraint schema.
type TopMostGoodConstraint struct {
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
	// TopMostGoodID holds the value of the "top_most_good_id" field.
	TopMostGoodID uuid.UUID `json:"top_most_good_id,omitempty"`
	// Constraint holds the value of the "constraint" field.
	Constraint string `json:"constraint,omitempty"`
	// Index holds the value of the "index" field.
	Index uint8 `json:"index,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TopMostGoodConstraint) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case topmostgoodconstraint.FieldID, topmostgoodconstraint.FieldCreatedAt, topmostgoodconstraint.FieldUpdatedAt, topmostgoodconstraint.FieldDeletedAt, topmostgoodconstraint.FieldIndex:
			values[i] = new(sql.NullInt64)
		case topmostgoodconstraint.FieldConstraint:
			values[i] = new(sql.NullString)
		case topmostgoodconstraint.FieldEntID, topmostgoodconstraint.FieldTopMostGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TopMostGoodConstraint", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TopMostGoodConstraint fields.
func (tmgc *TopMostGoodConstraint) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topmostgoodconstraint.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tmgc.ID = uint32(value.Int64)
		case topmostgoodconstraint.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tmgc.CreatedAt = uint32(value.Int64)
			}
		case topmostgoodconstraint.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tmgc.UpdatedAt = uint32(value.Int64)
			}
		case topmostgoodconstraint.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tmgc.DeletedAt = uint32(value.Int64)
			}
		case topmostgoodconstraint.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				tmgc.EntID = *value
			}
		case topmostgoodconstraint.FieldTopMostGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field top_most_good_id", values[i])
			} else if value != nil {
				tmgc.TopMostGoodID = *value
			}
		case topmostgoodconstraint.FieldConstraint:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field constraint", values[i])
			} else if value.Valid {
				tmgc.Constraint = value.String
			}
		case topmostgoodconstraint.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				tmgc.Index = uint8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TopMostGoodConstraint.
// Note that you need to call TopMostGoodConstraint.Unwrap() before calling this method if this TopMostGoodConstraint
// was returned from a transaction, and the transaction was committed or rolled back.
func (tmgc *TopMostGoodConstraint) Update() *TopMostGoodConstraintUpdateOne {
	return (&TopMostGoodConstraintClient{config: tmgc.config}).UpdateOne(tmgc)
}

// Unwrap unwraps the TopMostGoodConstraint entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tmgc *TopMostGoodConstraint) Unwrap() *TopMostGoodConstraint {
	_tx, ok := tmgc.config.driver.(*txDriver)
	if !ok {
		panic("ent: TopMostGoodConstraint is not a transactional entity")
	}
	tmgc.config.driver = _tx.drv
	return tmgc
}

// String implements the fmt.Stringer.
func (tmgc *TopMostGoodConstraint) String() string {
	var builder strings.Builder
	builder.WriteString("TopMostGoodConstraint(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tmgc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", tmgc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tmgc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", tmgc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", tmgc.EntID))
	builder.WriteString(", ")
	builder.WriteString("top_most_good_id=")
	builder.WriteString(fmt.Sprintf("%v", tmgc.TopMostGoodID))
	builder.WriteString(", ")
	builder.WriteString("constraint=")
	builder.WriteString(tmgc.Constraint)
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", tmgc.Index))
	builder.WriteByte(')')
	return builder.String()
}

// TopMostGoodConstraints is a parsable slice of TopMostGoodConstraint.
type TopMostGoodConstraints []*TopMostGoodConstraint

func (tmgc TopMostGoodConstraints) config(cfg config) {
	for _i := range tmgc {
		tmgc[_i].config = cfg
	}
}
