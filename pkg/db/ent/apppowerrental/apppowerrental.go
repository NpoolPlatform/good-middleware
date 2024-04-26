// Code generated by ent, DO NOT EDIT.

package apppowerrental

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the apppowerrental type in the database.
	Label = "app_power_rental"
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
	// FieldAppGoodID holds the string denoting the app_good_id field in the database.
	FieldAppGoodID = "app_good_id"
	// FieldServiceStartAt holds the string denoting the service_start_at field in the database.
	FieldServiceStartAt = "service_start_at"
	// FieldCancelMode holds the string denoting the cancel_mode field in the database.
	FieldCancelMode = "cancel_mode"
	// FieldCancelableBeforeStartSeconds holds the string denoting the cancelable_before_start_seconds field in the database.
	FieldCancelableBeforeStartSeconds = "cancelable_before_start_seconds"
	// FieldEnableSetCommission holds the string denoting the enable_set_commission field in the database.
	FieldEnableSetCommission = "enable_set_commission"
	// FieldMinOrderAmount holds the string denoting the min_order_amount field in the database.
	FieldMinOrderAmount = "min_order_amount"
	// FieldMaxOrderAmount holds the string denoting the max_order_amount field in the database.
	FieldMaxOrderAmount = "max_order_amount"
	// FieldMaxUserAmount holds the string denoting the max_user_amount field in the database.
	FieldMaxUserAmount = "max_user_amount"
	// FieldMinOrderDurationSeconds holds the string denoting the min_order_duration_seconds field in the database.
	FieldMinOrderDurationSeconds = "min_order_duration_seconds"
	// FieldMaxOrderDurationSeconds holds the string denoting the max_order_duration_seconds field in the database.
	FieldMaxOrderDurationSeconds = "max_order_duration_seconds"
	// FieldUnitPrice holds the string denoting the unit_price field in the database.
	FieldUnitPrice = "unit_price"
	// FieldSaleStartAt holds the string denoting the sale_start_at field in the database.
	FieldSaleStartAt = "sale_start_at"
	// FieldSaleEndAt holds the string denoting the sale_end_at field in the database.
	FieldSaleEndAt = "sale_end_at"
	// FieldSaleMode holds the string denoting the sale_mode field in the database.
	FieldSaleMode = "sale_mode"
	// FieldFixedDuration holds the string denoting the fixed_duration field in the database.
	FieldFixedDuration = "fixed_duration"
	// FieldPackageWithRequireds holds the string denoting the package_with_requireds field in the database.
	FieldPackageWithRequireds = "package_with_requireds"
	// Table holds the table name of the apppowerrental in the database.
	Table = "app_power_rentals"
)

// Columns holds all SQL columns for apppowerrental fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppGoodID,
	FieldServiceStartAt,
	FieldCancelMode,
	FieldCancelableBeforeStartSeconds,
	FieldEnableSetCommission,
	FieldMinOrderAmount,
	FieldMaxOrderAmount,
	FieldMaxUserAmount,
	FieldMinOrderDurationSeconds,
	FieldMaxOrderDurationSeconds,
	FieldUnitPrice,
	FieldSaleStartAt,
	FieldSaleEndAt,
	FieldSaleMode,
	FieldFixedDuration,
	FieldPackageWithRequireds,
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
	// DefaultAppGoodID holds the default value on creation for the "app_good_id" field.
	DefaultAppGoodID func() uuid.UUID
	// DefaultServiceStartAt holds the default value on creation for the "service_start_at" field.
	DefaultServiceStartAt uint32
	// DefaultCancelMode holds the default value on creation for the "cancel_mode" field.
	DefaultCancelMode string
	// DefaultCancelableBeforeStartSeconds holds the default value on creation for the "cancelable_before_start_seconds" field.
	DefaultCancelableBeforeStartSeconds uint32
	// DefaultEnableSetCommission holds the default value on creation for the "enable_set_commission" field.
	DefaultEnableSetCommission bool
	// DefaultMinOrderAmount holds the default value on creation for the "min_order_amount" field.
	DefaultMinOrderAmount decimal.Decimal
	// DefaultMaxOrderAmount holds the default value on creation for the "max_order_amount" field.
	DefaultMaxOrderAmount decimal.Decimal
	// DefaultMaxUserAmount holds the default value on creation for the "max_user_amount" field.
	DefaultMaxUserAmount decimal.Decimal
	// DefaultMinOrderDurationSeconds holds the default value on creation for the "min_order_duration_seconds" field.
	DefaultMinOrderDurationSeconds uint32
	// DefaultMaxOrderDurationSeconds holds the default value on creation for the "max_order_duration_seconds" field.
	DefaultMaxOrderDurationSeconds uint32
	// DefaultUnitPrice holds the default value on creation for the "unit_price" field.
	DefaultUnitPrice decimal.Decimal
	// DefaultSaleStartAt holds the default value on creation for the "sale_start_at" field.
	DefaultSaleStartAt uint32
	// DefaultSaleEndAt holds the default value on creation for the "sale_end_at" field.
	DefaultSaleEndAt uint32
	// DefaultSaleMode holds the default value on creation for the "sale_mode" field.
	DefaultSaleMode string
	// DefaultFixedDuration holds the default value on creation for the "fixed_duration" field.
	DefaultFixedDuration bool
	// DefaultPackageWithRequireds holds the default value on creation for the "package_with_requireds" field.
	DefaultPackageWithRequireds bool
)
