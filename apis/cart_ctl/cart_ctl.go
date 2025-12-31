package cartctl

import (
	"shop_server/internal/logics"
	"shop_server/nets"
	reqs "shop_server/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddToCart(c *gin.Context) {
	req := &reqs.AddToCartReq{}
	if err := c.ShouldBind(req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	id, err := logics.AddToCart(req)
	if err != nil {
		nets.Fail(c, 500, "添加失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"cart_item_id": id})
}

func RemoveFromCart(c *gin.Context) {
	// 从路径参数获取购物车项 ID
	cartIdStr := c.Param("cartId")
	if cartIdStr == "" {
		nets.Fail(c, 400, "缺少 cartId 参数")
		return
	}
	// 转换为整数
	cartId, err := strconv.ParseInt(cartIdStr, 10, 64)
	if err != nil {
		nets.Fail(c, 400, "无效的 cartId 参数: "+err.Error())
		return
	}
	success, err := logics.RemoveFromCart(cartId)
	if err != nil {
		nets.Fail(c, 500, "移除失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "移除失败")
		return
	}
	nets.Success(c, gin.H{"msg": "移除成功"})
}

func GetCartItems(c *gin.Context) {
	// 从路径参数读取 userId
	userIdStr := c.Param("userId")
	if userIdStr == "" {
		nets.Fail(c, 400, "缺少 userId 参数")
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		nets.Fail(c, 400, "无效的 userId 参数: "+err.Error())
		return
	}
	cartItems, err := logics.GetCartItems(userId)
	if err != nil {
		nets.Fail(c, 500, "获取失败: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"cart_list": cartItems})

}

// UpdateCartItem 更新购物车商品数量
func UpdateCartItem(c *gin.Context) {
	cartId := c.Param("cartId")
	req := &reqs.UpdateCartItemReq{}
	if err := c.ShouldBindJSON(req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}

	// 将cartId设置到请求结构中
	req.CartID = cartId

	success, err := logics.UpdateCartItem(req)
	if err != nil {
		nets.Fail(c, 500, "更新购物车失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "更新购物车失败")
		return
	}
	nets.Success(c, gin.H{"msg": "更新购物车成功"})
}

// ClearCart 清空用户购物车
func ClearCart(c *gin.Context) {
	userId := c.Param("userId")

	success, err := logics.ClearCart(userId)
	if err != nil {
		nets.Fail(c, 500, "清空购物车失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "清空购物车失败")
		return
	}
	nets.Success(c, gin.H{"msg": "清空购物车成功"})
}
