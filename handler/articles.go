/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-29 17:14:26
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-29 17:21:55
 */

package handler

import (
	"strconv"

	"github.com/HYX-1999/starsky/service"
	"github.com/HYX-1999/starsky/types"
	"github.com/HYX-1999/starsky/utils"
	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	title, ok := c.GetPostForm("title")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	desc := c.PostForm("desc")
	content, ok := c.GetPostForm("content")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	author, ok := c.GetPostForm("author")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	tags := c.PostForm("tags")
	category := c.PostForm("category")

	newCategory, err := strconv.Atoi(category)
	if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	ok = service.CreateArticleService(title, desc, content, author, tags, newCategory)
	if !ok {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}
