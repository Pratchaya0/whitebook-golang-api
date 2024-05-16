package paymentInfoCreates

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

func CreatePaymentInfo(c *gin.Context) {
	var paymentInfo entities.PaymentInfo

	if err := c.ShouldBindJSON(&paymentInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	paymentInfoCreate := entities.PaymentInfo{
		PaymentInfoName:     paymentInfo.PaymentInfoName,
		PaymentInfoCode:     paymentInfo.PaymentInfoCode,
		PaymentInfoImageUrl: paymentInfo.PaymentInfoImageUrl,
	}

	if err := entities.DB().Create(&paymentInfoCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": paymentInfoCreate})
}

func GetPaymentInfo(c *gin.Context) {
	paymentInfoId := c.Param("paymentInfoId")

	var paymentInfo entities.PaymentInfo

	if tx := entities.DB().Where("id = ?", paymentInfoId).First(&paymentInfo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment info role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentInfo})
}

func GetListPaymentInfos(c *gin.Context) {
	var paymentInfos []entities.PaymentInfo

	if err := entities.DB().Raw("SELECT * FROM payment_infos").Scan(&paymentInfos).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentInfos})
}

func UpdatePaymentInfo(c *gin.Context) {
	var paymentInfo entities.PaymentInfo

	if err := c.ShouldBindJSON(&paymentInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", paymentInfo.ID).First(&paymentInfo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment info role not found"})
		return
	}

	if err := entities.DB().Save(&paymentInfo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentInfo})
}

func DeletePaymentInfo(c *gin.Context) {
	paymentInfoId := c.Param("paymentInfoId")

	if tx := entities.DB().Exec("DELETE FROM payment_infos WHERE id = ?", paymentInfoId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": paymentInfoId})
}
