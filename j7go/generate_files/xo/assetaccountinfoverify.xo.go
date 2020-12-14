// Package xo contains the types for schema 'ddg_local'.
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

// AssetAccountInfoVerify represents a row from 'ddg_local.asset_account_info_verify'.
type AssetAccountInfoVerify struct {
	ID                 int64          `json:"id"`                    // id
	AssetID            sql.NullInt64  `json:"asset_id"`              // asset_id
	AssetAccountInfoID sql.NullInt64  `json:"asset_account_info_id"` // asset_account_info_id
	IdcardFrontImg     sql.NullString `json:"idcard_front_img"`      // idcard_front_img
	IdcardReverseImg   sql.NullString `json:"idcard_reverse_img"`    // idcard_reverse_img
	VerifyStatus       sql.NullInt64  `json:"verify_status"`         // verify_status
	RefuseContent      sql.NullString `json:"refuse_content"`        // refuse_content
	IdcardShowStatus   sql.NullInt64  `json:"idcard_show_status"`    // idcard_show_status
	Idcard             sql.NullString `json:"idcard"`                // idcard
	CreatedAt          mysql.NullTime `json:"created_at"`            // created_at
	UpdateAt           mysql.NullTime `json:"update_at"`             // update_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AssetAccountInfoVerify exists in the database.
func (aaiv *AssetAccountInfoVerify) Exists() bool { //asset_account_info_verify
	return aaiv._exists
}

// Deleted provides information if the AssetAccountInfoVerify has been deleted from the database.
func (aaiv *AssetAccountInfoVerify) Deleted() bool {
	return aaiv._deleted
}

// Get table name
func GetAssetAccountInfoVerifyTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "asset_account_info_verify", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the AssetAccountInfoVerify to the database.
func (aaiv *AssetAccountInfoVerify) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if aaiv._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAssetAccountInfoVerifyTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`asset_id, asset_account_info_id, idcard_front_img, idcard_reverse_img, verify_status, refuse_content, idcard_show_status, idcard, created_at, update_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aaiv.AssetID, aaiv.AssetAccountInfoID, aaiv.IdcardFrontImg, aaiv.IdcardReverseImg, aaiv.VerifyStatus, aaiv.RefuseContent, aaiv.IdcardShowStatus, aaiv.Idcard, aaiv.CreatedAt, aaiv.UpdateAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, aaiv.AssetID, aaiv.AssetAccountInfoID, aaiv.IdcardFrontImg, aaiv.IdcardReverseImg, aaiv.VerifyStatus, aaiv.RefuseContent, aaiv.IdcardShowStatus, aaiv.Idcard, aaiv.CreatedAt, aaiv.UpdateAt)
	} else {
		res, err = dbConn.Exec(sqlstr, aaiv.AssetID, aaiv.AssetAccountInfoID, aaiv.IdcardFrontImg, aaiv.IdcardReverseImg, aaiv.VerifyStatus, aaiv.RefuseContent, aaiv.IdcardShowStatus, aaiv.Idcard, aaiv.CreatedAt, aaiv.UpdateAt)
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
	aaiv.ID = int64(id)
	aaiv._exists = true

	return nil
}

// Update updates the AssetAccountInfoVerify in the database.
func (aaiv *AssetAccountInfoVerify) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aaiv._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAssetAccountInfoVerifyTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`asset_id = ?, asset_account_info_id = ?, idcard_front_img = ?, idcard_reverse_img = ?, verify_status = ?, refuse_content = ?, idcard_show_status = ?, idcard = ?, created_at = ?, update_at = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aaiv.AssetID, aaiv.AssetAccountInfoID, aaiv.IdcardFrontImg, aaiv.IdcardReverseImg, aaiv.VerifyStatus, aaiv.RefuseContent, aaiv.IdcardShowStatus, aaiv.Idcard, aaiv.CreatedAt, aaiv.UpdateAt, aaiv.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aaiv.AssetID, aaiv.AssetAccountInfoID, aaiv.IdcardFrontImg, aaiv.IdcardReverseImg, aaiv.VerifyStatus, aaiv.RefuseContent, aaiv.IdcardShowStatus, aaiv.Idcard, aaiv.CreatedAt, aaiv.UpdateAt, aaiv.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aaiv.AssetID, aaiv.AssetAccountInfoID, aaiv.IdcardFrontImg, aaiv.IdcardReverseImg, aaiv.VerifyStatus, aaiv.RefuseContent, aaiv.IdcardShowStatus, aaiv.Idcard, aaiv.CreatedAt, aaiv.UpdateAt, aaiv.ID)
	}
	return err
}

// Save saves the AssetAccountInfoVerify to the database.
func (aaiv *AssetAccountInfoVerify) Save(ctx context.Context) error {
	if aaiv.Exists() {
		return aaiv.Update(ctx)
	}

	return aaiv.Insert(ctx)
}

// Delete deletes the AssetAccountInfoVerify from the database.
func (aaiv *AssetAccountInfoVerify) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aaiv._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAssetAccountInfoVerifyTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aaiv.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aaiv.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aaiv.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	aaiv._deleted = true

	return nil
}

// AssetAccountInfoVerifyByID retrieves a row from 'ddg_local.asset_account_info_verify' as a AssetAccountInfoVerify.
//
// Generated from index 'asset_account_info_verify_id_pkey'.
func AssetAccountInfoVerifyByID(ctx context.Context, id int64, key ...interface{}) (*AssetAccountInfoVerify, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAssetAccountInfoVerifyTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, asset_id, asset_account_info_id, idcard_front_img, idcard_reverse_img, verify_status, refuse_content, idcard_show_status, idcard, created_at, update_at ` +
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
	aaiv := AssetAccountInfoVerify{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&aaiv.ID, &aaiv.AssetID, &aaiv.AssetAccountInfoID, &aaiv.IdcardFrontImg, &aaiv.IdcardReverseImg, &aaiv.VerifyStatus, &aaiv.RefuseContent, &aaiv.IdcardShowStatus, &aaiv.Idcard, &aaiv.CreatedAt, &aaiv.UpdateAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&aaiv.ID, &aaiv.AssetID, &aaiv.AssetAccountInfoID, &aaiv.IdcardFrontImg, &aaiv.IdcardReverseImg, &aaiv.VerifyStatus, &aaiv.RefuseContent, &aaiv.IdcardShowStatus, &aaiv.Idcard, &aaiv.CreatedAt, &aaiv.UpdateAt)
		if err != nil {
			return nil, err
		}
	}

	return &aaiv, nil
}
