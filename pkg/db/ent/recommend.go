// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/good-middleware/pkg/db/ent/recommend"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Recommend is the model entity for the Recommend schema.
type Recommend struct {
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
	// AppID holds the value of the "app_id" field.
	AppID uuid.UUID `json:"app_id,omitempty"`
	// GoodID holds the value of the "good_id" field.
	GoodID uuid.UUID `json:"good_id,omitempty"`
	// RecommenderID holds the value of the "recommender_id" field.
	RecommenderID uuid.UUID `json:"recommender_id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// RecommendIndex holds the value of the "recommend_index" field.
	RecommendIndex decimal.Decimal `json:"recommend_index,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Recommend) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case recommend.FieldRecommendIndex:
			values[i] = new(decimal.Decimal)
		case recommend.FieldID, recommend.FieldCreatedAt, recommend.FieldUpdatedAt, recommend.FieldDeletedAt:
			values[i] = new(sql.NullInt64)
		case recommend.FieldMessage:
			values[i] = new(sql.NullString)
		case recommend.FieldEntID, recommend.FieldAppID, recommend.FieldGoodID, recommend.FieldRecommenderID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Recommend", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Recommend fields.
func (r *Recommend) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recommend.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint32(value.Int64)
		case recommend.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = uint32(value.Int64)
			}
		case recommend.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = uint32(value.Int64)
			}
		case recommend.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				r.DeletedAt = uint32(value.Int64)
			}
		case recommend.FieldEntID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field ent_id", values[i])
			} else if value != nil {
				r.EntID = *value
			}
		case recommend.FieldAppID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field app_id", values[i])
			} else if value != nil {
				r.AppID = *value
			}
		case recommend.FieldGoodID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field good_id", values[i])
			} else if value != nil {
				r.GoodID = *value
			}
		case recommend.FieldRecommenderID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field recommender_id", values[i])
			} else if value != nil {
				r.RecommenderID = *value
			}
		case recommend.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				r.Message = value.String
			}
		case recommend.FieldRecommendIndex:
			if value, ok := values[i].(*decimal.Decimal); !ok {
				return fmt.Errorf("unexpected type %T for field recommend_index", values[i])
			} else if value != nil {
				r.RecommendIndex = *value
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Recommend.
// Note that you need to call Recommend.Unwrap() before calling this method if this Recommend
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Recommend) Update() *RecommendUpdateOne {
	return (&RecommendClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Recommend entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Recommend) Unwrap() *Recommend {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Recommend is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Recommend) String() string {
	var builder strings.Builder
	builder.WriteString("Recommend(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", r.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", r.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("ent_id=")
	builder.WriteString(fmt.Sprintf("%v", r.EntID))
	builder.WriteString(", ")
	builder.WriteString("app_id=")
	builder.WriteString(fmt.Sprintf("%v", r.AppID))
	builder.WriteString(", ")
	builder.WriteString("good_id=")
	builder.WriteString(fmt.Sprintf("%v", r.GoodID))
	builder.WriteString(", ")
	builder.WriteString("recommender_id=")
	builder.WriteString(fmt.Sprintf("%v", r.RecommenderID))
	builder.WriteString(", ")
	builder.WriteString("message=")
	builder.WriteString(r.Message)
	builder.WriteString(", ")
	builder.WriteString("recommend_index=")
	builder.WriteString(fmt.Sprintf("%v", r.RecommendIndex))
	builder.WriteByte(')')
	return builder.String()
}

// Recommends is a parsable slice of Recommend.
type Recommends []*Recommend

func (r Recommends) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
