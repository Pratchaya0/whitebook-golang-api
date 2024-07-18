package carts

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /cart/create [post]
func CreateCart(c *gin.Context) {
	var cart entities.Cart
	var book entities.Book
	var user entities.User

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", cart.CartBookId).First(&book); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	if tx := entities.DB().Where("id = ?", cart.UserId).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	cartCreate := entities.Cart{
		Book: book,
		User: user,
	}

	if err := entities.DB().Create(&cartCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   cartCreate,
	}

	c.JSON(http.StatusCreated, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /cart/{cartId} [post]
func GetCart(c *gin.Context) {
	cartId := c.Param("cartId")

	var cart entities.Cart

	if err := entities.DB().Preload("Book").Preload("User").Raw("SELECT * FROM carts WHERE id = ?", cartId).First(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   cart,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /carts [post]
func GetListCarts(c *gin.Context) {
	var carts []entities.Cart

	if err := entities.DB().Preload("Book").Preload("User").Raw("SELECT * FROM carts").Find(&carts).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   carts,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /cart/update [patch]
func UpdateCart(c *gin.Context) {
	var cart entities.Cart

	if err := c.ShouldBindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", cart.ID).First(&cart); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
		return
	}

	if err := entities.DB().Save(&cart).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   cart,
	}

	c.JSON(http.StatusOK, webResponse)
}

// @Summary Get a list of books in the the store
// @Tag BookPreviewImages
// @Security bearerToken
// @Produce  json
// @Success 200 {object} responses.Response{} "ok"
// @Router /cart/delete/{cartId} [delete]
func DeleteCart(c *gin.Context) {
	cartId := c.Param("cartId")

	if tx := entities.DB().Exec("DELETE FROM carts WHERE id = ?", cartId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   gin.H{"message": "cart id: " + cartId + " has been deleted"},
	}

	c.JSON(http.StatusOK, webResponse)
}
