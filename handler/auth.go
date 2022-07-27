/*
 * @Description:
 * @Author: hyx
 * @Date: 2022-07-27 10:41:06
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-27 11:48:32
 */

package handler

import (
	"errors"
	"fmt"
	"time"

	core "github.com/HYX-1999/starsky/db"
	"github.com/HYX-1999/starsky/models"
	"github.com/HYX-1999/starsky/types"
	"github.com/HYX-1999/starsky/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type JWTClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

const TokenExpireDuration = time.Hour * 2

var Secret = []byte("gin-api-blog")

/**
 * @description: GenToken
 * @param {string} username
 * @return {*}
 */
func GenToken(username string) (string, error) {
	c := JWTClaims{
		"username", // 自定义字段
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(), // 过期时间
			Issuer:    "heyuxin",                                  // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodES256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(Secret)
}

/**
 * @description: ParseToken 解析JWT
 * @param {string} tokenString
 * @return {*}
 */
func ParseToken(tokenString string) (*JWTClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(t *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func AuthLogin(c *gin.Context) {
	// 用户发送用户名和密码过来
	username := c.PostForm("username")
	password := utils.EncryMd5(c.PostForm("password"))

	db := core.GetDB()

	type UserInfo struct {
		Username string
		Password string
	}
	userInfo := make([]UserInfo, 2)
	db.First(&models.User{}).Where("username = ? AND password = ?", username, password).Scan(&userInfo)
	fmt.Println(userInfo)
	// 校验用户名和密码是否正确
	if username == userInfo[0].Username && password == userInfo[0].Password {
		tokenString, _ := GenToken(username)
		utils.Success(c, tokenString)
		return
	}
	utils.Error(c, int(types.ApiCode.AUTHFAILED), types.ApiCode.GetMessage(types.ApiCode.AUTHFAILED))
	return
}
