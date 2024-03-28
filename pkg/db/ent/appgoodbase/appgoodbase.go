// Code generated by ent, DO NOT EDIT.

package appgoodbase

import (
	"entgo.io/ent"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the appgoodbase type in the database.
	Label = "app_good_base"
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
	// FieldAppID holds the string denoting the app_id field in the database.
	FieldAppID = "app_id"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldPurchasable holds the string denoting the purchasable field in the database.
	FieldPurchasable = "purchasable"
	// FieldEnableProductPage holds the string denoting the enable_product_page field in the database.
	FieldEnableProductPage = "enable_product_page"
	// FieldProductPage holds the string denoting the product_page field in the database.
	FieldProductPage = "product_page"
	// FieldOnline holds the string denoting the online field in the database.
	FieldOnline = "online"
	// FieldVisible holds the string denoting the visible field in the database.
	FieldVisible = "visible"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDisplayIndex holds the string denoting the display_index field in the database.
	FieldDisplayIndex = "display_index"
	// FieldBanner holds the string denoting the banner field in the database.
	FieldBanner = "banner"
	// Table holds the table name of the appgoodbase in the database.
	Table = "app_good_bases"
)

// Columns holds all SQL columns for appgoodbase fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldEntID,
	FieldAppID,
	FieldGoodID,
	FieldPurchasable,
	FieldEnableProductPage,
	FieldProductPage,
	FieldOnline,
	FieldVisible,
	FieldName,
	FieldDisplayIndex,
	FieldBanner,
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
	// DefaultAppID holds the default value on creation for the "app_id" field.
	DefaultAppID func() uuid.UUID
	// DefaultGoodID holds the default value on creation for the "good_id" field.
	DefaultGoodID func() uuid.UUID
	// DefaultPurchasable holds the default value on creation for the "purchasable" field.
	DefaultPurchasable bool
	// DefaultEnableProductPage holds the default value on creation for the "enable_product_page" field.
	DefaultEnableProductPage bool
	// DefaultProductPage holds the default value on creation for the "product_page" field.
	DefaultProductPage string
	// DefaultOnline holds the default value on creation for the "online" field.
	DefaultOnline bool
	// DefaultVisible holds the default value on creation for the "visible" field.
	DefaultVisible bool
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// DefaultDisplayIndex holds the default value on creation for the "display_index" field.
	DefaultDisplayIndex int32
	// DefaultBanner holds the default value on creation for the "banner" field.
	DefaultBanner string
)