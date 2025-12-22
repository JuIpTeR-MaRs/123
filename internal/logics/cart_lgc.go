package logics

import (
	"fmt"
	"shop_server/internal/models"
	"shop_server/pkg/mysqldb"
	reqs "shop_server/requests"
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

func RemoveFromCart(productID int64) (bool, error) {
	cart := &models.CartItem{}
	query := mysqldb.Mysql.Where("product_id = ?", productID).First(&cart)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("商品不存在")
	}
	err := mysqldb.Mysql.Delete(&cart).Error
	return err == nil, err
}

func GetCartItems(userID int64) ([]models.CartItem, error) {
	var cartItems []models.CartItem
	query := mysqldb.Mysql.Where("user_id = ?", userID).Find(&cartItems)
	return cartItems, query.Error
}
