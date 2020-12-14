package bzj

import (
	"fmt"
	"github.com/apptut/rsp"
	"github.com/apptut/validator"
	"github.com/gin-gonic/gin"
	"github.com/joselee214/j7f/components/http/server"
	"j7go/models/ddg"
	"j7go/utils"
)

func Init(s *gin.Engine) {
	b := &BController{}
	s.POST("/bzj/query", b.Query)              // 查询保证金
	s.POST("/bzj/create", b.Create)            // 初始化保证金
	s.POST("/bzj/update_price", b.UpdatePrice) // 更新保证金
}

type BController struct {
	server.Controller
}

func (s *BController) UpdatePrice(ctx *gin.Context) {
	err := ValidUpdatePriceParams(ctx)
	if err != nil {
		s.ResponseError(ctx, err.Error())
		return
	}
	var query ddg.QueryReq
	ctx.Bind(&query)

	oldData, err := ddg.DdgBzjByFactoryID(ctx, query.FactoryId)

	if err != nil {
		s.ResponseError(ctx, err.Error())
		return
	}

	userInfo, _err := ddg.DdgAdminUserByID(ctx, query.Uid)
	if _err != nil {
		s.ResponseError(ctx, _err.Error())
		return
	}

	oldPrice := oldData.Price
	TotalPrice := oldData.TotalPrice
	opType := query.Type
	price := query.Price
	now := utils.NowTimestamp()

	switch opType {
	case 1, 2:
		oldPrice = oldPrice + price
		TotalPrice = TotalPrice + price

	case 3, 4:

		if price > oldPrice {
			s.ResponseError(ctx, "余额不足")
			return
		}

		oldPrice = oldPrice - price
		//TotalPrice = TotalPrice - _price
	}

	var ddgbzj = &ddg.DdgBzj{
		FactoryID:  query.FactoryId,
		Price:      oldPrice,
		TotalPrice: TotalPrice,
		Updated:    now,
	}
	err = ddgbzj.Update(ctx)
	if err != nil {
		s.ResponseError(ctx, err.Error())
		return
	}

	var bzjlog = &ddg.DdgBzjLog{
		Type:         opType,
		OperatorID:   query.Uid,
		OperatorName: userInfo.Name,
		RelatedID:    query.RelatedId,
		Price:        price,
		Action:       2,
		Created:      now,
		Updated:      now,
		BzjID:        query.FactoryId,
	}
	_ = bzjlog.Insert(ctx)
	s.Data = ddgbzj
	s.ResponseSuccess(ctx)
}

func (s *BController) Create(ctx *gin.Context) {

	now := utils.NowTimestamp()
	err := ValidCreateParams(ctx)
	if err != nil {
		//rsp.JsonErr(ctx, err.Error())
		s.ResponseError(ctx, err.Error())
		return
	}
	var query ddg.QueryReq
	err = ctx.Bind(&query)
	var ddgbzj ddg.DdgBzj
	ddgbzj.FactoryID = query.FactoryId
	ddgbzj.Price = query.Price
	ddgbzj.Created = now
	ddgbzj.Updated = now
	ddgbzj.TotalPrice = ddgbzj.Price
	if err != nil {
		s.ResponseError(ctx, err.Error())
		return
	}
	err = ddgbzj.Insert(ctx)
	if err != nil {
		s.ResponseError(ctx, err.Error())
		return
	}
	s.Data = ddgbzj
	s.ResponseSuccess(ctx)
}

func ValidUpdatePriceParams(ctx *gin.Context) error {

	data := map[string][]string{
		"price":      {ctx.PostForm("price")},
		"factory_id": {ctx.PostForm("factory_id")},
		"uid":        {ctx.PostForm("uid")},
		"related_id": {ctx.PostForm("related_id")},
		"type":       {ctx.PostForm("type")},
	}

	rules := map[string]string{
		"price":      "min:0|numeric",
		"factory_id": "int",
		"uid":        "int",
		"related_id": "nullable|int",
		"type":       "int",
	}

	msg := map[string]string{
		"price.numeric": "price参数格式不正确",
		"factory_id":    "缺少参数FactoryId",
		"uid":           "缺少参数Uid",
		"type":          "缺少参数Type",
	}

	_, err := validator.New(data, rules, msg)
	if err != nil {
		return rsp.NewErr(err)
	}
	return nil

}

func ValidCreateParams(ctx *gin.Context) error {

	data := map[string][]string{
		"factory_id": {ctx.PostForm("factory_id")},
		"price":      {ctx.PostForm("price")},
		"uid":        {ctx.PostForm("uid")},
	}

	rules := map[string]string{
		"factory_id": "int",
		"price":      "numeric",
		"uid":        "int",
	}
	msg := map[string]string{
		"factory_id": "缺少参数FactoryId",
		"price":      "缺少参数Price",
		"uid":        "缺少参数Uid",
	}
	_, err := validator.New(data, rules, msg)

	if err != nil {
		return err
	}
	return nil
}

func (s *BController) Query(ctx *gin.Context) {
	//s.Data = fmt.Sprintf("num fo chan : %d",1)
	//s.ResponseSuccess(ctx)
	var lp ddg.QueryReq
	ctx.Bind(&lp)
	fmt.Println(lp)

	data, err := ddg.QueryByCondition(ctx, &lp)
	if err != nil {
		s.ResponseError(ctx, err)
		return
	}
	s.Data = data
	s.ResponseSuccess(ctx)
}
