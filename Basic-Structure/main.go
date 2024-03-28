package main

import "github.com/gin-gonic/gin"

var loginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func main() {
	// Inisialisasi Gin Router
	router := gin.Default()

	// Middleware : Logger
	router.Use(gin.Logger())

	// Middleware : Recovery
	router.Use(gin.Recovery())

	// Router Definition
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"message": "Hello " + name + "!",
		})
	})

	router.POST("/login", func(c *gin.Context) {
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid Error Body",
			})
			return
		}

		// Validasi To login
		// Cek Email Dan Password Apakah Cocok
		if loginData.Email == "example@gmail.com" && loginData.Password == "password123" {
			c.JSON(200, gin.H{
				"message": "Login Successful",
			})
		} else {
			c.JSON(401, gin.H{
				"error": "Invalid Credentials",
			})
		}
	})

	// Mengambil Data Dengan Parameter Query
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Query("name")

		if name == "" {
			c.JSON(400, gin.H{
				"error": "Name Parameter Is Missing",
			})
			return
		}

		c.JSON(200, gin.H{
			"message": "Hello" + name + "!",
		})
	})

	// Jalankan Server
	router.Run(":8080")
}
