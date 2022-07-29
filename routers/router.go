/*
 * @Description: 路由
 * @Author: hyx
 * @Date: 2022-07-22 17:02:01
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-29 17:22:09
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

	tagsGroup := r.Group("/tags/v1")
	tagsGroup.Use(middleware.AuthMiddleware())
	{
		tagsGroup.POST("/create-tags", handler.CreateTags)
	}

	articlesGroup := r.Group("/articles/v1")
	articlesGroup.Use(middleware.AuthMiddleware())
	{
		articlesGroup.POST("/create-article", handler.CreateArticle)
	}

	r.Run(":8080")
}
