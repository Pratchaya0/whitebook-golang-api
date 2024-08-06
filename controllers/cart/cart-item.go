package cart

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func CreateCartItem(c *gin.Context) {
	var book entities.Book
	var cart entities.Cart
	var cartItem *entities.CartItem
	var status entities.CartItemStatus

	if err := c.ShouldBind(&cartItem); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check book is exist?
	if tx := entities.DB().Where("id = ?", cartItem.BookID).First(&book); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Book [id: %d] not found.", cartItem.BookID), nil)
		return
	}

	// check cart is exist?
	if tx := entities.DB().Where("id = ?", cartItem.CartID).First(&cart); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Cart [id: %d] not found.", cartItem.CartID), nil)
		return
	}

	if tx := entities.DB().Where("id = ?", cartItem.CartItemStatusID).First(&status); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Cate item status [id: %d] not found.", cartItem.CartItemStatusID), nil)
		return
	}

	// check cartItem already added?
	if tx := entities.DB().Where(&entities.CartItem{CartID: cartItem.CartID, BookID: cartItem.BookID}).First(&cartItem); tx.RowsAffected > 0 {
		helpers.RespondWithJSON(c, http.StatusOK, "This product already added.", cartItem)
		return
	}

	// create cart item.
	cartItemCreate := entities.CartItem{
		CartID:         cartItem.CartID,
		Book:           book,
		CartItemStatus: status,
	}

	if err := entities.DB().Create(&cartItemCreate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", cartItemCreate)
}

func UpdateCartItem(c *gin.Context) {
	// TODO Update status
}

// not use
func DeleteCartItem(c *gin.Context) {
	id := c.Param("id")
	var cartItem entities.CartItem

	if tx := entities.DB().Where("id = ?", id).First(&cartItem); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Cart Item [id: %s] not found.", id), nil)
		return
	}

	if err := entities.DB().Exec("DELETE FROM books WHERE id = ?", id).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", cartItem)
}
