package adminctl

import (
	"shop_server/internal/logics"
	reqs "shop_server/requests"

	"github.com/gin-gonic/gin"
)

// 管理员登录
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员

func AdminLogin(c *gin.Context) {
	//获取前端登录参数
	req := &reqs.AdminLoginReq{}
	if err := c.ShouldBind(req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	admin, err := logics.AdminLogin(req.Username, req.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{
		"message": "登录成功",
		"admin":   admin,
	})
}

func GetUserListByAdmin(c *gin.Context) {
	req := reqs.GetUsesrListReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	userList, err := logics.GetUserList(req.PageNum, req.PageSize)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
	c.JSON(200, gin.H{"user_list": userList})
}

func QueryUserDetail(c *gin.Context) {
	req := reqs.GetUserDetailReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	user, err := logics.GetUserDetail(req.Uid)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
	c.JSON(200, gin.H{"user": user})
}

func BlockUser(c *gin.Context) {
	req := reqs.BlockUserReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	success, err := logics.BlockUser(req.Uid, req.Reason)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
	c.JSON(200, gin.H{"success": success})
}

func UnblockUser(c *gin.Context) {
	req := reqs.BlockUserReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	success, err := logics.UnBlockUser(req.Uid)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
	c.JSON(200, gin.H{"success": success})
}

func SearchUser(c *gin.Context) {
	req := reqs.SearchUserReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	userList, err := logics.SearchUser(req.Uid, req.NickName, req.Phone)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
	c.JSON(200, gin.H{"user_list": userList})
}
