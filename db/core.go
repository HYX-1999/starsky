/*
 * @Description: 连接mysql
 * @Author: hyx
 * @Date: 2022-07-22 16:48:16
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-22 17:27:21
 */

package core

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// 定义db全局变量
var Db *gorm.DB

func init() {
	var err error
	dsn := "root:admin5698@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名加s
		},
		Logger:                                   logger.Default.LogMode(logger.Info), // 打印sql语句
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("Connecting database failed: " + err.Error())
	}
}

/**
 * @description: GetDB
 * @return {*}
 */
func GetDB() *gorm.DB {
	return Db
}
