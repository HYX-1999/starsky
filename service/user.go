/*
 * @Description: user服务处理业务逻辑
 * @Author: hyx
 * @Date: 2022-07-22 17:10:41
 * @LastEditors: hyx
 * @LastEditTime: 2022-07-26 11:30:28
 */

package service

import (
	"fmt"

	core "github.com/HYX-1999/starsky/db"
	"github.com/HYX-1999/starsky/models"
)

/**
 * @description: GetUsersService
 * @return {*}
 */
func GetUsersService() (userList []*models.User) {
	core.GetDB().Find(&userList)
	return userList
}

/**
 * @description: CreateUserService
 * @param {*} username
 * @param {*} password
 * @param {*} phone
 * @param {string} email
 * @param {int} state
 * @return {*}
 */
func CreateUserService(username, password, phone, email string, state int) bool {
	db := core.GetDB()

	user := models.User{
		Username: username,
		Password: password,
		Phone:    phone,
		Email:    email,
		State:    state,
	}

	tx := db.Create(&user)
	return tx.RowsAffected > 0
}

/**
 * @description: UpdateUserByIdService
 * @param {int} id
 * @param {*} username
 * @param {*} password
 * @param {*} phone
 * @param {string} email
 * @return {*}
 */
func UpdateUserByIdService(id int, username, password, phone, email string) (user models.User, ok bool, err error) {
	err = core.GetDB().First(&user, id).Error
	if err != nil {
		return user, false, err
	}

	user.Username = username
	user.Password = password
	user.Phone = phone
	user.Email = email

	err = core.GetDB().Save(&user).Error

	if err != nil {
		return user, false, err
	}

	return user, true, err
}

/**
 * @description: DeleteUserByIdService
 * @param {int} id
 * @return {*}
 */
func DeleteUserByIdService(id int) (err error, ok bool) {
	err = core.GetDB().First(&models.User{}, id).Error
	if err != nil {
		return err, false
	}

	err = core.GetDB().Delete(&models.User{}, id).Error
	return err, true
}

/**
 * @description: GetUserByIdService
 * @param {int} id
 * @return {*}
 */
func GetUserByIdService(id int) (user []*models.User) {
	core.GetDB().First(&user, id)
	return
}

/**
 * @description: ChangeUserStateService
 * @param {int} id
 * @param {int} state
 * @return {*}
 */
func ChangeUserStateService(id int, state int) (user models.User, err error) {
	err = core.GetDB().First(&user, id).Error
	fmt.Println(user)
	if err != nil {
		return user, err
	}

	core.GetDB().Model(&user).Update("state", state)

	return user, nil
}
