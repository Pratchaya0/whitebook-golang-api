package category

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func GetListCategories(c *gin.Context) {
	var paginationDto requests.PaginationDto
	var categories []entities.Category

	if err := c.ShouldBind(&paginationDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Scopes(entities.Paginate(categories, &paginationDto)).Find(&categories); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "No data categories.", nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", categories)
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category entities.Category

	if tx := entities.DB().Where("id = ?", id).First(&category); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Category [id: %s] not found.", id), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", category)
}

func CreateCategory(c *gin.Context) {
	var category entities.Category

	if err := c.ShouldBind(&category); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	categoryCreate := entities.Category{
		Name:        category.Name,
		Description: category.Description,
	}

	if err := entities.DB().Create(&categoryCreate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", categoryCreate)
}

func UpdateCategory(c *gin.Context) {
	var categoryRequest entities.Category
	var categoryUpdate entities.Category

	if err := c.ShouldBind(&categoryRequest); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Where("id = ?", categoryRequest.ID).First(&categoryUpdate); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusBadRequest, fmt.Sprintf("Category [id: %d] not found.", categoryRequest.ID), nil)
		return
	}

	categoryUpdate.Name = categoryRequest.Name
	categoryUpdate.Description = categoryRequest.Description

	if err := entities.DB().Where("id = ?", categoryRequest.ID).Save(&categoryUpdate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", categoryUpdate)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	var category entities.Category

	if tx := entities.DB().Where("id = ?", id).First(&category); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Category [id: %s] not found.", id), nil)
		return
	}

	if err := entities.DB().Where("id = ?", id).Delete(&category).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", category)
}
