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

// GetReviewById 获取单个评价详情
func GetReviewById(c *gin.Context) {
	reviewId := c.Param("reviewId")

	review, err := logics.GetReviewById(reviewId)
	if err != nil {
		nets.Fail(c, 500, "获取评价详情失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"review": review})
}

// UpdateReview 更新评价
func UpdateReview(c *gin.Context) {
	reviewId := c.Param("reviewId")
	req := &reqs.UpdateReviewReq{}
	if err := c.ShouldBindJSON(&req); err != nil {
		nets.Fail(c, 400, "参数错误: "+err.Error())
		return
	}

	req.Id = reviewId

	success, err := logics.UpdateReview(req)
	if err != nil {
		nets.Fail(c, 500, "更新评价失败: "+err.Error())
		return
	}
	if !success {
		nets.Fail(c, 500, "更新评价失败")
		return
	}
	nets.Success(c, gin.H{"msg": "更新评价成功"})
}

// GetUserReviews 获取用户的所有评价
func GetUserReviews(c *gin.Context) {
	userId := c.Param("userId")
	pageNum := c.DefaultQuery("pageNum", "1")
	pageSize := c.DefaultQuery("pageSize", "10")

	reviews, err := logics.GetUserReviews(userId, pageNum, pageSize)
	if err != nil {
		nets.Fail(c, 500, "获取用户评价失败: "+err.Error())
		return
	}
	nets.Success(c, gin.H{"reviews": reviews})
}
