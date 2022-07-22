/*
 * @Description: 统一APT返回格式
 * @Author: hyx
 * @Date: 2022-07-22 16:34:29
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-22 16:41:20
 */

package utils

import (
	"net/http"
	"time"

	"github.com/HYX-1999/starsky/types"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Time time.Time   `json: "time"`
	Code int         `json: "code"`
	Msg  string      `json: "msg"`
	Data interface{} `json: "data"`
}

/**
 * @description: 成功
 * @param {*gin.Context} c
 * @param {interface{}} data
 * @return {*}
 */
func Success(c *gin.Context, data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := Result{}
	res.Time = time.Now()
	res.Code = int(types.ApiCode.SUCCESS)
	res.Msg = types.ApiCode.GetMessage(types.ApiCode.SUCCESS)
	res.Data = data

	c.JSON(http.StatusOK, res)
}

/**
 * @description: 出错
 * @param {*gin.Context} c
 * @param {int} code
 * @param {string} msg
 * @return {*}
 */
func Error(c *gin.Context, code int, msg string) {
	res := Result{}
	res.Time = time.Now()
	res.Code = code
	res.Msg = msg
	res.Data = gin.H{}

	c.JSON(http.StatusOK, res)
}
