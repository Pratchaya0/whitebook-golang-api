package orders

// import (
// 	"net/http"

// 	"github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	"github.com/gin-gonic/gin"
// )

// func CreateOrderBookDetail(c *gin.Context) {
// 	var orderBookDetail entities.OrderBookDetail
// 	var order entities.Order
// 	var book entities.Book

// 	if err := c.ShouldBindJSON(&orderBookDetail); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// แทรกการ validate ไว้ช่วงนี้ของ controller
// 	if _, err := govalidator.ValidateStruct(orderBookDetail); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", orderBookDetail.OrderId).First(&order); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", orderBookDetail.OrderBookDetailBookId).First(&book); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
// 		return
// 	}

// 	orderBookDetailCreate := entities.OrderBookDetail{
// 		Order: order,
// 		Book:  book,
// 	}

// 	if err := entities.DB().Create(&orderBookDetailCreate).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{"data": orderBookDetailCreate})
// }

// func GetOrderBookDetail(c *gin.Context) {
// 	orderBookDetailId := c.Param("orderBookDetailId")

// 	var orderBookDetail entities.OrderBookDetail

// 	if err := entities.DB().Preload("Order").Preload("Book").Raw("SELECT * FROM order_book_details WHERE id = ?", orderBookDetailId).First(&orderBookDetail).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": orderBookDetail})
// }

// func GetListOrderBookDetails(c *gin.Context) {
// 	var orderBookDetails []entities.OrderBookDetail

// 	if err := entities.DB().Preload("Order").Preload("Book").Raw("SELECT * FROM order_book_details").Find(&orderBookDetails).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": orderBookDetails})
// }

// func UpdateOrderBookDetail(c *gin.Context) {
// 	var orderBookDetail entities.OrderBookDetail

// 	if err := c.ShouldBindJSON(&orderBookDetail); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// แทรกการ validate ไว้ช่วงนี้ของ controller
// 	if _, err := govalidator.ValidateStruct(orderBookDetail); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", orderBookDetail.ID).First(&orderBookDetail); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "book not found"})
// 		return
// 	}

// 	if err := entities.DB().Save(&orderBookDetail).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": orderBookDetail})
// }

// func DeleteOrderBookDetail(c *gin.Context) {
// 	orderBookDetailId := c.Param("orderBookDetailId")

// 	if tx := entities.DB().Exec("DELETE FROM order_book_details WHERE id = ?", orderBookDetailId); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": "order book detail id: " + orderBookDetailId + " has been deleted"})
// }
