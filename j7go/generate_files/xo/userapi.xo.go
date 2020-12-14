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
	"time"

	"go.uber.org/zap"
)

// UserAPI represents a row from 'aypcddg.user_api'.
type UserAPI struct {
	Skey     string    `json:"skey"`      // skey
	Secret   string    `json:"secret"`    // secret
	UID      int       `json:"uid"`       // uid
	Fid      int       `json:"fid"`       // fid
	UpdateAt time.Time `json:"update_at"` // update_at
	Token    string    `json:"token"`     // token

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserAPI exists in the database.
func (ua *UserAPI) Exists() bool { //user_api
	return ua._exists
}

// Deleted provides information if the UserAPI has been deleted from the database.
func (ua *UserAPI) Deleted() bool {
	return ua._deleted
}

// Get table name
func GetUserAPITableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "user_api", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserAPI to the database.
func (ua *UserAPI) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ua._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserAPITableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`skey, secret, uid, fid, update_at, token` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ua.Skey, ua.Secret, ua.UID, ua.Fid, ua.UpdateAt, ua.Token)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, ua.Skey, ua.Secret, ua.UID, ua.Fid, ua.UpdateAt, ua.Token)
	} else {
		res, err = dbConn.Exec(sqlstr, ua.Skey, ua.Secret, ua.UID, ua.Fid, ua.UpdateAt, ua.Token)
	}

	if err != nil {
		return err
	}

	// set existence
	ua._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	ua.Skey = string(id)
	ua._exists = true

	return nil
}

// Update updates the UserAPI in the database.
func (ua *UserAPI) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ua._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserAPITableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`secret = ?, uid = ?, fid = ?, update_at = ?, token = ?` +
		` WHERE skey = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ua.Secret, ua.UID, ua.Fid, ua.UpdateAt, ua.Token, ua.Skey)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ua.Secret, ua.UID, ua.Fid, ua.UpdateAt, ua.Token, ua.Skey)
	} else {
		_, err = dbConn.Exec(sqlstr, ua.Secret, ua.UID, ua.Fid, ua.UpdateAt, ua.Token, ua.Skey)
	}
	return err
}

// Save saves the UserAPI to the database.
func (ua *UserAPI) Save(ctx context.Context) error {
	if ua.Exists() {
		return ua.Update(ctx)
	}

	return ua.Insert(ctx)
}

// Delete deletes the UserAPI from the database.
func (ua *UserAPI) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ua._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserAPITableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE skey = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ua.Skey)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ua.Skey)
	} else {
		_, err = dbConn.Exec(sqlstr, ua.Skey)
	}

	if err != nil {
		return err
	}

	// set deleted
	ua._deleted = true

	return nil
}

// UserAPIsByFid retrieves a row from 'aypcddg.user_api' as a UserAPI.
//
// Generated from index 'fid'.
func UserAPIsByFid(ctx context.Context, fid int, key ...interface{}) ([]*UserAPI, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAPITableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`skey, secret, uid, fid, update_at, token ` +
		`FROM ` + tableName +
		` WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, fid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, fid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserAPI, 0)
	for queryData.Next() {
		ua := UserAPI{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ua.Skey, &ua.Secret, &ua.UID, &ua.Fid, &ua.UpdateAt, &ua.Token)
		if err != nil {
			return nil, err
		}

		res = append(res, &ua)
	}

	return res, nil
}

// UserAPIBySkey retrieves a row from 'aypcddg.user_api' as a UserAPI.
//
// Generated from index 'skey'.
func UserAPIBySkey(ctx context.Context, skey string, key ...interface{}) (*UserAPI, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAPITableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`skey, secret, uid, fid, update_at, token ` +
		`FROM ` + tableName +
		` WHERE skey = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, skey)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ua := UserAPI{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, skey).Scan(&ua.Skey, &ua.Secret, &ua.UID, &ua.Fid, &ua.UpdateAt, &ua.Token)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, skey).Scan(&ua.Skey, &ua.Secret, &ua.UID, &ua.Fid, &ua.UpdateAt, &ua.Token)
		if err != nil {
			return nil, err
		}
	}

	return &ua, nil
}

// UserAPIsByUID retrieves a row from 'aypcddg.user_api' as a UserAPI.
//
// Generated from index 'uid'.
func UserAPIsByUID(ctx context.Context, uid int, key ...interface{}) ([]*UserAPI, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAPITableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`skey, secret, uid, fid, update_at, token ` +
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
	res := make([]*UserAPI, 0)
	for queryData.Next() {
		ua := UserAPI{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ua.Skey, &ua.Secret, &ua.UID, &ua.Fid, &ua.UpdateAt, &ua.Token)
		if err != nil {
			return nil, err
		}

		res = append(res, &ua)
	}

	return res, nil
}

// UserAPIBySkey retrieves a row from 'aypcddg.user_api' as a UserAPI.
//
// Generated from index 'user_api_skey_pkey'.
func UserAPIBySkey(ctx context.Context, skey string, key ...interface{}) (*UserAPI, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserAPITableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`skey, secret, uid, fid, update_at, token ` +
		`FROM ` + tableName +
		` WHERE skey = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, skey)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ua := UserAPI{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, skey).Scan(&ua.Skey, &ua.Secret, &ua.UID, &ua.Fid, &ua.UpdateAt, &ua.Token)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, skey).Scan(&ua.Skey, &ua.Secret, &ua.UID, &ua.Fid, &ua.UpdateAt, &ua.Token)
		if err != nil {
			return nil, err
		}
	}

	return &ua, nil
}
