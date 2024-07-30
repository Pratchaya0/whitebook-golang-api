package user

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func GetListUsers(c *gin.Context) {

	// !!! only admin or higher
	var users []entities.User

	if err := entities.DB().Preload("UserRoles").Find(&users).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", users)
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")

	var user entities.User

	if err := entities.DB().Preload("UserRoles").Where("id = ?", userId).First(&user).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", user)
}

func UpdateUser(c *gin.Context) {
	var user entities.User

	if err := c.ShouldBind(&user); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("User [id: %d] not found.", user.ID), nil)
		return
	}

	if err := entities.DB().Save(&user).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", user)
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	var user entities.User

	if tx := entities.DB().Where("id = ?", userId).Delete(&user); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("User [id: %s] not found.", userId), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", user)
}
