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

	"go.uber.org/zap"
)

// DpShopInfo represents a row from 'aypcddg.dp_shop_info'.
type DpShopInfo struct {
	ID          int           `json:"id"`           // id
	ShopName    string        `json:"shop_name"`    // shop_name
	OperatorUID int           `json:"operator_uid"` // operator_uid
	Status      int8          `json:"status"`       // status
	Created     sql.NullInt64 `json:"created"`      // created
	Updated     sql.NullInt64 `json:"updated"`      // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DpShopInfo exists in the database.
func (dsi *DpShopInfo) Exists() bool { //dp_shop_info
	return dsi._exists
}

// Deleted provides information if the DpShopInfo has been deleted from the database.
func (dsi *DpShopInfo) Deleted() bool {
	return dsi._deleted
}

// Get table name
func GetDpShopInfoTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "dp_shop_info", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DpShopInfo to the database.
func (dsi *DpShopInfo) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dsi._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDpShopInfoTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`shop_name, operator_uid, status, created, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dsi.ShopName, dsi.OperatorUID, dsi.Status, dsi.Created, dsi.Updated)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dsi.ShopName, dsi.OperatorUID, dsi.Status, dsi.Created, dsi.Updated)
	} else {
		res, err = dbConn.Exec(sqlstr, dsi.ShopName, dsi.OperatorUID, dsi.Status, dsi.Created, dsi.Updated)
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
	dsi.ID = int(id)
	dsi._exists = true

	return nil
}

// Update updates the DpShopInfo in the database.
func (dsi *DpShopInfo) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dsi._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDpShopInfoTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`shop_name = ?, operator_uid = ?, status = ?, created = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dsi.ShopName, dsi.OperatorUID, dsi.Status, dsi.Created, dsi.Updated, dsi.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dsi.ShopName, dsi.OperatorUID, dsi.Status, dsi.Created, dsi.Updated, dsi.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dsi.ShopName, dsi.OperatorUID, dsi.Status, dsi.Created, dsi.Updated, dsi.ID)
	}
	return err
}

// Save saves the DpShopInfo to the database.
func (dsi *DpShopInfo) Save(ctx context.Context) error {
	if dsi.Exists() {
		return dsi.Update(ctx)
	}

	return dsi.Insert(ctx)
}

// Delete deletes the DpShopInfo from the database.
func (dsi *DpShopInfo) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dsi._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDpShopInfoTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dsi.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dsi.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dsi.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dsi._deleted = true

	return nil
}

// DpShopInfoByID retrieves a row from 'aypcddg.dp_shop_info' as a DpShopInfo.
//
// Generated from index 'dp_shop_info_id_pkey'.
func DpShopInfoByID(ctx context.Context, id int, key ...interface{}) (*DpShopInfo, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDpShopInfoTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, shop_name, operator_uid, status, created, updated ` +
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
	dsi := DpShopInfo{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dsi.ID, &dsi.ShopName, &dsi.OperatorUID, &dsi.Status, &dsi.Created, &dsi.Updated)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dsi.ID, &dsi.ShopName, &dsi.OperatorUID, &dsi.Status, &dsi.Created, &dsi.Updated)
		if err != nil {
			return nil, err
		}
	}

	return &dsi, nil
}

// DpShopInfosByStatus retrieves a row from 'aypcddg.dp_shop_info' as a DpShopInfo.
//
// Generated from index 's_index'.
func DpShopInfosByStatus(ctx context.Context, status int8, key ...interface{}) ([]*DpShopInfo, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDpShopInfoTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, shop_name, operator_uid, status, created, updated ` +
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
	res := make([]*DpShopInfo, 0)
	for queryData.Next() {
		dsi := DpShopInfo{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&dsi.ID, &dsi.ShopName, &dsi.OperatorUID, &dsi.Status, &dsi.Created, &dsi.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &dsi)
	}

	return res, nil
}
