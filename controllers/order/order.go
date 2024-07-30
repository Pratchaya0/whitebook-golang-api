package order

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func calculateAmount(cartItems []entities.CartItem) float64 {
	if len(cartItems) == 0 {
		log.Println("Cart is empty or nil")
		return 0
	}

	var amount float64 = 0
	for _, item := range cartItems {
		amount += item.Book.Price
	}
	return amount
}

func CreateOrder(c *gin.Context) {
	id := c.Param("id")

	var user entities.User
	var cart entities.Cart
	var cartItem []entities.CartItem

	userID, _ := strconv.ParseUint(id, 10, 32)

	if tx := entities.DB().Where("id = ?", userID).First(&user); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "User [id: %s] not found.", userID)
		return
	}

	if tx := entities.DB().Where(&entities.Cart{UserID: uint(userID)}).First(&cart); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "Cart not found.", nil)
		return
	}

	if tx := entities.DB().Where(&entities.CartItem{CartID: cart.ID}).Find(&cartItem); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "Cart is empty.", nil)
		return
	}

	var refCode = helpers.Hash(fmt.Sprintf("%s%d%s", id, cart.ID, time.Now().Format(time.RFC3339)))
	var amount = calculateAmount(cartItem)

	orderCreate := entities.Order{
		RefCode: refCode,
		Amount:  amount,

		User:  user,
		Items: cartItem,
	}

	if err := entities.DB().Create(&orderCreate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", orderCreate)
}

func UpdateOrder(c *gin.Context) {
	var orderRequest requests.OrderUpdate
	var order entities.Order
	var user entities.User
	var paymentMethod entities.PaymentMethod
	var cart entities.Cart
	var cartItem []entities.CartItem

	if err := c.ShouldBind(&orderRequest); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// check order is exist?
	if tx := entities.DB().Where("id = ?", orderRequest.ID).First(&order); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusBadRequest, fmt.Sprintf("Order [id: %d] not found.", order.ID), nil)
		return
	}

	// check user is exist?
	if tx := entities.DB().Where("id = ?", orderRequest.UserID).First(&user); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("User [id: %d] not found.", user.ID), nil)
		return
	}

	// check payment method is exist?
	if tx := entities.DB().Where("id = ?", orderRequest.PaymentMethodID).First(&paymentMethod); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Payment method [id: %d] not found.", order.PaymentMethodID), nil)
		return
	}

	// check item in cart
	if tx := entities.DB().Where(&entities.Cart{UserID: orderRequest.UserID}).First(&cart); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "Cart not found.", nil)
		return
	}

	if tx := entities.DB().Where(&entities.CartItem{CartID: cart.ID}).Find(&cartItem); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "Cart is empty.", nil)
		return
	}

	orderUpdate := entities.Order{
		RefCode:   order.RefCode,
		Amount:    order.Amount,
		SlipImage: order.SlipImage,

		PaymentMethodID: paymentMethod.ID,
		User:            user,
		Items:           cartItem,
	}

	if err := entities.DB().Create(&orderUpdate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", orderUpdate)
}
