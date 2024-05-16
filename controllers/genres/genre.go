package genres

import (
	"net/http"
	"time"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

func CreateGenre(c *gin.Context) {
	var genre entities.Genre

	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	genreCreate := entities.Genre{
		GenreName:       genre.GenreName,
		GenreCreateDate: time.Now(),
	}

	if err := entities.DB().Create(&genreCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": genreCreate})
}

func GetGenre(c *gin.Context) {
	genreId := c.Param("genreId")

	var genre entities.Genre

	if tx := entities.DB().Where("id = ?", genreId).First(&genre); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func GetListGenres(c *gin.Context) {
	var genres []entities.Genre

	if err := entities.DB().Raw("SELECT * FROM genres").Scan(&genres).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genres})
}

func UpdateGenre(c *gin.Context) {
	var genre entities.Genre

	if err := c.ShouldBindJSON(&genre); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", genre.ID).First(&genre); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
		return
	}

	if err := entities.DB().Save(&genre).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genre})
}

func DeleteGenre(c *gin.Context) {
	genreId := c.Param("genreId")

	if tx := entities.DB().Exec("DELETE FROM genres WHERE id = ?", genreId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genreId})
}
