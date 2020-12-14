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

// Adv2PlacementMaterialLog represents a row from 'aypcddg.adv2_placement_material_log'.
type Adv2PlacementMaterialLog struct {
	ID              int            `json:"id"`                // id
	PlaceMaterialID sql.NullInt64  `json:"place_material_id"` // place_material_id
	PreValue        sql.NullString `json:"pre_value"`         // pre_value
	PostValue       sql.NullString `json:"post_value"`        // post_value
	Type            sql.NullInt64  `json:"type"`              // type
	Field           sql.NullString `json:"field"`             // field
	Operator        sql.NullInt64  `json:"operator"`          // operator
	OperatorName    sql.NullString `json:"operator_name"`     // operator_name
	Created         sql.NullInt64  `json:"created"`           // created
	Updated         sql.NullInt64  `json:"updated"`           // updated

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Adv2PlacementMaterialLog exists in the database.
func (aml *Adv2PlacementMaterialLog) Exists() bool { //adv2_placement_material_log
	return aml._exists
}

// Deleted provides information if the Adv2PlacementMaterialLog has been deleted from the database.
func (aml *Adv2PlacementMaterialLog) Deleted() bool {
	return aml._deleted
}

// Get table name
func GetAdv2PlacementMaterialLogTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "adv2_placement_material_log", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Adv2PlacementMaterialLog to the database.
func (aml *Adv2PlacementMaterialLog) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if aml._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdv2PlacementMaterialLogTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`place_material_id, pre_value, post_value, type, field, operator, operator_name, created, updated` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aml.PlaceMaterialID, aml.PreValue, aml.PostValue, aml.Type, aml.Field, aml.Operator, aml.OperatorName, aml.Created, aml.Updated)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, aml.PlaceMaterialID, aml.PreValue, aml.PostValue, aml.Type, aml.Field, aml.Operator, aml.OperatorName, aml.Created, aml.Updated)
	} else {
		res, err = dbConn.Exec(sqlstr, aml.PlaceMaterialID, aml.PreValue, aml.PostValue, aml.Type, aml.Field, aml.Operator, aml.OperatorName, aml.Created, aml.Updated)
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
	aml.ID = int(id)
	aml._exists = true

	return nil
}

// Update updates the Adv2PlacementMaterialLog in the database.
func (aml *Adv2PlacementMaterialLog) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aml._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdv2PlacementMaterialLogTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`place_material_id = ?, pre_value = ?, post_value = ?, type = ?, field = ?, operator = ?, operator_name = ?, created = ?, updated = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aml.PlaceMaterialID, aml.PreValue, aml.PostValue, aml.Type, aml.Field, aml.Operator, aml.OperatorName, aml.Created, aml.Updated, aml.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aml.PlaceMaterialID, aml.PreValue, aml.PostValue, aml.Type, aml.Field, aml.Operator, aml.OperatorName, aml.Created, aml.Updated, aml.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aml.PlaceMaterialID, aml.PreValue, aml.PostValue, aml.Type, aml.Field, aml.Operator, aml.OperatorName, aml.Created, aml.Updated, aml.ID)
	}
	return err
}

// Save saves the Adv2PlacementMaterialLog to the database.
func (aml *Adv2PlacementMaterialLog) Save(ctx context.Context) error {
	if aml.Exists() {
		return aml.Update(ctx)
	}

	return aml.Insert(ctx)
}

// Delete deletes the Adv2PlacementMaterialLog from the database.
func (aml *Adv2PlacementMaterialLog) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if aml._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdv2PlacementMaterialLogTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aml.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, aml.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, aml.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	aml._deleted = true

	return nil
}

// Adv2PlacementMaterialLogByID retrieves a row from 'aypcddg.adv2_placement_material_log' as a Adv2PlacementMaterialLog.
//
// Generated from index 'adv2_placement_material_log_id_pkey'.
func Adv2PlacementMaterialLogByID(ctx context.Context, id int, key ...interface{}) (*Adv2PlacementMaterialLog, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAdv2PlacementMaterialLogTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, place_material_id, pre_value, post_value, type, field, operator, operator_name, created, updated ` +
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
	aml := Adv2PlacementMaterialLog{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&aml.ID, &aml.PlaceMaterialID, &aml.PreValue, &aml.PostValue, &aml.Type, &aml.Field, &aml.Operator, &aml.OperatorName, &aml.Created, &aml.Updated)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&aml.ID, &aml.PlaceMaterialID, &aml.PreValue, &aml.PostValue, &aml.Type, &aml.Field, &aml.Operator, &aml.OperatorName, &aml.Created, &aml.Updated)
		if err != nil {
			return nil, err
		}
	}

	return &aml, nil
}