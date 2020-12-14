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

// FactoryActivityGoodsMapping represents a row from 'aypcddg.factory_activity_goods_mapping'.
type FactoryActivityGoodsMapping struct {
	ID                uint           `json:"id"`                  // id
	FactoryActivityID int            `json:"factory_activity_id"` // factory_activity_id
	GoodsID           int            `json:"goods_id"`            // goods_id
	IsTop             int8           `json:"is_top"`              // is_top
	Status            int8           `json:"status"`              // status
	Sort              int            `json:"sort"`                // sort
	CreatedAt         mysql.NullTime `json:"created_at"`          // created_at
	UpdatedAt         mysql.NullTime `json:"updated_at"`          // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryActivityGoodsMapping exists in the database.
func (fagm *FactoryActivityGoodsMapping) Exists() bool { //factory_activity_goods_mapping
	return fagm._exists
}

// Deleted provides information if the FactoryActivityGoodsMapping has been deleted from the database.
func (fagm *FactoryActivityGoodsMapping) Deleted() bool {
	return fagm._deleted
}

// Get table name
func GetFactoryActivityGoodsMappingTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "factory_activity_goods_mapping", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryActivityGoodsMapping to the database.
func (fagm *FactoryActivityGoodsMapping) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fagm._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryActivityGoodsMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`factory_activity_id, goods_id, is_top, status, sort, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fagm.FactoryActivityID, fagm.GoodsID, fagm.IsTop, fagm.Status, fagm.Sort, fagm.CreatedAt, fagm.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, fagm.FactoryActivityID, fagm.GoodsID, fagm.IsTop, fagm.Status, fagm.Sort, fagm.CreatedAt, fagm.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, fagm.FactoryActivityID, fagm.GoodsID, fagm.IsTop, fagm.Status, fagm.Sort, fagm.CreatedAt, fagm.UpdatedAt)
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
	fagm.ID = uint(id)
	fagm._exists = true

	return nil
}

// Update updates the FactoryActivityGoodsMapping in the database.
func (fagm *FactoryActivityGoodsMapping) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fagm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryActivityGoodsMappingTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`factory_activity_id = ?, goods_id = ?, is_top = ?, status = ?, sort = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fagm.FactoryActivityID, fagm.GoodsID, fagm.IsTop, fagm.Status, fagm.Sort, fagm.CreatedAt, fagm.UpdatedAt, fagm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fagm.FactoryActivityID, fagm.GoodsID, fagm.IsTop, fagm.Status, fagm.Sort, fagm.CreatedAt, fagm.UpdatedAt, fagm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fagm.FactoryActivityID, fagm.GoodsID, fagm.IsTop, fagm.Status, fagm.Sort, fagm.CreatedAt, fagm.UpdatedAt, fagm.ID)
	}
	return err
}

// Save saves the FactoryActivityGoodsMapping to the database.
func (fagm *FactoryActivityGoodsMapping) Save(ctx context.Context) error {
	if fagm.Exists() {
		return fagm.Update(ctx)
	}

	return fagm.Insert(ctx)
}

// Delete deletes the FactoryActivityGoodsMapping from the database.
func (fagm *FactoryActivityGoodsMapping) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fagm._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryActivityGoodsMappingTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fagm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fagm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, fagm.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	fagm._deleted = true

	return nil
}

// FactoryActivityGoodsMappingsByFactoryActivityID retrieves a row from 'aypcddg.factory_activity_goods_mapping' as a FactoryActivityGoodsMapping.
//
// Generated from index 'factory_activity_goods_mapping_factory_activity_id_index'.
func FactoryActivityGoodsMappingsByFactoryActivityID(ctx context.Context, factoryActivityID int, key ...interface{}) ([]*FactoryActivityGoodsMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryActivityGoodsMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, factory_activity_id, goods_id, is_top, status, sort, created_at, updated_at ` +
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
	res := make([]*FactoryActivityGoodsMapping, 0)
	for queryData.Next() {
		fagm := FactoryActivityGoodsMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&fagm.ID, &fagm.FactoryActivityID, &fagm.GoodsID, &fagm.IsTop, &fagm.Status, &fagm.Sort, &fagm.CreatedAt, &fagm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &fagm)
	}

	return res, nil
}

// FactoryActivityGoodsMappingsByGoodsID retrieves a row from 'aypcddg.factory_activity_goods_mapping' as a FactoryActivityGoodsMapping.
//
// Generated from index 'factory_activity_goods_mapping_goods_id_index'.
func FactoryActivityGoodsMappingsByGoodsID(ctx context.Context, goodsID int, key ...interface{}) ([]*FactoryActivityGoodsMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryActivityGoodsMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, factory_activity_id, goods_id, is_top, status, sort, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE goods_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, goodsID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, goodsID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, goodsID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*FactoryActivityGoodsMapping, 0)
	for queryData.Next() {
		fagm := FactoryActivityGoodsMapping{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&fagm.ID, &fagm.FactoryActivityID, &fagm.GoodsID, &fagm.IsTop, &fagm.Status, &fagm.Sort, &fagm.CreatedAt, &fagm.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &fagm)
	}

	return res, nil
}

// FactoryActivityGoodsMappingByID retrieves a row from 'aypcddg.factory_activity_goods_mapping' as a FactoryActivityGoodsMapping.
//
// Generated from index 'factory_activity_goods_mapping_id_pkey'.
func FactoryActivityGoodsMappingByID(ctx context.Context, id uint, key ...interface{}) (*FactoryActivityGoodsMapping, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryActivityGoodsMappingTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, factory_activity_id, goods_id, is_top, status, sort, created_at, updated_at ` +
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
	fagm := FactoryActivityGoodsMapping{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&fagm.ID, &fagm.FactoryActivityID, &fagm.GoodsID, &fagm.IsTop, &fagm.Status, &fagm.Sort, &fagm.CreatedAt, &fagm.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&fagm.ID, &fagm.FactoryActivityID, &fagm.GoodsID, &fagm.IsTop, &fagm.Status, &fagm.Sort, &fagm.CreatedAt, &fagm.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &fagm, nil
}
