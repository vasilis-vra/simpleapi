package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	// Initialize the database
	db = initDB()
	if db == nil {
		log.Fatal("failed to initialize database")
	}

	seedDB(db)

	// Set up Gin router
	r := gin.Default()

	// Define routes
	r.POST("/products", createProduct)
	r.GET("/products/:id", getProduct)
	r.GET("/products", getProducts)
	r.PUT("/products/:id", updateProduct)
	r.DELETE("/products/:id", deleteProduct)

	// Start the server
	r.Run(":8080")
}
