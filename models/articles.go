/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-29 16:59:19
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-29 17:08:02
 */

package models

import (
	core "github.com/HYX-1999/starsky/db"
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	Title    string `gorm:"column:title;NOT NULL"`
	Desc     string `gorm:"column:desc"`
	Content  string `gorm:"column:content;NOT NULL"`
	Author   string `gorm:"column:author"`
	Tags     string `gorm:"column:tags"`
	Category int    `gorm:"column:category"`
	Love     int    `gorm:"column:love"`
	Watch    int    `gorm:"column:watch"`
}

func (a *Articles) Articles() string {
	return "articles"
}

func init() {
	core.GetDB().AutoMigrate(&Articles{})
}
