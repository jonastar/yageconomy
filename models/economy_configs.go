// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// EconomyConfig is an object representing the database table.
type EconomyConfig struct {
	GuildID                        int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Enabled                        bool             `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	Admins                         types.Int64Array `boil:"admins" json:"admins,omitempty" toml:"admins" yaml:"admins,omitempty"`
	CurrencyName                   string           `boil:"currency_name" json:"currency_name" toml:"currency_name" yaml:"currency_name"`
	CurrencyNamePlural             string           `boil:"currency_name_plural" json:"currency_name_plural" toml:"currency_name_plural" yaml:"currency_name_plural"`
	CurrencySymbol                 string           `boil:"currency_symbol" json:"currency_symbol" toml:"currency_symbol" yaml:"currency_symbol"`
	DailyFrequency                 int64            `boil:"daily_frequency" json:"daily_frequency" toml:"daily_frequency" yaml:"daily_frequency"`
	DailyAmount                    int64            `boil:"daily_amount" json:"daily_amount" toml:"daily_amount" yaml:"daily_amount"`
	ChatmoneyFrequency             int64            `boil:"chatmoney_frequency" json:"chatmoney_frequency" toml:"chatmoney_frequency" yaml:"chatmoney_frequency"`
	ChatmoneyAmountMin             int64            `boil:"chatmoney_amount_min" json:"chatmoney_amount_min" toml:"chatmoney_amount_min" yaml:"chatmoney_amount_min"`
	ChatmoneyAmountMax             int64            `boil:"chatmoney_amount_max" json:"chatmoney_amount_max" toml:"chatmoney_amount_max" yaml:"chatmoney_amount_max"`
	AutoPlantChannels              types.Int64Array `boil:"auto_plant_channels" json:"auto_plant_channels,omitempty" toml:"auto_plant_channels" yaml:"auto_plant_channels,omitempty"`
	AutoPlantMin                   int64            `boil:"auto_plant_min" json:"auto_plant_min" toml:"auto_plant_min" yaml:"auto_plant_min"`
	AutoPlantMax                   int64            `boil:"auto_plant_max" json:"auto_plant_max" toml:"auto_plant_max" yaml:"auto_plant_max"`
	AutoPlantChance                types.Decimal    `boil:"auto_plant_chance" json:"auto_plant_chance" toml:"auto_plant_chance" yaml:"auto_plant_chance"`
	StartBalance                   int64            `boil:"start_balance" json:"start_balance" toml:"start_balance" yaml:"start_balance"`
	FishingMaxWinAmount            int64            `boil:"fishing_max_win_amount" json:"fishing_max_win_amount" toml:"fishing_max_win_amount" yaml:"fishing_max_win_amount"`
	FishingMinWinAmount            int64            `boil:"fishing_min_win_amount" json:"fishing_min_win_amount" toml:"fishing_min_win_amount" yaml:"fishing_min_win_amount"`
	FishingCooldown                int              `boil:"fishing_cooldown" json:"fishing_cooldown" toml:"fishing_cooldown" yaml:"fishing_cooldown"`
	RobFine                        int              `boil:"rob_fine" json:"rob_fine" toml:"rob_fine" yaml:"rob_fine"`
	RobCooldown                    int              `boil:"rob_cooldown" json:"rob_cooldown" toml:"rob_cooldown" yaml:"rob_cooldown"`
	HeistServerCooldown            int              `boil:"heist_server_cooldown" json:"heist_server_cooldown" toml:"heist_server_cooldown" yaml:"heist_server_cooldown"`
	HeistFailedGamblingBanDuration int              `boil:"heist_failed_gambling_ban_duration" json:"heist_failed_gambling_ban_duration" toml:"heist_failed_gambling_ban_duration" yaml:"heist_failed_gambling_ban_duration"`
	HeistLastUsage                 time.Time        `boil:"heist_last_usage" json:"heist_last_usage" toml:"heist_last_usage" yaml:"heist_last_usage"`

	R *economyConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L economyConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EconomyConfigColumns = struct {
	GuildID                        string
	Enabled                        string
	Admins                         string
	CurrencyName                   string
	CurrencyNamePlural             string
	CurrencySymbol                 string
	DailyFrequency                 string
	DailyAmount                    string
	ChatmoneyFrequency             string
	ChatmoneyAmountMin             string
	ChatmoneyAmountMax             string
	AutoPlantChannels              string
	AutoPlantMin                   string
	AutoPlantMax                   string
	AutoPlantChance                string
	StartBalance                   string
	FishingMaxWinAmount            string
	FishingMinWinAmount            string
	FishingCooldown                string
	RobFine                        string
	RobCooldown                    string
	HeistServerCooldown            string
	HeistFailedGamblingBanDuration string
	HeistLastUsage                 string
}{
	GuildID:                        "guild_id",
	Enabled:                        "enabled",
	Admins:                         "admins",
	CurrencyName:                   "currency_name",
	CurrencyNamePlural:             "currency_name_plural",
	CurrencySymbol:                 "currency_symbol",
	DailyFrequency:                 "daily_frequency",
	DailyAmount:                    "daily_amount",
	ChatmoneyFrequency:             "chatmoney_frequency",
	ChatmoneyAmountMin:             "chatmoney_amount_min",
	ChatmoneyAmountMax:             "chatmoney_amount_max",
	AutoPlantChannels:              "auto_plant_channels",
	AutoPlantMin:                   "auto_plant_min",
	AutoPlantMax:                   "auto_plant_max",
	AutoPlantChance:                "auto_plant_chance",
	StartBalance:                   "start_balance",
	FishingMaxWinAmount:            "fishing_max_win_amount",
	FishingMinWinAmount:            "fishing_min_win_amount",
	FishingCooldown:                "fishing_cooldown",
	RobFine:                        "rob_fine",
	RobCooldown:                    "rob_cooldown",
	HeistServerCooldown:            "heist_server_cooldown",
	HeistFailedGamblingBanDuration: "heist_failed_gambling_ban_duration",
	HeistLastUsage:                 "heist_last_usage",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertypes_Int64Array struct{ field string }

func (w whereHelpertypes_Int64Array) EQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_Int64Array) NEQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_Int64Array) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_Int64Array) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpertypes_Int64Array) LT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Int64Array) LTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Int64Array) GT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Int64Array) GTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertypes_Decimal struct{ field string }

func (w whereHelpertypes_Decimal) EQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertypes_Decimal) NEQ(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertypes_Decimal) LT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Decimal) LTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Decimal) GT(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Decimal) GTE(x types.Decimal) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var EconomyConfigWhere = struct {
	GuildID                        whereHelperint64
	Enabled                        whereHelperbool
	Admins                         whereHelpertypes_Int64Array
	CurrencyName                   whereHelperstring
	CurrencyNamePlural             whereHelperstring
	CurrencySymbol                 whereHelperstring
	DailyFrequency                 whereHelperint64
	DailyAmount                    whereHelperint64
	ChatmoneyFrequency             whereHelperint64
	ChatmoneyAmountMin             whereHelperint64
	ChatmoneyAmountMax             whereHelperint64
	AutoPlantChannels              whereHelpertypes_Int64Array
	AutoPlantMin                   whereHelperint64
	AutoPlantMax                   whereHelperint64
	AutoPlantChance                whereHelpertypes_Decimal
	StartBalance                   whereHelperint64
	FishingMaxWinAmount            whereHelperint64
	FishingMinWinAmount            whereHelperint64
	FishingCooldown                whereHelperint
	RobFine                        whereHelperint
	RobCooldown                    whereHelperint
	HeistServerCooldown            whereHelperint
	HeistFailedGamblingBanDuration whereHelperint
	HeistLastUsage                 whereHelpertime_Time
}{
	GuildID:                        whereHelperint64{field: `guild_id`},
	Enabled:                        whereHelperbool{field: `enabled`},
	Admins:                         whereHelpertypes_Int64Array{field: `admins`},
	CurrencyName:                   whereHelperstring{field: `currency_name`},
	CurrencyNamePlural:             whereHelperstring{field: `currency_name_plural`},
	CurrencySymbol:                 whereHelperstring{field: `currency_symbol`},
	DailyFrequency:                 whereHelperint64{field: `daily_frequency`},
	DailyAmount:                    whereHelperint64{field: `daily_amount`},
	ChatmoneyFrequency:             whereHelperint64{field: `chatmoney_frequency`},
	ChatmoneyAmountMin:             whereHelperint64{field: `chatmoney_amount_min`},
	ChatmoneyAmountMax:             whereHelperint64{field: `chatmoney_amount_max`},
	AutoPlantChannels:              whereHelpertypes_Int64Array{field: `auto_plant_channels`},
	AutoPlantMin:                   whereHelperint64{field: `auto_plant_min`},
	AutoPlantMax:                   whereHelperint64{field: `auto_plant_max`},
	AutoPlantChance:                whereHelpertypes_Decimal{field: `auto_plant_chance`},
	StartBalance:                   whereHelperint64{field: `start_balance`},
	FishingMaxWinAmount:            whereHelperint64{field: `fishing_max_win_amount`},
	FishingMinWinAmount:            whereHelperint64{field: `fishing_min_win_amount`},
	FishingCooldown:                whereHelperint{field: `fishing_cooldown`},
	RobFine:                        whereHelperint{field: `rob_fine`},
	RobCooldown:                    whereHelperint{field: `rob_cooldown`},
	HeistServerCooldown:            whereHelperint{field: `heist_server_cooldown`},
	HeistFailedGamblingBanDuration: whereHelperint{field: `heist_failed_gambling_ban_duration`},
	HeistLastUsage:                 whereHelpertime_Time{field: `heist_last_usage`},
}

// EconomyConfigRels is where relationship names are stored.
var EconomyConfigRels = struct {
}{}

// economyConfigR is where relationships are stored.
type economyConfigR struct {
}

// NewStruct creates a new relationship struct
func (*economyConfigR) NewStruct() *economyConfigR {
	return &economyConfigR{}
}

// economyConfigL is where Load methods for each relationship are stored.
type economyConfigL struct{}

var (
	economyConfigColumns               = []string{"guild_id", "enabled", "admins", "currency_name", "currency_name_plural", "currency_symbol", "daily_frequency", "daily_amount", "chatmoney_frequency", "chatmoney_amount_min", "chatmoney_amount_max", "auto_plant_channels", "auto_plant_min", "auto_plant_max", "auto_plant_chance", "start_balance", "fishing_max_win_amount", "fishing_min_win_amount", "fishing_cooldown", "rob_fine", "rob_cooldown", "heist_server_cooldown", "heist_failed_gambling_ban_duration", "heist_last_usage"}
	economyConfigColumnsWithoutDefault = []string{"guild_id", "enabled", "admins", "currency_name", "currency_name_plural", "currency_symbol", "daily_frequency", "daily_amount", "chatmoney_frequency", "chatmoney_amount_min", "chatmoney_amount_max", "auto_plant_channels", "auto_plant_min", "auto_plant_max", "auto_plant_chance", "start_balance", "fishing_max_win_amount", "fishing_min_win_amount", "fishing_cooldown", "rob_fine", "rob_cooldown"}
	economyConfigColumnsWithDefault    = []string{"heist_server_cooldown", "heist_failed_gambling_ban_duration", "heist_last_usage"}
	economyConfigPrimaryKeyColumns     = []string{"guild_id"}
)

type (
	// EconomyConfigSlice is an alias for a slice of pointers to EconomyConfig.
	// This should generally be used opposed to []EconomyConfig.
	EconomyConfigSlice []*EconomyConfig

	economyConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	economyConfigType                 = reflect.TypeOf(&EconomyConfig{})
	economyConfigMapping              = queries.MakeStructMapping(economyConfigType)
	economyConfigPrimaryKeyMapping, _ = queries.BindMapping(economyConfigType, economyConfigMapping, economyConfigPrimaryKeyColumns)
	economyConfigInsertCacheMut       sync.RWMutex
	economyConfigInsertCache          = make(map[string]insertCache)
	economyConfigUpdateCacheMut       sync.RWMutex
	economyConfigUpdateCache          = make(map[string]updateCache)
	economyConfigUpsertCacheMut       sync.RWMutex
	economyConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single economyConfig record from the query using the global executor.
func (q economyConfigQuery) OneG(ctx context.Context) (*EconomyConfig, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single economyConfig record from the query.
func (q economyConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EconomyConfig, error) {
	o := &EconomyConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for economy_configs")
	}

	return o, nil
}

// AllG returns all EconomyConfig records from the query using the global executor.
func (q economyConfigQuery) AllG(ctx context.Context) (EconomyConfigSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all EconomyConfig records from the query.
func (q economyConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (EconomyConfigSlice, error) {
	var o []*EconomyConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EconomyConfig slice")
	}

	return o, nil
}

// CountG returns the count of all EconomyConfig records in the query, and panics on error.
func (q economyConfigQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all EconomyConfig records in the query.
func (q economyConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count economy_configs rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q economyConfigQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q economyConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if economy_configs exists")
	}

	return count > 0, nil
}

// EconomyConfigs retrieves all the records using an executor.
func EconomyConfigs(mods ...qm.QueryMod) economyConfigQuery {
	mods = append(mods, qm.From("\"economy_configs\""))
	return economyConfigQuery{NewQuery(mods...)}
}

// FindEconomyConfigG retrieves a single record by ID.
func FindEconomyConfigG(ctx context.Context, guildID int64, selectCols ...string) (*EconomyConfig, error) {
	return FindEconomyConfig(ctx, boil.GetContextDB(), guildID, selectCols...)
}

// FindEconomyConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEconomyConfig(ctx context.Context, exec boil.ContextExecutor, guildID int64, selectCols ...string) (*EconomyConfig, error) {
	economyConfigObj := &EconomyConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"economy_configs\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(query, guildID)

	err := q.Bind(ctx, exec, economyConfigObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from economy_configs")
	}

	return economyConfigObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *EconomyConfig) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EconomyConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_configs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(economyConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	economyConfigInsertCacheMut.RLock()
	cache, cached := economyConfigInsertCache[key]
	economyConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			economyConfigColumns,
			economyConfigColumnsWithDefault,
			economyConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(economyConfigType, economyConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(economyConfigType, economyConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"economy_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"economy_configs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into economy_configs")
	}

	if !cached {
		economyConfigInsertCacheMut.Lock()
		economyConfigInsertCache[key] = cache
		economyConfigInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single EconomyConfig record using the global executor.
// See Update for more documentation.
func (o *EconomyConfig) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the EconomyConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EconomyConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	economyConfigUpdateCacheMut.RLock()
	cache, cached := economyConfigUpdateCache[key]
	economyConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			economyConfigColumns,
			economyConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update economy_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"economy_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, economyConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(economyConfigType, economyConfigMapping, append(wl, economyConfigPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update economy_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for economy_configs")
	}

	if !cached {
		economyConfigUpdateCacheMut.Lock()
		economyConfigUpdateCache[key] = cache
		economyConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q economyConfigQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q economyConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for economy_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for economy_configs")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o EconomyConfigSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EconomyConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"economy_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, economyConfigPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in economyConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all economyConfig")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *EconomyConfig) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EconomyConfig) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no economy_configs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(economyConfigColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	economyConfigUpsertCacheMut.RLock()
	cache, cached := economyConfigUpsertCache[key]
	economyConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			economyConfigColumns,
			economyConfigColumnsWithDefault,
			economyConfigColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			economyConfigColumns,
			economyConfigPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert economy_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(economyConfigPrimaryKeyColumns))
			copy(conflict, economyConfigPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"economy_configs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(economyConfigType, economyConfigMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(economyConfigType, economyConfigMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert economy_configs")
	}

	if !cached {
		economyConfigUpsertCacheMut.Lock()
		economyConfigUpsertCache[key] = cache
		economyConfigUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single EconomyConfig record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *EconomyConfig) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single EconomyConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EconomyConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EconomyConfig provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), economyConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"economy_configs\" WHERE \"guild_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from economy_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for economy_configs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q economyConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no economyConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economy_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_configs")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o EconomyConfigSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EconomyConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EconomyConfig slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"economy_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyConfigPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from economyConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for economy_configs")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *EconomyConfig) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no EconomyConfig provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *EconomyConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEconomyConfig(ctx, exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyConfigSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty EconomyConfigSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EconomyConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EconomyConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), economyConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"economy_configs\".* FROM \"economy_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, economyConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EconomyConfigSlice")
	}

	*o = slice

	return nil
}

// EconomyConfigExistsG checks if the EconomyConfig row exists.
func EconomyConfigExistsG(ctx context.Context, guildID int64) (bool, error) {
	return EconomyConfigExists(ctx, boil.GetContextDB(), guildID)
}

// EconomyConfigExists checks if the EconomyConfig row exists.
func EconomyConfigExists(ctx context.Context, exec boil.ContextExecutor, guildID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"economy_configs\" where \"guild_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if economy_configs exists")
	}

	return exists, nil
}
