/*
 * @Description: user控制器
 * @Author: hyx
 * @Date: 2022-07-22 17:09:33
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-22 17:15:36
 */

package handler

import (
	"fmt"

	"github.com/HYX-1999/starsky/service"
	"github.com/HYX-1999/starsky/utils"
	"github.com/gin-gonic/gin"
)

/**
 * @description: GetUsers
 * @param {*gin.Context} c
 * @return {*}
 */
func GetUsers(c *gin.Context) {
	userList := service.GetUsersService()
	fmt.Println(userList)

	utils.Success(c, userList)
}
