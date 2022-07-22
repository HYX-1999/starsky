/*
 * @Description: 路由
 * @Author: hyx
 * @Date: 2022-07-22 17:02:01
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-22 17:49:47
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
	}

	r.Run(":8080")
}
