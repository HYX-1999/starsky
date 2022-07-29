/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-29 16:23:36
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-29 16:43:18
 */

package models

import (
	core "github.com/HYX-1999/starsky/db"
	"gorm.io/gorm"
)

type Tags struct {
	gorm.Model
	ID   uint   `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name string `gorm:"column:name;NOT NULL"`
}

func (t *Tags) Tags() string {
	return "tags"
}

func init() {
	core.GetDB().AutoMigrate(&Tags{})
}
