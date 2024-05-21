package webinfos

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

func GetWebInfo(c *gin.Context) {
	webInfoId := c.Param("webInfoId")

	var webInfo entities.WebInfo

	if tx := entities.DB().Where("id = ?", webInfoId).First(&webInfo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "web info not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": webInfo})
}

func UpdateWebInfo(c *gin.Context) {
	var webInfo entities.WebInfo

	if err := c.ShouldBindJSON(&webInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", webInfo.ID).First(&webInfo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "web info not found"})
		return
	}

	if err := entities.DB().Save(&webInfo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": webInfo})
}
