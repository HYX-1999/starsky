/*
 * @Description: user控制器
 * @Author: hyx
 * @Date: 2022-07-22 17:09:33
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-26 11:36:34
 */

package handler

import (
	"fmt"
	"strconv"

	"github.com/HYX-1999/starsky/service"
	"github.com/HYX-1999/starsky/types"
	"github.com/HYX-1999/starsky/utils"
	"github.com/gin-gonic/gin"
)

/**
 * @description: GetUsers
 * @param {*gin.Context} c
 * @return {*}
 */
func GetUsers(c *gin.Context) {
	userList := service.GetUsersService()
	fmt.Println(userList)

	utils.Success(c, userList)
}

/**
 * @description: CreateUser
 * @param {*gin.Context} c
 * @return {*}
 */
func CreateUser(c *gin.Context) {
	username, ok := c.GetPostForm("username")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	newPassword := utils.EncryMd5(password)

	phone := c.PostForm("phone")
	email := c.PostForm("email")
	state, _ := strconv.Atoi(c.PostForm("state"))

	ok = service.CreateUserService(username, newPassword, phone, email, state)
	if !ok {
		utils.Error(c, int(types.ApiCode.CREATEUSERFAILED), types.ApiCode.GetMessage(types.ApiCode.CREATEUSERFAILED))
		return
	}
	utils.Success(c, nil)
}

/**
 * @description: UpdateUserById
 * @param {*gin.Context} c
 * @return {*}
 */
func UpdateUserById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	username, ok := c.GetPostForm("username")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	password, ok := c.GetPostForm("password")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newPassword := utils.EncryMd5(password)
	phone := c.PostForm("phone")
	email := c.PostForm("email")

	_, ok, err = service.UpdateUserByIdService(newId, username, newPassword, phone, email)
	if !ok {
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

/**
 * @description: DeleteUserById
 * @param {*gin.Context} c
 * @return {*}
 */
func DeleteUserById(c *gin.Context) {
	id, ok := c.GetPostForm("id")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}
	fmt.Println("-------" + id + "-------")

	newId, err := strconv.Atoi(id)
	fmt.Printf("%T", newId)
	if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	err, ok = service.DeleteUserByIdService(newId)
	if !ok {
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	if err != nil {
		utils.Error(c, int(types.ApiCode.FAILED), types.ApiCode.GetMessage(types.ApiCode.FAILED))
		return
	}

	utils.Success(c, nil)
}

/**
 * @description: GetUserById
 * @param {*gin.Context} c
 * @return {*}
 */
func GetUserById(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		utils.Error(c, int(types.ApiCode.LCAKPARAMETERS), types.ApiCode.GetMessage(types.ApiCode.LCAKPARAMETERS))
		return
	}

	newId, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	user := service.GetUserByIdService(newId)

	utils.Success(c, user)
}

func ChangeUserState(c *gin.Context) {
	id := c.Param("id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		utils.Error(c, int(types.ApiCode.CONVERTFAILED), types.ApiCode.GetMessage(types.ApiCode.CONVERTFAILED))
		return
	}

	state := c.Param("state")
	newState, _ := strconv.Atoi(state)
	_, err = service.ChangeUserStateService(newId, newState)
	if err != nil {
		utils.Error(c, int(types.ApiCode.NOSUCHID), types.ApiCode.GetMessage(types.ApiCode.NOSUCHID))
		return
	}

	utils.Success(c, nil)
}
