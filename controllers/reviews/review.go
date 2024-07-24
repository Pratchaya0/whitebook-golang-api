package reviews

// import (
// 	"net/http"

// 	"github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	"github.com/gin-gonic/gin"
// )

// func CreateReview(c *gin.Context) {
// 	var review entities.Review
// 	var user entities.User
// 	var book entities.Book

// 	if err := c.ShouldBindJSON(&review); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// แทรกการ validate ไว้ช่วงนี้ของ controller
// 	if _, err := govalidator.ValidateStruct(review); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", review.UserId).First(&user); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", review.ReviewBookId).First(&book); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
// 		return
// 	}

// 	reviewCreate := entities.Review{
// 		ReviewComment: review.ReviewComment,
// 		ReviewRating:  review.ReviewRating,
// 		User:          user,
// 		Book:          book,
// 	}

// 	if err := entities.DB().Create(&reviewCreate).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"data": reviewCreate})
// }

// func GetReview(c *gin.Context) {
// 	reviewId := c.Param("reviewId")

// 	var review entities.Review

// 	if err := entities.DB().Preload("User").Preload("Book").Raw("SELECT * FROM reviews WHERE id = ?", reviewId).First(&review).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": review})
// }

// func GetListReviews(c *gin.Context) {
// 	var reviews []entities.Review

// 	if err := entities.DB().Preload("User").Preload("Book").Raw("SELECT * FROM reviews").Find(&reviews).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": reviews})
// }

// func UpdateReview(c *gin.Context) {
// 	var review entities.Review
// 	var user entities.User
// 	var book entities.Book

// 	if err := c.ShouldBindJSON(&review); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// แทรกการ validate ไว้ช่วงนี้ของ controller
// 	if _, err := govalidator.ValidateStruct(review); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", review.ID).First(&review); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "review not found"})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", review.UserId).First(&user); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", review.ReviewBookId).First(&book); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
// 		return
// 	}

// 	reviewUpdate := entities.Review{
// 		ReviewComment: review.ReviewComment,
// 		ReviewRating:  review.ReviewRating,
// 		User:          user,
// 		Book:          book,
// 	}

// 	if err := entities.DB().Save(&reviewUpdate).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": reviewUpdate})
// }

// func DeleteReview(c *gin.Context) {
// 	reviewId := c.Param("reviewId")

// 	if tx := entities.DB().Exec("DELETE FROM reviews WHERE id = ?", reviewId); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "review not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": "review id: " + reviewId + " has been deleted"})
// }
