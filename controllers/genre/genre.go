package genre

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func GetListGenres(c *gin.Context) {
	var paginationDto requests.PaginationDto
	var genres []entities.Genre

	if err := c.ShouldBind(&paginationDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Scopes(entities.Paginate(genres, &paginationDto)).Find(&genres); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "No data genres.", nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", genres)
}

func GetGenre(c *gin.Context) {
	id := c.Param("id")
	var genre entities.Genre

	if tx := entities.DB().Where("id = ?", id).First(&genre); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Genre [id: %s] not found.", id), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", genre)
}

func CreateGenre(c *gin.Context) {
	var genre entities.Genre

	if err := c.ShouldBind(&genre); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	genreCreate := entities.Genre{
		Name:        genre.Name,
		Description: genre.Description,
	}

	if err := entities.DB().Create(&genreCreate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", genreCreate)
}

func UpdateGenre(c *gin.Context) {
	var genreRequest entities.Genre
	var genreUpdate entities.Genre

	if err := c.ShouldBind(&genreRequest); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Where("id = ?", genreRequest.ID).First(&genreUpdate); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Genre [id: %d] not found.", genreRequest.ID), nil)
		return
	}

	genreUpdate.Name = genreRequest.Name
	genreUpdate.Description = genreRequest.Description

	if err := entities.DB().Where("id = ?", genreRequest.ID).Save(&genreUpdate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", genreUpdate)
}

func DeleteGenre(c *gin.Context) {
	id := c.Param("id")
	var genre entities.Genre

	if tx := entities.DB().Where("id = ?", id).First(&genre); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Genre [id: %s] not found.", id), nil)
		return
	}

	if err := entities.DB().Where("id = ?", id).Delete(&genre).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", genre)
}
