package books

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @Summary Get a list of books in the the store
// @Tag BookDetail
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-user-detail/create [post]
func CreateBookUserDetail(c *gin.Context) {
	var bookUserDetail entities.BookUserDetail
	var book entities.Book
	var user entities.User

	if err := c.ShouldBindJSON(&bookUserDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(bookUserDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", bookUserDetail.BookUserDetailBookId).First(&book); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	if tx := entities.DB().Where("id = ?", bookUserDetail.UserId).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	bookUserDetailCreate := entities.BookUserDetail{
		BookUserDetailIsAvailable: bookUserDetail.BookUserDetailIsAvailable,
		Book:                      book,
		User:                      user,
	}

	if err := entities.DB().Create(&bookUserDetailCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookUserDetailCreate,
	}

	c.JSON(http.StatusCreated, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookDetail
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-user-detail/{bookUserDetailId} [get]
func GetBookUserDetail(c *gin.Context) {
	bookUserDetailId := c.Param("bookUserDetailId")

	var bookUserDetail entities.BookUserDetail

	if err := entities.DB().Preload("Book").Preload("User").Raw("SELECT * FROM book_user_details WHERE id = ?", bookUserDetailId).First(&bookUserDetail).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookUserDetail,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookDetail
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-user-details [get]
func GetListBookUserDetails(c *gin.Context) {
	var bookUserDetails []entities.BookUserDetail

	if err := entities.DB().Preload("Book").Preload("User").Raw("SELECT * FROM book_user_details").Find(&bookUserDetails).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookUserDetails,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookDetail
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-user-detail/update [post]
func UpdateBookUserDetail(c *gin.Context) {
	var bookUserDetail entities.BookUserDetail
	var book entities.Book
	var user entities.User

	if err := c.ShouldBindJSON(&bookUserDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(bookUserDetail); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", bookUserDetail.BookUserDetailBookId).First(&book); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	if tx := entities.DB().Where("id = ?", bookUserDetail.UserId).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	bookUserDetailUpdate := entities.BookUserDetail{
		BookUserDetailIsAvailable: bookUserDetail.BookUserDetailIsAvailable,
		Book:                      book,
		User:                      user,
	}

	if err := entities.DB().Save(&bookUserDetailUpdate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   bookUserDetailUpdate,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookDetail
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /book-user-detail/delete/{bookUserDetailId} [post]
func DeleteBookUserDetail(c *gin.Context) {
	bookUserDetailId := c.Param("bookUserDetailId")

	if tx := entities.DB().Exec("DELETE FROM book_user_details WHERE id = ?", bookUserDetailId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book user detail not found"})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   gin.H{"message": "book id: " + bookUserDetailId + " has been deleted"},
	}

	c.JSON(http.StatusOK, webResponse)
}
