package books

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @Summary Get a list of books in the the store
// @Tag Book
// @Description -
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/create [post]
func CreateBook(c *gin.Context) {
	var book entities.Book
	var category entities.Category

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", book.BookCategoryId).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}

	bookCreate := entities.Book{
		BookName:          book.BookName,
		BookDescription:   book.BookDescription,
		BookPrice:         book.BookPrice,
		BookWriter:        book.BookWriter,
		BookPublisher:     book.BookPublisher,
		BookIsOnSale:      book.BookIsOnSale,
		BookCoverImageUrl: book.BookCoverImageUrl,
		BookUrl:           book.BookUrl,
		Category:          category,
	}

	if err := entities.DB().Create(&bookCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookCreate,
	}

	c.JSON(http.StatusCreated, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag Book
// @Description -
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book [get]
func GetBook(c *gin.Context) {
	bookId := c.Param("bookId")

	var book entities.Book

	if err := entities.DB().Preload("Category").Raw("SELECT * FROM books WHERE id = ?", bookId).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   book,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag Book
// @Description get string by ID
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /books [get]
func GetListBooks(c *gin.Context) {
	var books []entities.Book

	if err := entities.DB().Preload("Category").Raw("SELECT * FROM books").Find(&books).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   books,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag Book
// @Description -
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/update [patch]
func UpdateBook(c *gin.Context) {
	var book entities.Book
	var category entities.Category

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", book.BookCategoryId).First(&category); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "category not found"})
		return
	}

	bookUpdate := entities.Book{
		BookName:          book.BookName,
		BookDescription:   book.BookDescription,
		BookPrice:         book.BookPrice,
		BookWriter:        book.BookWriter,
		BookPublisher:     book.BookPublisher,
		BookIsOnSale:      book.BookIsOnSale,
		BookCoverImageUrl: book.BookCoverImageUrl,
		BookUrl:           book.BookUrl,
		Category:          category,
	}

	if err := entities.DB().Save(&bookUpdate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookUpdate,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag Book
// @Description -
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book/delete [delete]
func DeleteBook(c *gin.Context) {
	bookId := c.Param("bookId")

	if tx := entities.DB().Exec("DELETE FROM books WHERE id = ?", bookId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "book id: " + bookId + " has been deleted",
	}

	c.JSON(http.StatusOK, webResponse)
}
