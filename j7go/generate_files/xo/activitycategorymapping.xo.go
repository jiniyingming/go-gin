// Package xo contains the types for schema 'aypcddg'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// ActivityCategoryMapping represents a row from 'aypcddg.activity_category_mapping'.
type ActivityCategoryMapping struct {
	ID                uint           `json:"id"`                  // id
	ActivityID        int            `json:"activity_id"`         // activity_id
	FactoryActivityID int            `json:"factory_activity_id"` // factory_activity_id
	Class1ID          int            `json:"class1_id"`           // class1_id
	Class2ID          int            `json:"class2_id"`           // class2_id
	Class3ID          int            `json:"class3_id"`           // class3_id
	Status            int8           `json:"status"`              // status
	Sort              int            `json:"sort"`                // sort
	CreatedAt         mysql.NullTime `json:"created_at"`          // created_at
	UpdatedAt         mysql.NullTime `json:"updated_at"`          // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ActivityCategoryMapping exists in the database.
func (acm *ActivityCategoryMapping) Exists() bool { //activity_category_mapping
	return acm._exists
}

// Deleted provides information if the ActivityCategoryMapping has been deleted from the database.
func (acm *ActivityCategoryMapping) Deleted() bool {
	return acm._deleted
}

// Get table name
func GetActivityCategoryMappingTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "activity_category_mapping", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ActivityCategoryMapping to the database.
func (acm *ActivityCategoryMapping) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if acm._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityCategoryMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`activity_id, factory_activity_id, class1_id, class2_id, class3_id, status, sort, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, acm.ActivityID, acm.FactoryActivityID, acm.Class1ID, acm.Class2ID, acm.Class3ID, acm.Status, acm.Sort, acm.CreatedAt, acm.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, acm.ActivityID, acm.FactoryActivityID, acm.Class1ID, acm.Class2ID, acm.Class3ID, acm.Status, acm.Sort, acm.CreatedAt, acm.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, acm.ActivityID, acm.FactoryActivityID, acm.Class1ID, acm.Class2ID, acm.Class3ID, acm.Status, acm.Sort, acm.CreatedAt, acm.UpdatedAt)
	}

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	acm.ID = uint(id)
	acm._exists = true

	return nil
}

// Update updates the ActivityCategoryMapping in the database.
func (acm *ActivityCategoryMapping) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if acm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityCategoryMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`activity_id = ?, factory_activity_id = ?, class1_id = ?, class2_id = ?, class3_id = ?, status = ?, sort = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, acm.ActivityID, acm.FactoryActivityID, acm.Class1ID, acm.Class2ID, acm.Class3ID, acm.Status, acm.Sort, acm.CreatedAt, acm.UpdatedAt, acm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, acm.ActivityID, acm.FactoryActivityID, acm.Class1ID, acm.Class2ID, acm.Class3ID, acm.Status, acm.Sort, acm.CreatedAt, acm.UpdatedAt, acm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, acm.ActivityID, acm.FactoryActivityID, acm.Class1ID, acm.Class2ID, acm.Class3ID, acm.Status, acm.Sort, acm.CreatedAt, acm.UpdatedAt, acm.ID)
	}
	return err
}

// Save saves the ActivityCategoryMapping to the database.
func (acm *ActivityCategoryMapping) Save(ctx context.Context) error {
	if acm.Exists() {
		return acm.Update(ctx)
	}

	return acm.Insert(ctx)
}

// Delete deletes the ActivityCategoryMapping from the database.
func (acm *ActivityCategoryMapping) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if acm._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityCategoryMappingTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, acm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, acm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, acm.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	acm._deleted = true

	return nil
}

// ActivityCategoryMappingsByActivityID retrieves a row from 'aypcddg.activity_category_mapping' as a ActivityCategoryMapping.
//
// Generated from index 'a_id'.
func ActivityCategoryMappingsByActivityID(ctx context.Context, activityID int, key ...interface{}) ([]*ActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, class1_id, class2_id, class3_id, status, sort, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE activity_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, activityID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, activityID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, activityID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*ActivityCategoryMapping, 0)
	for queryData.Next() {
		acm := ActivityCategoryMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&acm.ID, &acm.ActivityID, &acm.FactoryActivityID, &acm.Class1ID, &acm.Class2ID, &acm.Class3ID, &acm.Status, &acm.Sort, &acm.CreatedAt, &acm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &acm)
	}

	return res, nil
}

// ActivityCategoryMappingByID retrieves a row from 'aypcddg.activity_category_mapping' as a ActivityCategoryMapping.
//
// Generated from index 'activity_category_mapping_id_pkey'.
func ActivityCategoryMappingByID(ctx context.Context, id uint, key ...interface{}) (*ActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, class1_id, class2_id, class3_id, status, sort, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, id)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	acm := ActivityCategoryMapping{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&acm.ID, &acm.ActivityID, &acm.FactoryActivityID, &acm.Class1ID, &acm.Class2ID, &acm.Class3ID, &acm.Status, &acm.Sort, &acm.CreatedAt, &acm.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&acm.ID, &acm.ActivityID, &acm.FactoryActivityID, &acm.Class1ID, &acm.Class2ID, &acm.Class3ID, &acm.Status, &acm.Sort, &acm.CreatedAt, &acm.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &acm, nil
}

// ActivityCategoryMappingsByStatus retrieves a row from 'aypcddg.activity_category_mapping' as a ActivityCategoryMapping.
//
// Generated from index 'activity_category_mapping_status_index'.
func ActivityCategoryMappingsByStatus(ctx context.Context, status int8, key ...interface{}) ([]*ActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, class1_id, class2_id, class3_id, status, sort, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, status)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*ActivityCategoryMapping, 0)
	for queryData.Next() {
		acm := ActivityCategoryMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&acm.ID, &acm.ActivityID, &acm.FactoryActivityID, &acm.Class1ID, &acm.Class2ID, &acm.Class3ID, &acm.Status, &acm.Sort, &acm.CreatedAt, &acm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &acm)
	}

	return res, nil
}

// ActivityCategoryMappingsByFactoryActivityID retrieves a row from 'aypcddg.activity_category_mapping' as a ActivityCategoryMapping.
//
// Generated from index 'f_a_id'.
func ActivityCategoryMappingsByFactoryActivityID(ctx context.Context, factoryActivityID int, key ...interface{}) ([]*ActivityCategoryMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityCategoryMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, class1_id, class2_id, class3_id, status, sort, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE factory_activity_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, factoryActivityID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, factoryActivityID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, factoryActivityID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*ActivityCategoryMapping, 0)
	for queryData.Next() {
		acm := ActivityCategoryMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&acm.ID, &acm.ActivityID, &acm.FactoryActivityID, &acm.Class1ID, &acm.Class2ID, &acm.Class3ID, &acm.Status, &acm.Sort, &acm.CreatedAt, &acm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &acm)
	}

	return res, nil
}
