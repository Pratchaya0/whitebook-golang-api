package review

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func GetListReviewsByBookID(c *gin.Context) {
	var paginationDto requests.PaginationDto
	var reviews []entities.Review

	bookId, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := c.ShouldBind(&paginationDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Scopes(entities.Paginate(reviews, &paginationDto)).Where(&entities.Review{BookID: uint(bookId)}).Find(&reviews); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "No data reviews.", nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", reviews)
}

func GetListReviewsByUserID(c *gin.Context) {
	var paginationDto requests.PaginationDto
	var reviews []entities.Review

	userId, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if err := c.ShouldBind(&paginationDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Scopes(entities.Paginate(reviews, &paginationDto)).Where(&entities.Review{UserID: uint(userId)}).Find(&reviews); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "No data reviews.", nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", reviews)
}

func GetReview(c *gin.Context) {
	id := c.Param("id")
	var review entities.Review

	if tx := entities.DB().Where("id = ?", id).First(&review); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Review [id: %s] not found.", id), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", review)
}

func CreateReview(c *gin.Context) {
	var book entities.Book
	var user entities.User
	var review entities.Review

	if err := c.ShouldBind(&review); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check book is exist?
	if tx := entities.DB().Where("id = ?", review.BookID).First(&book); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book [id: %d] not found.", review.BookID), nil)
		return
	}

	// check user is exist?
	if tx := entities.DB().Where("id = ?", review.UserID).First(&user); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("User [id: %d] not found.", review.UserID), nil)
		return
	}

	reviewCreate := entities.Review{
		Rating: review.Rating,
		Detail: review.Detail,

		BookID: review.BookID,
		UserID: review.UserID,
	}

	if err := entities.DB().Create(&reviewCreate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", reviewCreate)
}

func UpdateReview(c *gin.Context) {
	var book entities.Book
	var user entities.User
	var reviewRequest entities.Review
	var reviewUpdate entities.Review

	if err := c.ShouldBind(&reviewRequest); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check book is exist?
	if tx := entities.DB().Where("id = ?", reviewRequest.BookID).First(&book); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book [id: %d] not found.", reviewRequest.BookID), nil)
		return
	}

	// check user is exist?
	if tx := entities.DB().Where("id = ?", reviewRequest.UserID).First(&user); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("User [id: %d] not found.", reviewRequest.UserID), nil)
		return
	}

	if tx := entities.DB().Where("id = ?", reviewRequest.ID).First(&reviewUpdate); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Review [id: %d] not found.", reviewRequest.ID), nil)
		return
	}

	reviewUpdate.Rating = reviewRequest.Rating
	reviewUpdate.Detail = reviewRequest.Detail

	if err := entities.DB().Where("id = ?", reviewRequest.ID).Save(&reviewUpdate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", reviewUpdate)
}

func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	var review entities.Review

	if tx := entities.DB().Where("id = ?", id).First(&review); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Review [id: %s not found.]", id), nil)
		return
	}

	if err := entities.DB().Where("id = ?", id).Delete(&review).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", review)
}
