package controller

import (
	"Academic-Training-System/serializer"
	"Academic-Training-System/service/systemService"
	"github.com/gin-gonic/gin"
)

func Baseinfo(ctx *gin.Context) {

	res, err := systemService.GetBaseinfo(ctx)
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
func DeleteBaseinfo(ctx *gin.Context) {
	_, err := systemService.DeleteBaseinfo(ctx)
	if err != nil {
		ctx.JSON(200, serializer.Response{
			Code:  0,
			Msg:   "校区删除失败",
			Error: err.Error(),
		})
	} else {
		ctx.JSON(200, serializer.Response{
			Code:  1,
			Msg:   "校区删除成功",
			Error: "",
		})
	}

}
