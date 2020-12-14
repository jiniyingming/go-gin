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

// RefundOrder represents a row from 'aypcddg.refund_orders'.
type RefundOrder struct {
	ID                 int64           `json:"id"`                   // id
	RefundNo           sql.NullString  `json:"refund_no"`            // refund_no
	OrderID            sql.NullInt64   `json:"order_id"`             // order_id
	OrderNo            sql.NullString  `json:"order_no"`             // order_no
	FactoryID          sql.NullInt64   `json:"factory_id"`           // factory_id
	UID                sql.NullInt64   `json:"uid"`                  // uid
	OrderGsrid         sql.NullInt64   `json:"order_gsrid"`          // order_gsrid
	Gid                sql.NullInt64   `json:"gid"`                  // gid
	GoodsNums          sql.NullInt64   `json:"goods_nums"`           // goods_nums
	RefundAmount       sql.NullFloat64 `json:"refund_amount"`        // refund_amount
	OriginAmount       sql.NullFloat64 `json:"origin_amount"`        // origin_amount
	ActualRefundAmount sql.NullFloat64 `json:"actual_refund_amount"` // actual_refund_amount
	RefundExplain      sql.NullString  `json:"refund_explain"`       // refund_explain
	FactoryRemark      sql.NullString  `json:"factory_remark"`       // factory_remark
	RefundType         sql.NullInt64   `json:"refund_type"`          // refund_type
	RefundPayType      sql.NullInt64   `json:"refund_pay_type"`      // refund_pay_type
	EvidencePhoto      JSON            `json:"evidence_photo"`       // evidence_photo
	ExpressNo          sql.NullString  `json:"express_no"`           // express_no
	ExpressName        sql.NullString  `json:"express_name"`         // express_name
	ExpressID          sql.NullInt64   `json:"express_id"`           // express_id
	ExpressPhoto       sql.NullString  `json:"express_photo"`        // express_photo
	CauseID            sql.NullInt64   `json:"cause_id"`             // cause_id
	RefundCause        sql.NullString  `json:"refund_cause"`         // refund_cause
	Status             sql.NullInt64   `json:"status"`               // status
	FinishAt           mysql.NullTime  `json:"finish_at"`            // finish_at
	CreatedAt          mysql.NullTime  `json:"created_at"`           // created_at
	UpdatedAt          mysql.NullTime  `json:"updated_at"`           // updated_at
	ShippingReturn     JSON            `json:"shipping_return"`      // shipping_return
	OrderFrom          sql.NullInt64   `json:"order_from"`           // order_from
	ErpRefundNo        sql.NullString  `json:"erp_refund_no"`        // erp_refund_no
	CollectAt          mysql.NullTime  `json:"collect_at"`           // collect_at
	IsGyl              sql.NullInt64   `json:"is_gyl"`               // is_gyl

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the RefundOrder exists in the database.
func (ro *RefundOrder) Exists() bool { //refund_orders
	return ro._exists
}

// Deleted provides information if the RefundOrder has been deleted from the database.
func (ro *RefundOrder) Deleted() bool {
	return ro._deleted
}

// Get table name
func GetRefundOrderTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "refund_orders", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the RefundOrder to the database.
func (ro *RefundOrder) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if ro._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ro.RefundNo, ro.OrderID, ro.OrderNo, ro.FactoryID, ro.UID, ro.OrderGsrid, ro.Gid, ro.GoodsNums, ro.RefundAmount, ro.OriginAmount, ro.ActualRefundAmount, ro.RefundExplain, ro.FactoryRemark, ro.RefundType, ro.RefundPayType, ro.EvidencePhoto, ro.ExpressNo, ro.ExpressName, ro.ExpressID, ro.ExpressPhoto, ro.CauseID, ro.RefundCause, ro.Status, ro.FinishAt, ro.CreatedAt, ro.UpdatedAt, ro.ShippingReturn, ro.OrderFrom, ro.ErpRefundNo, ro.CollectAt, ro.IsGyl)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, ro.RefundNo, ro.OrderID, ro.OrderNo, ro.FactoryID, ro.UID, ro.OrderGsrid, ro.Gid, ro.GoodsNums, ro.RefundAmount, ro.OriginAmount, ro.ActualRefundAmount, ro.RefundExplain, ro.FactoryRemark, ro.RefundType, ro.RefundPayType, ro.EvidencePhoto, ro.ExpressNo, ro.ExpressName, ro.ExpressID, ro.ExpressPhoto, ro.CauseID, ro.RefundCause, ro.Status, ro.FinishAt, ro.CreatedAt, ro.UpdatedAt, ro.ShippingReturn, ro.OrderFrom, ro.ErpRefundNo, ro.CollectAt, ro.IsGyl)
	} else {
		res, err = dbConn.Exec(sqlstr, ro.RefundNo, ro.OrderID, ro.OrderNo, ro.FactoryID, ro.UID, ro.OrderGsrid, ro.Gid, ro.GoodsNums, ro.RefundAmount, ro.OriginAmount, ro.ActualRefundAmount, ro.RefundExplain, ro.FactoryRemark, ro.RefundType, ro.RefundPayType, ro.EvidencePhoto, ro.ExpressNo, ro.ExpressName, ro.ExpressID, ro.ExpressPhoto, ro.CauseID, ro.RefundCause, ro.Status, ro.FinishAt, ro.CreatedAt, ro.UpdatedAt, ro.ShippingReturn, ro.OrderFrom, ro.ErpRefundNo, ro.CollectAt, ro.IsGyl)
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
	ro.ID = int64(id)
	ro._exists = true

	return nil
}

// Update updates the RefundOrder in the database.
func (ro *RefundOrder) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ro._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`refund_no = ?, order_id = ?, order_no = ?, factory_id = ?, uid = ?, order_gsrid = ?, gid = ?, goods_nums = ?, refund_amount = ?, origin_amount = ?, actual_refund_amount = ?, refund_explain = ?, factory_remark = ?, refund_type = ?, refund_pay_type = ?, evidence_photo = ?, express_no = ?, express_name = ?, express_id = ?, express_photo = ?, cause_id = ?, refund_cause = ?, status = ?, finish_at = ?, created_at = ?, updated_at = ?, shipping_return = ?, order_from = ?, erp_refund_no = ?, collect_at = ?, is_gyl = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ro.RefundNo, ro.OrderID, ro.OrderNo, ro.FactoryID, ro.UID, ro.OrderGsrid, ro.Gid, ro.GoodsNums, ro.RefundAmount, ro.OriginAmount, ro.ActualRefundAmount, ro.RefundExplain, ro.FactoryRemark, ro.RefundType, ro.RefundPayType, ro.EvidencePhoto, ro.ExpressNo, ro.ExpressName, ro.ExpressID, ro.ExpressPhoto, ro.CauseID, ro.RefundCause, ro.Status, ro.FinishAt, ro.CreatedAt, ro.UpdatedAt, ro.ShippingReturn, ro.OrderFrom, ro.ErpRefundNo, ro.CollectAt, ro.IsGyl, ro.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ro.RefundNo, ro.OrderID, ro.OrderNo, ro.FactoryID, ro.UID, ro.OrderGsrid, ro.Gid, ro.GoodsNums, ro.RefundAmount, ro.OriginAmount, ro.ActualRefundAmount, ro.RefundExplain, ro.FactoryRemark, ro.RefundType, ro.RefundPayType, ro.EvidencePhoto, ro.ExpressNo, ro.ExpressName, ro.ExpressID, ro.ExpressPhoto, ro.CauseID, ro.RefundCause, ro.Status, ro.FinishAt, ro.CreatedAt, ro.UpdatedAt, ro.ShippingReturn, ro.OrderFrom, ro.ErpRefundNo, ro.CollectAt, ro.IsGyl, ro.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ro.RefundNo, ro.OrderID, ro.OrderNo, ro.FactoryID, ro.UID, ro.OrderGsrid, ro.Gid, ro.GoodsNums, ro.RefundAmount, ro.OriginAmount, ro.ActualRefundAmount, ro.RefundExplain, ro.FactoryRemark, ro.RefundType, ro.RefundPayType, ro.EvidencePhoto, ro.ExpressNo, ro.ExpressName, ro.ExpressID, ro.ExpressPhoto, ro.CauseID, ro.RefundCause, ro.Status, ro.FinishAt, ro.CreatedAt, ro.UpdatedAt, ro.ShippingReturn, ro.OrderFrom, ro.ErpRefundNo, ro.CollectAt, ro.IsGyl, ro.ID)
	}
	return err
}

// Save saves the RefundOrder to the database.
func (ro *RefundOrder) Save(ctx context.Context) error {
	if ro.Exists() {
		return ro.Update(ctx)
	}

	return ro.Insert(ctx)
}

// Delete deletes the RefundOrder from the database.
func (ro *RefundOrder) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if ro._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, ro.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, ro.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, ro.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	ro._deleted = true

	return nil
}

// RefundOrdersByFactoryID retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'factory_id'.
func RefundOrdersByFactoryID(ctx context.Context, factoryID sql.NullInt64, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE factory_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, factoryID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, factoryID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, factoryID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}

// RefundOrdersByOrderGsrid retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'order_gsrid'.
func RefundOrdersByOrderGsrid(ctx context.Context, orderGsrid sql.NullInt64, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE order_gsrid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, orderGsrid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, orderGsrid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, orderGsrid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}

// RefundOrdersByOrderID retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'order_id'.
func RefundOrdersByOrderID(ctx context.Context, orderID sql.NullInt64, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE order_id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, orderID)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, orderID)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, orderID)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}

// RefundOrdersByOrderNo retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'order_no'.
func RefundOrdersByOrderNo(ctx context.Context, orderNo sql.NullString, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE order_no = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, orderNo)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, orderNo)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, orderNo)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}

// RefundOrdersByRefundNo retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'refund_no'.
func RefundOrdersByRefundNo(ctx context.Context, refundNo sql.NullString, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE refund_no = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, refundNo)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, refundNo)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, refundNo)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}

// RefundOrderByID retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'refund_orders_id_pkey'.
func RefundOrderByID(ctx context.Context, id int64, key ...interface{}) (*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
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
	ro := RefundOrder{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}
	}

	return &ro, nil
}

// RefundOrdersByRefundType retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'refund_type'.
func RefundOrdersByRefundType(ctx context.Context, refundType sql.NullInt64, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE refund_type = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, refundType)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, refundType)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, refundType)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}

// RefundOrdersByStatus retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'status'.
func RefundOrdersByStatus(ctx context.Context, status sql.NullInt64, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, status)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, status)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}

// RefundOrdersByUID retrieves a row from 'aypcddg.refund_orders' as a RefundOrder.
//
// Generated from index 'uid'.
func RefundOrdersByUID(ctx context.Context, uid sql.NullInt64, key ...interface{}) ([]*RefundOrder, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetRefundOrderTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, refund_no, order_id, order_no, factory_id, uid, order_gsrid, gid, goods_nums, refund_amount, origin_amount, actual_refund_amount, refund_explain, factory_remark, refund_type, refund_pay_type, evidence_photo, express_no, express_name, express_id, express_photo, cause_id, refund_cause, status, finish_at, created_at, updated_at, shipping_return, order_from, erp_refund_no, collect_at, is_gyl ` +
		`FROM ` + tableName +
		` WHERE uid = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, uid)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	var queryData *sql.Rows
	if tx != nil {
		queryData, err = tx.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	} else {
		queryData, err = dbConn.Query(sqlstr, uid)
		if err != nil {
			return nil, err
		}
	}

	defer queryData.Close()

	// load results
	res := make([]*RefundOrder, 0)
	for queryData.Next() {
		ro := RefundOrder{
			_exists: true,
		}

		// scan
		err = queryData.Scan(&ro.ID, &ro.RefundNo, &ro.OrderID, &ro.OrderNo, &ro.FactoryID, &ro.UID, &ro.OrderGsrid, &ro.Gid, &ro.GoodsNums, &ro.RefundAmount, &ro.OriginAmount, &ro.ActualRefundAmount, &ro.RefundExplain, &ro.FactoryRemark, &ro.RefundType, &ro.RefundPayType, &ro.EvidencePhoto, &ro.ExpressNo, &ro.ExpressName, &ro.ExpressID, &ro.ExpressPhoto, &ro.CauseID, &ro.RefundCause, &ro.Status, &ro.FinishAt, &ro.CreatedAt, &ro.UpdatedAt, &ro.ShippingReturn, &ro.OrderFrom, &ro.ErpRefundNo, &ro.CollectAt, &ro.IsGyl)
		if err != nil {
			return nil, err
		}

		res = append(res, &ro)
	}

	return res, nil
}
