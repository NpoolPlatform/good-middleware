// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/requiredgood"
	"github.com/google/uuid"
)

// RequiredGood is the model entity for the RequiredGood schema.
type RequiredGood struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt uint32 `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt uint32 `json:"updated_at,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt uint32 `json:"deleted_at,omitempty"`
	// MainGoodID holds the value of the "main_good_id" field.
	MainGoodID uuid.UUID `json:"main_good_id,omitempty"`
	// RequiredGoodID holds the value of the "required_good_id" field.
	RequiredGoodID uuid.UUID `json:"required_good_id,omitempty"`
	// Must holds the value of the "must" field.
	Must bool `json:"must,omitempty"`
	// Commission holds the value of the "commission" field.
	Commission bool `json:"commission,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RequiredGood) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case requiredgood.FieldMust, requiredgood.FieldCommission:
			values[i] = new(sql.NullBool)
		case requiredgood.FieldCreatedAt, requiredgood.FieldUpdatedAt, requiredgood.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case requiredgood.FieldID, requiredgood.FieldMainGoodID, requiredgood.FieldRequiredGoodID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RequiredGood", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RequiredGood fields.
func (rg *RequiredGood) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case requiredgood.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				rg.ID = *value
			}
		case requiredgood.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rg.CreatedAt = uint32(value.Int64)
			}
		case requiredgood.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				rg.UpdatedAt = uint32(value.Int64)
			}
		case requiredgood.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				rg.DeletedAt = uint32(value.Int64)
			}
		case requiredgood.FieldMainGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field main_good_id", values[i])
			} else if value != nil {
				rg.MainGoodID = *value
			}
		case requiredgood.FieldRequiredGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field required_good_id", values[i])
			} else if value != nil {
				rg.RequiredGoodID = *value
			}
		case requiredgood.FieldMust:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field must", values[i])
			} else if value.Valid {
				rg.Must = value.Bool
			}
		case requiredgood.FieldCommission:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field commission", values[i])
			} else if value.Valid {
				rg.Commission = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RequiredGood.
// Note that you need to call RequiredGood.Unwrap() before calling this method if this RequiredGood
// was returned from a transaction, and the transaction was committed or rolled back.
func (rg *RequiredGood) Update() *RequiredGoodUpdateOne {
	return (&RequiredGoodClient{config: rg.config}).UpdateOne(rg)
}

// Unwrap unwraps the RequiredGood entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rg *RequiredGood) Unwrap() *RequiredGood {
	_tx, ok := rg.config.driver.(*txDriver)
	if !ok {
		panic("ent: RequiredGood is not a transactional entity")
	}
	rg.config.driver = _tx.drv
	return rg
}

// String implements the fmt.Stringer.
func (rg *RequiredGood) String() string {
	var builder strings.Builder
	builder.WriteString("RequiredGood(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rg.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", rg.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", rg.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", rg.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("main_good_id=")
	builder.WriteString(fmt.Sprintf("%v", rg.MainGoodID))
	builder.WriteString(", ")
	builder.WriteString("required_good_id=")
	builder.WriteString(fmt.Sprintf("%v", rg.RequiredGoodID))
	builder.WriteString(", ")
	builder.WriteString("must=")
	builder.WriteString(fmt.Sprintf("%v", rg.Must))
	builder.WriteString(", ")
	builder.WriteString("commission=")
	builder.WriteString(fmt.Sprintf("%v", rg.Commission))
	builder.WriteByte(')')
	return builder.String()
}

// RequiredGoods is a parsable slice of RequiredGood.
type RequiredGoods []*RequiredGood

func (rg RequiredGoods) config(cfg config) {
	for _i := range rg {
		rg[_i].config = cfg
	}
}
