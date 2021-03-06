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

// FactoryUser represents a row from 'aypcddg.factory_user'.
type FactoryUser struct {
	UID             uint           `json:"uid"`               // uid
	Fid             uint           `json:"fid"`               // fid
	Permissions     sql.NullString `json:"permissions"`       // permissions
	Status          int8           `json:"status"`            // status
	Created         uint           `json:"created"`           // created
	MidAdminID      sql.NullInt64  `json:"mid_admin_id"`      // mid_admin_id
	MidAdminName    sql.NullString `json:"mid_admin_name"`    // mid_admin_name
	Department      sql.NullString `json:"department"`        // department
	CreatedAt       mysql.NullTime `json:"created_at"`        // created_at
	UpdatedAt       mysql.NullTime `json:"updated_at"`        // updated_at
	FactoryUserName sql.NullString `json:"factory_user_name"` // factory_user_name

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryUser exists in the database.
func (fu *FactoryUser) Exists() bool { //factory_user
	return fu._exists
}

// Deleted provides information if the FactoryUser has been deleted from the database.
func (fu *FactoryUser) Deleted() bool {
	return fu._deleted
}

// Get table name
func GetFactoryUserTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "factory_user", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryUser to the database.
func (fu *FactoryUser) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fu._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryUserTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, fid, permissions, status, created, mid_admin_id, mid_admin_name, department, created_at, updated_at, factory_user_name` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fu.UID, fu.Fid, fu.Permissions, fu.Status, fu.Created, fu.MidAdminID, fu.MidAdminName, fu.Department, fu.CreatedAt, fu.UpdatedAt, fu.FactoryUserName)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, fu.UID, fu.Fid, fu.Permissions, fu.Status, fu.Created, fu.MidAdminID, fu.MidAdminName, fu.Department, fu.CreatedAt, fu.UpdatedAt, fu.FactoryUserName)
	} else {
		res, err = dbConn.Exec(sqlstr, fu.UID, fu.Fid, fu.Permissions, fu.Status, fu.Created, fu.MidAdminID, fu.MidAdminName, fu.Department, fu.CreatedAt, fu.UpdatedAt, fu.FactoryUserName)
	}

	if err != nil {
		return err
	}

	// set existence
	fu._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	fu.Fid = uint(id)
	fu._exists = true

	return nil
}

// Update updates the FactoryUser in the database.
func (fu *FactoryUser) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fu._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryUserTableName(key...)
	if err != nil {
		return err
	}

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`permissions = ?, status = ?, created = ?, mid_admin_id = ?, mid_admin_name = ?, department = ?, created_at = ?, updated_at = ?, factory_user_name = ?` +
		` WHERE uid = ? AND fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fu.Permissions, fu.Status, fu.Created, fu.MidAdminID, fu.MidAdminName, fu.Department, fu.CreatedAt, fu.UpdatedAt, fu.FactoryUserName, fu.UID, fu.Fid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fu.Permissions, fu.Status, fu.Created, fu.MidAdminID, fu.MidAdminName, fu.Department, fu.CreatedAt, fu.UpdatedAt, fu.FactoryUserName, fu.UID, fu.Fid)
	} else {
		_, err = dbConn.Exec(sqlstr, fu.Permissions, fu.Status, fu.Created, fu.MidAdminID, fu.MidAdminName, fu.Department, fu.CreatedAt, fu.UpdatedAt, fu.FactoryUserName, fu.UID, fu.Fid)
	}
	return err
}

// Save saves the FactoryUser to the database.
func (fu *FactoryUser) Save(ctx context.Context) error {
	if fu.Exists() {
		return fu.Update(ctx)
	}

	return fu.Insert(ctx)
}

// Delete deletes the FactoryUser from the database.
func (fu *FactoryUser) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fu._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryUserTableName(key...)
	if err != nil {
		return err
	}
	//2

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fu.Fid)))

	if tx != nil {
		_, err = tx.Exec(sqlstr, fu.Fid)
	} else {
		_, err = dbConn.Exec(sqlstr, fu.Fid)
	}
	if err != nil {
		return err
	}

	// set deleted
	fu._deleted = true

	return nil
}

// FactoryUserByUIDFid retrieves a row from 'aypcddg.factory_user' as a FactoryUser.
//
// Generated from index 'asd'.
func FactoryUserByUIDFid(ctx context.Context, uid uint, fid uint, key ...interface{}) (*FactoryUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uid, fid, permissions, status, created, mid_admin_id, mid_admin_name, department, created_at, updated_at, factory_user_name ` +
		`FROM ` + tableName +
		` WHERE uid = ? AND fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uid, fid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	fu := FactoryUser{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, uid, fid).Scan(&fu.UID, &fu.Fid, &fu.Permissions, &fu.Status, &fu.Created, &fu.MidAdminID, &fu.MidAdminName, &fu.Department, &fu.CreatedAt, &fu.UpdatedAt, &fu.FactoryUserName)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, uid, fid).Scan(&fu.UID, &fu.Fid, &fu.Permissions, &fu.Status, &fu.Created, &fu.MidAdminID, &fu.MidAdminName, &fu.Department, &fu.CreatedAt, &fu.UpdatedAt, &fu.FactoryUserName)
		if err != nil {
			return nil, err
		}
	}

	return &fu, nil
}

// FactoryUserByFid retrieves a row from 'aypcddg.factory_user' as a FactoryUser.
//
// Generated from index 'factory_user_fid_pkey'.
func FactoryUserByFid(ctx context.Context, fid uint, key ...interface{}) (*FactoryUser, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryUserTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`uid, fid, permissions, status, created, mid_admin_id, mid_admin_name, department, created_at, updated_at, factory_user_name ` +
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
	fu := FactoryUser{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fid).Scan(&fu.UID, &fu.Fid, &fu.Permissions, &fu.Status, &fu.Created, &fu.MidAdminID, &fu.MidAdminName, &fu.Department, &fu.CreatedAt, &fu.UpdatedAt, &fu.FactoryUserName)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fid).Scan(&fu.UID, &fu.Fid, &fu.Permissions, &fu.Status, &fu.Created, &fu.MidAdminID, &fu.MidAdminName, &fu.Department, &fu.CreatedAt, &fu.UpdatedAt, &fu.FactoryUserName)
		if err != nil {
			return nil, err
		}
	}

	return &fu, nil
}
