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

// DpShopOwnerRelManager represents a row from 'aypcddg.dp_shop_owner_rel_manager'.
type DpShopOwnerRelManager struct {
	ID             int  `json:"id"`              // id
	ShopOwnerID    uint `json:"shop_owner_id"`   // shop_owner_id
	ShopManagerID  uint `json:"shop_manager_id"` // shop_manager_id
	RelationStatus int8 `json:"relation_status"` // relation_status

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DpShopOwnerRelManager exists in the database.
func (dsorm *DpShopOwnerRelManager) Exists() bool { //dp_shop_owner_rel_manager
	return dsorm._exists
}

// Deleted provides information if the DpShopOwnerRelManager has been deleted from the database.
func (dsorm *DpShopOwnerRelManager) Deleted() bool {
	return dsorm._deleted
}

// Get table name
func GetDpShopOwnerRelManagerTableName(key ...interface{}) (string, error) {
	tableName, err := components.M.GetTable(components.E.Opts.DBConfig.Name, "dp_shop_owner_rel_manager", key...)
	if err != nil {
		return "", err
	}
	return tableName, nil
}

// Insert inserts the DpShopOwnerRelManager to the database.
func (dsorm *DpShopOwnerRelManager) Insert(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB
	var res sql.Result
	// if already exist, bail
	if dsorm._exists {
		return errors.New("insert failed: already exists")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDpShopOwnerRelManagerTableName(key...)
	if err != nil {
		return err
	}

	// sql insert query, primary key provided by autoincrement
	sqlstr := `INSERT INTO ` + tableName +
		` (` +
		`shop_owner_id, shop_manager_id, relation_status` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dsorm.ShopOwnerID, dsorm.ShopManagerID, dsorm.RelationStatus)))
	if err != nil {
		return err
	}
	if tx != nil {
		res, err = tx.Exec(sqlstr, dsorm.ShopOwnerID, dsorm.ShopManagerID, dsorm.RelationStatus)
	} else {
		res, err = dbConn.Exec(sqlstr, dsorm.ShopOwnerID, dsorm.ShopManagerID, dsorm.RelationStatus)
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
	dsorm.ID = int(id)
	dsorm._exists = true

	return nil
}

// Update updates the DpShopOwnerRelManager in the database.
func (dsorm *DpShopOwnerRelManager) Update(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dsorm._deleted {
		return errors.New("update failed: marked for deletion")
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDpShopOwnerRelManagerTableName(key...)
	if err != nil {
		return err
	}

	// sql query
	sqlstr := `UPDATE ` + tableName + ` SET ` +
		`shop_owner_id = ?, shop_manager_id = ?, relation_status = ?` +
		` WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dsorm.ShopOwnerID, dsorm.ShopManagerID, dsorm.RelationStatus, dsorm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dsorm.ShopOwnerID, dsorm.ShopManagerID, dsorm.RelationStatus, dsorm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dsorm.ShopOwnerID, dsorm.ShopManagerID, dsorm.RelationStatus, dsorm.ID)
	}
	return err
}

// Save saves the DpShopOwnerRelManager to the database.
func (dsorm *DpShopOwnerRelManager) Save(ctx context.Context) error {
	if dsorm.Exists() {
		return dsorm.Update(ctx)
	}

	return dsorm.Insert(ctx)
}

// Delete deletes the DpShopOwnerRelManager from the database.
func (dsorm *DpShopOwnerRelManager) Delete(ctx context.Context, key ...interface{}) error {
	var err error
	var dbConn *sql.DB

	// if deleted, bail
	if dsorm._deleted {
		return nil
	}

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetMasterConn()
		if err != nil {
			return err
		}
	}

	tableName, err := GetDpShopOwnerRelManagerTableName(key...)
	if err != nil {
		return err
	}
	//1

	// sql query with composite primary key
	sqlstr := `UPDATE ` + tableName + ` SET is_del = 1 WHERE id = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, dsorm.ID)))
	if tx != nil {
		_, err = tx.Exec(sqlstr, dsorm.ID)
	} else {
		_, err = dbConn.Exec(sqlstr, dsorm.ID)
	}

	if err != nil {
		return err
	}

	// set deleted
	dsorm._deleted = true

	return nil
}

// DpShopOwnerRelManagerByID retrieves a row from 'aypcddg.dp_shop_owner_rel_manager' as a DpShopOwnerRelManager.
//
// Generated from index 'dp_shop_owner_rel_manager_id_pkey'.
func DpShopOwnerRelManagerByID(ctx context.Context, id int, key ...interface{}) (*DpShopOwnerRelManager, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDpShopOwnerRelManagerTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, shop_owner_id, shop_manager_id, relation_status ` +
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
	dsorm := DpShopOwnerRelManager{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, id).Scan(&dsorm.ID, &dsorm.ShopOwnerID, &dsorm.ShopManagerID, &dsorm.RelationStatus)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, id).Scan(&dsorm.ID, &dsorm.ShopOwnerID, &dsorm.ShopManagerID, &dsorm.RelationStatus)
		if err != nil {
			return nil, err
		}
	}

	return &dsorm, nil
}

// DpShopOwnerRelManagerByShopManagerIDRelationStatus retrieves a row from 'aypcddg.dp_shop_owner_rel_manager' as a DpShopOwnerRelManager.
//
// Generated from index 'shop_manager_id'.
func DpShopOwnerRelManagerByShopManagerIDRelationStatus(ctx context.Context, shopManagerID uint, relationStatus int8, key ...interface{}) (*DpShopOwnerRelManager, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDpShopOwnerRelManagerTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, shop_owner_id, shop_manager_id, relation_status ` +
		`FROM ` + tableName +
		` WHERE shop_manager_id = ? AND relation_status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, shopManagerID, relationStatus)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	dsorm := DpShopOwnerRelManager{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, shopManagerID, relationStatus).Scan(&dsorm.ID, &dsorm.ShopOwnerID, &dsorm.ShopManagerID, &dsorm.RelationStatus)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, shopManagerID, relationStatus).Scan(&dsorm.ID, &dsorm.ShopOwnerID, &dsorm.ShopManagerID, &dsorm.RelationStatus)
		if err != nil {
			return nil, err
		}
	}

	return &dsorm, nil
}

// DpShopOwnerRelManagerByShopManagerIDRelationStatus retrieves a row from 'aypcddg.dp_shop_owner_rel_manager' as a DpShopOwnerRelManager.
//
// Generated from index 'shop_manager_id_2'.
func DpShopOwnerRelManagerByShopManagerIDRelationStatus(ctx context.Context, shopManagerID uint, relationStatus int8, key ...interface{}) (*DpShopOwnerRelManager, error) {
	var err error
	var dbConn *sql.DB

	tableName, err := GetDpShopOwnerRelManagerTableName(key...)
	if err != nil {
		return nil, err
	}

	// sql query
	sqlstr := `SELECT ` +
		`id, shop_owner_id, shop_manager_id, relation_status ` +
		`FROM ` + tableName +
		` WHERE shop_manager_id = ? AND relation_status = ?`

	// run query
	utils.GetTraceLog(ctx).Debug("DB", zap.String("SQL", fmt.Sprint(sqlstr, shopManagerID, relationStatus)))

	tx, err := components.M.GetConnFromCtx(ctx)
	if err != nil {
		dbConn, err = components.M.GetSlaveConn()
		if err != nil {
			return nil, err
		}
	}
	dsorm := DpShopOwnerRelManager{
		_exists: true,
	}

	if tx != nil {
		err = tx.QueryRow(sqlstr, shopManagerID, relationStatus).Scan(&dsorm.ID, &dsorm.ShopOwnerID, &dsorm.ShopManagerID, &dsorm.RelationStatus)
		if err != nil {
			return nil, err
		}
	} else {
		err = dbConn.QueryRow(sqlstr, shopManagerID, relationStatus).Scan(&dsorm.ID, &dsorm.ShopOwnerID, &dsorm.ShopManagerID, &dsorm.RelationStatus)
		if err != nil {
			return nil, err
		}
	}

	return &dsorm, nil
}
