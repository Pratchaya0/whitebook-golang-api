package books

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

func CreateBookPreviewImage(c *gin.Context) {
	var bookPreviewImage entities.BookPreviewImage
	var book entities.Book

	if err := c.ShouldBindJSON(&bookPreviewImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", bookPreviewImage.BookId).First(&book); tx.RowsAffected == 0 {
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

func GetBookPreviewImage(c *gin.Context) {
	bookPreviewImageId := c.Param("bookPreviewImageId")

	var bookPreviewImage entities.BookPreviewImage

	if err := entities.DB().Preload("Book").Raw("SELECT * FROM book_preview_images WHERE id = ?", bookPreviewImageId).First(&bookPreviewImage).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookPreviewImage})
}

func GetListBookPreviewImages(c *gin.Context) {
	var bookPreviewImages []entities.BookPreviewImage

	if err := entities.DB().Preload("Book").Raw("SELECT * FROM book_preview_images").Find(&bookPreviewImages).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": bookPreviewImages})
}

func UpdateBookPreviewImage(c *gin.Context) {
	var bookPreviewImage entities.BookPreviewImage

	if err := c.ShouldBindJSON(&bookPreviewImage); err != nil {
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

func DeleteBookPreviewImage(c *gin.Context) {
	bookPreviewImageId := c.Param("bookPreviewImageId")

	if tx := entities.DB().Exec("DELETE FROM book_preview_images WHERE id = ?", bookPreviewImageId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book preview image not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "book preview image id: " + bookPreviewImageId + " has been deleted"})
}
