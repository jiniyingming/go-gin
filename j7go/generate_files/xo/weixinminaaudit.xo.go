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

// WeixinMinaAudit represents a row from 'aypcddg.weixin_mina_audit'.
type WeixinMinaAudit struct {
	ID          uint64    `json:"id"`           // id
	UID         int64     `json:"uid"`          // uid
	Appid       string    `json:"appid"`        // appid
	Auditid     string    `json:"auditid"`      // auditid
	Status      int8      `json:"status"`       // status
	Reason      string    `json:"reason"`       // reason
	CreatedTime time.Time `json:"created_time"` // created_time
	UpdatedTime time.Time `json:"updated_time"` // updated_time

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the WeixinMinaAudit exists in the database.
func (wma *WeixinMinaAudit) Exists() bool { //weixin_mina_audit
	return wma._exists
}

// Deleted provides information if the WeixinMinaAudit has been deleted from the database.
func (wma *WeixinMinaAudit) Deleted() bool {
	return wma._deleted
}

// Get table name
func GetWeixinMinaAuditTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "weixin_mina_audit", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the WeixinMinaAudit to the database.
func (wma *WeixinMinaAudit) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if wma._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWeixinMinaAuditTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, appid, auditid, status, reason, created_time, updated_time` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wma.UID, wma.Appid, wma.Auditid, wma.Status, wma.Reason, wma.CreatedTime, wma.UpdatedTime)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, wma.UID, wma.Appid, wma.Auditid, wma.Status, wma.Reason, wma.CreatedTime, wma.UpdatedTime)
	} else {
		res, err = dbConn.Exec(sqlstr, wma.UID, wma.Appid, wma.Auditid, wma.Status, wma.Reason, wma.CreatedTime, wma.UpdatedTime)
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
	wma.ID = uint64(id)
	wma._exists = true

	return nil
}

// Update updates the WeixinMinaAudit in the database.
func (wma *WeixinMinaAudit) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if wma._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWeixinMinaAuditTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, appid = ?, auditid = ?, status = ?, reason = ?, created_time = ?, updated_time = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wma.UID, wma.Appid, wma.Auditid, wma.Status, wma.Reason, wma.CreatedTime, wma.UpdatedTime, wma.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, wma.UID, wma.Appid, wma.Auditid, wma.Status, wma.Reason, wma.CreatedTime, wma.UpdatedTime, wma.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, wma.UID, wma.Appid, wma.Auditid, wma.Status, wma.Reason, wma.CreatedTime, wma.UpdatedTime, wma.ID)
	}
	return err
}

// Save saves the WeixinMinaAudit to the database.
func (wma *WeixinMinaAudit) Save(ctx context.Context) error {
	if wma.Exists() {
		return wma.Update(ctx)
	}

	return wma.Insert(ctx)
}

// Delete deletes the WeixinMinaAudit from the database.
func (wma *WeixinMinaAudit) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if wma._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetWeixinMinaAuditTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, wma.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, wma.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, wma.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	wma._deleted = true

	return nil
}

// WeixinMinaAuditsByUID retrieves a row from 'aypcddg.weixin_mina_audit' as a WeixinMinaAudit.
//
// Generated from index 'uid'.
func WeixinMinaAuditsByUID(ctx context.Context, uid int64, key ...interface{}) ([]*WeixinMinaAudit, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetWeixinMinaAuditTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, appid, auditid, status, reason, created_time, updated_time ` +
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
	res := make([]*WeixinMinaAudit, 0)
	for queryData.Next() {
		wma := WeixinMinaAudit{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&wma.ID, &wma.UID, &wma.Appid, &wma.Auditid, &wma.Status, &wma.Reason, &wma.CreatedTime, &wma.UpdatedTime)
		if err != nil {
			return nil, err
		}

		res = append(res, &wma)
	}

	return res, nil
}

// WeixinMinaAuditByID retrieves a row from 'aypcddg.weixin_mina_audit' as a WeixinMinaAudit.
//
// Generated from index 'weixin_mina_audit_id_pkey'.
func WeixinMinaAuditByID(ctx context.Context, id uint64, key ...interface{}) (*WeixinMinaAudit, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetWeixinMinaAuditTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, uid, appid, auditid, status, reason, created_time, updated_time ` +
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
	wma := WeixinMinaAudit{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&wma.ID, &wma.UID, &wma.Appid, &wma.Auditid, &wma.Status, &wma.Reason, &wma.CreatedTime, &wma.UpdatedTime)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&wma.ID, &wma.UID, &wma.Appid, &wma.Auditid, &wma.Status, &wma.Reason, &wma.CreatedTime, &wma.UpdatedTime)
		if err != nil {
			return nil, err
		}
	}

	return &wma, nil
}
