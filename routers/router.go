package routers

import (
	adminctl "shop_server/apis/admin_ctl"
	cart_ctl "shop_server/apis/cart_ctl"
	order_ctl "shop_server/apis/order_ctl"
	productctl "shop_server/apis/products_ctl"
	review_ctl "shop_server/apis/review_ctl"
	usersctl "shop_server/apis/users_ctl"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	//http://localhost:8080/user/
	user := r.Group("/user")
	user.POST("/register", usersctl.UserCreate)
	user.POST("/login", usersctl.UserLogin)
	user.GET("/list", usersctl.GetUserList)
	//更新用户信息
	user.POST("/update", usersctl.UpdateUserInfo)
	//修改用户密码
	user.POST("/modify_psw", usersctl.UserModityPassword)

	//http://localhost:8080/admin/
	admin := r.Group("/admin")
	admin.POST("/login", adminctl.AdminLogin)
	admin.POST("/userlist", adminctl.GetUserListByAdmin)
	admin.POST("/queryuser", adminctl.QueryUserDetail)
	admin.POST("/blockuser", adminctl.BlockUser)
	admin.POST("/unblockuser", adminctl.UnblockUser)
	admin.GET("/searchuser", adminctl.SearchUser)

	//http://localhost:8080/product/
	//商品
	//商品类别管理
	product := r.Group("/product")
	product.GET("/getcategory", productctl.GetCategoryList)
	product.POST("/addcategory", productctl.AddCategory)
	product.POST("/updatecategory", productctl.UpdateCategory)
	product.POST("/deletecategory", productctl.DeleteCategory)
	//商品管理
	product.GET("/getproduct", productctl.GetProductList)
	product.POST("/addproduct", productctl.AddProduct)
	product.POST("/updateproduct", productctl.UpdateProduct)
	product.POST("/deleteproduct", productctl.DeleteProduct)

	//http://localhost:8080/cart/
	//购物车
	cart := r.Group("/cart")
	cart.POST("/add", cart_ctl.AddToCart)
	cart.POST("/remove", cart_ctl.RemoveFromCart)
	cart.POST("/list", cart_ctl.GetCartItems)

	//http://localhost:8080/order/
	//订单
	order := r.Group("/order")
	order.POST("/create", order_ctl.CreateOrder)
	order.POST("/cancel", order_ctl.CancelOrder)
	//根据用户ID获取订单详情
	order.POST("/detailbyuser", order_ctl.GetOrderDetailByUserID)
	//根据订单号获取订单详情
	order.POST("/detailbyorderno", order_ctl.GetOrderDetailByOrderNo)
	//订单列表
	order.POST("/list", order_ctl.GetOrderList)
	//支付
	order.POST("/pay", order_ctl.PayOrder)
	//发货
	order.POST("/delivery", order_ctl.DeliveryOrder)
	//完成交易
	order.POST("/complete", order_ctl.CompleteOrder)
	//申请退款
	order.POST("/refund", order_ctl.RefundOrder)

	//http://localhost:8080/review/
	//用户评价
	review := r.Group("/review")
	review.POST("/add", review_ctl.AddReview)
	//获取某个商品的全部评价
	review.POST("/list", review_ctl.GetReviews)
	//根据评论id删除评价
	review.POST("/delete", review_ctl.DeleteReview)
}
