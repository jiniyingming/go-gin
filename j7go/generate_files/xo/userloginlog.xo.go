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

// UserLoginLog represents a row from 'aypcddg.user_login_log'.
type UserLoginLog struct {
	Ulid    uint64         `json:"ulid"`    // ulid
	UID     uint           `json:"uid"`     // uid
	Created uint           `json:"created"` // created
	IP      sql.NullString `json:"ip"`      // ip
	Type    string         `json:"type"`    // type

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserLoginLog exists in the database.
func (ull *UserLoginLog) Exists() bool { //user_login_log
	return ull._exists
}

// Deleted provides information if the UserLoginLog has been deleted from the database.
func (ull *UserLoginLog) Deleted() bool {
	return ull._deleted
}

// Get table name
func GetUserLoginLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "user_login_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserLoginLog to the database.
func (ull *UserLoginLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ull._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserLoginLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, created, ip, type` +
		`) VALUES (` +
		`?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ull.UID, ull.Created, ull.IP, ull.Type)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ull.UID, ull.Created, ull.IP, ull.Type)
	} else {
		res, err = dbConn.Exec(sqlstr, ull.UID, ull.Created, ull.IP, ull.Type)
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
	ull.Ulid = uint64(id)
	ull._exists = true

	return nil
}

// Update updates the UserLoginLog in the database.
func (ull *UserLoginLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ull._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserLoginLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, created = ?, ip = ?, type = ?` +
		` WHERE ulid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ull.UID, ull.Created, ull.IP, ull.Type, ull.Ulid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ull.UID, ull.Created, ull.IP, ull.Type, ull.Ulid)
	} else {
		_, err = dbConn.Exec(sqlstr, ull.UID, ull.Created, ull.IP, ull.Type, ull.Ulid)
	}
	return err
}

// Save saves the UserLoginLog to the database.
func (ull *UserLoginLog) Save(ctx context.Context) error {
	if ull.Exists() {
		return ull.Update(ctx)
	}

	return ull.Insert(ctx)
}

// Delete deletes the UserLoginLog from the database.
func (ull *UserLoginLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ull._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserLoginLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE ulid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ull.Ulid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ull.Ulid)
	} else {
		_, err = dbConn.Exec(sqlstr, ull.Ulid)
	}

	if err != nil {
		return err
	}

	// set deleted
	ull._deleted = true

	return nil
}

// UserLoginLogsByUID retrieves a row from 'aypcddg.user_login_log' as a UserLoginLog.
//
// Generated from index 'uid'.
func UserLoginLogsByUID(ctx context.Context, uid uint, key ...interface{}) ([]*UserLoginLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserLoginLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`ulid, uid, created, ip, type ` +
		`FROM ` + tableName +
		` WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserLoginLog, 0)
	for queryData.Next() {
		ull := UserLoginLog{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ull.Ulid, &ull.UID, &ull.Created, &ull.IP, &ull.Type)
		if err != nil {
			return nil, err
		}

		res = append(res, &ull)
	}

	return res, nil
}

// UserLoginLogByUlid retrieves a row from 'aypcddg.user_login_log' as a UserLoginLog.
//
// Generated from index 'user_login_log_ulid_pkey'.
func UserLoginLogByUlid(ctx context.Context, ulid uint64, key ...interface{}) (*UserLoginLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserLoginLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`ulid, uid, created, ip, type ` +
		`FROM ` + tableName +
		` WHERE ulid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ulid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ull := UserLoginLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, ulid).Scan(&ull.Ulid, &ull.UID, &ull.Created, &ull.IP, &ull.Type)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, ulid).Scan(&ull.Ulid, &ull.UID, &ull.Created, &ull.IP, &ull.Type)
		if err != nil {
			return nil, err
		}
	}

	return &ull, nil
}
