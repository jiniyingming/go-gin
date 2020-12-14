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

// Receipt represents a row from 'aypcddg.receipt'.
type Receipt struct {
	Rid     uint           `json:"rid"`     // rid
	Oid     uint64         `json:"oid"`     // oid
	UID     uint           `json:"uid"`     // uid
	Type    int8           `json:"type"`    // type
	Title   sql.NullString `json:"title"`   // title
	Content sql.NullString `json:"content"` // content

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Receipt exists in the database.
func (r *Receipt) Exists() bool { //receipt
	return r._exists
}

// Deleted provides information if the Receipt has been deleted from the database.
func (r *Receipt) Deleted() bool {
	return r._deleted
}

// Get table name
func GetReceiptTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "receipt", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Receipt to the database.
func (r *Receipt) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if r._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetReceiptTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`oid, uid, type, title, content` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, r.Oid, r.UID, r.Type, r.Title, r.Content)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, r.Oid, r.UID, r.Type, r.Title, r.Content)
	} else {
		res, err = dbConn.Exec(sqlstr, r.Oid, r.UID, r.Type, r.Title, r.Content)
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
	r.Rid = uint(id)
	r._exists = true

	return nil
}

// Update updates the Receipt in the database.
func (r *Receipt) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if r._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetReceiptTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`oid = ?, uid = ?, type = ?, title = ?, content = ?` +
		` WHERE rid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, r.Oid, r.UID, r.Type, r.Title, r.Content, r.Rid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, r.Oid, r.UID, r.Type, r.Title, r.Content, r.Rid)
	} else {
		_, err = dbConn.Exec(sqlstr, r.Oid, r.UID, r.Type, r.Title, r.Content, r.Rid)
	}
	return err
}

// Save saves the Receipt to the database.
func (r *Receipt) Save(ctx context.Context) error {
	if r.Exists() {
		return r.Update(ctx)
	}

	return r.Insert(ctx)
}

// Delete deletes the Receipt from the database.
func (r *Receipt) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if r._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetReceiptTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE rid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, r.Rid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, r.Rid)
	} else {
		_, err = dbConn.Exec(sqlstr, r.Rid)
	}

	if err != nil {
		return err
	}

	// set deleted
	r._deleted = true

	return nil
}

// ReceiptByRid retrieves a row from 'aypcddg.receipt' as a Receipt.
//
// Generated from index 'receipt_rid_pkey'.
func ReceiptByRid(ctx context.Context, rid uint, key ...interface{}) (*Receipt, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetReceiptTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`rid, oid, uid, type, title, content ` +
		`FROM ` + tableName +
		` WHERE rid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, rid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	r := Receipt{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, rid).Scan(&r.Rid, &r.Oid, &r.UID, &r.Type, &r.Title, &r.Content)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, rid).Scan(&r.Rid, &r.Oid, &r.UID, &r.Type, &r.Title, &r.Content)
		if err != nil {
			return nil, err
		}
	}

	return &r, nil
}