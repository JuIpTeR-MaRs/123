package reqs

type AddProductReq struct {
	Name        string   `json:"name" binding:"required"`
	CategoryId  string   `json:"category_id" binding:"required"`
	Description string   `json:"description"`
	Price       float64  `json:"price" binding:"required"`
	Sku         string   `json:"sku" binding:"required"`
	Stock       int64    `json:"stock" binding:"required"`
	CoverImage  string   `json:"cover_image"`
	Images      []string `json:"images"`
	IsOnSale    bool     `json:"is_on_sale"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
	DeletedAt   string   `json:"deleted_at"`
}

type GetCategoryListReq struct {
	PageNum  int `form:"page_num" json:"page_num" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}

type AddCategoryReq struct {
	CategoryName string `json:"name" binding:"required"`
	Description  string `json:"description"`
}

type UpdateCategoryReq struct {
	CategoryID   int64  `json:"category_id" binding:"required"`
	CategoryName string `json:"name" binding:"required"`
	Description  string `json:"description"`
}

type DeleteCategoryReq struct {
	CategoryID int64 `json:"category_id" binding:"required"`
}

type GetProductListReq struct {
	PageNum  int `form:"page_num" json:"page_num" binding:"required"`
	PageSize int `form:"page_size" json:"page_size" binding:"required"`
}

type UpdateProductReq struct {
	ProductID   int64   `json:"product_id" binding:"required"`
	Name        string  `json:"name" binding:"required"`
	CategoryId  string  `json:"category_id" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required"`
	Sku         string  `json:"sku" binding:"required"`
	Stock       int64   `json:"stock" binding:"required"`
}

type DeleteProductReq struct {
	ProductID int64 `json:"product_id" binding:"required"`
}
