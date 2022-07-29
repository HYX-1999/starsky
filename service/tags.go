/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-29 16:44:38
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-29 16:49:20
 */

package service

import (
	"fmt"

	core "github.com/HYX-1999/starsky/db"
	"github.com/HYX-1999/starsky/models"
)

/**
 * @description: CreateTagsService
 * @param {string} name
 * @return {*}
 */
func CreateTagsService(name string) (tags *models.Tags, ok bool, res bool) {
	db := core.GetDB()

	err := db.Where("name = ?", name).First(&tags).Error
	if err != nil {
		return tags, false, false
	}
	tags.Name = name

	tx := db.Create(&tags)
	fmt.Println(tx.RowsAffected)
	if tx.RowsAffected > 0 {
		return tags, true, true
	}

	return tags, false, false
}
