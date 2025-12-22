package orderctl

import (
	"shop_server/internal/logics"
	"shop_server/nets"
	reqs "shop_server/requests"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	req := &reqs.CreateOrderReq{}
	if err := c.ShouldBind(req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	id, err := logics.CreateOrder(req)
	if err != nil {
		nets.Fail(c, 500, "创建订单失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"order_id": id})
}

func CancelOrder(c *gin.Context) {
	req := &reqs.CancelOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.CancelOrder(req.OrderNo)
	if err != nil {
		nets.Fail(c, 500, "取消订单失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "取消订单失败")
		return
	}
	nets.Success(c, gin.H{"msg": "取消订单成功"})
}

func GetOrderDetailByUserID(c *gin.Context) {
	req := &reqs.GetOrderDetailByUserIDReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	orderDetail, err := logics.GetOrderDetailByUserID(req.UserID)
	if err != nil {
		nets.Fail(c, 500, "获取订单详情失败: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"order_item": orderDetail})
}

func GetOrderDetailByOrderNo(c *gin.Context) {
	req := &reqs.GetOrderDetailByOrderNoReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	orderDetail, err := logics.GetOrderDetailByOrderNo(req.OrderNo)
	if err != nil {
		nets.Fail(c, 500, "获取订单详情失败: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"order_item": orderDetail})
}

func GetOrderList(c *gin.Context) {
	req := &reqs.GetOrderListReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	orderList, err := logics.GetOrderList(req.UserID, req.PageNum, req.PageSize)
	if err != nil {
		nets.Fail(c, 500, "获取订单列表失败: "+err.Error())
		return
	}
	c.JSON(200, gin.H{"order_list": orderList})
}

func PayOrder(c *gin.Context) {
	req := &reqs.PayOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.PayOrder(req.OrderNo)
	if err != nil {
		nets.Fail(c, 500, "支付订单失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "支付订单失败")
		return
	}
	nets.Success(c, gin.H{"msg": "支付订单成功，待发货"})
}

func DeliveryOrder(c *gin.Context) {
	req := &reqs.DeliveryOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.DeliveryOrder(req.OrderNo)
	if err != nil {
		nets.Fail(c, 500, "发货失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "发货失败")
		return
	}
	nets.Success(c, gin.H{"msg": "发货成功"})
}

func CompleteOrder(c *gin.Context) {
	req := &reqs.CompleteOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.CompleteOrder(req.OrderNo)
	if err != nil {
		nets.Fail(c, 500, "完成订单失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "完成订单失败")
		return
	}
	nets.Success(c, gin.H{"msg": "订单已完成: " + req.OrderNo})
}

func RefundOrder(c *gin.Context) {
	req := &reqs.RefundOrderReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.RefundOrder(req.OrderNo)
	if err != nil {
		nets.Fail(c, 500, "申请退款失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "申请退款失败")
		return
	}
	nets.Success(c, gin.H{"msg": "申请退款成功"})
}
