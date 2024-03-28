// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/topmostposter"
	"github.com/google/uuid"
)

// TopMostPoster is the model entity for the TopMostPoster schema.
type TopMostPoster struct {
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
	// TopMostID holds the value of the "top_most_id" field.
	TopMostID uuid.UUID `json:"top_most_id,omitempty"`
	// Poster holds the value of the "poster" field.
	Poster string `json:"poster,omitempty"`
	// Index holds the value of the "index" field.
	Index uint8 `json:"index,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TopMostPoster) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case topmostposter.FieldID, topmostposter.FieldCreatedAt, topmostposter.FieldUpdatedAt, topmostposter.FieldDeletedAt, topmostposter.FieldIndex:
			values[i] = new(sql.NullInt64)
		case topmostposter.FieldPoster:
			values[i] = new(sql.NullString)
		case topmostposter.FieldEntID, topmostposter.FieldTopMostID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type TopMostPoster", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TopMostPoster fields.
func (tmp *TopMostPoster) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case topmostposter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			tmp.ID = uint32(value.Int64)
		case topmostposter.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				tmp.CreatedAt = uint32(value.Int64)
			}
		case topmostposter.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				tmp.UpdatedAt = uint32(value.Int64)
			}
		case topmostposter.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				tmp.DeletedAt = uint32(value.Int64)
			}
		case topmostposter.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				tmp.EntID = *value
			}
		case topmostposter.FieldTopMostID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field top_most_id", values[i])
			} else if value != nil {
				tmp.TopMostID = *value
			}
		case topmostposter.FieldPoster:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field poster", values[i])
			} else if value.Valid {
				tmp.Poster = value.String
			}
		case topmostposter.FieldIndex:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field index", values[i])
			} else if value.Valid {
				tmp.Index = uint8(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this TopMostPoster.
// Note that you need to call TopMostPoster.Unwrap() before calling this method if this TopMostPoster
// was returned from a transaction, and the transaction was committed or rolled back.
func (tmp *TopMostPoster) Update() *TopMostPosterUpdateOne {
	return (&TopMostPosterClient{config: tmp.config}).UpdateOne(tmp)
}

// Unwrap unwraps the TopMostPoster entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (tmp *TopMostPoster) Unwrap() *TopMostPoster {
	_tx, ok := tmp.config.driver.(*txDriver)
	if !ok {
		panic("ent: TopMostPoster is not a transactional entity")
	}
	tmp.config.driver = _tx.drv
	return tmp
}

// String implements the fmt.Stringer.
func (tmp *TopMostPoster) String() string {
	var builder strings.Builder
	builder.WriteString("TopMostPoster(")
	builder.WriteString(fmt.Sprintf("id=%v, ", tmp.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", tmp.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", tmp.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", tmp.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", tmp.EntID))
	builder.WriteString(", ")
	builder.WriteString("top_most_id=")
	builder.WriteString(fmt.Sprintf("%v", tmp.TopMostID))
	builder.WriteString(", ")
	builder.WriteString("poster=")
	builder.WriteString(tmp.Poster)
	builder.WriteString(", ")
	builder.WriteString("index=")
	builder.WriteString(fmt.Sprintf("%v", tmp.Index))
	builder.WriteByte(')')
	return builder.String()
}

// TopMostPosters is a parsable slice of TopMostPoster.
type TopMostPosters []*TopMostPoster

func (tmp TopMostPosters) config(cfg config) {
	for _i := range tmp {
		tmp[_i].config = cfg
	}
}
