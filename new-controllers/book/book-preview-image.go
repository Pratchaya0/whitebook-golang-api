package book

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

// @Summary Get a list of book preview images
// @Tag Book
// @Security bearerToken
// @Produce json
// @Success 200 {object} responses.Response{} "ok"
// @Router /bookPreviewImages [get]
func GetListBookPreviewImages(c *gin.Context) {
	var bookPreviewImages []entities.Book

	if err := entities.DB().Raw("SELECT * FROM book_preview_images").Find(&bookPreviewImages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookPreviewImages,
	}

	c.JSON(http.StatusOK, webResponse)
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

	if err := entities.DB().Where("id = ?", id).First(&bookPreviewImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookPreviewImage,
	}

	c.JSON(http.StatusOK, webResponse)
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", bookPreviewImageDto.ID).First(&bookPreviewImage); tx.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Book preview image [id: %d] not found.", bookPreviewImageDto.ID)})
		return
	}

	bookPreviewImage = entities.BookPreviewImage{
		ImageLink: bookPreviewImageDto.ImageLink,
	}

	if err := entities.DB().Save(&bookPreviewImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookPreviewImage,
	}

	c.JSON(http.StatusOK, webResponse)
}
