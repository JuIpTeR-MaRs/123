package routers

import (
	adminctl "shop_server/apis/admin_ctl"
	cartctl "shop_server/apis/cart_ctl"
	order_ctl "shop_server/apis/order_ctl"
	productctl "shop_server/apis/products_ctl"
	reviewctl "shop_server/apis/review_ctl"
	usersctl "shop_server/apis/users_ctl"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {

	// 简单根路由，避免访问根路径时返回 404
	// 提供静态资源并将根路由指向静态首页
	r.Static("/static", "./static")
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/static/index.html")
	})

	// 忽略 favicon 请求，返回 204 减少日志噪音
	r.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204)
	})

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

	//http://localhost:8080/carts/
	//购物车 - 使用RESTful风格
	cart := r.Group("/carts")
	{
		cart.POST("", cartctl.AddToCart)                 // 添加商品到购物车
		cart.GET("/:userId", cartctl.GetCartItems)       // 获取用户购物车商品
		cart.PUT("/:cartId", cartctl.UpdateCartItem)     // 更新购物车商品数量
		cart.DELETE("/:cartId", cartctl.RemoveFromCart)  // 从购物车移除商品
		cart.DELETE("/clear/:userId", cartctl.ClearCart) // 清空用户购物车
	}

	//http://localhost:8080/orders/
	//订单 - 使用RESTful风格
	order := r.Group("/orders")
	{
		order.POST("", order_ctl.CreateOrder)                     // 创建订单
		order.GET("/:orderNo", order_ctl.GetOrderDetailByOrderNo) // 获取订单详情
		order.GET("/user/:userId", order_ctl.GetOrderList)        // 获取用户订单列表
		order.PUT("/cancel/:orderNo", order_ctl.CancelOrder)      // 取消订单
		order.PUT("/pay/:orderNo", order_ctl.PayOrder)            // 支付订单
		order.PUT("/deliver/:orderNo", order_ctl.DeliveryOrder)   // 发货
		order.PUT("/complete/:orderNo", order_ctl.CompleteOrder)  // 完成订单
		order.PUT("/refund/:orderNo", order_ctl.RefundOrder)      // 申请退款
	}

	//http://localhost:8080/reviews/
	//用户评价 - 使用RESTful风格
	review := r.Group("/reviews")
	{
		review.POST("", reviewctl.AddReview)                    // 添加评价
		review.GET("/product/:productId", reviewctl.GetReviews) // 获取商品评价
		review.GET("/:reviewId", reviewctl.GetReviewById)       // 获取单个评价详情
		review.PUT("/:reviewId", reviewctl.UpdateReview)        // 更新评价
		review.DELETE("/:reviewId", reviewctl.DeleteReview)     // 删除评价
		review.GET("/user/:userId", reviewctl.GetUserReviews)   // 获取用户的所有评价
	}
}
