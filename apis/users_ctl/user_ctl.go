package usersctl

import (
	"fmt"
	"shop_server/internal/logics"
	"shop_server/nets"
	reqs "shop_server/requests"

	"github.com/gin-gonic/gin"
)

// @Summary 获取用户列表
// @Description 获取用户列表
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param user query models.GetUserInfoReq
func GetUserList(c *gin.Context) {
	getUserListReq := &reqs.GetUsesrListReq{}
	if err := c.ShouldBind(getUserListReq); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	userList, err := logics.GetUserList(getUserListReq.PageNum, getUserListReq.PageSize)
	if err != nil {
		nets.Fail(c, 500, "获取用户列表失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"user_list": userList})
}

// @Summary 用户登录
// @Description 用户登录
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param user body models.UserLoginReq
func UserLogin(c *gin.Context) {
	loginReq := &reqs.UserLoginReq{}
	if err := c.ShouldBind(loginReq); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	user, err := logics.UserLogin(loginReq.Username, loginReq.Password)
	if err != nil {
		nets.Fail(c, 400, "用户登录失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"user": user})
}

// @Summary 用户注册
// @Description 用户注册
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param user body models.User
func UserCreate(c *gin.Context) {
	useReq := &reqs.UserCreateReq{}
	if err := c.ShouldBind(useReq); err != nil {
		fmt.Println(useReq.Username, useReq.Email, useReq.Password)
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	uid, err := logics.UserCreate(useReq.Username, useReq.Email, useReq.Password)
	if err != nil {
		nets.Fail(c, 400, "用户创建失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"uid": uid, "username": useReq.Username, "email": useReq.Email})
}

// @Summary 修改用户密码
// @Description 修改用户密码
// @Tags 用户相关接口
// @Accept json
// @Produce json
// @Param user body models.UserModifyPasswordReq
func UserModityPassword(c *gin.Context) {
	modifyReq := &reqs.UserModifyPasswordReq{}
	if err := c.ShouldBind(modifyReq); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.UpdateUserPassword(modifyReq.Uid, modifyReq.OldPassword, modifyReq.NewPassword)
	if err != nil {
		nets.Fail(c, 400, "修改密码失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 400, "修改密码失败: 用户不存在或旧密码错误")
		return
	}
	nets.Success(c, gin.H{"msg": "修改密码成功"})
}

func UpdateUserInfo(c *gin.Context) {
	updateReq := &reqs.UpdateUserInfo{}
	if err := c.ShouldBind(updateReq); err != nil {
		fmt.Printf("参数绑定错误: %v\n", err)
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.UpdateUserInfo(updateReq.Uid, updateReq.Password, updateReq.NewUsername, updateReq.NewEmail)
	if err != nil {
		nets.Fail(c, 400, "修改失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 400, "修改失败: 用户不存在或密码错误")
		return
	}
	nets.Success(c, gin.H{"msg": "修改成功"})
}
