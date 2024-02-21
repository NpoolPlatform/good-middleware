// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppDefaultGoodsColumns holds the columns for the "app_default_goods" table.
	AppDefaultGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppDefaultGoodsTable holds the schema information for the "app_default_goods" table.
	AppDefaultGoodsTable = &schema.Table{
		Name:       "app_default_goods",
		Columns:    AppDefaultGoodsColumns,
		PrimaryKey: []*schema.Column{AppDefaultGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appdefaultgood_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppDefaultGoodsColumns[4]},
			},
		},
	}
	// AppGoodsColumns holds the columns for the "app_goods" table.
	AppGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "online", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "visible", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "good_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "unit_price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "package_price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "display_index", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "sale_start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "sale_end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "service_start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "technical_fee_ratio", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "electricity_fee_ratio", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "descriptions", Type: field.TypeJSON, Nullable: true},
		{Name: "good_banner", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "display_names", Type: field.TypeJSON, Nullable: true},
		{Name: "enable_purchase", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "enable_product_page", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "cancel_mode", Type: field.TypeString, Nullable: true, Default: "Uncancellable"},
		{Name: "display_colors", Type: field.TypeJSON, Nullable: true},
		{Name: "cancellable_before_start", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "product_page", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "enable_set_commission", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
		{Name: "min_order_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "max_order_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "max_user_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "min_order_duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "max_order_duration", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "package_with_requireds", Type: field.TypeBool, Nullable: true, Default: true},
		{Name: "enable_simulate", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// AppGoodsTable holds the schema information for the "app_goods" table.
	AppGoodsTable = &schema.Table{
		Name:       "app_goods",
		Columns:    AppGoodsColumns,
		PrimaryKey: []*schema.Column{AppGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appgood_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppGoodsColumns[4]},
			},
			{
				Name:    "appgood_good_id_app_id_online",
				Unique:  false,
				Columns: []*schema.Column{AppGoodsColumns[6], AppGoodsColumns[5], AppGoodsColumns[7]},
			},
		},
	}
	// AppSimulateGoodsColumns holds the columns for the "app_simulate_goods" table.
	AppSimulateGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "coin_type_id", Type: field.TypeUUID, Nullable: true},
	}
	// AppSimulateGoodsTable holds the schema information for the "app_simulate_goods" table.
	AppSimulateGoodsTable = &schema.Table{
		Name:       "app_simulate_goods",
		Columns:    AppSimulateGoodsColumns,
		PrimaryKey: []*schema.Column{AppSimulateGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appsimulategood_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppSimulateGoodsColumns[4]},
			},
		},
	}
	// AppStocksColumns holds the columns for the "app_stocks" table.
	AppStocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "reserved", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "spot_quantity", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "locked", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "in_service", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "wait_start", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "sold", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// AppStocksTable holds the schema information for the "app_stocks" table.
	AppStocksTable = &schema.Table{
		Name:       "app_stocks",
		Columns:    AppStocksColumns,
		PrimaryKey: []*schema.Column{AppStocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appstock_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppStocksColumns[4]},
			},
		},
	}
	// AppStockLocksColumns holds the columns for the "app_stock_locks" table.
	AppStockLocksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_stock_id", Type: field.TypeUUID, Nullable: true},
		{Name: "units", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "app_spot_units", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "lock_state", Type: field.TypeString, Nullable: true, Default: "AppStockLocked"},
		{Name: "charge_back_state", Type: field.TypeString, Nullable: true, Default: "DefaultAppStockLockState"},
	}
	// AppStockLocksTable holds the schema information for the "app_stock_locks" table.
	AppStockLocksTable = &schema.Table{
		Name:       "app_stock_locks",
		Columns:    AppStockLocksColumns,
		PrimaryKey: []*schema.Column{AppStockLocksColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appstocklock_ent_id",
				Unique:  true,
				Columns: []*schema.Column{AppStockLocksColumns[4]},
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "order_id", Type: field.TypeUUID, Nullable: true},
		{Name: "content", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "reply_to_id", Type: field.TypeUUID, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "comment_ent_id",
				Unique:  true,
				Columns: []*schema.Column{CommentsColumns[4]},
			},
			{
				Name:    "comment_app_id_user_id_good_id_order_id",
				Unique:  false,
				Columns: []*schema.Column{CommentsColumns[5], CommentsColumns[6], CommentsColumns[7], CommentsColumns[9]},
			},
		},
	}
	// DeviceInfosColumns holds the columns for the "device_infos" table.
	DeviceInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "type", Type: field.TypeString, Nullable: true, Size: 64, Default: ""},
		{Name: "manufacturer", Type: field.TypeString, Nullable: true, Size: 64, Default: ""},
		{Name: "power_consumption", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "shipment_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
	}
	// DeviceInfosTable holds the schema information for the "device_infos" table.
	DeviceInfosTable = &schema.Table{
		Name:       "device_infos",
		Columns:    DeviceInfosColumns,
		PrimaryKey: []*schema.Column{DeviceInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "deviceinfo_ent_id",
				Unique:  true,
				Columns: []*schema.Column{DeviceInfosColumns[4]},
			},
		},
	}
	// ExtraInfosColumns holds the columns for the "extra_infos" table.
	ExtraInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
		{Name: "labels", Type: field.TypeJSON, Nullable: true},
		{Name: "likes", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "dislikes", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "recommend_count", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "comment_count", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "score_count", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "score", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// ExtraInfosTable holds the schema information for the "extra_infos" table.
	ExtraInfosTable = &schema.Table{
		Name:       "extra_infos",
		Columns:    ExtraInfosColumns,
		PrimaryKey: []*schema.Column{ExtraInfosColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "extrainfo_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ExtraInfosColumns[4]},
			},
			{
				Name:    "extrainfo_good_id",
				Unique:  false,
				Columns: []*schema.Column{ExtraInfosColumns[5]},
			},
		},
	}
	// GoodsColumns holds the columns for the "goods" table.
	GoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "device_info_id", Type: field.TypeUUID},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "inherit_from_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "vendor_location_id", Type: field.TypeUUID},
		{Name: "unit_price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "benefit_type", Type: field.TypeString, Nullable: true, Default: "DefaultBenefitType"},
		{Name: "good_type", Type: field.TypeString, Nullable: true, Default: "DefaultGoodType"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "quantity_unit", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "unit_amount", Type: field.TypeInt32, Nullable: true, Default: 0},
		{Name: "quantity_unit_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "delivery_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "start_mode", Type: field.TypeString, Nullable: true, Default: "GoodStartModeNextDay"},
		{Name: "test_only", Type: field.TypeBool, Nullable: true, Default: false},
		{Name: "benefit_interval_hours", Type: field.TypeUint32, Nullable: true, Default: 24},
		{Name: "unit_lock_deposit", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "unit_type", Type: field.TypeString, Nullable: true, Default: "GoodUnitByDurationAndQuantity"},
		{Name: "quantity_calculate_type", Type: field.TypeString, Nullable: true, Default: "GoodUnitCalculateBySelf"},
		{Name: "duration_type", Type: field.TypeString, Nullable: true, Default: "GoodDurationByDay"},
		{Name: "duration_calculate_type", Type: field.TypeString, Nullable: true, Default: "GoodUnitCalculateBySelf"},
		{Name: "settlement_type", Type: field.TypeString, Nullable: true, Default: "GoodSettledByCash"},
	}
	// GoodsTable holds the schema information for the "goods" table.
	GoodsTable = &schema.Table{
		Name:       "goods",
		Columns:    GoodsColumns,
		PrimaryKey: []*schema.Column{GoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "good_ent_id",
				Unique:  true,
				Columns: []*schema.Column{GoodsColumns[4]},
			},
		},
	}
	// GoodRewardsColumns holds the columns for the "good_rewards" table.
	GoodRewardsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "reward_state", Type: field.TypeString, Nullable: true, Default: "BenefitWait"},
		{Name: "last_reward_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "reward_tid", Type: field.TypeUUID, Nullable: true},
		{Name: "next_reward_start_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "last_reward_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "last_unit_reward_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "total_reward_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// GoodRewardsTable holds the schema information for the "good_rewards" table.
	GoodRewardsTable = &schema.Table{
		Name:       "good_rewards",
		Columns:    GoodRewardsColumns,
		PrimaryKey: []*schema.Column{GoodRewardsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "goodreward_ent_id",
				Unique:  true,
				Columns: []*schema.Column{GoodRewardsColumns[4]},
			},
		},
	}
	// GoodRewardHistoriesColumns holds the columns for the "good_reward_histories" table.
	GoodRewardHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "reward_date", Type: field.TypeUint32, Nullable: true},
		{Name: "tid", Type: field.TypeUUID, Nullable: true},
		{Name: "amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "unit_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "unit_net_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// GoodRewardHistoriesTable holds the schema information for the "good_reward_histories" table.
	GoodRewardHistoriesTable = &schema.Table{
		Name:       "good_reward_histories",
		Columns:    GoodRewardHistoriesColumns,
		PrimaryKey: []*schema.Column{GoodRewardHistoriesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "goodrewardhistory_ent_id",
				Unique:  true,
				Columns: []*schema.Column{GoodRewardHistoriesColumns[4]},
			},
		},
	}
	// LikesColumns holds the columns for the "likes" table.
	LikesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "like", Type: field.TypeBool},
	}
	// LikesTable holds the schema information for the "likes" table.
	LikesTable = &schema.Table{
		Name:       "likes",
		Columns:    LikesColumns,
		PrimaryKey: []*schema.Column{LikesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "like_ent_id",
				Unique:  true,
				Columns: []*schema.Column{LikesColumns[4]},
			},
		},
	}
	// RecommendsColumns holds the columns for the "recommends" table.
	RecommendsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "recommender_id", Type: field.TypeUUID, Nullable: true},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "recommend_index", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// RecommendsTable holds the schema information for the "recommends" table.
	RecommendsTable = &schema.Table{
		Name:       "recommends",
		Columns:    RecommendsColumns,
		PrimaryKey: []*schema.Column{RecommendsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "recommend_ent_id",
				Unique:  true,
				Columns: []*schema.Column{RecommendsColumns[4]},
			},
			{
				Name:    "recommend_app_id_good_id",
				Unique:  false,
				Columns: []*schema.Column{RecommendsColumns[5], RecommendsColumns[6]},
			},
		},
	}
	// RequiredGoodsColumns holds the columns for the "required_goods" table.
	RequiredGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "main_good_id", Type: field.TypeUUID},
		{Name: "required_good_id", Type: field.TypeUUID},
		{Name: "must", Type: field.TypeBool, Nullable: true, Default: false},
	}
	// RequiredGoodsTable holds the schema information for the "required_goods" table.
	RequiredGoodsTable = &schema.Table{
		Name:       "required_goods",
		Columns:    RequiredGoodsColumns,
		PrimaryKey: []*schema.Column{RequiredGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "requiredgood_ent_id",
				Unique:  true,
				Columns: []*schema.Column{RequiredGoodsColumns[4]},
			},
			{
				Name:    "requiredgood_main_good_id_required_good_id",
				Unique:  false,
				Columns: []*schema.Column{RequiredGoodsColumns[5], RequiredGoodsColumns[6]},
			},
		},
	}
	// ScoresColumns holds the columns for the "scores" table.
	ScoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID, Nullable: true},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "app_good_id", Type: field.TypeUUID, Nullable: true},
		{Name: "score", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// ScoresTable holds the schema information for the "scores" table.
	ScoresTable = &schema.Table{
		Name:       "scores",
		Columns:    ScoresColumns,
		PrimaryKey: []*schema.Column{ScoresColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "score_ent_id",
				Unique:  true,
				Columns: []*schema.Column{ScoresColumns[4]},
			},
		},
	}
	// StocksV1Columns holds the columns for the "stocks_v1" table.
	StocksV1Columns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "total", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "spot_quantity", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "locked", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "in_service", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "wait_start", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "sold", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "app_reserved", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// StocksV1Table holds the schema information for the "stocks_v1" table.
	StocksV1Table = &schema.Table{
		Name:       "stocks_v1",
		Columns:    StocksV1Columns,
		PrimaryKey: []*schema.Column{StocksV1Columns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "stock_ent_id",
				Unique:  true,
				Columns: []*schema.Column{StocksV1Columns[4]},
			},
		},
	}
	// TopMostsColumns holds the columns for the "top_mosts" table.
	TopMostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "top_most_type", Type: field.TypeString, Nullable: true, Default: "DefaultGoodTopMostType"},
		{Name: "title", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "message", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
		{Name: "start_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "end_at", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "threshold_credits", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "register_elapsed_seconds", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "threshold_purchases", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "threshold_payment_amount", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "kyc_must", Type: field.TypeBool, Nullable: true, Default: true},
	}
	// TopMostsTable holds the schema information for the "top_mosts" table.
	TopMostsTable = &schema.Table{
		Name:       "top_mosts",
		Columns:    TopMostsColumns,
		PrimaryKey: []*schema.Column{TopMostsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "topmost_ent_id",
				Unique:  true,
				Columns: []*schema.Column{TopMostsColumns[4]},
			},
		},
	}
	// TopMostGoodsColumns holds the columns for the "top_most_goods" table.
	TopMostGoodsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "good_id", Type: field.TypeUUID},
		{Name: "app_good_id", Type: field.TypeUUID},
		{Name: "coin_type_id", Type: field.TypeUUID},
		{Name: "top_most_id", Type: field.TypeUUID},
		{Name: "display_index", Type: field.TypeUint32, Nullable: true, Default: 0},
		{Name: "posters", Type: field.TypeJSON, Nullable: true},
		{Name: "unit_price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
		{Name: "package_price", Type: field.TypeOther, Nullable: true, SchemaType: map[string]string{"mysql": "decimal(37,18)"}},
	}
	// TopMostGoodsTable holds the schema information for the "top_most_goods" table.
	TopMostGoodsTable = &schema.Table{
		Name:       "top_most_goods",
		Columns:    TopMostGoodsColumns,
		PrimaryKey: []*schema.Column{TopMostGoodsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "topmostgood_ent_id",
				Unique:  true,
				Columns: []*schema.Column{TopMostGoodsColumns[4]},
			},
		},
	}
	// VendorBrandsColumns holds the columns for the "vendor_brands" table.
	VendorBrandsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "logo", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
	}
	// VendorBrandsTable holds the schema information for the "vendor_brands" table.
	VendorBrandsTable = &schema.Table{
		Name:       "vendor_brands",
		Columns:    VendorBrandsColumns,
		PrimaryKey: []*schema.Column{VendorBrandsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vendorbrand_ent_id",
				Unique:  true,
				Columns: []*schema.Column{VendorBrandsColumns[4]},
			},
		},
	}
	// VendorLocationsColumns holds the columns for the "vendor_locations" table.
	VendorLocationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint32, Increment: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "ent_id", Type: field.TypeUUID, Unique: true},
		{Name: "country", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "province", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "city", Type: field.TypeString, Nullable: true, Size: 128, Default: ""},
		{Name: "address", Type: field.TypeString, Nullable: true, Size: 256, Default: ""},
		{Name: "brand_id", Type: field.TypeUUID, Nullable: true},
	}
	// VendorLocationsTable holds the schema information for the "vendor_locations" table.
	VendorLocationsTable = &schema.Table{
		Name:       "vendor_locations",
		Columns:    VendorLocationsColumns,
		PrimaryKey: []*schema.Column{VendorLocationsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "vendorlocation_ent_id",
				Unique:  true,
				Columns: []*schema.Column{VendorLocationsColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppDefaultGoodsTable,
		AppGoodsTable,
		AppSimulateGoodsTable,
		AppStocksTable,
		AppStockLocksTable,
		CommentsTable,
		DeviceInfosTable,
		ExtraInfosTable,
		GoodsTable,
		GoodRewardsTable,
		GoodRewardHistoriesTable,
		LikesTable,
		RecommendsTable,
		RequiredGoodsTable,
		ScoresTable,
		StocksV1Table,
		TopMostsTable,
		TopMostGoodsTable,
		VendorBrandsTable,
		VendorLocationsTable,
	}
)

func init() {
	StocksV1Table.Annotation = &entsql.Annotation{
		Table: "stocks_v1",
	}
}
