/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-29 16:49:51
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-29 16:54:52
 */

package handler

import (
	"github.com/HYX-1999/starsky/service"
	"github.com/HYX-1999/starsky/types"
	"github.com/HYX-1999/starsky/utils"
	"github.com/gin-gonic/gin"
)

/**
 * @description: CreateTags
 * @param {*gin.Context} c
 * @return {*}
 */
func CreateTags(c *gin.Context) {
	name, ok := c.GetPostForm("name")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	_, ok, res := service.CreateTagsService(name)
	if !ok {
		utils.Error(c, int(types.ApiCode.EXISTSNAME), types.ApiCode.GetMessage(types.ApiCode.EXISTSNAME))
		return
	}

	if !res {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}
