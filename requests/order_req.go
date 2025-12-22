package reqs

type CreateOrderReq struct {
	UserID          int64   `json:"user_id" binding:"required"`
	ProductID       int64   `json:"product_id" binding:"required"`
	ProductName     string  `json:"product_name" binding:"required"`
	Quantity        int64   `json:"quantity" binding:"required"`
	TotalAmount     float64 `json:"total_amount" binding:"required"`
	PayAmount       float64 `json:"pay_amount" binding:"required"`
	UnitPrice       float64 `json:"unit_price" binding:"required"`
	ReceiverName    string  `json:"receiver_name" binding:"required"`
	ReceiverPhone   string  `json:"receiver_phone" binding:"required"`
	ReceiverAddress string  `json:"receiver_address" binding:"required"`
}

type CancelOrderReq struct {
	OrderNo string `json:"order_no" binding:"required"`
}

type GetOrderDetailByUserIDReq struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type GetOrderDetailByOrderNoReq struct {
	OrderNo string `json:"order_no" binding:"required"`
}

type GetOrderListReq struct {
	UserID   int64 `json:"user_id" binding:"required"`
	PageNum  int   `json:"page_num" binding:"required"`
	PageSize int   `json:"page_size" binding:"required"`
}

type PayOrderReq struct {
	OrderNo string `json:"order_no" binding:"required"`
}

type DeliveryOrderReq struct {
	OrderNo string `json:"order_no" binding:"required"`
}

type CompleteOrderReq struct {
	OrderNo string `json:"order_no" binding:"required"`
}

type RefundOrderReq struct {
	OrderNo string `json:"order_no" binding:"required"`
}
