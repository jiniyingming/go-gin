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

// TmpActivity represents a row from 'aypcddg.tmp_activity'.
type TmpActivity struct {
	ID      int            `json:"id"`       // id
	UID     uint           `json:"uid"`      // uid
	PicURL  sql.NullString `json:"pic_url"`  // pic_url
	ReferID int            `json:"refer_id"` // refer_id
	AddTime uint           `json:"add_time"` // add_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TmpActivity exists in the database.
func (ta *TmpActivity) Exists() bool { //tmp_activity
	return ta._exists
}

// Deleted provides information if the TmpActivity has been deleted from the database.
func (ta *TmpActivity) Deleted() bool {
	return ta._deleted
}

// Get table name
func GetTmpActivityTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "tmp_activity", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the TmpActivity to the database.
func (ta *TmpActivity) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ta._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTmpActivityTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, pic_url, refer_id, add_time` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ta.UID, ta.PicURL, ta.ReferID, ta.AddTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ta.UID, ta.PicURL, ta.ReferID, ta.AddTime)
	} else {
		res, err = dbConn.Exec(sqlstr, ta.UID, ta.PicURL, ta.ReferID, ta.AddTime)
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
	ta.ID = int(id)
	ta._exists = true

	return nil
}

// Update updates the TmpActivity in the database.
func (ta *TmpActivity) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ta._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTmpActivityTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, pic_url = ?, refer_id = ?, add_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ta.UID, ta.PicURL, ta.ReferID, ta.AddTime, ta.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ta.UID, ta.PicURL, ta.ReferID, ta.AddTime, ta.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ta.UID, ta.PicURL, ta.ReferID, ta.AddTime, ta.ID)
	}
	return err
}

// Save saves the TmpActivity to the database.
func (ta *TmpActivity) Save(ctx context.Context) error {
	if ta.Exists() {
		return ta.Update(ctx)
	}

	return ta.Insert(ctx)
}

// Delete deletes the TmpActivity from the database.
func (ta *TmpActivity) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ta._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTmpActivityTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ta.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ta.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ta.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ta._deleted = true

	return nil
}

// TmpActivityByID retrieves a row from 'aypcddg.tmp_activity' as a TmpActivity.
//
// Generated from index 'tmp_activity_id_pkey'.
func TmpActivityByID(ctx context.Context, id int, key ...interface{}) (*TmpActivity, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetTmpActivityTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, pic_url, refer_id, add_time ` +
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
	ta := TmpActivity{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ta.ID, &ta.UID, &ta.PicURL, &ta.ReferID, &ta.AddTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ta.ID, &ta.UID, &ta.PicURL, &ta.ReferID, &ta.AddTime)
		if err != nil {
			return nil, err
		}
	}

	return &ta, nil
}
