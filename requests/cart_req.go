package reqs

type AddToCartReq struct {
	UserID    int64 `json:"user_id" binding:"required"`
	ProductID int64 `json:"product_id" binding:"required"`
	Quantity  int64 `json:"quantity" binding:"required"`
}

type RemoveFromCartReq struct {
	ProductID int64 `json:"product_id" binding:"required"`
}

type GetCartItemsReq struct {
	UserID int64 `json:"user_id" binding:"required"`
}

type UpdateCartItemReq struct {
	CartID   string `json:"-"` // 通过URL参数传递
	Quantity int64  `json:"quantity" binding:"required"`
}
