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

// UserCoupon represents a row from 'aypcddg.user_coupons'.
type UserCoupon struct {
	Ucid      uint           `json:"ucid"`       // ucid
	UID       sql.NullInt64  `json:"uid"`        // uid
	Cid       sql.NullInt64  `json:"cid"`        // cid
	UsedNum   sql.NullInt64  `json:"used_num"`   // used_num
	TotalNum  sql.NullInt64  `json:"total_num"`  // total_num
	Oids      sql.NullString `json:"oids"`       // oids
	BeginTime sql.NullInt64  `json:"begin_time"` // begin_time
	EndTime   sql.NullInt64  `json:"end_time"`   // end_time
	Fid       sql.NullInt64  `json:"fid"`        // fid
	Sid       sql.NullInt64  `json:"sid"`        // sid

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the UserCoupon exists in the database.
func (uc *UserCoupon) Exists() bool { //user_coupons
	return uc._exists
}

// Deleted provides information if the UserCoupon has been deleted from the database.
func (uc *UserCoupon) Deleted() bool {
	return uc._deleted
}

// Get table name
func GetUserCouponTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "user_coupons", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the UserCoupon to the database.
func (uc *UserCoupon) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if uc._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserCouponTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`uid, cid, used_num, total_num, oids, begin_time, end_time, fid, sid` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uc.UID, uc.Cid, uc.UsedNum, uc.TotalNum, uc.Oids, uc.BeginTime, uc.EndTime, uc.Fid, uc.Sid)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, uc.UID, uc.Cid, uc.UsedNum, uc.TotalNum, uc.Oids, uc.BeginTime, uc.EndTime, uc.Fid, uc.Sid)
	} else {
		res, err = dbConn.Exec(sqlstr, uc.UID, uc.Cid, uc.UsedNum, uc.TotalNum, uc.Oids, uc.BeginTime, uc.EndTime, uc.Fid, uc.Sid)
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
	uc.Ucid = uint(id)
	uc._exists = true

	return nil
}

// Update updates the UserCoupon in the database.
func (uc *UserCoupon) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if uc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserCouponTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`uid = ?, cid = ?, used_num = ?, total_num = ?, oids = ?, begin_time = ?, end_time = ?, fid = ?, sid = ?` +
		` WHERE ucid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uc.UID, uc.Cid, uc.UsedNum, uc.TotalNum, uc.Oids, uc.BeginTime, uc.EndTime, uc.Fid, uc.Sid, uc.Ucid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, uc.UID, uc.Cid, uc.UsedNum, uc.TotalNum, uc.Oids, uc.BeginTime, uc.EndTime, uc.Fid, uc.Sid, uc.Ucid)
	} else {
		_, err = dbConn.Exec(sqlstr, uc.UID, uc.Cid, uc.UsedNum, uc.TotalNum, uc.Oids, uc.BeginTime, uc.EndTime, uc.Fid, uc.Sid, uc.Ucid)
	}
	return err
}

// Save saves the UserCoupon to the database.
func (uc *UserCoupon) Save(ctx context.Context) error {
	if uc.Exists() {
		return uc.Update(ctx)
	}

	return uc.Insert(ctx)
}

// Delete deletes the UserCoupon from the database.
func (uc *UserCoupon) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if uc._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetUserCouponTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE ucid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uc.Ucid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, uc.Ucid)
	} else {
		_, err = dbConn.Exec(sqlstr, uc.Ucid)
	}

	if err != nil {
		return err
	}

	// set deleted
	uc._deleted = true

	return nil
}

// UserCouponByUcid retrieves a row from 'aypcddg.user_coupons' as a UserCoupon.
//
// Generated from index 'user_coupons_ucid_pkey'.
func UserCouponByUcid(ctx context.Context, ucid uint, key ...interface{}) (*UserCoupon, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetUserCouponTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`ucid, uid, cid, used_num, total_num, oids, begin_time, end_time, fid, sid ` +
		`FROM ` + tableName +
		` WHERE ucid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ucid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	uc := UserCoupon{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, ucid).Scan(&uc.Ucid, &uc.UID, &uc.Cid, &uc.UsedNum, &uc.TotalNum, &uc.Oids, &uc.BeginTime, &uc.EndTime, &uc.Fid, &uc.Sid)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, ucid).Scan(&uc.Ucid, &uc.UID, &uc.Cid, &uc.UsedNum, &uc.TotalNum, &uc.Oids, &uc.BeginTime, &uc.EndTime, &uc.Fid, &uc.Sid)
		if err != nil {
			return nil, err
		}
	}

	return &uc, nil
}