package logics

import (
	"shop_server/internal/models"
	"shop_server/pkg/mysqldb"
	reqs "shop_server/requests"
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
