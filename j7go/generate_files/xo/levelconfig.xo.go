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

// LevelConfig represents a row from 'ddg_local.level_config'.
type LevelConfig struct {
	ID                          int64           `json:"id"`                              // id
	UserType                    sql.NullInt64   `json:"user_type"`                       // user_type
	LevelContent                sql.NullString  `json:"level_content"`                   // level_content
	LeveTitle                   sql.NullString  `json:"leve_title"`                      // leve_title
	LevelSeq                    sql.NullInt64   `json:"level_seq"`                       // level_seq
	DoStatus                    sql.NullInt64   `json:"do_status"`                       // do_status
	UpdateAt                    mysql.NullTime  `json:"update_at"`                       // update_at
	CreatedAt                   mysql.NullTime  `json:"created_at"`                      // created_at
	LevelThresholdSellMoney     sql.NullFloat64 `json:"level_threshold_sell_money"`      // level_threshold_sell_money
	LevelDownThresholdSellMoney sql.NullFloat64 `json:"level_down_threshold_sell_money"` // level_down_threshold_sell_money
	TeamCommissionProportion    sql.NullInt64   `json:"team_commission_proportion"`      // team_commission_proportion
	DirectCommissionProportion  sql.NullInt64   `json:"direct_commission_proportion"`    // direct_commission_proportion
	LevelIconLengthURL          string          `json:"level_icon_length_url"`           // level_icon_length_url
	LevelIconGreyURL            string          `json:"level_icon_grey_url"`             // level_icon_grey_url
	LevelIconURL                string          `json:"level_icon_url"`                  // level_icon_url

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the LevelConfig exists in the database.
func (lc *LevelConfig) Exists() bool { //level_config
	return lc._exists
}

// Deleted provides information if the LevelConfig has been deleted from the database.
func (lc *LevelConfig) Deleted() bool {
	return lc._deleted
}

// Get table name
func GetLevelConfigTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable("ddg_local", "level_config", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the LevelConfig to the database.
func (lc *LevelConfig) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if lc._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetLevelConfigTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`user_type, level_content, leve_title, level_seq, do_status, update_at, created_at, level_threshold_sell_money, level_down_threshold_sell_money, team_commission_proportion, direct_commission_proportion, level_icon_length_url, level_icon_grey_url, level_icon_url` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, lc.UserType, lc.LevelContent, lc.LeveTitle, lc.LevelSeq, lc.DoStatus, lc.UpdateAt, lc.CreatedAt, lc.LevelThresholdSellMoney, lc.LevelDownThresholdSellMoney, lc.TeamCommissionProportion, lc.DirectCommissionProportion, lc.LevelIconLengthURL, lc.LevelIconGreyURL, lc.LevelIconURL)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, lc.UserType, lc.LevelContent, lc.LeveTitle, lc.LevelSeq, lc.DoStatus, lc.UpdateAt, lc.CreatedAt, lc.LevelThresholdSellMoney, lc.LevelDownThresholdSellMoney, lc.TeamCommissionProportion, lc.DirectCommissionProportion, lc.LevelIconLengthURL, lc.LevelIconGreyURL, lc.LevelIconURL)
	} else {
		res, err = dbConn.Exec(sqlstr, lc.UserType, lc.LevelContent, lc.LeveTitle, lc.LevelSeq, lc.DoStatus, lc.UpdateAt, lc.CreatedAt, lc.LevelThresholdSellMoney, lc.LevelDownThresholdSellMoney, lc.TeamCommissionProportion, lc.DirectCommissionProportion, lc.LevelIconLengthURL, lc.LevelIconGreyURL, lc.LevelIconURL)
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
	lc.ID = int64(id)
	lc._exists = true

	return nil
}

// Update updates the LevelConfig in the database.
func (lc *LevelConfig) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if lc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetLevelConfigTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`user_type = ?, level_content = ?, leve_title = ?, level_seq = ?, do_status = ?, update_at = ?, created_at = ?, level_threshold_sell_money = ?, level_down_threshold_sell_money = ?, team_commission_proportion = ?, direct_commission_proportion = ?, level_icon_length_url = ?, level_icon_grey_url = ?, level_icon_url = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, lc.UserType, lc.LevelContent, lc.LeveTitle, lc.LevelSeq, lc.DoStatus, lc.UpdateAt, lc.CreatedAt, lc.LevelThresholdSellMoney, lc.LevelDownThresholdSellMoney, lc.TeamCommissionProportion, lc.DirectCommissionProportion, lc.LevelIconLengthURL, lc.LevelIconGreyURL, lc.LevelIconURL, lc.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, lc.UserType, lc.LevelContent, lc.LeveTitle, lc.LevelSeq, lc.DoStatus, lc.UpdateAt, lc.CreatedAt, lc.LevelThresholdSellMoney, lc.LevelDownThresholdSellMoney, lc.TeamCommissionProportion, lc.DirectCommissionProportion, lc.LevelIconLengthURL, lc.LevelIconGreyURL, lc.LevelIconURL, lc.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, lc.UserType, lc.LevelContent, lc.LeveTitle, lc.LevelSeq, lc.DoStatus, lc.UpdateAt, lc.CreatedAt, lc.LevelThresholdSellMoney, lc.LevelDownThresholdSellMoney, lc.TeamCommissionProportion, lc.DirectCommissionProportion, lc.LevelIconLengthURL, lc.LevelIconGreyURL, lc.LevelIconURL, lc.ID)
	}
	return err
}

// Save saves the LevelConfig to the database.
func (lc *LevelConfig) Save(ctx context.Context) error {
	if lc.Exists() {
		return lc.Update(ctx)
	}

	return lc.Insert(ctx)
}

// Delete deletes the LevelConfig from the database.
func (lc *LevelConfig) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if lc._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetLevelConfigTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, lc.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, lc.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, lc.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	lc._deleted = true

	return nil
}

// LevelConfigByID retrieves a row from 'ddg_local.level_config' as a LevelConfig.
//
// Generated from index 'level_config_id_pkey'.
func LevelConfigByID(ctx context.Context, id int64, key ...interface{}) (*LevelConfig, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetLevelConfigTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, user_type, level_content, leve_title, level_seq, do_status, update_at, created_at, level_threshold_sell_money, level_down_threshold_sell_money, team_commission_proportion, direct_commission_proportion, level_icon_length_url, level_icon_grey_url, level_icon_url ` +
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
	lc := LevelConfig{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&lc.ID, &lc.UserType, &lc.LevelContent, &lc.LeveTitle, &lc.LevelSeq, &lc.DoStatus, &lc.UpdateAt, &lc.CreatedAt, &lc.LevelThresholdSellMoney, &lc.LevelDownThresholdSellMoney, &lc.TeamCommissionProportion, &lc.DirectCommissionProportion, &lc.LevelIconLengthURL, &lc.LevelIconGreyURL, &lc.LevelIconURL)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&lc.ID, &lc.UserType, &lc.LevelContent, &lc.LeveTitle, &lc.LevelSeq, &lc.DoStatus, &lc.UpdateAt, &lc.CreatedAt, &lc.LevelThresholdSellMoney, &lc.LevelDownThresholdSellMoney, &lc.TeamCommissionProportion, &lc.DirectCommissionProportion, &lc.LevelIconLengthURL, &lc.LevelIconGreyURL, &lc.LevelIconURL)
		if err != nil {
			return nil, err
		}
	}

	return &lc, nil
}