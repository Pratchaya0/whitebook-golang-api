package advertisements

import (
	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

func CreateAdvertisement(c *gin.Context) {
	var advertisement entities.Advertisement

	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	advertisementCreate := entities.Advertisement{
		AdvertisementTitle:       advertisement.AdvertisementTitle,
		AdvertisementDescription: advertisement.AdvertisementDescription,
		AdvertisementHighlight:   advertisement.AdvertisementHighlight,
		AdvertisementImageUrl:    advertisement.AdvertisementImageUrl,
	}

	if err := entities.DB().Create(&advertisementCreate).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": advertisement})
}

func GetAdvertisement(c *gin.Context) {
	advertisementId := c.Param("advertisementsId")

	var advertisement entities.Advertisement

	if tx := entities.DB().Where("id = ?", advertisementId).First(&advertisement); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user role not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func GetListAdvertisements(c *gin.Context) {
	var advertisements []entities.Advertisement

	if err := entities.DB().Raw("SELECT * FROM advertisements").Scan(&advertisements).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisements})
}

func UpdateAdvertisement(c *gin.Context) {
	var advertisement entities.Advertisement

	if err := c.ShouldBindJSON(&advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	if _, err := govalidator.ValidateStruct(advertisement); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entities.DB().Where("id = ?", advertisement.ID).First(&advertisement); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "advertisement not found"})
		return
	}

	if err := entities.DB().Save(&advertisement).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisement})
}

func DeleteAdvertisement(c *gin.Context) {
	advertisementId := c.Param("advertisementId")

	if tx := entities.DB().Exec("DELETE FROM advertisements WHERE id = ?", advertisementId); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "advertisement not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": advertisementId})
}
