/*
 * @Description: 路由
 * @Author: hyx
 * @Date: 2022-07-22 17:02:01
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-26 11:39:03
 */

package routers

import (
	"github.com/HYX-1999/starsky/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	userGroup := r.Group("user/v1")
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
