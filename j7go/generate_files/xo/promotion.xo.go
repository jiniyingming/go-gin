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

// Promotion represents a row from 'aypcddg.promotions'.
type Promotion struct {
	Pid       uint           `json:"pid"`        // pid
	Fid       sql.NullInt64  `json:"fid"`        // fid
	Sid       sql.NullInt64  `json:"sid"`        // sid
	Type      int8           `json:"type"`       // type
	Rule      sql.NullString `json:"rule"`       // rule
	Content   sql.NullString `json:"content"`    // content
	Title     sql.NullString `json:"title"`      // title
	AddWord   sql.NullString `json:"add_word"`   // add_word
	Note      sql.NullString `json:"note"`       // note
	Des       string         `json:"des"`        // des
	BeginTime sql.NullInt64  `json:"begin_time"` // begin_time
	EndTime   sql.NullInt64  `json:"end_time"`   // end_time
	Created   sql.NullInt64  `json:"created"`    // created

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Promotion exists in the database.
func (p *Promotion) Exists() bool { //promotions
	return p._exists
}

// Deleted provides information if the Promotion has been deleted from the database.
func (p *Promotion) Deleted() bool {
	return p._deleted
}

// Get table name
func GetPromotionTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "promotions", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Promotion to the database.
func (p *Promotion) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPromotionTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`fid, sid, type, rule, content, title, add_word, note, des, begin_time, end_time, created` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, p.Fid, p.Sid, p.Type, p.Rule, p.Content, p.Title, p.AddWord, p.Note, p.Des, p.BeginTime, p.EndTime, p.Created)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, p.Fid, p.Sid, p.Type, p.Rule, p.Content, p.Title, p.AddWord, p.Note, p.Des, p.BeginTime, p.EndTime, p.Created)
	} else {
		res, err = dbConn.Exec(sqlstr, p.Fid, p.Sid, p.Type, p.Rule, p.Content, p.Title, p.AddWord, p.Note, p.Des, p.BeginTime, p.EndTime, p.Created)
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
	p.Pid = uint(id)
	p._exists = true

	return nil
}

// Update updates the Promotion in the database.
func (p *Promotion) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if p._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPromotionTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`fid = ?, sid = ?, type = ?, rule = ?, content = ?, title = ?, add_word = ?, note = ?, des = ?, begin_time = ?, end_time = ?, created = ?` +
		` WHERE pid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, p.Fid, p.Sid, p.Type, p.Rule, p.Content, p.Title, p.AddWord, p.Note, p.Des, p.BeginTime, p.EndTime, p.Created, p.Pid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, p.Fid, p.Sid, p.Type, p.Rule, p.Content, p.Title, p.AddWord, p.Note, p.Des, p.BeginTime, p.EndTime, p.Created, p.Pid)
	} else {
		_, err = dbConn.Exec(sqlstr, p.Fid, p.Sid, p.Type, p.Rule, p.Content, p.Title, p.AddWord, p.Note, p.Des, p.BeginTime, p.EndTime, p.Created, p.Pid)
	}
	return err
}

// Save saves the Promotion to the database.
func (p *Promotion) Save(ctx context.Context) error {
	if p.Exists() {
		return p.Update(ctx)
	}

	return p.Insert(ctx)
}

// Delete deletes the Promotion from the database.
func (p *Promotion) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if p._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetPromotionTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE pid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, p.Pid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, p.Pid)
	} else {
		_, err = dbConn.Exec(sqlstr, p.Pid)
	}

	if err != nil {
		return err
	}

	// set deleted
	p._deleted = true

	return nil
}

// PromotionByPid retrieves a row from 'aypcddg.promotions' as a Promotion.
//
// Generated from index 'promotions_pid_pkey'.
func PromotionByPid(ctx context.Context, pid uint, key ...interface{}) (*Promotion, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetPromotionTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`pid, fid, sid, type, rule, content, title, add_word, note, des, begin_time, end_time, created ` +
		`FROM ` + tableName +
		` WHERE pid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, pid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	p := Promotion{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, pid).Scan(&p.Pid, &p.Fid, &p.Sid, &p.Type, &p.Rule, &p.Content, &p.Title, &p.AddWord, &p.Note, &p.Des, &p.BeginTime, &p.EndTime, &p.Created)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, pid).Scan(&p.Pid, &p.Fid, &p.Sid, &p.Type, &p.Rule, &p.Content, &p.Title, &p.AddWord, &p.Note, &p.Des, &p.BeginTime, &p.EndTime, &p.Created)
		if err != nil {
			return nil, err
		}
	}

	return &p, nil
}
