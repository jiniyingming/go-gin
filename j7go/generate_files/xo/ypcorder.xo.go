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

// YpcOrder represents a row from 'aypcddg.ypc_orders'.
type YpcOrder struct {
	ID            int64          `json:"id"`              // id
	ParentOrderNo sql.NullString `json:"parent_order_no"` // parent_order_no
	Gid           sql.NullInt64  `json:"gid"`             // gid
	Number        sql.NullInt64  `json:"number"`          // number
	StockID       sql.NullInt64  `json:"stock_id"`        // stock_id
	CreatedAt     mysql.NullTime `json:"created_at"`      // created_at
	OrderNo       sql.NullString `json:"order_no"`        // order_no
	Fid           sql.NullInt64  `json:"fid"`             // fid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the YpcOrder exists in the database.
func (yo *YpcOrder) Exists() bool { //ypc_orders
	return yo._exists
}

// Deleted provides information if the YpcOrder has been deleted from the database.
func (yo *YpcOrder) Deleted() bool {
	return yo._deleted
}

// Get table name
func GetYpcOrderTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "ypc_orders", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the YpcOrder to the database.
func (yo *YpcOrder) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if yo._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetYpcOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`parent_order_no, gid, number, stock_id, created_at, order_no, fid` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, yo.ParentOrderNo, yo.Gid, yo.Number, yo.StockID, yo.CreatedAt, yo.OrderNo, yo.Fid)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, yo.ParentOrderNo, yo.Gid, yo.Number, yo.StockID, yo.CreatedAt, yo.OrderNo, yo.Fid)
	} else {
		res, err = dbConn.Exec(sqlstr, yo.ParentOrderNo, yo.Gid, yo.Number, yo.StockID, yo.CreatedAt, yo.OrderNo, yo.Fid)
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
	yo.ID = int64(id)
	yo._exists = true

	return nil
}

// Update updates the YpcOrder in the database.
func (yo *YpcOrder) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if yo._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetYpcOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`parent_order_no = ?, gid = ?, number = ?, stock_id = ?, created_at = ?, order_no = ?, fid = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, yo.ParentOrderNo, yo.Gid, yo.Number, yo.StockID, yo.CreatedAt, yo.OrderNo, yo.Fid, yo.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, yo.ParentOrderNo, yo.Gid, yo.Number, yo.StockID, yo.CreatedAt, yo.OrderNo, yo.Fid, yo.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, yo.ParentOrderNo, yo.Gid, yo.Number, yo.StockID, yo.CreatedAt, yo.OrderNo, yo.Fid, yo.ID)
	}
	return err
}

// Save saves the YpcOrder to the database.
func (yo *YpcOrder) Save(ctx context.Context) error {
	if yo.Exists() {
		return yo.Update(ctx)
	}

	return yo.Insert(ctx)
}

// Delete deletes the YpcOrder from the database.
func (yo *YpcOrder) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if yo._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetYpcOrderTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, yo.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, yo.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, yo.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	yo._deleted = true

	return nil
}

// YpcOrderByID retrieves a row from 'aypcddg.ypc_orders' as a YpcOrder.
//
// Generated from index 'ypc_orders_id_pkey'.
func YpcOrderByID(ctx context.Context, id int64, key ...interface{}) (*YpcOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetYpcOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, parent_order_no, gid, number, stock_id, created_at, order_no, fid ` +
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
	yo := YpcOrder{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&yo.ID, &yo.ParentOrderNo, &yo.Gid, &yo.Number, &yo.StockID, &yo.CreatedAt, &yo.OrderNo, &yo.Fid)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&yo.ID, &yo.ParentOrderNo, &yo.Gid, &yo.Number, &yo.StockID, &yo.CreatedAt, &yo.OrderNo, &yo.Fid)
		if err != nil {
			return nil, err
		}
	}

	return &yo, nil
}
