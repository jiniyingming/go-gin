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

// UserInvitationCodeAudit represents a row from 'aypcddg.user_invitation_code_audit'.
type UserInvitationCodeAudit struct {
	ID           int64          `json:"id"`             // id
	UID          int64          `json:"uid"`            // uid
	Status       int8           `json:"status"`         // status
	JSONData     JSON           `json:"json_data"`      // json_data
	AuditAt      sql.NullInt64  `json:"audit_at"`       // audit_at
	AuditDesc    sql.NullString `json:"audit_desc"`     // audit_desc
	MidAdminID   sql.NullInt64  `json:"mid_admin_id"`   // mid_admin_id
	MidAdminName sql.NullString `json:"mid_admin_name"` // mid_admin_name
	Created      sql.NullInt64  `json:"created"`        // created
	Updated      sql.NullInt64  `json:"updated"`        // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserInvitationCodeAudit exists in the database.
func (uica *UserInvitationCodeAudit) Exists() bool { //user_invitation_code_audit
	return uica._exists
}

// Deleted provides information if the UserInvitationCodeAudit has been deleted from the database.
func (uica *UserInvitationCodeAudit) Deleted() bool {
	return uica._deleted
}

// Get table name
func GetUserInvitationCodeAuditTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "user_invitation_code_audit", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserInvitationCodeAudit to the database.
func (uica *UserInvitationCodeAudit) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if uica._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserInvitationCodeAuditTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, status, json_data, audit_at, audit_desc, mid_admin_id, mid_admin_name, created, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uica.UID, uica.Status, uica.JSONData, uica.AuditAt, uica.AuditDesc, uica.MidAdminID, uica.MidAdminName, uica.Created, uica.Updated)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, uica.UID, uica.Status, uica.JSONData, uica.AuditAt, uica.AuditDesc, uica.MidAdminID, uica.MidAdminName, uica.Created, uica.Updated)
	} else {
		res, err = dbConn.Exec(sqlstr, uica.UID, uica.Status, uica.JSONData, uica.AuditAt, uica.AuditDesc, uica.MidAdminID, uica.MidAdminName, uica.Created, uica.Updated)
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
	uica.ID = int64(id)
	uica._exists = true

	return nil
}

// Update updates the UserInvitationCodeAudit in the database.
func (uica *UserInvitationCodeAudit) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if uica._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserInvitationCodeAuditTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, status = ?, json_data = ?, audit_at = ?, audit_desc = ?, mid_admin_id = ?, mid_admin_name = ?, created = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uica.UID, uica.Status, uica.JSONData, uica.AuditAt, uica.AuditDesc, uica.MidAdminID, uica.MidAdminName, uica.Created, uica.Updated, uica.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, uica.UID, uica.Status, uica.JSONData, uica.AuditAt, uica.AuditDesc, uica.MidAdminID, uica.MidAdminName, uica.Created, uica.Updated, uica.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, uica.UID, uica.Status, uica.JSONData, uica.AuditAt, uica.AuditDesc, uica.MidAdminID, uica.MidAdminName, uica.Created, uica.Updated, uica.ID)
	}
	return err
}

// Save saves the UserInvitationCodeAudit to the database.
func (uica *UserInvitationCodeAudit) Save(ctx context.Context) error {
	if uica.Exists() {
		return uica.Update(ctx)
	}

	return uica.Insert(ctx)
}

// Delete deletes the UserInvitationCodeAudit from the database.
func (uica *UserInvitationCodeAudit) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if uica._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserInvitationCodeAuditTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uica.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, uica.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, uica.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	uica._deleted = true

	return nil
}

// UserInvitationCodeAuditsByStatus retrieves a row from 'aypcddg.user_invitation_code_audit' as a UserInvitationCodeAudit.
//
// Generated from index 'status'.
func UserInvitationCodeAuditsByStatus(ctx context.Context, status int8, key ...interface{}) ([]*UserInvitationCodeAudit, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserInvitationCodeAuditTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, status, json_data, audit_at, audit_desc, mid_admin_id, mid_admin_name, created, updated ` +
		`FROM ` + tableName +
		` WHERE status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, status)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*UserInvitationCodeAudit, 0)
	for queryData.Next() {
		uica := UserInvitationCodeAudit{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&uica.ID, &uica.UID, &uica.Status, &uica.JSONData, &uica.AuditAt, &uica.AuditDesc, &uica.MidAdminID, &uica.MidAdminName, &uica.Created, &uica.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &uica)
	}

	return res, nil
}

// UserInvitationCodeAuditsByUID retrieves a row from 'aypcddg.user_invitation_code_audit' as a UserInvitationCodeAudit.
//
// Generated from index 'uid'.
func UserInvitationCodeAuditsByUID(ctx context.Context, uid int64, key ...interface{}) ([]*UserInvitationCodeAudit, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserInvitationCodeAuditTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, status, json_data, audit_at, audit_desc, mid_admin_id, mid_admin_name, created, updated ` +
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
	res := make([]*UserInvitationCodeAudit, 0)
	for queryData.Next() {
		uica := UserInvitationCodeAudit{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&uica.ID, &uica.UID, &uica.Status, &uica.JSONData, &uica.AuditAt, &uica.AuditDesc, &uica.MidAdminID, &uica.MidAdminName, &uica.Created, &uica.Updated)
		if err != nil {
			return nil, err
		}

		res = append(res, &uica)
	}

	return res, nil
}

// UserInvitationCodeAuditByID retrieves a row from 'aypcddg.user_invitation_code_audit' as a UserInvitationCodeAudit.
//
// Generated from index 'user_invitation_code_audit_id_pkey'.
func UserInvitationCodeAuditByID(ctx context.Context, id int64, key ...interface{}) (*UserInvitationCodeAudit, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserInvitationCodeAuditTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, status, json_data, audit_at, audit_desc, mid_admin_id, mid_admin_name, created, updated ` +
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
	uica := UserInvitationCodeAudit{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&uica.ID, &uica.UID, &uica.Status, &uica.JSONData, &uica.AuditAt, &uica.AuditDesc, &uica.MidAdminID, &uica.MidAdminName, &uica.Created, &uica.Updated)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&uica.ID, &uica.UID, &uica.Status, &uica.JSONData, &uica.AuditAt, &uica.AuditDesc, &uica.MidAdminID, &uica.MidAdminName, &uica.Created, &uica.Updated)
		if err != nil {
			return nil, err
		}
	}

	return &uica, nil
}
