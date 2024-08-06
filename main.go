package main

import (
	// "github.com/Pratchaya0/whitebook-golang-api/docs"

	"net/http"

	"github.com/Pratchaya0/whitebook-golang-api/controllers"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/book"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/cart"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/category"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/genre"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/order"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/payment"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/review"
	"github.com/Pratchaya0/whitebook-golang-api/controllers/user"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/Pratchaya0/whitebook-golang-api/middlewares"
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

	router := r.Group("/")
	{
		protected := router.Use(middlewares.Authorizes())
		{
			protected.GET("/books", book.GetListBooks)
			protected.GET("/book/:id", book.GetBook)
			protected.POST("/book/create", book.CreateBook)
			protected.PUT("/book/update", book.UpdateBook)
			protected.DELETE("/book/:id/delete", book.DeleteBook)

			protected.GET("/categories", category.GetListCategories)
			protected.GET("/category/:id", category.GetCategory)
			protected.POST("/category/create", category.CreateCategory)
			protected.PUT("/category/update", category.UpdateCategory)
			protected.DELETE("/category/:id/delete", category.DeleteCategory)

			protected.GET("/genres", genre.GetListGenres)
			protected.GET("/genre/:id", genre.GetGenre)
			protected.POST("/genre/create", genre.CreateGenre)
			protected.PUT("/genre/update", genre.UpdateGenre)
			protected.DELETE("/genre/:id/delete", genre.DeleteGenre)

			protected.POST("/order/create", order.CreateOrder)
			protected.PUT("/order/update", order.UpdateOrder)

			protected.GET("/paymentMethods", payment.GetListPaymentMethods)
			protected.GET("/paymentMethod/:id", payment.GetPaymentMethod)
			protected.POST("/paymentMethod/create", payment.CreatePaymentMethod)
			protected.PUT("/paymentMethod/update", payment.UpdatePaymentMethod)
			protected.DELETE("/paymentMethod/:id/delete", payment.DeletePaymentMethod)

			protected.GET("/reviews/:id/book", review.GetListReviewsByBookID)
			protected.GET("/reviews/:id/user", review.GetListReviewsByUserID)
			protected.GET("/review/:id", review.GetReview)
			protected.POST("/review/create", review.CreateReview)
			protected.PUT("/review/update", review.UpdateReview)
			protected.DELETE("/review/:id/delete", review.DeleteReview)

			protected.GET("/cart/list/:id", cart.GetListCartsByUserId)

			protected.POST("/cartItem/add", cart.CreateCartItem)
			protected.DELETE("/cartItem/:id/delete", cart.DeleteCartItem)

			protected.GET("/users", user.GetListUsers)
			protected.GET("/user/:id", user.GetUser)
			protected.PUT("/user/update", user.UpdateUser)
			protected.DELETE("/user/delete", user.DeleteUser)

			protected.GET("/userRoles", user.GetListUserRoles)
			protected.GET("/userRole", user.GetUserRole)
			protected.POST("/userRole/create", user.CreateUserRole)
		}
	}

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/healthCheck", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})

	// Sign Up User Route
	r.POST("/signup", controllers.SignUp)
	// login User Route
	r.POST("/login", controllers.Login)

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
