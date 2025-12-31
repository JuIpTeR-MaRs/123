package reqs

type AddReviewReq struct {
	UserID    int64  `json:"user_id" binding:"required"`
	ProductID int64  `json:"product_id" binding:"required"`
	Rating    int    `json:"rating" binding:"required,min=1,max=5"`
	Comment   string `json:"comment" binding:"required"`
}

type GetReviewsReq struct {
	ProductID int64 `json:"product_id" binding:"required"`
	PageNum   int   `json:"page_num" binding:"required,min=1"`
	PageSize  int   `json:"page_size" binding:"required,min=1,max=100"`
}
type DeleteReviewReq struct {
	Id int64 `json:"review_id" binding:"required"`
}

type UpdateReviewReq struct {
	Id      string `json:"-"` // 通过URL参数传递
	Rating  int    `json:"rating" binding:"min=1,max=5"`
	Comment string `json:"comment"`
}
