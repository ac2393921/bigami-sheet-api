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

// LowerSchool is an object representing the database table.
type LowerSchool struct {
	ID       int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	SchoolID int    `boil:"school_id" json:"school_id" toml:"school_id" yaml:"school_id"`
	Name     string `boil:"name" json:"name" toml:"name" yaml:"name"`
	Style    string `boil:"style" json:"style" toml:"style" yaml:"style"`
	Enemy    string `boil:"enemy" json:"enemy" toml:"enemy" yaml:"enemy"`

	R *lowerSchoolR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L lowerSchoolL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var LowerSchoolColumns = struct {
	ID       string
	SchoolID string
	Name     string
	Style    string
	Enemy    string
}{
	ID:       "id",
	SchoolID: "school_id",
	Name:     "name",
	Style:    "style",
	Enemy:    "enemy",
}

var LowerSchoolTableColumns = struct {
	ID       string
	SchoolID string
	Name     string
	Style    string
	Enemy    string
}{
	ID:       "lower_schools.id",
	SchoolID: "lower_schools.school_id",
	Name:     "lower_schools.name",
	Style:    "lower_schools.style",
	Enemy:    "lower_schools.enemy",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var LowerSchoolWhere = struct {
	ID       whereHelperint
	SchoolID whereHelperint
	Name     whereHelperstring
	Style    whereHelperstring
	Enemy    whereHelperstring
}{
	ID:       whereHelperint{field: "`lower_schools`.`id`"},
	SchoolID: whereHelperint{field: "`lower_schools`.`school_id`"},
	Name:     whereHelperstring{field: "`lower_schools`.`name`"},
	Style:    whereHelperstring{field: "`lower_schools`.`style`"},
	Enemy:    whereHelperstring{field: "`lower_schools`.`enemy`"},
}

// LowerSchoolRels is where relationship names are stored.
var LowerSchoolRels = struct {
	School string
}{
	School: "School",
}

// lowerSchoolR is where relationships are stored.
type lowerSchoolR struct {
	School *School `boil:"School" json:"School" toml:"School" yaml:"School"`
}

// NewStruct creates a new relationship struct
func (*lowerSchoolR) NewStruct() *lowerSchoolR {
	return &lowerSchoolR{}
}

func (r *lowerSchoolR) GetSchool() *School {
	if r == nil {
		return nil
	}
	return r.School
}

// lowerSchoolL is where Load methods for each relationship are stored.
type lowerSchoolL struct{}

var (
	lowerSchoolAllColumns            = []string{"id", "school_id", "name", "style", "enemy"}
	lowerSchoolColumnsWithoutDefault = []string{"school_id", "name", "style", "enemy"}
	lowerSchoolColumnsWithDefault    = []string{"id"}
	lowerSchoolPrimaryKeyColumns     = []string{"id"}
	lowerSchoolGeneratedColumns      = []string{}
)

type (
	// LowerSchoolSlice is an alias for a slice of pointers to LowerSchool.
	// This should almost always be used instead of []LowerSchool.
	LowerSchoolSlice []*LowerSchool
	// LowerSchoolHook is the signature for custom LowerSchool hook methods
	LowerSchoolHook func(context.Context, boil.ContextExecutor, *LowerSchool) error

	lowerSchoolQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	lowerSchoolType                 = reflect.TypeOf(&LowerSchool{})
	lowerSchoolMapping              = queries.MakeStructMapping(lowerSchoolType)
	lowerSchoolPrimaryKeyMapping, _ = queries.BindMapping(lowerSchoolType, lowerSchoolMapping, lowerSchoolPrimaryKeyColumns)
	lowerSchoolInsertCacheMut       sync.RWMutex
	lowerSchoolInsertCache          = make(map[string]insertCache)
	lowerSchoolUpdateCacheMut       sync.RWMutex
	lowerSchoolUpdateCache          = make(map[string]updateCache)
	lowerSchoolUpsertCacheMut       sync.RWMutex
	lowerSchoolUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var lowerSchoolAfterSelectHooks []LowerSchoolHook

var lowerSchoolBeforeInsertHooks []LowerSchoolHook
var lowerSchoolAfterInsertHooks []LowerSchoolHook

var lowerSchoolBeforeUpdateHooks []LowerSchoolHook
var lowerSchoolAfterUpdateHooks []LowerSchoolHook

var lowerSchoolBeforeDeleteHooks []LowerSchoolHook
var lowerSchoolAfterDeleteHooks []LowerSchoolHook

var lowerSchoolBeforeUpsertHooks []LowerSchoolHook
var lowerSchoolAfterUpsertHooks []LowerSchoolHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LowerSchool) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *LowerSchool) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *LowerSchool) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *LowerSchool) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *LowerSchool) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *LowerSchool) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *LowerSchool) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *LowerSchool) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *LowerSchool) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range lowerSchoolAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLowerSchoolHook registers your hook function for all future operations.
func AddLowerSchoolHook(hookPoint boil.HookPoint, lowerSchoolHook LowerSchoolHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		lowerSchoolAfterSelectHooks = append(lowerSchoolAfterSelectHooks, lowerSchoolHook)
	case boil.BeforeInsertHook:
		lowerSchoolBeforeInsertHooks = append(lowerSchoolBeforeInsertHooks, lowerSchoolHook)
	case boil.AfterInsertHook:
		lowerSchoolAfterInsertHooks = append(lowerSchoolAfterInsertHooks, lowerSchoolHook)
	case boil.BeforeUpdateHook:
		lowerSchoolBeforeUpdateHooks = append(lowerSchoolBeforeUpdateHooks, lowerSchoolHook)
	case boil.AfterUpdateHook:
		lowerSchoolAfterUpdateHooks = append(lowerSchoolAfterUpdateHooks, lowerSchoolHook)
	case boil.BeforeDeleteHook:
		lowerSchoolBeforeDeleteHooks = append(lowerSchoolBeforeDeleteHooks, lowerSchoolHook)
	case boil.AfterDeleteHook:
		lowerSchoolAfterDeleteHooks = append(lowerSchoolAfterDeleteHooks, lowerSchoolHook)
	case boil.BeforeUpsertHook:
		lowerSchoolBeforeUpsertHooks = append(lowerSchoolBeforeUpsertHooks, lowerSchoolHook)
	case boil.AfterUpsertHook:
		lowerSchoolAfterUpsertHooks = append(lowerSchoolAfterUpsertHooks, lowerSchoolHook)
	}
}

// One returns a single lowerSchool record from the query.
func (q lowerSchoolQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LowerSchool, error) {
	o := &LowerSchool{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: failed to execute a one query for lower_schools")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LowerSchool records from the query.
func (q lowerSchoolQuery) All(ctx context.Context, exec boil.ContextExecutor) (LowerSchoolSlice, error) {
	var o []*LowerSchool

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "db: failed to assign all query results to LowerSchool slice")
	}

	if len(lowerSchoolAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LowerSchool records in the query.
func (q lowerSchoolQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to count lower_schools rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q lowerSchoolQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "db: failed to check if lower_schools exists")
	}

	return count > 0, nil
}

// School pointed to by the foreign key.
func (o *LowerSchool) School(mods ...qm.QueryMod) schoolQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.SchoolID),
	}

	queryMods = append(queryMods, mods...)

	return Schools(queryMods...)
}

// LoadSchool allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (lowerSchoolL) LoadSchool(ctx context.Context, e boil.ContextExecutor, singular bool, maybeLowerSchool interface{}, mods queries.Applicator) error {
	var slice []*LowerSchool
	var object *LowerSchool

	if singular {
		var ok bool
		object, ok = maybeLowerSchool.(*LowerSchool)
		if !ok {
			object = new(LowerSchool)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeLowerSchool)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeLowerSchool))
			}
		}
	} else {
		s, ok := maybeLowerSchool.(*[]*LowerSchool)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeLowerSchool)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeLowerSchool))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &lowerSchoolR{}
		}
		args = append(args, object.SchoolID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &lowerSchoolR{}
			}

			for _, a := range args {
				if a == obj.SchoolID {
					continue Outer
				}
			}

			args = append(args, obj.SchoolID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`schools`),
		qm.WhereIn(`schools.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load School")
	}

	var resultSlice []*School
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice School")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for schools")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for schools")
	}

	if len(schoolAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.School = foreign
		if foreign.R == nil {
			foreign.R = &schoolR{}
		}
		foreign.R.LowerSchools = append(foreign.R.LowerSchools, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SchoolID == foreign.ID {
				local.R.School = foreign
				if foreign.R == nil {
					foreign.R = &schoolR{}
				}
				foreign.R.LowerSchools = append(foreign.R.LowerSchools, local)
				break
			}
		}
	}

	return nil
}

// SetSchool of the lowerSchool to the related item.
// Sets o.R.School to related.
// Adds o to related.R.LowerSchools.
func (o *LowerSchool) SetSchool(ctx context.Context, exec boil.ContextExecutor, insert bool, related *School) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `lower_schools` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"school_id"}),
		strmangle.WhereClause("`", "`", 0, lowerSchoolPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.SchoolID = related.ID
	if o.R == nil {
		o.R = &lowerSchoolR{
			School: related,
		}
	} else {
		o.R.School = related
	}

	if related.R == nil {
		related.R = &schoolR{
			LowerSchools: LowerSchoolSlice{o},
		}
	} else {
		related.R.LowerSchools = append(related.R.LowerSchools, o)
	}

	return nil
}

// LowerSchools retrieves all the records using an executor.
func LowerSchools(mods ...qm.QueryMod) lowerSchoolQuery {
	mods = append(mods, qm.From("`lower_schools`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`lower_schools`.*"})
	}

	return lowerSchoolQuery{q}
}

// FindLowerSchool retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindLowerSchool(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*LowerSchool, error) {
	lowerSchoolObj := &LowerSchool{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `lower_schools` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, lowerSchoolObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "db: unable to select from lower_schools")
	}

	if err = lowerSchoolObj.doAfterSelectHooks(ctx, exec); err != nil {
		return lowerSchoolObj, err
	}

	return lowerSchoolObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *LowerSchool) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("db: no lower_schools provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(lowerSchoolColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	lowerSchoolInsertCacheMut.RLock()
	cache, cached := lowerSchoolInsertCache[key]
	lowerSchoolInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			lowerSchoolAllColumns,
			lowerSchoolColumnsWithDefault,
			lowerSchoolColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(lowerSchoolType, lowerSchoolMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(lowerSchoolType, lowerSchoolMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `lower_schools` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `lower_schools` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `lower_schools` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, lowerSchoolPrimaryKeyColumns))
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
		return errors.Wrap(err, "db: unable to insert into lower_schools")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == lowerSchoolMapping["id"] {
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
		return errors.Wrap(err, "db: unable to populate default values for lower_schools")
	}

CacheNoHooks:
	if !cached {
		lowerSchoolInsertCacheMut.Lock()
		lowerSchoolInsertCache[key] = cache
		lowerSchoolInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the LowerSchool.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *LowerSchool) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	lowerSchoolUpdateCacheMut.RLock()
	cache, cached := lowerSchoolUpdateCache[key]
	lowerSchoolUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			lowerSchoolAllColumns,
			lowerSchoolPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("db: unable to update lower_schools, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `lower_schools` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, lowerSchoolPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(lowerSchoolType, lowerSchoolMapping, append(wl, lowerSchoolPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "db: unable to update lower_schools row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by update for lower_schools")
	}

	if !cached {
		lowerSchoolUpdateCacheMut.Lock()
		lowerSchoolUpdateCache[key] = cache
		lowerSchoolUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q lowerSchoolQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all for lower_schools")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected for lower_schools")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o LowerSchoolSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lowerSchoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `lower_schools` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, lowerSchoolPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to update all in lowerSchool slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to retrieve rows affected all in update all lowerSchool")
	}
	return rowsAff, nil
}

var mySQLLowerSchoolUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *LowerSchool) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("db: no lower_schools provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(lowerSchoolColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLLowerSchoolUniqueColumns, o)

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

	lowerSchoolUpsertCacheMut.RLock()
	cache, cached := lowerSchoolUpsertCache[key]
	lowerSchoolUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			lowerSchoolAllColumns,
			lowerSchoolColumnsWithDefault,
			lowerSchoolColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			lowerSchoolAllColumns,
			lowerSchoolPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("db: unable to upsert lower_schools, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`lower_schools`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `lower_schools` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(lowerSchoolType, lowerSchoolMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(lowerSchoolType, lowerSchoolMapping, ret)
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
		return errors.Wrap(err, "db: unable to upsert for lower_schools")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == lowerSchoolMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(lowerSchoolType, lowerSchoolMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "db: unable to retrieve unique values for lower_schools")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "db: unable to populate default values for lower_schools")
	}

CacheNoHooks:
	if !cached {
		lowerSchoolUpsertCacheMut.Lock()
		lowerSchoolUpsertCache[key] = cache
		lowerSchoolUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single LowerSchool record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *LowerSchool) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("db: no LowerSchool provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), lowerSchoolPrimaryKeyMapping)
	sql := "DELETE FROM `lower_schools` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete from lower_schools")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by delete for lower_schools")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q lowerSchoolQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("db: no lowerSchoolQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from lower_schools")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for lower_schools")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o LowerSchoolSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(lowerSchoolBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lowerSchoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `lower_schools` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, lowerSchoolPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "db: unable to delete all from lowerSchool slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "db: failed to get rows affected by deleteall for lower_schools")
	}

	if len(lowerSchoolAfterDeleteHooks) != 0 {
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
func (o *LowerSchool) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindLowerSchool(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *LowerSchoolSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := LowerSchoolSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), lowerSchoolPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `lower_schools`.* FROM `lower_schools` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, lowerSchoolPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "db: unable to reload all in LowerSchoolSlice")
	}

	*o = slice

	return nil
}

// LowerSchoolExists checks if the LowerSchool row exists.
func LowerSchoolExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `lower_schools` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "db: unable to check if lower_schools exists")
	}

	return exists, nil
}

// Exists checks if the LowerSchool row exists.
func (o *LowerSchool) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return LowerSchoolExists(ctx, exec, o.ID)
}
