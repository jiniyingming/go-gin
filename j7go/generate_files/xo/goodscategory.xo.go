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

	"github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

// GoodsCategory represents a row from 'aypcddg.goods_category'.
type GoodsCategory struct {
	Gcid                    uint           `json:"gcid"`                      // gcid
	Title                   sql.NullString `json:"title"`                     // title
	Description             sql.NullString `json:"description"`               // description
	Keywords                sql.NullString `json:"keywords"`                  // keywords
	Seq                     int16          `json:"seq"`                       // seq
	Slug                    sql.NullString `json:"slug"`                      // slug
	Pid                     uint           `json:"pid"`                       // pid
	Status                  sql.NullBool   `json:"status"`                    // status
	MidAdminID              sql.NullInt64  `json:"mid_admin_id"`              // mid_admin_id
	MidAdminName            sql.NullString `json:"mid_admin_name"`            // mid_admin_name
	Image                   sql.NullString `json:"image"`                     // image
	CreatedAt               mysql.NullTime `json:"created_at"`                // created_at
	UpdatedAt               mysql.NullTime `json:"updated_at"`                // updated_at
	GlobalMarkupPercent     float64        `json:"global_markup_percent"`     // global_markup_percent
	GlobalCommissionPercent float64        `json:"global_commission_percent"` // global_commission_percent

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the GoodsCategory exists in the database.
func (gc *GoodsCategory) Exists() bool { //goods_category
	return gc._exists
}

// Deleted provides information if the GoodsCategory has been deleted from the database.
func (gc *GoodsCategory) Deleted() bool {
	return gc._deleted
}

// Get table name
func GetGoodsCategoryTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "goods_category", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the GoodsCategory to the database.
func (gc *GoodsCategory) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if gc._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategoryTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`title, description, keywords, seq, slug, pid, status, mid_admin_id, mid_admin_name, image, created_at, updated_at, global_markup_percent, global_commission_percent` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gc.Title, gc.Description, gc.Keywords, gc.Seq, gc.Slug, gc.Pid, gc.Status, gc.MidAdminID, gc.MidAdminName, gc.Image, gc.CreatedAt, gc.UpdatedAt, gc.GlobalMarkupPercent, gc.GlobalCommissionPercent)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, gc.Title, gc.Description, gc.Keywords, gc.Seq, gc.Slug, gc.Pid, gc.Status, gc.MidAdminID, gc.MidAdminName, gc.Image, gc.CreatedAt, gc.UpdatedAt, gc.GlobalMarkupPercent, gc.GlobalCommissionPercent)
	} else {
		res, err = dbConn.Exec(sqlstr, gc.Title, gc.Description, gc.Keywords, gc.Seq, gc.Slug, gc.Pid, gc.Status, gc.MidAdminID, gc.MidAdminName, gc.Image, gc.CreatedAt, gc.UpdatedAt, gc.GlobalMarkupPercent, gc.GlobalCommissionPercent)
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
	gc.Gcid = uint(id)
	gc._exists = true

	return nil
}

// Update updates the GoodsCategory in the database.
func (gc *GoodsCategory) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gc._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategoryTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, description = ?, keywords = ?, seq = ?, slug = ?, pid = ?, status = ?, mid_admin_id = ?, mid_admin_name = ?, image = ?, created_at = ?, updated_at = ?, global_markup_percent = ?, global_commission_percent = ?` +
		` WHERE gcid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gc.Title, gc.Description, gc.Keywords, gc.Seq, gc.Slug, gc.Pid, gc.Status, gc.MidAdminID, gc.MidAdminName, gc.Image, gc.CreatedAt, gc.UpdatedAt, gc.GlobalMarkupPercent, gc.GlobalCommissionPercent, gc.Gcid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gc.Title, gc.Description, gc.Keywords, gc.Seq, gc.Slug, gc.Pid, gc.Status, gc.MidAdminID, gc.MidAdminName, gc.Image, gc.CreatedAt, gc.UpdatedAt, gc.GlobalMarkupPercent, gc.GlobalCommissionPercent, gc.Gcid)
	} else {
		_, err = dbConn.Exec(sqlstr, gc.Title, gc.Description, gc.Keywords, gc.Seq, gc.Slug, gc.Pid, gc.Status, gc.MidAdminID, gc.MidAdminName, gc.Image, gc.CreatedAt, gc.UpdatedAt, gc.GlobalMarkupPercent, gc.GlobalCommissionPercent, gc.Gcid)
	}
	return err
}

// Save saves the GoodsCategory to the database.
func (gc *GoodsCategory) Save(ctx context.Context) error {
	if gc.Exists() {
		return gc.Update(ctx)
	}

	return gc.Insert(ctx)
}

// Delete deletes the GoodsCategory from the database.
func (gc *GoodsCategory) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if gc._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetGoodsCategoryTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE gcid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gc.Gcid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, gc.Gcid)
	} else {
		_, err = dbConn.Exec(sqlstr, gc.Gcid)
	}

	if err != nil {
		return err
	}

	// set deleted
	gc._deleted = true

	return nil
}

// GoodsCategoryByGcid retrieves a row from 'aypcddg.goods_category' as a GoodsCategory.
//
// Generated from index 'goods_category_gcid_pkey'.
func GoodsCategoryByGcid(ctx context.Context, gcid uint, key ...interface{}) (*GoodsCategory, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategoryTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcid, title, description, keywords, seq, slug, pid, status, mid_admin_id, mid_admin_name, image, created_at, updated_at, global_markup_percent, global_commission_percent ` +
		`FROM ` + tableName +
		` WHERE gcid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, gcid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	gc := GoodsCategory{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, gcid).Scan(&gc.Gcid, &gc.Title, &gc.Description, &gc.Keywords, &gc.Seq, &gc.Slug, &gc.Pid, &gc.Status, &gc.MidAdminID, &gc.MidAdminName, &gc.Image, &gc.CreatedAt, &gc.UpdatedAt, &gc.GlobalMarkupPercent, &gc.GlobalCommissionPercent)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, gcid).Scan(&gc.Gcid, &gc.Title, &gc.Description, &gc.Keywords, &gc.Seq, &gc.Slug, &gc.Pid, &gc.Status, &gc.MidAdminID, &gc.MidAdminName, &gc.Image, &gc.CreatedAt, &gc.UpdatedAt, &gc.GlobalMarkupPercent, &gc.GlobalCommissionPercent)
		if err != nil {
			return nil, err
		}
	}

	return &gc, nil
}

// GoodsCategoriesBySeq retrieves a row from 'aypcddg.goods_category' as a GoodsCategory.
//
// Generated from index 'seq'.
func GoodsCategoriesBySeq(ctx context.Context, seq int16, key ...interface{}) ([]*GoodsCategory, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategoryTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcid, title, description, keywords, seq, slug, pid, status, mid_admin_id, mid_admin_name, image, created_at, updated_at, global_markup_percent, global_commission_percent ` +
		`FROM ` + tableName +
		` WHERE seq = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, seq)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, seq)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, seq)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*GoodsCategory, 0)
	for queryData.Next() {
		gc := GoodsCategory{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&gc.Gcid, &gc.Title, &gc.Description, &gc.Keywords, &gc.Seq, &gc.Slug, &gc.Pid, &gc.Status, &gc.MidAdminID, &gc.MidAdminName, &gc.Image, &gc.CreatedAt, &gc.UpdatedAt, &gc.GlobalMarkupPercent, &gc.GlobalCommissionPercent)
		if err != nil {
			return nil, err
		}

		res = append(res, &gc)
	}

	return res, nil
}

// GoodsCategoryBySlug retrieves a row from 'aypcddg.goods_category' as a GoodsCategory.
//
// Generated from index 'slug'.
func GoodsCategoryBySlug(ctx context.Context, slug sql.NullString, key ...interface{}) (*GoodsCategory, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetGoodsCategoryTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`gcid, title, description, keywords, seq, slug, pid, status, mid_admin_id, mid_admin_name, image, created_at, updated_at, global_markup_percent, global_commission_percent ` +
		`FROM ` + tableName +
		` WHERE slug = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, slug)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	gc := GoodsCategory{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, slug).Scan(&gc.Gcid, &gc.Title, &gc.Description, &gc.Keywords, &gc.Seq, &gc.Slug, &gc.Pid, &gc.Status, &gc.MidAdminID, &gc.MidAdminName, &gc.Image, &gc.CreatedAt, &gc.UpdatedAt, &gc.GlobalMarkupPercent, &gc.GlobalCommissionPercent)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, slug).Scan(&gc.Gcid, &gc.Title, &gc.Description, &gc.Keywords, &gc.Seq, &gc.Slug, &gc.Pid, &gc.Status, &gc.MidAdminID, &gc.MidAdminName, &gc.Image, &gc.CreatedAt, &gc.UpdatedAt, &gc.GlobalMarkupPercent, &gc.GlobalCommissionPercent)
		if err != nil {
			return nil, err
		}
	}

	return &gc, nil
}
