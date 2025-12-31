package orderctl

import (
	"shop_server/internal/logics"
	"shop_server/nets"
	"shop_server/pkg/logs"
	reqs "shop_server/requests"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateOrder(c *gin.Context) {
	var in struct {
		UserID          int64   `json:"user_id"`
		ProductID       int64   `json:"product_id"`
		ProductName     string  `json:"product_name"`
		Quantity        int64   `json:"quantity"`
		TotalAmount     float64 `json:"total_amount"`
		PayAmount       float64 `json:"pay_amount"`
		UnitPrice       float64 `json:"unit_price"`
		ReceiverName    string  `json:"receiver_name"`
		ReceiverPhone   string  `json:"receiver_phone"`
		ReceiverAddress string  `json:"receiver_address"`
	}
	if err := c.ShouldBindJSON(&in); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	if in.UserID == 0 || in.ProductID == 0 || in.Quantity == 0 || in.ReceiverName == "" || in.ReceiverPhone == "" || in.ReceiverAddress == "" {
		nets.Fail(c, 400, "缺少必要参数: user_id/product_id/quantity/receiver_*")
		return
	}
	req := &reqs.CreateOrderReq{
		UserID:          in.UserID,
		ProductID:       in.ProductID,
		ProductName:     in.ProductName,
		Quantity:        in.Quantity,
		TotalAmount:     in.TotalAmount,
		PayAmount:       in.PayAmount,
		UnitPrice:       in.UnitPrice,
		ReceiverName:    in.ReceiverName,
		ReceiverPhone:   in.ReceiverPhone,
		ReceiverAddress: in.ReceiverAddress,
	}
	logs.ZapLogger.Info("CreateOrder handler received request", zap.Any("payload", req))
	id, err := logics.CreateOrder(req)
	if err != nil {
		nets.Fail(c, 500, "创建订单失败: "+err.Error())
		return
	}
	logs.ZapLogger.Info("CreateOrder created order", zap.Int64("order_id", id))
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
	// 先绑定可能的 query 参数（page_num, page_size）
	if err := c.ShouldBind(req); err != nil {
		// 不立即返回，后面尝试解析 path param
	}
	// 如果 userId 未通过绑定获得，则尝试从路径参数读取
	if req.UserID == 0 {
		userIdStr := c.Param("userId")
		if userIdStr == "" {
			nets.Fail(c, 400, "缺少 userId 参数")
			return
		}
		uid, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			nets.Fail(c, 400, "无效的 userId 参数: "+err.Error())
			return
		}
		req.UserID = uid
	}
	// 设置默认分页，避免 page_size=0 导致 Limit(0)
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 50
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
