// Code generated by ent, DO NOT EDIT.

package topmostconstraint

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the topmostconstraint type in the database.
	Label = "top_most_constraint"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldEntID holds the string denoting the ent_id field in the database.
	FieldEntID = "ent_id"
	// FieldTopMostID holds the string denoting the top_most_id field in the database.
	FieldTopMostID = "top_most_id"
	// FieldConstraint holds the string denoting the constraint field in the database.
	FieldConstraint = "constraint"
	// FieldTargetValue holds the string denoting the target_value field in the database.
	FieldTargetValue = "target_value"
	// FieldIndex holds the string denoting the index field in the database.
	FieldIndex = "index"
	// Table holds the table name of the topmostconstraint in the database.
	Table = "top_most_constraints"
)

// Columns holds all SQL columns for topmostconstraint fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldTopMostID,
	FieldConstraint,
	FieldTargetValue,
	FieldIndex,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/NpoolPlatform/good-middleware/pkg/db/ent/runtime"
//
var (
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint32
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint32
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() uint32
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt func() uint32
	// DefaultEntID holds the default value on creation for the "ent_id" field.
	DefaultEntID func() uuid.UUID
	// DefaultTopMostID holds the default value on creation for the "top_most_id" field.
	DefaultTopMostID func() uuid.UUID
	// DefaultConstraint holds the default value on creation for the "constraint" field.
	DefaultConstraint string
	// DefaultTargetValue holds the default value on creation for the "target_value" field.
	DefaultTargetValue decimal.Decimal
	// DefaultIndex holds the default value on creation for the "index" field.
	DefaultIndex uint8
)
