package user

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

// The request body is described by the @Param annotation, which has the following syntax:

//  @Param [param_name] [param_type] [data_type] [required/mandatory] [description]
// The param_type can be one of the following values:

// query (indicates a query param)
// path (indicates a path param)
// header (indicates a header param)
// body
// formData

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

func UpdateUserRole(c *gin.Context) {

}
