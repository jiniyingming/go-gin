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

// TuanCategory represents a row from 'aypcddg.tuan_category'.
type TuanCategory struct {
	Tcid    int            `json:"tcid"`    // tcid
	Title   sql.NullString `json:"title"`   // title
	Addby   sql.NullInt64  `json:"addby"`   // addby
	Created sql.NullInt64  `json:"created"` // created
	Updated sql.NullInt64  `json:"updated"` // updated
	Status  sql.NullBool   `json:"status"`  // status

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the TuanCategory exists in the database.
func (tc *TuanCategory) Exists() bool { //tuan_category
	return tc._exists
}

// Deleted provides information if the TuanCategory has been deleted from the database.
func (tc *TuanCategory) Deleted() bool {
	return tc._deleted
}

// Get table name
func GetTuanCategoryTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "tuan_category", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the TuanCategory to the database.
func (tc *TuanCategory) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if tc._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTuanCategoryTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`title, addby, created, updated, status` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tc.Title, tc.Addby, tc.Created, tc.Updated, tc.Status)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, tc.Title, tc.Addby, tc.Created, tc.Updated, tc.Status)
	} else {
		res, err = dbConn.Exec(sqlstr, tc.Title, tc.Addby, tc.Created, tc.Updated, tc.Status)
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
	tc.Tcid = int(id)
	tc._exists = true

	return nil
}

// Update updates the TuanCategory in the database.
func (tc *TuanCategory) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if tc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTuanCategoryTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, addby = ?, created = ?, updated = ?, status = ?` +
		` WHERE tcid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tc.Title, tc.Addby, tc.Created, tc.Updated, tc.Status, tc.Tcid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, tc.Title, tc.Addby, tc.Created, tc.Updated, tc.Status, tc.Tcid)
	} else {
		_, err = dbConn.Exec(sqlstr, tc.Title, tc.Addby, tc.Created, tc.Updated, tc.Status, tc.Tcid)
	}
	return err
}

// Save saves the TuanCategory to the database.
func (tc *TuanCategory) Save(ctx context.Context) error {
	if tc.Exists() {
		return tc.Update(ctx)
	}

	return tc.Insert(ctx)
}

// Delete deletes the TuanCategory from the database.
func (tc *TuanCategory) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if tc._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetTuanCategoryTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE tcid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tc.Tcid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, tc.Tcid)
	} else {
		_, err = dbConn.Exec(sqlstr, tc.Tcid)
	}

	if err != nil {
		return err
	}

	// set deleted
	tc._deleted = true

	return nil
}

// TuanCategoryByTcid retrieves a row from 'aypcddg.tuan_category' as a TuanCategory.
//
// Generated from index 'tuan_category_tcid_pkey'.
func TuanCategoryByTcid(ctx context.Context, tcid int, key ...interface{}) (*TuanCategory, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetTuanCategoryTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`tcid, title, addby, created, updated, status ` +
		`FROM ` + tableName +
		` WHERE tcid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, tcid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	tc := TuanCategory{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, tcid).Scan(&tc.Tcid, &tc.Title, &tc.Addby, &tc.Created, &tc.Updated, &tc.Status)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, tcid).Scan(&tc.Tcid, &tc.Title, &tc.Addby, &tc.Created, &tc.Updated, &tc.Status)
		if err != nil {
			return nil, err
		}
	}

	return &tc, nil
}
