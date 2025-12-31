package logics

import (
	"fmt"
	"shop_server/internal/models"
	"shop_server/pkg/mysqldb"
	reqs "shop_server/requests"
	"strconv"
)

func AddToCart(req *reqs.AddToCartReq) (int64, error) {
	if IsProductExistById(req.ProductID) {
		return 0, fmt.Errorf("商品已存在")
	}
	cartItem := &models.CartItem{}
	cartItem.UserID = req.UserID
	cartItem.ProductID = req.ProductID
	cartItem.Quantity = req.Quantity

	result := mysqldb.Mysql.Create(cartItem)
	return cartItem.Id, result.Error
}

func IsProductExistById(productID int64) bool {
	cart := models.CartItem{}
	query := mysqldb.Mysql.Where("id = ?", productID).First(&cart)
	return query.RowsAffected > 0
}

func RemoveFromCart(cartID int64) (bool, error) {
	cart := &models.CartItem{}
	query := mysqldb.Mysql.Where("id = ?", cartID).First(&cart)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("购物车项不存在")
	}
	err := mysqldb.Mysql.Delete(&cart).Error
	return err == nil, err
}

func GetCartItems(userID int64) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	query := mysqldb.Mysql.Where("user_id = ?", userID).Find(&cartItems)
	return cartItems, query.Error
}

// UpdateCartItem 更新购物车商品数量
func UpdateCartItem(req *reqs.UpdateCartItemReq) (bool, error) {
	cartId, err := strconv.ParseInt(req.CartID, 10, 64)
	if err != nil {
		return false, fmt.Errorf("无效的购物车ID")
	}

	cart := &models.CartItem{}
	query := mysqldb.Mysql.Where("id = ?", cartId).First(&cart)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("购物车项不存在")
	}

	cart.Quantity = req.Quantity
	err = mysqldb.Mysql.Save(&cart).Error
	return err == nil, err
}

// ClearCart 清空用户购物车
func ClearCart(userId string) (bool, error) {
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return false, fmt.Errorf("无效的用户ID")
	}

	err = mysqldb.Mysql.Where("user_id = ?", userIdInt).Delete(&models.CartItem{}).Error
	return err == nil, err
}
