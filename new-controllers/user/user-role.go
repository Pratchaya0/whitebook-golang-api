package user

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

func GetListUserRoles(c *gin.Context) {
	var userRoles []entities.UserRole

	if err := entities.DB().Raw("SELECT * FROM user_roles").Find(&userRoles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userRoles,
	}

	c.JSON(http.StatusOK, webResponse)
}

func GetUserRole(c *gin.Context) {
	var userRoleId = c.Param("id")
	var userRole entities.UserRole

	if err := entities.DB().Where("id = ?", userRoleId).First(&userRole).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userRole,
	}

	c.JSON(http.StatusOK, webResponse)
}

func CreateUserRole(c *gin.Context) {
	var userRole entities.UserRole

	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := responses.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   userRole,
	}

	c.JSON(http.StatusOK, webResponse)
}
