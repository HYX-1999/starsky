/*
 * @Description: user服务处理业务逻辑
 * @Author: hyx
 * @Date: 2022-07-22 17:10:41
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-22 17:13:13
 */

package service

import (
	core "github.com/HYX-1999/starsky/db"
	"github.com/HYX-1999/starsky/models"
)

/**
 * @description: GetUsersService
 * @return {*}
 */
func GetUsersService() (userList []*models.User) {
	core.GetDB().Find(&userList)
	return userList
}
