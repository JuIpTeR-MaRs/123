package logics

import (
	"shop_server/internal/models"
	"shop_server/pkg/mysqldb"
	reqs "shop_server/requests"
	"strconv"
)

func AddReview(req *reqs.AddReviewReq) (int64, error) {
	review := &models.ProductReview{}
	review.UserID = req.UserID
	review.ProductID = req.ProductID
	review.Rating = req.Rating
	review.Comment = req.Comment

	result := mysqldb.Mysql.Create(review)
	return review.Id, result.Error
}

func GetReviews(productID int64, pageNum, pageSize int) ([]models.ProductReview, error) {
	var reviewList []models.ProductReview
	query := mysqldb.Mysql.Where("product_id = ?", productID).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&reviewList)
	return reviewList, query.Error
}

func DeleteReview(reviewID int64) (bool, error) {
	review := &models.ProductReview{}
	query := mysqldb.Mysql.Where("id = ?", reviewID).First(&review)
	if query.RowsAffected == 0 {
		return false, nil
	}
	err := mysqldb.Mysql.Delete(&review).Error
	return err == nil, err
}

// GetReviewById 获取单个评价详情
func GetReviewById(reviewId string) (*models.ProductReview, error) {
	id, err := strconv.ParseInt(reviewId, 10, 64)
	if err != nil {
		return nil, err
	}

	review := &models.ProductReview{}
	query := mysqldb.Mysql.Where("id = ?", id).First(&review)
	if query.RowsAffected == 0 {
		return nil, nil
	}
	return review, query.Error
}

// UpdateReview 更新评价
func UpdateReview(req *reqs.UpdateReviewReq) (bool, error) {
	id, err := strconv.ParseInt(req.Id, 10, 64)
	if err != nil {
		return false, err
	}

	review := &models.ProductReview{}
	query := mysqldb.Mysql.Where("id = ?", id).First(&review)
	if query.RowsAffected == 0 {
		return false, nil
	}

	review.Rating = req.Rating
	review.Comment = req.Comment
	err = mysqldb.Mysql.Save(&review).Error
	return err == nil, err
}

// GetUserReviews 获取用户的所有评价
func GetUserReviews(userId string, pageNumStr, pageSizeStr string) ([]models.ProductReview, error) {
	pageNum, _ := strconv.Atoi(pageNumStr)
	if pageNum == 0 {
		pageNum = 1
	}
	pageSize, _ := strconv.Atoi(pageSizeStr)
	if pageSize == 0 {
		pageSize = 10
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, err
	}

	var reviews []models.ProductReview
	query := mysqldb.Mysql.Where("user_id = ?", userIdInt).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&reviews)
	return reviews, query.Error
}
