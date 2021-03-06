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

// AlipayUserInfo represents a row from 'aypcddg.alipay_user_info'.
type AlipayUserInfo struct {
	ID           uint64         `json:"id"`            // id
	UID          int64          `json:"uid"`           // uid
	AliUserID    string         `json:"ali_user_id"`   // ali_user_id
	Avatar       sql.NullString `json:"avatar"`        // avatar
	Province     sql.NullString `json:"province"`      // province
	City         sql.NullString `json:"city"`          // city
	NickName     sql.NullString `json:"nick_name"`     // nick_name
	Gender       sql.NullString `json:"gender"`        // gender
	AccessToken  sql.NullString `json:"access_token"`  // access_token
	ExpiresIn    sql.NullInt64  `json:"expires_in"`    // expires_in
	RefreshToken sql.NullString `json:"refresh_token"` // refresh_token
	ReExpiresIn  sql.NullInt64  `json:"re_expires_in"` // re_expires_in
	CreatedAt    mysql.NullTime `json:"created_at"`    // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`    // updated_at
	AuthStatus   sql.NullInt64  `json:"auth_status"`   // auth_status

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AlipayUserInfo exists in the database.
func (aui *AlipayUserInfo) Exists() bool { //alipay_user_info
	return aui._exists
}

// Deleted provides information if the AlipayUserInfo has been deleted from the database.
func (aui *AlipayUserInfo) Deleted() bool {
	return aui._deleted
}

// Get table name
func GetAlipayUserInfoTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "alipay_user_info", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the AlipayUserInfo to the database.
func (aui *AlipayUserInfo) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if aui._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAlipayUserInfoTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, ali_user_id, avatar, province, city, nick_name, gender, access_token, expires_in, refresh_token, re_expires_in, created_at, updated_at, auth_status` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aui.UID, aui.AliUserID, aui.Avatar, aui.Province, aui.City, aui.NickName, aui.Gender, aui.AccessToken, aui.ExpiresIn, aui.RefreshToken, aui.ReExpiresIn, aui.CreatedAt, aui.UpdatedAt, aui.AuthStatus)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, aui.UID, aui.AliUserID, aui.Avatar, aui.Province, aui.City, aui.NickName, aui.Gender, aui.AccessToken, aui.ExpiresIn, aui.RefreshToken, aui.ReExpiresIn, aui.CreatedAt, aui.UpdatedAt, aui.AuthStatus)
	} else {
		res, err = dbConn.Exec(sqlstr, aui.UID, aui.AliUserID, aui.Avatar, aui.Province, aui.City, aui.NickName, aui.Gender, aui.AccessToken, aui.ExpiresIn, aui.RefreshToken, aui.ReExpiresIn, aui.CreatedAt, aui.UpdatedAt, aui.AuthStatus)
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
	aui.ID = uint64(id)
	aui._exists = true

	return nil
}

// Update updates the AlipayUserInfo in the database.
func (aui *AlipayUserInfo) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aui._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAlipayUserInfoTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, ali_user_id = ?, avatar = ?, province = ?, city = ?, nick_name = ?, gender = ?, access_token = ?, expires_in = ?, refresh_token = ?, re_expires_in = ?, created_at = ?, updated_at = ?, auth_status = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aui.UID, aui.AliUserID, aui.Avatar, aui.Province, aui.City, aui.NickName, aui.Gender, aui.AccessToken, aui.ExpiresIn, aui.RefreshToken, aui.ReExpiresIn, aui.CreatedAt, aui.UpdatedAt, aui.AuthStatus, aui.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aui.UID, aui.AliUserID, aui.Avatar, aui.Province, aui.City, aui.NickName, aui.Gender, aui.AccessToken, aui.ExpiresIn, aui.RefreshToken, aui.ReExpiresIn, aui.CreatedAt, aui.UpdatedAt, aui.AuthStatus, aui.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aui.UID, aui.AliUserID, aui.Avatar, aui.Province, aui.City, aui.NickName, aui.Gender, aui.AccessToken, aui.ExpiresIn, aui.RefreshToken, aui.ReExpiresIn, aui.CreatedAt, aui.UpdatedAt, aui.AuthStatus, aui.ID)
	}
	return err
}

// Save saves the AlipayUserInfo to the database.
func (aui *AlipayUserInfo) Save(ctx context.Context) error {
	if aui.Exists() {
		return aui.Update(ctx)
	}

	return aui.Insert(ctx)
}

// Delete deletes the AlipayUserInfo from the database.
func (aui *AlipayUserInfo) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aui._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAlipayUserInfoTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aui.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aui.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aui.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	aui._deleted = true

	return nil
}

// AlipayUserInfoByID retrieves a row from 'aypcddg.alipay_user_info' as a AlipayUserInfo.
//
// Generated from index 'alipay_user_info_id_pkey'.
func AlipayUserInfoByID(ctx context.Context, id uint64, key ...interface{}) (*AlipayUserInfo, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAlipayUserInfoTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, ali_user_id, avatar, province, city, nick_name, gender, access_token, expires_in, refresh_token, re_expires_in, created_at, updated_at, auth_status ` +
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
	aui := AlipayUserInfo{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&aui.ID, &aui.UID, &aui.AliUserID, &aui.Avatar, &aui.Province, &aui.City, &aui.NickName, &aui.Gender, &aui.AccessToken, &aui.ExpiresIn, &aui.RefreshToken, &aui.ReExpiresIn, &aui.CreatedAt, &aui.UpdatedAt, &aui.AuthStatus)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&aui.ID, &aui.UID, &aui.AliUserID, &aui.Avatar, &aui.Province, &aui.City, &aui.NickName, &aui.Gender, &aui.AccessToken, &aui.ExpiresIn, &aui.RefreshToken, &aui.ReExpiresIn, &aui.CreatedAt, &aui.UpdatedAt, &aui.AuthStatus)
		if err != nil {
			return nil, err
		}
	}

	return &aui, nil
}

// AlipayUserInfoByUID retrieves a row from 'aypcddg.alipay_user_info' as a AlipayUserInfo.
//
// Generated from index 'alipay_user_info_uid_unique'.
func AlipayUserInfoByUID(ctx context.Context, uid int64, key ...interface{}) (*AlipayUserInfo, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAlipayUserInfoTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, ali_user_id, avatar, province, city, nick_name, gender, access_token, expires_in, refresh_token, re_expires_in, created_at, updated_at, auth_status ` +
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
	aui := AlipayUserInfo{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, uid).Scan(&aui.ID, &aui.UID, &aui.AliUserID, &aui.Avatar, &aui.Province, &aui.City, &aui.NickName, &aui.Gender, &aui.AccessToken, &aui.ExpiresIn, &aui.RefreshToken, &aui.ReExpiresIn, &aui.CreatedAt, &aui.UpdatedAt, &aui.AuthStatus)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, uid).Scan(&aui.ID, &aui.UID, &aui.AliUserID, &aui.Avatar, &aui.Province, &aui.City, &aui.NickName, &aui.Gender, &aui.AccessToken, &aui.ExpiresIn, &aui.RefreshToken, &aui.ReExpiresIn, &aui.CreatedAt, &aui.UpdatedAt, &aui.AuthStatus)
		if err != nil {
			return nil, err
		}
	}

	return &aui, nil
}
