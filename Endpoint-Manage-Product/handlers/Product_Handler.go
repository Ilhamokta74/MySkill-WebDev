package handlers

import (
	"Endpoint-Manage-Product/models"
	"sync"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup

		// Menambahkan 1 Goroutine ke WaitGroup
		wg.Add(1)

		// Memulai Goroutine Untuk Melakukan Operai Yang Membutuhkan Waktu Lama
		go func() {
			defer wg.Done() // Menandai bahwa goroutine telah selesai
			db.Find(&products)
		}()

		// Menunggu Goroutine Selesai
		wg.Wait()

		c.JSON(200, products)
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var products models.Product

		if err := db.First(&products, id).Error; err != nil {
			c.JSON(404, gin.H{
				"message": "Product Not Found",
			})
			return
		}
		c.JSON(200, products)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		db.Create(&input)
		c.JSON(201, input)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var products models.Product

		if err := db.First(&products, id).Error; err != nil {
			c.JSON(404, gin.H{
				"message": "Product Not Found",
			})
			return
		}

		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(400, gin.H{
				"message": "Invalid Input",
			})
			return
		}

		db.Model(&products).Updates(input)
		c.JSON(200, products)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var products models.Product
		if err := db.First(&products, id).Error; err != nil {
			c.JSON(404, gin.H{
				"message": "Product Not Found",
			})
			return
		}

		db.Delete(&products)
		c.JSON(200, gin.H{
			"message": "Product Deleted",
		})
	}
}
