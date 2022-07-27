/*
 * @Description: 路由
 * @Author: hyx
 * @Date: 2022-07-22 17:02:01
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-27 11:49:48
 */

package routers

import (
	"github.com/HYX-1999/starsky/handler"
	"github.com/HYX-1999/starsky/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	// 管理后台
	// 管理员登陆
	r.POST("/api/login", handler.AuthLogin)

	userGroup := r.Group("user/v1")
	userGroup.Use(middleware.AuthMiddleware())
	{
		userGroup.GET("/get-users", handler.GetUsers)
		userGroup.POST("/create-user", handler.CreateUser)
		userGroup.POST("/update-user-by-id", handler.UpdateUserById)
		userGroup.POST("/del-user-by-id", handler.DeleteUserById)
		userGroup.GET("/get-user-by-id", handler.GetUserById)
		userGroup.GET("/change-user-state/:id/:state", handler.ChangeUserState)
	}

	r.Run(":8080")
}
