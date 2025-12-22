package reqs

type AdminLoginReq struct {
	Username string `form:"username" json:"username" binding:"required,min=3,max=50"`
	Password string `form:"password" json:"password" binding:"required,min=6,max=50"`
}

type GetUserDetailReq struct {
	Uid int64 `form:"uid" json:"uid" binding:"required"`
}

type BlockUserReq struct {
	Uid    int64  `form:"uid" json:"uid" binding:"required"`
	Reason string `form:"reason" json:"reason" binding:"required"`
}

// 根据uid搜索 或者 nickname 或者 phone
type SearchUserReq struct {
	Uid      int64  `form:"uid" json:"uid"`
	NickName string `form:"nick_name" json:"nick_name"`
	Phone    string `form:"phone" json:"phone"`
}
