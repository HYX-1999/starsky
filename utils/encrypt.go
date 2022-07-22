/*
 * @Description: md5加密
 * @Author: hyx
 * @Date: 2022-07-22 16:41:53
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-22 16:42:58
 */

package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/**
 * @description: EncryMd5
 * @param {string} s
 * @return {*}
 */
func EncryMd5(s string) string {
	ctx := md5.New()
	ctx.Write([]byte(s))
	return hex.EncodeToString(ctx.Sum(nil))
}
