// Code generated by ent, DO NOT EDIT.

package extrainfo

import (
	"entgo.io/ent"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

const (
	// Label holds the string label denoting the extrainfo type in the database.
	Label = "extra_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldGoodID holds the string denoting the good_id field in the database.
	FieldGoodID = "good_id"
	// FieldPosters holds the string denoting the posters field in the database.
	FieldPosters = "posters"
	// FieldLabels holds the string denoting the labels field in the database.
	FieldLabels = "labels"
	// FieldLikes holds the string denoting the likes field in the database.
	FieldLikes = "likes"
	// FieldDislikes holds the string denoting the dislikes field in the database.
	FieldDislikes = "dislikes"
	// FieldRecommendCount holds the string denoting the recommend_count field in the database.
	FieldRecommendCount = "recommend_count"
	// FieldScoreCount holds the string denoting the score_count field in the database.
	FieldScoreCount = "score_count"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// Table holds the table name of the extrainfo in the database.
	Table = "extra_infos"
)

// Columns holds all SQL columns for extrainfo fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldGoodID,
	FieldPosters,
	FieldLabels,
	FieldLikes,
	FieldDislikes,
	FieldRecommendCount,
	FieldScoreCount,
	FieldScore,
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
	// DefaultPosters holds the default value on creation for the "posters" field.
	DefaultPosters []string
	// DefaultLabels holds the default value on creation for the "labels" field.
	DefaultLabels []string
	// DefaultLikes holds the default value on creation for the "likes" field.
	DefaultLikes uint32
	// DefaultDislikes holds the default value on creation for the "dislikes" field.
	DefaultDislikes uint32
	// DefaultRecommendCount holds the default value on creation for the "recommend_count" field.
	DefaultRecommendCount uint32
	// DefaultScoreCount holds the default value on creation for the "score_count" field.
	DefaultScoreCount uint32
	// DefaultScore holds the default value on creation for the "score" field.
	DefaultScore decimal.Decimal
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
