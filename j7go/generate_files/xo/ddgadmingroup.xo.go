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

// DdgAdminGroup represents a row from 'aypcddg.ddg_admin_group'.
type DdgAdminGroup struct {
	ID            uint64         `json:"id"`             // id
	Name          string         `json:"name"`           // name
	PermissionIds sql.NullString `json:"permission_ids"` // permission_ids
	Status        int8           `json:"status"`         // status

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DdgAdminGroup exists in the database.
func (dag *DdgAdminGroup) Exists() bool { //ddg_admin_group
	return dag._exists
}

// Deleted provides information if the DdgAdminGroup has been deleted from the database.
func (dag *DdgAdminGroup) Deleted() bool {
	return dag._deleted
}

// Get table name
func GetDdgAdminGroupTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "ddg_admin_group", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DdgAdminGroup to the database.
func (dag *DdgAdminGroup) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dag._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDdgAdminGroupTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`name, permission_ids, status` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dag.Name, dag.PermissionIds, dag.Status)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dag.Name, dag.PermissionIds, dag.Status)
	} else {
		res, err = dbConn.Exec(sqlstr, dag.Name, dag.PermissionIds, dag.Status)
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
	dag.ID = uint64(id)
	dag._exists = true

	return nil
}

// Update updates the DdgAdminGroup in the database.
func (dag *DdgAdminGroup) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dag._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDdgAdminGroupTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`name = ?, permission_ids = ?, status = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dag.Name, dag.PermissionIds, dag.Status, dag.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dag.Name, dag.PermissionIds, dag.Status, dag.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dag.Name, dag.PermissionIds, dag.Status, dag.ID)
	}
	return err
}

// Save saves the DdgAdminGroup to the database.
func (dag *DdgAdminGroup) Save(ctx context.Context) error {
	if dag.Exists() {
		return dag.Update(ctx)
	}

	return dag.Insert(ctx)
}

// Delete deletes the DdgAdminGroup from the database.
func (dag *DdgAdminGroup) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dag._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDdgAdminGroupTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dag.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dag.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dag.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dag._deleted = true

	return nil
}

// DdgAdminGroupByID retrieves a row from 'aypcddg.ddg_admin_group' as a DdgAdminGroup.
//
// Generated from index 'ddg_admin_group_id_pkey'.
func DdgAdminGroupByID(ctx context.Context, id uint64, key ...interface{}) (*DdgAdminGroup, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDdgAdminGroupTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, name, permission_ids, status ` +
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
	dag := DdgAdminGroup{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dag.ID, &dag.Name, &dag.PermissionIds, &dag.Status)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dag.ID, &dag.Name, &dag.PermissionIds, &dag.Status)
		if err != nil {
			return nil, err
		}
	}

	return &dag, nil
}
