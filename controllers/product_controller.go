package controllers

import (
	"errors"
	"log"
	"simplerapi/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	service services.ProductService
}

func NewProductController(service services.ProductService) *ProductController {
	return &ProductController{service: service}
}

func SetupProductController(r *gin.Engine, service services.ProductService) {
	productController := NewProductController(service)

	r.POST("/products", productController.CreateProduct)
	r.GET("/products/:id", productController.GetProductById)
	r.GET("/products", productController.GetProducts)
	r.PUT("/products/:id", productController.UpdateProduct)
	r.DELETE("/products/:id", productController.DeleteProduct)
}

func (p *ProductController) CreateProduct(c *gin.Context) {
	createdProduct, err := p.service.CreateProduct(c)
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(409, gin.H{"error": "Product already exists"})
			return
		}
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println("Product created successfully")
	c.JSON(201, createdProduct)
}

func (p *ProductController) GetProductById(c *gin.Context) {
	product, err := p.service.GetProductById(c)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, product)
}

func (p *ProductController) GetProducts(c *gin.Context) {
	products, err := p.service.GetProducts(c)
	if err != nil {
		// Log the error for debugging
		log.Printf("Error getting products: %v", err)

		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, products)
}

func (p *ProductController) UpdateProduct(c *gin.Context) {

	updatedProduct, err := p.service.UpdateProduct(c)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.Println("Product updated successfully")

	c.JSON(200, updatedProduct)
}

func (p *ProductController) DeleteProduct(c *gin.Context) {
	err := p.service.DeleteProduct(c)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	log.Println("Product deleted successfully")

	c.JSON(204, gin.H{})
}
