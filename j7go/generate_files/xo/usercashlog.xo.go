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

// UserCashLog represents a row from 'aypcddg.user_cash_log'.
type UserCashLog struct {
	ID          int             `json:"id"`           // id
	UID         int             `json:"uid"`          // uid
	CashLeft    float64         `json:"cash_left"`    // cash_left
	CashBefore  float64         `json:"cash_before"`  // cash_before
	Action      sql.NullString  `json:"action"`       // action
	ActionCash  sql.NullFloat64 `json:"action_cash"`  // action_cash
	RelatedType sql.NullString  `json:"related_type"` // related_type
	RelatedID   sql.NullInt64   `json:"related_id"`   // related_id
	CreatedAt   mysql.NullTime  `json:"created_at"`   // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserCashLog exists in the database.
func (ucl *UserCashLog) Exists() bool { //user_cash_log
	return ucl._exists
}

// Deleted provides information if the UserCashLog has been deleted from the database.
func (ucl *UserCashLog) Deleted() bool {
	return ucl._deleted
}

// Get table name
func GetUserCashLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "user_cash_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserCashLog to the database.
func (ucl *UserCashLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ucl._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserCashLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, cash_left, cash_before, action, action_cash, related_type, related_id, created_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ucl.UID, ucl.CashLeft, ucl.CashBefore, ucl.Action, ucl.ActionCash, ucl.RelatedType, ucl.RelatedID, ucl.CreatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ucl.UID, ucl.CashLeft, ucl.CashBefore, ucl.Action, ucl.ActionCash, ucl.RelatedType, ucl.RelatedID, ucl.CreatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, ucl.UID, ucl.CashLeft, ucl.CashBefore, ucl.Action, ucl.ActionCash, ucl.RelatedType, ucl.RelatedID, ucl.CreatedAt)
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
	ucl.ID = int(id)
	ucl._exists = true

	return nil
}

// Update updates the UserCashLog in the database.
func (ucl *UserCashLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ucl._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserCashLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, cash_left = ?, cash_before = ?, action = ?, action_cash = ?, related_type = ?, related_id = ?, created_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ucl.UID, ucl.CashLeft, ucl.CashBefore, ucl.Action, ucl.ActionCash, ucl.RelatedType, ucl.RelatedID, ucl.CreatedAt, ucl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ucl.UID, ucl.CashLeft, ucl.CashBefore, ucl.Action, ucl.ActionCash, ucl.RelatedType, ucl.RelatedID, ucl.CreatedAt, ucl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ucl.UID, ucl.CashLeft, ucl.CashBefore, ucl.Action, ucl.ActionCash, ucl.RelatedType, ucl.RelatedID, ucl.CreatedAt, ucl.ID)
	}
	return err
}

// Save saves the UserCashLog to the database.
func (ucl *UserCashLog) Save(ctx context.Context) error {
	if ucl.Exists() {
		return ucl.Update(ctx)
	}

	return ucl.Insert(ctx)
}

// Delete deletes the UserCashLog from the database.
func (ucl *UserCashLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ucl._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserCashLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ucl.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ucl.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ucl.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ucl._deleted = true

	return nil
}

// UserCashLogByID retrieves a row from 'aypcddg.user_cash_log' as a UserCashLog.
//
// Generated from index 'user_cash_log_id_pkey'.
func UserCashLogByID(ctx context.Context, id int, key ...interface{}) (*UserCashLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserCashLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, cash_left, cash_before, action, action_cash, related_type, related_id, created_at ` +
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
	ucl := UserCashLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ucl.ID, &ucl.UID, &ucl.CashLeft, &ucl.CashBefore, &ucl.Action, &ucl.ActionCash, &ucl.RelatedType, &ucl.RelatedID, &ucl.CreatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ucl.ID, &ucl.UID, &ucl.CashLeft, &ucl.CashBefore, &ucl.Action, &ucl.ActionCash, &ucl.RelatedType, &ucl.RelatedID, &ucl.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &ucl, nil
}
