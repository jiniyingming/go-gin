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

// FactoryGroupSale represents a row from 'aypcddg.factory_group_sales'.
type FactoryGroupSale struct {
	Fgid int64 `json:"fgid"` // fgid
	Fsid int64 `json:"fsid"` // fsid
	Fbid int64 `json:"fbid"` // fbid
	Fid  int   `json:"fid"`  // fid
	Sid  int   `json:"sid"`  // sid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the FactoryGroupSale exists in the database.
func (fgs *FactoryGroupSale) Exists() bool { //factory_group_sales
	return fgs._exists
}

// Deleted provides information if the FactoryGroupSale has been deleted from the database.
func (fgs *FactoryGroupSale) Deleted() bool {
	return fgs._deleted
}

// Get table name
func GetFactoryGroupSaleTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "factory_group_sales", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the FactoryGroupSale to the database.
func (fgs *FactoryGroupSale) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if fgs._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryGroupSaleTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key must be provided
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fgid, fsid, fbid, fid, sid` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fgs.Fgid, fgs.Fsid, fgs.Fbid, fgs.Fid, fgs.Sid)))
	if tx != nil {
		res, err = tx.Exec(sqlstr, fgs.Fgid, fgs.Fsid, fgs.Fbid, fgs.Fid, fgs.Sid)
	} else {
		res, err = dbConn.Exec(sqlstr, fgs.Fgid, fgs.Fsid, fgs.Fbid, fgs.Fid, fgs.Sid)
	}

	if err != nil {
		return err
	}

	// set existence
	fgs._exists = true

	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	fgs.Fbid = int64(id)
	fgs._exists = true

	return nil
}

// Update updates the FactoryGroupSale in the database.
func (fgs *FactoryGroupSale) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fgs._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryGroupSaleTableName(key...)
	if err != nil {
		return err
	}

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fid = ?, sid = ?` +
		` WHERE fgid = ? AND fsid = ? AND fbid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fgs.Fid, fgs.Sid, fgs.Fgid, fgs.Fsid, fgs.Fbid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, fgs.Fid, fgs.Sid, fgs.Fgid, fgs.Fsid, fgs.Fbid)
	} else {
		_, err = dbConn.Exec(sqlstr, fgs.Fid, fgs.Sid, fgs.Fgid, fgs.Fsid, fgs.Fbid)
	}
	return err
}

// Save saves the FactoryGroupSale to the database.
func (fgs *FactoryGroupSale) Save(ctx context.Context) error {
	if fgs.Exists() {
		return fgs.Update(ctx)
	}

	return fgs.Insert(ctx)
}

// Delete deletes the FactoryGroupSale from the database.
func (fgs *FactoryGroupSale) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if fgs._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryGroupSaleTableName(key...)
	if err != nil {
		return err
	}
	//3

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fbid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fgs.Fbid)))

	if tx != nil {
		_, err = tx.Exec(sqlstr, fgs.Fbid)
	} else {
		_, err = dbConn.Exec(sqlstr, fgs.Fbid)
	}
	if err != nil {
		return err
	}

	// set deleted
	fgs._deleted = true

	return nil
}

// FactoryGroupSaleByFgidFsidFbid retrieves a row from 'aypcddg.factory_group_sales' as a FactoryGroupSale.
//
// Generated from index 'all'.
func FactoryGroupSaleByFgidFsidFbid(ctx context.Context, fgid int64, fsid int64, fbid int64, key ...interface{}) (*FactoryGroupSale, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryGroupSaleTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fgid, fsid, fbid, fid, sid ` +
		`FROM ` + tableName +
		` WHERE fgid = ? AND fsid = ? AND fbid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fgid, fsid, fbid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	fgs := FactoryGroupSale{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fgid, fsid, fbid).Scan(&fgs.Fgid, &fgs.Fsid, &fgs.Fbid, &fgs.Fid, &fgs.Sid)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fgid, fsid, fbid).Scan(&fgs.Fgid, &fgs.Fsid, &fgs.Fbid, &fgs.Fid, &fgs.Sid)
		if err != nil {
			return nil, err
		}
	}

	return &fgs, nil
}

// FactoryGroupSaleByFbid retrieves a row from 'aypcddg.factory_group_sales' as a FactoryGroupSale.
//
// Generated from index 'factory_group_sales_fbid_pkey'.
func FactoryGroupSaleByFbid(ctx context.Context, fbid int64, key ...interface{}) (*FactoryGroupSale, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryGroupSaleTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fgid, fsid, fbid, fid, sid ` +
		`FROM ` + tableName +
		` WHERE fbid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fbid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	fgs := FactoryGroupSale{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fbid).Scan(&fgs.Fgid, &fgs.Fsid, &fgs.Fbid, &fgs.Fid, &fgs.Sid)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fbid).Scan(&fgs.Fgid, &fgs.Fsid, &fgs.Fbid, &fgs.Fid, &fgs.Sid)
		if err != nil {
			return nil, err
		}
	}

	return &fgs, nil
}
