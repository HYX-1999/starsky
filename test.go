/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-15 17:40:00
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-15 17:42:10
 */

package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
