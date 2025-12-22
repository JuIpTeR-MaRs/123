package logics

import (
	"fmt"
	"shop_server/internal/models"
	"shop_server/pkg/mysqldb"
	reqs "shop_server/requests"

	"github.com/golang-module/carbon"
)

func AddProduct(req *reqs.AddProductReq) (int64, error) {
	if IsProductExistByName(req.Name) {
		return 0, fmt.Errorf("已存在该商品")
	}
	product := &models.Product{}
	product.Name = req.Name
	product.CategoryId = req.CategoryId
	product.Sku = req.Sku
	product.Price = req.Price
	product.Stock = req.Stock
	product.IsOnSale = req.IsOnSale
	product.Description = req.Description
	product.CreatedAt = carbon.Now().ToDateTimeString()

	result := mysqldb.Mysql.Create(product)
	return product.Id, result.Error
}

func IsProductExistByName(name string) bool {
	product := models.Product{}
	query := mysqldb.Mysql.Where("name = ?", name).First(&product)
	return query.RowsAffected > 0
}

func GetCategoryList(pageNum, pageSize int) ([]models.ProductCategory, error) {
	var categoryList []models.ProductCategory
	query := mysqldb.Mysql.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&categoryList)
	return categoryList, query.Error
}

func AddCategory(categoryname, description string) (int64, error) {
	if IsCategoryExistByName(categoryname) {
		return 0, fmt.Errorf("已存在该商品类别")
	}
	category := &models.ProductCategory{
		Name:        categoryname,
		Description: description,
		CreatedAt:   carbon.Now().ToDateTimeString(),
	}
	result := mysqldb.Mysql.Create(category)
	return category.Id, result.Error
}

func IsCategoryExistByName(categoryname string) bool {
	name := models.ProductCategory{}
	query := mysqldb.Mysql.Where("name = ?", categoryname).First(&name)
	return query.RowsAffected > 0
}

func UpdateCategory(categoryID int64, categoryname string, description string) (bool, error) {
	category := &models.ProductCategory{}
	query := mysqldb.Mysql.Where("id = ?", categoryID).First(&category)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("商品类别不存在")
	}
	category.Name = categoryname
	category.Description = description
	err := mysqldb.Mysql.Save(category).Error
	return err == nil, err
}

func DeleteCategory(categoryID int64) (bool, error) {
	category := &models.ProductCategory{}
	query := mysqldb.Mysql.Where("id = ?", categoryID).First(&category)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("商品类别不存在")
	}
	err := mysqldb.Mysql.Delete(category).Error
	return err == nil, err
}

func GetProductList(pageNum, pageSize int) ([]models.Product, error) {
	var productList []models.Product
	query := mysqldb.Mysql.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&productList)
	return productList, query.Error
}

func UpdateProduct(req *reqs.UpdateProductReq) (bool, error) {
	product := &models.Product{}
	query := mysqldb.Mysql.Where("id = ?", req.ProductID).First(&product)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("商品不存在")
	}
	product.Name = req.Name
	product.CategoryId = req.CategoryId
	product.Description = req.Description
	product.Price = req.Price
	product.Sku = req.Sku
	product.Stock = req.Stock
	err := mysqldb.Mysql.Save(product).Error
	return err == nil, err
}

func DeleteProduct(productID int64) (bool, error) {
	product := &models.Product{}
	query := mysqldb.Mysql.Where("id = ?", productID).First(&product)
	if query.RowsAffected == 0 {
		return false, fmt.Errorf("商品不存在")
	}
	err := mysqldb.Mysql.Delete(product).Error
	return err == nil, err
}
