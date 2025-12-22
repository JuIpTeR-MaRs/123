package logics

import (
	"fmt"
	"shop_server/internal/models"
	"shop_server/pkg/mysqldb"
	"shop_server/utils"
)

// adminLogin 登录
func AdminLogin(username, password string) (*models.Admin, error) {
	admin := models.Admin{}
	query := mysqldb.Mysql.Where("username = ?", username).First(&admin)
	if query.RowsAffected == 0 {
		return nil, fmt.Errorf("用户不存在")
	}
	if admin.Password != utils.MD5String(password) {
		return nil, fmt.Errorf("密码错误")
	}
	admin.Password = ""
	return &admin, nil
}

func GetUserDetail(uid int64) (*models.User, error) {
	user := &models.User{}
	query := mysqldb.Mysql.Where("uid = ?", uid).First(user)
	if query.RowsAffected == 0 {
		return nil, fmt.Errorf("用户不存在")
	}
	user.Password = ""
	return user, nil
}

// 拉黑用户
func BlockUser(uid int64, reason string) (bool, error) {
	query := mysqldb.Mysql.Create(&models.BlackListUid{Uid: uid, Reason: reason})
	return query.RowsAffected > 0, query.Error
}

// 解禁用户
func UnBlockUser(uid int64) (bool, error) {
	query := mysqldb.Mysql.Where("uid = ?", uid).Delete(&models.BlackListUid{})
	return query.RowsAffected > 0, query.Error
}

func SearchUser(uid int64, nickName string, phone string) ([]*models.User, error) {
	users := []*models.User{}
	user := &models.User{}
	query := mysqldb.Mysql.Table(user.TableName())
	if uid != 0 {
		query = query.Where("uid = ?", uid)
	}
	if nickName != "" {
		query = query.Where("nickname like ?", "%"+nickName+"%")
	}
	if phone != "" {
		query = query.Where("phone like ?", "%"+phone+"%")
	}
	query = query.Find(&users)
	return users, query.Error
}
