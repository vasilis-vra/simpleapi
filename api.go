package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func createProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//if product already exists return resource conflict error code
	var existingProduct Product
	if err := db.Where("name = ?", product.Name).First(&existingProduct).Error; err == nil {
		c.JSON(409, gin.H{"error": "Product already exists"})
		return
	}

	db.Create(&product)
	c.JSON(201, product)
}

func getProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(200, product)
}

func getProducts(c *gin.Context) {
	var products []Product
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("page_size", "10")

	var offset int
	if p, err := strconv.Atoi(page); err == nil {
		offset = (p - 1) * 10 // Default to 10 if invalid
	}

	if ps, err := strconv.Atoi(pageSize); err == nil {
		db.Limit(ps).Offset(offset).Find(&products)
	} else {
		db.Limit(10).Offset(offset).Find(&products) // Default to 10
	}

	c.JSON(200, products)
}

func updateProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := db.First(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return

	}

	if product.ID == 0 {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	db.Save(&product)
	c.JSON(200, product)
}

func deleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product Product

	if err := db.Delete(&product, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(204, gin.H{})
}
