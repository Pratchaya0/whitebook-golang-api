package payment

import (
	"fmt"
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/dtos/requests"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/helpers"
	"github.com/gin-gonic/gin"
)

func GetListPaymentMethods(c *gin.Context) {
	var paginationDto requests.PaginationDto
	var paymentMethods []entities.PaymentMethod

	if err := c.ShouldBind(&paginationDto); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Scopes(entities.Paginate(paymentMethods, &paginationDto)).Find(&paymentMethods); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, "No data payment methods.", nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", paymentMethods)
}

func GetPaymentMethods(c *gin.Context) {
	id := c.Param("id")
	var paymentMethods entities.PaymentMethod

	if tx := entities.DB().Where("id = ?", id).First(&paymentMethods); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Payment method [id: %s] not found.", id), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", paymentMethods)
}

func CreatePaymentMethod(c *gin.Context) {
	var paymentMethod entities.PaymentMethod

	if err := c.ShouldBind(&paymentMethod); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	paymentMethodCreate := entities.PaymentMethod{
		Code:         paymentMethod.Code,
		ProviderName: paymentMethod.ProviderName,
		AccountName:  paymentMethod.AccountName,
	}

	if err := entities.DB().Create(&paymentMethodCreate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", paymentMethodCreate)
}

func UpdatePaymentMethod(c *gin.Context) {
	var paymentMethodRequest entities.PaymentMethod
	var paymentMethodUpdate entities.PaymentMethod

	if err := c.ShouldBind(&paymentMethodRequest); err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if tx := entities.DB().Where("id = ?", paymentMethodRequest.ID).First(&paymentMethodUpdate); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Payment method [id: %d] not found.", paymentMethodRequest.ID), nil)
		return
	}

	paymentMethodUpdate.Code = paymentMethodRequest.Code
	paymentMethodUpdate.ProviderName = paymentMethodRequest.ProviderName
	paymentMethodUpdate.AccountName = paymentMethodRequest.AccountName

	if err := entities.DB().Where("id = ?", paymentMethodRequest.ID).Save(&paymentMethodUpdate).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", paymentMethodUpdate)
}

func DeletePaymentMethod(c *gin.Context) {
	id := c.Param("id")
	var paymentMethod entities.PaymentMethod

	if tx := entities.DB().Where("id = ?", id).First(&paymentMethod); tx.RowsAffected == 0 {
		helpers.RespondWithJSON(c, http.StatusNotFound, fmt.Sprintf("Genre [id: %s] not found.", id), nil)
		return
	}

	if err := entities.DB().Where("id = ?", id).Delete(&paymentMethod).Error; err != nil {
		helpers.RespondWithJSON(c, http.StatusBadRequest, err.Error(), nil)
		return
	}

	helpers.RespondWithJSON(c, http.StatusOK, "OK", paymentMethod)
}
