package logics

import (
	"fmt"
	"shop_server/internal/models"
	"shop_server/pkg/mysqldb"
	"shop_server/utils"

	"github.com/golang-module/carbon"
)

// 修改用户密码逻辑
func UpdateUserPassword(uid int64, oldPassword, newPassword string) (bool, error) {
	user := models.User{}
	query := mysqldb.Mysql.Where("uid = ?", uid).First(&user)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("用户不存在")
	}
	if user.Password != utils.MD5String(oldPassword) {
		return false, fmt.Errorf("旧密码错误")
	}
	user.Password = utils.MD5String(newPassword)
	err := mysqldb.Mysql.Save(&user).Error
	return err == nil, err
}

// 获取用户列表逻辑
func GetUserList(pageNum, pageSize int) ([]models.User, error) {
	var userList []models.User
	query := mysqldb.Mysql.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&userList)
	for i := range userList {
		userList[i].Password = ""
	}
	return userList, query.Error
}

// 用户登录逻辑
func UserLogin(username, password string) (*models.User, error) {
	user := models.User{}
	query := mysqldb.Mysql.Where("username = ?", username).First(&user)
	if query.RowsAffected == 0 {
		return nil, fmt.Errorf("用户不存在")
	}
	if user.Password != utils.MD5String(password) {
		return nil, fmt.Errorf("密码错误")
	}
	return &user, nil
}

// 创建用户逻辑
func UserCreate(username, email, password string) (int64, error) {
	//先判断用户名或者邮箱是否被注册
	if IsUserExistByName(username) || IsUserExistByEmail(email) {
		return 0, fmt.Errorf("用户名或者邮箱已被注册")
	}
	passwordmd5 := utils.MD5String(password)
	user := models.User{
		Username:  username,
		Email:     email,
		Password:  passwordmd5,
		CreatedAt: carbon.Now().ToDateTimeString(),
	}
	result := mysqldb.Mysql.Create(&user)
	return user.Uid, result.Error
}

// 判断邮箱是否存在
func IsUserExistByEmail(email string) bool {
	user := models.User{}
	query := mysqldb.Mysql.Where("email = ?", email).First(&user)
	return query.RowsAffected > 0
}

// 判断用户名是否存在
func IsUserExistByName(username string) bool {
	user := models.User{}
	query := mysqldb.Mysql.Where("username = ?", username).First(&user)
	return query.RowsAffected > 0
}

func UpdateUserInfo(uid int64, password, NewUsername, NewEmail string) (bool, error) {
	user := models.User{}
	query := mysqldb.Mysql.Where("uid = ?", uid).First(&user)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("用户不存在")
	}
	if user.Password != utils.MD5String(password) {
		return false, fmt.Errorf("密码错误")
	}
	user.Username = NewUsername
	user.Email = NewEmail
	err := mysqldb.Mysql.Save(&user).Error
	return err == nil, err
}
