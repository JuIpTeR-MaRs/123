package logics

import (
	"fmt"
	"math/rand"
	"shop_server/internal/models"
	"shop_server/pkg/logs"
	"shop_server/pkg/mysqldb"
	reqs "shop_server/requests"
	"time"

	"go.uber.org/zap"
)

func CreateOrder(req *reqs.CreateOrderReq) (int64, error) {
	order := &models.Order{}
	order.UserID = req.UserID
	order.OrderNo = GenerateOrderNo()
	order.ProductID = req.ProductID
	// 如果前端未提供 product name / unit price，则从产品表查询
	order.Quantity = req.Quantity
	if req.ProductName == "" || req.UnitPrice == 0 {
		// 尝试查询产品信息
		prod := &models.Product{}
		q := mysqldb.Mysql.Where("id = ?", req.ProductID).First(prod)
		if q.RowsAffected > 0 {
			order.ProductName = prod.Name
			order.UnitPrice = prod.Price
		} else {
			order.ProductName = req.ProductName
			order.UnitPrice = req.UnitPrice
		}
	} else {
		order.ProductName = req.ProductName
		order.UnitPrice = req.UnitPrice
	}
	// 计算金额
	order.TotalAmount = order.UnitPrice * float64(order.Quantity)
	order.PayAmount = order.TotalAmount
	order.ReceiverName = req.ReceiverName
	order.ReceiverPhone = req.ReceiverPhone
	order.ReceiverAddress = req.ReceiverAddress

	result := mysqldb.Mysql.Create(order)
	if result.Error != nil {
		logs.ZapLogger.Error("CreateOrder failed to insert order", zap.Any("order", order), zap.Error(result.Error))
	} else {
		logs.ZapLogger.Info("CreateOrder inserted order", zap.Int64("id", order.Id))
	}
	return order.Id, result.Error
}
func GenerateOrderNo() string {
	return "SO" + time.Now().Format("20060102") + fmt.Sprintf("%04d", rand.Intn(10000))
}

func CancelOrder(orderNo string) (bool, error) {
	order := &models.Order{}
	query := mysqldb.Mysql.Where("order_no = ?", orderNo).First(&order)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("订单不存在")
	}
	err := mysqldb.Mysql.Delete(&order).Error
	return err == nil, err
}

func GetOrderDetailByUserID(UserID int64) (*models.Order, error) {
	order := &models.Order{}
	// 按 user_id 查询用户的最新订单（如果需要返回列表，可改为 Find）
	query := mysqldb.Mysql.Where("user_id = ?", UserID).First(&order)
	if query.RowsAffected == 0 {
		return nil, fmt.Errorf("订单不存在")
	}
	return order, query.Error
}

func GetOrderDetailByOrderNo(orderNo string) (*models.Order, error) {
	order := &models.Order{}
	query := mysqldb.Mysql.Where("order_no = ?", orderNo).First(&order)
	if query.RowsAffected == 0 {
		return nil, fmt.Errorf("订单不存在")
	}
	return order, query.Error
}

func GetOrderList(userID int64, pageNum, pageSize int) ([]models.Order, error) {
	var orderList []models.Order
	query := mysqldb.Mysql.Where("user_id = ?", userID).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&orderList)
	return orderList, query.Error
}

func PayOrder(orderNo string) (bool, error) {
	order := &models.Order{}
	query := mysqldb.Mysql.Where("order_no = ?", orderNo).First(&order)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("订单不存在")
	}
	order.Status = "1" //1表示已支付
	err := mysqldb.Mysql.Save(&order).Error
	return err == nil, err
}

func DeliveryOrder(orderNo string) (bool, error) {
	order := &models.Order{}
	query := mysqldb.Mysql.Where("order_no = ?", orderNo).First(&order)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("订单不存在")
	}
	order.Status = "2" //2表示已发货
	err := mysqldb.Mysql.Save(&order).Error
	return err == nil, err
}

func CompleteOrder(orderNo string) (bool, error) {
	order := &models.Order{}
	query := mysqldb.Mysql.Where("order_no = ?", orderNo).First(&order)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("订单不存在")
	}
	order.Status = "3" //3表示已完成
	err := mysqldb.Mysql.Save(&order).Error
	return err == nil, err
}

func RefundOrder(orderNo string) (bool, error) {
	order := &models.Order{}
	query := mysqldb.Mysql.Where("order_no = ?", orderNo).First(&order)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("订单不存在")
	}
	order.Status = "4" //4表示申请退款
	err := mysqldb.Mysql.Save(&order).Error
	return err == nil, err
}
