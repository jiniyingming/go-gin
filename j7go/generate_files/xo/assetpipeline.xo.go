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

	"go.uber.org/zap"
)

// AssetPipeline represents a row from 'ddg_local.asset_pipeline'.
type AssetPipeline struct {
	ID             int64          `json:"id"`              // id
	AssetID        sql.NullInt64  `json:"asset_id"`        // asset_id
	Title          sql.NullString `json:"title"`           // title
	Content        sql.NullString `json:"content"`         // content
	Remark         sql.NullString `json:"remark"`          // remark
	OperationType  sql.NullInt64  `json:"operation_type"`  // operation_type
	OperationValue sql.NullString `json:"operation_value"` // operation_value

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the AssetPipeline exists in the database.
func (ap *AssetPipeline) Exists() bool { //asset_pipeline
	return ap._exists
}

// Deleted provides information if the AssetPipeline has been deleted from the database.
func (ap *AssetPipeline) Deleted() bool {
	return ap._deleted
}

// Get table name
func GetAssetPipelineTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "asset_pipeline", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the AssetPipeline to the database.
func (ap *AssetPipeline) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ap._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAssetPipelineTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`asset_id, title, content, remark, operation_type, operation_value` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ap.AssetID, ap.Title, ap.Content, ap.Remark, ap.OperationType, ap.OperationValue)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ap.AssetID, ap.Title, ap.Content, ap.Remark, ap.OperationType, ap.OperationValue)
	} else {
		res, err = dbConn.Exec(sqlstr, ap.AssetID, ap.Title, ap.Content, ap.Remark, ap.OperationType, ap.OperationValue)
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
	ap.ID = int64(id)
	ap._exists = true

	return nil
}

// Update updates the AssetPipeline in the database.
func (ap *AssetPipeline) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ap._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAssetPipelineTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`asset_id = ?, title = ?, content = ?, remark = ?, operation_type = ?, operation_value = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ap.AssetID, ap.Title, ap.Content, ap.Remark, ap.OperationType, ap.OperationValue, ap.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ap.AssetID, ap.Title, ap.Content, ap.Remark, ap.OperationType, ap.OperationValue, ap.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ap.AssetID, ap.Title, ap.Content, ap.Remark, ap.OperationType, ap.OperationValue, ap.ID)
	}
	return err
}

// Save saves the AssetPipeline to the database.
func (ap *AssetPipeline) Save(ctx context.Context) error {
	if ap.Exists() {
		return ap.Update(ctx)
	}

	return ap.Insert(ctx)
}

// Delete deletes the AssetPipeline from the database.
func (ap *AssetPipeline) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ap._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAssetPipelineTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ap.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ap.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ap.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ap._deleted = true

	return nil
}

// AssetPipelineByID retrieves a row from 'ddg_local.asset_pipeline' as a AssetPipeline.
//
// Generated from index 'asset_pipeline_id_pkey'.
func AssetPipelineByID(ctx context.Context, id int64, key ...interface{}) (*AssetPipeline, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAssetPipelineTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, asset_id, title, content, remark, operation_type, operation_value ` +
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
	ap := AssetPipeline{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ap.ID, &ap.AssetID, &ap.Title, &ap.Content, &ap.Remark, &ap.OperationType, &ap.OperationValue)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ap.ID, &ap.AssetID, &ap.Title, &ap.Content, &ap.Remark, &ap.OperationType, &ap.OperationValue)
		if err != nil {
			return nil, err
		}
	}

	return &ap, nil
}
