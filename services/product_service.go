package services

import (
	"simplerapi/models"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	CreateProduct(c *gin.Context) (models.Product, error)
	GetProductById(c *gin.Context) (models.Product, error)
	UpdateProduct(c *gin.Context) (models.Product, error)
	DeleteProduct(c *gin.Context) error
	GetProducts(pc *gin.Context) ([]models.Product, error)
}
