package categories

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @Summary Get a list of books in the the store
// @Tag Category
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /category/create [post]
func CreateCategory(c *gin.Context) {
	var category entities.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryCreate := entities.Category{
		CategoryName: category.CategoryName,
		CategoryIcon: category.CategoryIcon,
	}

	if err := entities.DB().Create(&categoryCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": categoryCreate})
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /category/{categoryId} [get]
func GetCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")

	var category entities.Category

	if tx := entities.DB().Where("id = ?", categoryId).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// @Summary Get a list of books in the the store
// @Tag Category
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /categories [get]
func GetListCategories(c *gin.Context) {
	var categories []entities.Category

	if err := entities.DB().Raw("SELECT * FROM categories").Scan(&categories).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categories})
}

// @Summary Get a list of books in the the store
// @Tag Category
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /category/update [patch]
func UpdateCategory(c *gin.Context) {
	var category entities.Category

	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", category.ID).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category role not found"})
		return
	}

	if err := entities.DB().Save(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": category})
}

// @Summary Get a list of books in the the store
// @Tag Category
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /category/delete/{categoryId} [delete]
func DeleteCategory(c *gin.Context) {
	categoryId := c.Param("categoryId")

	if tx := entities.DB().Exec("DELETE FROM categories WHERE id = ?", categoryId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": categoryId})
}
