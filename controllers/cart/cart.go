package cart

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context, userId uint) {
	var user entities.User
	var cart entities.Cart

	if err := entities.DB().Where("id = ?", userId).First(&user).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusNotFound, err.Error(), nil)
		return
	}

	if tx := entities.DB().Where("user_id = ?", userId).First(&cart); tx.RowsAffected > 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "Cart already exist.", nil)
		return
	}

	cartCreate := entities.Cart{
		UserID: userId,
	}

	if err := entities.DB().Create(&cartCreate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", cartCreate)
}

func GetListCartsByUserId(c *gin.Context) {
	userId := c.Param("userId")
	var carts []entities.Cart

	if tx := entities.DB().Preload("Items").Where("user_id = ?", userId).Find(&carts); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "Cart is empty.", nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", carts)
}
