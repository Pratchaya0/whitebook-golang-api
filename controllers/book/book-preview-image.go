package book

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

// @Summary Get a list of book preview images
// @Tag Book
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /bookPreviewImages [get]
func GetListBookPreviewImages(c *gin.Context) {
	var paginationDto requests.PaginationDto
	var bookPreviewImages []entities.Book

	if err := c.ShouldBind(&paginationDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := entities.DB().Scopes(entities.Paginate(bookPreviewImages, &paginationDto)).Find(&bookPreviewImages).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", bookPreviewImages)
}

// @Summary Get a list of book preview images
// @Tag Book
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /bookPreviewImages/{id} [get]
func GetBookPreviewImage(c *gin.Context) {
	id := c.Param("id")
	var bookPreviewImage entities.BookPreviewImage

	if tx := entities.DB().Where("id = ?", id).First(&bookPreviewImage); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book Preview Image [id: %s] not found.", id), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", bookPreviewImage)
}

// @Summary Get a list of book preview images
// @Tag Book
// @Param bookPreviewImageUpdateDto query requests.BookPreviewImageUpdateDto true "BookPreviewImageUpdateDto"
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /bookPreviewImages/{id} [get]
func UpdateBookPreviewImage(c *gin.Context) {
	var bookPreviewImageDto requests.BookPreviewImageUpdateDto
	var bookPreviewImage entities.BookPreviewImage

	if err := c.ShouldBind(&bookPreviewImageDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Where("id = ?", bookPreviewImageDto.ID).First(&bookPreviewImage); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book preview image [id: %d] not found.", bookPreviewImageDto.ID), nil)
		return
	}

	bookPreviewImage = entities.BookPreviewImage{
		ImageLink: bookPreviewImageDto.ImageLink,
	}

	if err := entities.DB().Save(&bookPreviewImage).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", bookPreviewImage)
}

// @Summary Get a list of book preview images
// @Tag Book
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /bookPreviewImages/delete/{id} [get]
func DeleteBookPreviewImage(c *gin.Context) {
	id := c.Param("id")
	var bookPreviewImage entities.BookPreviewImage

	if tx := entities.DB().Where("id = ?", id).First(&bookPreviewImage); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book [id: %s] not found.", id), nil)
		return
	}

	if err := entities.DB().Where("id = ?", id).Delete(&bookPreviewImage).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", bookPreviewImage)
}
