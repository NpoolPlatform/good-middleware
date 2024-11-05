// Code generated by ent, DO NOT EDIT.

package goodbase

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the goodbase type in the database.
	Label = "good_base"
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
	// FieldGoodType holds the string denoting the good_type field in the database.
	FieldGoodType = "good_type"
	// FieldBenefitType holds the string denoting the benefit_type field in the database.
	FieldBenefitType = "benefit_type"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldServiceStartAt holds the string denoting the service_start_at field in the database.
	FieldServiceStartAt = "service_start_at"
	// FieldStartMode holds the string denoting the start_mode field in the database.
	FieldStartMode = "start_mode"
	// FieldTestOnly holds the string denoting the test_only field in the database.
	FieldTestOnly = "test_only"
	// FieldBenefitIntervalHours holds the string denoting the benefit_interval_hours field in the database.
	FieldBenefitIntervalHours = "benefit_interval_hours"
	// FieldPurchasable holds the string denoting the purchasable field in the database.
	FieldPurchasable = "purchasable"
	// FieldOnline holds the string denoting the online field in the database.
	FieldOnline = "online"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// Table holds the table name of the goodbase in the database.
	Table = "good_bases"
)

// Columns holds all SQL columns for goodbase fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldGoodType,
	FieldBenefitType,
	FieldName,
	FieldServiceStartAt,
	FieldStartMode,
	FieldTestOnly,
	FieldBenefitIntervalHours,
	FieldPurchasable,
	FieldOnline,
	FieldState,
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
	// DefaultGoodType holds the default value on creation for the "good_type" field.
	DefaultGoodType string
	// DefaultBenefitType holds the default value on creation for the "benefit_type" field.
	DefaultBenefitType string
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultServiceStartAt holds the default value on creation for the "service_start_at" field.
	DefaultServiceStartAt uint32
	// DefaultStartMode holds the default value on creation for the "start_mode" field.
	DefaultStartMode string
	// DefaultTestOnly holds the default value on creation for the "test_only" field.
	DefaultTestOnly bool
	// DefaultBenefitIntervalHours holds the default value on creation for the "benefit_interval_hours" field.
	DefaultBenefitIntervalHours uint32
	// DefaultPurchasable holds the default value on creation for the "purchasable" field.
	DefaultPurchasable bool
	// DefaultOnline holds the default value on creation for the "online" field.
	DefaultOnline bool
	// DefaultState holds the default value on creation for the "state" field.
	DefaultState string
)
