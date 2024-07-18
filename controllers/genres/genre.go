package genres

// import (
// 	"net/http"

// 	"github.com/Pratchaya0/whitebook-golang-api/dtos/responses"
// 	"github.com/Pratchaya0/whitebook-golang-api/entities"
// 	"github.com/asaskevich/govalidator"
// 	"github.com/gin-gonic/gin"
// )

// // @Summary Get a list of books in the the store
// // @Tag Genre
// // @Security bearerToken
// // @Produce  json
// // @Success 200 {object} responses.Response{} "ok"
// // @Router /genre/create [post]
// func CreateGenre(c *gin.Context) {
// 	var genre entities.Genre

// 	if err := c.ShouldBindJSON(&genre); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// แทรกการ validate ไว้ช่วงนี้ของ controller
// 	if _, err := govalidator.ValidateStruct(genre); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	genreCreate := entities.Genre{
// 		GenreName: genre.GenreName,
// 	}

// 	if err := entities.DB().Create(&genreCreate).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "OK",
// 		Data:   genreCreate,
// 	}

// 	c.JSON(http.StatusCreated, webResponse)
// }

// // @Summary Get a list of books in the the store
// // @Tag BookPreviewImages
// // @Security bearerToken
// // @Produce  json
// // @Success 200 {object} responses.Response{} "ok"
// // @Router /genre [post]
// func GetGenre(c *gin.Context) {
// 	genreId := c.Param("genreId")

// 	var genre entities.Genre

// 	if tx := entities.DB().Where("id = ?", genreId).First(&genre); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
// 		return
// 	}

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "OK",
// 		Data:   genre,
// 	}

// 	c.JSON(http.StatusOK, webResponse)
// }

// // @Summary Get a list of books in the the store
// // @Tag BookPreviewImages
// // @Security bearerToken
// // @Produce  json
// // @Success 200 {object} responses.Response{} "ok"
// // @Router /genres [post]
// func GetListGenres(c *gin.Context) {
// 	var genres []entities.Genre

// 	if err := entities.DB().Raw("SELECT * FROM genres").Scan(&genres).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "OK",
// 		Data:   genres,
// 	}

// 	c.JSON(http.StatusOK, webResponse)
// }

// // @Summary Get a list of books in the the store
// // @Tag BookPreviewImages
// // @Security bearerToken
// // @Produce  json
// // @Success 200 {object} responses.Response{} "ok"
// // @Router /genre/update [post]
// func UpdateGenre(c *gin.Context) {
// 	var genre entities.Genre

// 	if err := c.ShouldBindJSON(&genre); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// แทรกการ validate ไว้ช่วงนี้ของ controller
// 	if _, err := govalidator.ValidateStruct(genre); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entities.DB().Where("id = ?", genre.ID).First(&genre); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
// 		return
// 	}

// 	if err := entities.DB().Save(&genre).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "OK",
// 		Data:   genre,
// 	}

// 	c.JSON(http.StatusOK, webResponse)
// }

// // @Summary Get a list of books in the the store
// // @Tag BookPreviewImages
// // @Security bearerToken
// // @Produce  json
// // @Success 200 {object} responses.Response{} "ok"
// // @Router /genre/delete/{genreId} [post]
// func DeleteGenre(c *gin.Context) {
// 	genreId := c.Param("genreId")

// 	if tx := entities.DB().Exec("DELETE FROM genres WHERE id = ?", genreId); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "genre not found"})
// 		return
// 	}

// 	webResponse := responses.Response{
// 		Code:   http.StatusOK,
// 		Status: "OK",
// 		Data:   genreId,
// 	}

// 	c.JSON(http.StatusOK, webResponse)
// }
