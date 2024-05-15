package main

import (
	"github.com/Pratchaya0/whitebook-golang-api/controllers"
	"github.com/Pratchaya0/whitebook-golang-api/entities"
	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entities.SetupDatabase()

	r := gin.Default()

	r.Use(CORSMiddleware())

	// router := r.Group("/")
	// {
	// 	// router.Use(middlewares.Authorizes()){
	// 	// 	//
	// 	// 	router.
	// 	// }
	// }

	// Sign Up User Route
	r.POST("/signup", controllers.SignUp)
	// login User Route
	r.POST("/login", controllers.Login)
	// Run the server go run main.go
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
