package productsctl

import (
	"fmt"
	logics "shop_server/internal/logics"
	"shop_server/nets"
	reqs "shop_server/requests"

	"github.com/gin-gonic/gin"
)

func GetCategoryList(c *gin.Context) {
	req := reqs.GetCategoryListReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	categoryList, err := logics.GetCategoryList(req.PageNum, req.PageSize)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
	c.JSON(200, gin.H{"category_list": categoryList})

}
func AddCategory(c *gin.Context) {
	req := &reqs.AddCategoryReq{}
	if err := c.ShouldBind(req); err != nil {
		fmt.Println(req.CategoryName, req.Description)
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	id, err := logics.AddCategory(req.CategoryName, req.Description)
	if err != nil {
		nets.Fail(c, 500, "添加商品类别失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"category_id": id, "category_name": req.CategoryName, "description": req.Description})

}

func UpdateCategory(c *gin.Context) {
	req := &reqs.UpdateCategoryReq{}
	if err := c.ShouldBind(req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.UpdateCategory(req.CategoryID, req.CategoryName, req.Description)
	if err != nil {
		nets.Fail(c, 500, "更新商品类别失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "更新商品类别失败")
		return
	}
	nets.Success(c, gin.H{"msg": "更新商品类别成功"})

}

func DeleteCategory(c *gin.Context) {
	req := &reqs.DeleteCategoryReq{}
	if err := c.ShouldBind(req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.DeleteCategory(req.CategoryID)
	if err != nil {
		nets.Fail(c, 500, "删除商品类别失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "删除商品类别失败")
		return
	}
	nets.Success(c, gin.H{"msg": "删除商品类别成功"})
}

func GetProductList(c *gin.Context) {
	req := reqs.GetProductListReq{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}
	productList, err := logics.GetProductList(req.PageNum, req.PageSize)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
	}
	c.JSON(200, gin.H{"product_list": productList})
}

func AddProduct(c *gin.Context) {
	req := &reqs.AddProductReq{}
	if err := c.ShouldBind(req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	id, err := logics.AddProduct(req)
	if err != nil {
		nets.Fail(c, 500, "添加商品失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"product_id": id})
}

func UpdateProduct(c *gin.Context) {
	req := &reqs.UpdateProductReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.UpdateProduct(req)
	if err != nil {
		nets.Fail(c, 500, "更新商品失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "更新商品失败")
		return
	}
	nets.Success(c, gin.H{"msg": "更新商品成功"})
}

func DeleteProduct(c *gin.Context) {
	req := &reqs.DeleteProductReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.DeleteProduct(req.ProductID)
	if err != nil {
		nets.Fail(c, 500, "删除商品失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "删除商品失败")
		return
	}
	nets.Success(c, gin.H{"msg": "删除商品成功"})
}
