package main

import (
	"log"
	"simplerapi/controllers"
	"simplerapi/postgresDB"
	"simplerapi/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database := postgresDB.InitDB()
	if database == nil {
		log.Fatal("failed to initialize database")
	}

	///setup db
	postgresDB.DB = database

	//inject services
	productService := services.NewProductService(postgresDB.DB)

	//seed database
	postgresDB.SeedDB(postgresDB.DB)

	// Set up Gin router
	r := gin.Default()

	//setup controller
	controllers.SetupProductController(r, productService)

	r.Run(":8080")
}
