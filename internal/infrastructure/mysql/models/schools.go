// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// School is an object representing the database table.
type School struct {
	ID    int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name  string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Style string `boil:"style" json:"style" toml:"style" yaml:"style"`
	Enemy string `boil:"enemy" json:"enemy" toml:"enemy" yaml:"enemy"`

	R *schoolR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L schoolL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SchoolColumns = struct {
	ID    string
	Name  string
	Style string
	Enemy string
}{
	ID:    "id",
	Name:  "name",
	Style: "style",
	Enemy: "enemy",
}

var SchoolTableColumns = struct {
	ID    string
	Name  string
	Style string
	Enemy string
}{
	ID:    "schools.id",
	Name:  "schools.name",
	Style: "schools.style",
	Enemy: "schools.enemy",
}

// Generated where

var SchoolWhere = struct {
	ID    whereHelperint
	Name  whereHelperstring
	Style whereHelperstring
	Enemy whereHelperstring
}{
	ID:    whereHelperint{field: "`schools`.`id`"},
	Name:  whereHelperstring{field: "`schools`.`name`"},
	Style: whereHelperstring{field: "`schools`.`style`"},
	Enemy: whereHelperstring{field: "`schools`.`enemy`"},
}

// SchoolRels is where relationship names are stored.
var SchoolRels = struct {
	LowerSchools string
}{
	LowerSchools: "LowerSchools",
}

// schoolR is where relationships are stored.
type schoolR struct {
	LowerSchools LowerSchoolSlice `boil:"LowerSchools" json:"LowerSchools" toml:"LowerSchools" yaml:"LowerSchools"`
}

// NewStruct creates a new relationship struct
func (*schoolR) NewStruct() *schoolR {
	return &schoolR{}
}

func (r *schoolR) GetLowerSchools() LowerSchoolSlice {
	if r == nil {
		return nil
	}
	return r.LowerSchools
}

// schoolL is where Load methods for each relationship are stored.
type schoolL struct{}

var (
	schoolAllColumns            = []string{"id", "name", "style", "enemy"}
	schoolColumnsWithoutDefault = []string{"name", "style", "enemy"}
	schoolColumnsWithDefault    = []string{"id"}
	schoolPrimaryKeyColumns     = []string{"id"}
	schoolGeneratedColumns      = []string{}
)

type (
	// SchoolSlice is an alias for a slice of pointers to School.
	// This should almost always be used instead of []School.
	SchoolSlice []*School
	// SchoolHook is the signature for custom School hook methods
	SchoolHook func(context.Context, boil.ContextExecutor, *School) error

	schoolQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	schoolType                 = reflect.TypeOf(&School{})
	schoolMapping              = queries.MakeStructMapping(schoolType)
	schoolPrimaryKeyMapping, _ = queries.BindMapping(schoolType, schoolMapping, schoolPrimaryKeyColumns)
	schoolInsertCacheMut       sync.RWMutex
	schoolInsertCache          = make(map[string]insertCache)
	schoolUpdateCacheMut       sync.RWMutex
	schoolUpdateCache          = make(map[string]updateCache)
	schoolUpsertCacheMut       sync.RWMutex
	schoolUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var schoolAfterSelectHooks []SchoolHook

var schoolBeforeInsertHooks []SchoolHook
var schoolAfterInsertHooks []SchoolHook

var schoolBeforeUpdateHooks []SchoolHook
var schoolAfterUpdateHooks []SchoolHook

var schoolBeforeDeleteHooks []SchoolHook
var schoolAfterDeleteHooks []SchoolHook

var schoolBeforeUpsertHooks []SchoolHook
var schoolAfterUpsertHooks []SchoolHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *School) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *School) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *School) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *School) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *School) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *School) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *School) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *School) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *School) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range schoolAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSchoolHook registers your hook function for all future operations.
func AddSchoolHook(hookPoint boil.HookPoint, schoolHook SchoolHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		schoolAfterSelectHooks = append(schoolAfterSelectHooks, schoolHook)
	case boil.BeforeInsertHook:
		schoolBeforeInsertHooks = append(schoolBeforeInsertHooks, schoolHook)
	case boil.AfterInsertHook:
		schoolAfterInsertHooks = append(schoolAfterInsertHooks, schoolHook)
	case boil.BeforeUpdateHook:
		schoolBeforeUpdateHooks = append(schoolBeforeUpdateHooks, schoolHook)
	case boil.AfterUpdateHook:
		schoolAfterUpdateHooks = append(schoolAfterUpdateHooks, schoolHook)
	case boil.BeforeDeleteHook:
		schoolBeforeDeleteHooks = append(schoolBeforeDeleteHooks, schoolHook)
	case boil.AfterDeleteHook:
		schoolAfterDeleteHooks = append(schoolAfterDeleteHooks, schoolHook)
	case boil.BeforeUpsertHook:
		schoolBeforeUpsertHooks = append(schoolBeforeUpsertHooks, schoolHook)
	case boil.AfterUpsertHook:
		schoolAfterUpsertHooks = append(schoolAfterUpsertHooks, schoolHook)
	}
}

// One returns a single school record from the query.
func (q schoolQuery) One(ctx context.Context, exec boil.ContextExecutor) (*School, error) {
	o := &School{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for schools")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all School records from the query.
func (q schoolQuery) All(ctx context.Context, exec boil.ContextExecutor) (SchoolSlice, error) {
	var o []*School

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to School slice")
	}

	if len(schoolAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all School records in the query.
func (q schoolQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count schools rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q schoolQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if schools exists")
	}

	return count > 0, nil
}

// LowerSchools retrieves all the lower_school's LowerSchools with an executor.
func (o *School) LowerSchools(mods ...qm.QueryMod) lowerSchoolQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`lower_schools`.`school_id`=?", o.ID),
	)

	return LowerSchools(queryMods...)
}

// LoadLowerSchools allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (schoolL) LoadLowerSchools(ctx context.Context, e boil.ContextExecutor, singular bool, maybeSchool interface{}, mods queries.Applicator) error {
	var slice []*School
	var object *School

	if singular {
		var ok bool
		object, ok = maybeSchool.(*School)
		if !ok {
			object = new(School)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSchool)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSchool))
			}
		}
	} else {
		s, ok := maybeSchool.(*[]*School)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSchool)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSchool))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &schoolR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &schoolR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`lower_schools`),
		qm.WhereIn(`lower_schools.school_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load lower_schools")
	}

	var resultSlice []*LowerSchool
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice lower_schools")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on lower_schools")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for lower_schools")
	}

	if len(lowerSchoolAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.LowerSchools = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &lowerSchoolR{}
			}
			foreign.R.School = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.SchoolID {
				local.R.LowerSchools = append(local.R.LowerSchools, foreign)
				if foreign.R == nil {
					foreign.R = &lowerSchoolR{}
				}
				foreign.R.School = local
				break
			}
		}
	}

	return nil
}

// AddLowerSchools adds the given related objects to the existing relationships
// of the school, optionally inserting them as new records.
// Appends related to o.R.LowerSchools.
// Sets related.R.School appropriately.
func (o *School) AddLowerSchools(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*LowerSchool) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.SchoolID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `lower_schools` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"school_id"}),
				strmangle.WhereClause("`", "`", 0, lowerSchoolPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.SchoolID = o.ID
		}
	}

	if o.R == nil {
		o.R = &schoolR{
			LowerSchools: related,
		}
	} else {
		o.R.LowerSchools = append(o.R.LowerSchools, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &lowerSchoolR{
				School: o,
			}
		} else {
			rel.R.School = o
		}
	}
	return nil
}

// Schools retrieves all the records using an executor.
func Schools(mods ...qm.QueryMod) schoolQuery {
	mods = append(mods, qm.From("`schools`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`schools`.*"})
	}

	return schoolQuery{q}
}

// FindSchool retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSchool(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*School, error) {
	schoolObj := &School{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `schools` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, schoolObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from schools")
	}

	if err = schoolObj.doAfterSelectHooks(ctx, exec); err != nil {
		return schoolObj, err
	}

	return schoolObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *School) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no schools provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(schoolColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	schoolInsertCacheMut.RLock()
	cache, cached := schoolInsertCache[key]
	schoolInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			schoolAllColumns,
			schoolColumnsWithDefault,
			schoolColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(schoolType, schoolMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(schoolType, schoolMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `schools` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `schools` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `schools` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, schoolPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db: unable to insert into schools")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == schoolMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "db: unable to populate default values for schools")
	}

CacheNoHooks:
	if !cached {
		schoolInsertCacheMut.Lock()
		schoolInsertCache[key] = cache
		schoolInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the School.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *School) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	schoolUpdateCacheMut.RLock()
	cache, cached := schoolUpdateCache[key]
	schoolUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			schoolAllColumns,
			schoolPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update schools, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `schools` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, schoolPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(schoolType, schoolMapping, append(wl, schoolPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update schools row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for schools")
	}

	if !cached {
		schoolUpdateCacheMut.Lock()
		schoolUpdateCache[key] = cache
		schoolUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q schoolQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for schools")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for schools")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SchoolSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("db: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `schools` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, schoolPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in school slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all school")
	}
	return rowsAff, nil
}

var mySQLSchoolUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *School) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no schools provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(schoolColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLSchoolUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
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
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	schoolUpsertCacheMut.RLock()
	cache, cached := schoolUpsertCache[key]
	schoolUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			schoolAllColumns,
			schoolColumnsWithDefault,
			schoolColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			schoolAllColumns,
			schoolPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("db: unable to upsert schools, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`schools`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `schools` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(schoolType, schoolMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(schoolType, schoolMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "db: unable to upsert for schools")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == schoolMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(schoolType, schoolMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "db: unable to retrieve unique values for schools")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "db: unable to populate default values for schools")
	}

CacheNoHooks:
	if !cached {
		schoolUpsertCacheMut.Lock()
		schoolUpsertCache[key] = cache
		schoolUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single School record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *School) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no School provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), schoolPrimaryKeyMapping)
	sql := "DELETE FROM `schools` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from schools")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for schools")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q schoolQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no schoolQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from schools")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for schools")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SchoolSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(schoolBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `schools` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, schoolPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from school slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for schools")
	}

	if len(schoolAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *School) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindSchool(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SchoolSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SchoolSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), schoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `schools`.* FROM `schools` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, schoolPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in SchoolSlice")
	}

	*o = slice

	return nil
}

// SchoolExists checks if the School row exists.
func SchoolExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `schools` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if schools exists")
	}

	return exists, nil
}

// Exists checks if the School row exists.
func (o *School) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return SchoolExists(ctx, exec, o.ID)
}
