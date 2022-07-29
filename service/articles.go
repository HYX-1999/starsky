/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-29 17:09:07
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-29 17:13:35
 */

package service

import (
	core "github.com/HYX-1999/starsky/db"
	"github.com/HYX-1999/starsky/models"
)

/**
 * @description: CreateArticleService
 * @param {*} title
 * @param {*} desc
 * @param {*} content
 * @param {*} author
 * @param {string} tags
 * @param {int} category
 * @return {*}
 */
func CreateArticleService(title, desc, content, author, tags string, category int) (ok bool) {
	db := core.GetDB()

	articles := models.Articles{
		Title:    title,
		Desc:     desc,
		Content:  content,
		Author:   author,
		Tags:     tags,
		Category: category,
	}

	tx := db.Create(&articles)
	return tx.RowsAffected > 0
}
