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

// FactoryFare represents a row from 'aypcddg.factory_fare'.
type FactoryFare struct {
	FareID         int            `json:"fare_id"`           // fare_id
	FareTid        sql.NullInt64  `json:"fare_tid"`          // fare_tid
	FareCitys      sql.NullString `json:"fare_citys"`        // fare_citys
	FareType       int            `json:"fare_type"`         // fare_type
	FareFirstNum   int            `json:"fare_first_num"`    // fare_first_num
	FareFirst      float64        `json:"fare_first"`        // fare_first
	FareAddNum     int            `json:"fare_add_num"`      // fare_add_num
	FareAdd        float64        `json:"fare_add"`          // fare_add
	FareDelivery   sql.NullString `json:"fare_delivery"`     // fare_delivery
	FareDeliveryTp int            `json:"fare_delivery_tp"`  // fare_delivery_tp
	FareFreeGapNum float64        `json:"fare_free_gap_num"` // fare_free_gap_num
	FareFreeNum    float64        `json:"fare_free_num"`     // fare_free_num
	FareExpressID  int            `json:"fare_express_id"`   // fare_express_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryFare exists in the database.
func (ff *FactoryFare) Exists() bool { //factory_fare
	return ff._exists
}

// Deleted provides information if the FactoryFare has been deleted from the database.
func (ff *FactoryFare) Deleted() bool {
	return ff._deleted
}

// Get table name
func GetFactoryFareTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "factory_fare", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryFare to the database.
func (ff *FactoryFare) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ff._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryFareTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fare_tid, fare_citys, fare_type, fare_first_num, fare_first, fare_add_num, fare_add, fare_delivery, fare_delivery_tp, fare_free_gap_num, fare_free_num, fare_express_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ff.FareTid, ff.FareCitys, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareFreeGapNum, ff.FareFreeNum, ff.FareExpressID)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ff.FareTid, ff.FareCitys, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareFreeGapNum, ff.FareFreeNum, ff.FareExpressID)
	} else {
		res, err = dbConn.Exec(sqlstr, ff.FareTid, ff.FareCitys, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareFreeGapNum, ff.FareFreeNum, ff.FareExpressID)
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
	ff.FareID = int(id)
	ff._exists = true

	return nil
}

// Update updates the FactoryFare in the database.
func (ff *FactoryFare) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ff._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryFareTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fare_tid = ?, fare_citys = ?, fare_type = ?, fare_first_num = ?, fare_first = ?, fare_add_num = ?, fare_add = ?, fare_delivery = ?, fare_delivery_tp = ?, fare_free_gap_num = ?, fare_free_num = ?, fare_express_id = ?` +
		` WHERE fare_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ff.FareTid, ff.FareCitys, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareFreeGapNum, ff.FareFreeNum, ff.FareExpressID, ff.FareID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ff.FareTid, ff.FareCitys, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareFreeGapNum, ff.FareFreeNum, ff.FareExpressID, ff.FareID)
	} else {
		_, err = dbConn.Exec(sqlstr, ff.FareTid, ff.FareCitys, ff.FareType, ff.FareFirstNum, ff.FareFirst, ff.FareAddNum, ff.FareAdd, ff.FareDelivery, ff.FareDeliveryTp, ff.FareFreeGapNum, ff.FareFreeNum, ff.FareExpressID, ff.FareID)
	}
	return err
}

// Save saves the FactoryFare to the database.
func (ff *FactoryFare) Save(ctx context.Context) error {
	if ff.Exists() {
		return ff.Update(ctx)
	}

	return ff.Insert(ctx)
}

// Delete deletes the FactoryFare from the database.
func (ff *FactoryFare) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ff._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryFareTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fare_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ff.FareID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ff.FareID)
	} else {
		_, err = dbConn.Exec(sqlstr, ff.FareID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ff._deleted = true

	return nil
}

// FactoryFareByFareID retrieves a row from 'aypcddg.factory_fare' as a FactoryFare.
//
// Generated from index 'factory_fare_fare_id_pkey'.
func FactoryFareByFareID(ctx context.Context, fareID int, key ...interface{}) (*FactoryFare, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryFareTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fare_id, fare_tid, fare_citys, fare_type, fare_first_num, fare_first, fare_add_num, fare_add, fare_delivery, fare_delivery_tp, fare_free_gap_num, fare_free_num, fare_express_id ` +
		`FROM ` + tableName +
		` WHERE fare_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fareID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	ff := FactoryFare{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fareID).Scan(&ff.FareID, &ff.FareTid, &ff.FareCitys, &ff.FareType, &ff.FareFirstNum, &ff.FareFirst, &ff.FareAddNum, &ff.FareAdd, &ff.FareDelivery, &ff.FareDeliveryTp, &ff.FareFreeGapNum, &ff.FareFreeNum, &ff.FareExpressID)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fareID).Scan(&ff.FareID, &ff.FareTid, &ff.FareCitys, &ff.FareType, &ff.FareFirstNum, &ff.FareFirst, &ff.FareAddNum, &ff.FareAdd, &ff.FareDelivery, &ff.FareDeliveryTp, &ff.FareFreeGapNum, &ff.FareFreeNum, &ff.FareExpressID)
		if err != nil {
			return nil, err
		}
	}

	return &ff, nil
}
