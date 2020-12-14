// Package xo contains the types for schema 'aypcddg'.
package xo

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"fmt"
	"j7go/components"
	"j7go/utils"

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// MidAdminRolePermission represents a row from 'aypcddg.mid_admin_role_permissions'.
type MidAdminRolePermission struct {
	RoleID       int            `json:"role_id"`       // role_id
	PermissionID int            `json:"permission_id"` // permission_id
	CreatedAt    mysql.NullTime `json:"created_at"`    // created_at
	UpdatedAt    mysql.NullTime `json:"updated_at"`    // updated_at
}

// MidAdminRolePermissionsByRoleIDPermissionID retrieves a row from 'aypcddg.mid_admin_role_permissions' as a MidAdminRolePermission.
//
// Generated from index 'mid_admin_role_permissions_role_id_permission_id_index'.
func MidAdminRolePermissionsByRoleIDPermissionID(ctx context.Context, roleID int, permissionID int, key ...interface{}) ([]*MidAdminRolePermission, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetMidAdminRolePermissionTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`role_id, permission_id, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE role_id = ? AND permission_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, roleID, permissionID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, roleID, permissionID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, roleID, permissionID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*MidAdminRolePermission, 0)
	for queryData.Next() {
		marp := MidAdminRolePermission{}

		// scan
		err = queryData.Scan(&marp.RoleID, &marp.PermissionID, &marp.CreatedAt, &marp.UpdatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &marp)
	}

	return res, nil
}
