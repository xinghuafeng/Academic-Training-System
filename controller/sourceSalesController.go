package controller

import (
	"Academic-Training-System/serializer"
	"Academic-Training-System/service/systemService"
	"github.com/gin-gonic/gin"
)

func GetSourceSales(ctx *gin.Context) {

	res, err := systemService.GetSourceSales(ctx)
	if err == nil {
		ctx.JSON(200, serializer.Response{
			Code:  1,
			Msg:   "获取列表成功",
			Data:  res,
			Error: "",
		})
	} else {
		ctx.JSON(200, serializer.Response{
			Code:  0,
			Msg:   "获取列表失败",
			Data:  res,
			Error: err.Error(),
		})
	}

}
func DeleteSourceSales(ctx *gin.Context) {
	_, err := systemService.DeleteSourceSales(ctx)
	if err != nil {
		ctx.JSON(200, serializer.Response{
			Code:  0,
			Msg:   "销售来源删除失败",
			Error: err.Error(),
		})
	} else {
		ctx.JSON(200, serializer.Response{
			Code:  1,
			Msg:   "销售来源删除成功",
			Error: "",
		})
	}

}
func UpdateOrInsertSourceSales(ctx *gin.Context) {
	_, err := systemService.UpdateOrInsertSourceSales(ctx)
	if err == nil {
		ctx.JSON(200, serializer.Response{
			Code:  1,
			Msg:   "校区更新成功",
			Error: "",
		})
	} else {
		ctx.JSON(200, serializer.Response{
			Code:  0,
			Msg:   "校区更新失败",
			Error: err.Error(),
		})
	}
}
