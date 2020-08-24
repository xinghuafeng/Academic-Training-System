package router

import (
	"Academic-Training-System/controller"
	"Academic-Training-System/controller/appController"
	"Academic-Training-System/controller/functionController"
	"Academic-Training-System/controller/orgController"
	"Academic-Training-System/controller/roleController"
	"Academic-Training-System/controller/roleFunctionController"
	"Academic-Training-System/controller/userController"
	"Academic-Training-System/controller/userRoleController"

	"github.com/gin-gonic/gin"
)

/*
	访问路由
*/
func Router(engine *gin.Engine) {

	v1 := engine.Group("Academic/api/v1") // 路由分组
	{
		//系统设置-校区管理
		v1.GET("baseinfo/getbaseinfo", controller.Baseinfo)
		v1.POST("baseinfo/delete", controller.DeleteBaseinfo)
		v1.POST("baseinfo/Insertorupdatebaseinfo", controller.UpdateOrInsertBaseinfo)
		//系统设置-销售来源

		//系统设置-时间段管理

		// ping
		v1.GET("/ping", controller.Ping)

		// 应用
		v1.GET("/app/all", appController.All)
		v1.GET("/app/find", appController.Find)
		v1.POST("app/save", appController.Save)
		v1.POST("/app/delete", appController.Delete)
		v1.GET("/app/options", appController.Options)

		// 用户
		v1.POST("/user/save", userController.Save)
		v1.POST("/user/login", userController.Login)
		v1.POST("/user/logout", userController.Logout)
		v1.POST("/user/delete", userController.Delete)
		v1.GET("/user/all", userController.All)
		v1.GET("/user/find", userController.Find)
		v1.GET("/user/api/own", userController.GetOwnInfo) // 写回 cookie

		// 角色
		v1.GET("/role/all", roleController.All)
		v1.GET("/role/find", roleController.Find)
		v1.POST("/role/save", roleController.Save)
		v1.POST("/role/delete", roleController.Delete)

		// 功能
		v1.GET("/function/all", functionController.All)
		v1.GET("/function/find", functionController.Find)

		// 角色-功能
		v1.POST("/roleFunction/save", roleFunctionController.Save)
		v1.GET("/roleFunction/find", roleFunctionController.Find)

		// 机构
		v1.GET("/org/find", orgController.Find)
		v1.POST("/org/save", orgController.Save)
		v1.POST("/org/delete", orgController.Delete)

		// 用户-角色
		v1.POST("/userRole/save", userRoleController.Save)
		v1.GET("/userRole/find", userRoleController.Find)

	}

}
