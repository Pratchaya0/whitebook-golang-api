package books

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-preview-image/create [post]
func CreateBookPreviewImage(c *gin.Context) {
	var bookPreviewImage entities.BookPreviewImage
	var book entities.Book

	if err := c.ShouldBindJSON(&bookPreviewImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(bookPreviewImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", bookPreviewImage.BookPreviewImageBookId).First(&book); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	bookPreviewImageCreate := entities.BookPreviewImage{
		BookPreviewImageUrl: bookPreviewImage.BookPreviewImageUrl,
		Book:                book,
	}

	if err := entities.DB().Create(&bookPreviewImageCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": bookPreviewImageCreate})
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-preview-image/{bookPreviewImageId} [get]
func GetBookPreviewImage(c *gin.Context) {
	bookPreviewImageId := c.Param("bookPreviewImageId")

	var bookPreviewImage entities.BookPreviewImage

	if err := entities.DB().Preload("Book").Raw("SELECT * FROM book_preview_images WHERE id = ?", bookPreviewImageId).First(&bookPreviewImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookPreviewImage})
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-preview-images [get]
func GetListBookPreviewImages(c *gin.Context) {
	var bookPreviewImages []entities.BookPreviewImage

	if err := entities.DB().Preload("Book").Raw("SELECT * FROM book_preview_images").Find(&bookPreviewImages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookPreviewImages})
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-preview-image/update [patch]
func UpdateBookPreviewImage(c *gin.Context) {
	var bookPreviewImage entities.BookPreviewImage

	if err := c.ShouldBindJSON(&bookPreviewImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(bookPreviewImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", bookPreviewImage.ID).First(&bookPreviewImage); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book preview image not found"})
		return
	}

	if err := entities.DB().Save(&bookPreviewImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookPreviewImage})
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-preview-image/delete/{bookPreviewImageId} [delete]
func DeleteBookPreviewImage(c *gin.Context) {
	bookPreviewImageId := c.Param("bookPreviewImageId")

	if tx := entities.DB().Exec("DELETE FROM book_preview_images WHERE id = ?", bookPreviewImageId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book preview image not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "book preview image id: " + bookPreviewImageId + " has been deleted"})
}
