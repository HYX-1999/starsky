/*
 * @Description: user账号模型
 * @Author: hyx
 * @Date: 2022-07-22 17:04:32
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-22 17:08:52
 */

package models

import (
	core "github.com/HYX-1999/starsky/db"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"column:username;NOT NULL"`
	Password string `gorm:"column:password;NOT NULL"`
	Phone    string `gorm:"column:phone"`
	Email    string `gorm:"column:email"`
	State    int    `gorm:"column:state;default:1;NOT NULL"`
}

func (u *User) User() string {
	return "user"
}

func init() {
	core.Db.AutoMigrate(&User{})
}
