package users

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

func CreateUserRole(c *gin.Context) {
	var userRole entities.UserRole

	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userRoleCreate := entities.UserRole{
		UserRoleName: userRole.UserRoleName,
	}

	if err := entities.DB().Create(&userRoleCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": userRoleCreate})
}

func GetUserRole(c *gin.Context) {
	userRoleId := c.Param("userRoleId")

	var userRole entities.UserRole

	if tx := entities.DB().Where("id = ?", userRoleId).First(&userRole); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userRole})
}

func GetListUserRoles(c *gin.Context) {
	var userRoles []entities.UserRole

	if err := entities.DB().Raw("SELECT * FROM user_roles").Scan(&userRoles).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userRoles})
}

func UpdateUserRole(c *gin.Context) {
	var userRole entities.UserRole

	if err := c.ShouldBindJSON(&userRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", userRole.ID).First(&userRole); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user role not found"})
		return
	}

	if err := entities.DB().Save(&userRole).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userRole})
}

func DeleteUserRole(c *gin.Context) {
	userRoleId := c.Param("userRoleId")

	if tx := entities.DB().Exec("DELETE FROM user_roles WHERE id = ?", userRoleId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": userRoleId})
}
