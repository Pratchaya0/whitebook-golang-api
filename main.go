package main

import (
	// "github.com/Pratchaya0/whitebook-golang-api/docs"

	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/controllers/book"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/category"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/genre"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

// @BasePath /

const PORT = "8080"

// @title Documenting Whitebook API
// @version 1
// @Description -

// @securityDefinitions.apikey bearerToken
// @in header
// @name Authorization

func main() {
	entities.SetupDatabaseII()

	r := gin.Default()

	// docs.SwaggerInfo.BasePath = "/"

	r.Use(CORSMiddleware())

	// router := r.Group("/")
	// {
	// 	protected := router.Use(middlewares.Authorizes())
	// 	{
	// 		// Advertise
	// 		protected.GET("advertisements", advertisements.GetListAdvertisements)
	// 		protected.GET("advertisement/:id", advertisements.GetAdvertisement)
	// 		protected.POST("advertisement/create", advertisements.CreateAdvertisement)
	// 		protected.PATCH("advertisement/update", advertisements.UpdateAdvertisement)
	// 		protected.DELETE("advertisement/delete/:id", advertisements.DeleteAdvertisement)

	// 		// Book
	// 		protected.GET("books", books.GetListBooks)
	// 		protected.GET("book/:id", books.GetBook)
	// 		protected.POST("book/create", books.CreateBook)
	// 		protected.PATCH("book/update", books.UpdateBook)
	// 		protected.DELETE("book/delete/:id", books.DeleteBook)

	// 		// Book Preview Image
	// 		protected.GET("book-preview-images", books.GetListBookPreviewImages)
	// 		protected.GET("book-preview-image/:id", books.GetBookPreviewImage)
	// 		protected.POST("book-preview-image/create", books.CreateBookPreviewImage)
	// 		protected.PATCH("book-preview-image/update", books.UpdateBookPreviewImage)
	// 		protected.DELETE("book-preview-image/delete/:id", books.DeleteBookPreviewImage)

	// 		// Book User Detail
	// 		protected.GET("book-user-details", books.GetListBookUserDetails)
	// 		protected.GET("book-user-detail", books.GetBookUserDetail)
	// 		protected.POST("book-user-detail/create", books.CreateBookUserDetail)
	// 		protected.PATCH("book-user-detail/update", books.UpdateBookUserDetail)
	// 		protected.DELETE("book-user-detail/delete/:id", books.DeleteBookUserDetail)

	// 		// Cart
	// 		protected.GET("carts", carts.GetListCarts)
	// 		protected.GET("cart/:id", carts.GetCart)
	// 		protected.POST("cart/create", carts.CreateCart)
	// 		protected.PATCH("cart/update", carts.UpdateCart)
	// 		protected.DELETE("cart/delete/:id", carts.DeleteCart)

	// 		// Category
	// 		protected.GET("categories", categories.GetListCategories)
	// 		protected.GET("category/:id", categories.GetCategory)
	// 		protected.POST("category/create", categories.CreateCategory)
	// 		protected.PATCH("category/update", categories.UpdateCategory)
	// 		protected.DELETE("category/delete/:id", categories.DeleteCategory)

	// 		// Genre
	// 		protected.GET("genres", genres.GetListGenres)
	// 		protected.GET("genre/:id", genres.GetGenre)
	// 		protected.POST("genre/create", genres.CreateGenre)
	// 		protected.PATCH("genre/update", genres.UpdateGenre)
	// 		protected.DELETE("genre/delete/:id", genres.DeleteGenre)

	// 		// Genre Book
	// 		protected.GET("genre-books", genres.GetListGenreBooks)
	// 		protected.GET("genre-book/:id", genres.GetGenreBook)
	// 		protected.POST("genre-book/create", genres.CreateGenreBook)
	// 		protected.PATCH("genre-book/update", genres.UpdateGenreBook)
	// 		protected.DELETE("genre-book/delete/:id", genres.DeleteGenreBook)

	// 		// Order
	// 		protected.GET("orders", orders.GetListOrders)
	// 		protected.GET("order/:id", orders.GetOrder)
	// 		protected.POST("order/create", orders.CreateOrder)
	// 		protected.PATCH("order/update", orders.UpdateOrders)
	// 		protected.DELETE("order/delete/:id", orders.DeleteOrder)

	// 		// Order Detail
	// 		protected.GET("order-details", orders.GetListOrderBookDetails)
	// 		protected.GET("order-detail/:id", orders.GetOrderBookDetail)
	// 		protected.POST("order-detail/create", orders.CreateOrderBookDetail)
	// 		protected.PATCH("order-detail/update", orders.UpdateOrderBookDetail)
	// 		protected.DELETE("order-detail/delete/:id", orders.DeleteOrderBookDetail)

	// 		// Payment information
	// 		protected.GET("payment-infos", paymentInfos.GetListPaymentInfos)
	// 		protected.GET("payment-info/:id", paymentInfos.GetPaymentInfo)
	// 		protected.POST("payment-info/create", paymentInfos.CreatePaymentInfo)
	// 		protected.PATCH("payment-info/update", paymentInfos.UpdatePaymentInfo)
	// 		protected.DELETE("payment-info/delete/:id", paymentInfos.DeletePaymentInfo)

	// 		// Review
	// 		protected.GET("reviews", reviews.GetListReviews)
	// 		protected.GET("review/:id", reviews.GetReview)
	// 		protected.POST("review/create", reviews.CreateReview)
	// 		protected.PATCH("review/update", reviews.UpdateReview)
	// 		protected.DELETE("review/delete/:id", reviews.DeleteReview)

	// 		// User
	// 		protected.GET("users", users.GetListUsers)
	// 		protected.GET("user/:id", users.GetUser)
	// 		protected.PATCH("user/update", users.UpdateUser)
	// 		protected.DELETE("user/delete", users.DeleteUser)

	// 		// User Role
	// 		protected.GET("user-roles", users.GetListUserRoles)
	// 		protected.GET("user-role/:id", users.GetListUsers)
	// 		protected.POST("user-role/create", users.CreateUserRole)
	// 		protected.PATCH("user-role/update", users.UpdateUserRole)
	// 		protected.DELETE("user-role/delete", users.DeleteUserRole)

	// 		// Web Information
	// 		protected.GET("web-info", webInfos.GetWebInfo)
	// 		protected.PATCH("web-info/update", webInfos.UpdateWebInfo)
	// 	}
	// }

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// // Sign Up User Route
	// r.POST("/signup", controllers.SignUp)
	// // login User Route
	// r.POST("/login", controllers.Login)
	// Run the server go run main.go
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	r.GET("/books", book.GetListBooks)
	r.GET("/book/:id", book.GetBook)
	r.POST("/book/create", book.CreateBook)
	r.PATCH("/book/update", book.UpdateBook)
	r.DELETE("/book/:id/delete", book.DeleteBook)

	r.GET("/categories", category.GetListCategories)
	r.GET("/category/:id", category.GetCategory)
	r.POST("/category/create", category.CreateCategory)
	r.PATCH("/category/update", category.UpdateCategory)
	r.DELETE("/category/:id/delete", category.DeleteCategory)

	r.GET("/genres", genre.GetListGenres)
	r.GET("/genre/:id", genre.GetGenre)
	r.POST("/genre/create", genre.CreateGenre)
	r.PATCH("/genre/update", genre.UpdateGenre)
	r.DELETE("/genre/:id/delete", genre.DeleteGenre)

	r.Run() // "localhost: " + PORT
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
