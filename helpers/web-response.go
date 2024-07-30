package helpers

import (
	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
	"github.com/gin-gonic/gin"
)

func RespondWithJSON(c *gin.Context, code int, status string, data interface{}) {
	webResponse := responses.Response{
		Code:   code,
		Status: status,
		Data:   data,
	}
	c.JSON(code, webResponse)
}
