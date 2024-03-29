package main

import (
	"Authentication/handlers"
	"Authentication/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := initDB()
	if err != nil {
		panic(err)
	}

	router := gin.Default()

	product := router.Group("/product")
	{
		product.GET("/", middlewares.AuthMiddleware(), handlers.ListProducts(db))
		product.GET("/:id", handlers.GetProduct(db))
		product.POST("/", handlers.CreateProduct(db))
		product.PUT("/:id", handlers.UpdateProduct(db))
		product.DELETE("/:id", handlers.DeleteProduct(db))
	}

	user := router.Group("/user")
	{
		user.POST("/register", handlers.Register(db))
	}

	router.Run(":8080")
}
