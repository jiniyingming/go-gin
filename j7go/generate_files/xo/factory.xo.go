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

// Factory represents a row from 'aypcddg.factory'.
type Factory struct {
	Fid           uint           `json:"fid"`            // fid
	Title         sql.NullString `json:"title"`          // title
	Products      sql.NullString `json:"products"`       // products
	Content       sql.NullString `json:"content"`        // content
	SaleNum       uint           `json:"sale_num"`       // sale_num
	BrandNum      uint           `json:"brand_num"`      // brand_num
	ContractBegin uint           `json:"contract_begin"` // contract_begin
	ContractEnd   uint           `json:"contract_end"`   // contract_end
	Fpid          uint           `json:"fpid"`           // fpid
	AddonFeatures sql.NullString `json:"addon_features"` // addon_features
	ProductRank   float64        `json:"product_rank"`   // product_rank
	Created       uint           `json:"created"`        // created
	CreatorID     uint           `json:"creator_id"`     // creator_id
	Updated       uint           `json:"updated"`        // updated
	Status        int8           `json:"status"`         // status
	ServicePhone  sql.NullString `json:"service_phone"`  // service_phone
	ShopName      sql.NullString `json:"shop_name"`      // shop_name
	ShopSlogan    sql.NullString `json:"shop_slogan"`    // shop_slogan
	Auid          sql.NullInt64  `json:"auid"`           // auid
	Goodstatus    int8           `json:"goodstatus"`     // goodstatus
	ShopSet       string         `json:"shop_set"`       // shop_set
	WxMchID       sql.NullString `json:"wx_mch_id"`      // wx_mch_id
	Flogo         sql.NullString `json:"flogo"`          // flogo
	DescPics      sql.NullString `json:"desc_pics"`      // desc_pics
	MidAdminID    sql.NullInt64  `json:"mid_admin_id"`   // mid_admin_id
	MidAdminName  sql.NullString `json:"mid_admin_name"` // mid_admin_name
	CreatedAt     mysql.NullTime `json:"created_at"`     // created_at
	UpdatedAt     mysql.NullTime `json:"updated_at"`     // updated_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Factory exists in the database.
func (f *Factory) Exists() bool { //factory
	return f._exists
}

// Deleted provides information if the Factory has been deleted from the database.
func (f *Factory) Deleted() bool {
	return f._deleted
}

// Get table name
func GetFactoryTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "factory", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the Factory to the database.
func (f *Factory) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if f._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`title, products, content, sale_num, brand_num, contract_begin, contract_end, fpid, addon_features, product_rank, created, creator_id, updated, status, service_phone, shop_name, shop_slogan, auid, goodstatus, shop_set, wx_mch_id, flogo, desc_pics, mid_admin_id, mid_admin_name, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, f.Title, f.Products, f.Content, f.SaleNum, f.BrandNum, f.ContractBegin, f.ContractEnd, f.Fpid, f.AddonFeatures, f.ProductRank, f.Created, f.CreatorID, f.Updated, f.Status, f.ServicePhone, f.ShopName, f.ShopSlogan, f.Auid, f.Goodstatus, f.ShopSet, f.WxMchID, f.Flogo, f.DescPics, f.MidAdminID, f.MidAdminName, f.CreatedAt, f.UpdatedAt)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, f.Title, f.Products, f.Content, f.SaleNum, f.BrandNum, f.ContractBegin, f.ContractEnd, f.Fpid, f.AddonFeatures, f.ProductRank, f.Created, f.CreatorID, f.Updated, f.Status, f.ServicePhone, f.ShopName, f.ShopSlogan, f.Auid, f.Goodstatus, f.ShopSet, f.WxMchID, f.Flogo, f.DescPics, f.MidAdminID, f.MidAdminName, f.CreatedAt, f.UpdatedAt)
	} else {
		res, err = dbConn.Exec(sqlstr, f.Title, f.Products, f.Content, f.SaleNum, f.BrandNum, f.ContractBegin, f.ContractEnd, f.Fpid, f.AddonFeatures, f.ProductRank, f.Created, f.CreatorID, f.Updated, f.Status, f.ServicePhone, f.ShopName, f.ShopSlogan, f.Auid, f.Goodstatus, f.ShopSet, f.WxMchID, f.Flogo, f.DescPics, f.MidAdminID, f.MidAdminName, f.CreatedAt, f.UpdatedAt)
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
	f.Fid = uint(id)
	f._exists = true

	return nil
}

// Update updates the Factory in the database.
func (f *Factory) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if f._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`title = ?, products = ?, content = ?, sale_num = ?, brand_num = ?, contract_begin = ?, contract_end = ?, fpid = ?, addon_features = ?, product_rank = ?, created = ?, creator_id = ?, updated = ?, status = ?, service_phone = ?, shop_name = ?, shop_slogan = ?, auid = ?, goodstatus = ?, shop_set = ?, wx_mch_id = ?, flogo = ?, desc_pics = ?, mid_admin_id = ?, mid_admin_name = ?, created_at = ?, updated_at = ?` +
		` WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, f.Title, f.Products, f.Content, f.SaleNum, f.BrandNum, f.ContractBegin, f.ContractEnd, f.Fpid, f.AddonFeatures, f.ProductRank, f.Created, f.CreatorID, f.Updated, f.Status, f.ServicePhone, f.ShopName, f.ShopSlogan, f.Auid, f.Goodstatus, f.ShopSet, f.WxMchID, f.Flogo, f.DescPics, f.MidAdminID, f.MidAdminName, f.CreatedAt, f.UpdatedAt, f.Fid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, f.Title, f.Products, f.Content, f.SaleNum, f.BrandNum, f.ContractBegin, f.ContractEnd, f.Fpid, f.AddonFeatures, f.ProductRank, f.Created, f.CreatorID, f.Updated, f.Status, f.ServicePhone, f.ShopName, f.ShopSlogan, f.Auid, f.Goodstatus, f.ShopSet, f.WxMchID, f.Flogo, f.DescPics, f.MidAdminID, f.MidAdminName, f.CreatedAt, f.UpdatedAt, f.Fid)
	} else {
		_, err = dbConn.Exec(sqlstr, f.Title, f.Products, f.Content, f.SaleNum, f.BrandNum, f.ContractBegin, f.ContractEnd, f.Fpid, f.AddonFeatures, f.ProductRank, f.Created, f.CreatorID, f.Updated, f.Status, f.ServicePhone, f.ShopName, f.ShopSlogan, f.Auid, f.Goodstatus, f.ShopSet, f.WxMchID, f.Flogo, f.DescPics, f.MidAdminID, f.MidAdminName, f.CreatedAt, f.UpdatedAt, f.Fid)
	}
	return err
}

// Save saves the Factory to the database.
func (f *Factory) Save(ctx context.Context) error {
	if f.Exists() {
		return f.Update(ctx)
	}

	return f.Insert(ctx)
}

// Delete deletes the Factory from the database.
func (f *Factory) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if f._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetFactoryTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, f.Fid)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, f.Fid)
	} else {
		_, err = dbConn.Exec(sqlstr, f.Fid)
	}

	if err != nil {
		return err
	}

	// set deleted
	f._deleted = true

	return nil
}

// FactoryByFid retrieves a row from 'aypcddg.factory' as a Factory.
//
// Generated from index 'factory_fid_pkey'.
func FactoryByFid(ctx context.Context, fid uint, key ...interface{}) (*Factory, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetFactoryTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`fid, title, products, content, sale_num, brand_num, contract_begin, contract_end, fpid, addon_features, product_rank, created, creator_id, updated, status, service_phone, shop_name, shop_slogan, auid, goodstatus, shop_set, wx_mch_id, flogo, desc_pics, mid_admin_id, mid_admin_name, created_at, updated_at ` +
		`FROM ` + tableName +
		` WHERE fid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, fid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	f := Factory{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, fid).Scan(&f.Fid, &f.Title, &f.Products, &f.Content, &f.SaleNum, &f.BrandNum, &f.ContractBegin, &f.ContractEnd, &f.Fpid, &f.AddonFeatures, &f.ProductRank, &f.Created, &f.CreatorID, &f.Updated, &f.Status, &f.ServicePhone, &f.ShopName, &f.ShopSlogan, &f.Auid, &f.Goodstatus, &f.ShopSet, &f.WxMchID, &f.Flogo, &f.DescPics, &f.MidAdminID, &f.MidAdminName, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, fid).Scan(&f.Fid, &f.Title, &f.Products, &f.Content, &f.SaleNum, &f.BrandNum, &f.ContractBegin, &f.ContractEnd, &f.Fpid, &f.AddonFeatures, &f.ProductRank, &f.Created, &f.CreatorID, &f.Updated, &f.Status, &f.ServicePhone, &f.ShopName, &f.ShopSlogan, &f.Auid, &f.Goodstatus, &f.ShopSet, &f.WxMchID, &f.Flogo, &f.DescPics, &f.MidAdminID, &f.MidAdminName, &f.CreatedAt, &f.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}

	return &f, nil
}
