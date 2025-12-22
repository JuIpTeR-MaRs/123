package reviewctl

import (
	"shop_server/internal/logics"
	"shop_server/nets"
	reqs "shop_server/requests"

	"github.com/gin-gonic/gin"
)

func AddReview(c *gin.Context) {
	req := &reqs.AddReviewReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	id, err := logics.AddReview(req)
	if err != nil {
		nets.Fail(c, 500, "添加评价失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"msg": "评价成功", "review_id": id})
}

func GetReviews(c *gin.Context) {
	req := &reqs.GetReviewsReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	reviews, err := logics.GetReviews(req.ProductID, req.PageNum, req.PageSize)
	if err != nil {
		nets.Fail(c, 500, "获取评价列表失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"reviews": reviews})
}

func DeleteReview(c *gin.Context) {
	req := &reqs.DeleteReviewReq{}
	if err := c.ShouldBind(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}
	success, err := logics.DeleteReview(req.Id)
	if err != nil {
		nets.Fail(c, 500, "删除评价失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "删除评价失败")
		return
	}
	nets.Success(c, gin.H{"msg": "删除评价成功"})
}
