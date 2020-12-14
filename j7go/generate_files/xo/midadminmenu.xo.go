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

// MidAdminMenu represents a row from 'aypcddg.mid_admin_menu'.
type MidAdminMenu struct {
	ID         uint           `json:"id"`         // id
	ParentID   int            `json:"parent_id"`  // parent_id
	Order      int            `json:"order"`      // order
	Title      string         `json:"title"`      // title
	Icon       string         `json:"icon"`       // icon
	URI        sql.NullString `json:"uri"`        // uri
	Permission sql.NullString `json:"permission"` // permission
	CreatedAt  mysql.NullTime `json:"created_at"` // created_at
	UpdatedAt  mysql.NullTime `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the MidAdminMenu exists in the database.
func (mam *MidAdminMenu) Exists() bool { //mid_admin_menu
	return mam._exists
}

// Deleted provides information if the MidAdminMenu has been deleted from the database.
func (mam *MidAdminMenu) Deleted() bool {
	return mam._deleted
}

// Get table name
func GetMidAdminMenuTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "mid_admin_menu", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the MidAdminMenu to the database.
func (mam *MidAdminMenu) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if mam._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminMenuTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`parent_id, order, title, icon, uri, permission, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mam.ParentID, mam.Order, mam.Title, mam.Icon, mam.URI, mam.Permission, mam.CreatedAt, mam.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, mam.ParentID, mam.Order, mam.Title, mam.Icon, mam.URI, mam.Permission, mam.CreatedAt, mam.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, mam.ParentID, mam.Order, mam.Title, mam.Icon, mam.URI, mam.Permission, mam.CreatedAt, mam.UpdatedAt)
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
	mam.ID = uint(id)
	mam._exists = true

	return nil
}

// Update updates the MidAdminMenu in the database.
func (mam *MidAdminMenu) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mam._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminMenuTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`parent_id = ?, order = ?, title = ?, icon = ?, uri = ?, permission = ?, created_at = ?, updated_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mam.ParentID, mam.Order, mam.Title, mam.Icon, mam.URI, mam.Permission, mam.CreatedAt, mam.UpdatedAt, mam.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mam.ParentID, mam.Order, mam.Title, mam.Icon, mam.URI, mam.Permission, mam.CreatedAt, mam.UpdatedAt, mam.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mam.ParentID, mam.Order, mam.Title, mam.Icon, mam.URI, mam.Permission, mam.CreatedAt, mam.UpdatedAt, mam.ID)
	}
	return err
}

// Save saves the MidAdminMenu to the database.
func (mam *MidAdminMenu) Save(ctx context.Context) error {
	if mam.Exists() {
		return mam.Update(ctx)
	}

	return mam.Insert(ctx)
}

// Delete deletes the MidAdminMenu from the database.
func (mam *MidAdminMenu) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if mam._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetMidAdminMenuTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, mam.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, mam.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, mam.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	mam._deleted = true

	return nil
}

// MidAdminMenuByID retrieves a row from 'aypcddg.mid_admin_menu' as a MidAdminMenu.
//
// Generated from index 'mid_admin_menu_id_pkey'.
func MidAdminMenuByID(ctx context.Context, id uint, key ...interface{}) (*MidAdminMenu, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminMenuTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, parent_id, order, title, icon, uri, permission, created_at, updated_at ` +
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
	mam := MidAdminMenu{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&mam.ID, &mam.ParentID, &mam.Order, &mam.Title, &mam.Icon, &mam.URI, &mam.Permission, &mam.CreatedAt, &mam.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&mam.ID, &mam.ParentID, &mam.Order, &mam.Title, &mam.Icon, &mam.URI, &mam.Permission, &mam.CreatedAt, &mam.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &mam, nil
}