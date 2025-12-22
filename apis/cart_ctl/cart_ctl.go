package cartctl

import (
	"shop_server/internal/logics"
	"shop_server/nets"
	reqs "shop_server/requests"

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
		nets.Fail(c, 500, "添加到购物车失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"cart_item_id": id})
}

func RemoveFromCart(c *gin.Context) {
	req := &reqs.RemoveFromCartReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.RemoveFromCart(req.ProductID)
	if err != nil {
		nets.Fail(c, 500, "从购物车移除商品失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "从购物车移除商品失败")
		return
	}
	nets.Success(c, gin.H{"msg": "从购物车移除商品成功"})
}

func GetCartItems(c *gin.Context) {
	req := &reqs.GetCartItemsReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	cartItems, err := logics.GetCartItems(req.UserID)
	if err != nil {
		nets.Fail(c, 500, "获取购物车商品失败: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"cart_list": cartItems})

}
