package orders

import (
	"net/http"
	"time"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var order entities.Order
	var paymentInfo entities.PaymentInfo
	var user entities.User

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", order.PaymentInfoId).First(&paymentInfo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment information not found"})
		return
	}

	if tx := entities.DB().Where("id = ?", order.UserId).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	orderCreate := entities.Order{
		OrderAmount:          order.OrderAmount,
		OrderPaymentImageUrl: order.OrderPaymentImageUrl,
		OrderIsPaid:          order.OrderIsPaid,
		OrderCreateDate:      time.Now(),
		PaymentInfo:          paymentInfo,
		User:                 user,
	}

	if err := entities.DB().Create(&orderCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": orderCreate})
}

func GetOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	var order entities.Order

	if err := entities.DB().Preload("PaymentInfo").Preload("User").Raw("SELECT * FROM orders WHERE id = ?", orderId).First(&order).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})
}

func GetListOrders(c *gin.Context) {
	var orders []entities.Order

	if err := entities.DB().Preload("PaymentInfo").Preload("User").Raw("SELECT * FROM orders").Find(&orders).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})
}

func UpdateOrders(c *gin.Context) {
	var order entities.Order
	var paymentInfo entities.PaymentInfo
	var user entities.User

	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", order.ID).First(&order); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	if tx := entities.DB().Where("id = ?", order.PaymentInfoId).First(&paymentInfo); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payment information not found"})
		return
	}

	if tx := entities.DB().Where("id = ?", order.UserId).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	orderUpdate := entities.Order{
		OrderAmount:          order.OrderAmount,
		OrderPaymentImageUrl: order.OrderPaymentImageUrl,
		OrderIsPaid:          order.OrderIsPaid,
		OrderUpdateDate:      time.Now(),
		PaymentInfo:          paymentInfo,
		User:                 user,
	}

	if err := entities.DB().Save(&orderUpdate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": orderUpdate})
}

func DeleteOrder(c *gin.Context) {
	orderId := c.Param("orderId")

	if tx := entities.DB().Exec("DELETE FROM orders WHERE id = ?", orderId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "order id: " + orderId + " has been deleted"})
}
