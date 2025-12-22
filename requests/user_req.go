package reqs

type GetUsesrListReq struct {
	PageNum  int `form:"page_num" json:"page_num" binding:"required,min=1"`
	PageSize int `form:"page_size" json:"page_size" binding:"required,min=1,max=100"`
}

type UserLoginReq struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=50"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=50"`
}

type UserCreateReq struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Username string `form:"username" json:"username" binding:"required,min=3,max=50"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=50"`
}

type UserModifyPasswordReq struct {
	Uid         int64  `form:"uid" json:"uid" binding:"required"`
	OldPassword string `form:"old_password" json:"old_password" binding:"required,min=6,max=50"`
	NewPassword string `form:"new_password" json:"new_password" binding:"required,min=6,max=50"`
}

type UpdateUserInfo struct {
	Uid         int64  `form:"uid" json:"uid" binding:"required"`
	Password    string `form:"password" json:"password" binding:"required,min=6,max=50"`
	NewUsername string `form:"new_username" json:"new_username" binding:"required,min=3,max=50"`
	NewEmail    string `form:"new_email" json:"new_email" binding:"required,email"`
}
