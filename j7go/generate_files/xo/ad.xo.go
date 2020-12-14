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

// Ad represents a row from 'aypcddg.ad'.
type Ad struct {
	Aid         int            `json:"aid"`         // aid
	Aiid        int            `json:"aiid"`        // aiid
	Apid        int            `json:"apid"`        // apid
	Starttime   int            `json:"starttime"`   // starttime
	Endtime     int            `json:"endtime"`     // endtime
	Dailybudget float64        `json:"dailybudget"` // dailybudget
	Chargingway sql.NullString `json:"chargingway"` // chargingway
	Price       float64        `json:"price"`       // price
	Addby       int            `json:"addby"`       // addby
	Updated     int            `json:"updated"`     // updated
	Created     int            `json:"created"`     // created
	Status      bool           `json:"status"`      // status

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Ad exists in the database.
func (a *Ad) Exists() bool { //ad
	return a._exists
}

// Deleted provides information if the Ad has been deleted from the database.
func (a *Ad) Deleted() bool {
	return a._deleted
}

// Get table name
func GetAdTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "ad", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Ad to the database.
func (a *Ad) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if a._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`aiid, apid, starttime, endtime, dailybudget, chargingway, price, addby, updated, created, status` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, a.Aiid, a.Apid, a.Starttime, a.Endtime, a.Dailybudget, a.Chargingway, a.Price, a.Addby, a.Updated, a.Created, a.Status)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, a.Aiid, a.Apid, a.Starttime, a.Endtime, a.Dailybudget, a.Chargingway, a.Price, a.Addby, a.Updated, a.Created, a.Status)
	} else {
		res, err = dbConn.Exec(sqlstr, a.Aiid, a.Apid, a.Starttime, a.Endtime, a.Dailybudget, a.Chargingway, a.Price, a.Addby, a.Updated, a.Created, a.Status)
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
	a.Aid = int(id)
	a._exists = true

	return nil
}

// Update updates the Ad in the database.
func (a *Ad) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if a._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`aiid = ?, apid = ?, starttime = ?, endtime = ?, dailybudget = ?, chargingway = ?, price = ?, addby = ?, updated = ?, created = ?, status = ?` +
		` WHERE aid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, a.Aiid, a.Apid, a.Starttime, a.Endtime, a.Dailybudget, a.Chargingway, a.Price, a.Addby, a.Updated, a.Created, a.Status, a.Aid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, a.Aiid, a.Apid, a.Starttime, a.Endtime, a.Dailybudget, a.Chargingway, a.Price, a.Addby, a.Updated, a.Created, a.Status, a.Aid)
	} else {
		_, err = dbConn.Exec(sqlstr, a.Aiid, a.Apid, a.Starttime, a.Endtime, a.Dailybudget, a.Chargingway, a.Price, a.Addby, a.Updated, a.Created, a.Status, a.Aid)
	}
	return err
}

// Save saves the Ad to the database.
func (a *Ad) Save(ctx context.Context) error {
	if a.Exists() {
		return a.Update(ctx)
	}

	return a.Insert(ctx)
}

// Delete deletes the Ad from the database.
func (a *Ad) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if a._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetAdTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE aid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, a.Aid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, a.Aid)
	} else {
		_, err = dbConn.Exec(sqlstr, a.Aid)
	}

	if err != nil {
		return err
	}

	// set deleted
	a._deleted = true

	return nil
}

// AdByAid retrieves a row from 'aypcddg.ad' as a Ad.
//
// Generated from index 'ad_aid_pkey'.
func AdByAid(ctx context.Context, aid int, key ...interface{}) (*Ad, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetAdTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`aid, aiid, apid, starttime, endtime, dailybudget, chargingway, price, addby, updated, created, status ` +
		`FROM ` + tableName +
		` WHERE aid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, aid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	a := Ad{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, aid).Scan(&a.Aid, &a.Aiid, &a.Apid, &a.Starttime, &a.Endtime, &a.Dailybudget, &a.Chargingway, &a.Price, &a.Addby, &a.Updated, &a.Created, &a.Status)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, aid).Scan(&a.Aid, &a.Aiid, &a.Apid, &a.Starttime, &a.Endtime, &a.Dailybudget, &a.Chargingway, &a.Price, &a.Addby, &a.Updated, &a.Created, &a.Status)
		if err != nil {
			return nil, err
		}
	}

	return &a, nil
}
