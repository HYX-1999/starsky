/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-27 10:29:11
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-27 11:48:59
 */

package middleware

import (
	"strings"

	"github.com/HYX-1999/starsky/handler"
	"github.com/HYX-1999/starsky/types"
	"github.com/HYX-1999/starsky/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			utils.Error(c, int(types.ApiCode.NOAUTH), types.ApiCode.GetMessage(types.ApiCode.NOAUTH))
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			utils.Error(c, int(types.ApiCode.AUTHFORMATERROR), types.ApiCode.GetMessage(types.ApiCode.AUTHFORMATERROR))
			c.Abort()
			return
		}

		mc, err := handler.ParseToken(parts[1])
		if err != nil {
			utils.Error(c, int(types.ApiCode.INVALIDTOKEN), types.ApiCode.GetMessage(types.ApiCode.INVALIDTOKEN))
			c.Abort()
			return
		}
		c.Set("username", mc.Username)
		c.Next()
	}
}
