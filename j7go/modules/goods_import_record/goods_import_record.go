package goods_import_record

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joselee214/j7f/components/http/server"
	"j7go/models/ddg"
	"os"
	"path"
	_ "time"
)

func Init(s *gin.Engine) {
	serve := &ServerController{}
	s.POST("goodsFactory/import", serve.Imports)
}

type ServerController struct {
	server.Controller
}

var (
	uploadFileKey = "file"
)

func (s *ServerController) Imports(ctx *gin.Context) {
	files, err := ctx.FormFile(uploadFileKey)
	if err != nil {
		s.ResponseError(ctx, "上传文件不能为空")
		return
	}
	guid := uuid.New().String()
	dst := path.Join("/opt/wwwroot/golang/j7go/log", guid+".xls")
	fmt.Println(dst)
	// gin 简单做了封装,拷贝了文件流
	if err := ctx.SaveUploadedFile(files, dst); err != nil {
		s.ResponseError(ctx, "文件保存失败")
		return
	}
	lastMap, err := ddg.SelectLastRow(ctx)
	batchNo := 1
	if err == nil {
		batchNo = lastMap.BatchNo
		batchNo += 1
	}

	xlsx, err := excelize.OpenFile(dst)
	if err != nil {
		fmt.Println(err.Error())
		s.ResponseError(ctx, "打开失败")
		return
	}
	rows := xlsx.GetRows("Sheet" + "1")
	for key, row := range rows {
		if key > 0 {

			var goodsImport = &ddg.GoodsDataImportRecord{
				GoodsName:     row[0],
				BarCode:       row[1],
				BrandName:     row[2],
				GoodsFormat:   row[3],
				SupplierName:  row[4],
				SupplierPrice: row[5],
				GoodsNo:       row[6],
				Color:         row[7],
				Size:          row[8],
				ShippingPrice: row[9],
				LowSellPrice:  row[10],
				BatchNo:       batchNo,
			}
			err = goodsImport.Insert(ctx)
			if err != nil {
				s.ResponseError(ctx, err.Error())
				return
			}
		}
	}
	_ = os.Remove(dst)
	s.ResponseSuccess(ctx)
}
