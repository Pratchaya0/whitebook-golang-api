package genres

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateGenreBook(c *gin.Context) {
	var genreBook entities.GenreBook
	var genre entities.Genre
	var book entities.Book

	if err := c.ShouldBindJSON(&genreBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(genreBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", genreBook.GenreId).First(&genre); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
		return
	}

	if tx := entities.DB().Where("id = ?", genreBook.GenreBookBookId).First(&book); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	genreBookCreate := entities.GenreBook{
		Genre: genre,
		Book:  book,
	}

	if err := entities.DB().Create(&genreBookCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": genreBookCreate})
}

func GetGenreBook(c *gin.Context) {
	genreBookId := c.Param("genreBookId")

	var genreBook entities.GenreBook

	if err := entities.DB().Preload("Genre").Preload("Book").Raw("SELECT * FROM genre_books WHERE id = ?", genreBookId).First(&genreBook).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genreBook})
}

func GetListGenreBooks(c *gin.Context) {
	var genreBooks []entities.GenreBook

	if err := entities.DB().Preload("Genre").Preload("Book").Raw("SELECT * FROM genre_books").Find(&genreBooks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genreBooks})
}

func UpdateGenreBook(c *gin.Context) {
	var genreBooks entities.GenreBook

	if err := c.ShouldBindJSON(&genreBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(genreBooks); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", genreBooks.ID).First(&genreBooks); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre book not found"})
		return
	}

	if err := entities.DB().Save(&genreBooks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": genreBooks})
}

func DeleteGenreBook(c *gin.Context) {
	genreBookId := c.Param("genreBookId")

	if tx := entities.DB().Exec("DELETE FROM genre_books WHERE id = ?", genreBookId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "genre book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "genre book id: " + genreBookId + " has been deleted"})
}
