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

// ActivityFactoryActivityMapping represents a row from 'aypcddg.activity_factory_activity_mapping'.
type ActivityFactoryActivityMapping struct {
	ID                uint           `json:"id"`                  // id
	ActivityID        int            `json:"activity_id"`         // activity_id
	FactoryActivityID int            `json:"factory_activity_id"` // factory_activity_id
	CreatedAt         mysql.NullTime `json:"created_at"`          // created_at
	UpdatedAt         mysql.NullTime `json:"updated_at"`          // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the ActivityFactoryActivityMapping exists in the database.
func (afam *ActivityFactoryActivityMapping) Exists() bool { //activity_factory_activity_mapping
	return afam._exists
}

// Deleted provides information if the ActivityFactoryActivityMapping has been deleted from the database.
func (afam *ActivityFactoryActivityMapping) Deleted() bool {
	return afam._deleted
}

// Get table name
func GetActivityFactoryActivityMappingTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "activity_factory_activity_mapping", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the ActivityFactoryActivityMapping to the database.
func (afam *ActivityFactoryActivityMapping) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if afam._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityFactoryActivityMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`activity_id, factory_activity_id, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, afam.ActivityID, afam.FactoryActivityID, afam.CreatedAt, afam.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, afam.ActivityID, afam.FactoryActivityID, afam.CreatedAt, afam.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, afam.ActivityID, afam.FactoryActivityID, afam.CreatedAt, afam.UpdatedAt)
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
	afam.ID = uint(id)
	afam._exists = true

	return nil
}

// Update updates the ActivityFactoryActivityMapping in the database.
func (afam *ActivityFactoryActivityMapping) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if afam._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityFactoryActivityMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`activity_id = ?, factory_activity_id = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, afam.ActivityID, afam.FactoryActivityID, afam.CreatedAt, afam.UpdatedAt, afam.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, afam.ActivityID, afam.FactoryActivityID, afam.CreatedAt, afam.UpdatedAt, afam.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, afam.ActivityID, afam.FactoryActivityID, afam.CreatedAt, afam.UpdatedAt, afam.ID)
	}
	return err
}

// Save saves the ActivityFactoryActivityMapping to the database.
func (afam *ActivityFactoryActivityMapping) Save(ctx context.Context) error {
	if afam.Exists() {
		return afam.Update(ctx)
	}

	return afam.Insert(ctx)
}

// Delete deletes the ActivityFactoryActivityMapping from the database.
func (afam *ActivityFactoryActivityMapping) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if afam._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetActivityFactoryActivityMappingTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, afam.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, afam.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, afam.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	afam._deleted = true

	return nil
}

// ActivityFactoryActivityMappingsByActivityID retrieves a row from 'aypcddg.activity_factory_activity_mapping' as a ActivityFactoryActivityMapping.
//
// Generated from index 'activity_factory_activity_mapping_activity_id_index'.
func ActivityFactoryActivityMappingsByActivityID(ctx context.Context, activityID int, key ...interface{}) ([]*ActivityFactoryActivityMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityFactoryActivityMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, created_at, updated_at ` +
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
	res := make([]*ActivityFactoryActivityMapping, 0)
	for queryData.Next() {
		afam := ActivityFactoryActivityMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&afam.ID, &afam.ActivityID, &afam.FactoryActivityID, &afam.CreatedAt, &afam.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &afam)
	}

	return res, nil
}

// ActivityFactoryActivityMappingsByFactoryActivityID retrieves a row from 'aypcddg.activity_factory_activity_mapping' as a ActivityFactoryActivityMapping.
//
// Generated from index 'activity_factory_activity_mapping_factory_activity_id_index'.
func ActivityFactoryActivityMappingsByFactoryActivityID(ctx context.Context, factoryActivityID int, key ...interface{}) ([]*ActivityFactoryActivityMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityFactoryActivityMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, created_at, updated_at ` +
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
	res := make([]*ActivityFactoryActivityMapping, 0)
	for queryData.Next() {
		afam := ActivityFactoryActivityMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&afam.ID, &afam.ActivityID, &afam.FactoryActivityID, &afam.CreatedAt, &afam.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &afam)
	}

	return res, nil
}

// ActivityFactoryActivityMappingByID retrieves a row from 'aypcddg.activity_factory_activity_mapping' as a ActivityFactoryActivityMapping.
//
// Generated from index 'activity_factory_activity_mapping_id_pkey'.
func ActivityFactoryActivityMappingByID(ctx context.Context, id uint, key ...interface{}) (*ActivityFactoryActivityMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetActivityFactoryActivityMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, activity_id, factory_activity_id, created_at, updated_at ` +
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
	afam := ActivityFactoryActivityMapping{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&afam.ID, &afam.ActivityID, &afam.FactoryActivityID, &afam.CreatedAt, &afam.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&afam.ID, &afam.ActivityID, &afam.FactoryActivityID, &afam.CreatedAt, &afam.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &afam, nil
}
